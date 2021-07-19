package middleware

import (
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := app.NewResponse(c)

		token, err := app.ExtractToken(c)
		if err != nil {
			resp.ToError(errcode.TokenInvalid.WithDetails(err.Error()))
			return
		}

		claims, err := app.VerifyToken(token, app.AccessTokenType)
		if err != nil {
			resp.ToError(errcode.TokenInvalid.WithDetails(err.Error()))
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
