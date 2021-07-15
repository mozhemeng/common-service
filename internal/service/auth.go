package service

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"github.com/pkg/errors"
)

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (svc *Service) CheckAuth(param *SignInRequest) (*model.User, error) {
	user, err := svc.dao.GetUserByUsername(param.Username)
	if err != nil {
		return nil, errors.Wrap(err, "svc.dao.GetUserByUsername")
	}

	if user.Status == 2 {
		return nil, errcode.UserNotActive
	}

	checked := app.CheckPasswordHash(param.Password, user.PasswordHashed)
	if !checked {
		return nil, errcode.PasswordWrong
	}

	return user, nil
}
