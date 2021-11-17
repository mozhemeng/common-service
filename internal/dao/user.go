package dao

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

func (d *Dao) ExistsUserByUsername(username string) (bool, error) {
	return d.commonExists(UserTableName, sq.Eq{"username": username})
}

func (d *Dao) ExistsUserByRoleId(roleId int64) (bool, error) {
	return d.commonExists(UserTableName, sq.Eq{"role_id": roleId})
}

func (d *Dao) getUser(condition interface{}, args ...interface{}) (*model.User, error) {
	one := model.User{}

	builder := sq.
		Select("user.*, role.name AS role_name").
		From(UserTableName).
		Join("role ON user.role_id=role.id").
		Where(condition, args...)
	err := d.getSql(builder, &one)
	if err != nil {
		return nil, fmt.Errorf("dao.getSql: %w", err)
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

	err := d.selectSql(builder, &many)
	if err != nil {
		return nil, fmt.Errorf("dao.selectSql: %w", err)
	}

	return many, nil
}

func (d *Dao) ListUserId(nickname string, status *uint, roleId int64) ([]int64, error) {
	many := make([]int64, 0)

	builder := sq.
		Select("id").
		From(UserTableName)
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

	err := d.selectSql(builder, &many)
	if err != nil {
		return nil, fmt.Errorf("dao.selectSql: %w", err)
	}

	return many, nil
}

func (d *Dao) CountUser(nickname string, status *uint, roleId int64) (int, error) {
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
	columns := []string{"username", "password_hashed", "nickname", "status", "role_id"}
	values := []interface{}{username, passwordHashed, nickname, status, roleId}
	return d.commonCreate(UserTableName, columns, values)
}

func (d *Dao) UpdateUser(id int64, nickname string, status uint, roleId int64) (int64, error) {
	setMap := map[string]interface{}{
		"nickname": nickname,
		"status":   status,
		"role_id":  roleId,
	}
	return d.commonUpdate(UserTableName, setMap, sq.Eq{"id": id})
}

func (d *Dao) DeleteUser(id int64) (int64, error) {
	return d.commonDelete(UserTableName, sq.Eq{"id": id})
}

func (d *Dao) GetUserInCache(id int64) (*model.User, error) {
	u := model.User{}
	key := fmt.Sprintf("user:%d", id)
	err := d.getCache(key, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *Dao) SetUserInCache(user *model.User, ttl time.Duration) error {
	key := fmt.Sprintf("user:%d", user.ID)
	return d.setCache(key, user, ttl)
}

func (d *Dao) DeleteUserInCache(id int64) error {
	key := fmt.Sprintf("user:%d", id)
	return d.deleteCache(key)
}
