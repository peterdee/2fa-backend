package utilities

import (
	"encoding/json"
	"time"

	"github.com/pascaldekloe/jwt"
)

type extraHeaders struct {
	Type string `json:"typ"`
}

func CreateToken(userId uint, clientType, tokenSecret string) (string, error) {
	var claims jwt.Claims

	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"id": userId, "client": clientType}

	var extra extraHeaders
	extra.Type = "JWT"
	extraString, _ := json.Marshal(extra)

	token, signError := claims.HMACSign(jwt.HS256, []byte(tokenSecret), extraString)
	return string(token), signError
}

func DecodeToken(token, tokenSecret string) {

}
