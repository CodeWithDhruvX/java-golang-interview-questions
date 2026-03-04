# 🔍 Docker Production Debugging — Advanced Techniques (Product-Based Companies)

This document focuses on real-world production debugging scenarios with Docker. These questions are asked in senior/staff-level interviews to assess hands-on operational experience.

---

### Q1: A containerized service is hanging — it's not responding to requests but the container is still running. Walk through your systematic debugging process.

**Answer:**
This is a "live debugging under pressure" scenario. A structured approach:

**Step 1: Check process state inside the container**
```bash
# Get container stats first
docker stats <container_id> --no-stream
# Look for: CPU at 100% (stuck loop), or 0% + high MEM (stuck on I/O wait)

# Get a shell inside the container (if shell available)
docker exec -it <container_id> sh

# Check running processes
ps aux
# Look for: zombie processes (Z state), processes in D (uninterruptible sleep — disk/network IO block)
```

**Step 2: Check the process state without entering the container via `/proc`**
```bash
# Find container's main PID on the host
docker inspect <container_id> --format '{{.State.Pid}}'
# e.g., returns 4521

# Check its status on the host
cat /proc/4521/status
# Look for: "State: D (disk sleep)" → blocked on IO
# Look for: "Threads: 1000+" → thread exhaustion

# Check its open file descriptors
ls -la /proc/4521/fd | wc -l
# Too many? → file descriptor leak
```

**Step 3: Attach `strace` to the running process**
```bash
# Without entering the container — from the host
strace -p 4521 -e trace=network,io -f

# With nsenter (enter the container's namespace from the host)
nsenter --target 4521 --pid --net --mount
# Now you're inside the container's namespaces AS ROOT, regardless of container user
strace -p 1 -e trace=network
```

**Step 4: Check network connections**
```bash
# Inside container or via nsenter
ss -tulnp       # socket stats
ss -s           # summary — look for CLOSE_WAIT flood (upstream not closing connections)

# Check if it can reach its dependencies
curl -v --max-time 5 http://database-service:5432
```

**Step 5: Get a thread dump (for JVM) or goroutine dump (Go)**
```bash
# JVM: send SIGQUIT for thread dump
docker exec <id> kill -3 1

# Go: enable pprof and curl
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# Node.js
docker exec <id> kill -USR1 1   # triggers inspector
```

---

### Q2: Explain `nsenter` — what is it, and give three concrete use cases for production Docker debugging.

**Answer:**
`nsenter` is a Linux tool that **enters an existing process's namespaces** without going through the container runtime (no `docker exec` needed). It's invaluable when:
- The container image has no shell (distroless)
- `docker exec` is refusing due to unhealthy container state
- You need host-level tools inside the container's namespace

**Syntax:**
```bash
nsenter --target <PID> --pid --net --mount --uts --ipc [-- command]
```

**Use Case 1: Debug a distroless container with no shell**
```bash
# Container runs distroless/static:latest — no bash, no sh
PID=$(docker inspect myapp --format '{{.State.Pid}}')

# Enter its network namespace and use HOST tools
nsenter --target $PID --net -- \
  ss -tulnp    # Use host's ss, inside container's net namespace

nsenter --target $PID --mount --pid -- \
  ls /proc/1/fd   # List open file descriptors
```

**Use Case 2: Investigate DNS resolution issues from inside the container's namespace**
```bash
PID=$(docker inspect myapp --format '{{.State.Pid}}')

# Enter net namespace — uses container's /etc/resolv.conf view
nsenter --target $PID --net --mount -- \
  cat /etc/resolv.conf

nsenter --target $PID --net -- \
  nslookup database-service   # Uses container's DNS config
```

**Use Case 3: Run a performance profiler inside a container's namespace**
```bash
PID=$(docker inspect myapp --format '{{.State.Pid}}')

# Run perf inside the container's PID namespace
nsenter --target $PID --pid --net --mount -- \
  perf top -p 1   # Profile PID 1 (container's main process)
```

**Key advantage over `docker exec`:**
`nsenter` runs with **host root privileges** regardless of the container's user restrictions. It's a recovery tool — use it when normal debugging paths are unavailable.

---

