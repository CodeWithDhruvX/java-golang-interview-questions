## 🧠 Docker Expert-Level Concepts (Questions 301-310)

### Question 301: What is the role of Linux namespaces in Docker containers?

**Answer:**
Namespaces provide the fundamental **isolation** of containers. They partition kernel resources such that one set of processes sees one set of resources while another sees a different set.
- **PID:** Process isolation.
- **MNT:** Mount points / Filesystem.
- **NET:** Network stack.
- **IPC:** Inter-process communication.
- **UTS:** Hostnames.
- **USER:** User IDs.

---

### Question 302: How do control groups (cgroups) work in Docker?

**Answer:**
Cgroups provide **resource limitation and accounting**.
- **Hierarchy:** Organized in a tree structure.
- **Controllers:** CPU, Memory, BlkIO, PIDs.
- Docker creates a cgroup for each container to enforce `docker run --cpus 1 --memory 1g`.

---

### Question 303: Explain PID, network, and mount namespaces in containerization.

**Answer:**
- **PID:** The container's INIT process is PID 1. It cannot see Host PID 1 or other containers' PIDs.
- **Network:** Container has its own `eth0`, IP, Routing table, and localhost (`lo`), distinct from Host.
- **Mount:** Container has its own root filesystem (`/`). A `mount` inside container does not affect Host.

---

### Question 304: How does Docker handle IPC between containers?

**Answer:**
By default, each container has its own IPC namespace (System V IPC, POSIX message queues are isolated).
**Sharing:**
`docker run --ipc=container:other_container ...`
Allows shared memory segments (SHM) for high-performance communication.

---

### Question 305: Can Docker containers run GUI apps? How?

**Answer:**
Yes.
1.  **X11 Forwarding:** Share X11 socket.
    `docker run -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=$DISPLAY ...`
2.  **VNC/RDP:** Run a VNC server inside the container.

---

### Question 306: What is the lifecycle of a Docker image from source to deployment?

**Answer:**
1.  **Code:** Developer commits code.
2.  **Build:** `docker build` reads Dockerfile, pulls base layers, executes instructions, creates new layers.
3.  **Tag:** Identify version.
4.  **Push:** Upload manifest + blobs to Registry.
5.  **Pull:** Deploy target downloads manifest + blobs.
6.  **Run:** Extracts layers via UnionFS, starts container.

---

### Question 307: How does Docker enforce isolation between containers on the same host?

**Answer:**
Primarily via **Kernel Namespaces** (Hiding resources) and **Seccomp/AppArmor** (Preventing access to kernel). Cgroups ensure one container doesn't starve others.

---

### Question 308: Can a container change its own cgroup settings?

**Answer:**
Generally **No**.
The cgroup filesystem (`/sys/fs/cgroup`) is mounted Read-Only (RO) in most secure configurations.
If mounted RW (or `--privileged`), it can modify its own limits (Security risk).

---

### Question 309: What happens when a Docker container forks too many processes?

**Answer:**
It can trigger a **Fork Bomb**, exhausting Host PIDs.
**Prevention:**
`docker run --pids-limit 100` uses the PIDs cgroup controller to limit process count.

---

### Question 310: How do you inspect a container's internal Linux kernel view?

**Answer:**
Since the Kernel is shared:
- `uname -a` inside container shows Host Kernel version.
- `ls /proc` inside container shows process info for *that namespace*.
- `sysctl -a` shows kernel parameters (some namespaced, some global).

---

## 🔐 Advanced Security Scenarios (Questions 311-320)

### Question 311: What are the risks of mounting `/var/run/docker.sock` inside a container?

**Answer:**
**Critical Risk.**
If a container has access to the socket, it can talk to the Docker Daemon. It can start a new `--privileged` container, mount the Host Root Filesystem, and take over the entire machine.
*Common in "Docker-in-Docker" CI setups (DooD).*

---

### Question 312: How do you implement rootless containers in Docker?

**Answer:**
**Rootless Mode:** Run the Docker Daemon itself as a non-root user.
- Uses `user_namespaces` to map user to root inside container.
- Prevents root access to host even if container breakout occurs.
- Requires `dockerd-rootless-setuptool.sh`.

---

### Question 313: How does Docker use `seccomp` profiles to prevent syscalls?

**Answer:**
Docker passes a BPF program to the kernel that filters syscalls.
If a container tries `kexec_load` (load new kernel), seccomp intercepts and denies it (`EPERM`).

---

### Question 314: What is a container escape vulnerability? Give an example.

**Answer:**
A bug allowing code to break out of the namespace isolation.
**Example:** "Shocker" (Old), runC CVE-2019-5736.
Often involves overwriting host binaries via `/proc/self/exe` or using kernel exploits (Dirty COW).

---

