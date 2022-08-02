package secrets

import (
	"github.com/gofiber/fiber/v2"

	"backend2fa/database"
	"backend2fa/database/models"
	"backend2fa/utilities"
)

func deleteAllSecretsController(context *fiber.Ctx) error {
	userId := context.Locals("userId").(uint)

	// find all of the existing Secrets
	var records []models.Secrets
	result := database.Connection.Where("user_id = ?", userId).Find(&records)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	if result.RowsAffected == 0 {
		return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
	}

	// get EntryId from all of the found Secrets
	var entryIds []string
	for _, entry := range records {
		entryIds = append(entryIds, entry.EntryID)
	}

	// create DeletedSecretID record for every Secret record
	for _, entryId := range entryIds {
		newDeletedSecretId := models.DeletedSecretIDs{
			EntryID: entryId,
			UserID:  userId,
		}
		result = database.Connection.Create(&newDeletedSecretId)
		if result.Error != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}
	}

	// delete Secret records
	result = database.Connection.Where("user_id = ?", userId).Delete(&models.Secrets{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponsePayloadStruct{Context: context})
}
