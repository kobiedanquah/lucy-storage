package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/kobiedanquah/lucy/internal/services"
)

type ServiceHandler struct {
	users *services.UserService
}

var validate = validator.New()

func NewServiceHandler(us *services.UserService) *ServiceHandler {
	return &ServiceHandler{
		users: us,
	}
}
