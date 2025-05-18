package models

import (
	"time"

	"github.com/it-gress/itg-go-template/internal/entities"
)

// Role represents the relationship between users and roles.
type Role struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

// ToDTO converts a Role entity to a RoleDTO.
func (r *Role) ToDTO() *entities.RoleDTO {
	return &entities.RoleDTO{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
	}
}
