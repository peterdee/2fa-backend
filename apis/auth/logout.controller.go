package auth

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func logoutController(context *fiber.Ctx) error {
	userId := context.Locals("userId").(uint)

	var tokenSecretRecord models.TokenSecrets
	result := database.Connection.Where("user_id = ?", userId).Find(&tokenSecretRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	newTokenSecret, secretHashError := utilities.CreateHash(utilities.CreateTokenKey(userId))
	if secretHashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	result = database.Connection.Model(&models.TokenSecrets{}).
		Where("id = ?", tokenSecretRecord.ID).
		Update("secret", newTokenSecret)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
