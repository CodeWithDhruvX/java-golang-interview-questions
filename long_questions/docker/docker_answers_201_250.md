## 🧱 Core Docker Internals & Low-Level Details (Questions 201-210)

### Question 201: What is the difference between containers and sandboxing?

**Answer:**
- **Sandboxing:** A restrictive environment to run code securely (e.g., Chrome tabs). Focus is solely on isolation and security.
- **Containers:** A packaging format that includes dependencies, OS, and settings. Containers *use* sandboxing techniques (Namespaces/Cgroups) but also provide portability and environmental consistency.
*All containers are sandboxes, but not all sandboxes are containers.*

---

### Question 202: What is the Docker Engine API and how do you use it?

**Answer:**
It is a REST API that the Docker Daemon exposes.
**Usage:**
- **Socket:** Unix socket `/var/run/docker.sock`.
- **Curl Example:**
  ```bash
  curl --unix-socket /var/run/docker.sock http://localhost/containers/json
  ```
- **SDK:** Use Go, Python, or Java clients to control Docker programmatically (e.g., CI tools, Orchestrators).

---

### Question 203: How does Docker manage container lifecycle states internally?

**Answer:**
Docker uses a state machine managed by `containerd`.
States:
1.  **Created:** Config created, no process.
2.  **Running:** Process active.
3.  **Paused:** Cgroup freeze.
4.  **Restarting:** In loop.
5.  **Exited:** Process died.
6.  **Dead:** Unremovable (filesystem error).

---

### Question 204: What happens under the hood when you run `docker run`?

**Answer:**
1.  **CLI** sends REST request to `Dockerd`.
2.  **Dockerd** checks if image exists (pulls if needed).
3.  **Dockerd** calls `containerd` to create a container object.
4.  **containerd** uses `runc` (OCI runtime) to set up Namespaces/Cgroups.
5.  **runc** executes the entrypoint process.
6.  **runc** exits, leaving the process running (managed by `containerd-shim`).

---

### Question 205: Explain the Docker image digest.

**Answer:**
A **Digest** is an immutable identifier (SHA256 hash) for the image content.
Unlike tags (`v1`, `latest`) which can be overwritten, a digest always points to exactly the same bits.
`ubuntu@sha256:45b23d...`

---

### Question 206: What is the Docker storage driver?

**Answer:**
The component responsible for managing the filesystem layers and the writable container layer.
**Drivers:**
- **overlay2:** (Standard)
- **fuse-overlayfs:** (Rootless)
- **btrfs/zfs:** (Advanced features)
It handles the Copy-on-Write (CoW) operations.

---

### Question 207: Compare OverlayFS and aufs in Docker.

**Answer:**
- **AUFS:** The original driver. Multiple directories stacked. Efficient memory usage. Removed from mainline kernel. Merged only if specific patchsets applied.
- **OverlayFS (Overlay2):** Mainline Linux kernel support. Faster inode lookup (page cache sharing). Simpler design. The industry standard now.

---

### Question 208: What is a Docker union mount?

**Answer:**
A mechanism to combine multiple directories (layers) into a single virtual directory.
- **LowerDir:** Read-only image layers.
- **UpperDir:** Writable container layer.
- **Merged:** What the container sees (Files in Upper "mask" files in Lower).

---

### Question 209: How does Docker manage layered files in copy-on-write systems?

**Answer:**
When you edit a file that exists in the image:
1.  Docker searches for the file in the image layers.
2.  It copies the **entire file** from the read-only layer to the writable upper layer.
3.  The write operation happens on the copy.
4.  The lower file is hidden.
*Note: Modifying large files can be slow due to this copy operation.*

---

### Question 210: What is the role of containerd in Docker's architecture?

**Answer:**
`containerd` acts as the daemon that manages the complete container lifecycle.
- It abstracted the logic away from the rigid Docker Daemon.
- It interacts with `runc` to spawn containers.
- It pushes/pulls images to registries.
- It allows Kubernetes to run containers without the full Docker Engine.

---

## 🚧 Advanced Dockerfile Mechanics (Questions 211-220)

