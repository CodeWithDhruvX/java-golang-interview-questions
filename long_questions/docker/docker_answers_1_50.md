## 🐳 Basic Docker Interview Questions (Questions 1-20)

### Question 1: What is Docker?

**Answer:**
Docker is an open-source platform that automates the deployment of applications inside lightweight, portable, and self-sufficient containers.

**Key Concepts:**
- **Containerization:** Packages code and dependencies together.
- **Portability:** Runs consistently across different environments (Dev, Test, Prod).
- **Efficiency:** Uses OS-level virtualization to share the host kernel, making it lighter than VMs.

**Example:**
```bash
# Verify Docker installation
docker --version

# Run a simple container
docker run hello-world
```

---

### Question 2: What are the benefits of using Docker?

**Answer:**
Docker offers several advantages for modern software development:

**Benefits:**
1.  **Consistency:** "It works on my machine" is no longer an issue.
2.  **Isolation:** Applications run in isolated environments without conflicting dependencies.
3.  **Portability:** Containers run anywhere Docker is installed.
4.  **Microservices:** Ideal for breaking monolithic apps into smaller services.
5.  **CI/CD:** Simplifies automated testing and deployment pipelines.
6.  **Resource Efficiency:** High density of containers on a single host.

---

### Question 3: What is the difference between Docker and a virtual machine?

**Answer:**
The main difference lies in architecture and resource usage:

| Feature | Docker Container | Virtual Machine (VM) |
| :--- | :--- | :--- |
| **Architecture** | Shares Host OS Kernel | Has its own Guest OS |
| **Size** | Megabytes (MB) | Gigabytes (GB) |
| **Boot Time** | Seconds | Minutes |
| **Isolation** | Process-level isolation | OS-level isolation |
| **Performance** | Native performance | Overhead due to hypervisor |

**Visual Representation:**
```text
[ App A ] [ App B ]     [ App A ] [ App B ]
[ Bin/Libs ] [ Bin/Libs ]     [ Bin/Libs ] [ Bin/Libs ]
[ Docker Engine ]       [ Guest OS ] [ Guest OS ]
[ Host OS ]             [ Hypervisor ]
[ Server Infrastructure ]     [ Host OS / Server ]
```

---

### Question 4: What is a Docker container?

**Answer:**
A Docker container is a runnable instance of a Docker image. It is a lightweight, standalone, and executable package of software that includes everything needed to run an application: code, runtime, system tools, system libraries, and settings.

**Lifecycle:**
1.  **Create:** Created from an image.
2.  **Start:** Processes begin execution.
3.  **Stop:** Processes are halted.
4.  **Remove:** Container is deleted.

**Command:**
```bash
# Create and start a container
docker run -d --name my-web-server nginx
```

---

### Question 5: What is a Docker image?

**Answer:**
A Docker image is a read-only template used to create containers. It consists of multiple layers, where each layer represents an instruction in the image's Dockerfile.

**Characteristics:**
- **Immutable:** Cannot be changed once built.
- **Layered:** Uses a Union File System.
- **Base:** Often starts from a base OS (e.g., Ubuntu, Alpine).

**Command:**
```bash
# List available images
docker images

# Pull an image from registry
docker pull python:3.9-alpine
```

---

### Question 6: What is Docker Hub?

**Answer:**
Docker Hub is a cloud-based registry service provided by Docker for finding and sharing container images.

**Features:**
- **Repositories:** Store and manage Docker images.
- **Official Images:** Curated images by Docker and upstream vendors (e.g., Node, Mongo).
- **Webhooks:** Trigger actions after a successful push.
- **Organizations:** Manage team access.

**Usage:**
```bash
# Login to Docker Hub
docker login

# Tag an image
docker tag my-app:v1 username/my-app:v1

# Push to Docker Hub
docker push username/my-app:v1
```

---

### Question 7: What is a Dockerfile?

**Answer:**
A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image.

**Example `Dockerfile`:**
```dockerfile
# Base image
FROM node:18-alpine

# Working directory
WORKDIR /app

# Copy dependencies
COPY package.json .

# Install dependencies
RUN npm install

# Copy source code
COPY . .

# Expose port
EXPOSE 3000

# Startup command
CMD ["npm", "start"]
```

