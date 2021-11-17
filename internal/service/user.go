package service

import (
	"common_service/global"
	"common_service/internal/dao"
	"common_service/internal/model"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"fmt"
	"github.com/go-redis/cache/v8"
)

type GetUserByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Nickname   string `json:"nickname" binding:"required"`
	Status     *uint  `json:"status" binding:"required,oneof=0 1"`
	RoleId     int64  `json:"role_id" binding:"required"`
}

type ListUserRequest struct {
	Nickname string `form:"nickname" json:"nickname"`
	Status   *uint  `form:"status" json:"status"`
	RoleId   int64  `form:"role_id" json:"role_id"`
}

type UpdateUserUriRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateUserBodyRequest struct {
	Nickname string `json:"nickname" binding:"required"`
	Status   *uint  `json:"status" binding:"required,oneof=0 1"`
	RoleId   int64  `json:"role_id" binding:"required"`
}

type DeleteUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (svc *Service) GetUserByID(param *GetUserByIDRequest) (*model.User, error) {
	u, err := svc.dao.GetUserInCache(param.ID)
	if err == nil {
		return u, nil
	}
	if err != cache.ErrCacheMiss {
		global.Logger.Error(fmt.Errorf("svc.dao.GetUserInCache: %w", err))
	}

	u, err = svc.dao.GetUserById(param.ID)
	if err != nil {
		if dao.IsNoRowFound(err) {
			return nil, errcode.UserNotExists
		}
		return nil, fmt.Errorf("svc.dao.GetUserByID: %w", err)
	}

	err = svc.dao.SetUserInCache(u, 0)
	if err != nil {
		global.Logger.Error(fmt.Errorf("svc.dao.SetUserInCache: %w", err))
	}

	return u, nil
}

func (svc *Service) CreateUser(param *CreateUserRequest) (*model.User, error) {
	exists, err := svc.dao.ExistsUserByUsername(param.Username)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.ExistsUserByUsername: %w", err)
	}
	if exists {
		return nil, errcode.UserAlreadyExists
	}

	roleExists, err := svc.dao.ExistsRoleById(param.RoleId)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.ExistsRoleById: %w", err)
	}
	if !roleExists {
		return nil, errcode.RoleNotExists
	}

	passwordHashed, err := app.HashPassword(param.Password)
	if err != nil {
		return nil, fmt.Errorf("app.HashPassword: %w", err)
	}
	newID, err := svc.dao.CreateUser(param.Username, passwordHashed, param.Nickname, *param.Status, param.RoleId)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.CreateUser: %w", err)
	}

	u, err := svc.dao.GetUserById(newID)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.GetUserById: %w", err)
	}

	err = svc.dao.SetUserInCache(u, 0)
	if err != nil {
		global.Logger.Error(fmt.Errorf("svc.dao.SetUserInCache: %w", err))
	}

	return u, nil
}

func (svc *Service) ListUser(param *ListUserRequest, pager *app.Pager) ([]*model.User, int, error) {
	users, err := svc.dao.ListUser(param.Nickname, param.Status, param.RoleId, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("svc.dao.ListUser: %w", err)
	}
	count, err := svc.dao.CountUser(param.Nickname, param.Status, param.RoleId)
	if err != nil {
		return nil, 0, fmt.Errorf("svc.dao.CountUser: %w", err)
	}

	return users, count, nil
}

func (svc *Service) UpdateUser(uriParam *UpdateUserUriRequest, bodyParam *UpdateUserBodyRequest) (*model.User, error) {
	roleExists, err := svc.dao.ExistsRoleById(bodyParam.RoleId)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.ExistsRoleById: %w", err)
	}
	if !roleExists {
		return nil, errcode.RoleNotExists
	}

	_, err = svc.dao.UpdateUser(uriParam.ID, bodyParam.Nickname, *bodyParam.Status, bodyParam.RoleId)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.UpdateUser: %w", err)
	}

	u, err := svc.dao.GetUserById(uriParam.ID)
	if err != nil {
		return nil, fmt.Errorf("svc.dao.GetUserById: %w", err)
	}

	err = svc.dao.SetUserInCache(u, 0)
	if err != nil {
		global.Logger.Error(fmt.Errorf("svc.dao.SetUserInCache: %w", err))
	}

	return u, nil
}

func (svc *Service) DeleteUser(param *DeleteUserRequest) error {
	_, err := svc.dao.DeleteUser(param.ID)
	if err != nil {
		return fmt.Errorf("svc.dao.DeleteUser: %w", err)
	}
	err = svc.dao.DeleteUserInCache(param.ID)
	if err != nil {
		global.Logger.Error(fmt.Errorf("svc.dao.DeleteUserInCache: %w", err))
	}
	return nil
}
