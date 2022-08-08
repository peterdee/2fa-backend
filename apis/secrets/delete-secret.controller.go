package secrets

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend2fa/configuration"
	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func deleteSecretController(context *fiber.Ctx) error {
	entryId := context.Params("id")
	userId := context.Locals("userId").(uint)

	result := database.Connection.
		Where("entry_id = ? AND user_id = ?", entryId, userId).
		Find(&models.DeletedSecretIDs{})
	if result.RowsAffected > 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.SecretAlreadyDeleted,
		)
	}

	var record models.Secrets
	result = database.Connection.Where(
		"entry_id = ? AND user_id = ?",
		entryId,
		userId,
	).Find(&record)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			configuration.RESPONSE_MESSAGES.InvalidEntryId,
		)
	}

	transactionError := database.Connection.Transaction(func(instance *gorm.DB) error {
		newDeletedSecretID := models.DeletedSecretIDs{
			EntryID: record.EntryID,
			UserID:  record.UserID,
		}
		result = instance.Create(&newDeletedSecretID)
		if result.Error != nil {
			return result.Error
		}
		result = instance.Where("entry_id = ?", record.EntryID).Delete(&models.Secrets{})
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}
		return nil
	})
	if transactionError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
