# 🟢 Go Theory Questions: 321–340 DevOps, Docker, and Deployment

## 321. How do you containerize a Go application?

**Answer:**
Go programs compile to a single static binary, making containerization essentially trivial.

We create a `Dockerfile`.
Step 1 (Build): Use `FROM golang:1.22-alpine AS builder` to compile the app.
Step 2 (Run): Use `FROM scratch` (an empty image). Copy *only* the binary and root CA certificates from the builder: `COPY --from=builder /app/main .`.

This results in a container image that is indistinguishable in size from the binary itself (usually < 20MB) and has zero OS vulnerability surface area.

---

## 322. What is a multi-stage Docker build and how does it help with Go?

**Answer:**
Multi-stage builds allow us to separate the **Build Environment** from the **Runtime Environment**.

We need the Go compiler, git, and dependencies to *build* the app. We strictly do *not* want those tools in production.
In the Dockerfile, we label the first `FROM` as `builder`. We run `go build`.
Then we start a new `FROM scratch` or `alpine`. We `COPY --from=builder` only the compiled artifact.

This reduces the final image size from ~800MB (full Go toolchain) to ~10MB (just the binary), improving security and deployment speed.

---

## 323. How do you reduce the size of a Go Docker image?

**Answer:**
1.  **Use Scratch**: The smallest possible base image.
2.  **Strip Debug Symbols**: Build with `go build -ldflags="-w -s"`. This removes the symbol table and DWARF debug info, shrinking the binary by ~30%.
3.  **No CGO**: `CGO_ENABLED=0`. This ensures a purely static binary with no dependency on `libc`, allowing it to run on `scratch` without runtime link errors.
4.  **UPX**: Optionally, use UPX to compress the binary executable itself, though this adds startup CPU cost.

---

## 324. How do you handle secrets in Go apps deployed via Docker?

**Answer:**
We strictly follow the **12-Factor App** methodology: Config (including secrets) comes from the Environment.

In Docker, we rely on environment variables injected at runtime: `docker run -e DB_PASS=secret ...`.
However, passing secrets in commands is visible in `docker history`.

In orchestration (Kubernetes), we mount secrets as files (tmpfs volumes) or export them as env vars from a K8s `Secret` object. The Go app simply reads `os.Getenv("DB_PASS")`. It remains agnostic to *how* the secret got there.

---

## 325. How do you use environment variables in Go?

**Answer:**
We use the standard library `os.Getenv("KEY")`.

However, `Getenv` returns an empty string if missing, which is ambiguous (is it missing or really empty?).
So we typically write a helper wrapper: `GetEnv(key, defaultVal)`.

For complex apps, we use libraries like `kelseyhightower/envconfig`. This maps environment variables directly to a struct `type Config struct { Port int \`envconfig:"PORT"\` }`. This centralizes config parsing and handles type conversion (string "8080" to int 8080) automatically on startup.

---

## 326. How do you compile a static Go binary for Alpine Linux?

**Answer:**
Alpine uses **musl libc**, whereas standard Linux uses **glibc**. If you link dynamically, your binary will crash on Alpine (`not found`).

The fix is **Static Compilation**.
Command: `CGO_ENABLED=0 GOOS=linux go build -o main .`

Disabling CGO tells the compiler: "Do not link against C libraries. Re-implement net/http and DNS resolvers properly in pure Go." The resulting binary is self-contained and runs on Alpine, Debian, CentOS, or even bare metal without modification.

---

## 327. What is `scratch` image in Docker and why is it used with Go?

**Answer:**
`scratch` is a special Docker reserved keyword indicating an **empty image**. No shell, no users, no folders, no libraries.

