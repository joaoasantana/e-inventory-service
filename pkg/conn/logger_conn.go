package conn

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func DebugLogger() *zap.Logger {
	encoder := ecszap.NewDefaultEncoderConfig()
	encoderCore := ecszap.NewCore(encoder, os.Stdout, zapcore.DebugLevel)

	return zap.New(encoderCore, zap.AddCaller())
}
