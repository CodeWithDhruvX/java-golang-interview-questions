# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 6)
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, IBM, Capgemini, etc.
> **Focus:** Multi-Container Pod Patterns, Startup Probes, Pod Priority & Preemption, Namespace Strategies, Cluster Upgrades, and Image Security.

---

## 🔹 Multi-Container Pod Design Patterns (Questions 66-70)

### Question 66: What are the three well-known multi-container pod patterns? Give a real-world use case for each.

**Answer:**
Multi-container pods share the same network namespace (`localhost`) and can share volumes.

**1. Sidecar Pattern**
A helper container augments the main container's functionality without changing it.
- **Example:** Main container is a `Node.js` web app. A `Fluentd` sidecar reads log files from a shared volume and forwards them to Elasticsearch.
```
[App Container] --writes logs--> [Shared Volume] <--reads-- [Fluentd Sidecar] --> Elasticsearch
```

**2. Ambassador Pattern**
The ambassador container acts as a proxy, translating or abstracting network communication on behalf of the main container.
- **Example:** Main app always connects to `localhost:5432`. The ambassador container proxies that connection to the correct environment's database (Dev → dev-db, Prod → prod-db), so the app code never needs to know the actual database address.
```
[App Container] --> localhost:5432 --> [Ambassador Container] --> Correct DB
```

**3. Adapter Pattern**
The adapter container transforms the main container's output into a standardized format.
- **Example:** A legacy app exposes metrics in a proprietary format. The adapter container reads those metrics and transforms them into Prometheus format on `/metrics`, making the pod visible to the Prometheus scraper without touching legacy code.
```
[Legacy App] -- proprietary /stats --> [Adapter Container] -- /metrics (Prometheus) --> Prometheus
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the three well-known multi-container pod patterns? Give a real-world use case for each.
**Your Response:** "There are three main patterns. First, the Sidecar pattern - like having a helper container that assists the main app. A common example is a Fluentd sidecar that reads log files from the main app and forwards them to Elasticsearch. Second, the Ambassador pattern - a proxy container that handles network connections. For example, the main app always connects to localhost, and the ambassador figures out which actual database to connect to based on the environment. Third, the Adapter pattern - transforms output into standard formats. Like a legacy app with custom metrics, and an adapter container converts them to Prometheus format. All three patterns let containers share localhost and volumes, making them work together as a team."

---

### Question 67: What is a Startup Probe and when should you use it over a Liveness Probe?

**Answer:**

| Probe | Purpose | Restart on failure |
|---|---|---|
| **Liveness** | Is the app alive and not deadlocked? | Yes — restarts container |
| **Readiness** | Is the app ready to serve traffic? | No — removes from Service endpoints |
| **Startup** | Has the app finished initial startup? | Yes — restarts container |

**The problem Startup Probe solves:**
Slow-starting applications (e.g., a Java Spring Boot app that takes 60-90 seconds to load) would repeatedly fail Liveness checks during startup, causing K8s to restart the container in an infinite loop before it's ever ready.

**Solution:**
```yaml
startupProbe:
  httpGet:
    path: /health
    port: 8080
  failureThreshold: 30      # Allow up to 30 failures
  periodSeconds: 10         # Check every 10 seconds
  # = allows up to 30 × 10 = 300 seconds (5 min) for startup
