## 🔹 Kubernetes Edge Use Cases & Advanced Scheduling (Questions 901-910)

### Question 901: How do you deploy Kubernetes on edge devices?

**Answer:**
Use lightweight distributions: **K3s, MicroK8s, K0s**.
- Low resource usage (<512MB RAM).
- Connect via simple dial-back tunnels (if behind NAT).
- **Fleet Management:** Rancher Fleet / Azure Arc.

---

### Question 902: What is K3s and how does it differ from K8s?

**Answer:**
Certified Kubernetes distro by Rancher.
- Single binary.
- Stripped of legacy cloud provider plugins.
- Uses SQLite instead of Etcd by default.
- Ideal for IoT.

---

### Question 903: How do you run Kubernetes on Raspberry Pi?

**Answer:**
1.  Install OS (Ubuntu/Raspbian).
2.  Enable cgroups (`cgroup_memory=1 cgroup_enable=cpuset`).
3.  `curl -sfL https://get.k3s.io | sh -`.

---

### Question 904: What is a topology spread constraint?

**Answer:**
Control skew of pods across Failure Workloads.
- "Don't put all 10 pods in Zone A. Spread them: 3 in A, 3 in B, 4 in C."

---

### Question 905: How do you ensure node affinity in mixed hardware clusters?

**Answer:**
Label nodes: `gpu=nvidia`, `arch=arm64`.
Pod Spec:
```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
      - matchExpressions:
        - key: arch
          operator: In
          values: [arm64]
```

---

### Question 906: What’s the difference between `affinity` and `antiAffinity`?

**Answer:**
- **Affinity:** Attraction (Run close).
- **AntiAffinity:** Repulsion (Run apart).

---

### Question 907: What is schedulingPolicy in K8s?

**Answer:**
Deprecated term, usually refers to **Scheduling Profiles** now (Kube-scheduler config).
- Defining which Plugins (Filter/Score) are enabled.

---

### Question 908: How do you schedule pods based on node labels?

**Answer:**
`nodeSelector` (Simple) or `nodeAffinity` (Complex).

---

### Question 909: What’s the difference between `preferredDuringScheduling` and `requiredDuringScheduling`?

**Answer:**
- **Required:** Hard. "No match? Pending forever."
- **Preferred:** Soft. "No match? Put it anywhere else."

---

### Question 910: How do you debug pod placement failures?

**Answer:**
`kubectl describe pod`.
- Look at **Events**. "0/5 nodes available: 3 Taint, 2 Insufficient CPU."
- Scheduler explicity tells you why it failed.

---

## 🔹 Cloud-Specific Kubernetes Implementations (Questions 911-920)

### Question 911: What is GKE Autopilot mode?

**Answer:**
Fully managed GKE.
- You don't see content of Nodes. You just pay per Pod resource requests.
- "Serverless-feel" K8s. Google manages node scaling.

---

### Question 912: How does AWS EKS manage control plane?

**Answer:**
- AWS runs the Masters (API/Etcd) in their VPC.
- High Availability (3 AZs).
- You trigger upgrades via API. AWS rotates instances seamlessly.

---

### Question 913: What are node groups in Amazon EKS?

**Answer:**
Managed Auto Scaling Groups.
- "NodeGroup A": t3.medium.
- EKS manages the draining/upgrading of these nodes via EC2 logic.

---

### Question 914: How does Azure AKS handle node pool scaling?

**Answer:**
**Virtual Machine Scale Sets (VMSS).**
- AKS Cluster Autoscaler adds VM instances to the Scale Set.

---

### Question 915: What is workload identity federation in GKE?

**Answer:**
IAM integration.
- K8s ServiceAccount acts as IAM Service Account.
- No keys needed.

---

### Question 916: What are the differences between Fargate and EKS worker nodes?

**Answer:**
- **EC2 Node:** You manage OS patching, DaemonSets allowed, cheaper.
- **Fargate:** No OS access, No DaemonSets (No Sidecars mostly), Pay per vCPU, higher isolation.

---

### Question 917: How do you use ARM64 nodes in cloud clusters?

**Answer:**
- Add a NodePool with ARM instances (AWS Graviton / Azure Altra).
- Multi-arch images required for Pods!

---

### Question 918: What is Anthos and how does it extend Kubernetes?

**Answer:**
Google's multi-cloud platform.
- Manage clusters on GKE, AWS, Azure, and On-Prem from one Google Console.

---

### Question 919: What is Azure Arc-enabled Kubernetes?

