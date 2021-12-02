package middleware

import (
	"common_service/global"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := app.NewResponse(c)

		obj := c.Request.URL.Path
		act := c.Request.Method
		claims := c.MustGet("claims").(*app.Claims)
		sub := claims.RoleName

		global.Logger.Debug().
			Str("sub", sub).
			Str("obj", obj).
			Str("act", act).
			Msg("perm policy")

		pass, err := global.Enforcer.Enforce(sub, obj, act)
		if err != nil {
			global.Logger.Err(fmt.Errorf("global.Enforcer.Enforce: %w", err)).Send()
			resp.ToError(errcode.PermissionDeny)
			return
		}
		if pass {
			c.Next()
		} else {
			resp.ToError(errcode.PermissionDeny)
			return
		}
	}
}
