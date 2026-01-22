## 🔹 Security (Questions 51-55)

### Question 51: What is RBAC?

**Answer:**
**Role-Based Access Control (RBAC)** is a method of regulating access to computer or network resources based on the roles of individual users within an enterprise.
- In K8s, it restricts **who** (User/ServiceAccount) can do **what** (Verbs: get, create, delete) to **which** (Resources: pods, secrets).
- **Core Objects:** Role, ClusterRole, RoleBinding, ClusterRoleBinding.

---

### Question 52: What are ServiceAccounts?

**Answer:**
A ServiceAccount provides an identity for **processes that run in a Pod**.
- Users have normal accounts (Usernames).
- Pods have ServiceAccounts.
- When a Pod accesses the API Server (e.g., a custom controller), it authenticates using the mounted ServiceAccount token.

---

### Question 53: How does Kubernetes handle authentication?

**Answer:**
Kubernetes does not have an internal database of users (like IAM). It relies on external auth plugins.
**Strategies:**
1.  **X509 Client Certs:** Used by admins/kubelet.
2.  **Bearer Tokens:** ServiceAccounts.
3.  **OIDC (OpenID Connect):** Connects to Google, Keycloak, Okta.
4.  **Webhook Token Auth:** Validate via external HTTP service.

---

### Question 54: What are NetworkPolicies used for?

**Answer:**
NetworkPolicies specify how groups of Pods are allowed to communicate with each other and other network endpoints.
- **Default:** All traffic allowed.
- **Policy:** "Deny All" or "Allow specific Ingress/Egress".
- REQUIRED: A network plugin (CNI) that enforces this (e.g., Calico, Weave). Flannel does NOT support it.

---

### Question 55: How do you restrict access to the API server?

**Answer:**
1.  **Authentication:** Ensure strong auth (Certificates/OIDC).
2.  **RBAC:** Authorize with least privilege.
3.  **Network:** Configure firewall/security groups to allow access to port 6443 only from trusted IPs (Bastion host/VPN).
4.  **Private Endpoint:** In cloud (EKS/GKE), disable public endpoint and access via private VPC peering.

---

## 🔹 Helm & Operators (Questions 56-60)

### Question 56: What is Helm?

**Answer:**
Helm is the **Package Manager** for Kubernetes (like `apt` or `npm`).
- It manages Kubernetes applications.
- Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.

---

### Question 57: What are Helm Charts?

**Answer:**
A Helm Chart is a collection of files that describe a related set of Kubernetes resources.
- **Structure:**
  - `Chart.yaml`: Meta-data.
  - `values.yaml`: Default configuration values.
  - `templates/`: Manifest files with placeholders (`{{ .Values.image }}`).

---

### Question 58: How do you install applications using Helm?

**Answer:**
```bash
# Add repo
helm repo add bitnami https://charts.bitnami.com/bitnami

# Install chart
helm install my-release bitnami/nginx

# Install with custom values
helm install my-db bitnami/mysql --set auth.rootPassword=secret
```

---

### Question 59: What is a Helm repository?

**Answer:**
An HTTP server that houses an `index.yaml` file and packaged charts (`.tgz`).
- Examples: Artifact Hub, Bitnami, internal Harbor registry.

---

### Question 60: What is a Kubernetes Operator?

**Answer:**
An Operator is a software extension to Kubernetes that makes use of **Custom Resources** to manage applications and their components.
- It encodes human operational knowledge (how to backup, how to upgrade, how to failover) into software.
- **Example:** Prometheus Operator, Postgres Operator.

---

## 🔹 Monitoring & Logging (Questions 61-65)

### Question 61: How do you monitor a Kubernetes cluster?

**Answer:**
You need to monitor:
1.  **Node Level:** CPU/RAM/Disk (Node Exporter).
2.  **K8s Level:** Pod status, Deployment health (Kube-state-metrics).
3.  **App Level:** HTTP latency, errors (Instrumentation).
**Standard Stack:** Prometheus + Grafana.

---

### Question 62: What is Prometheus?

**Answer:**
An open-source systems monitoring and alerting toolkit.
- **Pull Model:** It scrapes metrics from HTTP endpoints (`/metrics`) exposed by services.
- **TSDB:** Time Series Database.
- **PromQL:** Powerful query language.

---

### Question 63: What is Grafana?

**Answer:**
A visualization tool used to create dashboards from data stored in Prometheus (and other sources).
- It provides the "Glass" to look at the metrics.

---

### Question 64: How do you collect logs in Kubernetes?

