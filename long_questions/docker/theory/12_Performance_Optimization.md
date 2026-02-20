# 🚀 **Docker Performance & Optimization (131–140)**

---

### 131. How do you analyze Docker container performance?
"My performance analysis toolkit:

**Real-time**: `docker stats` — shows CPU %, memory usage, network I/O, block I/O for running containers. `docker stats --no-stream` for a one-time snapshot.

**Historical**: cAdvisor + Prometheus for time-series metrics. `docker system events` for lifecycle events.

**Process-level**: `docker top container` — shows processes inside the container. `docker exec container ps aux` for more detail.

**I/O**: `docker exec container iostat -x 1` — disk I/O stats. For network: `docker exec container sar -n DEV 1`."

#### In-depth
`docker stats` CPU% shows container CPU relative to **all available CPUs**. A container on a 4-core host running at 100% CPU is using 1 full core (25% of total). If limited via `--cpus=1`, 100% means it's at the limit of 1 core. Understand this distinction when setting alerts — a 100% CPU alert on a CPU-limited container means the app is resource-constrained; the same on an unlimited container means it's genuinely maxing out a core.

---

### 132. How do you reduce image size significantly?
"Four high-impact techniques:

1. **Multi-stage builds**: only ship the compiled binary, not the build tools. Go app: 1GB builder → 15MB final
2. **Use distroless or scratch base**: no shell, no package manager, no utilities
3. **RUN cleanup in the same layer**: `RUN apt-get install -y pkg && apt-get clean && rm -rf /var/lib/apt/lists/*`
4. **Avoid unnecessary files**: proper `.dockerignore`, don't copy test files, docs, or dev configs

Measure progress: `docker image history myimage:latest` and `dive myimage:latest` to see which instruction adds the most weight."

#### In-depth
The most impactful step is usually the base image choice. Real numbers: Ubuntu (29MB) → Alpine (3MB) → distroless (1.5MB) → scratch (0MB). For interpreted languages (Python, Node.js, Ruby), you can't easily go to scratch — use Alpine or slim variants. For compiled languages (Go, Rust, C++), scratch + CA certificates is achievable. A Go binary with `CGO_ENABLED=0 GOOS=linux go build` produces a fully static binary with no shared library dependencies.

---

### 133. How does using Alpine base image impact performance?
"Alpine uses `musl libc` instead of `glibc`. For most applications this has **zero performance difference** — the C standard library calls are the same at the API level.

However, there are subtle impacts:
- **DNS resolution**: musl's DNS resolver doesn't use `nscd` caching — each DNS lookup is a fresh system call. Under high volume, this can matter
- **Memory allocator**: musl's allocator is simpler and may perform worse for malloc-heavy workloads
- **Compatibility**: binaries compiled against glibc don't run on Alpine without workarounds (use `gcompat` or rebuild)

For Go apps with `CGO_ENABLED=0`: Alpine and Ubuntu perform identically."

#### In-depth
The musl DNS issue is the most common Alpine gotcha in production. Applications making many short-lived connections (each requiring a DNS lookup) perform better with a DNS cache (dnsmasq sidecar or enabling ndots adjustments). In Go, the resolver is configurable: set `GODEBUG=netdns=cgo` or use the pure Go resolver which caches results internally. Kubernetes also has ndots:5 as default which causes many unnecessary DNS lookups — adjusting to ndots:1 significantly reduces DNS traffic.

---

### 134. What is the performance difference between containers and VMs?
"Containers have near-native performance in most scenarios:

**CPU**: No virtualization overhead. Containers use CPU directly — same performance as bare metal. VMs add 5-10% hypervisor overhead.

**Memory**: Containers use the host memory directly. Less overhead and better cache utilization than VMs which have their own memory hierarchy.

**Startup**: Containers start in milliseconds (process fork). VMs start in 10-60 seconds (OS boot).

**Network**: Container bridge networking adds ~10-15% overhead vs. native. VMs with virtio NICs add ~5-10%. Host networking (container) = bare metal."

#### In-depth
The performance parity between containers and bare metal is real for CPU and memory. The main overhead areas are: **storage** (overlay filesystem adds copy-on-write overhead) and **network** (NAT/bridge adds a few microseconds). For latency-sensitive workloads (sub-millisecond trading, real-time gaming), use `--network=host` and local volumes (or bind mounts). For typical web APIs, the container overhead is completely negligible — usually <1%.

---

### 135. How do you enable Docker build cache effectively?
"Layer cache is controlled by **instruction order** and **cache keys**.

Key principles:
1. **Install dependencies before copying code**: `COPY package.json . && RUN npm install` then `COPY . .`
2. **Group rarely-changing operations**: OS package install in one `RUN` block early
3. **Use `--cache-from`** in CI: `docker build --cache-from=registry/app:latest .`
4. **BuildKit cache mounts**: `RUN --mount=type=cache,target=/root/.npm npm install`
5. **Pin base image versions**: change from `node:20` to `node:20.11.0` avoids unexpected cache invalidation"

