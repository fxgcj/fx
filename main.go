package main

import (
	"fmt"
	"github.com/ckeyer/fx/conf"
	"github.com/ckeyer/fx/lib"
	"net/http"
)

var (
	log    = lib.GetLogger()
	config *conf.Config
)

func main() {
	config = conf.GetConfig()
	if config == nil {
		panic("config is nil ")
	}
	log.Debugf("%#v\n", *config)
	mux := http.NewServeMux()
	mux.HandleFunc("/", Filter)
	//	mux.HandleFunc("/upload", Receive)
	addr := fmt.Sprintf(":%d", config.App.Port)
	log.Notice("Http is running at ", addr)
	err := http.ListenAndServe(addr, mux)
	panic(err)
}
