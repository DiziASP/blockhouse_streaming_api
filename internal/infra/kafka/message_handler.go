package kafka

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/file/json"
	"blockhouse_streaming_api/pkg/kafka"
	"context"
	"github.com/google/uuid"
	"log"
)

type MessageHandler struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
	kafkaAdm *kafka.Admin
}

func NewMessageHandler(kafkaAdm *kafka.Admin, producer *kafka.Producer, consumer *kafka.Consumer) repository.MessageRepository {
	return &MessageHandler{
		producer: producer,
		consumer: consumer,
		kafkaAdm: kafkaAdm,
	}
}

func (m MessageHandler) Publish(ctx context.Context, message entity.MessageEntity) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	topic := message.StreamID.String()
	err = m.producer.Produce(ctx, topic, data)
	if err != nil {
		return err
	}

	log.Printf("Message published to topic %s\n", topic)
	return nil
}

func (m MessageHandler) Consume(ctx context.Context, streamID uuid.UUID, handler func(data []byte)) error {
	topic := streamID.String()
	err := m.consumer.Consume(ctx, topic, handler)
	if err != nil {
		return err
	}

	log.Printf("Message consumed from topic %s\n", topic)
	return nil
}
