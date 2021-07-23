package dao

import (
	"common_service/internal/model"
	"common_service/pkg/app"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (d *Dao) ExistsRoleById(id int64) (bool, error) {
	return d.commonExists(RoleTableName, sq.Eq{"id": id})
}

func (d *Dao) ExistsRoleByName(name string) (bool, error) {
	return d.commonExists(RoleTableName, sq.Eq{"name": name})
}

func (d *Dao) getRole(condition interface{}, args ...interface{}) (*model.Role, error) {
	one := model.Role{}

	builder := sq.Select("*").From(RoleTableName).Where(condition, args...)
	err := d.getSql(builder, &one)
	if err != nil {
		return nil, errors.Wrap(err, "getSQL")
	}

	return &one, nil
}

func (d *Dao) GetRoleById(id int64) (*model.Role, error) {
	return d.getRole(sq.Eq{"id": id})
}

func (d *Dao) ListRole(name string, page, pageSize int) ([]*model.Role, error) {
	many := make([]*model.Role, 0)

	builder := sq.Select("*").From(RoleTableName)
	if name != "" {
		builder = builder.Where(sq.Like{"name": "%" + name + "%"})
	}
	pageOffSet := app.GetPageOffset(page, pageSize)
	builder = builder.Offset(uint64(pageOffSet)).Limit(uint64(pageSize))

	err := d.selectSql(builder, &many)
	if err != nil {
		return nil, errors.Wrap(err, "selectSql")
	}

	return many, nil
}

func (d *Dao) CountRole(name string) (int, error) {
	condition := sq.And{}
	if name != "" {
		condition = append(condition, sq.Like{"name": "%" + name + "%"})
	}
	return d.commonCount(RoleTableName, condition)
}

func (d *Dao) CreateRole(name, description string) (int64, error) {
	columns := []string{"name", "description"}
	values := []interface{}{name, description}
	return d.commonCreate(RoleTableName, columns, values)
}

func (d *Dao) UpdateRole(id int64, description string) (int64, error) {
	setMap := map[string]interface{}{
		"description": description,
	}
	return d.commonUpdate(RoleTableName, setMap, sq.Eq{"id": id})
}

func (d *Dao) DeleteRole(id int64) (int64, error) {
	return d.commonDelete(RoleTableName, sq.Eq{"id": id})
}

func (d *Dao) DeleteRoleWithUser(roleId int64, userIdList []int64) ([]sql.Result, error) {
	builders := []sq.Sqlizer{
		sq.Delete(RoleTableName).Where(sq.Eq{"id": roleId}),
		sq.Delete(UserTableName).Where(sq.Eq{"id": userIdList}),
	}
	return d.txExecSql(builders)
}
