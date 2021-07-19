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

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
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

func (svc *Service) GenerateAllToken(user *model.User) (string, string, error) {
	RefreshToken, err := app.GenerateToken(user.ID, user.RoleName, app.RefreshTokenType)
	if err != nil {
		return "", "", errors.Wrap(err, "app.GenerateToken(refresh)")
	}
	AccessToken, err := app.GenerateToken(user.ID, user.RoleName, app.AccessTokenType)
	if err != nil {
		return "", "", errors.Wrap(err, "app.GenerateToken(access)")
	}

	return RefreshToken, AccessToken, nil
}

func (svc *Service) RefreshAccessToken(param *RefreshAccessTokenRequest) (string, error) {
	claims, err := app.VerifyToken(param.RefreshToken, app.RefreshTokenType)
	if err != nil {
		return "", errors.Wrap(err, "app.VerifyToken")
	}

	return app.GenerateToken(claims.UserId, claims.RoleName, app.AccessTokenType)
}
