# 🔐 Docker Security — Advanced Interview Questions (Product-Based Companies)

This document covers security-focused Docker questions that are commonly asked at product-based companies. These interviews probe your understanding of container isolation boundaries, CVEs, kernel-level controls, and production hardening strategies.

---

### Q1: What is "container escape," and what kernel-level mechanisms does Docker rely on to prevent it?

**Answer:**
A container escape is when a process inside a container gains unauthorized access to the host filesystem, network, or process table — essentially breaking out of the isolation sandbox.

Docker relies on a **layered defence** of Linux kernel features:

1. **Namespaces:** Each container gets its own PID, NET, MNT, UTS, IPC, and USER namespace. A process in the container literally cannot *see* host-level resources even if it tries.
2. **Cgroups:** Prevent resource exhaustion attacks (fork bombs, memory bombs) from affecting the host.
3. **Seccomp (Secure Computing Mode):** Docker applies a default seccomp profile that **blocks ~44 dangerous syscalls** (e.g., `ptrace`, `kexec_load`, `perf_event_open`). An attacker who exploits a container-level bug can't make those syscalls to further escalate.
4. **AppArmor / SELinux:** Mandatory Access Control (MAC) frameworks that add per-process rules about which files, sockets, and capabilities a process can access, even if it's running as root inside the container.
5. **Linux Capabilities:** Docker drops many capabilities by default (e.g., `CAP_SYS_ADMIN`, `CAP_NET_RAW`). The principle: even if a process becomes root, it lacks the *capabilities* to reconfigure the host kernel.

**Common escape vectors that all these fight:** CVE-2019-5736 (runc), Dirty Cow, CVE-2022-0847 (Dirty Pipe).

---

### Q2: Explain Linux Capabilities in the context of Docker. What does `--cap-drop ALL --cap-add NET_BIND_SERVICE` mean in practice?

**Answer:**
Traditional Unix splits privilege into "root" and "non-root." Root can do everything. Linux **capabilities** split root's powers into ~40 distinct units, each granting specific powers:

- `CAP_NET_BIND_SERVICE` — Bind to ports < 1024
- `CAP_SYS_ADMIN` — Mount filesystems, manipulate namespaces, many syscalls (the "god" capability)
- `CAP_CHOWN` — Change file ownership
- `CAP_KILL` — Send signals to any process
- `CAP_NET_RAW` — Use raw/packet sockets (e.g., ping)
- `CAP_DAC_OVERRIDE` — Bypass filesystem permission checks

**Docker's defaults:** Docker starts containers with a reduced set — it drops `CAP_SYS_ADMIN`, `CAP_NET_ADMIN`, `CAP_SYS_MODULE`, etc., by default.

**`--cap-drop ALL --cap-add NET_BIND_SERVICE` in practice:**
1. `--cap-drop ALL`: Remove **every** capability from the container — the process cannot mount, kill foreign processes, change ownership, use raw sockets, etc.
2. `--cap-add NET_BIND_SERVICE`: Re-grant only the single capability needed to bind the web server to port 80 or 443.

This is the **principle of least privilege** in action. Even if an attacker exploits the app, they gain a shell in a process with almost zero kernel-level power. Compare this to running without `--cap-drop ALL`, where the process inherits `CAP_NET_RAW` and can craft ARP spoofing packets.

---

### Q3: What is a Seccomp profile in Docker, and how would you write a custom one for a Go microservice?

**Answer:**
Seccomp is a Linux kernel feature that acts as a syscall filter. A Docker Seccomp profile is a **JSON whitelist/blacklist of syscalls** that the kernel will or will not allow for a given container.

Docker ships with a hardened default profile that blocks ~44 syscalls (the `seccomp-default.json`). For a production Go microservice, you can go further.

**How to profile which syscalls your app actually needs:**
1. Run `strace -f -c ./your-binary` locally.
2. Run in a privileged container with `--security-opt seccomp=unconfined` and audit syscalls with `auditd`.
3. Tools like `oci-seccomp-bpf-hook` can generate a profile automatically.

