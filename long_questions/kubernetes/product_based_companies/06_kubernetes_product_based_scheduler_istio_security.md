# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 6)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Stripe, Datadog, etc.
> **Focus deep-dive:** Custom Scheduler Framework, Istio Service Mesh (Traffic Shaping), Kubernetes Audit Logging, Falco Runtime Security, FinOps & Cost Optimization, and Multi-Cluster Strategies.

---

## 🔹 Custom Scheduler & Scheduling Framework (Questions 59-62)

### Question 59: What is the Kubernetes Scheduling Framework and how does it differ from a Scheduler Extender?

**Answer:**
The **Scheduling Framework** is the modern (v1.19+) plugin architecture for extending the Kubernetes Scheduler natively.

**Scheduler Extender (old approach):**
- The default scheduler calls an **external HTTP webhook** at specific scheduling points.
- **Cons:** Network round-trips on every scheduling decision → massive latency at scale. Hard to deploy, maintain.

**Scheduling Framework (new approach):**
- You write a **Go plugin** compiled directly into a custom scheduler binary.
- Zero network overhead — runs in-process with the scheduler.
- Plugins hook into well-defined extension points:

```
PreFilter → Filter → PostFilter → PreScore → Score → Reserve → Permit → PreBind → Bind → PostBind
```

| Extension Point | Purpose |
|---|---|
| `Filter` | Eliminate ineligible nodes (like predicates) |
| `Score` | Rank remaining nodes (like priorities) |
| `Reserve` | Reserve node resources before binding |
| `Permit` | Delay or reject binding (e.g., gang scheduling) |
| `PreBind` / `Bind` | Perform the actual binding to a node |

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Kubernetes Scheduling Framework and how does it differ from a Scheduler Extender?
**Your Response:** "The Scheduling Framework is the modern plugin architecture for extending the scheduler, while Scheduler Extenders were the old approach. Extenders use external HTTP webhooks which cause network latency on every scheduling decision and are hard to maintain. The Framework uses Go plugins compiled directly into the scheduler binary, eliminating network overhead. Plugins hook into extension points like Filter to eliminate nodes, Score to rank them, and Permit for gang scheduling. It's like the difference between calling an expert consultant for every decision versus having an in-house specialist who's always available - much faster and more reliable."

---

### Question 60: Write a custom scheduler plugin that prefers nodes with an SSD label.

**Answer:**
```go
package ssdpreference

import (
    "context"
    v1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "sigs.k8s.io/scheduler-plugins/pkg/apis/config"
    framework "k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "SSDPreference"

type SSDPreference struct{}

func (s *SSDPreference) Name() string { return Name }

// Score: Give nodes labeled disk=ssd a higher score (100), others get 0
func (s *SSDPreference) Score(ctx context.Context, state *framework.CycleState,
    p *v1.Pod, nodeName string) (int64, *framework.Status) {

    nodeInfo, err := state.Read(framework.NodeInfoSnapshotKey)
    // In practice, retrieve from the informer cache
    node, _ := framework.NodeInfoFromCycleState(nodeName, state)

    if diskType, ok := node.Node().Labels["disk"]; ok && diskType == "ssd" {
        return 100, nil  // Prefer SSD nodes
    }
    return 0, nil        // Non-SSD nodes get score 0
}

func (s *SSDPreference) ScoreExtensions() framework.ScoreExtensions { return nil }
```

**Registration in scheduler config:**
```yaml
apiVersion: kubescheduler.config.k8s.io/v1
kind: KubeSchedulerConfiguration
profiles:
  - schedulerName: custom-scheduler
    plugins:
      score:
        enabled:
          - name: SSDPreference
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** Write a custom scheduler plugin that prefers nodes with an SSD label.
**Your Response:** "I'd write a Go plugin that implements the Score interface. The plugin checks if a node has the label 'disk=ssd' and gives it a score of 100, otherwise 0. This makes the scheduler prefer SSD nodes for workloads that need fast storage. I register the plugin in the scheduler configuration under the score extension point. The plugin runs in-process with the scheduler, so there's no network overhead. It's like giving the scheduler a preference checklist - when it sees an SSD-labeled node, it marks it as highly preferred for workloads that need fast I/O."

---

### Question 61: What is Gang Scheduling and why is it needed for ML workloads?

**Answer:**
**Gang Scheduling** means: either ALL pods of a job are scheduled simultaneously, or NONE are.

**Problem without it:**
A distributed TensorFlow training job requires 8 GPUs across 8 pods. If only 7 pods get scheduled (because node 8 ran out of GPU), those 7 pods sit idle, consuming GPU resources, while waiting forever for pod 8. This is called **resource deadlock** or **head-of-line blocking**.

**Solution — `Permit` extension point:**
The `Coscheduling` plugin (from scheduler-plugins repo) implements gang scheduling using the `Permit` phase:
1. Pods in the same `PodGroup` are held in the `Permit` phase (not yet bound).
2. Only when `minMember` pods have all passed Filter + Score simultaneously does the plugin release the whole group to `Bind`.
3. If not enough nodes are available, the whole group is queued (no partial waste).

```yaml
apiVersion: scheduling.sigs.k8s.io/v1alpha1
kind: PodGroup
metadata:
  name: tf-training-job
