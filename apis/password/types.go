package password

type changePasswordPayload struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
