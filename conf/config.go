package conf

import (
	"encoding/json"
	"github.com/ckeyer/commons/db"
	"io/ioutil"
	"os"
)

var (
	config      *Config
	config_file = "conf/v1.json"
)

type Config struct {
	App    AppConfig      `json:"app"`
	Mysql  db.MySqlConfig `json:"mysql"`
	WeChat WeChat         `json:"wechat"`
}

type AppConfig struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Level    string `json:"level"`
	LogLevel string `json:"log_level"`
}

type WeChat struct {
	Token          string `json:"token"`
	AppId          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	EncodingAESKey string `json:"encoding_aes_key"`
}

// init 初始化：加载配置内容
func init() {
	if config == nil {

	}
}

// LoadConfigFile 从文件加载配置内容
func LoadConfigFile(file string) (c *Config, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	c = new(Config)
	err = json.Unmarshal(bs, c)
	return
}

func GetConfig(file ...string) *Config {
	var config_fpath string

	if len(file) == 0 {
		config_fpath = config_file
	} else {
		config_fpath = file[0]
	}
	var err error
	config, err = LoadConfigFile(config_fpath)
	if err != nil {
		println("Load Config file failed, ", err.Error())
		return nil
	}
	return config
}
