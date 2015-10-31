package wechat

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Msg struct {
	content []byte
	w       http.ResponseWriter
	req     *http.Request

	Id int64 `xml:"-"`

	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`

	CreatedLocal time.Time `orm:"auto_now_add;type(datetime)"`
}

func Receive(w http.ResponseWriter, req *http.Request) {
	msg := &Msg{}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = xml.Unmarshal(bs, msg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	msg.content = bs
	msg.w = w
	msg.req = req

	switch msg.MsgType {
	case "event":
		msg.ReceiveEvent()
	default:
		msg.ReceiveMsg()
	}
}

// ReceiveMsg 消息的接受
func (m *Msg) ReceiveMsg() {
	switch m.MsgType {
	case "text":
		m.ReceiveTextMsg()
	case "image":
		m.ReceiveImageMsg()
	case "voice":
		m.ReceiveVoiceMsg()
	case "video":
		m.ReceiveVideoMsg()
	case "location":
		m.ReceiveLocationMsg()
	case "link":
		m.ReceiveLinkMsg()
	default:
		m.WriteText("观察君出海去啦~~~")
	}
	return
}

func (m *Msg) ReceiveEvent() {
	switch m.Event {
	case "subscribe":
		m.ReceiveSubscribeEvent()
	case "unsubscribe":
		m.ReceiveUnsubscribeEvent()
	case "SCAN":
		m.ReceiveScanEvent()
	case "LOCATION":
		m.ReceiveLocationEvent()
	case "CLICK":
		m.ReceiveClickEvent()
	case "VIEW":
		m.ReceiveViewEvent()
	}
}

func (m *Msg) WriteText(data string) {
	xmlfmt := `<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>`
	body := fmt.Sprintf(xmlfmt, m.FromUserName, m.ToUserName, time.Now().Unix(), data)
	log.Debug("send body: ", body)
	m.w.WriteHeader(http.StatusOK)
	m.w.Write([]byte(body))
}
