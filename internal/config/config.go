package config

import "os"

type Config struct {
	HTTPAddr string
	DBPath   string
}

func FromEnv() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	db := os.Getenv("DB_PATH")
	if db == "" {
		db = "data/app.db"
	}
	return Config{HTTPAddr: addr, DBPath: db}
}
