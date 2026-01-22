## 🧩 Docker Compose (Advanced Usage) (Questions 151-160)

### Question 151: How do you override Docker Compose configs for staging vs prod?

**Answer:**
Docker Compose allows merging multiple YAML files.
**Strategy:**
1.  **Base:** `docker-compose.yml` (Common configs).
2.  **Overrides:** `docker-compose.prod.yml` (Production limits, restarts).
3.  **Command:**
    ```bash
    docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
    ```

---

### Question 152: What is the purpose of `.env` in Docker Compose?

**Answer:**
The `.env` file is automatically read by Docker Compose to set **variables inside the YAML file itself** (variable substitution).
**Note:** It does NOT automatically pass environment variables to the container (unless referenced in the `environment` section or `env_file`).

**Example:**
```yaml
# docker-compose.yml
services:
  web:
    image: my-app:${TAG} # TAG comes from .env
```

---

### Question 153: How do you handle environment-specific containers in Compose?

**Answer:**
Use the `-f` flag to include specific files.
- **Dev:** `docker-compose.yml` + `docker-compose.override.yml` (includes Redis/DB locally).
- **Prod:** `docker-compose.yml` + `docker-compose.prod.yml` (omits DB if using RDS).

---

### Question 154: How do you connect external containers to Docker Compose network?

**Answer:**
Define an **external network** in your compose file.

```yaml
networks:
  default:
    external:
      name: my-pre-existing-network
```
Now, services in this compose stack will attach to `my-pre-existing-network` and can talk to other containers on it.

---

### Question 155: Can you define multiple Docker Compose files?

**Answer:**
Yes. You can chain them.
`docker-compose -f base.yml -f extension.yml config` shows the merged result.
Merging rules:
- Single value keys (image) are overridden.
- Multi-value keys (ports, volumes) are concatenated/merged.

---

### Question 156: How do you use `build` vs `image` in Compose?

**Answer:**
- **`image: redis:alpine`**: Pulls from registry.
- **`build: .`**: Builds from Dockerfile at runtime.
- **Both:**
  ```yaml
  image: my-registry/my-app:v1
  build: .
  ```
  Compose builds the image using `build` config, but tags it with the name provided in `image`.

---

### Question 157: How do you restart services automatically in Compose?

**Answer:**
Use the `restart` policy.
- `no`: Never restart (default).
- `always`: Always restart.
- `on-failure`: Restart only if exit code != 0.
- `unless-stopped`: Restart unless manually stopped.

```yaml
services:
  web:
    restart: always
```

---

### Question 158: How do you define health checks in Compose?

**Answer:**
Allows defining the check in YAML instead of Dockerfile.

```yaml
services:
  web:
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost"]
      interval: 30s
      timeout: 10s
      retries: 3
```

---

### Question 159: What are Compose profiles and how are they used?

**Answer:**
Profiles allow starting only specific services conditionally.
**Example:**
```yaml
services:
  frontend:
    profiles: ["ui"]
  backend:
    # default profile
```
**Run:**
`docker-compose --profile ui up` (Starts both).
`docker-compose up` (Starts only backend).

---

### Question 160: How do you upgrade a running Compose stack with zero downtime?

**Answer:**
Docker Compose primarily recreates containers, which causes brief downtime.
**Mitigation:**
1.  **Scale Up:** `docker-compose up -d --scale web=2 --no-recreate`
2.  **External Load Balancer:** Point to new instance.
3.  **Remove Old:** Scale down.
**Better:** Use Swarm (`docker stack deploy`) which handles rolling updates natively.

---

## 🛠️ Docker BuildKit & Build Process (Questions 161-170)

### Question 161: What is Docker BuildKit and why is it better?

**Answer:**
BuildKit is the modern builder toolkit for Docker.
- **Parallelism:** Builds independent layers simultaneously.
- **Secrets:** Secure secret handling.
- **Cache:** Better caching mechanisms.
- **Performance:** Much faster builds.

---

### Question 162: How do you enable BuildKit in Docker CLI?

