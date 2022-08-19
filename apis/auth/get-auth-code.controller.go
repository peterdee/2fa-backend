package auth

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/gohelpers"

	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func getAuthCodeController(context *fiber.Ctx) error {
	userId := context.Locals("userId").(uint)

	sourceString := fmt.Sprint(userId) + gohelpers.RandomString(32)
	codeHash, hashError := utilities.CreateHash(sourceString)
	if hashError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	code := strings.ToUpper(hex.EncodeToString([]byte(codeHash)))

	var existingRecord models.AuthCodes
	result := database.Connection.Where("user_id = ?", userId).Find(&existingRecord)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected > 0 {
		result = database.Connection.Model(&models.AuthCodes{}).
			Where("user_id = ?", userId).
			Update("code", code)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}
	}
	if result.RowsAffected == 0 {
		newAuthCode := models.AuthCodes{Code: code, UserID: userId}
		result = database.Connection.Create(&newAuthCode)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
		Data: fiber.Map{
			"authCode": code,
		},
	})
}
