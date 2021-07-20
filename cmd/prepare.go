package cmd

import (
	"common_service/global"
	"common_service/internal/dao"
	"common_service/pkg/app"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
)

func runInitTables() {
	_, err := sqlx.LoadFile(global.DB, global.AppSetting.InitTablesSqlPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "sqlx.LoadFile"))
	}
	log.Println("init db success")
}

func runCreateRootUser() {
	// 新建用户只用到sql db
	d := dao.New(global.DB, nil, nil)
	if global.AppSetting.RootPassword == "" {
		log.Fatal("must set root password first")
	}
	passwordHashed, err := app.HashPassword(global.AppSetting.RootPassword)
	if err != nil {
		log.Fatal(errors.Wrap(err, "app.HashPassword"))
	}

	_, err = d.CreateUser(global.AppSetting.RootUsername, passwordHashed, global.AppSetting.RootUsername, 1, 1)
	if err != nil {
		log.Fatal(errors.Wrap(err, "dao.CreateUser"))
	}
}

var prepareCmd = &cobra.Command{
	Use:     "prepare",
	Aliases: []string{"init", "setup"},
	Short:   "prepare for server",
	Run: func(cmd *cobra.Command, args []string) {
		runInitTables()
		runCreateRootUser()
	},
}
