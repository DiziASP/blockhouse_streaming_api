package controller

import (
	"blockhouse_streaming_api/internal/app/service"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type MessageController interface {
	SendMessage(ctx *fiber.Ctx) error
	FetchMessage(c *websocket.Conn)
}

type messageController struct {
	messageService service.MessageService
}

// NewMessageController Constructor
func NewMessageController(messageService service.MessageService) MessageController {
	return &messageController{
		messageService: messageService,
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
	panic("implement me")
}

// FetchMessage Fetch new messages
// @Summary Fetch new messages
// @Tags Message
// @Accept stream
// @Produce stream
// @Router /:id/results [post]
// FetchMessage streams real-time messages from Kafka to WebSocket clients using Fiber context
func (m *messageController) FetchMessage(c *websocket.Conn) {
	panic("implement me")
}
