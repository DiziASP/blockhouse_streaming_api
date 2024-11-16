package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

// RequestIDMiddleware generates a unique request ID for each request
func RequestIDMiddleware(c *fiber.Ctx) error {
	reqID := uuid.New().String()

	c.Locals(RequestIDKey, reqID)

	ctx := context.WithValue(c.Context(), RequestIDKey, reqID)
	c.SetUserContext(ctx)

	c.Set("X-Request-ID", reqID)

	return c.Next()
}
