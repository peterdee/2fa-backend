package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

func main() {
	app := fiber.New()

	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"info":   "OK",
			"status": 200,
		})
	})

	app.Listen(":2244")
}
