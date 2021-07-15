package global

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
)

var Trans ut.Translator
var UniTrans *ut.UniversalTranslator

func SetupUniTrans() error {
	var err error

	e := en.New()
	UniTrans = ut.New(e, e, zh.New())

	err = UniTrans.Import(ut.FormatJSON, "locales")
	if err != nil {
		return errors.Wrap(err, "UniTrans.Import")
	}

	err = UniTrans.VerifyTranslations()
	if err != nil {
		return errors.Wrap(err, "UniTrans.VerifyTranslations")
	}

	return nil
}
