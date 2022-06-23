package models

type Passwords struct {
	Generic
	Hash  string
	Users Users
}
