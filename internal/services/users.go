package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/primekobie/lucy/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store models.UserStore
}

func NewUserService(store models.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (us *UserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		return nil, ErrFailedOperation
	}
	now := time.Now().UTC()
	user := &models.User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    now,
		LastModified: now,
		Verified:     false,
	}

	if err := us.store.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) UpdateUser(ctx context.Context, userData map[string]any) (*models.User, error) {
	id, ok := userData["id"]
	if !ok {
		return nil, errors.New("user id not found")
	}

	user, err := us.store.GetByID(ctx, id.(uuid.UUID))
	if err != nil {
		return nil, err
	}

	name, ok := userData["name"]
	if ok {
		user.Name = name.(string)
	}

	profilePhoto, ok := userData["profilePhoto"]
	if ok {
		user.ProfilePhoto = profilePhoto.(string)
	}

	password, ok := userData["password"]
	if ok {
		if len(password.(string)) < 8 || len(password.(string)) > 20 {
			return nil, ErrInvalidPassword
		}
		err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password.(string)))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
				if err != nil {
					return nil, ErrFailedOperation
				}

				user.PasswordHash = hash
			} else {
				slog.Error("failed to compare password and hash", "error", err.Error())
				return nil, ErrFailedOperation
			}
		}

	}

	user.LastModified = time.Now().UTC()

	err = us.store.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FetchUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.store.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.store.Delete(ctx, id)
}
