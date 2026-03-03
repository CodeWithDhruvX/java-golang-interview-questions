# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 4)
> **Target:** Google, Amazon, Microsoft, Uber, Netflix, Robinhood, etc.
> **Focus deep-dive:** Platform Security, Multi-tenancy Isolation, OPA/Gatekeeper, mTLS Service Mesh, and Secrets Architecture.

---

## 🔹 Advanced Cluster Security & Multi-tenancy (Questions 36-40)

### Question 36: What is a "Hard Multi-tenancy" environment and why doesn't native Kubernetes provide it out-of-the-box?

**Answer:**
**Hard Multi-tenancy** means hostile tenants (e.g., competing customers) running code on the exact same cluster cannot break out of their pods or affect each other via kernel exploits or resource starvation.
- **Native K8s Flaw:** Containers share the same underlying Linux Kernel natively. Even with Namespaces and network policies, a zero-day kernel exploit could allow a tenant to gain host-node root.
- **Solution:** Companies achieve hard multi-tenancy by using **vClusters** (isolated control planes), or more critically, employing MicroVM Sandboxing runtimes (like **Kata Containers** or AWS **Firecracker** natively within K8s arrays) which wrap every pod in heavily restricted hypervisor boundaries instead of standard namespace slicing.

---

### Question 37: Explain the architecture of OPA Gatekeeper vs native Pod Security Admission (PSA).

**Answer:**
- **Pod Security Admission (PSA):** A native, built-in standard admission controller. It's statically configured using 3 profiles (`Privileged`, `Baseline`, `Restricted`). It's fast, but extremely inflexible. You cannot say "Allow hostNetwork *only* if the image comes from internal-registry".
- **OPA Gatekeeper:** A dynamic validating admissions webhook. It uses the `Rego` language to evaluate custom policies exactly. It intercepts the API payload tree and runs an infinite variety of logic: ensuring images aren't on blocklists, enforcing exact label/billing presence, or verifying ingress objects do not attempt malicious regex overlap paths across namespaces.

---

### Question 38: Are Network Policies capable of deeply inspecting HTTPS traffic to block specific HTTP headers or SQL queries?

**Answer:**
No. Native Kubernetes **NetworkPolicies** strictly operate at **Layer 3 / Layer 4** (IP and Port ranges). They act as a basic packet firewall.
If you need **Layer 7** inspection (blocking specific HTTP endpoints, routing based on headers, evaluating gRPC frames natively), you must deploy a **Service Mesh** (like Istio/Envoy or Linkerd), or leverage advanced L7 capabilities inside an eBPF-based CNI like Cilium.

---

### Question 39: We need to implement mTLS between all microservices within the cluster. How does a Service Mesh like Istio achieve this?

**Answer:**
1. The **Istio Control Plane (istiod)** acts internally as a Certificate Authority (CA).
2. Through Mutating Webhooks, a proxy container (Envoy) is injected as a **sidecar** into every application Pod in the mesh.
3. Both sidecars continuously rotate and fetch temporary TLS certificates natively mapped to their Pod's `ServiceAccount` identity.
4. When `App-A` talks to `App-B`, the connections do not go app-to-app. The traffic goes `App-A -> ProxyA -> mTLS Tunnel -> ProxyB -> App-B` entirely natively, enforcing military-grade encryption in-transit seamlessly to the application itself!

---

### Question 40: Assume you've given a Pod a Service Account token. How would you prevent the Pod from leveraging that token off-cluster?

**Answer:**
Historically, tokens were static `Secrets` that simply functioned until manually revoked. If an attacker stole it (via an app-level SSRF vulnerability), they could curl the APIServer from their external laptop.
- Modern K8s defaults to **Bound Service Account Tokens**. 
- These are temporary tokens projected dynamically by kubelet into the Pod's volume. They have expiration loops (e.g., 1 hour natively), and are strictly bound to the specific Pod UID. If the exact Pod dies, the token is instantly invalid. 
- You can further secure the APIServer by configuring `Anonymous-Auth=false` and restricting the API Endpoint to internal VPC/VPN CIDR limits exclusively.

---

## 🔹 Operational Scenarios & Best Practices (Questions 41-45)

### Question 41: How would you safely design the cluster architecture for a globally distributed stateless application (spanning US, EU, and Asia)?

**Answer:**
Avoid massively stretched, single cross-continental K8s clusters spanning 200 ms ping times because the `etcd` raft quorum writes would massively choke and fail.
Instead, operate **federated distinct clusters**, one in each geographical region natively:
1. Deploy 3 isolated K8s clusters natively.
2. Rely on a global ANYCAST DNS/Loadbalancer routing system (like AWS Route53 Geo-Routing or Cloudflare) at the edge.
3. Manage them all concurrently via a centralized GitOps engine (ArgoCD) pointing to a single source repo configuring multi-cluster helm rollouts.

---

### Question 42: What is the most reliable way to back up a Kubernetes cluster natively?

**Answer:**
Do not simply attempt to scrape the ETCD `db` file or run `kubectl get all` scripts (which misses core API boundaries, secrets or CRD state).
Use **Velero**.
It is a highly specialized K8s backup agent that cleanly pulls the exact API states seamlessly (backing them into S3 native Object storage) and additionally has hooks into the cloud provider's storage APIs (EC2/GCP) to trigger synchronous volume snapshots natively mapped to those API objects for total disaster recovery.

---

### Question 43: Describe an architecture pattern for securely pulling sensitive images from external registries without putting image secrets into every namespace.

**Answer:**
Instead of copying a specific `kubernetes.io/dockerconfigjson` Secret into exactly every namespace and configuring every pod with `imagePullSecrets` iteratively:
- **Solution 1:** Attach the IAM Role directly to the worker nodes (NodeInstanceRole). The native `Kubelet` on AWS/GCP can securely natively impersonate the instance hardware to authenticate directly to ECR/GCR transparently without K8s secrets!
- **Solution 2 (Operator):** Use a mutating webhook like Kyverno to automatically identify pods launching and seamlessly inject the `imagePullSecret` reference dynamically.

---

### Question 44: What are the risks of using `latest` tag in production, explicitly in Kubernetes?

**Answer:**
1. **Unpredictable Rollouts:** `kubectl apply` doesn't know the image actually changed inside the repository register if the tag still reads `latest`. Thus, it literally won’t trigger a RollingUpdate native cycle.
2. **Node Caching Failure:** If 2 worker nodes already pulled an old `latest`, and a 3rd node boots and pulls a newly pushed `latest`, your application is now actively running two completely different software versions concurrently causing intermittent obscure bugs.
3. **Improper Rollbacks:** If `latest` is broken, what exact previous SHA do you revert to natively? Always use explicit, immutable SHA or semantic versioning tags natively!!

---

### Question 45: A developer wants to run `sysctl -w net.ipv4.ip_forward=1` inside a container. Is this possible?

**Answer:**
By default, no. K8s prevents overriding sensitive Linux kernel parameters that affect the entire Node's network stack.
If explicitly required:
1. The container requires `securityContext.privileged: true`.
2. The Pod must specifically whitelist the `sysctl` in its payload.
3. The Kubelet itself on that worker node must be specifically booted with `--allowed-unsafe-sysctls` passing that exact parameter, which is highly scrutinized natively!
