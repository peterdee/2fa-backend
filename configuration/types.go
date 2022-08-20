package configuration

type AuthTypesStruct struct {
	HOTP string
	TOTP string
}

type ClientTypesStruct struct {
	Mobile string
	Web    string
}

type EnvironmentsStruct struct {
	Development string
	Heroku      string
	Production  string
}

type ResponseMessagesStruct struct {
	AccountSuspended       string
	InternalServerError    string
	InvalidAccountName     string
	InvalidClientType      string
	InvalidData            string
	InvalidEntryId         string
	InvalidIssuer          string
	InvalidLogin           string
	InvalidRecoveryAnswer  string
	InvalidToken           string
	LoginAlreadyInUse      string
	LoginIsTooLong         string
	MissingData            string
	MissingToken           string
	OK                     string
	OldPasswordIsInvalid   string
	PasswordContainsSpaces string
	PasswordIsTooShort     string
	SecretAlreadyDeleted   string
	SecretAlreadyExists    string
	Unauthorized           string
}
