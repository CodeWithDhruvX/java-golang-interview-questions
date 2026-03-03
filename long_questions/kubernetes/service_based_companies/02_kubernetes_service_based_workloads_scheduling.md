# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 2)
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, etc.
> **Focus:** Workload Management, Application Deployments, Stateful VS Stateless, and Pod Lifecycle.

---

## 🔹 Application Deployments & Updates (Questions 16-20)

### Question 16: What is a Rolling Update in Kubernetes, and how does it compare to a Recreate deployment strategy?

**Answer:**
- **Rolling Update (Strategy: `RollingUpdate`):** The default deployment strategy. It slowly replaces old Pods with new ones, ensuring there is always a certain number of Pods running (controlled by `maxSurge` and `maxUnavailable`). **Result:** Zero-downtime deployment.
- **Recreate (Strategy: `Recreate`):** Kills all existing Pods first before creating any new ones. **Result:** Application downtime during the update period. Used strictly when old and new code cannot connect to the database schema simultaneously.

---

### Question 17: How would you rollback a Deployment in Kubernetes?

**Answer:**
Kubernetes Deployments maintain a revision history by keeping old ReplicaSets.
1. Check history: `kubectl rollout history deployment/<name>`
2. Undo to the previous version: `kubectl rollout undo deployment/<name>`
3. Undo to a specific revision: `kubectl rollout undo deployment/<name> --to-revision=2`
Warning: Uncommitted changes in ConfigMaps/Secrets are not rolled back via this command.

---

### Question 18: What is a DaemonSet and when would you use it?

**Answer:**
A **DaemonSet** guarantees that exactly *one* instance of a Pod runs on *all* (or a subset selected via nodeSelector) Nodes in the cluster.
**Use Cases:**
1. **Log Aggregation:** Fluentd or Filebeat running on every node to scoop up container logs.
2. **Monitoring:** Prometheus Node Exporter DaemonSet running on all nodes to report CPU/Med.
3. **Networking:** The CNI plugin itself (e.g. `kube-flannel.ds` or `calico-node`) runs as a DaemonSet.

---

### Question 19: Explain Jobs and CronJobs in Kubernetes.

**Answer:**
- **Job:** Creates Pods to perform a specific finite task (like a database migration or a one-time script). Once the process exits successfully, the Pod moves to a `Completed` state successfully rather than being restarted (unlike Deployments where a pod must run forever).
- **CronJob:** A Job wrapped in a time-schedule using standard cron-expression format (e.g., `0 0 * * *` for midnight). Used for periodic backups, daily email digests, etc.

---

### Question 20: Can you modify a Pod's image directly once it's created? What about modifying a Deployment?

**Answer:**
- No, for an independent **Pod**, the image is an immutable field. You must delete the Pod and recreate it, or use standard update mechanisms if wrapped in a controller.
- Yes, for a **Deployment/ReplicaSet**, you can modify its template: `kubectl set image deployment/nginx-app nginx-app=nginx:1.20`. The deployment will natively generate a new ReplicaSet and spin up new Pods cleanly.

---

## 🔹 Advanced Scheduling Basics (Questions 21-25)

### Question 21: What are Taints and Tolerations?

**Answer:**
They are used to restrict what Pods can run on what Nodes.
- **Taint (placed on the Node):** Repels Pods. (e.g., `kubectl taint nodes node-1 key=value:NoSchedule`)
- **Toleration (placed on the Pod spec):** Allows (but does not mandate) a Pod to bypass a specified taint and be scheduled on that specific Node.
**Example:** K8s master nodes are heavily tainted automatically so normal application Pods are never scheduled there.

---

### Question 22: What is nodeSelector vs Node Affinity?

