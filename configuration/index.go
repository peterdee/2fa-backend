package configuration

var CLIENT_TYPES = ClientTypesStruct{
	Mobile: "mobile",
	Web:    "web",
}

var DEFAULT_PORT string = "2244"

var ENVS = EnvironmentsStruct{
	Development: "development",
	Heroku:      "heroku",
	Production:  "production",
}

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidData:         "INVALID_DATA",
	MissingData:         "MISSING_DATA",
	OK:                  "OK",
}
