package utilities

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/pascaldekloe/jwt"
)

type extraHeaders struct {
	Type string `json:"typ"`
}

type TokenClaims struct {
	ClientType string          `json:"client"`
	ID         uint            `json:"id"`
	Issued     jwt.NumericTime `json:"iat"`
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

func DecodeToken(token string) (TokenClaims, error) {
	invalidToken := errors.New("invalid token")
	tokenClaims := TokenClaims{}
	claims, parseError := jwt.ParseWithoutCheck([]byte(token))
	if parseError != nil {
		return tokenClaims, parseError
	}

	if claims.Issued == nil {
		return tokenClaims, invalidToken
	}
	tokenClaims.Issued = *claims.Issued
	customClaimsString, _ := json.Marshal(claims.Set)
	parseError = json.Unmarshal(customClaimsString, &tokenClaims)
	if parseError != nil {
		return tokenClaims, parseError
	}
	if tokenClaims.ID == 0 || tokenClaims.ClientType == "" {
		return tokenClaims, invalidToken
	}

	return tokenClaims, nil
}

func VerifyToken(token, tokenSecret string) bool {
	_, verificationError := jwt.HMACCheck([]byte(token), []byte(tokenSecret))
	return verificationError == nil
}
