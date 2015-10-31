package main

import (
	"net/http"
	"net/url"

	"crypto/sha1"
	"fmt"
	"github.com/ckeyer/fx/wechat"
	"io"
	"sort"
)

func Filter(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Error(err)
	}
	if Auth(w, req) {
		log.Notice("auth failed")
		return
	}
	switch req.Method {
	case "Get":
		Get(w, req)
	case "Post":
		Post(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Auth 服务器验证
func Auth(w http.ResponseWriter, req *http.Request) bool {
	query := req.URL.RawQuery
	u, err := url.ParseQuery(query)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("url values: %#v\n", u)

	signature := u.Get("signature")
	timestamp := u.Get("timestamp")
	nonce := u.Get("nonce")
	echostr := u.Get("echostr")

	tmps := []string{config.WeChat.Token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]

	tmp := func(data string) string {
		t := sha1.New()
		io.WriteString(t, data)
		return fmt.Sprintf("%x", t.Sum(nil))
	}(tmpStr)
	if tmp == signature {
		w.Write([]byte(echostr))
		return true
	}
	return false
}

func Get(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
}

func Post(w http.ResponseWriter, req *http.Request) {
	wechat.Receive(w, req)
}
