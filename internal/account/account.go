package account

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/account/repo"
	"github.com/greed-verse/greed/internal/shared"
)

type Account struct {
    logger     *shared.Logger
    router     fiber.Router
    repo       *repo.Queries
    appleConfig AppleAuthConfig 
}

func InitModule(context *shared.AppContext) {
    repo := repo.New(context.Repo)
    router := context.API.Router().Group("/account")
    
    config := AppleAuthConfig{
        TeamID:     os.Getenv("APPLE_TEAM_ID"),
        ClientID:   os.Getenv("APPLE_CLIENT_ID"), 
        KeyID:      os.Getenv("APPLE_KEY_ID"),
        PrivateKey: os.Getenv("APPLE_PRIVATE_KEY"),
    }
    
    var module *Account = &Account{
        logger: context.Logger,
        router: router,
        repo:   repo,
        appleConfig: config,
    }
    module.Serve()
}

func (a *Account) Serve() {
    a.router.Get("/health", a.Health)
    a.router.Post("/auth/apple", a.HandleAppleAuth) 
}
