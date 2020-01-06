package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, "?")
	})
	return router
}
