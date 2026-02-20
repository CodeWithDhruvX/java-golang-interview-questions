# 🔧 **Docker Production Debugging (381–410)**

---

### 381. How do you debug a container that crashes immediately on start?
"Immediate crash debugging approach:

```bash
# 1. Check exit code and last logs
docker ps -a  # See exit code
docker logs container_name

# 2. Override entrypoint to get a shell
docker run -it --entrypoint /bin/sh myimage

# 3. Check if binary exists and is executable
docker run --entrypoint ls myimage -la /app/

# 4. Run with strace to see syscalls
docker run --cap-add SYS_PTRACE myimage strace -e trace=all /app/server

# 5. Check for missing shared libraries
docker run --entrypoint ldd myimage /app/server
```"

#### In-depth
Exit code analysis: code 1 = application error, code 126 = permission denied (binary not executable), code 127 = command not found (binary missing or wrong path), code 137 = SIGKILL (OOM killed or `docker kill`), code 139 = SIGSEGV (segfault), code 143 = SIGTERM. Each tells a different story. Code 127 in a multi-stage build usually means the binary was copied from the wrong stage (`COPY --from=builder`) or to the wrong path. Always verify with `docker run --entrypoint ls`.

---

### 382. How do you investigate OOM (Out of Memory) killed containers?
"OOM (Out of Memory) kills happen when a container exceeds its memory limit.

**Detection**:
```bash
# Check for OOM kills in recent events
docker events --filter event=oom
docker inspect container_name | jq '.[0].State.OOMKilled'
# Returns true if OOM killed

# Check kernel OOM log
dmesg | grep -E 'oom|killed'
```

**Analysis**:
```bash
# Current memory usage
docker stats --no-stream container_name

# Historical: check if limit was too low
docker inspect container_name | jq '.[0].HostConfig.Memory'
```

**Fix**: increase limit or profile memory usage to find leaks."

#### In-depth
OOM kills in production are silent unless you monitor for them. Add a Prometheus alert: `rate(container_oom_events_total[5m]) > 0` — any OOM event pages the on-call team. Before simply increasing limits, profile the memory: `docker exec container jmap -histo:live <java_pid>` for JVM, `valgrind --tool=massif` for C/C++, or memory profiler middleware for Python/Node. OOM kills are often memory leaks that grow over days. Time-series memory metrics (Prometheus cAdvisor) reveal the growth pattern.

---

### 383. How do you analyze Docker container CPU throttling?
"CPU throttling happens when a container hits its CPU limit (cgroup quota) — processes are paused.

**Detect throttling**:
```bash
# Check cgroup CPU stats
cat /sys/fs/cgroup/cpu/docker/<container_id>/cpu.stat
# nr_throttled: count of times throttled
# throttled_time: nanoseconds spent throttled

# Via Prometheus (cAdvisor):
# container_cpu_cfs_throttled_periods_total
# container_cpu_cfs_periods_total
# Throttle rate = throttled / total periods
```

**Fix options**:
1. Increase CPU limit: `docker update --cpus=2.0 container`
2. Optimize CPU-intensive code
3. Scale horizontally (more replicas)"

#### In-depth
CPU throttling is invisible without metrics but causes high latency — a request that normally takes 10ms takes 100ms when the container is throttled 90% of the time. The throttle rate formula: `throttled_periods / total_periods`. >25% throttle rate is a problem. A common surprise: a container limited to 0.5 CPU that's fine under normal load gets throttled under bursts (nightly batch jobs, traffic spikes). Solution: use CPU `requests` (guaranteed reservation) and higher `limits` to allow bursting without constant throttling.

---

### 384. How do you perform a heap dump from a running Java container?
"Heap dump extraction from a running Java container:

```bash
# Find Java PID inside container
docker exec mycontainer jps

# Generate heap dump to container filesystem
docker exec mycontainer jmap -dump:format=b,file=/tmp/heapdump.hprof <pid>

# Copy out of container
docker cp mycontainer:/tmp/heapdump.hprof ./heapdump.hprof

# Analyze locally with Eclipse MAT, JProfiler, or jhat
jhat -port 7000 ./heapdump.hprof
```

For OOM-triggered dumps: add to JVM: `-XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/dumps/`
Then mount a volume at `/dumps/` to retrieve dumps after OOM kills."

