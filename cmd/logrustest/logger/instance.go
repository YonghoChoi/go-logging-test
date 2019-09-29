package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
)

var (
	gLogger *logrus.Logger
)

func GetInstance() *logrus.Logger {
	if gLogger == nil {
		logLevel := "debug"
		fileName := "./test.log"
		maxSize := 100   // mb : 파일로그가 설정된 용량을 초과하면 로그 파일 로테이트
		maxBackups := 10 // 로그 로테이트 된 파일의 개수가 설정된 값을 초과하면 가장 오래된 파일 제거
		maxAge := 1      // day : 로그 로테이트 주기

		level, err := logrus.ParseLevel(logLevel)
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
			Filename:   fileName,
			MaxSize:    maxSize, // megabytes
			MaxBackups: maxBackups,
			MaxAge:     maxAge, //days
			Compress:   true,   // 로그 로테이트 시 압축 여부
		}

		logger.SetOutput(io.MultiWriter(l, os.Stdout))
		entry := getLoggerWithRuntimeContext(logger, 3)
		if entry == nil {
			panic("log entry is nil")
		}

		gLogger = entry.Logger
	}

	return gLogger
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

func isLimitLength(size int) bool {
	const MaxLogLength = 1000
	if size > MaxLogLength {
		return true
	}

	return false
}

func Debug(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Debug(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Debug(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}

func Info(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Info(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Info(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}

func Warn(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Warn(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Warn(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}

func Error(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Error(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Error(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}

func Panic(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Panic(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Panic(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}

func Fatal(logType string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if isLimitLength(len(msg)) { // 로그 길이가 너무 길어지는 경우 제한 ( 현재 임의로 1000)
		GetInstance().Fatal(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"log message size over. msg : %s ... \"}", logType, msg[0:100]))
	} else {
		GetInstance().Fatal(fmt.Sprintf("{\"type\": \"%s\", \"msg\": \"%s\"}", logType, msg))
	}
}
