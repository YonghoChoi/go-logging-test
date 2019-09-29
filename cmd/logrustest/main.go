package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
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

	logger.SetOutput(l)
	logger.Debugf("test")
	logger.Debugf("test")
	logger.Debugf("test")
	if err := l.Rotate(); err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Debugf("test")
}
