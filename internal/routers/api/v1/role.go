package v1

import (
	"common_service/global"
	"common_service/internal/service"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Role struct{}

func NewRole() Role {
	return Role{}
}

// @Summary 新建角色
// @Description 新建角色
// @Tags role
// @Accept  json
// @Produce json
// @Param role body service.CreateRoleRequest true "新建角色"
// @Success 200 {object} app.Result{data=model.Role} "成功"
// @Security JWT
// @Router /api/v1/roles [post]
func (r Role) Create(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.CreateRoleRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	newRole, err := svc.CreateRole(&param)
	switch errors.Cause(err) {
	case nil:
		resp.Success(newRole)
	case errcode.RoleAlreadyExists:
		resp.ToError(errcode.RoleAlreadyExists)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.CreateRole"))
		resp.ToError(errcode.InternalError)
	}
}

// @Summary 查询角色列表
// @Description 查询角色列表
// @Tags role
// @Produce json
// @Param role query service.ListRoleRequest true "查询角色列表"
// @Success 200 {object} app.PagedResult{data=[]model.Role}
// @Security JWT
// @Router /api/v1/roles [get]
func (r Role) List(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.ListRoleRequest{}
	if err := c.ShouldBindQuery(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	roles, totalCount, err := svc.ListRole(&param, &pager)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.ListRole"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.SuccessList(roles, totalCount)
}

// @Summary 更新角色
// @Description 更新角色
// @Tags role
// @Accept  json
// @Produce json
// @Param role_id path int true "角色id"
// @Param role body service.UpdateRoleBodyRequest true "更新角色"
// @Success 200 {object} app.Result{data=model.Role}
// @Security JWT
// @Router /api/v1/roles/{role_id} [put]
func (r Role) Update(c *gin.Context) {
	resp := app.NewResponse(c)

	uriParam := service.UpdateRoleUriRequest{}
	if err := c.ShouldBindUri(&uriParam); err != nil {
		resp.ToValidationError(err)
		return
	}

	bodyParam := service.UpdateRoleBodyRequest{}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	role, err := svc.UpdateRole(&uriParam, &bodyParam)
	switch errors.Cause(err) {
	case nil:
		resp.Success(role)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.UpdateRole"))
		resp.ToError(errcode.InternalError)
	}
}

// @Summary 删除角色
// @Description 删除角色
// @Tags role
// @Produce json
// @Param role_id path int true "角色id"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/roles/{role_id} [delete]
func (r Role) Delete(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.DeleteRoleRequest{}
	if err := c.ShouldBindUri(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	err := svc.DeleteRole(&param)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.DeleteRole"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.Success(nil)
}
