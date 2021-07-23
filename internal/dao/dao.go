package dao

import (
	"common_service/global"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	UserTableName = "user"
	RoleTableName = "role"
	PermTableName = "casbin_rule"
)

type Dao struct {
	db    *sqlx.DB
	rdb   *redis.Client
	cache *cache.Cache
}

func New(db *sqlx.DB, rdb *redis.Client, cache *cache.Cache) *Dao {
	return &Dao{
		db:    db,
		rdb:   rdb,
		cache: cache,
	}
}

// sql
func IsNoRowFound(err error) bool {
	return errors.Cause(err) == sql.ErrNoRows
}

func logSql(s time.Time, query string, args []interface{}) {
	global.Logger.WithFields(logrus.Fields{
		"sql":      query,
		"args":     args,
		"duration": time.Now().Sub(s),
	}).Debug("sql log")
}

func (d *Dao) selectSql(builder sq.Sqlizer, dest interface{}) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "builder.ToSql")
	}

	defer logSql(time.Now(), query, args)
	err = d.db.Select(dest, query, args...)
	if err != nil {
		return errors.Wrap(err, "db.Select")
	}

	return nil
}

func (d *Dao) getSql(builder sq.Sqlizer, dest interface{}) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "builder.ToSql")
	}

	defer logSql(time.Now(), query, args)
	err = d.db.Get(dest, query, args...)
	if err != nil {
		return errors.Wrap(err, "db.Get")
	}

	return nil
}

func (d *Dao) execSql(builder sq.Sqlizer) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "builder.ToSql")
	}

	defer logSql(time.Now(), query, args)
	res, err := d.db.Exec(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.Exec")
	}

	return res, nil
}

func (d *Dao) txExec(tx *sql.Tx, builder sq.Sqlizer) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "builder.ToSql")
	}

	defer logSql(time.Now(), query, args)
	res, err := tx.Exec(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "tx.Exec")
	}

	return res, nil
}

func (d *Dao) txExecSql(builders []sq.Sqlizer) ([]sql.Result, error) {
	var resList []sql.Result

	tx, err := d.db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "db.Begin")
	}
	for _, builder := range builders {
		res, err := d.txExec(tx, builder)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		resList = append(resList, res)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, errors.Wrap(err, "tx.Commit")
	}
	return resList, nil
}

func (d *Dao) commonExists(tableName string, condition interface{}, args ...interface{}) (bool, error) {
	var exists bool

	builder := sq.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(tableName).
		Where(condition, args...).
		Suffix(")")
	err := d.getSql(builder, &exists)
	if err != nil {
		return exists, errors.Wrap(err, "d.getSql")
	}

	return exists, nil
}

func (d *Dao) commonCount(tableName string, condition interface{}, args ...interface{}) (int, error) {
	var count int

	builder := sq.Select("count(*)").From(tableName).Where(condition, args...)
	err := d.getSql(builder, &count)
	if err != nil {
		return count, errors.Wrap(err, "d.getSql")
	}

	return count, nil
}

func (d *Dao) commonCreate(tableName string, columns []string, values []interface{}) (int64, error) {
	var lastId int64

	builder := sq.Insert(tableName).Columns(columns...).Values(values...)
	res, err := d.execSql(builder)
	if err != nil {
		return 0, errors.Wrap(err, "d.execSql")
	}
	lastId, err = res.LastInsertId()
	if err != nil {
		return lastId, errors.Wrap(err, "LastInsertId")
	}

	return lastId, nil
}

func (d *Dao) commonUpdate(tableName string, setMap map[string]interface{}, condition interface{}, args ...interface{}) (int64, error) {
	var affectRows int64

	builder := sq.Update(tableName).Where(condition, args...)
	for k, v := range setMap {
		builder = builder.Set(k, v)
	}
	res, err := d.execSql(builder)
	if err != nil {
		return 0, errors.Wrap(err, "d.execSql")
	}
	affectRows, err = res.RowsAffected()
	if err != nil {
		return affectRows, errors.Wrap(err, "RowsAffected")
	}

	return affectRows, nil
}

func (d *Dao) commonDelete(tableName string, condition interface{}, args ...interface{}) (int64, error) {
	var affectRows int64

	builder := sq.Delete(tableName).Where(condition, args...)
	res, err := d.execSql(builder)
	if err != nil {
		return 0, errors.Wrap(err, "d.execSql")
	}
	affectRows, err = res.RowsAffected()
	if err != nil {
		return affectRows, errors.Wrap(err, "RowsAffected")
	}

	return affectRows, nil
}

// redis(cache)
func (d *Dao) setCache(key string, value interface{}, ttl time.Duration) error {
	return d.cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   key,
		Value: value,
		TTL:   ttl,
	})
}

func (d *Dao) getCache(key string, dest interface{}) error {
	return d.cache.Get(context.Background(), key, dest)
}

func (d *Dao) deleteCache(key string) error {
	return d.cache.Delete(context.Background(), key)
}
