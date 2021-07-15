package global

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func SetupDB() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local&multiStatements=%t",
		DatabaseSetting.Username,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.DBName,
		DatabaseSetting.Charset,
		DatabaseSetting.ParseTime,
		DatabaseSetting.MultiStatements)

	var err error
	DB, err = sqlx.Connect(DatabaseSetting.DBType, dsn)
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	DB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)

	return nil
}
