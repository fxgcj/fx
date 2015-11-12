package main

import (
	"net/http"
	"net/url"

	"crypto/sha1"
	"fmt"
	"github.com/fxgcj/fx/wechat"
	"io"
	"sort"
	"strings"
)

func Filter(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Error(err)
	}
	if !Auth(w, req) {
		log.Notice("auth failed")
	}
	log.Debug("method: ", req.Method)
	log.Debug("url: ", req.URL.String())

	switch strings.ToUpper(req.Method) {
	case "GET":
		Get(w, req)
	case "POST":
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
		return false
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
	log.Debug("auth receive: ", tmp)
	log.Debug("auth should be: ", signature)
	return false
}

func Get(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("风险观察君出海去了呢~.~/e@n"))
}

func Post(w http.ResponseWriter, req *http.Request) {
	wechat.Receive(w, req)
}