**Answer:**
Connect any K8s cluster (running anywhere) to Azure Portal.
- Apply Azure Policy and Monitoring to your on-prem cluster.

---

### Question 920: How does IAM integrate with K8s on different cloud providers?

**Answer:**
Standard is **OIDC**.
- Cloud Provider trusts the K8s OIDC issuer.
- Pod Token (JWT) -> Cloud Token exchange.

---

## 🔹 Advanced Autoscaling Techniques (Questions 921-930)

### Question 921: What is HPA based on custom metrics?

**Answer:**
Scaling on non-resource metrics.
- `queue_length`, `http_requests_per_second`.
- Requires Adapter (Prometheus Adapter).

---

### Question 922: What is Kubernetes VPA and how does it work?

**Answer:**
- **Recommender:** Watches usage history.
- **Updater:** Evicts pods with old limits.
- **Admission Plugin:** Mutates new pod with correct limits.

---

### Question 923: What is a predictive autoscaler?

**Answer:**
AI-based scaling.
- "Traffic rises every Monday at 9AM. Scale out at 8:55AM."
- KEDA supports some predictive logic.

---

### Question 924: How does KEDA autoscale based on Kafka or Redis?

**Answer:**
**Scalers.**
- Redis Scaler checks `LLEN list`.
- Kafka Scaler checks `Consumer Lag`.
- If Lag > X, Scale deployment.

---

### Question 925: What are the limitations of HPA?

**Answer:**
- Reactive (Lags behind spike).
- Thrashing (Oscillation).
- Requires Metrics Server.

---

### Question 926: What is resource-based autoscaling?

**Answer:**
Standard HPA (CPU/Mem).

---

### Question 927: How do you configure autoscaling with Prometheus metrics?

**Answer:**
1.  Prometheus collects metrics.
2.  Prometheus Adapter exposes them via `custom.metrics.k8s.io`.
3.  HPA targets that API.

---

### Question 928: What is cluster autoscaler expander?

**Answer:**
Strategy when scaling UP.
- **random:** Pick any node group.
- **least-waste:** Pick node group that will be most utilized after scale up.
- **priority:** User priority.

---

### Question 929: How do you avoid scaling flaps in autoscalers?

**Answer:**
**Stabilization Window.**
- "Don't scale down until usage is low for 5 minutes".

---

### Question 930: How do you autoscale StatefulSets?

**Answer:**
Possible with HPA but **Risk**.
- Scaling DB replicas usually requires data rebalancing. HPA doesn't know how to do that.
- Use Operators instead.

---

## 🔹 Performance & Tuning (Questions 931-940)

### Question 931: How do you tune kubelet performance?

**Answer:**
- `--kube-reserved`: Protect System processes.
- `--image-pull-progress-deadline`: Fail fast.
- `--event-qps`: Limit chatter.

---

### Question 932: What are cgroups and how do they affect performance?

**Answer:**
Linux control groups.
- Enforce CPU/Mem limits.
- **v2:** Better memory management and storage IO isolation.

---

### Question 933: How do you manage NUMA-aware scheduling?

**Answer:**
TopologyManager. Align CPU and Memory on same socket.

---

### Question 934: What are CPU/memory requests vs limits best practices?

**Answer:**
- Java: Set Request=Limit (Guaranteed).
- Node: Request < Limit (Burstable).

---

### Question 935: How does kubelet evict pods under resource pressure?

**Answer:**
Rank by QoS.
- 1. BestEffort (Delete first).
- 2. Burstable (Delete if using > Request).
- 3. Guaranteed (Delete last).

---

### Question 936: What are image pull strategies and how do they affect cold starts?

**Answer:**
- **Pre-baking:** AMI with images.
- **Node Caching:** DaemonSet pre-pulling.

---

### Question 937: How do you improve pod startup time?

**Answer:**
- Probes: Reduce `initialDelay`.
- Image: Smaller size.
- Lazy Loading: (Stargz).

---

### Question 938: How do you benchmark pod performance?

**Answer:**
`sysbench` inside pod, or `k6` outside pod.

---

### Question 939: What is the impact of container logging on performance?

**Answer:**
High stdout churn = High Docker load / IOPS.
- Can crash node.
- Use Asynchronous logging.

---

### Question 940: How do you detect noisy neighbor issues?

**Answer:**
- **CPU Throttling:** Metric `container_cpu_cfs_throttled_seconds`.
- **Latency:** Spikes when another heavy pod lands on node.

---

