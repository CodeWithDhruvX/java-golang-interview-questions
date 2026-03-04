# 🚀 Docker CI/CD, BuildKit & Advanced Build Strategies (Product-Based Companies)

This document covers Docker's advanced build system, CI/CD patterns, and pipeline integration — topics frequently tested at product-based companies that run large-scale build infrastructure.

---

### Q1: What is Docker BuildKit, and how is it fundamentally different from the legacy builder?

**Answer:**
**BuildKit** is Docker's next-generation image build engine, introduced as opt-in in Docker 18.09 and enabled by default in Docker 23+. It's a complete rewrite of the build daemon.

**Key architectural differences vs. legacy builder:**

| Feature | Legacy Builder | BuildKit |
|---|---|---|
| Parallelism | Sequential layer execution | Concurrent stage execution |
| Cache | Local only | Exportable, importable, registry-backed |
| Secrets | Baked into layers (dangerous!) | Ephemeral, never stored in layers |
| SSH forwarding | Not supported | Native `--ssh` flag |
| Build output | Always image | Image, local files, tar, OCI layout |
| Pruning | Manual | Automatic GC with configurable limits |
| Cross-platform | Limited | Native `--platform` multi-arch |
| Garbage collection | Docker prune | Configurable policies |

**Enable BuildKit (legacy method):**
```bash
DOCKER_BUILDKIT=1 docker build .
# OR
export DOCKER_BUILDKIT=1
```

**Using the new `docker buildx` (recommended):**
```bash
docker buildx build --platform linux/amd64,linux/arm64 -t myapp:latest --push .
```

**Parallel stage execution example:**
```dockerfile
# Stage A and Stage B run in PARALLEL with BuildKit
FROM golang:1.22-alpine AS test-stage
RUN go test ./...

FROM golang:1.22-alpine AS build-stage
RUN go build -o /app ./cmd/server
```
BuildKit analyses the dependency graph and runs non-dependent stages concurrently, cutting build times significantly in complex multi-stage Dockerfiles.

---

### Q2: How do you use BuildKit's secret mount to inject sensitive credentials during build without leaking them into a layer?

**Answer:**
The classic anti-pattern: developers run `RUN npm install` using a private registry that needs authentication credentials, so they `ENV NPM_TOKEN=secret` — that token is now **baked into the image layer permanently**, visible to anyone who pulls the image.

**BuildKit's `--mount=type=secret` solves this:**

**Dockerfile:**
```dockerfile
# syntax=docker/dockerfile:1.4
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./

# The secret is mounted as a read-only tmpfs at build time
# It is NEVER written to any image layer
RUN --mount=type=secret,id=npm_token \
    NPM_TOKEN=$(cat /run/secrets/npm_token) \
    npm config set "//registry.npmjs.org/:_authToken=$NPM_TOKEN" && \
    npm ci --only=production

COPY . .
RUN npm run build
```

**Build command:**
```bash
# Pass secret from file
docker build --secret id=npm_token,src=~/.npm_token -t myapp:latest .

# Pass secret from env var
docker build --secret id=npm_token,env=NPM_TOKEN -t myapp:latest .
```

**Verification — the secret is NOT in the image:**
```bash
docker history myapp:latest    # No token visible
docker inspect myapp:latest    # No env var present
dive myapp:latest              # No secret in any layer filesystem
```

**SSH forwarding for private Git repos:**
```dockerfile
RUN --mount=type=ssh \
    git clone git@github.com:myorg/private-lib.git /deps/private-lib
```
```bash
docker build --ssh default=$SSH_AUTH_SOCK -t myapp:latest .
```

---

### Q3: Compare Docker-in-Docker (DinD) vs. Docker socket mounting vs. Kaniko for CI/CD image builds. When would you use each?

**Answer:**
Building Docker images inside a CI runner (which is itself a container) is a core DevOps challenge. Three main approaches exist, each with different security postures:

