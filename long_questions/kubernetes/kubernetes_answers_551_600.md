## 🔹 Kubernetes at Scale (Questions 551-560)

### Question 551: What is the maximum number of pods per node?

**Answer:**
Default is **110**.
- Configurable via `--max-pods` in Kubelet.
- Limited by IP addresses in the subnet (if using AWS CNI) and Kubelet CPU/RAM overhead to manage them.

---

### Question 552: What are best practices for scaling to thousands of nodes?

**Answer:**
1.  **Etcd:** Running on dedicated, high-performance NVMe machines.
2.  **API Server:** Horizontal scaling (behind L4 LB).
3.  **Controllers:** Tune `--kube-api-qps` and `--kube-api-burst`.
4.  **Segmentation:** Use multiple smaller clusters if possible (Blast radius).

---

### Question 553: How does Kubernetes schedule high-density workloads?

**Answer:**
- **Bin Packing:** Scoring strategy (`LeastAllocated` vs `MostAllocated`).
- **MostAllocated:** Tries to fill Node A completely before moving to Node B. Saves money (Scaledown ready) but risky (No buffer).

---

### Question 554: How do you isolate workloads in large clusters?

**Answer:**
- **Node Selectors/Taints:** "AI Team gets GPU Nodes, Web Team gets General Nodes".
- **Namespaces:** Logical separation.
- **MicroVMs:** Kata Containers for strong isolation.

---

### Question 555: What is resource bin-packing in Kubernetes?

**Answer:**
The efficient arrangement of pods on nodes to minimize wasted space (fragmentation).
- Scheduler acts as the bin packer.

---

### Question 556: How do you configure taints and tolerations at scale?

**Answer:**
Use **Labels** and **Admission Controllers**.
- "If project=secret, auto-inject toleration for secure-node".
- Don't manage manually per pod.

---

### Question 557: What is topology-aware scheduling?

**Answer:**
Scheduler understands Zones/Regions.
- `topologySpreadConstraints`.
- Ensures High Availability by NOT putting all replicas in `us-east-1a`.

---

### Question 558: How do you scale CRDs and operators in large clusters?

**Answer:**
- **Limit Watches:** Operators should filter namespaces.
- **Cache:** Optimize Informer usage.
- **Sharding:** Run multiple operator instances (each managing subset of namespaces).

---

### Question 559: What is a horizontal pod autoscaler’s impact at scale?

**Answer:**
Heavy load on **Metrics Server** and **API Server**.
- If 10,000 pods scale at once, API server gets hammered with Updates.
- Tune HPA sync period.

---

### Question 560: How do you avoid API server overload?

**Answer:**
- **API Priority and Fairness (APF):** Throttles low-priority requests (e.g., logging agents) to ensure critical requests (Kubelet, Scheduler) get through.
- **Reflector optimization:** Use Protobuf instead of JSON.

---

## 🔹 Security – Advanced Concepts (Questions 561-570)

### Question 561: What is the Kubernetes Pod Security Admission (PSA)?

**Answer:**
Built-in controller (v1.23+) replacing PSP.
- Defines 3 levels: `Privileged`, `Baseline`, `Restricted`.
- Enforce via Namespace Label: `pod-security.kubernetes.io/enforce: restricted`.

---

### Question 562: What are seccomp profiles and how are they used?

**Answer:**
**Secure Computing Mode.**
- System call filter (Kernel level).
- "Profile: RuntimeDefault" blocks dangerous syscalls.
- prevents container breakout.

---

### Question 563: What is AppArmor in Kubernetes?

**Answer:**
Linux Security Module (Ubuntu/Debian).
- Profiles limit file access (`/etc/shadow`) and network capabilities.
- Annotation: `container.apparmor.security.beta.kubernetes.io/my-container: localhost/my-profile`.

---

### Question 564: How do you configure SELinux with Kubernetes?

**Answer:**
Security Enhanced Linux (RHEL/CentOS).
- Uses Labels (Contexts).
- K8s handles mounting volumes with correct SELinux context (`:z` or `:Z`) so container can read them.

---

### Question 565: What are PodSecurityPolicies (PSP) and their alternatives?

**Answer:**
**Deprecated & Removed.**
- Old way to secure pods.
- Replaced by **PSA** (Simple, Native) or **OPA/Kyverno** (Complex, Flexible).

---

### Question 566: What is a container escape and how can it happen?

