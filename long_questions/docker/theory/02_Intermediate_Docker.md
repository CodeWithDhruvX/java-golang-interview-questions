# 🐳 **Intermediate Docker Interview Questions (21–40)**

---

### 21. What is the difference between a bind mount and a volume?
"Both persist data outside a container's writable layer, but they differ in **who manages them**.

A **bind mount** maps a specific host path into the container: `docker run -v /home/user/code:/app`. The host path must exist. Great for development — I can edit code on the host and see changes live in the container.

A **volume** is managed by Docker: `docker run -v mydata:/app/data`. Docker handles the path. Volumes are isolated, portable, and the correct choice for production databases or stateful services."

#### In-depth
Bind mounts on Docker Desktop (macOS/Windows) pass through a file-sharing layer (osxfs, gRPC-FUSE, WSL2), adding latency. For write-heavy workloads like databases, this overhead is unacceptable. Named volumes bypass this — Docker uses a native Linux filesystem via the VM's storage, achieving near-native I/O.

---

### 22. How do you copy files into a Docker container?
"I use `docker cp <source> <container>:<destination>`.

For example: `docker cp ./config.json mycontainer:/app/config.json`. It works bidirectionally — I can also copy files out: `docker cp mycontainer:/app/logs ./logs`.

This is useful for debugging: if a container is misbehaving I can copy out its log files or configs without stopping it. However, for build-time file inclusion, I use `COPY` in the Dockerfile."

#### In-depth
`docker cp` bypasses the container filesystem layers — it directly accesses the container's writable layer via the Docker daemon. It doesn't require the container to have file utilities (`cp`, `tar`) installed. This makes it ideal for minimal containers (Alpine, distroless) where shell tools may not be available.

---

### 23. What is the purpose of multi-stage builds?
"Multi-stage builds allow you to use **multiple `FROM` instructions** in a single Dockerfile — each defines a stage. You can selectively copy artifacts from one stage to another.

The classic pattern: a `builder` stage with the full SDK (Go compiler, Node.js toolchain) that compiles the app, and a final `runner` stage based on a tiny image (Alpine, distroless) containing only the compiled binary.

This eliminates the need for separate build scripts and gives you small, secure production images without shipping the entire build toolchain."

#### In-depth
Named stages (`FROM golang:1.22 AS builder`) allow selective targeting. `docker build --target builder .` builds only up to that stage — useful for running tests in CI without building the final image. The build cache is still applied per stage, so incremental rebuilds remain fast.

---

### 24. How do you pass environment variables to a container?
"I use `-e` or `--env` at runtime: `docker run -e DATABASE_URL=postgres://... myapp`.

For multiple variables, a `.env` file is cleaner: `docker run --env-file .env myapp`.

In the Dockerfile, `ENV` sets a default value at build time. `ARG` is a build-time variable only — it doesn't exist in the running container. I never hardcode secrets in Dockerfiles or `-e` flags in scripts; instead I use a secret manager (Vault, AWS SSM) and inject at runtime."

#### In-depth
`ARG` values can be inspected via `docker history` — they are not secret. `ENV` values are embedded in the image metadata and visible via `docker inspect`. For actual secrets, use BuildKit secret mounts (`RUN --mount=type=secret`) which don't persist in image layers at all.

---

### 25. What is the use of `--rm` flag?
"`--rm` tells Docker to **automatically remove the container when it exits**.

I use it for one-off tasks: `docker run --rm alpine echo hello`. Without `--rm`, every `docker run` creates a new container that persists after exit, accumulating as stopped containers and wasting disk space.

For development scripts, testing, and debugging I always add `--rm`. For long-lived services I omit it so Docker can restart them with preserved state information."

#### In-depth
`--rm` is incompatible with `-d` (detach) in some older Docker versions, but works cleanly in modern ones. Under the hood, Docker registers a listener on container exit events and triggers container removal. It also removes anonymous volumes associated with the container (but not named volumes).

---

### 26. What are Docker namespaces?
"Linux **namespaces** are the kernel feature that provides isolation for containers. Docker uses six main namespaces:

**PID**: container gets its own process tree (PID 1 is the entrypoint). **Network**: its own network stack. **Mount**: its own view of the filesystem. **UTS**: its own hostname. **IPC**: isolated inter-process communication. **User**: (optional) maps container UIDs to host UIDs.

Namespaces are what makes containers feel like separate systems while sharing the host kernel."

#### In-depth
Unlike VMs which virtualize hardware, namespaces are a kernel-level illusion. A process in a PID namespace sees processes 1, 2, 3... but on the host, it's processes 5432, 5433, 5434... The kernel translates transparently. This is why container startup is milliseconds — no OS boot, just namespace creation.

---

### 27. What are Docker cgroups?
"**cgroups** (control groups) are the Linux kernel feature Docker uses to **limit and meter resources** per container.

