package cmd

import (
	"common_service/global"
	"github.com/spf13/cobra"
	"log"
)

func setup() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.SetupSetting err: %v", err)
	}
	err = global.SetupLogger()
	if err != nil {
		log.Fatalf("init.SetupLogger err: %v", err)
	}
	err = global.SetupDB()
	if err != nil {
		log.Fatalf("init.SetupDB err: %v", err)
	}
	err = global.SetupEnforcer()
	if err != nil {
		log.Fatalf("init.SetupEnforcer err: %v", err)
	}
	err = global.SetupRedis()
	if err != nil {
		log.Fatalf("init.SetupRedis err: %v", err)
	}
	err = global.SetupUniTrans()
	if err != nil {
		log.Fatalf("init.SetupUniTrans err: %v", err)
	}
}

func init() {
	setup()

	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(schedulerCmd)
	RootCmd.AddCommand(prepareCmd)
}

var RootCmd = &cobra.Command{
	Use:   "common_service",
	Short: "Common Service",
}
