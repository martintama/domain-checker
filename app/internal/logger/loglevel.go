package logger

type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug          = "debug"
)

func ParseLogLevel(level string) LogLevel {
	switch level {
	case "debug":
		return LogLevelDebug
	case "info":
		return LogLevelInfo
	// Add other cases as needed
	default:
		return LogLevelInfo // Default
	}
}
