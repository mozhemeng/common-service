package middleware

import (
	"common_service/global"
	"common_service/pkg/app"
	"common_service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
)

func IpRateLimiter(rateFormat string) gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(rateFormat)
	if err != nil {
		log.Fatal(err)
	}

	store, err := sredis.NewStoreWithOptions(global.RedisDB, limiter.StoreOptions{
		Prefix: "limiter_ip_" + rateFormat,
	})
	if err != nil {
		log.Fatal(err)
	}

	l := limiter.New(store, rate)

	// default use c.ClientIP() as key
	return mgin.NewMiddleware(
		l,
		mgin.WithErrorHandler(errorHandler),
		mgin.WithLimitReachedHandler(reachedHandler))
}

// must use after Authorization middleware
func UserRateLimiter(rateFormat string) gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(rateFormat)
	if err != nil {
		log.Fatal(err)
	}

	store, err := sredis.NewStoreWithOptions(global.RedisDB, limiter.StoreOptions{
		Prefix: "limiter_user_" + rateFormat,
	})
	if err != nil {
		log.Fatal(err)
	}

	l := limiter.New(store, rate)

	return mgin.NewMiddleware(
		l,
		mgin.WithErrorHandler(errorHandler),
		mgin.WithLimitReachedHandler(reachedHandler),
		mgin.WithKeyGetter(userKeyGetter),
		mgin.WithExcludedKey(userKeyExcluded))
}

func errorHandler(c *gin.Context, err error) {
	global.Logger.Error(fmt.Errorf("rate limit: %w", err))
	resp := app.NewResponse(c)
	resp.ToError(errcode.RateLimitExceeded)
}

func reachedHandler(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.ToError(errcode.RateLimitExceeded)
}

func userKeyGetter(c *gin.Context) string {
	claims := app.GetClaims(c)
	return claims.Username
}

func userKeyExcluded(key string) bool {
	if key == global.AppSetting.RootUsername {
		return true
	}
	return false
}
