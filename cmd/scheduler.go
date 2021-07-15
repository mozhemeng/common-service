package cmd

import (
	"common_service/pkg/schedule"
	"github.com/spf13/cobra"
	"log"
)

func runScheduler() {
	err := schedule.SetupTaskScheduler()
	if err != nil {
		log.Fatalf("init.SetupTaskScheduler err: %v", err)
	}
	schedule.TaskScheduler.StartBlocking()
}

var schedulerCmd = &cobra.Command{
	Use: "scheduler",
	Aliases: []string{"task", "job"},
	Short: "run background scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		runScheduler()
	},
}