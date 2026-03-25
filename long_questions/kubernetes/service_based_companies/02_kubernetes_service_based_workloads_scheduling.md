# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 2)
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, etc.
> **Focus:** Workload Management, Application Deployments, Stateful VS Stateless, and Pod Lifecycle.

---

## 🔹 Application Deployments & Updates (Questions 16-20)

### Question 16: What is a Rolling Update in Kubernetes, and how does it compare to a Recreate deployment strategy?

**Answer:**
- **Rolling Update (Strategy: `RollingUpdate`):** The default deployment strategy. It slowly replaces old Pods with new ones, ensuring there is always a certain number of Pods running (controlled by `maxSurge` and `maxUnavailable`). **Result:** Zero-downtime deployment.
- **Recreate (Strategy: `Recreate`):** Kills all existing Pods first before creating any new ones. **Result:** Application downtime during the update period. Used strictly when old and new code cannot connect to the database schema simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Rolling Update in Kubernetes, and how does it compare to a Recreate deployment strategy?
**Your Response:** "Rolling Update is like renovating a restaurant while keeping it open - you close a few tables at a time, renovate them, then reopen while working on the next section. Customers can still dine, so there's no downtime. Recreate is like closing the entire restaurant for renovation - everyone has to leave, you do all the work, then reopen. Rolling Update is the default because it gives zero downtime, but sometimes you need Recreate when old and new versions can't coexist, like major database changes."

---

### Question 17: How would you rollback a Deployment in Kubernetes?

**Answer:**
Kubernetes Deployments maintain a revision history by keeping old ReplicaSets.
1. Check history: `kubectl rollout history deployment/<name>`
2. Undo to the previous version: `kubectl rollout undo deployment/<name>`
3. Undo to a specific revision: `kubectl rollout undo deployment/<name> --to-revision=2`
Warning: Uncommitted changes in ConfigMaps/Secrets are not rolled back via this command.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you rollback a Deployment in Kubernetes?
**Your Response:** "Kubernetes automatically keeps previous versions of your deployment by maintaining old ReplicaSets. To rollback, I first check the history with `kubectl rollout history` to see what versions are available. Then I simply run `kubectl rollout undo` to go back to the previous version, or I can specify a particular revision if needed. It's like having a time machine for your deployments - one command and you're back to the previous working version. Just remember that changes to ConfigMaps or Secrets won't rollback automatically - those need to be handled separately."

---

### Question 18: What is a DaemonSet and when would you use it?

**Answer:**
A **DaemonSet** guarantees that exactly *one* instance of a Pod runs on *all* (or a subset selected via nodeSelector) Nodes in the cluster.
**Use Cases:**
1. **Log Aggregation:** Fluentd or Filebeat running on every node to scoop up container logs.
2. **Monitoring:** Prometheus Node Exporter DaemonSet running on all nodes to report CPU/Med.
3. **Networking:** The CNI plugin itself (e.g. `kube-flannel.ds` or `calico-node`) runs as a DaemonSet.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a DaemonSet and when would you use it?
**Your Response:** "A DaemonSet is like assigning a security guard to every building in a city - it ensures exactly one pod runs on every node in the cluster. I use it for things that need to be everywhere: log collectors to gather logs from all nodes, monitoring agents to track node health, or networking plugins that need to run on each server. Unlike regular deployments where you specify how many replicas you want, DaemonSet automatically puts one pod on each node and keeps it running there."

---

### Question 19: Explain Jobs and CronJobs in Kubernetes.

**Answer:**
- **Job:** Creates Pods to perform a specific finite task (like a database migration or a one-time script). Once the process exits successfully, the Pod moves to a `Completed` state successfully rather than being restarted (unlike Deployments where a pod must run forever).
- **CronJob:** A Job wrapped in a time-schedule using standard cron-expression format (e.g., `0 0 * * *` for midnight). Used for periodic backups, daily email digests, etc.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Jobs and CronJobs in Kubernetes.
**Your Response:** "Jobs are like hiring temporary workers for a specific task - they do their work and then they're done. A database migration script is a perfect example - it runs once, completes successfully, and the pod stops. CronJobs are like scheduling those temporary workers to show up regularly - same as a regular job but on a schedule. Think of it as setting up automatic backups that run every night at midnight. Unlike regular deployments that run forever, Jobs and CronJobs are designed to finish their work and disappear."

---

### Question 20: Can you modify a Pod's image directly once it's created? What about modifying a Deployment?

**Answer:**
- No, for an independent **Pod**, the image is an immutable field. You must delete the Pod and recreate it, or use standard update mechanisms if wrapped in a controller.
- Yes, for a **Deployment/ReplicaSet**, you can modify its template: `kubectl set image deployment/nginx-app nginx-app=nginx:1.20`. The deployment will natively generate a new ReplicaSet and spin up new Pods cleanly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you modify a Pod's image directly once it's created? What about modifying a Deployment?
**Your Response:** "For individual pods, the image is set in stone - you can't change it once the pod is created. You'd have to delete and recreate it. But for Deployments, absolutely! That's the whole point of deployments - you can update the image version and Kubernetes handles the rolling update automatically. I'd use `kubectl set image` to change the deployment's template, and Kubernetes will create new pods with the new image while gracefully retiring the old ones. This is why we always use deployments instead of managing pods directly."

