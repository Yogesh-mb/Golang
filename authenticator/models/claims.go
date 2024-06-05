package models

import "github.com/dgrijalva/jwt-go"

type claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