It is ideal for Go because Go binaries are static. We copy *just* the binary into scratch.
**Pros**: Security (hackers can't run `/bin/sh` or `apt-get` because they don't exist). Performance (minimal download).
**Gotcha**: You must manually copy `/etc/ssl/certs/ca-certificates.crt` from the builder image, otherwise your Go app cannot make HTTPS calls (it won't verify SSL certs).

---

## 328. How do you manage config files in Go across environments?

**Answer:**
We adhere to **Environment Parity**. The code should be identical; config changes.

We often use **Viper**. It allows a hierarchy:
1.  Default values in code.
2.  Read from `config.yaml` (if present).
3.  Override with **Environment Variables**.

In production (K8s), we rely exclusively on the Env Vars override feature (`VIPER_DB_HOST=prod-db`). This allows us to ship the exact same binary to Dev, Staging, and Prod, changing behavior solely via the infrastructure platform.

---

## 329. How do you build Go binaries for different OS/arch?

**Answer:**
Go has world-class **Cross-Compilation** built-in.

We use the variables `GOOS` (Operating System) and `GOARCH` (Architecture).
To build for Windows from a Mac: `GOOS=windows GOARCH=amd64 go build`.
To build for Raspberry Pi (ARM) from Linux: `GOOS=linux GOARCH=arm64 go build`.

Unlike C++, you don't need to install "cross-compilers" or virtual machines. The standard Go toolchain includes support for all platforms out of the box.

---

## 330. How do you use GoReleaser?

**Answer:**
GoReleaser is an automation tool that simplifies the release process.

With a simple `.goreleaser.yaml`, running `goreleaser release` will:
1.  Cross-compile binaries for Windows, Linux, Darwin (AMD64 + ARM).
2.  Create `.tar.gz` and `.zip` archives.
3.  Generate checksums (`sha256sum`).
4.  Create a GitHub Release and upload artifacts.
5.  Build and push Docker images.

It turns a complex 2-hour manual release checklist into a single command that runs in CI.

---

## 331. What is a Docker healthcheck for a Go app?

**Answer:**
A healthcheck tells Docker if the container is *functioning*, not just *running*. A deadlock might leave the process running but unresponsive.

In the Dockerfile: `HEALTHCHECK CMD curl -f http://localhost:8080/health || exit 1`.

The Go app must expose a lightweight `/health` endpoint.
Ideally, the Go app itself implements the check logic (e.g., using `tiny-health-check` binary) rather than relying on `curl` (which requires installing curl in the image, defeating the purpose of `scratch` images).

---

## 332. How do you log container stdout/stderr from Go?

**Answer:**
In container world, we **Log to Stdout**, not to files.

In Go, we configure our logger (slog/zap/logrus) to write to `os.Stdout` (for info) and `os.Stderr` (for errors).
`log.SetOutput(os.Stdout)`

Docker captures these streams. Kubernetes collects them, rotates the log files on the node, and Fluentd ships them to Elastic/Splunk. The application itself should *never* manage log files or rotation, as ephemeral containers can vanish at any moment.

---

## 333. How do you set up autoscaling for Go services?

**Answer:**
We rely on the platform (Kubernetes HPA - Horizontal Pod Autoscaler).

The Go app exposes metrics (e.g., CPU, Memory, or custom Request Rate) via an endpoint (Prometheus format).
The HPA watches these metrics. If CPU > 70%, it spawns more replicas.

Go's fast startup time (< 500ms) makes it perfect for autoscaling. Unlike Java (which might take 30s to warm up), a Go pod can receive traffic almost immediately, allowing us to handle traffic spikes aggressively.

---

## 334. How would you containerize a gRPC Go service?

**Answer:**
It is nearly identical to an HTTP service, but we must expose the **HTTP/2** port.

`EXPOSE 50051`.
The complexity usually lied in the health check, as standard tools are HTTP-based.
Nowadays, we implement key standard `grpc.health.v1.Health` service inside the Go app. Ideally, we use the `grpc-health-probe` binary in the container for the Kubernetes `livenessProbe` to verify the gRPC server is actually accepting RPCs.

---

## 335. How to deploy Go microservices in Kubernetes?

**Answer:**
We define a **Deployment** YAML and a **Service** YAML.

The Deployment manages the ReplicaSet (running 3 copies of the Go binary).
The Service provides a stable ClusterIP and load balancing across those pods.

For configuration, we use a **ConfigMap** mounted as environment variables.
For secrets, we use a **Secret** resource.
We rely on `readinessProbe` (calls `/health`) to ensure traffic isn't sent to a pod until it has established its database connections.

---

## 336. How do you write Helm charts for a Go app?

**Answer:**
Helm is a template engine for Kubernetes YAMLs.

We create a `values.yaml` defining variables like `replicaCount: 3` and `image: myapp:v1`.
The templates (`deployment.yaml`) use these: `replicas: {{ .Values.replicaCount }}`.

This allows us to deploy the same app to different environments with one command:
`helm install my-app ./chart -f values-prod.yaml`. It manages the upgrade lifecycle and allows easy rollbacks if a deployment fails.

---

## 337. How do you monitor a Go service in production?

**Answer:**
We follow the **Three Pillars of Observability**:

1.  **Metrics**: (Prometheus) Time-series data. "How much RAM?" "Requests per second?"
2.  **Logs**: (Elastic/Loki) Discrete events. "Error connecting to DB on line 42."
3.  **Traces**: (Jaeger/Tempo) Request lifecycles. "This request took 200ms in Service A and 50ms in Service B."

We instrument the Go code using OpenTelemetry libraries to emit all three automatically without polluting the business logic.

---

## 338. How do you use Prometheus with a Go app?

**Answer:**
We use the official `github.com/prometheus/client_golang` library.

We register metrics: `var requests = promauto.NewCounter(...)`.
We increment them in handlers: `requests.Inc()`.
We expose an HTTP handler: `http.Handle("/metrics", promhttp.Handler())`.

Prometheus scrapes this endpoint. The power is in **Labels**: `requests.WithLabelValues("POST", "200").Inc()`. This allows us to graph error rates purely by query, without changing code.

---

## 339. How do you enable structured logging in production?

**Answer:**
We enforce JSON logging via the logger configuration.

In Dev: `logger = slog.New(slog.NewTextHandler(os.Stdout))` (Pretty text).
In Prod: `logger = slog.New(slog.NewJSONHandler(os.Stdout))` (Machine implementation).

JSON is non-negotiable in production because it handles multi-line stack traces correctly (as a single JSON field) and allows log aggregation tools to index fields like `user_id` or `request_id` for instant searchability.

---

## 340. How do you handle log rotation in containerized Go apps?

**Answer:**
**We Don't.**

Log rotation is a file-system concern. Since Go containers write to `stdout`, they just output a stream.

The **Container Runtime** (Docker Engine / CRI-O) handles the file storage on the host node. We configure the Docker Daemon daemon.json: `{"log-driver": "json-file", "log-opts": {"max-size": "10m", "max-file": "3"}}`. The Go app stays simple and oblivious to how its logs are stored or rotated.
