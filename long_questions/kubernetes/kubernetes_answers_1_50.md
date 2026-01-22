## 🔹 Kubernetes Basics (Questions 1-10)

### Question 1: What is Kubernetes?

**Answer:**
Kubernetes (K8s) is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

**Key Features:**
- **Automated Rollouts/Rollbacks:** Progressively rollout changes or rollback if something breaks.
- **Service Discovery & Load Balancing:** Gives containers their own IP and a single DNS name for a set of containers.
- **Storage Orchestration:** Automatically mounts storage systems (local, cloud, NFS).
- **Self-Healing:** Restarts failed containers, replaces/reschedules them.

**Example Command:**
```bash
kubectl version --short
```

---

### Question 2: Why do we use Kubernetes?

**Answer:**
Kubernetes solves the problem of managing containers at scale.
1.  **Scalability:** Automatically scale up/down based on traffic.
2.  **High Availability:** Ensures zero downtime deployments.
3.  **Portability:** Runs on any cloud (AWS, Azure, GCP) or on-prem.
4.  **Efficiency:** Packs applications efficiently onto hardware nodes to save costs.

---

### Question 3: What is a Pod in Kubernetes?

**Answer:**
A Pod is the smallest deployable object in Kubernetes. It represents a single instance of a running process in the cluster.
- A Pod can contain one or more containers (e.g., main app + sidecar).
- Containers in a Pod share the same **Network IP**, **Storage Volumes**, and **localhost**.

