// File: internal/logger/logger.go
package logger

import (
	"io"
	"os"
	"sync"

	"github.com/martintama/domain-checker/internal/types"

	"github.com/sirupsen/logrus"
)

var (
	// Log is the default logger instance
	Log  *logrus.Logger
	once sync.Once
)

// Config holds logger configuration
type Config struct {
	// LogLevel sets the logging level (debug, info, warn, error)
	LogLevel string
	// Output is where logs are written (defaults to stdout)
	Output io.Writer
	// RunMode specifies how the app is running to alter the logging configuration appropriately
	RunMode types.RunMode
}

// Initialize sets up the logger with configuration
// This should be called early in your application startup
func Initialize(config Config) {
	once.Do(func() {
		// Create logger if it doesn't exist
		if Log == nil {
			Log = logrus.New()
		}

		// Set output
		if config.Output == nil {
			Log.SetOutput(os.Stdout)
		} else {
			Log.SetOutput(config.Output)
		}

		// Set formatter based on environment
		if config.RunMode == types.RunModeLambda {
			// Use JSON formatter for Lambda (better for CloudWatch)
			Log.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
			})
		} else {
			// Use text formatter for CLI with minimal output
			Log.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "15:04:05",
				DisableTimestamp:       true,
				DisableLevelTruncation: false,
				DisableColors:          false,
				ForceColors:            true,
			})
		}

		// Set log level
		level, err := logrus.ParseLevel(config.LogLevel)
		if err != nil {
			level = logrus.InfoLevel // Default to info
		}
		Log.SetLevel(level)
	})
}

// Fields is a shorthand for logrus.Fields
type Fields = logrus.Fields

// GetLogger returns the configured logger instance
func GetLogger() *logrus.Logger {
	if Log == nil {
		// Initialize with defaults if not done yet
		Initialize(Config{
			LogLevel: "info",
		})
	}
	return Log
}

// WithField returns a logger entry with a field
func WithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

// WithFields returns a logger entry with fields
func WithFields(fields Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Debugf logs a formatted debug message
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

// Info logs an info message
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof logs a formatted info message
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal logs a fatal message and exits
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf logs a formatted fatal message and exits
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

// Panic logs a panic message and panics
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// Panicf logs a formatted panic message and panics
func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}
