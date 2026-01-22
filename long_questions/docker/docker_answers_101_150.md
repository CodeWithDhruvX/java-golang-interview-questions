## 🧠 Conceptual & Deep-Dive Docker Questions (Questions 101-110)

### Question 101: What is the OCI (Open Container Initiative)?

**Answer:**
The OCI is a Linux Foundation project to design open standards for operating-system-level virtualization, most importantly Linux containers.
**Two Main Specifications:**
1.  **Runtime Specification (runtime-spec):** Defines how to run a "filesystem bundle" (lifecycle of a container). `runc` is the reference implementation.
2.  **Image Specification (image-spec):** Defines the archive format of OCI container images.

---

### Question 102: How is Docker related to containerd and runc?

**Answer:**
Docker has been modularized over time:
- **Docker Engine (dockerd):** The high-level daemon that handles the API, image management, and orchestration.
- **containerd:** An industry-standard container runtime (CNCF graduated) that manages the complete container lifecycle (pulling images, storage, execution). Docker uses containerd.
- **runc:** The CLI tool for spawning and running containers according to the OCI spec. containerd uses runc to actually create the containers.

**Chain:** `Docker CLI` -> `dockerd` -> `containerd` -> `containd-shim` -> `runc`.

---

### Question 103: What is the architecture of Docker Engine?

**Answer:**
Client-Server architecture:
1.  **Docker Daemon (`dockerd`):** The server process that manages Docker objects (images, containers, networks, volumes).
2.  **REST API:** The interface used by the CLI and other tools to talk to the daemon.
3.  **Docker CLI:** The command-line client (`docker` command) that uses the REST API to control the daemon.

---

### Question 104: What is the difference between Docker CLI and Docker API?

**Answer:**
- **Docker CLI:** The user-facing command-line tool (`docker run`). It is a wrapper that makes HTTP requests.
- **Docker API:** The underlying RESTful API provided by the Docker daemon. You can use `curl` or libraries (Python/Go SDKs) to interact with it directly.
  - Example: `GET /containers/json` (Lists containers).

---

### Question 105: What is the role of dockerd?

**Answer:**
`dockerd` is the persistent process (daemon) that manages containers.
**Responsibilities:**
- Listening for API requests.
- Managing images, containers, networks, and volumes.
- Delegating container execution to `containerd`.

---

### Question 106: How does Docker handle layered filesystems?

**Answer:**
Docker uses a **Union File System (UnionFS)**.
- **Images** are composed of multiple read-only layers.
- **Containers** add a thin "Read-Write" layer on top.
- When you modify a file in a container, Docker uses a **Copy-on-Write (CoW)** strategy: it copies the file from the read-only layer to the writable layer before modifying it.

---

### Question 107: Explain union file systems in Docker (AUFS, OverlayFS, etc.).

**Answer:**
Union filesystems allow files and directories of separate file systems, known as branches, to be transparently overlaid, forming a single coherent file system.
- **Overlay2:** The current default and recommended driver for Linux. It is fast and simpler than AUFS.
- **AUFS:** Older, originally used by Docker.
- **btrfs/zfs:** Snapshot-capable filesystems (less common).
- **VFS:** For testing (slow, no copy-on-write).

---

### Question 108: What are Docker manifests?

**Answer:**
A Docker Manifest gives information about the image, such as layers, size, and digest.
- **Manifest List (Multi-arch):** A "Fat Manifest" that lists different image variants for different architectures (amd64, arm64). When you pull `node:alpine`, Docker checks the manifest list to pull the correct binary for your CPU.

---

### Question 109: What is the difference between Docker manifests and Docker images?

**Answer:**
- **Image:** The actual binary data (layers, config, filesystem).
- **Manifest:** Metadata description JSON that points to the image layers.
  - While an "image" is what you run, the "manifest" is what the registry stores and serves to the client to describe *how* to pull that image.

---

### Question 110: How does the Docker image ID work?

**Answer:**
The Image ID is a SHA256 hash of the image's JSON configuration object.
- It uniquely identifies the image content.
- If the configuration or any layer changes, the hash changes (Immutability).
- Also known as `Image Digest` when referring to the manifest hash in a registry.

---

## ⚙️ Docker Networking (Advanced) (Questions 111-120)

### Question 111: What is a MACVLAN network in Docker?

**Answer:**
MACVLAN allows you to assign a MAC address to a container, making it appear as a physical device on your network.
- **Use Case:** Legacy apps that require direct connection to the physical network, or monitoring traffic/VLANs.
- **Disadvantage:** Requires promiscuous mode on the host network interface (often blocked by cloud providers).

---

### Question 112: How does bridge networking work in Docker?