---

### Question 8: How do you build a Docker image?

**Answer:**
You use the `docker build` command pointing to the location of the Dockerfile.

**Syntax:**
```bash
docker build -t <image_name>:<tag> <path_to_context>
```

**Example:**
```bash
# Build image from current directory
docker build -t my-node-app:1.0 .

# Build with no cache (force rebuild)
docker build --no-cache -t my-node-app:1.0 .
```

---

### Question 9: How do you run a Docker container?

**Answer:**
Use the `docker run` command to create and start a container from an image.

**Common Flags:**
- `-d`: Run in detached mode (background).
- `-p`: Map ports (Host:Container).
- `--name`: Assign a name.
- `-v`: Mount volumes.
- `-e`: Set environment variables.

**Example:**
```bash
docker run -d \
  --name my-app \
  -p 8080:80 \
  -e ENV=production \
  nginx
```

---

### Question 10: How do you stop a running container?

**Answer:**
You can stop a container using `docker stop` (graceful shutdown) or `docker kill` (immediate termination).

**Commands:**
```bash
# Stop a container (sends SIGTERM)
docker stop my-app

# Stop with a timeout (default 10s)
docker stop -t 30 my-app

# Kill a container (sends SIGKILL)
docker kill my-app
```

---

### Question 11: What is the difference between `CMD` and `ENTRYPOINT`?

**Answer:**
Both instructions define what command gets executed when a container starts, but they behave differently.

| Feature | `CMD` | `ENTRYPOINT` |
| :--- | :--- | :--- |
| **Purpose** | Default arguments/command | Executable to run |
| **Overridable** | Easily overridden by CLI args | Arguments append to it |
| **Common Use** | Default parameters | Main application binary |

**Example:**
```dockerfile
# Dockerfile
ENTRYPOINT ["python"]
CMD ["app.py"]
```
- `docker run my-image` executes: `python app.py`
- `docker run my-image script.py` executes: `python script.py` (CMD overridden)

---

### Question 12: How do you list all Docker containers?

**Answer:**
Use the `docker ps` command.

**Commands:**
```bash
# List running containers
docker ps

# List ALL containers (running and stopped)
docker ps -a

# Show only container IDs
docker ps -q
```

---

### Question 13: How do you list all Docker images?

**Answer:**
Use `docker images` or `docker image ls`.

**Commands:**
```bash
# List top-level images
docker images

# List all images including intermediate layers
docker images -a

# Filter images (e.g., dangling)
docker images -f "dangling=true"
```

---

### Question 14: How do you delete a Docker container?

**Answer:**
Use `docker rm`. The container must be stopped first unless you use force.

**Commands:**
```bash
# Delete a stopped container
docker rm my-container

# Force delete a running container
docker rm -f my-container

# Remove all stopped containers
docker container prune
```

---

### Question 15: How do you delete a Docker image?

**Answer:**
Use `docker rmi`.

**Commands:**
```bash
# Delete an image
docker rmi my-image:latest

# Force remove (if used by a stopped container)
docker rmi -f my-image:latest

# Remove all unused images
docker image prune
```

---

### Question 16: What is the purpose of `.dockerignore` file?

**Answer:**
Similar to `.gitignore`, it specifies files and directories in the build context that should be excluded from the Docker build.

**Benefits:**
- Reduces image size.
- Speeds up build process (less data validation).
- Security (prevents secrets/keys from copying).

**Example `.dockerignore`:**
```text
node_modules
.git
Dockerfile
.env
```

---

### Question 17: How can you expose a port from a Docker container?

**Answer:**
There are two parts:
1.  `EXPOSE` in Dockerfile (Documentation/Metadata).
2.  `-p` or `-P` flag in `docker run` (Actual mapping).

**Usage:**
```bash
# Dockerfile
EXPOSE 80

# Command Line
# Map Host 8080 to Container 80
docker run -p 8080:80 nginx

# Map random high port to exposed port
docker run -P nginx
```

---

### Question 18: What is the default Docker network?

**Answer:**
The default network is **Bridge**.

