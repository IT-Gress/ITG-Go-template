package utils

import (
	"errors"

	"github.com/lib/pq"
)

const uniqueViolationCode = "23505"

// IsUniqueViolation checks if the error is a unique violation error from PostgreSQL.
func IsUniqueViolation(err error) bool {
	pqErr := &pq.Error{}
	if errors.As(err, &pqErr) {
		if pqErr.Code == uniqueViolationCode {
			return true
		}
	}

	return false
}
