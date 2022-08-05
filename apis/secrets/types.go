package secrets

type addSecretPayload struct {
	AccountName string `json:"accountName"`
	Algorithm   string `json:"algorithm"`
	AuthType    string `json:"authType"`
	Counter     uint   `json:"counter"`
	Digits      uint   `json:"digits"`
	EntryID     string `json:"entryId"`
	Issuer      string `json:"issuer"`
	Period      uint   `json:"period"`
	ScannedAt   int64  `json:"scannedAt"`
	Secret      string `json:"secret"`
}

type updateSecretPayload struct {
	AccountName string `json:"accountName"`
	Issuer      string `json:"issuer"`
}
