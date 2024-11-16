package kafka

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProduceMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Use the generated mock implementation
	mockProducer := NewMockProducer(ctrl)

	ctx := context.Background()
	topic := "test-topic"
	key := "test-key"
	data := []byte("test-message")

	// Define the expected behavior for the mock
	mockProducer.EXPECT().Produce(ctx, topic, key, data).Return(nil)

	// Use the mockProducer directly as it implements the Producer interface
	err := mockProducer.Produce(ctx, topic, key, data)
	assert.NoError(t, err)
}
