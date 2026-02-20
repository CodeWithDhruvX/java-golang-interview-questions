# 📌 **Miscellaneous / Real-World Scenarios (191–200)**

---

### 191. How do you identify which containers are consuming the most CPU?
"Use `docker stats` sorted by CPU:

```bash
# Live view, sorted by CPU (re-sort manually)
docker stats --format 'table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}'

# One-shot, sort by CPU descending
docker stats --no-stream --format '{{.Name}}\t{{.CPUPerc}}' | sort -t$'\t' -k2 -rn | head -10
```

For automated alerting: cAdvisor exports `container_cpu_usage_seconds_total` to Prometheus — alert when the rate exceeds the container's CPU limit for more than 5 minutes.

`docker exec highcpu-container top` — see which process inside is the cause."

#### In-depth
`docker stats` CPU% is computed as a ratio of the container's CPU usage to all available CPU time across all cores. On a 4-core machine, 400% CPU means all cores saturated. When troubleshooting high CPU: check if it's expected load growth (scale horizontally), a CPU-bound bug (profiling needed), or a runaway process (infinite loop). Go's pprof, Java's async-profiler, or Node.js `--inspect` with Chrome DevTools are my go-to in-container profilers.

---

### 192. What do you do when a container uses 100% disk space?
"Systematic diagnosis:

1. `docker system df` — see Docker's total disk usage (images, containers, volumes, build cache)
2. `docker system df -v` — per-image and per-volume detail
3. `docker exec container df -h` — check disk usage inside the container
4. `docker exec container du -sh /*` — find which directory is large
5. Check logs: `docker exec container du -sh /var/log`

Common causes: runaway logs (set `--log-opt max-size=10m max-file=3`), leftover temp files from app bugs, build cache growth.

Fix: `docker system prune` for Docker artifacts; fix the app for runtime disk usage."

#### In-depth
Container log files are the most common disk exhaustion cause. Docker's `json-file` driver stores logs in `/var/lib/docker/containers/<id>/<id>-json.log`. Without limits, a verbose container can fill the disk in hours. Always set log rotation: in `daemon.json`: `{"log-opts": {"max-size": "10m", "max-file": "3"}}`. This limits logging to 30MB per container. For production, use a centralized logging driver (fluentd, awslogs) that ships logs off the host entirely.

---

### 193. How do you roll out updates to a fleet of containers?
"For Swarm: `docker service update --image myapp:v2 --update-parallelism 2 --update-delay 30s myservice`.

This updates 2 tasks at a time with 30 seconds between batches. Swarm health-checks each new task before proceeding to the next batch.

For bare Docker (no orchestration): rolling update script:
```bash
for host in $(cat hosts.txt); do
  ssh $host "docker pull myapp:v2 && docker stop myapp && docker rm myapp && docker run -d --name myapp myapp:v2"
  sleep 10
  curl -f http://$host/health || echo "FAILED on $host"
done
```

For Kubernetes: `kubectl rollout` with configurable `maxSurge` and `maxUnavailable`."

#### In-depth
The key to safe fleet rollouts: **canary deployments**. Update 1-2% of instances first, monitor error rates and p99 latency for 15-30 minutes, then proceed to the rest. This catches issues that weren't visible in staging: production-specific load patterns, data issues, and third-party API behavior. Tools like Argo Rollouts (Kubernetes) automate canary analysis — automatically promoting or rolling back based on Prometheus metrics.

---

### 194. How do you run scheduled jobs in Docker (cron)?
"Several approaches:

**1. Dedicated cron container**:
```dockerfile
FROM alpine
RUN apk add --no-cache curl
COPY crontab /etc/crontabs/root
CMD ["crond", "-f", "-d", "8"]
```

**2. Docker API trigger (preferred)**: keep cron on the host or in a scheduler container that uses the Docker API to spin up ephemeral containers: `docker run --rm myapp ./run-job.sh`.

**3. Kubernetes CronJobs**: if on K8s, use the native CronJob resource.

**4. Tools**: **Ofelia** — a cron-like scheduler that runs Docker containers on a schedule, defined in Compose labels."

#### In-depth
The ephemeral container approach (`docker run --rm`) is significantly better than a long-running cron container: each job run starts fresh (no state pollution from previous runs), resource usage is bounded (container exits after job completes), and failures are clearly logged (exit code, `docker ps -a` shows stopped containers). The host cron entry: `0 2 * * * docker run --rm myregistry/myapp:latest ./nightly-job.sh`. One concern: image pull latency — pre-pull the image on the schedule that triggers before the job.

---

### 195. How do you detect and avoid container zombie processes?
"Zombies occur when a child process exits but its parent hasn't called `wait()` on it. The process stays in the process table as a zombie (defunct).

**Detection**: `docker exec container ps aux | grep '<defunct>'`

**Prevention**:
1. **Use `--init` flag**: `docker run --init myapp` — tini becomes PID 1 and reaps zombies
2. **Use `dumb-init`**: `ENTRYPOINT ["/usr/bin/dumb-init", "--"]`
3. **Fix the parent**: if your app forks children, implement proper `wait()` or `SIGCHLD` handlers
4. **Use exec form** for CMD/ENTRYPOINT: prevents shell from being PID 1"