#### In-depth
`-XX:+HeapDumpOnOutOfMemoryError` is essential for production JVM containers. Without it, the container OOM-kills and you have no evidence of what caused it. Mount a persistent volume at `/dumps/` — the dump persists after the container is restarted. In Kubernetes: use an emptyDir volume shared with a debug sidecar that compresses and uploads dumps to S3. For containers that are distroless (no `jmap`): add the `jattach` binary to the image as a debug tool, or use async-profiler's `jattach` to trigger heap dumps remotely.

---

### 385. How do you trace system calls in Docker containers?
"System call tracing reveals exactly what a container process is doing at the kernel level.

```bash
# Attach strace to a running container process
docker exec container_name sh -c 'strace -p $(pidof myapp) -e trace=network,file'

# Or: strace from outside the container
docker_pid=$(docker inspect -f '{{.State.Pid}}' container)
strace -p $docker_pid -e trace=network

# For container startup (strace the entrypoint):
docker run --cap-add SYS_PTRACE --security-opt seccomp=unconfined \
  --entrypoint strace myimage -f -e trace=all /app/server
```

`-f` follows forks. `-e trace=network` shows only network syscalls."

#### In-depth
strace in containers requires `SYS_PTRACE` capability — not available by default. `--security-opt seccomp=unconfined` disables seccomp (which blocks ptrace). This is acceptable in debugging scenarios but never in production. For production debugging without relaxing security: use eBPF-based tools that don't require ptrace: `bpftrace -e 'tracepoint:syscalls:sys_enter_* /pid == $1/ { @[probe] = count(); }' -v $CONTAINER_PID`. This provides full syscall tracing with zero privilege escalation.

---

### 386. How do you debug a slow container application?
"Slow container debugging methodology:

**Step 1: Isolate the bottleneck tier**:
```bash
docker stats --no-stream  # See CPU, memory, I/O per container
```

**Step 2: Application-level profiling** (language-specific):
- Go: `curl localhost:6060/debug/pprof/profile > cpu.prof && go tool pprof cpu.prof`
- Node.js: `docker exec container node --prof app.js`, then `node --prof-process`
- Python: `docker exec container py-spy top --pid $(pidof python)`
- Java: `docker exec container async-profiler -e cpu -d 30 -o flamegraph <pid>`

**Step 3: Check external calls**: trace database queries, HTTP calls, cache hit rates."

#### In-depth
Flame graphs are the most effective visualization for CPU profiling — they show the call stack hierarchy and relative time spent. `py-spy` and `async-profiler` are particularly good because they're safe-mode profilers (no instrumentation, minimal overhead, work on production containers without code changes). For I/O-bound slowness: `iostat -x 1` on the host or `docker exec container cat /proc/<pid>/io` to check read/write bytes. Database query analysis: enable slow query logging in the DB container and analyze with `pt-query-digest`.

---

### 387. How do you find which process inside a container is consuming resources?
"Process-level resource analysis inside containers:

```bash
# CPU usage per process inside container
docker exec container top -b -n1

# Memory per process
docker exec container cat /proc/1/status | grep VmRSS
docker exec container ps aux --sort=-%mem

# File descriptors (common leak)
docker exec container ls -la /proc/1/fd | wc -l

# Open network connections
docker exec container ss -tunap

# I/O per process
docker exec container cat /proc/$(pidof myapp)/io
```

For advanced: `pidstat` from the `sysstat` package shows per-process CPU/IO over time."

#### In-depth
File descriptor leaks are a common production container issue — manifests as 'too many open files' errors. Each open socket, file, and pipe consumes an fd. `ls /proc/<pid>/fd | wc -l` shows current fd count. `cat /proc/sys/fs/file-max` shows system limit. Default limit is 1024 per process (check with `ulimit -n` inside container). Docker CLI allows: `docker run --ulimit nofile=65536:65536` to raise the limit. For production: monitor `process_open_fds` Prometheus metric — alert when it approaches the limit.

---

### 388. How do you debug a container that's in an infinite restart loop?
"CrashLoopBackOff debug process:

```bash
# 1. View logs from the crashed container (not current)
docker logs --tail=100 container_name  # May catch logs before crash

# For Kubernetes CrashLoopBackOff:
kubectl logs pod/myapp --previous  # Previous container's logs

# 2. Override entrypoint to prevent crash
docker run -it --entrypoint /bin/sh myimage
# Manually run the startup script to see what fails

# 3. Check exit code
docker inspect container --format='{{.State.ExitCode}}'

# 4. Increase start timeout in health check:
# start_period: 120s  -- give more time before failing
```"