**Answer:**
Set the environment variable:
```bash
export DOCKER_BUILDKIT=1
docker build .
```
*(Enabled by default in recent Docker versions).*

---

### Question 163: How does BuildKit handle secret mounts?

**Answer:**
It allows mounting secrets during build **without** persisting them in the image layer.

```dockerfile
# Dockerfile
RUN --mount=type=secret,id=mysecret cat /run/secrets/mysecret
```
```bash
# Build command
docker build --secret id=mysecret,src=token.txt .
```

---

### Question 164: How do you run parallel builds with BuildKit?

**Answer:**
BuildKit automatically constructs a dependency graph. If stage B does not depend on Stage A, they run in parallel.
Multi-stage builds benefit significantly from this.

---

### Question 165: What is the `RUN --mount=type=cache` used for?

**Answer:**
It creates a cache directory that persists **between builds** but is not part of the final image.
**Use Case:** Caching compiler dependencies (Maven `.m2`, Go modules).

```dockerfile
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -v ./...
```

---

### Question 166: How do you debug Docker builds effectively?

**Answer:**
1.  **Progress=plain:** `docker build --progress=plain .` to see full logs.
2.  **Stop at stage:** `docker build --target builder .`
3.  **Echo Debug:** Add `RUN ls -la /path` to see file structure.

---

### Question 167: What are `ONBUILD` instructions used for?

**Answer:**
`ONBUILD` executes a command when the image is used as a base for **another** build.
**Example:**
A generic Node.js image might have `ONBUILD COPY package.json .` so that child images automatically copy it.

---

### Question 168: How do you pass build arguments securely?

**Answer:**
- **Standard:** `ARG` (Visible in history, not secure).
- **Secure:** BuildKit Secrets (`--secret`).

---

### Question 169: What is a base image vs parent image?

**Answer:**
- **Base Image:** Has `FROM scratch` (starts from nothing).
- **Parent Image:** The image your Dockerfile references in `FROM` (e.g., `FROM ubuntu`).
*In practice, terms are often used interchangeably to mean "the image I am starting from".*

---

### Question 170: How do you prevent sensitive data leakage in builds?

**Answer:**
1.  **Use .dockerignore.**
2.  **Don't use COPY . .** carelessly (might copy `.env`).
3.  **Use BuildKit Secrets** for tokens.
4.  **Squash Layers:** (Historic method) to merge layers and hide deletion of secrets (Still risky).

---

## 🌐 Docker with Cloud Providers & Infrastructure (Questions 171-180)

### Question 171: How do you run Docker containers in AWS ECS?

**Answer:**
ECS (Elastic Container Service) is AWS's orchestrator.
1.  **Task Definition:** JSON file (like Compose) defining image, CPU, RAM.
2.  **Service:** ensures N copies of the task are running.
3.  **Launch Type:**
    - **EC2:** Manage your own VMs.
    - **Fargate:** Serverless (AWS manages infrastructure).

---

### Question 172: How does Docker integrate with AWS Fargate?

**Answer:**
Fargate provides "Serverless Containers". You just define the container requirements (CPU/RAM), and AWS allocates the compute. No need to patch OS or manage EC2 instances.

---

### Question 173: How do you deploy Docker containers to Azure App Service?

**Answer:**
Azure App Service can pull directly from Docker Hub or ACR (Azure Container Registry).
- Supports Webhooks (Continuous Deployment).
- Handles SSL/Scaling automatically.

---

### Question 174: What are container instances in GCP (Cloud Run)?

**Answer:**
Google Cloud Run:
- Fully managed serverless platform.
- Takes a Docker container -> Generates HTTPS endpoint.
- Scales to zero when not used (Pay per use).

---

### Question 175: How does Docker integrate with Terraform?

**Answer:**
Terraform uses the `docker` provider.
- Can spin up containers on a host with Docker API exposed.
- More commonly, Terraform provisions the **infrastructure** (EKS, ECS, GKE) where Docker containers eventually run.

