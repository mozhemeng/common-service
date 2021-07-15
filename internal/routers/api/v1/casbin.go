package v1

import (
	"common_service/global"
	"common_service/internal/service"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type CasbinPolicy struct{}

func NewCasbinPolicy() CasbinPolicy {
	return CasbinPolicy{}
}

// @Summary 新建权限规则
// @Description 新建权限规则
// @Tags permission
// @Accept  json
// @Produce json
// @Param policy body service.CreateCasbinPolicyRequest true "新建权限规则"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/casbin/policies [post]
func (p CasbinPolicy) Create(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.CreateCasbinPolicyRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	err := svc.CreateCasbinPolicy(&param)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.CreateCasbinPolicy"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.Success(nil)
}

// @Summary 查询权限规则列表
// @Description 查询权限规则列表
// @Tags permission
// @Produce json
// @Param policy query service.ListCasbinPolicyRequest true "查询权限规则列表"
// @Success 200 {object} app.Result{data=[]model.CasbinPolicy}
// @Security JWT
// @Router /api/v1/casbin/policies [get]
func (p CasbinPolicy) List(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.ListCasbinPolicyRequest{}
	if err := c.ShouldBindQuery(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	policies, totalCount, err := svc.ListCasbinPolicy(&param, &pager)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.ListCasbinPolicy"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.SuccessList(policies, totalCount)
}

// @Summary 删除权限规则
// @Description 删除权限规则
// @Tags permission
// @Produce json
// @Param policy body service.DeleteCasbinPolicyRequest true "删除权限规则"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/casbin/policies [delete]
func (p CasbinPolicy) Delete(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.DeleteCasbinPolicyRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	err := svc.DeleteCasbinPolicy(&param)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.DeleteCasbinPolicy"))
		resp.ToError(errcode.InternalError)
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
// @Router /api/v1/casbin/policies/reload [get]
func (p CasbinPolicy) ReLoad(c *gin.Context) {
	resp := app.NewResponse(c)
	svc := service.New(c)
	err := svc.ReLoadCasbin()
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.ReLoadCasbin"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.Success(nil)
}
