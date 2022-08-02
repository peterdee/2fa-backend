package configuration

var AUTH_TYPES = AuthTypesStruct{
	HOTP: "hotp",
	TOTP: "totp",
}

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

var MAX_FAILED_SIGN_IN_ATTEMPTS int = 10

var PASSWORD_MIN_LENGTH int = 8

var RESPONSE_MESSAGES = ResponseMessagesStruct{
	AccountSuspended:       "ACCOUNT_SUSPENDED",
	InternalServerError:    "INTERNAL_SERVER_ERROR",
	InvalidData:            "INVALID_DATA",
	InvalidEntryId:         "INVALID_ENTRY_ID",
	InvalidLogin:           "INVALID_LOGIN",
	InvalidRecoveryAnswer:  "INVALID_RECOVERY_ANSWER",
	InvalidToken:           "INVALID_TOKEN",
	LoginAlreadyInUse:      "LOGIN_ALREADY_IN_USE",
	LoginIsTooLong:         "LOGIN_IS_TOO_LONG",
	MissingData:            "MISSING_DATA",
	MissingToken:           "MISSING_TOKEN",
	OK:                     "OK",
	OldPasswordIsInvalid:   "OLD_PASSWORD_IS_INVALID",
	PasswordContainsSpaces: "PASSWORD_CONTAINS_SPACES",
	PasswordIsTooShort:     "PASSWORD_IS_TOO_SHORT",
	SecretAlreadyExists:    "SECRET_ALREADY_EXISTS",
	Unauthorized:           "UNAUTHORIZED",
}
