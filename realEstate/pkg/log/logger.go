package log

//logger содержжит структуру и функционал логирования программой

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logrusLog *logrus.Logger

func GetLogger() {
	logrusLog = logrus.New()
	logrusLog.Formatter = new(logrus.JSONFormatter)
	logrusLog.Level = logrus.DebugLevel
	//logrusLog.Out = os.Stdout
	file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err == nil {
		logrusLog.Out = file
	} else {
		logrusLog.Info("Failed to log to file, using default stderr")
	}

	wrt := io.MultiWriter(os.Stdout, file)
	logrusLog.SetOutput(wrt)
	logrusLog.Info("Logger get started")
}


func WithFields(fields logrus.Fields) {
	logrus.WithFields(fields)
}
func Debug(...interface{}) {
	logrus.Debug()
}

func Panic(args ...interface{}) {
	logrus.Panic()
}

func Error(args ...interface{}) {
	logrus.Error()
}

func Trace(args ...interface{}) {
	logrus.Trace()
}

func Info(args ...interface{}) {
	logrus.Info()
}
func Fatal(args ...interface{}) {
	logrus.Fatal()
}

func Warning(args ...interface{}) {
	logrus.Warning()
}


