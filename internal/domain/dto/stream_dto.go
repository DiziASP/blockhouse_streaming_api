package dto

import "github.com/google/uuid"

type CreateStreamDTO struct {
	StreamID uuid.UUID `json:"stream_id"`
}
