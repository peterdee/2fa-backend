package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/utilities"
)

func signUpController(context *fiber.Ctx) error {
	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
