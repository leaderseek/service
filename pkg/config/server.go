package config

type ServerConfig struct {
	DBConnection string `env:"DB_CONNECTION,required"`
}