### Q3: How do you investigate a Docker container that is consuming unexpectedly large amounts of disk space?

**Answer:**
This is a common production incident — a container fills the node's disk causing OOM or write failures.

**Step 1: Identify what's consuming space**
```bash
# Overall Docker disk usage
docker system df -v

# Output shows:
# Images:     total, active, size, reclaimable
# Containers: total, running, size, reclaimable
# Volumes:    total, linked, size, reclaimable
# Build cache: total, active, size, reclaimable
```

**Step 2: Check the container's writable layer**
```bash
docker inspect <container_id> --format '{{.GraphDriver.Data}}'
# Returns: MergedDir, UpperDir, WorkDir, LowerDir paths

# Check size of the writable (upper) layer
du -sh /var/lib/docker/overlay2/<layer-id>/diff/

# Inside container — find large files
docker exec <id> find / -xdev -type f -size +100M 2>/dev/null
# -xdev: don't cross filesystem boundaries (won't scan volumes)
```

**Step 3: Common culprits and fixes**

| Root Cause | Detection | Fix |
|---|---|---|
| Application writing logs to container filesystem | `find / -name "*.log" -size +1M` | Use logging driver (fluentd, json-file with rotation), or write to a mounted volume |
| Core dumps accumulating | `find / -name "core*"` inside container | Configure `ulimit -c 0` or redirect to a monitored directory |
| Temp files not cleaned up | `du -sh /tmp` inside container | Add tmpfs mount for `/tmp` |
| Unoptimized Docker build cache | `docker system df` | `docker builder prune --filter 'until=24h'` |
| Old stopped container layers | `docker ps -a` | `docker container prune` |
| Dangling images | `docker images -f dangling=true` | `docker image prune` |

**Step 4: Prevent log file growth**
Configure log rotation in `/etc/docker/daemon.json`:
```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "50m",
    "max-file": "3"
  }
}
```

---

### Q4: A container that passed all load tests is throwing `connection refused` errors in production but not in staging. How do you debug this?

**Answer:**
"Works in staging, fails in production" is a classic network/configuration mismatch problem.

**Hypothesis-driven debugging approach:**

**Hypothesis 1: Network policy or firewall difference**
```bash
# From inside the container, test connectivity to the target
docker exec <id> nc -zv target-service 5432
# If "connection refused" → port is closed/filtered
# If "connection timed out" → firewall is silently dropping (different from refused)

# Test from prod vs staging
nsenter --target <PID> --net -- traceroute target-service
```

**Hypothesis 2: DNS resolution difference**
```bash
# Container may resolve to different IP in prod
docker exec <id> nslookup target-service
docker exec <id> cat /etc/resolv.conf
# Prod may use a different search domain or DNS server
```

**Hypothesis 3: Environment variable misconfiguration**
```bash
docker inspect <prod_container> --format '{{json .Config.Env}}' | jq
docker inspect <staging_container> --format '{{json .Config.Env}}' | jq
# Diff these — wrong DATABASE_URL, SERVICE_HOST pointing to staging endpoint?
```

**Hypothesis 4: Connection pool exhaustion**
```bash
# If prod gets 10x traffic, the pool may be legitimately exhausted
docker exec <id> sh -c 'ss -s'
# Look for: high CLOSE_WAIT count → upstream not acknowledging close (TCP half-close bug)
# Or:       high TIME_WAIT → too many short-lived connections (needs connection pooling or keepalive)
```

**Hypothesis 5: TLS certificate mismatch**
```bash
# Staging might use self-signed cert; prod uses real CA
docker exec <id> openssl s_client -connect target-service:443 -servername target-service
# Look for certificate chain errors
```

**Hypothesis 6: Resource limits**
```bash
# Prod might have stricter ulimits
docker inspect <id> --format '{{json .HostConfig.Ulimits}}'
# Too-low nofile limit → can't open new sockets 
```

---

### Q5: How would you live-debug memory growth in a running container without restarting it?

**Answer:**
Restarting to fix a memory leak destroys the evidence. Live debugging is essential at scale.

**Step 1: Confirm memory growth trend**
```bash
# Watch memory stats in real time
docker stats <container_id>
# Look for: steadily climbing MEM USAGE over time = leak

# More detail from host
PID=$(docker inspect <id> --format '{{.State.Pid}}')
cat /proc/$PID/status | grep -E "VmRSS|VmPeak|VmSwap"
```

