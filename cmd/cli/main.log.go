package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// logger
	//logger := zap.NewExample()
	//logger.Info("hello world")
	//
	//// Development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("hello world")
	//
	//// Production
	//logger, _ = zap.NewProduction()
	//logger.Info("hello world")

	// lvl 3.
	encoderConfig := GetEncoderLog()
	sync := GetWriterSync()
	core := zapcore.NewCore(encoderConfig, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("hello world", zap.Int("num", 123))
	logger.Error("hello world", zap.Int("num", 123))
}

func GetEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func GetWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
