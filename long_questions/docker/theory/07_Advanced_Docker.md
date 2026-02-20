# ⚙️ **Advanced Docker Interview Questions (81–90)**

---

### 81. What is the lifecycle of a Docker container?
"A container goes through these states:

**Created** → container is created but not started. **Running** → process is executing. **Paused** → process is frozen (cgroup freezer, SIGSTOP). **Restarting** → container is restarting due to restart policy. **Removing** → `docker rm` is in progress. **Exited** — process has finished or was killed. **Dead** — container couldn't be removed (rare, filesystem error).

Transitions: `docker create` → Created, `docker start` → Running, `docker pause` → Paused, `docker stop` → Exited, `docker rm` → removed."

#### In-depth
The exit code of a container reveals how it terminated. Exit code 0: clean exit. Exit code 1: application error. Exit code 137: killed by OOM killer (exit 128 + signal 9). Exit code 143: killed by SIGTERM (128 + 15). Exit code 125/126/127: Docker itself couldn't run the command. Monitoring exit codes in your orchestrator is crucial for distinguishing between app bugs, resource limits, and deployment issues.

---

### 82. What is the difference between `exec` and `attach`?
"`docker exec` **runs a new process** inside the container's namespaces. It spawns a fresh process (like `/bin/sh`) alongside the container's PID 1. Killing the exec'd process has no effect on the container.

`docker attach` **connects your terminal to PID 1's streams** (stdin/stdout/stderr). It's like joining the container's initial session. Sending SIGINT (`Ctrl+C`) via attach sends the signal to PID 1, potentially terminating the container.

I always use `exec` for debugging — `docker exec -it container sh`. `attach` is rarely needed."

#### In-depth
`docker exec` is implemented by the Docker daemon calling the container runtime (runc) to `exec` a new process into the existing container namespaces (PID, network, mount, UTS, IPC). The new process is a sibling of PID 1, sharing all namespaces. This is exactly how Kubernetes' `kubectl exec` works — it calls the CRI runtime's `ExecSync` RPC.

---

### 83. How do you handle stateful applications in Docker?
"The key principle: **separate state from the container**.

State lives in **volumes** (named Docker volumes or cloud storage), never in the container's writable layer. The app container is stateless and replaceable — volumes persist through container restarts, updates, and replacements.

For databases: use a named volume for data directory, use a stable container name for DNS consistency, and run database backups independently (`docker run --volumes-from db postgres-backup`). Consider managed cloud databases (RDS, Cloud SQL) over self-hosted containers for critical production data."

#### In-depth
Running stateful apps (databases, message queues) in containers is legitimate but requires more operational care than stateless services. Key concerns: **data integrity** during container replacement (ensure clean shutdown, WAL flushing), **storage performance** (use local NVMe volumes for databases, not network storage), **backup/restore runbooks**, and **version upgrade paths** that include data migration steps.

---

### 84. Can you run a database inside Docker? Pros/cons?
"**Yes**, and it's very common. Postgres, MySQL, Redis, MongoDB all have official Docker images.

**Pros**: fast local development setup, version consistency across team, easy in CI for integration tests, isolated from host OS, easy to reset (rm + recreate).

**Cons**: persistent storage adds complexity (volumes), I/O performance depends on storage driver (can be slower than bare metal), container crashes can mean downtime (no automatic failover), upgrades require careful data migration, not ideal for production DBs requiring high availability.

My rule: Docker databases in dev/test, managed cloud DBs (RDS, Cloud SQL) in production."

#### In-depth
I/O performance is the critical concern. Docker volumes on Linux use the OverlayFS or direct mount, which adds some overhead. For write-intensive databases on Docker, use **bind mounts to an SSD** or **volume drivers that bypass the storage driver** (like `local` with `bind` options). Benchmarks show Docker postgres can achieve 80-95% of bare-metal performance with proper configuration.

---

### 85. What is Init system in containers?
"By default, Docker containers run your app as **PID 1**. PID 1 has special responsibilities: it must reap zombie processes (orphaned child processes). Most applications aren't written to do this — they just run and ignore zombie children.

An init system inside the container handles this. Docker provides `--init` flag which uses **tini** as a minimal init process. It becomes PID 1 and reaps zombie children, then exec's your app.

`docker run --init myapp` is the easiest fix for zombie process accumulation."

#### In-depth
The zombie problem is real: if your app forks child processes (many servers and workers do) and those children die before being `wait()`-ed on, they become zombies. Zombies hold a PID forever — with enough zombies, the system runs out of PIDs. `tini` or `dumb-init` solve this by being a proper PID 1 that calls `wait()` on all children. Kubernetes' containerd defaults to `pause` as PID 1 but apps still need to reap their own children.

