package middleware

import (
	"common_service/global"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := app.NewResponse(c)

		obj := c.Request.URL.Path
		act := c.Request.Method
		claims := c.MustGet("claims").(*app.Claims)
		sub := claims.RoleName

		global.Logger.WithFields(logrus.Fields{
			"sub": sub,
			"obj": obj,
			"act": act,
		}).Debug("perm policy")

		pass, err := global.Enforcer.Enforce(sub, obj, act)
		if err != nil {
			global.Logger.Error(errors.Wrap(err, "Enforce"))
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
