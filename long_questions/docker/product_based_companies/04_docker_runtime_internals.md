# ⚙️ Docker Runtime Internals — containerd, runc & the OCI Ecosystem (Product-Based Companies)

This document covers Docker's container runtime stack — a topic deeply tested at product-based companies building platforms, PaaS systems, or operating at infrastructure scale.

---

### Q1: Describe the full container runtime stack from `docker run` to a running process. What is the role of each component?

**Answer:**
When you execute `docker run`, a chain of processes is orchestrated. The Docker CLI itself does very little:

```
User: docker run nginx
  ↓
Docker CLI (docker) 
  → HTTP API call to Docker Daemon
  ↓
Docker Daemon (dockerd)
  → Handles image pulling, Dockerfile parsing, volume/network setup
  → Delegates container lifecycle to containerd
  ↓
containerd (High-level runtime)
  → Manages container snapshots (image layers), distribution (pulling/pushing)
  → Manages container lifecycle (start, stop, pause, delete)
  → Uses snapshotter plugins (overlayfs, btrfs, native)
  → Delegates actual container creation to a shim
  ↓
containerd-shim-runc-v2 (Shim process per container)
  → Acts as the parent process for the container
  → Prevents containerd from being the container's parent (critical for restarts without containerd restart)
  → Handles stdio piping from container to containerd
  ↓
runc (Low-level runtime, OCI runtime)
  → Reads the OCI Runtime Spec (config.json)
  → Makes kernel syscalls: clone() with namespace flags, mount(), pivot_root()
  → Sets up cgroups, drops capabilities, applies seccomp filter
  → Execs the container's init process
  → runc exits after container process starts
  ↓
Container process (PID 1 in the new namespace — e.g., nginx)
```

**Why the shim matters:** If containerd restarts, containers don't die because the shim (which is the actual OS-level parent) is still running. The shim re-attaches to containerd when it comes back up.

---

### Q2: What is the OCI (Open Container Initiative) specification, and why was it created?

**Answer:**
Before OCI, Docker was the only viable container runtime and defined the de-facto standard unilaterally. As Kubernetes grew, alternative runtimes (rkt, LXC-based) emerged, and the ecosystem needed standardization.

**OCI was founded in 2015 (Linux Foundation)** with two core specifications:

**1. OCI Image Specification:**
Defines what a container image *is* — the format for layer tarballs, a manifest JSON describing them, and a configuration JSON with entrypoint/env/labels. Any tool producing an OCI-compliant image can be run by any OCI-compliant runtime.

```
image/
├── blobs/sha256/
│   ├── abc123...   # Config JSON (entrypoint, env, labels)
│   ├── def456...   # Layer 1 tarball
│   └── ghi789...   # Layer 2 tarball  
├── index.json      # Manifest list (multi-platform)
└── oci-layout      # OCI version marker
```

**2. OCI Runtime Specification:**
Defines how to take an OCI image "bundle" (extracted filesystem + `config.json`) and *run it*. The spec mandates:
- How namespaces are created
- How mounts are set up
- How capabilities are applied
- How hooks (prestart, poststart, poststop) work

**runc is the OCI reference runtime.** Any tool that produces a `config.json` + rootfs and calls `runc run` will get a container. This is what enables Kubernetes to swap between containerd, CRI-O, or other runtimes without changing Pod specs.

**The CRI (Container Runtime Interface):**
Kubernetes talks to runtimes via gRPC using the CRI spec (separate from OCI). containerd and CRI-O both implement CRI. containerd uses an OCI runtime (runc) internally.

---

### Q3: What is `overlay2` and how do union filesystems enable Docker's layered image model?

**Answer:**
Docker images are composed of **read-only layers** stacked on top of each other. When a container runs, a thin **read-write layer** is added on top. The union filesystem (UFM) merges these into a single coherent view for the container process.

**`overlay2` (OverlayFS) — Docker's default storage driver:**

OverlayFS uses two directories:
- **Lower dir(s):** Read-only image layers (can be many, stacked)
- **Upper dir:** Read-write container layer
- **Work dir:** OverlayFS internal use for atomic operations
- **Merged view:** The unified view presented to the container

```
Merged view (container sees):
├── /etc/nginx/nginx.conf  ← from image layer 3
├── /usr/bin/nginx         ← from image layer 2
├── /lib/...               ← from image layer 1 (base OS)
└── /app/runtime.log       ← written to Upper (container layer)
```

**Copy-on-Write (CoW) mechanics:**
- When a container reads a file from a lower layer, it reads it directly (no copy).
- When a container **writes** to an existing file from a lower layer, OverlayFS **copies the file up** to the upper layer first, then modifies it there. The original lower layer is unchanged.
- When a container **deletes** a file from a lower layer, OverlayFS creates a "whiteout" file in the upper layer, which hides the lower layer file.

**Performance implication:** The first write to a large file from a lower layer has a one-time CoW overhead (copy-up latency). For databases writing large data files, this is why **Docker Volumes** (which bypass OverlayFS entirely and write directly to the host filesystem) are used instead.

**`overlay2` directory structure on disk:**
```bash
ls /var/lib/docker/overlay2/
# Each directory is a layer:
# <sha256-hash>/
#   ├── diff/        # The actual filesystem changes for this layer
#   ├── link         # Short ID (OverlayFS has a path length limit)
#   ├── lower        # Colon-separated IDs of parent layers
#   └── work/        # OverlayFS work dir
```

