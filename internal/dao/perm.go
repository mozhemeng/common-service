package dao

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (d *Dao) ListPermPolicy(roleName string, page, pageSize int) ([]*model.PermPolicy, error) {
	many := make([]*model.PermPolicy, 0)

	builder := sq.
		Select("v0, v1, v2").
		From(PermTableName)
	condition := sq.And{}
	if roleName != "" {
		condition = append(condition, sq.Eq{"v0": roleName})
	}
	builder = builder.Where(condition)
	pageOffSet := app.GetPageOffset(page, pageSize)
	builder = builder.Offset(uint64(pageOffSet)).Limit(uint64(pageSize))

	err := d.selectSql(builder, &many)
	if err != nil {
		return nil, errors.Wrap(err, "selectSql")
	}

	return many, nil
}

func (d *Dao) CountPermPolicy(roleName string) (int, error) {
	condition := sq.And{}
	if roleName != "" {
		condition = append(condition, sq.Eq{"v0": roleName})
	}
	return d.commonCount(PermTableName, condition)
}
