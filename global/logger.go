package global

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Logger *logrus.Logger

func SetupLogger() error {
	logFileName := AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt
	lumber := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100,
		MaxBackups: 5,
	}
	multiOut := io.MultiWriter(os.Stdout, lumber)

	Logger = logrus.New()
	Logger.SetOutput(multiOut)
	if ServerSetting.RunMode == "debug" {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}

	return nil
}
