# 🧱 **Docker Core Internals & Advanced Dockerfile (201–220)**

---

### 201. What is the difference between containers and sandboxing?
"**Containers** use Linux kernel features (namespaces, cgroups) to isolate processes. They share the host kernel — the boundary is enforced by the kernel itself, which can be bypassed via kernel vulnerabilities.

**Sandboxing** is a broader term for any isolation mechanism. It includes containers, but also: browser sandboxes (Chrome's site isolation, WebAssembly sandboxes), VMs (hardware-level isolation), and language-level sandboxes (JVM security manager, V8 isolates).

The key difference: containers share the kernel (kernel exploits = container escape). True sandboxes (gVisor, Kata) add an additional isolation layer above the host kernel."

#### In-depth
**gVisor** implements sandboxing for containers by intercepting all syscalls with a user-space kernel (Sentry). The container's process talks to Sentry, which validates and translates syscalls — the host kernel sees only a small set of upcalls. A container exploit can compromise Sentry but not the host kernel. Kata Containers take a different approach: each container runs in a lightweight VM, so the host kernel is fully isolated. Both have overhead vs. native containers but provide significantly stronger isolation guarantees.

---

### 202. What is the Docker Engine API and how do you use it?
"The Docker Engine API is a RESTful HTTP API exposed by dockerd. It's the same API the Docker CLI uses internally.

By default on Linux: `unix:///var/run/docker.sock`. Direct access:
```bash
curl --unix-socket /var/run/docker.sock http://localhost/containers/json
curl --unix-socket /var/run/docker.sock http://localhost/images/json
```

SDK examples (Python):
```python
import docker
client = docker.from_env()
containers = client.containers.list()
client.containers.run('ubuntu', 'echo hello', remove=True)
```

API documentation: https://docs.docker.com/engine/api/"

#### In-depth
The Docker API is versioned: `/v1.43/containers/json`. Clients negotiate the highest mutually supported version. For production automation, use the official SDK clients rather than raw HTTP — they handle versioning, TLS, error parsing, and stream handling correctly. Common automation uses: CI/CD orchestration, container management dashboards (Portainer), monitoring agents (cAdvisor), and custom auto-scaling scripts.

---

### 203. How does Docker manage container lifecycle states internally?
"Internally, Docker tracks containers through state transitions via the containerd state machine.

States stored in containerd's metadata store (boltDB): **created**, **running**, **paused**, **stopped**, **dead**. Each transition triggers events that Docker daemon subscribes to and propagates to the Docker event stream.

The **restart policy controller** watches for `stopped` events and re-queues start operations based on the policy (always, on-failure, unless-stopped). OOMKill events from cgroups similarly trigger restart logic.

You can observe the event stream: `docker events --filter type=container --format '{{.Status}} {{.Actor.Attributes.name}}'`"

#### In-depth
The containerd daemon stores all container metadata in `/var/lib/containerd/` using BoltDB (an embedded key-value store). Docker daemon caches some state but the ground truth is containerd's metadata. After a dockerd crash and restart, it recovers state by re-reading containerd's metadata — this is why containers persist through daemon restarts (unlike when using raw `containerd` without Docker's management layer, which doesn't persist state across daemon restarts by default).

---

### 204. What happens under the hood when you run `docker run`?
"The sequence:

1. Docker CLI parses the command and sends a POST to `/containers/create`
2. dockerd allocates a container ID, creates the container spec, and tells containerd to create a container
3. containerd creates the OCI runtime bundle (config.json + rootfs)
4. runc is invoked to create the container: sets up namespaces, cgroups, mounts the rootfs
5. dockerd networking: creates a veth pair, attaches to bridge, configures iptables
6. POST to `/containers/{id}/start`: runc executes the entrypoint in the new container
7. PID 1 (your process) runs inside the container

All this happens in ~100-300ms for a typical container start."

#### In-depth
The OCI runtime bundle is a directory with `config.json` (capabilities, namespaces, cgroups, mounts, process spec) and a `rootfs/` directory (the merged OverlayFS view). runc reads `config.json` and calls Linux kernel APIs: `clone()` with namespace flags, `cgroupv2` resource controllers, `pivot_root()` to change the filesystem root. The actual container creation is purely a series of system calls — no VMs, no hypervisors.

---

### 205. Explain the Docker image digest.
"An image digest is the **SHA256 hash of the image manifest**.

The manifest JSON describes all the image layers and config. Any change to any layer changes its digest → changes the manifest → changes the manifest digest. The digest is therefore a **cryptographic fingerprint** of the exact image content.

Reference by digest: `nginx@sha256:a3a96a...`. This reference is immutable — unlike a tag which can be repointed, `@sha256:...` always refers to the same content.

`docker images --digests` shows digests for locally cached images."

#### In-depth
The manifest digest is computed at the registry after all layers are uploaded. Two identical builds (same Dockerfile, same base, same context) may produce the same image ID (image config digest) but different manifest digests if built at different times — because manifest timestamps differ. For truly reproducible images, use `SOURCE_DATE_EPOCH` and `--reproducible` flags with BuildKit (alpha feature) to control timestamps and filesystem metadata, enabling identical digests for identical builds.

---

### 206. What is the Docker storage driver?
"The storage driver manages how image layers are stored on disk and presented to containers as a unified filesystem.

Available drivers:
- **overlay2**: Default on modern Linux. Uses OverlayFS — fastest and most efficient
- **devicemapper**: Legacy, used on RHEL/CentOS before overlay2 support
- **btrfs**: Uses Btrfs filesystem features for layering
- **zfs**: Uses ZFS for layering (ZFS-on-Linux)
- **aufs**: Legacy, requires patched kernel, used on old Ubuntu systems
- **vfs**: No union mount, full copy per layer — slowest, used for testing

Check current driver: `docker info | grep 'Storage Driver'`"

#### In-depth
**overlay2** is universally recommended. It requires a Linux kernel ≥4.0 and a compatible backing filesystem (ext4, xfs). The `overlay2` driver stores layers as directories under `/var/lib/docker/overlay2/`. Each layer has `diff/` (the layer contents), `lower` (a file listing parent layer IDs), and `merged/` (the union view when a container uses this layer). The simplicity and kernel integration of OverlayFS makes it 2-5x faster than devicemapper for typical workloads.

---

### 207. Compare OverlayFS and AUFS in Docker.
"Both are union filesystems that stack read-only layers for Docker images.

**AUFS** (Another Union FS): was Docker's original driver. Requires kernel patches (not in mainline Linux for years). Multi-layer support built-in. Still used on legacy Ubuntu deployments.

**OverlayFS** (overlay2): in mainline Linux kernel since 3.18 (2014). Much simpler implementation. Originally supported only 2 directories (lower + upper); since kernel 5.11 supports multiple lower directories natively. Faster than AUFS for most operations, especially stat() calls.

**Winner**: overlay2 on any modern system. AUFS is end-of-life."

#### In-depth
The OverlayFS performance advantage over AUFS is most pronounced for **rename operations** and **metadata operations** (stat, inode lookups). AUFS's multi-lower-layer support had to be emulated in older OverlayFS via chaining (each layer's `merged` becomes the next layer's `lower`). Modern Linux with native multi-lower overlayfs eliminates this chain, improving performance for images with many layers.

---

### 208. What is a Docker union mount?
"A union mount combines multiple filesystem directories into a single unified view.

For Docker: the image consists of read-only layers (L1, L2, L3...). When a container starts, Docker creates a union mount that merges all layers into what appears as a single coherent filesystem.

If a file exists in multiple layers, the **topmost layer wins** (most recently written). When a container writes to a file from a lower layer, the file is copied up to the container's writable layer first (copy-on-write), then modified. The read-only layers remain unchanged and shared."

#### In-depth
The copy-on-write (CoW) mechanism has performance implications for large files that are frequently modified. The first write to a large file (e.g., a 2GB database file) requires copying the entire file to the upper layer before the write proceeds — this can take seconds. For databases, this argues strongly against storing data in the container's copy-on-write filesystem and in favor of volumes (which bypass the union mount entirely and write directly to a backing filesystem).

---

### 209. How does Docker manage layered files in copy-on-write systems?
"Copy-on-write (CoW) in docker: containers never modify image layers directly. Instead:

**Read**: read from whichever layer the file exists in (checked top to bottom)
**Write**: on first write, copy the full file from its image layer to the container's writable layer (upperdir). Subsequent writes modify the upperdir copy.
**Delete**: create a 'whiteout' file in the upperdir — a special file that marks the path as deleted in the union view

This means: reads are free (no copying), only writes trigger the copy. Read-heavy workloads benefit most from CoW; write-heavy workloads (databases) should use volumes instead."

#### In-depth
Whiteout files (`.wh.filename`) are how deletions work in union filesystems. When Docker's `RUN rm /etc/unnecessary-file` executes, it doesn't delete the file from the lower layer — it creates `/.wh.unnecessary-file` in the new layer. The union mount hides files masked by whiteouts. This is why `RUN apt-get install && RUN rm -rf /var/lib/apt/lists/*` in separate instructions doesn't actually reduce image size — the files exist in the install layer, just masked by the delete layer. Combine them in one `RUN` to avoid this.

---

### 210. What is the role of containerd in Docker's architecture?
"containerd is the **container runtime daemon** that sits between dockerd and runc.

It handles: container lifecycle (create, start, stop, delete), image management (pull, store, manage layers via snapshotter), container networking interface, and resource management (cgroups).

dockerd delegates actual container operations to containerd via its gRPC API. containerd then uses runc (or any OCI-compliant runtime) to create containers.

containerd is also used directly by Kubernetes (via CRI plugin) without dockerd — this is the configuration used in most production Kubernetes clusters today."

#### In-depth
containerd uses **snapshotters** (abstraction over storage) instead of direct storage drivers. The default snapshotter is `overlayfs`. Containerd's snapshotters are pluggable — `stargz-snapshotter` enables lazy pulling of images (pull only the layers needed for the running process on demand), dramatically improving cold start times. This is a major research area for large images in FaaS/serverless containers.

---

### 211. How do you use `ARG` vs `ENV` in Dockerfile?
"`ARG` is a **build-time variable** — only available during `docker build`. Not present in the running container.

`ENV` is an **image variable** — persists in the image and is available in all subsequent build steps AND in the running container.

```dockerfile
ARG VERSION=1.0.0          # Build-time, used in Dockerfile
ENV APP_VERSION=$VERSION   # Captured into the image, available at runtime
ARG SECRET                 # Build-time only - don't use for real secrets
```

Use `ARG` for: version numbers, platform targets, build flags. Use `ENV` for: app configuration that the running process needs."

#### In-depth
A common pattern: `ARG TAG; ENV APP_TAG=$TAG` — capture a build argument into an environment variable so the running app knows what version it is. But remember: `ARG` values appear in `docker history`! Never use `ARG` for passwords, tokens, or private keys. ARGs set before `FROM` can be used in `FROM` for dynamic base images: `ARG BASE=alpine; FROM $BASE`. This enables `docker build --build-arg BASE=debian .` for different base image variants.

---

### 212. How can you dynamically set labels in a Dockerfile?
"Use `ARG` to pass values at build time, then assign them to `LABEL`:

```dockerfile
ARG GIT_SHA=unknown
ARG BUILD_DATE=unknown
ARG VERSION=0.0.0

LABEL org.opencontainers.image.version=$VERSION \
      org.opencontainers.image.created=$BUILD_DATE \
      org.opencontainers.image.revision=$GIT_SHA \
      org.opencontainers.image.source='https://github.com/myorg/myapp'
```

Build with:
```bash
docker buildx build \
  --build-arg GIT_SHA=$(git rev-parse HEAD) \
  --build-arg BUILD_DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  --build-arg VERSION=1.2.3 .
```"

#### In-depth
OCI image annotations (the `org.opencontainers.image.*` labels) are the standard way to embed metadata in images. They're queryable via `docker inspect`  and registry APIs, enabling: **traceability** (SHA → commit → author), **audit** (when was this image built), and **policy enforcement** (only images with certain labels pass deployment gates). Kubernetes admission webhooks or OPA policies can require specific labels before allowing image deployment.

---

### 213. How do you split Dockerfiles for better caching?
"The key insight: **split Dockerfiles at cache boundaries** — instructions that change frequently must come after instructions that change rarely.

**Optimal ordering**:
1. Base image (never changes)
2. System dependencies (changes when new packages needed - rare)
3. Application dependencies (`package.json`, `go.mod`, `requirements.txt`)
4. Application code (changes on every commit)

Two files approach: a `Dockerfile.deps` that only installs dependencies (rarely rebuilds), pushed to registry as a cache-from source. Then `Dockerfile.app` that adds app code. The deps image is shared across all team members."

#### In-depth
The 'split Dockerfile' approach with registry-cached deps images is common in monorepos with shared library dependencies. If `shared-lib` changes, only the services that depend on it need to rebuild their deps layer. Services that only changed application code reuse the deps layer from cache. This transforms the worst case (rebuild all 20 services on a lib change) into: (rebuild shared-lib deps layer, then rebuild affected service app layers).

---

### 214. What is a good strategy for base image versioning?
"**Pin to a specific version, never `latest`**:
- Instead of `FROM node:20` → use `FROM node:20.11.0-alpine3.19`
- Or digest pinning: `FROM node:20@sha256:abc123...`

Use **Renovate Bot** or **Dependabot** to automatically create PRs when new base image versions are available — you get automated update proposals without losing control.

Semantic strategy:
- `major.minor` (e.g., `node:20`) — get patch updates automatically
- `major.minor.patch-distro` — fully pinned, maximum stability
- digest — absolute immutability, requires bot to update"

#### In-depth
The trade-off: looser pins get security patches automatically but risk unexpected behavior changes; tighter pins require manual updates but are predictable. My recommendation: pin to `major.minor-distro` (e.g., `node:20-alpine`) for base images — you get automatic patch releases (security fixes) while being protected from major or minor version compatibility breaks. Use Renovate Bot to alert and auto-PR when the minor version changes.

---

### 215. Why should you pin image tags instead of using `latest`?
"`latest` is mutable — it can point to different image content every time you pull. This causes several problems:

1. **Reproducibility broken**: `FROM ubuntu:latest` in a Dockerfile built today vs. next month may use different Ubuntu versions
2. **Unexpected behavior**: a dependency update in `latest` could break your build silently
3. **No rollback path**: you can't easily return to 'the version that worked' without a specific tag
4. **CI inconsistency**: different CI runs with the same Dockerfile may produce different images

Always use specific tags: `FROM ubuntu:22.04` or `FROM node:20.11.0-alpine3.19`."

#### In-depth
The `latest` tag trap is especially painful for base images. `node:latest` has jumped from Node 18 to 20 to 22 across different years — each jump would break any app requiring a specific Node version. Using `latest` for your own images in production Kubernetes deployments is even more dangerous: `imagePullPolicy: Always` with `latest` means every pod restart could pull a different image — you might deploy v1.5 to some pods and v1.6 to others without any explicit action.

---

### 216. How do you ensure idempotent Dockerfile builds?
"Idempotent: same inputs → same outputs, every time.

Strategies:
1. **Pin all version numbers**: base image, package versions, tool versions
2. **Pin package lock files**: `COPY package-lock.json .` then `RUN npm ci` (not `npm install`)
3. **Use `COPY` not `ADD` with URLs** — URLs are non-deterministic
4. **Set `SOURCE_DATE_EPOCH`** for timestamp control in builds
5. **No internet calls in `RUN`** without checksums for downloaded artifacts
6. **Use `--no-cache`** periodically to verify: `docker build --no-cache .` should produce the same image"

#### In-depth
Truly reproducible builds (same binary/layer content across rebuilds) require controlling all sources of non-determinism: timestamps (file mtimes in layers), sort order of filesystem entries, randomness used during compilation. BuildKit's `--reproducible` flag (experimental) attempts to normalize timestamps. Language-level tools: Go's `-trimpath` flag removes absolute paths from binaries, making identical Go binaries across different build machines.

---

### 217. How can you minimize Docker build context?
"The build context is everything sent to the Docker daemon when you run `docker build .`. Minimize it by:

1. **`.dockerignore`**: exclude everything not needed:
```
.git
node_modules
*.log
.env
dist
coverage
*.md
.DS_Store
```

2. **Specify an explicit context**: `docker build -f Dockerfile ./src` — only sends `./src`

3. **Use BuildKit inline Dockerfiles**: `docker build - <<EOF` to pass Dockerfile from stdin with no context

4. Monitor context size: `docker build . 2>&1 | head -5` — shows 'Sending build context to Docker daemon X.XMB'"

#### In-depth
The build context check: before any build, Docker shows 'Sending build context to Docker daemon X.XMB'. If this is more than a few MB, investigate. The most common offenders: `.git` directory (can be 50-200MB for old repos), `node_modules` (often 200-500MB), compiled artifacts (`dist`, `build`, `target`), test fixtures with large files. A well-configured `.dockerignore` reduces the context from hundreds of MB to a handful of KB, dramatically speeding up builds.

---

### 218. What's the use of `USER` in Dockerfile?
"`USER` sets the user and optionally group for all subsequent `RUN`, `CMD`, and `ENTRYPOINT` instructions.

```dockerfile
RUN addgroup -S app && adduser -S app -G app
USER app:app

# Now RUN commands execute as 'app' (non-root)
RUN ./build-as-non-root.sh
CMD ["./app"]  # App starts as 'app' user
```

Without `USER`: all instructions run as root. The container starts as root. Security vulnerability: if the app is compromised, the attacker has root inside the container (and potentially on the host)."

#### In-depth
Placement of `USER` matters for layer caching and for file ownership. Create the user after installing system packages (which require root) but before copying application files — so the `COPY` layer and app files are owned by the app user from the start. If you `COPY . /app` as root then `USER app`, the app user can't write to `/app` (owned by root). Alternative: `COPY --chown=app:app . /app` to set ownership during the COPY step.

---

### 219. How can you multi-target builds in Docker?
"Multi-target builds use named stages in a multi-stage Dockerfile. Target specific stages with `--target`:

```dockerfile
FROM base AS deps
RUN npm install --production

FROM deps AS test-deps
RUN npm install  # includes dev deps

FROM node:20 AS test
COPY --from=test-deps /app .
RUN npm test

FROM nginx:alpine AS production
COPY --from=deps /app/dist /usr/share/nginx/html
```

Build commands:
- `docker build --target test .` — run tests stage
- `docker build --target production .` — production image
- `docker build .` — builds the final stage (production)"

#### In-depth
Multi-target builds enable: **matrix testing** (build same code against multiple Node/Python versions simultaneously), **separate CI steps** (test stage in CI, production stage for deployment), and **artifact extraction** (build test reports in test stage, copy them out with `docker cp`). Each `--target` build is independently cacheable — test stage cache is separate from production stage cache, so changing test code doesn't invalidate the production image cache.

---

### 220. How do you debug failed `RUN` commands during build?
"Built-in approaches:

1. **`--progress=plain`**: shows all stdout/stderr from RUN: `docker buildx build --progress=plain .`

2. **Target the failing stage**: build up to the failing stage, then run interactively:
```bash
docker build --target failing-stage -t debug-image .
docker run --rm -it debug-image sh
# manually run the failing command inside
```

3. **`docker buildx debug` with `--invoke`**:
```bash
docker buildx debug --invoke /bin/sh build .
# Drops into shell at the point of failure
```

4. **Split the RUN**: break into smaller steps to isolate which command fails. Then recombine."

#### In-depth
The `docker buildx debug --invoke` approach is the most powerful. When a RUN instruction fails, `--invoke /bin/sh` gives you an interactive shell in the filesystem state just before the failure. You can exactly reproduce the failed command, inspect the environment, check file permissions, and iterate until you fix it — all without rebuilding. This saves enormous time vs. add-debug-print → rebuild → check output → repeat cycle.

---
