package dao

import (
	"common_service/global"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (d *Dao) selectSQL(builder sq.SelectBuilder, dest interface{}) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	err = d.db.Select(dest, sql, args...)
	if err != nil {
		return errors.Wrap(err, "Get")
	}
	return nil
}

func (d *Dao) getSQL(builder sq.SelectBuilder, dest interface{}) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	err = d.db.Get(dest, sql, args...)
	if err != nil {
		return errors.Wrap(err, "Get")
	}
	return nil
}

func (d *Dao) insertSQL(builder sq.InsertBuilder, dest *int64) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	res, err := d.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}
	*dest, err = res.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "LastInsertId")
	}
	return nil
}

func (d *Dao) updateSQL(builder sq.UpdateBuilder, dest *int64) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	res, err := d.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}
	*dest, err = res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "RowsAffected")
	}
	return nil
}

func (d *Dao) deleteSQL(builder sq.DeleteBuilder, dest *int64) error {
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	res, err := d.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "Exec")
	}
	*dest, err = res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "RowsAffected")
	}
	return nil
}

func (d *Dao) commonExists(tableName string, condition interface{}, args ...interface{}) (bool, error) {
	var exists bool

	builder := sq.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(tableName).
		Where(condition, args...).
		Suffix(")")
	err := d.getSQL(builder, &exists)
	if err != nil {
		return exists, errors.Wrap(err, "dao.getSQL")
	}
	return exists, nil
}

func (d *Dao) commonCount(tableName string, condition interface{}, args ...interface{}) (int, error) {
	var count int

	builder := sq.Select("count(*)").From(tableName).Where(condition, args...)

	err := d.getSQL(builder, &count)
	if err != nil {
		return count, errors.Wrap(err, "getSQL")
	}

	return count, nil
}

func (d *Dao) commonCreate(tableName string, columns []string, values []interface{}) (int64, error) {
	var lastId int64

	builder := sq.Insert(tableName).Columns(columns...).Values(values...)
	err := d.insertSQL(builder, &lastId)
	if err != nil {
		return lastId, errors.Wrap(err, "getSQL")
	}

	return lastId, nil
}

func (d *Dao) commonUpdate(tableName string, setMap map[string]interface{}, condition interface{}, args ...interface{}) (int64, error) {
	var affectRows int64
	builder := sq.Update(tableName).Where(condition, args...)
	for k, v := range setMap {
		builder = builder.Set(k, v)
	}
	err := d.updateSQL(builder, &affectRows)
	if err != nil {
		return affectRows, errors.Wrap(err, "updateSQL")
	}

	return affectRows, nil
}

func (d *Dao) commonDelete(tableName string, condition interface{}, args ...interface{}) (int64, error) {
	var affectRows int64

	builder := sq.Delete(tableName).Where(condition, args...)
	err := d.deleteSQL(builder, &affectRows)
	if err != nil {
		return affectRows, errors.Wrap(err, "deleteSQL")
	}

	return affectRows, nil
}
