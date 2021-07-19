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

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

type TokenResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary 登录
// @Description 登录
// @Tags auth
// @Accept  json
// @Produce json
// @Param auth body service.SignInRequest true "登录"
// @Success 200 {object} app.Result{data=TokenResult}
// @Router /api/v1/sign_in [post]
func (a Auth) SignIn(c *gin.Context) {
	resp := app.NewResponse(c)
	param := service.SignInRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}

	svc := service.New(c)
	user, err := svc.CheckAuth(&param)
	if err != nil {
		switch errors.Cause(err) {
		case sql.ErrNoRows:
			resp.ToError(errcode.UserNotExists)
		case errcode.PasswordWrong:
			resp.ToError(errcode.PasswordWrong)
		case errcode.UserNotActive:
			resp.ToError(errcode.UserNotActive)
		default:
			global.Logger.Error(errors.Wrap(err, "svc.CheckAuth"))
			resp.ToError(errcode.InternalError)
		}
		return
	}

	RefreshToken, AccessToken, err := svc.GenerateAllToken(user)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "svc.GenerateAllToken"))
		resp.ToError(errcode.TokenGenerate)
		return
	}

	resp.Success(TokenResult{
		RefreshToken: RefreshToken,
		AccessToken:  AccessToken,
	})
}

// @Summary 刷新access_token
// @Description 刷新access_token
// @Tags auth
// @Accept  json
// @Produce json
// @Param auth body service.RefreshAccessTokenRequest true "刷新access_token"
// @Success 200 {object} app.Result{data=TokenResult}
// @Router /api/v1/refresh_token [post]
func (a Auth) RefreshAccessToken(c *gin.Context) {
	resp := app.NewResponse(c)
	param := service.RefreshAccessTokenRequest{}
	if err := c.ShouldBindJSON(&param); err != nil {
		resp.ToValidationError(err)
		return
	}
	svc := service.New(c)

	AccessToken, err := svc.RefreshAccessToken(&param)
	switch errors.Cause(err) {
	case nil:
		resp.Success(TokenResult{
			RefreshToken: param.RefreshToken,
			AccessToken:  AccessToken,
		})
	case errcode.TokenInvalid:
		resp.ToError(errcode.TokenInvalid)
	default:
		global.Logger.Error(errors.Wrap(err, "svc.RefreshAccessToken"))
		resp.ToError(errcode.TokenGenerate)
	}
}
