package password

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"
	"gorm.io/gorm"

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

	if gohelpers.IncludesString(strings.Split(newPassword, ""), " ") {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordContainsSpaces,
		)
	}
	if len(newPassword) < configuration.PASSWORD_MIN_LENGTH {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordIsTooShort,
		)
	}

	clientType := context.Locals("client").(string)
	userId := context.Locals("userId").(uint)
	var passwordRecord models.Passwords
	result := database.Connection.Where("user_id = ?", userId).Find(&passwordRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}
	var tokenSecretRecord models.TokenSecrets
	result = database.Connection.Where("user_id = ?", userId).Find(&tokenSecretRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	oldPasswordIsValid, compareError := utilities.CompareValueWithHash(
		oldPassword,
		passwordRecord.Hash,
	)
	if compareError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if !oldPasswordIsValid {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.OldPasswordIsInvalid,
		)
	}

	newPasswordHash, hashError := utilities.CreateHash(newPassword)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	newTokenSecret, secretHashError := utilities.CreateHash(utilities.CreateTokenKey(userId))
	if secretHashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	transactionError := database.Connection.Transaction(func(instance *gorm.DB) error {
		result = instance.Model(&models.Passwords{}).
			Where("id = ?", passwordRecord.ID).
			Update("hash", newPasswordHash)
		if result.Error != nil {
			return result.Error
		}
		result = instance.Model(&models.TokenSecrets{}).
			Where("id = ?", tokenSecretRecord.ID).
			Update("secret", newTokenSecret)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if transactionError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	token, signError := utilities.CreateToken(userId, clientType, newTokenSecret)
	if signError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Data: fiber.Map{
			"token": token,
		},
		Context: context,
	})
}
