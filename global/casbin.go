package global

import (
	"fmt"
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
)

var Enforcer *casbin.Enforcer

func SetupEnforcer() error {
	var err error
	a, err := sqlxadapter.NewAdapter(DB, "casbin_rule")
	if err != nil {
		return fmt.Errorf("sqlxadapter.NewAdapter: %w", err)
	}
	Enforcer, err = casbin.NewEnforcer(CasbinSetting.ModelFilePath, a)
	if err != nil {
		return fmt.Errorf("casbin.NewEnforcer: %w", err)
	}

	err = Enforcer.LoadPolicy()
	if err != nil {
		return fmt.Errorf("Enforcer.LoadPolicy: %w", err)
	}

	return nil
}
