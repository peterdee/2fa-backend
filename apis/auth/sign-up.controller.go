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
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	clientType := strings.Trim(payload.ClientType, " ")
	login := strings.ToLower(strings.Trim(payload.Login, " "))
	password := strings.Trim(payload.Password, " ")
	recoveryAnswer := strings.Trim(payload.RecoveryAnswer, " ")
	recoveryQuestion := strings.Trim(payload.RecoveryQuestion, " ")
	if clientType == "" || login == "" || password == "" ||
		recoveryAnswer == "" || recoveryQuestion == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}
	if len(login) > configuration.LOGIN_MAX_LENGTH {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.LoginIsTooLong,
		)
	}
	if len(password) < configuration.PASSWORD_MIN_LENGTH {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordIsTooShort,
		)
	}
	if !utilities.IsAlphanumeric(login) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidLogin,
		)
	}
	if gohelpers.IncludesString(strings.Split(password, ""), " ") {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.PasswordContainsSpaces,
		)
	}

	clients := gohelpers.ObjectValues(configuration.CLIENT_TYPES)
	if !gohelpers.IncludesString(clients, clientType) {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidData,
		)
	}

	var existingUsers []models.Users
	result := database.Connection.Where("login = ?", login).Find(&existingUsers)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected > 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.LoginAlreadyInUse,
		)
	}

	recoveryAnswerHash, hashError := utilities.CreateHash(recoveryAnswer)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	newUser := models.Users{
		Login:              login,
		RecoveryAnswerHash: recoveryAnswerHash,
		RecoveryQuestion:   recoveryQuestion,
	}
	result = database.Connection.Create(&newUser)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	passwordHash, hashError := utilities.CreateHash(password)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	newPassword := models.Passwords{Hash: passwordHash, UserID: newUser.ID}
	result = database.Connection.Create(&newPassword)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	tokenKey := utilities.CreateTokenKey(newUser.ID)
	tokenSecret, secretError := utilities.CreateHash(tokenKey)
	if secretError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	newTokenSecret := models.TokenSecrets{Secret: tokenSecret, UserID: newUser.ID}
	result = database.Connection.Create(&newTokenSecret)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	token, signError := utilities.CreateToken(newUser.ID, clientType, tokenSecret)
	if signError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

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
