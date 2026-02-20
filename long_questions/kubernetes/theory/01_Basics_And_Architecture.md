# 🟢 Kubernetes Basics & Architecture

---

### 1. What is Kubernetes?

"Kubernetes — or K8s — is an **open-source container orchestration platform** originally developed at Google and now maintained by the CNCF.

At its core, it automates the deployment, scaling, and management of containerized workloads. Think of it as a 'data center operating system' — you declare *what* you want (e.g., 3 replicas of a service), and Kubernetes figures out *how* to make that happen and keeps it that way.

I use it in production because it handles the hard parts: health checks, self-healing, rolling updates, service discovery, and scaling — things that are painful and brittle to wire up manually."

#### In Depth
Kubernetes was born from Google's internal Borg system, which ran billions of containers per week. K8s implements the **control loop pattern**: observe current state → compare to desired state → act to reconcile. This declarative model is why it's so resilient.

---

### 2. Why do we use Kubernetes?

"Before K8s, teams ran containers with Docker Compose or custom scripts — which worked fine for one machine but fell apart at scale.

K8s solves **distributed system problems**: Where do I place this container? What happens if the node dies? How do I roll out a new version without downtime? How do I auto-scale under load?

In my experience, the biggest win is **operational consistency** — every team uses the same YAML manifests, the same `kubectl` tooling, and the same deployment pipeline regardless of whether they're running 2 pods or 2000."

#### In Depth
Kubernetes abstracts away the underlying infrastructure (VMs, bare metal, cloud nodes) behind a unified API. This lets platform teams standardize deployments while development teams focus on writing code. The 12-factor app principles map naturally to K8s constructs.

---

### 3. What is a Pod in Kubernetes?

"A Pod is the **smallest deployable unit** in Kubernetes. It wraps one or more containers that share the same network namespace and storage.

Containers in a Pod share a single IP address and can communicate over `localhost`. They're co-scheduled and co-located on the same node.

I think of a Pod like a logical host — the containers inside are tight collaborators. A common pattern is a main app container with a sidecar for logging or proxying (like Envoy in a service mesh)."

#### In Depth
Pods are **ephemeral by design**. When a Pod dies, K8s creates a new one — it doesn't restart the old one in place. This is important for understanding why you don't store state inside a Pod directly. The PodSpec `restartPolicy` controls whether containers inside are restarted on failure.

---

### 4. What is a Node?

"A Node is a **worker machine** in Kubernetes — it can be a VM or bare-metal server. Nodes are where your actual containers run.

Each node runs three critical components: the **kubelet** (the node agent), the **kube-proxy** (networking rules), and a **container runtime** (like containerd).

In cloud environments like GKE or EKS, nodes are usually managed EC2 or GCE instances that the platform spins up and down for you. In on-prem setups, you own the hardware."

#### In Depth
Kubernetes tracks node health through **node conditions**: `Ready`, `DiskPressure`, `MemoryPressure`, etc. The node controller marks a node `NotReady` if the kubelet doesn't report back within the `node-monitor-grace-period`. After `pod-eviction-timeout`, pods are evicted and rescheduled elsewhere.

---

### 5. What is a Cluster?

"A Kubernetes cluster is the **full environment** — a control plane plus one or more worker nodes.

The control plane manages the overall cluster state (via etcd), issues scheduling decisions, and exposes the API. Worker nodes run workloads.

In a production setup, we typically run **3 or 5 control plane nodes** for HA (to maintain quorum), and as many worker nodes as needed for capacity."

#### In Depth
Cluster-wide communication happens over TLS. The API server is the single entry point — all other components (kubelet, controller-manager, scheduler) communicate through it rather than directly with each other. This design makes the cluster auditable and extensible.

---

### 6. What is a Namespace?

"A Namespace is a **virtual cluster** inside a physical cluster. It provides a scope for resource names and is used to isolate teams or environments.

Common patterns: one namespace per team (`team-payments`), one per environment (`staging`, `production`), or one per service (`namespace-auth`).

I've used namespaces extensively for **multi-tenancy** — applying separate ResourceQuotas, NetworkPolicies, and RBAC rules per namespace without spinning up separate clusters."

