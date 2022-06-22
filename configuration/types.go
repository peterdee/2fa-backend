package configuration

type ClientTypesStruct struct {
	Mobile string
	Web    string
}

type ResponseMessagesStruct struct {
	InternalServerError string
	InvalidData         string
	MissingData         string
	OK                  string
}
