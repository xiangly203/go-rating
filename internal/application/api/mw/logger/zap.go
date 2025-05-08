package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	Logger *zap.Logger
	once   sync.Once
)

func GetLogInstance() *zap.Logger {
	once.Do(func() {
		Logger, _ = NewLogger()
	})
	return Logger
}

func LogInit() {
	GetLogInstance()
}

func NewLogger() (*zap.Logger, error) {
	stdoutSyncer := zapcore.AddSync(os.Stdout)
	file, err := os.OpenFile("logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	fileSyncer := zapcore.AddSync(file)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Using ISO8601 Time Encoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	multiSyncWriter := zapcore.NewMultiWriteSyncer(stdoutSyncer, fileSyncer)
	core := zapcore.NewCore(encoder, multiSyncWriter, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return logger, nil
}
