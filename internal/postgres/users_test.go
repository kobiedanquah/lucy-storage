package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/primekobie/lucy/internal/models"
	"github.com/primekobie/lucy/internal/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := os.Getenv("TEST_DB_URL")
	if connStr == "" {
		connStr = "postgres://userone:pa88word@localhost:5432/lucy_test?sslmode=disable"
	}

	pool, err := sql.Open("pgx", connStr)
	require.NoError(t, err)
	t.Cleanup(func() { pool.Close() })
	return pool
}

func createTestUser(name, email string) *models.User {
	return &models.User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: []byte("hashedpassword"),
		CreatedAt:    time.Now().UTC(),
		LastModified: time.Now().UTC(),
		Verified:     true,
	}
}

func generateTestEmail() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return fmt.Sprintf("%s@example.com", string(b))
}

func TestUserStore_CreateUser(t *testing.T) {
	pool := setupTestDB(t)
	store := postgres.NewUserStore(pool)
	ctx := context.Background()

	email := generateTestEmail()

	tests := []struct {
		name    string
		user    *models.User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    createTestUser("Test User", email),
			wantErr: false,
		},
		{
			name:    "duplicate email",
			user:    createTestUser("Another User", email),
			wantErr: true,
		},
		{
			name: "empty name",
			user: &models.User{
				ID:           uuid.New(),
				Email:        generateTestEmail(),
				PasswordHash: []byte("hashedpassword"),
				CreatedAt:    time.Now().UTC(),
				LastModified: time.Now().UTC(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Create(ctx, tt.user)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			got, err := store.GetByID(ctx, tt.user.ID)
			require.NoError(t, err)
			assert.Equal(t, tt.user.ID, got.ID)
			assert.Equal(t, tt.user.Name, got.Name)
			assert.Equal(t, tt.user.Email, got.Email)
		})
	}
}

func TestUserStore_GetByID(t *testing.T) {
	pool := setupTestDB(t)
	store := postgres.NewUserStore(pool)
	ctx := context.Background()

	user := createTestUser("Get Test", generateTestEmail())
	require.NoError(t, store.Create(ctx, user))

	tests := []struct {
		name    string
		id      uuid.UUID
		wantErr bool
	}{
		{
			name:    "existing user",
			id:      user.ID,
			wantErr: false,
		},
		{
			name:    "non-existent user",
			id:      uuid.New(),
			wantErr: true,
		},
		{
			name:    "invalid ID",
			id:      uuid.Nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetByID(ctx, tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, user.ID, got.ID)
		})
	}
}
func TestUserStore_GetByEmail(t *testing.T) {
	pool := setupTestDB(t)
	store := postgres.NewUserStore(pool)
	ctx := context.Background()

	user := createTestUser("Get Test", generateTestEmail())
	require.NoError(t, store.Create(ctx, user))

	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "existing user",
			email:   user.Email,
			wantErr: false,
		},
		{
			name:    "non-existent user",
			email:   generateTestEmail(),
			wantErr: true,
		},
		{
			name:    "invalid ID",
			email:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetByEmail(ctx, tt.email)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, user.ID, got.ID)
		})
	}
}

func TestUserStore_UpdateUser(t *testing.T) {
	pool := setupTestDB(t)
	store := postgres.NewUserStore(pool)
	ctx := context.Background()

	user := createTestUser("Update Test", generateTestEmail())
	require.NoError(t, store.Create(ctx, user))

	tests := []struct {
		name    string
		user    *models.User
		updates func(*models.User)
		wantErr bool
	}{
		{
			name: "valid update",
			user: user,
			updates: func(u *models.User) {
				u.Name = "Updated Name"
				u.Email = generateTestEmail()
			},
			wantErr: false,
		},
		{
			name: "non-existent user",
			user: createTestUser("Non-existent", "Kwame@email.com"),
			updates: func(u *models.User) {
				u.Name = "Updated Name"
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.updates != nil {
				tt.updates(tt.user)
			}
			tt.user.LastModified = time.Now().UTC()

			err := store.Update(ctx, tt.user)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			got, err := store.GetByID(ctx, tt.user.ID)
			require.NoError(t, err)
			assert.Equal(t, tt.user.Name, got.Name)
			assert.Equal(t, tt.user.Email, got.Email)
		})
	}
}

func TestUserStore_DeleteUser(t *testing.T) {
	pool := setupTestDB(t)
	store := postgres.NewUserStore(pool)
	ctx := context.Background()

	user := createTestUser("Delete Test", generateTestEmail())
	require.NoError(t, store.Create(ctx, user))

	tests := []struct {
		name    string
		id      uuid.UUID
		wantErr bool
	}{
		{
			name:    "existing user",
			id:      user.ID,
			wantErr: false,
		},
		{
			name:    "non-existent user",
			id:      uuid.New(),
			wantErr: true,
		},
		{
			name:    "invalid ID",
			id:      uuid.Nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Delete(ctx, tt.id)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			assert.NoError(t, err)

			_, err = store.GetByID(ctx, tt.id)
			assert.Error(t, err)
		})
	}
}
