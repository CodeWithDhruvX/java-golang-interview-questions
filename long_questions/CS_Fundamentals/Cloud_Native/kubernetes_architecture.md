# Kubernetes Architecture

## 1. High Level Overview
Kubernetes (K8s) follows a **Master-Slave** (or Control Plane - Worker Node) architecture.
*   **Control Plane**: The brain. Manages the cluster state.
*   **Worker Nodes**: The brawn. Runs the actual applications (Pods).

## 2. Control Plane Components (The Master)
These components usually run on a separate server (or set of servers for HA).

### A. kube-apiserver
*   **Role**: The Front Door.
*   **Function**:
    *   Exposes the K8s API (REST).
    *   Authenticates/Authorizes requests (kubectl, UI).
    *   The *only* component that talks to `etcd`.
    *   All other components talk to the API Server, they never talk to each other directly.

### B. etcd
*   **Role**: The Database (Brain).
*   **Function**: A highly available Key-Value store. It stores the *entire* configuration and state of the cluster (Secrets, Pods, ConfigMaps).
*   **Protocol**: Uses **Raft** consensus algorithm for consistency.

### C. kube-scheduler
*   **Role**: The Decision Maker.
*   **Function**: Watches for new Pods with no assigned Node. It selects the best Node for the Pod based on resources (CPU/RAM), taints/tolerations, and affinity rules.

### D. kube-controller-manager
*   **Role**: The Regulator.
*   **Function**: Runs control loops that regulate the state of the system.
    *   *Node Controller*: Nnotices if a node goes down.
    *   *ReplicaSet Controller*: Ensures correct number of pods are running (Current State == Desired State).

## 3. Worker Node Components
These run on every node in the cluster.

### A. kubelet
*   **Role**: The Captain.
*   **Function**: An agent that runs on each node. It receives instructions from the API Server ("Run this Pod") and ensures the containers are running and healthy.

### B. kube-proxy
*   **Role**: The Networking Guy.
*   **Function**: Maintains network rules on the node (using IPtables or IPVS). It handles routing traffic (Services) to the correct Pods.

### C. Container Runtime
*   **Role**: The Engine.
*   **Function**: Software that actually runs the containers (e.g., **containerd**, **CRI-O**, Docker Engine).

## 4. The Request Flow (Example)
What happens when you run `kubectl run nginx`?
1.  **kubectl** sends HTTP POST to **API Server**.
2.  **API Server** validates req, persists "Desired: 1 Nginx Pod" to **etcd**. Returns ID.
3.  **Scheduler** (watching API) sees new unassigned Pod. Assigns it to `Node-1`. Writes binding to API.
4.  **API Server** writes binding to **etcd**.
5.  **Kubelet** (on `Node-1`, watching API) sees it has been assigned a Pod.
6.  **Kubelet** asks **Container Runtime** (containerd) to start the container.
7.  **Kubelet** reports status "Running" back to API Server.
8.  **API Server** writes status to **etcd**.

## 5. Interview Questions
1.  **What happens if etcd goes down?**
    *   *Ans*: The cluster creates no new changes. Existing pods keep running, but you can't deploy/scale anything. If `etcd` data is corrupt, the cluster is lost. Hence, backup `etcd`.
2.  **Role of Kubelet?**
    *   *Ans*: The primary "node agent". It registers the node with the apiserver and manages the lifecycle of Pods on that specific machine.
3.  **Does Kubernetes use Docker?**
    *   *Ans*: Not directly anymore. It uses any runtime that implements **CRI** (Container Runtime Interface), like `containerd` or `CRI-O`. Docker Engine is just one option (and K8s removed the special "dockershim" layer in v1.24).
