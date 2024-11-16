package controller

import (
	"blockhouse_streaming_api/internal/app/service"
	"blockhouse_streaming_api/internal/common/errors"
	responses "blockhouse_streaming_api/internal/common/response"
	"blockhouse_streaming_api/internal/domain/dto"
	"blockhouse_streaming_api/pkg/file/json"
	zapLogger "blockhouse_streaming_api/pkg/logger"
	"context"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type MessageController interface {
	SendMessage(ctx *fiber.Ctx) error
	FetchMessage(c *websocket.Conn)
}

type messageController struct {
	messageService service.MessageService
	logger         zapLogger.Logger
}

// NewMessageController Constructor
func NewMessageController(messageService service.MessageService, logger zapLogger.Logger) MessageController {
	return &messageController{
		messageService: messageService,
		logger:         logger,
	}
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
func (m *messageController) SendMessage(ctx *fiber.Ctx) error {
	// Parse request body
	var request dto.SendMessageDTO

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	ctx.SetUserContext(context.WithValue(ctx.UserContext(), "request_body", request))

	var errInstance error

	_streamId := ctx.Params("id", "")
	if _streamId == "" {
		errInstance := errors.ErrBadRequest
		errInstance.Message = "Stream ID is missing"
		return errInstance
	}

	streamId, err := uuid.Parse(_streamId)
	if err != nil {
		return err
	}
	request.StreamID = streamId

	// Publish message to Kafka
	if err := m.messageService.SendMessage(ctx.Context(), &request); err != nil {
		errInstance := errors.ErrInternalServer
		errInstance.Message = err.Error()
		return errInstance
	}

	defer func() {
		requestBody := ctx.UserContext().Value("request_body")
		requestBodyJSON, _ := json.Marshal(requestBody)

		reqID, _ := ctx.UserContext().Value("request_id").(string)

		if errInstance != nil {
			// Log the error
			m.logger.WithFields(
				zap.String("request_id", reqID),
				zap.String("request", string(requestBodyJSON)),
				zap.String("error", errInstance.Error()),
			).Errorf("Failed to send message")
		} else {
			// Log the success
			m.logger.WithFields(
				zap.String("request_id", reqID),
				zap.String("request", string(requestBodyJSON)),
			).Infof("Message sent successfully")
		}

	}()

	res := responses.DefaultSuccessResponse
	res.Message = "Message sent successfully"
	res.Data = request

	return res.JSON(ctx)
}

// FetchMessage Fetch new messages
// @Summary Fetch new messages
// @Tags Message
// @Accept stream
// @Produce stream
// @Router /:id/results [post]
// FetchMessage streams real-time messages from Kafka to WebSocket clients using Fiber context
func (m *messageController) FetchMessage(c *websocket.Conn) {
	ctx := context.Background()
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			m.logger.Errorf("Failed to close websocket connection")
			return
		} else {
			m.logger.Infof("Websocket connection has been closed by client")
			return
		}
	}(c)

	streamID := c.Params("id")

	err := m.messageService.FetchMessage(ctx, uuid.MustParse(streamID), func(msg dto.FetchMessageDTO) {
		if err := c.WriteJSON(msg); err != nil {
			m.logger.Errorf("Failed to fetch message: " + err.Error())
			return
		}
	})

	if err != nil {
		err := c.WriteMessage(websocket.CloseMessage, []byte("Error fetching messages"))
		if err != nil {
			m.logger.Errorf("Failed to close websocket connection")
			return
		}
	}

}