**1. Docker-in-Docker (DinD)**
Runs a nested Docker daemon inside the CI container.
```yaml
# GitLab CI
build:
  image: docker:24
  services:
    - docker:24-dind
  variables:
    DOCKER_HOST: tcp://docker:2376
    DOCKER_TLS_CERTDIR: "/certs"
  script:
    - docker build -t myapp:$CI_COMMIT_SHA .
```
*Pros:* Clean isolation, builds don't interfere with host.  
*Cons:* Requires `--privileged` flag on the CI container (significant security risk), nested layer storage overhead.

**2. Docker Socket Mounting**
Mounts the host's Docker socket into the CI container.
```yaml
build:
  image: docker:24
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock
  script:
    - docker build -t myapp:$CI_COMMIT_SHA .
```
*Pros:* No privilege escalation, fast — uses host daemon directly.  
*Cons:* **Major security risk** — the CI container gains full root on the host through the socket. Build containers can see all host containers.

**3. Kaniko**
A purpose-built tool (by Google) that builds images without a Docker daemon, running entirely as a normal (non-root, non-privileged) container.
```yaml
build:
  image:
    name: gcr.io/kaniko-project/executor:latest
    entrypoint: [""]
  script:
    - /kaniko/executor
      --context $CI_PROJECT_DIR
      --dockerfile $CI_PROJECT_DIR/Dockerfile
      --destination myrepo/myapp:$CI_COMMIT_SHA
```
*Pros:* No Docker daemon, no `--privileged`, works in locked-down Kubernetes clusters.  
*Cons:* Slower than native Docker builds, some Dockerfile features may behave differently, cache management is more complex.

**Decision Matrix:**

| Scenario | Recommendation |
|---|---|
| Security-sensitive production K8s cluster | **Kaniko** or **Buildah** |
| Internal trusted CI/CD with performance priority | **Socket mount** (if host is trusted) |
| Full isolation needed, security not a top priority | **DinD** |
| Multi-platform builds | **`docker buildx` with BuildKit** |

---

### Q4: How does Docker layer caching work with a registry-backed cache in a distributed CI/CD system?

**Answer:**
In a distributed CI system where multiple ephemeral runners spin up fresh environments for each build, the local filesystem cache is lost between runs. **Registry-backed cache** solves this.

**How it works:**
BuildKit can export the build cache as OCI image manifests to a container registry, and import it on the next build:

```bash
# First build — export cache to registry
docker buildx build \
  --cache-to type=registry,ref=myrepo/myapp:buildcache,mode=max \
  --cache-from type=registry,ref=myrepo/myapp:buildcache \
  -t myrepo/myapp:latest --push .

# Subsequent builds on any runner — imports cache from registry
docker buildx build \
  --cache-from type=registry,ref=myrepo/myapp:buildcache \
  --cache-to type=registry,ref=myrepo/myapp:buildcache,mode=max \
  -t myrepo/myapp:latest --push .
```

**`mode=max` vs `mode=min`:**
- `mode=min` (default): Only exports the cache for the final stage.
- `mode=max`: Exports cache for **all intermediate stages**. More data, but dramatically faster rebuilds for complex multi-stage Dockerfiles.

**GitHub Actions example:**
```yaml
- name: Build and push
  uses: docker/build-push-action@v5
  with:
    context: .
    push: true
    tags: myrepo/myapp:${{ github.sha }}
    cache-from: type=gha          # GitHub Actions Cache service
    cache-to: type=gha,mode=max   # Built-in free cache (~10GB)
```

**Other cache backends:** `type=s3` (AWS S3), `type=azblob` (Azure Blob), `type=local` (shared NFS mount).

---

### Q5: Describe multi-platform image builds with `docker buildx` and what happens under the hood with QEMU emulation.

**Answer:**
**The problem:** Your CI runs on `linux/amd64` but you deploy to AWS Graviton (`linux/arm64`) or Raspberry Pi (`linux/arm/v7`). The compiled binary must match the target architecture.

