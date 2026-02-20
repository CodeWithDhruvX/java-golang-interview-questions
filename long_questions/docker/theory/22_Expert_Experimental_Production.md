# 🧪 **Expert Docker: Experimental Tech & Advanced Production (301–350)**

---

### 301. What is WASM (WebAssembly) and how does it relate to Docker?
"WebAssembly (WASM + WASI) is emerging as a **lightweight alternative to containers** for some workloads.

Docker now supports running WASM containers alongside OCI containers:
```bash
docker run --runtime=io.containerd.wasmedge.v1 ghcr.io/containerd/runwasi/wasi-demo-app:latest /wasi-demo-app.wasm
```

WASM advantages vs Docker containers:
- **Smaller**: binaries are KBs vs. MB+ images
- **Faster startup**: milliseconds vs. seconds
- **Better sandbox**: no OS syscall surface, language-agnostic security model
- **Cross-platform**: compile once, run anywhere (any OS, any CPU)

WASM is not replacing Docker — it targets edge computing, short-lived functions, and plugin architectures."

#### In-depth
Docker + WASM uses the `containerd-shim-wasmedge-v1` runtime (or wasmtime, wasmer). WASM containers share the OCI image format and can be pushed to OCI registries — but they run the WASM runtime instead of Linux containers. For Kubernetes, SpiderLightning and the krustlet project bring WASM to K8s pods. The realistic near-term use: WASM for short-lived, untrusted user code execution (plugin systems, user-defined functions in databases), while standard Docker containers handle long-running services.

---

### 302. What is the future of Docker and container runtimes?
"The container ecosystem trajectory:

**Consolidation around OCI**: images are portable across all runtimes. The runtime (Docker, containerd, CRI-O, Podman) is becoming an implementation detail.

**Kubernetes dominance**: Docker Swarm adoption declining. Most new production deployments use Kubernetes.

**Rootless and security-first**: rootless Docker, user-namespace containers, gVisor/Kata becoming more mainstream.

**Supply chain security**: SBOM, cosign signing, and SLSA compliance becoming standard practices.

**WASM integration**: edge and serverless use cases attracting WASM alongside containers.

**eBPF revolution**: Cilium (eBPF CNI) replacing iptables-based networking, Falco (eBPF) for security monitoring."

