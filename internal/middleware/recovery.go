package middleware

import (
	"common_service/global"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := getStack()
				global.Logger.Error(fmt.Sprintf("panic recover: %v; ==> %s\n", err, string(stack)))
				resp := app.NewResponse(c)
				resp.ToError(errcode.InternalError)
			}
		}()
		c.Next()
	}
}

func getStack() []byte {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return buf[:n]
}
