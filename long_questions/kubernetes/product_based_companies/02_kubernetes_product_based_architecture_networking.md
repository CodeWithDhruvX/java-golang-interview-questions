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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the difference between Mutating and Validating Admission Webhooks. Give a real-world use case for each.
**Your Response:** "Admission webhooks are like security checkpoints that intercept API requests before they're saved. Mutating webhooks can actually change the request - like automatically injecting an Istio sidecar proxy into pods, or changing the image registry path to use an internal mirror. Validating webhooks are the final gatekeepers - they can only approve or reject requests. For example, rejecting pods that try to run as root, or ensuring every deployment has cost-center labels for billing. The key difference is that mutating webhooks modify the request, while validating webhooks just say yes or no. Mutating always runs first, then validating checks the modified request to ensure it still follows the rules."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you design a custom Kubernetes Operator? Explain the control loop architecture.
**Your Response:** "An Operator is basically a smart controller that manages custom resources. The design follows a reconciliation loop pattern. First, I define a CRD to extend the Kubernetes API with my custom resource like RedisCluster. Then I write a controller that watches for changes to these resources using Informers - these are efficient caches that use websockets instead of constant polling. When someone creates or updates a RedisCluster, it triggers a Reconcile function that compares the desired state (what's in the CRD) with the actual state (what's running in the cluster) and makes API calls to fix any differences. It's like having a smart assistant that constantly checks if reality matches your blueprint and fixes any mismatches automatically."

---

### Question 18: What is API Aggregation in Kubernetes, and how does it differ from a CRD?

**Answer:**
While CRDs are the standard way to add custom objects directly into the existing Kube-API, **API Aggregation** allows you to attach an entirely independent secondary REST API Server alongside the primary Kube-API server safely.
- **CRDs:** The apiserver handles all saving/validation directly into K8s' primary `etcd` store. Simply defines data structures natively.
- **API Aggregation (e.g., `metrics-server` or `custom-apiservers`):** You register a `APIService` object. The K8s front-facing APIServer proxies requested API paths (e.g., `v1beta1.custom.example.com`) back out to your custom POD running its own REST API. That custom API handles how to retrieve/store that data, letting you back it with an entirely different database if needed (like relational SQL).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is API Aggregation in Kubernetes, and how does it differ from a CRD?
**Your Response:** "CRDs and API Aggregation both extend Kubernetes but in different ways. CRDs are like adding new tables to the existing database - they store data directly in etcd and use the built-in API server. API Aggregation is like attaching a completely separate API server that handles its own data storage. With CRDs, Kubernetes manages everything. With API Aggregation, I can run my own API server that might store data in PostgreSQL or any other database, and Kubernetes just proxies requests to it. I'd use CRDs for simple custom resources, but API Aggregation when I need custom logic, different storage backends, or want to integrate with existing systems. It's the difference between extending Kubernetes versus connecting external systems to Kubernetes."

---
### Question 19: In a massive cluster, the APIServer becomes severely delayed. How do you troubleshoot and tune API performance?

**Answer:**
1. **Metrics Check:** Look at `/metrics` (Prometheus) `apiserver_request_duration_seconds`. Find which APIs/verbs are slow (e.g., LIST on Secrets vs POST on Pods).
2. **Watch Exhaustion:** Too many unoptimized controllers calling expensive `List` without `resourceVersion=0` (which fetches entirely from etcd memory cache instead of hitting disk). Ensure controllers use Informers with caching.
3. **Paginating:** Use `limit=500` and `continue` tokens in heavy scripts so large queries aren’t built completely in-memory.
4. **API Priority & Fairness (APF):** K8s maintains queues. Tune APF objects to ensure system critical nodes (`system-node` user) are prioritized ahead of arbitrary user `kubectl` CLI queries.

### How to Explain in Interview (Spoken style format)
**Interviewer:** In a massive cluster, the APIServer becomes severely delayed. How do you troubleshoot and tune API performance?
**Your Response:** "I'd start by checking Prometheus metrics to see which APIs are slow - LIST operations on large resources are usually the culprits. The most common issue is watch exhaustion where too many controllers are doing expensive LIST operations instead of using Informers with caching. I'd ensure all controllers use Informers and resourceVersion=0 to hit the etcd cache. For large queries, I'd implement pagination with limit and continue tokens. Finally, I'd tune API Priority & Fairness to prioritize critical system traffic over user queries. It's like optimizing a busy restaurant - identify which orders are taking longest, use better ordering systems, and prioritize VIP customers over casual diners."

