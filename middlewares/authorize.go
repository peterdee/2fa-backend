package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func Authorize(context *fiber.Ctx) error {
	token := context.Get(fiber.HeaderAuthorization)
	if token == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.MissingToken,
		)
	}

	tokenClaims, decodeError := utilities.DecodeToken(token)
	if decodeError != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.InvalidToken,
		)
	}

	var tokenSecretRecord = models.TokenSecrets{UserID: tokenClaims.ID}
	result := database.Connection.Find(&tokenSecretRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	isValid := utilities.VerifyToken(token, tokenSecretRecord.Secret)
	if !isValid {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			configuration.RESPONSE_MESSAGES.Unauthorized,
		)
	}

	context.Locals("client", tokenClaims.ClientType)
	context.Locals("userId", tokenClaims.ID)
	return context.Next()
}
