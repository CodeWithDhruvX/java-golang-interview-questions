# Kubernetes Deployment for Go Microservices

## Overview

Kubernetes provides powerful orchestration capabilities for Go microservices. This guide covers deployment patterns, best practices, and integration strategies specifically for Go applications.

## Core Concepts

### 1. Kubernetes Objects for Go Services

- **Pod**: Basic deployment unit for Go containers
- **Deployment**: Manages Pod replicas and updates
- **Service**: Network endpoint for Go services
- **ConfigMap/Secret**: Configuration and sensitive data
- **Ingress**: External access to services

### 2. Go-Specific Considerations

- **Binary Size**: Optimize Go binaries for smaller containers
- **Concurrency**: Leverage Go's concurrency with Kubernetes scaling
- **Health Checks**: Implement proper readiness/liveness probes
- **Graceful Shutdown**: Handle SIGTERM signals properly

## Containerization Patterns

### 1. Multi-Stage Dockerfile

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main .

# Final stage
FROM scratch

# Copy CA certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy binary
COPY --from=builder /app/main /main

# Expose port
EXPOSE 8080

# Run the application
CMD ["/main"]
```

### 2. Optimized Alpine Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev git

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Runtime image
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
```

## Kubernetes Manifests

### 1. Basic Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-microservice
  labels:
    app: go-microservice
    version: v1
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-microservice
  template:
    metadata:
      labels:
        app: go-microservice
        version: v1
    spec:
      containers:
      - name: go-microservice
        image: your-registry/go-microservice:latest
        ports:
        - containerPort: 8080
          name: http
        env:
        - name: PORT
          value: "8080"
        - name: LOG_LEVEL
          value: "info"
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
        lifecycle:
          preStop:
            exec:
              command: ["/bin/sh", "-c", "sleep 15"]
```

### 2. Service Configuration

```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-microservice-service
  labels:
    app: go-microservice
spec:
  selector:
    app: go-microservice
  ports:
  - name: http
    port: 80
    targetPort: 8080
    protocol: TCP
  type: ClusterIP
---
apiVersion: v1
kind: Service
metadata:
  name: go-microservice-headless
  labels:
    app: go-microservice
spec:
  selector:
    app: go-microservice
  ports:
  - name: http
    port: 8080
    targetPort: 8080
  clusterIP: None
```

### 3. ConfigMap and Secret

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: go-microservice-config
data:
  config.yaml: |
    server:
      port: 8080
      timeout: 30s
    database:
      host: "postgres-service"
      port: 5432
      name: "myapp"
    redis:
      host: "redis-service"
      port: 6379
    logging:
      level: "info"
      format: "json"
---
apiVersion: v1
kind: Secret
metadata:
  name: go-microservice-secrets
type: Opaque
data:
  database-password: <base64-encoded-password>
  jwt-secret: <base64-encoded-jwt-secret>
  api-key: <base64-encoded-api-key>
```

### 4. Ingress Configuration

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-microservice-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
    nginx.ingress.kubernetes.io/rate-limit: "100"
spec:
  tls:
  - hosts:
    - api.yourdomain.com
    secretName: api-tls
  rules:
  - host: api.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: go-microservice-service
            port:
              number: 80
```

## Go Application Integration

### 1. Health Check Endpoints

```go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "time"
)

type HealthChecker struct {
    db       *sql.DB
    redis    *redis.Client
    startTime time.Time
}

type HealthResponse struct {
    Status    string            `json:"status"`
    Timestamp time.Time         `json:"timestamp"`
    Uptime    string            `json:"uptime"`
    Checks    map[string]string `json:"checks"`
}

func NewHealthChecker(db *sql.DB, redis *redis.Client) *HealthChecker {
    return &HealthChecker{
        db:        db,
        redis:     redis,
        startTime: time.Now(),
    }
}