```

Once the Startup Probe **succeeds once**, it hands off control to the Liveness and Readiness probes. Until then, Liveness is completely disabled.

> **Rule of thumb:** Always add a `startupProbe` for Java, .NET, or any app with known slow JVM/runtime warm-up.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Startup Probe and when should you use it over a Liveness Probe?
**Your Response:** "Startup Probe is like giving slow-starting applications extra time to wake up. The problem is that apps like Java Spring Boot can take 60-90 seconds to start, but regular liveness probes might start checking after just 10 seconds. If the app isn't ready yet, Kubernetes keeps restarting it in an infinite loop. Startup Probe disables liveness checks during startup and gives the app a grace period - like 5 minutes. Once the startup probe succeeds once, it hands off to the regular liveness and readiness probes. I always use startup probes for Java, .NET, or any application with known slow startup times."

---

### Question 68: Explain Pod Priority and Preemption. What happens when a cluster runs out of resources?

**Answer:**
**Pod Priority** lets you define the relative importance of Pods. Kubernetes uses this during scheduling when resources are scarce.

**Setup:**
```yaml
# Define priority classes
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: critical-production
value: 1000000          # Highest priority
globalDefault: false

---
apiVersion: scheduling.k8s.io/v1
kind: PriorityClass
metadata:
  name: batch-low
value: 100              # Low priority — can be evicted
```

**What happens during resource shortage:**
1. A `critical-production` pod is submitted but no node has enough CPU/Memory.
2. The scheduler looks for nodes where **evicting lower-priority pods** would free enough resources.
3. Lower-priority pods (`batch-low`) are **gracefully terminated** (SIGTERM → grace period → SIGKILL).
4. Once the node has sufficient capacity, the high-priority pod is scheduled.

**Real-world use case:** A batch ML training job (low priority) is running at night. During a traffic spike, a new payment-service pod (high priority) needs to be scheduled urgently — the batch job is preempted to make room.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Pod Priority and Preemption. What happens when a cluster runs out of resources?
**Your Response:** "Pod Priority is like having different priority levels for pods - VIPs versus regular passengers. When the cluster runs out of resources, high-priority pods can preempt lower-priority ones. I define PriorityClasses with different values - critical production pods get high values, batch jobs get low values. When a critical pod needs to be scheduled but there's no space, Kubernetes looks for nodes where evicting low-priority pods would free up enough resources. The low-priority pods get gracefully terminated with SIGTERM, then the high-priority pod gets scheduled. It's like bumping economy class passengers for first-class passengers when the flight is overbooked."

---

### Question 69: How do you design a Namespace strategy for a team of 50 developers working across 3 projects?

**Answer:**
A good namespace strategy creates isolation without over-fragmenting the cluster.

**Recommended: Environment × Project structure:**
```
project-alpha-dev
project-alpha-staging
project-alpha-prod
project-beta-dev
project-beta-prod
shared-services        (databases, message queues)
monitoring             (Prometheus, Grafana)
```

**Per-namespace controls to apply:**
```yaml
# 1. ResourceQuota — cap resource usage per team/env
apiVersion: v1
kind: ResourceQuota
metadata:
  name: quota
  namespace: project-alpha-dev
spec:
  hard:
    requests.cpu: "4"
    requests.memory: 8Gi
    pods: "20"

