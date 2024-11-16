package kafka

import (
	"blockhouse_streaming_api/config"
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer interface {
	Produce(ctx context.Context, topic string, key string, data []byte) error
	Close()
}

type producer struct {
	client *kgo.Client
}

func NewProducer(cfg *config.Configuration) Producer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Kafka.Brokers...),
	)
	if err != nil {
		panic("Failed to create producer: " + err.Error())
	}
	return &producer{client: client}
}

func (p *producer) Produce(ctx context.Context, topic string, key string, data []byte) error {
	record := &kgo.Record{
		Topic: topic,
		Key:   []byte(key),
		Value: data,
	}

	// Define a channel to capture the result of the Produce call
	errChan := make(chan error, 1)

	// Produce the message with a callback to capture the result
	p.client.Produce(ctx, record, func(rec *kgo.Record, err error) {
		if err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	})

	// Wait for the result and return the error if any
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err() // Handle context cancellation
	}
}

func (p *producer) Close() {
	p.client.Close()
}
