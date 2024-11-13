package errors

import (
	responses "blockhouse_streaming_api/internal/common/response"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := responses.DefaultErrorResponse

	// retrieve the custom status code if it's a fiber.*Error
	var e *fiber.Error
	if errors.As(err, &e) {
		msg.StatusCode = e.Code
		msg.Message = e.Message
	}
	var customErr *Error
	if errors.As(err, &customErr) {
		msg = responses.BindingGeneral(customErr)
		msg.StatusCode = customErr.StatusCode
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return msg.JSON(ctx)
}