**Step 2: JVM applications (heap dump)**
```bash
# Get process PID inside container
docker exec <id> jps

# Trigger heap dump (writes to container filesystem)
docker exec <id> jmap -dump:format=b,file=/tmp/heap.hprof <app_pid>

# Copy dump to host for analysis
docker cp <id>:/tmp/heap.hprof ./heap.hprof

# Analyze with Eclipse MAT or:
docker run --rm -v $(pwd):/dumps eclipse-mat /dumps/heap.hprof
```

**Step 3: Go applications (pprof)**
```bash
# If pprof endpoint is enabled (:6060)
curl http://<container_ip>:6060/debug/pprof/heap > heap.out
go tool pprof -http=:8080 heap.out   # Interactive visual profiler
```

**Step 4: Native/C applications**
```bash
# Use valgrind inside the container
docker exec <id> valgrind --leak-check=full /app/binary

# Or attach GDB using gcore for a memory snapshot
nsenter --target $PID --pid -- gcore $PID
# Analyze the core file
```

**Step 5: Check kernel perspective**
```bash
# smaps shows detailed memory breakdown per mapping
cat /proc/$PID/smaps | grep -E "^(Anonymous|Swap):" | awk '{sum += $2} END {print sum " kB"}'
```

**Step 6: Container-level memory limit debugging**
```bash
# Check if process was OOM killed previously (without restarting)
docker inspect <id> --format '{{.State.OOMKilled}}'

# Check kernel's oom score
cat /proc/$PID/oom_score
cat /proc/$PID/oom_score_adj
```

---

### Q6: Your team wants to implement continuous profiling of Docker containers in production. What tools and approaches would you recommend?

**Answer:**
Continuous profiling (always-on, low-overhead profiling) is essential for catching performance regressions before users report them.

**OSS Tooling:**

**1. Parca / Parca Agent (eBPF-based)**
```bash
# Parca Agent runs on each node and uses eBPF to sample ALL container processes
# Zero instrumentation required — works on any language
docker run --privileged \
  -v /proc:/host/proc:ro \
  -v /sys:/sys:ro \
  --pid=host \
  parca/parca-agent:latest \
  --remote-store-address=parca.company.com:7070
```
- Collects CPU flame graphs from running containers using `perf_event_open` + eBPF
- Zero code changes required
- Overhead: ~1-3%

**2. Pyroscope**
```bash
# Application-side SDK (for language runtimes)
# Or eBPF agent (same as Parca, no instrumentation)
docker run -d \
  -e PYROSCOPE_ADHOC_DATA_PATH=/tmp/pyroscope \
  -v /proc:/proc:ro \
  --pid=host --privileged \
  grafana/pyroscope:latest
```

**3. Grafana Alloy + Beyla (for HTTP services)**
```bash
# Beyla uses eBPF to instrument HTTP/gRPC automatically
# Gets request latency p95/p99 without code modification
docker run --privileged \
  --pid=host \
  -e BEYLA_SERVICE_NAME=myapp \
  -e OTEL_EXPORTER_OTLP_ENDPOINT=http://otel-collector:4318 \
  grafana/beyla:latest
```

**Architecture pattern for production:**
```
[Parca/Pyroscope eBPF Agent] ← runs on each K8s/Docker node
    ↓ streams profiles via OTLP or Parca protocol
[Central Profile Store] ← Parca server or Grafana Pyroscope OSS
    ↓ query via web UI or Grafana datasource
[Engineering] ← drills into function-level flame graphs for any time range
```

**When to use what:**
| Tool | Best For | Overhead |
|---|---|---|
| Parca Agent | Multi-language, no instrumentation | ~1-3% CPU |
| Pyroscope SDK | Deep language integration (Go, Python, JVM) | ~2-5% |
| Beyla | HTTP/gRPC latency without code changes | ~1-2% |
| cAdvisor + Prometheus | Container resource metrics (not profiles) | Minimal |

---

*Prepared for Senior/Staff Software Engineering and Platform Engineering interviews at product-based companies.*
