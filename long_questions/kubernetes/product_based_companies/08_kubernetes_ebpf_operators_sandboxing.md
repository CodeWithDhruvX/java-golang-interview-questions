# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 8)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Stripe, Datadog, etc.
> **Focus deep-dive:** eBPF Internals & Cilium, Writing CRDs + Controllers (Operators), RuntimeClass & gVisor/Kata Sandboxing, Multi-Cluster Federation & Submariner.

---

## 🔹 eBPF Internals & Cilium (Questions 73-77)

### Question 73: What is eBPF and how does Kubernetes use it?

**Answer:**
**eBPF (extended Berkeley Packet Filter)** is a Linux kernel technology that lets you run sandboxed programs in the kernel without changing kernel source code or loading kernel modules.

**How it works:**
1. You write an eBPF program (usually in restricted C).
2. The **eBPF verifier** checks it — no infinite loops, bounded memory access, safe register writes.
3. The program is JIT-compiled to native machine code and attached to a **hook point** in the kernel.
4. Programs communicate with userspace via **eBPF maps** (key-value stores shared between kernel and userspace).

**Hook points relevant to Kubernetes:**

| Hook Type | Use Case |
|---|---|
| `XDP` (eXpress Data Path) | Highest-performance packet filtering — runs before the kernel network stack (DDoS mitigation, load balancing) |
| `TC` (Traffic Control) | Attach to network interfaces for per-packet policy enforcement |
| `kprobes` / `tracepoints` | Instrument kernel functions — used by Falco, Pixie for observability |
| `socket filters` | Intercept data going in/out of a socket — used for L7 policy |

**Kubernetes uses eBPF via Cilium:**
- Replaces kube-proxy entirely — Service VIPs resolved with eBPF maps instead of iptables rules.
- Scales to 100,000+ services without iptables chain bloat.
- Enforces NetworkPolicy at L3/L4/L7 (HTTP method, gRPC service name) using eBPF programs on TC hooks.

---

### Question 74: What are eBPF Maps and how does Cilium use them for service load balancing?

**Answer:**
**eBPF Maps** are shared memory structures (backed by the kernel) that allow data exchange between:
- **eBPF programs** running in the kernel
- **Userspace applications** (Cilium agent, kubectl-based tools)

**Common map types:**

| Map Type | Description |
|---|---|
| `BPF_MAP_TYPE_HASH` | Key-value lookup table — O(1) lookup |
| `BPF_MAP_TYPE_ARRAY` | Fixed-size array indexed by integer |
| `BPF_MAP_TYPE_LRU_HASH` | Least-recently-used hash — good for connection tracking |
| `BPF_MAP_TYPE_PROG_ARRAY` | Tail-call table — lets eBPF programs call other programs |
| `BPF_MAP_TYPE_PERF_EVENT_ARRAY` | Ring buffer for streaming events to userspace |

**Cilium's Service Load Balancing with eBPF Maps:**

```
Service VIP:Port → [BPF_MAP_TYPE_HASH] → Backend Pod IPs
```

When a pod sends traffic to a Service ClusterIP:
1. An eBPF program on the **socket layer** (before packet leaves the pod) intercepts the connect/sendmsg syscall.
2. It looks up the destination VIP in the **cilium_lb4_services** map.
3. It selects a backend pod from the **cilium_lb4_backends** map using consistent hashing.
4. It **rewrites the destination IP in-place** (socket-level DNAT) before the packet ever hits the network stack.

**vs. iptables kube-proxy:**
- iptables: Linear chain scan — O(n) for every packet, n = number of Service rules.
- Cilium eBPF: Hash map lookup — O(1) regardless of Service count.
- At 10,000 Services: iptables takes **~100ms per rule update**, Cilium takes **~1ms**.

---

### Question 75: How does Cilium enforce L7 (HTTP/gRPC) NetworkPolicy?

**Answer:**
Standard Kubernetes NetworkPolicy only works at Layer 3/4 (IP addresses, ports). Cilium extends this to **Layer 7** — HTTP paths, methods, gRPC service names — without any application code changes.

