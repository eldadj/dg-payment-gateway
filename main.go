package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/middleware"
	"github.com/eldadj/dgpg/models"
	_ "github.com/eldadj/dgpg/routers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if err := models.InitDB(); err != nil {
		panic(err)
	}
	defer models.CloseDB()
	web.BConfig.Log.AccessLogs = true
	middleware.Register()
	web.Run()
}
