package password

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func changePasswordController(context *fiber.Ctx) error {
	payload := new(changePasswordPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	newPassword := strings.Trim(payload.NewPassword, " ")
	oldPassword := strings.Trim(payload.OldPassword, " ")
	if newPassword == "" || oldPassword == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	userId := context.Locals("userId").(uint)
	var passwordRecord = models.Passwords{UserID: userId}
	result := database.Connection.Find(&passwordRecord)
	if result.Error != nil {

	}
	if result.RowsAffected == 0 {

	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
	})
}
