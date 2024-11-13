package repository

import "blockhouse_streaming_api/internal/domain/entity"

type MessageRepository interface {
	SendMessage(msg entity.MessageEntity) error
	FetchMessage(streamId string) error
}
