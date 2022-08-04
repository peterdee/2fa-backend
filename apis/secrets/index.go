package secrets

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/secrets")

	api.Delete("/:id", middlewares.Authorize, deleteSecretController)
	api.Delete("/delete/all", middlewares.Authorize, deleteAllSecretsController)
	api.Get("/", middlewares.Authorize, getSecretsController)
	api.Patch("/:id", middlewares.Authorize, updateSecretController)
	api.Post("/", middlewares.Authorize, addSecretController)
}
