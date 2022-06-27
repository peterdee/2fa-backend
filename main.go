package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"

	authAPI "backend2fa/apis/auth"
	indexAPI "backend2fa/apis/index"
	passwordAPI "backend2fa/apis/password"
	profileAPI "backend2fa/apis/profile"
	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/utilities"
)

func main() {
	env := os.Getenv("ENV")
	if env != configuration.ENVS.Heroku {
		environmentError := godotenv.Load()
		if environmentError != nil {
			log.Fatal("Could not load environment variables!")
		}
	}

	database.Connect()

	app := fiber.New(fiber.Config{
		ErrorHandler: utilities.CustomErrorHandler,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))
	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(helmet.New())
	app.Use(logger.New())

	authAPI.Initialize(app)
	indexAPI.Initialize(app)
	passwordAPI.Initialize(app)
	profileAPI.Initialize(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = configuration.DEFAULT_PORT
	}

	log.Fatal(app.Listen(":" + port))
}
