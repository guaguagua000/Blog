package middleware

import (
	"Blog/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProxyOption struct {
	//目标服务, 如http://localhost:8080
	Target string
	//要被重写为空字符串的子路径,无特殊情况,可以设置为空字符串
	PathRewrite string
}

/**
反向代理中间件: 用于代理其他平台的接口
*/
func ReverseProxyMiddleware(apiPrefix string, proxyOption ProxyOption) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Index(ctx.Request.RequestURI, apiPrefix) == 0 {
			client := &http.Client{}
			requestUrl := strings.Replace(ctx.Request.RequestURI, proxyOption.PathRewrite, "", -1)
			url := proxyOption.Target + requestUrl
			req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
			if err != nil {
				log.Logger.Error("反向代理中间件, 请求转发异常", zap.Error(err))
				return
			}
			req.Header = ctx.Request.Header
			resp, err := client.Do(req)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			for key, value := range resp.Header {
				if len(value) == 1 {
					ctx.Writer.Header().Add(key, value[0])
				}
			}
			ctx.Status(resp.StatusCode)
			ctx.Writer.Write(body)

		} else {
			ctx.Next()
		}
	}
}
