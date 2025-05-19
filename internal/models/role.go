package models

import (
	"time"

	"github.com/it-gress/itg-go-template/internal/entity"
)

// Role represents the relationship between users and roles.
type Role struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

// ToDTO converts a Role entity to a RoleDTO.
func (r *Role) ToDTO() *entity.RoleDTO {
	return &entity.RoleDTO{
		ID:        r.ID,
		Name:      r.Name,
		UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
	}
}
