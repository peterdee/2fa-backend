package auth

type codeSignInPayload struct {
	ClientType string `json:"clientType"`
	Code       string `json:"code"`
}

type signInPayload struct {
	ClientType string `json:"clientType"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type signUpPayload struct {
	signInPayload
	RecoveryAnswer   string `json:"recoveryAnswer"`
	RecoveryQuestion string `json:"recoveryQuestion"`
}
