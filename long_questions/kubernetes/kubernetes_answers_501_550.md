## 🔹 Kubernetes Internals & Architecture (Questions 501-510)

### Question 501: How does Kubernetes maintain consistency across distributed components?

**Answer:**
Through the **Level Triggered** logic and **Etcd**.
- **Etcd:** The single source of truth (Strong consistency, Raft protocol).
- **Controllers:** They don't just react to "Create Event". They constantly "Reconcile" (loop) to ensure Current State matches Desired State known in Etcd.

---

### Question 502: What is the role of etcd in Kubernetes?

**Answer:**
A distributed key-value store.
- Persists all cluster data (Pods, Secrets, ConfigMaps).
- Uses **RAFT algorithm** for consensus (Leader election).
- If Etcd is down, the API server is down, and the cluster is frozen.

---

### Question 503: How is data stored and secured in etcd?

**Answer:**
- **Structure:** Hierarchical keys (`/registry/pods/default/nginx`).
- **Security:**
  - **Auth:** Client Certs (Mutual TLS).
  - **Encryption at Rest:** Enable encryption provider config to encrypt values (secrets) on disk.

---

### Question 504: What is the informer pattern in Kubernetes controllers?

**Answer:**
A client-go pattern to reduce API Server load.
- Instead of Polling (`LIST` every 5s), it uses a **Watch** connection.
- Maintains a local **Cache** (In-Memory).
- Controller reads from Cache (Lister) and processes events (Add/Update/Delete).

---

### Question 505: How do reconciliation loops work?

**Answer:**
```go
for {
  desired := getFromEtcd()
  current := getFromSystem()
  if current != desired {
    applyChanges()
  }
  sleep()
}
```
It is a continuous feedback loop attempting to drive the system to the desired state.

---

### Question 506: What is optimistic concurrency in K8s resource updates?

**Answer:**
K8s uses **Resource Versions**.
- Client reads Obj (v1).
- Client modifies Obj.
- Client writes Obj.
- IF Obj in Etcd is still (v1) -> Success (v2).
- IF Obj in Etcd is now (v2) -> **Conflict Error (409)**. Client must retry.

---

### Question 507: What is leader election in Kubernetes controllers?

**Answer:**
Ensures only one replica of the Controller Manager (SCHEDULER, etc.) is active.
- Uses a `Lease` object in `kube-system` namespace.
- Active leader updates the `renewTime` every 2s.
- If leader dies, backups see `renewTime` expire and grab the lease.

---

### Question 508: How does Kubernetes handle clock drift across nodes?

**Answer:**
Kubernetes (specifically Etcd and Certificates) is sensitive to clocks.
- **TLS:** Valid NotBefore/NotAfter checks fail.
- **NTP:** You MUST run NTP on all nodes. Kubelet may become "NotReady" if large drift occurs.

---

### Question 509: How does Kubernetes achieve eventual consistency?

**Answer:**
The system doesn't guarantee instant updates.
- If you delete a pod, it takes time to terminate.
- The `Reconcile` loop will keep retrying until it succeeds.
- Eventually, the state converges.

---

### Question 510: What is the function of a finalizer in resource deletion?

**Answer:**
Prevents "Hard Delete".
- If `metadata.finalizers` is present, `delete` only sets `deletionTimestamp`.
- The object stays "Terminating" until a controller handles cleanup (e.g., delete AWS ELB) and removes the finalizer string.

---

## 🔹 CRDs & Operator Pattern (Questions 511-520)

### Question 511: What is a Custom Resource Definition (CRD)?

**Answer:**
Extends the Kubernetes API.
- Allows you to define `Kind: MySQLCluster`.
- K8s stores it, validates it, and serves it via kubectl.
- Does nothing logic-wise without a Controller.

---

### Question 512: How do you create a CRD?

**Answer:**
YAML definition:
```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
spec:
  group: stable.example.com
  names:
    kind: CronTab
```

---

### Question 513: What is an operator in Kubernetes?

**Answer:**
A software pattern: **CRD + Custom Controller**.
- Encodes "Human Operational Knowledge" into code.
- Example: "Prometheus Operator" knows how to upgrade Prometheus versions without losing data.

---

### Question 514: What is the Operator SDK?

**Answer:**
A toolkit to build Operators.
- Supports **Go, Ansible, Helm** based operators.
- Scaffolds code, handles leader election, metrics, and CRD generation.

---

### Question 515: How does an operator differ from a controller?

**Answer:**
- **Controller:** Generic loop (e.g., Deployment Controller).
- **Operator:** A controller specifically built for a *specific application* (e.g., Kafka Operator) using Custom Resources.

