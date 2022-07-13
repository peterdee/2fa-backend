package recovery

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

func updatePasswordController(context *fiber.Ctx) error {
	payload := new(updatePasswordPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	clientType := strings.Trim(payload.ClientType, " ")
	newPassword := strings.Trim(payload.NewPassword, " ")
	recoveryAnswer := strings.Trim(payload.RecoveryAnswer, " ")
	userId := payload.UserID
	if clientType == "" || newPassword == "" ||
		recoveryAnswer == "" || userId == 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	clients := gohelpers.ObjectValues(configuration.CLIENT_TYPES)
	if !gohelpers.IncludesString(clients, clientType) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidData,
		)
	}

	if len(newPassword) < configuration.PASSWORD_MIN_LENGTH {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordIsTooShort,
		)
	}
	if gohelpers.IncludesString(strings.Split(newPassword, ""), " ") {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordContainsSpaces,
		)
	}

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

	var passwordRecord models.Passwords
	result = database.Connection.Where("user_id = ?", user.ID).Find(&passwordRecord)
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
	result = database.Connection.Where("user_id = ?", user.ID).Find(&tokenSecretRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	isValid, comparisonError := utilities.CompareValueWithHash(
		recoveryAnswer,
		user.RecoveryAnswerHash,
	)
	if comparisonError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if !isValid {
		return fiber.NewError(
			fiber.StatusForbidden,
			configuration.RESPONSE_MESSAGES.InvalidRecoveryAnswer,
		)
	}

	newPasswordHash, hashError := utilities.CreateHash(newPassword)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	newTokenKey := utilities.CreateTokenKey(user.ID)
	newTokenSecret, hashError := utilities.CreateHash(newTokenKey)
	if hashError != nil {
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
		result = instance.Model(&models.Users{}).
			Where("id = ?", user.ID).
			Update("failed_sign_in_attempts", 0)
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
		Context: context,
		Data: fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":    user.ID,
				"login": user.Login,
			},
		},
	})
}
