## 🚀 Production Practices & Pitfalls (Questions 351-360)

### Question 351: What is a zombie container and how do you handle it?

**Answer:**
A container that is technically "Running" but the main process is defunct or unresponsive to signals.
**Fix:**
- Send `SIGKILL` (`docker kill`).
- Investigate why it trapped signals.
- Check for Host hardware issues (Disk hang).

---

### Question 352: How do you gracefully shutdown containers?

**Answer:**
`docker stop` sends `SIGTERM`. The app should catch this, finish requests, close DB connections, and exit.
If it doesn't exit in 10s (default), Docker sends `SIGKILL`.
**Config:** change timeout `docker stop --time=30`.

---

### Question 353: What are exponential backoff and retries in container healthchecks?

**Answer:**
If a container fails, restarting it immediately might overload the dependency (e.g., DB) that caused the failure.
**Exponential Backoff:** Wait 1s, then 2s, 4s, 8s...
Docker Swarm/K8s implement this automatically.

---

### Question 354: How do you debug random container crashes in production?

**Answer:**
1.  **OOM:** Check `dmesg` on host or `docker inspect`.
2.  **Logs:** Send logs to external aggregation (Splunk) to see last message before death.
3.  **Core Dumps:** Configure system to save core dumps for C++ processes.

---

### Question 355: How do you limit disk usage per container?

**Answer:**
**Storage Opts:**
`docker run --storage-opt size=10G ...`
*(Requires underlying filesystem support like xfs or specific overlay2 configuration).*

---

### Question 356: How do you limit concurrent connections in a containerized service?

**Answer:**
Do this at the **Application Layer** (e.g., Nginx `worker_connections`) or **Orchestrator Layer** (Ingress Controller rate limiting). Docker cgroups don't limit TCP connection counts directly.

---

### Question 357: What is container orphaning and how do you prevent it?

**Answer:**
Orphans are containers managed by tools (like Compose) that lost trak of them.
**Prevention:**
Use consistent Project Names (`-p myproject`) in Compose.

---

### Question 358: What’s the best way to run cron jobs in Docker?

**Answer:**
**Don't** run a long-lived container just for cron.
**Use Orchestrator capabilities:**
- **K8s:** `CronJob` resource.
- **Swarm:** Service is harder. Use a sidecar cron or host cron triggering `docker run`.

---

### Question 359: How do you avoid data corruption in stateful containers?

**Answer:**
1.  **Graceful Shutdown:** Ensure DB flushes to disk on SIGTERM.
2.  **Volume:** Use reliable storage backend.
3.  **Locks:** Ensure only one container writes to the volume at a time.

---

### Question 360: What is a container lifecycle policy?

**Answer:**
A policy determining how long images/containers live.
- "Delete containers older than 7 days".
- "Retain last 5 images".
Implemented via scripts or tools like `docker-gc`.

---

## 🛠️ Docker DevOps Tooling (Questions 361-370)

### Question 361: What is Docker layer caching and how do you visualize it?

**Answer:**
Visualizing: `docker build --progress=plain`.
You see `CACHED [2/5] RUN npm install`.
If a step is not cached, all subsequent steps show as fresh builds.

---

### Question 362: How does `docker diff` help during debugging?

**Answer:**
Shows files changed in the container's writable layer compared to the image.
- **A:** Added.
- **C:** Changed.
- **D:** Deleted.
Helps identify if valid config files were accidentally overwritten or logs are filling up root.

---

### Question 363: What is the use of `.docker/config.json`?

**Answer:**
Stores client-side configuration:
- Auth credentials (for registries).
- Alias/Proxy settings.
- CLI plugin settings.
Location: `~/.docker/config.json`.

---

### Question 364: What are the performance implications of mounting large volumes?

**Answer:**
- **Bind Mounts:** Near native speed.
- **Docker Desktop (Mac/Win):** **Very Slow** due to file sharing overhead across VM boundary.
  - *Fix:* Use Named Volumes or Mutagen.

---

### Question 365: How do you run containers inside containers (dind)?

**Answer:**
**Docker-in-Docker (DinD).**
Use the `docker:dind` image. It runs a new Child Docker Daemon inside the container.
- Requires `--privileged`.
- Alternative: Socket Binding (DooD) - reuse host daemon.

---

