package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/greed-verse/greed/internal/shared"
)

type Payment struct {
	logger *shared.Logger
	router fiber.Router
	repo   *repo.Queries
}


func (p *Payment) GetRepo() *repo.Queries {
	return p.repo
}

func New(context *shared.AppContext) *Payment {
	router := context.API.Router().Group("/payment")
	repo := repo.New(context.Repo)

	var module *Payment = &Payment{
		logger: context.Logger,
		router: router,
		repo:   repo,
	}

	module.RegisterRoutes()
    return module
}

func (p *Payment) RegisterRoutes() {
	p.router.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Module: OK")
	})
}
