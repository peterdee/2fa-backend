package models

type UserDevices struct {
	Generic
	DeviceName string
	LastUsed   uint
	UserID     uint
}