---
# 2. NetworkPolicy — default deny, allow only intra-namespace + ingress
# 3. RBAC — developers can deploy in dev, only CI/CD can deploy in prod
# 4. LimitRange — inject default resource limits for developer pods
```

**Access pattern:**
- Developers: `edit` ClusterRole scoped to `*-dev` namespaces only.
- CI/CD service account: `edit` on `*-staging` and `*-prod`.
- Platform team: `cluster-admin`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you design a Namespace strategy for a team of 50 developers working across 3 projects?
**Your Response:** "I'd use an environment-by-project structure - like project-alpha-dev, project-alpha-staging, project-alpha-prod for each project, plus shared namespaces for databases and monitoring. This gives isolation without creating too many namespaces. For each namespace, I'd apply ResourceQuotas to cap resource usage, NetworkPolicies for security, RBAC so developers can only deploy in dev environments, and LimitRanges for default resource limits. Developers get edit access only to dev namespaces, while CI/CD handles staging and production deployments. This balances isolation with manageability for a large team."

---

### Question 70: How does Kubernetes DNS resolve cross-namespace service discovery?

**Answer:**
Every Kubernetes Service gets a DNS entry in the format:
```
<service-name>.<namespace>.svc.cluster.local
```

**Within the same namespace:**
```bash
# Short name works inside the same namespace
curl http://user-service:8080/api
# CoreDNS resolves: user-service → user-service.team-a.svc.cluster.local
```

**Cross-namespace (requires full FQDN):**
```bash
# Service in team-b calling a service in shared-services namespace
curl http://db-service.shared-services.svc.cluster.local:5432
```

**Common mistake:** Forgetting the namespace in cross-namespace calls causes `Could not resolve host` errors that are often mistakenly blamed on the network.

**CoreDNS search domains** (defined in each Pod's `/etc/resolv.conf`):
```
search team-a.svc.cluster.local svc.cluster.local cluster.local
```
This is why short names resolve *within* the same namespace — CoreDNS appends the search domains automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does Kubernetes DNS resolve cross-namespace service discovery?
**Your Response:** "Every service gets a DNS name like 'service-name.namespace.svc.cluster.local'. Within the same namespace, I can use the short name like 'user-service:8080' because CoreDNS automatically appends search domains. But for cross-namespace calls, I need the full qualified domain name like 'db-service.shared-services.svc.cluster.local'. The common mistake is forgetting the namespace when calling services in other namespaces, which causes 'Could not resolve host' errors. CoreDNS handles the resolution - it's like the phone directory of the cluster, making sure services can find each other by name regardless of where they're running."

---

## 🔹 Cluster Operations & Upgrades (Questions 71-75)

### Question 71: Walk through a safe Kubernetes cluster upgrade procedure step by step.

**Answer:**
Kubernetes supports upgrading **one minor version at a time** (cannot skip from 1.24 → 1.26 directly).

**Pre-upgrade checklist:**
```bash
# 1. Check current version
kubectl version

# 2. Scan for deprecated APIs that will be removed in the target version
kubectl get --raw /metrics | grep apiserver_requested_deprecated_apis

# 3. Review the changelog for breaking changes
# https://kubernetes.io/releases/

# 4. Back up etcd (or use Velero for K8s objects)
ETCDCTL_API=3 etcdctl snapshot save /backup/etcd-$(date +%Y%m%d).db \
  --endpoints=https://127.0.0.1:2379 \
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \
  --cert=/etc/kubernetes/pki/etcd/server.crt \
  --key=/etc/kubernetes/pki/etcd/server.key
```

**Upgrade sequence (self-managed cluster with kubeadm):**
```bash
# Step 1: Upgrade control plane node(s) first
apt-get update && apt-get install -y kubeadm=1.27.x-00
kubeadm upgrade plan          # Preview changes
kubeadm upgrade apply v1.27.x # Apply upgrade
apt-get install -y kubelet=1.27.x-00 kubectl=1.27.x-00
systemctl restart kubelet

# Step 2: Drain and upgrade worker nodes one at a time
kubectl drain node-worker-1 --ignore-daemonsets --delete-emptydir-data
# SSH into worker node:
apt-get install -y kubeadm=1.27.x-00 kubelet=1.27.x-00 kubectl=1.27.x-00
kubeadm upgrade node
systemctl restart kubelet
# Back on control plane:
kubectl uncordon node-worker-1   # Re-enable scheduling on this node