**Two strategies for multi-platform builds:**

**Strategy 1: QEMU Emulation (Simpler, Slower)**
```bash
# Enable QEMU support (one-time setup)
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes

# Create a multi-platform builder
docker buildx create --name multibuilder --use
docker buildx inspect --bootstrap

# Build for multiple platforms simultaneously
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t myrepo/myapp:latest \
  --push .
```

**What happens with QEMU:** The BuildKit daemon uses QEMU (a processor emulator registered with Linux `binfmt_misc`) to execute ARM instructions on an x86 host. Every RUN instruction for the ARM target runs under emulation — functional but **5-10x slower** for CPU-intensive build steps.

**Strategy 2: Native Nodes (Faster, More Complex)**
```bash
# Register actual ARM64 machine as a builder node
docker buildx create --name multibuilder \
  --platform linux/amd64 \
  --node amd64-node ssh://ci-runner-amd64

docker buildx create --name multibuilder --append \
  --platform linux/arm64 \
  --node arm64-node ssh://ci-runner-arm64

# Now each platform builds natively on matching hardware
docker buildx build --platform linux/amd64,linux/arm64 \
  -t myrepo/myapp:latest --push .
```

**Output — OCI Image Index:**
The push creates a **manifest list** in the registry. When a client runs `docker pull myrepo/myapp:latest`, Docker inspects the manifest list and automatically pulls the correct architecture variant.

```json
{
  "mediaType": "application/vnd.oci.image.index.v1+json",
  "manifests": [
    { "platform": { "os": "linux", "architecture": "amd64" }, "digest": "sha256:abc..." },
    { "platform": { "os": "linux", "architecture": "arm64" }, "digest": "sha256:def..." }
  ]
}
```

---

### Q6: How would you implement a GitOps-style Docker image promotion pipeline across dev → staging → production with immutable tags?

**Answer:**
A robust promotion strategy uses **immutable, content-addressed tags** rather than mutable ones like `latest` or `staging`.

**Tagging strategy:**
```
myrepo/myapp:git-{short-sha}      # e.g., git-a3b9f12 — immutable, tied to source
myrepo/myapp:build-{build-number} # e.g., build-1234 — immutable, CI run
myrepo/myapp:v1.4.2               # Semantic version — immutable release
myrepo/myapp:latest                # Mutable pointer — anti-pattern for GitOps, avoid in prod
```

**Pipeline stages:**

```
[Dev Branch Push]
    ↓ Build + Unit Tests
    → Tag: git-{sha}, push to dev registry (dev.registry.company.com)
    
[PR Merge to main]
    ↓ Build + Integration Tests + Security Scan (Trivy --exit-code 1)
    → Tag: git-{sha}, push to staging registry
    → Update GitOps repo: staging/values.yaml image.tag = git-{sha}
    → ArgoCD/FluxCD detects change → deploys to staging
    
[Manual/Automated Promotion Gate]
    ↓ E2E tests pass + approval
    → Retag: docker buildx imagetools create 
              --tag prod.registry.company.com/myapp:v1.4.2
              staging.registry.company.com/myapp:git-a3b9f12
    → Update GitOps repo: prod/values.yaml image.tag = v1.4.2
    → ArgoCD → deploys to production
```

**Key principle:** The **same artifact** (same SHA256 digest) is promoted through environments. Nothing is rebuilt. This guarantees that what you tested in staging is exactly what runs in production — diff only the config, never the binary.

**Retagging without rebuild (using `imagetools`):**
```bash
# This doesn't pull/push the layers, only manipulates manifests
docker buildx imagetools create \
  --tag prod.registry.company.com/myapp:v1.4.2 \
  staging.registry.company.com/myapp:git-a3b9f12
```

---

*Prepared for DevOps Architecture and Platform Engineering interviews at product-based companies.*
