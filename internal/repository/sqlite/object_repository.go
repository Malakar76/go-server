package sqlite

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"go-server/internal/domain"
	"go-server/internal/repository"

	"github.com/google/uuid"
)

type ObjectRepository struct {
	db *sql.DB
}

func NewObjectRepository(db *sql.DB) *ObjectRepository {
	return &ObjectRepository{db: db}
}

func (r *ObjectRepository) Create(ctx context.Context, name string, userId string) (*domain.Object, error) {
	o := &domain.Object{
		ID:        uuid.NewString(),
		Name:      name,
		UserId:    userId,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO objects (id, name, user_id, created_at) VALUES (?, ?, ?, ?)`,
		o.ID, o.Name, o.UserId, o.CreatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "FOREIGN KEY constraint failed") {
			return nil, repository.ErrFKConstraint
		}
		return nil, err
	}
	return o, nil
}

func (r *ObjectRepository) GetByID(ctx context.Context, id string) (*domain.Object, error) {
	var o domain.Object
	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, user_id, created_at FROM objects WHERE id = ?`,
		id,
	).Scan(&o.ID, &o.Name, &o.UserId, &o.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
