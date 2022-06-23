package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"

	"backend2fa/configuration"
	"backend2fa/utilities"
)

func signInController(context *fiber.Ctx) error {
	payload := new(signInPayload)
	if err := context.BodyParser(payload); err != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	if payload.ClientType == "" || payload.Login == "" || payload.Password == "" {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.MissingData,
			Status:  fiber.StatusBadRequest,
		})
	}

	clients := gohelpers.ObjectValues(configuration.CLIENT_TYPES)
	if !gohelpers.IncludesString(clients, payload.ClientType) {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InvalidData,
			Status:  fiber.StatusBadRequest,
		})
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}