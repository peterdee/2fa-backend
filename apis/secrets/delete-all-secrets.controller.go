package secrets

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func deleteAllSecretsController(context *fiber.Ctx) error {
	userId := context.Locals("userId").(uint)

	var records []models.Secrets
	result := database.Connection.Where("user_id = ?", userId).Find(&records)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
	}

	var entryIds []string
	for _, entry := range records {
		entryIds = append(entryIds, entry.EntryID)
	}

	transactionError := database.Connection.Transaction(func(instance *gorm.DB) error {
		for _, entryId := range entryIds {
			newDeletedSecretId := models.DeletedSecretIDs{
				EntryID: entryId,
				UserID:  userId,
			}
			result = instance.Create(&newDeletedSecretId)
			if result.Error != nil {
				return result.Error
			}
		}

		result = instance.Where("user_id = ?", userId).Delete(&models.Secrets{})
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if transactionError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
