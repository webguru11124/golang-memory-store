
# Golang In-Memory Data Store (Redis-like)

This project is a simple in-memory data store implemented in **Go (Golang)** with support for:
- Strings and Lists
- Time-To-Live (TTL) keys
- Authentication (JWT)
- Data Persistence (File-based & Database Support)
- REST API
- Go Client Library
- Unit & Integration Tests
- Docker Deployment

## Features
- ✅ String & List operations (Set, Get, Delete, Push, Pop)
- ✅ TTL (Automatic expiration of keys)
- ✅ Authentication using JWT (JSON Web Tokens)
- ✅ Data Persistence (Saving data to disk & loading on startup or saving to a database)
- ✅ Modular architecture (Separate core, API, client, and persistence modules)
- ✅ Docker Support for deployment
- ✅ Test coverage with Unit & Integration tests

---

## Project Structure
```
/golang-memory-store
├── /cmd
│   └── /server             # Main Application Entry Point
├── /internal               
│   ├── /api                # REST API Handlers
│   ├── /auth               # Authentication Layer (JWT)
│   ├── /core               # Core Logic (In-memory Store)
│   ├── /client             # Go Client Library (SDK)
│   ├── /persistence        # Data Persistence Layer (File & Database Storage)
├── /tests                  # Integration Tests
├── /Dockerfile             
├── /Makefile               
├── /go.mod                 
├── /go.sum                 
└── /README.md              
```

---

## Requirements
- Go 1.20 or above
- Docker (Optional)
- `github.com/golang-jwt/jwt/v4`
- `github.com/gorilla/mux`
- `gorm.io/gorm`
- `gorm.io/driver/sqlite`
- `gorm.io/driver/postgres`

---

## Installation
```bash
git clone https://github.com/techlando/golang-memory-store.git
cd golang-memory-store
go mod tidy
```

---

## Usage

### Run the Server Locally
```bash
go run ./cmd/server
```

---

## Persistence Configuration

### File-Based Persistence (Default)
```bash
export ENABLE_PERSISTENCE=true
```

### SQLite Persistence
```bash
export DB_TYPE=sqlite
export DB_DSN=gorm.db
```

### PostgreSQL Persistence
```bash
export DB_TYPE=postgres
export DB_DSN="user=youruser password=yourpass dbname=yourdb port=5432 sslmode=disable"
```

---

## Testing

### Unit Tests (Core Module)
```bash
go test -v ./internal/core
```

### Integration Tests (API Module)
Ensure the server is running on `http://localhost:8080`.
```bash
go test -v ./tests
```

---

## Manual API Testing (Using `curl`)

### Generate JWT Token
```bash
curl -X POST http://localhost:8080/token -H "Content-Type: application/json" -d '{"username":"testuser"}'
```

### Set a Key
```bash
curl -X POST http://localhost:8080/set -H "Authorization: Bearer YOUR_JWT_TOKEN" -H "Content-Type: application/json" -d '{"key": "foo", "value": "bar", "ttl": 60}'
```

### Get a Key
```bash
curl -X GET http://localhost:8080/get/foo -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Delete a Key
```bash
curl -X DELETE http://localhost:8080/delete/foo -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Push to List
```bash
curl -X POST http://localhost:8080/list/push -H "Authorization: Bearer YOUR_JWT_TOKEN" -H "Content-Type: application/json" -d '{"key": "mylist", "value": "item1"}'
```

### Pop from List
```bash
curl -X POST http://localhost:8080/list/pop/mylist -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## Using Docker

### Build Docker Image
```bash
docker build -t golang-memory-store .
```

### Run Docker Container
```bash
docker run -p 8080:8080 -d golang-memory-store
```

---

## Makefile Commands

| Command         | Description                       |
|-----------------|-----------------------------------|
| `make build`    | Builds the application           |
| `make run`      | Runs the application locally     |
| `make test`     | Runs all tests                   |
| `make docker-build` | Builds the Docker image       |
| `make docker-run`   | Runs the Docker container      |
| `make clean`    | Cleans up build files            |

---

## API Documentation
See [InMemoryDataStoreAPI.md](./docs/api_docs.md)

## Client Library Usage
See [ClientLibraryAndDeploymentDocs.md](./docs/client_docs.md)

## Docker Usage
See [DockerDeployDocs.md](./docs/deploy_docs.md)

---