**Custom profile snippet:**
```json
{
  "defaultAction": "SCMP_ACT_ERRNO",
  "architectures": ["SCMP_ARCH_X86_64", "SCMP_ARCH_AARCH64"],
  "syscalls": [
    {
      "names": [
        "read", "write", "open", "close", "fstat", "mmap", "mprotect",
        "munmap", "brk", "pread64", "accept4", "listen", "bind",
        "connect", "socket", "getsockname", "getpeername", "setsockopt",
        "getsockopt", "clone", "futex", "nanosleep", "epoll_wait",
        "epoll_ctl", "epoll_create1", "eventfd2", "pipe2", "signalfd4",
        "exit", "exit_group", "rt_sigaction", "rt_sigprocmask"
      ],
      "action": "SCMP_ACT_ALLOW"
    }
  ]
}
```

**Apply it:**
```bash
docker run --security-opt seccomp=./my-service-seccomp.json my-go-service:latest
```

**Effect:** If the compromised process attempts `ptrace`, `kexec_load`, or `perf_event_open`, the kernel immediately returns `EPERM` instead of executing it, making kernel exploitation dramatically harder.

---

### Q4: What is Rootless Docker, and why is it architecturally significant compared to standard Docker?

**Answer:**
**Standard Docker** runs a daemon (`dockerd`) as **root** on the host. Even though containers are isolated, the Docker socket (`/var/run/docker.sock`) is root-owned. Anyone who can write to this socket can effectively escalate to root on the host (mount `/`, run privileged containers, etc.).

**Rootless Docker** runs the entire Docker daemon — and all containers it spawns — under a **non-root user** using:
- **User Namespaces:** The container's "root" (UID 0) is mapped to the unprivileged user's UID on the host (e.g., UID 1000).
- **`slirp4netns`** (or `pasta`): A userspace networking stack that doesn't require `CAP_NET_ADMIN`.
- **`fuse-overlayfs`:** A FUSE-based overlay filesystem for non-root environments where kernel overlay requires privileges.

**Why it matters architecturally:**
| Aspect | Standard Docker | Rootless Docker |
|---|---|---|
| Daemon runs as | root | non-root user |
| Socket compromise | Full host root access | Limited to that user |
| CVE blast radius | Host-wide | Contained to user namespace |
| Networking | Fast kernel bridge | Slightly slower (userspace) |
| Shared namespaces | Native | Via user namespace mapping |

**Trade-offs:** Some capabilities like AppArmor integration, `--network=host`, and certain storage drivers are unavailable or limited in rootless mode. Companies like Cloudflare and Shopify use rootless Docker or similar approaches in multi-tenant CI environments.

---

### Q5: How would you scan Docker images for vulnerabilities in a CI/CD pipeline, and what are the key vulnerability databases used?

**Answer:**
Image scanning is the practice of analysing container image layers against known CVE databases. The key tools and their approaches:

**Tools:**
- **Trivy (Aqua Security):** Scans OS packages (Alpine apk, Debian dpkg, RHEL rpm) AND language packages (Go modules, npm, Maven, PyPI). Fast, CI-friendly.
- **Grype (Anchore):** Similar scope, uses NIST NVD + GitHub Advisory DB.
- **Snyk Container:** Cloud-based, integrates with Docker Hub and registries.
- **Docker Scout:** Native Docker integration for vulnerability scanning.
- **Clair:** Open-source, designed for self-hosted registry integration.

**Vulnerability Databases:**
- **NVD (National Vulnerability Database)** — NIST maintained, canonical CVE data.
- **GitHub Advisory Database** — Language ecosystem CVEs for npm, Go, PyPI, Maven.
- **Alpine SecDB, Debian Security Tracker** — Distro-specific package-level CVE data.

**CI/CD pipeline integration (GitLab CI example):**
```yaml
scan-image:
  stage: security
  image: aquasec/trivy:latest
  script:
    - trivy image --severity HIGH,CRITICAL --exit-code 1 myapp:$CI_COMMIT_SHA
  allow_failure: false
```

