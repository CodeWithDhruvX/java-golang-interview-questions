# 🔗 **Docker + Kubernetes, Tooling & Expert Production (251–300)**

---

### 251. How does a Docker image become a Kubernetes pod?
"The journey: image → pod spec → kubelet → containerd → running container.

1. `kubectl apply` sends a Pod spec (with `image: registry/app:v1`) to the API server
2. Scheduler assigns the pod to a node
3. Kubelet on that node receives the pod spec
4. Kubelet instructs containerd (via CRI) to pull the image from the registry
5. containerd downloads the OCI image layers, stores them via the snapshotter
6. Kubelet tells containerd to create and start the container per the pod spec
7. Container runs; kubelet monitors it and reports back to API server

Docker images work transparently — they're OCI-compliant, which containerd understands natively."

#### In-depth
The image pull happens on the node where the pod is scheduled. First pull is slow (full layer download); subsequent pods using the same image on the same node are instant (layers cached by containerd). In large clusters, image pull storms (many pods scheduling simultaneously on nodes without the cached image) can cause slow rollouts. Mitigation: pre-pull critical images with a DaemonSet job, or use stargz-snapshotter for lazy pulling.

---

### 252. What Docker practices translate directly to Kubernetes?
"Most Docker best practices carry over to Kubernetes directly:

**Translates directly**:
- Multi-stage builds → still use multi-stage builds for K8s images
- Small images → even more important for faster pod scheduling
- Non-root user (`USER app`) → set in `securityContext.runAsNonRoot: true`
- Health checks → become liveness/readiness probes
- Resource limits → map to `resources.limits` in pod spec
- No hardcoded secrets → K8s Secrets or external secret managers

**Kubernetes-specific equivalents**:
- `docker run -e` → K8s env vars in pod spec
- Docker volumes → PersistentVolumeClaims
- Docker secrets → K8s Secrets (or HashiCorp Vault + CSI driver)
- Docker networks → K8s Services + NetworkPolicies"

#### In-depth
The biggest conceptual shift from Docker to Kubernetes: in Docker, you manage containers directly. In Kubernetes, you manage desired state (Deployments, StatefulSets) and the control loop reconciles reality. You don't `kubectl start container` — you declare `replicas: 3` and Kubernetes ensures 3 replicas exist. This shift requires rethinking operational habits: instead of 'ssh into container and fix', you 'update the deployment spec and let Kubernetes reconcile'.

---

### 253. How do you troubleshoot Docker containers in a Kubernetes environment?
"Kubernetes adds layers, but Docker debugging skills apply:

```bash
# Get pod details and events
kubectl describe pod myapp-abc123

# Get logs (equivalent to docker logs)
kubectl logs myapp-abc123 -c mycontainer
kubectl logs myapp-abc123 --previous  # Crashed container's logs

# Exec into container (equivalent to docker exec)
kubectl exec -it myapp-abc123 -- /bin/sh

# Copy files out (equivalent to docker cp)
kubectl cp myapp-abc123:/app/logs ./local-logs

# Port forward for debugging (no docker run -p equivalent)
kubectl port-forward pod/myapp-abc123 8080:8080
```"

#### In-depth
When a pod won't start: `kubectl describe pod` shows Events at the bottom — these are the most useful for diagnosing: image pull failures, insufficient CPU/memory on all nodes, volume mount failures, security policy rejections. For running pods with issues: `kubectl logs` with `--previous` gives the last crashed container's logs. For intermittent issues: `kubectl logs -f --since=1h` streams recent logs. The ephemeral container feature (`kubectl debug pod/xxx --image=busybox`) adds a debug sidecar without modifying the pod spec.

---

### 254. How do Docker security contexts compare to Kubernetes security contexts?
"Docker security flags map directly to Kubernetes pod/container security contexts:

| Docker flag | Kubernetes equivalent |
|-------------|----------------------|
| `--user 1000:1000` | `securityContext.runAsUser: 1000` |
| `--read-only` | `securityContext.readOnlyRootFilesystem: true` |
| `--cap-drop=ALL` | `securityContext.capabilities.drop: [ALL]` |
| `--cap-add=NET_BIND` | `securityContext.capabilities.add: [NET_BIND_SERVICE]` |
| `--security-opt no-new-privileges` | `securityContext.allowPrivilegeEscalation: false` |
| `--privileged` | `securityContext.privileged: true` |

In K8s, security contexts can be set at pod level (applies to all containers) or per-container."

#### In-depth
Kubernetes adds **Pod Security Standards** (PSS) at the cluster level: Privileged (no restrictions), Baseline (prevents obvious privilege escalations), and Restricted (enforces security best practices). `kubectl label namespace mynamespace pod-security.kubernetes.io/enforce=restricted` requires all pods in the namespace to meet the Restricted policy. Enforce restricted in all namespaces except system ones — it catches misconfigured pods at admission time rather than after deployment.

---

### 255. What is Kubernetes and why was it created to replace Docker Swarm?
"Kubernetes wasn't created to replace Docker Swarm — they address the same problem independently. Kubernetes was developed at Google (released 2014) based on their internal Borg system, while Docker Swarm was Docker Inc.'s orchestration solution.

Kubernetes gained dominance because:
- Google's backing and CNCF stewardship
- More powerful APIs (extensibility via CRDs, operators)
- Richer ecosystem (Helm, Istio, Argo CD, Prometheus Operator)
- Better scaling and scheduling primitives
- Industry adoption → vendor support → career demand

Docker Swarm is simpler but lacks the extensibility that enterprises needed for complex workloads."

#### In-depth
The 'orchestration wars' (2016-2019) saw Docker Swarm, Kubernetes, and Apache Mesos competing. Kubernetes won due to ecosystem momentum more than purely technical superiority — CNCF foundation provided governance, major cloud providers launched managed K8s (GKE 2014, AKS 2017, EKS 2018), and tooling ecosystem exploded. Docker Inc. eventually adopted Kubernetes support in Docker Desktop and Docker Enterprise, conceding the enterprise orchestration market.

---

### 256. How do you connect Docker containers across hosts without Kubernetes?
"Options without Kubernetes:

1. **Docker Swarm overlay networks**: deploy a Swarm cluster, use the overlay network for cross-host container communication

2. **WireGuard tunnel**: connect Docker hosts via a WireGuard VPN mesh. Containers on each host can reach each other via the WireGuard subnet. Tools: Tailscale, Netbird automate WireGuard mesh setup.

3. **Flannel**: a lightweight overlay network daemon that manages cross-host pod-like networking outside of Kubernetes

4. **Consul + Fabio**: Consul service discovery across hosts, Fabio as a Layer 4/7 load balancer. Each host runs a Consul agent."

