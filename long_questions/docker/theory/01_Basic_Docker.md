# 🐳 **Basic Docker Interview Questions (1–20)**

---

### 1. What is Docker?
"Docker is an open-source **containerization platform** that lets you package an application and all its dependencies into a standardized unit called a **container**.

Before Docker, the classic problem was 'it works on my machine'. Docker solves this by packaging the app, runtime, libraries, and config together. The container runs identically on a developer's laptop, a CI server, or a production cloud instance.

I think of it as a lightweight, portable, self-sufficient unit that isolates your app from the host OS while still sharing the OS kernel — unlike a full VM."

#### In-depth
Docker uses Linux kernel features like **namespaces** (for isolation) and **cgroups** (for resource limits). Unlike VMs that virtualize hardware, containers virtualize the OS. This is why containers start in milliseconds and use MBs instead of GBs.

---

### 2. What are the benefits of using Docker?
"Docker gives you **portability**, **consistency**, and **speed**.

Portability: a container runs the same everywhere. Consistency: no more environment drift between dev, staging, and prod. Speed: containers start in seconds, images build in minutes using layer caching.

It also encourages a **microservices mindset** — each service gets its own container, making teams independent. And CI/CD becomes simpler because you're shipping the same tested image from pipeline to production."

#### In-depth
Docker's layer-based image system enables efficient storage and transfer — only changed layers are downloaded/uploaded. Combined with Docker Hub as a registry, it creates a complete artifact distribution system, similar to npm for Node.js but for entire runtime environments.

---

### 3. What is the difference between Docker and a virtual machine?
"Both provide **isolation**, but they do it at different levels.

A VM virtualizes hardware: you get a full OS (kernel + userland) inside a hypervisor (VMware, VirtualBox). This means GBs of overhead and minutes to start. A Docker container virtualizes only the OS: it shares the host kernel but isolates the process space, filesystem, and network.

I use the analogy: a VM is like renting a separate house; a container is like renting a room in the same house — you share electricity and plumbing (the kernel), but your space is private."

