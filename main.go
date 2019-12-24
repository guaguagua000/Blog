package main

import (
	"gin_learning/config"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.Run(config.GetPort())
}
