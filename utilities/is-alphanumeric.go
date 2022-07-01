package utilities

import "regexp"

func IsAlphanumeric(value string) bool {
	var test = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	return test(value)
}
