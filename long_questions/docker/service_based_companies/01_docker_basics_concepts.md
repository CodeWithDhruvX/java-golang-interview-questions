# 🐳 Docker Basics & Core Concepts — Interview Questions (Service-Based Companies)

This document covers foundational Docker concepts commonly tested in service-based company interviews, including TCS, Infosys, Wipro, Capgemini, HCL, and similar. These questions appear in 1–5 years experience rounds.

---

### Q1: What is Docker and how is a container different from a Virtual Machine (VM)?

**Answer:**
**Docker** is an open-source platform that packages an application and all of its dependencies (libraries, runtime, configuration) into a standardised unit called a **container**, ensuring it runs identically regardless of the environment.

**Container vs. VM — core difference:**

| Feature | Container | Virtual Machine |
|---|---|---|
| Virtualizes | OS (user space only) | Entire hardware |
| Kernel sharing | Shares the **host kernel** | Each VM has its **own kernel** |
| Size | Megabytes (usually 50–500 MB) | Gigabytes (usually 5–20 GB) |
| Startup time | Seconds (often < 1s) | Minutes |
| Performance | Near-native | ~10-20% overhead from hypervisor |
| Isolation | Process/namespace isolation | Full hardware isolation |
| Tool | Docker, Podman | VMware, VirtualBox, KVM, Hyper-V |
| Use case | Microservices, CI/CD | Full OS isolation, legacy apps |

**Key insight:** A container **is not a mini-VM**. It's just a set of Linux processes running in isolated namespaces (PID, NET, MNT, UTS) with cgroup-enforced resource limits. The processes share the host OS kernel directly — there is no hypervisor overhead.

**When to use each:**
- Use **containers** for stateless microservices, API servers, batch jobs, CI runners.
- Use **VMs** when you need full OS isolation, different kernels, or regulatory compliance for hard tenant separation.

---

### Q2: What are Docker images, containers, and Docker Hub? How are they related?

**Answer:**

**Docker Image:**
- A **read-only template** containing everything needed to run an application: OS files, runtime, app code, libraries, and configuration.
- Composed of stacked **layers** — each instruction in a Dockerfile adds a new layer.
- Images are immutable. You can never change a running image's layers.
- Identified by a name and tag: `nginx:1.25`, `python:3.12-slim`, `myapp:v2.1`.

**Docker Container:**
- A **running instance of an image**.
- Think of an image as a class definition, and a container as an object (instance) of that class.
- Containers are ephemeral — their writable layer (the changes made during runtime) is lost when the container is removed.
- Multiple containers can run from the same image simultaneously.

**Docker Hub:**
- The default public **container registry** — a repository for storing and distributing Docker images.
- Similar to GitHub, but for Docker images.
- Contains official images (`nginx`, `postgres`, `redis`) and community images.
- You `docker pull` from it and `docker push` to it.

**Relationship:**
```
Dockerfile --[docker build]--> Image --[stored on]--> Docker Hub (Registry)
                                 ↓
                        [docker run]
                                 ↓
                           Container (running instance)
```

---

### Q3: Explain the most commonly used Dockerfile instructions.

**Answer:**

```dockerfile
# Base image to start from
FROM node:20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy files from host build context to container filesystem
COPY package.json package-lock.json ./

# Execute commands during the image BUILD (creates a new layer)
RUN npm ci --only=production

# Copy application source code
COPY . .

# Document which port the container listens on (informational only — doesn't open the port)
EXPOSE 3000

# Set environment variables available at runtime
ENV NODE_ENV=production PORT=3000

# Set default environment variables that can be overridden at build time
ARG BUILD_VERSION=1.0.0

# Add metadata labels
LABEL maintainer="team@company.com" version="1.0"

# Define the default executable when container starts
# ENTRYPOINT (non-overridable) + CMD (overridable default args)
ENTRYPOINT ["node"]
CMD ["server.js"]

# Specify a non-root user to run the container as
USER nonroot
```

**`ENTRYPOINT` vs `CMD` — the key distinction:**
- `CMD` is the **default argument** — easily overridden by appending to `docker run myimage <new-command>`.
- `ENTRYPOINT` is the **executable** — harder to override (requires `--entrypoint` flag).
- Best practice: `ENTRYPOINT ["node"]` + `CMD ["server.js"]`. Override just the script while keeping node as executor.

---

### Q4: What is the difference between `docker run`, `docker start`, `docker stop`, and `docker rm`?

