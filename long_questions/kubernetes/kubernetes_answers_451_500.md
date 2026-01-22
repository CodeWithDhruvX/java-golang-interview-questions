## 🔹 Authentication & Authorization (Questions 451-460)

### Question 451: How does Kubernetes handle authentication?

**Answer:**
Kubernetes does not manage Users. It assumes a User is authenticated by an outside source.
- **Strategies:** X509 Client Certs, Static Token File, Bootstrap Tokens, Service Account Tokens, OpenID Connect, Webhook Token Auth.
- API Server checks the `Authorization` header.

---

### Question 452: What are authentication plugins?

**Answer:**
Modules configured in the API Server implementation to handle auth.
- `--authentication-token-webhook-config-file`: Validation of bearer tokens.
- OIDC Plugin: Logic to validate JWTs from Google/Azure.

---

### Question 453: What is service account token projection?

**Answer:**
A security improvement (`projected` volume).
- Instead of a static secret, the kubelet requests a token from TokenRequest API with specific **Audience** and **Duration**.
- Rotates automatically. If pod dies, token expires quickly.

---

### Question 454: How do you authenticate with OIDC?

**Answer:**
1.  **Configure IdP:** Create Client in Okta.
2.  **Configure API Server:** `oidc-issuer-url`, `oidc-client-id`.
3.  **Config Kubectl:** Use `kubelogin` plugin to fetch token from Okta and send it as `Bearer` token to K8s.

---

### Question 455: What is impersonation in Kubernetes?

**Answer:**
Allows a user (typically admin) to act as another user.
- Headers: `Impersonate-User: jane`.
- API Server checks if original user has `impersonate` verb on `users` resource.
- Helpful avoiding re-logging in.

---

### Question 456: What is ABAC and is it still supported?

**Answer:**
**Attribute-Based Access Control.**
- Policy defined in a JSON file on the master node.
- **Status:** Legacy. Hard to manage (requires API server restart to update policies).
- RBAC is the standard replacement.

---

### Question 457: How can you limit a user to read-only access?

**Answer:**
Create a **ClusterRole** with `verbs: ["get", "list", "watch"]`.
Bind it to the user.
```yaml
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["list", "get"]
```

---

### Question 458: How do you audit RBAC permissions?

**Answer:**
**Tools:**
- `kubectl auth can-i create pods --as jane`.
- **Krakk8s / Rakess:** Visualizes access matrix.
- **Kubiscan:** Scans for risky permissions (e.g., who has `secrets` access?).

---

### Question 459: What is RBAC escalation and how to prevent it?

**Answer:**
A user creating a Role with permissions **higher** than they possess.
- Kubernetes API blocks this.
- You cannot create a Role binding granting Admin rights if you are not Admin.

---

### Question 460: What are roles vs clusterroles?

**Answer:**
- **Role:** Namespaced (Can read Pods in Dev).
- **ClusterRole:** Global (Can read Nodes / PersistentVolumes).
- **Binding:** You can bind a ClusterRole (View) to a RoleBinding (in Dev namespace) to give "View access only in Dev".

---

## 🔹 Performance & Optimization (Questions 461-470)

### Question 461: How do you profile CPU and memory usage of pods?

**Answer:**
- **Metrics:** Prometheus (Usage vs Request).
- **Profiling:** `pprof` (if Go app exposes it).
- **Continuous Profiling:** Tools like **Parca** or **Pyroscope**.

---

### Question 462: What are best practices for container resource limits?

**Answer:**
- **Requests:** Set to expected usage (Guarantees capacity).
- **Limits:** Set higher (to allow burst), but not infinite (to protect node).
- **Memory Limit:** Critical (Prevents Node OOM).
- **CPU Limit:** Controversial (Can cause throttling latency). Some disable it for latency-sensitive apps.

---

### Question 463: What happens if you don’t set resource limits?

**Answer:**
The Pod is **BestEffort** QoS.
- It can consume all CPU/RAM on the node.
- If node is full, it is the **first** to be evicted.
- "Noisy Neighbor" risk.

---

### Question 464: How does the scheduler choose nodes based on resources?

**Answer:**
Checks **Allocatable** resources.
- `Node Allocatable = Node Capacity - Kube Reserved - System Reserved`.
- If Pod Request < Available Allocatable => Schedule.

---

### Question 465: What is the eviction threshold for memory pressure?

**Answer:**
Kubelet settings: `memory.available<100Mi`.
- If Node Free Memory drops below this, Kubelet starts killing pods to reclaim memory.

---

### Question 466: How do you minimize pod startup time?

