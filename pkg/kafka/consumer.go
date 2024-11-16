package kafka

import (
	"blockhouse_streaming_api/config"
	"context"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

type Consumer struct {
	client *kgo.Client
	wg     sync.WaitGroup
}

func NewConsumer(cfg *config.Configuration) *Consumer {
	groupID := uuid.New().String()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Kafka.Brokers...),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		panic(err)
	}
	return &Consumer{client: client}
}

func (c *Consumer) Consume(ctx context.Context, topic string, handler func([]byte)) error {
	panic("implement me")
}

func (c *Consumer) Close() {
	c.client.Close()
}
