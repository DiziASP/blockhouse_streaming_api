package repository

import (
	"blockhouse_streaming_api/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

type MessageRepository interface {
	// Publish publishes a message to the Kafka topic.
	Publish(ctx context.Context, message entity.MessageEntity) error

	// Consume subscribes to the Kafka topic and streams message.
	Consume(ctx context.Context, streamID uuid.UUID, handler func(data []byte)) error
}
