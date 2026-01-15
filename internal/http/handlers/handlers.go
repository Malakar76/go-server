package handlers

import "go-server/internal/service"

type Deps struct {
	UserSvc   *service.UserService
	ObjectSvc *service.ObjectService
}

type Handlers struct {
	userSvc   *service.UserService
	objectSvc *service.ObjectService
}

func New(d Deps) *Handlers {
	return &Handlers{
		userSvc:   d.UserSvc,
		objectSvc: d.ObjectSvc,
	}
}
