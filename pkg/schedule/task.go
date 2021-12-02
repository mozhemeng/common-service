package schedule

import "common_service/global"

func MyTask() {
	global.Logger.Info().
		Msg("I am running task.")
}

func MyTaskWithParams(a int, b string) {
	global.Logger.Info().
		Msgf("I am running task. This is int: %d and string: %s", a, b)
}
