## 🔹 Ingress & Traffic Control (Questions 151-160)

### Question 151: What is the difference between Ingress Controller and Ingress Resource?

**Answer:**
- **Ingress Resource:** A K8s configuration object (YAML) defining routing rules (e.g., "Send `/api` to `backend-svc`"). It does nothing by itself.
- **Ingress Controller:** The actual software (Nginx, Traefik container) that watches these Resources and updates its config files to route the traffic.

---

### Question 152: What is the NGINX Ingress Controller?

**Answer:**
A popular Ingress Controller maintained by the Kubernetes community.
- It runs as a Pod (and Service/LoadBalancer).
- It reads Ingress Resources and builds `nginx.conf` dynamically.
- It reloads Nginx when configuration changes.

---

### Question 153: How do you implement TLS termination in Ingress?

**Answer:**
1.  **Create Secret:** Store `tls.crt` and `tls.key` in a generic Secret.
2.  **Edit Ingress:**
    ```yaml
    spec:
      tls:
      - hosts:
        - myapp.com
        secretName: myapp-tls
    ```
3.  The Ingress Controller handles the SSL handshake and forwards HTTP to the pod.

---

### Question 154: What is path-based routing in Ingress?

**Answer:**
Routing requests to different microservices based on the URL path.
- `example.com/shop` -> Shop Service
- `example.com/blog` -> Blog Service
Allows hosting multiple apps behind a single LoadBalancer IP.

---

### Question 155: How to handle rate limiting in Ingress?

**Answer:**
Using **Annotations** specific to the Controller.
For Nginx:
```yaml
metadata:
  annotations:
    nginx.ingress.kubernetes.io/limit-rps: "5"
```
Limits client to 5 requests per second.

---

### Question 156: What is sticky session in Kubernetes?

**Answer:**
Ensures that multiple requests from the same user (session) go to the **same Pod**.
- **Why?** If the app stores session state in memory (not Redis).
- **Ingress Annotation:** `nginx.ingress.kubernetes.io/affinity: "cookie"`.

---

### Question 157: How to configure custom headers in Ingress?

**Answer:**
Annotations or ConfigMap.
- `nginx.ingress.kubernetes.io/configuration-snippet: | more_set_headers "X-Custom: Header";`
- Useful for security headers (HSTS, X-Frame-Options).

---

### Question 158: What is external DNS in Kubernetes?

**Answer:**
An addon that synchronizes exposed Kubernetes Services/Ingresses with external DNS providers (AWS Route53, Google Cloud DNS).
- If you create Ingress `myapp.com`, it automatically creates an A Record in Route53 pointing to the LoadBalancer.

---

### Question 159: How do you manage HTTPS certificates in Kubernetes?

**Answer:**
Manually managing secrets is hard (expiry issues).
**Cert-Manager** is the standard solution.
- Automates certificate issuance from Let's Encrypt or internal PKI.
- Automatically renews certificates near expiry.

---

### Question 160: What is cert-manager?

**Answer:**
A native Kubernetes certificate management controller.
- **CRDs:** `Issuer`, `Certificate`.
- Watches `Ingress` resources.
- Performs ACME challenge (DNS01 or HTTP01) to validate domain ownership.
- Saves the certificate as a K8s Secret.

---

## 🔹 Logging & Monitoring - Advanced (Questions 161-170)

### Question 161: What are the best practices for centralized logging in Kubernetes?

**Answer:**
- **Stdout/Stderr:** Apps should write logs here (The 12-Factor App).
- **Node-Level Agent:** Run Fluentd/Filebeat as DaemonSet.
- **No Sidecars (Preferred):** Resource heavy to run sidecar per pod.
- **Structured Logging:** Use JSON logs for easier parsing.

---

### Question 162: What is the EFK stack?

**Answer:**
**Elasticsearch + Fluentd + Kibana.**
- **Fluentd:** Collects logs from nodes.
- **Elasticsearch:** Stores indices.
- **Kibana:** UI to search logs.
(Often simpler than ELK which uses Logstash).

