package dao

import (
	"common_service/global"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (d *Dao) ExistsRoleById(id uint64) (bool, error) {
	var exists bool
	builder := sq.
		Select("1").
		Prefix("SELECT EXISTS (").
		From("role").
		Where(sq.Eq{"id": id}).
		Suffix(")")
	sql, args, err := builder.ToSql()
	if err != nil {
		return false, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	err = d.db.QueryRow(sql, args...).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sql exec")
	}
	return exists, nil
}
