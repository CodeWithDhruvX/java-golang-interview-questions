# 🐳 08 — DevOps & Deployment in Go
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Building Go binaries
- Dockerizing a Go app
- Environment variable configuration
- Go modules (`go.mod`, `go.sum`)
- CI/CD basics with Go
- Cross-platform compilation

---

## ❓ Most Asked Questions

### Q1. What are the key Go build commands?

```bash
# Run directly (no binary created)
go run main.go

# Build binary for current OS
go build -o app ./cmd/main.go

# Strip debug info (smaller binary) — production
go build -ldflags="-s -w" -o app .

# Cross-compile for Linux from Windows/Mac
GOOS=linux GOARCH=amd64 go build -o app-linux .

# Cross-compile for Windows
GOOS=windows GOARCH=amd64 go build -o app.exe .

# Install to $GOPATH/bin
go install ./...
```

---

### Q2. How do you Dockerize a Go application?

```dockerfile
# Multi-stage Dockerfile — recommended for production

# ---- Build Stage ----
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Download dependencies first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/main.go

# ---- Run Stage ----
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
```

```bash
# Build and run
docker build -t my-go-app .
docker run -p 8080:8080 my-go-app
```

> **Key:** `CGO_ENABLED=0` creates a statically linked binary that runs in minimal containers like `alpine` or even `scratch`.

---

### Q3. How do you manage configuration with environment variables?

```go
import "os"

// Basic
port := os.Getenv("PORT")
if port == "" { port = "8080" }

// Using a config struct
type Config struct {
    Port     string
    DBHost   string
    DBPort   string
    DBName   string
    DBUser   string
    DBPass   string
    Debug    bool
}

func LoadConfig() Config {
    return Config{
        Port:   getEnv("PORT", "8080"),
        DBHost: getEnv("DB_HOST", "localhost"),
        DBPort: getEnv("DB_PORT", "5432"),
        DBName: getEnv("DB_NAME", "mydb"),
        DBUser: getEnv("DB_USER", "postgres"),
        DBPass: getEnv("DB_PASS", ""),
        Debug:  os.Getenv("DEBUG") == "true",
    }
}

func getEnv(key, defaultVal string) string {
    if val := os.Getenv(key); val != "" { return val }
    return defaultVal
}

// Better: use godotenv for .env files
import "github.com/joho/godotenv"
godotenv.Load(".env")  // loads .env into os environment
```

---

### Q4. How do Go modules work (`go.mod`, `go.sum`)?

```bash
# Initialize a new module
go mod init github.com/myuser/myapp

# Add a dependency
go get github.com/gin-gonic/gin@v1.9.1

# Remove unused dependencies
go mod tidy

# Vendor dependencies locally
go mod vendor
```

```
# go.mod structure
module github.com/myuser/myapp

go 1.22

require (
    github.com/gin-gonic/gin v1.9.1
    gorm.io/gorm v1.25.0
    gorm.io/driver/mysql v1.5.0
)
```

> `go.sum` contains cryptographic hashes of module versions — ensures reproducible builds.

---

### Q5. How do you set up a basic GitHub Actions CI pipeline for Go?

```yaml
# .github/workflows/ci.yml
name: Go CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Download dependencies
        run: go mod download

      - name: Run linter
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
          golangci-lint run ./...

      - name: Run tests
        run: go test -v -race -coverprofile=coverage.out ./...

      - name: Build
        run: go build -ldflags="-s -w" -o app ./cmd/main.go
```

---

### Q6. How do you use build tags (build constraints)?

```go
//go:build linux
// +build linux  (old syntax, keep for compatibility)

package main

// This code ONLY compiled on Linux
func getOSInfo() string {
    return "Running on Linux"
}
```

```bash
# Build only for specific OS
GOOS=linux go build -tags linux .

# Build with custom tag
go build -tags debug .
```

Common use cases:
- Platform-specific code (Linux/Windows/Mac)
- Including/excluding test utilities
- Feature flags at compile time
- Debug vs production builds

---

### Q7. What is `go vet` and `golangci-lint`?

```bash
# go vet — built-in static analysis
go vet ./...
# Catches: unreachable code, wrong format verbs, suspicious constructs

# golangci-lint — runs 50+ linters at once
golangci-lint run ./...
```

Key linters from `golangci-lint`:
| Linter | What it checks |
|--------|---------------|
| `staticcheck` | Bugs, unused code |
| `errcheck` | Unchecked errors |
| `govet` | Go vet checks |
| `gofmt` | Code formatting |
| `gosec` | Security issues |
| `ineffassign` | Ineffectual assignments |

---

### Q8. How do you use `go generate`?

```go
// In a Go source file, add a generate directive:
//go:generate mockery --name=UserRepository --output=./mocks

//go:generate stringer -type=Weekday

// Run all directives in current package
// go generate ./...
```
> Common uses: generate mocks from interfaces, generate string methods for enums, compile protobuf files.
