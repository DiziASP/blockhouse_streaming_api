package service

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/common/utils"
	"blockhouse_streaming_api/internal/domain/dto"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/file/json"
	"blockhouse_streaming_api/pkg/logger"
	"context"
	"github.com/google/uuid"
	"time"
)

type MessageService struct {
	logger logger.Logger
	repo   repository.MessageRepository
}

// NewMessageService creates a new instance of MessageService.
func NewMessageService(repo repository.MessageRepository, logger logger.Logger) MessageService {
	return MessageService{
		repo:   repo,
		logger: logger,
	}
}

// SendMessage publishes a message to the specified stream.
func (s *MessageService) SendMessage(ctx context.Context, req *dto.SendMessageDTO) error {
	message := entity.MessageEntity{
		StreamID:  req.StreamID,
		Message:   req.Message,
		Timestamp: time.Now(),
	}

	// Publish the message to Kafka
	return s.repo.Publish(ctx, message)
}

// FetchMessage subscribes to a stream and streams messages using a callback function.
func (s *MessageService) FetchMessage(ctx context.Context, streamID uuid.UUID, handler func(dto.FetchMessageDTO)) error {
	return s.repo.Consume(ctx, streamID, func(data []byte) {
		var message dto.FetchMessageDTO

		// Unmarshal JSON data into the FetchMessageDTO struct
		if err := json.Unmarshal(data, &message); err != nil {
			s.logger.Errorf("Failed to get messages: " + err.Error())
			return
		}

		message.Message = utils.ProfanityFilter(message.Message)

		handler(message)
	})
}