### Question 20: Explain split-brain in etcd. How does Kubernetes protect against it?

**Answer:**
**Split-brain** occurs when a network partition cuts cluster nodes off from each other, causing two sides to think they are the "leader" and diverging data sets.
- **Protection via Raft:** etcd's `Raft` algorithm *requires* a strict mathematical majority (quorum) to execute a write. In a 5-node etcd cluster, if a partition splits it into (3 Nodes) and (2 Nodes), the 3-node side retains quorum and continues normally. The 2-node side immediately recognizes it lost quorum and transitions to read-only/error, physically preventing two leaders from writing simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain split-brain in etcd. How does Kubernetes protect against it?
**Your Response:** "Split-brain happens when a network partition separates cluster nodes, and both sides think they're the leader, leading to inconsistent data. Kubernetes prevents this through etcd's Raft consensus algorithm which requires a strict majority for any writes. In a 5-node cluster, if the network splits into 3 and 2 nodes, only the 3-node side has quorum and can continue writing. The 2-node side immediately goes into read-only mode since it lost the majority vote. This mathematical guarantee prevents two leaders from writing simultaneously. It's like a board meeting where decisions only pass with majority vote - if half the members leave, the remaining half can't make decisions until they have majority support."

---

## 🔹 Deep-Dive Networking (CNI & eBPF) (Questions 21-25)

### Question 21: How does Calico use BGP (Border Gateway Protocol) compared to Flannel's VxLAN encapsulation?

**Answer:**
- **Flannel (VxLAN):** Uses overlay networking. It encapsulates every Pod's IP packet inside a larger Node IP packet (UDP). This requires CPU overhead to encrypt/encapsulate at the host level, but works on *any* network.
- **Calico (BGP/Routed):** Can operate entirely unencapsulated (direct routing). It assigns Pod IPs natively and uses BGP to configure the physical routers in the datacenter (or cloud routing tables) to natively know exactly which Node MAC address holds which Pod IPs, resulting in near-bare-metal line-rate throughput capabilities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Calico use BGP (Border Gateway Protocol) compared to Flannel's VxLAN encapsulation?
**Your Response:** "Flannel uses VxLAN overlay networking - it wraps pod packets inside node packets like putting letters in envelopes. This adds CPU overhead but works anywhere. Calico can use BGP for direct routing without encapsulation - it tells the network routers exactly which node has which pod IPs, achieving near-bare-metal performance. It's like the difference between sending mail through a central sorting facility (Flannel) versus having direct delivery routes to every address (Calico). BGP requires more network configuration but gives much better performance, while VxLAN is simpler to set up but has performance overhead. I'd choose Calico with BGP for high-performance environments and Flannel for simplicity."

---

### Question 22: What is eBPF and why is it replacing standard iptables in modern CNIs like Cilium?

**Answer:**
- **Iptables:** Is a linear set of routing rules parsed iteratively. In a cluster with 5000 services, kube-proxy creates ~20,000+ iptables rules. Every packet transverses these sequentially, causing significant CPU latency (O(n) complexity).
- **eBPF (Extended Berkeley Packet Filter):** Allows running sandboxed custom programs statically inside the Linux kernel itself. Cilium (a popular CNI) compiles network policies into eBPF bytecode using BPF maps (O(1) hash lookups). It completely bypasses the standard Linux network stack routing table, offering massively superior performance, latency, and deep L7 observability metrics silently.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is eBPF and why is it replacing standard iptables in modern CNIs like Cilium?
**Your Response:** "Iptables is like a long list of rules that every packet has to check sequentially - with 5000 services, you might have 20,000 rules to go through, causing significant latency. eBPF is revolutionary - it lets me run custom programs directly in the Linux kernel that make routing decisions in O(1) time using hash lookups. Cilium uses eBPF to compile network policies into efficient kernel bytecode, completely bypassing the traditional iptables chain. This gives massive performance improvements and enables deep L7 visibility. It's like the difference between searching through a phone book page by page versus using instant lookup - eBPF is the future of Kubernetes networking."

