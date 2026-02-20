# 🛠️ **Docker BuildKit & Build Process (161–170)**

---

### 161. What is Docker BuildKit and why is it better?
"BuildKit is the **next-generation Docker build engine**, replacing the classic builder (`docker build`).

Key advantages over the classic builder:
- **Parallel execution**: independent stages run in parallel (classic builder runs sequentially)
- **Better caching**: finer cache invalidation, cache mounts for package managers
- **Secret mounts**: secrets available during RUN without persisting in layers
- **SSH forwarding**: clone private repos in builds without embedding keys
- **Garbage collection**: automatic build cache management with configurable size limits
- **Improved performance**: 2-10x faster builds in practice

Enabled by default since Docker 23.0. For older: `DOCKER_BUILDKIT=1 docker build .`"

#### In-depth
BuildKit's **parallel stage execution** is the biggest performance win for multi-stage builds. In a classic builder, all stages run sequentially. In BuildKit, stages with no dependencies on each other run in parallel. A Dockerfile with separate test and production stages can run both simultaneously — halving the total build time. The dependency graph is inferred from `COPY --from=stage` and `FROM stage AS` references.

---

### 162. How do you enable BuildKit in Docker CLI?
"Multiple methods:

**Environment variable (temporary)**:
```bash
DOCKER_BUILDKIT=1 docker build .
```

**Daemon config (permanent)** — in `/etc/docker/daemon.json`:
```json
{
  "features": {
    "buildkit": true
  }
}
```
Restart daemon: `systemctl restart docker`.

**Docker Desktop**: BuildKit is on by default since Docker Desktop 4.x.

**`docker buildx`**: always uses BuildKit — `docker buildx build .` uses BuildKit regardless of the DOCKER_BUILDKIT variable."

#### In-depth
`docker buildx` is the modern CLI plugin for BuildKit builds. It supports additional features that `docker build` doesn't even with BuildKit enabled: multi-platform builds, multiple builder backends (local, docker-container, kubernetes), and the full cache API. For any serious CI pipeline in 2024, use `docker buildx build` with the `--platform`, `--cache-from`, `--cache-to`, and `--push` flags — it's the complete build solution.

---

### 163. How does BuildKit handle secret mounts?
"BuildKit secrets are passed to RUN commands during build **without persisting in image layers**.

**Usage**:
```dockerfile
RUN --mount=type=secret,id=npmrc cat /run/secrets/npmrc
```

**Pass the secret** at build time:
```bash
docker buildx build --secret id=npmrc,src=$HOME/.npmrc .
```

The secret file appears at `/run/secrets/npmrc` inside the `RUN` command only. It's backed by `tmpfs` — never written to disk. It doesn't appear in `docker history`, `docker inspect`, or any image layer."

#### In-depth
Before BuildKit secrets, teams used multi-stage builds to remove secrets: `COPY --chown=root .npmrc /root/.npmrc && RUN npm install && rm -f /root/.npmrc` — but the `.npmrc` still exists in an intermediate layer visible via `docker history`. BuildKit secret mounts solve this correctly at the kernel level. Common use cases: private npm/pip/Maven package registry auth, SSH keys for private git dependencies, API keys needed during build.

---

### 164. How do you run parallel builds with BuildKit?
"BuildKit automatically parallelizes **independent stages** in a multi-stage Dockerfile.

```dockerfile
FROM node:20 AS test
COPY . .
RUN npm test

FROM node:20 AS build
COPY . .
RUN npm run build

FROM nginx:alpine AS final
COPY --from=build /app/dist /usr/share/nginx/html
```

