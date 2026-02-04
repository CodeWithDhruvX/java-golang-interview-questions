## 🟡 DevOps, Docker, and Deployment (Questions 321-340)

### Question 321: How do you containerize a Go application?

**Answer:**
Containerizing a Go app involves creating a `Dockerfile` to build and run the binary.

**Basic Dockerfile:**
```dockerfile
# Start from a small base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Command to run
CMD ["./main"]
```

**Build and Run:**
```bash
docker build -t my-go-app .
docker run -p 8080:8080 my-go-app
```

---

### Question 322: What is a multi-stage Docker build and how does it help with Go?

**Answer:**
Multi-stage builds allow you to use a large image for building (with compilers/tools) and a tiny image for running (just the binary). This drastically reduces image size.

**Example:**
```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```

**Benefits:**
- Removes source code, compilers, and cache from final image.
- Reduces image size from ~800MB (golang image) to ~10MB (alpine + binary).
- Improves security (fewer tools installed in runtime container).

---

### Question 323: How do you reduce the size of a Go Docker image?

**Answer:**
1. **Use Multi-stage builds:** Separate build and runtime environments.
2. **Use `scratch` or `alpine` base images:**
   - `scratch`: Empty image (smallest possible).
   - `alpine`: Minimal Linux (~5MB).
3. **Strip debug information:**
   ```bash
   go build -ldflags="-s -w" -o myapp
   ```
   - `-s`: Disable symbol table.
   - `-w`: Disable DWARF generation.
4. **Compress binary (optional):** Use `upx`.

**Optimized Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp .

FROM scratch
COPY --from=builder /app/myapp /myapp
# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/myapp"]
```

---

### Question 324: How do you handle secrets in Go apps deployed via Docker?

**Answer:**
Never hardcode secrets. Use environment variables or secret managers.

**1. Environment Variables:**
```go
func main() {
    dbPass := os.Getenv("DB_PASSWORD")
    if dbPass == "" {
        log.Fatal("DB_PASSWORD not set")
    }
}
```
Run with: `docker run -e DB_PASSWORD=secret myapp`

**2. Docker Secrets (Swarm/K8s):**
Mount secrets as files (e.g., `/run/secrets/db_password`).
```go
func readSecret(name string) string {
    content, _ := os.ReadFile("/run/secrets/" + name)
    return strings.TrimSpace(string(content))
}
```

**3. External Secret Managers:**
Use SDKs for AWS Secrets Manager, HashiCorp Vault, etc.

---

### Question 325: How do you use environment variables in Go?

**Answer:**
Use the `os` package or libraries like `godotenv` or `viper`.

**Standard library:**
```go
import "os"

// Get value
port := os.Getenv("PORT")
if port == "" {
    port = "8080" // Default
}

// Set value (for current process)
os.Setenv("APP_ENV", "production")

// Expand string
conn := os.ExpandEnv("postgres://user:$DB_PASSWORD@localhost:5432/db")
```

**Using `godotenv` (Load from .env file):**
```go
import "github.com/joho/godotenv"

func main() {
    godotenv.Load() // Loads .env file
    s3Bucket := os.Getenv("S3_BUCKET")
}
```

---

### Question 326: How do you compile a static Go binary for Alpine Linux?

**Answer:**
Alpine uses `musl` libc instead of `glibc`. To ensure compatibility or avoid dependency issues, build a statically linked binary.

**Command:**
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .
```

- `CGO_ENABLED=0`: Disables CGO, ensuring a pure Go binary (static by default).
- `GOOS=linux`: Targets Linux.
- `-a`: Force rebuild of packages.

**Why?**
- Allows the binary to run on `scratch` images.
- Avoids "no such file or directory" errors when moving binaries between distros.

---

### Question 327: What is scratch image in Docker and why is it used with Go?

**Answer:**
`scratch` is an explicitly empty Docker image. It contains no filesystem, no shell, no libraries.

**Why use it with Go?**
- **Size:** Results in the smallest possible Docker image (size = binary size).
- **Security:** Small attack surface (no shell, no tools to exploit).

**Gotchas:**
- **No shell:** Can't use `sh`, `bash`, or `exec` inside.
- **Missing SSL certs:** HTTPS requests will fail. You must copy `/etc/ssl/certs/ca-certificates.crt` from a builder stage.
- **Missing Timezone data:** Time functions might work, but loading locations (e.g., `time.LoadLocation("America/New_York")`) requires copying `zoneinfo`.

---

### Question 328: How do you manage config files in Go across environments?

**Answer:**
Best practice is to prioritize: Flag > Env Var > Config File > Default.