**Answer:**
Both tell the Scheduler where to place a Pod.
- **`nodeSelector`:** The older, simpler method requiring an exact match of key-value labels.
- **`Node Affinity`:** Much more expressive. You can use operators (`In`, `NotIn`, `Exists`) and set rules to either strictly mandate placement (`requiredDuringSchedulingIgnoredDuringExecution`) or just prefer it broadly (`preferredDuringSchedulingIgnoredDuringExecution`).

---

### Question 23: How do you prevent two replicas of the same Web API from being scheduled on the exact same worker node?

**Answer:**
You would use **Pod Anti-Affinity**. 
You configure the Deployment specification such that the Pod rejects scheduling if a Pod with the exact same application label already exists in the same topology domain (like `kubernetes.io/hostname`). This ensures High Availability across the nodes.

---

### Question 24: What are Resource Requests and Limits dynamically?

**Answer:**
- **Requests:** What the container *must* have to be scheduled on a node. The scheduler sums all requests to see if a node has room. If no node has enough resources matching the Request, the Pod stays `Pending`.
- **Limits:** The maximum CPU or Memory the application is allowed to spike to. If a container breaches its memory Limit, it is immediately `OOMKilled`. If it hits its CPU limit, it gets throttled (but not killed).

---

### Question 25: What is a Static Pod?

**Answer:**
A Static Pod is managed directly by the `kubelet` daemon on a specific node, without the API server managing it.
You create it by placing its YAML file inside a specific directory on a node (usually `/etc/kubernetes/manifests/`). If you delete the manifest, kubelet stops the pod. 
**Usage:** Used fundamentally to bootstrap the K8s control plane itself (etcd, kube-apiserver, etc.).

---

## 🔹 Configs & Troubleshooting Depth (Questions 26-30)

### Question 26: You inject a ConfigMap containing user config into a Pod via environment variables. If you update the ConfigMap, does the Pod get the new values?

**Answer:**
No. If injected via **Environment Variables**, the application Pod does NOT see updates to the ConfigMap. The pod must be restarted (`kubectl rollout restart deployment/<name>`) to absorb the new variables. 
However, if mounted as a **Volume File**, kubelet eventually updates the underlying file, though the running application must be capable of hot-reloading configurations by watching the file.

---

### Question 27: Explain the concept of Init Containers.

**Answer:**
An `initContainer` is a specialized container that runs and completes strictly *before* the main application containers start.
If an Init Container fails, Kubernetes repeatedly restarts the Pod until it succeeds.
**Use case:** Delaying the startup of a web service until the database pod is ready, executing database schema migrations, or pulling external configuration from a vault before the app boots.

---

### Question 28: A service fails to connect to its database. Where exactly do you start troubleshooting?

**Answer:**
1. Determine if the app pod itself is healthy using `kubectl logs` and `kubectl get pods`.
2. Ensure you are targeting the right Service name. Check core-dns resolution: `nslookup <db-service>`.
3. Verify the Endpoint binding! `kubectl get endpoints <db-service>`. If it is `<none>`, the labels on the Service do not match the labels on the DB pod.
4. If Endpoint exists, check Network Policies restricting cross-namespace ingress.

---

### Question 29: What happens when an OOMKilled application is managed by a Deployment? Will it recover?

**Answer:**
When a Pod is `OOMKilled` (Out of Memory), the container stops abruptly. Because its `.spec.restartPolicy` defaults to `Always`, the `kubelet` simply restarts the container inside the same Pod (triggering a `CrashLoopBackOff` loop if it repeatedly balloons memory too fast). The Deployment notices the crashing behavior, but ultimately it's relying on `kubelet` attempting to resurrect it.

---

### Question 30: What is helm and why is it called a package manager for Kubernetes?

**Answer:**
**Helm** abstracts away complex Kubernetes YAMLs into templated charts. 
Instead of maintaining 10 manually linked YAML files (Deployment, Service, Ingress, Secrets) across 3 environments (Dev/Test/Prod), Helm allows you to package them all as one "application" and pass in different variables (values) via a `values.yaml` file. It installs, updates, and deletes all associated resources reliably in one command.
