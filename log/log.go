package log

import (
	"Blog/config"
	"Blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.GetString("log.logFilePath")
	logFileName := config.GetString("log.logFileName")
	fileName := path.Join(logFilePath, logFileName)

	fileExist, err := util.PathExists(fileName)
	if !fileExist && err == nil {
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			fmt.Println("err:", err)
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = src
	//日志级别
	logger.SetLevel(logrus.DebugLevel)
	//日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
