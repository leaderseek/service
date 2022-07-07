package config

import (
	"context"

	"github.com/friendsofgo/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	HTTP           *HTTPConfig           `env:",prefix=HTTP_"`
	Logger         *LoggerConfig         `env:",prefix=LOGGER_"`
	TemporalClient *TemporalClientConfig `env:",prefix=TEMPORAL_CLIENT_"`
	Server         *ServerConfig         `env:",prefix=SERVER_"`
}

func (config *AppConfig) Validate() error {
	if err := validation.ValidateStruct(config,
		validation.Field(&config.HTTP, validation.Required),
		validation.Field(&config.Logger, validation.Required),
		validation.Field(&config.TemporalClient, validation.Required),
		validation.Field(&config.Server, validation.Required)); err != nil {
		return errors.Wrap(err, "failed to find config fields")
	}

	if err := config.HTTP.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate http")
	}

	return nil
}

func NewAppConfigFromEnv() (*AppConfig, error) {
	cfg := new(AppConfig)
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		return nil, errors.Wrap(err, "failed to process config")
	}

	return cfg, nil
}
