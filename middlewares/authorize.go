package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Authorize(context *fiber.Ctx) error {
	token := context.Get(fiber.HeaderAuthorization)
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "TEST")
		// return utilities.Response(utilities.ResponsePayloadStruct{
		// 	Context: context,
		// 	Info:    configuration.RESPONSE_MESSAGES.MissingToken,
		// 	Status:  fiber.StatusUnauthorized,
		// })
	}

	// return context.Next(*fiber.NewError(fiber.StatusUnauthorized, "TEST"))
	return fiber.NewError(fiber.StatusUnauthorized, "TEST")
}