**Network Types:**
- **Bridge:** Default. Containers can talk to each other if on same bridge.
- **Host:** Container shares host's network stack.
- **None:** No networking.
- **Overlay:** Used for multi-host (Swarm).

**Check Network:**
```bash
docker network ls
docker network inspect bridge
```

---

### Question 19: What is a volume in Docker?

**Answer:**
A volume is the preferred mechanism for persisting data generated by and used by Docker containers.

**Characteristics:**
- Stored outside the container’s filesystem (usually in `/var/lib/docker/volumes`).
- Managed by Docker.
- Can be shared between containers.
- Survives container deletion.

**Creation:**
```bash
docker volume create my-vol
```

---

### Question 20: How do you persist data in Docker?

**Answer:**
Data inside a container is ephemeral. To persist it, use:

1.  **Volumes:** Managed by Docker (Best for DBs).
2.  **Bind Mounts:** Maps a host file/directory to container.

**Example:**
```bash
# Using Volume
docker run -v my-data:/var/lib/mysql mysql

# Using Bind Mount
docker run -v /home/user/project:/app nginx
```

---

## 🛠️ Intermediate Docker Interview Questions (Questions 21-40)

### Question 21: What is the difference between a bind mount and a volume?

**Answer:**

| Feature | Volume | Bind Mount |
| :--- | :--- | :--- |
| **Location** | Managed by Docker (`/var/lib/docker/volumes`) | Anywhere on Host Filesystem |
| **Portability** | High (Host OS agnostic) | Low (Dependent on Host paths) |
| **Management** | `docker volume` commands | Manual OS commands |
| **Use Case** | Databases, Persistent storage | Dev code sync, Config files |

**Example:**
```bash
# Volume
docker run -v my_vol:/app/data my-image

# Bind Mount
docker run -v $(pwd)/data:/app/data my-image
```

---

### Question 22: How do you copy files into a Docker container?

**Answer:**
Use the `docker cp` command. It works both ways (Host ↔ Container).

**Commands:**
```bash
# Host to Container
docker cp ./local-file.txt my-container:/app/

# Container to Host
docker cp my-container:/app/log.txt ./local-logs/
```
*Note: This can be done on running or stopped containers.*

---

### Question 23: What is the purpose of multi-stage builds?

**Answer:**
Multi-stage builds allow you to use multiple `FROM` instructions in a single Dockerfile. They are primarily used to reduce image size by separating the build environment from the runtime environment.

**Example:**
```dockerfile
# Stage 1: Build
FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

# Stage 2: Runtime
FROM alpine:3.14
WORKDIR /root/
# Copy only the binary from builder
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```

---

### Question 24: How do you pass environment variables to a container?

**Answer:**
You can pass them using the `-e` flag or an env file.

**Methods:**
```bash
# Inline
docker run -e DB_HOST=localhost -e DB_PORT=5432 my-app

# Using a file (.env)
docker run --env-file .env my-app
```
*Inside Dockerfile, use `ENV` to set default values.*

---

### Question 25: What is the use of `--rm` flag?

**Answer:**
The `--rm` flag automatically creates a cleanup rule. It removes the container (and its associated anonymous volumes) immediately after it exits.

**Use Case:**
- One-off tasks or scripts.
- Testing images.
- Preventing accumulation of stopped containers.

**Example:**
```bash
docker run --rm busybox echo "Hello World"
# Container is deleted after echo
```

---

### Question 26: What are Docker namespaces?

**Answer:**
Namespaces are a Linux kernel feature that Docker uses to provide isolation. Each aspect of a container runs in a separate namespace and its access is limited to that namespace.

**Types:**
- **PID:** Process isolation.
- **NET:** Network stacks.
- **MNT:** Mount points.
- **IPC:** Inter-Process Communication.
- **UTS:** Hostname and domain name.
- **USER:** User ID mapping.

---

### Question 27: What are Docker cgroups?

**Answer:**
Control Groups (cgroups) are a Linux kernel feature that limits, accounts for, and isolates the usage of physical resources (CPU, Memory, Disk I/O, Network) of a collection of processes.

**Function:**
- **Resource Limiting:** Limit memory to 512MB.
- **Prioritization:** Give more CPU shares to specific containers.
- **Accounting:** Measure resource usage.

