package sqlite

import (
	"context"
	"database/sql"
	"time"

	"go-server/internal/domain"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, name string) (*domain.User, error) {
	u := &domain.User{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (id, name, created_at) VALUES (?, ?, ?)`,
		u.ID, u.Name, u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, created_at FROM users WHERE id = ?`,
		id,
	).Scan(&u.ID, &u.Name, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
