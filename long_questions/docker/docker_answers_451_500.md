## ⚡ Docker + Dev Environments (Questions 451-460)

### Question 451: How do you use Docker for local development with live code reloading?

**Answer:**
Mount the source code as a volume.
**Node.js Example:**
`docker run -v $(pwd):/app -w /app node:14 nodemon index.js`
Nodemon inside container watches the file changes on host (propagated via volume) and restarts the app.

---

### Question 452: How do you mount a local directory into a container with read-only access?

**Answer:**
`docker run -v $(pwd):/data:ro my-image`
Useful for configuration files you don't want the container to modify.

---

### Question 453: What is Docker Desktop's Dev Environments feature?

**Answer:**
Allows sharing reproducible coding environments.
Includes code + tools + editor setup. Allows "One-click" setup of dev environment from a Git URL.

---

### Question 454: How do you manage host-to-container volume sync on Windows?

**Answer:**
Windows uses WSL 2 backend.
**Best Performance:** Store code **inside** the WSL 2 filesystem (`\\wsl$\Ubuntu\home\user`).
Mounting from Windows C: drive (`/mnt/c`) is significantly slower due to file system translation 9P protocol.

---

### Question 455: How do you containerize a microservice for fast iteration?

**Answer:**
Use a "Development Dockerfile".
- Install debuggers.
- Use `CMD ["tail", "-f", "/dev/null"]` to keep it running.
- `docker exec` to run build/test manually.

---

### Question 456: How can Docker improve onboarding for dev teams?

**Answer:**
"Clone & Run".
`git clone repo` -> `docker-compose up`.
No need to install Python/Node/Postgres versions locally. Eliminates environmental drift.

---

### Question 457: How do you debug file permission issues when mounting volumes?

**Answer:**
Common issue: Host user (1000) != Container user (root or other).
**Fix:**
- Pass UID: `docker run -u $(id -u):$(id -g)`.
- Use Entrypoint to `chown` the volume directory (Slow).
- Use User Namespaces.

---

### Question 458: How do you persist PostgreSQL or Redis data locally?

**Answer:**
Use named volumes or local bind mounts in Compose.
```yaml
volumes:
  - ./postgres-data:/var/lib/postgresql/data
```
Data remains after `down`.

---

### Question 459: How do you use watch scripts with Docker for hot reload?

**Answer:**
Similar to nodemon.
- **Go:** `air`.
- **Python:** `flask run --reload`.
Run the watcher *inside* the container entrypoint.

---

### Question 460: How do you isolate dev environments with custom networks?

**Answer:**
`docker-compose --project-name client1 up`.
`docker-compose --project-name client2 up`.
Creates `client1_default` and `client2_default` networks.
They run identical stacks isolated from each other.

---

## 🐳 Docker + Serverless & MicroVMs (Questions 461-470)

### Question 461: What is Firecracker and how does it relate to containers?

**Answer:**
AWS open-source virtualization technology.
Creates lightweight MicroVMs.
Docker can use Firecracker via `containerd` to launch MicroVMs instead of Linux Containers for higher isolation (Multi-tenant security).

---

### Question 462: What is Docker’s position in serverless architecture?

**Answer:**
Docker allows "Packaged Functions".
Platforms like AWS Lambda now support deploying functions as **Container Images**.
Docker standardizes the packaging; Serverless standardizes the scaling/billing.

---

### Question 463: How are microVMs different from containers?

**Answer:**
- **Container:** Shared Kernel. Process Isolation.
- **MicroVM:** Own Kernel. Hardware Virtualization.
MicroVMs (Firecracker/Kata) offer VM-level security with Container-like speed.

---

### Question 464: What is Kata Containers and how does it improve security?

**Answer:**
An OCI runtime (like runc) but it launches KVMs.
If an attacker breaks out of the "Container" (VM), they are trapped in the Hypervisor, not on Host Kernel.

---

### Question 465: How do you run AWS Lambda locally with Docker?

**Answer:**
Use AWS SAM CLI or `docker-lambda` images.
`docker run -v $PWD:/var/task lambci/lambda:python3.8 handler.py`
Simulates the Lambda runtime environment.

---

### Question 466: What are sandboxed containers?

