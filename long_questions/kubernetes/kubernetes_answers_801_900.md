## 🔹 Kubernetes Cost Optimization (Questions 801-810)

### Question 801: What tools help monitor Kubernetes cost per namespace?

**Answer:**
**Kubecost (OpenCost).**
- Breaks down cost by Namespace, Deployment, Service.
- Integrates with AWS/GCP Billing API to give real dollar values.

---

### Question 802: What is Kubecost and how does it work?

**Answer:**
Tooling installed in cluster.
- Prometheus Metrics (CPU seconds used) * Cloud Price Provider (AWS API) = Cost.
- Helps chargeback/showback for multi-tenant teams.

---

### Question 803: How can you identify overprovisioned resources?

**Answer:**
Comparing **Requests** vs **Usage**.
- If Request=4GB and Usage=500MB -> Overprovisioned.
- **Goldilocks:** Tool that visualizes VPA recommendations to right-size requests.

---

### Question 804: How do you right-size CPU and memory?

**Answer:**
- **Vertical Pod Autoscaler (Recommendation Mode):** Tells you the "correct" request.
- **Load Testing:** Stress test app to find saturation point, add 20% buffer.

---

### Question 805: What are the pros and cons of spot instances in Kubernetes?

**Answer:**
- **Pros:** 70-90% Cheaper.
- **Cons:** Can be terminated with 2-minute warning.
- **Strategy:** Use for Stateless, Batch, or Replicated workloads. Handle termination via `node-termination-handler`.

---

### Question 806: What is cluster autoscaler vs node autoscaler?

**Answer:**
Both refer to the same component: **Cluster Autoscaler** (CA).
- CA adds/removes cloud instances (Nodes) based on pod demand.
- HPA scales Pods. CA scales Nodes.

---

### Question 807: How does vertical pod autoscaler help with cost control?

**Answer:**
Prevents Overprovisioning.
- Automatically shrinks "Fat" requests to fit reality.
- Allows packing more pods onto the same node.

---

### Question 808: What are idle pods and how to detect them?

**Answer:**
Pods doing nothing (0 CPU/Network).
- Abandoned Dev deployments.
- Detect via Prometheus metrics (`container_cpu_usage` near zero for 7 days).

---

### Question 809: How do you use taints/tolerations for cost-based scheduling?

**Answer:**
- Taint expensive nodes (GPU/HighMem): `expensive=true:NoSchedule`.
- Only workloads that *need* it get the toleration.
- Prevents random "Hello World" pod landing on a $10/hour instance.

---

### Question 810: What is FinOps in the context of Kubernetes?

**Answer:**
Financial Operations.
- Cultural shift: "Engineers own their costs".
- Using Metadata (Labels) to tag every cost.
- Optimization loops.

---

## 🔹 Cloud-Native Security (Advanced) (Questions 811-820)

### Question 811: What is runtime security in Kubernetes?

**Answer:**
Detecting attacks **after** deployment (while running).
- **Falco:** "Shell spawned in container!", "File /etc/shadow read!".
- Static scanning (Trivy) misses zero-day exploits or runtime config drift.

---

### Question 812: How does Falco detect suspicious container behavior?

**Answer:**
- Hooks into kernel (eBPF or Kernel Module).
- Watches System Calls (`execve`, `open`, `connect`).
- Compares against Rules ("Unexpected process spawning shell").

---

### Question 813: What is Aqua Trivy and what does it scan?

**Answer:**
Comprehensive Scanner.
- **Images:** OS packages (CVEs) + Language deps (npm, pip).
- **Filesystem:** Secrets, Misconfigs (IaC).
- **K8s:** Cluster scanning (RoleBindings).

---

### Question 814: How do you protect against container escape vulnerabilities?

**Answer:**
- **Isolation:** gVisor / Kata Containers (Sandboxing).
- **Proactive:** Seccomp profiles, AppArmor, ReadOnlyRootFilesystem.
- **Update:** Patch Kernel/Kubelet frequently.

---

