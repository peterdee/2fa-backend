package secrets

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func updateSecretController(context *fiber.Ctx) error {
	payload := new(updateSecretPayload)
	if parsingError := context.BodyParser(payload); parsingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	accountName := strings.Trim(payload.AccountName, " ")
	issuer := strings.Trim(payload.Issuer, " ")
	if accountName == "" || issuer == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	entryId := context.Params("id")
	userId := context.Locals("userId").(uint)
	result := database.Connection.Model(&models.Secrets{}).
		Where("entry_id = ? AND user_id = ?", entryId, userId).
		Updates(models.Secrets{Issuer: issuer, AccountName: accountName})
	if result.Error != nil {
		return result.Error
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
