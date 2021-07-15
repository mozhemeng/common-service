package dao

import (
	"common_service/global"
	"common_service/internal/model"
	"common_service/pkg/app"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/cache/v8"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

func (d *Dao) ExistsUserByUsername(username string) (bool, error) {
	var exists bool
	builder := sq.
		Select("1").
		Prefix("SELECT EXISTS (").
		From("user").
		Where(sq.Eq{"username": username}).
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

func (d *Dao) GetUserById(id uint64) (*model.User, error) {
	u := model.User{}
	builder := sq.
		Select("user.*, role.name AS role_name").
		From("user").
		Join("role ON user.role_id=role.id").
		Where(sq.Eq{"user.id": id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	err = d.db.Get(&u, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "sql exec")
	}
	return &u, nil
}

func (d *Dao) GetUserByUsername(username string) (*model.User, error) {
	u := model.User{}
	builder := sq.
		Select("user.*, role.name AS role_name").
		From("user").
		Join("role ON user.role_id=role.id").
		Where(sq.Eq{"user.username": username})
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	err = d.db.Get(&u, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "sql exec")
	}
	return &u, nil
}

func (d *Dao) CreateUser(username, passwordHashed, nickname string, status uint, roleId uint64) (uint64, error) {
	builder := sq.
		Insert("user").
		Columns("username", "password_hashed", "nickname", "status", "role_id").
		Values(username, passwordHashed, nickname, status, roleId)
	sql, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	res, err := d.db.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sql exec")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "sql exec")
	}
	return uint64(id), nil
}

func (d *Dao) ListUser(nickname string, status *uint, roleId uint64, page, pageSize int) ([]*model.User, error) {
	builder := sq.
		Select("user.*, role.name AS role_name").
		From("user").
		Join("role ON user.role_id=role.id")
	if nickname != "" {
		builder = builder.Where(sq.Like{"user.nickname": "%" + nickname + "%"})
	}
	if status != nil {
		builder = builder.Where(sq.Eq{"user.status": *status})
	}
	if roleId > 0 {
		builder = builder.Where(sq.Eq{"user.role_id": roleId})
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
	var users []*model.User
	err = d.db.Select(&users, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "sql exec")
	}
	if users == nil {
		users = make([]*model.User, 0)
	}
	return users, nil
}

func (d *Dao) CountUser(nickname string, status *uint, roleId uint64) (int, error) {
	builder := sq.Select("count(*)").From("user")
	if nickname != "" {
		builder = builder.Where(sq.Like{"user.nickname": "%" + nickname + "%"})
	}
	if status != nil {
		builder = builder.Where(sq.Eq{"user.status": *status})
	}
	if roleId > 0 {
		builder = builder.Where(sq.Eq{"user.role_id": roleId})
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

func (d *Dao) UpdateUser(id uint64, nickname string, status uint, roleId uint64) error {
	builder := sq.Update("user").
		Where(sq.Eq{"id": id}).
		Set("nickname", nickname).
		Set("status", status).
		Set("role_id", roleId)

	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	_, err = d.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}

func (d *Dao) DeleteUser(id uint64) error {
	builder := sq.Delete("user").Where(sq.Eq{"id": id})
	sql, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "sql builder")
	}
	global.Logger.WithFields(logrus.Fields{
		"sql":  sql,
		"args": args,
	}).Debug("sql builder")
	_, err = d.db.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "sql exec")
	}
	return nil
}

func (d *Dao) GetUserInCache(id uint64) (*model.User, error) {
	u := model.User{}
	key := fmt.Sprintf("user:%d", id)
	err := d.cache.Get(context.Background(), key, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *Dao) SetUserInCache(user *model.User, ttl time.Duration) error {
	key := fmt.Sprintf("user:%d", user.ID)
	err := d.cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   key,
		Value: user,
		TTL:   ttl,
	})
	if err != nil {
		return err
	}
	return nil
}