**YAML Example:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
```

---

### Question 4: What is a Node?

**Answer:**
A Node is a physical or virtual machine that is part of the Kubernetes cluster.
- **Master Node:** Manages the cluster (Control Plane).
- **Worker Node:** Runs the actual applications (Pods).
Each node contains services like `kubelet`, `kube-proxy`, and a Container Runtime (Docker/containerd).

---

### Question 5: What is a Cluster?

**Answer:**
A Kubernetes Cluster is a set of nodes interacting with each other. It consists of:
1.  **Control Plane (Master):** Makes decisions (scheduling, detecting events).
2.  **Data Plane (Workers):** Runs the workloads.
A cluster must have at least one worker node.

---

### Question 6: What is a Namespace?

**Answer:**
Namespaces provide a mechanism to isolate resources within a single cluster. They are like "virtual clusters."
- **Use Case:** Separate environments like `dev`, `test`, `prod` or separate teams.
- **Default Namespaces:** `default`, `kube-system` (for K8s internals), `kube-public`.

**Commands:**
```bash
kubectl create namespace dev
kubectl get pods -n dev
```

---

### Question 7: What is the difference between Docker and Kubernetes?

**Answer:**
They are complementary technologies.

| Docker | Kubernetes |
| :--- | :--- |
| Container Runtime | Container Orchestrator |
| Builds and runs a single container | Manages a cluster of containers |
| "A single brick" | "A building made of bricks" |
| `docker run` | `kubectl apply` |

*Note: Docker Swarm is a competitor to Kubernetes, but Docker itself is the underlying technology K8s often uses.*

---

### Question 8: What is Minikube?

**Answer:**
Minikube is a lightweight tool that runs a single-node Kubernetes cluster inside a Virtual Machine (VM) or container on your local laptop.
- Used for learning and local development.
- Supports standard K8s features (DNS, NodePorts, ConfigMaps).

**Command:**
```bash
minikube start
```

---

### Question 9: What is kubelet?

**Answer:**
Kubelet is an agent that runs on **every worker node**.
- It ensures that containers defined in PodSpecs are running and healthy.
- It communicates with the API Server to receive instructions (e.g., "Start this pod").
- It does **not** manage containers that were not created by Kubernetes.

---

### Question 10: What is kubectl?

**Answer:**
`kubectl` is the command-line interface (CLI) tool for communicating with the Kubernetes API server.
- It converts user commands into REST API calls.

**Common Commands:**
```bash
kubectl get nodes
kubectl describe pod my-pod
kubectl logs my-pod
```

---

## 🔹 Architecture (Questions 11-20)

### Question 11: Explain Kubernetes architecture.

**Answer:**
Kubernetes follows a Master-Worker architecture.
1.  **Control Plane (Master):**
    - **API Server:** Frontend for K8s users.
    - **Etcd:** Key-value store for cluster data.
    - **Scheduler:** Assigns pods to nodes.
    - **Controller Manager:** Maintains cluster state.
2.  **Worker Node:**
    - **Kubelet:** Agent managing pods.
    - **Kube-proxy:** Network proxy.
    - **Container Runtime:** Runs containers.

---

### Question 12: What are the components of the Master Node?

**Answer:**
The Control Plane components are:
1.  **kube-apiserver:** The brain; handles all REST requests.
2.  **etcd:** Distributed storage for all cluster data (State).
3.  **kube-scheduler:** Decides which node a pod goes to.
4.  **kube-controller-manager:** Runs loops to ensure Desired State matches Current State (e.g., Node Controller, Replicas Controller).
5.  **cloud-controller-manager:** Connects to cloud provider APIs (AWS/GCP).

---

### Question 13: What is the role of etcd?

**Answer:**
Etcd is a consistent and highly-available key-value store used as Kubernetes' backing store for all cluster data.
- Stores configs, secrets, pod states, etc.
- **Critical:** If etcd is lost, the cluster is lost. Always backup etcd.

---

### Question 14: What is the Scheduler?

**Answer:**
The `kube-scheduler` watches for newly created Pods that have no Node assigned.
- It selects the best node for the Pod based on:
  - Resource requirements (CPU/RAM).
  - Taints and Tolerations.
  - Node Affinity.
  - Data Locality.

---

### Question 15: What is the Controller Manager?

**Answer:**
A daemon that runs core control loops. Even though they are separate logically (Node Controller, Job Controller), they are compiled into a single binary.
- **Job:** Watch state → if difference → make changes to reach desired state.

---

### Question 16: What is the API Server?

**Answer:**
`kube-apiserver` is the front end of the Kubernetes control plane.
- It exposes the Kubernetes API (usually on port 443).
- **Functions:** Authenticates users, validates requests, retrieves data from etcd, and updates etcd.
- Only component that talks to etcd directly.

---

### Question 17: What is a worker node?

**Answer:**
A machine (VM or bare metal) that performs the actual work.
**Components:**
1.  **Kubelet:** Talks to Master.
2.  **Kube-proxy:** Handles networking rules.
3.  **Container Runtime:** (Docker/containerd) Runs the images.

---

### Question 18: What are DaemonSets?

**Answer:**
A DaemonSet ensures that **all** (or some specific) Nodes run a copy of a Pod.
- **Use Cases:**
  - Log collectors (Fluentd) on every node.
  - Monitoring agents (Prometheus Node Exporter).
  - Cluster storage daemons (Glusterd).
- When a new node joins, the DaemonSet automatically adds the pod to it.

---

### Question 19: What is kube-proxy?

**Answer:**
`kube-proxy` runs on each node and implements Kubernetes Service concept.
- It maintains network rules on the host (using iptables or IPVS).
- It allows network communication to your Pods from inside or outside the cluster.

---

### Question 20: How does service discovery work in Kubernetes?

**Answer:**
Two main methods:
1.  **DNS:** Kubernetes has an internal DNS server (CoreDNS).
    - Service `my-service` in namespace `default` gets entry `my-service.default.svc.cluster.local`.
2.  **Environment Variables:** When a Pod starts, K8s injects environment variables for all currently running services (e.g., `MY_SERVICE_Host=10.0.0.1`).

---

## 🔹 Pods & ReplicaSets (Questions 21-25)

### Question 21: What is a ReplicaSet?

**Answer:**
A ReplicaSet ensures that a specified number of pod replicas are running at any given time.
- It replaces Pods that are deleted or fail.
- Usually managed by a **Deployment**, not used directly.

**Selector:**
```yaml
selector:
  matchLabels:
    app: web
```

---

### Question 22: How is a ReplicaSet different from a ReplicationController?

**Answer:**
- **ReplicationController (Old):** Only supports equality-based selectors (`env = prod`).
- **ReplicaSet (New):** Supports set-based selectors (`env in (prod, stage)`).
- ReplicationControllers are deprecated.

---

### Question 23: How do you scale Pods in Kubernetes?

**Answer:**
1.  **Imperative:**
    ```bash
    kubectl scale deployment my-app --replicas=5
    ```
2.  **Declarative:** Edit YAML `replicas: 5` and apply.
3.  **Autoscaling (HPA):**
    ```bash
    kubectl autoscale deployment my-app --min=2 --max=10 --cpu-percent=80
    ```

---

### Question 24: What is a multi-container Pod?

**Answer:**
A Pod containing more than one container.
- They share the same network (localhost) and storage.
- **Patterns:**
  - **Sidecar:** Use logging agent alongside app.
  - **Ambassador:** Proxy to external world.
  - **Adapter:** Convert output to standard format.

---

### Question 25: How do Pods communicate with each other?

**Answer:**
1.  **Same Pod:** Via `localhost`.
2.  **Different Pods (Same Node):** Via Bridge Network/CNI Plugin without NAT.
3.  **Different Pods (Different Nodes):** Via Overlay Network (Flannel, Calico) so Pod IP is reachable across nodes.
4.  **Service:** Best practice implies using Service IP/DNS to abstract Pod IPs.

---

## 🔹 Deployments & Rollouts (Questions 26-30)

### Question 26: What is a Deployment?

**Answer:**
A Deployment provides declarative updates for Pods and ReplicaSets.
- You describe a *desired state* in a Deployment, and the Controller changes the actual state to the desired state at a controlled rate.
- Allows **Rolling Updates** and **Rollbacks**.

**YAML:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
```