**Using `viper`:**
```go
import "github.com/spf13/viper"

func initConfig() {
    viper.SetConfigName("config") // config.yaml
    viper.AddConfigPath(".")      // Look in current dir
    viper.AutomaticEnv()          // Read env vars

    // Set defaults
    viper.SetDefault("port", 8080)

    if err := viper.ReadInConfig(); err != nil {
        log.Println("No config file found, using defaults")
    }
}

func main() {
    port := viper.GetInt("port")
}
```

**Pattern:**
- Local dev: use `config.yaml`
- Production: use Environment Variables (override config file).

---

### Question 329: How do you build Go binaries for different OS/arch?

**Answer:**
Go supports cross-compilation out of the box using `GOOS` and `GOARCH` environment variables.

**Examples:**
```bash
# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o myapp-linux

# macOS (M1/M2 - ARM64)
GOOS=darwin GOARCH=arm64 go build -o myapp-mac

# Linux (ARM - Raspberry Pi)
GOOS=linux GOARCH=arm go build -o myapp-pi
```

**Check supported platforms:**
```bash
go tool dist list
```

---

### Question 330: How do you use GoReleaser?

**Answer:**
GoReleaser automates building, packaging, and releasing Go binaries to GitHub/GitLab.

**Steps:**
1. **Install:** `brew install goreleaser`
2. **Init:** `goreleaser init` (creates `.goreleaser.yaml`)
3. **Configure `.goreleaser.yaml`:**
   ```yaml
   builds:
     - env:
         - CGO_ENABLED=0
       goos:
         - linux
         - windows
         - darwin
       goarch:
         - amd64
         - arm64
   ```
4. **Release:**
   ```bash
   git tag -a v1.0.0 -m "First release"
   git push origin v1.0.0
   goreleaser release --clean
   ```

**What it does:**
- Cross-compiles for defined targets.
- Creates archives (.tar.gz, .zip).
- Generates checksums.
- Creates GitHub Release and uploads artifacts.
- Can build Docker images and update Homebrew taps.

---

### Question 331: What is a Docker healthcheck for a Go app?

**Answer:**
A healthcheck ensures the container is actually ready to serve traffic, not just running.

**In Dockerfile:**
```dockerfile
# Check every 30s, timeout 3s
HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1
```
*Note: This requires `curl` to be installed in the image.*

**Go Implementation:**
```go
func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Check DB connection, cache, etc.
        if err := db.Ping(); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    http.ListenAndServe(":8080", nil)
}
```

**For `scratch` images:**
Use a separate Go binary for healthchecking or the built-in `grpc-health-probe` if using gRPC.

---

### Question 332: How do you log container stdout/stderr from Go?

**Answer:**
By default, Docker captures stdout and stderr. In Go, simply print to these streams.

**Standard Log:**
```go
log.Println("This goes to stderr") // Default logger uses stderr
```

**Explicit Streams:**
```go
fmt.Fprintln(os.Stdout, "Info log")
fmt.Fprintln(os.Stderr, "Error log")
```

**Structured Logging (JSON):**
For production/parsing (ELK, Datadog), use JSON logs.
```go
import "log/slog" // Go 1.21+

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("User logged in", "user_id", 42)
// Output: {"time":"...","level":"INFO","msg":"User logged in","user_id":42}
```

**Docker logs command:**
```bash
docker logs -f my-container
```

---

### Question 333: How do you set up autoscaling for Go services?

**Answer:**
Autoscaling is typically handled by the orchestrator (Kubernetes), not the Go app itself.

**Horizontal Pod Autoscaler (HPA) in K8s:**
Scales pods based on CPU/Memory or custom metrics.

1. **Expose Metrics:**
   Use specific endpoints (e.g., `/metrics` with **Prometheus**).
   ```go
   // Expose "active_requests" metric
   ```

2. **Configure HPA:**
   ```yaml
   apiVersion: autoscaling/v2
   kind: HorizontalPodAutoscaler
   metadata:
     name: my-go-app
   spec:
     scaleTargetRef:
       apiVersion: apps/v1
       kind: Deployment
       name: my-go-app
     minReplicas: 2
     maxReplicas: 10
     metrics:
     - type: Resource
       resource:
         name: cpu
         target:
           type: Utilization
           averageUtilization: 50
   ```
   *Scales up if CPU usage > 50%.*

---

### Question 334: How would you containerize a gRPC Go service?

**Answer:**
Similar to HTTP, but needs to expose the gRPC port.

**Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o grpc-server server.go

