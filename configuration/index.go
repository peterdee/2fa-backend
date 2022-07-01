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

var LOGIN_MAX_LENGTH int = 16

var PASSWORD_MIN_LENGTH int = 8

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	InternalServerError:    "INTERNAL_SERVER_ERROR",
	InvalidData:            "INVALID_DATA",
	InvalidToken:           "INVALID_TOKEN",
	LoginAlreadyInUse:      "LOGIN_ALREADY_IN_USE",
	LoginContainsSpaces:    "LOGIN_CONTAINS_SPACES",
	LoginIsTooLong:         "LOGIN_IS_TOO_LONG",
	MissingData:            "MISSING_DATA",
	MissingToken:           "MISSING_TOKEN",
	OK:                     "OK",
	OldPasswordIsInvalid:   "OLD_PASSWORD_IS_INVALID",
	PasswordContainsSpaces: "PASSWORD_CONTAINS_SPACES",
	PasswordIsTooShort:     "PASSWORD_IS_TOO_SHORT",
	Unauthorized:           "UNAUTHORIZED",
}