### Question 815: What is container image signing?

**Answer:**
**Cosign / Notary.**
- Cryptographically signing the image digest.
- Verification: Kubernetes Admission Controller verifies signature before allowing `Pull`. "If not signed by MyCompany -> Reject".

---

### Question 816: What is Notary v2 and how is it used with Kubernetes?

**Answer:**
(Now **Notation** within CNCF).
- Standard for signing OCI artifacts.
- Integrates with Registries (AWS ECR, Azure ACR).

---

### Question 817: How do you implement workload identity for cloud access?

**Answer:**
**OIDC Federation.**
- K8s ServiceAccount Token (JWT) is exchanged for AWS Cloud Token (STS).
- No long-lived Access Keys in secrets.

---

### Question 818: What is SPIFFE/SPIRE in zero-trust Kubernetes?

**Answer:**
**SPIFFE:** Standard for Identity (SVID).
**SPIRE:** Implementation.
- Attests the identity of workloads based on generic properties (it runs on this node, signed by K8s).
- Issues short-lived mTLS certs.

---

### Question 819: How do you protect service meshes from L7 attacks?

**Answer:**
- **WAF:** Envoy supports WAF filters (ModSecurity/Coraza).
- **AuthZ:** Istio AuthorizationPolicy (Reject requests missing specific headers).

---

### Question 820: How do you enforce egress traffic restrictions?

**Answer:**
**NetworkPolicy.**
```yaml
egress:
- to:
  - ipBlock: { cidr: 10.0.0.0/8 }
```
- Egress Gateway (Istio) allows whitelisting "google.com" but blocking "malware.com".

---

## 🔹 API Server Deep Dive (Questions 821-830)

### Question 821: How does the Kubernetes API server handle incoming requests?

**Answer:**
Chain of handlers:
1.  **Authentication:** (Who are you?)
2.  **Authorization:** (Can you do this?)
3.  **Admission Control:** (Should we allow this?)
4.  **Validation:** (Is YAML valid?)
5.  **Persistence:** (Write to Etcd).

---

### Question 822: What are admission phases in request processing?

**Answer:**
- **Mutating Phase:** (Defaults, Sidecars).
- **Schema Validation:** (OpenAPI).
- **Validating Phase:** (Policy checks).

---

### Question 823: What is the role of API aggregation layer?

**Answer:**
Allows extending core K8s API with extra servers (Metrics Server).
- Path `/apis/metrics.k8s.io` is proxied to the Metrics Server pod.
- Path `/apis/batch` is handled by Kube-Apiserver.

---

### Question 824: How do CRDs extend the Kubernetes API?

**Answer:**
Dynamically.
- You post a CRD definition.
- API Server updates its internal routing table to accept REST calls for that new Path.
- No reboot needed.

---

### Question 825: How do you secure the API server endpoint?

**Answer:**
- **Private Access:** Disable Public IP.
- **CIDR Restriction:** Allow only VPN IP.
- **Auth:** Disable Anonymous Auth.

---

### Question 826: What is RBAC escalation and how do you prevent it?

**Answer:**
Allowing a user to grant rights they don't have.
- K8s prevents this natively (`user` check).

---

### Question 827: What is the difference between built-in and custom resources?

**Answer:**
- **Built-in:** (Pod, Svc). Hardcoded in Go. High performance.
- **CRD:** Defined in YAML (Metacode). Slightly slower (stored as JSON blob).

---

### Question 828: How does the watch mechanism work in the API server?

**Answer:**
Efficient HTTP Long Polling / Chunked encoding.
- Client says "Tell me updates to Pods from resourceVersion 100".
- API Server holds connection open and pushes events.

---

### Question 829: What is an informer and how does it optimize traffic?

**Answer:**
Client-side cache helper (Go).
- Handles the Watch connection.
- Re-populates Store.
- Users verify against Store (Memory), not API Server (Network).

---

### Question 830: How do you troubleshoot slow API server responses?