#### In Depth
Namespaces don't provide strong isolation — a pod in `namespace-a` can still reach a pod in `namespace-b` unless a NetworkPolicy blocks it. For hard isolation (e.g., different security domains), you need separate clusters or tools like Virtual Clusters (vCluster).

---

### 7. What is the difference between Docker and Kubernetes?

"Docker is a **container runtime and toolchain** — it builds images and runs containers on a single machine.

Kubernetes is an **orchestrator** — it manages containers *across a fleet of machines*, handling scheduling, networking, scaling, and failure recovery.

The analogy I use: Docker is like an individual airport gate; Kubernetes is the entire air traffic control system managing thousands of flights simultaneously."

#### In Depth
K8s deprecated Docker as its runtime in 1.20+ and shifted to the CRI (Container Runtime Interface). Tools like **containerd** or **CRI-O** are now the defaults. Docker images still work perfectly — it's just the runtime layer that changed. Docker itself uses containerd under the hood.

---

### 8. What is Minikube?

"Minikube is a tool to run a **single-node Kubernetes cluster locally** — perfect for development and testing.

It creates a local VM (or uses Docker) and bootstraps a full K8s cluster. I use it to validate manifests, test Helm charts, and experiment with K8s features before deploying to a real cluster.

For CI pipelines, I prefer `kind` (Kubernetes in Docker) because it starts faster and doesn't need a hypervisor."

#### In Depth
Minikube supports add-ons like the NGINX ingress controller, dashboard, and metrics-server out of the box via `minikube addons enable`. It also supports multi-node clusters now (`--nodes=3`), making it versatile for local testing of affinity rules.

---

### 9. What is kubelet?

"The kubelet is the **node agent** — the process that runs on every worker node and ensures containers are running as specified.

It watches the API server for PodSpecs assigned to its node, talks to the container runtime to start/stop containers, runs health probes, and reports node and pod status back to the control plane.

I think of the kubelet as the 'enforcer' on each node — it's responsible for reconciling the desired state (what the scheduler said should run here) with reality."

#### In Depth
The kubelet also manages volume mounting, hostPath access, device plugins (like GPUs), and node resource reporting. It does NOT manage pods created directly via the container runtime (these are called **static pods** and are placed in `/etc/kubernetes/manifests/`). This is how control plane components like the API server itself are run in `kubeadm` clusters.

---

### 10. What is kubectl?

"kubectl is the **CLI tool** for interacting with the Kubernetes API server.

It's how I do everything: deploying workloads, checking pod status, getting logs, exec-ing into containers, applying manifests, managing RBAC, etc.

My most-used commands daily are `kubectl get pods -A`, `kubectl logs`, `kubectl describe pod`, and `kubectl apply -f`. For production investigations, `kubectl top` and `kubectl events` are essential."

#### In Depth
kubectl reads your `~/.kube/config` to determine which cluster to connect to and uses the API server's REST interface. Every `kubectl` command is just a REST call under the hood — you can add `--v=8` to see the raw HTTP requests. This is useful for debugging what a command actually does.

---

### 11. Explain Kubernetes Architecture.

"Kubernetes has two major layers: the **Control Plane** and the **Worker Nodes**.

The **Control Plane** consists of: the API Server (the frontend), etcd (the distributed key-value store for all cluster state), the Scheduler (assigns pods to nodes), and the Controller Manager (runs reconciliation loops for deployments, replicasets, etc.).

The **Worker Nodes** run the actual workloads. Each has a kubelet (node agent), kube-proxy (networking), and a container runtime.

The golden rule: the control plane declares *what should happen*, worker nodes make it *actually happen*."

#### In Depth
All inter-component communication goes through the API server using the **watch mechanism** — components watch for changes to their resource types rather than polling. This event-driven architecture makes K8s scalable. The API server itself is stateless; all state is in etcd, which is why etcd backup is critical for disaster recovery.

---

### 12. What are the components of the Master Node?

"The master (control plane) has four main components:

