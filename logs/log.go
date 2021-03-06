package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var S *zap.SugaredLogger

func init() {
	GetLog("warn")
}

func GetLog(level string) {
	var atomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	switch strings.ToLower(level) {
	case "info":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		atomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}
	cfg := zap.Config{
		Encoding:         "console",
		Level:            atomicLevel,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := cfg.Build()
	S = logger.Sugar()
}
