## ⚡ Docker and Kubernetes (Questions 251-260)

### Question 251: Why was Docker deprecated in Kubernetes 1.20+?

**Answer:**
Kubernetes deprecated the **Dockershim**, the adapter that allowed Kubelet to communicate with Docker Daemon.
- **Reason:** Docker is not CRI (Container Runtime Interface) compliant natively. Maintaining the shim was a burden.
- **Impact:** K8s now communicates directly with `containerd` or `CRI-O`. You can still *build* images with Docker, but K8s nodes don't need the full Docker Daemon installed.

---

### Question 252: What is the difference between Docker and CRI-O?

**Answer:**
- **Docker:** A full-featured platform (CLI, API, Build, Swarm). Heavyweight.
- **CRI-O:** A lightweight container runtime specifically capable of implementing the K8s CRI. It exists *only* to run pods for Kubernetes. It cannot "build" images.

---

### Question 253: Can you use Docker with Minikube? How?

**Answer:**
Yes. Minikube usually runs a Docker daemon inside its VM.
**Tip:** You can reuse the Minikube Docker daemon to build images locally without pushing to a registry.
```bash
eval $(minikube docker-env)
docker build . -t my-app
```

---

### Question 254: What is the role of Docker images in Kubernetes clusters?

**Answer:**
The OCI Image Format (which Docker produces) is the universal standard. Kubernetes (via containerd/CRI-O) can pull and run any standard Docker image. The `Dockerfile` remains the standard way to define K8s workloads.

---

### Question 255: What are some migration steps from Docker to containerd?

**Answer:**
1.  **Nodes:** Uninstall Docker, install containerd. Configure Kubelet to use containerd socket in `--container-runtime-endpoint`.
2.  **Logs:** Docker logs were JSON. Containerd logs are in CRI format. Ensure Fluentd/Logstash parsers are updated.
3.  **Docker-in-Docker:** If CI relied on mounting `/var/run/docker.sock`, it will break. Switch to `Kaniko` or `Buildah`.

---

### Question 256: What tools help convert Docker Compose to Kubernetes YAML?

**Answer:**
- **Kompose:** Official tool. `kompose convert -f docker-compose.yml`.
- **Helm:** Often manual rewrite is preferred to leverage Helm templating.

---

### Question 257: Why is container runtime interface important in Kubernetes?

**Answer:**
CRI allows K8s to be runtime-agnostic. It can plug in Docker, containerd, CRI-O, or gVisor without changing K8s source code.

---

### Question 258: What happens when you use `kubectl run --image`?

**Answer:**
1.  API Server receives request.
2.  Scheduler assigns Pod to a Node.
3.  Kubelet on Node instructs Runtime (containerd) to pull image.
4.  Runtime starts container.
*It mimics `docker run` but goes through the K8s control plane.*

---

### Question 259: How do you push Docker images to be used by Kubernetes pods?

**Answer:**
1.  Push to a registry (Docker Hub/ECR).
2.  Configure K8s `imagePullSecrets` if private.
3.  Define `image: my-registry/app:tag` in Pod spec.

---

### Question 260: How do you debug image pull issues in Kubernetes using Docker?

**Answer:**
If K8s says `ImagePullBackOff`:
1.  SSH into the Node.
2.  Try `docker pull <image>` manually.
3.  Often reveals Auth errors, Network issues, or Typo in tag that K8s events don't show clearly.

---

## 🚢 Docker Swarm Mode Deep Dive (Questions 261-270)

### Question 261: What is the role of Raft consensus in Swarm?

**Answer:**
Manager nodes use the Raft Consensus Algorithm to agree on the state of the cluster (leader election, service definitions).
- Requires (N/2)+1 quorum.
- Ensures consistency even if a manager fails.

---

### Question 262: What is a replicated vs global service?

**Answer:**
- **Replicated:** N copies of the container distributed across the cluster (`--replicas 3`).
- **Global:** One copy of the container on **every single node** (good for agents like Datadog, Splunk).

---

### Question 263: How does Docker Swarm handle node failure?

**Answer:**
1.  Manager detects node is `Down`.
2.  It reschedules the tasks (containers) that were on that node to other healthy nodes to meet the replica count.

---

### Question 264: What are Swarm secrets and how are they encrypted?

