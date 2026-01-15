package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-server/internal/domain"
	"go-server/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, name string) (*domain.User, error) {
	if len(name) < 2 {
		return nil, fmt.Errorf("name too short")
	}
	return s.repo.Create(ctx, name)
}

func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}