#### In-depth
For most small-scale cross-host Docker needs, Docker Swarm's overlay network is the right choice — it's built-in, well-maintained, and requires no additional setup. For non-Swarm cross-host connectivity (e.g., separate Docker hosts that aren't clustered), WireGuard/Tailscale mesh networks are the modern solution: zero-trust networking, encrypted, with automatic topology management. Containers get routable IPs across the mesh without any Docker-specific networking changes.

---

### 257. What is the `docker buildx` command and its use cases?
"`docker buildx` is the Docker CLI plugin for extended build capabilities via BuildKit.

Key use cases:
- **Multi-platform builds**: `docker buildx build --platform linux/amd64,linux/arm64 .`
- **Multiple builder backends**: local, docker-container (isolated BuildKit daemon), Kubernetes (distributed builds in K8s)
- **Advanced cache API**: `--cache-from type=registry` and `--cache-to type=registry`
- **Build attestations**: SBOM and provenance data embedded in image with `--sbom=true --provenance=true`
- **OCI output**: `--output type=oci,dest=image.tar` for OCI-format archives

Primary value: multi-platform images and registry-based distributed build cache."

#### In-depth
The docker-container builder backend creates an isolated BuildKit container: `docker buildx create --name mybuilder --driver docker-container`. This builder doesn't share the local Docker cache — useful for reproducible builds or when you need multiple isolated build contexts. The Kubernetes driver runs builds as Kubernetes pods: `docker buildx create --driver kubernetes --driver-opt replicas=2` — scales build parallelism using K8s infrastructure. Combined with BuildKit's parallel stage execution, this enables massive distributed builds.

---

### 258. How do you create multi-platform Docker images?
"Use `docker buildx build` with `--platform`:

```bash
# Create a multi-platform builder (required)
docker buildx create --name multibuilder --use

# Build and push for multiple platforms
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  --tag myregistry/app:latest \
  --push \
  .
```

The `--push` flag is required for multi-platform — local Docker can only load single-platform images. The result in the registry: a manifest list pointing to platform-specific image manifests."

#### In-depth
Multi-platform builds run the compilation for each platform in QEMU emulation (for cross-compilation) or on native-platform builders (for performance). QEMU emulation is slow — building an arm64 image on x86_64 with QEMU can be 5-10x slower than native. For production CI, use a multi-node builder with both x86 and ARM nodes: `docker buildx create --append --name multibuilder --platform linux/arm64 ssh://arm-builder-host`. The build runs natively on each platform and the manifest is assembled at the end.

---

### 259. How does Docker Scout work for image security?
"Docker Scout is Docker's container image vulnerability analysis tool, integrated into Docker Desktop and Docker Hub.

It analyzes image layers against multiple CVE databases (NVD, GitHub Advisory, vendor advisories) and provides:
- **SBOM generation**: complete software bill of materials for the image
- **CVE matching**: which CVEs affect which packages in the image
- **Fix recommendations**: 'upgrade base image from ubuntu:22.04 to ubuntu:22.04.3 to fix CVE-...'
- **Policy evaluation**: define policies ('no critical CVEs allowed') and fail CI if violated

Usage: `docker scout cves myapp:latest` in CLI, or integrated into Docker Hub image pushes."

#### In-depth
Docker Scout differentiates from Trivy/Snyk by offering the 'compare' feature: `docker scout compare myapp:v1.0 myapp:v2.0` shows which vulnerabilities were added or fixed between versions. This is especially useful for understanding the security impact of a base image update. The integration with GitHub Actions provides PR-level vulnerability comments: 'This PR introduces 2 new HIGH CVEs from the new npm dependency'. Making security actionable at code review time has the highest impact on mean time to remediate.

---

### 260. How do you use Trivy for container scanning?
"Trivy is an open-source vulnerability scanner for containers, filesystems, git repos, and Kubernetes clusters.

**Scan an image**:
```bash
trivy image myapp:latest

# Exit non-zero on CRITICAL vulnerabilities (for CI gates)
trivy image --exit-code 1 --severity CRITICAL myapp:latest

# Output as table, JSON, or SARIF
trivy image --format sarif --output results.sarif myapp:latest
```

**Scan a Dockerfile** before building:
```bash
trivy config ./Dockerfile
```

**Scan a Kubernetes cluster**:
```bash
trivy k8s --report summary cluster
```"

#### In-depth
Trivy integrates brilliantly with CI: the SARIF output uploads to GitHub Security tab, showing vulnerabilities inline in PRs. Trivy's database includes OS packages (apt, rpm), language packages (npm, pip, go modules), and infrastructure misconfigurations. `trivy fs .` scans the current directory's dependency files (go.sum, package-lock.json) without building an image first — catching vulnerabilities even earlier in the development pipeline.

---

### 261. How do you use Dive to analyze Docker image layers?
"Dive is a CLI tool for exploring Docker image layers and their contents.

```bash
# Install
brew install dive  # macOS
# or: go install github.com/wagoodman/dive@latest

# Analyze an image
dive myapp:latest
```

Dive shows: the Dockerfile instruction that created each layer, the files added/changed/deleted in each layer, the cumulative size, and an efficiency score (% of image space that's unique vs. duplicated).

Use `dive --ci` in CI pipelines: fails if image efficiency drops below a threshold or wasted space exceeds a limit."

#### In-depth
Dive's efficiency score catches common mistakes: copying files in one layer and deleting them in another (both layers contribute to image size), installing packages and their cache in separate RUN commands, or accidentally including `.git` or `node_modules` in the final image. I use `dive --ci --lowestEfficiency=0.95` as a gate in PR checks — any PR that introduces wasted layers fails and requires optimization. This keeps images lean throughout the development lifecycle, not just at release time.

---

### 262. How do you implement Docker image governance?
"Image governance ensures only approved, secure images run in production.

**Registry-level**: use Docker Hub Teams & Organizations, ECR resource policies, or Harbor with project-level RBAC to restrict who can push to production registries.

**Deployment-level** (Kubernetes): OPA/Gatekeeper admission webhooks: only images from approved registries, only images with current CVE scans, only signed images (cosign verification).

```yaml
# OPA policy example
deny[msg] {
  input.request.object.spec.containers[_].image
  not startswith(input.request.object.spec.containers[_].image, "myregistry.ecr.aws/")
  msg := "Only images from approved registry allowed"
}
```"

#### In-depth
Harbor is the most complete open-source image governance platform: push/pull RBAC per project, Trivy-based vulnerability scanning on every push, cosign signature verification, and replication policies (mirror from Docker Hub to reduce rate limits). For enterprises, Harbor running on Kubernetes (via the official Helm chart) provides a complete private registry with governance features that docker.io and even ECR don't offer out of the box.

---

### 263. What is Docker Extensions and how does it work?
"Docker Extensions adds third-party tools and integrations directly into Docker Desktop's UI.

Extensions are distributed as Docker images containing: a `metadata.json` describing the extension, a web frontend (React/Vue app), and optionally a backend service.

```bash
docker extension install snyk/snyk-docker-desktop-extension:latest
docker extension ls
docker extension rm snyk/snyk-docker-desktop-extension:latest
```

Available extensions: Disk usage visualizer, Logs explorer, Kubernetes cluster viewer, Portainer, Snyk vulnerability scanning, Grafana dashboards."

#### In-depth
Docker Extensions is strategically significant: it turns Docker Desktop from a container runtime into a platform for developer tools. Third-party ISVs (Snyk, Portainer, JFrog) ship their tools as extensions, making them discoverable and easy to install for millions of Docker Desktop users. The extension API exposes Docker Desktop's backend (socket communication) and the Extension SDK provides React components for consistent UI. For teams with custom internal tooling: building a private Docker Extension can standardize developer tool access across the team.

---

### 264. How do you build immutable infrastructure with Docker?
"Immutable infrastructure: servers (or containers) are never modified after creation. To update, replace with a newly built container.

Docker naturally enables this:
1. Never `docker exec container apt-get install` or modify a running container
2. Build a new image with the change: `docker build -t app:v1.1 .`
3. Replace the running container: `docker stop app && docker run -d --name app app:v1.1`
4. Old container is discarded, new one starts from a clean state

Enforcement: use `--read-only` containers to prevent runtime filesystem modifications. Any change requires a new image build and deployment."

#### In-depth
Immutability extends beyond the application container to the host OS. Tools like Flatcar Container Linux or CoreOS provide immutable OS images — the base OS is replaced atomically on update, not patched. Combined with immutable Docker containers, you get a fully immutable stack: application changes → new image → new container; OS changes → new OS image → new host. This eliminates configuration drift and 'snowflake' servers that diverge from their defined state over time.

---

### 265. How do you implement container auto-healing?
"Container auto-healing: detect failures and automatically restart/replace containers without human intervention.

**Docker restart policies**: `--restart=on-failure:5` – auto-restart on non-zero exit, up to 5 times.

**Health check + restart policy combination**: a health check that fails causes Docker to consider the container unhealthy. In Swarm, unhealthy tasks are replaced. In Kubernetes, unhealthy pods are restarted (liveness probe failure).

**Active monitoring approach**: Watchdog container (Shepherd, Ouroboros) watches for new image versions or failed health checks and automatically pulls + restarts containers."

#### In-depth
True auto-healing requires: **detection** (health checks), **replacement** (new container), and **prevention** (circuit breakers to stop cascading failures). Kubernetes does this with liveness probes (restart unhealthy pods) and readiness probes (remove unready pods from service endpoints without restarting them). The readiness probe is often more important: a pod that's temporarily overloaded (not crashing, just slow) should be removed from load balancing temporarily — a liveness-only approach would restart it unnecessarily, discarding in-flight requests.

---

### 266. What is Portainer and how does it help with Docker management?
"Portainer is a **web-based Docker management UI** that provides a visual interface for managing Docker hosts and Swarm clusters.

Features: container management (start/stop/exec/logs), image management (pull/build/push), volume and network management, Docker Compose deployment via editor, Swarm service management, user/team RBAC, and environment management (connect multiple Docker hosts).

Deploy: `docker run -d -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock portainer/portainer-ce`

Portainer BE (Business Edition) adds Kubernetes support, LDAP/SSO, audit logging, and enterprise features."

#### In-depth
Portainer fills the gap for teams transitioning to containerization who aren't comfortable with the CLI yet. It provides operational visibility without requiring CLI expertise: developers can view logs, restart containers, and inspect environment variables without SSH access. The RBAC model is important for production access control: developers get read-only access to view logs; ops engineers get the ability to restart; only CI has push access. Portainer's Docker Socket exposure is a security consideration — secure it with TLS and strong authentication.

---

### 267. What is Watchtower and how does it help manage container updates?
"Watchtower is a container that **automatically monitors and updates running containers** when new image versions are published.

```bash
docker run -d \
  --name watchtower \
  -v /var/run/docker.sock:/var/run/docker.sock \
  containrrr/watchtower \
  --schedule '0 */6 * * *' \
  --cleanup
```

Every 6 hours, Watchtower pulls the latest image for each running container. If a newer image is found, it stops the container, pulls, and restarts with the same options.

Best for: personal projects, lab environments, non-critical services where always-running-latest is acceptable."

#### In-depth
Watchtower is a double-edged sword for production. The appeal: automatic security patch delivery without manual intervention. The risk: an upstream base image update breaks your app at 2AM. In production, I never use Watchtower — updates should go through CI: build → test → deploy. For security patch automation, I use Renovate Bot to create PRs that update the Dockerfile base image tag, run CI (including tests), and auto-merge if CI passes. That's automated updates with safety nets.

---

### 268. How do you use lazy loading of Docker images (stargz)?
"**Stargz** (Seekable tar.gz) is an OCI-compatible image format that enables lazy pulling — containers can start before the entire image is downloaded.

Standard pull: download all layers → extract → start container.
Stargz: start container → fetch only the files needed as they're accessed → stream remaining files in background.

Setup with containerd (eStargz snapshotter):
```bash
# Enable stargz snapshotter in containerd config
[proxy_plugins]
  [proxy_plugins.stargz]
    type = "snapshot"
    address = "/run/containerd-stargz-grpc/containerd-stargz-grpc.sock"
```
Google Cloud Run uses a similar on-demand layer loading system."

#### In-depth
The performance gain depends on what % of the image is actually accessed at startup. For a 1GB Node.js image, startup might only need 100MB of layers — with lazy pulling, the container starts in 1/10th the normal time. Files accessed for the first time after startup stream transparently. This is especially impactful for FaaS/serverless (Cloud Run, AWS Lambda container images) where cold start time directly impacts user-facing latency.

---

### 269. How do you ensure Docker image reproducibility?
"Reproducible images: the same Dockerfile + context → the same image content (same layer digests) across builds.

Challenges:
1. **Timestamps**: file modification times vary per build
2. **Package manager versions**: `apt-get install curl` may get a different version tomorrow
3. **Random ordering**: some package managers output files in non-deterministic order

Solutions:
1. **Pin all versions**: base image by digest, all packages with exact versions
2. **Use BuildKit `--reproducible`** (experimental): normalizes timestamps
3. **Lock files**: `npm ci`, `pip install -r requirements.txt`, `go mod download`
4. **Set `SOURCE_DATE_EPOCH`** for timestamp normalization in builds"

#### In-depth
Fully reproducible Docker images are an advanced goal — most teams aim for functionally reproducible (same functionality, no need for bit-for-bit identical digests). The primary practical goal: ensure a rebuild 6 months later produces the same running behavior, not necessarily the same byte sequence. Version pinning + lock files achieves this. For supply chain security (SLSA framework), generate a build provenance attestation: `docker buildx build --provenance=true --sbom=true` — this cryptographically proves what inputs produced the image.

---

### 270. What are OCI artifacts and how do Docker images relate?
"**OCI artifacts** are the generalization of OCI images — any content (not just container images) can be stored in an OCI-compliant registry.

The OCI specification defines: an **image manifest** (references layers), an **image index** (manifest list for multi-platform), and an **artifact manifest** (for non-image content).

Examples of OCI artifacts: Helm charts (stored in Helm OCI registries), cosign signatures and attestations, SBOMs (CycloneDX, SPDX), WASM modules, ML models.

Docker images ARE OCI artifacts — a subset conforming to the container image specification. All OCI registries (Docker Hub, ECR, GCR, Harbor) can store any OCI artifact."

#### In-depth
The OCI artifact ecosystem is growing rapidly. Helm chart OCI registries (`helm push mychart.tgz oci://registry/charts`) store charts alongside their images in the same registry — single source of truth for deployment. cosign stores image signatures as OCI artifacts attached to the signed image — `cosign verify myimage:v1.0` resolves the signature artifact from the registry. This co-location with the image makes signature verification more reliable than separate signature stores.

---

### 271. How do you manage Docker containers without root access?
"**Rootless Docker**: runs the Docker daemon as a non-root user.

Setup:
```bash
dockerd-rootless-setuptool.sh install
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock
docker run hello-world  # Works without sudo
```

The daemon runs in a user namespace with the current user mapped to root inside. Containers run without real root privileges — even if the container process is 'root', it's the user's UID on the host.

**Podman**: natively rootless — no daemon, no root required. `podman run ubuntu` works without sudo or any setup."

#### In-depth
Rootless Docker has some limitations: port numbers <1024 require additional privilege (handled by `rootlesskit`'s `--port-driver slirp4netns`), some container networking modes are unavailable, and block device access is restricted. For CI environments (GitHub Actions runs as a non-root user), rootless Docker is essential. The runner's Docker socket is rootless by default in some CI environments — understanding rootless Docker avoids mysterious 'permission denied' container failures in CI.

---

### 272. How do you handle Docker image layers and caching in large teams?
"Cross-team cache challenges: individual developer machines have different local caches. Each CI runner starts cold without sharing cache with others.

**Solution: registry-based cache**:
```bash
docker buildx build \
  --cache-from type=registry,ref=registry/cache:main \
  --cache-to type=registry,ref=registry/cache:$BRANCH,mode=max \
  .
```

`mode=max` exports ALL intermediate layer caches (not just the final stage). Any team member's build benefits from layers cached by the CI server.

Also: **cache image as a base**: push the dependencies layer as a tagged image (`app:deps`) — others `FROM app:deps` to reuse it."

#### In-depth
`mode=max` is critical for multi-stage builds — without it, only the final stage's layers are cached, and all intermediate stages (test runner, builder) are rebuilt every time. With `mode=max`, every stage of every branch is cached in the registry — a developer switching from `main` to a feature branch gets cache hits from main's cached layers up to the first divergence point. This is the most effective way to share build caches across large distributed teams and multiple CI machines.

---

### 273. How do you handle Dockerfile complexity at scale?
"At scale, Dockerfiles become complex. Strategies:

**Modular base images**: create organization-specific base images with common tooling, security configurations, and non-root users:
```dockerfile
FROM company/base-python:3.11  # Internal base with company defaults
```

**Dockerfile linting with Hadolint** in CI: catches common mistakes automatically.

**Templates with ERB or Jinja**: generate Dockerfiles programmatically for many similar services.

**Docker Bake with HCL**: `docker-bake.hcl` defines build matrices (multiple services, platforms) as code — version-controlled, reviewed, consistent."

#### In-depth
The 'golden base image' pattern is essential for large organizations: security teams define base images with hardened defaults (non-root user, seccomp profiles, up-to-date OS packages, security tools), and application teams build FROM these bases. When a CVE requires updating the base, one change to the golden base (+ rebuild and redistribution) propagates to all application images on their next build. Without this pattern, each team independently manages base images — creating inconsistency and slow CVE remediation times.

---

### 274. How do you use Docker for local development environments?
"Docker enables reproducible local dev environments that match production closely.

Dev environment patterns:

**1. Bind mount for hot reload**:
```yaml
services:
  api:
    build: { context: ., target: dev }  # Dev stage with dev tools
    volumes:
      - ./src:/app/src  # Changes reflect immediately
    command: npm run dev  # Watcher rebuilds on change
```

**2. dev containers (VS Code)**: `.devcontainer/devcontainer.json` defines a container as the full development environment. VS Code runs inside the container — same runtime as production.

**3. Compose profiles**: `--profile dev` starts local DBs, mail catchers, mock services."

#### In-depth
Dev containers (`.devcontainer`) are increasingly popular because they solve the 'works on my machine' problem at the development environment level. The container becomes the dev environment: same Node/Python/Go version, same OS, same extension and tool versions. New team members start contributing in minutes (just open the repo in VS Code and accept 'Reopen in Container') instead of spending half a day setting up their environment. GitHub Codespaces runs dev containers in the cloud — no local setup at all.

---

### 275. How do you use Docker in a microservices architecture?
"Docker is the natural deployment unit for microservices. Each microservice:

- Has its own Dockerfile (decoupled build)
- Builds to its own image (independently deployable)
- Gets its own container (isolated runtime)
- Communicates via HTTP/gRPC/messaging (well-defined APIs)
- Has its own data store (database-per-service pattern)

In Compose for local dev: all services in one `docker-compose.yml` for easy orchestration. In production (Swarm/K8s): services deployed individually, with independent scaling, versioning, and failure domains."

#### In-depth
The microservices antipattern with Docker: sharing volumes or bind mounts between services as the communication mechanism. Services should communicate via APIs or message queues — shared volumes create coupling. Each service's data (database, files) should be owned by that service alone. The Docker networking layer (service discovery via DNS) provides the right abstraction: services know each other by name, not by shared filesystem paths.

---

### 276. What is Docker Content Trust and cosign?
"**Docker Content Trust (DCT)**: Docker's original image signing system based on The Update Framework (TUF) and Notary. Enabled with `DOCKER_CONTENT_TRUST=1`. Signs and verifies image tags using cryptographic keys stored in `~/.docker/trust/`.

**cosign**: newer, Sigstore-based image signing tool. Works with OIDC-based keyless signing (identity-based, no long-lived keys). Signatures stored as OCI artifacts in the registry alongside the image.

cosign keyless signing in CI:
```bash
cosign sign --yes myregistry/app:v1.0
cosign verify --certificate-identity github-actions@myorg.github.com myregistry/app:v1.0
```"

#### In-depth
cosign keyless signing works by: CI pipeline authenticates with GitHub/Google OIDC → Fulcio CA issues a short-lived certificate bound to the OIDC identity → cosign uses this certificate to sign the image → the transparency log (Rekor) records the signing event → Rekor provides a verifiable audit trail without managing long-lived keys. This solves the hardest problem in signing: key management. The short-lived certificates expire in minutes — no key rotation, no key compromise risk.

---

### 277. How do you sign Docker images and verify them at deployment?
"**Signing**:
```bash
# Using cosign keyless
cosign sign --yes myregistry/app:$SHA

# Using cosign with key file
cosign sign --key cosign.key myregistry/app:$SHA
```

**Verification** at deployment (Kubernetes):

**Policy Controller** (Sigstore): install in K8s, define which images must be signed by which identities. Unsigned images rejected at admission.

```yaml
apiVersion: policy.sigstore.dev/v1beta1
kind: ClusterImagePolicy
spec:
  images: [{glob: 'myregistry/app:**'}]
  authorities: [{keyless: {url: 'https://fulcio.sigstore.dev'}}]
```"

#### In-depth
Image signature verification at the admission controller level is the gold standard for supply chain security. An attacker who compromises the CI/CD pipeline can push malicious images to the registry — but cannot forge a signature from the legitimate OIDC identity without compromising the OIDC provider. Requiring signatures from a specific GitHub Actions workflow identity: `github-actions@<org>.github.com` via OIDC ensures only your CI pipeline can produce deployable images.

---

### 278. What is a Software Bill of Materials (SBOM) for containers?
"An SBOM is a **complete inventory of all software components** in a container image: OS packages, language packages (npm, pip, go), and their versions.

Generate with Docker buildx:
```bash
docker buildx build --sbom=true --push -t myapp:latest .
```

Or with Syft:
```bash
syft packages docker:myapp:latest -o spdx-json > sbom.json
```

SBOM format standards: **SPDX** (Linux Foundation), **CycloneDX** (OWASP).

Use cases: vulnerability tracking (which images contain log4j?), license compliance, supply chain transparency."

#### In-depth
SBOMs have become regulatory requirements in certain industries. The US government's Executive Order on Cybersecurity (14028) requires SBOMs for software sold to federal agencies. In practice, SBOMs enable: rapid impact assessment when a new CVE is disclosed ('which of our 200 images contain log4shell?'), license audit ('are we distributing GPL code in our commercial product?'), and SBOM attestation as proof of diligence in security audits. Store SBOMs in your OCI registry alongside the image using cosign's `attest` command.

---

### 279. How do you prevent supply chain attacks in Docker pipelines?
"Multi-layered defense for supply chain security:

1. **Pin all dependencies by digest**: base images by SHA256, npm/pip by lockfile hash
2. **Scan base images before use**: trivy/Snyk scan as the first build step
3. **Use private mirrors**: mirror Docker Hub + npm/pip to an internal registry — reduces dependency on external availability
4. **Sign all images** with cosign and verify before deploy
5. **Generate and attest SBOMs** for auditability
6. **Use minimal base images**: smaller surface = fewer vulnerable packages
7. **Reproducible builds**: detect if a 'same-version' package produces different outputs (registry poisoning)"

#### In-depth
The SolarWinds and event-stream attacks demonstrated how supply chain compromise works: modify a widely-used package, wait for downstream projects to update their dependencies, collect credentials/data from running code. Docker-specific mitigations: never pull from Docker Hub in CI without caching (rate limits force using mirrors which you control), pin all `FROM` to digests (prevents tag-squatting attacks where an attacker overwrites a tag after you've tested against it), and use Go module proxy / npm Enterprise for internal package mirrors you control.

---

### 280. How do you implement Docker in a zero-trust architecture?
"Zero-trust: **verify everyone, trust no one**, regardless of network location.

Docker implementation:

1. **mTLS between services**: every container authenticates to every other container. Tools: Consul Connect, Istio mTLS, or SPIFFE/SPIRE for identity.
2. **No implicit trust based on network location**: containers on the same bridge network can't talk unless explicitly allowed (use internal networks + explicit connections)
3. **Secret injection via identity**: containers get secrets based on their SPIFFE SVID (identity), not their network location. Vault + SPIFFE agent auto-issues short-lived credentials.
4. **Container identity**: cosign-signed images + Pod identity (Kubernetes ServiceAccounts with IAM roles)"

#### In-depth
SPIFFE (Secure Production Identity Framework For Everyone) and SPIRE (its implementation) mint X.509 certificates (SVIDs) for each workload based on its identity (namespace, service account, pod labels). These SVIDs enable mTLS without managing PKI manually. Vault's auto-auth with Kubernetes can issue database credentials, API keys, and TLS certs automatically to containers just by verifying their Kubernetes identity — no static secrets, no secret sprawl. This closes the hardest gap in Docker security: credential delivery.

---

### 281. What are Kubernetes ephemeral containers and how do they relate to Docker?
"Kubernetes ephemeral containers are temporary containers added to a **running pod** for debugging purposes — without restarting the pod or modifying its spec.

```bash
kubectl debug -it pod/myapp --image=busybox --target=myapp
```

This adds a `busybox` container to the running `myapp` pod, sharing its process namespace, network, and volumes. You can inspect the namespace, run diagnostic tools, check files, and debug — then delete the ephemeral container.

This mirrors `docker exec` but for distroless containers that have no shell — you attach a debug container to inspect a shell-less production container."

#### In-depth
Ephemeral containers solve the debug-vs-security tradeoff for distroless images. Distroless containers are maximally secure (no shell, no utilities, tiny attack surface) but impossible to `exec` into for debugging. Ephemeral containers let you attach `gcr.io/google-cloud-tools/busybox` (or even a custom debug image with your specific tools) to the running distroless container, share its PID namespace (see all processes), and diagnose issues. Then remove the ephemeral container and the production container is back to its secure state.

---

### 282. How does Docker differ when run on CoreOS/Flatcar Linux?
"CoreOS (now Flatcar Container Linux) is an immutable, container-optimized OS designed specifically for Docker/container workloads.

Key differences vs. standard Linux:
- **Immutable root filesystem**: the OS itself is read-only. Updates are atomic OS swaps via A/B partition scheme.
- **No package manager**: you can't `apt-get install` anything. All tools run in containers.
- **Auto-updates**: OS updates are applied automatically and silently.
- **Optimized**: minimal OS (~300MB), only the Docker runtime and SSH pre-installed.
- **Ignition config**: describes the desired OS state (units, users, files) applied only on first boot.

Docker runs the same — just the host OS is hardened and immutable."

#### In-depth
Flatcar is ideal for Swarm clusters and Kubernetes nodes. The immutability eliminates configuration drift — you can't accidentally install something on a node that makes it different from others. Forced auto-updates mean nodes always run recent kernel versions (important for container security). The Ignition config approach (declarative OS configuration applied at first boot) combines with Terraform for fully automated node provisioning: `terraform apply` creates a Flatcar VM with a specific Ignition config → node joins the Swarm cluster automatically.

---

### 283. What is Buildpacks and how does it compare to Dockerfiles?
"**Cloud Native Buildpacks** (CNBs) automatically detect your language and create optimized container images — without a Dockerfile.

```bash
# Install pack CLI
pack build myapp --builder paketobuildpacks/builder:base
```

Pack detects: Go module → uses Go buildpack, Node.js package.json → uses Node buildpack, Python requirements.txt → uses Python buildpack. It handles: caching, layer optimization, metadata, base image management.

**vs Dockerfile**: Buildpacks are opinionated (less control) but automatic (no Dockerfile maintenance). Great for polyglot platforms and developer self-service. Dockerfiles give fine-grained control but require maintenance."

#### In-depth
Buildpacks are the default for Heroku, Cloud Foundry, and Google Cloud Run's source-to-image flow. The advantage: Buildpack maintainers handle OS security patches — when a CVE hits the runtime layer, the buildpack updates automatically. Teams using Buildpacks don't need to know about Docker or manage base image updates — they just push code. The trade-off: less control over the final image. Buildpacks produce OCI-compliant images that work with any container runtime, so you're not locked into the platform.

---

### 284. How do you effectively use Docker in a data science workflow?
"Data science Docker patterns:

**1. Jupyter container with pinned packages**:
```dockerfile
FROM python:3.11-slim
RUN pip install numpy==1.26 pandas==2.1 scikit-learn==1.3 jupyter==7.0
WORKDIR /notebooks
CMD ['jupyter', 'notebook', '--ip=0.0.0.0', '--no-browser', '--allow-root']
```

**2. GPU support** (for model training):
```bash
docker run --gpus all nvidia/cuda:12.0-cudnn8-devel myml:latest
```

**3. Reproducibility**: pin EVERY package version in requirements.txt. The container ensures the same environment months later.

**4. Large model files**: mount a volume for models, don't bake them into images."

#### In-depth
GPU containers require the **NVIDIA Container Toolkit** (`nvidia-container-toolkit`) installed on the host. It adds a custom OCI hook that gives containers access to the GPU device. The `--gpus all` flag exposes all GPUs; `--gpus device=0` exposes specific GPU 0. For ML training in the cloud: prebuilt GPU-optimized base images (AWS Deep Learning Containers, NVIDIA NGC catalog) include CUDA, cuDNN, and popular frameworks — starting from these saves hours of setup time.

---

### 285. How do you handle database migrations in Docker environments?
"Database migration strategies:

**1. Init container** (Kubernetes): run migrations as an init container before the app starts.

**2. Startup check in entrypoint**:
```bash
#!/bin/sh
./wait-for db:5432
./run-migrations.sh
exec "$@"  # Start the main application
```

**3. Separate migration job** (preferred): run migrations as a `docker run --rm` job in CI before deploying the new app version.

**Best practice**: use the expand-contract pattern — migrations are backward-compatible. Old app + new schema works. New app + new schema works. This enables blue-green deployments and canary releases."

#### In-depth
The separate migration job approach is most robust. CI pipeline: (1) run `docker run migrate:v2 ./migrate up`, (2) verify migration success, (3) deploy app v2. If migration fails, deployment doesn't happen. If app deployment fails, rollback to v1 (which still works with the migrated schema, because migrations are backward-compatible). The entrypoint approach has a race condition: multiple app replicas spinning up simultaneously all try to run migrations, requiring distributed locking (Flyway and Liquibase have advisory lock support).

---

### 286. How do you implement Circuit Breaker patterns in Dockerized microservices?
"Circuit breaker prevents cascade failures: if Service B is failing, Service A opens its circuit, stops calling B, and returns a fallback response immediately.

**Implementation options**:
1. **Application-level**: use resilience4j (Java), `go-resilience` (Go), `circuitbreaker` (Node.js) libraries. Each app instance manages its own circuit state.

2. **Sidecar proxy**: deploy Envoy or Haproxy as a sidecar container. Configure circuit breaker in the proxy — no app code changes needed.

3. **Service mesh** (Istio): `DestinationRule` with `outlierDetection` automatically ejects unhealthy hosts.

In Docker Compose: Envoy sidecar configuration via `envoy.yaml`."

#### In-depth
The sidecar circuit breaker pattern with Envoy works without modifying application code: traffic flows APP → ENVOY → TARGET. Envoy tracks success/failure rates and opens the circuit if error threshold is exceeded. The app continues making requests normally — Envoy returns local error responses without reaching the failing service, giving the failed service time to recover. This is especially powerful for languages without good resilience libraries (legacy PHP, older Node.js apps).

---

### 287. How do you implement distributed tracing for Dockerized services?
"Distributed tracing follows a request across multiple microservices.

**Setup**:
1. Deploy Jaeger or Zipkin as a sidecar/shared service: `docker run -d jaegertracing/all-in-one`
2. Instrument services using OpenTelemetry SDK: auto-instrumented for common frameworks (Flask, Express, Spring Boot)
3. Services send trace spans to the tracing collector via gRPC/HTTP

In Compose:
```yaml
services:
  jaeger:
    image: jaegertracing/all-in-one
    ports:
      - '16686:16686'  # Jaeger UI
  api:
    environment:
      OTEL_EXPORTER_OTLP_ENDPOINT: http://jaeger:4318
```"

#### In-depth
OpenTelemetry (OTel) is the industry standard for instrumentation — one SDK supports traces, metrics, and logs. Auto-instrumentation for many frameworks means zero code changes: run the app with the OTel agent. Traces reveal: which DB queries are slow, which downstream services add latency, where errors propagate. The correlation ID is propagated via HTTP headers (`traceparent`) between services — all services must participate for complete traces. Missing OTel in even one service breaks the trace at that hop.

---

### 288. How do you implement graceful shutdown in Docker containers?
"Graceful shutdown: when a container receives `SIGTERM`, it completes in-flight requests before exiting.

**Essential practices**:
1. **App handles SIGTERM**: implement signal handler that stops accepting new connections and finishes current requests
2. **Use exec form** for CMD/ENTRYPOINT: `CMD ['node', 'server.js']` — SIGTERM goes to node (PID 1), not to a shell
3. **Set `stop-signal`** if app uses a different signal: `STOPSIGNAL SIGUSR1`
4. **Set `stop_grace_period`** in Compose: give the app enough time to finish: `stop_grace_period: 30s`

After `stop_grace_period` expires, Docker sends SIGKILL — brutal kill."

#### In-depth
HTTP servers need drain time: `net/http` (Go) has `Shutdown(ctx)` with a context deadline — it stops accepting new connections, waits for in-flight requests to complete (or context deadline), then exits. The common mistake: using shell form (`CMD node server.js`) — Docker sends SIGTERM to `sh`, not to `node`. The shell doesn't forward signals by default, so node never receives SIGTERM and gets SIGKILL after the timeout. Always use exec form and verify with `docker stop && docker logs` to see 'graceful shutdown' logs appearing.

---

### 289. How do you manage Docker container logs in production?
"Production log management principles:

1. **Log to stdout/stderr** (not files): Docker captures stdout/stderr automatically. Apps should not write log files inside containers.

2. **Use a proper log driver**: ship logs to a centralized system at the Docker level — no need for log agents inside containers.

3. **Configure rotation** for the json-file driver: `{"log-opts": {"max-size": "10m", "max-file": "5"}}`

4. **Structured logging (JSON)**: machine-parseable logs are easier to query and filter in Kibana/CloudWatch/Datadog.

5. **Include context**: trace ID, service name, version in every log line — critical for distributed debugging."

#### In-depth
The twelve-factor app principle of treating logs as event streams is perfectly aligned with Docker's stdout logging. By logging to stdout, the container's logging behavior is configurable at runtime (change the Docker log driver without rebuilding the image). Production log pipeline: app → stdout → Docker log driver (fluentd/awslogs) → central log system (CloudWatch/Elasticsearch). Logs never touch the container's disk, preventing disk pressure from log accumulation. Alert on missing logs: a container that stops producing logs is usually unhealthy.

---

### 290. How do you audit Docker container activity?
"Auditing tracks who changed what and when in the container environment.

**Docker daemon audit**: Linux auditd rules for Docker socket:
```bash
-a always,exit -F arch=b64 -F path=/var/run/docker.sock -k docker-socket
-w /etc/docker -p wa -k docker-config
```

**Container runtime audit**: Falco monitors container activity (syscalls, file access, network) in real-time and raises alerts for suspicious behavior (`curl` executed in a prod container, /etc/shadow read, unexpected process).

**Swarm**: `docker service ls` audit via API logs. Kubernetes: API server audit logs, capturing every kubectl action with user identity, timestamp, and resource modified."

#### In-depth
Falco is the CNCF standard for container runtime security monitoring. It uses kernel eBPF probes to monitor every syscall from every container. Predefined rules detect: shell spawned in container (`docker exec`), unexpected network connections from production containers, sensitive file access, privilege escalation attempts. Falco rules can trigger: log to syslog (for SIEM), call a webhook (for PagerDuty), or block the syscall (enforcement mode). Pair Falco with a SIEM (Splunk, Elastic SIEM) for compliance audit trails.

---

### 291. What is the CRI (Container Runtime Interface) and how does Docker fit in?
"CRI (Container Runtime Interface) is the Kubernetes standard API for plugins between kubelet and container runtimes.

CRI defines: `CreateContainer`, `StartContainer`, `StopContainer`, `RemoveContainer`, `ListContainers`, `RunPodSandbox` (creates the pod network namespace), etc.

Docker does NOT implement CRI natively. Previously, Kubernetes used the `dockershim` (a CRI adapter for Docker). Since K8s 1.24, dockershim was removed. The current CRI implementations: **containerd** (with the `cri` plugin enabled) and **CRI-O** (a lightweight runtime for K8s only).

Docker images still work because images are OCI-compliant — the image format is separate from the runtime API."

#### In-depth
containerd's CRI plugin translates Kubernetes CRI calls to containerd API calls. The full chain: `kubelet → containerd CRI plugin → containerd → runc → container`. CRI-O is even more streamlined — it was designed specifically for Kubernetes without the extra Docker compatibility layer. For K8s clusters: containerd is the default runtime (EKS, GKE, AKS all default to containerd). CRI-O is common in OpenShift (Red Hat's K8s distribution). The runtime choice is transparent to application developers — only cluster admins care.

---

### 292. What is the SPIFFE/SPIRE framework for container identity?
"SPIFFE (Secure Production Identity Framework For Everyone) defines a standard for cryptographically verifiable workload identity.

SPIRE (its reference implementation) runs as: a Server (trust authority) + Node Agents (attestors on each host). SPIRE attests workload identity based on: node identity (AWS instance metadata, K8s node) + workload attributes (K8s service account, Docker container labels).

Each workload receives an **SVID** (SPIFFE Verifiable Identity Document) — an X.509 certificate. SVIDs enable mTLS between services without managing PKI.

```
Service A's SVID: spiffe://myorg.com/service-a
Service B's SVID: spiffe://myorg.com/service-b
Service A → mTLS → Service B: mutual identity verification
```"

#### In-depth
SPIFFE/SPIRE is the foundational layer for zero-trust in Kubernetes and Docker environments. When Service A connects to Service B, B verifies A's SVID (is it really service-a.myorg.com?) and A verifies B's SVID — no trust based on network location. SPIRE rotates SVIDs every hour by default — short-lived credentials limit the blast radius of a certificate compromise. Vault can use SPIRE attestation for secret delivery: only containers with a valid SVID for `spiffe://myorg.com/api` receive database credentials.

---

### 293. How do you implement observability in Docker environments?
"The three pillars of observability for Docker:

**Metrics (Prometheus + cAdvisor)**: per-container CPU, memory, network, disk I/O.
**Logs (structured JSON to centralized store)**: stdout → fluentd/filebeat → Elasticsearch/CloudWatch.
**Traces (OpenTelemetry + Jaeger)**: distributed request tracing across services.

Plus: **Events** (`docker events` stream) for lifecycle events — container starts/stops, image pulls, volume mounts.

The golden triangle correlation: link a trace ID from traces to log lines (structured logging with `trace_id` field) to identify exactly which logs belong to which traces. Sentry, Datadog, and Grafana Tempo provide this correlation out of the box."

#### In-depth
End-to-end observability requires buy-in from all service teams — consistently applying correlation IDs, structured logging, and OTel instrumentation. The most impactful single investment: correlation IDs. Generate a unique ID for each incoming request, propagate it via HTTP header (`X-Request-ID`) to all downstream service calls, and include it in every log line. This alone enables: reconstructing a transaction across 10 services from a single request ID, correlating user complaints to specific trace IDs, and understanding cascade failures.

---

### 294. What are Pod Disruption Budgets and how do they relate to Docker deployments?
"Pod Disruption Budgets (PDBs) are Kubernetes resources that **limit voluntary disruptions** (drains, rolling updates, node maintenance) to ensure a minimum number of pods remain available.

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
spec:
  maxUnavailable: 1  # At most 1 pod unavailable during disruption
  selector:
    matchLabels:
      app: myapi
```

Equivalent Docker Swarm concept: `--update-parallelism` and `--rollback-parallelism` in service definitions. Swarm limits simultaneous task disruptions during updates.

Use PDBs for any production service with strict availability requirements."

#### In-depth
PDBs are enforced by the Kubernetes eviction API. When `kubectl drain node` is run (for maintenance), it calls the eviction API for each pod. The eviction API checks PDBs — if evicting the pod would violate the PDB (too many already unavailable), the eviction is denied. The drain operation retries until the pod can be evicted safely. This prevents maintenance windows from accidentally taking down services by draining nodes too aggressively. Always pair PDBs with more than 1 replica — a PDB with `minAvailable: 1` and `replicas: 1` creates an undrain-able node.

---

### 295. How do you implement container network policies?
"Network policies control allowed traffic between pods/containers.

**Kubernetes NetworkPolicy**:
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
spec:
  podSelector:
    matchLabels:
      app: db
  ingress:
    - from:
      - podSelector:
          matchLabels:
            app: api  # Only api pods can connect to db
  policyTypes: [Ingress]
```

**Docker**: separate network per service tier (frontend network, backend network, db network). Containers on different networks can't communicate — explicit `network connect` required."

#### In-depth
Kubernetes NetworkPolicies are enforced by the CNI plugin (not all CNI plugins support them — Flannel doesn't, Calico and Cilium do). Without a policy-capable CNI, NetworkPolicy resources are created but not enforced. This is a silent failure mode — verify your CNI supports NetworkPolicies before relying on them for security. Cilium's eBPF-based enforcement provides both L3/L4 (pod-to-pod, IP-based) and L7 (HTTP, gRPC method-level) policies — far more granular than traditional iptables-based enforcement.

---

### 296. How do you implement automated rollback in Docker deployments?
"Automated rollback triggers when a deployment fails health checks or error rate spikes.

**Docker Swarm**: `--update-failure-action rollback` — automatic rollback if any task fails to start healthy. `docker service rollback myservice` for manual rollback.

**Kubernetes**: `kubectl rollout undo deployment/myapp` for manual rollback. With Argo Rollouts: automatic rollback based on Prometheus metrics:
```yaml
analysis:
  metrics:
    - name: success-rate
      thresholdRange:
        min: 95
      provider:
        prometheus:
          query: |
            rate(http_requests_total{status!~'5..'}[5m]) /
            rate(http_requests_total[5m])
```"

#### In-depth
Metric-based automated rollback (Argo Rollouts + Prometheus) is the gold standard. It answers the question engineers answer manually during deployments: 'is the error rate higher than baseline?' If yes, roll back automatically. The rollback decision is based on objective metrics, not human monitoring fatigue. Define your SLO (e.g., 99.5% success rate) and let automation enforce it during deployments. Combined with progressive delivery (canary → 10% → 50% → 100%), you catch problems before they affect all users.

---

### 297. How do you implement service discovery without Kubernetes?
"Service discovery without Kubernetes (standalone Docker or Swarm):

**Docker Swarm DNS**: built-in service discovery — `http://service-name` resolves automatically within overlay networks.

**Consul**: deploy Consul cluster, register each container on start, de-register on stop. Services query Consul API or DNS for healthy instances of other services.

**Registrator + Consul**: Registrator auto-registers Docker containers on start by watching Docker events and calling Consul API. No app changes needed.

**Traefik**: auto-discovers containers via Docker labels, maintains routing table for HTTP services."

#### In-depth
Registrator + Consul is the most widely deployed pattern for Docker service discovery outside of orchestrators. Registrator watches Docker events (`docker events`), when a container starts it reads labels/environment variables to determine the service name and port, then registers it in Consul. When the container stops, Registrator deregisters it. Services query `service.consul` DNS names to find healthy instances. Consul's health checking (HTTP endpoint checks) ensures only healthy instances receive traffic.

---

### 298. What are the limitations of Docker Compose for production?
"Docker Compose limitations in production:

1. **Single host**: all containers run on one machine. No multi-host load distribution or fault tolerance.

2. **No auto-healing across hosts**: if the host dies, all containers die. No automatic rescheduling on another host.

3. **Limited service discovery**: DNS works within a single Compose network, but no cross-host discovery.

4. **No built-in load balancing**: scale with `--scale api=3` but no external LB to route traffic to all 3.

5. **No secret rotation**: secrets must be rebuilt into .env files. No dynamic secret delivery.

6. **No resource scheduling**: doesn't bin-pack containers across nodes based on available resources."

#### In-depth
The production threshold: Docker Compose is appropriate for single-machine deployments where the business can tolerate the machine as a single point of failure. For true production (HA), you need Docker Swarm or Kubernetes. The decision criteria: if your service has an SLA requiring >99.9% uptime, you need multi-host orchestration. Many small SaaS applications run successfully on a single $50/month VM with Docker Compose — for them, the simplicity trade-off is correct. Know your SLA requirements before over-engineering.

---

### 299. What is Docker Hub's image pull rate limit and how do you work around it?
"Docker Hub rate limits (since November 2020):
- Anonymous pulls: 100 pulls per 6 hours per IP
- Authenticated free accounts: 200 pulls per 6 hours
- Pro/Team/Business: unlimited authenticated pulls

These limits cause CI pipeline failures when multiple jobs run simultaneously and exhaust the limit.

**Workarounds**:
1. **Authenticate in CI**: `docker login -u $DOCKERHUB_USER -p $DOCKERHUB_TOKEN` before pulls (200 limit instead of 100)
2. **Registry mirror**: configure a Docker Hub pull-through cache registry in `daemon.json`
3. **Use alternative registries**: pull official images from GitHub Container Registry (`ghcr.io/library/`) or Quay.io mirrors
4. **ECR Public**: AWS mirrors popular images at `public.ecr.aws/`"

#### In-depth
The registry mirror approach is most effective at scale. In `daemon.json`: `{"registry-mirrors": ["https://myharbor.internal.company.com"]}`. Harbor's proxy cache pulls from Docker Hub on first request, caches for 7 days, serves from cache on subsequent requests. CI instances hit the internal Harbor — no Docker Hub rate limits. Harbor also provides vulnerability scanning of cached images. For Kubernetes: set the Kubelet's `--pod-infra-container-image` and image registry config to use the internal mirror for all image pulls.

---

### 300. How do you implement disaster recovery for Docker-based applications?
"DR for Docker applications has several layers:

**Application data** (most critical):
- Databases: automated backups to S3/GCS with PITR (Point-in-Time Recovery)
- Volumes: snapshot schedules on cloud volumes (EBS snapshots, Azure Disk snapshots)

**Container state**:
- Images in registry: cross-region ECR replication
- Compose/Stack definitions: version-controlled in Git

**Infrastructure** (for VM-based Docker):
- Packer images for Docker hosts with pre-installed Docker
- Terraform for infra reproduction
- `docker stack deploy` restores the application stack in minutes

**RTO/RPO targets**: define recovery time and point objectives, test DR quarterly."

#### In-depth
The biggest DR misconception for Docker: thinking container immutability means you don't need DR. Containers are ephemeral — image layers can be rebuilt. But STATE (database, user data, uploaded files) is irreplaceable without backups. Test DR quarterly at minimum: actually delete a volume and restore from backup. The backup that has never been tested is not a backup — it's hope. Chaos engineering (intentional DR drills) using tools like Chaos Monkey or Gremlin validates DR procedures under realistic conditions.

---