**How it works:**
1. The Cilium agent programs an eBPF TC (Traffic Control) hook on the pod's `eth0` interface.
2. When a packet matches a policy requiring L7 inspection, the eBPF program **redirects** it to a per-pod **Envoy proxy** (not Istio — Cilium manages its own Envoy instance).
3. Envoy applies the L7 policy and routes or drops the request.
4. For allowed traffic, Envoy forwards it to the destination pod.

**L7 policy example:**
```yaml
apiVersion: cilium.io/v2
kind: CiliumNetworkPolicy
metadata:
  name: allow-payment-api-only
spec:
  endpointSelector:
    matchLabels:
      app: payment-api
  ingress:
    - fromEndpoints:
        - matchLabels:
            app: checkout-service
      toPorts:
        - ports:
            - port: "8080"
              protocol: TCP
          rules:
            http:
              - method: "POST"
                path: "/api/v1/payment"    # Only allow POST to this exact path
              - method: "GET"
                path: "/api/v1/payment/status"
```

> **Key insight for interviews:** This allows you to enforce "checkout-service can POST to /payment but NOT GET /admin" — a security guarantee impossible with vanilla NetworkPolicy.

**Hubble — eBPF-based observability:**
```bash
# Real-time L7 flow visibility (no sidecar proxies needed)
hubble observe --namespace payments --protocol http
# Timestamp  Source              Destination          Status  Method  URL
# 09:12:01   checkout/pod-abc    payment-api/pod-xyz  200     POST    /api/v1/payment
# 09:12:02   fraud-svc/pod-def   payment-api/pod-xyz  403     GET     /api/v1/admin  ← DROPPED by policy
```

---

### Question 76: What is XDP and how is it used for DDoS mitigation in Kubernetes?

**Answer:**
**XDP (eXpress Data Path)** is the earliest hook point in the Linux network stack — it runs an eBPF program **directly in the NIC driver** before a packet is even allocated a socket buffer (`sk_buff`). This means it's processed before any kernel networking (no routing, no iptables, no conntrack).

**Performance:** XDP can process **14–24 million packets/second per core** (line rate on 10Gbps NICs), vs. iptables which saturates at ~1M pps.

**XDP actions:**
```c
XDP_DROP      // Drop the packet immediately — cheapest possible operation
XDP_PASS      // Let it through to the normal kernel stack
XDP_TX        // Bounce it back out the same interface
XDP_REDIRECT  // Send it to another interface or CPU
```

**DDoS mitigation in K8s (Cilium + XDP):**
```yaml
# Cilium enables XDP-accelerated load balancing automatically on supported NICs
apiVersion: v1
kind: ConfigMap
metadata:
  name: cilium-config
data:
  enable-bpf-masquerade: "true"
  loadbalancer-mode: "dsr"       # Direct Server Return — XDP-level LB
  bpf-lb-algorithm: "maglev"     # Consistent hashing across backends
```

When a SYN flood hits the cluster's ingress node:
1. The XDP eBPF program inspects each packet in the NIC driver.
2. Known-bad source IPs (populated in a BPF_MAP_TYPE_LRU_HASH by Cilium) cause `XDP_DROP` **before** the packet enters the kernel — consuming near-zero CPU.
3. Legitimate traffic gets `XDP_PASS` or redirected via `XDP_TX` to backend pods directly.

---

### Question 77: How does kube-proxy replacement with Cilium affect cluster architecture?

**Answer:**
When Cilium runs in **kube-proxy replacement mode** (`kubeProxyReplacement: strict`), it takes over **all** Service networking responsibilities:

**What changes:**
```
Before (with kube-proxy):
  iptables PREROUTING → DNAT → Pod

After (Cilium kube-proxy replacement):
  Socket-level eBPF hook → Hash map lookup → Pod
  (Packets never leave the pod's socket as Service VIP)
```

**Implications:**

