package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	log.Printf("generate JWT_SECRET: %s", os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	// return token.SignedString([]byte("rahasia"))
}