### Question 366: What’s the difference between `COPY` and `ADD`?

**Answer:**
*Refresher:*
- `COPY`: Local files only.
- `ADD`: Local files + URLs + Auto-extraction of tarballs.
**Best Practice:** Always use `COPY` unless you explicitly need extraction.

---

### Question 367: How do you use `docker cp` effectively?

**Answer:**
Snapshotting config or DB dumps.
`docker cp db:/var/lib/mysql/dump.sql ./backup.sql`
Works even on stopped containers.

---

### Question 368: How do you run an interactive session inside a production container?

**Answer:**
`docker exec -it <id> /bin/bash`.
*Warning:* Any changes made are ephemeral but can break the running app. Avoid in strict prod.

---

### Question 369: How do you tag a container build with a dynamic version?

**Answer:**
Shell expansion.
`docker build -t myapp:$(date +%Y%m%d) .`
Or use Git SHA:
`docker build -t myapp:$(git rev-parse --short HEAD) .`

---

### Question 370: How do you chain build tools with Docker volumes?

**Answer:**
Volume-based pipeline.
1.  `docker run -v src:/src maven mvn build` (Outputs jar to /src/target).
2.  `docker run -v src:/src java run ...`
The volume acts as the shared workspace.

---

## 🗂️ Registries and Image Distribution (Deep) (Questions 371-380)

### Question 371: What is the schema of a Docker image manifest?

**Answer:**
JSON structure (Schema V2).
Contains:
- **Config:** Digest of config JSON.
- **Layers:** Array of layer digests (blobs).

---

### Question 372: How does layer deduplication work across registries?

**Answer:**
Registry stores blobs by hash.
If `App A` and `App B` both require `Ubuntu Base Layer (Hash X)`, the registry stores Blob X once.
Both manifests point to Blob X.

---

### Question 373: How do you replicate images across multiple regions?

**Answer:**
- **AWS ECR:** Cross-region replication rules.
- **Harbor:** Replication policies.
- **Manual:** `docker pull`, `docker tag`, `docker push`.

---

### Question 374: What is OCI-compliant image layout and why does it matter?

**Answer:**
Standard directory structure (`blobs/`, `index.json`, `oci-layout`).
Allows tools (Skopeo, Buildah) to work with images on disk as if they were in a registry, without needing a Docker Daemon.

---

### Question 375: How do you sign Docker images and verify integrity?

**Answer:**
Use **Docker Content Trust (Notary)**.
or **Cosign (Sigstore)**.
`cosign sign --key cosign.key my-image:tag`.

---

### Question 376: What are cosign and Notary v2?

**Answer:**
Modern alternatives to the original Docker Notary.
- **Notary v2/Notation:** MS/AWS backed standard.
- **Cosign:** Part of Sigstore. Keyless signing using OIDC (OpenID Connect).

---

### Question 377: How do you mirror DockerHub to a local registry?

**Answer:**
Run the standard Registry container in **proxy mode**.
Env Var: `REGISTRY_PROXY_REMOTEURL=https://registry-1.docker.io`.
Pull from local; it caches from Hub.

---

### Question 378: How do you serve a Docker registry over S3?

**Answer:**
Run the `registry:2` container.
Configure storage driver:
```yaml
storage:
  s3:
    bucket: my-bucket
```
The registry becomes a stateless API tier; data lives in S3.

---

### Question 379: How do you purge unused images from a private registry?

**Answer:**
The Registry API does not delete blobs automatically.
1.  **Soft Delete:** Delete Manifest via API.
2.  **Hard Delete:** Run Garbage Collection (GC) tool on the registry storage.

---

### Question 380: What’s the difference between Docker Hub and GHCR?

**Answer:**
- **Hub:** The default. Huge public library.
- **GHCR (GitHub Container Registry):** Integrated tightly with GitHub Actions/Packages. Better permission mapping to Repo access.

---

## 🌀 Container Lifecycle, Management, and Orchestration (Questions 381-390)

### Question 381: What happens to volume data when a container is force removed?

**Answer:**
`docker rm -f` removes the **Container**.
It does **NOT** remove the **Volume**.
Data is safe.

---

### Question 382: What’s the difference between restart policies: always, unless-stopped, and on-failure?

