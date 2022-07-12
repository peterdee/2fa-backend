package account

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/account")

	api.Delete("/", middlewares.Authorize, deleteAccountController)
}
