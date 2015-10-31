package wechat

import (
	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/fx/conf"
	"github.com/ckeyer/fx/lib"
	"github.com/ckeyer/fx/reply"
	_ "github.com/go-sql-driver/mysql"
)

var (
	config       *conf.Config
	log          = lib.GetLogger()
	access_token string
	replier      = reply.GetReplier()
)

func init() {
	config = conf.GetConfig()
	if config == nil {
		log.Panic("config is nil")
	}
	//	RegDB()
}

func RegDB() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", config.Mysql.GetConnStr())

	orm.RegisterModel(new(TextMsg),
		new(ImageMsg),
		new(LinkMsg),
		new(LocationMsg),
		new(VideoMsg),
		new(VoiceMsg))
}
