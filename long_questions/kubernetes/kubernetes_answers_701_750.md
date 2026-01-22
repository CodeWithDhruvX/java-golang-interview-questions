## 🔹 Kubernetes Storage & CSI Drivers (Questions 701-710)

### Question 701: What is a Container Storage Interface (CSI)?

**Answer:**
A standard interface for exposing block and file storage systems to container orchestration systems.
- Before CSI, storage plugins were "in-tree".
- CSI allows third-party vendors (AWS, NetApp) to develop plugins independently.

---

### Question 702: How do you install a CSI driver in Kubernetes?

**Answer:**
Usually via **Helm** or **Manifests**.
- It installs a DaemonSet (node-plugin) on every node.
- It installs a Deployment (controller) for API interactions.
- Creates `CSIDriver` object.

---

### Question 703: What is dynamic volume provisioning?

**Answer:**
Automatic creation of storage capability.
- User creates `PersistentVolumeClaim`.
- StorageClass provisions the real volume (e.g., EBS) on demand.
- Binds them together.

---

### Question 704: What is the difference between ReadWriteOnce and ReadWriteMany?

**Answer:**
- **RWO:** Block storage (EBS/Disk). Can be mounted by one Node.
- **RWX:** File storage (NFS/EFS). Can be mounted by multiple Nodes.

---

### Question 705: How do you mount the same volume across multiple pods?

**Answer:**
If Pods are on the **same node**, Standard RWO works.
If Pods are on **different nodes**, you MUST use **RWX** (NFS/CephFS).

---

### Question 706: What is a VolumeSnapshot and how is it used?

**Answer:**
A CRD to trigger a storage-level snapshot.
- `kind: VolumeSnapshot`.
- Source: `persistentVolumeClaimName: my-pvc`.
- Used for Backups and Cloning.

---

### Question 707: How do you clone a PVC in Kubernetes?

**Answer:**
Create a new PVC with `dataSource`.
```yaml
spec:
  dataSource:
    name: original-pvc
    kind: PersistentVolumeClaim
```
The CSI driver creates a new volume populated with data from the original.

---

### Question 708: What’s the difference between ephemeral and persistent volumes?

**Answer:**
- **Ephemeral:** `emptyDir`, `configMap`. Data is lost when Pod dies.
- **Persistent:** `hostPath` (Node durable), `nfs/ebs` (Cluster durable). Data survives Pod restart vs Data survives Cluster restart.

---

### Question 709: How do you implement volume resizing?

**Answer:**
1.  Verify StorageClass `allowVolumeExpansion: true`.
2.  Edit PVC `request.storage: 10Gi` -> `20Gi`.
3.  CSI Driver expands backend volume.
4.  File system expands (online or on pod restart).

---

### Question 710: What is storageClass and how is it used?

**Answer:**
Defines the "Tier" of storage.
- `slow`: HDD.
- `fast`: NVMe.
- Includes `parameters` passed to the provisioner (type, iops, encryptionKey).

---

## 🔹 Logging, Monitoring & Observability (Questions 711-720)

### Question 711: What is the recommended logging architecture for Kubernetes?

**Answer:**
**Node-Level Logging Agent.**
- DaemonSet (Fluentd/FluentBit) on every node.
- Reads `/var/log/containers/*.log`.
- Enriches metadata.
- Sends to Central Storage (ES/Loki).

---

### Question 712: How does Fluent Bit differ from Fluentd?

**Answer:**
- **Fluentd:** Ruby/C. Feature rich, plugin heavy. 50MB+ RAM.
- **Fluent Bit:** C. Extremely lightweight (1MB). Designed for embedded/high-scale.
- Often use Bit as collector -> Forward to Fluentd aggregator.

---

### Question 713: What is Loki and how does it integrate with Promtail?

**Answer:**
Grafana's Log System.
- **Promtail:** The Agent (grep logs).
- **Loki:** The Server (Stores chunks).
- Integration: Promtail pushes to Loki HTTP API.

---

### Question 714: How do you collect application logs in Kubernetes?

**Answer:**
- **Stdout/Stderr:** Standard method. Kubelet handles rotation.
- **File:** If app writes to `/var/log/app.log`, use a Sidecar to `tail` it to stdout or ship it.

---

### Question 715: What is the ELK stack and how does it relate to K8s?

