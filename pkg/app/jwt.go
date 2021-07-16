package app

import (
	"common_service/global"
	"common_service/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	RoleName string `json:"role_name"`
	jwt.StandardClaims
}

func GenerateToken(userId int64, roleName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire * time.Hour)
	claims := Claims{
		UserId:   userId,
		RoleName: roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.JWTSetting.Secret))
	return token, err
}

func VerifyToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.Secret), nil
	})
	if err != nil {
		return nil, errcode.TokenInvalid.WithDetails(errors.Wrap(err, "jwt.ParseWithClaims").Error())
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errcode.TokenInvalid
}

func ExtractToken(c *gin.Context) string {
	var token string
	if s, exist := c.GetQuery("token"); exist {
		token = s
	} else {
		authorization := c.GetHeader("Authorization")
		a := strings.Split(authorization, " ")
		if len(a) == 2 {
			token = a[1]
		}
	}
	return token
}
