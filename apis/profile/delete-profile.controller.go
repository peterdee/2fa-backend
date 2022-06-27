package profile

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func deleteProfileController(context *fiber.Ctx) error {
	userId := context.Locals("userId").(uint)

	result := database.Connection.Delete(&models.Users{}, userId)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