**Answer:**
- Docker creates a virtual bridge (usually `docker0`) on the host.
- Each container gets a virtual ethernet interface (`veth`) connected to this bridge.
- Docker uses `iptables` NAT (Network Address Translation) to route traffic from the host port to the container port.

---

### Question 113: What is host networking, and when should you use it?

**Answer:**
The container shares the host's networking namespace directly. It does not get its own IP; it uses the host's IP.
- **Performance:** low latency (no NAT overhead).
- **Cons:** Port conflicts (cannot run two containers on port 80).
- **Use Case:** High-performance networking apps, or protocol complexity (like VoIP/SIP).
  ```bash
  docker run --network host nginx
  ```

---

### Question 114: How do you create a custom Docker network?

**Answer:**
Using `docker network create`. Custom bridge networks provide automatic DNS resolution (unlike default bridge).

**Command:**
```bash
docker network create \
  --driver bridge \
  --subnet 172.18.0.0/16 \
  my-net
```

---

### Question 115: How do you inspect a Docker network?

**Answer:**
Use `docker network inspect` to see configuration and connected containers.

**Command:**
```bash
docker network inspect bridge
```
*Useful to find the internal IP address assigned to a container.*

---

### Question 116: Can two containers on different networks communicate?

**Answer:**
No, not by default. This is part of Docker's isolation.
**Solution:**
1.  Connect a container to both networks: `docker network connect net2 container1`.
2.  Route traffic through the host (expose ports).

---

### Question 117: What is port mapping, and how is it done?

**Answer:**
Port mapping (Publishing) forwards a port on the host machine to a port inside the container using NAT.

**Syntax:** `-p host_port:container_port`
```bash
# Host 8080 -> Container 80
docker run -p 8080:80 nginx
```

---

### Question 118: How do you debug Docker network issues?

**Answer:**
1.  **Inspect:** `docker network inspect` (Check IP/Gateways).
2.  **Netshoot:** Run a container with networking tools attached to the target container's namespace.
    ```bash
    docker run -it --net container:<target_id> nicolaka/netshoot
    ```
3.  **Ping:** Test connectivity.
4.  **Logs:** Check daemon logs for IP conflicts.

---

### Question 119: How do you restrict container-to-container communication?

**Answer:**
In a custom bridge network, all containers can talk. To restrict:
- **Iptables:** Use Docker's `iptables` rules (`DOCKER-USER` chain) to drop traffic.
- **icc=false:** Set `com.docker.network.bridge.enable_icc=false` (Inter-Container Communication) when creating the network (deprecated in favor of network policies in K8s, but valid in Docker).

---

### Question 120: How do you implement DNS-based service discovery in Docker?

**Answer:**
On any user-defined network (not default bridge), Docker runs an embedded DNS server (127.0.0.11).
- It automatically resolves container names to their private IPs.
- **Example:** In a `web` container, `ping db` works if the other container is named `db`.
- **Aliases:** You can give network aliases (`--network-alias`).

---

## 🔒 Docker Security & Compliance (Questions 121-130)

### Question 121: What is a Docker security profile (seccomp)?

**Answer:**
**Seccomp (Secure Computing Mode):** Filters the system calls (syscalls) a process can make to the kernel.
- Docker applies a default seccomp profile that blocks ~44 syscalls (like `reboot`, `swapoff`) out of ~300+.
- You can pass a custom JSON profile using `--security-opt seccomp=profile.json`.

---

### Question 122: What is AppArmor and how does it work with Docker?

**Answer:**
AppArmor (Application Armor) is a Linux kernel security module that allows you to restrict programs' capabilities with per-program profiles.
- Docker auto-generates and loads a default profile: `docker-default`.
- It restricts file access, network access, and capability usage.

---

### Question 123: What is SELinux and how is it used with Docker?

**Answer:**
SELinux (Security-Enhanced Linux) provides a labeling system for mandatory access control (MAC).
- Docker supports SELinux labeling.
- If enabled, the Docker daemon labels container processes/files (`svirt_sandbox_file_t`) to prevent them from accessing host files unless explicitly allowed (the `z` or `Z` bind mount option).

---

### Question 124: What is the purpose of capabilities in Docker?

**Answer:**
Linux "Capabilities" break down the power of the "root" user into small, distinct privileges (e.g., `CAP_CHOWN`, `CAP_NET_ADMIN`).
- Docker drops most capabilities by default.
- You can Add or Drop specific ones.
  ```bash
  docker run --cap-add=NET_ADMIN nginx
  ```

---

### Question 125: How do you restrict Linux capabilities for a container?

**Answer:**
Use `--cap-drop`. Good security practice is to drop ALL and add only what's needed.

**Example:**
```bash
docker run --cap-drop=ALL --cap-add=CHOWN my-app
```

