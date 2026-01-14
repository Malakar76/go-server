// @title           Go Server API
// @version         1.0
// @description     Simple Gin API with SQLite
// @host            localhost:8080
// @BasePath        /
// @schemes         http
package main

import (
	"go-server/internal/app"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	defer sentry.Flush(2 * time.Second)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
