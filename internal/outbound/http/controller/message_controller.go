package controller

import (
	"blockhouse_streaming_api/internal/app/service"
	"blockhouse_streaming_api/internal/common/errors"
	responses "blockhouse_streaming_api/internal/common/response"
	"blockhouse_streaming_api/internal/common/utils"
	"blockhouse_streaming_api/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MessageController interface {
	SendMessage(ctx *fiber.Ctx) error
	FetchMessage(ctx *fiber.Ctx) error
}

type messageController struct {
	messageService service.MessageService
}

// NewMessageController Constructor
func NewMessageController(messageService service.MessageService) MessageController {
	return &messageController{messageService: messageService}
}

// SendMessage Send a new message
// @Summary Send a new message
// @Tags Message
// @Accept json
// @Produce json
// @Param stream body dto.SendMessageDTO true "message data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /:id/send [post]
func (m messageController) SendMessage(ctx *fiber.Ctx) error {
	context := ctx.Context()

	_streamId := ctx.Params("id", "")
	if _streamId == "" {
		return errors.ErrBadRequest
	}

	streamId, err := uuid.Parse(_streamId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := new(dto.SendMessageDTO)
	if err := ctx.BodyParser(&req); err != nil {
		return errors.ErrBadRequest
	}
	req.StreamID = streamId

	if err := utils.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}

	if err := m.messageService.SendMessage(context, req); err != nil {
		return err
	}
	res := responses.DefaultSuccessResponse
	res.Message = "Message sent!"
	return res.JSON(ctx)
}

// FetchMessage Fetch new messages
// @Summary Fetch new messages
// @Tags Message
// @Accept json
// @Produce json
// @Param stream body dto.FetchMessageDTO true "The message data"
// @Success 200 {object} responses.General
// @Failure 500 {object} responses.General
// @Failure 400 {object} responses.General
// @Router /:id/results [post]
func (m messageController) FetchMessage(ctx *fiber.Ctx) error {
	//context := ctx.Context()

	_streamId := ctx.Params("id", "")
	if _streamId == "" {
		return errors.ErrBadRequest
	}

	streamId, err := uuid.Parse(_streamId)
	if err != nil {
		return errors.ErrBadRequest
	}

	req := new(dto.FetchMessageDTO)
	req.StreamID = streamId

	if err := utils.GetValidator().Validate(req); err != nil {
		return errors.ErrBadRequest
	}

	// TODO implement Fetch in realtime using SSE/WebSocket

	res := responses.DefaultSuccessResponse
	res.Message = "Message sent!"
	return res.JSON(ctx)
}
