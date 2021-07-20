package app

import (
	"common_service/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type TokenType string

const (
	RefreshTokenType TokenType = "refresh"
	AccessTokenType  TokenType = "access"
)

type Claims struct {
	UserId   int64     `json:"user_id"`
	Username string    `json:"username"`
	RoleName string    `json:"role_name"`
	Type     TokenType `json:"type"` // refresh or access
	jwt.StandardClaims
}

func GenerateToken(userId int64, username string, roleName string, tokenType TokenType) (string, error) {
	nowTime := time.Now()
	var expireTime time.Time
	if tokenType == RefreshTokenType {
		expireTime = nowTime.Add(global.JWTSetting.RefreshExpire)
	} else {
		expireTime = nowTime.Add(global.JWTSetting.AccessExpire)
	}
	claims := Claims{
		UserId:   userId,
		Username: username,
		RoleName: roleName,
		Type:     tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.JWTSetting.Secret))
	return token, err
}

func VerifyToken(token string, tokenType TokenType) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.Secret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt.ParseWithClaims")
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid && claims.Type == tokenType {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}

func ExtractToken(c *gin.Context) (string, error) {
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
	if token == "" {
		return token, errors.New("no token found")
	} else {
		return token, nil
	}
}
