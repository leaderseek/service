package config

import (
	"context"

	"github.com/friendsofgo/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sethvargo/go-envconfig"
	"go.temporal.io/sdk/worker"
)

// TODO
type WorkerOptionsConfig struct{}

// TODO
func (cfg *WorkerOptionsConfig) ToTemporal() *worker.Options {
	return new(worker.Options)
}

type WorkerConfig struct {
	Logger                *LoggerConfig         `env:",prefix=LOGGER_"`
	TemporalClient        *TemporalClientConfig `env:",prefix=TEMPORAL_CLIENT_"`
	TemporalWorkerOptions *WorkerOptionsConfig  `env:",prefix=TEMPORAL_WORKER_OPTIONS_"`
}

func (config *WorkerConfig) Validate() error {
	if err := validation.ValidateStruct(config,
		validation.Field(&config.Logger, validation.Required),
		validation.Field(&config.TemporalClient, validation.Required),
		validation.Field(&config.TemporalWorkerOptions, validation.Required)); err != nil {
		return errors.Wrap(err, "failed to find config fields")
	}

	return nil
}

func NewWorkerConfigFromEnv() (*WorkerConfig, error) {
	cfg := new(WorkerConfig)
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		return nil, errors.Wrap(err, "failed to process config")
	}

	return cfg, nil
}
