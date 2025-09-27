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

1. Add the below .env file into the root of the backend folder:
```
PORT=8080
LOG_LEVEL=debug
ENV=local
DATABASE_URL=postgres://arritechappuser:arritech_2025@db:5432/arritech_db?sslmode=disable
```

2. Add the below .env file into the root of the frontend folder:
```
VITE_API_BASE_URL=http://localhost:8085/api/v1
```

3. Run the below command from the root of the repository
```bash
docker-compose up --build
```

Frontend will be available at http://localhost:5173.

API will be available at http://localhost:8080.