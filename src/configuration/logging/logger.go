package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	log_output = os.Getenv("LOG_OUTPUT")

	log_level = os.Getenv("LOG_LEVELS")
)

func init() {
	log_config := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			MessageKey:   "message",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = log_config.Build()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)

	log.Sync()
}

func Error(message string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("err", err))

	log.Info(message, fields...)

	log.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(log_output)

	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(log_level) {
	case "ERROR":
		return zapcore.ErrorLevel
	case "DEBUG":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
