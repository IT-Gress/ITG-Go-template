package utils

import "time"

// FormatTime formats a time.Time pointer to a string pointer in RFC3339 format.
func FormatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}

	formatted := t.Format(time.RFC3339)

	return &formatted
}