**Example:**
```bash
docker run --cpus=".5" --memory="512m" nginx
```

---

### Question 28: What is Docker Compose?

**Answer:**
Docker Compose is a tool for defining and running multi-container Docker applications. It uses a YAML file (`docker-compose.yml`) to configure the application's services.

**Workflow:**
1.  Define environment in Dockerfile.
2.  Define services in `docker-compose.yml`.
3.  Run `docker-compose up`.

---

### Question 29: How do you define services in Docker Compose?

**Answer:**
Services are defined under the `services` key in the YAML file.

**Example `docker-compose.yml`:**
```yaml
version: '3.8'
services:
  web:
    build: .
    ports:
      - "5000:5000"
  redis:
    image: "redis:alpine"
```

---

### Question 30: How do you scale services in Docker Compose?

**Answer:**
Use the `--scale` flag with `docker-compose up`.

**Command:**
```bash
# Start 3 instances of the 'web' service
docker-compose up -d --scale web=3
```
*Note: Requires the usage of a load balancer or removing host port binding to avoid conflict.*

---

### Question 31: What is the `depends_on` in Docker Compose?

**Answer:**
`depends_on` expresses start and shutdown options between services. It ensures that a specific service starts before another.

**Example:**
```yaml
services:
  web:
    depends_on:
      - db
  db:
    image: postgres
```
*Note: It only waits for the container to start, not for the application (e.g., DB ready) to be ready. Use `healthcheck` for that.*

---

### Question 32: What is the difference between `COPY` and `ADD` in Dockerfile?

**Answer:**
Both copy files from host to container, but `ADD` has extra features.

| Feature | COPY | ADD |
| :--- | :--- | :--- |
| **Basic Copy** | Yes (Local files) | Yes (Local files) |
| **URL Support** | No | Yes (Download from URL) |
| **Extraction** | No | Yes (Extract tar/zip auto) |
| **Best Practice** | Preferred | Use only when needed |

**Reference:**
```dockerfile
COPY . .                  # Recommended
ADD https://example.com/file.tar.gz .  # Use ADD for this
```

---

### Question 33: How do you connect containers using Docker networks?

**Answer:**
You create a user-defined bridge network and attach containers to it. Containers on the same network can resolve each other by container name.

**Steps:**
```bash
# 1. Create network
docker network create my-net

# 2. Run containers on network
docker run -d --name db --network my-net redis
docker run -d --name app --network my-net my-node-app

# 3. Application connects to 'db:6379'
```

---

### Question 34: What are the different types of Docker networks?

**Answer:**
1.  **Bridge:** Default, for containers on same host.
2.  **Host:** Removes network isolation, binds directly to host.
3.  **Overlay:** For multi-host communication (Swarm/K8s).
4.  **Macvlan:** Assigns a MAC address to container, making it look like a physical device.
5.  **None:** Disabled networking.

---

### Question 35: What is Docker Swarm?

**Answer:**
Docker Swarm is Docker's native container orchestration tool. It turns a pool of Docker hosts into a single, virtual Docker host.

**Features:**
- Decentralized design.
- Declarative service model.
- Scaling.
- Rolling updates.
- Load balancing.

---

### Question 36: How do you create a Docker Swarm cluster?

**Answer:**
1.  **Initialize Manager:**
    ```bash
    docker swarm init --advertise-addr <MANAGER_IP>
    ```
2.  **Join Workers:**
    Run the command outputted by the init command on worker nodes:
    ```bash
    docker swarm join --token <TOKEN> <MANAGER_IP>:2377
    ```

---

### Question 37: What is the difference between a service and a container in Docker Swarm?

**Answer:**
- **Container:** An isolated process (atomic unit).
- **Service:** High-level abstraction in Swarm. It defines the desired state (image, replicas, ports) for a group of containers (tasks). Swarm ensures the actual state matches the desired state.

**Command:**
```bash
# Create a service (5 replicas)
docker service create --replicas 5 --name my-web nginx
```

---

### Question 38: What is a manager node in Docker Swarm?

