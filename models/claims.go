package models

import "github.com/dgrijalva/jwt-go"

// custom claims
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}