**Answer:**
Process breaking out of isolation to Host OS.
- **Causes:** Privileged mode, Mounting `/var/run/docker.sock`, Kernel exploit (Dirty COW).

---

### Question 567: How do you enforce read-only root filesystem?

**Answer:**
`securityContext: readOnlyRootFilesystem: true`.
- Attackers cannot modify binaries or install malware (curl/wget).
- App must write temp data to `/tmp` (EmptyDir volume).

---

### Question 568: What is a minimal base image and why use it?

**Answer:**
Distroless or Alpine.
- Contains NO Shell, No Package Manager.
- Reduces attack surface. (Attacker cannot run `yum install` or `sh`).

---

### Question 569: How do you scan container images for vulnerabilities?

**Answer:**
**Trivy, Clair, Anchore.**
- Scan in CI/CD (Build time).
- Scan in Registry (Harbor).
- Scan in Cluster (Admission Controller blocking High CVEs).

---

### Question 570: What is Kube-bench and how does it help with security?

**Answer:**
Go tool that checks if Kubernetes is deployed according to **CIS Benchmark**.
- Checks file permissions (`/etc/kubernetes/pki`), API arguments (`--anonymous-auth=false`), etc.

---

## 🔹 Edge & IoT Integration (Questions 571-580)

### Question 571: How does Kubernetes manage edge workloads with limited resources?

**Answer:**
- **Lightweight Distros:** K3s, K0s.
- **Tuning:** Reduce Eviction Thresholds. Disable unneeded controllers.
- **Scheduling:** PriorityClasses to ensure critical IoT app runs over logging agent.

---

### Question 572: What is the role of MQTT in edge applications with K8s?

**Answer:**
Messaging Protocol (Pub/Sub) for IoT.
- K8s hosts the MQTT Broker (Mosquitto/VerneMQ).
- Sensors publish to Broker. Pods consume from Broker.

---

### Question 573: How can you use K3s on Raspberry Pi devices?

**Answer:**
- ARM64 architecture support.
- Low memory footprint (512MB RAM).
- `curl -sfL https://get.k3s.io | sh -`.

---

### Question 574: What is the role of cloud sync in edge clusters?

**Answer:**
Edge devices are often offline.
- **KubeEdge / Akri:** Syncs Desired State from Cloud -> Edge when online.
- Reports status Edge -> Cloud.

---

### Question 575: What are the challenges in autoscaling edge applications?

**Answer:**
- **Finite Capacity:** You cannot "Add a node" to a Raspberry Pi cluster stuck on a pole.
- **Solution:** Vertical Scaling or Priority Preemption (Kill non-critical tasks).

---

### Question 576: How do you implement offline-first apps with Kubernetes?

**Answer:**
- Dependencies (Images) must be cached locally.
- App logic must queue data locally (SQLite/Redpanda) and sync when network returns.

---

### Question 577: What is the role of WASM in IoT applications?

**Answer:**
Super lightweight.
- Container: 100MB startup.
- WASM: 1MB startup.
- Perfect for constrained resource devices.

---

### Question 578: What is a lightweight ingress solution for edge clusters?

**Answer:**
**Traefik** or **Nginx (Micro).**
- Avoid heavy setups like Istio Gateway.
- Sometimes usage of `HostPort` is acceptable in simple Edge scenarios.

---

### Question 579: How do you implement OTA (Over-the-Air) updates using Kubernetes?

**Answer:**
Simply `kubectl set image`.
- Kubelet on edge pulls new image (differential download).
- Restarts pod.
- Native OTA mechanism.

---

### Question 580: How does 5G/Edge integration affect Kubernetes architecture?

**Answer:**
**MEC (Multi-Access Edge Computing).**
- K8s nodes sit at the Cell Tower.
- Extremely low latency (<5ms).
- Requires intent-based scheduling ("Run on tower closest to User X").

---

## 🔹 WebAssembly (WASM) & Modern Workloads (Questions 581-590)

### Question 581: What is WebAssembly (WASM) in cloud-native applications?

**Answer:**
Binary instruction format. W3C standard.
- "Write once, run anywhere" (Browser, Server, Edge).
- Sandboxed, Near-native performance.

---

### Question 582: How does WASM differ from containers?

**Answer:**
- **Container:** Virtualizes User Space (Linux Userland).
- **WASM:** Virtualizes the Instruction Set (Application Process protection).
- WASM startup is microseconds.

---

### Question 583: What is Krustlet?

