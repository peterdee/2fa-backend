package models

type AuthCodes struct {
	Generic
	Code   string `json:"code"`
	UserID uint   `json:"userId"`
}