## 🔹 GitOps Advanced Concepts (Questions 941-950)

### Question 941: What is drift detection in GitOps?

**Answer:**
Continuous comparison of Git vs Etcd.

---

### Question 942: What’s the difference between sync and reconcile?

**Answer:**
They are often synonymous in GitOps context.
- **Sync:** The act of applying Git to Cluster.

---

### Question 943: What are ArgoCD sync waves?

**Answer:**
Ordering within an App.
- `argocd.argoproj.io/sync-wave: "1"` (DB).
- `argocd.argoproj.io/sync-wave: "5"` (App).
- Ensures DB created first.

---

### Question 944: How do you implement image update automation in GitOps?

**Answer:**
**Argo Image Updater / Flux Image Automation.**
- Watches registry. New tag -> Commits change to Git -> Syncs.

---

### Question 945: How does Flux support multi-tenancy?

**Answer:**
`ServiceAccount` impersonation.
- Flux Controller (Admin) applies changes using restricted Tenant SA permissions.

---

### Question 946: What are health checks in ArgoCD applications?

**Answer:**
Lua scripts assessing resource health.

---

### Question 947: How do you securely manage secrets in GitOps?

**Answer:**
External Secrets.

---

### Question 948: How do you promote apps across environments using GitOps?

**Answer:**
PR (Pull Request).
- Merge `dev` branch to `stage`.

---

### Question 949: What is a GitOps pipeline with Helm + ArgoCD?

**Answer:**
Git contains `HelmRelease` CR or `Application` pointing to Helm Chart + Values.

---

### Question 950: How do you implement GitOps fallback mechanisms?

**Answer:**
Git Revert.

---

## 🔹 Policy as Code & Governance (Questions 951-960)

### Question 951: How does Kyverno differ from OPA Gatekeeper?

**Answer:**
Native YAML vs Rego.

---

### Question 952: How do you write policies for image signature enforcement?

**Answer:**
Kyverno Policy `verifyImages`. Checks Cosign signature.

---

### Question 953: How do you restrict privileged containers using policies?

**Answer:**
Validate `securityContext.privileged == false`.

---

### Question 954: What is constraint templating in Gatekeeper?

**Answer:**
Rego logic template ("Must have label X").
Constraint instance ("Must have label 'cost-center'").

---

### Question 955: How do you apply policies per namespace?

**Answer:**
Selectors: `match: namespaces: ["prod"]`.

---

### Question 956: How do you validate ingress/egress policies?

**Answer:**
NetworkPolicy testing tools (Cilium Editor).

---

### Question 957: What is policy violation alerting?

**Answer:**
Policy Engine emit events/metrics.
- AlertManager watches `gatekeeper_violations`.

---

### Question 958: How do you manage exception handling in policies?

**Answer:**
Exclusion list.
- "Enforce everywhere EXCEPT namespace 'kube-system'".

---

### Question 959: What is mutating vs validating policy?

**Answer:**
- **Mutating:** Change it (Add default label).
- **Validating:** Block it.

---

### Question 960: How do you audit policy violations retrospectively?

**Answer:**
Scan existing cluster report (Audit).

---

## 🔹 Disaster Recovery & HA (Questions 961-970)

### Question 961: How do you set up an etcd backup and restore?

**Answer:**
CronJob running `etcdctl snapshot save`.
Push to S3.

---

### Question 962: What is a quorum loss in etcd and how do you recover?

**Answer:**
>50% nodes down.
- Cluster Read-Only.
- Recover: Restore from snapshot on a new node (create new cluster of 1, join others).

---

### Question 963: What is etcd defragmentation?

**Answer:**
Reclaiming disk space.

---

### Question 964: How do you implement multi-master HA in on-prem clusters?

**Answer:**
Load Balancer (HAProxy/Keepalived) in front of 3 API Servers.

---

### Question 965: What’s the role of external load balancers for HA API servers?

**Answer:**
Provide single endpoint (`k8s.example.com`) to workers/users.

---

### Question 966: How do you restore a cluster from an etcd snapshot?

**Answer:**
Disaster recovery procedure. Restores state. Does NOT restore persisted data volumes (need Velero for that).

---

### Question 967: What is Velero and how does it help with backup?

**Answer:**
Backups K8s resources + PVCs.

---

### Question 968: How do you test K8s disaster recovery scenarios?

**Answer:**
Game Days. Deliberately destroy cluster and time restoration.

---

### Question 969: What is kubeadm high availability setup?

**Answer:**
`kubeadm init` -> `kubeadm join --control-plane`. Stacks control plane nodes.

