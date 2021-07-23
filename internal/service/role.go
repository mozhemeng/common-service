package service

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/pkg/errors"
)

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ListRoleRequest struct {
	Name string `form:"name" json:"name"`
}

type UpdateRoleUriRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateRoleBodyRequest struct {
	Description string `json:"description" binding:"required"`
}

type DeleteRoleRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (svc *Service) CreateRole(param *CreateRoleRequest) (*model.Role, error) {
	exists, err := svc.dao.ExistsRoleByName(param.Name)
	if err != nil {
		return nil, errors.Wrap(err, "dao.ExistsRoleByName")
	}
	if exists {
		return nil, errcode.RoleAlreadyExists
	}

	newID, err := svc.dao.CreateRole(param.Name, param.Description)
	if err != nil {
		return nil, errors.Wrap(err, "dao.CreateRole")
	}

	r, err := svc.dao.GetRoleById(newID)
	if err != nil {
		return nil, errors.Wrap(err, "dao.GetRoleById")
	}

	return r, nil
}

func (svc *Service) ListRole(param *ListRoleRequest, pager *app.Pager) ([]*model.Role, int, error) {
	roles, err := svc.dao.ListRole(param.Name, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "dao.ListRole")
	}
	count, err := svc.dao.CountRole(param.Name)
	if err != nil {
		return nil, 0, errors.Wrap(err, "dao.CountRole")
	}

	return roles, count, nil
}

func (svc *Service) UpdateRole(uriParam *UpdateRoleUriRequest, bodyParam *UpdateRoleBodyRequest) (*model.Role, error) {
	_, err := svc.dao.UpdateRole(uriParam.ID, bodyParam.Description)
	if err != nil {
		return nil, errors.Wrap(err, "dao.UpdateRole")
	}

	r, err := svc.dao.GetRoleById(uriParam.ID)
	if err != nil {
		return nil, errors.Wrap(err, "dao.GetRoleById")
	}

	return r, nil
}

func (svc *Service) DeleteRole(param *DeleteRoleRequest) error {

	exists, err := svc.dao.ExistsUserByRoleId(param.ID)
	if err != nil {
		return errors.Wrap(err, "dao.ExistsUserByRoleId")
	}

	if exists {
		return errcode.HaveRelativeUser
	}

	_, err = svc.dao.DeleteRole(param.ID)
	if err != nil {
		return errors.Wrap(err, "dao.DeleteRole")
	}

	return nil
}
