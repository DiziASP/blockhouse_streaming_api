package route

import (
	"blockhouse_streaming_api/internal/common/utils"
	"blockhouse_streaming_api/internal/outbound/http/controller"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewMainRouter,
)

type MainRouter interface {
	Init(root *fiber.Router)
}

type mainRouter struct {
	messageApi controller.MessageController
	streamApi  controller.StreamController
}

func (v mainRouter) Init(root *fiber.Router) {
	mainRouter := (*root).Group("/stream")
	{
		mainRouter.Post("/start", v.streamApi.CreateStream)
		mainRouter.Post("/:id/send", v.messageApi.SendMessage)
		mainRouter.Get("/:id/results", utils.WebsocketHandler, websocket.New(v.messageApi.FetchMessage))
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