**Answer:**
Kubernetes does not provide a native storage solution for log data.
**Pattern (Sidecar or DaemonSet):**
1.  App writes to `stdout/stderr`.
2.  Docker/Runtime writes to `/var/log/containers/*.log`.
3.  **Log Agent (Fleet/Fluentd) run as DaemonSet** reads these files.
4.  Pushes logs to centralized backend (Elasticsearch/Splunk/Loki).

---

### Question 65: What is Fluentd?

**Answer:**
An open-source data collector for unified logging layer.
- Actively parses, filters, buffers, and outputs logs.
- Commonly used in **EFK Stack** (Elasticsearch, Fluentd, Kibana).

---

## 🔹 Autoscaling (Questions 66-70)

### Question 66: What is Horizontal Pod Autoscaler (HPA)?

**Answer:**
Automatically scales the number of Pods in a replication controller, deployment, replica set, or stateful set based on observed CPU utilization (or custom metrics).
- **Scale Out:** Adds pods when load increases.
- **Scale In:** Removes pods when load decreases.

---

### Question 67: What is Vertical Pod Autoscaler (VPA)?

**Answer:**
Automatically sets Container resource requests and limits based on usage.
- It restarts pods with new CPU/Memory configurations.
- Helpful for rightsizing workloads.

---

### Question 68: What is Cluster Autoscaler?

**Answer:**
Automatically adjusts the size of the Kubernetes cluster (Nodes).
- **Scale Up:** If there are pods pending because of insufficient resources.
- **Scale Down:** If nodes have been underutilized for an extended period and pods can be placed elsewhere.
- Works with Cloud Providers (AWS ASG, GKE Node pools).

---

### Question 69: How does HPA work internally?

**Answer:**
1.  HPA queries the **Metrics Server** every 15 seconds (default).
2.  Compares usage (e.g., current CPU 80%) vs Target (e.g., target CPU 50%).
3.  Calculates new replica count: `TargetNum = CurrentReplicas * (CurrentMetric / TargetMetric)`.
4.  Updates the `DesiredReplicas` in the Deployment/ReplicaSet.

---

### Question 70: Can you autoscale based on custom metrics?

**Answer:**
Yes.
- The standard Metrics Server only provides CPU/Memory.
- For business metrics (Queue length, Requests per sec), you need **Prometheus Adapter**.
- It exposes custom metrics to the K8s Custom Metrics API, which HPA can consume.

---

## 🔹 Troubleshooting (Questions 71-75)

### Question 71: How do you troubleshoot a CrashLoopBackOff error?

**Answer:**
Indicates the pod starts, crashes, starts again, and loops.
1.  **Check Logs:** `kubectl logs <pod>`. Look for app panic/exit errors.
2.  **Describe Pod:** `kubectl describe pod <pod>`. Check "LastState", exit code (137=OOM, 1=Error).
3.  **Previous Logs:** `kubectl logs <pod> --previous`.
4.  **Config:** Check for missing env vars or config maps.

---

### Question 72: What command do you use to view Pod logs?

**Answer:**
```bash
# Basic
kubectl logs my-pod

# If pod has multiple containers
kubectl logs my-pod -c my-container

# Follow live
kubectl logs -f my-pod

# View previous instance logs (after restart)
kubectl logs -p my-pod
```

---

### Question 73: How do you debug a Kubernetes node issue?

**Answer:**
1.  **Check Node Status:** `kubectl get nodes`.
2.  **Describe Node:** `kubectl describe node <node>`. Look for `Conditions` (DiskPressure, MemoryPressure).
3.  **Check Kubelet Logs:** SSH into node and run `journalctl -u kubelet`.
4.  **Check Disk Space:** `df -h`.

---

### Question 74: What are common reasons for Pending Pods?

**Answer:**
Pending means it hasn't been scheduled to a node yet.
1.  **Insufficient Resources:** No node has enough CPU/RAM to meet user Request.
2.  **Taints/Tolerations:** Pod cannot tolerate node taints.
3.  **Affinity:** Constraints prevent scheduling.
4.  **PVC Pending:** Waiting for storage volume provisioning.

---

### Question 75: How to check resource usage per Pod?

**Answer:**
Requires **Metrics Server** installed.
```bash
kubectl top pod
kubectl top nodes
```

---

## 🔹 CI/CD with Kubernetes (Questions 76-80)

### Question 76: How do you implement CI/CD in Kubernetes?

