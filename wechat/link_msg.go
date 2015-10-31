package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type LinkMsg struct {
	Msg
	Title       string `xml: "Title"`
	Description string `xml: "Description"`
	Url         string `xml: "Url"`
	MsgId       int64  `xml: "MsgId"`
}

func (m *Msg) ReceiveLinkMsg() string {
	var msg LinkMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (l *LinkMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(l)
	if e != nil {
		return e
	}
	return nil
}
