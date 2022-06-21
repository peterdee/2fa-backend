package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"

	"backend2fa/configuration"
)

func main() {
	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(helmet.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"info":   configuration.RESPONSE_MESSAGES.OK,
			"status": fiber.StatusOK,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = configuration.DEFAULT_PORT
	}
	app.Listen(":" + port)
}
