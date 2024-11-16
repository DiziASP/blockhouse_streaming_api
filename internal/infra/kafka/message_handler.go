package kafka

import (
	"blockhouse_streaming_api/config"
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/file/json"
	"blockhouse_streaming_api/pkg/kafka"
	"context"
	"github.com/google/uuid"
	"strconv"
)

type MessageHandler struct {
	cfg      *config.Configuration
	kafkaAdm *kafka.Admin
}

func NewMessageHandler(cfg *config.Configuration, kafkaAdm *kafka.Admin) repository.MessageRepository {
	return &MessageHandler{
		cfg:      cfg,
		kafkaAdm: kafkaAdm,
	}
}

func (m *MessageHandler) Publish(ctx context.Context, message entity.MessageEntity) error {
	topic := message.StreamID.String() // Use StreamID as the topic name
	key := strconv.FormatInt(message.Timestamp.UnixMilli(), 10)
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	producer := kafka.NewProducer(m.cfg)
	defer producer.Close()

	return producer.Produce(ctx, topic, key, data)
}

func (m *MessageHandler) Consume(ctx context.Context, streamID uuid.UUID, handler func(data []byte)) error {
	topic := streamID.String()
	consumer := kafka.NewConsumer(m.cfg, topic)
	defer consumer.Close()

	return consumer.Consume(ctx, handler)
}
