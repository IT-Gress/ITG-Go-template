package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/it-gress/itg-go-template/internal/apierror"
	"github.com/it-gress/itg-go-template/internal/models"
	"github.com/it-gress/itg-go-template/internal/utils"
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

// InsertUser inserts a new user into the database.
func (r *UserRepository) InsertUser(ctx context.Context, user *models.User) (*models.User, *apierror.APIError) {
	query := `
	INSERT INTO users (
		name, username, password_hash, email, role_id
	)
	VALUES (
		:name, :username, :password_hash, :email, :role_id
	)
	RETURNING *`

	stmt, err := r.DB.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, "Failed to prepare query", err)
	}

	err = stmt.GetContext(ctx, user, user)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return nil, apierror.New(http.StatusConflict, "User already exists", err)
		}

		return nil, apierror.New(http.StatusInternalServerError, "Failed to insert user", err)
	}

	return user, nil
}

// FindAllUsers retrieves all users from the database.
func (r *UserRepository) FindAllUsers(ctx context.Context) ([]*models.User, *apierror.APIError) {
	var users []*models.User

	err := r.DB.SelectContext(ctx, &users, "SELECT * FROM users")
	if err != nil {
		return nil, apierror.New(http.StatusInternalServerError, "Failed to query users", err)
	}

	return users, nil
}

// FindUserByID retrieves a user by their ID.
func (r *UserRepository) FindUserByID(ctx context.Context, userID int) (*models.User, *apierror.APIError) {
	var user models.User

	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.New(http.StatusNotFound, "User not found", err)
		}

		return nil, apierror.New(http.StatusInternalServerError, "Failed to query user", err)
	}

	return &user, nil
}

// FindUserByUsername retrieves a user by their username.
func (r *UserRepository) FindUserByUsername(ctx context.Context, username string) (*models.User, *apierror.APIError) {
	var user models.User

	err := r.DB.GetContext(ctx, &user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.New(http.StatusNotFound, "User not found", err)
		}

		return nil, apierror.New(http.StatusInternalServerError, "Failed to query user", err)
	}

	return &user, nil
}

// FindPermissionsByUserID retrieves permissions for a user by their ID.
func (r *UserRepository) FindPermissionsByUserID(
	ctx context.Context, userID int) ([]*models.Permission, *apierror.APIError) {
	query := `SELECT p.* FROM permissions p
	JOIN role_permissions rp ON p.id = rp.permission_id
	JOIN roles r ON rp.role_id = r.id
	JOIN users u ON u.role_id = r.id
	WHERE u.id = $1;`

	var permissions []*models.Permission

	err := r.DB.SelectContext(ctx, &permissions, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.New(http.StatusNotFound, "No permissions found for user", err)
		}

		return nil, apierror.New(http.StatusInternalServerError, "Failed to query permissions", err)
	}

	return permissions, nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, *apierror.APIError) {
	query := `UPDATE users
	SET name = :name, username = :username, email = :email, role_id = :role_id, password_hash = :password_hash
	, is_active = :is_active
	, last_login = :last_login
	, updated_at = NOW()
	WHERE id = :id
	RETURNING *`

	_, err := r.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		if utils.IsUniqueViolation(err) {
			return nil, apierror.New(http.StatusConflict, "User with this username or email already exists", err)
		}

		return nil, apierror.New(http.StatusInternalServerError, "Failed to update user", err)
	}

	return user, nil
}
