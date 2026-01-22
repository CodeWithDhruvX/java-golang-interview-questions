## 🔹 Networking Deep Dive (Questions 251-260)

### Question 251: What are endpoints in Kubernetes?

**Answer:**
Endpoints (or EndpointSlices) tracks the IP addresses of the Pods that match a Service's selector.
- When you create a Service, it creates an Endpoints object.
- `kubectl get endpoints my-service`.
- Lists `10.244.1.5:8080`, `10.244.2.3:8080`.

---

### Question 252: How does the EndpointSlice resource differ from Endpoints?

**Answer:**
- **Endpoints:** Single object holding ALL IPs. Big clusters hitting 1000s of pods hit size limit/performance issues.
- **EndpointSlice:** Shards the list into multiple smaller objects (e.g., 100 IPs per slice).
- **Scalability:** Essential for large clusters.

---

### Question 253: What is kube-dns vs CoreDNS?

**Answer:**
- **kube-dns:** Legacy DNS server.
- **CoreDNS:** The current standard (Go-based, highly extensible via plugins).
- Replaced kube-dns in K8s v1.13+. It acts as the cluster name server.

---

### Question 254: What is the role of DNSPolicy?

**Answer:**
Controls how a Pod resolves DNS.
- **ClusterFirst:** (Default) Query Cluster DNS first. If failing, query Node upstream.
- **Default:** Query Node upstream directly (No cluster DNS).
- **None:** Custom config.

---

### Question 255: What is headless service discovery?

**Answer:**
A service with `clusterIP: None`.
- Kube-proxy does not handle it (No Virtual IP).
- DNS triggers A Record lookups returning the set of **Pod IPs** directly.
- Used for StatefulSets (database clusters) needing direct peer-to-peer discovery.

---

### Question 256: What are readiness gates?

**Answer:**
An extension mechanism for Pod readiness.
- Allows external controllers (like LoadBalancer Controller) to inject extra conditions into `status.conditions`.
- The Pod is not considered "Ready" until both the inner readiness probe AND the external gate are true.
- Use Case: Wait for AWS ALB to register target before sending flow.

---

### Question 257: What is ExternalName service?

**Answer:**
Maps a Service to a DNS name, not a selector.
- Returns a CNAME record.
- **Use Case:** Migrate to K8s but keep Database on legacy AWS RDS.
- `my-db-svc` -> `rds.aws.amazon.com`.

---

### Question 258: How do services resolve within a namespace?

**Answer:**
By simple name.
- Pod in `default` calls `http://my-service`.
- DNS resolves `my-service.default.svc.cluster.local`.

---

### Question 259: How do you troubleshoot DNS issues in Kubernetes?

**Answer:**
1.  **Launch Debug Pod:** `kubectl run -it --rm --image=busybox:1.28 dns-test -- sh`.
2.  **Lookup:** `nslookup my-service`.
3.  **Check CoreDNS:** `kubectl get pods -n kube-system -l k8s-app=kube-dns`.
4.  **Check Config:** `/etc/resolv.conf` in pod.

---

### Question 260: What is `dnsPolicy: ClusterFirstWithHostNet`?

**Answer:**
Used for pods running with `hostNetwork: true`.
- Normally `hostNetwork` pods see host DNS configuration.
- This setting forces them to use the **Cluster DNS** (CoreDNS) anyway.

---

## 🔹 Security & Compliance (Questions 261-270)

### Question 261: What is a PodSecurityPolicy (PSP)?

**Answer:**
**Deprecated** in v1.21, Removed in v1.25.
- It was a cluster-level resource controlling security sensitive aspects of pod specification (Root user, Privileged mode, Volume types).

---

### Question 262: What replaced PSPs in newer Kubernetes versions?

**Answer:**
**Pod Security Admission (PSA) Standards.**
- Built-in Admission Controller.
- Configuration via Namespace Labels.
- Levels: `privileged`, `baseline`, `restricted`.
- `kubectl label ns my-ns pod-security.kubernetes.io/enforce=restricted`.

---

### Question 263: What is PodSecurityAdmission?

**Answer:**
The successor to PSP.
- Validates pods against predefined standards.
- Low config: Just verify/label namespaces.
- No custom CRDs needed.

---

### Question 264: How does Open Policy Agent (OPA) work?

**Answer:**
A general-purpose policy engine.
- **Gatekeeper:** The K8s controller for OPA.
- Uses **Rego** language.
- "Deny creation of Service type LoadBalancer in dev namespace."
- Much more flexible than PSA.

