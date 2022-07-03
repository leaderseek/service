package config

import "os"

type Config struct {
	DBConnectionString string `json:"dbConnectionString" yaml:"dbConnectionString"`
}

func NewConfigFromEnv() *Config {
	return &Config{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}
}
