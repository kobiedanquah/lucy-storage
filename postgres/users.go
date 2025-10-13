package postgres

import (
	"database/sql"

	"github.com/primekobie/lucy/models"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) models.UserStore{
	return UserStore{
		db: db,
	}
}