**Shift-left strategy:**
1. Scan the **base image** at Dockerfile linting time.
2. Scan the **built image** before pushing to the registry.
3. Scan **running images** periodically for newly disclosed CVEs on images already in production (runtime scanning).

**Key interview point:** Distinguish between *fixing vulnerabilities* (update base image) vs. *accepting/deferring* them (CVE triage). Not every HIGH CVE is exploitable in your context.

---

### Q6: What is the Docker Content Trust (DCT) mechanism, and how does Notary work under the hood?

**Answer:**
**Docker Content Trust (DCT)** is an image signing and verification framework. When enabled (`DOCKER_CONTENT_TRUST=1`), Docker will **refuse to pull or run any image that doesn't have a valid cryptographic signature**.

**Under the hood — The Update Framework (TUF) + Notary:**

DCT is built on the [Notary](https://github.com/notaryproject/notary) project, which implements TUF (The Update Framework):

1. **Root Key:** The publisher's master key stored offline (air-gapped ideally). Used to sign all other keys.
2. **Targets Key:** Signs the actual image metadata (image digest ↔ tag mapping).
3. **Snapshot Key:** Signs the current state of all targets (prevents replay attacks).
4. **Timestamp Key:** Short-lived, rotated frequently, prevents replay of stale metadata.

**Flow:**
```
Publisher push:
  docker push myrepo/myimage:v1.2
  Notary signs the (tag → sha256 digest) mapping with the Targets key
  Signature stored in Notary server

Consumer pull:
  docker pull myrepo/myimage:v1.2
  Docker contacts Notary server
  Verifies signature chain: Timestamp → Snapshot → Targets → Root
  If valid → pull proceeds; if tampered → pull rejected
```

**Why this matters:** It prevents supply chain attacks where an attacker compromises the registry and replaces a legitimate image with a backdoored one. The attacker can't forge signatures without the private key.

**Modern alternative:** **Sigstore/Cosign** is rapidly becoming the industry standard, using OCI artifact specs to store signatures alongside images in the registry itself, without a separate Notary server.

---

### Q7: A security audit reveals your production Docker containers are running as root. Walk me through the remediation process.

**Answer:**
Running as root in production is a significant security anti-pattern. Here's the systematic remediation:

**Step 1: Add a USER directive to your Dockerfiles**
```dockerfile
FROM node:20-alpine

# Create a non-root user with a specific UID
RUN addgroup --system --gid 1001 appgroup && \
    adduser --system --uid 1001 --ingroup appgroup --no-create-home appuser

WORKDIR /app
COPY --chown=appuser:appgroup package*.json ./
RUN npm ci --only=production
COPY --chown=appuser:appgroup . .

# Switch to non-root before the final CMD
USER appuser
EXPOSE 3000
CMD ["node", "server.js"]
```

**Step 2: Fix file ownership issues**
Files copied before the `USER` switch must be explicitly `--chown`'d. Files written at runtime (logs, uploads) must be in directories the non-root user owns.

**Step 3: Handle privileged port binding**
If your app binds port 80, either:
- Use `CAP_NET_BIND_SERVICE` capability specifically, or
- Bind to port 3000 internally and use a reverse proxy (nginx/Traefik) or load balancer for port 80 externally.

**Step 4: Verify with:**
```bash
docker run --rm myimage:latest whoami
# Should return: appuser (not root)
```

**Step 5: Add pod-level security in Kubernetes:**
```yaml
securityContext:
  runAsNonRoot: true
  runAsUser: 1001
  readOnlyRootFilesystem: true  # Also mount /tmp as emptyDir if needed
  allowPrivilegeEscalation: false
```

**Step 6: Audit running containers:**
```bash
docker ps -q | xargs -I{} docker exec {} id
# Should show uid=1001, not uid=0
```

---

*Prepared for Security Architecture and Platform Engineering interviews at product-based companies.*
