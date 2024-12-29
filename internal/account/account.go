package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/account/repo"
	"github.com/greed-verse/greed/internal/shared"
)

type Account struct {
	logger *shared.Logger
	router fiber.Router
	repo   *repo.Queries

	wallet shared.WalletService
}

func New(context *shared.AppContext, walletService shared.WalletService) *Account {
	repo := repo.New(context.Repo)
	router := context.API.Router().Group("/account")

	var module *Account = &Account{
		logger: context.Logger,
		router: router,
		repo:   repo,
	}
	module.Serve()
	return module
}

func (a *Account) Serve() {
	a.router.Get("/health", a.Health)
	a.router.Post("/auth/apple", a.HandleAppleAuth)
	a.router.Get("/auth/logout", a.HandleAppleAuth)
	a.router.Get("/user/profile/:resource?", a.HandleAppleAuth)
}
