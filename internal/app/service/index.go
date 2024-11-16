package service

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStreamService,
	NewMessageService,
)
