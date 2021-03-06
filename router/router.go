package router

import (
	"Blog/log"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(log.LoggerToFile())
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, "萱仔很胖")
	})
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.File("public\\img\\favicon.ico")
	})
	return router
}
