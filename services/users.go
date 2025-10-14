package services

import "github.com/primekobie/lucy/models"

type UserService struct {
	store models.UserStore
}

func NewUserSerive(store models.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}
