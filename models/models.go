package models

import "time"

// User represents a user in the system.
type User struct {
	ID           int        `db:"id"`
	Name         string     `db:"name"`
	Username     string     `db:"username"`
	Email        string     `db:"email"`
	PasswordHash string     `db:"password_hash"`
	LastLogin    *time.Time `db:"last_login"`
	IsActive     bool       `db:"is_active"`
	RoleID       int        `db:"role_id"`
	UpdatedAt    *time.Time `db:"updated_at"`
	CreatedAt    *time.Time `db:"created_at"`
}

// Role represents the relationship between users and roles.
type Role struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
}

// Permission represents the relationship between users and roles.
type Permission struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Value     string     `db:"value"`
	CreatedAt *time.Time `db:"created_at"`
}
