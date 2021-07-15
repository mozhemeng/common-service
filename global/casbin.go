package global

import (
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/pkg/errors"
)

var Enforcer *casbin.Enforcer

func SetupEnforcer() error {
	var err error
	a, err := sqlxadapter.NewAdapter(DB, "casbin_rule")
	if err != nil {
		return errors.Wrap(err, "NewAdapter")
	}
	Enforcer, err = casbin.NewEnforcer(CasbinSetting.ModelFilePath, a)
	if err != nil {
		return errors.Wrap(err, "NewEnforcer")
	}

	err = Enforcer.LoadPolicy()
	if err != nil {
		return errors.Wrap(err, "LoadPolicy")
	}

	return nil
}