**Answer:**
- **Metrics:** `apiserver_request_duration_seconds`.
- **Logs:** Look for "Throttling request" or "Etcd slow".
- **Audit Logs:** Check who is hammering the API (e.g., a broken CI script listing pods every 10ms).

---

## 🔹 Advanced Networking & DNS (Questions 831-840)

### Question 831: How does CoreDNS work inside Kubernetes?

**Answer:**
Deployment.
- Reads `Corefile`.
- Uses `kubernetes` plugin to watch API.
- Serves DNS requests from Pods on Port 53 (UDP/TCP).

---

### Question 832: What is stubDomains in CoreDNS?

**Answer:**
Conditional forwarding.
- "If domain ends in `consul.local`, forward to `10.1.2.3`".
- Used to integrate K8s with legacy corporate DNS.

---

### Question 833: How do you debug DNS issues in Kubernetes?

**Answer:**
- Check `/etc/resolv.conf`.
- `nslookup kubernetes.default`.
- Check CoreDNS logs `kubectl logs -n kube-system -l k8s-app=kube-dns`.

---

### Question 834: How does headless service DNS resolution work?

**Answer:**
Returns **Multiple A Records** (list of Pod IPs).
- Client side Load Balancing (Round Robin).

---

### Question 835: What is kube-proxy and how does it implement networking?

**Answer:**
- Watches Services/Endpoints.
- Updates Kernel Rules (Iptables/IPVS).
- Implements the "Virtual IP" magic.

---

### Question 836: What is IPVS mode and how is it configured?

**Answer:**
Kernel Load Balancer.
- Config: `kube-proxy --proxy-mode=ipvs`.
- Scales to 100k services. Iptables degrades at 5k.

---

### Question 837: How do you enable IPv6 in Kubernetes?

**Answer:**
- Requires CNI support (Calico/Cilium).
- Requires Dual-Stack configuration on API Server/Controller Manager/Kubelet.
- `--service-cluster-ip-range=IPv4CIDR,IPv6CIDR`.

---

### Question 838: How does dual-stack networking work in clusters?

**Answer:**
Pods get **Two IPs** (one v4, one v6).
- Services get two ClusterIPs.
- Allows native communication with IPv6 internet.

---

### Question 839: What is the `externalTrafficPolicy` setting?

**Answer:**
Controls how traffic reaches nodes.
- **Cluster (Default):** NodePort on Node A forwards to Pod on Node B (Second hop, SNAT, internal traffic).
- **Local:** NodePort on Node A drops packet if Pod not on Node A. (Preserves Client IP, No extra hop).

---

### Question 840: How do you configure DNS caching in nodes?

**Answer:**
**NodeLocal DNSCache.**
- DaemonSet runs a mini-DNS on every node.
- Pods query `169.254.20.10` (Localhost).
- Reduces Conntrack entries and UDP latency to CoreDNS.

---

## 🔹 Multi-Cluster & Federation (Questions 841-850)

### Question 841: What is Kubernetes Federation v2?

**Answer:**
See previous sections. (Evolution of KubeFed, focus on propagation policies).

---

### Question 842: How does KubeFed synchronize resources across clusters?

**Answer:**
- **Templates:** The resource definition.
- **Placement:** Which clusters?
- **Overrides:** Specific values for Cluster B.

---

### Question 843: What is Submariner and how does it enable inter-cluster networking?

**Answer:**
CNCF project.
- Connects overlay networks of different clusters using **VPN Tunnels** (IPSec/Wireguard).
- Pod in Cluster A can ping Pod in Cluster B directly.

---

### Question 844: What are typical use cases for multi-cluster deployments?

**Answer:**
- Hybrid Cloud (Effectively connecting On-Prem and AWS).
- Compliance (Data Residency).
- Blast Radius reduction.

---

### Question 845: How do you manage failover between clusters?

**Answer:**
**Global Load Balancer (GSLB).**
- Monitoring Probe detects Cluster A failure.
- Update DNS to point only to Cluster B.
- App must be stateless or have replicated DB.

---