| Aspect | With kube-proxy | With Cilium (no kube-proxy) |
|---|---|---|
| Rule updates | iptables chain rebuild — O(n) | BPF map update — O(1), sub-millisecond |
| Session affinity | iptables `statistic` module | BPF `LRU_HASH` for source-IP affinity |
| NodePort | iptables + conntrack | XDP or TC — 2–5x faster |
| Health checking | kube-proxy polls endpoints | Cilium integrates with Endpoint health APIs |
| Service discovery scale | Degrades at >5,000 services | Flat — works the same at 100,000 services |

**Deployment:**
```yaml
# Helm values for Cilium with full kube-proxy replacement
kubeProxyReplacement: strict
k8sServiceHost: "10.0.0.1"     # API server IP
k8sServicePort: "6443"
hostServices:
  enabled: true
nodePort:
  enabled: true
  mode: dsr                    # Direct Server Return for NodePort — better perf
```

---

## 🔹 Writing CRDs and Kubernetes Operators (Questions 78-82)

### Question 78: What is a CRD and when should you write a Kubernetes Operator?

**Answer:**
A **Custom Resource Definition (CRD)** extends the Kubernetes API with your own resource types. Once defined, you can `kubectl get`, `kubectl apply`, and `kubectl describe` your custom resources just like built-in Pod or Deployment objects.

**When to use a CRD (without controller):**
- As a structured config store — e.g., store feature flags, tenant configs, environment-specific settings in etcd via CRDs.
- When other controllers (ArgoCD, Crossplane) will reconcile the CRD for you.

**When to write a full Operator:**
- You need **operational runbook logic** in code. Examples:
  - A `PostgresCluster` CRD that provisions a Primary, creates Replicas, runs pg_basebackup for backups, handles failover.
  - A `RedisCluster` CRD that manages sharding, rebalancing, and rolling upgrades.
- The rule of thumb: **if a human operator would need to follow a runbook to manage the lifecycle, write a Kubernetes Operator to automate that runbook.**

**Operator maturity levels (OperatorHub):**
1. Basic Install
2. Seamless Upgrades
3. Full Lifecycle (backup/restore)
4. Deep Insights (metrics, alerts, dashboards)
5. Auto Pilot (auto-tuning, anomaly detection)

---

### Question 79: Define a CRD for a `DatabaseBackup` resource.

**Answer:**
```yaml
# Step 1: Define the CRD schema
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: databasebackups.mycompany.io
spec:
  group: mycompany.io
  versions:
    - name: v1alpha1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              required: ["databaseRef", "schedule", "retentionDays"]
              properties:
                databaseRef:
                  type: string
                  description: "Name of the database to back up"
                schedule:
                  type: string
                  description: "Cron expression for backup schedule"
                retentionDays:
                  type: integer
                  minimum: 1
                  maximum: 365
                storageLocation:
                  type: string
                  description: "S3 bucket URI"
            status:
              type: object
              properties:
                lastBackupTime:
                  type: string
                  format: date-time
                lastBackupStatus:
                  type: string
                  enum: ["Pending", "Running", "Succeeded", "Failed"]
                lastBackupSizeBytes:
                  type: integer
      subresources:
        status: {}    # Enables /status subresource for controller to update
      additionalPrinterColumns:
        - name: Database
          type: string
          jsonPath: .spec.databaseRef
        - name: Schedule
          type: string
          jsonPath: .spec.schedule
        - name: Last Status
          type: string
          jsonPath: .status.lastBackupStatus
        - name: Last Backup
          type: string
          jsonPath: .status.lastBackupTime
  scope: Namespaced
  names:
    plural: databasebackups
    singular: databasebackup
    kind: DatabaseBackup
    shortNames: ["dbb"]
```

```yaml
# Step 2: Create a CR instance
apiVersion: mycompany.io/v1alpha1
kind: DatabaseBackup
metadata:
  name: payments-db-backup
  namespace: production
spec:
  databaseRef: payments-postgres
  schedule: "0 2 * * *"        # 2 AM daily
  retentionDays: 30
  storageLocation: "s3://my-backups/payments/"
```

