package v1

import (
	"common_service/internal/service"
	"common_service/pkg/app"
	"github.com/gin-gonic/gin"
)

type PermPolicy struct{}

func NewPermPolicy() PermPolicy {
	return PermPolicy{}
}

// @Summary 新建权限规则
// @Description 新建权限规则
// @Tags permission
// @Accept  json
// @Produce json
// @Param policy body service.CreatePermPolicyRequest true "新建权限规则"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/perm/policies [post]
func (p PermPolicy) Create(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.CreatePermPolicyRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToError(err)
		return
	}

	svc := service.New(c)
	err := svc.CreatePermPolicy(&param)
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.Success(nil)
}

// @Summary 查询权限规则列表
// @Description 查询权限规则列表
// @Tags permission
// @Produce json
// @Param policy query service.ListPermPolicyRequest true "查询权限规则列表"
// @Success 200 {object} app.Result{data=[]model.PermPolicy}
// @Security JWT
// @Router /api/v1/perm/policies [get]
func (p PermPolicy) List(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.ListPermPolicyRequest{}
	if err := c.ShouldBindQuery(&param); err != nil {
		resp.ToError(err)
		return
	}

	svc := service.New(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	policies, totalCount, err := svc.ListPermPolicy(&param, &pager)
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.SuccessList(policies, totalCount)
}

// @Summary 删除权限规则
// @Description 删除权限规则
// @Tags permission
// @Produce json
// @Param policy body service.DeletePermPolicyRequest true "删除权限规则"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/perm/policies [delete]
func (p PermPolicy) Delete(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.DeletePermPolicyRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToError(err)
		return
	}

	svc := service.New(c)
	err := svc.DeletePermPolicy(&param)
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.Success(nil)
}

// @Summary 重新载入权限规则
// @Description 重新载入权限规则
// @Tags permission
// @Produce json
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/perm/policies/reload [get]
func (p PermPolicy) ReLoad(c *gin.Context) {
	resp := app.NewResponse(c)

	svc := service.New(c)
	err := svc.ReLoadPerm()
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.Success(nil)
}