spec:
  minMember: 8        # All 8 must be schedulable before any binds
  scheduleTimeoutSeconds: 300
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Gang Scheduling and why is it needed for ML workloads?
**Your Response:** "Gang scheduling means either all pods of a job are scheduled together, or none are. This is critical for ML workloads like distributed TensorFlow training that need 8 GPUs across 8 pods. Without gang scheduling, if only 7 pods get scheduled, they sit idle consuming GPU resources while waiting for the 8th pod - this is called resource deadlock. The Coscheduling plugin uses the Permit extension point to hold pods in a PodGroup until all minMember pods can be scheduled simultaneously. If not enough nodes are available, the whole group waits. It's like booking a restaurant table - either everyone gets a seat together, or nobody goes, rather than having some people wait awkwardly while others are already eating."

---

### Question 62: How does the Kubernetes Scheduler handle preemption when a high-priority pod is pending?

**Answer:**
**Preemption** is triggered when a high-priority pod cannot be scheduled due to insufficient resources.

**Flow:**
1. Scheduler runs Filter/Score — no node has enough resources.
2. Since the pod has a higher **`PriorityClass`**, the scheduler enters the **PostFilter** phase.
3. It finds nodes where evicting one or more *lower-priority* pods would free enough resources.
4. The scheduler **nominates** that node on the pending pod's status (`nominatedNodeName`).
5. The preempted pods receive a graceful termination (respecting their `terminationGracePeriodSeconds`).
6. Once resources are freed, the high-priority pod binds to the target node.

**PriorityClass example:**
```yaml
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: high-priority-prod
value: 1000000         # Higher = more important
globalDefault: false
preemptionPolicy: PreemptLowerPriority   # Can evict lower-priority pods
description: "For production critical services"
---
# Low-priority workload that CAN be evicted
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: low-priority-batch
value: 1000
preemptionPolicy: Never     # This class itself cannot preempt others
```

> **Key interview point:** PDB (`minAvailable`) is respected during preemption — if evicting a low-priority pod would violate its PDB, the scheduler skips it and looks at other candidates.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the Kubernetes Scheduler handle preemption when a high-priority pod is pending?
**Your Response:** "When a high-priority pod can't be scheduled due to insufficient resources, the scheduler triggers preemption. It first tries normal scheduling, then enters the PostFilter phase where it looks for nodes where evicting lower-priority pods would free enough resources. The scheduler nominates a target node and gracefully terminates the lower-priority pods, respecting their termination grace periods. Once resources are freed, the high-priority pod binds. I use PriorityClasses to define the hierarchy - high-priority production pods can preempt batch jobs. Importantly, PDBs are respected during preemption, so the scheduler won't evict pods that would violate their availability guarantees. It's like having emergency vehicles that can clear traffic, but only if it doesn't break critical road safety rules."

---

## 🔹 Istio Service Mesh — Traffic Shaping & Resilience (Questions 63-67)

### Question 63: Explain Istio's VirtualService and DestinationRule. How do they work together?

**Answer:**
These two CRDs are the foundation of all Istio traffic control.

**DestinationRule** — defines how traffic is routed **to** a service (subsets, load balancing, circuit breaking):
```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: payment-api
spec:
  host: payment-api     # K8s Service name
  trafficPolicy:
    loadBalancer:
      simple: ROUND_ROBIN
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
    - name: canary
      labels:
        version: canary
      trafficPolicy:
        connectionPool:
          http:
            http1MaxPendingRequests: 100
```

