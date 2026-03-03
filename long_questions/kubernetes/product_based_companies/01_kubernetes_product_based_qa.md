# 🚀 Kubernetes Interview Questions - Product-Based Companies
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Snowflake etc.
> **Focus deep-dive:** etcd internals, consensus, Pod lifecycle events, CNI/Iptables internals, CRDs, multi-tenant security, scheduler algorithms, and high-level platform engineering scenarios.

---

## 🔹 Advanced Kubernetes Architecture & Internals (Questions 1-5)

### Question 1: How does etcd maintain concensus across the control plane controllers, and what happens if a master node crashes in highly available environments?

**Answer:** 
etcd uses the **Raft consensus algorithm** out of the strict nodes configured for cluster quorum.
- **Raft** ensures that changes are committed to the key-value store only when a majority of etcd nodes (n/2 + 1 to avoid split-brain) respond with successful appends.
- If a master goes down, existing pods on worker nodes continue running because they rely only on components logic local to them (`kubelet` & container-runtime). The API server becomes unreachable via that master node structure, load balancers will fail calls over to other masters APIs. As soon as etcd maintains a quorum, K8s is still schedulable and responsive.

---

### Question 2: Can you explain the exact sequence of events from when a user runs `kubectl apply` to create a Deployment, down to the Pod actually running?

**Answer:** 
1. **API Server:** Receives the apply payload, authenticates/authorizes, validates the definition, and stores it in `etcd`.
2. **Deployment Controller:** Watches the API server, notices the new Deployment, and generates a corresponding **ReplicaSet**.
3. **ReplicaSet Controller:** Watches the API server. Upon seeing the new ReplicaSet definition, it determines X Pods are needed. It requests Pod allocations via API interactions.
4. **Scheduler:** Runs a loop watching for 'Unassigned' Pods. It runs its filtering + scoring algorithms to select the best Node, and updates the Pod to signal 'assigned to Node X' to the APIServer.
5. **Kubelet:** On Node X, sees a new Pod assigned to its Node. It issues instructions to the underlying Container Runtime (via CRI over gRPC) to download images, set up networking, and begin containers.

---

### Question 3: How does the Kubernetes Scheduler decide where to place a Pod? What are predicates and priorities?

**Answer:**
It uses a two-phase process:
- **Filtering (Predicates):** Filters out nodes that do not meet the pod's hard constraints (e.g., node is out of CPU/Memory, Taints and Tolerations don't align, Pod anti-affinity isn't met).
- **Scoring (Priorities):** Ranks the remaining valid nodes. It gives higher scores to nodes based on softer rules (e.g., distributing pods to different availability zones, least requested CPU/Mem, image locality).
The node with the highest score is eventually picked.

---

### Question 4: How does a pod get an IP address, and what is CNI?

**Answer:**
A Pod gets an IP via the **Container Network Interface (CNI)** plugin.
- CNI is a standard specification defining how networking plugins interface with the container runtime. Calico, Flannel, and AWS VPC CNI are popular plugin implementations.
- When `kubelet` asks to configure networking, the runtime invokes the CNI plugin on the node.
- The CNI assigns an IP from a predefined IPAM (IP Address Management) block designated for that node and sets up the Linux network namespaces, veth pairs (virtual ethernet bridges connecting the pod namespace to the root node space), and iptables/eBPF routes.

---

### Question 5: What is the difference between a Custom Resource Definition (CRD) and an Operator?

**Answer:**
- **CRD (Custom Resource Definition):** Extends the Kubernetes REST API. It allows you to declare a custom YAML kind (e.g., `kind: PostgresCluster`). By itself, it only tells the APIServer to store and validate the data; it doesn't **do** anything actively.
- **Operator:** Is the "brain" (a custom controller running as a Pod) that actively watches the API for occurrences of that CRD. When a `PostgresCluster` object is created, the Operator takes action to build the underlying stateful components matching that requirement.

---

## 🔹 Security, Storage & Stateful Workloads (Questions 6-10)

### Question 6: How do you achieve true isolation in a multi-tenant Kubernetes cluster?

**Answer:**
Namespaces alone do NOT provide hard isolation. Multi-tenancy implies combining:
1. **RBAC:** Strictly scope ServiceAccounts and users down to minimum necessary permissions within their namespace.
2. **Network Policies:** Implement default-deny network rules at the namespace boundary (Zero trust within the cluster). 
3. **Resource Quotas & LimitRanges:** Prevent neighbor resource starvation (CPU/Mem limits).
4. **Pod Security Admission (or OPA Gatekeeper):** To reject privileged pods, hostPath mountings, or running as root.
5. **Node Isolation (Taints & Tolerations):** To physically bind tenant pods to separated hardware instances if required.
6. Alternatively, use lightweight VMs (like Kata Containers) or vClusters.

---

### Question 7: Explain how StatefulSets differ from Deployments, and where you would use them.

