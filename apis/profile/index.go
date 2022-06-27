package profile

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/profile")

	api.Delete("/", middlewares.Authorize, deleteProfileController)
}
