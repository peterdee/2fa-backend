package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func signInController(context *fiber.Ctx) error {
	payload := new(signInPayload)
	if err := context.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	clientType := strings.Trim(payload.ClientType, " ")
	login := strings.ToLower(strings.Trim(payload.Login, " "))
	password := strings.Trim(payload.Password, " ")
	if clientType == "" || login == "" || password == "" {
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
	if user.FailedSignInAttempts >= configuration.MAX_FAILED_SIGN_IN_ATTEMPTS {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.AccountSuspended,
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

	isValid, comparisonError := utilities.CompareValueWithHash(password, passwordRecord.Hash)
	if comparisonError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if !isValid {
		result = database.Connection.Model(&models.Users{}).
			Where("id = ?", user.ID).
			Update("failed_sign_in_attempts", user.FailedSignInAttempts+1)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}

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

	token, signError := utilities.CreateToken(user.ID, clientType, tokenSecretRecord.Secret)
	if signError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	result = database.Connection.Model(&models.Users{}).
		Where("id = ?", user.ID).
		Update("failed_sign_in_attempts", 0)
	if result.Error != nil {
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