```bash
# Now works like built-in resources!
kubectl get dbb -n production
# NAME                   DATABASE           SCHEDULE      LAST STATUS   LAST BACKUP
# payments-db-backup     payments-postgres  0 2 * * *     Succeeded     2026-03-05T02:00:00Z
```

---

### Question 80: Write a Kubernetes Controller (Operator) in Go using controller-runtime.

**Answer:**
This is the minimal structure of a production-ready operator using `sigs.k8s.io/controller-runtime`:

```go
package controllers

import (
    "context"
    "fmt"
    "time"

    batchv1 "k8s.io/api/batch/v1"
    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/log"

    mycompanyv1 "mycompany.io/api/v1alpha1"
)

// DatabaseBackupReconciler reconciles DatabaseBackup objects
type DatabaseBackupReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

// Reconcile is the core control loop — called whenever a DatabaseBackup changes
// or is re-queued. MUST be idempotent.
func (r *DatabaseBackupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := log.FromContext(ctx)

    // 1. Fetch the DatabaseBackup CR
    backup := &mycompanyv1.DatabaseBackup{}
    if err := r.Get(ctx, req.NamespacedName, backup); err != nil {
        if errors.IsNotFound(err) {
            // CR was deleted — nothing to do (cleanup via ownerReferences)
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

    log.Info("Reconciling DatabaseBackup", "name", backup.Name, "schedule", backup.Spec.Schedule)

    // 2. Check if a backup Job is already running for this CR
    existingJob := &batchv1.Job{}
    jobName := fmt.Sprintf("backup-%s-%d", backup.Name, time.Now().Unix())
    err := r.Get(ctx, client.ObjectKey{Namespace: backup.Namespace, Name: jobName}, existingJob)

    if errors.IsNotFound(err) {
        // 3. Create a new backup Job
        job := r.buildBackupJob(backup, jobName)

        // Set owner reference — Job will be garbage-collected when DatabaseBackup is deleted
        if err := ctrl.SetControllerReference(backup, job, r.Scheme); err != nil {
            return ctrl.Result{}, err
        }

        if err := r.Create(ctx, job); err != nil {
            return ctrl.Result{}, err
        }

        log.Info("Created backup Job", "job", jobName)

        // 4. Update status
        backup.Status.LastBackupStatus = "Running"
        backup.Status.LastBackupTime = metav1.Now().UTC().Format(time.RFC3339)
        if err := r.Status().Update(ctx, backup); err != nil {
            return ctrl.Result{}, err
        }
    } else if err != nil {
        return ctrl.Result{}, err
    } else {
        // 5. Job exists — check its status
        if existingJob.Status.Succeeded > 0 {
            backup.Status.LastBackupStatus = "Succeeded"
        } else if existingJob.Status.Failed > 0 {
            backup.Status.LastBackupStatus = "Failed"
            log.Error(nil, "Backup Job failed", "job", jobName)
        }
        if err := r.Status().Update(ctx, backup); err != nil {
            return ctrl.Result{}, err
        }
    }

    // Re-queue after 1 hour (or parse the cron schedule)
    return ctrl.Result{RequeueAfter: time.Hour}, nil
}

func (r *DatabaseBackupReconciler) buildBackupJob(backup *mycompanyv1.DatabaseBackup, jobName string) *batchv1.Job {
    return &batchv1.Job{
        ObjectMeta: metav1.ObjectMeta{
            Name:      jobName,
            Namespace: backup.Namespace,
        },
        Spec: batchv1.JobSpec{
            Template: corev1.PodTemplateSpec{
                Spec: corev1.PodSpec{
                    RestartPolicy: corev1.RestartPolicyOnFailure,
                    Containers: []corev1.Container{
                        {
                            Name:  "backup",
                            Image: "mycompany/pg-backup:latest",
                            Env: []corev1.EnvVar{
                                {Name: "DB_REF", Value: backup.Spec.DatabaseRef},
                                {Name: "STORAGE_LOCATION", Value: backup.Spec.StorageLocation},
                                {Name: "RETENTION_DAYS", Value: fmt.Sprintf("%d", backup.Spec.RetentionDays)},
                            },
                        },
                    },
                },
            },
        },
    }
}

// SetupWithManager registers the controller with the manager
func (r *DatabaseBackupReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&mycompanyv1.DatabaseBackup{}).  // Watch DatabaseBackup CRs
        Owns(&batchv1.Job{}).               // Also watch Jobs owned by DatabaseBackup
        Complete(r)
}
```