func (hc *HealthChecker) LivenessHandler(w http.ResponseWriter, r *http.Request) {
    response := HealthResponse{
        Status:    "healthy",
        Timestamp: time.Now(),
        Uptime:    time.Since(hc.startTime).String(),
        Checks:    make(map[string]string),
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func (hc *HealthChecker) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    checks := make(map[string]string)
    status := "ready"
    
    // Check database connection
    if err := hc.db.PingContext(ctx); err != nil {
        checks["database"] = "unavailable: " + err.Error()
        status = "not_ready"
    } else {
        checks["database"] = "available"
    }
    
    // Check Redis connection
    if _, err := hc.redis.Ping(ctx).Result(); err != nil {
        checks["redis"] = "unavailable: " + err.Error()
        status = "not_ready"
    } else {
        checks["redis"] = "available"
    }
    
    response := HealthResponse{
        Status:    status,
        Timestamp: time.Now(),
        Uptime:    time.Since(hc.startTime).String(),
        Checks:    checks,
    }
    
    statusCode := http.StatusOK
    if status != "ready" {
        statusCode = http.StatusServiceUnavailable
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}
```

### 2. Graceful Shutdown

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Server struct {
    httpServer *http.Server
    db         *sql.DB
    redis      *redis.Client
}

func NewServer() *Server {
    return &Server{
        httpServer: &http.Server{
            Addr:         ":8080",
            ReadTimeout:  15 * time.Second,
            WriteTimeout: 15 * time.Second,
            IdleTimeout:  60 * time.Second,
        },
    }
}

func (s *Server) Start() error {
    // Setup routes
    mux := http.NewServeMux()
    mux.HandleFunc("/health", s.healthHandler)
    mux.HandleFunc("/ready", s.readinessHandler)
    mux.HandleFunc("/api/v1/", s.apiHandler)
    
    s.httpServer.Handler = mux
    
    // Start server
    log.Printf("Starting server on %s", s.httpServer.Addr)
    return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
    log.Println("Shutting down server...")
    
    // Shutdown HTTP server
    if err := s.httpServer.Shutdown(ctx); err != nil {
        return err
    }
    
    // Close database connections
    if s.db != nil {
        if err := s.db.Close(); err != nil {
            return err
        }
    }
    
    // Close Redis connection
    if s.redis != nil {
        if err := s.redis.Close(); err != nil {
            return err
        }
    }
    
    log.Println("Server shutdown complete")
    return nil
}

func main() {
    server := NewServer()
    
    // Start server in goroutine
    go func() {
        if err := server.Start(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()
    
    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    // Graceful shutdown with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Server shutdown error: %v", err)
    }
}
```

### 3. Configuration Management

```go
package config

import (
    "fmt"
    "os"
    "time"
    
    "github.com/spf13/viper"
    "k8s.io/client-go/rest"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    Logging  LoggingConfig  `mapstructure:"logging"`
}

type ServerConfig struct {
    Port           string        `mapstructure:"port"`
    Timeout        time.Duration `mapstructure:"timeout"`
    ReadTimeout    time.Duration `mapstructure:"read_timeout"`
    WriteTimeout   time.Duration `mapstructure:"write_timeout"`
    IdleTimeout    time.Duration `mapstructure:"idle_timeout"`
    ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Name     string `mapstructure:"name"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    SSLMode  string `mapstructure:"ssl_mode"`
    MaxOpenConns int `mapstructure:"max_open_conns"`
    MaxIdleConns int `mapstructure:"max_idle_conns"`
    ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type LoggingConfig struct {
    Level  string `mapstructure:"level"`
    Format string `mapstructure:"format"`
    Output string `mapstructure:"output"`
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    
    // Add config paths
    viper.AddConfigPath("/etc/app")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")
    
    // Set environment variable prefix
    viper.SetEnvPrefix("APP")
    viper.AutomaticEnv()
    
    // Set defaults
    setDefaults()
    
    // Read config file
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            log.Println("Config file not found, using defaults and environment variables")
        } else {
            return nil, fmt.Errorf("error reading config file: %w", err)
        }
    }
    
    // Unmarshal config
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("error unmarshaling config: %w", err)
    }
    
    return &config, nil
}

