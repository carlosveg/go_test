package models

import "github.com/golang-jwt/jwt"

// Claims para generar el token
type TokenClaims struct {
	Data string `json:"datos"`
	jwt.StandardClaims
}

// Cuando se genere el token la respuesta tendrá la forma de este struct
type TokenResponse struct {
	Token string `json:"token"`
}