---

### Question 126: How do you enable Docker Content Trust (DCT)?

**Answer:**
DCT enforces digital signatures.
1.  **Enable Environment Variable:**
    ```bash
    export DOCKER_CONTENT_TRUST=1
    ```
2.  **Effect:** `docker pull` or `docker run` will fail if the image is not signed by a trusted publisher.

---

### Question 127: What is the `--security-opt` flag used for?

**Answer:**
It is a generic flag to pass security options.
- Load AppArmor profile: `--security-opt apparmor=my-profile`
- Load Seccomp profile: `--security-opt seccomp=profile.json`
- Disable SELinux: `--security-opt label=disable`
- No New Privileges: `--security-opt no-new-privileges`

---

### Question 128: How do you isolate sensitive data from logs in containers?

**Answer:**
1.  **Don't Log Secrets:** Ensure code doesn't print passwords/tokens to `stdout`.
2.  **Log Driver Splunk/Syslog:** Send logs to a secure, encrypted destination.
3.  **RBAC:** Restrict access to `docker logs`.
4.  **Sidecar:** Use a sidecar to filter logs before shipping.

---

### Question 129: How does Docker protect against container escape?

**Answer:**
Multiple layers of defense:
1.  **Namespaces:** Hides host processes.
2.  **Cgroups:** Limits resources.
3.  **Capabilities:** Restricts root power.
4.  **Seccomp/AppArmor:** Filters syscalls.
5.  **User Namespaces:** Maps container root to a non-privileged user on host.

---

### Question 130: What are common CVEs related to Docker?

**Answer:**
- **runC vulnerability (CVE-2019-5736):** Allowed overwriting the host runc binary to gain root.
- **Danging Image vulnerabilities.**
- **HTTP/2 Rapid Reset.**
- **Mitigation:** Keep Docker Engine updated and scan images.

---

## 🚀 Docker Performance & Optimization (Questions 131-140)

### Question 131: How do you analyze Docker container performance?

**Answer:**
1.  **Live:** `docker stats` (CPU/Mem/Net I/O).
2.  **Historical:** Prometheus + cAdvisor + Grafana.
3.  **Inside Container:** `top`, `htop`, `vmstat`.
4.  **Load Testing:** Tools like `Apache Bench (ab)` or `k6` against the containerized app.

---

### Question 132: How do you reduce image size significantly?

**Answer:**
1.  **Alpine Linux:** Use `alpine` base.
2.  **Multi-Stage Build:** Copy only compiled artifacts.
3.  **Chaining RUN:** `RUN apt update && apt install ... && rm -rf /var/lib/apt/lists/*`.
4.  **Docker Slim:** Tool that auto-minifies images.
5.  **Distroless:** Images containing only the app and its runtime dependencies (no shell).

---

### Question 133: How does using Alpine base image impact performance?

**Answer:**
- **Pros:** Smaller size = Faster pull/startup time. Less disk I/O.
- **Cons:** Uses `musl libc` instead of `glibc`. Some C/C++ apps might be slower or buggy. DNS resolution implementation differs.
- **Verdict:** Generally performs well, but verify compatibility.

---

### Question 134: What is the performance difference between containers and VMs?

**Answer:**
- **CPU/RAM:** Containers have near-native performance (negligible overhead). VMs have hypervisor overhead.
- **I/O:** Container NAT can introduce minor network latency. Disk I/O (especially with OverlayFS) adds slight overhead compared to native, but is much faster than VM virtual disks.
- **Startup:** Containers (ms) vs VMs (minutes).

---

### Question 135: How do you enable Docker build cache effectively?

**Answer:**
Order instructions from **least changed** to **most changed**.

**Bad:**
```dockerfile
COPY . .
RUN npm install
```

**Good:**
```dockerfile
COPY package.json .
RUN npm install    # Cached unless package.json changes
COPY . .           # Only this layer rebuilds on code change
```

---

### Question 136: What are the drawbacks of having large Docker images?

**Answer:**
1.  **Slow Deployment:** Longer `docker pull` time on nodes.
2.  **Disk Usage:** Fills up node storage quickly.
3.  **Security:** More packages = Larger attack surface.
4.  **Bandwidth:** Higher network costs for transfer.

---

### Question 137: How do you optimize Docker builds with `--cache-from`?

**Answer:**
In CI/CD, build agents are often ephemeral (no local cache).
Use `--cache-from` to download existing layers from a registry to use as cache source.

```bash
docker pull my-app:latest || true
docker build --cache-from my-app:latest -t my-app:new .
```

---

### Question 138: How do you measure startup time of containers?

**Answer:**
You can inspect the container data.