1. **kube-apiserver** — the only component that reads/writes from etcd. All other components communicate through it.
2. **etcd** — distributed, consistent key-value store. The 'source of truth' for the entire cluster.
3. **kube-scheduler** — watches for unscheduled pods and assigns them to suitable nodes based on resource requirements, affinity, taints, etc.
4. **kube-controller-manager** — runs a suite of controllers: Node controller, Deployment controller, ReplicaSet controller, Job controller, etc."

#### In Depth
In a high-availability setup, you run **3 or 5 control plane replicas**. The scheduler and controller-manager use **leader election** — only one instance is active at a time. The API server is load balanced and all instances can be active simultaneously since it's stateless.

---

### 13. What is the role of etcd?

"etcd is the **brain's memory** — it's a distributed key-value store where Kubernetes persists every piece of cluster state: pods, deployments, configmaps, secrets, RBAC rules, everything.

If etcd loses data, the cluster loses everything. That's why **etcd backup is non-negotiable in production**.

I've seen clusters fail catastrophically because etcd ran out of disk space or suffered split-brain. Regular snapshots (`etcdctl snapshot save`) and monitoring etcd latency/disk I/O are critical production responsibilities."

#### In Depth
etcd uses the **Raft consensus algorithm** for distributed consistency. All writes go through the leader, which replicates to followers. A cluster of 3 can tolerate 1 failure; 5 can tolerate 2. etcd is particularly sensitive to disk I/O latency — even P99 spikes above 100ms can cause API server responsiveness issues.

---

### 14. What is the Scheduler?

"The Scheduler is the component that **assigns Pods to Nodes**.

When a Pod is created without a node assignment, the Scheduler picks it up. It runs through two phases: **Filtering** (which nodes can run this pod based on resources, taints, affinity?) and **Scoring** (which of the remaining nodes is the best fit?). It picks the highest-scored node and binds the Pod to it.

I've had to customize scheduler behavior using **pod affinity rules** and **topology spread constraints** to distribute pods evenly across availability zones."

#### In Depth
The default scheduler is extensible via **Scheduler Plugins** (the scheduler framework) and **Scheduling Extenders** (webhooks). You can also run **custom schedulers** alongside the default one. The scheduling decision is not final until the kubelet on the selected node acknowledges the pod — if the node becomes unavailable after scheduling, the pod stays in `Pending`.

---

### 15. What is the Controller Manager?

"The Controller Manager is a collection of **background control loops**, each watching a specific resource type and reconciling desired state with actual state.

Examples: the Deployment controller creates/updates ReplicaSets; the ReplicaSet controller ensures the right number of pods are running; the Node controller monitors node health.

I think of controllers as the 'janitors' of the cluster — constantly cleaning up and fixing things to match what you declared."

#### In Depth
Each controller is a goroutine in the `kube-controller-manager` binary. They use **list-watch** semantics: they cache the entire state locally and react to events (add, update, delete). The key property is **idempotent reconciliation** — running the loop many times has the same result as running it once.

---

### 16. What is the API Server?

"The API Server is the **gateway to the entire Kubernetes cluster**. Every operation — creating a pod, reading resources, deleting namespaces — goes through it.

It validates and authenticates requests, runs them through admission controllers (for extra validation/mutation), persists them to etcd, and notifies watchers.

In companies I've worked at, the API Server is load-balanced behind an NLB, and its audit logs feed into SIEM tools for compliance and security monitoring."

#### In Depth
The API server processes requests through a pipeline: **Authentication → Authorization (RBAC) → Admission Controllers → Validation → Persistence**. Admission controllers are critical for policy enforcement. Mutating webhooks run before validating webhooks, allowing you to inject defaults (like sidecar containers) before validation.

---

### 17. What is a worker node?

"A worker node is where your **actual application containers run**. It hosts the kubelet, kube-proxy, and container runtime.

The node registers itself with the control plane, receives pod assignments, runs them, and reports back health and resource usage.

In cloud environments, worker nodes are typically managed as **Node Groups** or **Node Pools** — groups of similarly configured machines that can be scaled in/out as a unit. You don't SSH into individual nodes — you manage them declaratively."

#### In Depth
Worker nodes are assigned a **CIDR block** for pod networking. The CNI plugin allocates individual IPs from this block to pods. Nodes report `Allocatable` resources (actual capacity minus OS/system overhead) which the scheduler uses for placement decisions.