**Answer:**
1.  **Image Size:** Use Alpine/Distroless.
2.  **Pull Policy:** `IfNotPresent`.
3.  **Probes:** Optimize `initialDelaySeconds`.
4.  **InitContainers:** Optimize their work.

---

### Question 467: What are tools to benchmark Kubernetes clusters?

**Answer:**
- **Kube-burner:** Cluster density/stress test.
- **Sonobuoy:** Compliance.
- **Netperf:** Network throughput.

---

### Question 468: How do you reduce image pull latency?

**Answer:**
- **Pre-pulling:** DaemonSet that pulls common images to all nodes.
- **Image Locality:** Schedule pod on node that already has image.
- **Registry Mirror:** Run a mirror inside the VPC.

---

### Question 469: What is `imagePullPolicy: IfNotPresent` vs `Always`?

**Answer:**
- **IfNotPresent:** Use local cache if available. (Fast).
- **Always:** Contact Registry to check digest. (Slower, but ensures freshness).
- Note: `:latest` tag implies `Always` by default.

---

### Question 470: How does the pod preemption mechanism work?

**Answer:**
If High Priority pod is Pending and no nodes are free:
- Scheduler Evicts Lower Priority pods from a node.
- Once resources freed, High Priority pod schedules there.
- Uses `PriorityClass`.

---

## 🔹 Upgrades & Versioning (Questions 471-480)

### Question 471: How do you upgrade a Kubernetes cluster safely?

**Answer:**
**Rolling Upgrade:**
1.  New Master (Control Plane).
2.  Cordon Node A.
3.  Drain Node A.
4.  Upgrade Kubelet/OS on Node A.
5.  Uncordon Node A.
6.  Repeat.

---

### Question 472: What is the upgrade path for minor/patch versions?

**Answer:**
- K8s supports **skew** of n-2 (mostly).
- Recommended: Upgrade one minor version at a time (1.24 -> 1.25 -> 1.26).
- **Do not skip versions** (1.24 -> 1.26 is dangerous).

---

### Question 473: What is the role of kubeadm in cluster upgrades?

**Answer:**
- `kubeadm upgrade plan`: Checks available versions.
- `kubeadm upgrade apply`: Upgrades static pods (API, Controller, Scheduler) manifests and certificates.

---

### Question 474: How do you perform a zero-downtime upgrade?

**Answer:**
Requires HA Control Plane.
Requires **PodDisruptionBudgets** (PDB) on workloads so users aren't impacted when nodes drain.
Requires Application with >1 Replica.

---

### Question 475: How do you upgrade CRDs?

**Answer:**
Carefully.
- Helm does NOT upgrade CRDs.
- Must manually `kubectl apply -f crds/`.
- Ensure new CRD schema is backward compatible with existing data in etcd.

---

### Question 476: How do you check for deprecated APIs before upgrading?

**Answer:**
- **Kubent (Kube No Trouble):** Scans cluster for objects using deprecated APIs (e.g., `batch/v1beta1 CronJob`).
- **Pluto:** Scans Helm releases.

---

### Question 477: What’s the role of feature gates in Kubernetes upgrades?

**Answer:**
Flags to enable/disable Alpha/Beta features.
- `--feature-gates="SomeFeature=true"`.
- As K8s upgrades, some gates become default (GA) and are removed.

---

### Question 478: How do you verify cluster health after an upgrade?

**Answer:**
1.  `kubectl get componentstatuses` (Deprecated but useful).
2.  `kubectl get nodes` (All Ready?).
3.  Run a "Smoke Test" workload.
4.  Check metrics for error spikes.

---

### Question 479: What is the impact of upgrading etcd?

**Answer:**
API Server connects to etcd.
- Usually etcd is upgraded via `kubeadm` or provider content.
- Must ensure data format compatibility.
- Backup beforehand is mandatory.

---

### Question 480: What’s the difference between kubelet and control plane upgrades?

**Answer:**
- Control Plane must be >= Kubelet.
- You upgrade CP first.
- Nodes (Kubelet) can trail behind (up to 2 minor versions).

---

## 🔹 Production Readiness (Questions 481-490)

### Question 481: What are readiness criteria for going to production in Kubernetes?

**Answer:**
Checklist:
- HA Masters.
- RBAC locked down.
- Monitoring/Alerting active.
- Backups tested.
- Resource Limits set.
- PDBs configured.
- Liveness/Readiness probes set.

---

### Question 482: How do you enforce policies for resource limits?

**Answer:**
**LimitRange.**
- Create `LimitRange` in namespace.
- If user creates pod without limits, it automatically injects valid default limits.

