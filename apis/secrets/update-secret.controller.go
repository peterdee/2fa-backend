package secrets

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/utilities"
)

// TODO: finish this
func updateSecretController(context *fiber.Ctx) error {
	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
