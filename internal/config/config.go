package config

import "os"

type Config struct {
	HTTPAddr  string
	DBPath    string
	SentryDSN string
}

func FromEnv() Config {
	cfg := Config{
		HTTPAddr:  getEnv("HTTP_ADDR", ":8080"),
		DBPath:    getEnv("DB_PATH", "data/app.db"),
		SentryDSN: getEnv("SENTRY_DSN", ""),
	}
	return cfg
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
