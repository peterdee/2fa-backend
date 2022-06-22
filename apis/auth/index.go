package auth

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Get("/sign-in", signInController)
	api.Get("/sign-up", signUpController)
}
