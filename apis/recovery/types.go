package recovery

type checkLoginPayload struct {
	Login string `json:"login"`
}

type updatePasswordPayload struct {
	ClientType     string `json:"clientType"`
	NewPassword    string `json:"newPassword"`
	RecoveryAnswer string `json:"recoveryAnswer"`
	UserID         uint   `json:"userId"`
}
