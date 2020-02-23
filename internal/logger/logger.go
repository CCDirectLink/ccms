package logger

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	// InfoLevel will be shown
	InfoLevel = 1
	// LogLevel will be shown
	LogLevel = 2
	// WarnLevel will be shown
	WarnLevel = 4
	// CriticalLevel will be shown
	CriticalLevel = 8
	// AllLevel will be shown
	AllLevel = 15
)

var logLevel int = 0

// SetLogLevel sets level
func SetLogLevel(level int) {
	logLevel = level
}

// Info as [INFO]prefix:message
func Info(prefix string, message string) string {
	msg := _log("INFO", prefix, message)

	if logLevel&InfoLevel != 0 {
		color.Cyan(msg)
	}

	return msg
}

// Log as [LOG]prefix:message
func Log(prefix string, message string) string {

	msg := _log("LOG", prefix, message)

	if logLevel&LogLevel != 0 {
		color.White(msg)
	}
	return msg
}

// Warn as [WARN]prefix:message
func Warn(prefix string, message string) string {
	msg := _log("WARN", prefix, message)

	if logLevel&WarnLevel != 0 {
		color.Yellow(msg)
	}
	return msg
}

// Critical as [CRITICAL]prefix:message
func Critical(prefix, message string) string {

	msg := _log("CRITICAL", prefix, message)

	if logLevel&CriticalLevel != 0 {
		color.Red(msg)
	}
	return msg
}

func _log(logType string, prefix string, message string) string {
	return fmt.Sprintf("[%s]%s:%s", logType, prefix, message)
}
