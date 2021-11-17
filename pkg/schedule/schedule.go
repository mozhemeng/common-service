package schedule

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

var TaskScheduler *gocron.Scheduler

func SetupTaskScheduler() error {
	TaskScheduler = gocron.NewScheduler(time.UTC)

	var err error
	errs := make([]error, 0)
	// 添加任务
	_, err = TaskScheduler.Every(3).Seconds().Do(MyTask)
	if err != nil {
		errs = append(errs, fmt.Errorf("task.MyTask: %w", err))
	}

	_, err = TaskScheduler.Every(5).Seconds().Do(MyTaskWithParams, 1, "hello")
	if err != nil {
		errs = append(errs, fmt.Errorf("task.MyTaskWithParams: %w", err))
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
