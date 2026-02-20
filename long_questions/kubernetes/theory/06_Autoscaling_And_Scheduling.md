# 🟣 Autoscaling & Scheduling

---

### 1. What is Horizontal Pod Autoscaler (HPA)?

"HPA automatically **scales the number of pod replicas** based on observed metrics — CPU, memory, or custom metrics from your monitoring stack.

You define: `minReplicas`, `maxReplicas`, and target utilization. If pods are at 80% CPU and the target is 50%, HPA scales up. If utilization drops below target (with a stabilization window), it scales down.

In production, I configure HPA based on **custom business metrics** (request rate, queue depth) rather than CPU alone — CPU is a lagging indicator."

#### In Depth
HPA v2 (the current API) supports multiple metrics simultaneously with `AND` semantics — all metrics must be within target for scale-down, but ANY metric above threshold triggers scale-up. The scale-up/down behavior is controlled by `behavior` settings: `scaleUp.stabilizationWindowSeconds` (default 0s — scales up immediately) and `scaleDown.stabilizationWindowSeconds` (default 300s — waits 5 min before scaling down). Adjust these to prevent thrashing.

---

### 2. What is Vertical Pod Autoscaler (VPA)?

"VPA automatically **adjusts the CPU and memory requests/limits** of containers based on actual usage.

Instead of scaling horizontally (more pods), it scales vertically (give each pod more resources).

I use VPA in `Off` mode first — just to see recommendations — then move to `Auto` for non-critical workloads. For production databases or anything with a strong memory requirement, VPA recommendations help right-size without manual tuning."

#### In Depth
VPA has three modes: `Off` (recommend only), `Initial` (apply on pod creation only, no future changes), and `Auto` (apply changes, which **restarts the pod**). The restart-on-resize is VPA's biggest limitation — it doesn't work with `PodDisruptionBudget`-protected workloads well. VPA and HPA cannot currently be used together on the same pod for the same metric type (CPU/memory) — they fight each other.

---

### 3. What is Cluster Autoscaler?

"Cluster Autoscaler automatically **adds or removes nodes** from the cluster based on pod scheduling needs.

If pods are `Pending` because there aren't enough resources on existing nodes, CA provisions new nodes. If nodes are consistently underutilized, CA drains and terminates them.

In GKE, EKS, and AKS, CA integrates with the cloud's node group APIs. I configure it with safe limits: minimum 2 nodes per zone (for HA) and maximum based on cost budget."

#### In Depth
CA has a `scan-interval` (default 10s) and a `scale-down-delay-after-add` (default 10 min). Nodes are only considered for scale-down if they've been underutilized (< 50% allocatable resources) for the `scale-down-unneeded-time` (default 10 min). Nodes annotated with `cluster-autoscaler.kubernetes.io/scale-down-disabled: "true"` are never terminated by CA. Critical control plane and monitoring nodes should have this annotation.

---

### 4. How does HPA work internally?

"HPA is implemented as a **control loop** running in the kube-controller-manager.

Every 15 seconds (configurable via `--horizontal-pod-autoscaler-sync-period`): it queries the metrics API (Metrics Server for CPU/memory, or custom metrics API for business metrics), calculates the desired replica count using the formula `desiredReplicas = ceil(currentReplicas * (currentMetricValue / desiredMetricValue))`, and then updates the Deployment's replica count if needed.

The Metrics Server aggregates kubelet cAdvisor data. For custom metrics, I deploy the Prometheus Adapter to expose Prometheus metrics through the custom metrics API."

#### In Depth
HPA ignores pods that are not yet Ready, and pods in `PodPending` state. For scale-up decisions, it uses the higher estimate (assumes pending pods would hit 100% utilization). The `--horizontal-pod-autoscaler-tolerance` (default 0.1, meaning 10%) prevents constant scaling from small fluctuations — scaling only happens if the metric deviates more than 10% from target.

---

### 5. What is KEDA (Kubernetes Event-Driven Autoscaler)?

"KEDA extends HPA to scale based on **external event sources** — Kafka topic lag, RabbitMQ queue depth, AWS SQS queue length, Azure Service Bus, HTTP request rate, and 60+ other sources.

Without KEDA, HPA can only use CPU, memory, and custom K8s metrics. KEDA bridges the gap.

My favorite use case: scale a Kafka consumer Deployment based on consumer group lag. When lag is 0, scale to 0 pods. When lag grows, scale up. This is true event-driven scaling and can reduce costs dramatically during off-peak hours."

#### In Depth
KEDA works by creating HPA objects under the hood and supplying custom metrics. `ScaledObject` is the KEDA CRD — it defines what to scale, the event source, trigger conditions, and min/max replicas. KEDA can scale to zero (`minReplicaCount: 0`), which standard HPA cannot do (minimum is 1). For scale-from-zero, KEDA uses a different mechanism than HPA to detect the first event.

---

### 6. How does the Kubernetes Scheduler work?

"The scheduler works in two phases:

**Filtering**: Find all nodes that can run this pod — checking resources, taints/tolerations, affinity rules, PVC binding topology, etc. Nodes that fail any check are eliminated.

**Scoring**: Among the passing nodes, score each one using multiple prioritization functions (resource balancing, pod affinity preference, image locality, etc.). Pick the highest-scored node.

If multiple nodes tie, one is picked randomly. The scheduler then creates a **Binding** object, which the kubelet picks up to start the pod."

