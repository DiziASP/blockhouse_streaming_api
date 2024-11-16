package kafka

import (
	"blockhouse_streaming_api/config"
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(cfg *config.Configuration) *Producer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Kafka.Brokers...),
	)
	if err != nil {
		panic(err)
	}
	return &Producer{client: client}
}

func (p *Producer) Produce(ctx context.Context, topic string, data []byte) error {
	panic("implement me")
}

func (p *Producer) Close() {
	p.client.Close()
}
