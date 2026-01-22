## 🔹 Cluster Management & Node Operations (Questions 101-110)

### Question 101: How do you upgrade a Kubernetes cluster?

**Answer:**
Typically handled by the provider (EKS/GKE), but for self-hosted (kubeadm):
1.  **Upgrade Control Plane:**
    - Usage `kubeadm upgrade plan`.
    - `kubeadm upgrade apply v1.xx.x`.
    - Upgrade CNI plugin if needed.
2.  **Upgrade Nodes:**
    - `kubectl drain <node> --ignore-daemonsets`.
    - Upgrade `kubelet` and `kube-proxy` packages.
    - Restart services.
    - `kubectl uncordon <node>`.

---

### Question 102: How do you drain a node safely?

**Answer:**
Draining evicts all pods from a node so it can be maintained (rebooted/upgraded).
```bash
kubectl drain node-1 --ignore-daemonsets --delete-emptydir-data
```
- **ignore-daemonsets:** Required because DaemonSet pods cannot be rescheduled (they need to be on that node).
- **delete-emptydir-data:** Required if pods use local storage, acknowledging data loss.

---

### Question 103: What happens during a `kubectl cordon`?

**Answer:**
`cordon` marks a node as **unschedulable**.
- Existing pods continue running.
- New pods will not be scheduled on this node.
- Typically the first step before draining.

---

### Question 104: How to add a node to an existing cluster?

**Answer:**
Using `kubeadm`:
1.  On the Master Node, generate a token:
    ```bash
    kubeadm token create --print-join-command
    ```
2.  Copy the output (it includes token and CA hash).
3.  Run that command on the new Worker Node (sudo required).

---

### Question 105: How to safely remove a node from a Kubernetes cluster?

**Answer:**
1.  **Drain** the node to move workloads: `kubectl drain <node>`.
2.  **Delete** the node object: `kubectl delete node <node>`.
3.  **Reset** the node (on the machine itself): `kubeadm reset` (cleans up iptables, manifests).
4.  Decommission the VM/Hardware.

---

### Question 106: What is the difference between cordon, drain, and delete node?

**Answer:**
- **Cordon:** "Stop sending new work here." (Maintenance prep).
- **Drain:** "Empty the building." (Evict running pods + Cordon).
- **Delete:** "Demolish the building." (Remove from cluster registry).

---

### Question 107: How do you handle node failures in Kubernetes?

**Answer:**
Kubernetes handles it automatically for stateless apps.
1.  Controller Manager detects node is `NotReady` (after --pod-eviction-timeout, default 5m).
2.  It deletes the Pods assigned to that node.
3.  Scheduler reschedules them on healthy nodes.
**Action:** Admin investigates node logs, restarts kubelet, or replaces the VM.

---

### Question 108: How does Kubernetes know a node is unhealthy?

**Answer:**
The **Node Controller** monitors the status reported by `kubelet`.
- Kubelet posts heartbeat to API server every 10s.
- If API server receives no heartbeat for 40s (node-monitor-grace-period), status becomes `Unknown`.
- After 5m (eviction-timeout), pods are evicted.

---

### Question 109: What are node conditions?

**Answer:**
Status fields reported by the node:
- **Ready:** Is it accepting pods?
- **DiskPressure:** Is disk cap low?
- **MemoryPressure:** Is RAM cap low?
- **PIDPressure:** Too many processes?
- **NetworkUnavailable:** Network configured correctly?

---

### Question 110: What is the purpose of a kubeconfig file?

**Answer:**
A YAML file used by `kubectl` to find and authenticate to the cluster.
**Contains:**
1.  **Clusters:** API server URL, CA cert.
2.  **Users:** Client certs, tokens.
3.  **Contexts:** Combines a Cluster + User + Namespace.
Default location: `~/.kube/config`.

---

## 🔹 Performance & Optimization (Questions 111-120)

### Question 111: How do you optimize resource requests and limits?

