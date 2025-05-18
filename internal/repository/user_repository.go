package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/it-gress/itg-go-template/internal/apierror"
	"github.com/it-gress/itg-go-template/models"
	"github.com/jmoiron/sqlx"
)

// UserRepository defines the methods for user-related database operations.
type UserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// FindAllUsers retrieves all users from the database.
func (r *UserRepository) FindAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	err := r.DB.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUserByUsername retrieves a user by their username.
func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.New(http.StatusNotFound, "User not found", nil)
		}

		return nil, err
	}

	return &user, nil
}

// FindPermissionsByUserID retrieves permissions for a user by their ID.
func (r *UserRepository) FindPermissionsByUserID(ctx context.Context, userID int) ([]*models.Permission, error) {
	query := `SELECT p.* FROM permissions p
	JOIN role_permissions rp ON p.id = rp.permission_id
	JOIN roles r ON rp.role_id = r.id
	JOIN users u ON u.role_id = r.id
	WHERE u.id = $1;`

	var permissions []*models.Permission

	err := r.DB.SelectContext(ctx, &permissions, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.New(http.StatusNotFound, "No permissions found for user", nil)
		}

		return nil, err
	}

	return permissions, nil
}

// UpdateLastLogin updates the last login timestamp for a user.
func (r *UserRepository) UpdateLastLogin(ctx context.Context, userID int) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE users SET last_login = NOW() WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
