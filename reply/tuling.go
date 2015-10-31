package reply

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	tulingApiURL = "http://www.tuling123.com/openapi/api"
	apiKey       = "328549a76aa96a44c66341791c67c40b"
)

const (
	codeTextReply   = 100000
	codeLinkReply   = 200000
	codeNewsReply   = 302000
	codeTrainReply  = 305000
	codeFlightReply = 306000
	codeMenuReply   = 308000
)

type TulingReply struct {
	Content      string
	Username     string
	DefaultReply string
}

type TulingRespond struct {
	Code int                 `json:"code"`
	Text string              `json:"text"`
	Url  string              `json:"url"`
	List []map[string]string `json:"list"`
}

// NetReplier
func NewTulingReplier() *TulingReply {
	return &TulingReply{}
}

// Reply
// *args[0]: content
// args[1]: username
func (t *TulingReply) Reply(args ...interface{}) (rep string) {
	if len(args) == 1 {
		if s, ok := args[0].(string); ok {
			rep = t.replyForContetn(s)
		}
	}
	if rep == "" {
		rep = t.defaultReply()
	}
	return strings.Replace(rep, "图灵机器人", "风险观察君", -1)
}

// defaultReply
func (t *TulingReply) defaultReply() string {
	if t.DefaultReply != "" {
		return t.DefaultReply
	}
	return "O_-qoo_!p~~~"
}

// (t *TulingReply)replyForContetn 对字符串内容进行回复
func (t *TulingReply) replyForContetn(content string) string {
	if content == "" {
		log.Error("content is nil")
		return ""
	}

	url := fmt.Sprintf("%s?key=%s&info=%s", tulingApiURL, apiKey, content)
	res, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil || len(bs) == 0 {
		log.Error(err.Error(), " or res body is nil")
		return ""
	}

	v := &TulingRespond{}
	err = json.Unmarshal(bs, v)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	log.Debugf("tuling api received: %#v", v)

	return v.filter()
}

//  (t *TulingRespond) 结果处理
func (t *TulingRespond) filter() (rep string) {
	if t == nil {
		return
	}
	var size int = 5

	switch t.Code {
	case codeTextReply:
		rep = t.Text
	case codeLinkReply:
		rep = fmt.Sprintf("%s\n%s", t.Text, t.Url)
	case codeNewsReply:
		rep = t.Text + "\n"
		for i, news := range t.List {
			if i >= size-1 {
				break
			}
			rep += fmt.Sprintf("<a href=\"%s\">%s</a>(来自%s)\n", news["detailurl"], news["article"], news["source"])
		}
	case codeTrainReply:
		for i, train := range t.List {
			if i >= size-1 {
				break
			}
			rep += fmt.Sprintf("%s: %s(%s)->%s{%s}\n", train["trainnum"], train["start"], train["starttime"], train["terminal"], train["endtime"])
		}
	case codeFlightReply:
		for i, flight := range t.List {
			if i >= size-1 {
				break
			}
			rep += fmt.Sprintf("%s(%s-%s)\n", flight["flight"], flight["starttime"], flight["endtime"])
		}
	case codeMenuReply:
		rep = t.Text + "\n"
		for i, menu := range t.List {
			if i >= 2 {
				break
			}
			rep += fmt.Sprintf("<a href=\"%s\">%s</a>(原料:%s)\n", menu["detailurl"], menu["name"], menu["info"])
		}
	default:
		rep = ""
	}
	return
}