**Key interview points on operator patterns:**
- **Idempotency is mandatory** — `Reconcile` may be called many times for the same event. Never assume it runs only once.
- **ownerReferences** — set via `ctrl.SetControllerReference` to enable cascading GC when the parent CR is deleted.
- **`Status()` subresource** — always update status via `.Status().Update()` not `.Update()` to avoid race conditions.
- **RequeueAfter** — use instead of timers; the controller-runtime will re-enqueue the object after the duration.

---

### Question 81: What is the difference between a Validating and Mutating Admission Webhook, and how do you write one?

**Answer:**

**Mutating Webhook:** Runs first. Can **modify** the incoming resource (inject defaults, add sidecars, set labels).
**Validating Webhook:** Runs after mutations. Can only **accept or reject** — no modifications.

**Lifecycle:**
```
kubectl apply → API Server → [MutatingWebhookConfiguration] → [ValidatingWebhookConfiguration] → etcd
```

**Write a Mutating Webhook in Go (using controller-runtime):**
```go
// Webhook that injects a cost-center label on all Pods
type PodInjector struct{}

func (p *PodInjector) Default(ctx context.Context, obj runtime.Object) error {
    pod, ok := obj.(*corev1.Pod)
    if !ok {
        return fmt.Errorf("expected a Pod, got %T", obj)
    }

    // Inject default label if missing
    if pod.Labels == nil {
        pod.Labels = make(map[string]string)
    }
    if _, ok := pod.Labels["cost-center"]; !ok {
        pod.Labels["cost-center"] = "unassigned"  // Default value
    }

    // Inject resource limits if unset (policy enforcement via mutation)
    for i := range pod.Spec.Containers {
        if pod.Spec.Containers[i].Resources.Limits == nil {
            pod.Spec.Containers[i].Resources.Limits = corev1.ResourceList{
                corev1.ResourceCPU:    resource.MustParse("500m"),
                corev1.ResourceMemory: resource.MustParse("512Mi"),
            }
        }
    }
    return nil
}
```

**WebhookConfiguration (registered with cluster):**
```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  name: pod-policy-injector
webhooks:
  - name: pod-injector.mycompany.io
    admissionReviewVersions: ["v1"]
    sideEffects: None
    clientConfig:
      service:
        name: admission-webhook-svc
        namespace: webhook-system
        path: /mutate-v1-pod
      caBundle: <base64-encoded-CA>    # Must match the webhook server's TLS cert
    rules:
      - operations: ["CREATE"]
        apiGroups: [""]
        apiVersions: ["v1"]
        resources: ["pods"]
    failurePolicy: Fail       # Reject pod if webhook is unreachable (strict mode)
    namespaceSelector:
      matchExpressions:
        - key: webhook-injection
          operator: In
          values: ["enabled"]  # Only apply to namespaces with this label
```

> **Prod tip:** Always set `failurePolicy: Fail` for security webhooks (reject if webhook is down), and `failurePolicy: Ignore` for non-critical mutation (accept pod even if webhook is unreachable).

---

### Question 82: What is the Informer pattern in Kubernetes and why is it better than polling?

**Answer:**
An **Informer** is the standard Kubernetes pattern for building controllers that watch API resources efficiently.

**Problem with naive polling:**
```go
// BAD: Polling every 10 seconds — hammers the API server under load
for {
    pods, _ := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
    processPods(pods)
    time.Sleep(10 * time.Second)
}
```