**Answer:**
StatefulSets provide guarantees about the **ordering and uniqueness** of Pods.
- Pods are created sequentially (e.g., web-0, web-1) and deployed/updated in order. 
- You get a stable, predictable network identity (sticky network naming independent of Pod rescheduling).
- Volumes are individually attached to replicas: `volumeClaimTemplates` give each Pod its own dedicated PVC binding safely. Examples: Kafka brokers, Cassandra nodes, Elasticsearch, etc.

---

### Question 8: How do you secure etcd?

**Answer:**
- **Encryption in transit:** Use mTLS (mutual TLS) between API server and etcd, and for etcd peer-to-peer traffic.
- **Encryption at rest:** Enable the `EncryptionConfiguration` on the API server to ensure secrets are stored encrypted inside etcd, preventing direct inspection if etcd's volume is compromised.
- **Access control:** Ensure etcd runs separated to isolated master node networks. Never expose it on user networks. Restrict file permissions of etcd's data directories (`/var/lib/etcd`).

---

### Question 9: Describe the difference between Headless Services and regular Services.

**Answer:**
- **Regular Service (ClusterIP):** Uses `kube-proxy` to distribute traffic identically across ready endpoints. The DNS server returns the *single* virtual ClusterIP, not the Pod IP.
- **Headless Service:** Defined with `clusterIP: None`. It bypasses kube-proxy. The DNS query returns an A-record list directly containing the IPs of the backing Pods. This is strictly required for StatefulSets (database clusters) where a client must connect specifically to *Node-1* vs *Node-2* (e.g., knowing who is Primary and who is Replica).

---

### Question 10: How do you handle secrets rotation and avoid leaking secrets in ConfigMaps?

**Answer:**
- Don't use `ConfigMap` for passwords; use `Secrets`. But native K8s Secrets are just Base64 mapped in etcd.
- Real enterprises use external secret management systems (like **HashiCorp Vault**, AWS Secrets Manager, Azure Key Vault). 
- Using the **External Secrets Operator (ESO)** syncing them securely or mounting dynamically using the **CSI Secrets Store Provider** (no secrets written to etcd).
- Restart pods transparently on Secret changes using tooling like `Reloader` (Stakater) to mount the new versions.

---

## 🔹 Advanced Production Issues & Networking (Questions 11-15)

### Question 11: Application pods consistently go into `ImagePullBackOff` intermittently. What are the potential causes?

**Answer:**
- Reaching Docker Hub / external registry rate limits (too many pulls on public IPs).
- Network routing issues or sporadic DNS resolution failures on the affected worker nodes.
- Missing or malformed `imagePullSecrets` leading to permissions dropping out intermittently.
- The Node's container storage runs out of allocated disk space when uncompressing massive `.tar` images, causing the CRI to reject it.

### Question 12: We have an application that uses websockets or long-polling, how does rolling updates affect these connections?

**Answer:**
During a rolling rollout, K8s sends a `SIGTERM` to your app, but older active connections on a websocket hold open until they complete or are forcefully killed (by `terminationGracePeriodSeconds`).
- Your Pod should catch `SIGTERM` and intentionally initiate a graceful teardown of WebSockets, notifying clients to reconnect cleanly.
- If it doesn't gracefully sever that connection within the default 30 seconds, `SIGKILL` aggressively drops the pod making client network streams fail abruptly without close frames.

### Question 13: Your cluster is facing IP Exhaustion and no new Pods can be scheduled because of networking. What do you do?

**Answer:**
This typically occurs internally within constraints like AWS VPC CNI where available private IPs are depleted by the underlying subnets limit. 
- Adjust the IPAM mapping configuration, and allocate secondary IP CIDRs entirely for the Pod scope to isolate pod routing.
- Enable specific ENI prefix delegations to allow nodes to absorb many smaller blocks compactly.
- If it's pure Flannel/Calico limited range configuration space (e.g., full /16 split to /24 per node), evaluate if you have too many micro-nodes taking IP blocks (and perhaps combine to fewer, bigger nodes).

### Question 14: What is Horizontal Pod Autoscaling (HPA) vs Vertical Pod Autoscaling (VPA), and can you use both together?

**Answer:**
- **HPA:** Scales the *number of Pods* out/in based on metrics like CPU%, Memory threshold, or custom metrics logic.
- **VPA:** Modifies the resource Requests/Limits (CPU & Mem allocations) of the *container dynamically in place* or restarts it scaling it up.
- **Using both:** Do NOT use VPA and HPA simultaneously on the *exact same metrics* (like CPU). They will conflict, causing scaling loops. VPA acts best on background tasks, HPA acts perfectly for stateless web APIs.

### Question 15: If an application has massive scaling demands that the cluster node capacity cannot currently handle, what comes into play?

**Answer:**
**Cluster Autoscaler (CA) or Karpenter.**
When HPA demands more Replica Pods, and the Scheduler detects they cannot fit anywhere due to CPU/Mem `requests` being blocked, the pods enter a `Pending` state.
The Cluster Autoscaler (watching for `Pending` pods failing filters) will issue an API request to the underlying cloud provisioner (AWS ASG/GCP NodePools) to spin up a new EC2/Compute node instance. Once the Node boots and registers (via kubelet), the pending Pod binds to it and runs.
