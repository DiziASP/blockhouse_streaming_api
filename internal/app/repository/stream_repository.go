package repository

import "blockhouse_streaming_api/internal/domain/entity"

type StreamRepository interface {
	// CreateStream initializes a new topic/partition in Redpanda and returns the StreamEntity.
	CreateStream() (*entity.StreamEntity, error)
}