**Informer pattern (list-watch):**
```go
// GOOD: Use a SharedInformerFactory — one connection to the API server, shared across all controllers

factory := informers.NewSharedInformerFactory(clientset, time.Minute*10)
podInformer := factory.Core().V1().Pods()

// Register event handlers
podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
    AddFunc: func(obj interface{}) {
        pod := obj.(*corev1.Pod)
        fmt.Printf("Pod added: %s/%s\n", pod.Namespace, pod.Name)
    },
    UpdateFunc: func(oldObj, newObj interface{}) {
        oldPod := oldObj.(*corev1.Pod)
        newPod := newObj.(*corev1.Pod)
        if oldPod.ResourceVersion != newPod.ResourceVersion {
            fmt.Printf("Pod updated: %s/%s\n", newPod.Namespace, newPod.Name)
        }
    },
    DeleteFunc: func(obj interface{}) {
        pod := obj.(*corev1.Pod)
        fmt.Printf("Pod deleted: %s/%s\n", pod.Namespace, pod.Name)
    },
})

// Start informers (does the initial List, then opens a Watch stream)
factory.Start(stopCh)
factory.WaitForCacheSync(stopCh)  // Block until the local cache is warm
```

**How it works internally:**
1. **List:** On startup, the informer does a full `GET /api/v1/pods` and populates a **local in-memory cache** (thread-safe store).
2. **Watch:** It opens a persistent HTTP streaming connection (`GET /api/v1/pods?watch=true`). The API server sends events (ADDED, MODIFIED, DELETED) as they happen.
3. **Reconnect:** If the watch stream breaks, the informer uses the last seen `resourceVersion` to resume without re-listing.
4. **SharedInformer:** Multiple controllers sharing the same informer share one API server connection and one cache — critical for scalability.

> **Key interview point:** This is why Kubernetes controllers are event-driven and scalable: they don't poll; they react to a push-based event stream, with a local cache as the source of truth.

---

## 🔹 RuntimeClass, gVisor & Pod Sandboxing (Questions 83-85)

### Question 83: What is RuntimeClass and why does it exist?

**Answer:**
By default, all Kubernetes pods use the **same container runtime** (containerd + runc). **RuntimeClass** allows you to select a **different runtime** for specific pods — enabling different security-performance trade-offs.

**Why it exists:**
Standard `runc` containers share the host kernel directly. A container escape (CVE) gives the attacker access to the host. Some workloads need **stronger isolation**:
- Multi-tenant SaaS (untrusted customer code)
- Running user-supplied container images
- Compliance requirements (PCI-DSS, HIPAA) mandating isolation

**Define a RuntimeClass:**
```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: gvisor           # Name referenced by pods
handler: runsc           # The OCI runtime binary on the node (gVisor's runsc)
scheduling:
  nodeSelector:
    sandbox: gvisor      # Only schedule on nodes where gVisor is installed
  tolerations:
    - key: sandbox
      operator: Equal
      value: gvisor
      effect: NoSchedule
overhead:
  podFixed:
    memory: "120Mi"      # gVisor sandbox overhead (used by scheduler for placement)
    cpu: "250m"
```

**Use the RuntimeClass in a Pod:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: untrusted-workload
spec:
  runtimeClassName: gvisor    # Use gVisor instead of runc
  containers:
    - name: sandbox-app
      image: customer-code:latest
