## 🔹 GitOps & CI/CD in Kubernetes (Questions 401-410)

### Question 401: What is GitOps in the context of Kubernetes?

**Answer:**
A set of practices where the entire system state is described in Git.
- An **Operator** (ArgoCD/Flux) in the cluster pulls changes from Git and applies them.
- **Rollback:** Revert commit in Git.
- **Audit:** Git history shows who changed what.

---

### Question 402: How does ArgoCD work?

**Answer:**
- It is a Kubernetes controller.
- It compares **Live State** (etcd) vs **Target State** (Git).
- If distinct, it marks the application `OutOfSync`.
- **Sync:** Applies the Git version to the cluster.

---

### Question 403: What is FluxCD?

**Answer:**
A GitOps toolset (Flux v2).
- **Source Controller:** Fetches artifacts from Git/Helm Repos.
- **Kustomize Controller:** Applies YAMLs.
- **Helm Controller:** Manages Helm releases.
- "Pull based" CD for Kubernetes.

---

### Question 404: How does GitOps differ from traditional CI/CD?

**Answer:**
- **Traditional (Push):** CI Server (Jenkins) has `kubectl` access strings. Runs `script -> kubectl apply`.
  - *Risk:* Security (CI has admin keys). Config Drift (Manual changes overwrite CI).
- **GitOps (Pull):** Cluster pulls from Git.
  - *Benefit:* Cluster doesn't expose credentials. Self-Corrects drift automatically.

---

### Question 405: What are the benefits of GitOps for K8s environments?

**Answer:**
1.  **Security:** No admin keys outside cluster.
2.  **Reliability:** Single source of truth.
3.  **Velocity:** Developers just use Git.
4.  **Compliance:** All changes are audited commits.

---

### Question 406: What are health checks in ArgoCD?

**Answer:**
ArgoCD assesses if an app is "Healthy".
- **Built-in:** Knows valid states for Deployment (AvailableReplicas == Replicas), PVC (Bound).
- **Custom:** You can write Lua scripts to define health for your CRDs (e.g., "MyCustomDb is healthy if status=running").

---

### Question 407: How do you manage secrets securely with GitOps?

**Answer:**
Since Git is public/shared, you can't put raw secrets there.
- **Approach 1:** Bitnami Sealed Secrets (Encrypted in Git).
- **Approach 2:** External Secrets Operator (Reference in Git: `secretStoreRef: aws-secrets`).

---

### Question 408: How do you implement multi-environment GitOps?

**Answer:**
**Structure:**
- `/base`: Common YAMLs.
- `/overlays/dev`: Patches for Dev.
- `/overlays/prod`: Patches for Prod.
- **Argo App 1:** Points to `dev`.
- **Argo App 2:** Points to `prod`.

---

### Question 409: What is the sync policy in ArgoCD?

**Answer:**
- **Manual:** User clicks "Sync" in UI.
- **Automatic:**
  - `selfHeal`: If someone runs `kubectl delete`, Argo recreates it immediately.
  - `prune`: If file deleted in Git, delete resource in K8s.

---

### Question 410: How do you track drift between Git and cluster state?

**Answer:**
GitOps tools monitor this continuously.
- ArgoCD UI shows a **Diff** view (Git says "Image:v2", Cluster has "Image:v1").
- Sending alerts on "OutOfSync" is a common pattern.

---

## 🔹 Chaos Engineering & Resilience (Questions 411-420)

### Question 411: What is chaos engineering in Kubernetes?

**Answer:**
The discipline of experimenting on a system to build confidence in its capability to withstand turbulent conditions.
- Intentionally breaking things (Deleting pods, killing nodes) to verify redundancy works.
- "Break it before production does."

---

### Question 412: How does LitmusChaos work?

**Answer:**
A Cloud-Native Chaos Engineering framework.
- **ChaosOperator:** Manages lifecycle.
- **ChaosExperiments (CRDs):** Defines the attack (e.g., `pod-delete`).
- It runs a "Runner Pod" that injects the fault and checks "Probes" for recovery.

---

### Question 413: What are the most common chaos experiments for K8s?

**Answer:**
1.  **Pod Delete:** Randomly kill pods. (Test ReplicaSet).
2.  **Node Drain:** Simulate node loss.
3.  **Network Latency:** Add 500ms delay to service. (Test timeouts).
4.  **CPU Stress:** Burn 100% CPU. (Test HPA/Limits).

---

### Question 414: How do you simulate node failure?

**Answer:**
- **Soft:** `kubectl drain node`.
- **Hard:** Terminate the EC2 instance / VM via Cloud API.
- **Chaos Mesh:** `AWSChaos` / `GCPChaos` integration.

---

### Question 415: What is a pod-kill scenario?

**Answer:**
"If I kill the DB leader, does a new leader get elected with zero data loss?"
- Verifies StatefulSet configuration and Leader Election logic.

---

### Question 416: What metrics determine resiliency of a K8s app?

