package account

import (
	"context"
	"fmt"

	"github.com/Timothylock/go-signin-with-apple/apple"
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
    RedirectURI string `json:"redirect_uri"`
}

func (a *Account) validateAppleAuthInput(input *appleAuthRequest) error {
    if input.Code == "" {
        return fmt.Errorf("code is required")
    }
    return nil
}

func (a *Account) HandleAppleAuth(ctx *fiber.Ctx) error {
    var input appleAuthRequest
    if err := ctx.BodyParser(&input); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if err := a.validateAppleAuthInput(&input); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    secret, err := apple.GenerateClientSecret(
        a.appleConfig.PrivateKey,
        a.appleConfig.TeamID, 
        a.appleConfig.ClientID,
        a.appleConfig.KeyID,
    )
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate client secret",
        })
    }

    client := apple.New()
    vReq := apple.AppValidationTokenRequest{
        ClientID:     a.appleConfig.ClientID,
        ClientSecret: secret,
        Code:        input.Code,
    }

    var resp apple.ValidationResponse
    if err := client.VerifyAppToken(context.Background(), vReq, &resp); err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid Apple token",
        })
    }

    // Get unique Apple ID
    appleUserID, err := apple.GetUniqueID(resp.IDToken) 
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get Apple user ID",
        })
    }

    // Get user claims like email
    claims, err := apple.GetClaims(resp.IDToken)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to get user claims",
        })
    }

    email := (*claims)["email"].(string)
    emailVerified := (*claims)["email_verified"].(bool)
    
    // Check if user exists
    exists, err := a.repo.CheckEmailExists(ctx.Context(), email)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Database error",
        })
    }

    var user repo.User
    if !exists {
        params := repo.CreateUserParams{
            Email: email,
            Name:  email, 
        }
        user, err = a.repo.CreateUser(ctx.Context(), params)
        if err != nil {
            return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to create user",
            })
        }
    } else {
        user, err = a.repo.GetUserByEmail(ctx.Context(), email)
        if err != nil {
            return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to get user",
            })
        }
    }

    return ctx.JSON(fiber.Map{
        "user": user,
        "apple_user_id": appleUserID,
        "email_verified": emailVerified,
    })
}