---

### Question 163: What is Loki in Kubernetes logging?

**Answer:**
"Prometheus for Logs."
- Unlike Elasticsearch, Loki **does not index the text of the logs**. It groups them by labels (like Prometheus).
- Highly efficient/cheap storage.
- Integrated tightly with Grafana.

---

### Question 164: What is the difference between container logs and system logs?

**Answer:**
- **Container Logs:** Output of the application (e.g., Apache access logs).
- **System Logs:** Logs from Kubelet, Docker Daemon, Kernel (`/var/log/syslog`, `journald`).
- Monitoring system logs is crucial to catch "OOMKilled" or "XFS Corruption".

---

### Question 165: How to monitor etcd health?

**Answer:**
Etcd exposes metrics at `/metrics`.
**Key Metrics:**
- `etcd_server_has_leader`: Must be 1.
- `etcd_disk_wal_fsync_duration_seconds`: High latency means slow disk (Cluster instability).
- `etcd_server_proposals_failed_total`.

---

### Question 166: How do you implement alerting in Kubernetes?

**Answer:**
**Prometheus AlertManager.**
1.  Define **PrometheusRule** (YAML): `if: node_down > 5m`.
2.  Prometheus sends alert to AlertManager.
3.  AlertManager groups/deduplicates and sends to Slack/PagerDuty/Email.

---

### Question 167: What is kube-state-metrics?

**Answer:**
A service that talks to the API Server and generates metrics about the **state of objects**.
- **Metrics:** "How many pods are pending?", "How many replicas desired?", "When was deployment created?".
- Unlike node-exporter (which checks CPU), this checks K8s Metadata.

---

### Question 168: What are some metrics to monitor for production clusters?

**Answer:**
- **Node:** CPU, Mem, Disk (Saturation).
- **Kubelet:** Health.
- **API Server:** Request latency (slowness implies overload).
- **Etcd:** Fsync latency.
- **Pods:** Restarts (Looping), Pending count.

---

### Question 169: How to handle high disk I/O alerts in Kubernetes?

**Answer:**
1.  Check which pod is causing it (`iotop` on node).
2.  Common culprits: Logging (verbose), Database, Ephemeral usage.
3.  **Fix:** Move high I/O app to faster node, or limit IOPS via limits, or optimize app logging.

---

### Question 170: What are application-level vs infrastructure-level metrics?

**Answer:**
- **Infrastructure:** CPU, Memory, Network TX/RX (Provided by default).
- **Application:** "Number of HTTP 500s", "Checkout Transaction Time".
  - Requires developers to instrument code (using Prom Client Library) to expose custom `/metrics`.

---

## 🔹 CI/CD - Advanced (Questions 171-180)

### Question 171: How do you implement GitOps in Kubernetes?

**Answer:**
1.  Store ALL YAML manifests in Git.
2.  Install Controller (ArgoCD) in Cluster.
3.  Point ArgoCD to Git Repo.
4.  ArgoCD applies YAMLs.
5.  Developers never run `kubectl apply`. They only `git commit`.

---

### Question 172: What is the role of Argo Workflows?

**Answer:**
A workflow engine to orchestrate parallel jobs on Kubernetes.
- Unlike ArgoCD (Deploy tool), Workflows is for **Pipelines**.
- **Use Case:** Steps in a CI pipeline (Build -> Test -> Scan), Data processing, ML pipelines.

---

### Question 173: What is a Kubernetes webhook?

**Answer:**
An HTTP callback mechanism.
- K8s API Server calls external service to valid/mutate objects (Admission Webhooks).
- Or, AlertManager calls a webhook to notify a system.

---

### Question 174: How do you trigger a deployment pipeline on config change?

**Answer:**
- **GitOps:** Automatically detects commit.
- **CI-Push:** CI pipeline (GitHub Actions) runs `sed` to update the Image Tag in the YAML repo, commits it. GitOps tool sees change and syncs.

---

### Question 175: What are pre-deployment and post-deployment hooks?