#### In-depth
The most significant near-term development: eBPF-based container networking and security. Cilium with eBPF processes network packets in the kernel without userspace overhead — 50-100% faster than iptables for L4, and supports L7 policies (HTTP, gRPC, Kafka) natively. Hubble (Cilium's observability layer) provides real-time network flow visibility for free as a side effect of eBPF's packet introspection. The iptables-based Docker networking model will gradually be replaced by eBPF-based solutions in production Kubernetes clusters.

---

### 303. What is the kata containers project?
"Kata Containers is an OCI-compatible container runtime that runs containers inside **lightweight VMs**.

Each container (or pod in Kubernetes) runs in its own micro-VM with a dedicated kernel. The host kernel is NOT shared — a container kernel exploit only compromises the VM, not the host.

Architecture: `containerD → kata-runtime → QEMU/KVM → micro-VM kernel → container process`

Docker-compatible: `docker run --runtime=kata-runtime myimage` uses Kata. Kubernetes-compatible: `RuntimeClass: kata` in pod specs.

Use cases: multi-tenant environments, security-sensitive workloads (financial data, PII), and running untrusted code."

#### In-depth
Kata Containers are used in production by major cloud providers for multi-tenant serverless: AWS Firecracker (used by Lambda and Fargate) is Amazon's Kata-inspired micro-VM, Google uses gVisor and KVM-based isolation for Cloud Run. The performance overhead of Kata: ~100-200ms cold start (vs ~10ms for containers, but much less than VMs), ~3-5% CPU overhead for the VMM layer, and higher memory overhead per container. For multi-tenant security, the tradeoff is clearly worth it.

---

### 304. What is gVisor and when should you use it?
"gVisor is a **user-space kernel** that intercepts container syscalls, validating and mediating them before passing a subset to the host kernel.

The container process makes syscalls → gVisor's Sentry (user-space kernel) intercepts → validates against allowed operations → if safe, proxies to host kernel via a small set of upcalls.

Benefits: container exploits compromise Sentry (not the host kernel), drastically reduced host kernel attack surface, OCI-compatible, lower overhead than full VMs.

Run with Docker: `docker run --runtime=runsc myimage` (runsc is gVisor's OCI runtime)."

#### In-depth
gVisor's trade-off vs Kata: gVisor intercepts ALL syscalls in user-space (higher per-syscall overhead, ~20% CPU overhead), while Kata uses hardware virtualization (lower per-syscall overhead but higher memory overhead per container). gVisor doesn't support all Linux syscalls — applications using unusual syscalls (`perf_event_open`, some ioctl variants) may not work. Test your application with gVisor before deploying — `dmesg` inside a gVisor container shows compatibility issues. Google Cloud Run uses gVisor, so any app deployed to Cloud Run is gVisor-validated.

---

### 305. What is eBPF and how is it used in container environments?
"eBPF (extended Berkeley Packet Filter) is a Linux kernel technology that allows running **sandboxed programs in the kernel** without changing kernel source or loading kernel modules.

In container environments:
- **Networking** (Cilium): replaces iptables/netfilter for pod networking. 50-100x faster packet processing, L7-aware
- **Security** (Falco with eBPF): monitors syscalls, file access, network connections at kernel level
- **Observability**: Pixie/Cloudflare's eBPF tools capture network flows and application metrics without instrumentation
- **Load balancing** (Cilium): kube-proxy replacement using XDP (eXpress Data Path) for kernel-native load balancing

eBPF programs are JIT-compiled and verified for safety before being loaded into the kernel."

#### In-depth
The eBPF revolution eliminates the need for privileged sidecars or kernel module loading for container monitoring and networking. Pixie (CNCF project) uses eBPF to automatically capture and analyze HTTP requests, SQL queries, Redis commands, and gRPC calls across your entire Kubernetes cluster — zero-instrumentation observability. Data stays in the kernel, reducing memory and CPU overhead compared to agent-per-pod approaches. eBPF requires a recent kernel (4.19+ for basic features, 5.15+ for full XDP/sk_msg support).

---

### 306. How do you use Docker with ARM architecture?
"ARM support requires building ARM-targeted images. Key approaches:

**1. Multi-platform builds with BuildKit**:
```bash
docker buildx build --platform linux/arm64,linux/amd64 -t app:latest --push .
```

**2. QEMU emulation** (slow, works anywhere):
```bash
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker run --platform linux/arm64 arm64v8/ubuntu
```

**3. Native ARM builders**: connect a Raspberry Pi or AWS Graviton instance as a BuildKit node.

**4. Cross-compilation** (Go/Rust): compile ARM binaries natively on x86 without emulation."

#### In-depth
AWS Graviton3 (arm64) instances offer 40% better price/performance than equivalent x86 for many workloads. The key requirement: your Docker images must be multi-platform. Most official Docker Hub images (nginx, postgres, node, python, golang) are already multi-platform. For your own apps, the multi-platform BuildKit workflow produces a manifest list — Kubernetes nodes on Graviton instances automatically pull the arm64 variant. Migrating to Graviton for non-GPU workloads is one of the best immediate cost optimizations available.

---

### 307. What is the difference between docker save, docker export, and docker push?
"`docker save myimage:tag > image.tar`: saves the **image** as a tar archive — includes all layers, image manifest, config.json. Can be loaded with `docker load < image.tar`. Preserves layer history and metadata. Used for air-gapped environments.

`docker export container > container.tar`: exports the **container filesystem** (current state including writes) as a flat tar. No layer information, no metadata. Loaded with `docker import` creating a single-layer image. Used for: exporting a container's modified filesystem, creating minimal images.

`docker push`: uploads image layers to a registry. More efficient (registry-to-registry layer deduplication, resumes on failure). The standard production method."

#### In-depth
`docker export` + `docker import` is a way to squash all layers into one — the imported image has a single layer with the container's complete filesystem. This removes layer history, makes `docker history` useless for analysis, but can reduce image size when many intermediate layers have redundant data. Before BuildKit multi-stage builds made squashing easy, `export + import` was the manual squash method. Today, multi-stage builds with a clean final stage achieve the same result properly.

---

### 308. How does Docker work with GPU workloads?
"GPU container support via **NVIDIA Container Toolkit** (for NVIDIA GPUs):

Install on the Docker host: `nvidia-container-toolkit`. Adds a custom OCI hook that grants GPU device access.

Run with GPU:
```bash
docker run --gpus all nvidia/cuda:12.3-base nvidia-smi
docker run --gpus '"device=0,1"' myml:latest  # Specific GPUs
docker run --gpus all --shm-size=16g myml:latest  # Large shared memory for ML
```

For AMD GPUs: ROCm (Radeon Open Compute) provides similar functionality. `docker run --device=/dev/kfd --device=/dev/dri rocm/tensorflow:latest`."

#### In-depth
GPU container isolation: each container gets access to specified GPU devices but the underlying GPU memory is NOT isolated between containers — if two containers share GPU 0, they share the VRAM address space (though context-switched). True GPU isolation (similar to vGPU) requires NVIDIA MIG (Multi-Instance GPU) on A100/H100 GPUs — each MIG instance gets dedicated compute and memory. For Kubernetes, the `nvidia.com/gpu` resource request allocates GPU resources via the NVIDIA GPU Operator, which manages MIG partitioning and device plugin automatically.

---

### 309. What is Docker's BuildKit daemon mode?
"BuildKit can run as a standalone daemon (`buildkitd`) independent of dockerd. This is the backend for `docker buildx`.

Run as a daemon:
```bash
buildkitd --addr unix:///run/buildkit/buildkitd.sock
```

Connect BuildKit from CLI:
```bash
export BUILDKIT_HOST=unix:///run/buildkit/buildkitd.sock
buildctl build --frontend dockerfile.v0 --local context=. --local dockerfile=.
```

Docker Buildx creates `docker-container` builders that run BuildKit inside a Docker container:
```bash
docker buildx create --name mybuilder --driver docker-container
docker buildx use mybuilder
```"

#### In-depth
The docker-container builder backend runs BuildKit in a privileged Docker container — isolating it from the host system. Benefits: multiple isolated build environments (one per project), custom resource limits for the BuildKit container, and the ability to run different BuildKit versions for different projects. In Kubernetes, the `kubernetes` driver runs BuildKit as a pod: `docker buildx create --driver kubernetes --driver-opt replicas=3` — distributes build load across multiple BuildKit pods, consuming cluster resources for builds.

---

### 310. How do you use Docker for edge computing?
"Edge computing deploys containers on resource-constrained devices close to data sources (IoT, factories, retail stores).

**Challenges**:
- Limited CPU/RAM (Raspberry Pi has 4GB RAM max)
- Intermittent connectivity to cloud
- Diverse hardware (ARM, x86, RISC-V)
- Remote management without physical access

**Solutions**:
- **k3s**: lightweight Kubernetes for edge (100MB binary, 512MB RAM)
- **Balena**: Docker-based IoT platform with OTA update management
- **Multi-platform images** (linux/arm/v7, linux/arm64): single image works on all ARM devices
- **Offline operation**: image pull-through cache on-premise, apps work with connectivity loss"

#### In-depth
Balena (balenaCloud) is the most complete Docker-for-IoT solution. It manages fleets of devices running balenaOS (Docker-optimized embedded Linux), handles OTA updates (new containers auto-deployed to devices), and provides secure remote access (SSH via HTTPS tunnel). The update model: push new images to balena registry → balena orchestration updates devices one-by-one with rollback on failure. This is Docker's rolling-update model applied to thousands of physical edge devices.

---

### 311. What is Docker's content-addressable storage?
"Docker uses content-addressed storage (CAS) for image layers — each layer is stored indexed by its SHA256 digest.

When you pull an image, each layer is stored as a compressed tar file named by its digest: `/var/lib/docker/overlay2/<layer-digest>/`. If two images share a layer (same base image), they share the same on-disk data — deduplicated automatically.

The layer digest is computed from the layer content: if any file changes, the digest changes. This ensures integrity: `docker pull` verifies the downloaded layer matches its digest — any tampering is detected."

#### In-depth
Content-addressable storage enables efficient multi-image storage: 100 images all based on `ubuntu:22.04` share the same ubuntu layers on disk. Only the unique application layers add storage per image. The deduplication is automatic — no explicit configuration. `docker system df` shows the effective storage savings: `SIZE` vs `SHARED SIZE`. The practical implication: rebuilding an image with only a code change (not a dependency change) shares all dependency layers with the previous build — the repo grows slowly with many images.

---

### 312. How do you implement feature flags with Docker deployments?
"Feature flags decouple deployment from feature activation — deploy code with the flag disabled, then enable without redeployment.

**Environment variable flags** (simplest):
```yaml
environment:
  FEATURE_NEW_UI: 'false'
```
Toggle via `docker service update --env-add FEATURE_NEW_UI=true myservice` — triggers a rolling restart.

**External flag service** (no restart needed):
Services poll/subscribe to a feature flag service (LaunchDarkly, Unleash, Flagsmith). Flag changes propagate in real-time without container restarts.

**ConfigMap-based K8s** (with file watch): mount flags as a ConfigMap volume — applications watch the file for changes."

#### In-depth
Real-time feature flags (external service, no restart) are the most powerful for production. LaunchDarkly evaluates flags per-user — enabling gradual rollout: 1% of users see the new feature first, then 10%, then 100%. Combined with Docker canary deployments (new image version for 10% of traffic) and user-level feature flags (10% of users per new instance see new feature), you can stage releases extremely granularly. This is how Netflix and Facebook release features to millions of users with zero downtime.

---

### 313. How do you use Docker for batch processing workloads?
"Batch processing: run a container, process data, exit. Key patterns:

**One-off jobs**:
```bash
docker run --rm -v data:/data myprocessor ./process.sh
```
`--rm` cleans up the container automatically.

**Scheduled batch processing** (cron + Docker):
```bash
0 2 * * * docker run --rm --name nightly-report registry/reporter:latest ./generate.sh
```

**Parallel batch processing**:
```bash
for item in $(cat items.txt); do
  docker run --rm -e ITEM=$item processor:latest &
done
wait  # Wait for all parallel containers
```

**Kubernetes Jobs/CronJobs**: the proper production-grade batch system."

#### In-depth
Docker batch processing's key advantage: each job runs in an isolated container (no state pollution between runs), uses the exact right version of the processing code (pinned image tag), and cleans up automatically (`--rm`). Resource limits prevent one job from starving others. The weakness: no built-in queue management, retry logic, or distributed coordination. For complex batch pipelines: Apache Airflow (with DockerOperator), Prefect, or Argo Workflows (K8s-native) add scheduling, dependencies, retries, and monitoring on top of Docker's container isolation.

---

### 314. How do you test Docker Compose files with automated tools?
"Several tools for testing Compose configurations:

**Compose validation**:
```bash
docker compose config  # Validates syntax and resolves variables
docker compose config --quiet  # Silent validation (exit code = pass/fail)
```

**Container-structure-test** (Google): test images post-build:
```yaml
schemaVersion: '2.0.0'
commandTests:
  - name: "App binary exists"
    command: "ls"
    args: ["/app/server"]
    expectedOutput: ["/app/server"]
fileExistenceTests:
  - name: "Config file exists"
    path: "/app/config.yaml"
    shouldExist: true
```

**Molecule** (Ansible) / **Terratest** (Go): test infrastructure including Docker containers."

#### In-depth
Container-structure-test is particularly valuable for catching regressions in multi-stage build final images: verify the binary exists and is executable, check that non-root user is set, confirm expected files are present and configuration files are in place. Run it in CI after `docker build` and before `docker push` — it's a lightweight sanity check that catches common Dockerfile mistakes (forgetting to copy files, wrong file permissions, missing executables) without requiring the container to run.

---

### 315. What is the Docker Desktop Dev Environments feature?
"Docker Desktop Dev Environments lets you define a **complete development environment** (OS, tools, dependencies, editor config) as a Docker container, shareable via GitHub.

Workflow:
1. Define environment in `.docker/docker-compose.yaml`
2. Share the GitHub repo URL with a colleague
3. Colleague opens Docker Desktop → New Dev Environment → pastes URL → clicks Create
4. Docker pulls the environment image, mounts the repo, and optionally opens VS Code connected to it

Features: environment snapshot (save and share the exact running state), environment cloning, team-shared environments."

#### In-depth
Dev Environments on Docker Desktop is Docker Inc.'s answer to the onboarding problem: instead of a 10-page SETUP.md, send a Docker Desktop link. The environment definition is code — version-controlled alongside the project. New team members start contributing as soon as the container starts. This is similar in concept to VS Code's dev containers but integrated into Docker Desktop's GUI rather than requiring VS Code. The limitation: requires Docker Desktop (not available on all Linux environments) — dev containers work everywhere Docker runs.

---

### 316. How do you use Docker for Windows workloads?
"Docker on Windows supports two types of containers:

**Linux containers** (default): run in a Linux VM (WSL2 or Hyper-V). Standard Docker images work here. This is what Docker Desktop provides on Windows.

**Windows containers**: run Windows application containers natively (requires Windows Server or Windows 10/11 Professional). Two isolation modes:
- `process`: containers share the Windows kernel (similar to Linux containers)
- `Hyper-V`: each container runs in a lightweight Hyper-V VM (stronger isolation)

Switch modes: `docker run --isolation=hyperv windowsservercore:ltsc2022 cmd`"

#### In-depth
Windows containers are necessary for: legacy Windows apps (classic ASP.NET, COM-based services) that can't run on Linux, Windows-specific executables ('.NET Framework', PowerShell scripts), or applications requiring Windows APIs. The limitation: Windows container images are HUGE (Windows Server Core base is ~5GB vs. alpine's 5MB). You cannot share layers between Windows and Linux images. Kubernetes supports Windows node pools (EKS, AKS both support Windows nodes) for mixed Linux/Windows clusters.

---

### 317. How do you handle persistent state in stateless container architectures?
"Stateless containers: any container instance can handle any request (no session state stored locally).

For state that must persist:

**Session state**: store in Redis or memcached (shared across all container instances).

**User uploads**: store in object storage (S3, GCS) — addressable by URL, independent of which container processed the upload.

**Database**: use a managed database service (RDS, Cloud SQL) or a stateful Kubernetes pod with PVC.

**Configuration**: inject at startup via environment variables or secret managers (no config files baked into images).

The goal: containers are cattle, not pets — any can die and be replaced without data loss."

#### In-depth
The 12-Factor App principle #6 (Processes: stateless and share-nothing) directly maps to Docker best practices. The pattern for session management: when container A handles request #1 (creates session), the session token maps to a Redis key. When container B handles request #2 (from the same user), it looks up the token in Redis — no stickiness needed. This enables true horizontal scaling: any container handles any request, load balancer can use pure round-robin.

---

### 318. How do you implement service-level agreements (SLAs) with Docker?
"SLA implementation requires: measurement, alerting, and automated action.

**Measurement**: Prometheus collects availability (uptime) and latency metrics per service.
**SLO configuration**: define SLOs from the SLA:
```yaml
- name: api-availability
  description: API should be up 99.9% (43 minutes downtime/month allowed)
  objective: 99.9
  indicator:
    http:
      goodThreshold: 200-299
      totalThreshold: all
  window: 30d
```

**Alerting**: burn rate alerts — alert when consuming the error budget too fast to meet the SLO.

**Automated remediation**: auto-scale, auto-restart, or trigger on-call if SLO is at risk."

#### In-depth
Error budget management changes the deployment culture: instead of 'is this change risky?' the question is 'do we have error budget for this risk?'. If the service has consumed 80% of its monthly error budget by the 10th, deployments freeze until budget recovers. This creates a quantitative, data-driven conversation between engineering and business about risk. Sloth (open source SLO framework) or Google's SLO monitoring integrates with Prometheus to track error budgets automatically.

---

### 319. How do you use Skopeo for container image management?
"Skopeo is a tool for **working with container images and registries** without requiring a Docker daemon.

```bash
# Copy image between registries (no docker pull + push needed)
skopeo copy docker://registry1/app:v1 docker://registry2/app:v1

# Inspect image metadata without pulling
skopeo inspect docker://registry/app:v1

# Delete an image from a registry
skopeo delete docker://registry/app:old-version

# Sync registry to another
skopeo sync --src docker --dest docker registry1/ registry2/

# Mirror multiple images
skopeo copy docker://mcr.microsoft.com/dotnet/aspnet:8.0 docker://internal-registry/aspnet:8.0
```"

#### In-depth
Skopeo is essential for: **registry-to-registry promotion** (dev → staging → prod without pulling to a build agent), **air-gapped environments** (pre-seed a private registry with all needed images using skopeo sync), and **registry auditing** (`skopeo list-tags docker://myregistry/app` lists all tags without needing `docker pull`). Unlike Docker CLI, Skopeo works with any OCI registry without a daemon — suitable for CI agents that don't have Docker installed or for use in minimal container environments.

---

### 320. How do you use crane for Docker image operations?
"crane is a Go tool from Google for interacting with container registries.

```bash
# Digest a tag (pinning)
crane digest nginx:latest

# Copy between registries (server-side, no local pull)
crane copy nginx:latest myregistry/nginx:latest

# Append new layer to existing image
crane append -b nginx:latest -f new-files.tar -t myregistry/nginx:custom

# List tags
crane ls myregistry/app

# Flatten an image (remove layer history)
crane flatten myimage:latest
```

crane's `append` command is powerful: add files to an existing image without rebuilding from scratch — useful for configuration injection."

#### In-depth
`crane append` enables a pattern called 'image mutation': take a base image, append configuration or CA certificates as a new layer, push as a custom image — all without a Dockerfile or build context. This is used by Kustomize's `images` transformer and by some CI pipelines that inject deployment-time configuration. Combined with `crane copy` (server-side, no bandwidth cost), you can maintain a fleet of customized images derived from upstream official images efficiently.

---

### 321. How do you work with OCI registries programmatically?
"OCI registries expose an HTTP API (OCI Distribution Spec) for programmatic access.

**Using Go ORAS library**:
```go
import "oras.land/oras-go/v2"
// Push any artifact to registry
err = oras.Copy(ctx, src, ref, dst, ref, oras.DefaultCopyOptions)
```

**Using crane Go library**:
```go
img, err := crane.Pull("gcr.io/myproject/myimage:latest")
layers, err := img.Layers()
```

**Direct HTTP** (for simple operations):
```bash
TOKEN=$(docker login ... && cat ~/.docker/config.json | jq -r ....)
curl -H "Authorization: Bearer $TOKEN" \
  https://registry.example.com/v2/myapp/tags/list
```"

#### In-depth
ORAS (OCI Registry As Storage) is the standard Go library for working with OCI registries as a generic artifact store. It's used by Helm (chart storage in OCI registries), cosign (signature storage), and projects storing ML models, SBOMs, or other artifacts in registries. Understanding the OCI Distribution API is valuable for building custom registry clients, mirroring tools, or implementing custom garbage collection and lifecycle management that registry products don't provide out of the box.

---

### 322. What is Docker's experimental `docker init` command?
"`docker init` is a CLI command that **analyzes your project and generates Docker assets** automatically.

It detects the language and framework, then creates:
- `Dockerfile` (optimized for the detected language)
- `.dockerignore` (appropriate exclusions for the language)
- `compose.yaml` (with database and other common service dependencies)

Supported: Go, Python (Django, Flask), Node.js, Rust, ASP.NET, Java, PHP, C.

Usage: `cd myproject && docker init`

Interactive prompts for customization (port, app requirements, node version, etc.)."

#### In-depth
`docker init` generates high-quality, production-ready Dockerfiles — multi-stage builds, pinned versions, non-root users, and optimized layer ordering. It's an excellent starting point even for experienced Docker users: the generated Dockerfile reflects current best practices and can be customized from there. The generated `compose.yaml` includes service health checks and proper `depends_on` conditions — common mistakes that the tool handles by default.

---

### 323. How do you implement contract testing with Docker?
"Contract testing verifies that service A's API matches service B's expectations — without requiring both to run simultaneously.

**With Pact** + Docker:
1. Service B (consumer) generates a Pact contract file (what requests it makes and what it expects)
2. Service A (provider) runs `pact-verifier` against the contract:
```bash
docker run --rm \
  -e PACT_BROKER_BASE_URL=http://pact-broker \
  -e PACT_BROKER_TOKEN=$TOKEN \
  pactfoundation/pact-cli pact-verifier \
  --provider=service-a \
  --provider-base-url=http://service-a:8080
```
3. Provider verification results published to Pact Broker"

#### In-depth
Contract testing is particularly valuable in microservices environments where integration testing requires spinning up many services. Contract tests catch API compatibility issues without full integration test suites. The Pact workflow: consumer-driven contracts (consumer defines what it needs → provider must satisfy) are more stable than provider-first (provider defines API → consumers adapt). In Docker Compose: run the provider service, run the pact verifier container, check the exit code. If the provider fails a contract, the deployment is blocked.

---

### 324. How do you use Docker for functional testing (Testcontainers)?
"**Testcontainers** is a library that manages Docker containers in integration tests — starting and stopping real dependencies programmatically.

```kotlin
@Testcontainers
class DatabaseTest {
  @Container
  val postgres = PostgreSQLContainer<>('postgres:15-alpine')
    .withDatabaseName('testdb')
    .withUsername('test')
    .withPassword('test')
  
  @Test
  fun insertAndQuery() {
    val url = postgres.jdbcUrl  // Dynamic port
    val conn = DriverManager.getConnection(url, 'test', 'test')
    // Test with real Postgres, not a mock
  }
}
```

Available in: Java, Go, Node.js, Python, .NET, Rust."

#### In-depth
Testcontainers changes testing philosophy from 'mock the database' to 'use the real database in every test'. The benefits are compelling: tests catch actual SQL syntax errors, trigger real constraint violations, and test actual transaction behavior. The trade-off: tests are slower (container startup time). Mitigation: use `@Container` at class level (not method level) for shared container, and use `Testcontainers.exposeHostPorts()` to share containers across test suites. With container reuse enabled (`testcontainers.reuse.enable=true`), containers persist between test runs — startup time drops to near zero.

---

### 325. How do you perform chaos engineering with Docker?
"Chaos engineering deliberately injects failures to test system resilience.

**Pumba** (Docker chaos tool):
```bash
# Kill random container every 30s
docker run -it gaiaadm/pumba --interval 30s kill docker-container-name

# Introduce 100ms network delay
docker run -it gaiaadm/pumba netem --tc-image gaiaadm/tc-alpine --duration 1m delay --time 100 docker-container-name

# Limit bandwidth to 1Mbps
docker run -it gaiaadm/pumba netem --tc-image gaiaadm/tc-alpine bandwidth --rate '1mbit' target-container
```

**toxiproxy**: a proxy with configurable failure modes (latency, bandwidth, connection drops):
```bash
docker run -d -p 8474:8474 shopify/toxiproxy
```"

#### In-depth
Chaos engineering in Docker focuses on: network partitions (simulate service-to-service connectivity loss), resource constraints (simulate CPU/memory pressure), and container failures (simulate crash loops). The goal: find resilience gaps before they occur in production. Start with hypothesis-driven experiments: 'we believe the api service handles database connection drops gracefully.' Run Pumba to kill the database container. Observe: does the API return meaningful errors? Does it reconnect automatically? Does the circuit breaker kick in? Document results and fix gaps.

---

### 326. How do you containerize legacy applications?
"Legacy app containerization strategy — the 'lift and shift' approach:

**Step 1: Analyze** the app's dependencies (OS version, libraries, external services, file paths).

**Step 2: Create a matching base image**: if the app requires RHEL 7, use UBI7. If it requires specific glibc versions, pin the distro version.

**Step 3: Identify startup and shutdown signals**: legacy apps often haven't handled SIGTERM. Add a wrapper script or use `dumb-init`.

**Step 4: Handle file paths and configs**: use environment variables to override hardcoded paths. Use bind mounts for config that operators must customize.

**Step 5: Handle stateful data**: identify what's on disk that must persist (logs, uploads, session data) and mount volumes."

#### In-depth
The hardest legacy apps to containerize: those with licensed software tied to MAC addresses or hardware fingerprints (the container's MAC changes on restart). Workaround: Docker's `--mac-address` flag sets a static MAC for the container. Apps that fork to daemonize (detach from the terminal and run in background) — they become PID 1 and then fork, leaving PID 1 as the dead parent, triggering Docker to stop the container. Fix: `--foreground` flag or wrapper script that runs the app in the foreground.

---

### 327. How do you containerize applications that require X11?
"Applications requiring a graphical X11 display can run in containers with X11 forwarding.

**Method 1: X11 socket sharing** (Linux):
```bash
docker run -e DISPLAY=$DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix myapp-gui
```
The container uses the host's X11 display socket. Allow connections: `xhost +local:docker`.

**Method 2: VNC inside container**:
```dockerfile
FROM ubuntu
RUN apt-get install -y xvfb x11vnc
CMD ['sh', '-c', 'Xvfb :1 & x11vnc -display :1 & /app/gui']
```
Access via VNC client.

**Method 3: noVNC** (browser-based X11): add noVNC to serve the VNC over a web browser."

#### In-depth
The browser-based noVNC approach is the most portable: the container runs a virtual X11 display (Xvfb), a VNC server captures it, and noVNC serves it over HTTP. Users access the GUI via a browser — no VNC client installation needed. This is how JupyterHub and online IDEs (VS Code Online, Gitpod) serve graphical development environments in browsers. For CI testing: headless Chrome and Firefox containers using Xvfb run browser UI tests without a physical display.

---

### 328. What is the difference between CMD and ENTRYPOINT in practice?
"The real decision: should the user be able to change the command, or only the arguments?

**ENTRYPOINT**: defines the fixed command that always runs. The container is treated as an executable.
`ENTRYPOINT ['nginx']` → `docker run nginx -g 'daemon off;'` passes `-g 'daemon off;'` as args to nginx.

**CMD**: defines default arguments. Can be completely replaced with `docker run image <newcmd>`.
`CMD ['redis-server']` → `docker run redis /bin/sh` runs a shell instead of redis.

**Together** (recommended pattern):
```dockerfile
ENTRYPOINT ['python', '-u', 'app.py']
CMD ['--port', '8080']  # Default port, overridable
```
`docker run myapp --port 9090` changes the port."

#### In-depth
The common mistake: using CMD when ENTRYPOINT is appropriate, or vice versa. A database image should use ENTRYPOINT to ensure the database always starts — you don't want `docker run postgres /bin/bash` to accidentally modify production data. A utility image (like a tool container) should use CMD — `docker run mytools` runs the default tool, `docker run mytools ls /data` runs a different command entirely. When in doubt: if your container has a primary purpose (a service), use ENTRYPOINT. If it's a utility container, use CMD.

---

### 329. What is the `HEALTHCHECK` instruction timeout behavior?
"HEALTHCHECK timeout determines when a check is considered failed:

```dockerfile
HEALTHCHECK \
  --interval=30s \
  --timeout=10s \
  --start-period=5s \
  --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```

- `--interval=30s`: how often to run the check
- `--timeout=10s`: if the check doesn't complete within 10 seconds, it counts as failed
- `--start-period=5s`: don't count failures in the first 5 seconds (startup grace)
- `--retries=3`: mark unhealthy only after 3 consecutive failures

Status values: `starting`, `healthy`, `unhealthy`."

#### In-depth
The `start_period` is the most commonly misconfigured option. For slow-starting applications (JVM warmup, database migrations that run at startup), set start_period to at least the maximum observed startup time. During start_period, failed health checks don't count toward the retries threshold — they're ignored. After start_period, normal health check counting begins. Without sufficient start_period, a JVM app that takes 40 seconds to warm up will be marked unhealthy and restarted before it ever becomes ready.

---

### 330. How do you implement hot reload in Docker during development?
"Hot reload: code changes reflect immediately without container restart.

**Pattern 1: Bind mount + nodemon/watchexec**:
```yaml
services:
  api:
    build:
      target: dev  # Dev stage with hot-reload tool
    volumes:
      - ./src:/app/src  # Mount source code
    command: nodemon src/index.js  # Watches for changes
```

**Pattern 2: Air** (Go hot reload):
```yaml
command: air  # Uses .air.toml for config
volumes:
  - .:/app
```

**Pattern 3: VS Code devcontainer with live share**: edit files that are part of the container filesystem directly."

#### In-depth
The performance of bind mounts on macOS/Windows Docker Desktop (which uses a VM) can be slow — file watching with inotify doesn't work across VM boundaries without special workarounds. Docker Desktop with VirtioFS (Apple Silicon) or WSL2 (Windows) significantly improves this. For Go specifically, Air's workflow: detect file change in the bind mount → recompile Go binary inside the container → restart the process. This is faster than rebuilding the container image. For production parity, the dev stage of the Dockerfile should include the hot-reload tool while the production stage is clean.

---

### 331. What is a sidecar container pattern and when is it used?
"A sidecar is an additional container that runs alongside the main application container, sharing its namespace (in Kubernetes pods, they share network and optionally PID namespace).

**Common sidecar patterns**:
1. **Log shipper**: Fluentd sidecar reads stdout and ships to central logging
2. **Proxy**: Envoy sidecar intercepts all network traffic for mTLS, observability, circuit breaking
3. **Secret manager**: retrieve and refresh secrets, write to shared volume
4. **Certificate rotator**: fetch and renew TLS certificates, notify main app via signal or volume update
5. **Istio/Linkerd**: service mesh sidecars injected by a mutating admission webhook in Kubernetes"

#### In-depth
The sidecar pattern's trade-off: operational consistency (every service gets the same proxy, the same log shipper, the same secret manager) at the cost of resource overhead (each sidecar needs CPU and memory). Istio's Envoy proxy sidecar adds ~50MB memory and ~5ms latency per hop to every service. For small services, this overhead is proportionally significant. Sidecarless service mesh (Cilium's eBPF-based approach, Istio Ambient Mesh) is emerging to eliminate sidecar overhead while retaining the capabilities.

---

### 332. What is the init container pattern in Kubernetes?
"Init containers run to completion BEFORE the main container starts. They're used for setup tasks that must complete first.

```yaml
initContainers:
  - name: wait-for-db
    image: busybox
    command: ['sh', '-c', 'until nc -z db 5432; do sleep 2; done']

  - name: run-migrations
    image: myapp-migrations:v1.2
    command: ['./migrate', 'up']
    env:
      - name: DB_URL
        valueFrom:
          secretKeyRef:
            name: db-secret
            key: url
```

Init containers: share volumes with the main container, run sequentially in order, must exit 0 for the main container to start."

#### In-depth
Init containers solve the startup ordering problem correctly — unlike `depends_on` in Compose (which is host-based), init containers run inside the pod and are guaranteed by Kubernetes before the main container starts. The migration init container pattern is battle-tested: database schema is ready before the first app request. If the migration fails (non-zero exit), the pod stays in Init state and Kubernetes retries — you see the failure immediately in `kubectl describe pod` before any bad state reaches users.

---

### 333. How do you handle secrets rotation in Docker environments?
"Secrets rotation: replace a secret without downtime.

**Swarm secrets**: create `db_password_v2`, update service to use it, remove `db_password_v1`:
```bash
echo "newpassword" | docker secret create db_password_v2 -
docker service update --secret-rm db_password_v1 --secret-add db_password_v2 myservice
docker secret rm db_password_v1
```

**HashiCorp Vault dynamic secrets**: Vault generates a fresh DB credential for each request with TTL. When TTL expires, Vault rotates automatically. The container uses a Vault agent sidecar that transparently retrieves refreshed credentials.

**AWS Secrets Manager**: rotate on schedule, Lambda triggers app to reload credentials."

#### In-depth
Vault's dynamic secrets are the gold standard for rotation: instead of 'change the DB password and update all services', Vault creates a unique credential for each service (api, worker, report-generator) with specific grants (api gets SELECT + INSERT, worker gets INSERT only). Each credential has a TTL (e.g., 24 hours). Vault rotates automatically by creating new credentials before the old ones expire. If a credential is leaked, it's useless in <24 hours. This eliminates credential sharing (separate creds per service) and manual rotation (automated, TTL-based).

---

### 334. How do you use Docker with HashiCorp Vault?
"Vault provides secrets management, dynamic secrets, and PKI for Docker environments.

**Pattern 1: Pull secrets at startup via entrypoint**:
```bash
#!/bin/sh
# Use Vault Agent or direct API
export DB_PASSWORD=$(vault kv get -field=password secret/myapp/db)
exec "$@"
```

**Pattern 2: Vault Agent sidecar** (Kubernetes): Vault Agent runs as an init container, generates a config file with secrets, and a main container reads them.

**Pattern 3: Vault Agent template** (Consul Template):
```hcl
template {
  contents = "{{ with secret 'database/creds/myapp' }}{{ .Data.password }}{{ end }}"
  destination = "/app/config/db_password"
}
```"

#### In-depth
The Vault Kubernetes Auth method is the most elegant integration: the Kubernetes pod gets a ServiceAccount token, which Vault verifies against the Kubernetes API (the pod is authentic). Vault then issues secrets or credentials appropriate for that ServiceAccount's policies. No long-lived Vault tokens, no secrets in environment variables, no Vault credentials to manage for each container — the pod identity IS the credential for Vault access. This closes the bootstrapping problem: how does the first secret get into the container?

---

### 335. What is multi-tenancy in Docker and how do you implement it?
"Multi-tenancy: multiple tenants (organizations, teams, users) share the same Docker infrastructure, isolated from each other.

**Levels of isolation**:
1. **Container isolation** (weakest): each tenant's workloads run in containers. Technically achieved but kernel is shared.
2. **Network isolation**: each tenant has a dedicated Docker network. No cross-tenant traffic.
3. **Namespace isolation** (Kubernetes): each tenant gets a namespace with NetworkPolicies, ResourceQuotas, RBAC.
4. **VM isolation** (Kata Containers, gVisor): each tenant's containers run in separate VMs. Strong isolation.
5. **Cluster isolation** (strongest): each tenant gets their own cluster. Maximum isolation, maximum cost."

#### In-depth
True multi-tenancy security in Kubernetes requires: separate namespaces (isolation unit), NetworkPolicies (network segmentation), ResourceQuotas (prevent resource hogging), LimitRanges (per-pod resource controls), and RBAC (namespace-scoped access). Harder: preventing noisy neighbors in I/O (PriorityClasses + PodDisruptionBudgets help). Hardest: absolute kernel isolation without VMs — container + namespace + seccomp + AppArmor + no privileged containers provides strong practical isolation but a kernel exploit breaks it. For financial services or healthcare tenants: Kata Containers or dedicated clusters.

---

### 336. How do you prevent container breakout via volume mounts?
"Volume mount-based container breakout vectors:

**1. Docker socket mount** (`-v /var/run/docker.sock:/var/run/docker.sock`): gives full Docker API access — attacker can start privileged containers → host root access. Never mount the socket unless absolutely necessary (Portainer, CI agents). Restrict with RBAC if you must (Docker API Authorization Plugin).

**2. Host path traversal**: `-v /=/host:ro` mounts the entire host filesystem. Avoid mounting `/`, `/proc`, `/sys`, `/dev`, or sensitive host paths.

**3. Write access to sensitive host directories**: `-v /etc:/etc` allows container to write `/etc/cron.d` for persistence. Use `:ro` whenever possible."

#### In-depth
Kubernetes OPA/Gatekeeper policies prevent dangerous volume mounts at admission: deny any pod with `hostPath: /var/run/docker.sock`, deny any pod mounting `/`, `/etc`, `/var/run`, etc. The policy library (Gatekeeper Policy Library) includes pre-built policies for common security violations. For Docker: `--mount type=bind,source=/data,target=/app/data,readonly` is safer than `-v /data:/app/data` — the explicit readonly flag prevents accidental write mounts.

---

### 337. What is the Docker API authorization plugin?
"Docker's authorization plugin system allows third-party plugins to intercept and approve/deny Docker API calls.

The plugin receives every API request (start/stop container, pull image, create volume, etc.) and returns: allow or deny + reason.

**Authorization plugins**: OPA (Open Policy Agent) as a Docker auth plugin, Opaque Docker Auth Plugin.

Example: allow developers to manage containers in their namespace but not run privileged containers or mount sensitive host paths.

```bash
--authorization-plugin=opa-docker-authz
```"

#### In-depth
Docker's authorization plugin is the enforcement point for Docker operations — but it's often overlooked in favor of Kubernetes admission control. For pure Docker/Swarm environments, the OPA Docker authorization plugin provides powerful policy enforcement: 'only allow images from approved registries', 'deny --privileged unless the container is in the CI whitelist', 'require resource limits on all containers'. Without this plugin, any `docker run` command with sufficient Docker socket access can bypass all security policies.

---

### 338. How do you implement Docker-based disaster recovery drills?
"DR drills verify recovery procedures work before a real disaster.

**Regular drill schedule**:
1. **Weekly**: restore a database backup to a test volume. Verify application can read data.
2. **Monthly**: simulate full service failure — delete the service's containers + volumes. Restore from backup, redeploy from registry. Time the recovery and document it.
3. **Quarterly**: full infrastructure DR — provision new Docker hosts from Terraform, deploy all services from registry, restore all databases. Test against RTO/RPO targets.

Tools for automated DR drills: Gremlin (chaos engineering with DR scenarios), or custom scripts."

#### In-depth
DR drill automation: write a drill script that: tears down a non-production replica of prod, restores from the latest backup, runs smoke tests to verify recovery, and measures RTO. Run this drill in CI on a schedule (monthly). The drill runs without human intervention — the results (pass/fail, recovery time) are posted to Slack and logged. If the drill fails, the team fixes the DR procedure before a real disaster occurs. This continuous DR validation is the most reliable way to ensure your DR procedure actually works.

---

### 339. How do you implement rate limiting with Docker?
"Rate limiting in Docker-based systems:

**Nginx rate limiting** (Compose with Nginx proxy):
```nginx
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=100r/m;
location /api/ {
  limit_req zone=api_limit burst=20 nodelay;
  proxy_pass http://api:8080;
}
```

**Traefik middleware** (Docker labels):
```yaml
- 'traefik.http.middlewares.ratelimit.ratelimit.average=100'
- 'traefik.http.middlewares.ratelimit.ratelimit.period=1m'
- 'traefik.http.middlewares.ratelimit.ratelimit.burst=50'
```

**Application-level** (recommended for user-specific limits): use Redis (rate limit bucket per user ID) — accurate across all container replicas."

#### In-depth
Proxy-level rate limiting (Nginx/Traefik) is simple but limits by IP — not effective against well-distributed attacks or for per-user quotas. Application-level rate limiting with Redis enables: per-user rate limits (100 requests/minute per API key), tiered limits (free vs. paid users), and distributed accuracy across all replicas. Redis's `INCR` + `EXPIRE` or the `RATE:LIMITER` Lua script provides atomic counters without race conditions. Libraries (rate-limiting Redis middleware for Express, Django REST Framework throttling) make implementation straightforward.

---

### 340. How do you set up horizontal pod autoscaling equivalent in Docker Swarm?
"Docker Swarm doesn't have native equivalent to Kubernetes HPA (which scales based on CPU/memory metrics). You must build it.

**Manual scaling**: `docker service scale api=10`

**Custom autoscaler** (conceptual):
```bash
#!/bin/bash
CPU=$(docker stats api --no-stream --format '{{.CPUPerc}}' | head -1 | sed 's/%//')
TARGET_REPLICAS=$(awk "BEGIN {print int($CPU / 20 + 1)}")
docker service scale api=$TARGET_REPLICAS
```

**Third-party**: `swarm-autoscaler` project watches Prometheus metrics and scales Swarm services accordingly.

**Kubernetes alternative**: use Kubernetes for metric-based autoscaling (HPA + KEDA for event-driven scaling)."

#### In-depth
The absence of native metric-based autoscaling is Docker Swarm's biggest operational gap vs. Kubernetes. Scaling in Swarm is always manual or custom-scripted. For production autoscaling needs, Kubernetes HPA is the correct tool: with the Metrics Server, HPA automatically scales Deployments based on CPU/memory. KEDA (Kubernetes Event-Driven Autoscaling) extends HPA to scale based on external metrics: queue depth (SQS, Kafka, RabbitMQ), HTTP request rate, database connections — covering virtually all autoscaling scenarios without custom scripts.

---

### 341. What is pod topology spread constraints in Kubernetes?
"Topology spread constraints control how pods are distributed across nodes, AZs, or regions.

```yaml
topologySpreadConstraints:
  - maxSkew: 1
    topologyKey: kubernetes.io/hostname
    whenUnsatisfiable: DoNotSchedule
    labelSelector:
      matchLabels:
        app: api
  - maxSkew: 1
    topologyKey: topology.kubernetes.io/zone
    whenUnsatisfiable: DoNotSchedule
```

`maxSkew: 1` means: at most 1 more replica in any topology segment than any other. This spreads replicas evenly across nodes AND availability zones — preventing all replicas from running in the same AZ."

#### In-depth
Topology spread constraints replace the older pod anti-affinity pattern for replica distribution. Anti-affinity with `requiredDuringScheduling` is too strict (can fail scheduling), and soft anti-affinity is unpredictable. Topology spread gives precise control: `maxSkew: 2` allows 2-replica imbalance (practical for large deployments where strict balance wastes capacity), and `DoNotSchedule` vs `ScheduleAnyway` for the unschedulable case. For HA: spread across 3 AZs with `maxSkew: 1` ensures no single AZ failure takes down more than 1/3 of pods.

---

### 342. How do you run Docker in Docker (DinD) in CI?
"Docker in Docker: running Docker commands inside a Docker container.

**Method 1: DinD image** (privileged):
```bash
docker run --privileged docker:dind &
docker run --link dind:docker docker:latest docker ps
```
Risks: privileged container has host root access.

**Method 2: Docker socket pass-through** (most common in CI):
```bash
docker run -v /var/run/docker.sock:/var/run/docker.sock docker:latest docker ps
```
The inner Docker CLI connects to the host daemon. Containers created are siblings, not children.

**Method 3: Kaniko** (no Docker daemon needed for builds):
```bash
docker run gcr.io/kaniko-project/executor:latest --dockerfile=Dockerfile
```"

#### In-depth
Socket pass-through is the most common CI approach but creates a security boundary violation: the CI container can control any container on the host, start privileged containers, etc. GitHub Actions runners mitigate this by running each job in an isolated VM. In shared CI environments (GitLab shared runners, Jenkins agents), socket pass-through is dangerous — any CI job can affect other jobs' containers. Kaniko is the secure alternative: builds images without Docker daemon, no privileged access needed, designed for Kubernetes/CI environments.

---

### 343. What is Buildah and how does it differ from Docker?
"Buildah is a tool for building OCI/Docker-compatible container images **without requiring a Docker daemon**.

```bash
# Build from Dockerfile
buildah bud -f Dockerfile -t myimage:latest

# Build in script form (no Dockerfile)
container=$(buildah from fedora)
buildah run $container -- dnf install -y nginx
buildah config --cmd '/usr/sbin/nginx -g "daemon off;"' $container
buildah commit $container nginx-custom:latest
```

Key differences from Docker:
- No daemon required
- Rootless by default
- Separate tool (not bundled with container runtime)
- Supports OCI and Docker image formats"

#### In-depth
Buildah's scriptable image building is powerful for programmatic image construction without Dockerfiles. You can build images using shell scripts, Ansible playbooks, or any tool that can run commands. Buildah is the build component in Red Hat's container toolkit alongside Podman (run) and Skopeo (registry operations). On RHEL/CentOS systems without Docker, Buildah + Podman + Skopeo provides 100% of Docker's capabilities in a daemonless, rootless architecture. In CI: `buildah bud` + `buildah push` replaces `docker build` + `docker push` with no privilege requirements.

---

### 344. How do you use Packer to create Docker images?
"Packer is HashiCorp's tool for creating machine images across multiple platforms from a single template. It supports Docker as a builder.

```hcl
source 'docker' 'ubuntu' {
  image = 'ubuntu:22.04'
  export_path = 'myimage.tar'
}

build {
  source 'docker.ubuntu'
  
  provisioner 'shell' {
    inline = [
      'apt-get update',
      'apt-get install -y nginx'
    ]
  }
  
  post-processor 'docker-import' {
    repository = 'myimage'
    tag = 'latest'
  }
}
```"

#### In-depth
Packer's Docker builder creates containers, provisions them (Shell, Ansible, Chef, Puppet provisioners), and exports them as Docker images. Where it excels: teams already using Packer for AMIs can reuse the same Ansible playbooks for Docker images — one provisioner, multiple image types. Where it's inferior to Dockerfiles: no layer caching (each Packer build starts fresh), resulting in slower iteration during development. Packer Docker builds are best for: final image creation from well-tested provisioner scripts, creating golden base images, and teams with existing Packer/Ansible investment.

---

### 345. How do you run containers on AWS Lambda?
"AWS Lambda supports container images up to 10GB via the Lambda Container Image Support.

```bash
# Build and push to ECR
docker build -t myfunction .
aws ecr create-repository --repository-name myfunction
docker tag myfunction:latest <acct>.dkr.ecr.us-east-1.amazonaws.com/myfunction:latest
docker push <acct>.dkr.ecr.us-east-1.amazonaws.com/myfunction:latest

# Create Lambda from container
aws lambda create-function \
  --function-name myfunction \
  --package-type Image \
  --code ImageUri=<acct>.dkr.ecr.us-east-1.amazonaws.com/myfunction:latest \
  --role arn:aws:iam::<acct>:role/lambda-role
```"

#### In-depth
Lambda container images must implement the **Lambda Runtime Interface** — either by using an AWS base image (`aws/lambda/python:3.11`, `aws/lambda/nodejs:20`) that includes the runtime interface client, or by packaging the Runtime Interface Client (RIC) in a custom image and handling invocations via the Lambda Runtime API. The Lambda Runtime Interface Emulator (RIE) lets you test Lambda container images locally: `docker run -p 9000:8080 -v ~/.aws/lambda/runtime:/aws/lambda myfunction:latest`. This is excellent for local debugging before deploying.

---

### 346. How do you implement blue-green deployments with Docker Swarm?
"Swarm blue-green using two services with DNS alias switching:

```bash
# Current production: app-blue
docker service create --name app-blue --network frontend myapp:v1

# Deploy green (staging)
docker service create --name app-green --network frontend myapp:v2

# Verify green is healthy
docker service ls | grep app-green

# Switch traffic: update the routing alias
docker service update --network-add alias=app-production app-green
docker service update --network-rm alias=app-production app-blue

# Rollback: just swap back
docker service update --network-add alias=app-production app-blue
docker service update --network-rm alias=app-production app-green
```"

#### In-depth
With Traefik: both services are deployed simultaneously, and Traefik labels control which service receives traffic. Switching is a label update on the services — Traefik reconfigures instantly with zero dropped connections. `docker service update --label-add traefik.enable=true app-green --label-rm traefik.enable=true app-blue` — Traefik's Docker provider processes this in milliseconds. The alias stays in the DNS, but the routing table update is atomic — no request is served by both versions simultaneously.

---

### 347. What are some anti-patterns in Docker image creation?
"Common Dockerfile antipatterns to avoid:

1. **Running as root** (security risk — use USER instruction)
2. **Storing secrets in images** (use BuildKit secrets or runtime injection)
3. **Multiple RUN commands for cleanup**: `RUN apt && RUN rm -rf /var/cache` doesn't reduce size (whiteout files). Use: `RUN apt && rm -rf /var/cache` in one RUN.
4. **Using `latest` tags for base images** (non-reproducible builds)
5. **No `.dockerignore`** (bloated build context)
6. **Copying entire directory before package install**: `COPY . .` then `RUN npm install` — any code change invalidates npm cache. Instead: `COPY package.json .` → `RUN npm install` → `COPY . .`
7. **No health check** (no visibility into app readiness)"

#### In-depth
Antipattern #6 (poor layer ordering) is the most impactful for daily developer experience. A developer who changes a single JavaScript line and rebuilds shouldn't need to wait for all npm packages to re-download. With correct ordering: `COPY package*.json ./` → `RUN npm ci` → `COPY . .`, only the last COPY layer is invalidated by code changes. This turns a 3-minute rebuild (download 500 npm packages) into a 10-second rebuild (just copy code). Multiply this by 50 developer rebuilds per day × 5 developers → hours saved per sprint.

---

### 348. How do you handle Docker image lifecycle management?
"Image lifecycle management prevents registry bloat and controls costs.

**Registry side**:
- ECR lifecycle policies: delete untagged images after 7 days, keep last 10 tagged versions
- Harbor: built-in garbage collection scheduler
- `docker buildx prune`: clean unused BuildKit cache

**CI side**:
- Delete PR-specific images when PRs are merged/closed (GitHub Action: `on: pull_request: types: [closed]` → delete image)
- Keep main branch images (30 days retention for rollback)
- Keep release tags forever (compliance/audit)

**Host side**: `docker image prune -a --filter until=720h` — remove images unused for 30 days."

#### In-depth
ECR lifecycle policy example for production:
```json
{"rules": [
  {"rulePriority": 1, "selection": {"tagStatus": "tagged", "tagPrefixList": ["release-"], "countType": "imageCountMoreThan", "countNumber": 10}, "action": {"type": "expire"}},
  {"rulePriority": 2, "selection": {"tagStatus": "untagged", "countType": "sinceImagePushed", "countUnit": "days", "countNumber": 7}, "action": {"type": "expire"}},
  {"rulePriority": 3, "selection": {"tagStatus": "tagged", "countType": "sinceImagePushed", "countUnit": "days", "countNumber": 90}, "action": {"type": "expire"}}
]}
```
This keeps the last 10 release tags, deletes untagged in 7 days, and purges any tagged image older than 90 days.

---

### 349. How do you use Docker for testing microservices in isolation?
"Testing microservices in isolation using Docker:

**Consumer contract tests** (Pact):
Service A defines what it expects from Service B → Pact generates a contract → Docker container runs Service B against the contract without Service A needing to run.

**Mock server containers** (WireMock):
```bash
docker run -d -p 8080:8080 wiremock/wiremock
curl -X POST http://localhost:8080/__admin/mappings \
  -d '{"request": {"url": "/api/users", "method": "GET"}, "response": {"status": 200, "body": "[{\"id\":1}]"}}'
```

**Testcontainers for real dependencies**: run actual Postgres, Redis, Kafka as ephemeral test instances."

#### In-depth
The testing pyramid for microservices: unit tests (fast, no containers), service-level integration tests (Testcontainers for real DB/cache), contract tests (Pact for service-to-service API contracts), and minimal end-to-end tests (full stack in Docker Compose, run rarely). The contract testing layer is the most valuable for catching integration issues early — it runs in seconds (no real inter-service calls), and failing contracts block merges before deploying incompatible APIs together.

---

### 350. What are the best practices for Docker in financial services/regulated environments?
"Financial services require compliance (PCI-DSS, SOX, SOC 2) alongside standard security:

**Immutable audit trails**: all container lifecycle events logged and stored for 7+ years (consider AWS CloudTrail, Docker audit logs, Falco to Splunk pipeline).

**CIS Docker Benchmark**: run `docker-bench-security` against all hosts. Remediate all PASS/WARN/FAIL findings.

**Image signing mandatory**: cosign keyless signing + Policy Controller enforcement. Signed images only.

**FIPS-compliant base images**: use FIPS-validated cryptography libraries. Red Hat UBI FIPS, or Ubuntu with FIPS kernel.

**Air-gapped environments**: no direct internet access from container hosts. All images via internal registry. All traffic via proxies with TLS inspection."

#### In-depth
PCI-DSS requirement 10 (audit logs) requires: privileged user access logs, container start/stop events, image pull audit, network connection logs. Falco rules cover most of these: log every `docker exec --privileged`, every container started with `--cap-add`, every connection to non-approved external IPs. FIPS 140-2 validated cryptography: containers using OpenSSL must be built with FIPS/validated providers. Red Hat UBI with the FIPS kernel module (`fips=1` kernel argument) provides government-validated cryptography for applications that require it.

---
