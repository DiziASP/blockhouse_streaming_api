package entity

import (
	"github.com/google/uuid"
	"time"
)

type MessageEntity struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	StreamID  uuid.UUID `json:"stream_id"`
}
