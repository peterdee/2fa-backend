package utilities

import "github.com/alexedwards/argon2id"

func CompareValueWithHash(value, hash string) (bool, error) {
	match, matchError := argon2id.ComparePasswordAndHash(value, hash)
	return match, matchError
}

func CreateHash(value string) (string, error) {
	hash, hashError := argon2id.CreateHash(value, argon2id.DefaultParams)
	return hash, hashError
}
