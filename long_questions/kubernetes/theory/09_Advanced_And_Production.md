# 🔵 Advanced Topics, StatefulSets, Operators & Production

---

### 1. What is a StatefulSet?

"A StatefulSet manages **stateful applications** where pod identity, ordering, and persistent storage matter.

Unlike Deployments where pods are interchangeable, StatefulSet pods have:
- A **stable, unique hostname**: `db-0`, `db-1`, `db-2`
- **Sticky storage**: `db-0` always mounts the same PVC
- **Ordered startup/shutdown**: pods are created/deleted in order (0 → 1 → 2)

I use StatefulSets for Postgres, Kafka, Cassandra, Redis Sentinel, Elasticsearch — any distributed system that needs stable identity."

#### In Depth
The stable hostname format is `<pod-name>.<headless-service>.<namespace>.svc.cluster.local`. This FQDN is stable across pod restarts, even if the pod moves to a different node. The headless Service (ClusterIP: None) is a prerequisite for StatefulSets — it enables DNS-based discovery of individual pods. The `podManagementPolicy: Parallel` can be set to skip ordering when the application doesn't require sequential startup.

---

### 2. What are the alternatives to StatefulSets?

"For most distributed databases, there are now **Kubernetes Operators** that abstract away the StatefulSet complexity.

- Instead of manually managing a PostgreSQL StatefulSet, I use the **CloudNativePG operator** or **Zalando Postgres Operator**.
- For Kafka, I use the **Strimzi operator**.
- For Elasticsearch, the **ECK operator** from Elastic.

Operators encode operational knowledge (failover, backup, scaling) into code. They're the right abstraction for production databases — raw StatefulSets require you to build all that operational logic yourself."

#### In Depth
The Operator pattern (a controller watching a custom CRD) was coined by CoreOS in 2016. The **Operator Capability Model** defines 5 maturity levels from basic install to auto-pilot. Level 4-5 operators handle backup/restore, upgrades, and auto-remediation. Before choosing a raw StatefulSet for a database, check operatorhub.io — there's likely a mature operator that handles the hard parts.

---

### 3. What is a Job and CronJob?

"A **Job** runs a pod until it completes successfully. It handles retries and tracks completion status.

A **CronJob** creates Jobs on a cron schedule. It's Kubernetes' native way to run recurring batch tasks.

I use Jobs for database migrations (run once at deploy time), data processing pipelines (triggered by events), and CronJobs for nightly report generation, log cleanup, and health checks."

#### In Depth
CronJob's `concurrencyPolicy`: `Allow` (multiple runs can run concurrently), `Forbid` (skip if previous run is still going), `Replace` (kill the running job, start new). In production, I use `Forbid` for idempotent jobs to prevent duplicate runs. The `successfulJobsHistoryLimit` and `failedJobsHistoryLimit` prevent old Job resources from accumulating. `startingDeadlineSeconds` defines the window after which a missed schedule is skipped rather than run late.

---

### 4. What are DaemonSets used for?

"A DaemonSet ensures a pod runs on **every node** (or a selected subset) in the cluster.

Core use cases:
- **Log collection**: Fluent Bit reads `/var/log/containers` from each node
- **Monitoring**: Prometheus Node Exporter reports node metrics
- **Security**: Falco watches syscalls on every node
- **Networking**: CNI plugins and kube-proxy are DaemonSets

I also use node-problem-detector (a DaemonSet) to surface hardware issues (disk I/O errors, NTP failures) as K8s node conditions that can trigger alerts."

#### In Depth
DaemonSets by default don't run on control plane nodes due to their `NoSchedule` taint. For DaemonSets that need to run everywhere (like network plugins), add the toleration: `{key: node-role.kubernetes.io/control-plane, effect: NoSchedule}`. DaemonSet updates use rolling update strategy — one node at a time. Use `maxSurge` and `maxUnavailable` to control the rollout speed.

---

### 5. What is taint-based eviction?

"The `NoExecute` taint effect causes pods **already running on a node** to be evicted.