---

### Q4: How does containerd manage image snapshotting, and what is the difference between a snapshot and an image layer?

**Answer:**
**Image Layer:** A read-only, content-addressed tarball stored in the containerd content store (CAS — Content Addressable Storage). Identified by its SHA256 hash. Layers are immutable and shared across images.

**Snapshot:** A prepared, unpacked, filesystem representation of a layer stack — ready to be mounted. A snapshot is what actually gets mounted when a container starts.

**containerd's Snapshot flow:**
```
Image Pull:
  1. Download layer tarballs → store in /var/lib/containerd/io.containerd.content.v1.content/
  2. Unpack each layer as a snapshot in the snapshotter
  3. Each snapshot is a "child" of the previous one (chain)

Container Start:
  4. containerd asks the snapshotter for a "prepared" snapshot 
     (mounts the image layers as lower dirs)
  5. Creates an "active" snapshot with an upper/work dir for the container's R/W layer
  6. Mounts the overlay mount → BundleDir → runc → process
  
Container Stop:
  7. Active snapshot → committed snapshot (saved as new image layer if docker commit)
  8. Or discarded (if just a container that's removed)
```

**Snapshotters available:**
- `overlayfs` (default Linux) — Fast, mature, requires kernel ≥ 4.0
- `native` — Simple copy-based (slow, for compatibility)
- `devmapper` — LVM thin provisioning (used historically for database containers)
- `btrfs` — CoW at filesystem level (great for clones/snapshots)
- `zfs` — Enterprise features (dedup, compression)

**Why this matters at scale:** At a company with hundreds of microservices sharing the same base image, containerd's content store deduplicates shared layers — a 30-node cluster pulling `node:20-alpine` only downloads those base layers once per node, not per container.

---

### Q5: What is `containerd-shim` and why does each container get its own shim process?

**Answer:**
The `containerd-shim` is a small stub process that is spawned **per container** by containerd. Its key responsibilities:

**1. Parent Process Independence:**
In Linux, if a parent process dies, its children become orphans and are adopted by PID 1. If containerd was the direct parent of container processes, restarting containerd would cause all containers to be re-parented to init — or worse, to die (depending on kernel version and signal handling).

The shim is the **container's OS-level parent process.** When containerd restarts, the containers' parent (the shim) stays alive. Containerd re-attaches to existing shims via abstract Unix domain sockets on restart.

**2. stdio Multiplexing:**
The shim buffers/relays the container's stdin/stdout/stderr to containerd and client terminals. Without the shim, detached containers (`docker run -d`) couldn't have their stdio managed after Docker CLI disconnects.

**3. Exit Status Collection:**
The shim calls `wait()` to collect the container process's exit code and relays it to containerd. This prevents zombie processes at the OS level.

**4. Enabling Live Restore:**
Docker's `--live-restore` flag (in `daemon.json`) uses shims to keep containers running even while `dockerd` is restarting (for daemon upgrades without downtime).

**Process tree example:**
```
systemd (PID 1)
└── dockerd (PID 812)
    └── containerd (PID 919)
        ├── containerd-shim-runc-v2 (PID 1040) ← Container A's shim
        │   └── nginx (PID 1055)               ← Container A's PID 1
        └── containerd-shim-runc-v2 (PID 1101) ← Container B's shim
            └── node server.js (PID 1116)      ← Container B's PID 1
```

---

### Q6: What is `gVisor` and how does it compare to standard runc for tenant isolation in multi-tenant environments?

**Answer:**
**runc** (standard): When a container process makes a syscall (e.g., `read()`), it goes **directly to the host kernel**. The isolation is provided by namespaces and seccomp, but ultimately the same kernel is shared. A kernel exploit in the container can affect the host.

**gVisor** (Google): Implements a **user-space kernel** — a Go process called the "Sentry" that intercepts **all** container syscalls and re-implements them in user space. The actual host kernel only sees a small, well-defined set of syscalls from the Sentry.

```
Standard:
Container Process → syscall → Host Linux Kernel

gVisor:
Container Process → syscall → gVisor Sentry (user-space kernel) → limited syscalls → Host Linux Kernel
```

**Two interception modes:**
- **ptrace mode:** Uses `ptrace` to intercept syscalls (compatible everywhere but slow — ~50-70% overhead)
- **KVM mode:** Uses the KVM hypervisor for efficient interception (~5-15% overhead)

**gVisor as an OCI runtime:** It implements the OCI Runtime Spec via `runsc`, so it's a drop-in replacement for `runc`:

```json
// /etc/docker/daemon.json
{
  "runtimes": {
    "runsc": {
      "path": "/usr/local/bin/runsc"
    }
  }
}
```
```bash
docker run --runtime=runsc -it alpine sh  # Runs in gVisor
```

**When to use gVisor:**
- Multi-tenant SaaS — running untrusted user code (e.g., Replit, CI as a service)
- Compliance-heavy environments needing kernel surface reduction
- Google Cloud Run uses gVisor for container isolation

**Trade-offs vs. runc:**
- Stronger isolation, but ~5-15% performance overhead
- Some syscalls not implemented (niche Linux features)
- File I/O can be slower depending on mode

---

*Prepared for Platform Engineering and Infrastructure Architecture interviews at product-based companies.*
