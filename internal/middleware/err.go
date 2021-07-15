package middleware

import (
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var e *errcode.ApiError
		lastErr := c.Errors.Last()
		if lastErr != nil {
			err := lastErr.Err
			switch err.(type) {
			case *errcode.ApiError:
				e = err.(*errcode.ApiError)
			default:
				e = errcode.InternalError.WithDetails(err.Error())
			}
			c.JSON(e.StatusCode(), e.JSON())
		}
	}
}