Kubernetes uses this automatically: when a node becomes `NotReady`, K8s applies `node.kubernetes.io/not-ready:NoExecute` taint. Pods without a matching toleration are evicted after `tolerationSeconds` (default 300s).

I customize this based on workload criticality: critical monitoring pods get `tolerationSeconds: 600` (tolerate longer outages), regular apps get 0 (evict immediately for faster recovery on available nodes)."

#### In Depth
The `node.kubernetes.io/unreachable:NoExecute` taint is applied when the node controller loses contact with a node. Both `not-ready` and `unreachable` taints are added automatically since K8s 1.18. This is the mechanism behind Kubernetes' self-healing — pods are rescheduled elsewhere when nodes die. The eviction delay prevents unnecessary rescheduling during brief network blips.

---

### 6. What is a Custom Resource Definition (CRD)?

"A CRD extends the Kubernetes API with **your own resource types**.

Instead of inventing your own API server, you define a CRD (like defining a new schema), and kubectl and the API server treat it like any built-in resource: you can `kubectl get`, `kubectl apply`, `kubectl delete` it.

The power is combining CRDs with controllers — the controller watches these custom resources and acts on them. This is the foundation of every Kubernetes operator."

#### In Depth
CRDs are defined with an OpenAPI v3 schema for validation — this gives you type checking and required field enforcement before resources are persisted to etcd. Versioning CRDs is complex: you can have multiple versions (v1alpha1, v1beta1, v1) with **conversion webhooks** that translate between them. The `status` subresource should be separate from `spec` — this allows controllers to update status without triggering a reconciliation loop.

---

### 7. What is a Kubernetes Operator?

"An Operator is a **controller + CRD pair** that encodes operational knowledge about a specific application.

Pattern: You define a CRD like `PostgresCluster` with fields like `replicas`, `storageSize`, `backupSchedule`. The operator watches these resources and manages the underlying StatefulSets, Services, ConfigMaps, and backup Jobs to match the desired state.

The operator knows how to: create a primary and replicas, set up replication, run backups, perform failover, and handle upgrades. All you do is apply a YAML file."

#### In Depth
Writing an operator from scratch: use **controller-runtime** (Go) or **Java Operator SDK**. The Kubebuilder framework scaffolds the boilerplate. The core loop is: `Reconcile(ctx, req)` — it retrieves the current state from the API and takes actions to bring it toward the desired state. Critically, the reconcile loop must be **idempotent** — it may be called many times (errors, crashes, API watches) and must produce the same result.

---

### 8. What is a PodDisruptionBudget?

"A PodDisruptionBudget (PDB) defines **how many pods of a set can be unavailable simultaneously** during voluntary disruptions.

Example: `minAvailable: 2` for a 3-replica service means at most 1 pod can be down at a time — a node drain won't proceed if it would violate this.

I set PDBs for every production service. Without them, a `kubectl drain` for node maintenance could take down your entire service in one shot."

#### In Depth
PDB applies to **voluntary disruptions** (node drain, rolling updates, admission control evictions) — not involuntary ones (node crash, OOMKill). The `max-unavailable: 50%` alternative specifies that no more than 50% of pods can be unavailable. PDBs interact with the Cluster Autoscaler — a node with pods protected by a PDB that can't be evicted will block CA from deleting that node. This is a common reason CA gets stuck.

---

### 9. What is admission control in Kubernetes?

"Admission controllers are **plugins that intercept API requests** after authentication/authorization but before persistence to etcd.

They can **mutate** (change the request, e.g., inject a sidecar) or **validate** (reject invalid requests, e.g., block privileged pods).

Built-in admission controllers: `NamespaceLifecycle`, `LimitRanger`, `ServiceAccount`, `ResourceQuota`, `PodSecurity`. Extended via: `MutatingAdmissionWebhook` and `ValidatingAdmissionWebhook` (external HTTP servers called by the API server)."