### Question 846: What is service mesh expansion in multi-cluster?

**Answer:**
Running one logical mesh over 2 clusters.
- Istio Multicluster.
- Control Plane in Cluster A (Primary) or Replicated.

---

### Question 847: What is Crossplane multi-cluster composition?

**Answer:**
- One Crossplane Management Cluster.
- Provisions "Child Clusters" (EKS) and pushes apps to them.

---

### Question 848: How do you secure communication across clusters?

**Answer:**
- **mTLS:** via Service Mesh.
- **VPN/Peering:** via Cloud networking (VPC Peering).

---

### Question 849: How do you do service discovery across clusters?

**Answer:**
- MCS (Multi-Cluster Services) API.
- Imports `ServiceExport` from Cluster A -> `ServiceImport` in Cluster B.

---

### Question 850: What is a multi-region GKE/Amazon EKS strategy?

**Answer:**
Running identical stacks in `us-east` and `eu-west`.
- CI/CD pipelines deploy to all regions simultaneously (or staged).
- Database is the hard part (Global Spanner / Aurora Global).

---

## 🔹 CI/CD with Kubernetes (Questions 851-860)

### Question 851: What is a Kubernetes-native CI/CD pipeline?

**Answer:**
Pipeline runs **as Pods**.
- **Tekton:** The standard.
- **Jenkins X:** Built on Tekton.
- Steps are containers.

---

### Question 852: How does Tekton differ from Argo Workflows?

**Answer:**
- **Tekton:** Focused on CI (Building images). granular.
- **Argo Workflows:** Focused on General Workflows / Data / ML. (DAGs).

---

### Question 853: What is a pipelineRun in Tekton?

**Answer:**
An execution of a Pipeline.
- `Pipeline`: The definition (Build -> Test -> Push).
- `PipelineRun`: "Run it now with git-revision ABC".

---

### Question 854: How do you implement progressive delivery with Argo Rollouts?

**Answer:**
Argo Rollouts replaces `Deployment`.
- `kind: Rollout`.
- Strategy: Canary.
- `setWeight: 20` -> `pause: {duration: 1h}` -> `setWeight: 100`.

---

### Question 855: What are blue-green vs canary deployments in K8s?

**Answer:**
- **B/G:** Switch 100% traffic from V1 to V2 instantly. (Safe fallback).
- **Canary:** Slow drift (1% -> 5% -> 100%). (Low impact of errors).

---

### Question 856: How do you validate manifests before deploying?

**Answer:**
- `kubectl apply --dry-run=server`.
- `kubeval` / `kubeconform` (Schema check).
- `conftest` (Policy check).

---

### Question 857: How do you test Helm charts in CI?

**Answer:**
- `helm lint`.
- `helm test` (Runs Pods defined in `templates/tests/`).
- `chart-testing (ct)` tool.

---

### Question 858: What is GitHub Actions’ integration with Kubernetes?

**Answer:**
Runners can deploy to K8s.
- `azure/k8s-deploy`, `aws-actions/amazon-eks-kubeconfig`.
- Best practice: Don't let GHA touch K8s. Let GHA push to Git, let ArgoCD touch K8s.

---

### Question 859: What’s the role of a buildkit + Kaniko combo?

**Answer:**
Building Docker images **inside** Kubernetes.
- Docker Daemon is usually unavailable (rootless).
- **Kaniko:** Builds images in userspace without daemon.

---

### Question 860: How do you achieve zero-downtime deployments?

**Answer:**
- `RollingUpdate` strategy.
- `readinessProbe` configured correctly.
- `preStop` hook to drain connections.

---

## 🔹 Observability with OpenTelemetry (Questions 861-870)

### Question 861: What are traces vs metrics vs logs?

**Answer:**
The Three Pillars.
- **Metrics:** Aggregates (Counters for dashboard).
- **Logs:** Events (Text).
- **Traces:** Requests flow (Latency analysis).

---

### Question 862: What is an OpenTelemetry Collector?

