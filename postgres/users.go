package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/primekobie/lucy/models"
)

type UserStore struct {
	conn *sql.DB
}

func NewUserStore(db *sql.DB) models.UserStore {
	return UserStore{
		conn: db,
	}
}

// Create implements models.UserStore.
func (u UserStore) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (id, name, email, password_hash, profile_photo, created_at, last_modified, verified)
		VALUES ($1, NULLIF($2,''), $3, $4, $5, $6, $7, $8);`

	_, err := u.conn.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.ProfilePhoto,
		user.CreatedAt,
		user.LastModified,
		user.Verified,
	)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return models.ErrDuplicateUser
		}
		slog.Error("failed to insert user", "error", err)
		return err
	}
	return nil
}

// Delete implements models.UserStore.
func (u UserStore) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1;`

	result, err := u.conn.ExecContext(ctx, query, id)
	if err != nil {
		slog.Error("failed delete user", "error", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.ErrNotFound
	}
	return nil
}

// GetByEmail implements models.UserStore.
func (u UserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, profile_photo, created_at, last_modified, verified 
		FROM users 
		WHERE email = $1;`

	var user models.User
	err := u.conn.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.ProfilePhoto,
		&user.CreatedAt,
		&user.LastModified,
		&user.Verified,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, models.ErrNotFound
	}

	return &user, nil
}

// GetByID implements models.UserStore.
func (u UserStore) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, name, email, password_hash, profile_photo, created_at, last_modified, verified 
		FROM users 
		WHERE id = $1;`

	var user models.User
	err := u.conn.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.ProfilePhoto,
		&user.CreatedAt,
		&user.LastModified,
		&user.Verified,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, models.ErrNotFound
	}

	return &user, nil
}

// Update implements models.UserStore.
func (u UserStore) Update(ctx context.Context, user *models.User) error {
		query := `
		UPDATE users 
		SET name = $1, email = $2, password_hash = $3, profile_photo = $4, last_modified = $5, verified = $6
		WHERE id = $7;`

	result, err := u.conn.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.ProfilePhoto,
		user.LastModified,
		user.Verified,
		user.ID,
	)
	if err != nil {
		slog.Error("failed update user", "error", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.ErrNotFound
	}

	return nil
}