**Answer:**
Scripts run during the deployment lifecycle.
- **Helm Hooks:** `pre-install` (Database migration), `test` (Integration test).
- **ArgoCD Hooks:** `PreSync` (Backups), `PostSync` (Notification).

---

### Question 176: What is a Helm lifecycle hook?

**Answer:**
Annotations in templates:
`"helm.sh/hook": post-install`
Helm waits for this pod/job to complete before considering the release "Successful".

---

### Question 177: How do you use Kustomize in CI/CD?

**Answer:**
**Kustomize** is a template-free way to customize YAML.
- **Base:** `deployment.yaml` (Generic).
- **Overlays:** `prod/kustomization.yaml` (patches replica count to 10).
- CI runs `kubectl apply -k overlays/prod`. Built into kubectl.

---

### Question 178: How do you rollback failed deployments automatically?

**Answer:**
- **Deployment:** Manual (`rollout undo`).
- **Helm:** `helm rollback` (manual).
- **ArgoCD/Flagger:** Can enable **Progressive Delivery**.
  - If analysis (metrics) fails during canary, it automatically aborts and rolls back.

---

### Question 179: What is Kubernetes-native CI/CD?

**Answer:**
Pipelines that run **as Pods** inside the cluster.
- **Examples:** Tekton, Jenkins X.
- **Pros:** Scalable (just more pods), secure (no external access needed).

---

### Question 180: How do you manage secrets in GitOps?

**Answer:**
You cannot store raw secrets in Git.
1.  **Bitnami Sealed Secrets:** Store encrypted (safe) file in Git.
2.  **SOPS:** Mozilla tool to encrypt values.
3.  **External Secrets:** Store reference in Git, secret acts as pointer to AWS/Vault.

---

## 🔹 Kubernetes Internals (Questions 181-190)

### Question 181: What is a container runtime?

**Answer:**
The software responsible for running containers.
- Kubernetes Kubelet talks to the Runtime via CRI.
- **Examples:** `containerd` (Standard), `CRI-O` (RedHat), `Docker` (Deprecated shim).

---

### Question 182: What runtimes are supported in Kubernetes?

**Answer:**
Any runtime implementing the **CRI (Container Runtime Interface)**.
- containerd
- CRI-O
- Mirantis Container Runtime
- Kata Containers (VM based).

---

### Question 183: What is CRI and why is it important?

**Answer:**
**Container Runtime Interface.**
- A protocol API (gRPC).
- Before CRI, K8s hardcoded Docker logic.
- With CRI, K8s can use *any* runtime without recompiling.

---

### Question 184: How do controllers work internally?

**Answer:**
**Control Loop:**
1.  **Watch:** Observe the API Server for changes.
2.  **Compare:** Current State vs Desired State.
3.  **Reconcile:** Take action to fix the difference (Create a pod, Delete a pod).
4.  **Repeat:** Infinite loop.

---

### Question 185: How does the Kubernetes Scheduler work?

**Answer:**
1.  **Informer:** Sees unscheduled pod.
2.  **Predicate (Filter):** Eliminates nodes (Not enough RAM, Tainted).
3.  **Priority (Score):** Scores nodes (0-10) based on spreading, affinity.
4.  **Bind:** Assigns Pod to NodeName in API Server.

---

### Question 186: What is a reconcile loop?

**Answer:**
The logic inside a controller.
```go
func Reconcile(req Request) Result {
   // get resource
   // if missing, create it
   // if different, update it
   return Result{}
}
```
Ensures eventual consistency.

---

### Question 187: What is a finalizer in Kubernetes?

**Answer:**
A string in `metadata.finalizers` that prevents an object from being hard-deleted.
- When you delete, K8s sets `deletionTimestamp`.
- Controller sees timestamp, runs cleanup logic (e.g., delete AWS LoadBalancer).
- Controller removes finalizer.
- K8s deletes object.

---

### Question 188: What is a taint effect NoExecute?

**Answer:**
Stronger than `NoSchedule`.
- `NoSchedule`: New pods won't land here. Existing pods stay.
- `NoExecute`: New pods won't land here. **Existing pods are evicted** (unless they tolerate it).
- Used for draining nodes or quarantining bad nodes.

