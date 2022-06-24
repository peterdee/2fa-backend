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
	InternalServerError string
	InvalidData         string
	LoginAlreadyInUse   string
	MissingData         string
	OK                  string
}
