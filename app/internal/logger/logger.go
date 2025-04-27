package logger

import (
	"io"
	"os"
	"sync"

	"github.com/martintama/domain-checker/internal/types"

	"github.com/sirupsen/logrus"
)

var (
	// defaultLogger is the global default logger instance
	defaultLogger *AppLogger
	once          sync.Once
)

// AppLogger is our abstraction over the underlying logger implementation
type AppLogger struct {
	entry *logrus.Entry
}

// Config holds logger configuration
type Config struct {
	// LogLevel sets the logging level (debug, info, warn, error)
	LogLevel LogLevel
	// Output is where logs are written (defaults to stdout)
	Output io.Writer
	// RunMode specifies how the app is running to alter the logging configuration appropriately
	RunMode types.RunMode
	// DefaultFields stores default fields :)
	DefaultFields Fields
}

// Initialize sets up the logger with configuration
// This should be called early in your application startup
func Initialize(config Config) {
	once.Do(func() {
		// Create the base logrus logger
		base := logrus.New()

		// Set output
		if config.Output == nil {
			base.SetOutput(os.Stdout)
		} else {
			base.SetOutput(config.Output)
		}

		// Set formatter based on environment
		if config.RunMode == types.RunModeLambda {
			// Use JSON formatter for Lambda (better for CloudWatch)
			base.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
			})
		} else {
			// Use text formatter for CLI with minimal output
			base.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "15:04:05",
				DisableTimestamp:       true,
				DisableLevelTruncation: false,
				DisableColors:          false,
				ForceColors:            true,
			})
		}

		// Set log level
		level, err := logrus.ParseLevel(string(config.LogLevel))
		if err != nil {
			level = logrus.InfoLevel // Default to info
		}
		base.SetLevel(level)

		// Create the default logger with default fields
		entry := base.WithFields(config.DefaultFields)
		defaultLogger = &AppLogger{entry: entry}
	})
}

// Fields is a shorthand for logrus.Fields
type Fields = logrus.Fields

// GetLogger returns the configured default logger instance
func GetLogger() *AppLogger {
	if defaultLogger == nil {
		// Initialize with defaults if not done yet
		Initialize(Config{
			LogLevel: "info",
		})
	}
	return defaultLogger
}

// WithField creates a new logger with an additional field
func (l *AppLogger) WithField(key string, value interface{}) *AppLogger {
	return &AppLogger{
		entry: l.entry.WithField(key, value),
	}
}

// WithFields creates a new logger with additional fields
func (l *AppLogger) WithFields(fields Fields) *AppLogger {
	return &AppLogger{
		entry: l.entry.WithFields(fields),
	}
}

// WithError creates a new logger with an error field
func (l *AppLogger) WithError(err error) *AppLogger {
	return &AppLogger{
		entry: l.entry.WithError(err),
	}
}

// Debug logs a debug message
func (l *AppLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugf logs a formatted debug message
func (l *AppLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Info logs an info message
func (l *AppLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Infof logs a formatted info message
func (l *AppLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warn logs a warning message
func (l *AppLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warnf logs a formatted warning message
func (l *AppLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error logs an error message
func (l *AppLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Errorf logs a formatted error message
func (l *AppLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatal logs a fatal message and exits
func (l *AppLogger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Fatalf logs a formatted fatal message and exits
func (l *AppLogger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// WithField returns a new logger with the field added
func WithField(key string, value interface{}) *AppLogger {
	return GetLogger().WithField(key, value)
}

// WithFields returns a new logger with the fields added
func WithFields(fields Fields) *AppLogger {
	return GetLogger().WithFields(fields)
}

// WithError returns a new logger with the error field added
func WithError(err error) *AppLogger {
	return GetLogger().WithError(err)
}

// SetLogLevel changes the current logging level
func SetLogLevel(level LogLevel) {
	if defaultLogger != nil && defaultLogger.entry != nil && defaultLogger.entry.Logger != nil {
		// Convert your LogLevel to the equivalent logrus level
		var logrusLevel logrus.Level
		switch level {
		case LogLevelDebug:
			logrusLevel = logrus.DebugLevel
		case LogLevelInfo:
			logrusLevel = logrus.InfoLevel
		default:
			logrusLevel = logrus.InfoLevel
		}

		defaultLogger.entry.Logger.SetLevel(logrusLevel)
	}
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
