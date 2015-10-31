package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

// 点击事件结构体
type ClickEvent struct {
	Msg
	EventKey string `xml:"EventKey"`
}

func (m *Msg) ReceiveClickEvent() string {
	var msg ClickEvent
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (c *ClickEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(c)
	if err == nil {
		c.Id = id
	}
	return err
}
