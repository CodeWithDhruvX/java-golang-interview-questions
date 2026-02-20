# 🔐 **Docker Security & Best Practices (41–50)**

---

### 41. How do you secure a Docker container?
"Security is a layered approach. My checklist: **run as non-root**, **drop capabilities**, **use read-only filesystem**, **limit resources**, and **scan images**.

In practice: `docker run --user 1001 --cap-drop=ALL --cap-add=NET_BIND_SERVICE --read-only --memory=256m myapp`. Each flag incrementally reduces the attack surface.

Beyond runtime, I use minimal base images (distroless, Alpine), scan with **Docker Scout** or **Trivy**, and enforce signing with **Docker Content Trust**."

#### In-depth
The principle is **least privilege**. A container should have only what it needs to function. `--cap-drop=ALL` removes all Linux capabilities and `--cap-add` adds back only what's required. This means even if the app is compromised, the attacker has severely restricted kernel access — they can't load kernel modules, change network config, etc.

---

### 42. What is Docker Content Trust?
"Docker Content Trust (DCT) is a security feature that uses **cryptographic signatures** to verify the integrity and publisher of Docker images.

When enabled (`export DOCKER_CONTENT_TRUST=1`), Docker will only pull and run images that are signed. Unsigned images are rejected, preventing supply chain attacks.

Publishers sign images with private keys managed by **Notary**. Users verify signatures against the Notary server's root of trust. I enforce DCT in our CI/CD and production environments."

#### In-depth
DCT is built on **The Update Framework (TUF)**, a specification for securing software update systems. It protects against: image tampering in transit (MITM), unauthorized image pushes, and tag re-pointing attacks. However, adoption is low in practice — many official images on Docker Hub aren't signed, limiting real-world applicability.

---

### 43. What are Docker secrets?
"Docker secrets are a secure way to store and distribute **sensitive data** (passwords, API keys, certificates) to containers in a Swarm cluster.

Secrets are encrypted at rest (using the Raft log encryption) and in transit (TLS). They're only mounted in memory (`tmpfs`) inside the container at `/run/secrets/<secret-name>` — they never touch disk.

I create them with `docker secret create db_password ./password.txt` and reference them in a Swarm service config: `secrets: - db_password`."

#### In-depth
Docker secrets are a Swarm-only feature. In Compose (non-Swarm), `secrets` is supported but uses bind mounts — not the same secure delivery mechanism. For non-Swarm environments, I use external secret managers (HashiCorp Vault, AWS Secrets Manager) and inject secrets as environment variables at runtime via entrypoint scripts.

---

### 44. How do you scan Docker images for vulnerabilities?
"I use **Trivy** or **Docker Scout** to scan images for known CVEs in OS packages and application dependencies.

In CI: `trivy image myapp:latest --exit-code 1 --severity CRITICAL,HIGH`. This fails the pipeline if critical or high vulnerabilities exist. I scan in three places: on every image build in CI, on a schedule for images already in the registry, and before deployments.

Docker Scout is built into Docker Desktop and integrates with Docker Hub — it shows vulnerability counts on `docker scout quickview`."

#### In-depth
Image scanning checks a Software Bill of Materials (SBOM) — a manifest of all packages, libraries, and their versions — against vulnerability databases (NVD, OSV, vendor advisories). Key insight: most vulnerabilities come from the base OS layer. Switching from `ubuntu:latest` to `gcr.io/distroless/static` often eliminates 80%+ of CVEs because distroless has no shell or OS utilities for attackers to exploit.

---

### 45. What are best practices for writing Dockerfiles?
"My top practices:

1. **Use specific base image tags** (never `latest`)
2. **Order instructions from least to most frequently changing** (maximize cache)
3. **Combine related RUN commands** with `&&` to minimize layers
4. **Use multi-stage builds** to keep production images small
5. **Don't run as root** — add a non-root USER
6. **Use `.dockerignore`** to keep the build context lean
7. **No secrets** in Dockerfile or image layers — use BuildKit secrets or runtime injection

A well-crafted Dockerfile is reproducible, minimal, secure, and fast to build."