**VirtualService** — defines traffic routing **rules** (which subset gets what % of traffic):
```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: payment-api
spec:
  hosts:
    - payment-api
  http:
    - match:
        - headers:
            x-user-type:
              exact: "beta-tester"    # Beta users always get v2
      route:
        - destination:
            host: payment-api
            subset: v2
    - route:                          # Everyone else: 95%/5% split
        - destination:
            host: payment-api
            subset: v1
          weight: 95
        - destination:
            host: payment-api
            subset: canary
          weight: 5
```

**Together:** `DestinationRule` defines *what* subsets exist and their policies. `VirtualService` defines *how much traffic* goes to each subset.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Istio's VirtualService and DestinationRule. How do they work together?
**Your Response:** "DestinationRule and VirtualService work together to control traffic. DestinationRule defines how traffic routes to a service - it creates subsets based on labels, sets load balancing policies, and configures connection pooling. VirtualService defines the actual routing rules - what percentage of traffic goes to each subset, or which users get routed to specific versions. Think of DestinationRule as defining the available highways and their rules, while VirtualService is the traffic control system that directs cars onto specific highways. Together they enable sophisticated traffic management like canary deployments and A/B testing without changing application code."

---

### Question 64: How does Istio implement a Circuit Breaker pattern?

**Answer:**
Traditional circuit breakers (Hystrix) are implemented in application code. Istio implements circuit breaking **transparently at the proxy (Envoy) layer** — no code changes needed.

**Two main mechanisms in DestinationRule:**

**1. Connection Pool limits (protect the downstream service):**
```yaml
trafficPolicy:
  connectionPool:
    tcp:
      maxConnections: 100       # Max concurrent TCP connections
    http:
      http1MaxPendingRequests: 50   # Max queued requests
      maxRequestsPerConnection: 10   # Force connection cycling
```

**2. Outlier Detection (automatically eject unhealthy endpoints):**
```yaml
trafficPolicy:
  outlierDetection:
    consecutive5xxErrors: 5      # After 5 consecutive 5xx errors...
    interval: 30s                # ...evaluated every 30 seconds
    baseEjectionTime: 30s        # ...eject the pod for 30s
    maxEjectionPercent: 50       # Never eject more than 50% of the pool
```

**How it works:** If a specific pod instance (IP) returns 5 consecutive 5xx errors within the evaluation window, Envoy automatically stops routing traffic to it for `baseEjectionTime`. After the ejection period, the pod is gradually allowed back. This prevents cascading failures without any application-level code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Istio implement a Circuit Breaker pattern?
**Your Response:** "Istio implements circuit breaking transparently at the Envoy proxy layer, so no application code changes are needed. It uses two mechanisms in DestinationRule: connection pool limits to protect downstream services by limiting concurrent connections and queued requests, and outlier detection that automatically ejects unhealthy endpoints. If a pod returns 5 consecutive 5xx errors, Envoy stops routing traffic to it for 30 seconds, then gradually allows it back. This prevents cascading failures without writing any circuit breaker code in the application. It's like having an automatic traffic controller that redirects cars around accidents before they cause gridlock."

---

### Question 65: How do you implement fault injection to test resilience in a Kubernetes service mesh?

**Answer:**
Istio allows injecting **synthetic faults** to test how your services handle upstream failures — without touching any application code.

**Inject a 5-second delay for 20% of requests:**
```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: user-service
spec:
  hosts:
    - user-service
  http:
    - fault:
        delay:
          percentage:
            value: 20.0      # 20% of requests get delayed
          fixedDelay: 5s     # by 5 seconds
    - route:
        - destination:
            host: user-service
```

**Inject HTTP 503 errors for 10% of requests:**
```yaml
    - fault:
        abort:
          percentage:
            value: 10.0
          httpStatus: 503
```

**Why this matters at product companies:**
- Chaos engineering without a separate tool (Chaos Monkey).
- Validate that `timeout` and `retry` policies in dependent services actually work.
- Can be scoped to specific users via header matching (e.g., only inject faults for internal test accounts).

---

### Question 66: Describe Istio's retry and timeout policies at the mesh level.

