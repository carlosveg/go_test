package models

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	Data string `json:"datos"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}
