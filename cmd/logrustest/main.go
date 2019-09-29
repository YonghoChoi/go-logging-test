package main

import (
	"github.com/YonghoChoi/go-logging-test/cmd/logrustest/logger"
)

func main() {
	logger.Debug("Test", "test")
	logger.Debug("Test", "test")
	logger.Debug("Test", "test")
}