---

### Question 516: What are the common use cases for writing an operator?

**Answer:**
- **Databases:** Managing backups, sharding, failovers (Postgres, Cassandra).
- **Monitoring:** Configuring alerts dynamically (Prometheus).
- **Certificates:** Cert-Manager.

---

### Question 517: What is the reconciliation logic in a custom controller?

**Answer:**
The code that executes when a CR changes.
- `func Reconcile(req Request) (Result, Error)`
- Logic: "User wants MySQL 8.0. I see 5.7. Run upgrade script."

---

### Question 518: How do you watch for changes in custom resources?

**Answer:**
Using the `WATCH` API (HTTP Streaming).
- In Go `client-go`: `Informer.AddEventHandler()`.
- Alerts the operator immediately when a user does `kubectl edit my-crd`.

---

### Question 519: What are conversion webhooks in CRD versioning?

**Answer:**
Allow multiple API versions (`v1alpha1`, `v1`).
- When user requests `v1` but etcd has `v1alpha1`, the **Webhook** converts the JSON structure on the fly so the client gets what they asked for.

---

### Question 520: How do you handle backward compatibility in CRDs?

**Answer:**
- Use `deprecated: true` for fields.
- Use Conversion Webhooks.
- Support multiple versions in `versions[]` list in CRD definition.

---

## 🔹 Multi-Cluster & Federation (Questions 521-530)

### Question 521: What is Kubernetes Federation (KubeFed)?

**Answer:**
A project to coordinate configuration across multiple clusters.
- Define a `FederatedReplicaSet`.
- KubeFed creates `ReplicaSet` in Cluster A, Cluster B...

---

### Question 522: What are the use cases for multi-cluster deployment?

**Answer:**
- **Isolation:** Prod vs Dev clusters.
- **Latency:** US-East and EU-West clusters (closer to users).
- **Scale:** Avoiding the 5000 node limit.
- **Disaster Recovery:** Failover if one region burns.

---

### Question 523: How do you share secrets across clusters?

**Answer:**
Federation doesn't do this securely out of the box.
- Use **External Secrets** (Vault) connected to both clusters.
- Or use **GitOps** with SealedSecrets deployed to both.

---

### Question 524: What is a service mesh in a multi-cluster setup?

**Answer:**
**Istio Multi-Cluster.**
- Services in Cluster A can talk to Cluster B via mTLS gateways.
- Looks like local traffic to the application.
- `servicename.namespace.global`.

---

### Question 525: What are the challenges of multi-cluster observability?

**Answer:**
- Aggregating metrics from 10 clusters.
- **Solution:** Thanos or Cortex (Federated Prometheus).
- Tracing requests that jump between clusters (Need unified Trace ID propagation).

---

### Question 526: How do you manage DNS across multiple clusters?

**Answer:**
**ExternalDNS** or **Multi-Cluster Services (MCS)** API.
- Create a global DNS record `app.global.example.com` balancing traffic to LoadBalancers in both clusters.

---

### Question 527: How does ArgoCD support multi-cluster deployment?

**Answer:**
ArgoCD runs in a "Management Cluster".
- It has `kubeconfig` credentials for "Target Clusters".
- It pushes resources to targets.
- Single pane of glass for all apps everywhere.

---

### Question 528: What is a cluster registry?

**Answer:**
A database/CRD listing available clusters.
- `kind: Cluster`.
- `endpoint: https://1.2.3.4`.
- Used by tooling (like KubeFed, Argo) to discover targets.

---

### Question 529: How do you enforce policies across clusters?

**Answer:**
**Policy Engines (Kyverno / OPA / ACM).**
- Deploy the logical policy to a Git repo.
- GitOps agents on ALL clusters sync the policy.
- Ensures `Privileged: false` is enforced globally.

---

### Question 530: What is the difference between geo-replication and multi-cluster?

**Answer:**
- **Geo-Replication:** Application layer feature (Database syncing data US->EU).
- **Multi-Cluster:** Infrastructure layer feature (Running K8s nodes in US and EU).

---

## 🔹 Advanced Networking Plugins (Questions 531-540)

### Question 531: What is CNI in Kubernetes?

**Answer:**
**Container Network Interface.**
- A standard spec for network plugins.
- Kubelet calls CNI binary (`/opt/cni/bin/bridge`) to setup pod network namespace (IP, Route) when pod starts.

---

### Question 532: How does Calico differ from Flannel?

**Answer:**
- **Flannel:** Simple Overlay (VXLAN). Layer 2. Easy. No NetworkPolicies.
- **Calico:** Layer 3 (BGP). High performance. Supports NetworkPolicies (Security).