**Answer:**
**E**lasticsearch, **L**ogstash, **K**ibana.
- The "Classic" stack.
- In K8s, "Elastic Cloud on Kubernetes (ECK)" operator makes management easier.

---

### Question 716: How do you forward logs to external providers?

**Answer:**
Plugins in Fluentd/FluentBit.
- `output-cloudwatch`, `output-datadog`, `output-splunk`.

---

### Question 717: How does Prometheus service discovery work?

**Answer:**
Prometheus talks to API Server.
- `kubernetes_sd_configs`.
- Discovers `Node`, `Service`, `Pod`, `Ingress`, `Endpoint`.
- Automatically targets new pods as they scale up.

---

### Question 718: What are Prometheus relabeling rules?

**Answer:**
Powerful transformation logic during scraping.
- "Drop metric if pod name starts with `test-`".
- "Rename label `__meta_kubernetes_pod_name` to `pod`".

---

### Question 719: What is a blackbox exporter in Prometheus?

**Answer:**
Probes things from outside.
- Ping / HTTP Get / DNS Query.
- "Is google.com reachable from inside the cluster?"

---

### Question 720: How does OpenTelemetry enhance Kubernetes observability?

**Answer:**
Unified standard.
- Avoids vendor lock-in (Datadog agent vs NewRelic agent).
- Collects Traces, Metrics, Logs in one OTLP format and exports to any backend.

---

## 🔹 Helm Advanced Usage (Questions 721-730)

### Question 721: What is a Helm release lifecycle?

**Answer:**
1.  Install (Pending -> Deployed).
2.  Upgrade (Deployed -> Superseeded).
3.  Rollback.
4.  Uninstall.
- State stored in `Secret` (sh.helm.release.v1.my-app.v1).

---

### Question 722: How do you upgrade a Helm chart without downtime?

**Answer:**
Helm just updates K8s resources.
- If Deployment strategy is `RollingUpdate`, K8s handles the zero-downtime part.
- Helm acts as the trigger.

---

### Question 723: How do you manage secrets with Helm charts?

**Answer:**
- **Bad:** Plaintext in `values.yaml`.
- **Good:** Encrypted `secrets.yaml` (Helm Secrets plugin using SOPS).
- **Better:** Do not manage secrets in Helm at all. Use ExternalSecrets.

---

### Question 724: What are Helm hooks and how do they work?

**Answer:**
Special annotations on templates.
- `helm.sh/hook: pre-install, pre-upgrade`.
- Run a Job (DB Migration) *before* the Deployment is updated.
- If Hook fails, Upgrade fails.

---

### Question 725: What is the difference between `values.yaml`, `secrets.yaml`, and `Chart.yaml`?

**Answer:**
- **Chart.yaml:** Metadata (Name, Version, Dependencies).
- **values.yaml:** Default configuration.
- **secrets.yaml:** (Convention) Encrypted values.

---

### Question 726: How do you package and distribute a custom Helm chart?

**Answer:**
1.  `helm package my-chart` -> `.tgz` file.
2.  Upload to registry (ECR/Harbor) which supports OCI (Open Container Initiative).
3.  `helm install oci://registry/my-chart`.

---

### Question 727: What are subcharts in Helm?

**Answer:**
Dependencies.
- "WordPress Chart" depends on "MySQL Chart".
- `Chart.yaml`: `dependencies: - name: mysql`.
- `helm dependency update`.

---

### Question 728: What is `helm lint` and how is it used?

**Answer:**
Static analysis tool.
- Checks for formatting errors, missing values, invalid YAML.
- Run in CI before packaging.

---

### Question 729: How do you rollback a Helm release?

**Answer:**
`helm rollback <release> <revision>`.
- Reverts to the exact manifest state of that previous revision.

---

### Question 730: How do you test Helm charts using Helm unittest?

**Answer:**
Plugin (`helm-unittest`).
- Write YAML test cases.
- "Assert that if `values.replicaCount=3`, the generated Deployment has `replicas: 3`".
- Tests logic without needing a cluster.

---

## 🔹 Platform Engineering with Kubernetes (Questions 731-740)

### Question 731: What is an Internal Developer Platform (IDP)?

**Answer:**
Self-service layer on top of K8s.
- "I need a Database + API".
- IDP provisions AWS RDS + EKS Namespace + CI/CD.
- Developer doesn't see kubectl complexity.

---

### Question 732: How does Backstage integrate with Kubernetes?

