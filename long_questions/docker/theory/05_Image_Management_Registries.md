# 📦 **Docker Image Management & Registries (61–70)**

---

### 61. What is the difference between a private and public Docker registry?
"A **public registry** (like Docker Hub) is accessible to anyone. Images are free to pull without authentication. Great for open-source projects.

A **private registry** requires authentication. Only authorized users can push or pull. Organizations use private registries for proprietary application images, to avoid Docker Hub rate limits, and to keep images close to their deployments (reducing latency).

Popular private registry options: **AWS ECR**, **Google Artifact Registry**, **GitHub Container Registry (GHCR)**, **JFrog Artifactory**, **Harbor** (self-hosted open-source)."

#### In-depth
For production workloads, a private registry in the same cloud region as your cluster is essential. Pulling a 500MB image from Docker Hub to a US-East production server on every deployment introduces latency and Docker Hub rate limit risk. ECR in `us-east-1` pulls at 10Gbps with no rate limits. The proximity also reduces egress costs.

---

### 62. How do you push an image to Docker Hub?
"Three steps: **login**, **tag**, **push**.

```bash
docker login -u myusername
docker tag myapp:latest myusername/myapp:v1.0
docker push myusername/myapp:v1.0
```

In CI, I use non-interactive login with a token: `docker login -u $DOCKER_USER --password-stdin <<< $DOCKER_TOKEN`. I never log in with my password directly — I use Docker Hub access tokens with scoped permissions (push-only for CI)."

#### In-depth
Docker Hub access tokens can be scoped: `Read-only`, `Read & Write`, or `Read, Write & Delete`. For CI, I create a Read & Write token so it can push but not delete. For production pull credentials (Kubernetes `imagePullSecrets`), use Read-only tokens. Rotating tokens regularly and using separate tokens per service/team is good security hygiene.

---

### 63. How do you pull an image from a private registry?
"I authenticate first: `docker login myregistry.example.com -u user -p token`.

Then pull normally: `docker pull myregistry.example.com/myapp:latest`.

For cloud registries, I use cloud CLI tools that configure authentication automatically: `aws ecr get-login-password | docker login --username AWS --password-stdin <ecr-url>`. In Kubernetes, I create `imagePullSecrets` so pods can pull from private registries without manual auth."

#### In-depth
Authentication credentials are stored in `~/.docker/config.json`. Docker uses **credential helpers** (`docker-credential-ecr-login`, `docker-credential-gcr`) for cloud registries — they automatically refresh short-lived tokens so builds don't fail when tokens expire. Setting up credentials helpers in CI is a critical but often overlooked step.

---

### 64. What is the format of Docker image tags?
"The full format is: `registry/namespace/repository:tag@digest`

- **registry**: optional, defaults to `docker.io`. E.g., `gcr.io`, `mycompany.registry.io`
- **namespace**: Docker Hub username or org. E.g., `library` for official images
- **repository**: image name. E.g., `nginx`, `myapp`
- **tag**: human-readable label. E.g., `latest`, `v1.2.3`, `stable`
- **digest**: `sha256:abc123...` — immutable reference to a specific version

For production I always pin by digest: `nginx@sha256:abc...` — the tag can be rewritten, the digest never changes."

#### In-depth
The `latest` tag is a common pitfall. It's just a tag like any other — not a guarantee of actual freshness. If someone pushes `myapp:latest` on Friday and you pull it Monday, you get the Friday image — but if they push a new `latest` and restart your container, you get a different image. For reproducibility, always use specific version tags or digests in production configs.

---

### 65. What is a dangling image?
"A dangling image is an image that has **no tag and no container using it** — a leftover artifact.

They occur when you rebuild an image with the same tag: the old image loses its tag (becomes `<none>:<none>`) but stays on disk. Over time, dangling images accumulate and waste significant disk space.

List them: `docker images -f dangling=true`. Remove them: `docker image prune`. I run this in CI post-build cleanup steps."

#### In-depth
Dangling images are different from unused images. A dangling image has no tag; an unused image has a tag but no container using it. `docker image prune` removes only dangling images. `docker image prune -a` removes all unused images (including tagged ones not used by any container) — be careful with this in development environments where you keep cached images for faster iteration.

---

### 66. How do you clean up unused Docker images?
"My cleanup toolkit:

- `docker image prune` — remove dangling images only
- `docker image prune -a` — remove all unused images (not just dangling)
- `docker image prune -a --filter 'until=72h'` — remove images older than 72 hours
- `docker system prune` — remove everything: stopped containers, unused networks, dangling images, build cache
- `docker system prune -a --volumes` — the nuclear option: removes volumes too

For CI servers I run `docker system prune -f` after each job to prevent disk exhaustion."

#### In-depth
In heavy CI environments (building 50+ images per day), disk can fill up within hours without cleanup. Kubernetes nodes with containerd have their own garbage collection (`imageGCHighThresholdPercent`). Docker Buildx cache (BuildKit) grows separately — use `docker buildx prune --filter until=48h` to clean it. Monitor disk usage with `docker system df` as part of your ops dashboards.

---

### 67. What is an image layer?
"An image layer is a **read-only filesystem snapshot** corresponding to one instruction in the Dockerfile.

Each `RUN`, `COPY`, or `ADD` instruction creates a new layer. Layers are stacked — each layer contains only the diff from the previous one. When Docker runs a container, it adds a writable layer on top.

Layers are **content-addressable by SHA256 digest** — two identical layers share storage. This is why images download fast: Docker only fetches layers you don't already have."

#### In-depth
You can inspect layers with `docker history myimage:latest` — it shows each layer's size and the command that created it. Use `dive myimage:latest` for an interactive layer explorer that shows you exactly which files each layer adds, modifies, or deletes. This is invaluable for identifying image bloat.

---

### 68. How does Docker layer caching work?
"Docker caches the result of each instruction. On rebuild, it checks if the instruction and its inputs have changed. If not — **cache hit**, it reuses the layer instead of re-executing.

Cache invalidation rules: for `RUN` instructions, the cache key is the command string. For `COPY`/`ADD`, Docker checksums the copied files. Any cache miss **invalidates all subsequent layers** — even if those instructions didn't change.

This is why instruction order matters: frequently changing things (app code) must come AFTER rarely changing things (OS packages, dependencies)."

#### In-depth
BuildKit (Docker's modern build engine) improves caching further with: parallel stage execution, fine-grained cache mounts (`RUN --mount=type=cache`), and external cache backends (`--cache-to=type=registry`). The `--mount=type=cache` is a game-changer for package managers — `apt`, `npm`, `pip` caches survive between builds without adding to image layers.

---

### 69. What is the size of a typical Docker image?
"This varies enormously by base image choice:

- **Ubuntu**: ~29MB compressed; **Debian**: ~25MB
- **Alpine**: ~3MB — the go-to for minimal images
- **Distroless**: ~1-2MB for language-base variants
- **Scratch**: 0MB — completely empty

A Java Spring Boot app on Ubuntu (with JDK): ~400-600MB. Same app on Alpine with JRE-only: ~100-150MB. With GraalVM native image + distroless: ~15-30MB.

I always aim for the smallest functional image — less storage, less attack surface, faster pulls."

#### In-depth
Image size directly impacts deployment speed (especially in auto-scaling scenarios where new nodes must pull images urgently), storage costs (ECR, GCR charge per GB stored), and security (more packages = more CVEs). The `docker images` size column shows the uncompressed size; registry shows compressed size (roughly 30-50% smaller). The real-world number for pull speed is the compressed size.

---

### 70. How do you optimize image size?
"My optimization checklist:

1. **Choose the right base image** — Alpine or distroless over Ubuntu
2. **Multi-stage builds** — compile in one stage, ship only the binary
3. **Combine RUN commands** — `RUN apt-get update && apt-get install -y pkg && rm -rf /var/lib/apt/lists/*`
4. **Don't install unnecessary tools** — no debug tools in production images
5. **Use `.dockerignore`** — don't copy test files, docs, node_modules
6. **Pin and clean up** — `apt-get clean` and remove package lists after install

Measure with `dive` and `docker history` to identify which layer is contributing the most size."

#### In-depth
The most impactful optimization is **multi-stage builds**. A Go binary compiled from source might have a 1GB builder image (Go SDK + build tools + source code) but the final binary is 15MB. With multi-stage builds, the production image contains only the 15MB binary + distroless base — a 98% size reduction. This pattern applies to any compiled language: Java, Rust, C++, TypeScript (transpile to JS).

---