**Answer:**
- **Encryption:** Stored encrypted on disk (Raft log) and sent encrypted over network (mTLS).
- **Mounting:** Decrypted only in memory (`tmpfs`) within the container at `/run/secrets/`.
- Never written to disk on worker nodes.

---

### Question 265: How does Swarm auto-recovery work?

**Answer:**
Swarm reconciliation loop constantly compares **Current State** with **Desired State**.
If a container crashes (Current < Desired), Swarm starts a new one.

---

### Question 266: What is service placement constraint?

**Answer:**
Rules controlling where a service runs.
```bash
docker service create --constraint node.role==worker nginx
docker service create --constraint node.labels.region==us-east nginx
```

---

### Question 267: What’s the purpose of service update configuration?

**Answer:**
Controls how updates roll out.
- `--update-parallelism`: Update 2 containers at once.
- `--update-delay`: Wait 10s between updates.
- `--update-failure-action`: Pause or Rollback if update fails.

---

### Question 268: How do rolling updates work in Swarm?

**Answer:**
Swarm updates tasks in batches.
1.  Stop old task.
2.  Start new task.
3.  Wait for state `Running`.
4.  Proceed to next batch.
*Maintains availability during upgrade.*

---

### Question 269: How do you rollback in Docker Swarm?

**Answer:**
`docker service update --rollback <service_name>`.
Reverts to the `ServiceSpec` defined immediately prior to the last update.

---

### Question 270: How do you scale services up/down in Swarm via CLI?

**Answer:**
```bash
docker service scale my-web=5
```
Or update the service: `docker service update --replicas 5 my-web`.

---

## 🛡️ Container Isolation, Policies & Policies (Questions 271-280)

### Question 271: How does Docker provide process isolation?

**Answer:**
Via **PID Namespaces**.
The container sees itself as PID 1, and cannot see processes on the host or other containers (unless using `--pid=host`).

---

### Question 272: What is namespace sharing between containers?

**Answer:**
Containers can join each other's namespaces (Pod concept).
`docker run --net=container:other --ipc=container:other ...`
Used for sidecars (e.g., a log shipper reading shared volume or localhost network).

---

### Question 273: What is the danger of `--privileged` mode?

**Answer:**
It disables all security features (Namespaces, Cgroups, Seccomp, Caps).
The container has nearly full host access. Root in container = Root on Host.
**Never use in production.**

---

### Question 274: How do AppArmor profiles differ per container?

**Answer:**
You can apply different profiles per container.
- Container A (Web): Deny Write to /etc.
- Container B (Admin): Allow Write /etc.
`docker run --security-opt apparmor=my-strict-profile ...`

---

### Question 275: How do you set immutable filesystem for containers?

**Answer:**
`docker run --read-only`.
Forces the container's root filesystem to be Read-Only.
Any writes must go to mounted Volumes or Tmpfs.
*High security.*

---

### Question 276: How does Docker protect the host kernel?

**Answer:**
It doesn't "protect" the kernel itself (kernel is shared).
It uses **Seccomp** to prevent the container from making dangerous syscalls that could crash or exploit the kernel.

---

### Question 277: Can containers affect host performance? How to mitigate?

**Answer:**
Yes, by causing "Noisy Neighbor" issues (CPU starvation).
**Mitigate:**
- **CPU:** `--cpus`.
- **Memory:** `-m`.
- **I/O:** `--device-read-bps`.

---

### Question 278: How do Linux capabilities impact container permissions?

**Answer:**
They give granular permissions.
Instead of "Root" (All powerful), you give "Net Admin" (Can change IP) or "Chown" (Can change file owner).
Minimizes blast radius if compromised.

---

### Question 279: What’s the role of UID/GID mapping in Docker?

**Answer:**
User Namespaces (`userns-remap`).
Maps `Root (UID 0)` inside container to `User (UID 1000)` outside.
If attacker breaks out as Root, they find themselves as a nobody user on Host.

---

### Question 280: How can you set a read-only root filesystem?

**Answer:**
```bash
docker run --read-only --tmpfs /run --tmpfs /tmp nginx
```
*Note: Most apps need at least /tmp or /run writable.*

---

## 🧰 Tooling, Ecosystem, Plugins (Questions 281-290)

### Question 281: What is Docker Slim?

**Answer:**
A tool that analyzes your running container and generates a "Slim" version of the image, removing unused files/libs. Can reduce size by 30x.

---

### Question 282: How does Dive help in image optimization?

