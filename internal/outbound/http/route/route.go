package route

import (
	"blockhouse_streaming_api/internal/outbound/http/controller"
	"github.com/gofiber/fiber/v2"
)

type MainRouter interface {
	Init(root *fiber.Router)
}

type mainRouter struct {
	messageApi controller.MessageController
	streamApi  controller.StreamController
}

func (v mainRouter) Init(root *fiber.Router) {
	mainRouter := (*root).Group("/v1")
	{
		mainRouter.Post("/:id/send", v.messageApi.SendMessage)
		mainRouter.Get("/:id/results", v.messageApi.FetchMessage)
	}
}

func NewMainRouter(
	messageApi controller.MessageController,
	streamApi controller.StreamController,
) MainRouter {
	return &mainRouter{
		messageApi: messageApi,
		streamApi:  streamApi,
	}
}