`test` and `build` stages run in parallel (they don't depend on each other). `final` stage waits for `build`.

To build only specific targets: `docker buildx build --target test .` — only the test stage runs."

#### In-depth
For maximum parallelism, structure Dockerfiles as an explicit DAG (Directed Acyclic Graph). Stages that copy from the same base but have no mutual `COPY --from` dependencies run fully in parallel. Use `--build-arg BUILDKIT_INLINE_CACHE=1` to embed cache metadata in the pushed image — enabling cache reuse on `--cache-from`. BuildKit exposes a `--progress=plain` flag showing parallel stage execution clearly in the output.

---

### 165. What is the `RUN --mount=type=cache` used for?
"Cache mounts persist a directory **between builds** as a BuildKit cache volume — separate from image layers.

```dockerfile
RUN --mount=type=cache,target=/root/.npm \
    npm install

RUN --mount=type=cache,target=/root/.cache/pip \
    pip install -r requirements.txt

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update && apt-get install -y build-essential
```

The npm cache, pip cache, and apt cache directories persist between builds. `npm install` downloads packages on the first build, then reuses the cache on subsequent builds — even if package.json changed partially."

#### In-depth
Cache mounts are the single biggest build-time improvement for dependency-heavy builds. Traditional approach: change one line in `requirements.txt` → entire `pip install` re-downloads all packages from PyPI → 5 minutes. With cache mounts: change one line → pip downloads only the new/changed package from its persistent cache → 10 seconds. Cache mount data is stored in the BuildKit cache, separate from images. `docker buildx prune` manages it.

---

### 166. How do you debug Docker builds effectively?
"Built debugging techniques:

**1. Use `--progress=plain`**: shows all RUN output inline:
`docker buildx build --progress=plain .`

**2. Target a failing stage**:
`docker buildx build --target problematic-stage . && docker run --rm my-intermediate-stage`

**3. Add debug prints**:
`RUN ls -la /app && cat /app/config.json || echo "MISSING"`

**4. Interactive build debugging** (BuildKit):
`docker buildx debug build .` — opens an interactive shell where you can step through build instructions

**5. `--no-cache`**: bypass cache to ensure the failure is reproducible:
`docker buildx build --no-cache .`"

#### In-depth
`docker buildx debug` (Docker Desktop 4.24+) is a major debugging improvement. It provides a `--invoke` flag: `docker buildx debug --invoke /bin/sh build .` which, on a failed RUN instruction, drops you into an interactive shell in the state just before the failure occurred. You can run the failed command manually, inspect the filesystem, and diagnose the issue without modifying the Dockerfile or rebuilding.

---

### 167. What are `ONBUILD` instructions used for?
"`ONBUILD` registers a trigger instruction that executes **when a child image uses this image as a base**.

```dockerfile
# Base image Dockerfile
FROM node:20
ONBUILD COPY package*.json ./
ONBUILD RUN npm install
ONBUILD COPY . .
```

A child image that uses `FROM my-base` automatically runs those ONBUILD instructions during ITS build. The child's Dockerfile only needs: `FROM my-base && CMD ['node', 'server.js']`.

Useful for creating standardized base images for a team — the base image enforces certain build steps."

#### In-depth
`ONBUILD` is powerful but can be surprising — child image users may not expect hidden build steps. Docker's official language images previously used `ONBUILD` variants (e.g., `node:onbuild`) but they're deprecated now because the behavior was confusing and inflexible. Today, multi-stage builds are the preferred pattern for shared build logic. `ONBUILD` still has a niche: framework base images for internal teams with strict build conventions.

---

### 168. How do you pass build arguments securely?
"Build arguments (`ARG`) are for **non-secret build-time config** — they appear in `docker history`.

For secure values, use **BuildKit secret mounts** instead:
```dockerfile
# WRONG - secret exposed in history
ARG NPM_TOKEN
RUN npm config set //registry.npmjs.org/:_authToken $NPM_TOKEN

# CORRECT - secret never in layers
RUN --mount=type=secret,id=npm_token \
    NPM_TOKEN=$(cat /run/secrets/npm_token) npm install
```

For non-sensitive build args (build version, git SHA, target environment):
```bash
docker buildx build --build-arg GIT_SHA=$(git rev-parse HEAD) .
```"

#### In-depth
Build args are visible via `docker history --no-trunc myimage` — every ARG value that was set during build is stored in the image metadata. This is an accidental data leak vector: teams that pass access tokens, passwords, or API keys via `--build-arg` are unknowingly embedding those secrets in the image for anyone with `docker pull` access to see. Run `docker history myimage` in security audits to catch this pattern.

---

### 169. What is a base image vs parent image?
"**Parent image**: the image referenced in the `FROM` instruction. It provides the starting filesystem.

```dockerfile
FROM nginx:1.25  # nginx:1.25 is the parent image
```

**Base image**: the foundational image at the bottom of the chain. Usually: `scratch` (empty), `alpine`, `ubuntu`, `debian`. Every other image is derived from a base image via inheritance chains.

`scratch` is Docker's truly empty base image — no OS, no shell, nothing. Used for compiling minimalist images: `FROM scratch` + `COPY myapp /` + `CMD ["/myapp"]`. Most Go app production images should use scratch or distroless."

#### In-depth
The `scratch` image generates images with literally 0 layers beyond your binary. `FROM scratch` signals to Docker that there's no parent to fetch — the COPY instruction adds the first layer. The resulting image has no attack surface: no shell, no utilities, no OS CVEs. The required extras for most apps: CA certificates (`COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/`) and timezone data (`tzdata` package or copy from another image).

---

### 170. How do you prevent sensitive data leakage in builds?
"Multi-layered prevention:

1. **Use BuildKit `--secret`** instead of `--build-arg` for any credentials
2. **Use SSH forwarding** for git cloning: `docker buildx build --ssh default .`
3. **Verify with `docker history --no-trunc myimage`** — check no sensitive strings appear
4. **Multi-stage builds**: sensitive data in builder stage, only artifact in final stage
5. **`.dockerignore`**: exclude `.env`, `credentials`, `*.pem`, `*.key` from build context
6. **Automated detection**: scan images with **Trivy** (`--scanners secret`) or **Trufflehog** to detect leaked secrets in image layers"

#### In-depth
`docker history --no-trunc` reveals all Dockerfile instructions and their arguments embedded in image metadata. `docker save myimage | tar xf - | cat */layer.tar | tar tv` reveals all files in all layers. Leaked secrets are often: private keys copied in and then deleted (still visible in history layers), passwords in `ENV` instructions, or API tokens passed as `ARG`. Automated secret scanning as a CI gate catches these before pushing to a registry accessible by others.

---
