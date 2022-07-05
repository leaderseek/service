package worker

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/leaderseek/service/pkg/app"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

type Config struct {
	ZapLogger             *zap.Config
	TemporalClient        *app.TemporalClientConfig
	TemporalWorkerOptions *worker.Options
}

func NewConfigFromEnv() (*Config, error) {
	var cfg *Config

	cfg.ZapLogger = app.NewZapLoggerConfigFromEnv()
	cfg.TemporalClient = app.NewTemporalClientConfigFromEnv()

	// TODO
	cfg.TemporalWorkerOptions = new(worker.Options)

	return cfg, nil
}

func (cfg *Config) Validate() error {
	if err := validation.ValidateStruct(cfg,
		validation.Field(&cfg.ZapLogger, validation.Required)); err != nil {
		return fmt.Errorf("failed to find config fields, error = %v", err)
	}

	return nil
}
