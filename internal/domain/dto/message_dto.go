package dto

import "github.com/google/uuid"

type SendMessageDTO struct {
	Message  string    `json:"message"`
	StreamID uuid.UUID `json:"stream_id"`
}

type FetchMessageDTO struct {
	StreamID uuid.UUID `json:"stream_id"`
}