**Answer:**
Containers running with an extra layer of isolation.
Examples: gVisor (Google), Kata.
They intercept syscalls or use virtualization to prevent kernel exploits.

---

### Question 467: How can Docker be used to run serverless functions with openfaas?

**Answer:**
OpenFaaS (Function as a Service) runs on top of Docker Swarm/K8s.
It packages code as Docker containers and auto-scales them based on HTTP requests (=0 to N).

---

### Question 468: What’s the benefit of combining gVisor with Docker?

**Answer:**
**Defense in Depth.**
gVisor (runsc) acts as a user-space kernel.
If app tries to crash kernel, it only crashes gVisor sandbox.

---

### Question 469: Can Docker containers be orchestrated in FaaS platforms?

**Answer:**
Yes. Google Cloud Run and AWS Fargate are essentially "Container-as-a-Service" or Serverless Containers. You give them a container; they handle the rest.

---

### Question 470: What is OCI runtimes' role in supporting microVMs?

**Answer:**
The OCI spec allows swapping `runc` for `kata-runtime` or `runsc` (gVisor).
Docker doesn't care; it just says "Start Bundle". The Runtime decides *how* (Namespace vs VM).

---

## 🔐 Security Automation with Docker (Questions 471-480)

### Question 471: How do you run containers with read-only filesystems?

**Answer:**
`docker run --read-only`.
Force specific writable paths using `--tmpfs /run` or Volumes.
Prevents attackers from downloading/chmodding malware.

---

### Question 472: What is a security scanner in DockerHub?

**Answer:**
Automated Snyk scans on Official Images.
Checks binary signatures and package versions against CVE DB.

---

### Question 473: How do you configure `no-new-privileges` in Docker?

**Answer:**
`--security-opt no-new-privileges`.
Prevents processes from gaining more privileges (e.g., via `suid` binaries) during execution. Secure default.

---

### Question 474: What does `--cap-drop=ALL` do and when should you use it?

**Answer:**
Drops all Linux Capabilities.
Granularly add back only what is needed.
Best practice for high-security environments.

---

### Question 475: What’s the best way to rotate secrets inside a container?

**Answer:**
Containers are immutable.
**Don't rotate inside.**
Update the Secret in Orchestrator -> Redeploy/Restart Container.

---

### Question 476: How do you detect tampering in a running container?

**Answer:**
**Runtime Security Tools (Falco).**
Detects:
- Shell spawned in production.
- Modification of sensitive files (`/etc/passed`).
- Unexpected outbound network connection.

---

### Question 477: What is Docker content trust and how does it work?

**Answer:**
Uses The Update Framework (TUF).
Ensures the image digest matches the publisher's signature.
Prevents "Man in the Middle" attacks swapping image content.

---

### Question 478: How do you disable inter-container communication?

**Answer:**
On default bridge: `--icc=false`.
On Custom Network: Use Isolation segment or Network Policies (if K8s).

---

### Question 479: What are seccomp profiles and how do you customize them?

**Answer:**
JSON file defining "Whitelist of System Calls".
Customize by starting with default and removing calls (like `ptrace` or `mount`) or adding needed ones.

---

### Question 480: What is the role of Snyk in Docker security?

**Answer:**
Standard partner for `docker scan`. Requires login.
Provides remediations (e.g., "Upgrade base image to Alpine 3.14 to fix 5 Critical CVEs").

---

## 🌐 Hybrid and Multi-Platform Docker Use (Questions 481-490)

### Question 481: How do you build ARM images on x86 machines?

**Answer:**
**QEMU User Emulation.**
`docker buildx` automatically sets up QEMU binfmt interpreters.
Allows `RUN uname -m` inside build to return `aarch64`.

---

### Question 482: How do you emulate cross-architecture builds with QEMU?

**Answer:**
`docker run --rm --privileged multiarch/qemu-user-static --reset -p yes`.
Registers the QEMU interpreters in the Kernel.

---

### Question 483: How do you publish multi-platform images to registries?

**Answer:**
`docker buildx build --platform linux/amd64,linux/arm64 -t img:latest --push .`
Push creates a **Manifest List** pointing to both Blobs.

---

### Question 484: What is `platform=linux/arm64` used for?

