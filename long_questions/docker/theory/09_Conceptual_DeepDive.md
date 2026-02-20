# 🧠 **Conceptual & Deep-Dive Docker Questions (101–110)**

---

### 101. What is the OCI (Open Container Initiative)?
"OCI is the **Open Container Initiative** — an open governance project under the Linux Foundation that defines standard specifications for container runtimes and image formats.

It emerged in 2015 when Docker donated its container format and runtime to avoid fragmentation in the ecosystem. Two main specs: **OCI Runtime Specification** (how to run a container) and **OCI Image Specification** (how to format a container image).

Any OCI-compliant runtime (runc, crun, gVisor, Kata) can run any OCI-compliant image — Docker, Podman, Buildah, containerd all interoperate. This prevented a proprietary lock-in."

#### In-depth
Before OCI, Docker's image format and runtime were proprietary. Kubernetes first used Docker directly, then abstracted to CRI (Container Runtime Interface), which led to pluggable runtimes like containerd and CRI-O. The OCI spec is why you can build an image with `docker build`, push to any OCI-compatible registry (GitHub Container Registry, ECR), and run with `containerd` in Kubernetes — the entire pipeline interoperates.

---

### 102. How is Docker related to containerd and runc?
"Docker is built on a stack of components:

**Docker CLI** → **Docker daemon (dockerd)** → **containerd** → **runc**

**runc**: the low-level OCI runtime — actually creates the container (namespaces, cgroups, mounts). It's a CLI tool, not a daemon.
**containerd**: manages container lifecycle (start, stop, pull images, store images). A daemon that sits between dockerd and runc.
**dockerd**: the Docker daemon — provides the full Docker API, handles builds, volumes, networking, and delegates container operations to containerd.

Docker donated containerd to CNCF in 2017. Kubernetes now uses containerd directly via CRI, bypassing dockerd entirely."

#### In-depth
The architectural split matters for Kubernetes: when Kubernetes deprecated the Docker shim in 1.24, it didn't mean containers stopped working — it means Kubernetes talks to containerd directly via CRI instead of through dockerd. Docker images still work perfectly because they're OCI-compliant. Only the intermediate dockerd layer was removed.

---

### 103. What is the architecture of Docker Engine?
"Docker Engine has three main components:

1. **Docker CLI**: the command-line client that sends API requests to dockerd
2. **Docker daemon (dockerd)**: a long-running server that manages Docker objects (images, containers, networks, volumes). Exposes REST API (default: Unix socket `/var/run/docker.sock`)
3. **containerd**: the container runtime that manages container lifecycle. dockerd delegates container operations to it

The daemon listens on a Unix socket (or optionally TCP) and handles requests from the CLI. remotely-accessible Docker APIs use TLS certificates for security."

#### In-depth
The Unix socket `/var/run/docker.sock` is the Docker API endpoint on Linux. By default it's owned by root or the `docker` group. Users in the `docker` group have effective root access — because they can mount the host filesystem, start privileged containers, etc. This is why Docker rootless mode (user-namespace based Docker daemon) is increasingly recommended for development environments.

---

### 104. What is the difference between Docker CLI and Docker API?
"The Docker CLI is the **human-friendly interface** — `docker run`, `docker build`, `docker ps`. Internally, every CLI command translates to an **HTTP request against the Docker REST API**.

The Docker REST API is the **programmatic interface** — available over the Unix socket or TCP. You can call it directly: `curl --unix-socket /var/run/docker.sock http://localhost/containers/json`.

SDK clients (Go, Python, Java) wrap the REST API for programmatic container management. Portainer, Docker Desktop, and all third-party Docker management tools use the API directly."

#### In-depth
The Docker API is versioned (e.g., `/v1.43/containers/json`). The CLI negotiates the API version with the daemon. Running `DOCKER_API_VERSION=1.40 docker ...` forces a specific version. For automation, I prefer the SDK clients over raw HTTP because they handle version negotiation, authentication, and TLS automatically — the `docker-py` Python SDK and `moby/moby/client` Go package are the most common.

---

### 105. What is the role of dockerd?
"**dockerd** (Docker daemon) is the server-side process that manages all Docker objects.

It: accepts API requests (from CLI or SDK), builds images (`docker build`), manages image storage (layers, cache), creates/starts/stops containers (delegates to containerd), manages Docker networks (creates bridges, configures iptables), manages volumes (creates, mounts, removes), and serves the Docker API over the socket.

It's the central coordinator — without dockerd running, all Docker commands fail."