**Answer:**
CLI tool to explore layers.
- Shows what files were added in each layer.
- Highlights "wasted space" (files added then deleted in later layers).

---

### Question 283: What is Hadolint and how do you use it?

**Answer:**
A Dockerfile linter in Haskell.
`hadolint Dockerfile`
Warns about: `RUN cd ...`, `Using latest`, `Missing versions`.

---

### Question 284: What is Docker Scout?

**Answer:**
Docker's new supply chain security product.
Replaces `docker scan`. It analyzes image SBOMs (Software Bill of Materials) and CVEs.

---

### Question 285: How do you use Clair for container vulnerability scanning?

**Answer:**
Clair is an API-driven static analysis tool.
Usually integrated into registries (like Quay). It scans layers against CVE databases.

---

### Question 286: What are Docker plugins and what types are available?

**Answer:**
Plugins extend Docker functionality.
- **Volume Plugins:** (RexRay, Portworx) - Network storage.
- **Network Plugins:** (Weave, Calico) - Overlay networks.
- **Log Plugins:** (Splunk).

---

### Question 287: What are alternatives to Docker Compose?

**Answer:**
- **Podman Compose:** Backend for Podman.
- **Helm:** For Kubernetes (different but analogous role).
- **Skaffold/Tilt:** For Dev loops.

---

### Question 288: How can you use Skopeo with Docker registries?

**Answer:**
Skopeo can inspect, copy, and delete images on remote registries **without pulling them**.
`skopeo inspect docker://docker.io/library/ubuntu`

---

### Question 289: What are some tools for visualizing container topologies?

**Answer:**
- **Portainer:** Management UI.
- **Weave Scope:** Real-time visualization of network connections between containers.
- **Docker Desktop Dashboard.**

---

### Question 290: What’s the difference between Buildah and Docker?

**Answer:**
- **Buildah:** Dedicated tool purely for building OCI images. Daemonless. Great for CI.
- **Docker:** All-in-one tool (Build + Run).

---

## 📦 Registries & Image Distribution (Questions 291-300)

### Question 291: How does Docker pull layers from registries?

**Answer:**
Parallel download.
1.  Downloads Manifest.
2.  Identifies missing layers (by Digest).
3.  Downloads blobs (layers) concurrently.

---

### Question 292: What is a multi-arch Docker image?

**Answer:**
A single tag (`python:3.9`) that generally works on AMD64, ARM64 (M1 Mac), etc.
Supported by **Manifest Lists**.

---

### Question 293: How do you create multi-platform Docker images?

**Answer:**
Use `buildx`.
```bash
docker buildx build --platform linux/amd64,linux/arm64 -t my-app . --push
```

---

### Question 294: How does Docker Hub rate-limiting work?

**Answer:**
- **Anonymous:** 100 pulls / 6 hours.
- **Free Account:** 200 pulls / 6 hours.
- **Pro:** Unlimited.
*Impacts CI pipelines heavily if not authenticated.*

---

### Question 295: What is image replication and why is it useful?

**Answer:**
Syncing images between registries (e.g., Docker Hub -> Private ECR).
- **Latency:** Closer to deploy target.
- **Availability:** Resilience if Hub goes down.

---

### Question 296: How do you enforce private registry access rules?

**Answer:**
- **Authentication:** `docker login`.
- **RBAC:** Registry-side (e.g., AWS IAM for ECR) allows "Read Only" for specific user roles.

---

### Question 297: What is Docker registry garbage collection?

**Answer:**
A process on the registry server to delete blob data that is no longer referenced by any manifest (e.g., after overwriting tags).
Reclaims disk space.

---

### Question 298: How do you secure a private registry with TLS?

**Answer:**
1.  Get Certificate (Let's Encrypt).
2.  Configure Registry container (`registry:2`) with `REGISTRY_HTTP_TLS_CERTIFICATE` environment variables.
3.  Clients rely on CA trust.

---

### Question 299: How do you audit access logs for registries?

**Answer:**
Enable logging in the registry configuration.
Parse access logs (usually Nginx-style format) to see `pull` (GET) and `push` (PUT) actions.

---

### Question 300: How do you mirror a Docker registry for internal use?

**Answer:**
Configure Docker Daemon `daemon.json`:
```json
{
  "registry-mirrors": ["https://my-mirror.internal"]
}
```
Docker will try the mirror first before calling Docker Hub.
