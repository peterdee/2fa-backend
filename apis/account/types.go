package account

type updateRecoveryDataPayload struct {
	NewRecoveryAnswer   string `json:"newRecoveryAnswer"`
	NewRecoveryQuestion string `json:"newRecoveryQuestion"`
}
