package configuration

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
	InternalServerError  string
	InvalidData          string
	InvalidToken         string
	LoginAlreadyInUse    string
	LoginIsTooLong       string
	MissingData          string
	MissingToken         string
	OK                   string
	OldPasswordIsInvalid string
	PasswordIsTooShort   string
	Unauthorized         string
}
