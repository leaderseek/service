package config

type TemporalClientConfig struct {
	HostPort string `env:"HOST_PORT,required"`
}
