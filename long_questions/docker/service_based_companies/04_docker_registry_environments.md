# 📦 Docker Registries, Image Tagging & Environment Management (Service-Based Companies)

This document covers Docker registries, image lifecycle management, and multi-environment deployment strategies — commonly tested at service-based company interviews for senior/lead roles.

---

### Q1: What is a Docker registry? How do you push and pull images from Docker Hub and a private registry?

**Answer:**
A **Docker registry** is a storage and distribution system for Docker images. It's analogous to GitHub for source code — but for container images.

**Public registries:**
- **Docker Hub** (`hub.docker.com`) — The default registry; `docker pull nginx` implicitly pulls from `docker.io/library/nginx:latest`
- **GitHub Container Registry** (`ghcr.io`)
- **Google Artifact Registry** (`us-docker.pkg.dev`)
- **Amazon ECR Public** (`public.ecr.aws`)

**Private registries:**
- **Amazon ECR** (Elastic Container Registry)
- **Azure Container Registry (ACR)**
- **Google Artifact Registry (private)**
- **Harbor** (self-hosted, open-source)
- **JFrog Artifactory**
- **GitLab Container Registry**

**Push/pull workflow:**
```bash
# 1. Login to Docker Hub
docker login                          # Prompts for username/password
docker login -u myuser -p mypass      # Non-interactive (avoid — password in history)
echo $DOCKER_PASS | docker login -u myuser --password-stdin  # Secure way

# 2. Tag your image for the registry
docker tag myapp:latest myusername/myapp:v1.0.0
# Format: <registry>/<repository>:<tag>
# For Docker Hub: <dockerhub-username>/<image-name>:<tag>

# 3. Push
docker push myusername/myapp:v1.0.0

# 4. Pull (on another machine)
docker pull myusername/myapp:v1.0.0
```

**Private registry (e.g., AWS ECR):**
```bash
# Authenticate with ECR
aws ecr get-login-password --region ap-south-1 | \
  docker login --username AWS --password-stdin 123456789.dkr.ecr.ap-south-1.amazonaws.com

# Tag for ECR
docker tag myapp:latest 123456789.dkr.ecr.ap-south-1.amazonaws.com/myapp:v1.0.0

# Push
docker push 123456789.dkr.ecr.ap-south-1.amazonaws.com/myapp:v1.0.0
```

---

### Q2: What are the best practices for Docker image tagging and versioning?

**Answer:**
Proper tagging is essential for reproducible deployments and rollbacks.

**Common tagging strategies:**

| Tag Pattern | Example | Use Case |
|---|---|---|
| `latest` | `myapp:latest` | Development convenience — **avoid in production** (mutable, unpredictable) |
| Semantic version | `myapp:v1.4.2` | Production releases — immutable, meaningful |
| Git commit SHA | `myapp:git-a3b9f12` | CI/CD — tied directly to code, perfect traceability |
| Branch name | `myapp:main` | Auto-deployed branch preview |
| Build number | `myapp:build-1234` | CI platform-based tracking |
| Date-based | `myapp:2026.03.04` | Time-referenced deployments |

**Why `latest` is problematic in production:**
- It's mutable — `docker pull myapp:latest` today gets a different image than tomorrow
- No rollback capability — you can't `docker run myapp:latest` and know what version you're running
- No audit trail

**Best practice — multi-tag approach:**
```bash
# After CI build on main branch:
docker build -t myapp:git-$GIT_SHA .
docker tag myapp:git-$GIT_SHA myapp:v1.4.2       # Semver release
docker tag myapp:git-$GIT_SHA myapp:latest        # Convenience (optional)

docker push myapp:git-$GIT_SHA
docker push myapp:v1.4.2
docker push myapp:latest

# Production deployment references: myapp:v1.4.2 (never :latest)
# Rollback: change to myapp:v1.4.1 (old tag still in registry)
```

**Immutable tag enforcement:**
Most registries support immutable tags — once pushed, a tag like `v1.4.2` cannot be overwritten. Enable this in ECR or Harbor to prevent accidental overwrites.

---

### Q3: How do you manage different configurations for different environments (dev, staging, production) in Docker?

**Answer:**
The **12-Factor App** principle says: store config in the environment. Docker supports this with multiple approaches:

**Approach 1: Environment-specific env files**
```bash
# File structure
.env.development
.env.staging
.env.production

# .env.production
DB_HOST=prod-db.company.com
DB_PORT=5432
LOG_LEVEL=warn
MAX_POOL_SIZE=20

# Run with specific env file
docker run --env-file .env.production myapp:v1.4.2
```

**Approach 2: Docker Compose with environment-specific override files**
```bash
# Base config (shared)
compose.yml

# Environment overrides
compose.dev.yml      # Mount source code, enable debug ports
compose.staging.yml  # Use staging secrets and DB
compose.prod.yml     # Resource limits, restart policies

# Run for dev:
docker compose -f compose.yml -f compose.dev.yml up

# Run for production:
docker compose -f compose.yml -f compose.prod.yml up -d
```

**`compose.dev.yml` example:**
```yaml
services:
  app:
    build:
      context: .
      target: development    # Multi-stage Dockerfile target
    volumes:
      - ./src:/app/src       # Hot reload
    environment:
      - LOG_LEVEL=debug
    ports:
      - "9229:9229"          # Node.js debugger port
```

