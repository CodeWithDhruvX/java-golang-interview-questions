# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 5)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Stripe, Robinhood, etc.
> **Focus deep-dive:** Pod Disruption Budgets, Observability Stack (Prometheus/Grafana), eBPF tracing, Canary deployments, and cert-manager.

---

## 🔹 Reliability Engineering — PDBs & Disruption Control (Questions 46-50)

### Question 46: What is a Pod Disruption Budget (PDB) and why is it critical during node drains?

**Answer:**
A **PodDisruptionBudget (PDB)** is a Kubernetes API object that limits the number of Pods of a replicated application that may be voluntarily disrupted at the same time.

Without a PDB, running `kubectl drain node-1` during a rolling upgrade of nodes could evict ALL 3 replicas of a critical service simultaneously before any replace Pod becomes Ready, causing a full service outage.

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: api-pdb
spec:
  minAvailable: 2          # OR: maxUnavailable: 1
  selector:
    matchLabels:
      app: payment-api
```

- **`minAvailable: 2`:** At least 2 Pods must remain Running during any eviction operation.
- **`maxUnavailable: 1`:** At most 1 Pod can be unavailable at any point.
- The Cluster Autoscaler and `kubectl drain` both **respect** PDBs, blocking further evictions until the budget is satisfied.

> **Common gotcha:** PDBs only protect against *voluntary* disruptions (drains, rolling upgrades). They do NOT protect against involuntary disruptions (node hardware failures).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Pod Disruption Budget (PDB) and why is it critical during node drains?
**Your Response:** "A PodDisruptionBudget is like a safety net that ensures I always have enough pods running during maintenance. Without it, when I drain a node for upgrades, Kubernetes could evict all replicas of a service at once, causing an outage. With a PDB, I can specify that at least 2 pods must always be available, or that only 1 pod can be down at a time. Both kubectl drain and the Cluster Autoscaler respect these limits. It's crucial for high availability - like having a rule that at least 2 cashiers must always be working during a store remodel. The key thing to remember is PDBs only protect against planned disruptions like maintenance, not unexpected failures like hardware crashes."

---

### Question 47: If a node is being drained for maintenance and a PDB is blocking eviction, how do you safely proceed?

**Answer:**
This is a classic SRE scenario. You have several safe options:

1. **Wait for health:** The eviction is retried automatically. If the replacement pod will eventually become Ready, the drain will unblock by itself.
2. **Scale up temporarily:** Increase the deployment's `replicas` count *before* the drain, satisfying the `minAvailable` threshold during the disruption window.
3. **Evaluate (with caution):** `kubectl drain --disable-eviction` bypasses the Eviction API (and the PDB), using DELETE directly instead. This is a last resort, as it removes the safety guarantee.
4. **Fix the underlying pod issue:** If a pod is `Pending` or `CrashLoopBackOff`, it won't count towards `minAvailable`, so first fix the pod health to unblock the budget.

### How to Explain in Interview (Spoken style format)
**Interviewer:** If a node is being drained for maintenance and a PDB is blocking eviction, how do you safely proceed?
**Your Response:** "When a PDB blocks eviction during node drain, I have several options. First, I can wait - if the replacement pod will become healthy, the drain will unblock automatically. Second, I can temporarily scale up the deployment to satisfy the minAvailable requirement during maintenance. Third, as a last resort, I can use kubectl drain --disable-eviction to bypass the PDB, but this removes the safety guarantee so I'm very careful with this approach. Fourth, I check if unhealthy pods are causing the issue - pods in CrashLoopBackOff or Pending don't count toward availability, so fixing those might unblock the drain. It's like having a minimum staffing rule during store renovations - either wait for new staff to arrive, hire temporary staff, or as a last resort, override the rule but accept the risk."

---

### Question 48: How does Kubernetes leader election work for controllers like `kube-controller-manager` and `kube-scheduler`?

**Answer:**
In a highly available control plane (e.g., 3 master nodes), you do NOT want all 3 `kube-scheduler` instances making scheduling decisions simultaneously (race conditions).

Kubernetes uses **Lease-based leader election**:
1. Each controller instance tries to create (or renew) a **`Lease`** object in the `kube-system` namespace (e.g., `kube-scheduler`).
2. The instance that successfully writes to the Lease first becomes the **active leader**.
3. The leader continuously renews the Lease every few seconds (controlled by `leaseDuration`, `renewDeadline`, `retryPeriod` flags).
4. The other instances (standby) continuously watch the Lease. If the leader fails to renew within `renewDeadline`, a standby immediately acquires the Lease and promotes itself to leader.

This ensures only **one active scheduler** runs at any time, even across a flapping network partition.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Kubernetes leader election work for controllers like `kube-controller-manager` and `kube-scheduler`?
**Your Response:** "Kubernetes uses a lease-based leader election system to prevent multiple controllers from conflicting with each other. In a 3-master cluster, only one scheduler should be active at a time. Each controller tries to grab a Lease object in kube-system namespace - whoever gets it first becomes the leader. The leader continuously renews the lease like a heartbeat, and the other controllers watch it. If the leader fails to renew within the timeout, another controller immediately takes over. This ensures smooth failover without split-brain scenarios. It's like having a team where only one person holds the talking stick at a time - if they drop it, someone else immediately picks it up to continue the conversation."

---

### Question 49: Explain Karpenter and how it differs from the traditional Cluster Autoscaler.

**Answer:**

| Feature | Cluster Autoscaler (CA) | Karpenter |
|---|---|---|
| **Mechanism** | Works with existing Node Groups / ASGs | Directly calls cloud provider APIs (EC2, etc.) |
| **Speed** | Slower — must scale an ASG, then wait for node registration | Much faster — provisions individual nodes directly |
| **Flexibility** | Tied to pre-defined node group instance types | Selects the *optimal* instance type dynamically from any available pool |
| **Bin-packing** | Basic | Advanced — can consolidate underutilized nodes (`consolidation: true`) |
| **Drift Detection** | No | Yes — replaces nodes when they drift from their `NodePool` spec |

**Karpenter's core config:**
```yaml
apiVersion: karpenter.sh/v1beta1
kind: NodePool
metadata:
  name: default
