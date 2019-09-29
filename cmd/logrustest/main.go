package main

import "github.com/YonghoChoi/go-logging-test/cmd/logrustest/log"

func main() {
	l := log.New("debug", "./test.log", 1, 10, 1)
	l.Debug("Test", "test")
	l.Debug("Test", "test")
	l.Debug("Test", "test")
	l.Debug("Test", "test")
	l.Debug("Test", "test")
}
