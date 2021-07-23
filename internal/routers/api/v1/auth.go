package v1

import (
	"common_service/internal/service"
	"common_service/pkg/app"
	"github.com/gin-gonic/gin"
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
		resp.ToError(err)
		return
	}

	svc := service.New(c)
	user, err := svc.CheckAuth(&param)
	if err != nil {
		resp.ToError(err)
		return
	}

	RefreshToken, AccessToken, err := svc.GenerateAllToken(user)
	if err != nil {
		resp.ToError(err)
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
		resp.ToError(err)
		return
	}

	svc := service.New(c)
	AccessToken, err := svc.RefreshAccessToken(&param)
	if err != nil {
		resp.ToError(err)
		return
	}

	resp.Success(TokenResult{
		RefreshToken: param.RefreshToken,
		AccessToken:  AccessToken,
	})
}
