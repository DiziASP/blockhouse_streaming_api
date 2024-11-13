package service

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"context"
)

type StreamService struct {
	repo repository.StreamRepository
}

// NewStreamService creates a new instance of StreamService.
func NewStreamService(repo repository.StreamRepository) *StreamService {
	return &StreamService{repo: repo}
}

// CreateStream initializes a new topic/partition in Redpanda and returns the StreamEntity.
func (s *StreamService) CreateStream(ctx context.Context) (*entity.StreamEntity, error) {
	return s.repo.CreateStream()
}
