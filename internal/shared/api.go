package shared

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
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
		ErrorHandler:          errorHandler,
	}
	server := fiber.New(config)

	router := server.Group("/v1")
	return &API{core: server, router: router}
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = ctx.Status(code).JSON(fiber.Map{
		"error": e.Error(),
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return nil
}

func (api *API) Core() *fiber.App {
	return api.core
}

func (api *API) Router() fiber.Router {
	return api.router
}

func (api *API) InitMiddlewares(logger *Logger) {
	api.Core().Use(logger.Middleware())
}