---

### Question 483: What is a pod disruption budget?

**Answer:**
A guardrail for voluntary disruptions (Maintenance).
- `minAvailable: 1`: "Maintenance process (Drain) generally allowed, BUT you must ensure 1 is always up."

---

### Question 484: How do you manage blue/green deployment in production?

**Answer:**
Usually via **Ingress Controller** or **Service Mesh** traffic splitting.
- Service level switching (label selector) breaks connection draining. Mesh is smoother.

---

### Question 485: How do you ensure HA of critical workloads?

**Answer:**
1.  `replicas: >1`.
2.  `topologySpreadConstraints`: Spread across Zones (AZs).
3.  `podAntiAffinity`: Don't stack on same node.

---

### Question 486: What is the impact of control plane downtime?

**Answer:**
- **Running Pods:** Continue running fine.
- **Networking:** Kube-proxy rules remain.
- **Impact:** Cannot schedule new pods, cannot scale, cannot update deployments. API read-only or down.

---

### Question 487: What’s the difference between availability and resilience?

**Answer:**
- **Availability:** Uptime % (Is it working now?).
- **Resilience:** Ability to recover from failure (If it breaks, does it self-heal?).
- Autoscaling/Restart improves resilience -> leads to Availability.

---

### Question 488: How do you manage production secrets?

**Answer:**
Use an External Secret Store (Vault/AWS Secrets Manager) synced via ESO. Do not rely on base64 encoded yaml in git.

---

### Question 489: What are common production anti-patterns in Kubernetes?

**Answer:**
- Huge clusters (Blast radius).
- Shared namespace for everything.
- Ops team manually applying YAMLs (ClickOps).
- Ignoring resource requests.

---

### Question 490: How do you scale Kubernetes clusters in production?

**Answer:**
- **Cluster Autoscaler:** Adds nodes.
- **Node Pools:** Segment workloads (High Mem pool, GPU pool).
- **Federation:** Add more clusters.

---

## 🔹 Edge, Hybrid, & Specialized Use Cases (Questions 491-500)

### Question 491: What is K3s and where is it used?

**Answer:**
Lightweight Kubernetes (Rancher).
- Single binary (<100MB).
- Removes cloud providers, uses SQLite instead of etcd (optional).
- **Use Case:** IoT, Edge, Dev.

---

### Question 492: What is KubeEdge?

**Answer:**
Framework to extend K8s to Edge.
- **CloudCore:** Running in DC.
- **EdgeCore:** Running on device.
- Handles offline syncing (device disconnects, keeps running, syncs later).

---

### Question 493: What are microK8s?

**Answer:**
Canonical (Ubuntu) K8s distribution.
- Snap install.
- Zero-ops.
- Good for local dev or single-node edge.

---

### Question 494: How does Kubernetes support edge computing?

**Answer:**
By providing a standard API to manage distributed apps.
- Problem: Latency, Bandwidth.
- Solution: GitOps pulls config to edge. Apps run locally processing data, sending only summary to cloud.

---

### Question 495: How can Kubernetes be used for IoT workloads?

**Answer:**
Deploy monitoring agents or data processors as Pods on IoT gateways.
- Update firmware/app logic via `kubectl set image`.

---

### Question 496: How do you handle intermittent connectivity in edge K8s?

**Answer:**
- **Local Images:** Don't rely on `Always` pull.
- **Independent Autonomy:** Kubelet continues running existing pods even if API Server (Cloud) is unreachable.

---

### Question 497: What is the difference between bare-metal and cloud K8s?

**Answer:**
- **LB:** No ELB on bare-metal (Use MetalLB).
- **Storage:** No EBS (Use Longhorn/Ceph/OpenEBS).
- **Network:** Need to manage physical router/BGP.

---

### Question 498: How does multi-cloud Kubernetes deployment work?

**Answer:**
Usually via a Management Plane (Rancher/Anthos/Arc).
- A central pane of glass manages GKE cluster and EKS cluster.
- App deployment targets both.

---

### Question 499: What is Anthos and how does it relate to Kubernetes?

**Answer:**
Google Cloud's hybrid platform.
- Runs GKE on-prem (VMware) or on AWS/Azure.
- Unified management.

---

### Question 500: What’s the role of WASM (WebAssembly) in Kubernetes?

**Answer:**
Next-gen runtime.
- **WASM modules** are lighter/faster than containers.
- **Kruslet:** A Kubelet implementation that runs WASM instead of Containers.
- Better security/sandboxing for edge.
