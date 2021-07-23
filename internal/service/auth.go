package service

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"database/sql"
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
		if err == sql.ErrNoRows {
			return nil, errcode.UserNotExists
		}
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
	RefreshToken, err := app.GenerateToken(user.ID, user.Username, user.RoleName, app.RefreshTokenType)
	if err != nil {
		return "", "", errcode.TokenGenerate.WithDetails(err.Error())
	}
	AccessToken, err := app.GenerateToken(user.ID, user.Username, user.RoleName, app.AccessTokenType)
	if err != nil {
		return "", "", errcode.TokenGenerate.WithDetails(err.Error())
	}

	return RefreshToken, AccessToken, nil
}

func (svc *Service) RefreshAccessToken(param *RefreshAccessTokenRequest) (string, error) {
	claims, err := app.VerifyToken(param.RefreshToken, app.RefreshTokenType)
	if err != nil {
		return "", errcode.TokenInvalid.WithDetails(err.Error())
	}

	token, err := app.GenerateToken(claims.UserId, claims.Username, claims.RoleName, app.AccessTokenType)
	if err != nil {
		return "", errcode.TokenGenerate.WithDetails(err.Error())
	}
	return token, nil
}
