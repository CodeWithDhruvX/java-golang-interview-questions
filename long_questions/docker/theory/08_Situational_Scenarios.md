# 🧠 **Situational & Scenario-Based Questions (91–100)**

---

### 91. How would you troubleshoot a container that keeps restarting?
"I follow a systematic approach:

1. `docker ps` — check the STATUS and RESTARTS count
2. `docker inspect container --format='{{.State.ExitCode}}'` — check the exit code
3. `docker logs container` — see the last output before crash
4. `docker events` — see the restart events with timestamps

Exit code 137 = OOM killed (increase memory limit). Exit code 1 = app crash (check logs for error). Exit code 143 = graceful terminate received.

The most common causes: memory limits too low, app crashing on startup (config error, missing env var, DB connection failure), or a dependency service not ready yet."

#### In-depth
Add `--restart=on-failure:3` as a restart policy with a count limit — it prevents infinite restart loops from masking a real bug. Also add `--log-opt max-size=10m` to prevent log files from filling the disk during a crash loop. For proper debugging, temporarily override the entrypoint: `docker run --entrypoint sh myimage` to get a shell and manually investigate the startup sequence.

---

### 92. How would you update a running container to a new image?
"The correct approach is **never update in-place** — replace the container.

1. Pull the new image: `docker pull myapp:v2`
2. Stop the old container: `docker stop myapp`
3. Remove it: `docker rm myapp`
4. Start new container with the same run configuration: `docker run -d --name myapp ... myapp:v2`

For zero-downtime, use Docker Swarm: `docker service update --image myapp:v2 myservice` — it replaces tasks one by one, health-checking before moving to the next.

Or use Docker Compose: update the image tag in `docker-compose.yml` and run `docker compose up -d`."

#### In-depth
The danger of `docker update` is that it only changes resource limits — it doesn't change the image. A running container always uses the image it was started with. You cannot update the image of a running container — you must replace it. This is a feature, not a bug: immutability guarantees consistency. Tools like **Watchtower** automate this replacement cycle by polling registries for new image versions.

---

### 93. How would you handle a failed container in a CI/CD pipeline?
"A failed container in CI should: **surface the error clearly**, **preserve artifacts**, and **fail the pipeline**.

1. `docker run --rm myapp:ci run-tests; echo "Exit: $?"` — capture exit code
2. On failure: `docker logs --tail=100 myapp` — print logs to CI output
3. `docker cp myapp:/app/test-reports ./reports` — extract test artifacts
4. `docker rm myapp` — cleanup even on failure
5. The pipeline fails with a non-zero exit code and test reports are uploaded as artifacts

In GitHub Actions: use `if: always()` steps to collect logs/artifacts even when the build fails."

#### In-depth
One CI pattern I like: run the container with `--cidfile id.txt` to capture the container ID even if `docker run` fails. Then always-cleanup with `docker rm $(cat id.txt)`. This handles the case where the container starts, does some work, then fails — you still get the container ID for log extraction. Bash trap functions (`trap 'docker logs $CID' EXIT`) are also useful for guaranteed cleanup.

---

### 94. How would you ensure high availability of your Docker services?
"High availability requires redundancy at every layer:

**Multiple container replicas**: Swarm services with `--replicas=3` spread across multiple nodes.
**Multiple nodes**: at least 3 nodes so one can fail without impact.
**Health checks**: proper health checks ensure unhealthy replicas are replaced promptly.
**Load balancing**: Swarm's routing mesh or external LB distributes traffic.
**Persistent storage**: shared volumes (NFS, EFS) or replicated storage so any node can serve the data.
**Restart policies**: `--restart=unless-stopped` on all services."

#### In-depth
For stateless services, 3 replicas across 3 availability zones is the standard HA setup — tolerates one AZ failure. For stateful services (databases), HA is harder: use database-native replication (Postgres streaming replication, MySQL Group Replication) rather than trying to make containers themselves HA. The database handles the data safety; Docker handles the container lifecycle.

---

### 95. How would you manage secrets in a distributed Docker environment?
"In Swarm, use native **Docker secrets**: encrypted, memory-only delivery, only accessible to services that explicitly request them.

For non-Swarm environments, I use **external secret managers**: HashiCorp Vault, AWS Secrets Manager, or GCP Secret Manager. Inject secrets at container startup via an entrypoint script that fetches and exports them as env vars before starting the app.

Never store secrets in: environment variables baked into images, Dockerfiles, docker-compose.yml committed to git, or container labels. Use `.gitignore` for local `.env` files."

#### In-depth
Vault's **agent sidecar** pattern is the gold standard for secret injection in production: a Vault Agent container shares a tmpfs volume with the app container, fetches and refreshes secrets automatically, and writes them to files. The app reads from that file — it never sees secret values in environment variables (which can leak via `/proc/<pid>/environ`). Kubernetes Vault integration uses mutating webhooks to inject sidecar automatically.

