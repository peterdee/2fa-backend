package password

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
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
			fiber.StatusInternalServerError,
			configuration.RESPONSE_MESSAGES.MissingData,
		)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
	})
}
