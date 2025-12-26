package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	// Encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"

	// Encoders
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// File writer
	logFile, err := os.OpenFile(
		"app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	// Write syncers
	fileSyncer := zapcore.AddSync(logFile)
	consoleSyncer := zapcore.AddSync(os.Stdout)

	// Combine cores (console + file)
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileSyncer, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, consoleSyncer, zapcore.DebugLevel),
	)

	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	Logger.Info("logger initialized")
}