#### In Depth
The scheduler is extensible via the **Scheduling Framework** — you can write plugins that hook into different scheduling phases: PreFilter, Filter, PostFilter, PreScore, Score, Reserve, Permit, PreBind, Bind, PostBind. Multiple schedulers can run in parallel (each pod's `spec.schedulerName` determines which scheduler handles it). You can also use scheduling extenders (webhook-based) to integrate external scheduling logic.

---

### 7. What is pod affinity and anti-affinity?

"Pod affinity and anti-affinity control **where pods are scheduled relative to other pods**.

- **Affinity**: 'Schedule me near pods with label X' — useful for co-location (e.g., a frontend pod near its backend for low latency).
- **Anti-affinity**: 'Don't schedule me on a node (or zone) that already has a pod with label X' — essential for spreading replicas across failure domains.

I always configure anti-affinity for HA deployments: `topologyKey: topology.kubernetes.io/zone` ensures replicas spread across availability zones."

#### In Depth
Two preference modes: `requiredDuringSchedulingIgnoredDuringExecution` (hard requirement — if not satisfiable, pod stays `Pending`) and `preferredDuringSchedulingIgnoredDuringExecution` (soft preference — scheduler tries, but proceeds anyway). The `IgnoredDuringExecution` part means affinity rules aren't enforced after scheduling — if a node's labels change after the pod is placed, the pod isn't evicted. `RequiredDuringExecution` support was planned but not yet released.

---

### 8. What are taints and tolerations?

"Taints are placed on **nodes** to repel pods. Tolerations on **pods** allow them to land on tainted nodes.

Example: `kubectl taint nodes gpu-node-1 hardware=gpu:NoSchedule` — only pods with `tolerations: [{key: hardware, value: gpu}]` can be scheduled there.

I use taints to dedicate nodes for specific workloads: GPU nodes for ML jobs, high-memory nodes for Elasticsearch, on-call nodes for low-latency services."

#### In Depth
Taint effects: `NoSchedule` (new pods aren't scheduled, existing pods stay), `PreferNoSchedule` (soft version — scheduler tries to avoid), `NoExecute` (new pods not scheduled AND existing pods without toleration are evicted after `tolerationSeconds`). Control plane nodes have taints by default — `node-role.kubernetes.io/control-plane:NoSchedule`. This is why workload pods don't land there unless you explicitly add the toleration.

---

### 9. What is pod topology spread constraints?

"Topology spread constraints are a modern, declarative way to **distribute pods evenly** across topology domains (zones, regions, nodes).

Instead of using complex affinity rules, you say: 'I want my pods spread across zones with a maximum skew of 1'.

`maxSkew: 1, topologyKey: zone, whenUnsatisfiable: DoNotSchedule` means: across all zones, the difference between the most and least loaded zones must be ≤ 1 pod."

#### In Depth
Topology spread constraints are often combined with anti-affinity. They supersede the `podAntiAffinity + topologyKey` pattern for even spreading because they actively balance — they count pods and enforce distribution. The new `matchLabelKeys` field (beta in 1.27) uses deployment pod-template hash to correctly handle rolling updates without counting old pods toward the spread calculation.

---

### 10. What is a custom scheduler in Kubernetes?

"A custom scheduler is an alternative scheduler you write that implements your own placement logic.

You deploy it as a regular pod, and any pod with `spec.schedulerName: my-custom-scheduler` gets handled by it instead of the default scheduler.

Use cases: ML frameworks that need gang scheduling (all pods in a job must be scheduled simultaneously or not at all), or NUMA-aware placement that the default scheduler doesn't handle optimally."

#### In Depth
Writing a full custom scheduler from scratch is complex. The **Scheduling Framework** is a better approach — it lets you write plugins that extend the default scheduler. The community project `Volcano` is a popular production-ready custom scheduler for batch/ML workloads, supporting gang scheduling, queue management, and fair sharing.

---

### 11. What is a node selector?

"nodeSelector is the **simplest form of node scheduling constraint** — it requires pods to be scheduled on nodes with matching labels.

```yaml
nodeSelector:
  disktype: ssd
```

This is the easiest way to pin a pod to a specific type of node. However, `nodeAffinity` is strictly more powerful and should be preferred for new workloads."

#### In Depth
nodeSelector is implemented as a filtering step in the scheduler — nodes without the required labels are excluded from consideration. Labels on nodes can be set manually (`kubectl label node`) or automatically by the cloud provider (e.g., `topology.kubernetes.io/region`, `node.kubernetes.io/instance-type`). Node Feature Discovery (NFD) can automatically label nodes based on hardware capabilities (CPUs with AVX512, GPUs, FPGAs).

---

### 12. How does the Cluster Autoscaler work with node groups?

"Cluster Autoscaler works with **node groups** (AWS Auto Scaling Groups, GCP Managed Instance Groups, Azure VMSS).

When CA decides to scale up, it picks the most cost-efficient node group that can satisfy the pending pod's requests and simulates what would happen if a node from that group were added.

When CA decides to scale down, it checks if all pods on a node can be safely moved elsewhere (respecting PDB, local storage, etc.), then cordon and drain the node before terminating the cloud instance."

#### In Depth
CA's expander policy determines which node group gets a new node during scale-up: `random`, `least-waste`, `most-pods`, `price`, or `priority`. In mixed-workload clusters, I use `least-waste` to minimize resource fragmentation. CA can also be configured with **overprovisioning** (placeholder pods that keep extra capacity available), reducing the latency of scale-up from ~3 minutes to near-instant.

---
