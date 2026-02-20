# 🟡 **321–340: DevOps, Docker, and Deployment**

### 321. How do you containerize a Go application?
"I write a **Dockerfile** that leverages Go's static compilation.

I almost always use a **multi-stage build**.
Stage 1 (Builder): Uses `golang:alpine`. I copy `go.mod`, download dependencies, and build the binary using `go build -o app`.
Stage 2 (Runner): Uses `gcr.io/distroless/static` or `scratch`. I copy *only* the compiled binary from Stage 1.

This results in a production image that is incredibly small (often <15MB) and secure because it lacks a shell or package manager."

#### Indepth
Always use `ENTRYPOINT` instead of `CMD` for the binary. `ENTRYPOINT ["/app"]` makes the container executable. `CMD` provides default arguments. This allows users to run `docker run my-image --help` and have the `--help` flag passed directly to your Go binary.

---

### 322. What is a multi-stage Docker build and how does it help with Go?
"A multi-stage build allows me to use a heavy image for compilation and a lightweight image for execution, all in one Dockerfile.

For Go, this is critical. The `golang` image (with GCC, Git, etc.) is ~800MB. My binary is ~10MB.
If I shipped the builder image, I’d be wasting 790MB of space and exposing a massive attack surface. Multi-stage builds strip all that away, leaving just the executable."

#### Indepth
You can reuse the Go Module Cache in multi-stage builds. Mount a cache volume: `RUN --mount=type=cache,target=/go/pkg/mod go build ...`. This speeds up subsequent builds by 10x because it doesn't re-download the internet every time you change a single line of code.

---

### 323. How do you reduce the size of a Go Docker image?
"I target the **scratch** base image, which is an empty file system.

To make this work, I compile with `CGO_ENABLED=0` to ensure a statically linked binary (no dependency on `glibc`).
I also use linker flags `-ldflags="-s -w"` to strip debug information, which shrinks the binary size by about 25%.

The final result is an image that is literally the size of the binary itself. You can't get smaller than that."

#### Indepth
For extreme size reduction, you can use **UPX** to compress the binary. It can shrink a 10MB Go binary to 3MB. However, it increases startup time (decompression overhead) and can trigger false positives in antivirus software. Generally, the standard strip (`-s -w`) is sufficient.

---

### 324. How do you handle secrets in Go apps deployed via Docker?
"I strictly follow the rule: **Never bake secrets into the image.**

I inject them at runtime as **Environment Variables**.
In Docker Compose, I use an `.env` file. In Kubernetes, I use `Secret` objects mapped to env vars.

My Go app reads them using `os.Getenv("DB_PASS")`. This separation of config and code allows the same image to run in Dev, Staging, and Prod safely."

#### Indepth
Avoid passing secrets as **Build Args** (`docker build --build-arg SECRET=...`). They persist in the Docker image history and can be recovered by anyone who checks `docker history`. Use **Docker Secrets** or runtime environment variables only.

---

### 325. How do you use environment variables in Go?
"I use the standard library `os.Getenv(key)`.

However, raw `Getenv` is limited because it returns an empty string if the key is missing.
For production apps, I use a configuration library like **kelseyhightower/envconfig** or **Viper**.
These libraries automatically map environment variables (like `APP_PORT=8080`) to struct fields (`Config.Port`), handling type conversion and default values for me."

#### Indepth
When using env vars, define a precedent order. **Flag > Env > Config File > Default**. This allows you to override a specific setting (like `LOG_LEVEL=debug`) quickly when debugging a container without rebuilding the image or changing the deployment implementation.

---

### 326. How do you compile a static Go binary for Alpine Linux?
"I must explicitly disable CGO.

`CGO_ENABLED=0 GOOS=linux go build`.
Alpine Linux uses `musl` instead of the standard `glibc`. If I rely on Cgo, the binary will crash on Alpine with 'no such file or directory'. By disabling Cgo, I force the compiler to use pure Go implementations for DNS and User lookups, making the binary truly portable across any Linux distribution."

#### Indepth
Alpine's `musl` libc has some quirks compared to `glibc`, especially with DNS resolution. Startups sometimes fail in Alpine due to these differences. If you hit weird DNS issues, switch to `gcr.io/distroless/static-debian11`, which is slightly larger but uses standard glibc compatibility.

---

### 327. What is `scratch` image in Docker and why is it used with Go?
"`scratch` is a special, empty Docker image. It contains absolutely nothing—no shell, no `/bin/ls`, no libraries.