#### In-depth
The infinite restart loop with no logs is the hardest to debug. Causes: binary segfaults before writing any logs, missing required file that crashes at startup before logger initializes, or strict health check that fails too quickly. Approach: override CMD/ENTRYPOINT with `sleep 3600` — the container stays alive. Then exec in and manually run the startup command to see the error interactively. For Kubernetes: `kubectl debug pod/myapp --copy-to=debug-pod --set-image=myapp=/bin/sleep -- 3600` creates a copy of the pod with the command replaced by sleep.

---

### 389. How do you use nsenter for container debugging?
"`nsenter` enters a process's Linux namespaces — useful for debugging containers from the host.

```bash
# Get container PID
CPID=$(docker inspect -f '{{.State.Pid}}' mycontainer)

# Enter all namespaces (like docker exec but works for paused containers)
nsenter --target $CPID --mount --uts --ipc --net --pid

# Enter only network namespace (inspect container networking from host)
nsenter --target $CPID --net ip addr
nsenter --target $CPID --net ss -tunap
nsenter --target $CPID --net tcpdump -i eth0

# Enter filesystem namespace
nsenter --target $CPID --mount ls /app/
```"

#### In-depth
`nsenter` is more powerful than `docker exec` for debugging because: it works on containers in any state (including paused or containers with no shell), it can enter namespaces individually (useful for network debugging without entering the PID namespace), and it runs with the host's tools (not limited to what's in the container). For security-hardened containers (no shell in distroless): `nsenter --mount` to inspect the filesystem, `nsenter --net` to check network from the host's perspective, `nsenter --pid` to see processes in the container's PID namespace.

---

### 390. How do you use `docker diff` for debugging?
"`docker diff` shows changes made to a container's filesystem since it started from the image.

```bash
docker diff mycontainer
# Output:
# A /tmp/cachefile     -- Added
# C /etc/app.conf     -- Changed
# D /var/log/old.log  -- Deleted
```

Use cases:
1. **Auditing**: see what a container modified (security audit)
2. **Debugging**: see if config files were correctly written on startup
3. **Image optimization**: see what's being created at runtime — should it be in the image instead?
4. **Mutable state detection**: identify containers that modify their filesystem (breaks immutability)"

#### In-depth
`docker diff` reveals all filesystem mutations from image to current state — it reads the container's writable layer in OverlayFS. For multi-process containers that write runtime files (PID files, cache files), this shows exactly what was created. For debugging configuration issues: if a config file shows `C` (Changed) but shouldn't be, something modified it at runtime. If expected config shows no `A` (Added), the entrypoint failed to create it. This is faster than exec'ing into the container to manually inspect files.

---

### 391. How do you capture and analyze network traffic from a Docker container?
"Packet capture on a container's network interface:

```bash
# Method 1: tcpdump inside the container
docker exec -it mycontainer tcpdump -i eth0 -w /tmp/capture.pcap
docker cp mycontainer:/tmp/capture.pcap ./
# Open in Wireshark

# Method 2: tcpdump from host via nsenter
CPID=$(docker inspect -f '{{.State.Pid}}' mycontainer)
nsenter --target $CPID --net tcpdump -i eth0 -w capture.pcap

# Method 3: netshoot sidecar
docker run --rm --net container:mycontainer \
  nicolaka/netshoot tcpdump -i eth0 port 5432
```"

#### In-depth
Packet capture is indispensable for debugging: TLS handshake failures (verify certificate chain), unexpected connection resets (RST packets: which side is resetting?), slow queries (time from query packet to response packet), and DNS failures (no DNS response, NXDOMAIN). Use Wireshark filters: `http.response.code >= 400` for HTTP errors, `tcp.flags.reset == 1` for connection resets, `dns.flags.rcode != 0` for DNS failures. For TLS-encrypted traffic: capture at the application layer using `SSLKEYLOGFILE` env var to capture TLS session keys, then decrypt in Wireshark.

---

### 392. How do you debug Docker volumes and data persistence issues?
"Volume debugging methodology:

```bash
# 1. Verify volume is mounted correctly
docker inspect container | jq '.[0].Mounts'
# Check Source (host path) and Destination (container path)

# 2. Check volume content from host
docker volume inspect myvol
ls $(docker volume inspect myvol --format '{{.Mountpoint}}')

# 3. Check filesystem permissions
docker exec container ls -la /app/data/

# 4. Check if volume is read-only
docker inspect container | jq '.[0].Mounts[].RW'
# false = read-only mount

# 5. Test write access
docker exec container touch /app/data/testfile
```"

