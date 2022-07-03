package app

import (
	"fmt"

	"go.temporal.io/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapToTemporalLogger(zapLogger *zap.Logger) log.Logger {
	return appLogger{
		zapLogger: zapLogger,
	}
}

type appLogger struct {
	zapLogger *zap.Logger
}

func any(keyvals ...interface{}) (fields []zapcore.Field) {
	for i := range keyvals {
		fields = append(fields, zap.Any(fmt.Sprintf("%d", i), keyvals[i]))
	}
	return
}

func (logger appLogger) Debug(msg string, keyvals ...interface{}) {
	logger.zapLogger.Debug(msg, any(keyvals...)...)
}

func (logger appLogger) Error(msg string, keyvals ...interface{}) {
	logger.zapLogger.Error(msg, any(keyvals...)...)
}

func (logger appLogger) Info(msg string, keyvals ...interface{}) {
	logger.zapLogger.Info(msg, any(keyvals...)...)
}

func (logger appLogger) Warn(msg string, keyvals ...interface{}) {
	logger.zapLogger.Warn(msg, any(keyvals...)...)
}
