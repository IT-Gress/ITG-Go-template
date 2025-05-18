package auth

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/it-gress/itg-go-template/internal/apierror"
)

// JWTPayload represents the structure of the JWT token payload.
type JWTPayload struct {
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for the given user ID and scopes.
func GenerateToken(userID int, scopes []string) (string, *apierror.APIError) {
	if jwtSecret == "" {
		slog.Error("JWT not initialized")
		return "", apierror.New(http.StatusInternalServerError, "JWT not initialized", nil)
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
		return "", apierror.New(http.StatusInternalServerError, "Failed to generate token", err)
	}

	return token, nil
}

// ValidateToken checks if the token is valid and returns the claims if it is.
func ValidateToken(tokenString string) (*JWTPayload, *apierror.APIError) {
	if jwtSecret == "" {
		return nil, apierror.New(http.StatusInternalServerError, "JWT not initialized", nil)
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apierror.New(http.StatusUnauthorized, "unexpected signing method", nil)
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, apierror.New(http.StatusUnauthorized, "error decoding token", err)
	}

	claims, ok := token.Claims.(*JWTPayload)
	if !ok {
		return nil, apierror.New(http.StatusUnauthorized, "invalid token claims", nil)
	}

	if !token.Valid {
		return nil, apierror.New(http.StatusUnauthorized, "invalid token", nil)
	}

	return claims, nil
}
