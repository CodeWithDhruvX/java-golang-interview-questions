# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 2)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Robinhood, etc.
> **Focus deep-dive:** Admission Controllers, Custom API logic, Advanced CNI (eBPF, BGP), CoreDNS tuning, and Control Plane high availability.

---

## 🔹 Advanced API & Admission Control (Questions 16-20)

### Question 16: Explain the difference between Mutating and Validating Admission Webhooks. Give a real-world use case for each.

**Answer:**
Admission webhooks allow you to intercept requests to the Kube-APIServer *after* authorization but *before* the object gets persisted to `etcd`.
- **Mutating Admission Webhook:** Modifies the request payload dynamically. 
  - *Use Case:* Injecting a Sidecar proxy (like Istio Envoy `.yaml`) automatically into any Pod spawned in a namespace marked `istio-injection=enabled`. Or altering an image registry path to pull from a private intra-network mirror securely.
- **Validating Admission Webhook:** Cannot change the payload, but acts as a final Go/No-Go gate.
  - *Use Case:* PodSecurityAdmission. Rejecting any pod that specifies `allowPrivilegeEscalation: true`. Or ensuring every Deployment has specific `cost-center` billing labels.
- *(Note: Mutating runs first, then Validation prevents Mutating from bypassing rules).*

---

### Question 17: How would you design a custom Kubernetes Operator? Explain the control loop architecture.

**Answer:**
An Operator is a Custom Controller combined with a Custom Resource Definition (CRD).
The architecture relies on the **Reconciliation Loop**:
1. Define a struct `MyCustomResource` (e.g., `kind: RedisCluster`) and register it as a CRD.
2. The user executes `kubectl apply` creating desired State.
3. The Operator Pod (written typically in Go compiled via Kubebuilder) sets up **Informer** caches against the API server efficiently (avoiding polling by using WebSockets-based HTTP watches).
4. An event (`Add`/`Update`/`Delete`) queues a **Reconcile(req)** function.
5. `Reconcile()` reads the *desired state* (from CRD) and examines the *current cluster state*. It executes API calls to close the gap (e.g., creating 3 Redis pods).

---

### Question 18: What is API Aggregation in Kubernetes, and how does it differ from a CRD?

**Answer:**
While CRDs are the standard way to add custom objects directly into the existing Kube-API, **API Aggregation** allows you to attach an entirely independent secondary REST API Server alongside the primary Kube-API server safely.
- **CRDs:** The apiserver handles all saving/validation directly into K8s' primary `etcd` store. Simply defines data structures natively.
- **API Aggregation (e.g., `metrics-server` or `custom-apiservers`):** You register a `APIService` object. The K8s front-facing APIServer proxies requested API paths (e.g., `v1beta1.custom.example.com`) back out to your custom POD running its own REST API. That custom API handles how to retrieve/store that data, letting you back it with an entirely different database if needed (like relational SQL).

---

### Question 19: In a massive cluster, the APIServer becomes severely delayed. How do you troubleshoot and tune API performance?

**Answer:**
1. **Metrics Check:** Look at `/metrics` (Prometheus) `apiserver_request_duration_seconds`. Find which APIs/verbs are slow (e.g., LIST on Secrets vs POST on Pods).
2. **Watch Exhaustion:** Too many unoptimized controllers calling expensive `List` without `resourceVersion=0` (which fetches entirely from etcd memory cache instead of hitting disk). Ensure controllers use Informers with caching.
3. **Paginating:** Use `limit=500` and `continue` tokens in heavy scripts so large queries aren’t built completely in-memory.
4. **API Priority & Fairness (APF):** K8s maintains queues. Tune APF objects to ensure system critical nodes (`system-node` user) are prioritized ahead of arbitrary user `kubectl` CLI queries.

---

### Question 20: Explain split-brain in etcd. How does Kubernetes protect against it?

