package kafka

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAdmin,
	NewConsumer,
	NewProducer,
)