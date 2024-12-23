package shared

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

type API struct {
	core   *fiber.App
	router fiber.Router
}

func NewApi(logger *Logger) *API {
	defer logger.Core().Info().Msg("API Initiated")

	var config = fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           time.Second * 5,
		WriteTimeout:          time.Second * 5,
	}
	server := fiber.New(config)

	router := server.Group("/v1")
	return &API{core: server, router: router}
}

func (api *API) Core() *fiber.App {
	return api.core
}

func (api *API) Router() fiber.Router {
	return api.router
}

func (api *API) InitMiddlewares(logger *Logger) {
	api.Core().Use(logger.Middleware())
	api.Core().Use(helmet.New())
}
