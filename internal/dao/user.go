package dao

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/cache/v8"
	"github.com/pkg/errors"
	"time"
)

var UserTableName = "user"

func (d *Dao) ExistsUserByUsername(username string) (bool, error) {
	return d.commonExists(UserTableName, sq.Eq{"username": username})
}

func (d *Dao) getUser(condition interface{}, args ...interface{}) (*model.User, error) {
	one := model.User{}

	builder := sq.
		Select("user.*, role.name AS role_name").
		From(UserTableName).
		Join("role ON user.role_id=role.id").
		Where(condition, args...)
	err := d.getSQL(builder, &one)
	if err != nil {
		return nil, errors.Wrap(err, "getSQL")
	}

	return &one, nil
}

func (d *Dao) GetUserById(id int64) (*model.User, error) {
	return d.getUser(sq.Eq{"user.id": id})
}

func (d *Dao) GetUserByUsername(username string) (*model.User, error) {
	return d.getUser(sq.Eq{"user.username": username})
}

func (d *Dao) ListUser(nickname string, status *uint, roleId int64, page, pageSize int) ([]*model.User, error) {
	many := make([]*model.User, 0)

	builder := sq.
		Select("user.*, role.name AS role_name").
		From(UserTableName).
		Join("role ON user.role_id=role.id")
	condition := sq.And{}
	if nickname != "" {
		condition = append(condition, sq.Like{"user.nickname": "%" + nickname + "%"})
	}
	if status != nil {
		condition = append(condition, sq.Eq{"user.status": *status})
	}
	if roleId > 0 {
		condition = append(condition, sq.Eq{"user.role_id": roleId})
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

func (d *Dao) CountUser(nickname string, status *uint, roleId int64) (int, error) {
	//builder := sq.Select("count(*)").From("user")
	//if nickname != "" {
	//	builder = builder.Where(sq.Like{"user.nickname": "%" + nickname + "%"})
	//}
	//if status != nil {
	//	builder = builder.Where(sq.Eq{"user.status": *status})
	//}
	//if roleId > 0 {
	//	builder = builder.Where(sq.Eq{"user.role_id": roleId})
	//}
	//sql, args, err := builder.ToSql()
	//if err != nil {
	//	return 0, errors.Wrap(err, "sql builder")
	//}
	//global.Logger.WithFields(logrus.Fields{
	//	"sql":  sql,
	//	"args": args,
	//}).Debug("sql builder")
	//var count int
	//err = d.db.Get(&count, sql, args...)
	//if err != nil {
	//	return 0, errors.Wrap(err, "sql exec")
	//}
	//return count, nil
	condition := sq.And{}
	if nickname != "" {
		condition = append(condition, sq.Like{"user.nickname": "%" + nickname + "%"})
	}
	if status != nil {
		condition = append(condition, sq.Eq{"user.status": *status})
	}
	if roleId > 0 {
		condition = append(condition, sq.Eq{"user.role_id": roleId})
	}
	return d.commonCount(UserTableName, condition)
}

func (d *Dao) CreateUser(username, passwordHashed, nickname string, status uint, roleId int64) (int64, error) {
	//builder := sq.
	//	Insert("user").
	//	Columns("username", "password_hashed", "nickname", "status", "role_id").
	//	Values(username, passwordHashed, nickname, status, roleId)
	//sql, args, err := builder.ToSql()
	//if err != nil {
	//	return 0, errors.Wrap(err, "sql builder")
	//}
	//global.Logger.WithFields(logrus.Fields{
	//	"sql":  sql,
	//	"args": args,
	//}).Debug("sql builder")
	//res, err := d.db.Exec(sql, args...)
	//if err != nil {
	//	return 0, errors.Wrap(err, "sql exec")
	//}
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return 0, errors.Wrap(err, "sql exec")
	//}
	//return id, nil
	columns := []string{"username", "password_hashed", "nickname", "status", "role_id"}
	values := []interface{}{username, passwordHashed, nickname, status, roleId}
	return d.commonCreate(UserTableName, columns, values)
}

func (d *Dao) UpdateUser(id int64, nickname string, status uint, roleId int64) (int64, error) {
	//builder := sq.Update("user").
	//	Where(sq.Eq{"id": id}).
	//	Set("nickname", nickname).
	//	Set("status", status).
	//	Set("role_id", roleId)
	//
	//sql, args, err := builder.ToSql()
	//if err != nil {
	//	return errors.Wrap(err, "sql builder")
	//}
	//global.Logger.WithFields(logrus.Fields{
	//	"sql":  sql,
	//	"args": args,
	//}).Debug("sql builder")
	//_, err = d.db.Exec(sql, args...)
	//if err != nil {
	//	return errors.Wrap(err, "sql exec")
	//}
	//return nil
	setMap := map[string]interface{}{
		"nickname": nickname,
		"status":   status,
		"role_id":  roleId,
	}
	return d.commonUpdate(UserTableName, setMap, sq.Eq{"id": id})
}

func (d *Dao) DeleteUser(id int64) (int64, error) {
	//builder := sq.Delete("user").Where(sq.Eq{"id": id})
	//sql, args, err := builder.ToSql()
	//if err != nil {
	//	return errors.Wrap(err, "sql builder")
	//}
	//global.Logger.WithFields(logrus.Fields{
	//	"sql":  sql,
	//	"args": args,
	//}).Debug("sql builder")
	//_, err = d.db.Exec(sql, args...)
	//if err != nil {
	//	return errors.Wrap(err, "sql exec")
	//}
	//return nil
	return d.commonDelete(UserTableName, sq.Eq{"id": id})
}

func (d *Dao) GetUserInCache(id int64) (*model.User, error) {
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
