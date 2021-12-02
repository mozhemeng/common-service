package middleware

import (
	"bytes"
	"common_service/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 替换自带response writer
		//bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//c.Writer = bodyWriter
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 响应体（非application/json；swagger接口；等忽略显示响应体）
		//var respBody = "omitted"
		//if !strings.HasPrefix(reqUri, "/swagger") {
		//	respContentType, ok := bodyWriter.Header()["Content-Type"]
		//	if ok && len(respContentType) > 0 && strings.HasPrefix(respContentType[0], "application/json") {
		//		respBody = bodyWriter.body.String()
		//	}
		//}

		//日志格式
		global.Logger.Info().
			Int("http_status", statusCode).
			Str("total_time", latencyTime).
			Str("ip", clientIP).
			Str("method", reqMethod).
			Str("uri", reqUri).
			Msg("access")
	}
}
