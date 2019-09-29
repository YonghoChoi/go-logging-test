package main

import (
	"github.com/sirupsen/logrus"
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

	logger.Debugf("test")
}
