# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 3)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Robinhood, etc.
> **Focus deep-dive:** High Availability, Scale-out Tuning, Resource Evictions limits, and Advanced State Management (CSI & Operators).

---

## 🔹 Scaling & Reliability Tuning (Questions 26-30)

### Question 26: How does the Kubernetes Scheduler handle custom resources like GPUs during placement?

**Answer:**
Kubernetes supports **Device Plugins** (like the Nvidia GPU device plugin).
When deployed as a DaemonSet, the plugin registers the hardware resources on the node (e.g., `nvidia.com/gpu: 4`) with the `kubelet`. The kubelet communicates this to the APIServer. When a user defines a Pod requesting `nvidia.com/gpu: 1` under resources, the Scheduler treats it purely as an extended integer resource (like CPU) and successfully targets an eligible node, delegating physical GPU bindings to the local Device Plugin payload.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the Kubernetes Scheduler handle custom resources like GPUs during placement?
**Your Response:** "Kubernetes handles GPUs through Device Plugins. I deploy the Nvidia GPU device plugin as a DaemonSet, which registers the available GPUs on each node with the kubelet. The kubelet reports these as extended resources like 'nvidia.com/gpu: 4' to the API server. When a pod requests a GPU, the scheduler treats it like any other resource request and finds a node with available GPUs. The actual GPU binding happens locally on the node through the device plugin. It's like having specialized equipment reservations - the scheduler knows which rooms have projectors and books meetings accordingly, while the local facilities manager handles the actual equipment setup."

---

### Question 27: Explain Descheduler. Why is it needed when the Scheduler already picks the best node?

**Answer:**
The standard K8s Scheduler only looks at conditions *at the exact moment* a Pod is scheduled.
If conditions change wildly days later (nodes are added/removed, taints change, pods are terminated freeing resources), older Pods remain sub-optimally placed locally causing cluster hotspots.
The **Descheduler** routinely checks the cluster against optimization policies (e.g., `RemoveDuplicates`, `LowNodeUtilization`). If it finds Pods violating optimal placement over time, it aggressively targets them for eviction so they are forced into the Scheduler again to find a newly balanced home.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Descheduler. Why is it needed when the Scheduler already picks the best node?
**Your Response:** "The regular scheduler makes decisions based on conditions at the exact moment a pod is scheduled, but clusters are dynamic. Days later, conditions might change - new nodes added, pods terminated, resources freed up - leaving older pods in suboptimal locations. The Descheduler is like a periodic housekeeper that continuously checks if pods are still in the right places. If it finds pods that could be better placed elsewhere - like multiple copies of the same pod on one node, or nodes that are underutilized - it evicts them so the scheduler can place them optimally. It's the difference between one-time planning versus ongoing optimization of resource distribution."

---

### Question 28: A Pod stays `Pending` constantly with the error `FailedScheduling`. If resources (CPU/Mem) are available across the NodePool, what causes this?

**Answer:**
Several constraints can block the scheduler:
1. **Topology Spread Constraints:** Blocking placement if putting a pod on the available node would cause a zone imbalance.
2. **PodAffinity / AntiAffinity:** The node doesn't have other pods matching an affinity rule, or it violates a strict anti-affinity limit entirely.
3. **No Tolerations matches Taint:** The available nodes are fully tainted.
4. **PVC Volume Binding:** The node is in Availability Zone East, but the dynamically provisioned EBS Volume (PVC) was constrained/physically placed in AZ West. The pod cannot attach the remote physical drive across regions.

### How to Explain in Interview (Spoken style format)
**Interviewer:** A Pod stays `Pending` constantly with the error `FailedScheduling`. If resources (CPU/Mem) are available across the NodePool, what causes this?
**Your Response:** "Even with available CPU and memory, several constraints can block scheduling. Topology spread constraints might prevent placing a pod if it would unbalance zones. Pod affinity rules might require the pod to be near other specific pods that aren't on the available node. Anti-affinity rules might prevent placing it near certain pods. The node might be tainted without matching tolerations. A common issue is volume-zone mismatch - the pod needs a PVC that's physically tied to a different availability zone, so it can't attach the volume across regions. I'd check the pod events and describe output to see which constraint is failing. It's like having a hotel room available but not being able to book it due to various policy restrictions."

---

### Question 29: Describe how the `oomScore` kernel parameter interacts with Kubernetes Quality of Service (QoS) classes.

