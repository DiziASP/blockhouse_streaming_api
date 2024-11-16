package kafka

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConsumeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock consumer
	mockConsumer := NewMockConsumer(ctrl)

	ctx := context.Background()
	handler := func(data []byte) {
		assert.Equal(t, "test-message", string(data))
	}

	// Define expected behavior for the mock
	mockConsumer.EXPECT().Consume(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, handler func([]byte)) error {
		handler([]byte("test-message"))
		return nil
	})

	// Call the Consume method on the mockConsumer
	err := mockConsumer.Consume(ctx, handler)
	assert.NoError(t, err)
}
