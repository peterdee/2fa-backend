package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/auth")

	api.Get("/code", middlewares.Authorize, getAuthCodeController)
	api.Get("/logout", middlewares.Authorize, logoutController)
	api.Post("/code", codeSignInController)
	api.Post("/sign-in", signInController)
	api.Post("/sign-up", signUpController)
}