#### In-depth
BuildKit's `--cache-from` with type=registry enables **registry-based distributed cache** — CI agents pull cached layers from the registry before building. This solves the cold-start problem in ephemeral CI runners (GitHub Actions, GitLab shared runners) that don't have local Docker cache. The BuildKit build cache is separate from the image cache — `docker buildx build --cache-to=type=registry,ref=registry/cache:key` stores it in the registry.

---

### 136. What are the drawbacks of having large Docker images?
"Large images cause real operational pain:

**Slow deployments**: pulling a 2GB image takes minutes vs. seconds for a 50MB image. In auto-scaling scenarios, this delay is critical.

**More storage costs**: ECR, GCR, Artifact Registry charge per GB stored and transferred.

**Larger attack surface**: more packages = more CVEs. A 2GB Ubuntu image has 500+ installed packages, each a potential vulnerability.

**Slower CI**: agent disk fills up faster, cleanup more frequent.

**Docker Hub rate limits**: each large pull counts against your rate limit limit.

**Cold start latency**: in FaaS (AWS ECS, Cloud Run), image size directly impacts cold start time."

#### In-depth
AWS Lambda container images have a 10GB size limit but cold start time scales with image size. AWS publishes internal data showing a 500MB image cold starts in ~1-2 seconds vs. 50MB cold starting in <200ms. For any latency-sensitive serverless workload, image size optimization gives the most consistent cold-start improvement — more effective than provisioned concurrency in some cases.

---

### 137. How do you optimize Docker builds with `--cache-from`?
"`--cache-from` tells Docker to use an external image as the cache source before checking local cache.

**Usage**: `docker build --cache-from myregistry/app:latest -t myregistry/app:$SHA .`

**CI workflow**: before building, pull the latest image: `docker pull myregistry/app:latest || true`. Then build with `--cache-from`. The `|| true` prevents failure on cold builds.

With **BuildKit**: `docker buildx build --cache-from type=registry,ref=myregistry/cache:key --cache-to type=registry,ref=myregistry/cache:key .` — reads AND writes cache to registry."

#### In-depth
The classic `--cache-from` has a limitation: it uses image layers as cache, which requires the final image to contain the same layers as the cached version. BuildKit's type=registry cache separates the build cache from the image — you can store a flat cache manifest with all intermediate layers without including them in the final image. This is more efficient, especially with multi-stage builds where intermediate stages are large but not shipped.

---

### 138. How do you measure startup time of containers?
"Use `docker events` to capture precise timestamps:

```bash
docker run -d --name myapp myimage
docker events --since=60s --filter event=start --filter container=myapp --format '{{.Time}} {{.Action}}'
```

Or time it directly:
```bash
time docker run --rm myapp node -e 'console.log("ready")'
```

For production readiness time (including app initialization): measure time from container start to first successful health check. `docker inspect --format='{{.State.StartedAt}}' myapp` gives the start time."

#### In-depth
For production, startup time matters most in auto-scaling scenarios and crash-recovery. Track MTTR (Mean Time To Recovery) as: time from container crash to new container passing health check. This includes: scheduler/orchestrator reaction time (~1-2s for Kubernetes), image pull time (0s if cached, seconds/minutes if not), container startup, application initialization, and health check pass. Image pull time dominates in practice for large images.

---

### 139. What causes slow builds in Docker?
"The main causes:

1. **Large build context**: all files in the directory are sent to the daemon — missing `.dockerignore` sends node_modules, .git, etc.
2. **Cache invalidation high up**: `COPY . .` early in Dockerfile invalidates all subsequent cache
3. **Large base image**: pulling a 1GB base image on every CI build
4. **Missing package cache**: `npm install` or `pip install` re-downloads packages every build
5. **Sequential `RUN` commands** that could be parallelized (BuildKit supports parallel stages)
6. **No `--cache-from`** in CI: cold-start builds every time"

#### In-depth
Profile your build with `BUILDKIT_PROGRESS=plain docker buildx build .` — it shows exact timing per step. Usually, one or two steps dominate. For `apt-get install` or `npm install`, the fix is `RUN --mount=type=cache,target=/root/.npm npm install` — the package download cache persists between builds as a BuildKit volume, not an image layer. For `pip install`: `--mount=type=cache,target=/root/.cache/pip`.

---

### 140. How do you profile memory and CPU usage of containers?
"`docker stats` gives a live view per container: CPU %, MEM USAGE/LIMIT, NET I/O, BLOCK I/O.

For deeper analysis: run `cadvisor` which exports per-container metrics in Prometheus format. The `container_cpu_usage_seconds_total` and `container_memory_working_set_bytes` metrics give time-series data for analysis and alerting.

For heap profiling inside the container: use language-specific profilers. For Go: `pprof` over HTTP. For Java: JFR, async-profiler. For Node.js: `--inspect` + Chrome DevTools. Access them via `docker exec` or published ports."

#### In-depth
`container_memory_working_set_bytes` (from cAdvisor) is the most accurate memory usage metric — it includes all memory used by the container including kernel memory for its files. `container_memory_usage_bytes` includes file cache which the kernel reclaims under pressure — less useful for capacity planning. Set alerts on working set approaching 80% of the memory limit to proactively catch memory leaks before OOM kills.

---
