package secrets

type addSecretPayload struct {
	Algorithm   string `json:"algorithm"`
	AccountName string `json:"accountName"`
	AuthType    string `json:"authType"`
	Counter     uint   `json:"counter"`
	Digits      uint   `json:"digits"`
	EntryID     string `json:"entryId"`
	Issuer      string `json:"issuer"`
	Period      uint   `json:"period"`
	ScannedAt   int64  `json:"scannedAt"`
	Secret      string `json:"secret"`
}
