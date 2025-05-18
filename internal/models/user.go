package models

import (
	"time"

	"github.com/it-gress/itg-go-template/internal/entities"
	"github.com/it-gress/itg-go-template/internal/utils"
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
	UpdatedAt    time.Time  `db:"updated_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

// ToDTO converts a User entity to a UserDTO.
func (u *User) ToDTO() *entities.UserDTO {
	return &entities.UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		RoleID:    u.RoleID,
		LastLogin: utils.FormatTime(u.LastLogin),
		IsActive:  u.IsActive,
		CreatedAt: *utils.FormatTime(&u.CreatedAt),
		UpdatedAt: *utils.FormatTime(&u.UpdatedAt),
	}
}
