package repository

import (
	"context"
	"go-server/internal/domain"
)

type ObjectRepository interface {
	Create(ctx context.Context, name string, userId string) (*domain.Object, error)
	GetByID(ctx context.Context, id string) (*domain.Object, error)
}