---

### Question 533: What is the role of Canal in networking?

**Answer:**
A hybrid.
- Uses **Flannel** for easy networking (VXLAN).
- Uses **Calico** for Policy enforcement.
- Best of both worlds for some users.

---

### Question 534: What is Cilium and why is it popular?

**Answer:**
CNI based on **eBPF**.
- No iptables (super fast).
- Deep observability (L7 visibility without sidecars).
- High scale.

---

### Question 535: How does eBPF enhance Kubernetes networking?

**Answer:**
Runs sandboxed programs in the Linux Kernel.
- Allows packet filtering/routing without context switching to userspace or traversing long iptables chains.
- Used by Cilium.

---

### Question 536: What is a NetworkPolicy and how is it enforced?

**Answer:**
Firewall rules for Pods.
- `allow from app: frontend`.
- **Enforced by CNI Plugin** (Calico/Cilium).
- If you use Flannel (without plugin), NetworkPolicies are ignored!

---

### Question 537: What are ingress and egress policies?

**Answer:**
- **Ingress:** Incoming traffic (Who can call me?).
- **Egress:** Outgoing traffic (Who can I call?).
- Default is Allow-All.

---

### Question 538: What is a network plugin vs network proxy?

**Answer:**
- **Plugin (CNI):** Sets up IP/Wires (Calico). Run once per pod start.
- **Proxy (Kube-Proxy):** Sets up Service Balancing (IPTables). constantly syncing.

---

### Question 539: What is kube-router and when would you use it?

**Answer:**
A lean alternative to kube-proxy + CNI.
- Uses LVS/IPVS for services.
- Uses BGP for pod networking.

---

### Question 540: How does IPAM (IP Address Management) work in K8s?

**Answer:**
CNI plugin responsibility.
- **Host-Local:** Assigns range `10.1.1.0/24` to Node A. Node A assigns `.1`, `.2` locally.
- **DHCP:** Asks external server.

---

## 🔹 Ingress Controllers & Service Mesh (Questions 541-550)

### Question 541: What is the difference between an ingress and ingress controller?

**Answer:**
- **Ingress:** The Rule (YAML). "Route /foo to Service A".
- **Controller:** The Worker (Nginx/Traefik). Reads rule, updates config, handles traffic.

---

### Question 542: What is the NGINX ingress controller and how does it work?

**Answer:**
Official/Community controller using Nginx.
- Watches Ingress resources.
- Generates `nginx.conf`.
- Reloads Nginx.
- Provides LoadBalancing, SSL, Auth.

---

### Question 543: What is Istio and how does it relate to Kubernetes?

**Answer:**
A Service Mesh.
- Adds a logic layer *above* the network.
- **Sidecar (Envoy):** Intercepts all traffic.
- Features: mTLS, Retries, Circuit Breaking, Canary Splitting, Tracing.

---

### Question 544: What is Linkerd and how is it different from Istio?

**Answer:**
Another Service Mesh.
- **Focus:** Simplicity, Lightness, Speed (Rust-based proxy).
- **Difference:** Istio is feature-rich but complex/heavy. Linkerd is "Just works".

---

### Question 545: How does Envoy proxy work in a service mesh?

**Answer:**
- Deployed as sidecar in *every* pod.
- Application talks to localhost port.
- Envoy handles mTLS, discovery, routing to destination Envoy.

---

### Question 546: How do mutual TLS (mTLS) connections work in Istio?

**Answer:**
1.  Citadel (Istio Auth) issues certs to every pod.
2.  Envoy Proxy A talks to Envoy Proxy B.
3.  They handshake and verify identity (SPIFFE ID).
4.  Traffic encrypted. Transparent to App.

---

### Question 547: How do sidecars help with observability and security?

**Answer:**
- **Observability:** Envoy logs every request (Latency, 200/500 codes) -> Prometheus.
- **Security:** Enforces mTLS and AuthorizationPolicy (Service A can call B only on GET /api).

---

### Question 548: What is a virtual service in Istio?

**Answer:**
Custom resource to define routing rules.
- "If header `user=beta`, route to `subset: v2`".
- More powerful than standard K8s Ingress.

---

### Question 549: How do you implement traffic shifting with a service mesh?

**Answer:**
Define weights in VirtualService.
```yaml
route:
- destination: app-v1
  weight: 90
- destination: app-v2
  weight: 10
```

---

### Question 550: How do you manage circuit breaking in Istio?

**Answer:**
Configure `DestinationRule`.
- `connectionPool`: Max parallel connections.
- `outlierDetection`: If host returns 500s 3 times, eject it from pool for 1 minute.