They control CPU (time slices, cores), memory (hard limits, swap), disk I/O (bandwidth, IOPS), and network I/O. When you run `docker run --memory=512m --cpus=0.5 myapp`, Docker translates these into cgroup settings in `/sys/fs/cgroup/`.

Without cgroups, one container could consume all host CPU/memory and starve others — which is critical in multi-tenant environments."

#### In-depth
cgroup v2 (unified hierarchy) became the default in modern kernels and Docker 20.10+. It simplified management and added features like memory QoS and PSI (Pressure Stall Information). On systemd-based systems, Docker creates cgroup slices under `system.slice/docker-<id>.scope`.

---

### 28. What is Docker Compose?
"Docker Compose is a tool for defining and running **multi-container applications** using a `docker-compose.yml` file.

Instead of running multiple `docker run` commands, I define all services, networks, and volumes in one YAML file. Then `docker compose up -d` starts the full stack. It handles inter-service networking, dependency ordering, and scaling.

I use it for every local development setup. A team member can run `docker compose up` and have the entire stack (API, DB, cache, message broker) running instantly — no manual setup."

#### In-depth
Docker Compose v2 (bundled with Docker Desktop) replaced the standalone `docker-compose` binary. Key difference: it runs as a Docker CLI plugin (`docker compose`) instead of a separate Python binary, has better performance, and supports the Compose Specification — a vendor-neutral standard for multi-container definitions.

---

### 29. How do you define services in Docker Compose?
"In `docker-compose.yml`, each service is a key under `services:`.

```yaml
services:
  api:
    image: myapp:latest
    ports:
      - '3000:3000'
    environment:
      - DB_URL=postgres://db:5432/mydb
    depends_on:
      - db
  db:
    image: postgres:15
    volumes:
      - pgdata:/var/lib/postgresql/data
volumes:
  pgdata:
```

Each service can define image, build context, ports, environment, volumes, networks, and more."

#### In-depth
Service names act as DNS hostnames within the Compose-created network. So `api` can connect to `db` simply by using `db` as the hostname. This is why database URLs in Compose apps use the service name: `postgres://db:5432/mydb`. Docker's embedded DNS resolves service names to container IPs automatically.

---

### 30. How do you scale services in Docker Compose?
"I use `docker compose up --scale api=3` to start 3 replicas of the `api` service.

This is useful for testing load balancing locally. However, if services have a fixed host port mapping (e.g., `ports: - '3000:3000'`), scaling will fail because multiple containers can't bind the same host port. I remove the host port mapping and put a load balancer (nginx, Traefik) in front.

For production scaling, Docker Swarm or Kubernetes is more appropriate than Compose."

#### In-depth
When scaling with Compose, containers get names like `project-api-1`, `project-api-2`. Docker's DNS for the service name (`api`) returns all container IPs in round-robin. Services that connect to `api` will naturally load balance across instances. But this is client-side DNS load balancing — not as robust as a proper LB.

---

### 31. What is the `depends_on` in Docker Compose?
"`depends_on` controls **startup order** — it ensures one service starts before another.

```yaml
api:
  depends_on:
    - db
```

This starts `db` before `api`. But critically: it only waits for the container to **start**, not for the service inside it to be **ready**. Postgres takes a few seconds to initialize after the container starts.

I use `condition: service_healthy` with a healthcheck for true readiness: `depends_on: db: condition: service_healthy`."

#### In-depth
The naive `depends_on` trap causes startup failures in 90% of new Compose setups. The solution pattern with healthchecks: define a `healthcheck` on the db service using `pg_isready`, then use `condition: service_healthy` in `depends_on`. Now `api` won't start until Postgres is actually accepting connections.

---

### 32. What is the difference between `COPY` and `ADD` in Dockerfile?
"Both copy files from the build context into the image. But `ADD` has **extra magic** that makes it dangerous.

`COPY src dst` is explicit and predictable. I use it for 99% of cases.

`ADD` can: automatically extract `.tar.gz` archives, and download from URLs. These hidden behaviors make builds non-obvious. I only use `ADD` when I specifically need the tarball extraction feature — and document why."

#### In-depth
The Docker team itself recommends preferring `COPY` over `ADD` in their best practices documentation. `ADD` with a URL doesn't benefit from layer caching (each build re-downloads), and downloading from the internet in a build is a security risk. For URL downloads, use `RUN curl` with checksum verification instead.

---

### 33. How do you connect containers using Docker networks?
"I attach containers to the **same Docker network**.

```bash
docker network create mynet
docker run -d --name api --network mynet myapp
docker run -d --name db --network mynet postgres
```

Now `api` can reach `db` using the hostname `db`. Docker's embedded DNS handles the resolution.

In Compose, all services automatically share a project network — no manual steps needed."

#### In-depth
A container can belong to **multiple networks simultaneously**. You can attach/detach a running container to a network with `docker network connect/disconnect`. This is useful for adding a debugging sidecar to a running service's network without restarting it.

---

### 34. What are the different types of Docker networks?
"Docker has five network drivers:

