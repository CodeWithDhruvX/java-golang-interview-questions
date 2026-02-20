# 🧪 **Docker Testing, CI/CD & Logging (51–60)**

---

### 51. How do you test a Docker image?
"I test at multiple levels. First, **structure tests** using **Container Structure Test** (by Google) — they verify that specific files exist, commands return expected output, and environment variables are set.

Second, **integration tests** — I run the image and execute tests against it using Docker Compose. Third, **vulnerability scans** with Trivy or Docker Scout as part of the pipeline.

A simple smoke test: `docker run --rm myapp:latest --version` or `docker run --rm myapp:latest healthcheck-script.sh`."

#### In-depth
Container Structure Tests are defined in YAML and check: file existence, file contents, command output, environment variables, and metadata. They're fast (run in seconds), language-agnostic, and catch common Dockerfile mistakes — like a missing entry point, wrong working directory, or forgotten file copy.

---

### 52. How is Docker used in CI/CD pipelines?
"Docker is the backbone of modern CI/CD. The pattern: **build → test → scan → push → deploy**.

1. `docker build -t myapp:$GIT_SHA .` — build with commit SHA as tag
2. `docker run --rm myapp:$GIT_SHA run-tests` — run tests inside the image
3. `trivy image myapp:$GIT_SHA` — security scan
4. `docker push myregistry/myapp:$GIT_SHA` — push to registry
5. Deploy by updating the service to use the new image tag

The same image that passes tests is the exact image deployed — no 'it worked in CI' surprises."

#### In-depth
The immutability of Docker images is what makes this model reliable. The SHA256-tagged image is a content-addressed artifact — what you tested is exactly what runs in production. Contrast with traditional deployments where `git pull && npm install` on the server could differ from the CI environment due to transient dependencies or package registry changes.

---

### 53. How can you log container activity?
"Containers write logs to **stdout and stderr** — this is the Docker logging contract. Docker captures these streams and makes them available via `docker logs`.

The logging driver determines where Docker sends these logs: `json-file` (default, writes to host filesystem), `syslog`, `journald`, `fluentd`, `awslogs`, `gelf`, and more.

For production, I configure a centralized logging driver from the start — `awslogs` for AWS, `fluentd` for ELK/Loki stacks. Writing to files inside containers is an anti-pattern."

#### In-depth
The 12-factor app methodology says: **logs are streams**. Applications should write to stdout/stderr, and the infrastructure handles routing. This decouples the app from the logging backend. Changing from Elasticsearch to Loki becomes a config change in the Docker daemon, not an application code change.

---

### 54. How do you access logs from a Docker container?
"I use `docker logs <container>`.

Key flags: `-f` (follow/tail), `--since=1h` (last hour), `--tail=100` (last 100 lines), `--timestamps` (show timestamps).

Combined: `docker logs -f --since=5m --timestamps myapi` — follow logs from the last 5 minutes with timestamps.

For Docker Compose: `docker compose logs -f api db` — follow logs from specific services."

#### In-depth
`docker logs` only works with the default `json-file` and `journald` drivers. If you configure a remote logging driver (like `awslogs`), `docker logs` returns nothing — logs go directly to the remote backend. This surprises many developers. Always test your logging setup in a dev environment before going to production.

---

### 55. What is the difference between `docker logs` and `docker attach`?
"`docker logs` retrieves **past and live output** from a container's stdout/stderr without connecting to it. It's non-interactive and read-only.

`docker attach` **connects your terminal to the container's stdin/stdout/stderr**. You're essentially joining the same TTY session. If the process is interactive, you can type into it. Pressing `Ctrl+C` will send SIGINT to the container process — which may stop the container.

For debugging, I prefer `docker logs -f` or `docker exec -it container /bin/sh` over `attach`."

#### In-depth
`docker attach` connects to PID 1's standard streams. If PID 1 is your application (not a shell), pressing Enter in the attached session sends data to the application's stdin — which is usually unexpected. Use `--sig-proxy=false` with `attach` to prevent `Ctrl+C` from killing the container. `docker exec` is almost always safer for interactive debugging.

---

### 56. How do you debug a running container?
"My debugging toolkit:

1. `docker exec -it container /bin/sh` — open an interactive shell
2. `docker logs -f container` — follow logs
3. `docker inspect container` — full JSON metadata (network, mounts, env, etc.)
4. `docker stats container` — real-time CPU, memory, I/O metrics
5. `docker top container` — running processes inside
6. `docker cp container:/app/logs ./` — extract log files

For minimal containers without a shell, I use `docker debug` (Docker Desktop feature) which attaches a debug toolkit without modifying the image."

#### In-depth
`docker debug` (added in Docker Desktop 4.27) spins up a privileged debug container sharing the namespace of the target container — giving you `bash`, `curl`, `strace`, `tcpdump` etc. without modifying the production image. Same concept as Kubernetes ephemeral debug containers. For older setups, I add a second `FROM` stage in the Dockerfile specifically for debug builds (`FROM myapp AS debug; RUN apk add curl`).

---

### 57. What tools integrate with Docker for CI/CD?
"The major CI platforms all have first-class Docker support:

**GitHub Actions**: `docker/build-push-action`, `docker/login-action` — official Docker actions. **GitLab CI**: Docker-in-Docker or socket binding, registry built-in. **Jenkins**: Docker Pipeline plugin, Docker agents. **CircleCI**: Docker executors with layer caching. **Buildkite**: Docker Compose plugin.

Beyond CI, **Argo CD** does GitOps continuous delivery of Docker images to Kubernetes, and **Watchtower** auto-updates containers when new images are pushed."

#### In-depth
GitHub Actions + Docker's official actions is the most common modern setup. The `docker/build-push-action` with `cache-from` and `cache-to` using **GitHub's cache API** gives you cross-run layer caching without a separate cache registry. Build times drop from 5-10 minutes to 30-60 seconds for incremental changes.

---

### 58. What are some common Docker CI tools?
"Beyond the CI platforms themselves, key tools in a Docker CI pipeline:

**Hadolint**: Dockerfile linter — catches best practice violations. **Trivy/Grype**: Image vulnerability scanners. **Dive**: Image layer inspector — find unnecessary files bloating images. **Container Structure Test**: Automated image structure validation. **Docker Scout**: Docker's built-in vulnerability and SBOM tool. **Cosign**: Image signing for supply chain security.

I run Hadolint in the IDE (via VS Code extension) and in CI as a PR gate."

#### In-depth
Hadolint combines shell script linting (via ShellCheck) with Dockerfile-specific rules. It catches things like: using `apt-get` without `-y` (build hangs waiting for input), not pinning package versions, using `ADD` instead of `COPY`, missing `--no-install-recommends` (increases image size). These are all real production pitfalls.

---

### 59. What is the role of Docker in microservices architecture?
"Docker is the runtime foundation of microservices. Each microservice gets its **own container**: isolated runtime, independent scaling, independent deployment.

Docker enables: **technology heterogeneity** (Java service in one container, Go service in another), **independent deployments** (change service A without touching B), **resource isolation** (one service's memory leak doesn't crash others), and **local development parity** (Compose runs all 20 services locally).

Without Docker, microservices become an operational nightmare — each service needs its own VM or complex dependency management."

#### In-depth
The container model aligns naturally with microservices patterns. **Sidecar pattern**: attach a logging/monitoring container alongside the main service container, sharing its network namespace. **Ambassador pattern**: a proxy container handles service discovery/retry logic. **Init container pattern** (Kubernetes): run a setup container before the main service starts.

---

### 60. How do you manage configuration across environments in Docker?
"The 12-factor principle: **config in the environment, not in the image**. I build one image and promote it across dev → staging → prod with different configs injected at runtime.

Mechanisms: **environment variables** (simple key-value), **`.env` files** (Compose), **config files mounted as volumes** (`docker run -v ./config.prod.yaml:/app/config.yaml`), **Docker secrets** (Swarm, for sensitive values), **external config stores** (Consul, AWS Parameter Store, Vault).

Never bake environment-specific config into images. The same image must work in all environments."

#### In-depth
A common antipattern: building separate images for dev/staging/prod. This means the tested image isn't the deployed image — defeating the entire point. The correct pattern: one image, behavior driven by environment variables. Use feature flags and config backends (like LaunchDarkly or Consul) for complex multi-environment config management.

---
