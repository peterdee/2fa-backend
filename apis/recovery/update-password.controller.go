package recovery

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func updatePasswordController(context *fiber.Ctx) error {
	payload := new(updatePasswordPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	newPassword := strings.ToLower(strings.Trim(payload.NewPassword, " "))
	recoveryAnswer := strings.ToLower(strings.Trim(payload.RecoveryAnswer, " "))
	userId := payload.UserID
	if newPassword == "" || recoveryAnswer == "" || userId == 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	// TODO: add all of the necessary checks for the new password

	var user models.Users
	result := database.Connection.Where("id = ?", userId).Find(&user)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	// TODO: validate all of the data, change the password and issue a new token

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
	})
}