---

### 18. What are DaemonSets?

"A DaemonSet ensures that **exactly one pod runs on every node** (or a subset of nodes via node selectors).

I use DaemonSets for infrastructure-level agents that need to run on every machine: log collectors (Fluentd/Fluent Bit), monitoring agents (Prometheus node-exporter), security agents (Falco), or CNI plugins themselves.

When a new node joins the cluster, the DaemonSet controller automatically schedules the pod on it. When a node is removed, the pod is garbage-collected."

#### In Depth
By default, DaemonSet pods can be scheduled on control plane nodes too if you add the appropriate toleration (`node-role.kubernetes.io/control-plane:NoSchedule`). DaemonSets support rolling updates via `updateStrategy: RollingUpdate`, making it safe to update node agents without downtime.

---

### 19. What is kube-proxy?

"kube-proxy runs on every node and is responsible for **implementing Kubernetes Service networking**.

When you create a Service, kube-proxy programs networking rules (iptables or IPVS) on each node so that traffic to the Service's ClusterIP gets load-balanced to the backing Pods.

In modern clusters I prefer **IPVS mode** over iptables because it uses hashing rather than scanning chain rules — it scales much better when you have tens of thousands of services."

#### In Depth
kube-proxy doesn't proxy traffic in the traditional sense — it doesn't sit in the data path. Instead, it programs kernel-level rules (iptables/IPVS) that intercept packets and DNAT them to the correct pod IP. With CNI plugins like Cilium, kube-proxy can be replaced entirely with eBPF-based networking, which is even more efficient.

---

### 20. How does service discovery work in Kubernetes?

"Kubernetes provides **DNS-based service discovery** out of the box via CoreDNS.

Every Service gets a DNS name of the form `<service>.<namespace>.svc.cluster.local`. Pods can reach any service by name — e.g., `http://payments-service` (if in the same namespace) or `http://payments-service.finance.svc.cluster.local` (cross-namespace).

I rely on this heavily in microservice architectures. No hardcoded IPs, no custom discovery mechanism needed — just standard DNS."

#### In Depth
CoreDNS watches the API server for Service changes and updates its DNS records in real time. The `resolv.conf` in each pod is pre-configured with the cluster's DNS server and search domains. For headless Services (`.spec.clusterIP: None`), CoreDNS returns individual pod IPs instead of the service VIP, enabling direct pod addressing — critical for StatefulSets.

---

### 21. What is the difference between `kubectl apply` and `kubectl create`?

"`kubectl create` is **imperative** — it creates a resource and fails if it already exists.

`kubectl apply` is **declarative** — it creates or updates a resource based on the manifest, tracking changes via a `last-applied-configuration` annotation.

In all my CI/CD pipelines, I exclusively use `kubectl apply` because it's idempotent. `kubectl create` is fine for quick one-offs in development, but breaks automation workflows when a resource already exists."

#### In Depth
`kubectl apply` uses a **three-way merge patch**: it compares the current live object, the previous applied manifest, and the new manifest to compute a minimal diff. This is why it can intelligently handle cases where the live object was modified by the cluster (e.g., by an autoscaler) and still apply your intent correctly.

---

### 22. What is Minikube vs Kind vs K3s?

"These are three different ways to run K8s locally or at the edge:

- **Minikube**: Creates a full K8s cluster in a local VM or Docker. Easy to start, rich add-on ecosystem. Best for local development.
- **Kind** (Kubernetes in Docker): Runs K8s nodes as Docker containers. Extremely fast startup, ideal for CI pipelines.
- **K3s**: A lightweight, production-grade K8s by Rancher. Replaces etcd with SQLite by default, strips heavy controllers. Used in edge, IoT, and resource-constrained environments.

For CI I use kind. For edge deployments I use K3s. For developer laptops, either Minikube or kind depending on team preference."

#### In Depth
K3s is a CNCF Sandbox project. It runs in under 512MB RAM and is a single binary. It's perfect for environments where running full K8s would be wasteful. K3s supports multiple backends (SQLite, embedded etcd, external Postgres/MySQL) making it flexible for different scale requirements.

---
