package account

import (
	"errors"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/account/repo"
)

type AppleAuthConfig struct {
	TeamID     string
	ClientID   string
	KeyID      string
	PrivateKey string
}

type appleAuthRequest struct {
	Code        string `json:"code" validate:"required"`
	RedirectURI string `json:"redirect_uri" validate:"required"`
}

var config *AppleAuthConfig

func (a *Account) validateAppleAuthInput(input *appleAuthRequest) error {
	if input.Code == "" {
		return errors.New("Code is required")
	}
	return nil
}

func (a *Account) HandleAppleAuth(ctx *fiber.Ctx) error {
	email := "heyanantraj@gmail.com"
	emailVerified := true

	// Check if user exists
	exists, err := a.repo.CheckEmailExists(ctx.Context(), email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var user repo.User
	if !exists {
		params := repo.CreateUserParams{
			Email:    email,
			Username: email,
		}
		user, err = a.repo.CreateUser(ctx.Context(), params)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		var msg *message.Message = message.NewMessage(watermill.NewUUID(), []byte(user.ID))
		a.pubsub.Core().Publish("user-created.topic", msg)
	} else {
		user, err = a.repo.GetUserByEmail(ctx.Context(), email)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return ctx.JSON(fiber.Map{
		"user":           user,
		"email_verified": emailVerified,
	})
}