# Repeat for remaining worker nodes
```

> **Managed clusters (EKS/GKE/AKS):** Control plane is upgraded via the cloud console/CLI. Worker nodes are upgraded by rolling through node groups, which respect PDBs automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Walk through a safe Kubernetes cluster upgrade procedure step by step.
**Your Response:** "I start with preparation - check the current version, scan for deprecated APIs that will be removed, review the changelog for breaking changes, and back up etcd. For self-managed clusters, I upgrade the control plane first using kubeadm - upgrade kubeadm, apply the upgrade, then upgrade kubelet and kubectl. After the control plane is upgraded, I drain each worker node one by one, upgrade the components, then uncordon it. For managed clusters like EKS, the control plane is handled by the cloud provider, and I just upgrade the worker node groups. The key is upgrading one minor version at a time and doing control plane before workers."

---

### Question 72: What is `kubectl drain` vs `kubectl cordon`? When would you use each?

**Answer:**

| Command | What it does | Use case |
|---|---|---|
| `kubectl cordon <node>` | Marks node as `Unschedulable` — no NEW pods will be placed here. Existing pods keep running. | Temporarily stop new work before maintenance without disrupting existing workloads. |
| `kubectl drain <node>` | Cordons the node AND evicts all existing pods off it (respecting PDBs and grace periods). | Before taking a node fully offline for maintenance, hardware replacement, or upgrade. |
| `kubectl uncordon <node>` | Marks node as `Schedulable` again. | After maintenance is complete. |

**Common `kubectl drain` flags:**
```bash
kubectl drain node-1 \
  --ignore-daemonsets \       # DaemonSet pods can't be moved, so ignore them
  --delete-emptydir-data \    # Allow deletion of pods using emptyDir volumes
  --grace-period=60 \         # Give pods 60s to terminate gracefully
  --timeout=300s              # Abort if drain takes more than 5 minutes
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `kubectl drain` vs `kubectl cordon`? When would you use each?
**Your Response:** "Cordon is like putting a 'closed for business' sign on a node - it stops new pods from being scheduled there but lets existing pods keep running. I use this when I'm preparing for maintenance but don't want to disrupt running workloads yet. Drain is more aggressive - it cordons the node AND evicts all existing pods, respecting PodDisruptionBudgets and grace periods. I use drain when I need to take a node completely offline for maintenance or upgrades. After maintenance, I use uncordon to reopen the node for new work. Think of cordon as 'stop accepting new customers' and drain as 'close the store and move all customers elsewhere'."

---

### Question 73: What is a PodDisruptionBudget and how does it integrate with cluster upgrades?

**Answer:**
A **PodDisruptionBudget (PDB)** protects your application during *voluntary disruptions* (drains, rolling upgrades) by defining the minimum number of pods that must remain available.

```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: api-pdb
  namespace: production
spec:
  minAvailable: 2      # At least 2 pods must be Running at all times
  selector:
    matchLabels:
      app: payment-api
```

**Integration with `kubectl drain`:**
- When draining a node, K8s calls the **Eviction API** (not direct delete) for each pod.
- The Eviction API checks: "Would evicting this pod violate its PDB?"
- If yes → eviction is **denied** and drain waits/retries.
- If no → pod is safely evicted.

**Example:** You have 3 replicas of `payment-api` and `minAvailable: 2`. Draining a node that holds 1 replica is allowed (3 - 1 = 2, budget satisfied). Draining a node that holds 2 replicas is blocked (3 - 2 = 1, would violate budget).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a PodDisruptionBudget and how does it integrate with cluster upgrades?
**Your Response:** "PodDisruptionBudget is like a safety net that ensures my application stays available during maintenance. It specifies the minimum number of pods that must be running at all times. When I drain a node for upgrades, Kubernetes uses the Eviction API instead of just deleting pods. The Eviction API checks if evicting a pod would violate its PDB - if it would, the eviction is denied and the drain waits. For example, if I have 3 replicas and minAvailable: 2, draining a node with 1 replica is fine, but draining a node with 2 replicas would be blocked. This prevents cluster upgrades from taking down too many pods at once and causing downtime."

---

### Question 74: How do you scan container images for vulnerabilities in a Kubernetes CI/CD pipeline?

**Answer:**
Image scanning should happen at multiple stages:

**Stage 1 — In CI pipeline (before push to registry):**
```bash
# Using Trivy (most popular open-source scanner)
trivy image --exit-code 1 --severity HIGH,CRITICAL my-app:v1.5
# Fails the CI pipeline if HIGH or CRITICAL CVEs found
```