---

### Question 23: How do you identify and mitigate CoreDNS failures under high load?

**Answer:**
If CoreDNS latency spikes, apps fail to resolve external endpoints (`ImagePullErrors` or internal 5xx timeouts).
1. **Mitigation 1 (Replicas):** Use the `dns-autoscaler` to increase CoreDNS pod replicas dynamically based on cluster core counts.
2. **Mitigation 2 (Caching):** Enable `NodeLocal DNSCache`. This runs a tiny DNS agent as a DaemonSet on every node. Pods query their local daemon (via dummy IP) immediately caching requests, bypassing traversing the network entirely to reach CoreDNS for duplicate records.
3. **Troubleshoot:** Look at `coredns_dns_request_duration_seconds` in Prometheus. Check if specific UDP packets are being aggressively bounded/dropped by the CNI's conntrack tables limits (`nf_conntrack_max`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you identify and mitigate CoreDNS failures under high load?
**Your Response:** "CoreDNS issues show up as DNS resolution failures causing ImagePullErrors or 5xx timeouts. I mitigate this in multiple ways. First, I use the DNS autoscaler to automatically increase CoreDNS replicas based on cluster size. Second, I enable NodeLocal DNSCache which runs a DNS agent on each node as a DaemonSet - this caches DNS queries locally so pods don't have to go across the network for repeated lookups. For troubleshooting, I check CoreDNS metrics in Prometheus and watch for conntrack table limits that might be dropping UDP packets. It's like adding more cashiers and installing local ATMs to reduce wait times at the central bank."

---

### Question 24: What exactly is `conntrack`, and how does it relate to kube-proxy?

**Answer:**
`conntrack` (Connection Tracking) is a kernel module that tracks active network connections mappings.
Because `kube-proxy` heavily uses Network Address Translation (NAT) so Pod IPs can talk to Service IPs and external IPs, the kernel must remember the translation mapping strictly.
In massively scaled environments (e.g., 500 pods on a node making 10 outbound requests a second), the `conntrack` table fills up completely, resulting in incoming or outgoing packets immediately being dropped with "packet drop" errors, despite CPU and Memory being totally fine. You mitigate by increasing `net.netfilter.nf_conntrack_max` in the node's `sysctl`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What exactly is `conntrack`, and how does it relate to kube-proxy?
**Your Response:** "Conntrack is a kernel module that tracks network connections, and it's crucial for kube-proxy's NAT functionality. When pods communicate with services, kube-proxy translates pod IPs to service IPs, and conntrack remembers these translations. In large clusters with thousands of pods making many connections, the conntrack table can fill up completely. When this happens, packets get dropped even though CPU and memory are fine. I mitigate this by increasing the conntrack limit using sysctl. It's like a receptionist who can only remember a limited number of visitors - when too many people come and go, they start turning people away at the door even though there's plenty of space in the building."

---

### Question 25: How does Kubernetes handle IPv4 / IPv6 dual-stack networking?

**Answer:**
Dual-stack allows Pods and Services to hold **both** an IPv4 and IPv6 address simultaneously.
- You must configure the `kube-apiserver`, `kube-controller-manager`, and `kubelet` with `--service-cluster-ip-range` and `--cluster-cidr` flags containing comma-separated lists of the v4 and v6 blocks.
- Your chosen CNI (like Calico) must technically support managing dual-interfaces. 
- You specify `ipFamilyPolicy: RequireDualStack` in your Service definitions natively. It allows massive native internet-facing IPv6 clusters without relying entirely on NAT64 translation endpoints.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Kubernetes handle IPv4 / IPv6 dual-stack networking?
**Your Response:** "Dual-stack networking lets pods and services have both IPv4 and IPv6 addresses simultaneously. I configure the control plane components with both IPv4 and IPv6 CIDR ranges, and ensure my CNI supports dual interfaces. In service definitions, I set ipFamilyPolicy to RequireDualStack. This allows clusters to serve both IPv4 and IPv6 traffic natively, which is essential as the internet transitions to IPv6. It's like giving every house both a traditional street address and a new digital address - mail carriers can use either one. This enables true IPv6-only services while maintaining compatibility with existing IPv4 systems, future-proofing the cluster architecture."
