package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
)

func ClientTypeCheck(context *fiber.Ctx) error {
	clientType := context.Locals("client").(string)
	if clientType != configuration.CLIENT_TYPES.Mobile {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidClientType,
		)
	}

	return context.Next()
}
