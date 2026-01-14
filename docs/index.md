# Go-Server 🚀

Une API REST en **Go** construite avec **Gin**, pensée pour être simple, rapide et déployable facilement (ex: Raspberry Pi).

---

## ✅ Fonctionnalités

- API REST avec Gin
- Endpoints versionnés (`/v1`)
- Stockage avec SQLite (ou PostgreSQL)
- Tests unitaires (Go test)
- CI/CD via GitHub Actions (build multi-arch + release)

---

## 📦 Installation

### Prérequis
- Go >= 1.21
- (Optionnel) SQLite / PostgreSQL selon ton setup

### Lancer en local

```bash
go run ./cmd/api