package payment

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/shared"
)

type Payment struct {
	logger *shared.Logger
	router fiber.Router
	repo   *repo.Queries
}

func InitModule(context *shared.AppContext) {
    router := context.API.Router().Group("/payment")
    var module *Payment = &Payment{
        logger: context.Logger,
        router: router,
    }

}