---

### Question 265: What is Kyverno?

**Answer:**
A policy engine designed **specifically for Kubernetes**.
- Policies are written in **YAML** (easier than OPA's Rego).
- Can Validate, Mutate, and Generate resources.

---

### Question 266: What is container security context?

**Answer:**
Field in Pod/Container Spec (`securityContext`).
- Control user (`runAsUser`).
- Control capabilities (`capabilities: drop: [ALL]`).
- Control filesystem (`readOnlyRootFilesystem: true`).

---

### Question 267: What is `runAsUser` and `fsGroup`?

**Answer:**
- **runAsUser:** The UID the process runs as (e.g., 1000).
- **fsGroup:** The GID that owns the mounted Volumes. K8s recursively `chowns` the volume to this group so the user can write to it.

---

### Question 268: How do you enforce non-root containers?

**Answer:**
Set `runAsNonRoot: true`.
- If the image is built to run as root (default), the Kubelet refuses to start it.
- Essential for security.

---

### Question 269: How to limit container capabilities?

**Answer:**
```yaml
securityContext:
  capabilities:
    drop:
      - ALL
    add:
      - NET_BIND_SERVICE
```
Least privilege principle.

---

### Question 270: What are seccomp and AppArmor?

**Answer:**
Kernel-level security modules.
- **Seccomp:** Restricts System Calls (e.g., prevent `reboot`).
- **AppArmor:** Restricts file access/network access.
- Enabled via annotations or SecurityContext.

---

## 🔹 Backup & Disaster Recovery (Questions 271-280)

### Question 271: What is Velero?

**Answer:**
The standard tool for Kubernetes Backup & Migration.
- Backs up **YAMLs** (Cluster resources) to Object Storage (S3).
- Backs up **PV Data** via Restic/Kopia or Cloud Snapshots.

---

### Question 272: How do you back up etcd?

**Answer:**
Using `etcdctl snapshot save`.
- This is a full cluster state backup.
- **Note:** It does not backup PV data, only configuration/state.

---

### Question 273: What happens if etcd is corrupted?

**Answer:**
The API Server becomes Read-Only or fails. The cluster is unresponsive.
- Existing pods keep running (Kubelet works in cache mode).
- No new changes possible.
- **Fix:** Restore from snapshot.

---

### Question 274: How do you restore a Kubernetes cluster?

**Answer:**
1.  Restore Etcd (Recovers Control Plane).
2.  If replacing nodes, ensure IPs match or regenerate certs.
3.  Check Controller Manager brings state back.

---

### Question 275: How do you perform point-in-time recovery for etcd?

**Answer:**
Etcd snapshots are periodic. You restore to the *last available snapshot*.
- Etcd does not support "Transaction Log Playback" to specific seconds easily like Postgres PITR.

---

### Question 276: What are etcd snapshots?

**Answer:**
A frozen file containing the key-value store database at a moment in time (db file).

---

### Question 277: How to take a manual etcd backup?

**Answer:**
Exec into etcd pod:
```bash
etcdctl snapshot save /tmp/snapshot.db
```
Copy file out.

---

### Question 278: How do you monitor etcd disk usage?

**Answer:**
Metric: `etcd_mvcc_db_total_size_in_bytes`.
- Etcd has a quota (default 2GB or 8GB).
- If full, it goes into "Alarm" mode and becomes Read-Only.
- Must compact and defrag.

---

### Question 279: What’s the impact of etcd latency?

**Answer:**
High latency (>100ms) causes leader elections and cluster instability.
- Use fast SSDs (NVMe).
- Never use network storage (NFS) for etcd.

---

### Question 280: How do you back up Persistent Volumes?

**Answer:**
1.  **Cloud Snapshots:** `VolumeSnapshot` API triggers EBS Snapshot.
2.  **Velero/Restic:** Filesystem level backup (slower, generic).

---

## 🔹 Cluster Federation & Multi-tenancy (Questions 281-290)

### Question 281: What is Kubernetes federation?

**Answer:**
Coordinating multiple clusters from a single control plane.
- Sync resources (Deployments/Services) across clusters.
- **KubeFed** is the official project (status: alpha/beta).

---

### Question 282: How does kubefed work?

**Answer:**
- **Host Cluster:** Runs control plane.
- **Member Clusters:** Targets.
- You define `FederatedDeployment` on Host.
- KubeFed Controller pushes `Deployment` to Member Clusters.

---

### Question 283: What are the limitations of Kubernetes federation?

**Answer:**
- Complexity.
- Version drift between clusters.
- Network latency.
- Many organizations prefer **GitOps** (ArgCD pushing to multiple clusters) over Federation.

---

### Question 284: What is multi-tenancy in Kubernetes?

**Answer:**
Sharing a cluster between multiple Users/Teams/Customers.
- **Soft:** Namespaces (Shared kernel, risk of neighbor noise).
- **Hard:** Virtual Clusters (vCluster) or Separate Clusters.

---

### Question 285: How do you isolate teams in Kubernetes?

**Answer:**
1.  **Namespaces:** One per team.
2.  **RBAC:** Bind team to their namespace.
3.  **ResourceQuotas:** Limit CPU/RAM per team.
4.  **NetworkPolicies:** Isolate traffic.

---

### Question 286: What are Virtual Clusters?

**Answer:**
Running a simplified K8s control plane **inside a Pod** of the Host K8s cluster.
- The tenant thinks they are root.
- They can create CRDs/Namespaces without affecting Host.

---

### Question 287: What is vCluster and its use case?

**Answer:**
A popular tool for Virtual Clusters.
- **Use Case:** Dev environments. Give every developer their own "Cluster" (which is actually just pods in a namespace). Cheap and fast.

---

### Question 288: What is hierarchical namespace controller (HNC)?

**Answer:**
Allows managing hierarchy of namespaces.
- `Parent NS` -> `Child NS`.
- RBAC and Secrets propagate from Parent to Child.
- Solves the problem of applying same policy to 50 related namespaces.

---

### Question 289: How do you enforce network policies across namespaces?

**Answer:**
Use a **Default Deny** policy at cluster creation or use Admission Controllers (Kyverno) to inject policies into every new namespace automatically.

---

### Question 290: How do you implement resource quotas per namespace?

**Answer:**
Create `ResourceQuota` object.
```yaml
spec:
  hard:
    pods: "10"
    requests.cpu: "4"
```
Prevents one tenant from eating all cluster resources.

---

## 🔹 Real-world Scenarios & Misc (Questions 291-300)

### Question 291: How do you manage Kubernetes in air-gapped environments?

**Answer:**
1.  **Private Registry:** Mirror Docker Hub images to internal Harbor/Artifactory.
2.  **Bundled Install:** Use K3s airgap tarball or Kubeadm with pre-downloaded images.
3.  **Sneakernet:** Physically move data updates.

---

### Question 292: What is OpenEBS and how does it help with local storage?

**Answer:**
Containerized Storage.
- Turns local disks on nodes into a dynamic storage pool.
- Helps run Stateful apps on bare metal without a SAN.

---

### Question 293: What is Portworx?

**Answer:**
Enterprise storage solution for K8s.
- Provides HA, Snapshots, DR, Encryption for volumes.

---

### Question 294: How do you monitor API server performance?

**Answer:**
Prometheus metrics.
- `apiserver_request_duration_seconds`: Latency.
- `apiserver_request_total`: Rate/Errors (4xx, 5xx).

---

### Question 295: What is kube-bench?

**Answer:**
A tool that checks whether Kubernetes is deployed securely by running the checks documented in the **CIS Kubernetes Benchmark**.

---

### Question 296: What is the Kubernetes conformance test suite?

**Answer:**
Sonobuoy.
- Verifies that a cluster is truly "Kubernetes" and compliant with APIs.
- Required for vendor certification (CNCF Certified).

---

### Question 297: What is Cluster API (CAPI)?

**Answer:**
Declarative API for creating clusters.
- `kind: Cluster`.
- `kind: Machine`.
- "Kubernetes managing Kubernetes." Use K8s to spin up EC2 instances and bootstrap new K8s clusters.

---

### Question 298: What is a managed Kubernetes service?

**Answer:**
Provider manages Control Plane (Master).
- You verify Worker Nodes.
- Examples: EKS, AKS, GKE.
- Removes burden of etcd backups/upgrades.

---

### Question 299: What is the role of the CSI driver in Kubernetes?

**Answer:**
Allows third-party storage providers to create plugins without touching core K8s code.
- CSI Driver daemonset runs on nodes to mount drives.

---

### Question 300: How would you secure Kubernetes clusters on a public cloud?

**Answer:**
1.  Private Cluster (Private Endpoint).
2.  Restrict SG/Firewall.
3.  IAM Integration (IRSA on AWS).
4.  Encrypt secrets (KMS).
5.  Audit Logging enabled.
