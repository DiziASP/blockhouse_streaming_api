package entity

import "github.com/google/uuid"

type StreamEntity struct {
	StreamID uuid.UUID `json:"stream_id"`
}
