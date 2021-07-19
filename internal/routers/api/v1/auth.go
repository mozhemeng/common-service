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
	AccessToken string `json:"access_token"`
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

	token, err := app.GenerateToken(user.ID, user.RoleName)
	if err != nil {
		global.Logger.Error(errors.Wrap(err, "app.GenerateToken"))
		resp.ToError(errcode.TokenGenerate)
		return
	}

	resp.Success(TokenResult{AccessToken: token})

}