---

### 96. How would you troubleshoot network issues between containers?
"My network debugging toolkit:

1. Verify both containers are on the **same network**: `docker inspect container | grep -A20 Networks`
2. Enter one container and ping the other by name: `docker exec -it container1 ping container2`
3. `docker exec -it container1 nslookup container2` — check DNS resolution
4. `docker exec -it container1 curl http://container2:port/health` — test HTTP connectivity
5. `docker network inspect mynetwork` — see all connected containers and their IPs
6. Check firewall: `iptables -L DOCKER-USER` — Docker may be blocked by host firewall"

#### In-depth
The most common cause of inter-container networking failures: containers are on **different networks**. Compose creates a project-specific network by default — containers in different Compose projects are on different networks and cannot communicate. Solution: attach both services to a shared named network. Second most common: the app is listening on `127.0.0.1` inside the container (loopback only) instead of `0.0.0.0`. Fix: configure the app to bind to all interfaces.

---

### 97. What steps would you take to clean up Docker disk space?
"Systematic cleanup from safest to most aggressive:

1. `docker container prune` — remove all stopped containers
2. `docker image prune` — remove dangling images
3. `docker volume prune` — remove unused volumes (careful: data loss if miscategorized)
4. `docker network prune` — remove unused networks
5. `docker image prune -a` — remove all unused images (not just dangling)
6. `docker builder prune` — remove BuildKit build cache
7. `docker system prune -a --volumes` — nuclear option, removes everything unused

Monitor usage: `docker system df` and `docker system df -v` for details."

#### In-depth
BuildKit's build cache (`docker builder prune`) is often the biggest disk consumer and the safest to clear — it just means the next build is slower. Production CI servers should run `docker builder prune --filter until=24h` nightly. I also use `--storage-opt dm.basesize=20G` to limit the maximum size of a container's writable layer — prevents a runaway process from filling the host disk.

---

### 98. How would you migrate a legacy app to Docker?
"I follow a phased approach:

1. **Understand the app**: document all dependencies (OS packages, runtimes, services it needs)
2. **Create a basic Dockerfile**: start from the closest official base image, install deps, copy code
3. **Test locally**: `docker build && docker run` — fix issues iteratively
4. **Externalize state**: identify any files written by the app (logs, uploads, data) and move to volumes
5. **Externalize configuration**: move hardcoded configs to environment variables
6. **Write a Compose file**: add all required services (DB, cache, etc.)
7. **Update CI/CD**: build and test via Docker in CI
8. **Production deployment**: start with staging, monitor, then production"

#### In-depth
The biggest challenge is usually **configuration** — legacy apps often have config files with hardcoded paths, DB addresses, and environment-specific settings baked in. The containerization effort becomes a config refactoring effort. Tools like **`docker init`** (new in Docker 2023) can generate a starting Dockerfile and Compose file by detecting your language and framework, giving you a 80% complete starting point.

---

### 99. What challenges have you faced while working with Docker?
"Real challenges I've encountered:

**Volume permission errors**: container runs as UID 1000, host files owned by UID 0 — solution: match UIDs or use `chown` in entrypoint.

**Build cache invalidation**: a tiny change high up in the Dockerfile rebuilds all subsequent layers — solution: reorganize instruction order.

**Docker Desktop performance on macOS**: bind mounts are slow — solution: use named volumes, enable VirtioFS.

**Networking complexity**: containers can't reach each other after Compose restart — solution: use named networks explicitly in Compose.

**Image size bloat**: dev dependencies shipped in production — solution: multi-stage builds."

#### In-depth
The volume permission problem is the most painful for new Docker users. The root cause: Docker containers run with a specific UID, and bind-mounted host directories have their own ownership. The cleanest solutions: use named volumes (Docker handles permissions), or use `fixuid` in the entrypoint to remap the container UID to match the host user dynamically.

---

### 100. How would you explain Docker to a non-technical stakeholder?
"I use the **shipping container analogy**:

Before shipping containers, every cargo ship had custom loading methods — goods were loaded differently on every ship, causing delays, damage, and confusion at every port.

Docker is the standard shipping container for software. You pack your application and everything it needs into a container. It runs identically on any server — your laptop, the test server, or the production cloud. No more 'it works on my machine' problems.

The business benefits: **faster deployments**, **more reliable software**, **infrastructure flexibility** (ship to any cloud), and **lower infrastructure costs** (run more apps on the same hardware)."

#### In-depth
The shipping container analogy resonates well because it has real historical impact — standardized containers in the 1950s reduced shipping costs by 90% and transformed global trade. The parallel is intentional in Docker's naming and branding. For technical audiences, follow up with: Docker is to servers what VMs were to physical machines — better density, faster provisioning — but with containers, you're isolating processes, not entire operating systems.

---
