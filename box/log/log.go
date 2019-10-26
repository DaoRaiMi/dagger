package log

import (
	"github.com/daoraimi/dagger/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  config.GetString("log.fileName"),
		MaxSize:   config.GetInt("log.maxSize"), // megabytes
		MaxAge:    config.GetInt("log.maxAge"),  // days
		LocalTime: true,
	})

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "Level",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		MessageKey:     "Msg",
		StacktraceKey:  "Stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		config.GetLogLevel("log.level"),
	)

	logger = zap.New(core)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
