package schedule

import "common_service/global"

func MyTask() {
	global.Logger.Println("I am running task.")
}

func MyTaskWithParams(a int, b string) {
	global.Logger.Printf("I am running task. This is int: %d and string: %s", a, b)
}
