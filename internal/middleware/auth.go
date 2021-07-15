package middleware

import (
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var claims *app.Claims
		resp := app.NewResponse(c)
		token := app.ExtractToken(c)
		if token == "" {
			err = errcode.TokenInvalid
		} else {
			claims, err = app.VerifyToken(token)
		}
		if err != nil {
			resp.ToError(err.(*errcode.ApiError))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