spec:
  template:
    spec:
      requirements:
        - key: karpenter.sh/capacity-type
          operator: In
          values: ["spot", "on-demand"]
        - key: node.kubernetes.io/instance-type
          operator: In
          values: ["m5.large", "m5.xlarge", "m6i.large"]
      nodeClassRef:
        name: default
  limits:
    cpu: 1000
  disruption:
    consolidationPolicy: WhenUnderutilized
```

**Key advantage:** Karpenter can pick `spot` instances automatically and fall back to `on-demand`, saving up to 70% on compute costs while maintaining availability.

---

### Question 50: What is a Pod Topology Spread Constraint and how does it improve on Pod Anti-Affinity for spreading replicas?

**Answer:**
`topologySpreadConstraints` is the modern, more expressive replacement for Pod Anti-Affinity when distributing pods across zones/nodes.

**Anti-Affinity problem:** Hard anti-affinity blocks scheduling entirely if the constraint can't be met (e.g., can't place pod if another pod exists on same node). It gives you no control over *how uneven* the spread can be.

**Topology Spread Constraint:**
```yaml
topologySpreadConstraints:
  - maxSkew: 1
    whenUnsatisfiable: DoNotSchedule
    labelSelector:
      matchLabels:
        app: payment-api
```

- **`maxSkew: 1`:** The difference in pod counts between any two zones must not exceed 1.
- **`whenUnsatisfiable: ScheduleAnyway`:** Soft version — prefers even spread but won't block scheduling.
- **`topologyKey`:** Can be zone, region, node hostname, or any label key.

This provides fine-grained, mathematically bounded spreading across AZs without the all-or-nothing behavior of anti-affinity.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Pod Topology Spread Constraint and how does it improve on Pod Anti-Affinity for spreading replicas?
**Your Response:** "Topology Spread Constraints are the modern way to distribute pods across zones, replacing the old anti-affinity approach. Anti-affinity is rigid - if it can't satisfy the constraint, scheduling fails completely. Topology spread gives me much finer control - I can specify that pod counts between any two zones shouldn't differ by more than 1, and I can make it a soft preference rather than a hard requirement. The topologyKey lets me spread by zone, region, or any custom label. It's like the difference between a strict rule that says 'no two departments can be on same floor' versus a guideline that says 'try to keep departments balanced across floors'. This gives me predictable, mathematically bounded distribution without blocking deployments entirely."


**Answer:**
**cert-manager** is a Kubernetes controller that automates the management and issuance of TLS certificates from various sources (Let's Encrypt, HashiCorp Vault, self-signed CAs).

**Architecture:**
1. You define an `Issuer` or `ClusterIssuer` pointing to a CA (e.g., Let's Encrypt ACME).
2. You create a `Certificate` CRD referencing the Issuer and specifying the DNS name.
3. cert-manager's controller watches for `Certificate` objects, initiates the ACME challenge (HTTP-01 or DNS-01), proves ownership, and stores the issued cert as a Kubernetes `Secret`.
4. Ingress/pods reference the Secret for TLS termination.
5. cert-manager **automatically renews** the certificate before expiry (typically at 2/3 of the validity period).

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: api-tls
spec:
  secretName: api-tls-secret
  issuerRef:
    name: letsencrypt-prod
    kind: ClusterIssuer
  dnsNames:
    - api.example.com
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does cert-manager automate TLS certificate provisioning in Kubernetes?
**Your Response:** "Cert-manager automates the entire TLS certificate lifecycle. I define an Issuer pointing to Let's Encrypt or an internal CA, then create a Certificate object with the domain name. Cert-manager handles the ACME challenges automatically, proves domain ownership, and stores the certificate as a Kubernetes Secret. Ingress controllers or pods can then reference this Secret for TLS termination. The best part is automatic renewal - cert-manager renews certificates before they expire. It's like having an automated security team that handles all certificate paperwork for me - I just declare what domains I need certificates for, and cert-manager handles the rest without me worrying about expiring certificates."

---

### Question 52: Walk through writing a real OPA Rego policy that blocks images not from an approved registry.

**Answer:**
OPA Gatekeeper uses `Rego` language for policies. Here's a real policy that **blocks pods pulling images from outside `gcr.io/company-name`**:

**ConstraintTemplate (defines the policy logic):**
```rego
package k8srequiredregistry

