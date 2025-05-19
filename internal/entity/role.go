package entity

// RoleDTO is a struct that represents a role data transfer object.
type RoleDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
