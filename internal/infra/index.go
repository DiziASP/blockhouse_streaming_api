package infra

import (
	"blockhouse_streaming_api/internal/infra/kafka"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	kafka.NewMessageHandler,
	kafka.NewStreamHandler,
)