**Answer:**
Manager nodes handle cluster management tasks:
- Maintaining cluster state.
- Scheduling services.
- Serving Swarm API endpoints.

*There should be an odd number of managers for quorum (Raft consensus).*

---

### Question 39: What is a worker node in Docker Swarm?

**Answer:**
Worker nodes execute the tasks (containers) assigned by the Manager nodes. They do not participate in Raft consensus but report the state of tasks to managers.

---

### Question 40: How does Docker Swarm handle load balancing?

**Answer:**
Swarm uses **Ingress Load Balancing**:
- It exposes the port on **every** node in the cluster (Routing Mesh).
- Requests to any node are routed to an available running container.
- It also uses internal DNS for service discovery and load balancing between services.

---

## 🔐 Docker Security & Best Practices (Questions 41-50)

### Question 41: How do you secure a Docker container?

**Answer:**
Security is multi-layered:
1.  **Minimal Base Images:** Use Alpine or Distroless.
2.  **Non-Root User:** Don't run as root.
3.  **Read-Only Filesystem:** Use `--read-only`.
4.  **Limit Resources:** Use cgroups (memory/cpu limits).
5.  **Scan Images:** Check for CVEs.
6.  **Capabilities:** Drop unused Linux capabilities (`--cap-drop`).

---

### Question 42: What is Docker Content Trust?

**Answer:**
Docker Content Trust (DCT) uses digital signatures to verify the integrity and publisher of all data received from or sent to a registry.

**Enable:**
```bash
export DOCKER_CONTENT_TRUST=1
```
*When enabled, Docker only runs signed images.*

---

### Question 43: What are Docker secrets?

**Answer:**
Docker Secrets is a secure way to store and manage sensitive data (passwords, keys) in Swarm mode. Secrets are encrypted at rest and in transit.

**Usage:**
```bash
# Create secret
echo "my_password" | docker secret create db_pass -

# Use in service
docker service create --name db --secret db_pass postgres
```

---

### Question 44: How do you scan Docker images for vulnerabilities?

**Answer:**
Use the `docker scan` command (powered by Snyk) or other tools like Trivy or Clair.

**Command:**
```bash
docker scan my-image:latest
```
*Output lists CVEs and severity levels.*

---

### Question 45: What are best practices for writing Dockerfiles?

**Answer:**
1.  **Use specific tags:** Avoid `latest` (e.g., `node:14-alpine`).
2.  **Minimize Layers:** Combine RUN commands.
3.  **Leverage Cache:** Order commands effectively (copy package.json before code).
4.  **Multi-stage builds:** To reduce size.
5.  **Lint:** Use tools like Hadolint.
6.  **.dockerignore:** Exclude unnecessary files.

---

### Question 46: Why should you use non-root users in containers?

**Answer:**
Running as root (default) is a security risk. If a container breakout occurs, the attacker gains root access to the host.

**Implementation:**
```dockerfile
RUN adduser -D myuser
USER myuser
```

---

### Question 47: What are the risks of running privileged containers?

**Answer:**
Running with `--privileged` gives the container all capabilities of the host kernel.
- **Risk:** It bypasses container isolation.
- **Consequence:** Malicious code can modify host system configuration, load kernel modules, or access all devices.

---

### Question 48: What is image hardening?

**Answer:**
Hardening involves securing the container image by reducing the attack surface.
- Removing unnecessary packages (shells, debuggers).
- Ensuring latest security patches.
- Configuring secure defaults (files permissions).
- Using signed images.

---

### Question 49: How do you limit container resources?

**Answer:**
Use flags during `docker run` to prevent a single container from exhausting host resources (DoS).

**Example:**
```bash
docker run -d \
  --memory="512m" \
  --memory-swap="1g" \
  --cpus="1.5" \
  nginx
```

---

### Question 50: How can you prevent container breakout?

**Answer:**
Container breakout involves escaping the isolation to access the host.
**Prevention:**
1.  Keep Docker and Host OS updated.
2.  Do NOT run as root.
3.  Do NOT use `--privileged`.
4.  Use User Namespaces (`userns-remap`).
5.  Enforce AppArmor or SELinux profiles.
6.  Mount filesystems as Read-Only.

---
