package jwt

import (
	"errors"
	"github.com/Fabese/project1/models"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var (
	Email     string
	IDUsuario string
)

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("Invalid Token Format")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
	}
	if !tkn.Valid {
		return &claims, false, "", errors.New("Invalid Token")
	}
	return &claims, false, "", nil
}