It is the gold standard for Go deployments because Go binaries don't need an OS to run.
The security benefit is massive: even if an attacker finds a vulnerability in my app, they cannot 'shell out' to run commands because there makes no shell. The only caveat is needing to manually copy SSL certificates (`/etc/ssl/certs`) if I make HTTPS calls."

#### Indepth
Another missing piece in `scratch` is Timezone data. If your app relies on `time.LoadLocation("America/New_York")`, it will panic. You must manually copy `/usr/share/zoneinfo` from the builder stage to the scratch image.

---

### 328. How do you manage config files in Go across environments?
"I follow the **12-Factor App** principles, so valid config comes from the environment, not files.

However, for complex nested config, I use **Viper**.
It allows a layered approach:
1.  Read `config.yaml` (Base defaults).
2.  Override with Environment Variables (`VIPER_DB_HOST`).
3.  Override with Command Line flags.

This gives me the convenience of a file for local dev and the flexibility of env vars for Kubernetes."

#### Indepth
If you use Kubernetes ConfigMaps mounted as files, Viper allows **Hot Reloading** (`viper.WatchConfig()`). It detects when the file changes and updates the struct. This allows you to change the log level to DEBUG in production without restarting the pod!

---

### 329. How do you build Go binaries for different OS/arch?
"Cross-compilation is built right into the `go` tool.

I just set the `GOOS` and `GOARCH` environment variables.
To build for a standard server: `GOOS=linux GOARCH=amd64 go build`.
To build for a new MacBook: `GOOS=darwin GOARCH=arm64 go build`.

I don't need to install any extra toolchains or compilers. It’s one of Go’s superpowers given how painful this is in C++."

#### Indepth
Be careful with **CGO**. If you import `net` (which uses CGO for DNS by default on some OSs) or `sqlite3`, simple cross-compilation breaks. You typically need a designated "Cross-Compiler Docker Image" provided by projects like `goreleaser/goreleaser-cross` to handle the C dependencies.

---

### 330. How do you use GoReleaser?
"**GoReleaser** is my automation tool for shipping binaries.

I define a `.goreleaser.yaml` file. When I push a new Git tag (e.g., `v1.0.0`), GitHub Actions triggers GoReleaser.
It automatically builds binaries for every OS/Arch combination, zips them, generates checksums, creates a GitHub release, and even creates Homebrew recipes or pushes Docker images. It turns a manual 2-hour process into a background job."

#### Indepth
GoReleaser also supports **Snapshot Releases**. You can run `goreleaser release --snapshot --rm-dist` locally to build all artifacts and verify your release process works *without* actually pushing anything to GitHub. This is crucial for testing your build pipeline.

---

### 331. What is a Docker healthcheck for a Go app?
"It’s a command that runs inside the container to verify the app is working.

In my `Dockerfile`, I add:
`HEALTHCHECK --interval=30s CMD curl -f http://localhost:8080/health || exit 1`.

My Go app exposes a simple `/health` endpoint that returns 200 OK. If the database connection drops, I might return 500, causing Docker (or Kubernetes) to mark the container as unhealthy and restart it."

#### Indepth
Since `scratch` images don't have `curl`, you cannot use `CMD curl ...`. You must compile a tiny standalone binary (like `grpc-health-probe` or your own 50-line Go client) and copy it into the image to act as the healthheck executable.

---

### 332. How do you log container stdout/stderr from Go?
"I simply write to `os.Stdout` (for info) and `os.Stderr` (for errors).

In containerized environments, we **do not** write to log files.
The container runtime (Docker Engine) captures these streams. I then use a log shipper like **Fluentd** or **Filebeat** to read the Docker logs and forward them to a central system like Elasticsearch. This decouples the application from the logging infrastructure."

#### Indepth
Multi-line logs (like Stack Traces) are a pain in containers because Docker treats each line as a separate log event. Use structured logging (JSON) so that the entire stack trace is encapsulated in a single `"error": "..."` field, keeping the log entry atomic.

---

### 333. How do you set up autoscaling for Go services?
"I use the **Horizontal Pod Autoscaler (HPA)** in Kubernetes.

I define a CPU target, say 60%.
Because Go uses specialized goroutines, CPU usage is a very reliable proxy for load. When traffic spikes, CPU rises, and the HPA spins up new Pods.
Go apps start in milliseconds, so this scaling is highly responsive, unlike Java apps which might struggle with cold-start times during a surge."

#### Indepth
Don't just scale on CPU. Scale on **Custom Metrics** like "RabbitMQ Queue Depth". If the queue has 10,000 pending jobs, but CPU is low (because the workers are just waiting on IO), CPU scaling won't trigger. KEDA (Kubernetes Event-driven Autoscaling) is the standard tool for this.

---