#### In-depth
Since containers share the host kernel, they cannot run a different OS entirely (e.g., you can't run a Windows container on a Linux host without special bridge layers). However, this shared kernel means near-native performance with minimal overhead — containers typically have <5% CPU overhead vs. 5–20% for VMs.

---

### 4. What is a Docker container?
"A container is a **running instance of a Docker image**.

It's an isolated process (or group of processes) with its own filesystem, network interface, and process tree, all governed by Linux namespaces and cgroups. It's ephemeral by default — when it stops, its writable layer is gone unless you use volumes.

I think of it like a process with superpowers: isolated from the host, yet lightweight enough to run hundreds on a single machine."

#### In-depth
Each container gets a thin writable layer on top of the read-only image layers (Copy-on-Write). Changes inside a container (writing a file) modify only this writable layer. When the container is removed, that layer is deleted — making volumes essential for persistent data.

---

### 5. What is a Docker image?
"A Docker image is a **read-only template** used to create containers.

It's made of stacked layers — each instruction in a Dockerfile adds a layer. The layers are cached and reused, making builds fast. An image is like a class in OOP; a container is an instance of that class.

I push images to a registry (like Docker Hub or ECR) and pull them anywhere. The image ID is a SHA256 digest that guarantees integrity — the same ID always means the same content."

#### In-depth
Images use a **union filesystem** (OverlayFS on most modern systems). Layers are content-addressable by their SHA256 digest. If two images share a base layer, Docker stores it only once on disk. This deduplication is critical for storage efficiency at scale.

---

### 6. What is Docker Hub?
"Docker Hub is Docker's **official public container registry** — like GitHub, but for images.

You can push and pull images there. Official images (nginx, postgres, python) are vetted and maintained by Docker or the vendors. Anyone can publish public images; private repos require a subscription.

I use Docker Hub for public open-source images and for personal experiments. For production at companies, I typically use **AWS ECR**, **GCR**, or a self-hosted **Harbor** registry for security and access control."

#### In-depth
Docker Hub has rate limits for unauthenticated pulls (100 pulls/6 hours for anonymous, 200/6h for free accounts). In high-throughput CI/CD environments, this can cause build failures. The solution is to authenticate or mirror critical images to a private registry.

---

### 7. What is a Dockerfile?
"A Dockerfile is a **text file with instructions** that Docker uses to build an image.

Each instruction (`FROM`, `RUN`, `COPY`, `CMD`, etc.) creates a new layer. Docker reads it top-to-bottom and executes each step. The result is a reproducible, version-controlled image build process.

I treat Dockerfiles like code — they live in the repo, go through code review, and are linted with tools like **Hadolint**. A good Dockerfile is minimal, ordered for caching efficiency, and contains no secrets."

#### In-depth
The order of instructions matters heavily for **build caching**. Docker caches layers and invalidates cache from the point of the first change down. Putting frequently changing instructions (like `COPY . .`) after rarely changing ones (like `RUN apt-get install`) maximizes cache hits and reduces build time.

---

### 8. How do you build a Docker image?
"I use `docker build -t my-app:1.0 .`

The `-t` flag tags the image with a name and version. The final `.` is the **build context** — the directory Docker sends to the daemon. I always have a `.dockerignore` to exclude `node_modules`, `.git`, and build artifacts from this context.

For faster builds I use `--cache-from` to pull a previously built image and reuse its layers in CI pipelines."

#### In-depth
`docker build` sends the entire context directory to the Docker daemon. A large context (accidentally including large files) dramatically slows the build because it all gets transferred over the socket. Always verify your `.dockerignore` is working with `docker build --no-cache` to spot surprises.

---

### 9. How do you run a Docker container?
"The basic command is `docker run nginx`.

But in practice I always add flags: `-d` for detached mode (background), `-p 8080:80` to map host port 8080 to container port 80, `--name my-nginx` to give it a name, and `--rm` to clean it up on exit.

For a full local development run: `docker run -d --name api -p 3000:3000 -v $(pwd):/app my-app:latest`."

#### In-depth
`docker run` is actually `docker create` + `docker start` combined. Behind the scenes, Docker creates the container filesystem, sets up networking, and then starts the entrypoint process. Each freshly started container gets a new writable layer on top of the image.

---

### 10. How do you stop a running container?
"I use `docker stop <container>`.

`stop` sends a **SIGTERM** to the main process, giving it 10 seconds to shut down gracefully. If it doesn't exit, Docker sends **SIGKILL**. I can adjust the timeout with `--time` (e.g., `docker stop --time=30`).

For unresponsive containers I use `docker kill`, which sends SIGKILL immediately without a grace period. I always prefer `stop` because it allows clean shutdown — flushing buffers, closing connections, etc."

#### In-depth
The 10-second default timeout was chosen by Docker as a balance. Applications should handle SIGTERM for graceful shutdown. In Go, I typically use `signal.NotifyContext` or `os/signal` to listen for SIGTERM and run shutdown logic (draining HTTP in-flight requests, closing DB connections) before exiting.

---

### 11. What is the difference between `CMD` and `ENTRYPOINT`?
"`ENTRYPOINT` defines the **executable** — the main program to run. `CMD` provides **default arguments** to that executable.

When used together, `CMD` arguments are appended to `ENTRYPOINT`. This is the preferred pattern: `ENTRYPOINT ["python", "app.py"]` + `CMD ["--port", "8080"]`. The user can override CMD at runtime without changing the entrypoint.

Alone, `CMD` is simply the default command. I can override it completely: `docker run myimage --port 9090`."

#### In-depth
Always use the **exec form** (`["executable", "arg"]`) rather than shell form (`executable arg`). Shell form wraps your command in `sh -c`, meaning PID 1 is the shell, not your process. This causes issues with signal handling — SIGTERM goes to the shell, not your app, breaking graceful shutdown.

---

### 12. How do you list all Docker containers?
"I use `docker ps` to list **running** containers.

To see all containers including stopped ones, I add `-a`: `docker ps -a`.

For scripting I often use `--format` to get specific fields: `docker ps --format '{{.Names}}\t{{.Status}}'`. Or `-q` to get only container IDs, which I pipe into other commands: `docker rm $(docker ps -aq -f status=exited)`."

#### In-depth
`docker ps` reads from the Docker daemon's internal state store, not from `/proc`. Container states include: `created`, `running`, `paused`, `restarting`, `removing`, `exited`, and `dead`. The filter flags (`-f`) are powerful for scripting automated cleanup.

---

### 13. How do you list all Docker images?
"I use `docker images` or the equivalent `docker image ls`.

This shows repository, tag, image ID, creation time, and size. To see all images including intermediate/dangling ones I add `-a`. To filter only dangling images (untagged): `docker images -f dangling=true`.

I regularly audit my local images and clean up with `docker image prune` to reclaim disk space."

#### In-depth
Image sizes shown by `docker images` are **virtual sizes** — they include all shared layers. Actual disk usage is lower because layers are shared. Use `docker system df` to see real disk usage breakdown across images, containers, and volumes.

---

### 14. How do you delete a Docker container?
"I use `docker rm <container-name-or-id>`.

The container must be stopped first. To force-remove a running container: `docker rm -f mycontainer`. To remove all stopped containers at once: `docker container prune`.

I always add `--rm` to `docker run` for ephemeral containers I create for one-off tasks — it automatically removes them when they exit, keeping the system clean."

#### In-depth
`docker rm` only removes the container's writable layer and metadata. It does **not** remove volumes attached to the container by default. To also remove associated anonymous volumes: `docker rm -v mycontainer`. Named volumes must be removed separately with `docker volume rm`.

---

### 15. How do you delete a Docker image?
"I use `docker rmi <image>` or `docker image rm <image>`.

You can specify by name:tag or by image ID. You cannot remove an image that has running containers referencing it — you must stop and remove those containers first.

To clean up all dangling images: `docker image prune`. To remove all unused images (not just dangling): `docker image prune -a`. I run this periodically in CI to prevent disk exhaustion."

#### In-depth
Docker images are reference-counted. An image won't be deleted if any container (even stopped) references it. This is why `docker image prune` safely removes only unreferenced images. Use `docker system prune` as a nuclear option — it removes stopped containers, networks, dangling images, and build cache.

---

### 16. What is the purpose of `.dockerignore` file?
"`.dockerignore` tells Docker which files to **exclude from the build context** when running `docker build`.

It works exactly like `.gitignore`. Common entries: `.git`, `node_modules`, `*.log`, `.env`, `dist/`, `__pycache__`. Excluding these speeds up the build (less data sent to daemon) and prevents secrets from sneaking into the image.

I always include `**/.DS_Store`, `**/*.md` (for docs), and local config overrides. Missing this file is a performance and security red flag."

#### In-depth
The build context is sent as a tar archive to the Docker daemon. Without `.dockerignore`, a repo with `node_modules` (often 500MB+) gets sent on every build — even if those files never change. Build times can go from seconds to minutes just due to context transfer.

---

### 17. How can you expose a port from a Docker container?
"There are two steps: **`EXPOSE`** in the Dockerfile and **`-p`** (publish) at runtime.

`EXPOSE 8080` in a Dockerfile is documentation — it declares which port the app listens on. But it doesn't actually publish it. To make it accessible from the host: `docker run -p 8080:8080 myapp`.

The format is `-p host_port:container_port`. I can also bind to a specific IP: `-p 127.0.0.1:8080:8080` to prevent external access."

#### In-depth
`-P` (capital P) automatically publishes all `EXPOSE`d ports to random, ephemeral host ports. Useful for testing but not for production. For multi-container apps, containers on the same Docker network communicate over container ports directly — no need for port publishing between services.

---

### 18. What is the default Docker network?
"The default network is the **bridge** network, called `bridge`.

When you run a container without specifying a network, it joins the default bridge. Containers can communicate by IP, but NOT by name (unless you use a custom bridge network, which enables DNS-based discovery).

I always create custom bridge networks for applications: `docker network create myapp-net`. This gives me container-name DNS resolution and network isolation from other containers."

#### In-depth
The default bridge network is a legacy mode with several limitations: no automatic DNS, no ICC (inter-container communication) control per network, all containers can see each other. Custom bridge networks solve all this. Docker's user-defined bridge networks use an embedded DNS server at `127.0.0.11`.

---

### 19. What is a volume in Docker?
"A volume is Docker-managed persistent storage that exists **outside the container's writable layer**.

Unlike the container filesystem (which is deleted when the container is removed), volumes persist. They are stored at `/var/lib/docker/volumes/` on Linux. You create them with `docker volume create` or implicitly in `docker run -v`.

I use named volumes for databases: `docker run -v pgdata:/var/lib/postgresql/data postgres`. The data survives container restarts, upgrades, and replacements."

#### In-depth
Volumes are preferred over bind mounts for production because: Docker manages them, they can be shared between containers, they can be backed up easily with `docker run --volumes-from`, and they have better performance on Linux (no host filesystem translation overhead unlike bind mounts on macOS/Windows).

---

### 20. How do you persist data in Docker?
"There are three ways: **volumes**, **bind mounts**, and **tmpfs mounts**.

Volumes are the preferred production approach — Docker manages them, they're portable and easy to back up. Bind mounts map a host directory into the container (ideal for development: `docker run -v $(pwd)/code:/app myapp`). tmpfs mounts store data in host memory — useful for sensitive temporary data that shouldn't touch disk.

For databases I always use named volumes. For dev hot-reload I use bind mounts."

#### In-depth
Bind mounts have a critical pitfall: they expose host filesystem paths inside containers, creating tight coupling between host and container. They can also create permission issues (UID/GID mismatch between host user and container user). Volumes abstract this away, and Docker handles ownership correctly when creating volumes.

---
