package models

import (
	"time"

	"github.com/it-gress/itg-go-template/internal/entities"
)

// Permission represents the relationship between users and roles.
type Permission struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
}

// ToDTO converts a Permission entity to a PermissionDTO.
func (p *Permission) ToDTO() *entities.PermissionDTO {
	return &entities.PermissionDTO{
		ID:    p.ID,
		Name:  p.Name,
		Value: p.Value,
	}
}
