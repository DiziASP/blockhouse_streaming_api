package service

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/dto"
	"blockhouse_streaming_api/internal/domain/entity"
	timePkg "blockhouse_streaming_api/pkg/time"
	"context"
)

type MessageService struct {
	repo         repository.MessageRepository
	timeProvider timePkg.Provider
}

// NewMessageService creates a new instance of MessageService.
func NewMessageService(repo repository.MessageRepository, timeProvider timePkg.Provider) *MessageService {
	return &MessageService{
		repo:         repo,
		timeProvider: timeProvider,
	}
}

// SendMessage publishes a message to the specified stream.
func (s *MessageService) SendMessage(ctx context.Context, req *dto.SendMessageDTO) error {
	data := entity.MessageEntity{
		Message:   req.Message,
		Timestamp: s.timeProvider.Now(),
	}

	return s.repo.SendMessage(data)
}

// FetchMessage subscribes to a stream and read incoming messages.
func (s *MessageService) FetchMessage(ctx context.Context, req *dto.FetchMessageDTO) error {
	streamId := req.StreamID

	return s.repo.FetchMessage(streamId)
}