### Question 315: How do you prevent container breakout via `/proc` or `/sys`?

**Answer:**
Docker mounts generic `/proc` and `/sys` as Read-Only paths.
It masks sensitive paths (e.g., `/proc/kcore`, `/proc/sys`).
*Do not unlock these with `--privileged`.*

---

### Question 316: How do you implement container user namespace remapping?

**Answer:**
Configure `daemon.json`:
```json
{ "userns-remap": "default" }
```
Docker creates a user `dockremap`. Container UID 0 maps to Host UID 100000.

---

### Question 317: How can containers be isolated by AppArmor profile enforcement?

**Answer:**
AppArmor enforces mandatory access control on *file paths*.
A profile can say: "This container can only read `/etc/nginx` and write `/var/log/nginx`. Touch nothing else."

---

### Question 318: How do you ensure least-privilege containers in production?

**Answer:**
1.  User: Non-root.
2.  Filesystem: Read-only.
3.  Caps: Drop ALL, add specific.
4.  Network: Network Policy (Deny all ingress/egress by default).

---

### Question 319: What is a privileged container and why is it dangerous?

**Answer:**
`--privileged` gives access to all **Devices** (`/dev`).
Allows loading Kernel Modules.
Essentially a processes running on Host Root, just with a different mount namespace.

---

### Question 320: How do you configure Docker to disallow specific kernel capabilities?

**Answer:**
`docker run --cap-drop=NET_RAW ...`
(Prevents packet crafting/spoofing).

---

## 🔁 Edge-Case Dockerfile Practices (Questions 321-330)

### Question 321: How do you use conditional logic in Dockerfiles?

**Answer:**
Dockerfile language DOES NOT support `if/else`.
**Workarounds:**
1.  Run shell logic: `RUN if [ "$BUILD_ENV" = "prod" ]; then rm ...; fi`.
2.  Multi-stage with Build Args.

---

### Question 322: What are `.dockerignore` best practices?

**Answer:**
1.  Ignore `.git` (Huge size).
2.  Ignore `Dockerfile` and `docker-compose.yml`.
3.  Ignore secrets (`.env`, `*.key`).
4.  Ignore local build artifacts (`node_modules`, `target/`).

---

### Question 323: How do you prevent secrets from being baked into Docker images?

**Answer:**
**NEVER** use `COPY .env .` or `RUN export KEY=...`.
Use **BuildKit Secrets** (`--mount=type=secret`). The secret is available only during the `RUN` command execution and not saved in the layer.

---

### Question 324: How do you install a private SSH key in a secure Dockerfile build?

**Answer:**
```dockerfile
# Dockerfile
RUN --mount=type=ssh git clone ...
```
Build:
`docker build --ssh default .`
(Forward local SSH agent safely).

---

### Question 325: What happens when you COPY a symlink in Docker?

**Answer:**
`COPY` resolves symlinks by copying the **file content** it points to (dereference), NOT the link itself.

---

### Question 326: Why is ordering of instructions critical in Dockerfiles?

**Answer:**
**Cache invalidation.**
If Line 3 changes, Lines 4-100 must be rebuilt.
Place volatile instructions (COPY src) LAST. Place stable instructions (Install OS deps) FIRST.

---

### Question 327: Can `CMD` and `ENTRYPOINT` be combined? How?

**Answer:**
Yes.
- `ENTRYPOINT`: The executable (`/bin/my-app`).
- `CMD`: Default arguments (`--help`).
User runs `docker run image --version`.
Result: `/bin/my-app --version` (`--version` overrides CMD).

---

### Question 328: What are best practices for creating language runtime images?

**Answer:**
1.  Use official upstream images.
2.  Keep them slim.
3.  Install only runtime dependencies (no compilers in final image).
4.  Set standard environment variables (`PYTHONUNBUFFERED=1`).

---

### Question 329: How do you write a Dockerfile for a polyglot (multi-language) app?

**Answer:**
**Multi-stage:**
- Stage 1: Golang Builder.
- Stage 2: Node Builder.
- Stage 3: Ubuntu Runtime.
  - `COPY --from=0 /bin/app .`
  - `COPY --from=1 /build/static .`

---

### Question 330: How do you set up `.docker/config.json` for secure auth?

**Answer:**
Use **Credential Helpers** (`docker-credential-helper`).
The config file stores a reference to the system keychain (Keychain on Mac, WinCred on Windows) instead of base64 encoded passwords.

---

## 🧩 Container Networking (Niche & Deep) (Questions 331-340)

### Question 331: How do you assign a static IP to a container?

**Answer:**
Only possible on user-defined networks.
```bash
docker network create --subnet 172.20.0.0/16 mynet
docker run --net mynet --ip 172.20.0.5 nginx
```

---

### Question 332: What is the default subnet range used by Docker bridge?

