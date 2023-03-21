package logger

import (
	"io"
	"log"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

var logLevel = LogLevelInfo

func New(output io.Writer, prefix string, flag int) *log.Logger {
	return log.New(output, prefix, flag)
}

func LogDebug(logger *log.Logger, message string) {
	if logLevel <= LogLevelDebug {
		logger.Println("[DEBUG]", message)
	}
}

func LogInfo(logger *log.Logger, message string) {
	if logLevel <= LogLevelInfo {
		logger.Println("[INFO]", message)
	}
}

func LogWarn(logger *log.Logger, message string) {
	if logLevel <= LogLevelWarn {
		logger.Println("[WARN]", message)
	}
}

func LogError(logger *log.Logger, message string) {
	if logLevel <= LogLevelError {
		logger.Println("[ERROR]", message)
	}
}
