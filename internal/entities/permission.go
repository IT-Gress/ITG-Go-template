package entities

// PermissionDTO is a struct that represents a permission data transfer object.
type PermissionDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
}
