package service

import (
	"common_service/global"
	"common_service/internal/model"
	"common_service/pkg/app"
	"fmt"
)

type CreatePermPolicyRequest struct {
	RoleName string `json:"role_name" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required,oneof=GET POST PUT UPDATE DELETE"`
}

type ListPermPolicyRequest struct {
	RoleName string `form:"role_name" json:"role_name"`
}

type DeletePermPolicyRequest struct {
	RoleName string `json:"role_name" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required,oneof=GET POST PUT UPDATE DELETE"`
}

func (svc *Service) ListPermPolicy(param *ListPermPolicyRequest, pager *app.Pager) ([]*model.PermPolicy, int, error) {
	policies, err := svc.dao.ListPermPolicy(param.RoleName, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("svc.dao.ListPermPolicy: %w", err)
	}
	count, err := svc.dao.CountPermPolicy(param.RoleName)
	if err != nil {
		return nil, 0, fmt.Errorf("svc.dao.CountPermPolicy: %w", err)
	}

	return policies, count, nil
}

func (svc *Service) CreatePermPolicy(param *CreatePermPolicyRequest) error {
	_, err := global.Enforcer.AddPolicy(param.RoleName, param.Path, param.Method)
	if err != nil {
		return fmt.Errorf("global.Enforcer.AddPolicy: %w", err)
	}
	// 更新完后重载policy使之生效
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("global.Enforcer.LoadPolicy: %w", err)
	}
	return nil
}

func (svc *Service) DeletePermPolicy(param *DeletePermPolicyRequest) error {
	_, err := global.Enforcer.RemovePolicy(param.RoleName, param.Path, param.Method)
	if err != nil {
		return fmt.Errorf("global.Enforcer.RemovePolicy: %w", err)
	}
	// 删除完后重载policy使之生效
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("global.Enforcer.LoadPolicy: %w", err)
	}

	return nil
}

func (svc *Service) ReLoadPerm() error {
	var err error
	err = global.Enforcer.LoadModel()
	if err != nil {
		return fmt.Errorf("global.Enforcer.LoadModel: %w", err)
	}
	// LoadModel后需要重新LoadPolicy
	err = global.Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("global.Enforcer.LoadPolicy: %w", err)
	}

	return nil
}
