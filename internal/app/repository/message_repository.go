package repository

import (
	"blockhouse_streaming_api/internal/domain/entity"
	"github.com/google/uuid"
)

type MessageRepository interface {
	SendMessage(msg entity.MessageEntity) error
	FetchMessage(streamId uuid.UUID) error
}