#### In-depth
The most common volume issue: permission mismatch. A container running as UID 1000 can't write to a directory owned by root (UID 0). Check: `docker exec container stat /app/data/` — shows owner UID. Fix options: (1) `chown` in the Dockerfile's USER section, (2) init container that `chown`s the directory, (3) `fsGroup` in Kubernetes securityContext (sets group ownership of volume). For NFS volumes: check NFS export options (`no_root_squash` for containers that need root access to NFS), client mount options, and network connectivity between node and NFS server.

---

### 393. How do you recover data from a failed Docker container?
"Data recovery from a crashed/deleted container:

```bash
# Container still exists (even stopped):
docker cp failed_container:/app/data ./recovered_data

# Container deleted but volume persists:
docker volume ls  # Find orphaned volume
docker run --rm -v orphaned_vol:/data busybox tar czf /backup.tar.gz /data
docker cp <temp_container_id>:/backup.tar.gz .

# OverlayFS layer recovery (advanced):
# Containers' writable layers are in:
ls /var/lib/docker/overlay2/
# Find the container's layer by ID:
docker inspect failed_container | jq '.[0].GraphDriver.Data'
# Mount and copy files directly
```"

#### In-depth
Data recovery prevention is better than recovery: always use named volumes (`--mount type=volume`), never store critical data in the container's writable layer. OverlayFS layer recovery is a last resort — the writable layer exists until `docker rm` is called. If a container is stopped but not removed, all data in its writable layer is accessible via `docker cp`. After `docker rm`, the OverlayFS layer is deleted and unrecoverable. `docker run --rm` auto-deletes after exit — never use for containers with important state.

---

### 394. How do you diagnose Docker daemon issues?
"Docker daemon debugging:

```bash
# Check daemon status and recent logs
systemctl status docker
journalctl -u docker --since '1 hour ago' -f

# Increase daemon log level temporarily
dockerd --log-level debug  # Or set in daemon.json: {"log-level": "debug"}

# Daemon API connectivity
curl --unix-socket /var/run/docker.sock http://localhost/version

# Check daemon configuration
docker info
docker system info | grep -E 'Storage|Runtime|Warnings'

# Check for daemon panics/crashes
journalctl -u docker | grep -E 'panic|fatal|error'
```"

#### In-depth
Common daemon issues: (1) Disk full — daemon can't create containers, `df -h /var/lib/docker` reveals disk usage. Fix: `docker system prune -a`. (2) Too many open files — daemon hits system fd limit. Fix: set `LimitNOFILE=1048576` in docker systemd service. (3) iptables corruption — container networking fails. Fix: `systemctl restart docker` to reset iptables rules. (4) Hung containers — containers won't stop. Fix: `docker kill --signal=SIGKILL container`. If daemon itself hangs: `kill -SIGUSR1 $(pidof dockerd)` dumps goroutine stacks for debugging.

---

### 395. How do you use `docker system events` for real-time monitoring?
"`docker system events` (or `docker events`) streams real-time Docker lifecycle events.

```bash
# All events
docker events

# Filter to specific events
docker events --filter event=die --filter event=oom

# Filter to specific container
docker events --filter container=myapp

# Events since a time
docker events --since '2024-01-15T10:00:00'

# JSON format for parsing
docker events --format '{{json .}}'
```

Event types: `create`, `start`, `die`, `kill`, `oom`, `pause`, `unpause`, `destroy`, `pull`, `push`, `exec_create`, `exec_start`, `exec_die`, `health_status`."

#### In-depth
`docker events` is the foundation for event-driven automation and monitoring. Automation example: a script that watches for `die` events and sends a Slack alert if a service container exits unexpectedly. Combine with jq for parsing:
```bash
docker events --format '{{json .}}' | jq -r 'select(.status == "die") | "\(.Actor.Attributes.name) died with exit code \(.Actor.Attributes.exitCode)"'
```
This provides real-time container failure notification. For production: pipe events to a log aggregator (Fluentd, CloudWatch) for long-term retention and alerting.

---

### 396. How do you debug Docker image build failures?
"Build failure debugging:

```bash
# See the failed layer's ID in output:
# => ERROR [builder 3/5] RUN npm install
# => CACHED [1/5] FROM node:20-alpine

# Run the last successful layer interactively
docker run -it <last_successful_layer_id> /bin/sh
# Then manually run the failing command

# With BuildKit: target a specific stage
docker buildx build --target builder --load .
docker run -it myimage-builder /bin/sh
# Then inspect and run the failing step

# Print full build context size:
docker build --progress=plain . 2>&1 | grep -i context
```"

#### In-depth
BuildKit's `--progress=plain` output shows every command with full output — invaluable for debugging apt/npm/pip failures that are truncated in the interactive view. When `RUN npm install` fails: rebuild with `--no-cache` to rule out stale cache, add `--verbose` to npm install, or check network connectivity in the build environment (`RUN curl -v https://registry.npmjs.org` as a diagnostic step). For corporate networks with SSL inspection: the build environment may need the corporate CA certificate added to the image before package manager commands.

---

### 397. How do you handle core dumps from Docker containers?
"Core dumps capture process memory state at crash time — essential for segfault analysis.

**Enable core dumps**:
```bash
# On host: set core dump pattern
echo '/tmp/cores/core.%e.%p' > /proc/sys/kernel/core_pattern
mkdir -p /tmp/cores && chmod 777 /tmp/cores

# Container: set core dump size limit
docker run --ulimit core=-1 myapp  # Unlimited core dumps
```

**Mount a volume for dumps**:
```yaml
volumes:
  - /tmp/cores:/tmp/cores  # Accessible from host
```

**Analyze with GDB**:
```bash
gdb /app/server /tmp/cores/core.server.1234
(gdb) bt  # Backtrace
(gdb) info threads
```"

#### In-depth
Core patterns are global on the host — `core_pattern` applies to all processes including containers. Setting it to a path that's within a mounted volume is the only way to retrieve core dumps after container exit. Some systems use systemd-coredump (`|/lib/systemd/systemd-coredump %P %u %g %s %t %c %h`) as the core dump handler — core dumps go to the journal and can be retrieved with `coredumpctl`. For Go programs: `GOTRACEBACK=crash` environment variable produces goroutine stack traces + core dump on any panic. For production core dump analysis without rebuilding: ensure your images have debug symbols (separate `-dbg` packages for system libraries).

---

### 398. How do you monitor Docker container disk I/O?
"Disk I/O monitoring for containers:

```bash
# Real-time I/O per container
docker stats --format 'table {{.Name}}\t{{.BlockIO}}'
# Shows: block read/write since container start

# Per-process I/O inside container
docker exec container iotop -b -n3

# Host-level: which cgroup is doing I/O
iotop -b -n1 | grep docker

# Prometheus (cAdvisor):
# container_fs_reads_bytes_total
# container_fs_writes_bytes_total
# container_blkio_device_usage_total
```

For I/O-bound containers: check if using tmpfs for temp files, if writes are sequential vs random."

#### In-depth
Random I/O is significantly slower than sequential I/O — 100x on spinning disks, 10x on SSDs. Container workloads that do lots of small random writes (log files, temp files, SQLite databases) hit disk I/O bottlenecks quickly. Mitigation: use tmpfs for ephemeral data (`--mount type=tmpfs,destination=/tmp`), use a proper database (not SQLite) for production data, or use an NVMe SSD host with high IOPS. The `docker run --device-write-bps` flag limits write throughput per container — useful for preventing one heavy I/O container from impacting others.

---

### 399. How do you use FlameGraphs to diagnose container performance?
"Flame graphs visualize CPU profiling — the width of each bar shows relative CPU time.

**Generating a flame graph for a container process**:
```bash
# Install perf on the host
apt-get install linux-perf

# Profile a container process
CPID=$(docker inspect -f '{{.State.Pid}}' mycontainer)
perf record -F 99 -p $CPID -g -- sleep 30
perf script > perf.data.txt

# Convert to flame graph (Brendan Gregg's tools)
./stackcollapse-perf.pl perf.data.txt | ./flamegraph.pl > flamegraph.svg
```

For language-specific profiles: `async-profiler` (Java), `py-spy` (Python), `pprof` (Go) all generate flame graph SVGs directly."

#### In-depth
Flame graphs solve the core problem with numeric profiling output: too much data to understand intuitively. A flame graph immediately shows: the widest bars at the top are the hottest code paths. For a web server, seeing `json.Marshal` taking 40% of CPU is immediately obvious in a flame graph and hidden in tables. `pprof`'s HTTP server built into Go makes this zero-friction: `import _ "net/http/pprof"`, then `curl localhost:6060/debug/pprof/flamegraph > flamegraph.svg` — a single command from your workstation to a production container.

