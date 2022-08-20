package secrets

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/middlewares"
)

func Initialize(app *fiber.App) {
	api := app.Group("/api/secrets")

	api.Delete(
		"/:id",
		middlewares.Authorize,
		middlewares.ClientTypeCheck,
		deleteSecretController,
	)
	api.Delete(
		"/delete/all",
		middlewares.Authorize,
		middlewares.ClientTypeCheck,
		deleteAllSecretsController,
	)
	api.Get("/", middlewares.Authorize, getSecretsController)
	api.Patch(
		"/:id",
		middlewares.Authorize,
		middlewares.ClientTypeCheck,
		updateSecretController,
	)
	api.Post(
		"/",
		middlewares.Authorize,
		middlewares.ClientTypeCheck,
		addSecretController,
	)
}
