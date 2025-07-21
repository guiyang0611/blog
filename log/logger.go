package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout", "logfile.log"}
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}
