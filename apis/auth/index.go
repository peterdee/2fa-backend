package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Post("/sign-in", middlewares.Authorize, signInController)
	api.Post("/sign-up", signUpController)
}
