package auth

type signInPayload struct {
	ClientType string `json:"clientType"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}
