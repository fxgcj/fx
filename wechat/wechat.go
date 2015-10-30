package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/fx/conf"
	"github.com/ckeyer/fx/lib"
	_ "github.com/go-sql-driver/mysql"
)

var (
	config       *conf.Config
	log          = lib.GetLogger()
	access_token string
)

func init() {
	config = conf.GetConfig()
	if config == nil {
		log.Panic("config is nil")
	}

}

func RegDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", config.Mysql.GetConnStr())

	orm.RegisterModel()
}

func ReceiveMsg(content string) (r string) {
	var msgtype MsgType
	err := xml.Unmarshal([]byte(content), &msgtype)
	if err != nil {
		return
	}
	switch msgtype.MsgType {
	// case "text", "image", "voice", "video", "location", "link":
	case "event":
		r = event.ReceiveEvent(content, msgtype.Event)
	default:
		r = msg.ReceiveMsg(content, msgtype.Event)
	}
	return
}