**Answer:**
Usually `172.17.0.0/16`.
Docker increments the second octet for new networks (172.18, 172.19...).

---

### Question 333: How do you create an isolated VLAN for container networks?

**Answer:**
Use the **MACVLAN** driver with sub-interfaces (e.g., `eth0.10`).
It tags packets with VLAN ID 10 directly on the wire.

---

### Question 334: How do you resolve DNS inside Docker containers?

**Answer:**
1.  Container queries embedded DNS (127.0.0.11).
2.  Docker checks internal service names.
3.  If not found, forwards to Host's DNS (`/etc/resolv.conf`).

---

### Question 335: What is IP masquerading in Docker NAT?

**Answer:**
When container traffic leaves the host (to Internet), Docker uses `iptables` MASQUERADE.
It rewrites the Source IP (172.x) to the Host IP (192.x) so the return traffic can find its way back.

---

### Question 336: How do Docker containers talk to services running on host OS?

**Answer:**
1.  **Mac/Win:** `host.docker.internal` DNS name.
2.  **Linux:** Use the Docker Bridge Gateway IP (usually `172.17.0.1`).
3.  **Host Network:** `--net host` (Linux only).

---

### Question 337: How does Docker manage egress and ingress traffic?

**Answer:**
- **Ingress:** Port mapping (DNAT). `Host:80 -> Container:80`.
- **Egress:** Masquerading (SNAT). `Container -> Internet`.
- **Policy:** By default, all egress allowed.

---

### Question 338: What are overlay networks in multi-host Docker setups?

**Answer:**
Uses **VXLAN** encapsulation.
Creates a Virtual L2 network over Layer 3 infrastructure.
Packet Structure: `[ Host IP [ VXLAN Header [ Container IP [ Payload ] ] ] ]`.

---

### Question 339: What is the difference between `docker network inspect` and `docker container inspect`?

**Answer:**
- **Container Inspect:** Shows settings for ONE container (IP, Gateway, Mac).
- **Network Inspect:** Shows the subnet definition and LIST of all containers attached to it.

---

### Question 340: How does Docker interact with firewalls (iptables, ufw)?

**Answer:**
Docker manipulates `iptables` directly.
**Conflict:** UFW (Ubuntu Firewall) might block routing. Docker inserts rules *before* UFW chains often bypassing UFW rules unless configured carefully (`DOCKER-USER` chain).

---

## 📦 Image and Artifact Management (Questions 341-350)

### Question 341: How do you remove dangling and orphaned Docker images?

**Answer:**
`docker image prune`.
Orphaned images (dangling) are layers that have no tag associated.

---

### Question 342: What’s the difference between image ID, digest, and tag?

**Answer:**
- **Tag:** Mutable alias (`v1`, `latest`).
- **Image ID:** Local hash of config (Computed locally).
- **Digest:** Content-addressable hash found in Registry manifest (Global unique ID).

---

### Question 343: How do you inspect image history and layers?

**Answer:**
`docker history <image_name>`.
Shows the size and command used to create each layer.

---

### Question 344: How can you squash layers in a Docker image?

**Answer:**
**Squashing:** Merging multiple layers into one.
- **Experimental:** `docker build --squash`.
- **Why:** To reduce size and hide intermediate files.
- **Why Not:** Breaks caching (users must download the whole blob again).

---

### Question 345: What are image labels and how are they used in CI/CD?

**Answer:**
Metadata (Key-Value pairs).
`LABEL maintainer="me@corp.com" build_date="2023-01-01"`.
Used by automation tools to identify owners or expiry dates.

---

### Question 346: What tools analyze image bloat?

**Answer:**
**Dive.**
Shows efficiency score (e.g., 90%).
Identifies files that are overwritten or deleted in subsequent layers (wasted space).

---

### Question 347: How can you scan images for license violations?

**Answer:**
Tools like **Snyk**, **FOSSA**, or **Black Duck**.
They analyze the packages (npm, pip, apk) installed in the image and report license types (GPL, MIT).

---

### Question 348: What is a reproducible image build?

**Answer:**
A build that produces the exact same bit-for-bit Image Digest every time.
Requires: Fixed timestamps, deterministic package installation order. Hard to achieve perfectly.

---

### Question 349: How do you ensure determinism in Docker image builds?

**Answer:**
1.  **Lock Files:** `package-lock.json`, `go.sum`.
2.  **Base Images:** SHA-pinned (`ubuntu@sha256:...`).
3.  **Repositories:** Use internal mirror (snapshot of apt repo) instead of public `apt-get update`.

---

### Question 350: Can you rename a Docker image and preserve history?

**Answer:**
`docker tag old:tag new:tag`.
It doesn't rename/move data. It just adds a new pointer (Alias) to the same Image ID. History is preserved because it IS the same image.
