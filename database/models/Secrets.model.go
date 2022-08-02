package models

type Secrets struct {
	Generic
	AccountName    string
	Algorithm      string
	AuthType       string
	Counter        uint
	Digits         uint
	EntryID        string
	Issuer         string
	Period         uint
	ScannedAt      int64
	Secret         string
	SynchronizedAt int64
	UserID         uint
}
