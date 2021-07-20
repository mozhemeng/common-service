package app

import (
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) *Claims {
	return c.MustGet("claims").(*Claims)
}
