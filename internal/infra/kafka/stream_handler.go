package kafka

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/kafka"
	"blockhouse_streaming_api/pkg/uuid"
	"context"
	"errors"
)

type StreamHandler struct {
	kafkaAdm *kafka.Admin
}

func NewStreamHandler(kafkaAdm *kafka.Admin) repository.StreamRepository {
	return &StreamHandler{
		kafkaAdm: kafkaAdm,
	}
}

func (s StreamHandler) CreateStream(ctx context.Context) (*entity.StreamEntity, error) {
	streamId := uuid.NewUUIDProvider().NewUUID()

	if !s.kafkaAdm.TopicExists(streamId.String()) {
		s.kafkaAdm.CreateTopic(streamId.String())
		return &entity.StreamEntity{StreamID: streamId}, nil
	}

	return nil, errors.New("topic already exists")
}
