package models

type DeletedSecretIDs struct {
	Generic
	EntryID string `json:"entryId"`
	UserID  uint   `json:"userId"`
}