#### In-depth
The zombie accumulation problem is slow — it takes many zombie PIDs to exhaust the PID namespace. But in containers with low `ulimit -u` (PID limits), it can happen faster. `tini` (used by `--init`) is 2KB and handles reaping correctly. It also properly forwards signals to the child process — a secondary benefit. Docker's official `--init` implementation uses `tini` specifically because it's battle-tested, minimal, and correct.

---

### 196. How do you perform blue-green deployments with Docker?
"Blue-green uses two identical environments. Only one is live (receives traffic) at a time.

**With Nginx or Traefik as LB**:
1. Blue is live (`app-blue:latest`)
2. Deploy green: `docker run -d --name app-green --network internal myapp:v2`
3. Health check green until passing
4. Switch LB config from blue to green (atomic nginx reload or Traefik label update)
5. Blue stays running for instant rollback
6. After verifying green: remove blue

**With Swarm**: use service aliases — `docker service update --network-add alias=frontend myservice-green`."

#### In-depth
Blue-green requires 2x capacity during the transition, but provides: **instant rollback** (switch LB back to blue), **no mixed versions** (all traffic goes to one version — unlike rolling), and **safe database migrations** (run migration against green, test, then cutover). The database migration challenge: use backward-compatible schema changes (expand-contract pattern) so both blue and green can operate against the same database during the transition.

---

### 197. How do you restart failed containers automatically?
"Use Docker's **restart policy** at container creation:

```bash
docker run -d --restart=on-failure:5 myapp
# or
docker run -d --restart=always myapp
# or
docker run -d --restart=unless-stopped myapp
```

In Compose:
```yaml
services:
  api:
    restart: on-failure
```

In Swarm, services auto-restart failed tasks by design — replace a failed task on a healthy node. Configure: `--restart-condition any --restart-delay 5s --restart-max-attempts 3`."

#### In-depth
`on-failure` with a max attempt count is my preferred policy for non-critical services — it prevents infinite crash loops from obscuring a real bug. After `N` failures, the container stays stopped and alerts fire (assuming you monitor container states). `unless-stopped` is for production services that should always run — it survives host reboots while respecting manual `docker stop`. The exponential backoff built into Docker's restart prevents thundering herd after a transient failure.

---

### 198. How do you isolate builds between microservices using Docker?
"Each microservice has its own Dockerfile and builds to its own image — they're naturally isolated.

For deeper isolation in CI:
- Each service runs `docker build` in its own CI job with its own ephemeral runner
- No shared volumes between builds (each runner is clean)
- Separate image tags: `registry/service-a:sha1`, `registry/service-b:sha2`
- Use **Docker Buildx bake** for multi-service builds: `docker buildx bake` reads `docker-bake.hcl` and builds multiple targets in parallel

In monorepo CI: use path filters to trigger only affected service builds."

#### In-depth
`docker buildx bake` is the modern way to define multi-service builds. A `docker-bake.hcl` defines all service targets with their build contexts, tags, and cache settings. Run `docker buildx bake --print` to see the resolved build plan. Run `docker buildx bake` to build all services in parallel with shared BuildKit cache. It's particularly powerful in monorepos: `docker buildx bake --set *.cache-to=type=registry,ref=...` configures registry-based cache for all services in one command.

---

### 199. How do you implement canary releases using Docker?
"Canary releases route a small percentage of traffic to a new version to validate it before full deployment.

**With Nginx**:
```nginx
upstream backend {
    server app-stable:8080 weight=9;  # 90% traffic
    server app-canary:8080 weight=1;  # 10% traffic
}
```

**With Traefik** (Compose labels):
```yaml
- "traefik.http.services.canary.loadbalancer.sticky.cookie.name=canary"
- "traefik.http.middlewares.canary.ratelimit.average=10"
```

**With Swarm**: deploy new image as a separate service with low replica count alongside the main service on the same network. The load balancer splits traffic by replica count."

#### In-depth
True canary releases require traffic splitting + **metric comparison**. Simply routing 10% to v2 isn't enough — you must compare error rates, P99 latency, and business metrics (conversion rate) between canary and stable. Tools like **Argo Rollouts** (Kubernetes) automate this: define a metric threshold (`errorRate < 1%`), and Rollouts automatically promotes (increase canary traffic) or aborts (roll back canary). Manual canary with Nginx weights is fast to set up but requires human judgment on the metrics.

---

### 200. How do you document Docker workflows for your team?
"Documentation that actually gets used:

1. **`README.md` in each service**: quick start (`docker compose up`), environment variable reference, local development guide
2. **Annotated Dockerfiles**: comments explaining non-obvious choices (`# Alpine chosen for 3MB size; ensure musl compatibility`)
3. **Annotated `docker-compose.yml`**: comments for each service, why each env var exists
4. **Runbooks**: step-by-step for common ops tasks (deploy, rollback, debug a failing container, backup DB)
5. **Architecture diagrams**: use Mermaid in markdown to show service topology and network connections
6. **`.env.example`**: documented template with dummy values, committed to the repo"

#### In-depth
The most valuable Docker documentation is **`docker compose --help` supplemented runbooks** — scenario-specific guides like 'How to debug a failing container in production' or 'How to update an image in production'. Structured as: context → step-by-step commands → expected output → troubleshooting if it fails. Pair with demo recordings (Asciinema or similar) — watching someone `docker exec -it` into a container to diagnose an issue teaches more than reading about it. Store runbooks alongside code, not in a separate wiki.

---