**Answer:**
**CI (Continuous Integration):**
- Developer pushes code.
- CI Server (Jenkins/GitLab) builds Docker image.
- Runs tests.
- Pushes image to Registry.
**CD (Continuous Delivery):**
- CD Tool updates the K8s Deployment manifest with new image tag.
- Applies manifest to cluster (`kubectl apply` or Helm).

---

### Question 77: What tools integrate well with Kubernetes for CI/CD?

**Answer:**
- **Jenkins:** (Classic, flexible).
- **GitLab CI:** (Integrated container registry + Runner).
- **Tekton:** (Kubernetes-native CI pipelines).
- **ArgoCD / Flux:** (GitOps CD).
- **Spinnaker:** (Complex, multi-cloud).

---

### Question 78: What is ArgoCD?

**Answer:**
A declarative, GitOps continuous delivery tool for Kubernetes.
- Runs **inside** the cluster.
- Watches a Git repo for manifest changes.
- Automatically synchronizes the cluster state with Git state.
- Provides a UI for visualization.

---

### Question 79: What is FluxCD?

**Answer:**
Another GitOps tool created by Weaveworks.
- Works similarly to Argo: ensures that the state of a cluster matches the config in git.
- Flux v2 is based on the GitOps Toolkit components.

---

### Question 80: How does GitOps work in Kubernetes?

**Answer:**
**GitOps** = Infrastructure as Code + Pull Requests + CI/CD.
1.  **Source of Truth:** Git Repo contains K8s YAMLs.
2.  **Agent:** A controller (Argo/Flux) inside K8s pulls changes from Git.
3.  **Sync:** The agent applies the changes.
**Benefit:** No manual `kubectl apply`. Audit trail in Git history. Easy rollback (`git revert`).

---

## 🔹 Advanced Concepts (Questions 81-90)

### Question 81: What is a StatefulSet?

**Answer:**
Workload API object used to manage stateful applications (DBs like Cassandra, Mongo, ZooKeeper).
**Features:**
- **Stable, unique network identifiers:** `web-0`, `web-1` (not random hash).
- **Stable, persistent storage:** Each pod gets its own PVC that persists across rescheduling (sticky identity).
- **Ordered deployment/scaling:** 0 -> 1 -> 2.

---

### Question 82: What is a Job and CronJob?

**Answer:**
- **Job:** Creates one or more Pods and ensures that a specified number of them successfully terminate (Batch task, Database migration).
- **CronJob:** Creates Jobs on a repeating schedule (like Linux Cron). `0 23 * * *`.

---

### Question 83: What is taint and toleration?

**Answer:**
Mechanism to prevent pods from being scheduled onto inappropriate nodes.
- **Taint (Node):** "I have a GPU."
- **Toleration (Pod):** "I need a GPU."
- A pod is scheduled on a tainted node ONLY if it has a matching toleration.
- **Effect:** NoSchedule, PreferNoSchedule, NoExecute.

---

### Question 84: What is affinity and anti-affinity?

**Answer:**
Advancement over NodeSelector.
- **Node Affinity:** "Prefer running on High-Mem nodes" (Soft/Hard rules).
- **Pod Affinity:** "Run this App Pod near the Cache Pod" (Co-location).
- **Pod Anti-Affinity:** "Do NOT run two App Pods on same node" (High Availability / Spread).

---

### Question 85: What is a sidecar container?

**Answer:**
A utility container running in the **same Pod** as the main application container.
- **Role:** Enhances or helps main app without changing its code.
- **Examples:** Log shipper (reads log file, sends to Splunk), Proxy (Istio Envoy), Config refresher.

---

### Question 86: What is a headless service?

**Answer:**
A service that does **not** allocate a ClusterIP (`clusterIP: None`).
- **Use Case:** StatefulSets.
- Instead of load balancing, DNS query returns **Multiple A Records** (IPs of all individual pods).
- Allows client to connect directly to specific pod.

---

### Question 87: Explain init containers.

**Answer:**
Special containers that run **before** app containers are started.
- Must run to completion successfully.
- **Use Cases:**
  - Delay start until DB is up (`nc -z db 5432`).
  - Configure disk permissions.
  - Download secrets/configs.

---

### Question 88: What is a PodDisruptionBudget (PDB)?

**Answer:**
Limits the number of Pods of a replicated application that are down simultaneously from voluntary disruptions (Nodes draining, upgrades).
- Ensures High Availability.
- `minAvailable: 2` means K8s will block node drain if it causes available pods to drop below 2.

---

### Question 89: What is custom resource definition (CRD)?

