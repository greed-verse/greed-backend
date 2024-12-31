package payment

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/payment/repo"
	"github.com/greed-verse/greed/internal/shared"
)

type Payment struct {
	logger *shared.Logger
	pubsub *shared.PubSub
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
		pubsub: context.PubSub,
		router: router,
		repo:   repo,
	}

	module.serve()
	module.subscribe()
	return module
}

func (p *Payment) serve() {
	p.router.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Module: OK")
	})
}

func (p *Payment) subscribe() error {
	_, err := p.pubsub.Core().Subscribe(context.Background(), "user-created.topic")
	if err != nil {
		return err
	}

	handler := p.pubsub.Router().AddNoPublisherHandler("user-created.handler", "user-created.topic", p.pubsub.Core(), p.walletHandler)
	handler.AddMiddleware(p.walletMiddleware)
	return nil
}

func (p *Payment) walletMiddleware(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		p.logger.Core().Info().Msg("User was created... creating corresponding wallet")
		producedMsg, err := h(msg)
		return producedMsg, err
	}
}
