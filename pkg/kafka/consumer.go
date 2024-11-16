package kafka

import (
	"blockhouse_streaming_api/config"
	"context"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
	"sync"
)

type Consumer interface {
	Consume(ctx context.Context, handler func(data []byte)) error
	Close()
}

type consumer struct {
	client *kgo.Client
	wg     sync.WaitGroup
}

func NewConsumer(cfg *config.Configuration, topic string) Consumer {
	groupID := uuid.New().String()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Kafka.Brokers...),
		kgo.ConsumeTopics(topic),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		panic("Failed to create consumer: " + err.Error())
	}
	return &consumer{client: client}
}

func (c *consumer) Consume(ctx context.Context, handler func([]byte)) error {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() {
			return nil
		}

		fetches.EachRecord(func(record *kgo.Record) {
			handler(record.Value)
		})
	}
}

func (c *consumer) Close() {
	c.client.Close()
}