**Answer:**

| Command | What it does |
|---|---|
| `docker run` | Creates a **new** container from an image AND starts it |
| `docker start` | Starts an **existing stopped** container |
| `docker stop` | Gracefully stops a running container (sends SIGTERM, waits 10s, then SIGKILL) |
| `docker kill` | Immediately kills a container (sends SIGKILL, no grace period) |
| `docker rm` | **Removes** a stopped container permanently (data in writable layer is lost) |
| `docker rmi` | Removes an image from local storage |
| `docker pause` | Freezes a container (SIGSTOP — processes frozen, not killed) |
| `docker unpause` | Unfreezes a paused container |

**Common flags for `docker run`:**
```bash
docker run -d              # Detached mode (background)
           -it             # Interactive + pseudo-TTY (for shells)
           --name myapp    # Give the container a name
           -p 8080:3000    # Map host port 8080 to container port 3000
           -e "DB_URL=..."  # Set environment variable
           -v /data:/app/data  # Mount host directory
           --rm            # Auto-remove container when it exits
           --restart=always    # Auto-restart policy
           myimage:latest  # Image name:tag
```

**Useful commands:**
```bash
docker ps           # List running containers
docker ps -a        # List ALL containers (including stopped)
docker logs <id>    # View container log output
docker exec -it <id> sh   # Execute command in running container
docker inspect <id>  # Detailed JSON info about a container
docker top <id>      # Processes running inside container
```

---

### Q5: What are Docker Volumes and when should you use them vs. bind mounts?

**Answer:**
By default, data written inside a container is stored in its **writable layer**, which is **destroyed when the container is removed**. Volumes and bind mounts both solve this by persisting data outside the container lifecycle.

**Docker Volumes:**
- Managed by Docker, stored at `/var/lib/docker/volumes/` on the host.
- Completely decoupled from host filesystem structure.
- Can be shared between containers.
- Persisted across container stop/start/remove.

```bash
# Create a named volume
docker volume create mydata

# Use it
docker run -d -v mydata:/var/lib/postgresql/data postgres:15

# List volumes
docker volume ls

# Inspect
docker volume inspect mydata

# Remove
docker volume rm mydata
```

**Bind Mounts:**
- Map a **specific host directory** into the container.
- The host path must exist.
- Docker has no management over it.

```bash
# Mount current host directory ./src to /app in container
docker run -d -v $(pwd)/src:/app/src myapp:dev
```

**When to use which:**

| Scenario | Recommendation |
|---|---|
| Database data persistence | **Named Volume** |
| Development: hot-reloading source code | **Bind Mount** |
| Sharing config files with container | **Bind Mount** |
| Backup/restore of container data | **Named Volume** (easier to manage with `docker volume`) |
| Production databases | **Named Volume** |
| CI/CD artifact output | **Bind Mount** |

---

### Q6: What is the `docker-compose` / Docker Compose tool and when is it used?

**Answer:**
**Docker Compose** is a tool for defining and running **multi-container Docker applications** using a single YAML configuration file (`docker-compose.yml` or `compose.yml`).

Instead of running multiple `docker run` commands manually, you define all your services, networks, volumes, and dependencies in one file and start everything with a single command.

**Example `compose.yml`:**
```yaml
version: "3.9"

services:
  web:
    build: .
    ports:
      - "8080:3000"
    environment:
      - DATABASE_URL=postgres://user:pass@db:5432/mydb
    depends_on:
      db:
        condition: service_healthy

  db:
    image: postgres:15
    volumes:
      - pgdata:/var/lib/postgresql/data
    environment:
      - POSTGRES_PASSWORD=pass
      - POSTGRES_USER=user
      - POSTGRES_DB=mydb
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U user"]
      interval: 10s
      timeout: 5s
      retries: 5

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  pgdata:
```

**Key commands:**
```bash
docker compose up -d          # Start all services in background
docker compose down           # Stop and remove containers, networks
docker compose down -v        # Also remove volumes
docker compose logs -f web    # Follow logs for the web service
docker compose ps             # Status of all services
docker compose exec web sh    # Shell into the web container
docker compose build          # Rebuild images
docker compose restart web    # Restart specific service
```

**Primary use case:** Local development environment — spin up the full stack (app + DB + cache + message queue) with one command.

*Note: Docker Compose is for single-host orchestration. For multi-host production deployments, use Kubernetes.*

---

*Prepared for technical screening rounds at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant).*
