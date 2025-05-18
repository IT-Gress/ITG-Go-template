package auth

import (
	"net/http"

	"github.com/it-gress/itg-go-template/internal/apierror"
	"golang.org/x/crypto/bcrypt"
)

// CreateHash creates a hash from the given password using bcrypt.
func CreateHash(password string) (string, *apierror.APIError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+hashSalt), bcrypt.DefaultCost)
	if err != nil {
		return "", apierror.New(http.StatusInternalServerError, "Failed to hash password", err)
	}

	return string(hash), nil
}

// CompareHash compares the given password with the stored hash.
func CompareHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+hashSalt)) != nil
}
