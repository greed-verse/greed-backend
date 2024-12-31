package account

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/Timothylock/go-signin-with-apple/apple"
	"github.com/gofiber/fiber/v2"
	"github.com/greed-verse/greed/internal/account/repo"
	"github.com/greed-verse/greed/pkg/env"
	"github.com/greed-verse/greed/pkg/validator"
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
	if config == nil {
		config = &AppleAuthConfig{
			TeamID:     env.GetEnv().APPLE_TEAM_ID(),
			ClientID:   env.GetEnv().APPLE_CLIENT_ID(),
			KeyID:      env.GetEnv().APPLE_KEY_ID(),
			PrivateKey: env.GetEnv().APPLE_PRIVATE_KEY(),
		}
	}

	var input appleAuthRequest
	if err := validator.GetValidator().Struct(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Request is missing fields")
	}

	secret, err := apple.GenerateClientSecret(
		config.PrivateKey,
		config.TeamID,
		config.ClientID,
		config.KeyID,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	client := apple.New()
	vReq := apple.AppValidationTokenRequest{
		ClientID:     config.ClientID,
		ClientSecret: secret,
		Code:         input.Code,
	}

	var resp apple.ValidationResponse
	if err := client.VerifyAppToken(context.Background(), vReq, &resp); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	// Get unique Apple ID
	appleUserID, err := apple.GetUniqueID(resp.IDToken)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Get user claims like email
	claims, err := apple.GetClaims(resp.IDToken)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	email := (*claims)["email"].(string)
	emailVerified := (*claims)["email_verified"].(bool)

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
		"apple_user_id":  appleUserID,
		"email_verified": emailVerified,
	})
}
