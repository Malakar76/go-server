package app

import (
	"go-server/internal/config"
	"go-server/internal/http"
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
	if err := db.InitSchema(sqliteDB); err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepository(sqliteDB)
	userSvc := service.NewUserService(userRepo)

	srv := http.NewServer(cfg, userSvc)

	return &App{server: srv}, nil
}

func (a *App) Run() error {
	return a.server.Run()
}
