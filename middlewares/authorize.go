package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
)

func Authorize(context *fiber.Ctx) error {
	token := context.Get(fiber.HeaderAuthorization)
	if token == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.MissingToken,
		)
	}

	return fiber.NewError(fiber.StatusUnauthorized, "TEST")
}
