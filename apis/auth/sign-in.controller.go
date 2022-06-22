package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
)

func signInController(context *fiber.Ctx) error {
	payload := new(signInPayload)
	if err := context.BodyParser(payload); err != nil {

	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"info":   configuration.RESPONSE_MESSAGES.OK,
		"status": fiber.StatusOK,
	})
}