---

### Question 189: How do events propagate inside Kubernetes?

**Answer:**
Components (Scheduler, Kubelet) create **Event Objects**.
- Events are stored in etcd for 1 hour (TTL).
- Users view them via `kubectl get events`.
- Not for long-term auditing (use Audit Logs for that).

---

### Question 190: What are API groups in Kubernetes?

**Answer:**
To organize the huge API.
- Core: `/api/v1` (Pod, Service, Node).
- Named Groups: `/apis/apps/v1` (Deployment), `/apis/batch/v1` (Job).
- Helps versioning different parts of the system independently.

---

## 🔹 Real-world & Production Readiness (Questions 191-200)

### Question 191: How do you prepare a cluster for production?

**Answer:**
1.  **HA Control Plane:** 3+ Masters, Multi-AZ.
2.  **Security:** RBAC, NetworkPolicies, Private Endpoint.
3.  **Monitoring:** Prometheus/Grafana set up.
4.  **Backup:** Velero for etcd/PVs.
5.  **Logging:** Centralized.

---

### Question 192: What are anti-patterns in Kubernetes deployment?

**Answer:**
1.  **Naked Pods:** Using Pods without Deployments (No self-healing).
2.  **Latest Tag:** Using `:latest` image (Unpredictable).
3.  **No Limits:** Ignoring Resource Limits (Cluster instability).
4.  **Config in Image:** Baking config files into Docker image instead of ConfigMap.
5.  **Manual Changes:** Using `kubectl edit` instead of GitOps.

---

### Question 193: How do you manage multiple environments (dev/stage/prod)?

**Answer:**
1.  **Namespaces:** Same cluster, `ns-dev`, `ns-prod`. (Good for small teams).
2.  **Clusters:** Separate clusters. (Best isolation).
3.  **Config Management:** Use Helm/Kustomize to override values per env (`values-prod.yaml`).

---

### Question 194: How do you handle image versioning?

**Answer:**
- **Semantic Versioning:** `v1.2.3`.
- **Git SHA:** `app:a1b2c3d`. (Good for traceability).
- **Process:** CI builds image, pushes unique tag, updates Manifest repo with that tag.

---

### Question 195: How do you manage secrets securely in CI/CD pipelines?

**Answer:**
Inject them at runtime.
- GitHub Secrets / GitLab Variables.
- CI Runner passes them as `--build-arg` (Still risky) or injects into K8s Secrets during deployment step.

---

### Question 196: What is the best way to do blue/green deployment in Kubernetes?

**Answer:**
Standard Service Selector Switch.
- Requires 2x resources temporarily.
- Instant switch.
- **Argo Rollouts** automates the "Promotion" step.

---

### Question 197: How to implement canary deployments using Istio?

**Answer:**
Istio **VirtualService**:
```yaml
route:
- destination:
    host: my-app
    subset: v1
  weight: 90
- destination:
    host: my-app
    subset: v2
  weight: 10
```
Gradually shift weight 10 -> 50 -> 100.

---

### Question 198: How do you secure inter-pod communication?

**Answer:**
**mTLS (Mutual TLS).**
- Ensuring both client and server authenticate each other and encryption.
- Hard to manage manually.
- **Service Mesh (Istio/Linkerd)** does this automatically (transparent mTLS).

---

### Question 199: What is multi-cluster Kubernetes?

**Answer:**
Running workloads across multiple K8s clusters.
- **Reasons:** High Availability, Geo-latency, Compliance (EU vs US data), Scale limits (5000 nodes).
- **Management:** Rancher, Tanzu, Google Anthos.

---

### Question 200: What is federation in Kubernetes?

**Answer:**
Syncing resources across multiple clusters.
- **KubeFed (Federation v2):** Define a `FederatedDeployment`, and KubeFed propogates it to Cluster A, B, and C.
- Complex and arguably finding less traction than GitOps-driven multi-cluster.

---
