# ⚙️ Docker Daemon Production Tuning — `daemon.json` (Product-Based Companies)

This document covers Docker daemon configuration for production environments — a topic that bridges DevOps, SRE, and platform engineering interviews. While less common in pure coding rounds, it is frequently discussed in system design and infrastructure interviews.

---

### Q1: What is `daemon.json` and what are the most critical production settings in it?

**Answer:**
`/etc/docker/daemon.json` is the primary configuration file for the Docker daemon (`dockerd`). Changes here affect **all containers** on the host and require a daemon restart to take effect (`systemctl restart docker`).

**Full production-ready `daemon.json` with explanations:**
```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "50m",
    "max-file": "5",
    "compress": "true"
  },

  "storage-driver": "overlay2",

  "data-root": "/mnt/docker-data",

  "live-restore": true,

  "default-ulimits": {
    "nofile": {
      "Name": "nofile",
      "Hard": 65536,
      "Soft": 65536
    }
  },

  "max-concurrent-downloads": 10,
  "max-concurrent-uploads": 5,

  "registry-mirrors": [
    "https://mirror.company.com"
  ],

  "insecure-registries": [],

  "dns": ["8.8.8.8", "8.8.4.4"],

  "default-address-pools": [
    { "base": "172.80.0.0/16", "size": 24 }
  ],

  "experimental": false,

  "metrics-addr": "0.0.0.0:9323",
  "experimental": false,

  "features": {
    "buildkit": true
  },

  "no-new-privileges": true
}
```

---

### Q2: What is `live-restore` and why is it critical in production?

**Answer:**
```json
{ "live-restore": true }
```

**Without `live-restore`:** When you restart `dockerd` (e.g., to apply a config change, upgrade Docker, or after a crash), **all running containers are killed**. This means a daemon restart = application downtime.

**With `live-restore: true`:** The Docker daemon can be restarted or upgraded while containers **continue running** uninterrupted. The daemon re-attaches to the existing container processes (via the containerd shim) when it comes back up.

**How it works internally:**
- Containers run under `containerd-shim` processes, which are independent of the daemon.
- When `dockerd` restarts with `live-restore`, it re-reads the container state from containerd and re-attaches stdio, event streams, and health checks.

**Limitations:**
- Network changes during daemon downtime (e.g., iptables rules) may not be re-applied correctly.
- Cannot be combined with `--cluster-store` options (Docker Swarm HA mode).
- Docker stats/exec are unavailable while daemon is down.

**Why you care:** In bare-metal or VM deployments where Docker CE manages production containers directly (not Kubernetes), a `dockerd` crash or version upgrade without `live-restore` causes customer-facing downtime.

---

### Q3: What are `default-ulimits` and why do you need to configure them for high-throughput services?

**Answer:**
`ulimits` (user limits) are per-process OS resource limits. The most important for Docker workloads:

- **`nofile`** — Maximum open file descriptors (sockets, files, pipes)
- **`nproc`** — Maximum number of processes/threads
- **`memlock`** — Maximum locked memory (needed for Elasticsearch, some JVM apps)

**The default Linux `nofile` limit is 1024** — far too low for:
- High-concurrency web servers (each request = 1 socket = 1 fd)
- Database connection pools
- Message brokers (Kafka, RabbitMQ)

**Configure globally in `daemon.json`:**
```json
{
  "default-ulimits": {
    "nofile": {
      "Name": "nofile",
      "Hard": 65536,
      "Soft": 65536
    },
    "nproc": {
      "Name": "nproc",
      "Hard": 32768,
      "Soft": 32768
    }
  }
}
```

**Override per container:**
```bash
docker run --ulimit nofile=131072:131072 elasticsearch:8

# In Compose
services:
  kafka:
    image: confluentinc/cp-kafka:latest
    ulimits:
      nofile:
        soft: 65536
        hard: 65536
```

**Debugging ulimit issues:**
```bash
docker exec <container_id> ulimit -a
# Check: "open files" value

# Container hitting ulimit:
docker exec <container_id> sh -c 'ls /proc/1/fd | wc -l'
# Compare against: ulimit -n inside container
```

---

### Q4: What is the `default-address-pools` setting and why does it matter in large-scale deployments?

**Answer:**
Every Docker user-defined network gets assigned a **subnet** from which container IPs are allocated. By default, Docker picks subnets from `172.17.0.0/16` through `172.31.0.0/16`.

**The problem at scale:**
- A host running 20+ Compose stacks or many Docker networks exhausts the default pool.
- The default subnets may **conflict** with your company's VPN, on-premises network, or cloud VPC CIDR ranges (e.g., `172.20.0.0/16` is used by corporate VPN → containers can't reach VPN-connected hosts).

**Configure a custom pool:**
```json
{
  "default-address-pools": [
    { "base": "10.200.0.0/16", "size": 24 },
    { "base": "10.201.0.0/16", "size": 24 }
  ]
}
```

This allocates each new Docker network a `/24` subnet from the `10.200.x.x` and `10.201.x.x` ranges, avoiding conflicts with corporate infrastructure.

**Diagnose network exhaustion:**
```bash
# See all docker networks and their subnets
docker network ls -q | xargs docker network inspect --format '{{.Name}}: {{range .IPAM.Config}}{{.Subnet}}{{end}}'

# If you hit "could not find an available, non-overlapping IPv4 address pool":
docker network prune   # Remove unused networks
# Or reconfigure default-address-pools
```

---

### Q5: How do you configure Docker to use a registry mirror, and why is this important in a corporate or bandwidth-constrained environment?

