# 🏢 Kubernetes Interview Questions - Service-Based Companies
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, etc.
> **Focus:** Core K8s concepts, Deployments, Services, ConfigMaps, daily management commands, and basic architecture.

---

## 🔹 Kubernetes Basics & Architecture (Questions 1-5)

### Question 1: Explain the relationship between Docker and Kubernetes. Are they competitors?

**Answer:**
No, they are highly complementary technologies, though they serve different purposes.
- **Docker** is a platform to build, distribute, and run containers. It packages the application and its environment. (Analogy: A single brick).
- **Kubernetes** is an orchestration system for managing containerized applications across a cluster of nodes. It manages the lifecycle, scaling, and networking of these containers. (Analogy: A building managed by an architect).
While Docker Swarm is a competitor to Kubernetes, Docker itself is widely used as the Container Runtime Interface (CRI) underlying Kubernetes (though now containerd is often used directly).

---

### Question 2: Explain the main components of the Kubernetes Control Plane (Master Node).

**Answer:**
The Control Plane makes global decisions about the cluster and detects/responds to cluster events. Its components include:
1. **kube-apiserver:** The frontend for the Kubernetes control plane. All components communicate via this REST API.
2. **etcd:** A highly available key-value store used as the single source of truth for all cluster data and state.
3. **kube-scheduler:** Watches for newly created Pods with no assigned node, and selects a node for them to run on based on resource requirements and constraints.
4. **kube-controller-manager:** Runs controller processes (like Node Controller, Job Controller, ReplicaSet Controller) that constantly watch the state and move the current state towards the desired state.

---

### Question 3: What is the role of a kubelet?

**Answer:**
The `kubelet` is the primary "node agent" that runs on every worker node in the cluster.
- It registers the node with the API server.
- It receives PodSpecs (pod definitions) from the API server.
- It ensures that the containers described in those PodSpecs are up, running, and healthy by talking to the container runtime.
- It reports the status of the node and the pods back to the control plane.

---

### Question 4: How does a ReplicaSet differ from a Deployment?

**Answer:**
- **ReplicaSet:** Its sole purpose is to maintain a stable set of replica Pods running at any given time. However, it does not support rolling updates or rollbacks natively.
- **Deployment:** A higher-level concept that manages ReplicaSets. Deployments provide declarative updates for Pods. When you update a Deployment (e.g., change the image version), it automatically creates a new ReplicaSet and scales it up while scaling down the old ReplicaSet to ensure a zero-downtime rolling update.
**Best Practice:** Always use Deployments instead of managing ReplicaSets directly.

---

### Question 5: What is a Pod and why does Kubernetes use Pods instead of running containers directly?

**Answer:**
A Pod is the smallest deployable computing unit in Kubernetes. It encapsulates one or more containers.
**Why Pods?**
1. **Shared Resources:** Containers inside the same Pod share the same Network namespace (IP address and `localhost`) and can easily share Storage Volumes.
2. **Co-scheduling:** Certain background tasks (like a logging sidecar agent) need to be tightly coupled and co-located on the same physical VM as the main application container. Pods group them logically.

---

## 🔹 Services & Networking (Questions 6-10)

### Question 6: What is a Kubernetes Service and why is it needed?

**Answer:**
Pods are ephemeral—they are created, destroyed, and their IP addresses change dynamically.
A **Service** is an abstraction that defines a logical set of Pods (identified via labels) and provides them with a stable, static IP address and a DNS name. It ensures that traffic reaches the active Pods, providing built-in load balancing.

---

### Question 7: Explain the different types of Services in Kubernetes.

**Answer:**
1. **ClusterIP (Default):** Exposes the service on a cluster-internal IP. It makes the service only reachable from within the cluster.
2. **NodePort:** Exposes the service on a static port (between 30000-32767) on each Worker Node's IP. It allows external access to the service.
3. **LoadBalancer:** Uses the cloud provider's external load balancer (like AWS ELB) to expose the service to the internet.
4. **ExternalName:** Maps the service to a DNS name (e.g., `foo.external.com`) by returning a CNAME record.

