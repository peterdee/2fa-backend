package models

type Passwords struct {
	Generic
	Hash   string `json:"hash"`
	UserID uint   `json:"userId"`
}
