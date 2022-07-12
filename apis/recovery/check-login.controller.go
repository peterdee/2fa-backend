package recovery

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func checkLoginController(context *fiber.Ctx) error {
	payload := new(checkLoginPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	login := strings.ToLower(strings.Trim(payload.Login, " "))
	if login == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	var user models.Users
	result := database.Connection.Where("login = ?", login).Find(&user)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
		Data: fiber.Map{
			"recoveryQuestion": user.RecoveryQuestion,
			"userId":           user.ID,
		},
	})
}
