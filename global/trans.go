package global

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator
var UniTrans *ut.UniversalTranslator

func SetupUniTrans() error {
	var err error

	e := en.New()
	UniTrans = ut.New(e, e, zh.New())

	err = UniTrans.Import(ut.FormatJSON, "locales")
	if err != nil {
		return fmt.Errorf("UniTrans.Import: %w", err)
	}

	err = UniTrans.VerifyTranslations()
	if err != nil {
		return fmt.Errorf("UniTrans.VerifyTranslations: %w", err)
	}

	return nil
}