**Answer:**
- **always:** Restarts even if I manually stopped it (on daemon restart).
- **unless-stopped:** Restarts unless I manually stopped it (stays stopped on daemon restart).
- **on-failure:** Restarts only if crash (exit > 0).

---

### Question 383: How does Docker handle `OOMKilled` containers?

**Answer:**
Kernel kills the process. Docker sees Exit 137.
If Restart Policy is set, Docker restarts it.
*Warning:* Causes "CrashLoop" if memory barely suffices.

---

### Question 384: How do you restore a container from a committed image?

**Answer:**
If you `docker commit broken-container backup`.
You can `docker run -it backup /bin/bash` to explore the filesystem state at the moment of commit.

---

### Question 385: How do you batch restart services with downtime awareness?

**Answer:**
Use `docker service update --force --update-parallelism 1 ...` in Swarm.
Ensures minimal disruption.

---

### Question 386: What are the limitations of using Docker for system-level daemons?

**Answer:**
Docker containers usually lack `systemd` (PID 1 is the app).
Services relying on DBus, Systemd, or specific Kernel modules might fail or require extensive privileges.

---

### Question 387: What’s the difference between foreground and background containers?

**Answer:**
- **Foreground:** `docker run` (Stdout attched). Ctrl+C kills it.
- **Background:** `docker run -d` (Detached). Runs until `stop` called.

---

### Question 388: How do you enforce container exit traps and signals?

**Answer:**
In Shell scripts (Entrypoint):
```bash
trap 'echo Terminating; exit' SIGTERM
```
Ensures cleanup scripts run.

---

### Question 389: How do you manage container dependencies manually?

**Answer:**
Without Compose? Hard.
Scripts:
`docker run -d --name db ...`
`until nc -z localhost 5432...`
`docker run --link db ...` (Legacy).

---

### Question 390: How do you create auto-healing containers without Kubernetes?

**Answer:**
1.  **Restart Policies.**
2.  **Health Checks.**
3.  **Local sidecar:** Running a script that monitors `docker events` and acts.

---

## 🧪 Experimental and Future Docker Tech (Questions 391-400)

### Question 391: What is `docker init` (experimental)?

**Answer:**
New CLI command. Scans your project (Go, Node, Python) and **automatically generates** Dockerfile and compose.yaml.

---

### Question 392: How do you use Docker extensions with Docker Desktop?

**Answer:**
One-click install plugins in Docker Desktop GUI.
Examples: Snyk, Disk Usage, Logs Explorer.
They run as hidden containers mapped to the GUI.

---

### Question 393: What is Wasm+Docker and where is it headed?

**Answer:**
**WebAssembly.** Docker can now run Wasm modules alongside Linux containers (`containerd-wasm-shim`).
- **Pros:** Near instant start, total sandbox, arch-independent.
- **Cons:** Limited ecosystem.

---

### Question 394: Can Docker support GPU runtime isolation natively?

**Answer:**
Use the **NVIDIA Container Toolkit**.
`docker run --gpus all ...`
It injects the GPU drivers from host into container.

---

### Question 395: What’s Docker's future in serverless computing?

**Answer:**
Docker images are the packaging standard for Lambda/Cloud Run.
The "Runtime" varies, but the "Package" is Docker.

---

### Question 396: How are containers being extended for microVM use cases (Firecracker)?

**Answer:**
Using **Kata Containers** or **Firecracker**.
They look like Docker (OCI spec) but run inside a lightweight VM for hard isolation.

---

### Question 397: How does Docker interact with eBPF monitoring tools?

**Answer:**
eBPF runs on Host Kernel. It "peers" into the container capabilities/calls without needing agents inside the container.

---

### Question 398: What are the experimental features in Docker CE?

**Answer:**
Enable via `dockerd --experimental`.
Includes things like Wasm support, Checkpoint/Restore (CRIU).

---

### Question 399: What is BuildKit remote cache and how is it used?

**Answer:**
Pushing cache blobs to registry (`type=registry`).
Allows team members to share build cache layer.

---

### Question 400: What is the roadmap of Docker in the AI & ML ecosystem?

**Answer:**
Focus on **Reproducible ML**.
Docker provides pre-configured AI/ML stacks (PyTorch/Tensorflow) with GPU drivers pre-installed (`--gpus`).
GenAI Stack (partnership with Neo4j, LangChain).