---

### Question 8: How is an Ingress different from a LoadBalancer Service?

**Answer:**
- **LoadBalancer Service:** Operates at Layer 4 (TCP/UDP). For every service you expose, it creates a new cloud load balancer, which can get expensive.
- **Ingress:** Operates at Layer 7 (HTTP/HTTPS). It is an API object that provides routing rules to manage external access to multiple services in a cluster. You only need one Cloud Load Balancer for the Ingress Controller, which then performs path-based or host-based routing (e.g., `/api` -> api-service, `/web` -> web-service).

---

### Question 9: What is the purpose of kube-proxy?

**Answer:**
`kube-proxy` is a network proxy that runs on each node in your cluster. It implements the concept of Kubernetes Services.
It maintains network rules (using iptables or IPVS) on the host node, which allow network communication directly to the Pods from network sessions inside or outside of your cluster.

---

### Question 10: How do containers within the same pod communicate with each other?

**Answer:**
Because containers within the same Pod share the same network namespace, they can communicate with each other using standard inter-process communications (IPC) like SystemV semaphores or POSIX shared memory, and most commonly over **`localhost`**. For example, a web app container on port 8080 can talk to a sidecar container on port 9090 via `http://localhost:9090`.

---

## 🔹 Configs, Storage, & Troubleshooting (Questions 11-15)

### Question 11: How do you handle application configurations and sensitive data in K8s?

**Answer:**
Using **ConfigMaps** and **Secrets**.
- **ConfigMaps:** Used to store non-confidential data in key-value pairs (e.g., environment variables, settings files).
- **Secrets:** Used to store sensitive information like passwords, OAuth tokens, and SSH keys. They are base64-encoded.
Both can be injected into a Pod either as environment variables or mounted as files in a Volume.

---

### Question 12: Tell me the difference between a PersistentVolume (PV) and a PersistentVolumeClaim (PVC).

**Answer:**
- **PersistentVolume (PV):** Represents a piece of storage in the cluster (like an AWS EBS volume or NFS drive) provisioned by an administrator or dynamically. It is the actual physical resource.
- **PersistentVolumeClaim (PVC):** Represents a request for storage by a user/Pod. It claims a specific size and access mode from a PV. Pods mount a PVC, and Kubernetes binds the PVC to a suitable PV behind the scenes.

---

### Question 13: What happens if a worker node crashes?

**Answer:**
1. The `kube-controller-manager` (specifically the Node Controller) will notice the node isn't sending heartbeats.
2. The node is marked as `NotReady`.
3. After a timeout period (typically 5 minutes), the Pods on that node will be marked for eviction.
4. If those Pods are managed by a Deployment/ReplicaSet, new Pods will automatically be scheduled onto different, healthy worker nodes to maintain the desired replica count.

---

### Question 14: Explain Liveness and Readiness probes.

**Answer:**
Probes are used by the kubelet to understand the health of a container.
- **Readiness Probe:** Checks if the app is ready to accept traffic. If it fails, Kubernetes removes the Pod's IP from the Service endpoints (stops sending user traffic), but does not kill the container.
- **Liveness Probe:** Checks if the app is running smoothly. If the application is deadlocked or frozen and fails the liveness probe, Kubernetes will automatically restart the container.

---

### Question 15: What are the common kubectl commands you use to troubleshoot a failing Pod?

**Answer:**
1. `kubectl get pods`: Check the status (is it CrashLoopBackOff, ImagePullBackOff, Pending?).
2. `kubectl describe pod <pod-name>`: Check the Events section at the bottom for scheduling errors or probe failures.
3. `kubectl logs <pod-name>`: View the standard output/error logs of the application.
4. `kubectl logs <pod-name> -p`: View logs from the *previous* crashed instance of the container.
5. `kubectl exec -it <pod-name> -- /bin/sh`: Open a shell inside the container to manually check network connectivity or configuration files.

---
