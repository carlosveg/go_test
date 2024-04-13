package assets

import (
	"github.com/carlosveg/go_test/models"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secreto")

func GenerarTokenJWT(data string) (string, error) {
	claims := models.TokenClaims{
		Data: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
