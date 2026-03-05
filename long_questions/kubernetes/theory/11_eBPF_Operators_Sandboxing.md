# 🔬 Advanced Kubernetes Theory — eBPF, Operators & Container Sandboxing

---

## 🔹 Section 1: eBPF in the Linux Kernel

### What is eBPF?

**eBPF (extended Berkeley Packet Filter)** is a Linux kernel subsystem that lets you run safe, sandboxed programs in kernel context without changing kernel source code. Think of it as a "JavaScript engine for the kernel" — a JIT-compiled, verifier-checked execution environment.

**The verification guarantee:**
Before execution, every eBPF program passes through the **eBPF Verifier**, which statically analyzes the bytecode to prove:
- No infinite loops (bounded loops only, since Linux 5.3)
- No out-of-bounds memory access
- All stack frames properly sized
- Register types correct (no reading uninitialized memory)

If verification fails, the program is rejected. This is what makes eBPF programs safe to run in kernel space.

---

### eBPF Hook Points

```
User Space App
     │ syscall
     ▼
─────────────────────────────────────────────────── Kernel Boundary
     │
     ├── [kprobe/kretprobe] ← Attach to any kernel function (entry/exit)
     ├── [tracepoints]      ← Predefined stable kernel instrumentation points  
     ├── [syscall hooks]    ← Intercept specific syscalls
     │
     ├── Kernel Network Stack
     │     ├── [XDP]        ← In NIC driver, before sk_buff allocation
     │     ├── [TC ingress] ← After sk_buff, before routing
     │     ├── [TC egress]  ← After routing, before leaving NIC
     │     └── [socket]     ← At the socket layer (sendmsg/recvmsg)
     │
     └── [cgroup hooks]     ← Per-cgroup (per-pod) network/resource control
```

---

### eBPF Maps — The Data Plane

Maps are the primary way eBPF programs store state and communicate with userspace:

| Map Type | Use Case | Access |
|---|---|---|
| `HASH` | Service VIP → Backend pods lookup | O(1) get/set |
| `ARRAY` | Per-CPU counters, fixed indexed data | O(1), lock-free with per-CPU |
| `LRU_HASH` | Connection tracking, recent flow cache | O(1), auto-evicts oldest |
| `PERF_EVENT_ARRAY` | Stream events to userspace (Falco alerts, traces) | Ring buffer |
| `PROG_ARRAY` | Tail calls — eBPF programs calling other programs | Max 33 calls deep |
| `CGROUP_STORAGE` | Per-cgroup (per-pod) state | Hierarchical |

---

### eBPF in Kubernetes — Where It's Used

| Tool | Hook Used | Purpose |
|---|---|---|
| **Cilium** | TC, XDP, socket | kube-proxy replacement, L7 NetworkPolicy |
| **Falco** | kprobe, tracepoints | Runtime syscall monitoring (security) |
| **Pixie** | kprobe, uprobe | Zero-instrumentation distributed tracing |
| **Hubble** | TC (via Cilium) | Network flow observability |
| **Bpftrace** | kprobe, tracepoints | Ad-hoc kernel debugging/profiling |

---

## 🔹 Section 2: Kubernetes Operator Pattern — Conceptual Deep Dive

### The Reconciler Mental Model

Every Kubernetes controller — including your Operators — follows the same mental model:

```
Desired State (CR Spec)  ─────┐
                               ├──► Reconciler ──► Actions ──► Actual State
Actual State (external)  ─────┘
```

The Reconciler is called whenever:
1. A watched resource changes (CREATE, UPDATE, DELETE)
2. A `RequeueAfter` timer fires
3. A resource the controller **owns** changes (via `Owns()`)

**Golden rules of reconciliation:**
- **Idempotent**: Running the same reconcile loop 100 times must produce the same outcome as running it once.
- **Level-triggered, not edge-triggered**: Don't react to individual events — react to the current state. If you miss 3 events, the 4th reconcile must still reach the correct state.
- **Never assume order**: Events can arrive out of order or be duplicated.

---

### Operator SDK vs. controller-runtime vs. kubebuilder

| Tool | What It Is | Use When |
|---|---|---|
| `controller-runtime` | The core library. Used by most K8s ecosystem projects. | Core primitives — direct when you want control |
| `kubebuilder` | CLI scaffolding on top of controller-runtime | Fastest way to generate CRD+controller boilerplate |
| `Operator SDK` (Red Hat) | Wraps kubebuilder; also supports Ansible/Helm operators | Teams that want batteries-included tooling |
| `kopf` (Python) | Python framework for operators | Data science teams comfortable with Python |

---

### Status vs Spec — The Contract

| | **Spec** | **Status** |
|---|---|---|
| Set by | User / CI pipeline | Controller only |
| Represents | Desired state | Observed/actual state |
| Updated via | `r.Update(ctx, obj)` | `r.Status().Update(ctx, obj)` |
| Stored separately | No | Yes (via `/status` subresource) |