#### In-depth
dockerd's configuration lives in `/etc/docker/daemon.json` (Linux). Important settings: `storage-driver` (usually overlay2), `log-driver`, `log-opts` (log rotation), `insecure-registries` (private registries without TLS), `registry-mirrors` (Docker Hub mirror for rate limiting), `data-root` (where images/containers are stored — move to a larger disk). Restart dockerd after config changes: `systemctl restart docker`.

---

### 106. How does Docker handle layered filesystems?
"Docker images are built from **stacked, read-only layers**. Each layer contains only the filesystem diff from the previous layer — added, modified, and deleted files.

When a container starts, Docker adds a thin **writable layer** on top. Any writes the container makes go to this writable layer (copy-on-write). The underlying image layers remain unchanged and shared across all containers from the same image.

This is implemented by a **union filesystem** (most commonly OverlayFS on Linux)."

#### In-depth
OverlayFS merges the layer stack into a single filesystem view using two directories: `lowerdir` (all read-only layers, merged) and `upperdir` (the writable container layer). The `merged` mount point is what the container sees — a unified view combining all layers. When a file is modified, it's copied from `lowerdir` to `upperdir` on first write (copy-on-write) — then modifications happen in `upperdir`.

---

### 107. Explain union file systems in Docker (AUFS, OverlayFS, etc.).
"Union filesystems combine multiple directories into a single view — like transparencies stacked on an overhead projector.

**AUFS** (Another Union FS): the original Docker storage driver. Complex patches, not in the mainline kernel, still supported on Ubuntu for legacy.

**OverlayFS**: modern, mainline Linux kernel (3.18+), simpler design, better performance. Docker's default storage driver — `overlay2`. Uses `lowerdir` (read-only layers), `upperdir` (writable), `workdir` (temp for atomic operations), and `merged` (the unified view).

**Device Mapper** and **Btrfs**: alternatives for specific enterprise use cases."

#### In-depth
OverlayFS has a limitation: it can only merge two directories — a lower and upper. Docker works around this for multi-layer images by chaining overlay mounts (each layer's merged view becomes the lowerdir for the next). Starting from kernel 5.11, OverlayFS supports multiple lower directories natively (`overlay2` with `metacopy` and `redirect_dir` mount options), improving multi-layer performance.

---

### 108. What are Docker manifests?
"A Docker manifest is a **JSON document that describes an image**: its layers, the platform it supports (architecture, OS), and configuration.

A **manifest list** (or **multi-platform manifest**) points to multiple manifests for different platforms: `linux/amd64`, `linux/arm64`, `windows/amd64`. When you `docker pull nginx`, Docker pulls the manifest list, identifies your platform, then fetches the correct platform-specific manifest and layers.

I create multi-platform manifests with `docker buildx build --platform linux/amd64,linux/arm64 --push`."

#### In-depth
Manifest lists enable transparent multi-architecture support. The image tag `nginx:latest` on Docker Hub is a manifest list — the same tag works on Intel Macs, Apple Silicon Macs, Raspberry Pis, and AWS Graviton servers. The registry returns the appropriate layers based on the pulling client's architecture. This is essential for organizations moving workloads to ARM for cost savings (AWS Graviton is ~20% cheaper than equivalent Intel).

---

### 109. What is the difference between Docker manifests and Docker images?
"A Docker **image** is the actual set of layers (filesystem data) plus a config JSON (environment variables, CMD, labels).

A Docker **manifest** is the metadata document that describes an image: list of layer digests, the image config digest, and platform info. Think of it as the index pointing to the actual content.

A **manifest list** is one level above — it's a list of manifests, one per supported platform. When you reference an image by tag, you're actually referencing either a manifest or a manifest list."

#### In-depth
Understanding this matters for immutable deployments. A tag (`myapp:v1.0`) is mutable — it can be rewritten to point to a different manifest. A digest (`myapp@sha256:abc123`) is immutable — it always refers to the same content. For production Kubernetes deployments, pinning by digest ensures you get exactly the image you tested, even if someone overwrites the tag.

---

### 110. How does the Docker image ID work?
"The Docker image ID is the **SHA256 digest of the image's configuration JSON**.

The image config includes: environment variables, CMD/ENTRYPOINT, labels, and crucially the **ordered list of layer digests**. If any layer changes, the config changes, and the image ID changes.

This makes image IDs **content-addressable** — two identical builds produce the same ID. If you see the same pull producing a different ID, the image content changed. You can use image IDs to verify you're running exactly the image you think you are."

#### In-depth
The distinction between image ID, config digest, and manifest digest is subtle but important. `docker images --digests` shows the manifest digest (what's stored in the registry). `docker inspect image --format='{{.Id}}'` shows the image config digest. The manifest digest includes the image config digest, so both change if anything changes. For pinning in Kubernetes, use the manifest digest (`image: nginx@sha256:...`) because that's what the registry resolves.

---
