package main

import (
	"log"

	"gitlab.com/xpresslane1/xintegrations/internal/conf"
	xconf "gitlab.com/xpresslane1/xcommon/conf"
	"gitlab.com/xpresslane1/xcommon/middlewares"
	"gitlab.com/xpresslane1/xintegrations/internal/app"
)

var AppConfig *xconf.XAppConfig

func init() {
	AppConfig = conf.NewConfig()
	middlewares.Init(AppConfig, true, false, false)
}

func main() {
	log.Println("About to start the API server...")
	defer AppConfig.DB.Close()
	app.StartApplication(AppConfig)
}