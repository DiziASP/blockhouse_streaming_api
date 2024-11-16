package logger

import (
	"blockhouse_streaming_api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Logger methods interface
type Logger interface {
	InitLogger()
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	WithFiled(field zapcore.Field) *zap.Logger
}

// Logger Instance
type apiLogger struct {
	cfg         *config.Configuration
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

// NewApiLogger
// Logger constructor
func NewApiLogger(cfg ...*config.Configuration) Logger {
	apilg := &apiLogger{}
	if len(cfg) == 0 {
		apilg.DefaultInit()
	} else {
		apilg.cfg = cfg[0]
		apilg.InitLogger()
	}
	return apilg
}
