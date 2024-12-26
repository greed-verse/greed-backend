package account

import "github.com/gofiber/fiber/v2"

func (a *Account) Health(ctx *fiber.Ctx) error {
	return ctx.JSON("Server: OK")
}