**bridge**: Default for single-host setups. Creates a virtual switch. **host**: Container shares host's network stack directly — no isolation, maximum performance. **none**: No networking at all. **overlay**: For multi-host communication in Swarm. **macvlan**: Assigns a MAC address to the container — it appears as a physical device on the network.

For most applications: `bridge` for local, `overlay` for Swarm/production clusters."

#### In-depth
**Host networking** removes the NAT overhead and is the highest-performance option — useful for latency-sensitive services like real-time gaming servers or high-frequency trading. But it completely removes network isolation: the container can bind any host port and see all host network interfaces.

---

### 35. What is Docker Swarm?
"Docker Swarm is Docker's **built-in container orchestration** tool.

It turns a pool of Docker hosts into a single virtual Docker host. You define **services** (desired state), and Swarm schedules containers (**tasks**) across nodes, handles failures, and manages rolling updates.

I use it for simpler production workloads where Kubernetes overhead isn't justified. If you're already comfortable with Docker CLI, Swarm is the natural next step — it uses the same `docker-compose.yml` format with Swarm-specific extensions."

#### In-depth
Swarm uses **Raft consensus** for leader election among manager nodes. With 3 managers, it can tolerate 1 failure; with 5, it tolerates 2. For production, odd numbers of managers (3 or 5) are recommended. Workers are stateless — they execute tasks assigned by managers.

---

### 36. How do you create a Docker Swarm cluster?
"On the first node (the manager): `docker swarm init --advertise-addr <manager-ip>`.

This outputs a join token. On worker nodes: `docker swarm join --token <token> <manager-ip>:2377`.

To add more managers: `docker swarm join-token manager` to get the manager join token, then run it on the new node.

To verify: `docker node ls` on a manager shows all nodes and their roles."

#### In-depth
Port 2377 (TCP) is used for Swarm management traffic. Port 7946 (TCP/UDP) is for container network discovery. Port 4789 (UDP) is for the overlay network data path. All must be open between nodes. In cloud environments (AWS, GCP), ensure security groups allow these ports between nodes.

---

### 37. What is the difference between a service and a container in Docker Swarm?
"In Swarm, a **service** is the **desired state declaration** — 'I want 3 replicas of my API running'. A **task** is a slot for a container. A **container** is the actual running process instantiating the task.

A service is the management abstraction: I create, scale, update, and delete services. Swarm automatically creates/destroys the underlying containers to match the desired state.

Outside Swarm: containers are managed individually. Inside Swarm: I never touch containers directly — I manage services."

#### In-depth
When a container (task) fails, Swarm automatically schedules a replacement task on a healthy node. This self-healing is the core value proposition. The controller loop in the Swarm manager continuously reconciles actual state with desired state — similar to Kubernetes controllers.

---

### 38. What is a manager node in Docker Swarm?
"Manager nodes are the **control plane** of a Swarm cluster. They run the Raft consensus algorithm to maintain cluster state, schedule services onto nodes, and serve the Swarm API.

You can have multiple managers for high availability (minimum 3 for production). Managers can also run service tasks (by default), but for large clusters I drain manager nodes so they focus on scheduling: `docker node update --availability drain <manager-node>`."

#### In-depth
The Raft log on manager nodes stores all cluster state: services, tasks, networks, secrets. Quorum requires a majority of managers to agree. This means with 3 managers you need 2 alive (quorum = 2), with 5 you need 3. Avoid even numbers — with 4 managers you need 3 for quorum, which is the same fault tolerance as 3 managers with extra cost.

---

### 39. What is a worker node in Docker Swarm?
"Worker nodes **execute the tasks** (containers) assigned to them by manager nodes. They have no visibility into cluster state and cannot make scheduling decisions.

Workers communicate with managers over the management plane. They receive task assignments, run containers, and report status back. You can promote a worker to manager or demote a manager to worker at runtime.

For separation of concerns in production, I keep dedicated workers for application workloads and dedicated managers for control functions."

#### In-depth
Workers join the cluster with a token provisioned by the manager. This token can be rotated (`docker swarm join-token --rotate worker`) for security. Worker nodes only need connectivity to the manager's API port (2377) — they don't need to reach each other directly (that's handled at the overlay network layer).

---

### 40. How does Docker Swarm handle load balancing?
"Swarm has **two layers** of load balancing:

**Internal**: Service discovery via DNS. Containers within the overlay network resolve a service name to a virtual IP (VIP), which routes to a healthy backend task round-robin.

**External (ingress)**: The `ingress` network provides a mesh routing network. Any node in the Swarm can accept traffic on a published port, and Swarm routes it to any healthy task — even if that task isn't on the receiving node. This is called **routing mesh**."

#### In-depth
The routing mesh uses IPVS (IP Virtual Server) in the Linux kernel for load balancing — it operates at L4 (TCP/UDP), not HTTP. For L7 (HTTP/HTTPS) load balancing with host-based routing, I place a reverse proxy (Traefik, Caddy) in front that integrates with Docker labels to discover services and auto-configure routes.

---