**Answer:**
**The problem:** Every CI runner pulling `node:20-alpine` from Docker Hub downloads the same layers repeatedly. At scale (50 runners × 200MB image × 100 builds/day), this is:
- **Expensive** (egress bandwidth costs)
- **Slow** (download is the bottleneck)
- **Rate-limited** (Docker Hub limits unauthenticated pulls to 100/6h per IP; authenticated to 200/6h)

**Solution: Registry mirror (pull-through cache):**
A registry mirror is a local proxy that caches Docker Hub images. On first pull, it fetches from Docker Hub and caches. Subsequent pulls from any host hit the local cache.

**Configure in `daemon.json`:**
```json
{
  "registry-mirrors": [
    "https://docker-mirror.company.internal"
  ]
}
```

**Every `docker pull nginx:latest`** now transparently hits `docker-mirror.company.internal` first. If cached → instant. If not → fetched from Docker Hub once, then cached for future pulls.

**Popular mirror implementations:**
| Tool | Notes |
|---|---|
| **Harbor** | Full-featured private registry with pull-through cache proxy |
| **Sonatype Nexus** | Enterprise artifact manager (supports Docker proxy repos) |
| **JFrog Artifactory** | Same category as Nexus |
| **docker/distribution** | Open-source, lightweight pull-through cache |
| **AWS ECR Pull Through Cache** | Native AWS feature for caching Docker Hub/ECR Public |

**Important:** Mirrors only cache public images. Private image pulls still go to the original registry with credentials.

---

### Q6: What is `no-new-privileges` and how does it harden running containers at the daemon level?

**Answer:**
```json
{ "no-new-privileges": true }
```

**Background — the `setuid` problem:**
Linux programs can be `setuid` — when executed, they temporarily gain the file owner's privilege level (e.g., `/usr/bin/sudo` is owned by root, so running `sudo` gives your process temporary root). Inside a container, a non-root user could exploit `setuid` binaries to escalate to root.

**`no-new-privileges`** sets the `PR_SET_NO_NEW_PRIVS` prctl on the container's init process:
- A process **cannot gain more privileges** via `execve()` calls, regardless of `setuid`/`setgid` bits on executables.
- Even if a container has a vulnerable `setuid` binary (e.g., old `pkexec` — Polkit CVE-2021-4034), it cannot be exploited to escalate.

**Setting in `daemon.json` (applies globally):**
```json
{ "no-new-privileges": true }
```

**Setting per container:**
```bash
docker run --security-opt no-new-privileges:true myapp:latest
```

**Setting in Kubernetes:**
```yaml
securityContext:
  allowPrivilegeEscalation: false   # Same effect as no-new-privileges
```

**Interview note:** This is frequently paired with `runAsNonRoot: true` and `readOnlyRootFilesystem: true` as the "security triad" for hardened containers.

---

### Q7: How do you configure the Docker daemon for better observability with Prometheus metrics?

**Answer:**
Docker exposes a native Prometheus metrics endpoint when configured:

**Enable in `daemon.json`:**
```json
{
  "metrics-addr": "0.0.0.0:9323",
  "experimental": false
}
```

**Available metrics:**
- `engine_daemon_container_running_state_code` — Running/stopped/paused counts
- `engine_daemon_image_actions_seconds_*` — Image pull/push latency histograms
- `engine_daemon_network_actions_seconds_*` — Network create/delete latency
- `go_goroutines` — Concurrent goroutines in the daemon (monitor for leaks)
- `go_memstats_alloc_bytes` — Daemon memory usage

**Add to Prometheus scrape config:**
```yaml
scrape_configs:
  - job_name: 'docker-daemon'
    static_configs:
      - targets: ['docker-host:9323']
```

**cAdvisor for container-level metrics:**
```bash
# Run cAdvisor alongside Docker — exposes per-container CPU/RAM/network metrics
docker run -d \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:ro \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --publish=8080:8080 \
  --name=cadvisor \
  gcr.io/cadvisor/cadvisor:latest
```

**Full observability stack:**
```
Docker Daemon (port 9323) → Prometheus → Grafana
cAdvisor (port 8080)      → Prometheus → Grafana Dashboard (template ID 193)
```

---

### Q8: How do you safely apply `daemon.json` changes in production without downtime?

**Answer:**
Restarting `dockerd` kills all containers unless `live-restore: true` is set. Here's the safe procedure:

**Step 1: Verify `live-restore` is already enabled**
```bash
docker info | grep "Live Restore"
# Should show: Live Restore Enabled: true
```

**Step 2: Validate `daemon.json` syntax before applying**
```bash
# Parse and validate JSON syntax
python3 -m json.tool /etc/docker/daemon.json

# Or install dockerd itself to validate
dockerd --validate --config-file=/etc/docker/daemon.json
```

**Step 3: Take a snapshot of running containers**
```bash
docker ps --format '{{.Names}}\t{{.Status}}\t{{.Ports}}' > /tmp/containers_before.txt
```

**Step 4: Apply and reload**
```bash
# For some settings: SIGHUP is enough (no full restart needed)
kill -HUP $(pidof dockerd)

# For most settings: full restart required
systemctl restart docker
```

**Step 5: Verify containers are still running**
```bash
docker ps --format '{{.Names}}\t{{.Status}}' > /tmp/containers_after.txt
diff /tmp/containers_before.txt /tmp/containers_after.txt
# Should show no diff if live-restore worked
```

**Step 6: Verify new config took effect**
```bash
docker info | grep -E "Storage Driver|Logging Driver|Live Restore|Default Runtime"
```

---

*Prepared for Platform Engineering, SRE, and Infrastructure Architecture interviews at product-based companies.*