---

### 400. How do you implement SLO-based alerting for Docker workloads?
"SLO alerting uses multi-window, multi-burn-rate alerts to catch both fast and slow error budget consumption.

**Prometheus alerting rules** (based on Google SRE book):
```yaml
groups:
  - name: slo_alerts
    rules:
      - alert: ErrorBudgetBurnFast
        expr: |
          (rate(http_requests_total{status=~'5..'}[1h]) / rate(http_requests_total[1h])) > 14.4 * 0.001
          AND
          (rate(http_requests_total{status=~'5..'}[5m]) / rate(http_requests_total[5m])) > 14.4 * 0.001
        for: 2m
        annotations:
          summary: 'Burning error budget at 14.4x rate (fast burn)'
```"

#### In-depth
The multi-window burn rate approach catches two threat patterns: fast burns (a new deployment breaks 50% of requests — caught in minutes) and slow burns (a subtle bug affecting 0.1% of requests — caught over hours before it consumes the monthly budget). The burn rate threshold is derived from: `budget_consumption_rate = burn_rate × error_rate_target`. A burn rate of 14.4× on a 99.9% SLO (0.1% error budget) means the budget is consumed in ~2 hours (1 month / 14.4). Alert before budget is gone, not after.

---

### 401. How do you implement automated canary analysis for Docker deployments?
"Canary analysis compares metrics between the current production version and a new canary version.

**With Argo Rollouts + Prometheus**:
```yaml
analysis:
  metrics:
    - name: success-rate
      interval: 2m
      thresholdRange:
        min: 99
      provider:
        prometheus:
          address: http://prometheus:9090
          query: |
            sum(rate(http_requests_total{app='{{args.service-name}}', status!~'5..'}[2m]))
            /
            sum(rate(http_requests_total{app='{{args.service-name}}'}[2m])) * 100
    - name: p99-latency
      thresholdRange:
        max: 500  # ms
```"

#### In-depth
Automated canary analysis with Argo Rollouts is the safest production deployment approach: 5% of traffic → canary → automated statistical analysis → if healthy, promote to 10%, 25%, 50%, 100%. If the canary's metrics are worse than baseline (success rate down, latency up, error rate up), automatic rollback within minutes. Companies like Netflix (Spinnaker) and LinkedIn built internal canary analysis systems before Argo Rollouts open-sourced this capability. The key metric choices: always monitor success rate AND latency AND business metrics (conversion rate, checkout completions).

---

### 402. How do you profile Docker containers in production safely?
"Safe production profiling principles:

1. **Low-overhead profilers**: async-profiler (Java), py-spy (Python), pprof (Go), perf with low sample freq (99 Hz, not 1000 Hz)

2. **Short profiling windows**: 30-60 seconds is enough for most analyses. Long profiles add unnecessary overhead.

3. **Sample-based profiling** (not instrumentation): doesn't add code to every function call. Total overhead <3%.

4. **Profile one instance only**: in a fleet of 10 containers, profile one while others handle production load.

5. **Use HTTP endpoints in dev/staging too**: always run pprof endpoints — catch performance issues before production."

#### In-depth
`py-spy` is particularly impressive for production Python: it's a Rust binary that attaches to the Python process via ptrace (or proc filesystem) to sample the call stack. No Python code modification, no import, no overhead outside of the sampling. `py-spy top --pid 1234` gives an htop-like view of Python function call percentages. `py-spy record --pid 1234 -o profile.svg` produces a flame graph. The entire profiling session adds <1% CPU overhead. For Node.js: `clinic doctor` attaches to a running Node.js process and generates a flamegraph + bottleneck analysis report.

---

### 403. How do you detect and fix Docker container memory leaks?
"Memory leak detection process:

**1. Trend monitoring** (Prometheus):
`container_memory_working_set_bytes` growing steadily over hours = memory leak

**2. Heap profiling** (allocations over time):
```bash
# Go: two heap snapshots, compare allocations
curl localhost:6060/debug/pprof/heap > heap1.pprof
sleep 300
curl localhost:6060/debug/pprof/heap > heap2.pprof
go tool pprof -diff_base heap1.pprof heap2.pprof
```

**3. Java: heap histogram over time**:
```bash
docker exec container jmap -histo:live <pid> > histo1.txt
sleep 300
docker exec container jmap -histo:live <pid> > histo2.txt
diff histo1.txt histo2.txt
```"

