package kafka

import (
	"blockhouse_streaming_api/internal/app/repository"
	"blockhouse_streaming_api/internal/domain/entity"
	"blockhouse_streaming_api/pkg/kafka"
	"github.com/google/uuid"
)

type StreamHandler struct {
	kafkaAdm *kafka.Admin
}

func NewStreamHandler(kafkaAdm kafka.Admin) repository.StreamRepository {
	return &StreamHandler{
		kafkaAdm: &kafkaAdm,
	}
}

func (s StreamHandler) CreateStream() (*entity.StreamEntity, error) {
	streamId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	s.kafkaAdm.CreateTopic(streamId.String())

	return &entity.StreamEntity{StreamID: streamId}, nil
}