#### In-depth
Layer caching is the key to fast builds. The golden rule: instructions that change infrequently (install OS packages, install dependencies) must come BEFORE instructions that change frequently (copy application code). In Node.js: `COPY package.json .` + `RUN npm install` BEFORE `COPY . .`. This way dependency installation is cached unless `package.json` changes.

---

### 46. Why should you use non-root users in containers?
"By default, processes in Docker containers run as **root (UID 0)**. If an attacker exploits the application and breaks out of the container, they have root privileges on the host — catastrophic.

Adding a non-root user: `RUN addgroup -S app && adduser -S app -G app` + `USER app` means an exploited container gives the attacker only a low-privilege process on the host.

Many Kubernetes environments (PSP, OPA Gatekeeper, Pod Security Admission) now **require** non-root containers in production."

#### In-depth
Running as non-root also prevents the process from accessing files owned by root on bind-mounted volumes — a useful accidental protection. However, beware of the opposite problem: if the container runs as UID 1000 but the host files are owned by UID 0, the container can't write to them. UID mapping and volume permission alignment is a common DevOps headache.

---

### 47. What are the risks of running privileged containers?
"A **privileged container** (`--privileged`) runs with nearly all Linux capabilities enabled and with direct access to host devices.

The risks: it can mount host filesystems, load kernel modules, modify network rules (iptables), and access `/dev`. Essentially, a process that escapes a privileged container has full host root access — nullifying all container isolation.

I only use privileged containers for specific infrastructure tools (like a container that manages Docker itself — Docker-in-Docker) and only in isolated, trusted environments."

#### In-depth
The classic container escape CVE-2019-5736 (runc vulnerability) was exacerbated by privileged containers. A non-privileged container leak would yield a low-privilege shell; a privileged container escape yields full root. Security benchmarks (CIS Docker Benchmark) flag privileged containers as a critical finding in audits.

---

### 48. What is image hardening?
"Image hardening means **reducing the attack surface** of a Docker image by removing everything unnecessary.

Steps: start from a minimal base (Alpine, distroless), remove package managers, shells, and debug tools after installation, set strict file permissions, run as non-root, mark the filesystem read-only, and pin all dependency versions.

The philosophy: if the attacker can't find a shell, can't find `curl`, and can't write files — their options after exploiting the app are severely limited."

#### In-depth
**Distroless images** (from Google) take hardening to the extreme: they contain only the application runtime and its dependencies — no package manager, no shell, no system utilities. `gcr.io/distroless/java` contains just the JVM. Debugging distroless containers requires `docker debug` (with ephemeral debug containers) rather than `exec -it /bin/sh`.

---

### 49. How do you limit container resources?
"I use runtime flags for CPU and memory limits:

- `--memory=512m` — hard memory limit (container is OOM killed if exceeded)
- `--memory-reservation=256m` — soft limit / best-effort
- `--cpus=0.5` — limit to 50% of one CPU core
- `--cpu-shares=512` — relative weight (default 1024)
- `--pids-limit=100` — prevent fork bombs

In Compose/Swarm: `deploy.resources.limits.memory: 512M` and `cpus: '0.5'`."

#### In-depth
Without memory limits, a single container with a memory leak can cause the entire host to OOM, killing random processes including other containers. Setting limits with `--memory` enables the kernel's OOM killer to target only that container. The `--oom-kill-disable` flag (which disables OOM killing) is extremely dangerous and should never be used in production.

---

### 50. How can you prevent container breakout?
"Container breakout prevention is defense-in-depth:

1. **Don't run privileged** — `--privileged` is the #1 enabler of breakout
2. **Drop capabilities** — `--cap-drop=ALL`
3. **Use seccomp profiles** — restrict syscalls the container can make
4. **Use AppArmor/SELinux** — mandatory access control policies
5. **Read-only root filesystem** — `--read-only`
6. **No `/var/run/docker.sock` mounts** — gives container full Docker API access
7. **Use rootless Docker** or **gVisor/Kata Containers** for stronger isolation at the runtime layer"

#### In-depth
Mounting `/var/run/docker.sock` is the most common accidental privilege escalation — any container with this socket can spawn privileged containers, mount host filesystems, etc. It's essentially container escape by design. If you need Docker access in a container (CI runners), use **Docker-in-Docker** with a separate `dockerd` instead of socket mounting.

---
