package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	Log, _ = zap.NewDevelopment()
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}