**Answer:**
1.  **Measure:** Use Prometheus + Grafana or VPA in "Recommendation" mode to see actual usage.
2.  **Requests:** Set to actual average usage (Ensures scheduling).
3.  **Limits:** Set to max burst usage (Prevents starvation of neighbors).
4.  **Strategy:** Avoid setting limits too low (CPU throttling).

---

### Question 112: What happens if a container exceeds its memory limit?

**Answer:**
**OOMKilled**.
- Linux Kernel kills the process.
- Kubernetes sees the container exit code (137).
- If RestartPolicy allows, it restarts the container.
- If it keeps happening -> **CrashLoopBackOff**.

---

### Question 113: What is the effect of setting CPU limits too low?

**Answer:**
The container is **Throttled** (Slowed down).
- Unlike memory, CPU is compressible. The process is not killed.
- It essentially gets less CPU time slices than it wants.
- **Symptom:** App becomes slow, high latency requests.

---

### Question 114: How do you monitor pod resource utilization?

**Answer:**
- **Quick:** `kubectl top pod`.
- **Detailed:** Prometheus metrics (`container_memory_usage_bytes`, `container_cpu_usage_seconds_total`).
- **Dashboard:** Grafana or K8s Dashboard source consumption view.

---

### Question 115: What is the difference between CPU throttling and CPU limits?

**Answer:**
- **Limit:** The configuration value (`limits: cpu: 500m`).
- **Throttling:** The action taken by the CFS (Completely Fair Scheduler) Linux scheduler when the limit is reached. The thread is put to sleep until the next period.

---

### Question 116: What tools can you use for cluster performance tuning?

**Answer:**
1.  **Kube-bench:** Security/CIS benchmarks.
2.  **Kube-burner:** Stress testing / Load generation.
3.  **Goldilocks:** Visualizer for VPA recommendations (helps tune requests).
4.  **Popeye:** Sanitizer that reports resource over/under allocation.

---

### Question 117: How does Kubernetes scheduler select a node?

**Answer:**
Two step process:
1.  **Filtering:** Removes nodes that don't fit (Taints, Memory full, Selector mismatch).
2.  **Scoring:** Ranks remaining nodes (Least requested vs Most requested, Image locality).
The node with the highest score wins.

---

### Question 118: What is bin packing in scheduling?

**Answer:**
A strategy to pack pods tightly onto as few nodes as possible (ResourcePacking).
- **Pros:** Frees up other nodes for scale-down (saves cost).
- **Cons:** High density risks resource contention.
- Configured via scheduler scoring profiles (MostAllocated vs LeastAllocated).

---

### Question 119: How to avoid noisy neighbor issues in Kubernetes?

**Answer:**
1.  **Limits:** Always set Resource Limits.
2.  **QoS Classes:** Ensure critical apps are **Guaranteed** (Request = Limit). BestEffort pods are killed first.
3.  **Isolated Nodes:** Use Taints/Tolerations to dedicate nodes for sensitive workloads.

---

### Question 120: How do resource requests affect scheduling?

**Answer:**
The Scheduler uses **Requests**, not Limits, or current usage.
- If Node A has 4GB RAM total, and existing pods Request 3.5GB (even if using 0), the Scheduler considers the node "Full" for any pod requesting > 0.5GB.

---

## 🔹 Networking - Advanced (Questions 121-130)

### Question 121: What is CNI (Container Network Interface)?

**Answer:**
A CNCF project specifying a standard interface for configuring network interfaces for Linux containers.
- Kubernetes calls CNI plugins to set up networking when creating a pod.
- Allows swapping providers (Calico, Flannel, Cilium) without changing K8s code.

---

### Question 122: How does Kubernetes networking differ from Docker networking?

**Answer:**
- **Docker:** Default is Host-Private (Bridge). requires NAT/Port Mapping to access from outside. Pods on different hosts can't speak by default.
- **Kubernetes:** Flat network model.
  - Every Pod gets a unique IP.
  - Any Pod can talk to any Pod on any Node without NAT.
  - Agents (CNI) handle the routing across nodes.

---

### Question 123: What are common CNI plugins?