```

---

### Question 84: What is gVisor and how does it provide stronger isolation than runc?

**Answer:**
**gVisor** (Google's container sandbox, used in GKE Sandbox) interposes a **user-space kernel** between the container and the host kernel.

**Standard runc:**
```
Container App → Linux syscall → Host Kernel → Hardware
```

**gVisor (runsc):**
```
Container App → Linux syscall → gVisor Sentry (user-space kernel) → Host Kernel (minimal syscalls) → Hardware
```

**How it works:**
1. The container's processes make linux syscalls (open, read, connect, etc.).
2. gVisor's **Sentry** component intercepts ALL syscalls — the container never directly calls the host kernel.
3. Sentry implements the Linux syscall interface itself in Go. It translates safe operations to a restricted set of host syscalls.
4. Even if the container exploits a vulnerability, it can only compromise the Sentry — not the host kernel.

**Two kernel interception mechanisms:**
- **ptrace mode:** Sentry uses Linux `ptrace` to intercept syscalls. Simpler but slower.
- **KVM mode (production):** Sentry runs as a Type-2 hypervisor using KVM hardware virtualization. Faster, stronger isolation.

**Trade-offs:**

| | runc | gVisor (runsc) | Kata Containers |
|---|---|---|---|
| Isolation | Kernel namespaces only | User-space kernel (syscall interception) | Full VM (separate kernel) |
| Performance overhead | ~0% | ~10-20% CPU, ~120MB RAM | ~5-10% CPU, ~256MB RAM |
| Syscall compatibility | 100% | ~80% (some uncommon syscalls missing) | ~100% |
| Attack surface | Entire host kernel | Only Sentry process | VM boundary |
| Use case | Trusted workloads | Untrusted code, multi-tenant SaaS | Compliance-heavy, financial |

---

### Question 85: What is Kata Containers and how does it differ from gVisor?

**Answer:**
**Kata Containers** wraps each pod in a **lightweight Virtual Machine** with its own kernel — providing hardware-level isolation.

**Architecture:**
```
Container App → Container Kernel (per-pod, full Linux) → KVM hypervisor → Host Kernel
```

**vs. gVisor:**
- gVisor intercepts syscalls in **user-space** (no VM, lighter overhead, some syscall gaps).
- Kata uses a **real VM** (full syscall compatibility, slightly higher overhead, hardware-enforced isolation).

**When to prefer Kata over gVisor:**
1. You need **100% syscall compatibility** (applications that use obscure syscalls gVisor doesn't implement, e.g., eBPF inside the container, FUSE, certain kernel modules).
2. **Regulatory compliance** requiring hardware VM boundaries (some PCI-DSS audits require this).
3. Running **nested Kubernetes** (K8s inside a Kata pod) for CI platforms.

**RuntimeClass for Kata:**
```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: kata-containers
handler: kata
scheduling:
  nodeSelector:
    kata: "true"
overhead:
  podFixed:
    memory: "256Mi"
    cpu: "500m"
```

**GKE Sandbox** uses gVisor under the hood:
```yaml
# GKE node pool with sandbox enabled
spec:
  nodeConfig:
    sandboxConfig:
      sandboxType: gvisor
```

**Production pattern — defense in depth:**
```
Layer 1: OPA Gatekeeper / Kyverno (admission policy — prevent misconfig)
Layer 2: gVisor / Kata (runtime sandbox — contain exploits)
Layer 3: Falco (runtime detection — alert when something suspicious happens)
```
These three layers together address different phases of an attack.

---

## 🔹 Multi-Cluster Federation & Submariner (Questions 86-87)

### Question 86: What is Submariner and how does it enable cross-cluster service connectivity?

**Answer:**
**Submariner** (CNCF Sandbox) enables direct **pod-to-pod and service-to-service connectivity across different Kubernetes clusters** — without a full service mesh like Istio.

**Problem it solves:**
By default, pod CIDR ranges in different clusters overlap (both use 10.0.0.0/8), and cluster networks are completely isolated. A pod in Cluster A cannot reach a pod in Cluster B by IP.

**How Submariner works:**

```
Cluster A                    Cluster B
┌─────────────────┐         ┌─────────────────┐
│  Pod (10.1.1.5) │         │  Pod (10.2.1.8) │
│       ↓         │         │       ↑         │
│  Gateway Node   │ ←IPSec→ │  Gateway Node   │
│  (Submariner)   │ tunnel  │  (Submariner)   │
└─────────────────┘         └─────────────────┘
         ↑                           ↑
    Broker (Hub cluster — coordinates route exchange)
