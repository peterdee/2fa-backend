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
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	clientType := strings.Trim(payload.ClientType, " ")
	login := strings.Trim(payload.Login, " ")
	password := strings.Trim(payload.Password, " ")
	if clientType == "" || login == "" || password == "" {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.MissingData,
			Status:  fiber.StatusBadRequest,
		})
	}

	clients := gohelpers.ObjectValues(configuration.CLIENT_TYPES)
	if !gohelpers.IncludesString(clients, clientType) {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InvalidData,
			Status:  fiber.StatusBadRequest,
		})
	}

	var user models.Users
	result := database.Connection.Where("login = ?", login).Find(&user)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	if result.RowsAffected == 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.Unauthorized,
			Status:  fiber.StatusUnauthorized,
		})
	}

	var passwordRecord models.Passwords
	result = database.Connection.Where("user_id = ?", user.ID).Find(&passwordRecord)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	if result.RowsAffected == 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.Unauthorized,
			Status:  fiber.StatusUnauthorized,
		})
	}

	var tokenSecretRecord models.TokenSecrets
	result = database.Connection.Where("user_id = ?", user.ID).Find(&tokenSecretRecord)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	if result.RowsAffected == 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.Unauthorized,
			Status:  fiber.StatusUnauthorized,
		})
	}

	token, signError := utilities.CreateToken(user.ID, clientType, tokenSecretRecord.Secret)
	if signError != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Data: fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":    user.ID,
				"login": user.Login,
			},
		},
		Context: context,
	})
}
