package models

import (
	"time"

	"github.com/it-gress/itg-go-template/internal/entities"
)

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

// ToDTO converts a User entity to a UserDTO.
func (u *User) ToDTO() *entities.UserDTO {
	return &entities.UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		RoleID:    u.RoleID,
		LastLogin: u.LastLogin.Format(time.RFC3339),
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
}
