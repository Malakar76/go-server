package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"go-server/internal/domain"
	"go-server/internal/repository"
)

type fakeUserRepo struct {
	getFn    func(ctx context.Context, id string) (*domain.User, error)
	createFn func(ctx context.Context, name string) (*domain.User, error)
}

func (f fakeUserRepo) Create(ctx context.Context, name string) (*domain.User, error) {
	return f.createFn(ctx, name)
}
func (f fakeUserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return f.getFn(ctx, id)
}

// compile-time check que fakeUserRepo implémente l’interface
var _ repository.UserRepository = (*fakeUserRepo)(nil)

func TestUserService_GetByID_NotFound(t *testing.T) {
	svc := NewUserService(fakeUserRepo{
		getFn: func(ctx context.Context, id string) (*domain.User, error) {
			return nil, sql.ErrNoRows
		},
		createFn: func(ctx context.Context, name string) (*domain.User, error) {
			t.Fatal("Create should not be called")
			return nil, nil
		},
	})

	_, err := svc.GetByID(context.Background(), "123")
	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got: %v", err)
	}
}

func TestUserService_Create_Validation(t *testing.T) {
	svc := NewUserService(fakeUserRepo{
		createFn: func(ctx context.Context, name string) (*domain.User, error) {
			t.Fatal("repo.Create should not be called for invalid name")
			return nil, nil
		},
		getFn: func(ctx context.Context, id string) (*domain.User, error) { return nil, nil },
	})

	_, err := svc.Create(context.Background(), "a")
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
}
