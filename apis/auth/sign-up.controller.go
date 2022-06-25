package auth

import (
	"fmt"
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

	passwordHash, hashError := utilities.CreateHash(password)
	if hashError != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	newPassword := models.Passwords{Hash: passwordHash, UserID: newUser.ID}
	result = database.Connection.Create(&newPassword)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	tokenKey := utilities.CreateTokenKey(newUser.ID)
	tokenSecret, secretError := utilities.CreateHash(tokenKey)
	if secretError != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}
	newTokenSecret := models.TokenSecrets{Secret: tokenSecret, UserID: newUser.ID}
	result = database.Connection.Create(&newTokenSecret)
	if result.Error != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	token, signError := utilities.CreateToken(newUser.ID, clientType, tokenSecret)
	if signError != nil {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Info:    configuration.RESPONSE_MESSAGES.InternalServerError,
			Status:  fiber.StatusInternalServerError,
		})
	}

	decoded, _ := utilities.DecodeToken(token)
	fmt.Println(decoded)

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
		Data: fiber.Map{
			"token": token,
			"user": fiber.Map{
				"id":    newUser.ID,
				"login": newUser.Login,
			},
		},
	})
}
