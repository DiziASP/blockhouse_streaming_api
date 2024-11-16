//go:build wireinject
// +build wireinject

package internal

import (
	"blockhouse_streaming_api/config"
	"blockhouse_streaming_api/internal/app/service"
	"blockhouse_streaming_api/internal/common/logger"
	"blockhouse_streaming_api/internal/infra"
	"blockhouse_streaming_api/internal/outbound/http/controller"
	"blockhouse_streaming_api/internal/outbound/http/route"
	"blockhouse_streaming_api/pkg/file/json"
	"blockhouse_streaming_api/pkg/kafka"
	loggerPkg "blockhouse_streaming_api/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	"os"
	"time"
)

type Server struct {
	app    *fiber.App
	cfg    *config.Configuration
	logger loggerPkg.Logger
}

func New() (*Server, error) {
	panic(wire.Build(wire.NewSet(
		config.Set,
		infra.Set,
		route.Set,
		service.Set,
		logger.Set,
		controller.Set,
		kafka.Set,
		//sessions.Set,
		NewServerInstance,
	)))
}

func NewServerInstance(
	cfg *config.Configuration,
	logger loggerPkg.Logger,
	mainRouter route.MainRouter,
) *Server {
	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		AppName: cfg.Server.Name,
		Prefork: cfg.Server.Prefork,
		//ErrorHandler: errors.CustomErrorHandler,
		ReadTimeout:  time.Second * cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * cfg.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	// Fiber Logging
	app.Use(fiberLog.New(fiberLog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	api := app.Group("/") // Initialize Root Route
	mainRouter.Init(&api)

	return &Server{cfg: cfg, logger: logger, app: app}
}

func (server Server) App() *fiber.App {
	return server.app
}

func (server Server) Config() *config.Configuration {
	return server.cfg
}

func (server Server) Logger() loggerPkg.Logger {
	return server.logger
}
