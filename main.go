package main

import (
	"Blog/config"
	_ "Blog/config"
	"Blog/router"
)

func main() {
	app := router.InitRouter()
	app.Run(config.GetString("mainApp.host") + ":" + config.GetString("mainApp.port"))
}
