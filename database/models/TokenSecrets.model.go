package models

type TokenSecrets struct {
	Generic
	Secret string `json:"secret"`
	UserID uint   `json:"userId"`
}