**Answer:**
Extension of the Kubernetes API.
- Allows you to define your own object types (e.g., `Prometheus`, `KafkaCluster`).
- Used heavily by Operators to manage custom application logic.

---

### Question 90: How does Kubernetes perform self-healing?

**Answer:**
1.  **Restart:** If container crashes (PID 1 exit), kubelet restarts it.
2.  **Reschedule:** If Node dies, Controller Manager notices and creates replacements on healthy nodes.
3.  **Liveness Probes:** Kubelet kills unhealthy container.
4.  **Readiness Probes:** Service LB removes traffic from unhealthy pod.

---

## 🔹 Real-world & DevOps (Questions 91-100)

### Question 91: How do you do blue-green deployments?

**Answer:**
1.  **Blue:** Active version (v1). Service points here.
2.  **Green:** New version (v2). Deployed but idle.
3.  **Switch:** Update Service Selector to point to Green labels.
4.  **Verification:** If issues, switch Service back to Blue instantly.

---

### Question 92: How do you do canary deployments?

**Answer:**
Already covered in Deployments section, but in tools:
- **Istio/Service Mesh:** Fine-grained traffic splitting (Weight 5%).
- **Argo Rollouts:** Native Canary Controller.
- **Ingress Controller (Nginx):** Annotation based canary-by-header or canary-by-weight.

---

### Question 93: What are some best practices in Kubernetes security?

**Answer:**
1.  **Least Privilege:** Use RBAC carefully.
2.  **No Root:** `runAsNonRoot: true`.
3.  **Network Policies:** Deny all by default.
4.  **Image Scanning:** Scan for CVEs.
5.  **Secrets:** Encrypt etcd, use external vaults.
6.  **Read Only FS:** `readOnlyRootFilesystem: true`.

---

### Question 94: How do you manage secrets in production?

**Answer:**
- Do not check secrets into Git.
- **Tools:**
  - **Sealed Secrets (Bitnami):** Encrypts secret into Safe format for Git. Controller decrypts in cluster.
  - **External Secrets Operator:** Fetches secrets from AWS Secrets Manager / HashiCorp Vault and creates K8s Secrets.

---

### Question 95: How do you backup and restore etcd?

**Answer:**
Use `etcdctl`.
**Backup:**
```bash
ETCDCTL_API=3 etcdctl snapshot save snapshot.db \
  --cacert=/etc/kubernetes/pki/etcd/ca.crt \
  --cert=/etc/kubernetes/pki/etcd/server.crt \
  --key=/etc/kubernetes/pki/etcd/server.key
```
**Restore:**
Restoring requires stopping API server and replacing data directory with snapshot.

---

### Question 96: What are some production monitoring strategies?

**Answer:**
- **Golden Signals:** Latency, Traffic, Errors, Saturation.
- **USE Method:** Utilization, Saturation, Errors (Infrastructure).
- **RED Method:** Rate, Errors, Duration (Services).
- **Alerting:** Alert on symptoms ("Website down"), not causes ("CPU high").

---

### Question 97: What’s the difference between Kubernetes and OpenShift?

**Answer:**
- **Kubernetes:** The Engine (Kernel). Open Source project.
- **OpenShift:** The Car (Product). Enterprise distribution by Red Hat.
  - Comes with integrated Registry, CI/CD (Tekton), Monitoring, stricter security (SCC), and a Web Console.
  - "Opinionated Kubernetes".

---

### Question 98: What cloud providers support Kubernetes?

**Answer:**
All major ones offer Managed Kubernetes:
- **AWS:** EKS (Elastic Kubernetes Service).
- **Google:** GKE (Google Kubernetes Engine) - The pioneer.
- **Azure:** AKS (Azure Kubernetes Service).
- **DigitalOcean:** DOKS.
- **Others:** IBM, Oracle, Linode.

---

### Question 99: What is kubeadm?

**Answer:**
A tool to bootstrap Kubernetes clusters.
- `kubeadm init`: Initializes control plane.
- `kubeadm join`: Joins worker nodes.
- Used for self-hosted/bare-metal setups (The "Hard Way" made easy).

---

### Question 100: What is the future of Kubernetes?

**Answer:**
- **Standardization:** It is the OS of the Cloud.
- **Serverless K8s:** GKE Autopilot, AWS Fargate (No nodes to manage).
- **Edge:** K3s running on IoT devices/Satellites.
- **AI/ML:** Kubeflow for running training workloads.
- **Simplification:** Improving Developer Experience (less YAML).

---
