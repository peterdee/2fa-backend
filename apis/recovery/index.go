package recovery

import "github.com/gofiber/fiber/v2"

func Initialize(app *fiber.App) {
	api := app.Group("/api/recovery")

	api.Post("/check", checkLoginController)
	api.Patch("/update", updatePasswordController)
}
