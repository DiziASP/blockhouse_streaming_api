package repository

import (
	"blockhouse_streaming_api/internal/domain/entity"
	"context"
)

type StreamRepository interface {
	// CreateStream initializes a new topic/partition in Redpanda and returns the StreamEntity.
	CreateStream(ctx context.Context) (*entity.StreamEntity, error)
}