**Answer:**
A middleware proxy.
- **Receiver:** Receives data (OTLP, Jaeger, Prometheus).
- **Processor:** Filters/Batches/Anonymizes.
- **Exporter:** Sends to Backend (Splunk, Datadog).

---

### Question 863: How do you instrument Go applications for tracing?

**Answer:**
Import `go.opentelemetry.io/otel`.
- Start span `tracer.Start(ctx, "operation")`.
- `defer span.End()`.

---

### Question 864: What’s the role of exporters in OpenTelemetry?

**Answer:**
Plugins to send data to specific vendor.
- Code agnostic. Change exporter to switch from Jaeger to Zipkin without changing app code.

---

### Question 865: What are spans in a distributed trace?

**Answer:**
A single unit of work.
- "SQL Query", "HTTP Handler".
- Have Start/End time and metadata.
- Parent-Child relationship builds the Trace.

---

### Question 866: How do you view OpenTelemetry traces in Jaeger?

**Answer:**
Jaeger UI visualizes the Gantt chart of the Trace.

---

### Question 867: What is semantic convention in OpenTelemetry?

**Answer:**
Standard naming.
- `http.method` (not `method` or `verb`).
- `db.system` (postgres).
- Ensures consistency across tools.

---

### Question 868: How do you correlate logs and traces in Kubernetes?

**Answer:**
Inject `TraceID` and `SpanID` into the Log Message context.
- Logging tools (Loki) can then link from Log line to Trace view.

---

### Question 869: What’s the difference between OTLP/HTTP and OTLP/gRPC?

**Answer:**
Protocol transport.
- **gRPC:** More efficient, binary, persistent connection. Recommended.
- **HTTP:** Easier to debug/proxy through standard LBs.

---

### Question 870: How does OpenTelemetry integrate with Prometheus?

**Answer:**
- OTel Collector can scrape Prometheus endpoints.
- Or OTel Collector can export *to* Prometheus (Expose `/metrics` endpoint).

---

## 🔹 StatefulSets & Advanced Storage (Questions 871-880)

### Question 871: How does StatefulSet differ from Deployment?

**Answer:**
- **Identity:** Sticky (web-0, web-1).
- **Ordering:** Start 0 -> Ready -> Start 1.
- **Storage:** VolumeClaimTemplates (Each replica gets unique PVC).

---

### Question 872: What are stable network IDs in StatefulSets?

**Answer:**
DNS names predictible: `web-0.service.ns`.
- Allows peers to find each other for clustering (Raft/Paxos).

---

### Question 873: How do you scale StatefulSets gracefully?

**Answer:**
Scale down happens in reverse order (2 -> 1 -> 0).
- Pre-stop hooks crucial to Deregister from cluster.

---

### Question 874: What is `volumeClaimTemplates` in StatefulSets?

**Answer:**
The template.
- "For every pod created, create a PVC based on this template".

---

### Question 875: How do you handle quorum-sensitive apps like Kafka or Zookeeper?

**Answer:**
`PodDisruptionBudget` is mandatory (`minAvailable: 2`).
- Never drain more nodes than the quorum can handle.

---

### Question 876: How do you handle node failure in StatefulSet-based apps?

**Answer:**
- Pod `web-0` goes Unknown.
- **It does not auto-reschedule** quickly (Safety). K8s fears "Split Brain" (Two web-0 writing to same disk).
- Admin must `force delete` the dead pod to allow it to restart elsewhere.

---

### Question 877: What is `podManagementPolicy` and when to use it?

**Answer:**
- **OrderedReady (Default):** Strict sequential.
- **Parallel:** Start/Kill all at once. Use for stateless workers that just need Stable IDs.

---

### Question 878: How do you ensure state persistence after cluster upgrades?

**Answer:**
PV/PVCs persist independent of Pods/Nodes.
- Reattaching might be slow (Cloud detach/attach).

---

### Question 879: How do you implement leader election in Stateful workloads?