**Answer:**
- **MTTR:** Mean Time To Recovery (How fast did it come back?).
- **Error Rate:** Did 500s spike during the kill?
- **Availability:** Did the Service Endpoint remain reachable?

---

### Question 417: How can you test for cascading failure in Kubernetes?

**Answer:**
- Inject latency into Service B (Backend).
- Observe Service A (Frontend).
- If Service A consumes all threads/memory waiting for B -> **Cascading Failure**.
- **Fix:** Implement Circuit Breakers / Timeouts.

---

### Question 418: What is network partition chaos testing?

**Answer:**
Simulating a "Split Brain".
- Blocking traffic between Node Group A and Node Group B.
- Verifies quorum logic in databases (etcd/Consul/Zookeeper).

---

### Question 419: What is the role of probes in resilience?

**Answer:**
Probes are the **Auto-Healing trigger**.
- If chaos kills the app logic but process stays running (zombie), **Liveness Probe** detects it and restarts the pod.
- Crucial for recovery.

---

### Question 420: How do you test for auto-healing behavior?

**Answer:**
1.  Deploy App with Replicas=3.
2.  Script: `while true; do kubectl delete pod -l app=my-app --random; sleep 10; done`.
3.  Simultaneously run a load test (`hey` or `k6`).
4.  Measure success rate of requests.

---

## 🔹 Observability & Monitoring (Questions 421-430)

### Question 421: What is the role of Prometheus in Kubernetes?

**Answer:**
The de-facto standard for metrics.
- **Service Discovery:** It talks to K8s API to find all Pods/Nodes.
- **Scraping:** It pulls `/metrics` from them.
- **Storage:** Stores time-series data locally (or sends to Thanos/Cortrex).

---

### Question 422: How does kube-state-metrics help in monitoring?

**Answer:**
Prometheus scrapes the app. But who tells you about deployments?
**Kube-state-metrics** listens to API server and exposes metrics like:
- `kube_deployment_status_replicas_unavailable`
- `kube_pod_status_phase{phase="Pending"}`

---

### Question 423: What is Grafana used for in K8s environments?

**Answer:**
Visualization.
- It queries Prometheus.
- Displays dashboards: "Cluster Overview", "Node Performance", "Pod Status".

---

### Question 424: What is the difference between whitebox and blackbox monitoring?

**Answer:**
- **Whitebox:** Insights from *inside* (Logs, JVM Metrics, Stack Traces). "Why is it slow?"
- **Blackbox:** Insights from *outside* (HTTP Check, Ping). "Is it slow?"
  - **Blackbox Exporter:** Probes an endpoint `http://my-service/health` from outside.

---

### Question 425: How do you monitor node health?

**Answer:**
**Node Exporter.**
- A DaemonSet running on every node.
- Exposes OS metrics (CPU interrupts, Disk I/O, Load Average, Network Stats) to Prometheus.

---

### Question 426: What is the role of Alertmanager?

