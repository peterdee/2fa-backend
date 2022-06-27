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
	InternalServerError:  "INTERNAL_SERVER_ERROR",
	InvalidData:          "INVALID_DATA",
	InvalidToken:         "INVALID_TOKEN",
	LoginAlreadyInUse:    "LOGIN_ALREADY_IN_USE",
	MissingData:          "MISSING_DATA",
	MissingToken:         "MISSING_TOKEN",
	OK:                   "OK",
	OldPasswordIsInvalid: "OLD_PASSWORD_IS_INVALID",
	Unauthorized:         "UNAUTHORIZED",
}
