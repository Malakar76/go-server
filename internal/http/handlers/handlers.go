package handlers

import "go-server/internal/service"

type Handlers struct {
	userSvc *service.UserService
}

func New(userSvc *service.UserService) *Handlers {
	return &Handlers{userSvc: userSvc}
}