**Answer:**
Prometheus finds the problem -> Sends to AlertManager.
AlertManager:
1.  **Deduplicates:** (Don't send 100 emails for 100 failing pods).
2.  **Groups:** (Group by Cluster).
3.  **Route:** (Devs get Slack, Ops get PagerDuty).

---

### Question 427: How do you define custom metrics for HPA?

**Answer:**
1.  App exposes value (e.g., `queue_depth`).
2.  Prometheus scrapes it.
3.  Prometheus Adapter exposes it to K8s Custom Metrics API.
4.  HPA configured to scale on `pods/queue_depth`.

---

### Question 428: What is the difference between logs and metrics?

**Answer:**
- **Metrics:** Aggregatable numbers (Request Count, Memory Usage). Cheap to store. Good for "What/When".
- **Logs:** Discrete events (Text lines, Stack traces). Expensive to store. Good for "Why".

---

### Question 429: What is OpenTelemetry?

**Answer:**
A vendor-neutral standard for Traces, Metrics, and Logs.
- **SDK:** Libraries for code.
- **Collector:** Agent to receive data and export to Prometheus/Jaeger/Datadog.
- Replaces proprietary agents.

---

### Question 430: How do you collect traces in Kubernetes?

**Answer:**
**Distributed Tracing (Jaeger/Zipkin).**
- App instrumented with OTel SDK.
- Injects `Trace-ID` in HTTP headers.
- Services propagate headers.
- Spans sent to Jaeger Collector.
- Visualizes the full request journey across microservices.

---

## 🔹 Logging & Auditing (Questions 431-440)

### Question 431: What is Fluentd and how does it integrate with Kubernetes?

**Answer:**
A log routing and processing tool.
- Deployed as **DaemonSet**.
- Mounts `/var/log/containers` from host.
- Parses Docker/Containerd JSON logs.
- Adds metadata (Pod Name, Namespace).
- Forwards to backend (Elasticsearch).

---

### Question 432: How do you implement centralized logging?

**Answer:**
**The EFK Pattern:**
- **Nodes:** Fluentd (Collector).
- **Backend:** Elasticsearch (Storage).
- **Frontend:** Kibana (Viewer).
Allows searching logs from all pods in one place.

---

### Question 433: What is EFK/ELK stack?

**Answer:**
- **E**lasticsearch.
- **L**ogstash (Heavy) OR **F**luentd (Cloud Native).
- **K**ibana.

---

### Question 434: How can you collect logs from pods?

**Answer:**
If app writes to file (not stdout), use a **Sidecar container** (e.g., simple `tail -f file` or fluent-bit) to stream it to stdout or send it to the collector.
Best practice: Refactor app to write to stdout.

---

### Question 435: What is the audit log in Kubernetes?

**Answer:**
Records actions taken on the API.
- "User Alice deleted Pod-X at 10:00 AM".
- Configured via `--audit-policy-file` on API Server.
- Essential for forensic analysis.

---

### Question 436: How do you configure Kubernetes audit policies?

**Answer:**
YAML policy file.
- **None:** Don't log.
- **Metadata:** Log User, Timestamp, Resource (but not payload).
- **RequestResponse:** Log full body (Sensitive data warning!).
Example: "Log everything in `kube-system` at Metadata level."

---

### Question 437: How do you ensure logs are tamper-proof?

**Answer:**
- Stream them immediately to an external write-only storage (S3 with Object Lock or Splunk).
- Do not store logs locally on the node where an attacker could modify them.

---

### Question 438: How do you identify suspicious activity in logs?

**Answer:**
- Use **Falco** (Runtime Security) to generate alerts for suspicious logs.
- Search for "Unauthorized" (401/403) spikes in API Audit logs.
- Look for `exec` commands into pods.

---

### Question 439: How do you rotate and manage logs in Kubernetes?

**Answer:**
Docker/Containerd engine handles rotation of the local JSON files on the node.
- Config: `max-size: 10m`, `max-file: 3`.
- If not configured, disk fills up -> Node Failure.

---

### Question 440: What is Loki and how does it work with Grafana?

**Answer:**
PLG Stack (Promtail, Loki, Grafana).
- **Promtail:** Agent (DaemonSet) ships logs.
- **Loki:** Aggregates logs, indexes *only metadata* (labels).
- **Grafana:** Query logs using **LogQL** (e.g., `{app="nginx"} |= "error"`).

---

## 🔹 Storage Deep Dive (Questions 441-450)

### Question 441: What is CSI (Container Storage Interface)?

**Answer:**
Unified standard for block/file storage.
- Allowed K8s to move volume plugins **Out-of-Tree**.
- Storage vendors (NetApp, Pure, AWS) can ship updates independently of K8s releases.

---

### Question 442: How does dynamic volume provisioning work?

**Answer:**
1.  User creates PVC.
2.  CSI Controller watches PVC.
3.  Calls Provider API (e.g., `CreateVolume`).
4.  Creates PV object.
5.  Binds PV to PVC.

---

### Question 443: What are volume reclaim policies?

**Answer:**
`persistentVolumeReclaimPolicy`:
- **Retain:** Keep data after PVC deleted.
- **Delete:** Delete data (Cloud volume destroyed).
- **Recycle:** Scrub data (rm -rf).

---

### Question 444: What is the difference between Retain, Delete, and Recycle?

**Answer:**
Covered above. Key nuance:
- **Retain** requires manual admin intervention to make the PV reusable (remove `ClaimRef`).

---

### Question 445: What is ReadWriteOnce vs ReadWriteMany?

**Answer:**
**Access Modes:**
- **RWO:** Mounted by single Node (EBS, Azure Disk).
- **RWX:** Mounted by multiple Nodes simultaneously (NFS, EFS, GlusterFS).
- **ROX:** Read Only multiple nodes.

---

### Question 446: How do you mount a PVC across multiple pods?

**Answer:**
- If pods are on the **same node**, Standard RWO works (usually).
- If pods are on **different nodes**, you MUST use **ReadWriteMany (RWX)** capable storage (NFS).

---

### Question 447: How do StatefulSets work with persistent storage?

**Answer:**
Uses `volumeClaimTemplates`.
- Generates `pvc-0` for `pod-0`, `pvc-1` for `pod-1`.
- If `pod-0` restarts on another node, it re-attaches `pvc-0`.
- Ensures data locality and identity.

---

### Question 448: What are StorageClasses?

**Answer:**
The "Profile" for dynamic provisioning.
- `provisioner`: `ebs.csi.aws.com`.
- `parameters`: `type: io1`, `iops: 5000`.
- Users just ask for `storageClassName: fast`.

---

### Question 449: What’s the difference between ephemeral and persistent volumes?

**Answer:**
- **Ephemeral:** `emptyDir`, `configMap`, `secret`. Tie to Pod life.
- **Persistent:** `hostPath` (Node life), `nfs`, `ebs` (Independent life).

---

### Question 450: How do you backup data in PVCs?

**Answer:**
**VolumeSnapshots.**
- K8s CRD: `kind: VolumeSnapshot`.
- CSI Driver commands storage array to snapshot.
- Restore: Create PVC with `dataSource: name: my-snapshot`.