### Question 211: How do you use `ARG` vs `ENV` in Dockerfile?

**Answer:**
- **ARG:** Available **only during build time**. Gone in the running container.
  `docker build --build-arg VER=1.0 .`
- **ENV:** Available **during build AND runtime**. Persists in the container.
  `ENV PATH=/app/bin`

---

### Question 212: How can you dynamically set labels in a Dockerfile?

**Answer:**
Use `ARG` combined with `LABEL`.
```dockerfile
ARG COMMIT_HASH
LABEL version="${COMMIT_HASH}"
```
Build:
`docker build --build-arg COMMIT_HASH=$(git rev-parse --short HEAD) .`

---

### Question 213: How do you split Dockerfiles for better caching?

**Answer:**
Separate **Infrastructure** (rarely changes) from **Application** (frequently changes).
1.  Install OS packages.
2.  Copy dependency manifests (package.json/go.mod).
3.  Install dependencies.
4.  Copy Source Code.
5.  Build.

---

### Question 214: What is a good strategy for base image versioning?

**Answer:**
Use **Specific, Deterministic Tags**.
- `node:14` (Risky - changes minor versions).
- `node:14.17.0-alpine3.13` (Best - locks Node, OS, and Patch).
This ensures reproducible builds in the future.

---

### Question 215: Why should you pin image tags instead of using `latest`?

**Answer:**
`latest` is rolling.
- **Build Breakage:** A new `latest` might introduce breaking changes.
- **Debugging:** Hard to know which version you ran yesterday.
- **Security:** Might unknowingly pull a vulnerable version.

---

### Question 216: How do you ensure idempotent Dockerfile builds?

**Answer:**
1.  Pin base images (`ubuntu:20.04`, not `ubuntu`).
2.  Pin package versions (`apt-get install nginx=1.18`).
3.  Avoid downloading from `latest` URLs (download specific versioned tarballs).
4.  Use `copy` instead of `git clone`.

---

### Question 217: How can you minimize Docker build context?

**Answer:**
The build context is the directory sent to the daemon.
1.  **Root:** Don't build from `/` or `/home/user`.
2.  **.dockerignore:** Explicitly ignore `.git`, `node_modules`, `build` folders, and oversized assets.

---

### Question 218: What’s the use of `USER` in Dockerfile?

**Answer:**
Sets the UID/GID for subsequent instructions (`RUN`, `CMD`, `ENTRYPOINT`).
Essential for security (Running as non-root).
```dockerfile
RUN groupadd -r app && useradd -r -g app app
USER app
CMD ["node", "app.js"]
```

---

### Question 219: How can you multi-target builds in Docker?

**Answer:**
Use **BuildKit Stages**.
```dockerfile
FROM base AS dev
RUN install-dev-tools

FROM base AS prod
COPY . .
```
Build: `docker build --target prod .`

---

### Question 220: How do you debug failed `RUN` commands during build?

**Answer:**
Docker build (without BuildKit) leaves an intermediate container before the failure.
1.  Find the ID of the last successful layer.
2.  `docker run -it <last_layer_id> /bin/bash`.
3.  Run the failed command manually to see the error.

---

## 🧪 Advanced Testing and QA Practices in Docker (Questions 221-230)

### Question 221: How can you test container health during CI?

**Answer:**
1.  **Docker Compose:** `depends_on` with `condition: service_healthy`.
2.  **Wait Script:** Use `wait-for-it.sh` script to block until a port is open.
3.  **Healthcheck:** Inspect the container status after start.

---

### Question 222: How do you simulate network failure in a container?

**Answer:**
Use **Pumba** (Chaos testing tool) or `tc` (Traffic Control) inside the container.
- Add latency.
- Drop packets.
```bash
# Using Pumba to delay network by 3s
docker run -v /var/run/docker.sock:/var/run/docker.sock gaiaadm/pumba netem --duration 5m delay --time 3000 my-container
```

---

### Question 223: How can you perform chaos testing with containers?

