package password

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/password")

	api.Post("/", middlewares.Authorize, changePasswordController)
}