#### In-depth
The most insidious memory leaks in containerized applications: (1) Go goroutine leaks — goroutines created but never terminated, each holding stack memory. Monitor `runtime_goroutines_total` metric. (2) Node.js event listener leaks — adding listeners without removing them. Track `process_heap_bytes` metric trend. (3) Connection pool leaks — database connections opened but not returned to pool. Monitor `pg_stat_activity` count growing with connection age. The `pprof --alloc_objects` profile (Go) shows which functions are allocating the most objects — useful for finding allocation hotspots that strain the GC.

---

### 404. How do you implement rolling log analysis for Docker containers?
"Automated log analysis identifies patterns and anomalies without manual review.

**Log pattern monitoring** with Prometheus + Loki:
```yaml
# Loki alerting rule
groups:
  - name: log_alerts
    rules:
      - alert: ErrorSpikeDetected
        expr: |
          sum(rate({app='myapp'} |= 'ERROR' [5m])) by (app)
          > 10 * sum(rate({app='myapp'} |= 'ERROR' [1h])) by (app)
        for: 3m
        annotations:
          summary: 'Log error rate spiked 10x above baseline'
```

**Automated anomaly detection**: ElasticSearch ML (Machine Learning) jobs analyze log patterns and detect statistical anomalies automatically."

#### In-depth
Loki (Grafana's log aggregation) is designed for Docker logs: it stores log streams indexed by labels (container name, pod, namespace) and provides LogQL (similar to PromQL) for querying. It's lightweight compared to Elasticsearch (no indexing of log content, only labels) — dramatically lower storage and compute cost. For Docker: use the Loki Docker logging driver (`docker run --log-driver=loki --log-opt loki-url=http://loki:3100/loki/api/v1/push`). Grafana's unified dashboard shows Prometheus metrics + Loki logs + Jaeger traces side-by-side — correlation without switching tools.

---

### 405. How do you detect network anomalies in Docker environments?
"Network anomaly detection in Docker:

**Prometheus metrics-based detection**:
```yaml
# Alert on unusual outbound connection count
- alert: UnusualOutboundConnections
  expr: |
    rate(container_network_tcp_usage_total{state='established'}[5m]) > 1000
  for: 5m
```

**Falco rules** for unexpected network connections:
```yaml
- rule: Unexpected Network Connection
  desc: Container connects to unexpected external IP
  condition: |
    outbound and container and not fd.snet in (approved_cidrs)
  output: 'Unexpected connection from %container.name% to %fd.rip%'
  priority: WARNING
```

**Hubble** (Cilium): real-time network flow monitoring with policy violation alerts."

#### In-depth
Network anomaly detection catches: data exfiltration (unusual large outbound transfers), C2 communication (containers connecting to unknown external IPs), cryptomining (high outbound bandwidth to mining pools), and lateral movement (containers connecting to other containers they don't normally communicate with). Falco's network rules combined with a curated list of approved CIDR ranges (your internal services, known external dependencies) provide effective detection. False positive management: start in alert-only mode, tune approved lists, then enable blocking mode.

---

### 406. How do you implement dashboards for Docker container metrics?
"Grafana dashboards for Docker container observability:

**Recommended dashboard stack**:
1. Prometheus + cAdvisor: collect container metrics
2. Grafana: visualize with pre-built dashboards

**Key dashboards**:
- Docker Overview (ID: 893): per-container CPU, memory, network, I/O
- Node Exporter Full (ID: 1860): host-level resource usage
- Docker Swarm Services (ID: 609): Swarm service task status

**Custom panels to add**:
```
# CPU saturation (throttling percentage)
sum(rate(container_cpu_cfs_throttled_periods_total[5m]))
/ sum(rate(container_cpu_cfs_periods_total[5m]))
```"

#### In-depth
The most valuable custom Grafana panels for Docker: (1) CPU throttle rate by service — shows which services are CPU-constrained. (2) Memory working set vs limit — shows how close each service is to being OOM-killed. (3) Restart count over time — services restarting frequently indicate instability. (4) Image pull time per deployment — helps identify registry performance issues. Set up alerting directly in Grafana for these panels — `Alert when throttle_rate > 0.25 for any service for > 5 minutes` catches CPU-constrained services before they affect users.

---

### 407. How do you perform vulnerability scanning during image builds?
"Shift-left scanning: scan at build time before pushing to registry.

**CI pipeline with Trivy scan gate**:
```yaml
# GitHub Actions
- name: Build image
  run: docker build -t myapp:$GITHUB_SHA .

- name: Scan for vulnerabilities
  uses: aquasecurity/trivy-action@master
  with:
    image-ref: 'myapp:${{ github.sha }}'
    format: 'sarif'
    output: 'trivy-results.sarif'
    exit-code: '1'
    severity: 'CRITICAL,HIGH'
    ignore-unfixed: true  # Don't fail on unfixable CVEs

- name: Push image
  if: success()  # Only push if scan passes
  run: docker push myapp:$GITHUB_SHA
```"

#### In-depth
`ignore-unfixed: true` is key for practical CI gates — many CVEs are reported on packages that have no available fix yet. Failing CI on unfixable CVEs blocks deployments without a path to resolution. The workflow: fail on CRITICAL with available fix (can be addressed), alert on CRITICAL with no fix (track, remediate when fix available), ignore MEDIUM and LOW in the CI gate (manage via a separate vulnerability management process). Regularly review the ignored list — no-fix CVEs gain fixes over time.

---

### 408. How do you implement drift detection for Docker containers?
"Container drift: changes to a running container's state vs. the declared image (security concern).

**Detection approaches**:

1. **Falco runtime rule**:
```yaml
- rule: File Modified in Running Container
  condition: |
    open_write and container and not proc.name in (expected_writers)
  output: 'File modified in container: %container.name% %fd.name%'
```

2. **Periodic `docker diff` checks**:
```bash
docker diff mycontainer | grep -v 'A /tmp' | grep -v 'A /var/run'
# Alert if unexpected changes found
```

3. **Immutable container enforcement**: `docker run --read-only` prevents any runtime filesystem changes."

#### In-depth
Container drift is a security signal: production containers should be immutable — if files are being modified at runtime, either the application has a bug (writing to the container filesystem instead of a volume), or an attacker is modifying system files. Falco's write detection rules catch common attacker techniques: writing cron jobs (`/etc/cron.d/`), adding SSH keys (`/root/.ssh/authorized_keys`), or modifying `/etc/passwd`. Combine with alert routing: Falco → Webhook → PagerDuty for immediate incident response on suspicious write events in production containers.

---

### 409. How do you manage container configuration drift across environments?
"Configuration drift: dev, staging, and prod container configurations diverge over time.

**Prevent drift**:
1. **Infrastructure as Code**: all configurations in Git (Compose files, Helm values, Terraform). No manual changes.
2. **Immutable configuration**: environment-specific values in ConfigMaps/Secrets, never in images. The same image runs in all environments; only config differs.
3. **GitOps**: Argo CD or Flux continuously reconciles deployed state against Git. Any manual `kubectl apply` is reverted automatically.

**Detect drift** (Kubernetes):
```bash
kubectl diff -f ./k8s/  # Shows diff between running and declared state
```"

#### In-depth
GitOps (Argo CD) is the most effective drift prevention: Argo CD watches your Git repo and the Kubernetes cluster simultaneously. If anyone manually patches a deployment (`kubectl edit`), Argo CD detects the divergence within minutes and either alerts (sync policy: manual) or automatically reverts to the Git-declared state (sync policy: automatic). This makes Git the single source of truth and eliminates configuration drift by design. The audit trail in Git (PR history, commit authors) also satisfies compliance requirements for change management.

---

### 410. How do you use distributed tracing to debug microservice latency?
"Distributed tracing finds which service/operation is causing latency in a multi-service request.

**Implementation**:
```python
# Python with OpenTelemetry
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter

tracer = trace.get_tracer(__name__)

with tracer.start_as_current_span('process-order') as span:
    span.set_attribute('order.id', order_id)
    result = call_payment_service(order_id)  # Automatically creates child span
    span.set_attribute('payment.status', result.status)
```

**Analysis in Jaeger/Tempo**: view the full trace waterfall, identify which span takes the most time."

#### In-depth
The trace waterfall diagram immediately shows: is the total latency in serial calls (A→B→C each taking 100ms = 300ms total) or parallel calls (A calls B and C simultaneously, both take 100ms = 100ms total)? For serial bottlenecks: find the longest single span. For parallel: if total time > max parallel span, something is blocking. Common findings: database N+1 queries (10 spans to `db.query` each taking 5ms = 50ms total, fixable with one JOIN), synchronous calls that should be async, and cache misses adding round-trips. Tracing makes these invisible inefficiencies immediately visible.

---
