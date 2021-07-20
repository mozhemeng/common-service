package service

import (
	"common_service/global"
	"common_service/internal/dao"
	"common_service/internal/model"
	"common_service/pkg/app"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DB, global.RedisDB, global.RedisCache)
	return svc
}

func (svc *Service) GetCurrentUser() *model.User {
	claims := app.GetClaims(svc.ctx)

	user, err := svc.dao.GetUserById(claims.UserId)
	if err != nil {
		return &model.User{}
	} else {
		return user
	}
}
