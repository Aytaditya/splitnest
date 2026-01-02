package middleware

import (
	"time"

	"github.com/Aytaditya/splitnest-user-service/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Aditya Aryan")

func CreateToken(userId int64, username string, email string) (string, error) {
	claims := types.CustomClaims{
		ID:       userId,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day expiry
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