violation[{"msg": msg}] {
  container := input.review.object.spec.containers[_]
  not startswith(container.image, "gcr.io/company-name/")
  msg := sprintf("Container '%v' uses an unapproved image registry: %v", [container.name, container.image])
}
```

**Constraint (enforces the policy):**
```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredRegistry
metadata:
  name: require-internal-registry
spec:
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
  parameters: {}
```

When a Pod is submitted with `image: nginx:latest`, the Validating Webhook invokes OPA, which runs `violation[]`, finds a match, and the API server rejects the Pod with the custom `msg` string.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Walk through writing a real OPA Rego policy that blocks images not from an approved registry.
**Your Response:** "I write OPA policies using Rego language to enforce security rules. Here's a policy that blocks non-approved image registries. The ConstraintTemplate defines the logic using a package, and the Constraint enforces it. The policy checks each container's image field and rejects pods using images outside our approved gcr.io/company-name registry. When someone tries to deploy nginx:latest, OPA evaluates the policy, finds a violation, and Kubernetes rejects the pod with a custom error message. This gives me centralized policy enforcement across the cluster without modifying individual applications. It's like having a security guard at the door that checks everyone's ID against an approved list before letting them enter."

---

### Question 53: Describe how you set up end-to-end secret rotation without restarting pods, using External Secrets Operator (ESO).

**Answer:**
The **External Secrets Operator (ESO)** synchronizes secrets from external stores (AWS Secrets Manager, HashiCorp Vault, GCP Secret Manager) into native Kubernetes Secrets.

**Architecture:**
```
[AWS Secrets Manager] → [ESO Controller] → [K8s Secret] → [Pod Volume Mount]
```

**ExternalSecret resource:**
```yaml
apiVersion: external-secrets.io/v1beta1
kind: ExternalSecret
metadata:
  name: db-credentials
spec:
  refreshInterval: 1h           # ESO re-fetches from AWS every hour
  secretStoreRef:
    name: aws-secretsmanager
    kind: ClusterSecretStore
  target:
    name: db-credentials-k8s   # K8s Secret name
  data:
    - secretKey: password
      remoteRef:
        key: prod/db/password   # AWS Secrets Manager key
```

**Zero-restart rotation pattern:**
1. Rotate the secret in AWS Secrets Manager.
2. ESO detects the change at the next `refreshInterval` and updates the K8s Secret.
3. Deploy **Stakater Reloader** as a sidecar/DaemonSet — it watches K8s Secrets and triggers a **rolling restart** of Deployments that mount them automatically, providing near-zero-downtime rotation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Describe how you set up end-to-end secret rotation without restarting pods, using External Secrets Operator (ESO).
**Your Response:** "I use External Secrets Operator to rotate secrets without pod restarts. ESO syncs secrets from external stores like AWS Secrets Manager into Kubernetes Secrets automatically. I configure it to refresh every hour, and when the external secret changes, ESO updates the Kubernetes Secret. Then I use Stakater Reloader to watch for Secret changes and trigger rolling restarts of Deployments that mount those secrets. This gives me near-zero-downtime rotation because Stakater performs rolling restarts rather than killing all pods at once. It's like having an automated key replacement service - the locksmith changes the locks while people continue working, and only those who need the new keys have to restart."

---

## 🔹 Advanced Observability (Prometheus, eBPF & Hubble) (Questions 54-58)

### Question 54: How does Prometheus scrape metrics from Pods inside Kubernetes without static configuration?

**Answer:**
The **Prometheus Operator** (part of `kube-prometheus-stack` Helm chart) introduces custom CRDs that make scrape configuration dynamic:

- **`ServiceMonitor`:** Tells Prometheus to scrape all pods backing a specific Kubernetes Service.
```yaml
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: payment-api

