package account

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func updateRecoveryDataController(context *fiber.Ctx) error {
	payload := new(updateRecoveryDataPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	newRecoveryAnswer := strings.Trim(payload.NewRecoveryAnswer, " ")
	newRecoveryQuestion := strings.Trim(payload.NewRecoveryQuestion, " ")
	if newRecoveryAnswer == "" || newRecoveryQuestion == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	newRecoveryAnswerHash, hashError := utilities.CreateHash(newRecoveryAnswer)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	userId := context.Locals("userId").(uint)
	result := database.Connection.Model(&models.Users{}).
		Where("id = ?", userId).
		Updates(&models.Users{
			RecoveryAnswerHash: newRecoveryAnswerHash,
			RecoveryQuestion:   newRecoveryQuestion,
		})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