#### In Depth
The admission control flow: API request → Authentication → Authorization → Mutating webhooks → Object schema validation → Validating webhooks → Persisted to etcd. Mutating webhooks run first (they can add fields), then validation runs on the mutated object. If a webhook is unavailable and `failurePolicy: Fail`, the API request is rejected. Use `failurePolicy: Ignore` only for non-critical webhooks. Always set `timeoutSeconds` (max 30s) — a slow webhook can cause API server latency spikes.

---

### 10. What is a headless service?

"A headless Service (`spec.clusterIP: None`) returns **individual pod IPs** from DNS instead of a virtual IP.

When you query the DNS name of a headless service, you get an A record per ready pod. Clients do their own load balancing.

Critical use cases: StatefulSet pod addressing (each pod has its own stable DNS entry), gRPC client-side load balancing (gRPC maintains persistent connections and distributes calls itself), and service meshes that need direct pod connectivity."

#### In Depth
With a headless service, CoreDNS returns a round-robin list of pod IPs. The ordering is random per query (DNS shuffling). For StatefulSets, the DNS name `<pod-name>.<headless-service>` resolves to a specific pod's IP — this is what enables stable peer discovery in distributed systems. The headless service must have `spec.selector` matching the StatefulSet pods for the DNS records to be created.

---

### 11. What is a sidecar container?

"A sidecar runs in the **same Pod as the main application** but handles a cross-cutting concern — logging, proxying, metrics, configuration updates.

Classic examples:
- **Envoy sidecar** (in Istio): Handles all network traffic in/out, provides mTLS, circuit breaking, metrics.
- **Fluent Bit sidecar**: Reads app log files and ships to Elasticsearch.
- **Config reloader sidecar**: Watches ConfigMap changes and sends SIGHUP to the main app.

Sidecars share the pod's network namespace and volumes, but have their own process and resource limits."

#### In Depth
K8s 1.28+ introduced **sidecar containers as a first-class concept** via `initContainers` with `restartPolicy: Always`. Previously, sidecars were just regular containers in the pod with no lifecycle guarantees. The new sidecar feature ensures sidecars start before the main container and are terminated after it — solving the long-standing issue where Fluent Bit or Envoy would exit before the main app's preStop hook finished draining.

---

### 12. What is a multi-cluster Kubernetes strategy?

"Multi-cluster means running **multiple separate K8s clusters** for isolation, geographic distribution, or different security domains.

Common reasons: regulatory compliance (data sovereignty — EU data must stay in EU), blast radius reduction (an outage in cluster A doesn't affect cluster B), environment isolation (production and dev/staging in completely separate clusters).

Tools for multi-cluster management: ArgoCD multi-cluster support, Flux's cluster API integration, Google Anthos, Rancher, and the CNCF's Cluster API (CAPI)."

#### In Depth
Multi-cluster adds complexity: cross-cluster service discovery, consistent policy enforcement, unified observability, and cross-cluster secret management. **Submariner** provides L3 network connectivity between clusters. **Istio multi-cluster** extends the service mesh across clusters. **KubeFed** (Federation v2) can sync resources across clusters. Despite the tools, multi-cluster is operationally heavy — only add this complexity when you have a clear requirement.

---

### 13. What is progressive delivery?

"Progressive delivery is the practice of **gradually shifting traffic to a new version**, monitoring it, and automating promotion or rollback based on metrics.

It extends canary deployments with automation:
1. Deploy new version to 5% of traffic
2. Measure error rate, latency, business metrics automatically
3. If metrics pass thresholds, promote to 20% → 50% → 100%
4. If metrics fail, auto-rollback to 0%

**Argo Rollouts** is the K8s-native tool for this, integrating with Prometheus for metric analysis."

#### In Depth
Argo Rollouts introduces `Rollout` CRD (replaces Deployment) and `AnalysisTemplate` (defines the metrics to check and thresholds). The `analysisRun` automatically queries Prometheus during each canary step and makes the promotion/rollback decision. Argo Rollouts integrates with Istio, AWS ALB, NGINX, and Contour for precise traffic shifting — not the coarse pod-count-based approach of native K8s.

---
