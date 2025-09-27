# Arritech Take-Home Assignment

This is a full-stack take-home project built with:

- **Backend**: Go (Gin + GORM + PostgreSQL)
- **Frontend**: Vue 3 + Vite + Element Plus + Pinia + Vue Router
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose

The application provides basic **CRUD functionality** for managing users (list, search, create, edit, delete) with a simple frontend UI.

---

## 📦 Requirements

- [Go 1.22+](https://go.dev/)
- [Node.js 18+](https://nodejs.org/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)

---

## 🚀 Running with Docker (recommended)

You can spin up the entire stack (backend, frontend, and database) with:

```bash
docker-compose up --build
```

Frontend will be available at http://localhost:5173.

API will be available at http://localhost:8080.