**`compose.prod.yml` example:**
```yaml
services:
  app:
    image: myrepo/myapp:${APP_VERSION}   # No build — use pre-built image
    restart: unless-stopped
    deploy:
      resources:
        limits:
          memory: 512M
          cpus: "0.5"
    logging:
      driver: awslogs
      options:
        awslogs-group: /myapp/production
        awslogs-region: ap-south-1
```

---

### Q4: How does Docker handle multi-architecture images and why does this matter?

**Answer:**
Modern infrastructure uses a mix of CPU architectures:
- **`linux/amd64`** — Standard x86_64 servers, most cloud VMs
- **`linux/arm64`** — AWS Graviton, Apple M-series, Raspberry Pi 4
- **`linux/arm/v7`** — Raspberry Pi 3, IoT devices

A single Docker image tag can point to a **manifest list** (multi-arch image) that contains separate images for each architecture. When you `docker pull`, Docker automatically fetches the right one for the host's CPU.

**Check if an image is multi-arch:**
```bash
docker manifest inspect nginx:latest
# Shows list of manifests for amd64, arm64, s390x, etc.
```

**Why it matters in practice:**
- Developers on Apple M-series Macs (arm64) can't run amd64-only images natively without emulation
- AWS Graviton instances (arm64) are ~40% cheaper — if your image is amd64-only, you can't use them
- IoT deployments on Raspberry Pi need arm images

**Building multi-arch images:**
```bash
# Setup multi-arch builder
docker buildx create --use --name multi-builder
docker buildx inspect --bootstrap

# Build and push multi-arch image
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t myrepo/myapp:v1.4.2 \
  --push .
```

**Check your current image's architecture:**
```bash
docker inspect myapp:latest --format '{{.Architecture}}'
docker image inspect myapp:latest | grep Architecture
```

---

### Q5: How do you implement Docker in a CI/CD pipeline? Walk through a complete end-to-end example.

**Answer:**
Here's a complete Jenkins/GitHub Actions CI/CD pipeline for a Docker-based application:

**GitHub Actions example (`.github/workflows/docker.yml`):**
```yaml
name: Build and Deploy

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-test-push:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      # 1. Checkout code
      - uses: actions/checkout@v4

      # 2. Set up Docker BuildKit
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      # 3. Login to registry
      - name: Log in to GitHub Container Registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      # 4. Extract metadata (determines tags and labels)
      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
          tags: |
            type=sha,prefix=git-
            type=semver,pattern={{version}}
            type=raw,value=latest,enable={{is_default_branch}}

      # 5. Run unit tests (before building final image)
      - name: Run tests
        run: |
          docker build --target test -t myapp:test .
          docker run --rm myapp:test

      # 6. Build and push (with registry cache)
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: ${{ github.event_name == 'push' }}
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max

      # 7. Scan for vulnerabilities
      - name: Security scan
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}:git-${{ github.sha }}
          exit-code: 1
          severity: CRITICAL

  deploy-staging:
    needs: build-test-push
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Deploy to staging
        run: |
          # Update image tag in Kubernetes/Compose config
          # Trigger ArgoCD sync, or SSH + docker-compose pull + up
          ssh deployer@staging-server "
            docker compose pull &&
            docker compose up -d --no-deps app
          "
```

---

### Q6: What is the difference between `CMD` and `ENTRYPOINT`? When would you use both together?

**Answer:**
This is one of the most frequently asked Docker questions. The confusion arises because both define what runs when a container starts.

**`CMD`:**
- Provides **default arguments** for the container.
- Can be **completely overridden** by passing arguments to `docker run`.
- If multiple `CMD` instructions exist, only the last one takes effect.

```dockerfile
CMD ["nginx", "-g", "daemon off;"]
```
```bash
docker run myimage                   # Runs: nginx -g daemon off;
docker run myimage sh                # Runs: sh  ← CMD completely replaced
```

**`ENTRYPOINT`:**
- Sets the **fixed executable** for the container.
- Cannot be overridden by normal `docker run <args>` — the args are passed as arguments *to* the entrypoint.
- Can only be overridden with `docker run --entrypoint`.

```dockerfile
ENTRYPOINT ["nginx"]
```
```bash
docker run myimage                   # Runs: nginx
docker run myimage -t                # Runs: nginx -t  ← appended as arg
docker run --entrypoint sh myimage   # Runs: sh  ← explicit override
```

**Using both together (recommended pattern):**
```dockerfile
ENTRYPOINT ["python3", "app.py"]   # Fixed executable
CMD ["--port", "8080"]             # Default arguments — easily overridable
```
```bash
docker run myimage                           # Runs: python3 app.py --port 8080
docker run myimage --port 9090               # Runs: python3 app.py --port 9090
docker run --entrypoint sh myimage           # Runs: sh (full override)
```

**Shell form vs Exec form:**
```dockerfile
# Shell form — runs in /bin/sh -c, so signals (SIGTERM) may not reach your process
CMD python app.py
ENTRYPOINT python app.py

# Exec form (recommended) — runs directly, proper signal handling
CMD ["python", "app.py"]
ENTRYPOINT ["python", "app.py"]
```

Always prefer the **exec form** `["..."]` in production so Docker can properly send SIGTERM to your process on `docker stop`.

---

*Prepared for L2/L3 technical rounds and DevOps screening at service-based companies.*
