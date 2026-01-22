## 🧪 Docker Testing, CI/CD & Logging (Questions 51-60)

### Question 51: How do you test a Docker image?

**Answer:**
Testing a Docker image ensures it performs as expected before deployment.

**Testing Strategies:**
1.  **Linter:** Use tools like `Hadolint` to check Dockerfile syntax and best practices.
2.  **Container Structure Tests:** Verify file existence, command capability, and file content (e.g., Google's `container-structure-test`).
3.  **Integration Tests:** Run the container and verify endpoints (e.g., using `curl` or `Postman` against the running container).
4.  **Security Scanning:** Scan for vulnerabilities using `Trivy` or `Docker Scan`.

**Example (Structure Test):**
```yaml
schemaVersion: "2.0.0"
fileExistenceTests:
  - name: "File existence"
    path: "/app/main.js"
    shouldExist: true
```

---

### Question 52: How is Docker used in CI/CD pipelines?

**Answer:**
Docker standardizes the environment across the pipeline:

**Workflow:**
1.  **Build:** CI server builds the Docker image from the commit.
2.  **Test:** CI runs tests inside a container created from that image.
3.  **Push:** If tests pass, the image is pushed to a registry (Docker Hub/ECR).
4.  **Deploy:** CD triggers deployment (e.g., updates Kubernetes manifest) to pull the new image.

**Benefit:** Eliminates "works on my machine" issues and ensures consistency between Test and Production.

---

### Question 53: How can you log container activity?

**Answer:**
Docker captures the standard output (`STDOUT`) and standard error (`STDERR`) streams of the container's main process.

**Logging Drivers:**
Docker uses logging drivers to manage where logs go:
- `json-file` (Default): Logs are stored as JSON files on the host.
- `syslog`: Sends logs to a syslog server.
- `fluentd`: Forwards logs to Fluentd.
- `awslogs`: Sends logs to CloudWatch.

**Configuration:**
```bash
docker run --log-driver=syslog nginx
```

---

### Question 54: How do you access logs from a Docker container?

**Answer:**
Use the `docker logs` command.

**Commands:**
```bash
# View all logs
docker logs my-container

# Follow logs strictly (tail)
docker logs -f my-container

# View last 100 lines
docker logs --tail 100 my-container

# Show timestamps
docker logs --timestamps my-container
```

---

### Question 55: What is the difference between `docker logs` and `docker attach`?

**Answer:**

| Feature | `docker logs` | `docker attach` |
| :--- | :--- | :--- |
| **Purpose** | View historical output (stdout/stderr) | Attach terminal to running process |
| **Interaction** | Read-only | Interactive (input/output) |
| **Use Case** | Debugging, auditing | Interacting with shell/REPL |
| **Detaching** | Ctrl+C stops viewing | Ctrl+C might stop the container |

---

### Question 56: How do you debug a running container?

**Answer:**
There are several ways to inspect a running container:

1.  **Logs:** Check `docker logs container_id`.
2.  **Inspect:** Check configuration with `docker inspect container_id`.
3.  **Shell Access:** Enter the container using `exec`.
    ```bash
    docker exec -it my-container /bin/sh
    ```
4.  **Processes:** Check running processes with `docker top my-container`.
5.  **Filesystem:** Use `docker cp` to extract files for analysis.

---

### Question 57: What tools integrate with Docker for CI/CD?

**Answer:**
Almost all modern CI/CD tools integrate natively with Docker.

**Popular Tools:**
- **Jenkins:** Has Docker pipeline create/build/push plugins.
- **GitLab CI:** Uses `.gitlab-ci.yml` to define Docker-based jobs (Docker-in-Docker).
- **GitHub Actions:** Actions can run in containers or build containers.
- **CircleCI:** First-class Docker support.
- **Travis CI:** Built-in Docker service.

---

### Question 58: What are some common Docker CI tools?

**Answer:**
Apart from the Orchestrators (Jenkins, GitLab), specific tools help in the CI process:
1.  **Watchtower:** Automatically updates running containers when new images are pushed.
2.  **Hadolint:** Dockerfile linter.
3.  **Dive:** Tool for exploring a docker image, layer contents, and ways to shrink the size.
4.  **Skopeo:** Work with remote image registries - retrieving information, images, signing content.

---

### Question 59: What is the role of Docker in microservices architecture?

**Answer:**
Docker is the enabler of microservices by providing:
1.  **Isolation:** Each microservice runs in its own container with its own dependencies.
2.  **Polyglot:** Service A can be in Node.js, Service B in Go, running side-by-side.
3.  **Independence:** Services can be upgraded, scaled, and restarted independently.
4.  **Standardization:** All services are deployed as "Containers", simplifying the interface for Ops.

---

### Question 60: How do you manage configuration across environments in Docker?

**Answer:**
Configuration should be decoupled from the image (12-Factor App methodology).

**Strategies:**
1.  **Environment Variables:** Pass config via `-e` or `.env` files.
2.  **Bind Mounts:** Mount config files (e.g., `nginx.conf`) from the host.
3.  **Docker Secrets:** For sensitive data in Swarm.
4.  **Config Maps/Secrets:** If using Kubernetes.
5.  **Configuration Service:** Fetch config from a remote server (e.g., Consul, Spring Cloud Config) at startup.

---

## 📦 Docker Image Management & Registries (Questions 61-70)

### Question 61: What is the difference between a private and public Docker registry?

**Answer:**
- **Public Registry (e.g., Docker Hub, Quay.io):** Images are accessible to anyone on the internet. Used for open-source software and base OS images.
- **Private Registry (e.g., AWS ECR, Self-hosted Registry):** Images are restricted to authorized users. Used for proprietary company applications and sensitive code.

---

### Question 62: How do you push an image to Docker Hub?

**Answer:**
**Steps:**
1.  **Login:**
    ```bash
    docker login
    ```
2.  **Tag Image:** Must follow format `username/repository:tag`.
    ```bash
    docker tag my-app:latest dhruv/my-app:v1
    ```
3.  **Push:**
    ```bash
    docker push dhruv/my-app:v1
    ```

---

### Question 63: How do you pull an image from a private registry?

**Answer:**
First, you must authenticate against that specific registry URL.

**Example (AWS ECR):**
```bash
# 1. Login
aws ecr get-login-password | docker login --username AWS --password-stdin <id>.dkr.ecr.region.amazonaws.com

# 2. Pull
docker pull <id>.dkr.ecr.region.amazonaws.com/my-app:latest
```

---

### Question 64: What is the format of Docker image tags?

**Answer:**
Format: `registry_url/repository_namespace/image_name:tag`

- **Registry URL:** Defaults to `docker.io` if omitted.
- **Namespace:** User or Organization (e.g., `library` for official images).
- **Image:** Name of the app.
- **Tag:** Version identifier (defaults to `latest`).

**Example:** `ghcr.io/facebook/react:next`

---

### Question 65: What is a dangling image?

**Answer:**
A dangling image is an image that is not tagged and implies it is no longer referenced by a valid name. This usually happens when you build a new image with the same name and tag as an existing one; the old one becomes dangling (shows as `<none>:<none>`).

**Cleanup:**
```bash
docker image prune
```

---

### Question 66: How do you clean up unused Docker images?

**Answer:**
Docker provides the `prune` command.

**Commands:**
```bash
# Remove dangling images
docker image prune

# Remove ALL images not used by a running container
docker image prune -a

# Remove strictly specific image
docker rmi <image_id>
```

---

### Question 67: What is an image layer?

**Answer:**
A Docker image is built up from a series of layers. Each instruction in a `Dockerfile` (like `RUN`, `COPY`, `ADD`) creates a new layer.
- Layers are read-only.
- They are stacked on top of each other.
- When a container starts, a thin read-write layer is added on top.

---

### Question 68: How does Docker layer caching work?

**Answer:**
When building an image, Docker checks if a layer with the exact same instruction and parent layer already exists in the cache.
- If match found: Uses cached layer (Fast).
- If no match: Rebuilds that layer and **all subsequent layers**.

**Optimization:** Put frequently changing commands (like `COPY . .`) at the bottom of Dockerfile, and stable commands (like `npm install`) at the top.

---

### Question 69: What is the size of a typical Docker image?

**Answer:**
It varies greatly depending on the base image:
- **Ubuntu/Debian base:** ~100MB - 700MB.
- **Alpine base:** ~5MB (Base) to ~50MB (with app).
- **Distroless:** Very small, contains only runtime.

**Goal:** Keep it as small as possible for faster pull/deploy times.

---

### Question 70: How do you optimize image size?

**Answer:**
1.  **Use correct base image:** Use `alpine` versions (e.g., `python:3.9-alpine`).
2.  **Multi-stage builds:** Compile in one stage, copy binary to a minimal runtime stage.
3.  **Minimize layers:** Chain commands using `&&`.
    ```dockerfile
    RUN apt-get update && apt-get install -y vim && rm -rf /var/lib/apt/lists/*
    ```
4.  **Use `.dockerignore`:** Exclude `node_modules`, `.git`, temporary files.

---

## ☁️ Docker in Production & Orchestration (Questions 71-80)

### Question 71: How do you deploy Docker containers in production?

**Answer:**
In production, you rarely use plain `docker run`. You use Orchestrators or Managed Services:
1.  **Orchestrators:** Kubernetes (K8s), Docker Swarm, Nomad.
2.  **Cloud Services:** AWS ECS/Fargate, Google Cloud Run, Azure Container Instances.
3.  **Process:** CI pipeline builds image -> Pushes to Registry -> Orchestrator pulls and updates the deployment.

---

### Question 72: What is the difference between Docker Swarm and Kubernetes?

**Answer:**

| Feature | Docker Swarm | Kubernetes (K8s) |
| :--- | :--- | :--- |
| **Complexity** | Low (Easy to set up) | High (Steep learning curve) |
| **Scalability** | Good for small/medium | Excellent for massive scale |
| **Features** | Basic orchestration | Full ecosystem (ConfigMaps, PVC, Ingress) |
| **Installation** | Native (Built-in) | External tool needed (kubeadm, minikube) |
| **Market Share** | Declining | Industry Standard |

---

### Question 73: How do you monitor containers in production?

**Answer:**
You need to monitor 4 Golden Signals: Latency, Traffic, Errors, and Saturation (CPU/RAM).

**Methods:**
1.  **Docker Stats:** Real-time CLI stats (basic).
2.  **cAdvisor:** Google tool to expose container metrics.
3.  **Prometheus:** Scrapes metrics from cAdvisor/Docker.
4.  **Grafana:** Visualizes metrics from Prometheus.

---

### Question 74: What tools are used for Docker monitoring?

**Answer:**
- **Prometheus & Grafana:** Open-source standard.
- **Datadog:** Commercial, full-stack observability.
- **New Relic:** Application performance monitoring (APM).
- **Sysdig:** Deep container visibility and security.
- **Portainer:** GUI for managing and monitoring Docker environments.

---

### Question 75: What is container orchestration?

**Answer:**
Container orchestration is the automated management of container lifeca. It handles:
- **Provisioning & Deployment:** Scheduling containers on nodes.
- **Scaling:** Scaling up/down based on load.
- **Networking:** Load balancing and service discovery.
- **Health Monitoring:** Restarting failed containers.
- **Updates:** Rolling updates/rollbacks.

---

### Question 76: How do you perform zero-downtime deployments with Docker?

**Answer:**
Using orchestration strategies:
1.  **Rolling Updates (Swarm/K8s):** Update one container at a time. If health checks pass, proceed to next. If fail, rollback.
2.  **Blue/Green Deployment:** Spin up a new version (Green) alongside old (Blue). Switch traffic load balancer to Green.
3.  **Health Checks:** Essential to ensure traffic isn't sent to a container until it's fully ready.

---

### Question 77: How do you handle configuration management in Docker?

**Answer:**
configuration should be injected at runtime, not baked into the image.
- **Prod vs Dev:** Use specific `.env` files or orchestration config objects.
- **Swarm:** Use `docker config` objects.
- **Kubernetes:** Use `ConfigMaps` and `Secrets`.
- **Entrypoint Scripts:** Use a script to generate config files from ENV vars on startup.

---

### Question 78: What is container health check?

**Answer:**
A health check is a command configured in the Dockerfile or Compose file that keeps checking if the service inside the container is working correctly (e.g., `curl localhost:80`).
- **Liveness:** Is the container running?
- **Readiness:** Is the app ready to accept traffic?

---

### Question 79: How do you implement a health check in Dockerfile?

**Answer:**
Use the `HEALTHCHECK` instruction.

**Example:**
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost/ || exit 1
```
If the command fails 3 times (default retries), the container status becomes `unhealthy`.

---

### Question 80: How does Docker handle service discovery?

**Answer:**
Docker provides built-in DNS server.
- Every container gets an IP.
- Containers on the same user-defined network can reach each other by **container name** or **service name**.
- Docker's embedded DNS server resolves `db` to the internal IP of the database container.

---

## ⚙️ Advanced Docker Interview Questions (Questions 81-90)

### Question 81: What is the lifecycle of a Docker container?

**Answer:**
1.  **Created:** Image is instantiated, resources reserved, but not running.
2.  **Running:** Process is executing.
3.  **Paused:** Processes suspended (SIGSTOP).
4.  **Stopped:** Main process has exited (SIGTERM/SIGKILL).
5.  **Deleted:** Container removed from disk.

---

### Question 82: What is the difference between `exec` and `attach`?

**Answer:**
- **`docker attach`:** Connects your terminal to the **main process** (PID 1) of the container. If you hit Ctrl+C, you might kill the process and stop the container.
- **`docker exec`:** Starts a **new process** (e.g., `/bin/bash`) inside an existing container. Exiting this new process does not affect the main container process. `exec` is safer for debugging.

---

### Question 83: How do you handle stateful applications in Docker?

**Answer:**
Stateful apps (databases like MySQL, Postgres) require data persistence.
1.  **Volumes:** ALWAYS use named volumes or bind mounts for data directories (`/var/lib/mysql`).
2.  **Placement Constraints:** In orchestration, pin specific containers to specific nodes if using local storage.
3.  **StatefulSets:** In Kubernetes, use StatefulSets to manage ordering and stable network IDs.

---

### Question 84: Can you run a database inside Docker? Pros/cons?

**Answer:**
**Yes, but with considerations.**

| Pros | Cons |
| :--- | :--- |
| Easy Setup/Teardown | Data persistence complexity |
| Consistent Environment | Performance tuning (Disk I/O) |
| Isolation | Risk of data loss if not managed right |
| Good for Dev/Test | Backup/Restore complexity vs Managed DB |

*Verdict:* Excellent for non-production. For production, managed services (RDS) are often preferred unless you are an expert in DB ops.

---

### Question 85: What is Init system in containers?

**Answer:**
Containers often run a single process. However, some apps spawn child processes. If the main process (PID 1) doesn't handle signals (like SIGTERM) or reap zombie processes correctly, the container becomes unstable.
- **Solution:** Use a lightweight init system like `tini`.
- **Usage:** Add `--init` flag to `docker run`.
  ```bash
  docker run --init my-image
  ```

---

### Question 86: How do you limit CPU/memory for containers?

**Answer:**
Docker maps to Linux cgroups for this.

**Commands:**
```bash
# Hard memory limit
docker run -m 512m nginx

# Soft memory limit (reservation)
docker run --memory-reservation 256m nginx

# CPU limit (0.5 CPU)
docker run --cpus="0.5" nginx

# CPU shares (relative weight)
docker run --cpu-shares 1024 nginx
```

---

### Question 87: What are container registries alternatives to Docker Hub?

**Answer:**
1.  **Amazon ECR:** Elastic Container Registry.
2.  **Google GCR/Artifact Registry.**
3.  **Azure ACR:** Azure Container Registry.
4.  **Harbor:** Open-source, enterprise-class registry with security scanning.
5.  **JFrog Artifactory:** Universal artifact manager.
6.  **GitLab Container Registry.**

---

### Question 88: How do you rollback to a previous image version?

**Answer:**
- **Manual:** Stop current container `v2`, start container with image `v1`.
- **Swarm:** `docker service update --rollback my-service`.
- **Kubernetes:** `kubectl rollout undo deployment my-deployment`.

*Prerequisite:* You must have the previous image tag available (hence, never use `:latest` in prod).

---

### Question 89: What is overlay network?

**Answer:**
An overlay network creates a distributed network among multiple Docker daemon hosts. It sits on top of the host-specific networks.
- It allows containers running on different Swarm nodes to communicate securely.
- Only the containers attached to the overlay network can talk to each other.

---

### Question 90: What is the difference between `alpine`, `debian`, `ubuntu` images?

**Answer:**
- **Alpine:** Based on Alpine Linux (musl libc, busybox). Extremely small (<5MB). Good for Go/Node apps. Potential compatibility issues with glibc-dependent apps.
- **Debian:** Stable, standard Linux environment. Medium size (~120MB). Good compatibility.
- **Ubuntu:** Based on Debian but with more modern packages. Larger.

*Recommendation:* Start with Alpine; if issues arise, switch to Debian-slim.

---

## 🧠 Situational & Scenario-Based Questions (Questions 91-100)

### Question 91: How would you troubleshoot a container that keeps restarting?

**Answer:**
A "CrashLoopBackOff" scenario.
1.  **Check Logs:** `docker logs <container_id>`. Look for app errors, missing env vars, or syntax errors.
2.  **Inspect Exit Code:** `docker inspect <container_id>`. (e.g., Code 137 = OOM Killed).
3.  **Override Entrypoint:** Run with shell to debug manually.
    ```bash
    docker run -it --entrypoint /bin/sh my-image
    ```
4.  **Check Resources:** Is it hitting memory limits?

---

### Question 92: How would you update a running container to a new image?

**Answer:**
You cannot "update" a container in place. You must replace it.
1.  Pull new image: `docker pull my-app:v2`.
2.  Stop old container: `docker stop my-app`.
3.  Remove old container: `docker rm my-app`.
4.  Start new container: `docker run --name my-app my-app:v2`.
*Tools like Watchtower or Orchestrators automate this.*

---

### Question 93: How would you handle a failed container in a CI/CD pipeline?

**Answer:**
1.  **Fail Fast:** The pipeline should stop immediately.
2.  **Capture Logs:** Ensure the CI tool captures and displays `stdout/stderr`.
3.  **Cleanup:** Use `--rm` in CI to ensure failed containers don't persist.
4.  **Notification:** Send alert (Slack/Email) to the developer with the specific stage failure.

---

### Question 94: How would you ensure high availability of your Docker services?

**Answer:**
1.  **Redundancy:** Run multiple replicas of the service.
2.  **Orchestration:** Use Swarm/K8s to distribute replicas across different physical nodes (Anti-Affinity).
3.  **Health Checks:** Auto-restart unhealthy containers.
4.  **Load Balancing:** Distribute traffic evenly.
5.  **Multi-AZ:** If on cloud, spread nodes across Availability Zones.

---

### Question 95: How would you manage secrets in a distributed Docker environment?

**Answer:**
NEVER use environment variables for sensitive secrets (they are visible in `docker inspect`).
- **Swarm:** Use Docker Secrets (`docker secret create`). Mounted as files in `/run/secrets/`.
- **Kubernetes:** Use K8s Secrets.
- **External:** Use Vault or AWS Secrets Manager and fetch at runtime.

---

### Question 96: How would you troubleshoot network issues between containers?

**Answer:**
1.  **Check Network:** `docker network inspect my-net`.
2.  **Connectivity:** EXEC into Container A and `ping` Container B by name.
    ```bash
    docker exec -it app ping db
    ```
3.  **DNS:** Check `/etc/resolv.conf` inside container.
4.  **Firewall:** Check host firewall (iptables/ufw) interfering with Docker bridge.

---

### Question 97: What steps would you take to clean up Docker disk space?

**Answer:**
1.  **Analyze:** `docker system df` to see usage.
2.  **Prune:** `docker system prune` (remove unused data).
3.  **Deep Clean:** `docker system prune -a --volumes` (Be careful! Removes stopped containers and unused volumes).
4.  **Logs:** Check `/var/lib/docker/containers/` for large log files (configure log rotation).

---

### Question 98: How would you migrate a legacy app to Docker?

**Answer:**
1.  **Containerize Base:** Choose an OS image closest to the legacy server (e.g., Ubuntu 18.04).
2.  **Dependencies:** Identify all installed libs and add to `Dockerfile`.
3.  **Config:** Extract hardcoded configs to ENV variables/files.
4.  **State:** Identify where data is written and map to Volumes.
5.  **Process:** Ensure app runs in foreground (no systemd/background services in container).

---

### Question 99: What challenges have you faced while working with Docker?

**Answer:**
*Common interview answers:*
1.  **Networking:** Debugging communication between containers on different networks.
2.  **Image Size:** Reducing bloated images from 1GB to 100MB.
3.  **Permissions:** Handling UID/GID mismatch between host volumes and container users.
4.  **Zombie Processes:** PID 1 reaping issues (solved with `tini`).
5.  **Data Persistence:** Accidentally losing data by not mounting volumes correctly.

---

### Question 100: How would you explain Docker to a non-technical stakeholder?

**Answer:**
"Imagine shipping houses. Before Docker, each house (application) was custom-built on the site (server), which was slow and prone to errors (missing parts).

With Docker, we build the house in a factory and put it in a standard shipping container. This container has everything the house needs—furniture, plumbing, electricity. We can ship this container anywhere in the world (any server), and when we open it, the house looks and works exactly the same as it did in the factory. It’s faster, safer, and cheaper."
