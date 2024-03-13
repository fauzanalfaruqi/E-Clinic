package dto

import "github.com/dgrijalva/jwt-go"

type JWTClams struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}