**Answer:**
Use tools like **Chaos Mesh** or **Pumba**.
- Randomly kill containers.
- Stress CPU/Memory.
- Corrupt filesystem.
Verify if the orchestrator (Swarm/K8s) recovers the service.

---

### Question 224: How do you use `testcontainers` in Docker-based testing?

**Answer:**
`Testcontainers` is a library (Java, Go, Python, Node) that lets unit tests programmatically spin up Docker containers (e.g., throwing away a PostgreSQL DB for a test).
- **Pros:** Real environment, no mocks.
- **Cons:** Slower test execution.

---

### Question 225: How do you snapshot container states for rollback?

**Answer:**
- **Commit:** `docker commit <container_id> backup-image:tag`. (Creates an image from current state).
- **Note:** Not recommended for production (state should be in volumes, not container layer).

---

### Question 226: How do you implement container linting?

**Answer:**
Linting applies static analysis to the `Dockerfile`.
- **Hadolint:** Checks for best practices (e.g., "Pin versions", "Use trusted base").
- **CI Step:** Fail build if lint score < 10.

---

### Question 227: How do you benchmark containers across environments?

**Answer:**
Use a standardized containerized benchmark tool.
- **Sysbench:** CPU/Memory/File IO.
- **iPerf:** Network.
Run the exact same image on Dev, Staging, and Prop to compare baseline metrics.

---

### Question 228: How do you run integration tests across multiple services using Docker Compose?

**Answer:**
1.  Define `docker-compose.test.yml`.
2.  Include the SUT (System Under Test) and a "Tester" container (contains test scripts).
3.  `docker-compose up --exit-code-from tester`.
   - If tester passes (exit 0), CI passes.
   - If tester fails, CI fails.

---

### Question 229: What is Docker check and how is it different from `healthcheck`?

**Answer:**
- **Healthcheck:** Runtime check (is the app listening?).
- **Docker Check (Scout/Scan):** Static analysis check (does the image have vulnerabilities?). *Note: Terminology varies, usually refers to `docker scan` or `docker scout`.*

---

### Question 230: How do you test backward compatibility with old image versions?

**Answer:**
1.  Spin up database with **Old Image**.
2.  Populate data.
3.  Upgrade container to **New Image**.
4.  Run migration scripts.
5.  Verify data integrity.

---

## 📈 Monitoring & Observability (Questions 231-240)

### Question 231: How can you integrate Docker with Prometheus?

**Answer:**
1.  **Daemon Metrics:** Configure `dockerd` to expose metrics (`"metrics-addr" : "0.0.0.0:9323"`). Prometheus scrapes this for engine health.
2.  **Container Metrics:** Use **cAdvisor**. It runs as a container, reads cgroups, and exposes Prometheus-compatible metrics for ALL containers.

---

### Question 232: What are some logging drivers available in Docker?

**Answer:**
- **json-file:** (Default) Local JSON.
- **syslog:** Standard Syslog.
- **journald:** Linux Journal.
- **fluentd:** Fluentd forwarder.
- **splunk:** Direct to Splunk.
- **awslogs:** AWS CloudWatch.

---

### Question 233: What is the `json-file` log driver and its drawbacks?

**Answer:**
Writes logs to `/var/lib/docker/containers/.../*.log`.
**Drawbacks:**
- **Disk Space:** Can fill up disk if no rotation is configured.
- **Performance:** Reading large JSON files is slow.
- **Parsing:** JSON overhead.

---

### Question 234: How do you send logs to a remote syslog server?

**Answer:**
`docker run --log-driver=syslog --log-opt syslog-address=udp://1.2.3.4:514 my-app`

---

### Question 235: How do you monitor Docker container network metrics?

**Answer:**
- **Basic:** `docker stats` (RX/TX bytes).
- **Advanced:** cAdvisor exposes `container_network_...` metrics.
- **Deep:** Service Mesh (Linkerd/Istio) or eBPF tools (Cilium) for packet-level observability.

---

### Question 236: What tools can you use to monitor Docker container memory usage?

**Answer:**
1.  **Docker Stats.**
2.  **cAdvisor.**
3.  **Linux Tools:** `top -c -b -n 1` inside container (if allowed).
4.  **OOM Events:** Watch for `OOMKilled` in `docker inspect`.

