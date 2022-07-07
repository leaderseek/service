package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Debug              bool                 `env:"DEBUG,default=false"`
	Encoding           string               `env:"ENCODING,default=json"`
	Encoder            *LoggerEncoderConfig `env:",prefix=ENCODER_"`
	StandardOutputPath string               `env:"STDOUT_PATH,default=stdout"`
	StandardErrorPath  string               `env:"STDERR_PATH,default=stderr"`
}

type LoggerEncoderConfig struct {
	Message    string `env:"MESSAGE,default=message"`
	Level      string `env:"LEVEL,default=level"`
	Time       string `env:"TIME,default=time"`
	Name       string `env:"NAME,default=name"`
	Caller     string `env:"CALLER,default=caller"`
	Function   string `env:"FUNCTION,default=function"`
	Stacktrace string `env:"STACKTRACE,default=stacktrace"`
}

func (cfg *LoggerConfig) ToZap() *zap.Config {
	atomicLevel := zap.NewAtomicLevel()

	if cfg.Debug {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	encoderCfg := cfg.Encoder.ToZap()

	return &zap.Config{
		Level:            atomicLevel,
		Development:      cfg.Debug,
		Encoding:         cfg.Encoding,
		EncoderConfig:    *encoderCfg,
		OutputPaths:      []string{cfg.StandardOutputPath},
		ErrorOutputPaths: []string{cfg.StandardErrorPath},
	}
}

func (cfg *LoggerEncoderConfig) ToZap() *zapcore.EncoderConfig {
	encoderCfg := &zapcore.EncoderConfig{
		MessageKey:    cfg.Message,
		LevelKey:      cfg.Level,
		TimeKey:       cfg.Time,
		NameKey:       cfg.Name,
		CallerKey:     cfg.Caller,
		FunctionKey:   cfg.Function,
		StacktraceKey: cfg.Stacktrace,
	}

	encoderCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.EncodeDuration = zapcore.NanosDurationEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	return encoderCfg
}