func setDefaults() {
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("server.timeout", "30s")
    viper.SetDefault("server.read_timeout", "15s")
    viper.SetDefault("server.write_timeout", "15s")
    viper.SetDefault("server.idle_timeout", "60s")
    viper.SetDefault("server.shutdown_timeout", "30s")
    
    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", 5432)
    viper.SetDefault("database.ssl_mode", "disable")
    viper.SetDefault("database.max_open_conns", 25)
    viper.SetDefault("database.max_idle_conns", 25)
    viper.SetDefault("database.conn_max_lifetime", "5m")
    
    viper.SetDefault("redis.host", "localhost")
    viper.SetDefault("redis.port", 6379)
    viper.SetDefault("redis.db", 0)
    
    viper.SetDefault("logging.level", "info")
    viper.SetDefault("logging.format", "json")
    viper.SetDefault("logging.output", "stdout")
}
```

## Helm Charts

### 1. Helm Chart Structure

```
helm/go-microservice/
├── Chart.yaml
├── values.yaml
├── values-prod.yaml
├── values-dev.yaml
├── templates/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── configmap.yaml
│   ├── secret.yaml
│   ├── ingress.yaml
│   ├── hpa.yaml
│   └── serviceaccount.yaml
└── charts/
```

### 2. Chart.yaml

```yaml
apiVersion: v2
name: go-microservice
description: A Helm chart for Go microservice
type: application
version: 0.1.0
appVersion: "1.0.0"
keywords:
  - go
  - microservice
  - api
maintainers:
  - name: Your Name
    email: your.email@example.com
```

### 3. values.yaml

```yaml
replicaCount: 3

image:
  repository: your-registry/go-microservice
  pullPolicy: IfNotPresent
  tag: "latest"

nameOverride: ""
fullnameOverride: ""

serviceAccount:
  create: true
  annotations: {}
  name: ""

podAnnotations: {}

podSecurityContext:
  fsGroup: 2000

securityContext:
  allowPrivilegeEscalation: false
  runAsNonRoot: true
  runAsUser: 1000
  capabilities:
    drop:
    - ALL

service:
  type: ClusterIP
  port: 80
  targetPort: 8080

ingress:
  enabled: true
  className: "nginx"
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
  hosts:
    - host: api.yourdomain.com
      paths:
        - path: /
          pathType: Prefix
  tls:
    - secretName: api-tls
      hosts:
        - api.yourdomain.com

resources:
  limits:
    cpu: 500m
    memory: 128Mi
  requests:
    cpu: 250m
    memory: 64Mi

autoscaling:
  enabled: true
  minReplicas: 3
  maxReplicas: 10
  targetCPUUtilizationPercentage: 80
  targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}

config:
  server:
    port: "8080"
    timeout: "30s"
  database:
    host: "postgres-service"
    port: 5432
    name: "myapp"
  redis:
    host: "redis-service"
    port: 6379
  logging:
    level: "info"
    format: "json"

secrets:
  databasePassword: ""
  jwtSecret: ""
  apiKey: ""
```

### 4. Deployment Template

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "go-microservice.fullname" . }}
  labels:
    {{- include "go-microservice.labels" . | nindent 4 }}
spec:
  {{- if not .Values.autoscaling.enabled }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "go-microservice.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      annotations:
        checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}
        checksum/secret: {{ include (print $.Template.BasePath "/secret.yaml") . | sha256sum }}
        {{- with .Values.podAnnotations }}
        {{- toYaml . | nindent 8 }}
        {{- end }}
      labels:
        {{- include "go-microservice.selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "go-microservice.serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: 8080
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /health
              port: http
            initialDelaySeconds: 30
            periodSeconds: 10
          readinessProbe:
            httpGet:
              path: /ready
              port: http
            initialDelaySeconds: 5
            periodSeconds: 5
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
          env:
            - name: PORT
              value: "8080"
          volumeMounts:
            - name: config
              mountPath: /etc/app/config.yaml
              subPath: config.yaml
            - name: secrets
              mountPath: /etc/app/secrets
              readOnly: true
      volumes:
        - name: config
          configMap:
            name: {{ include "go-microservice.fullname" . }}-config
        - name: secrets
          secret:
            secretName: {{ include "go-microservice.fullname" . }}-secrets
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
```

## Advanced Patterns

