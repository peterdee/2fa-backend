package auth

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Post("/sign-in", signInController)
	api.Post("/sign-up", signUpController)
}
