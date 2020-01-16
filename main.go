package main

import (
	"Blog/config"
	"Blog/dao"
	"Blog/router"
)

func main() {
	config.Init()
	dao.InitMysql()
	defer dao.DB.Close()
	app := router.InitRouter()
	app.Run(config.GetString("mainApp.host") + ":" + config.GetString("mainApp.port"))
}
