package middleware

import (
	"common_service/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var locale string
		tags, _, err := language.ParseAcceptLanguage(c.GetHeader("Accept-Language"))
		if err != nil {
			global.Logger.Error(fmt.Errorf("language.ParseAcceptLanguage: %w", err))
		}
		localeList := make([]string, len(tags))
		for k, t := range tags {
			localeList[k] = t.String()
		}
		if len(localeList) > 0 {
			locale = localeList[0]
		}
		global.Trans, _ = global.UniTrans.FindTranslator(localeList...)

		// 注册参数校验翻译
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			switch locale {
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
			default:
				_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
			}
		}
		c.Next()
	}
}
