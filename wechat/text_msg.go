package wechat

import (
	"encoding/xml"
)

type TextMsg struct {
	Msg
	Content string `xml:"Content"`
	MsgId   int64  `xml:"MsgId"`
}

func (t *TextMsg) Insert() error {
	return nil
}

func (m *Msg) ReceiveTextMsg() {
	var msg TextMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	msg.Insert()
	ret := replier.Reply(msg.Content)
	log.Notice("received: ", msg.Content)
	log.Notice("replied: ", ret)
	m.WriteText(ret)
}