**Answer:**
- **Sidecar:** (zk-election sidecar).
- **App Logic:** (Zookeeper/Etcd locking).
- **K8s API:** Lease API.

---

### Question 880: What is etcd defragmentation and why is it important?

**Answer:**
Etcd storage space doesn't shrink automatically after deletes.
- "Space" remains allocated.
- Defrag releases it back to FS.
- Maintenance task.

---

## 🔹 Custom Controllers & Operators (Questions 881-890)

### Question 881: What is a Kubernetes operator?

**Answer:**
Pattern of building custom controllers.

---

### Question 882: What is the Operator SDK?

**Answer:**
Framework to accelerate Operator dev.

---

### Question 883: What is the role of controller-runtime library?

**Answer:**
Common Go library used by Kubebuilder / Operator SDK.
- Handles the heavy lifting (Informer, workqueue, LeaderElection).

---

### Question 884: How do you implement reconciliation loops?

**Answer:**
Idempotent logic.
- "Ensure deployment exists".
- "If exists, ensure replicas matches spec".
- "If not, update it".

---

### Question 885: What is the difference between a reconciler and a webhook?

**Answer:**
- **Webhook:** Sync (Before creation). Validation/Defaults.
- **Reconciler:** Async (Background). Enforcement.

---

### Question 886: What is the purpose of finalizers in a custom resource?

**Answer:**
Cleanup.
- "Wait until RDS instance is destroyed before deleting this CR YAML".

---

### Question 887: How do you test custom controllers?

**Answer:**
- **Unit:** Fake Client.
- **Integration:** EnvTest (Spins up local API Server).
- **E2E:** Kind cluster.

---

### Question 888: What are some examples of Kubernetes operators in the wild?

**Answer:**
Prometheus Operator, Strimzi (Kafka), ECK (Elastic), Postgres Operator (Zalando).

---

### Question 889: What is the difference between imperative and declarative controllers?

**Answer:**
- **Declarative:** "Make it look like X". (Ideal).
- **Imperative:** "Run backup NOW". (Usually modeled as a Job CR).

---

### Question 890: How do CRDs get versioned and evolved?

**Answer:**
Storage Versions / Served Versions.
- Conversion Webhooks migrate data.

---

## 🔹 Governance, Compliance & Audit (Questions 891-900)

### Question 891: What is audit logging in Kubernetes?

**Answer:**
Records all API requests.

---

### Question 892: How do you configure audit log policies?

**Answer:**
YAML file passed to API Server.
- `level: Metadata` (Who/When).
- `level: RequestResponse` (Full Payload).

---

### Question 893: What is the role of cloud-native policy engines?

**Answer:**
Centralized guardrails.

---

### Question 894: How do you enforce security standards like NIST or CIS?

**Answer:**
Kyverno Policy Packs / OPA Library.
- Pre-made policies mapping to NIST controls (e.g., "Require Limit Range").

---

### Question 895: What is policy as code?

**Answer:**
Treating Policies (Rego/YAML) as software (Git, CI/CD, Testing).

---

### Question 896: What are compliance risks in multi-tenant clusters?

**Answer:**
Data leakage (Shared volumes), Network hops (No NetworkPolicy), Node Privilege escalation.

---

### Question 897: How do you integrate policy enforcement in CI pipelines?

**Answer:**
Run `conftest` or `kyverno apply --dry-run` in CI.
- Fail build if YAML violates policy.

---

### Question 898: How do you implement least privilege in RBAC?

**Answer:**
- No wildcards `*`.
- Specific Resources (`pods`, not `*`).
- Automated scanning (Kubiscan).

---

### Question 899: What tools help detect misconfigured Kubernetes workloads?

**Answer:**
**Datree, Polaris, Kube-score.**
- Static analysis of YAMLs.

---

### Question 900: How do you automate audit remediation in Kubernetes?

**Answer:**
Policy Engines (Kyverno) can **Mutate** (Fix) bad resources automatically or just Block them.
- "If image tag missing, append :latest" (Bad practice, but example of mutation).
- Better: Block and Notify.
