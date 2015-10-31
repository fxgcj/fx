package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type TextMsg struct {
	Msg
	Content string `xml:"Content"`
	MsgId   int64  `xml:"MsgId"`
}

func (this *TextMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(this)
	if e != nil {
		return e
	}
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
	log.Debug(msg.Content)

	m.WriteText(`/::D/::D
服务器维护中
/::D/::D`)
}
