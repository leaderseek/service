package app

import (
	"fmt"
	"os"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	server_config "github.com/leaderseek/service/pkg/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	HTTP           *HTTPConfig           `json:"http" yaml:"http"`
	ZapLogger      *zap.Config           `json:"zapLogger" yaml:"zapLogger"`
	TemporalClient *TemporalClientConfig `json:"temporalClient" yaml:"temporalClient"`
	ServerConfig   *server_config.Config `json:"server" yaml:"server"`
}

func NewConfigFromEnv() (*Config, error) {
	var config *Config

	httpConfig, err := NewHTTPConfigFromEnv()
	if err != nil {
		return nil, err
	}

	config.HTTP = httpConfig

	config.ZapLogger = NewZapLoggerConfigFromEnv()
	config.TemporalClient = NewTemporalClientConfigFromEnv()

	config.ServerConfig = server_config.NewConfigFromEnv()

	return config, nil
}

func NewHTTPConfigFromEnv() (*HTTPConfig, error) {
	httpPort, err := strconv.ParseUint(os.Getenv("HTTP_PORT"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &HTTPConfig{
		Port: httpPort,
	}, nil
}

func NewZapLoggerConfigFromEnv() *zap.Config {
	atomicLevel := zap.NewAtomicLevel()

	if os.Getenv("LOGGER_LEVEL") == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    os.Getenv("LOGGER_ENCODER_CONFIG_MESSAGE_KEY"),
		LevelKey:      os.Getenv("LOGGER_ENCODER_CONFIG_LEVEL_KEY"),
		TimeKey:       os.Getenv("LOGGER_ENCODER_CONFIG_TIME_KEY"),
		NameKey:       os.Getenv("LOGGER_ENCODER_CONFIG_NAME_KEY"),
		CallerKey:     os.Getenv("LOGGER_ENCODER_CONFIG_CALLER_KEY"),
		FunctionKey:   os.Getenv("LOGGER_ENCODER_CONFIG_FUNCTION_KEY"),
		StacktraceKey: os.Getenv("LOGGER_ENCODER_CONFIG_STACKTRACE_KEY"),
	}

	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return &zap.Config{
		Level:            atomicLevel,
		Development:      os.Getenv("LOGGER_DEVELOPMENT") == "true",
		Encoding:         os.Getenv("LOGGER_ENCODING"),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{os.Getenv("LOGGER_STDOUT_PATH")},
		ErrorOutputPaths: []string{os.Getenv("LOGGER_STDERR_PATH")},
	}
}

func NewTemporalClientConfigFromEnv() *TemporalClientConfig {
	return &TemporalClientConfig{
		HostPort: os.Getenv("TEMPORAL_HOSTPORT"),
	}
}

func (config *Config) Validate() error {
	if err := validation.ValidateStruct(config,
		validation.Field(&config.HTTP, validation.Required),
		validation.Field(&config.ZapLogger, validation.Required)); err != nil {
		return fmt.Errorf("failed to find config fields, error = %v", err)
	}

	if err := config.HTTP.Validate(); err != nil {
		return fmt.Errorf("failed to validate http, error = %v", err)
	}

	return nil
}

type HTTPConfig struct {
	Port uint64 `json:"port" yaml:"port"`
}

func (httpConfig *HTTPConfig) Validate() error {
	port := httpConfig.Port

	if port != 80 && port != 443 && port < 1025 {
		return fmt.Errorf("config error: invalid http port")
	}

	return nil
}

func (httpConfig *HTTPConfig) Address() string {
	return fmt.Sprintf(":%d", httpConfig.Port)
}

type TemporalClientConfig struct {
	HostPort string `json:"hostPort" yaml:"hostPort"`
}
