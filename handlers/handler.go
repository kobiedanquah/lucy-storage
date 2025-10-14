package handlers

import "github.com/primekobie/lucy/services"

type ServiceHandler struct {
	users *services.UserService
}

func NewServiceHandler(us *services.UserService) *ServiceHandler {
	return &ServiceHandler{
		users: us,
	}
}
