# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 3)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Robinhood, etc.
> **Focus deep-dive:** High Availability, Scale-out Tuning, Resource Evictions limits, and Advanced State Management (CSI & Operators).

---

## 🔹 Scaling & Reliability Tuning (Questions 26-30)

### Question 26: How does the Kubernetes Scheduler handle custom resources like GPUs during placement?

**Answer:**
Kubernetes supports **Device Plugins** (like the Nvidia GPU device plugin).
When deployed as a DaemonSet, the plugin registers the hardware resources on the node (e.g., `nvidia.com/gpu: 4`) with the `kubelet`. The kubelet communicates this to the APIServer. When a user defines a Pod requesting `nvidia.com/gpu: 1` under resources, the Scheduler treats it purely as an extended integer resource (like CPU) and successfully targets an eligible node, delegating physical GPU bindings to the local Device Plugin payload.

---

### Question 27: Explain Descheduler. Why is it needed when the Scheduler already picks the best node?

**Answer:**
The standard K8s Scheduler only looks at conditions *at the exact moment* a Pod is scheduled.
If conditions change wildly days later (nodes are added/removed, taints change, pods are terminated freeing resources), older Pods remain sub-optimally placed locally causing cluster hotspots.
The **Descheduler** routinely checks the cluster against optimization policies (e.g., `RemoveDuplicates`, `LowNodeUtilization`). If it finds Pods violating optimal placement over time, it aggressively targets them for eviction so they are forced into the Scheduler again to find a newly balanced home.

---

### Question 28: A Pod stays `Pending` constantly with the error `FailedScheduling`. If resources (CPU/Mem) are available across the NodePool, what causes this?

**Answer:**
Several constraints can block the scheduler:
1. **Topology Spread Constraints:** Blocking placement if putting a pod on the available node would cause a zone imbalance.
2. **PodAffinity / AntiAffinity:** The node doesn't have other pods matching an affinity rule, or it violates a strict anti-affinity limit entirely.
3. **No Tolerations matches Taint:** The available nodes are fully tainted.
4. **PVC Volume Binding:** The node is in Availability Zone East, but the dynamically provisioned EBS Volume (PVC) was constrained/physically placed in AZ West. The pod cannot attach the remote physical drive across regions.

---

### Question 29: Describe how the `oomScore` kernel parameter interacts with Kubernetes Quality of Service (QoS) classes.

**Answer:**
When Linux hits physical memory exhaustion, the Kernel OOMKiller ranks processes using `oomScore` (higher score = killed first, [-1000 to +1000]).
Kubelet modifies the `oom_score_adj` based on the K8s QoS:
1. **Guaranteed (Requests == Limits):** Scored lowest (e.g., `-997`). Last to be killed natively. (Highest priority).
2. **Burstable (Requests < Limits):** Scored middle. Based on how much above their request they currently utilize.
3. **BestEffort (No Requests/Limits):** Scored highest (`+1000`). Immediate first targets for kernel slaughter when nodes run out of physical memory.

---

### Question 30: What is Keda and how does it compare to the default K8s Horizontal Pod Autoscaler (HPA)?

**Answer:**
**KEDA (Kubernetes Event-Driven Autoscaling)** is an advanced operator.
The default HPA struggles to scale applications based on *external events* (like RabbitMQ/Kafka queue depths or AWS SQS messages). Trying to configure the custom-metrics API for that is incredibly tedious.
KEDA natively extends the HPA by providing dozens of pre-built "Scalers". It reads external metrics safely and creates HPA objects injected dynamically. Crucially, KEDA supports scaling workloads completely down to **Zero** when there's no event queue, a feature native HPA fundamentally limits (it stops at `minReplicas: 1`).

---

## 🔹 State & Workload Optimization (Questions 31-35)

### Question 31: How does Kubernetes handle container image caching, and what are the security implications of `imagePullPolicy: IfNotPresent` in production?

**Answer:**
`IfNotPresent` caches the pulled image by its tag on the node.
- **Security Flaw:** If you deploy `my-app:v1` in `Dev` namespace, the node pulls and caches `my-app:v1`. If a user in `Prod` namespace spins up a pod explicitly specifying `my-app:v1`, they might inadvertently get that exact locally cached image without ever verifying registry permissions! (Though managed if `AlwaysPullImages` admission control is activated natively).
- **Optimization Strategy:** For massive image loads across 1000 nodes, use peer-to-peer image registries like Dragonfly or Kraken to prevent DDoS-ing the centralized registry concurrently.

---

### Question 32: Explain the Container Storage Interface (CSI) architectural separation.

**Answer:**
Before CSI, storage provisioners (like AWS EBS or Ceph) were "in-tree" (compiled directly into Kubernetes core codebase binary). This meant any EBS bug required a full K8s version release to patch natively.
**CSI (Container Storage Interface)** moved storage "out-of-tree" using standard gRPC endpoints. Cloud providers now write their independent CSI Driver Operators which run as DaemonSets on nodes. K8s simply tells the CSI driver via standard interfaces, "Attach volume X to Node Y", separating core cluster life cycles from external cloud vendors entirely.

---

### Question 33: How does a StatefulSet perform rolling updates without causing split-brain or data corruption?

**Answer:**
Unlike a Deployment that aggressively replaces subsets of Pods dynamically or spun up concurrently in parallel, a **StatefulSet updates pods strictly sequentially in reverse ordinal order**.
If you have `web-0`, `web-1`, `web-2`:
The controller terminates `web-2`, waits for it to completely terminate locally, spins up the new `web-2`, waits for it to become *Ready* (passes readiness probe entirely), and then—and only then—moves systematically to `web-1`, protecting database replication strict synchronization loops from mass parallel failures.

---

### Question 34: What is Headroom, and how do you calculate optimal cluster over-provisioning?

**Answer:**
**Headroom** refers to keeping intentional unused capacity ("dummy pods" usually via a priority-class trick) within a cluster so that when a massive traffic spike triggers the HPA, pods immediately deploy without waiting for the 3-to-4-minute delay of the Cluster Autoscaler booting new EC2 VMs linearly.
You run "Pause" containers with `LowPriority`. When high-priority app pods are scaled in, the scheduler immediately evicts the Pause pods (reclaiming their ready CPU/Mem space), causing the Cluster autoscaler to only boot instances *after* the fast-reaction occurs natively.

---

### Question 35: Describe the lifecycle hooks: PostStart and PreStop.

**Answer:**
- **PostStart:** Executes a command immediately after a container is created. Note: It executes asynchronously. There is no guarantee it completes before the container’s ENTRYPOINT natively starts.
- **PreStop:** Crucial for zero-downtime bounds! Executed synchronously immediately before a container is sent `SIGTERM`. The kubelet waits for this hook to complete completely before moving to terminate. 
**Usage:** Writing `sleep 10` in `PreStop` allows time for the API-server endpoints to cleanly deregister the IP across the cluster routing mesh, while the pod finishes its last inflight HTTP streams gracefully.
