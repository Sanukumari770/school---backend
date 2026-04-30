package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate Token
func GenerateJWT(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

//  Verify Token (function)
func VerifyToken(tokenStr string) (*jwt.Token, error) {

	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}