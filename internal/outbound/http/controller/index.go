package controller

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewStreamController,
	NewMessageController,
)
