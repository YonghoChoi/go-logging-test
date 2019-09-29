package log

type Logger interface {
	Debug(logType string, format string, args ...interface{})
	Info(logType string, format string, args ...interface{})
	Warn(logType string, format string, args ...interface{})
	Error(logType string, format string, args ...interface{})
	Fatal(logType string, format string, args ...interface{})
	Panic(logType string, format string, args ...interface{})
}
