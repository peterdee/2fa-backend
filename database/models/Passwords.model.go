package models

type Passwords struct {
	Generic
	Hash   string
	UserID uint
}