**Answer:**
When Linux hits physical memory exhaustion, the Kernel OOMKiller ranks processes using `oomScore` (higher score = killed first, [-1000 to +1000]).
Kubelet modifies the `oom_score_adj` based on the K8s QoS:
1. **Guaranteed (Requests == Limits):** Scored lowest (e.g., `-997`). Last to be killed natively. (Highest priority).
2. **Burstable (Requests < Limits):** Scored middle. Based on how much above their request they currently utilize.
3. **BestEffort (No Requests/Limits):** Scored highest (`+1000`). Immediate first targets for kernel slaughter when nodes run out of physical memory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Describe how the `oomScore` kernel parameter interacts with Kubernetes Quality of Service (QoS) classes.
**Your Response:** "When a node runs out of memory, the Linux OOMKiller needs to decide which processes to kill first. Kubernetes sets the oom_score_adj based on pod QoS classes. Guaranteed pods get the lowest score (around -997) so they're killed last - they have predictable resource usage. Burstable pods get middle scores based on how much they're using above their requests. BestEffort pods get the highest score (+1000) so they're killed first since they made no resource promises. This ensures critical workloads survive memory pressure. It's like having different priority levels on a lifeboat - guaranteed passengers get the best seats, burstable get middle priority, and best-effort passengers are first to be moved if the boat gets too heavy."

---

### Question 30: What is Keda and how does it compare to the default K8s Horizontal Pod Autoscaler (HPA)?

**Answer:**
**KEDA (Kubernetes Event-Driven Autoscaling)** is an advanced operator.
The default HPA struggles to scale applications based on *external events* (like RabbitMQ/Kafka queue depths or AWS SQS messages). Trying to configure the custom-metrics API for that is incredibly tedious.
KEDA natively extends the HPA by providing dozens of pre-built "Scalers". It reads external metrics safely and creates HPA objects injected dynamically. Crucially, KEDA supports scaling workloads completely down to **Zero** when there's no event queue, a feature native HPA fundamentally limits (it stops at `minReplicas: 1`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Keda and how does it compare to the default K8s Horizontal Pod Autoscaler (HPA)?
**Your Response:** "KEDA is an event-driven autoscaler that extends HPA for external events. Standard HPA works well for CPU/memory metrics but struggles with external triggers like Kafka queue depth or SQS messages. KEDA provides pre-built scalers for dozens of external systems - it reads queue lengths or other event metrics and automatically creates HPA objects. The killer feature is scaling to zero - when there are no events, KEDA can scale completely to zero pods, while standard HPA always keeps at least one pod running. I use KEDA for event-driven workloads like message processors, and standard HPA for web APIs. It's like having a smart thermostat that turns off completely when no one's home, versus a regular one that always keeps the house at minimum temperature."

---

## 🔹 State & Workload Optimization (Questions 31-35)

### Question 31: How does Kubernetes handle container image caching, and what are the security implications of `imagePullPolicy: IfNotPresent` in production?

**Answer:**
`IfNotPresent` caches the pulled image by its tag on the node.
- **Security Flaw:** If you deploy `my-app:v1` in `Dev` namespace, the node pulls and caches `my-app:v1`. If a user in `Prod` namespace spins up a pod explicitly specifying `my-app:v1`, they might inadvertently get that exact locally cached image without ever verifying registry permissions! (Though managed if `AlwaysPullImages` admission control is activated natively).
- **Optimization Strategy:** For massive image loads across 1000 nodes, use peer-to-peer image registries like Dragonfly or Kraken to prevent DDoS-ing the centralized registry concurrently.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Kubernetes handle container image caching, and what are the security implications of `imagePullPolicy: IfNotPresent` in production?
**Your Response:** "The IfNotPresent policy caches images locally on nodes, which is great for performance but has security implications. If I deploy my-app:v1 in dev namespace, it gets cached on the node. Later, someone in prod namespace could deploy the same tag and get the cached image without proper registry authentication. To mitigate this, I enable AlwaysPullImages admission control to force registry checks every time. For large-scale deployments, I use peer-to-peer image distribution like Dragonfly to avoid DDoS-ing the central registry. It's like having a local library cache - convenient but needs security controls to ensure people are always getting authorized, up-to-date books rather than whatever happens to be on the shelf."

---

### Question 32: Explain the Container Storage Interface (CSI) architectural separation.