**Answer:**
1.  **Flannel:** Simple, overlay network (VXLAN). No NetworkPolicy support.
2.  **Calico:** L3 routing (BGP). High performance. Supports NetworkPolicy.
3.  **Cilium:** eBPF based. High security and visibility.
4.  **Weave Net:** Mesh network.

---

### Question 124: What is Calico and how does it work?

**Answer:**
A popular CNI that provides networking and security.
- Uses **BGP (Border Gateway Protocol)** to route packets (like the internet) instead of encapsulation (overlay), providing near-native speed.
- Also implements the Kubernetes NetworkPolicy API.

---

### Question 125: What is Flannel in Kubernetes?

**Answer:**
One of the simplest CNI plugins.
- Creates an **Overlay Network** using VXLAN.
- Encapsulates packets to transport them across the node network.
- Easy to set up, but has encapsulation overhead and lacks security policy support.

---

### Question 126: What is the role of iptables in Kubernetes?

**Answer:**
`iptables` is used by `kube-proxy` (in iptables mode) to handle Service routing.
- When you call a Service ClusterIP, iptables rules on the node DNAT (Destination NAT) the request to one of the random backend Pod IPs.

---

### Question 127: How does service mesh work in Kubernetes?

**Answer:**
Adds a layer of infrastructure (Control Plane) and injects **Sidecar Proxies** (Data Plane) into every Pod.
- Traffic flows: App A -> Proxy A -> Proxy B -> App B.
- Enables: mTLS, tracing, advanced traffic splitting, retries, without code changes.

---

### Question 128: What is Istio and why use it?

**Answer:**
The most popular Service Mesh.
**Features:**
- **Traffic Management:** Canary, Timeout, Retry.
- **Security:** Automatic mTLS between services.
- **Observability:** Metrics/Logs for all traffic.

---

### Question 129: Explain the envoy proxy in Kubernetes.

**Answer:**
Envoy is a high-performance C++ proxy.
- It is the data plane for Istio.
- It runs as a sidecar in the pod.
- It intercepts all network traffic to apply policies.

---

### Question 130: How is network isolation enforced in Kubernetes?

**Answer:**
By default, it is **NOT** enforced.
- You must use **NetworkPolicies**.
- You must use a CNI that enforces them (Calico, Cilium).
- Without this, any compromised pod can scan/attack other pods in the cluster.

---

## 🔹 Storage - Advanced (Questions 131-140)

### Question 131: What is dynamic volume provisioning?

**Answer:**
The ability for K8s to create cloud storage volumes on-demand.
- **Without it:** Admin manually creates 10 EBS volumes (PVs), User claims one (PVC).
- **With it:** User creates PVC. StorageClass calls AWS API, creates EBS, creates PV, binds it.

---

### Question 132: What is the reclaim policy of a PersistentVolume?

**Answer:**
Determines what happens to the PV when the PVC is deleted.
1.  **Retain:** Volume remains. Admin must manually clean up data.
2.  **Delete:** Volume (e.g., AWS EBS) is deleted. Data lost. (Default for dynamic).
3.  **Recycle:** `rm -rf /volume/*` (Deprecated).

---

### Question 133: How do StatefulSets handle persistent storage?

**Answer:**
They use **volumeClaimTemplates**.
- Unlike Deployments (which share one PVC if defined), StatefulSet creates a **unique PVC** for each replica (`data-web-0`, `data-web-1`).
- Ensures each sticky identity has its own sticky storage.

---

### Question 134: How do you back up Kubernetes volumes?

**Answer:**
1.  **Volume Snapshots:** K8s API Standard (`VolumeSnapshot`).
2.  **Tools:** Velero (Takes disk snapshots of PVs + YAMLs).
3.  **Manual:** Run a pod mounting the volume, tar/rsync data to S3.

---

### Question 135: How do you resize a PVC?

**Answer:**
1.  Verify StorageClass has `allowVolumeExpansion: true`.
2.  Edit PVC: `kubectl edit pvc my-pvc`. Increase `storage` size.
3.  Apply.
   - Cloud provider resizes disk.
   - File system expands (often requires pod restart or online resize support).

---

### Question 136: What happens if a pod uses a PVC that no longer exists?

