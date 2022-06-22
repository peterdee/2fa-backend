package auth

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (payload *signInPayload) validateSignInPayload() error {
	return validation.ValidateStruct(
		&payload,
		validation.Field(&payload.Login, validation.Required, validation.Length(5, 64)),
		validation.Field(&payload.Password, validation.Required),
	)
}
