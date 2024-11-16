package decorator

import (
	"blockhouse_streaming_api/pkg/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type HandlerFunc func(ctx context.Context) error

func WithLogging(log logger.Logger, handler HandlerFunc, action string) HandlerFunc {
	return func(ctx context.Context) error {
		// Extract request ID from context
		reqID, _ := ctx.Value("request_id").(string)

		// Execute the handler function
		if err := handler(ctx); err != nil {
			// Log the error with the request ID
			log.WithFields(
				zap.String("request_id", reqID),
				zap.String("body", action),
				zap.String("error", err.Error()),
			).Errorf("Action failed")

			// Return a formatted error
			return fmt.Errorf("error during %s: %w", action, err)
		}

		// Log the successful completion of the action
		log.WithFields(
			zap.String("request_id", reqID),
			zap.String("action", action),
		).Infof("Action completed successfully")

		return nil
	}
}