---

### 86. How do you limit CPU/memory for containers?
"At runtime, I use Docker flags:

**Memory**: `--memory=512m` (hard limit, triggers OOM kill), `--memory-reservation=256m` (soft limit, best-effort), `--memory-swap=1g` (includes swap; set equal to `--memory` to disable swap).

**CPU**: `--cpus=1.5` (use up to 1.5 CPU cores), `--cpu-shares=512` (relative weight, default 1024), `--cpuset-cpus=0,1` (pin to specific CPU cores).

In Compose/Swarm: use `deploy.resources.limits` and `deploy.resources.reservations`."

#### In-depth
`--memory` vs `--memory-reservation`: the reservation is a soft guarantee — Docker will try to keep this much memory available for the container. When the host is memory-constrained, containers using more than their reservation face more aggressive reclamation. This is useful for mixed workloads: reserve a baseline for critical services while allowing burst beyond reservation when memory is available.

---

### 87. What are container registries alternatives to Docker Hub?
"The main alternatives:

- **AWS ECR** (Elastic Container Registry): native AWS integration, per-GB pricing, IAM auth, lifecycle policies
- **Google Artifact Registry**: multi-format (Docker, Maven, npm), GCP-native
- **GitHub Container Registry (GHCR)**: free for public, integrated with GitHub Actions
- **GitLab Container Registry**: built into GitLab, no extra configuration
- **Harbor**: open-source, self-hosted, supports Helm charts, RBAC, image scanning
- **JFrog Artifactory**: enterprise, multi-format, proxies public registries

I choose based on where the deployment target is — ECR for EKS, GCR for GKE."

#### In-depth
For CI/CD performance, the registry location matters enormously. An ECR registry in `us-east-1` pulling to an EKS cluster in `us-east-1` is free (no egress) and fast (AWS backbone). The same pull from Docker Hub costs egress fees and is subject to rate limits. Collocating your registry with your cluster is a critical production architecture decision.

---

### 88. How do you rollback to a previous image version?
"In Swarm: `docker service update --rollback myservice` — instantly reverts to the config before the last update (Docker stores the previous service spec).

Alternatively, update to a specific previous tag: `docker service update --image myregistry/app:v1.4 myservice`.

In Kubernetes: `kubectl rollout undo deployment/myapp`.

The key to reliable rollbacks: always tag images with immutable identifiers (Git SHA, semantic version) rather than overwriting `latest`. A rollback is simply `service update --image <old-tag>`."

#### In-depth
Configure Swarm rollback behavior: `docker service update --rollback-monitor=120s --rollback-max-failure-ratio=0.3`. This means: monitor for 120s after each task update; if 30% of tasks fail, automatically trigger rollback. This automated rollback guard prevents a bad deployment from propagating across all replicas — it stops and self-heals.

---

### 89. What is overlay network?
"An overlay network is a **virtual network that spans multiple Docker hosts**, enabling containers on different physical/virtual machines to communicate as if on the same LAN.

Overlay networks use **VXLAN** (Virtual Extensible LAN) to encapsulate container traffic in UDP packets tunneled over the host network. Docker Swarm creates overlay networks automatically for services.

I use overlays for multi-node Swarm deployments: `docker network create -d overlay mynet; docker service create --network mynet ...`. Each node's containers can reach each other by hostname."

#### In-depth
VXLAN encapsulates L2 Ethernet frames in L4 UDP datagrams (port 4789). This lets containers communicate at L2 (same Ethernet segment) even when the underlying hosts are on different L3 subnets — e.g., hosts in different availability zones or data centers. The overhead is a small header addition per packet. Kernel-based VXLAN is very efficient — Docker overlay networks add ~5% latency in benchmarks.

---

### 90. What is the difference between `alpine`, `debian`, `ubuntu` images?
"These are the common Linux base images with very different philosophies:

**Alpine Linux** (~3MB): uses `musl libc` and `busybox` instead of `glibc` and GNU coreutils. Minimal, security-focused, package managed with `apk`. Preferred for production containers.

**Debian** (~25MB): standard `glibc`, `apt` package manager, wide package availability. Stable and widely compatible.

**Ubuntu** (~29MB): based on Debian, user-friendly, newer package versions via PPAs. Familiar to most developers.

My default: Alpine for production images, Debian (slim variant) when Alpine compatibility is an issue."

#### In-depth
Alpine's use of `musl libc` instead of `glibc` can cause subtle incompatibilities with pre-compiled binaries that link against `glibc` (many commercial software, some Go cgo applications). The fix: use `debian:slim` instead, or for Go applications — compile statically (`CGO_ENABLED=0`) so there's no libc dependency at all. The true minimal choice for Go apps: scratch + CA certificates.

---