---

### Question 176: What is the difference between Docker and Podman?

**Answer:**
| Feature | Docker | Podman |
| :--- | :--- | :--- |
| **Architecture** | Daemon-based (`dockerd`) | Daemonless (Direct fork/exec) |
| **Root** | Usually requires root (daemon) | Rootless by design |
| **Compatibility** | Standard CLI | CLI compatible (`alias docker=podman`) |
| **K8s** | - | Can generate K8s YAML (`podman generate kube`) |

---

### Question 177: What is the role of Docker in Kubernetes?

**Answer:**
Historically, K8s used Docker as the Runtime.
Since K8s v1.20+, Docker is deprecated as the *runtime*, replaced by **containerd** or **CRI-O**.
*You still build images with Docker, but K8s runs them using containerd.*

---

### Question 178: How do you create a private image registry in the cloud?

**Answer:**
Use the cloud-native providers:
- **AWS:** Amazon ECR (`aws ecr create-repository`).
- **GCP:** Artifact Registry.
- **Azure:** ACR.
- **Docker Hub:** Private repository plan.

---

### Question 179: What is Docker Desktop vs Docker Engine?

**Answer:**
- **Docker Engine:** The Linux-native core (CLI + Daemon).
- **Docker Desktop:** GUI application for Windows/Mac. It runs a Linux VM (WSL2 or HyperKit) to host the Docker Engine. Includes K8s, GUI Dashboard, and Extensions.

---

### Question 180: How do you provision infrastructure with Docker Machine?

**Answer:**
*Deprecated Tool.*
Docker Machine allowed you to install Docker Engine on virtual hosts (VirtualBox, AWS EC2) from your local computer.
**Replacement:** Use Terraform, Ansible, or Cloud-init scripts.

---

## 📦 Container Storage & Volumes (Advanced) (Questions 181-190)

### Question 181: What is the difference between `tmpfs` and volume mounts?

**Answer:**
- **Volume/Bind:** Writes to Host Disk (Persistent).
- **Tmpfs:** Writes to Host Memory (RAM).
  - **Fastest I/O.**
  - **Secure:** Data vanishes on stop (good for secrets).
  - **Volatile:** Not persistent.

---

### Question 182: How do you persist database state in Docker?

**Answer:**
Always mount a **Volume** to the database's data directory.
**MySQL:** Mount to `/var/lib/mysql`.
**Postgres:** Mount to `/var/lib/postgresql/data`.
If you delete the container, the Volume remains.

---

### Question 183: How do you backup and restore Docker volumes?

**Answer:**
**Backup:** Run a temporary container that mounts the volume AND a local backup directory, then tar the contents.
```bash
docker run --rm -v my-vol:/data -v $(pwd):/backup ubuntu tar cvf /backup/backup.tar /data
```

---

### Question 184: How do you use NFS with Docker volumes?

**Answer:**
You can create a volume using the `local` driver with NFS options.
```bash
docker volume create --driver local \
  --opt type=nfs \
  --opt o=addr=192.168.1.1,rw \
  --opt device=:/path/to/share \
  my-nfs-vol
```

---

### Question 185: What are CSI (Container Storage Interface) drivers?

**Answer:**
CSI is a standard for exposing arbitrary block and file storage systems to container orchestration systems (like K8s). Docker uses its own plugin system, but K8s uses CSI to connect to AWS EBS, Azure Disk, etc.

---

### Question 186: How do you share volumes across containers?

**Answer:**
1.  **Multiple mounts:**
    ```bash
    docker run -v shared_data:/data app1
    docker run -v shared_data:/data app2
    ```
2.  **`--volumes-from` (Old way):**
    ```bash
    docker run --volumes-from container1 app2
    ```

---

### Question 187: How do you inspect data inside a volume?

**Answer:**
1.  **Mount it:** Spin up a dummy container.
    `docker run --rm -it -v my-vol:/data alpine ls -la /data`
2.  **Inspect host path:** `docker volume inspect my-vol` (Finds mount point on host, usually `/var/lib/docker/volumes/...`).