---

### Question 27: How do you roll back a Deployment?

**Answer:**
Kubernetes keeps a history of rollouts (ReplicaSets).

**Commands:**
```bash
# Check history
kubectl rollout history deployment/my-app

# Undo last deployment
kubectl rollout undo deployment/my-app

# Rollback to specific revision
kubectl rollout undo deployment/my-app --to-revision=2
```

---

### Question 28: What is a Rolling Update?

**Answer:**
The default deployment strategy.
- Updates Pods in a rolling fashion, not all at once.
- **MaxUnavailable:** Max pods that can be down during update (default 25%).
- **MaxSurge:** Max pods that can be created above desired count (default 25%).
- Ensures Zero Downtime.

---

### Question 29: How do you pause and resume Deployments?

**Answer:**
Useful for making multiple changes and triggering only one rollout.

```bash
# Pause
kubectl rollout pause deployment/my-app

# Update image, resources, env vars... (No rollout triggers yet)
kubectl set image deployment/my-app nginx=nginx:1.16

# Resume
kubectl rollout resume deployment/my-app
```

---

### Question 30: What is a canary deployment in Kubernetes?

**Answer:**
A strategy to release a new version to a small subset of users before full rollout.
**Implementation:**
1.  Create `Deployment-V1` (Stable) with 90 replicas.
2.  Create `Deployment-V2` (Canary) with 10 replicas.
3.  Both Deployments use the **same Service Label**.
4.  Service load balances traffic 90:10.
5.  If V2 is good, scale V2 up and V1 down.

---

## 🔹 Services & Networking (Questions 31-40)

### Question 31: What is a Service in Kubernetes?

**Answer:**
A Service is an abstraction which defines a logical set of Pods and a policy to access them.
- Pods are ephemeral (IPs change).
- Services provide a stable **ClusterIP** and **DNS name** that doesn't change.

---

### Question 32: What are the types of Kubernetes Services?

**Answer:**
1.  **ClusterIP:** (Default) Internal IP, reachable only within cluster.
2.  **NodePort:** Exposes service on a static port on each Node's IP.
3.  **LoadBalancer:** Provisions an external Cloud Load Balancer (AWS ELB).
4.  **ExternalName:** Maps service to a DNS name (CNAME), no proxying.

---

### Question 33: What is a ClusterIP service?

**Answer:**
The default Service type.
- Assigns a virtual IP address accessible only from within the cluster.
- Used for internal communication (e.g., Backend → Database).

---

### Question 34: What is a NodePort service?

**Answer:**
Exposes the Service on each Node’s IP at a static port (default range 30000-32767).
- **Access:** `<NodeIP>:<NodePort>`.
- **Flow:** User → Node:Port → Service → Pod.
- Useful for dev or on-prem without Load Balancer.

---

### Question 35: What is a LoadBalancer service?

**Answer:**
Exposes the Service externally using a cloud provider's load balancer.
- **AWS:** Creates an ELB/NLB.
- **GCP:** Creates a Cloud Load Balancer.
- Traffic comes into the LB and is routed to Kubernetes Nodes.

---

### Question 36: What is an Ingress?

**Answer:**
Ingress is NOT a Service type, but an API object that manages external access to services, typically HTTP/HTTPS.
- It provides load balancing, SSL termination, and name-based virtual hosting.
- Requires an **Ingress Controller** (e.g., Nginx, Traefik).

**Example:**
`api.myapp.com` -> Service A
`web.myapp.com` -> Service B

---

### Question 37: Difference between Ingress and LoadBalancer?