---

## 🔹 Advanced Scheduling Basics (Questions 21-25)

### Question 21: What are Taints and Tolerations?

**Answer:**
They are used to restrict what Pods can run on what Nodes.
- **Taint (placed on the Node):** Repels Pods. (e.g., `kubectl taint nodes node-1 key=value:NoSchedule`)
- **Toleration (placed on the Pod spec):** Allows (but does not mandate) a Pod to bypass a specified taint and be scheduled on that specific Node.
**Example:** K8s master nodes are heavily tainted automatically so normal application Pods are never scheduled there.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Taints and Tolerations?
**Your Response:** "Taints and Tolerations are like bouncers and VIP passes at a club. Taints are placed on nodes and say 'no regular pods allowed here' - like the master nodes which are tainted to prevent applications from running there. Tolerations are like VIP passes that certain pods can carry to bypass the bouncer. A pod with a toleration can still run on a tainted node, but pods without the matching toleration will be rejected. It's how we keep certain nodes dedicated to specific workloads."

---

### Question 22: What is nodeSelector vs Node Affinity?

**Answer:**
Both tell the Scheduler where to place a Pod.
- **`nodeSelector`:** The older, simpler method requiring an exact match of key-value labels.
- **`Node Affinity`:** Much more expressive. You can use operators (`In`, `NotIn`, `Exists`) and set rules to either strictly mandate placement (`requiredDuringSchedulingIgnoredDuringExecution`) or just prefer it broadly (`preferredDuringSchedulingIgnoredDuringExecution`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is nodeSelector vs Node Affinity?
**Your Response:** "nodeSelector is the basic version - it's like telling the scheduler 'this pod must go on a node with this exact label.' It's simple but limited. Node Affinity is the advanced version - it's like giving the scheduler sophisticated rules. I can say 'prefer nodes with SSD storage' or 'don't put this pod on nodes labeled as database servers.' Node Affinity gives me much more control with operators like 'not in' or 'exists', and I can make requirements mandatory or just preferences."

---

### Question 23: How do you prevent two replicas of the same Web API from being scheduled on the exact same worker node?

**Answer:**
You would use **Pod Anti-Affinity**. 
You configure the Deployment specification such that the Pod rejects scheduling if a Pod with the exact same application label already exists in the same topology domain (like `kubernetes.io/hostname`). This ensures High Availability across the nodes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent two replicas of the same Web API from being scheduled on the exact same worker node?
**Your Response:** "I'd use Pod Anti-Affinity, which is like telling the scheduler 'don't put all my eggs in one basket.' I configure the deployment so that pods of the same application refuse to be scheduled on the same node. This ensures high availability - if one node goes down, I still have pods running on other nodes. It's especially important for critical applications where you want to spread your replicas across different failure domains."

---

### Question 24: What are Resource Requests and Limits dynamically?

**Answer:**
- **Requests:** What the container *must* have to be scheduled on a node. The scheduler sums all requests to see if a node has room. If no node has enough resources matching the Request, the Pod stays `Pending`.
- **Limits:** The maximum CPU or Memory the application is allowed to spike to. If a container breaches its memory Limit, it is immediately `OOMKilled`. If it hits its CPU limit, it gets throttled (but not killed).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Resource Requests and Limits dynamically?
**Your Response:** "Requests are like reserving a table at a restaurant - you're guaranteed this much CPU and memory, and the scheduler uses these requests to decide which node has enough space. Limits are like saying 'I'll never eat more than this much' - it's the ceiling your application can hit. If you exceed memory limit, Kubernetes kills your pod immediately. If you exceed CPU limit, it just slows you down but doesn't kill you. This prevents noisy neighbor problems where one application starves others of resources."

---

### Question 25: What is a Static Pod?

**Answer:**
A Static Pod is managed directly by the `kubelet` daemon on a specific node, without the API server managing it.
You create it by placing its YAML file inside a specific directory on a node (usually `/etc/kubernetes/manifests/`). If you delete the manifest, kubelet stops the pod. 
**Usage:** Used fundamentally to bootstrap the K8s control plane itself (etcd, kube-apiserver, etc.).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Static Pod?
**Your Response:** "Static Pods are like the founding members of the cluster - they're managed directly by the kubelet on each node, not by the API server. You create them by dropping YAML files in a special directory on the node, and the kubelet just starts running them. They're special because they exist even before the API server is fully up - that's why we use them for the core components like etcd and the API server itself. Once the cluster is running, we mostly use regular deployments, but static pods are what bootstraps the whole system."

---

## 🔹 Configs & Troubleshooting Depth (Questions 26-30)

### Question 26: You inject a ConfigMap containing user config into a Pod via environment variables. If you update the ConfigMap, does the Pod get the new values?

**Answer:**
No. If injected via **Environment Variables**, the application Pod does NOT see updates to the ConfigMap. The pod must be restarted (`kubectl rollout restart deployment/<name>`) to absorb the new variables. 
However, if mounted as a **Volume File**, kubelet eventually updates the underlying file, though the running application must be capable of hot-reloading configurations by watching the file.

### How to Explain in Interview (Spoken style format)
**Interviewer:** You inject a ConfigMap containing user config into a Pod via environment variables. If you update the ConfigMap, does the Pod get the new values?
**Your Response:** "No, if I inject ConfigMap as environment variables, the pod won't see updates - those variables are set at pod startup and never change. I'd need to restart the deployment to pick up the new values. But if I mount the ConfigMap as a volume file, Kubernetes does update the file automatically - though my application needs to be smart enough to detect the file change and reload its configuration. Most applications need a restart for config changes, but some can hot-reload from mounted files."

---

### Question 27: Explain the concept of Init Containers.

**Answer:**
An `initContainer` is a specialized container that runs and completes strictly *before* the main application containers start.
If an Init Container fails, Kubernetes repeatedly restarts the Pod until it succeeds.
**Use case:** Delaying the startup of a web service until the database pod is ready, executing database schema migrations, or pulling external configuration from a vault before the app boots.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the concept of Init Containers.
**Your Response:** "Init Containers are like setup crew that prepare the stage before the main actors arrive. They run to completion before your main application containers even start, and if they fail, the whole pod restarts. I use them for things like running database migrations before my web app starts, or waiting for a database to be ready, or pulling secrets from a vault. They ensure all prerequisites are met before my application tries to run. Think of them as the 'checklist' that must complete successfully before the main show begins."

---

### Question 28: A service fails to connect to its database. Where exactly do you start troubleshooting?

**Answer:**
1. Determine if the app pod itself is healthy using `kubectl logs` and `kubectl get pods`.
2. Ensure you are targeting the right Service name. Check core-dns resolution: `nslookup <db-service>`.
3. Verify the Endpoint binding! `kubectl get endpoints <db-service>`. If it is `<none>`, the labels on the Service do not match the labels on the DB pod.
4. If Endpoint exists, check Network Policies restricting cross-namespace ingress.

### How to Explain in Interview (Spoken style format)
**Interviewer:** A service fails to connect to its database. Where exactly do you start troubleshooting?
**Your Response:** "I follow a systematic approach. First, I check if the application pod itself is healthy with `kubectl logs` and `kubectl get pods`. Then I verify DNS resolution with `nslookup` to make sure the service name resolves correctly. Next, I check the endpoints with `kubectl get endpoints` - if it shows `<none>`, that means the service selector labels don't match the pod labels. Finally, if endpoints look good, I check Network Policies that might be blocking traffic. This covers 90% of database connection issues in Kubernetes."

---

### Question 29: What happens when an OOMKilled application is managed by a Deployment? Will it recover?

**Answer:**
When a Pod is `OOMKilled` (Out of Memory), the container stops abruptly. Because its `.spec.restartPolicy` defaults to `Always`, the `kubelet` simply restarts the container inside the same Pod (triggering a `CrashLoopBackOff` loop if it repeatedly balloons memory too fast). The Deployment notices the crashing behavior, but ultimately it's relying on `kubelet` attempting to resurrect it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when an OOMKilled application is managed by a Deployment? Will it recover?
**Your Response:** "When a pod gets OOMKilled, the kubelet immediately restarts the container inside the same pod. If the application keeps consuming too much memory, you'll see a CrashLoopBackOff pattern where it keeps dying and restarting. The deployment controller notices this but relies on the kubelet to handle the restarts. The deployment won't create new pods on different nodes - it just keeps trying to restart the crashed container. To truly fix it, you need to either increase the memory limit or fix the memory leak in the application."

---

### Question 30: What is helm and why is it called a package manager for Kubernetes?

**Answer:**
**Helm** abstracts away complex Kubernetes YAMLs into templated charts. 
Instead of maintaining 10 manually linked YAML files (Deployment, Service, Ingress, Secrets) across 3 environments (Dev/Test/Prod), Helm allows you to package them all as one "application" and pass in different variables (values) via a `values.yaml` file. It installs, updates, and deletes all associated resources reliably in one command.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is helm and why is it called a package manager for Kubernetes?
**Your Response:** "Helm is like the Docker Hub or npm for Kubernetes applications. Instead of managing dozens of separate YAML files for deployments, services, and configmaps, Helm packages everything into a 'chart' - like a zip file for your entire application. I can have different values files for dev, test, and prod environments, and with one command install or update the whole stack. It's called a package manager because it does for Kubernetes what apt does for Linux or npm does for Node.js - it makes managing complex applications simple and repeatable."