**Answer:**
Without Istio, each microservice team must implement retries and timeouts in their application code (e.g., Resilience4j, go-retryablehttp). With Istio, these become **infrastructure-level policies**:

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: inventory-service
spec:
  hosts:
    - inventory-service
  http:
    - timeout: 3s               # Global timeout for this route
      retries:
        attempts: 3             # Retry up to 3 times
        perTryTimeout: 1s       # Each attempt must respond within 1s
        retryOn: "gateway-error,connect-failure,retriable-4xx"
      route:
        - destination:
            host: inventory-service
```

**`retryOn` conditions:**
- `gateway-error` — 502, 503, 504
- `connect-failure` — TCP connection refused
- `retriable-4xx` — 409 Conflict (safe to retry)
- `5xx` — any 5xx response

> **Key prod insight:** Retries + timeouts interact. With 3 retries and 1s per-try timeout, the *worst-case* latency is 3 seconds (not 1s). Your `timeout: 3s` must account for `attempts × perTryTimeout`.

---

### Question 67: What is Istio's Sidecar resource and why is it important at scale?

**Answer:**
By default, every Envoy sidecar in an Istio mesh receives the **full cluster's service registry** — configuration for every single service across all namespaces.

In a cluster with 500 services:
- Each Envoy holds routing rules, endpoints, and TLS certs for all 500 services.
- This wastes **~150MB RAM per pod** and slows xDS config propagation (istiod pushes updates to 10,000 proxies for every service change).

**`Sidecar` resource scopes the proxy's view:**
```yaml
apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: payment-sidecar
  namespace: payments
spec:
  egress:
    - hosts:
        - "./*"                  # All services in the payments namespace
        - "istio-system/*"       # Istio control plane
        - "shared-services/db-service.shared-services.svc.cluster.local"
```

**Result:** The payment pod's Envoy only knows about services in `payments`, `shared-services`, and `istio-system`. Instead of 500 service configs, it holds ~15. Memory drops from 150MB → 20MB per proxy, and config push times drop proportionally.

---

## 🔹 Audit Logging, Runtime Security & FinOps (Questions 68-72)

### Question 68: What is Kubernetes Audit Logging and what does a real audit policy look like?

**Answer:**
The **Kube-APIServer Audit Log** records every request made to the API server — who did what, when, and with what response. Essential for security compliance (SOC2, PCI-DSS).

**Audit Levels:**
- `None` — Don't log this request.
- `Metadata` — Log request metadata (user, verb, resource) but NOT the request body.
- `Request` — Metadata + request body.
- `RequestResponse` — Metadata + request + response body (most verbose).

**Real audit policy:**
```yaml
apiVersion: audit.k8s.io/v1
kind: Policy
rules:
  # Never log read-only system activity (noisy, low-value)
  - level: None
    users: ["system:kube-proxy"]
    verbs: ["watch"]
    resources:
      - group: ""
        resources: ["endpoints", "services", "secrets"]

  # Log secret access at Request level (who read which secret)
  - level: Request
    resources:
      - group: ""
        resources: ["secrets", "configmaps"]

  # Log all exec/port-forward at RequestResponse (security critical)
  - level: RequestResponse
    resources:
      - group: ""
        resources: ["pods/exec", "pods/portforward", "pods/attach"]

  # Default: log everything else at Metadata
  - level: Metadata
```

---

### Question 69: How does Falco detect runtime security threats in Kubernetes?

**Answer:**
**Falco** (CNCF project) is a runtime security tool that monitors **system calls** via eBPF probes on every node to detect anomalous container behavior.

**How it works:**
- Falco loads an eBPF program into the Linux kernel that intercepts every `syscall` (open, read, connect, execve, etc.) made by containers.
- These events are compared against a **rules engine** in user-space using a Lua-based rule language.
- Violations trigger alerts (stdout, Slack, Falcosidekick → PagerDuty, etc.).

**Example Falco rules:**
```yaml
- rule: Shell Spawned in Container
  desc: A shell was spawned inside a container — possible intrusion
  condition: >
    spawned_process and container and
    proc.name in (shell_binaries) and
    not proc.pname in (known_shell_parents)
  output: "Shell spawned in container (user=%user.name cmd=%proc.cmdline container=%container.name)"
  priority: WARNING

- rule: Sensitive File Read
  desc: A binary read /etc/shadow or /etc/passwd
  condition: >
    open_read and container and
    fd.name in (/etc/shadow, /etc/passwd, /etc/kubernetes/admin.conf)
  output: "Sensitive file read (file=%fd.name proc=%proc.name container=%container.name)"
  priority: CRITICAL