**Answer:**
Spotify's Developer Portal.
- **Plugin:** Shows Pod status inside the Backstage UI.
- Developers see "Deployment Status: Healthy" alongside their Jira tickets and Docs.

---

### Question 733: What is score.dev and how does it help platform teams?

**Answer:**
Workload specification standard.
- Developers write a platform-agnostic `score.yaml`.
- Tool translates it to `helm` or `docker-compose`.

---

### Question 734: What’s the role of a Kubernetes platform team?

**Answer:**
- **Curate:** Pick the best tools (Argo, Istio).
- **Abstract:** Hide complexity.
- **Operate:** Manage the cluster upgrades/security so App teams focus on code.

---

### Question 735: How do you define golden paths for developers?

**Answer:**
Standardized templates.
- "Spring Boot Service Template": Includes Dockerfile, Helm Chart, CI Pipeline, Prom Metrics.
- Developer just forks and codes.

---

### Question 736: What tools help automate environment creation?

**Answer:**
**vCluster / Crossplane / Terraform.**
- "Click button -> Get Ephemeral Namespace" via GitHub Actions.

---

### Question 737: What is a self-service deployment portal?

**Answer:**
Port / Backstage.
- Catalog of "Actions".
- Action: "Scaffold new Microservice".
- Action: "Provision S3 Bucket".

---

### Question 738: How do you secure multi-tenant platform services?

**Answer:**
- **OIDC:** Everyone logs in securely.
- **RBAC:** Team A cannot see Team B resources.
- **Quotas:** prevents resource hogging.

---

### Question 739: What is GitHub Copilot for Kubernetes manifests?

**Answer:**
AI assistance.
- "Write a deployment for redis with pvc".
- Generates 90% accurate YAML.

---

### Question 740: What is Kratix and how is it used?

**Answer:**
Framework for building Platforms.
- "Promise": Definition of a service (e.g., PostgreSQL).
- "Pipeline": Workflow to fulfill the promise (Terraform -> Helm).

---

## 🔹 CNCF Ecosystem Awareness (Questions 741-750)

### Question 741: What is the CNCF Landscape?

**Answer:**
The interactive map of 1000+ cloud-native projects.
- Grouped by category (Orchestration, Observability, Storage).

---

### Question 742: How do tools like ArgoCD and Flux differ?

**Answer:**
- **ArgoCD:** GUI-centric, Application abstraction, Push from controller.
- **Flux:** CLI/File-centric, Controller per functionality (Source/Kustomize/Helm), more native "GitOps" feel.

---

### Question 743: What is Dapr and how does it work with Kubernetes?

**Answer:**
**Distributed Application Runtime.**
- Sidecar pattern.
- Abstracts "State Store", "PubSub", "Secrets".
- App calls `localhost:3500/v1.0/state/redis`. Dapr talks to Redis.

---

### Question 744: What is OpenFunction?

**Answer:**
FaaS (Function as a Service) on Kubernetes.
- Influenced by Knative.
- Supports Dapr bindings.

---

### Question 745: What is LitmusChaos used for?

**Answer:**
Chaos Engineering.
- "ChaosCenter" UI.
- "ChaosHub" Catalog of experiments (Pod Delete, Node Drain).

---

### Question 746: What is Keptn and how does it enable SLO-based delivery?

**Answer:**
Event-driven Control Plane.
- **Quality Gates:** "Prometheus, is error rate < 1%?". If No, deny promotion.
- Automated validation.

---

### Question 747: What is OpenKruise?

**Answer:**
Advanced workload management (Alibaba).
- `CloneSet`: Better than StatefulSet (In-place update).
- `AdvancedDaemonSet`: Canary updates for DaemonSets.

---

### Question 748: What is Harbor and how does it help?

**Answer:**
Open Source Registry.
- Stores Images, Helm Charts.
- Scans for CVEs (Trivy/Clair).
- Replicates images between registries.

---

### Question 749: What is Crossplane’s role in infrastructure-as-code?

**Answer:**
Drift Correction.
- Terraform runs once.
- Crossplane watches forever. If someone changes AWS Security Group manually, Crossplane reverts it.

---

### Question 750: What is KEDA (Kubernetes Event-Driven Autoscaling)?

**Answer:**
Scales workloads based on events.
- Replaces HPA for non-CPU metrics.
- SQS Messages > 1000 -> Scale up to 50 pods.
