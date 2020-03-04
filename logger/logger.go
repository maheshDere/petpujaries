package logger

import (
	"log"
	"os"
	"petpujaris/config"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var logFile *os.File

func Setup() {
	appConfig := config.GetAppConfig()
	logFile, err := os.OpenFile(appConfig.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error opening log file, using Stdout", logFile, "Error:", err.Error())
		logFile = os.Stdout
	}

	logger = &logrus.Logger{
		Out:       logFile,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.ErrorLevel,
	}
}

func Close() {
	logFile.Close()
}

func LogInfo(fields logrus.Fields, info string) {
	logger.WithFields(logrus.Fields(fields)).Info(info)
}
func LogError(err error, where, errMsg string) {
	fields := logrus.Fields{"Error": err.Error(), "Where": where}
	logger.WithFields(logrus.Fields(fields)).Error(errMsg)
}
