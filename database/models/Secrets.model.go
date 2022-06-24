package models

type Secrets struct {
	Generic
	AccountName string
	Algorithm   string
	AuthType    string
	Counter     uint
	Digits      uint
	EntryID     string
	Issuer      string
	Period      uint
	Secret      string
	UserID      uint
}
