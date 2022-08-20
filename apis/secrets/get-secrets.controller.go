package secrets

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func getSecretsController(context *fiber.Ctx) error {
	clientType := context.Locals("client").(string)
	userId := context.Locals("userId").(uint)

	var secrets []models.Secrets
	result := database.Connection.Where("user_id = ?", userId).Find(&secrets)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if clientType == configuration.CLIENT_TYPES.Web {
		return utilities.Response(utilities.ResponsePayloadStruct{
			Context: context,
			Data: fiber.Map{
				"secrets": secrets,
			},
		})
	}

	var deletedIds []models.DeletedSecretIDs
	result = database.Connection.Where("user_id = ?", userId).Find(&deletedIds)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{
		Context: context,
		Data: fiber.Map{
			"deletedSecretIds": deletedIds,
			"secrets":          secrets,
		},
	})
}
