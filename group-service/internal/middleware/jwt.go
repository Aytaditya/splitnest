package middleware

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("Aditya Aryan")

func ValidateToken(authHeader string) (int64, error) {
	if authHeader == "" {
		return 0, errors.New("missing authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		return 0, errors.New("invalid authorization format")
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("user id missing in token")
	}

	return int64(idFloat), nil
}