```bash
# Time command
time docker run --rm alpine echo "start"

# More precision involves Application metrics (time from process start to Ready).
```

---

### Question 139: What causes slow builds in Docker?

**Answer:**
1.  **Large Context:** Sending huge context (files) to daemon (Check `.dockerignore`).
2.  **Network:** Slow download of base images/dependencies.
3.  **Cache Misses:** Poorly ordered Dockerfile instructions.
4.  **I/O:** Heavy disk write operations during build.

---

### Question 140: How do you profile memory and CPU usage of containers?

**Answer:**
1.  **Memory:** `docker stats` shows usage vs limit. Look for `OOMKilled` in `docker inspect`.
2.  **CPU:** Use `docker stats` or Linux `perf` tool on the PID of the container.
3.  **Pprof:** For Go apps, use `pprof` inside the container.

---

## 🔁 Docker CI/CD Integration & GitOps (Questions 141-150)

### Question 141: How do you version Docker images in a CI pipeline?

**Answer:**
Use specific tags linked to Git:
- **Git Commit SHA:** `app:a1b2c3d` (Precise).
- **Git Tag/Release:** `app:v1.0.0` (Stable).
- **Build Number:** `app:build-123`.
- **Avoid:** Relying on `latest`.

---

### Question 142: What is the best strategy to tag Docker images?

**Answer:**
**Dual Tagging Strategy:**
1.  Tag with unique ID (SHA/Build ID) for immutability.
2.  Tag with "Semantic Version" (`v1.2`) for user friendliness.
3.  Tag with `latest` (optional, for convenience).

`docker push myapp:1.2.3` AND `docker push myapp:latest`

---

### Question 143: How do you implement GitOps with Docker images?

**Answer:**
1.  **Code Repo:** Developer commits code.
2.  **CI:** Builds Docker image and pushes to registry.
3.  **Config Repo:** CI updates a manifest file (k8s yaml/helm) with the new image tag (`image: app:v2`).
4.  **CD (ArgoCD/Flux):** Detects change in Config Repo and syncs the cluster.

---

### Question 144: How do you ensure rollback safety in Docker-based deployments?

**Answer:**
1.  **Immutable Tags:** Never overwrite tags like `v1`. Always create `v2`.
2.  **Keep History:** Retain last N versions in registry.
3.  **Health Checks:** Automated rollback if new container fails health probes.

---

### Question 145: How do you automate security scanning in CI for Docker images?

**Answer:**
Integrate scanners as a pipeline step.

**Example (GitHub Actions with Trivy):**
```yaml
- name: Run Trivy vulnerability scanner
  uses: aquasecurity/trivy-action@master
  with:
    image-ref: 'my-app:${{ github.sha }}'
    format: 'table'
    exit-code: '1' # Fail pipeline on High severity
    severity: 'CRITICAL,HIGH'
```

---

### Question 146: What’s the difference between ephemeral containers and long-lived ones in CI?

**Answer:**
- **Ephemeral (Service Containers):** Spun up just for the test (e.g., a Redis DB used for integration testing), destroyed after.
- **Build Containers:** The container where the build runs (e.g., `maven:3-jdk-11`). Also ephemeral.
- **Long-lived:** The artifact produced (the final image) which is deployed to Prod.

---

### Question 147: What is an image promotion strategy in Docker pipelines?

**Answer:**
Moving a single image artifact through environments validation.
1.  **Build** -> Push to `dev-registry`.
2.  **Test in Dev.**
3.  **Promote:** Retag the *exact same image ID* from `dev` to `staging`. (Do NOT rebuild).
4.  **Deploy to Staging.**

---

### Question 148: How do you use Docker in a monorepo setup?

**Answer:**
1.  **Context:** Use complex build contexts or copy logic.
2.  **Selective Build:** Only build Docker images for services that changed.
3.  **Naming:** `repo/service-a:ver`, `repo/service-b:ver`.
4.  **Paths:**
    ```dockerfile
    COPY services/service-a/package.json .
    ```

---

### Question 149: How do you test multi-service containers before pushing to prod?

**Answer:**
Use **Docker Compose** in CI.
1.  Spin up exact stack: `docker-compose -f docker-compose.test.yml up -d`.
2.  Run integration test suite against the exposed ports.
3.  Tear down.

---

### Question 150: How do you manage secrets in CI/CD for Dockerized apps?

**Answer:**
1.  **CI Variables:** Store secrets in CI project settings (encrypted).
2.  **Injection:** Pass as `--build-arg` (Risky! secrets persist in history) or `--secret` (BuildKit specific, Safer).
    ```bash
    docker build --secret id=mysecret,src=key.txt .
    ```
3.  **Runtime:** Inject as ENV vars during deployment, not build.

---
