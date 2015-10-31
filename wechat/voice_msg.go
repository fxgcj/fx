package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type VoiceMsg struct {
	Msg
	MediaId string `xml:"MediaId"`
	Format  string `xml:"Format"`
	MsgId   int64  `xml:"MsgId"`
}

func (m *Msg) ReceiveVoiceMsg() string {
	var msg VoiceMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (v *VoiceMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(v)
	if e != nil {
		return e
	}
	return nil
}
