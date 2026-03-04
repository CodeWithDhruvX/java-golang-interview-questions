# 🛠️ Docker Operations, Logging & Container Lifecycle (Service-Based Companies)

This document covers Docker operational practices, logging, and container lifecycle management — frequently tested topics in service-based company interviews for DevOps and backend engineering roles.

---

### Q1: What are Docker restart policies and when would you use each one?

**Answer:**
Restart policies determine what Docker does when a container exits (due to crash, manual stop, or daemon restart).

**Available policies:**

| Policy | Behaviour | When to use |
|---|---|---|
| `no` (default) | Never restarts | Development/testing — you want to examine the exit state |
| `always` | Always restarts, even after daemon restart | Long-running production services |
| `unless-stopped` | Restarts unless explicitly stopped by user | Production — won't restart if you `docker stop` it manually |
| `on-failure[:max-retries]` | Restarts only on non-zero exit code | Batch jobs / retry-on-crash scenarios |

**Usage:**
```bash
docker run -d --restart=always nginx
docker run -d --restart=unless-stopped myapp:latest
docker run -d --restart=on-failure:5 myjob:latest   # Max 5 retries

# In Docker Compose
services:
  api:
    image: myapi:latest
    restart: unless-stopped
```

**`always` vs `unless-stopped`:**
- Both restart after daemon restarts.
- The difference: if you `docker stop` a container running with `always`, it will still restart when the daemon restarts. With `unless-stopped`, it stays stopped because it remembers you explicitly stopped it.

---

### Q2: Explain Docker logging drivers. What are the options and which should you use in production?

**Answer:**
Docker captures container stdout/stderr and routes it through a **logging driver** that determines how and where logs are stored/forwarded.

**Available logging drivers:**

| Driver | Description | When to use |
|---|---|---|
| `json-file` (default) | Logs stored as JSON files on host at `/var/lib/docker/containers/<id>/<id>-json.log` | Development, small deployments |
| `none` | No logs collected | Batch jobs where logs go to a volume |
| `syslog` | Send logs to syslog | Traditional Linux servers |
| `journald` | Send to systemd journal | systemd-based hosts |
| `fluentd` | Forward to Fluentd aggregator | Production with centralized logging |
| `awslogs` | Direct to Amazon CloudWatch Logs | AWS deployments |
| `gelf` | Graylog Extended Log Format | Graylog deployments |
| `splunk` | Forward to Splunk | Enterprise Splunk setups |
| `local` | Compressed, rotated local files | Better performance than json-file |

**Configure default driver for all containers (`/etc/docker/daemon.json`):**
```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "50m",
    "max-file": "5",
    "compress": "true"
  }
}
```

**Per-container logging driver:**
```bash
docker run --log-driver=awslogs \
  --log-opt awslogs-region=ap-south-1 \
  --log-opt awslogs-group=/myapp/production \
  myapp:latest
```

**Viewing logs:**
```bash
docker logs <container>             # All logs
docker logs -f <container>          # Follow (like tail -f)
docker logs --tail 100 <container>  # Last 100 lines
docker logs --since 1h <container>  # Last 1 hour
docker logs --until 30m <container> # Until 30 minutes ago
docker logs -t <container>          # Include timestamps
```

**Production recommendation:** Use `json-file` with rotation for non-ECS deployments, or the cloud provider's native driver (e.g., `awslogs` on AWS ECS). Avoid storing logs in containers without rotation — the log file can fill up the node's disk.

---

### Q3: How do environment variables and secrets work in Docker? What is the difference between `-e`, `--env-file`, and Docker secrets?

**Answer:**

**Method 1: `-e` flag (inline, good for dev)**
```bash
docker run -e DB_HOST=localhost -e DB_PORT=5432 -e DB_PASS=secret myapp:latest
```
*Con:* Commands with passwords are visible in shell history, `docker inspect`, and `ps aux`.

**Method 2: `--env-file` (better for dev)**
```bash
# .env file (add to .gitignore!)
DB_HOST=localhost
DB_PORT=5432
DB_PASS=supersecret

docker run --env-file .env myapp:latest
```
*Con:* File on disk, must be secured and distributed to each host.

**Method 3: Docker Compose `.env` file (devs love this)**
```
# .env file in project root
POSTGRES_PASSWORD=devpassword
APP_PORT=3000
```
```yaml
# compose.yml
services:
  app:
    ports:
      - "${APP_PORT}:3000"   # Automatically reads from .env
  db:
    environment:
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD}  # Substituted from .env
```

**Method 4: Docker Secrets (Docker Swarm — for production)**
Docker Secrets store sensitive data encrypted at rest in the Swarm Raft log and delivers them to containers as **files** at `/run/secrets/<secret_name>`.
```bash
# Create a secret
printf "supersecretpassword" | docker secret create db_password -

# Use in stack
services:
  db:
    image: postgres:15
    secrets:
      - db_password
    environment:
      - POSTGRES_PASSWORD_FILE=/run/secrets/db_password
secrets:
  db_password:
    external: true
```