### 334. How would you containerize a gRPC Go service?
"It’s similar to HTTP, but I expose the gRPC port (e.g., 50051).

The tricky part is the health check. standard `curl` doesn't speak `gRPC`.
I install **grpc_health_probe** in the container image.
My Go app implements the standard gRPC Health Checking Protocol. The probe calls this service, ensuring that not only is the TCP port open, but the gRPC server is ready to accept requests."

#### Indepth
When using `grpc_health_probe`, make sure to handle the **Shutdown Gracefully**. When a Pod terminates, it should set its health status to `SERVING_STATUS_NOT_SERVING` immediately, so the Load Balancer stops sending traffic *while* the existing requests finish (Grace Period).

---

### 335. How to deploy Go microservices in Kubernetes?
"I define three core Kubernetes objects:
1.  **Deployment**: Manages the Pods (replicas, image version, resources).
2.  **Service**: Provides a stable ClusterIP and DNS name for internal discovery.
3.  **ConfigMap/Secret**: Injects configuration via environment variables.

For zero-downtime deployments, I use a R ollingUpdate strategy, which Kubernetes handles natively because Go apps shut down gracefully on `SIGTERM`."

#### Indepth
Configure **Pod Disruption Budgets (PDB)**. This ensures that during a cluster upgrade (node draining), Kubernetes won't kill all your replicas at once. `minAvailable: 1` guarantees at least one pod is always up, preserving 100% availability.

---

### 336. How do you write Helm charts for a Go app?
"**Helm** allows me to template my Kubernetes manifests.

Instead of hardcoding `replicas: 3`, I write `replicas: {{ .Values.replicaCount }}`.
This lets me use the exact same templates for Dev (1 replica) and Prod (10 replicas) just by supplying a different `values.yaml` file. It packages my Go app deployment logic into a reusable artifact."

#### Indepth
Helm Charts can be complex. Keep them simple. Use **Library Charts** or **Subcharts** for shared logic (like standard probes, common labels, sidecars) to avoid copy-pasting code across 50 different microservice charts.

---

### 337. How do you monitor a Go service in production?
"I focus on the **RED** method: **Rate** (RPM), **Errors** (5xx), and **Duration** (Latency).

I instrument the code using **Prometheus**. I use middleware that wraps every HTTP handler, measuring how long it takes and recording the status code.
Prometheus scrapes my service’s `/metrics` endpoint every 15 seconds, and I visualize the RED metrics on Grafana dashboards."

#### Indepth
Be careful with **Cardinality Explosion**. If you add a label `path="/users/123"`, you create a new metric series for *every user*. This will crash Prometheus. Always normalize paths: `path="/users/:id"`.

---

### 338. How do you use Prometheus with a Go app?
"I use the `prometheus/client_golang` library.

I define metrics like `var httpRequests = promauto.NewCounter(...)`.
In my code, I call `httpRequests.Inc()`.
I then expose a standard HTTP handler `promhttp.Handler()` on a dedicated port (e.g., `:9090`).
This 'pull model' is robust because if my app is under heavy load, monitoring doesn't block critical request processing."

#### Indepth
For short-lived Batch Jobs (cron jobs), the Pull Model fails (the job dies before Prometheus scrapes it). Use the **Pushgateway**. The Go job pushes metrics to the Gateway just before exiting, and Prometheus scrapes the Gateway.

---

### 339. How do you enable structured logging in production?
"I use **Zap** or **slog** (Go 1.21+) with a JSON handler.

`logger := slog.New(slog.NewJSONHandler(os.Stdout))`
This outputs log lines as JSON objects: `{"level":"info", "msg":"user login", "user_id":42}`.
Structure is non-negotiable in production because it allows me to index and query logs by specific fields in tools like Splunk or Datadog, rather than grepping text."

#### Indepth
Include the **Trace ID** and **Span ID** in every log line automatically. This allows you to correlate "Error in DB" with "User Request X". In `slog`, you can use `slog.With("trace_id", ctx.TraceID())` to inject this context into the logger.

---

### 340. How do you handle log rotation in containerized Go apps?
"**I do not handle it in the app.**

The application should effectively write infinite logs to Stdout.
Log rotation is the responsibility of the execution environment. In Docker, I configure the `json-file` logging driver with `max-size` and `max-file` options.
If I tried to handle rotation inside the Go app, I’d run into concurrency issues and risk losing logs during rotation."

#### Indepth
In legacy environments (VMs) where you *must* log to a file, use an external tool like `logrotate` (standard Linux utility) combined with the `SIGHUP` signal. When `logrotate` moves the file, it sends `SIGHUP` to your app, telling it to reopen the file handle.
