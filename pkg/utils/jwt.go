package utils

import (
	"avengers-clinic/model/dto"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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

func GetJWT(c *gin.Context) *dto.JWTClams {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
	token, _ := VerifyJWT(tokenString)
	claims := token.Claims.(*dto.JWTClams)
	return claims
}