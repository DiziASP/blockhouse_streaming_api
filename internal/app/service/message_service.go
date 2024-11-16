package service

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/dto"
	"context"
	"github.com/google/uuid"
)

type MessageService struct {
	repo repository.MessageRepository
}

// NewMessageService creates a new instance of MessageService.
func NewMessageService(repo repository.MessageRepository) MessageService {
	return MessageService{
		repo: repo,
	}
}

// SendMessage publishes a message to the specified stream.
func (s *MessageService) SendMessage(ctx context.Context, req *dto.SendMessageDTO) error {
	panic("implement me")
}

// FetchMessage subscribes to a stream and streams messages using a callback function.
func (s *MessageService) FetchMessage(ctx context.Context, streamID uuid.UUID, handler func(dto.FetchMessageDTO)) error {
	panic("implement me")
}