### 1. Horizontal Pod Autoscaler

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: go-microservice-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: go-microservice
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 80
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
  - type: Pods
    pods:
      metric:
        name: http_requests_per_second
      target:
        type: AverageValue
        averageValue: "100"
  behavior:
    scaleDown:
      stabilizationWindowSeconds: 300
      policies:
      - type: Percent
        value: 10
        periodSeconds: 60
    scaleUp:
      stabilizationWindowSeconds: 0
      policies:
      - type: Percent
        value: 100
        periodSeconds: 15
      - type: Pods
        value: 4
        periodSeconds: 15
      selectPolicy: Max
```

### 2. Pod Disruption Budget

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: go-microservice-pdb
spec:
  minAvailable: 2
  selector:
    matchLabels:
      app: go-microservice
```

### 3. Network Policy

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: go-microservice-netpol
spec:
  podSelector:
    matchLabels:
      app: go-microservice
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: ingress-nginx
    ports:
    - protocol: TCP
      port: 8080
  - from:
    - podSelector:
        matchLabels:
          app: frontend
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to:
    - podSelector:
        matchLabels:
          app: postgres
    ports:
    - protocol: TCP
      port: 5432
  - to:
    - podSelector:
        matchLabels:
          app: redis
    ports:
    - protocol: TCP
      port: 6379
