package logger

import (
	"blockhouse_streaming_api/config"
	"blockhouse_streaming_api/pkg/logger"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewLoggerApplication,
)

// NewLoggerApplication Constructor
func NewLoggerApplication(cfg *config.Configuration) logger.Logger {
	return logger.NewApiLogger(cfg)
}