---

### Question 970: How do you migrate workloads between clusters?

**Answer:**
Export YAMLs (Velero) -> Import. Or GitOps sync to new cluster.

---

## 🔹 Service Mesh (Advanced) (Questions 971-980)

### Question 971: What is Istio ambient mesh?

**Answer:**
Sidecar-less Istio.
- Uses per-node proxy (ztunnel) and Waypoint proxies.
- Cheaper (No sidecar overhead per pod).

---

### Question 972: What’s the difference between Istio and Linkerd in terms of architecture?

**Answer:**
- Istio: Envoy (C++). Feature rich.
- Linkerd: Linkerd-proxy (Rust). Micro-proxy.

---

### Question 973: How do you visualize service-to-service communication in Istio?

**Answer:**
**Kiali.** Dashboard for mesh topology.

---

### Question 974: What are virtual services in Istio?

**Answer:**
Routing rules.

---

### Question 975: How do you implement mTLS with Istio?

**Answer:**
`PeerAuthentication: STRICT`.

---

### Question 976: How does Envoy proxy sidecar work?

**Answer:**
Iptables redirect traffic into Envoy port.

---

### Question 977: How do you do request-level routing in Istio?

**Answer:**
Match headers/path in VirtualService.

---

### Question 978: What is a destination rule in Istio?

**Answer:**
Policy applied *after* routing (LoadBalancing, CircuitBreaking, TLS settings).

---

### Question 979: How do you implement rate limiting in a service mesh?

**Answer:**
Envoy RateLimit filter (Global or Local).

---

### Question 980: What is sidecar injection and how is it controlled?

**Answer:**
Mutating Webhook.
- Enabled via label `istio-injection=enabled` on namespace.

---

## 🔹 Debugging Complex Issues (Questions 981-990)

### Question 981: How do you debug a stuck terminating pod?

**Answer:**
Check finalizers.

---

### Question 982: How do you handle OOMKilled errors?

**Answer:**
Increase Memory Limit. Check for leaks.

---

### Question 983: What is the `kubectl debug` command used for?

**Answer:**
Spins up Ephemeral Container attached to target pod.
- Useful for distroless images (where you can't `exec` sh).

---

### Question 984: How do you investigate container image pull failures?

**Answer:**
Events + Secrets check.

---

### Question 985: How do you analyze network policy blocks?

**Answer:**
Cilium logs / Network Policy Editor.

---

### Question 986: What are signs of etcd performance degradation?

**Answer:**
Slow API, Leader elections.

---

### Question 987: How do you troubleshoot webhook timeouts?

**Answer:**
Check if Webhook Service is reachable.
- If control plane cannot reach webhook pod (Firewall?), all creates fail.

---

### Question 988: What causes long scheduling delays?

**Answer:**
Resource crunch or Scheduler throughput limits.

---

### Question 989: How do you find pod restart reasons historically?

**Answer:**
Kube-state-metrics / Prom. Events are ephemeral (1 hour).

---

### Question 990: How do you troubleshoot DNS resolution failures?

**Answer:**
`dnsutils` pod, `nslookup`, `dig`.

---

## 🔹 Miscellaneous Real-World & Best Practices (Questions 991-1000)

### Question 991: What is the role of PodSecurity admission?

**Answer:**
Security Baseline.

---

### Question 992: How do you do rolling updates in production safely?

**Answer:**
Canary / PDBs / Readiness Gates.

---

### Question 993: How do you handle secrets at scale?

**Answer:**
Vault.

---

### Question 994: How do you ensure base image hygiene in builds?

**Answer:**
Rebuild frequently.

---

### Question 995: What’s a good readiness probe pattern for APIs?

**Answer:**
Check local dependencies (Config loaded?). DO NOT check external dependencies (DB) to prevent cascading failure.

---

### Question 996: How do you secure etcd traffic?

**Answer:**
mTLS peers.

---

### Question 997: What are common K8s misconfigurations in production?

**Answer:**
Default namespace usage, No limits, No pdb, Privileged pods.

---

### Question 998: How do you manage thousands of Helm releases?

**Answer:**
GitOps (ArgoCD ApplicationSets).

---

### Question 999: What are some anti-patterns in microservices on Kubernetes?

**Answer:**
Distributed Monolith (tight coupling), ignoring partial failure.

---

### Question 1000: How do you manage app versioning in multi-team clusters?

**Answer:**
Semantic Versioning + Helm Charts + GitOps.

---