```

**Falco vs Admission Controllers:**
- **Admission Controllers** prevent bad configs from being created (shift-left).
- **Falco** catches runtime threats that bypass admission gates (e.g., an attacker who already has container access).

---

### Question 70: How do you optimize Kubernetes costs (FinOps) in a cloud environment?

**Answer:**
FinOps in Kubernetes focuses on eliminating three types of waste:

**1. Right-sizing (Over-provisioned Requests/Limits):**
- Use **VPA in recommendation mode** (not enforcement) to suggest correct request values based on historical usage.
- Use **Goldilocks** (Fairwinds): runs VPA in recommendation mode across all deployments and provides a UI dashboard showing current vs. recommended requests.

**2. Idle Resource Elimination:**
- Use **KEDA** to scale event-driven workloads down to **zero** replicas during off-peak hours.
- Use **Karpenter's consolidation**: automatically replaces underutilized nodes with smaller, cheaper instances or terminates empty nodes.

**3. Spot/Preemptible Instances:**
```yaml
# Karpenter NodePool — prefer spot, fallback to on-demand
spec:
  template:
    spec:
      requirements:
        - key: karpenter.sh/capacity-type
          operator: In
          values: ["spot", "on-demand"]
      # Tolerate spot interruptions
      taints:
        - key: spot-interruption
          effect: NoSchedule
```

**4. Namespace Cost Attribution:**
- Use **OpenCost** or **Kubecost**: allocates cluster cost per namespace/pod/label using cloud pricing APIs. Enables team-level cost showback/chargeback.

**Typical savings: 40-60% on compute** by combining right-sizing + spot + zero-scale.

---

### Question 71: Describe a multi-cluster Kubernetes strategy for a global SaaS product.

**Answer:**
**Architecture: Hub-and-Spoke with GitOps**

```
                    ┌──────────────────┐
                    │   Git Repository │
                    │  (Single Source) │
                    └────────┬─────────┘
                             │ ArgoCD sync
          ┌──────────────────┼──────────────────┐
          ▼                  ▼                  ▼
  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
  │  US-EAST-1   │  │  EU-WEST-1   │  │  AP-SE-1     │
  │  K8s Cluster │  │  K8s Cluster │  │  K8s Cluster │
  └──────────────┘  └──────────────┘  └──────────────┘
          ▲                  ▲                  ▲
          └──────────────────┴──────────────────┘
                    CloudFlare / Route53
                    (Geo-Routing / Anycast DNS)
```

**Key design decisions:**

| Concern | Solution |
|---|---|
| Config management | ArgoCD `ApplicationSet` generates one App per cluster from a single chart |
| Secrets | External Secrets Operator syncs from a central Vault per-region |
| Observability | Thanos / Grafana Mimir aggregates Prometheus metrics from all clusters |
| Progressive rollout | Deploy to AP first (smallest user base), monitor, then EU, then US |
| Data residency (GDPR) | EU cluster NEVER routes to US-hosted databases |

---

### Question 72: What is Kubernetes API deprecation management, and how do you handle upgrades safely?

**Answer:**
Kubernetes deprecates API versions with each minor release. Example: `extensions/v1beta1` Ingress was removed in v1.22, replaced by `networking.k8s.io/v1`.

**Safe upgrade process:**

**Step 1 — Scan for deprecated APIs before upgrading:**
```bash
# Using Pluto (Fairwinds tool)
pluto detect-helm --target-versions k8s=v1.26
# NAME           NAMESPACE   KIND       VERSION              DEPRECATED  REMOVED
# nginx-ingress  default     Ingress    extensions/v1beta1   true        true
```

**Step 2 — Test on a non-production cluster first:**
- Use **Cluster API (CAPI)** or cloud provider tools (EKS Blue/Green node groups) to spin up a parallel cluster running the target version.

**Step 3 — Update manifests and Helm charts:**
```bash
# Convert old API versions automatically
kubectl-convert -f old-ingress.yaml --output-version networking.k8s.io/v1
```

**Step 4 — Upgrade control plane first, then node groups:**
- Kubernetes guarantees n-2 version skew between control plane and kubelet.
- Upgrade control plane → drain + upgrade node groups one group at a time (rolling).

**Step 5 — Validate post-upgrade:**
```bash
kubectl get --raw /metrics | grep apiserver_requested_deprecated_apis
```