---

### Question 188: How does Docker handle read-only volumes?

**Answer:**
Add `:ro` to the mount command. The container can read but not modify files.
```bash
docker run -v my-config:/etc/config:ro my-app
```

---

### Question 189: What’s the impact of file permission issues in volume mounts?

**Answer:**
If the Host User ID (UID) differs from the Container User ID, the container might get "Permission Denied".
**Fix:**
- Ensure UIDs match.
- Or use `chmod` in entrypoint script.
- Or use `user: "1000:1000"` in Compose.

---

### Question 190: How do you manage volume lifecycles?

**Answer:**
- Volumes persist until explicitly deleted.
- **Anonymous volumes:** Created if not named. Can accumulate trash.
- **Cleanup:** `docker volume prune` removes unused volumes.

---

## 📌 Miscellaneous / Real-World Scenarios (Questions 191-200)

### Question 191: How do you identify which containers are consuming the most CPU?

**Answer:**
`docker stats` gives a real-time table of CPU/Memory usage for running containers.
For valid sorting, you might need external monitoring (Grafana).

---

### Question 192: What do you do when a container uses 100% disk space?

**Answer:**
1.  **Check:** Is it the container filesystem or a volume?
2.  **Logs:** Often JSON logs fill the disk (`/var/lib/docker/containers`). Truncate them.
3.  **Tmp:** Check for temp files inside the writable layer using `docker diff`.
4.  **Prune:** Clean up unused images/builder cache.

---

### Question 193: How do you roll out updates to a fleet of containers?

**Answer:**
**Orchestration (Swarm/K8s):**
`docker service update --image new-image --update-parallelism 2 my-service`
This stops 2 old containers, starts 2 new ones, waits for health, repeats.

---

### Question 194: How do you run scheduled jobs in Docker (cron)?

**Answer:**
1.  **Host Cron:** Trigger `docker run ...` from host crontab.
2.  **Internal Cron:** Run `crond` inside the container alongside the app (requires Supervisor).
3.  **K8s CronJob:** Native resource in Kubernetes.
4.  **Ofelia:** A Docker based job scheduler.

---

### Question 195: How do you detect and avoid container zombie processes?

**Answer:**
- **Symptom:** `docker stop` hangs (process ignores SIGTERM).
- **Cause:** PID 1 doesn't handle signals/reap children.
- **Fix:** Use `--init` flag to wrap with `tini` (handles signal forwarding/reaping).

---

### Question 196: How do you perform blue-green deployments with Docker?

**Answer:**
Manual approach with Nginx (Load Balancer):
1.  Blue is running on port 8001.
2.  Start Green on port 8002.
3.  Update Nginx config to point upstream to 8002.
4.  Reload Nginx.
5.  Stop Blue.

---

### Question 197: How do you restart failed containers automatically?

**Answer:**
Use Restart Policies:
`docker run --restart on-failure:5 my-app`
(Retries 5 times if exit code is error).

---

### Question 198: How do you isolate builds between microservices using Docker?

**Answer:**
1.  **Separate Dockerfiles.**
2.  **Multi-stage:** Build dependencies in isolation.
3.  **CI Jobs:** Run separate CI jobs for each service.
4.  **Docker Compose:** `docker-compose build service-a` only builds A.

---

### Question 199: How do you implement canary releases using Docker?

**Answer:**
Requires a sophisticated Load Balancer (like Traefik, Istio, or Nginx).
1.  Deploy v2 container.
2.  Configure LB to send 5% of traffic to v2, 95% to v1.
3.  Monitor errors.
4.  If good, increase to 100%.

---

### Question 200: How do you document Docker workflows for your team?

**Answer:**
1.  **README:** Explain `docker-compose up` usage.
2.  **Makefile:** Wrap complex docker commands (`make dev`, `make test`).
3.  **Comments:** Comment the Dockerfile explaining weird environment setups.
4.  **Architecture Diagram:** Show network/volume interactions.

---