**Answer:**
Kubernetes Rust Kubelet.
- A Kubelet implementation that listens to API Server but runs WASM modules instead of Docker containers.

---

### Question 584: How do you run WASM workloads in Kubernetes?

**Answer:**
1.  **Container Shim:** runwasi / containerd-wasm-shim.
2.  Allows you to define `runtimeClassName: wasm` in PodSpec.
3.  Containerd spins up WASM runtime (WasmEdge/Wasmtime) instead of runc.

---

### Question 585: What are the benefits of WASM in K8s?

**Answer:**
- **Speed:** Instant cold starts (Serverless).
- **Security:** Memory safe sandbox.
- **Size:** Tiny binaries.
- **Portability:** Same binary runs on x86 and ARM.

---

### Question 586: What are the limitations of WASM today in production?

**Answer:**
- **Networking/Threading:** WASI (Interface) is still evolving.
- **Ecosystem:** Not all libraries (Python C-extensions) work yet.
- **Debugging:** Tools are immature compared to Docker.

---

### Question 587: How does WASI relate to Kubernetes?

**Answer:**
**WebAssembly System Interface.**
- Defines how WASM talks to OS (Files, Network, Environment).
- Without WASI, WASM can only do math (Calculator). With WASI, it becomes a Docker replacement.

---

### Question 588: What are use cases for WASM in microservices?

**Answer:**
- Serverless Functions.
- Sidecar Filters (Envoy Filters are WASM).
- Data processing at Edge.

---

### Question 589: How do you secure WASM modules?

**Answer:**
- **Signature verification:** Sign the `.wasm` file.
- **Capability model:** Grant explicit access to folders (WASI). It cannot touch anything else by design.

---

### Question 590: What is SpinKube?

**Answer:**
A project to run **Spin** (Fermyon's WASM framework) apps on Kubernetes easily.

---

## 🔹 Miscellaneous Deep-Dive Topics (Questions 591-600)

### Question 591: What is the difference between a Job and a CronJob?

**Answer:**
- **Job:** Run once until completion.
- **CronJob:** Creates Jobs on a time schedule (`*/5 * * * *`).

---

### Question 592: How do you manage third-party software lifecycle in Kubernetes?

**Answer:**
**Operators (OLM - Operator Lifecycle Manager).**
- Provides "App Store" experience.
- Handles Over-the-Air upgrades of the Operator itself.

---

### Question 593: What are Kubernetes Plugins vs Extensions?

**Answer:**
- **Plugins:** Extend core components (scheduler plugins, CNI plugins).
- **Extensions:** Extend API (CRDs, Aggregated APIServers).

---

### Question 594: How do you configure cluster DNS resolution?

**Answer:**
`CoreFILE` (ConfigMap in kube-system).
- Define upstream (8.8.8.8).
- Define rewrites, stub domains (consul.local).

---

### Question 595: What is the difference between external-dns and CoreDNS?

**Answer:**
- **CoreDNS:** Internal Cluster DNS.
- **ExternalDNS:** Controller that talks to AWS Route53 / Google Cloud DNS to create public A Records for Ingresses.

---

### Question 596: How do you implement immutable infrastructure with K8s?

**Answer:**
- NEVER SSH into nodes.
- If node is bad -> Terminate -> Auto Scaling Group replaces it.
- NEVER `kubectl edit` pod. Update Git -> CI/CD deploys new version.

---

### Question 597: What is kube-burner and how is it used?

**Answer:**
Performance testing tool.
- "Create 1000 pods with this template".
- Measures latency, throughput.

---

### Question 598: How do you test high availability in HA Kubernetes clusters?

**Answer:**
**Chaos Testing.**
- Shut down Master 1. API should work.
- Shut down Master 2. API should work (if 3 masters).
- Shut down Etcd leader. Cluster should pause briefly and recover.

---

### Question 599: What is the future of Kubernetes in AI/ML workflows?

**Answer:**
**Batch Scheduling (Volcano/Yunikorn).**
- Gang Scheduling (All or Nothing).
- GPU Slicing.
- Replacing Slurm/HPC schedulers.

---

### Question 600: What is the role of the CNCF in Kubernetes development?

**Answer:**
**Cloud Native Computing Foundation.**
- Holds the IP (Intellectual Property).
- Governance.
- Marketing (KubeCon).
- Certifications (CKA/CKAD).
- Kubernetes is a "Graduated" project.
