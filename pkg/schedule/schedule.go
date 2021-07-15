package schedule

import (
	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
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
		errs = append(errs, errors.Wrap(err, "Do task.MyTask"))
	}

	_, err = TaskScheduler.Every(5).Seconds().Do(MyTaskWithParams, 1, "hello")
	if err != nil {
		errs = append(errs, errors.Wrap(err, "Do task.MyTaskWithParams"))
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}
