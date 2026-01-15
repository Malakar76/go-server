package app

import (
	"go-server/internal/config"
	"go-server/internal/http"
	"go-server/internal/http/handlers"
	"go-server/internal/platform/db"
	repo "go-server/internal/repository/sqlite"
	"go-server/internal/service"
	"log"

	"github.com/getsentry/sentry-go"
)

type App struct {
	server *http.Server
}

func setupSentry(cfg config.Config) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		TracesSampleRate: 0.05,
		SendDefaultPII:   true,
	})
	if err != nil {
		log.Printf("sentry init failed: %v", err)
	}
}

func New() (*App, error) {
	cfg := config.FromEnv()
	setupSentry(cfg)

	sqliteDB, err := db.OpenSQLite(cfg.DBPath)
	if err != nil {
		return nil, err
	}
	if err := db.RunMigrations(sqliteDB, "migrations"); err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepository(sqliteDB)
	userSvc := service.NewUserService(userRepo)
	objectRepo := repo.NewObjectRepository(sqliteDB)
	objectSvc := service.NewObjectService(objectRepo)

	h := handlers.New(handlers.Deps{
		UserSvc:   userSvc,
		ObjectSvc: objectSvc,
	})

	srv := http.NewServer(cfg, h)

	return &App{server: srv}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