**Answer:**
In `FROM`:
`FROM --platform=linux/arm64 python:3.9`
Forces Docker to pull the ARM version even if you are on Intel (runs via QEMU).

---

### Question 485: How do you run a Windows container on a Linux host?

**Answer:**
You generally **cannot**. Containers share kernels.
Linux Host = Linux Kernel. Windows Apps need Windows Kernel.
*Exception:* VM based isolation (like Hyper-V) can theoretically bridge this, but standard Docker does not.

---

### Question 486: What’s the difference between LCOW and WCOW?

**Answer:**
- **WCOW:** Windows Containers on Windows.
- **LCOW:** Linux Containers on Windows (via Hyper-V lightweight Linux VM).

---

### Question 487: Can you share volumes across OS platforms?

**Answer:**
In Docker Desktop (LCOW), yes. Windows files are mounted into Linux containers via CIFS/9P.
Performance is the bottleneck.

---

### Question 488: How do Docker Desktop and WSL2 work together?

**Answer:**
Docker Desktop installs 2 distributions:
1.  `docker-desktop`: Runs the daemon.
2.  `docker-desktop-data`: Stores images/containers.
It integrates with your default distro (`Ubuntu`) via socket mapping.

---

### Question 489: How do you configure Docker for multi-architecture builds in CI?

**Answer:**
1.  Set up QEMU.
2.  Set up Buildx builder instance.
    `docker buildx create --use`
3.  Build.

---

### Question 490: How do you test performance differences on ARM vs x86 containers?

**Answer:**
Build image for both (`--platform`).
Deploy on AWS Graviton (ARM) and AWS regular (x86).
Run benchmarks (`sysbench`).
*Often ARM is cheaper/faster for cloud workloads.*

---

## 📈 Docker Observability Tools (Questions 491-500)

### Question 491: How do you use `cadvisor` to inspect container metrics?

**Answer:**
Run:
`docker run -p 8080:8080 gcr.io/cadvisor/cadvisor`
Access UI at `localhost:8080`.
Shows real-time graphs for every running container (without config).

---

### Question 492: How do you integrate Docker metrics into Prometheus?

**Answer:**
Scrape config in `prometheus.yml`:
```yaml
scrape_configs:
  - job_name: 'cadvisor'
    static_configs:
      - targets: ['cadvisor:8080']
```

---

### Question 493: What is `logspout` and how does it work?

**Answer:**
A log router container. Reads raw logs from Docker socket and forwards them to Syslog/HTTP/Logstash.
`docker run -v /var/run/docker.sock:/var/run/docker.sock gliderlabs/logspout syslog://dest:514`

---

### Question 494: How do you configure Fluentd with Docker containers?

**Answer:**
Use the `fluentd` logging driver.
`docker run --log-driver=fluentd ...`
Requires Fluentd agent running on localhost:24224.

---

### Question 495: What’s the best way to tail logs from all containers at once?

**Answer:**
- `docker-compose logs -f` (If in same project).
- **Stern:** (K8s tool but useful concept).
- **Dozle:** Lightweight web-based log viewer for Docker.

---

### Question 496: How do you set log file rotation and retention in Docker?

**Answer:**
In `daemon.json`:
```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  }
}
```
Prevents logs from filling disk.

---

### Question 497: How do you analyze container startup time?

**Answer:**
1.  `docker events` (Timestamp of Create vs Start).
2.  App Logs (Time to "Listening on port 80").

---

### Question 498: How can you inspect container DNS resolution failures?

**Answer:**
Use `dntools` image or just `alpine`.
`docker run --rm -it alpine nslookup google.com`
Check `cat /etc/resolv.conf`.

---

### Question 499: How do you identify container resource starvation?

**Answer:**
Check **Throttling metrics**.
cAdvisor: `container_cpu_cfs_throttled_seconds_total`.
If high, the container wants more CPU than Limit allows.

---

### Question 500: How can OpenTelemetry be used with Docker containers?

**Answer:**
Run an **OTel Collector** as a sidecar or daemon.
Configure App to push Traces/Metrics to the Collector.
The Collector exports to Backend (Jaeger/Prometheus).
Standardizes observability across polyglot containers.
