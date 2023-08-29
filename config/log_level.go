package config

type LogLevel string

var (
	LogLevelError LogLevel = "ERROR"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelFatal LogLevel = "FATAL"
)
