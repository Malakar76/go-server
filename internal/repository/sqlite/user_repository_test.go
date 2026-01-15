package sqlite

import (
	"context"
	"database/sql"
	"testing"

	"go-server/internal/platform/db"

	_ "modernc.org/sqlite"
)

func TestUserRepository_CRUD(t *testing.T) {
	// DB en mémoire
	sqldb, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	defer sqldb.Close()

	// pragmas + schema (réutilise tes fonctions si possible)
	if _, err := sqldb.Exec(`PRAGMA journal_mode=WAL;`); err != nil {
		t.Fatal(err)
	}
	if err := db.RunMigrations(sqldb, "../../../migrations"); err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(sqldb)

	u, err := repo.Create(context.Background(), "Robin")
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.GetByID(context.Background(), u.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Name != "Robin" {
		t.Fatalf("expected name Robin, got %q", got.Name)
	}
}
