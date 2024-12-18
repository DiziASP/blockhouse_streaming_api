package kafka

import (
	"blockhouse_streaming_api/config"
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Admin struct {
	cfg    *config.Configuration
	client *kadm.Client
}

func NewAdmin(cfg *config.Configuration) *Admin {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Kafka.Brokers...),
	)
	if err != nil {
		panic(err)
	}
	admin := kadm.NewClient(client)
	return &Admin{
		cfg:    cfg,
		client: admin,
	}
}
func (a *Admin) TopicExists(topic string) bool {
	ctx := context.Background()
	topicsMetadata, err := a.client.ListTopics(ctx)
	if err != nil {
		panic(err)
	}
	for _, metadata := range topicsMetadata {
		if metadata.Topic == topic {
			return true
		}
	}
	return false
}
func (a *Admin) CreateTopic(topic string) {
	ctx := context.Background()
	resp, err := a.client.CreateTopics(ctx, 1, 1, nil, topic)
	if err != nil {
		panic(err)
	}
	for _, ctr := range resp {
		if ctr.Err != nil {
			fmt.Printf("Unable to create topic '%s': %s", ctr.Topic, ctr.Err)
		} else {
			fmt.Printf("Created topic '%s'\n", ctr.Topic)
		}
	}
}
func (a *Admin) Close() {
	a.client.Close()
}
