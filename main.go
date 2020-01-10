package main

import (
	"Blog/config"
	"Blog/router"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadBaseConfig()
	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()
	router.Run(":" + config.GetPort())
}
