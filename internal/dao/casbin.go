package dao

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

var CasbinTableName = "casbin_rule"

func (d *Dao) ListCasbinPolicy(roleName string, page, pageSize int) ([]*model.CasbinPolicy, error) {
	many := make([]*model.CasbinPolicy, 0)

	builder := sq.
		Select("v0, v1, v2").
		From(CasbinTableName)
	condition := sq.And{}
	if roleName != "" {
		condition = append(condition, sq.Eq{"v0": roleName})
	}
	builder = builder.Where(condition)
	pageOffSet := app.GetPageOffset(page, pageSize)
	builder = builder.Offset(uint64(pageOffSet)).Limit(uint64(pageSize))

	err := d.selectSQL(builder, &many)
	if err != nil {
		return nil, errors.Wrap(err, "selectSQL")
	}

	return many, nil
}

func (d *Dao) CountCasbinPolicy(roleName string) (int, error) {
	condition := sq.And{}
	if roleName != "" {
		condition = append(condition, sq.Eq{"v0": roleName})
	}
	return d.commonCount(CasbinTableName, condition)
}
