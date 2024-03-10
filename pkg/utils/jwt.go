package utils

import (
	"avengers-clinic/model/dto"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey = []byte("rahasia banget")
	method = jwt.SigningMethodHS256
)

func GenerateJWT(id, username, role string) (string, error) {
	claims := dto.JWTClams{
		ID: id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "kelompok-3",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(method, claims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	claims := &dto.JWTClams{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return token, err
}