**Answer:**
**Split-brain** occurs when a network partition cuts cluster nodes off from each other, causing two sides to think they are the "leader" and diverging data sets.
- **Protection via Raft:** etcd's `Raft` algorithm *requires* a strict mathematical majority (quorum) to execute a write. In a 5-node etcd cluster, if a partition splits it into (3 Nodes) and (2 Nodes), the 3-node side retains quorum and continues normally. The 2-node side immediately recognizes it lost quorum and transitions to read-only/error, physically preventing two leaders from writing simultaneously.

---

## 🔹 Deep-Dive Networking (CNI & eBPF) (Questions 21-25)

### Question 21: How does Calico use BGP (Border Gateway Protocol) compared to Flannel's VxLAN encapsulation?

**Answer:**
- **Flannel (VxLAN):** Uses overlay networking. It encapsulates every Pod's IP packet inside a larger Node IP packet (UDP). This requires CPU overhead to encrypt/encapsulate at the host level, but works on *any* network.
- **Calico (BGP/Routed):** Can operate entirely unencapsulated (direct routing). It assigns Pod IPs natively and uses BGP to configure the physical routers in the datacenter (or cloud routing tables) to natively know exactly which Node MAC address holds which Pod IPs, resulting in near-bare-metal line-rate throughput capabilities.

---

### Question 22: What is eBPF and why is it replacing standard iptables in modern CNIs like Cilium?

**Answer:**
- **Iptables:** Is a linear set of routing rules parsed iteratively. In a cluster with 5000 services, kube-proxy creates ~20,000+ iptables rules. Every packet transverses these sequentially, causing significant CPU latency (O(n) complexity).
- **eBPF (Extended Berkeley Packet Filter):** Allows running sandboxed custom programs statically inside the Linux kernel itself. Cilium (a popular CNI) compiles network policies into eBPF bytecode using BPF maps (O(1) hash lookups). It completely bypasses the standard Linux network stack routing table, offering massively superior performance, latency, and deep L7 observability metrics silently.

---

### Question 23: How do you identify and mitigate CoreDNS failures under high load?

**Answer:**
If CoreDNS latency spikes, apps fail to resolve external endpoints (`ImagePullErrors` or internal 5xx timeouts).
1. **Mitigation 1 (Replicas):** Use the `dns-autoscaler` to increase CoreDNS pod replicas dynamically based on cluster core counts.
2. **Mitigation 2 (Caching):** Enable `NodeLocal DNSCache`. This runs a tiny DNS agent as a DaemonSet on every node. Pods query their local daemon (via dummy IP) immediately caching requests, bypassing traversing the network entirely to reach CoreDNS for duplicate records.
3. **Troubleshoot:** Look at `coredns_dns_request_duration_seconds` in Prometheus. Check if specific UDP packets are being aggressively bounded/dropped by the CNI's conntrack tables limits (`nf_conntrack_max`).

---

### Question 24: What exactly is `conntrack`, and how does it relate to kube-proxy?

**Answer:**
`conntrack` (Connection Tracking) is a kernel module that tracks active network connections mappings.
Because `kube-proxy` heavily uses Network Address Translation (NAT) so Pod IPs can talk to Service IPs and external IPs, the kernel must remember the translation mapping strictly.
In massively scaled environments (e.g., 500 pods on a node making 10 outbound requests a second), the `conntrack` table fills up completely, resulting in incoming or outgoing packets immediately being dropped with "packet drop" errors, despite CPU and Memory being totally fine. You mitigate by increasing `net.netfilter.nf_conntrack_max` in the node's `sysctl`.

---

### Question 25: How does Kubernetes handle IPv4 / IPv6 dual-stack networking?

**Answer:**
Dual-stack allows Pods and Services to hold **both** an IPv4 and IPv6 address simultaneously.
- You must configure the `kube-apiserver`, `kube-controller-manager`, and `kubelet` with `--service-cluster-ip-range` and `--cluster-cidr` flags containing comma-separated lists of the v4 and v6 blocks.
- Your chosen CNI (like Calico) must technically support managing dual-interfaces. 
- You specify `ipFamilyPolicy: RequireDualStack` in your Service definitions natively. It allows massive native internet-facing IPv6 clusters without relying entirely on NAT64 translation endpoints.
