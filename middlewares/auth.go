package middlewares

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
// var jwtSecret = []byte("rahasia")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("verify : ", os.Getenv("JWT_SECRET"))
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				return nil, errors.New("JWT_SECRET is empty")
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}
