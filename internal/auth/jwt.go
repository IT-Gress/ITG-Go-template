package auth

import (
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTPayload represents the structure of the JWT token payload.
type JWTPayload struct {
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for the given user ID and scopes.
func GenerateToken(userID int, scopes []string) (string, error) {
	if jwtSecret == "" {
		slog.Error("JWT not initialized")
		return "", errors.New("JWT not initialized")
	}

	userIDStr := strconv.Itoa(userID)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTPayload{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userIDStr,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
		Scopes: scopes,
	}).SignedString([]byte(jwtSecret))

	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return token, nil
}

// ValidateToken checks if the token is valid and returns the claims if it is.
func ValidateToken(tokenString string) (*JWTPayload, error) {
	if jwtSecret == "" {
		slog.Error("JWT not initialized")
		return nil, errors.New("JWT not initialized")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*JWTPayload)
	if !ok {
		slog.Error("invalid token claims")
		return nil, errors.New("invalid token claims")
	}

	if !token.Valid {
		slog.Error("invalid token")
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