**Answer:**
- **LoadBalancer Service:** High cost (One LB per service). 10 services = 10 ELBs. Layer 4 (TCP/UDP).
- **Ingress:** Cost effective (One LB for the Ingress Controller). Routes traffic to 10 services based on Host/Path. Layer 7 (HTTP/HTTPS).

---

### Question 38: How does DNS work in Kubernetes?

**Answer:**
Kubernetes runs a cluster-wide DNS server (CoreDNS).
- Every Service creates a DNS record.
- **Format:** `service-name.namespace.svc.cluster.local`.
- Pods interact with services using names instead of IPs.

---

### Question 39: What is a NetworkPolicy?

**Answer:**
Acts like a firewall for Pods.
- By default, all Pods can talk to all Pods (Flat network).
- **NetworkPolicy** allows you to restrict traffic.
  - "Allow traffic to DB-Pod only from Backend-Pod".
- Requires a CNI plugin that supports it (Calico, Cilium).

---

### Question 40: How do you expose a service to the outside world?

**Answer:**
Three ways:
1.  **NodePort:** Open port on server.
2.  **LoadBalancer:** Use Cloud LB.
3.  **Ingress:** Use an Ingress Controller to route HTTP/S traffic.
4.  **Port Forward (Dev):** `kubectl port-forward svc/myservice 8080:80`.

---

## 🔹 ConfigMaps & Secrets (Questions 41-45)

### Question 41: What is a ConfigMap?

**Answer:**
An API object used to store non-confidential data in key-value pairs.
- Decouples configuration artifacts from image content.
- **Usage:** Environment variables, command-line args, or config files.

**Creation:**
```bash
kubectl create configmap app-config --from-literal=DB_HOST=localhost
```

---

### Question 42: What is a Secret?

**Answer:**
Similar to ConfigMap but intended to hold sensitive information (passwords, OAuth tokens, SSH keys).
- Stored as base64 encoded strings in etcd.
- Mounted into pods as files or env vars.

---

### Question 43: How do you mount a ConfigMap to a Pod?

**Answer:**
1.  **As Environment Variable:**
    ```yaml
    env:
      - name: DB_HOST
        valueFrom:
          configMapKeyRef:
            name: app-config
            key: DB_HOST
    ```
2.  **As Volume:**
    ```yaml
    volumes:
      - name: config-vol
        configMap:
          name: app-config
    ```

---

### Question 44: How are Secrets stored in Kubernetes?

**Answer:**
- By default, they are stored **unencrypted** (base64 encoded) in etcd.
- **Security Best Practice:** Enable **Encryption at Rest** configuration in the API server to encrypt secrets in etcd.

---

### Question 45: How do you update a ConfigMap without restarting a Pod?

**Answer:**
- **Volume Mount:** If ConfigMap is mounted as a volume, Kubernetes automatically updates the file in the Pod (takes ~1 minute). The app must support file watching (hot reload).
- **Env Vars:** REQUIRES a Pod restart. Env vars are only set at startup.

---

## 🔹 Volumes & Storage (Questions 46-50)

### Question 46: What is a Volume in Kubernetes?

**Answer:**
A directory accessible to containers in a Pod.
- Unlike Docker volumes, a K8s Volume has an explicit lifetime same as the **Pod**.
- If container restarts, data is safe. If Pod is deleted, ephemeral volume types (`emptyDir`) are deleted.

---

### Question 47: Difference between emptyDir and hostPath?

**Answer:**
- **emptyDir:** Created when Pod starts, empty initially. Deleted when Pod is removed from node. Used for temporary scratch space.
- **hostPath:** Mounts a file/directory from the **Node's filesystem** into the Pod. Persists as long as data remains on Node. Used for system daemons / logs.

---

### Question 48: What is PersistentVolume (PV)?

**Answer:**
A piece of storage in the cluster that has been provisioned by an administrator or dynamically.
- It is a cluster resource (like a Node).
- It enables data to survive Pod restarts and Pod rescheduling to different nodes (if network storage).

---

### Question 49: What is PersistentVolumeClaim (PVC)?

**Answer:**
A request for storage by a user/Pod.
- A PVC consumes PV resources.
- PV is the "Hardware", PVC is the "Ticket/Request" to use it.
- Pods mount PVCs, not PVs directly.

---

### Question 50: What is StorageClass?

**Answer:**
Describes the "classes" of storage offered (e.g., "fast-ssd", "cheap-hdd").
- **Dynamic Provisioning:** Allows K8s to create a PV *automatically* when a PVC requests a specific StorageClass.
- Eliminates need for manual PV creation by admins.

**Example:**
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
```

---
