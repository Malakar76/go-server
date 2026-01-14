package repository

import (
	"context"
	"go-server/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, name string) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
}