```

1. Each cluster runs a **Submariner Gateway** pod on a designated node.
2. Gateways establish **IPSec or WireGuard tunnels** between clusters.
3. A centralized **Broker** cluster stores the route table that maps pod CIDRs and Service CIDRs to their cluster.
4. The **Route Agent** DaemonSet on each worker node programs the node's routing table to send inter-cluster traffic to the local gateway.

**ServiceExport / ServiceImport (Lighthouse CRDs):**
```yaml
# In Cluster A — export the payments service for other clusters to discover
apiVersion: multicluster.x-k8s.io/v1alpha1
kind: ServiceExport
metadata:
  name: payment-api
  namespace: production
```

```yaml
# In Cluster B — import is auto-created by Lighthouse; access via DNS
# payment-api.production.svc.clusterset.local
# → resolves to Cluster A's payment-api pods via the IPSec tunnel
```

**vs. Istio multi-cluster:**
| | Submariner | Istio Multi-cluster |
|---|---|---|
| Routing | IP-level (direct pod routing) | Proxy-level (Envoy sidecar) |
| Overhead | Very low (no sidecar) | Higher (Envoy per pod) |
| L7 policy | Not supported natively | Full Istio policy/observability |
| Use case | Cluster connectivity without service mesh | Full service mesh spanning clusters |

---

### Question 87: Describe a Cluster API (CAPI) workflow for declaratively managing cluster lifecycle.

**Answer:**
**Cluster API (CAPI)** is a CNCF project that extends Kubernetes to manage the lifecycle of other Kubernetes clusters using CRDs and controllers — applying the same declarative approach used for pods/deployments to entire clusters.

**Core philosophy:** A running Kubernetes cluster (the **Management Cluster**) reconciles `Cluster` CRDs that represent **Workload Clusters**.

**Key CRDs:**

| CRD | Purpose |
|---|---|
| `Cluster` | Top-level cluster definition |
| `MachineDeployment` | Like a Deployment, but for Cluster Nodes |
| `Machine` | Represents a single node VM |
| `AWSMachineTemplate` | Infrastructure-specific node config (AWS) |
| `KubeadmControlPlane` | Manages the control plane using kubeadm |

**Complete EKS cluster definition via CAPI:**
```yaml
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: prod-us-east-1
  namespace: clusters
spec:
  clusterNetwork:
    pods:
      cidrBlocks: ["10.128.0.0/16"]
    services:
      cidrBlocks: ["172.20.0.0/16"]
  controlPlaneRef:
    apiVersion: controlplane.cluster.x-k8s.io/v1beta1
    kind: AWSManagedControlPlane       # EKS control plane
    name: prod-us-east-1-control-plane
  infrastructureRef:
    apiVersion: infrastructure.cluster.x-k8s.io/v1beta1
    kind: AWSManagedCluster
    name: prod-us-east-1
---
apiVersion: cluster.x-k8s.io/v1beta1
kind: MachineDeployment
metadata:
  name: prod-us-east-1-workers
spec:
  clusterName: prod-us-east-1
  replicas: 5
  template:
    spec:
      bootstrap:
        configRef:
          kind: EKSConfigTemplate
          name: prod-worker-config
      infrastructureRef:
        kind: AWSMachineTemplate
        name: prod-worker-template
        # EC2 instance type, AMI, SGs, etc. defined here
```

**Cluster lifecycle operations via CAPI:**
```bash
# Scale worker nodes
kubectl patch machinedeployment prod-us-east-1-workers \
  -p '{"spec":{"replicas":8}}'

# Upgrade Kubernetes version (rolling, zero-downtime)
kubectl patch awsmanagedcontrolplane prod-us-east-1-control-plane \
  -p '{"spec":{"version":"v1.29"}}'

# Delete entire cluster
kubectl delete cluster prod-us-east-1
# → CAPI controller tears down all EC2 instances, VPCs, EKS cluster — clean

# ClusterClass — template for many clusters (like a class instantiated many times)
```

> **Key interview insight:** CAPI turns "create an EKS cluster" from a 3-hour Terraform/CLI workflow into a `kubectl apply` — and cluster upgrades from a high-risk manual procedure into a rolling, reconciled operation. At FAANG scale (managing hundreds of clusters), this GitOps-driven approach is how platform teams operate.
