package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
)

func signUpController(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"info":   configuration.RESPONSE_MESSAGES.OK,
		"status": fiber.StatusOK,
	})
}
