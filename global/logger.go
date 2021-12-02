package global

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Logger zerolog.Logger

func SetupLogger() error {
	logFileName := AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt
	lumber := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100,
		MaxBackups: 5,
	}
	multiOut := io.MultiWriter(os.Stdout, lumber)

	Logger = zerolog.New(multiOut).With().Timestamp().Logger()

	if ServerSetting.RunMode == "debug" {
		Logger = Logger.Level(zerolog.DebugLevel)
	} else {
		Logger = Logger.Level(zerolog.InfoLevel)
	}

	return nil
}
