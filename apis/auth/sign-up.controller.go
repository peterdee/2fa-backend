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

func signUpController(context *fiber.Ctx) error {
	payload := new(signUpPayload)
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
	recoveryAnswer := strings.Trim(payload.RecoveryAnswer, " ")
	recoveryQuestion := strings.Trim(payload.RecoveryQuestion, " ")
	if clientType == "" || login == "" || password == "" ||
		recoveryAnswer == "" || recoveryQuestion == "" {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.MissingData,
			Status:  fiber.StatusBadRequest,
		})
	}

	clients := gohelpers.ObjectValues(configuration.CLIENT_TYPES)
	if !gohelpers.IncludesString(clients, payload.ClientType) {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InvalidData,
			Status:  fiber.StatusBadRequest,
		})
	}

	var existingUsers []models.Users
	result := database.Connection.Where("login = ?", login).Find(&existingUsers)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	if result.RowsAffected > 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.LoginAlreadyInUse,
			Status:  fiber.StatusBadRequest,
		})
	}

	newUser := models.Users{Login: login}
	result = database.Connection.Create(&newUser)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	// TODO: create password hash, create token secret

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