**Answer:**
Before CSI, storage provisioners (like AWS EBS or Ceph) were "in-tree" (compiled directly into Kubernetes core codebase binary). This meant any EBS bug required a full K8s version release to patch natively.
**CSI (Container Storage Interface)** moved storage "out-of-tree" using standard gRPC endpoints. Cloud providers now write their independent CSI Driver Operators which run as DaemonSets on nodes. K8s simply tells the CSI driver via standard interfaces, "Attach volume X to Node Y", separating core cluster life cycles from external cloud vendors entirely.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the Container Storage Interface (CSI) architectural separation.
**Your Response:** "CSI moved storage drivers from being compiled into Kubernetes core to being external plugins. Before CSI, if AWS EBS had a bug, we needed a whole Kubernetes release to fix it. With CSI, cloud providers write their own storage drivers that run as DaemonSets. Kubernetes just makes standard gRPC calls like 'attach volume X to node Y' and the CSI driver handles the specifics. This separation means storage vendors can update their drivers independently of Kubernetes releases. It's like having universal USB ports instead of having to build devices directly into the computer - any storage vendor can plug in their driver using the standard interface."

---

### Question 33: How does a StatefulSet perform rolling updates without causing split-brain or data corruption?

**Answer:**
Unlike a Deployment that aggressively replaces subsets of Pods dynamically or spun up concurrently in parallel, a **StatefulSet updates pods strictly sequentially in reverse ordinal order**.
If you have `web-0`, `web-1`, `web-2`:
The controller terminates `web-2`, waits for it to completely terminate locally, spins up the new `web-2`, waits for it to become *Ready* (passes readiness probe entirely), and then—and only then—moves systematically to `web-1`, protecting database replication strict synchronization loops from mass parallel failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a StatefulSet perform rolling updates without causing split-brain or data corruption?
**Your Response:** "StatefulSets are much more careful than Deployments during updates. While Deployments might replace many pods at once, StatefulSets update pods one at a time in reverse order - starting with web-2, then web-1, then web-0. It waits for each pod to fully terminate and the new pod to become ready before moving to the next one. This sequential approach is crucial for databases where you need to maintain quorum and avoid split-brain scenarios. It's like renovating a bridge - you close one lane at a time and ensure it's fully functional before opening the next, rather than shutting down multiple lanes simultaneously and risking collapse."

---

### Question 34: What is Headroom, and how do you calculate optimal cluster over-provisioning?

**Answer:**
**Headroom** refers to keeping intentional unused capacity ("dummy pods" usually via a priority-class trick) within a cluster so that when a massive traffic spike triggers the HPA, pods immediately deploy without waiting for the 3-to-4-minute delay of the Cluster Autoscaler booting new EC2 VMs linearly.
You run "Pause" containers with `LowPriority`. When high-priority app pods are scaled in, the scheduler immediately evicts the Pause pods (reclaiming their ready CPU/Mem space), causing the Cluster autoscaler to only boot instances *after* the fast-reaction occurs natively.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Headroom, and how do you calculate optimal cluster over-provisioning?
**Your Response:** "Headroom is like keeping spare capacity ready for immediate use. I run low-priority 'pause' pods that consume resources but can be instantly evicted when real applications need to scale. When traffic spikes and HPA needs more pods, they immediately take the space from the pause pods instead of waiting 3-4 minutes for the Cluster Autoscaler to boot new nodes. The Cluster Autoscaler then replaces the evicted pause pods by adding new nodes. This gives me instant scaling response while still maintaining optimal resource utilization. It's like having a few backup servers on standby that can be repurposed instantly, while the procurement process orders new hardware in the background."

---

### Question 35: Describe the lifecycle hooks: PostStart and PreStop.

**Answer:**
- **PostStart:** Executes a command immediately after a container is created. Note: It executes asynchronously. There is no guarantee it completes before the container’s ENTRYPOINT natively starts.
- **PreStop:** Crucial for zero-downtime bounds! Executed synchronously immediately before a container is sent `SIGTERM`. The kubelet waits for this hook to complete completely before moving to terminate. 
**Usage:** Writing `sleep 10` in `PreStop` allows time for the API-server endpoints to cleanly deregister the IP across the cluster routing mesh, while the pod finishes its last inflight HTTP streams gracefully.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Describe the lifecycle hooks: PostStart and PreStop.
**Your Response:** "PostStart and PreStop are lifecycle hooks that help with graceful container management. PostStart runs right after a container starts, but asynchronously - so it might run alongside the main application. I use it for initialization tasks like warming up caches. PreStop is crucial for zero-downtime deployments - it runs synchronously right before SIGTERM, and Kubernetes waits for it to complete before killing the container. I often put a short sleep in PreStop to give time for the service endpoint deregistration to propagate, allowing the pod to finish serving existing requests gracefully. It's like having a welcome routine when someone starts work and a proper checkout process when they leave - ensuring clean transitions both ways."