---

### Question 237: How can you set up alerting for failing containers?

**Answer:**
1.  **Prometheus:** AlertManager rule: `container_last_seen < 30s`.
2.  **Docker Events:** Watch for `die` events and trigger a script (WebHook).
3.  **Orchestrator:** K8s/Swarm handles restarts, but alerts on "RestartLoop".

---

### Question 238: How can you use Grafana dashboards with Docker?

**Answer:**
Stack: Docker -> cAdvisor -> Prometheus -> Grafana.
Import a standard Docker Dashboard (e.g., ID 893) into Grafana to verify CPU, Memory, and Net usage per container.

---

### Question 239: How do you debug container I/O bottlenecks?

**Answer:**
1.  **iotop:** Run on host to see which process (container PID) is hammering disk.
2.  **docker stats:** Check Block I/O.
3.  **Volume type:** Ensure you aren't writing heavily to the Container Layer (slow CoW). Use Volumes.

---

### Question 240: How do you get the exit code of a container?

**Answer:**
Use `docker inspect`.
```bash
docker inspect my-container --format='{{.State.ExitCode}}'
```
- **0:** Success.
- **137:** SIGKILL (OOM).
- **1/255:** App Error.

---

## 💾 Advanced Volumes, Filesystems & Storage (Questions 241-250)

### Question 241: How can you mount config files as volumes?

**Answer:**
Use a **Bind Mount** for single files.
```bash
docker run -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro nginx
```
*Note: Bind mounting a single file hides the original file in the image.*

---

### Question 242: How do you share data between containers securely?

**Answer:**
1.  **Shared Volume:** Both containers mount the same named volume.
2.  **Permissions:** Set specific GID on the volume so only those containers can read it.

---

### Question 243: What happens to volumes when containers are removed?

**Answer:**
- **Named Volumes:** Persist. Must be manually removed (`docker volume rm`).
- **Anonymous Volumes:** Are **NOT** removed automatically unless you used `docker rm -v`. If not, they become "dangling volumes".
- **Bind Mounts:** Host files are untouched.

---

### Question 244: How do you manage storage for stateful containers across clusters?

**Answer:**
Swarm/Docker alone is bad at this.
**Solutions:**
1.  **Cloud Volumes:** EBS/EFS (Using plugins like RexRay).
2.  **NFS:** Mount NFS share on all nodes.
3.  **Kubernetes:** Use PV (Persistent Volume) and PVC.

---

### Question 245: How do you encrypt volume data in Docker?

**Answer:**
- Docker doesn't encrypt volumes natively.
- **Solution:** Use an encrypted filesystem on the host (LUKS) and mount that directory.
- **Enterprise:** Use storage plugins (Portworx) that support encryption at rest.

---

### Question 246: What happens to volume data in Docker swarm rolling update?

**Answer:**
Volumes are detached from the old container and re-attached to the new container on the **same node**.
*Risk:* If the task is rescheduled to a *different node*, local volumes are **LOST** (unless using Network Storage like NFS/EFS).

---

### Question 247: Can volumes be backed up directly? How?

**Answer:**
Yes. Since they are just directories on the host (`/var/lib/docker/volumes/name/_data`).
- **Root access:** `cp -r ...`
- **Docker way:** Spin up a helper container to `tar` the volume content.

---

### Question 248: How do you use third-party volume plugins?

**Answer:**
Install plugin: `docker plugin install rexray/ebs`.
Create volume:
```bash
docker volume create -d rexray/ebs --name my-ebs-vol
```

---

### Question 249: What is lazy loading of volumes?

**Answer:**
Some network volume drivers support lazy loading, where data is only fetched from the remote source when requested by the application, speeding up container start time.

---

### Question 250: Can you mount a volume from one container into another running container?

**Answer:**
Yes, using `--volumes-from`.
```bash
docker run --name data-store -v /data alpine
docker run --volumes-from data-store alpine ls /data
```
*Note: This copies the mount definition, not the data itself.*
