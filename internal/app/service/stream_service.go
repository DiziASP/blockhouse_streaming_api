package service

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/logger"
	"context"
)

type StreamService struct {
	logger logger.Logger
	repo   repository.StreamRepository
}

// NewStreamService creates a new instance of StreamService.
func NewStreamService(repo repository.StreamRepository, logger logger.Logger) StreamService {
	return StreamService{
		repo:   repo,
		logger: logger,
	}
}

// CreateStream initializes a new topic/partition in Redpanda and returns the StreamEntity.
func (s *StreamService) CreateStream(ctx context.Context) (*entity.StreamEntity, error) {
	return s.repo.CreateStream(ctx)
}