```

## Interview Questions

### Technical Questions
1. How do you optimize Go applications for Kubernetes deployment?
2. Explain the difference between liveness and readiness probes in Go.
3. How do you implement graceful shutdown in Go applications?
4. What are the best practices for containerizing Go applications?
5. How do you manage configuration in Kubernetes for Go services?

### Design Questions
1. Design a deployment strategy for a Go microservices architecture.
2. How would you implement zero-downtime deployments for Go services?
3. Explain how to handle secrets management in Kubernetes.
4. Design a scaling strategy for Go applications based on custom metrics.
5. How do you implement service mesh integration with Go services?

### Practical Coding
1. Write a Go application with proper health check endpoints.
2. Implement graceful shutdown handling in a Go web server.
3. Create a configuration management system for Kubernetes deployments.
4. Design a metrics collection system for Go applications in Kubernetes.
5. Implement a circuit breaker pattern for Go services.

## Best Practices

1. **Small Containers**: Use multi-stage builds to minimize image size
2. **Resource Limits**: Set appropriate CPU and memory limits
3. **Health Checks**: Implement comprehensive liveness and readiness probes
4. **Graceful Shutdown**: Handle SIGTERM signals properly
5. **Security**: Use non-root users and drop capabilities
6. **Configuration**: Externalize configuration using ConfigMaps and Secrets
7. **Logging**: Use structured logging and centralized log collection
8. **Monitoring**: Integrate with Prometheus and other monitoring tools




Since you’re prepping for high-level roles, the key is to sound like someone who has actually broken (and fixed) things in production. Here is how you’d "speak" these answers in an interview—balancing technical depth with practical experience.

---

## Technical Questions

### 1. How do you optimize Go applications for Kubernetes deployment?

"First, I focus on the **binary size**. I use multi-stage Docker builds with a `distroless` or `alpine` base to keep the image slim, which speeds up pull times during scaling. Second, I always set **GOMAXPROCS**. Since Go isn't naturally 'cgroup aware,' it might see all the host's cores instead of the K8s quota, leading to CPU throttling. I usually use the `uber-go/automaxprocs` library to handle this automatically. Finally, I make sure the heap is managed by setting `GOGC` or `GOMEMLIMIT` to avoid OOM kills when memory usage spikes near the limit."

### 2. Explain the difference between liveness and readiness probes in Go.

"Think of it this way: **Liveness** is about 'Are you alive?' If it fails, K8s restarts the pod. I use this for catching deadlocks. **Readiness** is about 'Are you ready to work?' If this fails, K8s just stops sending traffic to that pod. In Go, I’ll have my readiness probe check downstream dependencies like the DB or Redis. If the DB is down, the app is 'alive' but not 'ready' to serve requests."

### 3. How do you implement graceful shutdown in Go applications?

"I use a **signal notify** context. I listen for `SIGINT` or `SIGTERM` signals. When one hits, I trigger a shutdown with a timeout—usually 30 seconds. This stops the server from accepting new requests but gives existing ones time to finish. I also make sure to close database connections and cleanup any background goroutines or workers during this window so we don't leave data in a partial state."

### 4. What are the best practices for containerizing Go applications?

"Definitely **multi-stage builds**. Build the binary in a heavy 'Golang' image, then copy just the executable into a tiny 'scratch' or 'distroless' image. Also, I ensure I’m running as a **non-root user** for security. I’ll also bake in a `.dockerignore` to keep out the `.git` folder and local binaries to keep the build context light."

### 5. How do you manage configuration in Kubernetes for Go services?

"I treat the binary as immutable. I use **ConfigMaps** for general settings and **Secrets** for sensitive data, injecting them as environment variables or mounting them as files. I’m a fan of the `spf13/viper` library because it can watch for file changes—so if I update a ConfigMap, the Go app can actually hot-reload the config without a full restart."

---

## Design Questions

### 1. Design a deployment strategy for a Go microservices architecture.

"I’d lean toward a **Canary deployment** using an Ingress controller like Nginx or a Service Mesh. We’d roll out the new Go binary to 5% of users, monitor the error rates and p99 latency, and if the metrics look healthy, we automate the full rollout. This minimizes the blast radius if we've introduced a bug."

### 2. How would you implement zero-downtime deployments for Go services?

"The core of this is the **Rolling Update** strategy in K8s combined with proper **Graceful Shutdown**. You need a `preStop` hook or a slight delay in your Go code to let the K8s endpoints update so you don't get 'connection refused' errors while the old pod is terminating. If you have both probes and graceful handling, K8s won't kill the old pod until the new one is fully 'Ready'."

### 3. Explain how to handle secrets management in Kubernetes.

"K8s Secrets are just base64 encoded by default, so for production, I’d integrate with something like **HashiCorp Vault** or **AWS Secrets Manager**. I’d use an 'External Secrets Operator' to sync those into K8s, or use a sidecar pattern where the Go app fetches secrets directly via an SDK at startup."

### 4. Design a scaling strategy for Go applications based on custom metrics.

"Standard CPU/Memory scaling is often too late for Go. I prefer using the **Horizontal Pod Autoscaler (HPA)** with **Prometheus Adapter**. We can scale based on 'Request Per Second' or 'Goroutine count.' If the number of active goroutines or the queue depth in our message broker spikes, we spin up more pods before the CPU even hits 50%."

### 5. How do you implement service mesh integration with Go services?

"I’d use **Istio or Linkerd**. The Go app doesn't need to know the mesh exists—it just talks to `localhost`. The sidecar proxy handles the mTLS (security), retries, and circuit breaking. It’s great for Go because it moves the 'networking complexity' out of our application code and into the infrastructure layer."

---

## Practical Coding (Conceptual Overviews)

### 1. Go application with health check endpoints.

"I’d use a standard `http.ServeMux`. I’d have `/healthz` return a simple 200 OK immediately, and `/readyz` would actually ping the database. If the ping fails, I return a 503 so K8s knows to pull this pod out of the load balancer."

### 2. Graceful shutdown handling.

"I’d wrap my `http.Server` in a `context.WithTimeout`. I'd use `signal.Notify` to catch `os.Interrupt`. Once caught, I call `server.Shutdown(ctx)` which gracefully drains the connections. It’s a clean loop that prevents dropped packets."

### 3. Configuration management system.

"I’d build a `Config` struct and use `envconfig` or `viper`. The app should look for a `config.yaml` at a specific path. In K8s, I mount the ConfigMap to that path. If I need a secret, I look for an Environment Variable."

### 4. Metrics collection system.

"I’d use the `prometheus/client_golang` library. I’d expose a `/metrics` endpoint and define a `prometheus.Counter` for total requests and a `prometheus.Histogram` for response latency. Prometheus then scrapes this every 15 seconds."

### 5. Circuit breaker pattern.

"I’d use the `sony/gobreaker` library. I wrap my outgoing HTTP client calls. If the downstream service starts failing (say, 50% failure rate), the breaker 'opens' and I immediately return an error or a cached response instead of hanging and wasting goroutines."

---

**Would you like me to dive deeper into the specific Go code for any of these, or perhaps mock a specific interview scenario?**