package controller

import (
	"blockhouse_streaming_api/internal/app/service"
	responses "blockhouse_streaming_api/internal/common/response"
	zapLogger "blockhouse_streaming_api/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type StreamController interface {
	CreateStream(ctx *fiber.Ctx) error
}

type streamController struct {
	streamService service.StreamService
	logger        zapLogger.Logger
}

// NewStreamController Constructor
func NewStreamController(streamService service.StreamService, logger zapLogger.Logger) StreamController {
	return &streamController{streamService: streamService, logger: logger}
}

// CreateStream Create new channel stream
// @Summary Create new channel stream
// @Tags Stream
// @Accept json
// @Produce json
// @Param stream body dto.CreateStreamDTO true "message data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /:id/send [post]
func (s streamController) CreateStream(ctx *fiber.Ctx) error {
	context := ctx.Context()

	data, err := s.streamService.CreateStream(context)
	if err != nil {
		s.logger.Errorf("Failed to create new stream: " + err.Error())
		return err
	}

	resp := responses.DefaultSuccessResponse
	resp.Message = fmt.Sprintf("Stream %s created successfully", data.StreamID.String())
	resp.Data = data
	return resp.JSON(ctx)
}