**Stage 2 — In the registry (continuous scanning):**
- **AWS ECR:** Enable `scanOnPush` → ECR scans every pushed image automatically.
- **GCR / Artifact Registry:** Enable Container Analysis API.
- **Harbor (self-hosted):** Built-in Trivy integration, blocks pull of vulnerable images.

**Stage 3 — In Kubernetes (admission control):**
Using **Kyverno** or **OPA Gatekeeper** to block deployment of images with known CVEs:
```yaml
# Kyverno policy — block images not scanned within last 7 days
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: verify-image
spec:
  rules:
    - name: check-image-signature
      match:
        resources:
          kinds: ["Pod"]
      verifyImages:
        - image: "registry.company.com/*"
          attestors:
            - entries:
                - keyless:
                    subject: "https://github.com/my-org/*"
                    issuer: "https://token.actions.githubusercontent.com"
```

**Image signing:** Use **Cosign** (Sigstore) to cryptographically sign images in CI. Kyverno's `verifyImages` rule then enforces that only signed images can be deployed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you scan container images for vulnerabilities in a Kubernetes CI/CD pipeline?
**Your Response:** "I scan images at multiple stages for defense in depth. First in the CI pipeline before pushing to the registry using Trivy - this fails the build if high or critical vulnerabilities are found. Second in the registry itself with continuous scanning - ECR can scan on push, Harbor has built-in scanning. Third at runtime with admission controllers like Kyverno that can block deployment of vulnerable images. I also use Cosign to sign images cryptographically, and Kyverno can enforce that only signed images are deployed. This multi-layer approach ensures vulnerable images never make it to production."

---

### Question 75: What is Kyverno and how does it compare to OPA Gatekeeper for policy enforcement?

**Answer:**

| | **Kyverno** | **OPA Gatekeeper** |
|---|---|---|
| **Policy language** | Native Kubernetes YAML | Rego (custom language) |
| **Learning curve** | Low — familiar YAML syntax | High — Rego requires learning |
| **Capabilities** | Validate, Mutate, Generate, Verify Images | Validate only (via webhooks) |
| **Generate resources** | Yes — can auto-create NetworkPolicy on namespace creation | No |
| **Image verification** | Built-in Cosign support | Requires external setup |
| **Community** | CNCF Incubating | CNCF Graduated |

**Kyverno example — auto-create a NetworkPolicy whenever a new namespace is created:**
```yaml
apiVersion: kyverno.io/v1
kind: ClusterPolicy
metadata:
  name: add-default-deny-networkpolicy
spec:
  rules:
    - name: default-deny-ingress
      match:
        resources:
          kinds: ["Namespace"]
      generate:
        kind: NetworkPolicy
        name: default-deny-ingress
        namespace: "{{request.object.metadata.name}}"
        data:
          spec:
            podSelector: {}
            policyTypes: ["Ingress"]
```

**When to choose:**
- **Kyverno:** Teams new to policy-as-code; when you need mutation + generation, not just validation.
- **OPA Gatekeeper:** Teams already familiar with Rego; complex policy logic (cross-object validation); existing OPA investment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Kyverno and how does it compare to OPA Gatekeeper for policy enforcement?
**Your Response:** "Kyverno and OPA Gatekeeper are both policy enforcement tools but with different approaches. Kyverno uses native Kubernetes YAML, so it's easier to learn for teams already familiar with K8s. It can validate, mutate, generate resources, and verify images out of the box. OPA Gatekeeper uses Rego language which has a steeper learning curve but is more powerful for complex logic. Kyverno can automatically generate resources like NetworkPolicies when namespaces are created, while Gatekeeper only validates. I'd choose Kyverno for teams new to policy-as-code or when I need mutation and generation capabilities. I'd choose Gatekeeper for teams already comfortable with Rego or when I need complex cross-object validation."
