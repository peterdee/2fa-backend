package recovery

type checkLoginPayload struct {
	Login string `json:"login"`
}

type updatePasswordPayload struct {
	NewPassword    string `json:"newPassword"`
	RecoveryAnswer string `json:"recoveryAnswer"`
	UserID         uint   `json:"userId"`
}
