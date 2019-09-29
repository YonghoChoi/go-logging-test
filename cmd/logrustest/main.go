package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
)

func main() {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		panic(err)
	}

	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Formatter: &logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	l := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    100, // megabytes
		MaxBackups: 10,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
	}

	logger.SetOutput(io.MultiWriter(l, os.Stdout))
	logger.Debugf("test")
	logger.Debugf("test")
	logger.Debugf("test")
	if err := l.Rotate(); err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Debugf("test")

	entry := getLoggerWithRuntimeContext(logger, 3)
	entry.Debug(fmt.Sprint("test"))
}

func getLoggerWithRuntimeContext(logger *logrus.Logger, skip int) *logrus.Entry {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		funname := runtime.FuncForPC(pc).Name()
		logFields := logrus.Fields{}
		logFields["file"] = file
		logFields["func"] = funname
		logFields["line"] = line

		return logger.WithFields(logFields)
	}
	return nil
}
