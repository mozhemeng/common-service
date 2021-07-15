package dao

import (
	"common_service/global"
	"common_service/internal/model"
	"common_service/pkg/app"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (d *Dao) ListCasbinPolicy(roleName string, page, pageSize int) ([]*model.CasbinPolicy, error) {
	builder := sq.
		Select("v0, v1, v2").
		From("casbin_rule")
	if roleName != "" {
		builder = builder.Where(sq.Eq{"casbin_rule.v0": roleName})
	}
	pageOffSet := app.GetPageOffset(page, pageSize)
	builder = builder.Offset(uint64(pageOffSet)).Limit(uint64(pageSize))
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	var Policies []*model.CasbinPolicy
	err = d.db.Select(&Policies, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "sql exec")
	}
	if Policies == nil {
		Policies = make([]*model.CasbinPolicy, 0)
	}
	return Policies, nil
}

func (d *Dao) CountCasbinPolicy(roleName string) (int, error) {
	builder := sq.Select("count(*)").From("casbin_rule")
	if roleName != "" {
		builder = builder.Where(sq.Eq{"casbin_rule.v0": roleName})
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	var count int
	err = d.db.Get(&count, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sql exec")
	}
	return count, nil
}
