package secrets

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func addSecretController(context *fiber.Ctx) error {
	payload := new(addSecretPayload)
	if parsingError := context.BodyParser(payload); parsingError != nil {
		if strings.Contains(parsingError.Error(), "converting") {
			return fiber.NewError(
				fiber.StatusBadRequest,
				configuration.RESPONSE_MESSAGES.InvalidData,
			)
		}
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	accountName := strings.ToLower(strings.Trim(payload.AccountName, " "))
	algorithm := strings.Trim(payload.Algorithm, " ")
	authType := strings.Trim(payload.AuthType, " ")
	counter := payload.Counter
	digits := payload.Digits
	entryId := strings.Trim(payload.EntryID, " ")
	issuer := strings.Trim(payload.Issuer, " ")
	period := payload.Period
	scannedAt := payload.ScannedAt
	secret := strings.Trim(payload.Secret, " ")

	if algorithm == "" || accountName == "" || authType == "" ||
		digits == 0 || entryId == "" || issuer == "" ||
		period == 0 || scannedAt == 0 || secret == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	authTypes := gohelpers.ObjectValues(configuration.AUTH_TYPES)
	if !gohelpers.IncludesString(authTypes, authType) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidData,
		)
	}

	var existingRecord models.Secrets
	result := database.Connection.Where("entry_id = ?", entryId).Find(&existingRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected > 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.SecretAlreadyExists,
		)
	}

	userId := context.Locals("userId").(uint)
	newSecret := models.Secrets{
		AccountName:    accountName,
		Algorithm:      algorithm,
		AuthType:       authType,
		Counter:        counter,
		Digits:         digits,
		EntryID:        entryId,
		Issuer:         issuer,
		Period:         period,
		ScannedAt:      scannedAt,
		Secret:         secret,
		SynchronizedAt: gohelpers.MakeTimestamp(),
		UserID:         userId,
	}
	result = database.Connection.Create(&newSecret)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