---

### Question 55: What Kubernetes-specific metrics are most important to monitor? Name the "four golden signals" adapted for K8s.

**Answer:**
Adapted from Google SRE's "four golden signals":

| Signal | K8s Metric | PromQL Example |
|---|---|---|
| **Latency** | API server request duration | `histogram_quantile(0.99, apiserver_request_duration_seconds_bucket)` |
| **Traffic** | Requests per second to a service | `rate(http_requests_total[5m])` |
| **Errors** | Pod restart rate | `rate(kube_pod_container_status_restarts_total[5m])` |
| **Saturation** | Node memory/CPU pressure | `kube_node_status_condition{condition="MemoryPressure",status="true"}` |

**Essential cluster-level alerts:**
- `KubePodCrashLooping` — restarts > 5 in 10 min
- `KubeNodeNotReady` — node `NotReady` > 2 min
- `KubePersistentVolumeFillingUp` — PVC > 85% full
- `KubeDeploymentReplicasMismatch` — desired ≠ ready replicas

---

### Question 56: How does Cilium's Hubble provide deep observability without application code changes?

**Answer:**
**Hubble** is the observability layer built on top of **Cilium** (eBPF-based CNI).

**How it works:**
Because Cilium's eBPF programs run *inside the kernel* on every packet path, they can observe every TCP flow, HTTP request, DNS query, and gRPC call **without touching application code or injecting sidecars** (unlike Istio's Envoy model).

Hubble processes these kernel-level events and exposes them via:
- **Hubble UI:** A real-time service dependency graph showing which services talk to which, with latency and drop rates.
- **Hubble CLI:** `hubble observe --namespace prod --protocol http` streams live HTTP flows.
- **Hubble Relay:** Aggregates flows from all nodes into a centralized Prometheus-compatible metrics endpoint.

**Advantage over Istio Envoy sidecars:**
- No per-pod sidecar memory/CPU overhead.
- Works even for pods that don't support sidecar injection (Windows nodes, HostNetwork pods).
- Sub-microsecond latency overhead vs. Envoy's ~1ms-per-hop overhead.

---

### Question 57: What is Pixie and how does it enable no-instrumentation Kubernetes debugging?

**Answer:**
**Pixie** (by New Relic / CNCF) uses **eBPF** to auto-instrument applications without code changes or agents.

It attaches eBPF probes to:
- **System calls** (open, read, write, connect) to trace I/O and network activity.
- **uprobes** on language runtimes (Go, Python, Java JVM) to capture HTTP/gRPC/SQL spans automatically.

**Capabilities for Kubernetes:**
- Auto-generates per-pod flame graphs.
- Captures golden signals (latency, RPS, errors) per service without Prometheus exporters.
- Can replay full HTTP request/response bodies for failed requests.
- `pxl` (PXL script language) lets you write custom queries: "show me all failing gRPC calls in the `payments` namespace in the last 5 minutes".

**Key interview point:** Pixie ≠ Prometheus. Prometheus requires apps to expose `/metrics`. Pixie captures everything passively at the kernel level.

---

### Question 58: Canary deployments on Kubernetes — how do you implement them without a service mesh?

**Answer:**
Without a service mesh (Istio / Argo Rollouts), you can implement a basic canary using **two Deployments + one Service**:

**Step 1 — Stable deployment (9 replicas):**
```yaml
# deployment-stable.yaml
labels:
  app: payment-api
  version: stable
replicas: 9
```

**Step 2 — Canary deployment (1 replica):**
```yaml
# deployment-canary.yaml
labels:
  app: payment-api   # same label — Service selects both!
  version: canary
replicas: 1
image: payment-api:v2.0
```

**Step 3 — Service selects BOTH:**
```yaml
selector:
  app: payment-api   # matches stable AND canary pods
```
Traffic splits ~90% stable / ~10% canary based purely on replica count. This is **coarse-grained traffic splitting**.

**With Argo Rollouts (production-grade):**
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Rollout
spec:
  strategy:
    canary:
      steps:
        - setWeight: 10     # 10% to canary
        - pause: {duration: 10m}
        - setWeight: 50
        - pause: {duration: 10m}
        - setWeight: 100
      analysis:
        templates:
          - templateName: error-rate-check
```
Argo Rollouts integrates with Prometheus to **automatically abort** if the canary's error rate exceeds a threshold.
