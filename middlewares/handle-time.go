package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"
)

func HandlingTime(context *fiber.Ctx) error {
	context.Locals("handling", gohelpers.MakeTimestamp())
	return context.Next()
}
