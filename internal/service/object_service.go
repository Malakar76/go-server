package service

import (
	"context"
	"database/sql"
	"errors"

	"go-server/internal/domain"
	"go-server/internal/repository"
)

type ObjectService struct {
	repo repository.ObjectRepository
}

func NewObjectService(repo repository.ObjectRepository) *ObjectService {
	return &ObjectService{repo: repo}
}

func (s *ObjectService) Create(ctx context.Context, name string, userId string) (*domain.Object, error) {
	o, err := s.repo.Create(ctx, name, userId)
	if err != nil {
		if errors.Is(err, repository.ErrFKConstraint) {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return o, nil
}

func (s *ObjectService) GetByID(ctx context.Context, id string) (*domain.Object, error) {
	o, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, err
	}
	return o, nil
}
