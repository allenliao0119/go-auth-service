package config

type Config struct {
	Server     ServerConfig
	Postgres   PostgresConfig
	Redis      RedisConfig
	Logger     LoggerConfig
	Token      TokenConfig
	APIDocAuth APIAuthConfig
}

type ServerConfig struct{}

type PostgresConfig struct{}

type RedisConfig struct{}

type LogFormat string
type LogLevel string

const (
	// Format
	LogFormatJSON LogFormat = "json"
	LogFormatText LogFormat = "text"

	// Level
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type LoggerConfig struct {
	Format LogFormat
	Level  LogLevel
}

type TokenConfig struct{}

type APIAuthConfig struct{}
