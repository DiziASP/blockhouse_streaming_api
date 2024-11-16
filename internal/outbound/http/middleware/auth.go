package middleware

import (
	"blockhouse_streaming_api/config"
	"blockhouse_streaming_api/internal/common/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func APIKeyAuthMiddleware(cfg *config.Configuration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the API key from headers
		apiKey := c.Get("X-API-Key")
		if apiKey == "" {
			err := errors.ErrUnauthenticated
			err.Message = "Missing API key"
			return err
		}

		// Validate the API key
		key, err := uuid.Parse(apiKey)
		if err != nil {
			err := errors.ErrUnauthenticated
			err.Message = "Invalid API key"
			return err
		}

		// Check if the API key matches the expected value from the configuration
		if key.String() != cfg.Server.ApiKey {
			err := errors.ErrUnauthenticated
			err.Message = "Invalid API key"
			return err
		}

		// Continue to the next handler
		return c.Next()
	}
}
