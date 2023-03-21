package logger

import (
	"bytes"
	"log"
	"testing"
	"time"
)

func TestLoggingFunctions(t *testing.T) {
	tests := []struct {
		name           string
		logLevel       LogLevel
		message        string
		expectedOutput string
	}{
		{"Debug message", LogLevelDebug, "This is a debug message", time.Now().Format("2006/01/02 15:04:05") + " [DEBUG] This is a debug message\n"},
		{"Info message", LogLevelInfo, "This is an informational message", time.Now().Format("2006/01/02 15:04:05") + " [INFO] This is an informational message\n"},
		{"Warn message", LogLevelWarn, "This is a warning message", time.Now().Format("2006/01/02 15:04:05") + " [WARN] This is a warning message\n"},
		{"Error message", LogLevelError, "This is an error message", time.Now().Format("2006/01/02 15:04:05") + " [ERROR] This is an error message\n"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			myLogger := New(&buf, "", log.Ldate|log.Ltime)
			logLevel = test.logLevel

			switch test.logLevel {
			case LogLevelDebug:
				LogDebug(myLogger, test.message)
			case LogLevelInfo:
				LogInfo(myLogger, test.message)
			case LogLevelWarn:
				LogWarn(myLogger, test.message)
			case LogLevelError:
				LogError(myLogger, test.message)
			}

			logOutput := buf.String()

			if logOutput != test.expectedOutput {
				t.Errorf("Unexpected log output for '%s': \nGot '%s', \nExpected '%s'", test.name, logOutput, test.expectedOutput)
			}
		})
	}
}
