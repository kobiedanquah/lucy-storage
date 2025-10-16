package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrDuplicateUser = errors.New("user with email already exists")
	ErrNotFound      = errors.New("entity not found")
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	ProfilePhoto string    `json:"profilePhoto"`
	CreatedAt    time.Time `json:"createdAt"`
	LastModified time.Time `json:"lastModified"`
	Verified     bool      `json:"emailVerified"`
}

type UserToken struct {
	Hash      string
	UserId    uuid.UUID
	ExpiresAt time.Time
	Scope     string
}

type UserStore interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	InsertToken(ctx context.Context, token *UserToken) error
	GetUserForToken(ctx context.Context, tokenHash, scope, email string) (User, error)
	DeleteToken(ctx context.Context, tokenHash, scope string) error
}
