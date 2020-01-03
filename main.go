package main

import (
	"Blog/config"
	"Blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	config.LoadBaseConfig()
	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetPort()),
		Handler:        router,
		ReadTimeout:    config.GetReadTimeOut(),
		WriteTimeout:   config.GetWriteTimeOut(),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
