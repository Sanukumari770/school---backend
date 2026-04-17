package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secure JWT (best to import os )
var SECRET_KEY = []byte(os.Getenv("JWT_SECRET"))



func GenerateToken(userId string, role string) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // token expire in 24 hours for security //
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(SECRET_KEY)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
}