package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/account/repo"
	"github.com/greed-verse/greed/internal/shared"
)

type Account struct {
	logger *shared.Logger
	pubsub *shared.PubSub
	router fiber.Router
	repo   *repo.Queries
}

func New(context *shared.AppContext) {
	repo := repo.New(context.Repo)
	router := context.API.Router().Group("/account")

	var module *Account = &Account{
		logger: context.Logger,
		pubsub: context.PubSub,
		router: router,
		repo:   repo,
	}
	module.subscribe()
	module.serve()
}

func (a *Account) serve() {
	a.router.Get("/health", a.Health)
	a.router.Post("/auth/apple", a.HandleAppleAuth)
	a.router.Get("/auth/logout", a.HandleAppleAuth)
	a.router.Get("/user/profile/:resource?", a.HandleAppleAuth)
}

func (a *Account) subscribe() {
}
