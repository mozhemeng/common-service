package service

import (
	"common_service/global"
	"common_service/internal/model"
	"common_service/pkg/app"
	"github.com/pkg/errors"
)

type CreateCasbinPolicyRequest struct {
	RoleName string `json:"role_name" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required,oneof=GET POST PUT UPDATE DELETE"`
}

type ListCasbinPolicyRequest struct {
	RoleName string `form:"role_name" json:"role_name"`
}

type DeleteCasbinPolicyRequest struct {
	RoleName string `json:"role_name" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required,oneof=GET POST PUT UPDATE DELETE"`
}

func (svc *Service) ListCasbinPolicy(param *ListCasbinPolicyRequest, pager *app.Pager) ([]*model.CasbinPolicy, int, error) {
	policies, err := svc.dao.ListCasbinPolicy(param.RoleName, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "svc.dao.ListCasbinPolicy")
	}
	count, err := svc.dao.CountCasbinPolicy(param.RoleName)
	if err != nil {
		return nil, 0, errors.Wrap(err, "svc.dao.CountCasbinPolicy")
	}

	return policies, count, nil
}

func (svc *Service) CreateCasbinPolicy(param *CreateCasbinPolicyRequest) error {
	_, err := global.Enforcer.AddPolicy(param.RoleName, param.Path, param.Method)
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.AddPolicy")
	}
	// 更新完后重载policy使之生效
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.LoadPolicy")
	}
	return nil
}

func (svc *Service) DeleteCasbinPolicy(param *DeleteCasbinPolicyRequest) error {
	_, err := global.Enforcer.RemovePolicy(param.RoleName, param.Path, param.Method)
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.RemovePolicy")
	}
	// 删除完后重载policy使之生效
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.LoadPolicy")
	}

	return nil
}

func (svc *Service) ReLoadCasbin() error {
	var err error
	err = global.Enforcer.LoadModel()
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.LoadModel")
	}
	// LoadModel后需要重新LoadPolicy
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return errors.Wrap(err, "global.Enforcer.LoadPolicy")
	}

	return nil
}
