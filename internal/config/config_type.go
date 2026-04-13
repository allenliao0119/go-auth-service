package config

type Config struct {
	Server     ServerConfig
	Postgres   PostgresConfig
	Redis      RedisConfig
	Logger     LoggerConfig
	Token      TokenConfig
	APIDocAuth APIAuthConfig
}

type ServerConfig struct {}

type PostgresConfig struct {}

type RedisConfig struct {}

type LoggerConfig struct {}

type TokenConfig struct {}

type APIAuthConfig struct {}