**Comparison table:**

| Method | Security | Use case |
|---|---|---|
| `-e` flag | Low (visible in ps/inspect) | Quick local testing only |
| `--env-file` / `.env` | Medium (file-based) | Local development |
| Cloud secrets (SSM, Vault) | High | Production (fetch at startup) |
| Docker Secrets (Swarm) | High | Docker Swarm production |

---

### Q4: How do you optimize a Dockerfile to produce smaller, faster-building images?

**Answer:**

**Technique 1: Use a minimal base image**
```dockerfile
# BAD: Full Ubuntu = ~70MB OS overhead
FROM ubuntu:22.04

# GOOD: Alpine Linux = ~5MB
FROM alpine:3.19

# BEST for final stage: distroless (no shell, no package manager)
FROM gcr.io/distroless/static-debian12
```

**Technique 2: Multi-stage builds**
```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# Stage 2: Minimal runtime image
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
# Final image: ~10MB instead of ~900MB
```

**Technique 3: Order layers by change frequency (cache optimization)**
```dockerfile
# Dependencies change RARELY — put them at the top
COPY package.json package-lock.json ./
RUN npm ci

# Source code changes OFTEN — put it at the bottom
COPY . .
RUN npm run build
```
Docker caches each layer. If `package.json` doesn't change, the `npm ci` layer is reused (cache hit), saving minutes on every build.

**Technique 4: Combine RUN instructions**
```dockerfile
# BAD: 3 separate layers
RUN apt-get update
RUN apt-get install -y curl wget
RUN rm -rf /var/lib/apt/lists/*

# GOOD: Single layer — AND cleanup in same layer
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl wget && \
    rm -rf /var/lib/apt/lists/*
```

**Technique 5: Use `.dockerignore`**
```
# .dockerignore
node_modules/
.git/
.gitignore
*.md
tests/
.env
dist/
```
Prevents large/unnecessary directories from being sent to the Docker build context, speeding up `docker build`.

---

### Q5: How do you monitor resource usage of running Docker containers?

**Answer:**

**`docker stats` — Real-time resource monitoring**
```bash
docker stats                         # All running containers
docker stats <container1> <container2>  # Specific containers
docker stats --no-stream             # Snapshot (non-streaming output)
docker stats --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}"
```

**Output columns:**
| Column | Meaning |
|---|---|
| `CPU %` | Container's CPU usage % of total host CPU |
| `MEM USAGE / LIMIT` | Memory used vs. limit set |
| `MEM %` | Memory usage as % of limit |
| `NET I/O` | Network bytes in/out |
| `BLOCK I/O` | Disk bytes read/written |
| `PIDs` | Number of processes in container |

**Set resource limits:**
```bash
# Limit CPU (0.5 = half a core) and RAM
docker run -d \
  --cpus="0.5" \
  --memory="512m" \
  --memory-swap="512m" \    # Same as memory = no swap
  myapp:latest

# In Compose
services:
  api:
    deploy:
      resources:
        limits:
          cpus: "0.5"
          memory: 512M
        reservations:
          cpus: "0.25"
          memory: 256M
```

**External monitoring tools:**
- **cAdvisor** (Google) — Exposes Docker container metrics as Prometheus metrics
- **Prometheus + Grafana** — Industry standard for metric collection and dashboarding
- **Datadog / New Relic** — Commercial APM with Docker integration
- **Portainer** — Web UI for managing and monitoring containers

---

### Q6: What is `docker system prune` and how do you clean up Docker resources?

**Answer:**
Over time, Docker accumulates unused resources — old images, stopped containers, dangling volumes — which can fill up disk space significantly.

**Targeted cleanup commands:**
```bash
# Remove all stopped containers
docker container prune

# Remove all unused images (dangling = no tag; includes unused if -a)
docker image prune           # Only dangling images
docker image prune -a        # All images not used by any container

# Remove unused volumes
docker volume prune

# Remove unused networks
docker network prune

# Remove build cache
docker builder prune

# Nuclear option — removes EVERYTHING unused
docker system prune -a --volumes
# WARNING: This removes ALL unused images, not just dangling ones!
```

**Targeted filter-based pruning:**
```bash
# Remove images older than 7 days
docker image prune -a --filter "until=168h"

# Remove stopped containers older than 24h
docker container prune --filter "until=24h"

# Remove build cache older than 1 day
docker builder prune --filter "until=24h"
```

**Check disk usage:**
```bash
docker system df              # Summary
docker system df -v           # Verbose — lists individual items
```

**Automate pruning in production (cron job):**
```bash
# /etc/cron.d/docker-prune
0 2 * * 0 root docker system prune -f --filter "until=168h" >> /var/log/docker-prune.log 2>&1
```

---

*Prepared for DevOps and backend engineering screening rounds at service-based companies.*
