package utilities

import (
	"fmt"

	"github.com/julyskies/gohelpers"
)

func CreateTokenKey(userId uint) string {
	return fmt.Sprint(userId) + "&" + gohelpers.RandomString(32)
}