FROM alpine:latest
COPY --from=builder /src/grpc-server /app/
EXPOSE 50051
CMD ["/app/grpc-server"]
```

**Testing:**
Since gRPC is not HTTP (usually), you can't just `curl` it.
- Use **`grpcurl`** for command line testing.
- Include **gRPC Health Probe** in the container for K8s readiness probes.

```dockerfile
# Install health probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.19 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
```

---

### Question 335: How to deploy Go microservices in Kubernetes?

**Answer:**
Deployment involves creating K8s manifests:

**1. Deployment.yaml:**
Defines the pods, replicas, and container image.
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
      - name: go-api
        image: myregistry/go-api:v1
        ports:
        - containerPort: 8080
        env:
        - name: DB_URL
          valueFrom:
            secretKeyRef:
              name: db-secrets
              key: url
```

**2. Service.yaml:**
Exposes the pods to the cluster.
```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-api-svc
spec:
  selector:
    app: go-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

**Steps:**
1. Build image: `docker build -t ...`
2. Push image: `docker push ...`
3. Apply manifests: `kubectl apply -f k8s/`

---

### Question 336: How do you write Helm charts for a Go app?

**Answer:**
Helm creates reusable, templated K8s manifests.

**Structure:**
```
my-chart/
  Chart.yaml    # Meta info
  values.yaml   # Default config
  templates/    # Manifest templates
    deployment.yaml
    service.yaml
```

**values.yaml:**
```yaml
replicaCount: 2
image:
  repository: my-go-app
  tag: "1.0.0"
service:
  port: 80
```

**templates/deployment.yaml:**
```yaml
apiVersion: apps/v1
kind: Deployment
spec:
  replicas: {{ .Values.replicaCount }}
  ...
      containers:
        - image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
```

**Install:**
```bash
helm install my-release ./my-chart
```
This allows deploying the same Go app to Dev, Staging, and Prod with different `values.yaml` files.

---

### Question 337: How do you monitor a Go service in production?

**Answer:**
Comprehensive monitoring includes:

1. **Metrics (Quantitative):**
   - Use **Prometheus** for time-series data.
   - Track: Request rate, Error rate, Latency (RED method), Memory, Goroutines.

2. **Logs (Qualitative):**
   - Use Structured Logging (JSON).
   - Aggregate via ELK stack (Elasticsearch, Logstash, Kibana) or Loki.

3. **Tracing (Context):**
   - Use **OpenTelemetry** (Jaeger/Zipkin) to trace requests across microservices.

4. **Health Checks:**
   - Liveness/Readiness probes for K8s.

**Go Code:**
Expose `/metrics` endpoint for Prometheus scraping and inject trace IDs into logs.

---

### Question 338: How do you use Prometheus with a Go app?

**Answer:**
Use the `prometheus/client_golang` library.

**Code:**
```go
package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a metric
var reqCounter = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
)

func init() {
    // Register metric
    prometheus.MustRegister(reqCounter)
}

func handler(w http.ResponseWriter, r *http.Request) {
    reqCounter.Inc() // Increment
    w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/", handler)
    // Expose metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}
```
Prometheus server then scrapes `http://app:8080/metrics`.

---

### Question 339: How do you enable structured logging in production?

**Answer:**
Structured logging (JSON) is essential for machine parsing.

**Using `log/slog` (Go 1.21+):**
```go
import (
    "log/slog"
    "os"
)

func main() {
    // Determine format based on env
    var handler slog.Handler
    if os.Getenv("ENV") == "production" {
        handler = slog.NewJSONHandler(os.Stdout, nil)
    } else {
        handler = slog.NewTextHandler(os.Stdout, nil)
    }
    
    logger := slog.New(handler)
    slog.SetDefault(logger)

    // Log with fields
    slog.Info("Request processed", 
        "path", "/api/v1/user",
        "status", 200,
        "duration_ms", 45,
    )
}
```
**Output (Prod):**
`{"time":"...","level":"INFO","msg":"Request processed","path":"/api/v1/user","status":200,"duration_ms":45}`

---

### Question 340: How do you handle log rotation in containerized Go apps?

**Answer:**
**You usually don't.**

In containerized environments (Docker/K8s):
1. **App Responsibility:** The app should simply write to `stdout`/`stderr`. It should **not** manage files or rotation.
2. **Platform Responsibility:** The container runtime (Docker daemon, Kubelet) captures these logs.
   - **Docker:** Configure logging driver (e.g., `json-file` with `max-size` and `max-file`).
   - **K8s:** Often uses a sidecar (e.g., Fluentd, Filebeat) to ship logs to a central backend (Elasticsearch, Splunk).

**If running on raw VM (Systemd):**
Use **Lumberjack** library in Go to manage file rotation application-side.
```go
import "gopkg.in/natefinch/lumberjack.v2"

log.SetOutput(&lumberjack.Logger{
    Filename:   "/var/log/myapp/foo.log",
    MaxSize:    100, // megabytes
    MaxBackups: 3,
    MaxAge:     28,  // days
    Compress:   true,
})
```
But for Docker, stick to standard output.
