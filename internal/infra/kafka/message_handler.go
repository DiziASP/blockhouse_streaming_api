package kafka

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
)

type MessageHandler struct {
	producer *kgo.Client
	consumer *kgo.Client
}

func NewMessageHandler(brokers []string) repository.MessageRepository {
	producer, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}

	consumer, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		panic(err)
	}

	return &MessageHandler{
		producer: producer,
		consumer: consumer,
	}
}

func (m MessageHandler) SendMessage(msg entity.MessageEntity) error {
	//TODO implement me
	panic("implement me")
}

func (m MessageHandler) FetchMessage(streamId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