**Answer:**
The Pod will stick in **Pending** or **ContainerCreating** state.
Events will show `FailedMount`: "Volume not found".

---

### Question 137: What is volume binding mode?

**Answer:**
In StorageClass:
1.  **Immediate:** Create volume as soon as PVC is created. (Risky: might create in Zone A, but Pod needs to run in Zone B).
2.  **WaitForFirstConsumer:** Delay volume creation until the Pod is scheduled. Creates volume in the **same zone** as the Node.

---

### Question 138: How do you share storage between multiple pods?

**Answer:**
Requires a volume plugin that supports `ReadWriteMany` (RWX).
- **Supported:** NFS, CephFS, Azure File, AWS EFS.
- **Not Supported:** AWS EBS, Azure Disk (These are block devices, ReadWriteOnce only).

---

### Question 139: What is ephemeral storage?

**Answer:**
Local storage used by the container (logs, overlay fs, emptyDir).
- It consumes the Node's root disk.
- If ephemeral storage fills up, the node enters `DiskPressure` and evicts pods.

---

### Question 140: Difference between CSI and in-tree volume plugins?

**Answer:**
- **In-Tree:** Volume code (AWS, GCP) was part of K8s core binary. Difficult to update.
- **CSI (Container Storage Interface):** Standard plugin system. Storage vendors release their own drivers separately. K8s calls them via RPC.

---

## 🔹 Authentication & Authorization (Questions 141-150)

### Question 141: What is the difference between RBAC and ABAC?

**Answer:**
- **RBAC (Role Based):** "Admins can do everything." (Static Roles). Standard in K8s.
- **ABAC (Attribute Based):** "User 'Bob' can 'create pods' only if 'env=dev' and 'time=9am'." (Policy Based). Harder to configure, file-based.

---

### Question 142: What is a ClusterRole vs Role?

**Answer:**
- **Role:** Namespaced. Grants permissions within a specific namespace.
- **ClusterRole:** Cluster-wide. Grants permissions across all namespaces (e.g., read Nodes) OR can be reused in namespace bindings.

---

### Question 143: How do you create a custom ClusterRole?

**Answer:**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "list"]
```

---

### Question 144: How to bind a Role to a user?

**Answer:**
Use a **RoleBinding**.
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
- kind: User
  name: jane
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

---

### Question 145: What is a ServiceAccount token?

**Answer:**
A JWT (JSON Web Token) signed by the cluster CA.
- Stored in a Secret (old K8s) or projected volume (new K8s).
- Used by client libraries inside the pod to authenticate to the API Server.

---

### Question 146: How do you secure access to the Kubernetes Dashboard?

**Answer:**
1.  **Do NOT:** Expose via NodePort/LoadBalancer required public internet.
2.  **Do:** Use `kubectl proxy` (localhost access).
3.  **Do:** Integrate with OIDC (OAuth) / Ingress with Auth.
4.  **RBAC:** Ensure the user logging in has restricted rights.

---

### Question 147: What is OIDC authentication in Kubernetes?

**Answer:**
K8s asks an Identity Provider (Google/Okta/Keycloak) "Who is this?".
1.  User logins to Provider -> Gets ID Token.
2.  Sends ID Token to K8s API.
3.  K8s verifies signature and maps `sub` claim to User and `groups` claim to Groups.

---

### Question 148: How do you rotate API server certificates?

**Answer:**
If certificates expire (usually 1 year), K8s stops working.
- **Kubeadm:** `kubeadm certs renew all` -> Restart control plane static pods.
- **Managed (EKS):** Handled by cloud provider automatically.

---

### Question 149: What is impersonation in Kubernetes?

**Answer:**
Allows a user (Admin) to act as another user.
`kubectl get pods --as=jane --as-group=devs`.
- Useful for testing RBAC rules without logging out/in.

---

### Question 150: How do audit logs work in Kubernetes?

**Answer:**
Kubernetes can log every request to the API server (Who, What, When, Result).
- Configured via Audit Policy (YAML).
- Logs sent to backend (Json file / Webhook).
- Critical for compliance security (detecting unauthorized access attempts).

---