Using a status subresource means a user with only `update` permission on the resource CANNOT update `.status` — only the controller (with `update` on the `/status` subresource) can. This prevents users from lying about status.

---

### Finalizers — Clean Up Before Deletion

When a CR is deleted, Kubernetes doesn't immediately remove it if it has **finalizers** (a list of strings in `metadata.finalizers`). The controller gets a chance to clean up external resources first.

```go
const finalizer = "mycompany.io/backup-cleanup"

func (r *DatabaseBackupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    backup := &mycompanyv1.DatabaseBackup{}
    r.Get(ctx, req.NamespacedName, backup)

    // Object is being deleted
    if !backup.DeletionTimestamp.IsZero() {
        if controllerutil.ContainsFinalizer(backup, finalizer) {
            // Run cleanup — delete S3 bucket contents, DB connections, etc.
            if err := r.cleanupExternalResources(ctx, backup); err != nil {
                return ctrl.Result{}, err
            }
            // Remove finalizer to allow Kubernetes to delete the object
            controllerutil.RemoveFinalizer(backup, finalizer)
            r.Update(ctx, backup)
        }
        return ctrl.Result{}, nil
    }

    // Object is not being deleted — add finalizer if not present
    if !controllerutil.ContainsFinalizer(backup, finalizer) {
        controllerutil.AddFinalizer(backup, finalizer)
        r.Update(ctx, backup)
    }

    // ... normal reconciliation
}
```

---

## 🔹 Section 3: Container Sandboxing — Security Depth

### Threat Model — Why Standard Containers Aren't Enough

Standard Linux containers use **namespaces** (isolate process view) and **cgroups** (limit resources), but they share the **host kernel**. A kernel exploit (like Dirty COW, CVE-2016-5195) inside a container can directly compromise the host.

**Threat scenarios requiring stronger isolation:**
1. **CaaS (Container-as-a-Service)**: You let customers submit Docker images to run on your infrastructure. A malicious image could escape.
2. **CI/CD runner pods**: Running untrusted PR code in ephemeral pods — common attack vector.
3. **Multi-tenant SaaS**: Different enterprise customers' workloads running on shared nodes.

---

### The Isolation Spectrum

```
←── Less secure, faster ──────────────────────────── More secure, slower ──→

 runc         gVisor (KVM)        Kata Containers        Full VM (EC2)
  │                │                    │                     │
Namespaces    User-space kernel    Per-pod full VM       Separate machine
 + cgroups   (syscall interp.)   (hardware isolation)
  ~0%           ~15% overhead       ~10% overhead         100% overhead
```

---

### gVisor Internal Architecture

```
Container App
     │
     │ Linux syscalls (open, read, connect, execve...)
     ▼
┌─────────────────────────────────────────┐
│  gVisor Sentry (user-space process)     │
│  ┌────────────────────────────────────┐ │
│  │ Platform: KVM mode or ptrace mode  │ │
│  │ Implements Linux kernel interface  │ │
│  │ in Go                              │ │
│  └────────────────────────────────────┘ │
│  ┌────────────────────────────────────┐ │
│  │ Gofer: file system proxy           │ │
│  │ (mediates all file I/O)            │ │
│  └────────────────────────────────────┘ │
└─────────────────────────────────────────┘
     │
     │ ~50 allowed host syscalls (futex, mmap, etc.)
     ▼
Host Linux Kernel
```

**Key point:** The Sentry itself is a userspace process. Even if the container exploits a bug in the Sentry, it only compromises that single Sentry instance — not the host kernel or other pods.

---

### Pod Overhead declaration (for scheduler awareness)

RuntimeClasses declare overhead so the **scheduler accounts for the sandbox cost**:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: gvisor
handler: runsc
overhead:
  podFixed:
    memory: "120Mi"
    cpu: "250m"
```

When you schedule a pod with `runtimeClassName: gvisor` that requests `500m CPU / 256Mi memory`, the scheduler treats it as needing `750m CPU / 376Mi memory` (adding the overhead). This prevents the node from being overcommitted just because the sandbox overhead was hidden.

---

### Choosing the Right Sandbox

```
Is syscall compatibility 100% required?
    ├── YES → Use Kata Containers (full VM, real Linux kernel)
    └── NO →  Is overhead critical (high-frequency, latency-sensitive)?
                  ├── YES → Reconsider: use runc + strong admission policies + Falco monitoring
                  └── NO  → Use gVisor (good balance of security vs. performance)

Special cases:
  - Running nested containers/K8s → Kata
  - Financial services (PCI-DSS VM boundary) → Kata
  - General multi-tenant SaaS → gVisor
  - Internal trusted workloads → runc (fastest)
```
