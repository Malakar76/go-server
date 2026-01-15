package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-server/internal/domain"
	"go-server/internal/service"

	"github.com/gin-gonic/gin"
)

// Fake service minimal
type fakeUserService struct {
	getFn    func(ctx context.Context, id string) (*domain.User, error)
	createFn func(ctx context.Context, name string) (*domain.User, error)
}

func (f fakeUserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return f.getFn(ctx, id)
}
func (f fakeUserService) Create(ctx context.Context, name string) (*domain.User, error) {
	return f.createFn(ctx, name)
}

func TestGetUser_404(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := fakeRepoNotFound{}
	userSvc := service.NewUserService(repo)
	h := New(Deps{
		UserSvc:   userSvc,
		ObjectSvc: nil,
	})

	r := gin.New()
	r.GET("/v1/users/:id", h.GetUser)

	req := httptest.NewRequest(http.MethodGet, "/v1/users/123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d body=%s", w.Code, w.Body.String())
	}
}

type fakeRepoNotFound struct{}

func (fakeRepoNotFound) Create(ctx context.Context, name string) (*domain.User, error) {
	return nil, nil
}

func (fakeRepoNotFound) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, sql.ErrNoRows
}

// on simule sql.ErrNoRows sans importer database/sql dans ce fichier (optionnel)
type sqlErrNoRows struct{}

func (sqlErrNoRows) Error() string { return "sql: no rows in result set" }
