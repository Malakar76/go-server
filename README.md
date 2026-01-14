# Go Server

A small REST API server written in **Go** using **Gin**.

This project is designed to be lightweight and easy to run locally or deploy on small devices (ex: Raspberry Pi).

---

## 🚀 Features

- Fast HTTP server powered by [Gin](https://github.com/gin-gonic/gin)
- Simple REST endpoints (example: healtxh check)
- SQLite support via a local DB file (optional depending on your setup)
- Ready for CI/CD (tests + builds + releases)

---

## ✅ Requirements

- Go (1.21+ recommended)

---

## ⚙️ Environment Variables

You can configure the server using the following environment variables.

### `GIN_MODE`
Sets Gin runtime mode.

**Allowed values:** `debug`, `release`, `test`  
**Default:** `debug`

### `HTTP_ADDR`
Defines the address the HTTP server listens on.

**Format: :** `host:port`
**Default:** `:8080`

### `DB_PATH`
Path to the SQLite database file.

**Default:** `data/app.db`