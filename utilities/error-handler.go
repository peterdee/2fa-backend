package utilities

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
)

func CustomErrorHandler(context *fiber.Ctx, thrownError error) error {
	info := configuration.RESPONSE_MESSAGES.InternalServerError
	status := fiber.StatusInternalServerError

	if e, ok := thrownError.(*fiber.Error); ok {
		info = e.Message
		status = e.Code
	}

	return Response(ResponsePayloadStruct{
		Context: context,
		Info:    info,
		Status:  status,
	})
}
