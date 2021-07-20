package v1

import (
	"common_service/pkg/app"
	"github.com/gin-gonic/gin"
)

type Example struct{}

func NewExample() Example {
	return Example{}
}

// @Summary 用户访问速率限制
// @Description 用户访问速率限制
// @Tags example
// @Produce json
// @Success 200 {object} app.Result "成功"
// @Security JWT
// @Router /api/v1/example/rate-limit [get]
func (e Example) UserRateLimit(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.Success(gin.H{"msg": "This endpoint have rate limit"})
}
