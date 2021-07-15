package v1

import (
	"common_service/global"
	"common_service/internal/service"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @Summary 新建用户
// @Description 新建用户
// @Tags user
// @Accept  json
// @Produce json
// @Param user body service.CreateUserRequest true "新建用户"
// @Success 200 {object} app.Result{data=model.User} "成功"
// @Security JWT
// @Router /api/v1/users [post]
func (u User) Create(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.CreateUserRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	newUser, err := svc.CreateUser(&param)
	switch errors.Cause(err) {
	case nil:
		resp.Success(newUser)
	case errcode.UserAlreadyExists:
		resp.ToError(errcode.UserAlreadyExists)
	case errcode.RoleNotExists:
		resp.ToError(errcode.RoleNotExists)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.CreateUser"))
		resp.ToError(errcode.InternalError)
	}
}

// @Summary 查询单个用户
// @Description 查询单个用户
// @Tags user
// @Produce json
// @Param user_id path int true "用户id"
// @Success 200 {object} app.Result{data=model.User}
// @Security JWT
// @Router /api/v1/users/{user_id} [get]
func (u User) GetByID(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.GetUserByIDRequest{}
	if err := c.ShouldBindUri(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	user, err := svc.GetUserByID(&param)
	switch errors.Cause(err) {
	case nil:
		resp.Success(user)
	case sql.ErrNoRows:
		resp.ToError(errcode.UserNotExists)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.GetUserByID"))
		resp.ToError(errcode.InternalError)
	}
}

// @Summary 查询用户列表
// @Description 查询用户列表
// @Tags user
// @Produce json
// @Param user query service.ListUserRequest true "查询用户列表"
// @Success 200 {object} app.PagedResult{data=[]model.User}
// @Security JWT
// @Router /api/v1/users [get]
func (u User) List(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.ListUserRequest{}
	if err := c.ShouldBindQuery(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
	}
	users, totalCount, err := svc.ListUser(&param, &pager)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.ListUser"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.SuccessList(users, totalCount)
}

// @Summary 更新用户
// @Description 更新用户
// @Tags user
// @Accept  json
// @Produce json
// @Param user_id path int true "用户id"
// @Param user body service.UpdateUserBodyRequest true "更新用户"
// @Success 200 {object} app.Result{data=model.User}
// @Security JWT
// @Router /api/v1/users/{user_id} [put]
func (u User) Update(c *gin.Context) {
	resp := app.NewResponse(c)

	uriParam := service.UpdateUserUriRequest{}
	if err := c.ShouldBindUri(&uriParam); err != nil {
		resp.ToValidationError(err)
		return
	}

	bodyParam := service.UpdateUserBodyRequest{}
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	user, err := svc.UpdateUser(&uriParam, &bodyParam)
	switch errors.Cause(err) {
	case nil:
		resp.Success(user)
	case errcode.RoleNotExists:
		resp.ToError(errcode.RoleNotExists)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.UpdateUser"))
		resp.ToError(errcode.InternalError)
	}
}

// @Summary 删除用户
// @Description 删除用户
// @Tags user
// @Produce json
// @Param user_id path int true "用户id"
// @Success 200 {object} app.Result
// @Security JWT
// @Router /api/v1/users/{user_id} [delete]
func (u User) Delete(c *gin.Context) {
	resp := app.NewResponse(c)

	param := service.DeleteUserRequest{}
	if err := c.ShouldBindUri(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	err := svc.DeleteUser(&param)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.DeleteUser"))
		resp.ToError(errcode.InternalError)
		return
	}
	resp.Success(nil)
}
