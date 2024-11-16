package integration

import (
	"context"
	"testing"
	"time"

	"blockhouse_streaming_api/pkg/kafka"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRealTimeStreaming(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	streamID := uuid.New().String()

	// Create mock producer and consumer
	mockProducer := kafka.NewMockProducer(ctrl)
	mockConsumer := kafka.NewMockConsumer(ctrl)

	message := []byte("real-time message test")

	// Set up the mock behavior for the consumer
	mockConsumer.EXPECT().Consume(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, handler func(data []byte)) error {
		// Simulate receiving a message
		handler(message)
		return nil
	})

	// Set up the mock behavior for the producer
	mockProducer.EXPECT().Produce(ctx, streamID, "test-key", message).Return(nil)

	// Start consumer in a separate goroutine
	go func() {
		err := mockConsumer.Consume(ctx, func(data []byte) {
			assert.Equal(t, message, data)
		})
		assert.NoError(t, err)
	}()

	// Produce a message
	time.Sleep(1 * time.Second)
	err := mockProducer.Produce(ctx, streamID, "test-key", message)
	assert.NoError(t, err)

	time.Sleep(2 * time.Second)
}
