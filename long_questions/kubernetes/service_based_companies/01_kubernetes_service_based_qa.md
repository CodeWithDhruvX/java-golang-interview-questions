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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the relationship between Docker and Kubernetes. Are they competitors?
**Your Response:** "Docker and Kubernetes are not competitors - they work together perfectly. Think of Docker as the tool that creates individual containers, like building bricks. Kubernetes is the master architect that manages hundreds of these containers across multiple servers, ensuring they're running, scaled properly, and can communicate with each other. While Docker has its own orchestration tool called Docker Swarm, most companies use Docker for creating containers and Kubernetes for managing them at scale."

---

### Question 2: Explain the main components of the Kubernetes Control Plane (Master Node).

**Answer:**
The Control Plane makes global decisions about the cluster and detects/responds to cluster events. Its components include:
1. **kube-apiserver:** The frontend for the Kubernetes control plane. All components communicate via this REST API.
2. **etcd:** A highly available key-value store used as the single source of truth for all cluster data and state.
3. **kube-scheduler:** Watches for newly created Pods with no assigned node, and selects a node for them to run on based on resource requirements and constraints.
4. **kube-controller-manager:** Runs controller processes (like Node Controller, Job Controller, ReplicaSet Controller) that constantly watch the state and move the current state towards the desired state.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the main components of the Kubernetes Control Plane.
**Your Response:** "The Control Plane is the brain of the Kubernetes cluster. It has four key components: First, the API server acts as the front door - everything communicates through it. Second, etcd is the memory - it stores all cluster data reliably. Third, the scheduler is like a matchmaker - it decides which node should run each new pod based on resources. Finally, the controller manager is like the maintenance crew - it constantly watches everything and fixes issues to keep the cluster healthy."

---

### Question 3: What is the role of a kubelet?

**Answer:**
The `kubelet` is the primary "node agent" that runs on every worker node in the cluster.
- It registers the node with the API server.
- It receives PodSpecs (pod definitions) from the API server.
- It ensures that the containers described in those PodSpecs are up, running, and healthy by talking to the container runtime.
- It reports the status of the node and the pods back to the control plane.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of a kubelet?
**Your Response:** "The kubelet is like the foreman on each worker node. It's the agent that actually runs on every server in the cluster. Its job is to receive instructions from the control plane about what containers should be running, then make sure those containers are actually running and healthy. It also reports back to the control plane about the node's status - like saying 'Hey, I'm alive and here's what's running on me.'"

---

### Question 4: How does a ReplicaSet differ from a Deployment?

**Answer:**
- **ReplicaSet:** Its sole purpose is to maintain a stable set of replica Pods running at any given time. However, it does not support rolling updates or rollbacks natively.
- **Deployment:** A higher-level concept that manages ReplicaSets. Deployments provide declarative updates for Pods. When you update a Deployment (e.g., change the image version), it automatically creates a new ReplicaSet and scales it up while scaling down the old ReplicaSet to ensure a zero-downtime rolling update.
**Best Practice:** Always use Deployments instead of managing ReplicaSets directly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a ReplicaSet differ from a Deployment?
**Your Response:** "Think of a ReplicaSet as a simple tool that just keeps a certain number of pods running - if one dies, it starts another. But a Deployment is much smarter - it's like having a version control system for your pods. When you want to update your application, the Deployment creates new pods with the new version, gradually scales them up while scaling down the old ones, and can even rollback if something goes wrong. Always use Deployments - they give you rolling updates and rollbacks for free."

---

### Question 5: What is a Pod and why does Kubernetes use Pods instead of running containers directly?

**Answer:**
A Pod is the smallest deployable computing unit in Kubernetes. It encapsulates one or more containers.
**Why Pods?**
1. **Shared Resources:** Containers inside the same Pod share the same Network namespace (IP address and `localhost`) and can easily share Storage Volumes.
2. **Co-scheduling:** Certain background tasks (like a logging sidecar agent) need to be tightly coupled and co-located on the same physical VM as the main application container. Pods group them logically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Pod and why does Kubernetes use Pods instead of running containers directly?
**Your Response:** "A Pod is like a small apartment building - it's the smallest unit Kubernetes manages, but it can hold one or more containers as roommates. The reason Kubernetes uses pods instead of individual containers is for efficiency. Containers in the same pod share the same IP address and can talk to each other via localhost, just like roommates sharing a kitchen. This is perfect for when you have a main application container that needs a helper container - like a logging agent - that needs to be tightly coupled and always together."

---

## 🔹 Services & Networking (Questions 6-10)

### Question 6: What is a Kubernetes Service and why is it needed?

**Answer:**
Pods are ephemeral—they are created, destroyed, and their IP addresses change dynamically.
A **Service** is an abstraction that defines a logical set of Pods (identified via labels) and provides them with a stable, static IP address and a DNS name. It ensures that traffic reaches the active Pods, providing built-in load balancing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Kubernetes Service and why is it needed?
**Your Response:** "Pods are like temporary workers - they come and go, and their addresses keep changing. A Service is like giving your application a permanent phone number. Even when individual pods die and get replaced, the service keeps the same stable IP address and DNS name. It automatically routes traffic to whatever healthy pods are currently running, and even load balances between them. Without services, you'd have to constantly update your application every time a pod's IP changed."

---

### Question 7: Explain the different types of Services in Kubernetes.

**Answer:**
1. **ClusterIP (Default):** Exposes the service on a cluster-internal IP. It makes the service only reachable from within the cluster.
2. **NodePort:** Exposes the service on a static port (between 30000-32767) on each Worker Node's IP. It allows external access to the service.
3. **LoadBalancer:** Uses the cloud provider's external load balancer (like AWS ELB) to expose the service to the internet.
4. **ExternalName:** Maps the service to a DNS name (e.g., `foo.external.com`) by returning a CNAME record.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the different types of Services in Kubernetes.
**Your Response:** "Kubernetes offers four types of services. ClusterIP is the default - it's like an internal extension, only accessible within the cluster. NodePort opens a specific port on every node, like putting your service on a public extension that anyone can call. LoadBalancer is the premium option - it provisions a real cloud load balancer to give your service a proper public IP. ExternalName is special - it's like having a forwarding service that just points to another DNS name outside the cluster."

---

### Question 8: How is an Ingress different from a LoadBalancer Service?

**Answer:**
- **LoadBalancer Service:** Operates at Layer 4 (TCP/UDP). For every service you expose, it creates a new cloud load balancer, which can get expensive.
- **Ingress:** Operates at Layer 7 (HTTP/HTTPS). It is an API object that provides routing rules to manage external access to multiple services in a cluster. You only need one Cloud Load Balancer for the Ingress Controller, which then performs path-based or host-based routing (e.g., `/api` -> api-service, `/web` -> web-service).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How is an Ingress different from a LoadBalancer Service?
**Your Response:** "The key difference is that LoadBalancer works at Layer 4 - it's like a simple traffic cop that just forwards connections. Every LoadBalancer service creates its own expensive cloud load balancer. Ingress works at Layer 7 - it's like a smart receptionist that can read HTTP requests and route them based on the URL path or hostname. The best part is you only need one external load balancer for the Ingress controller, and it can handle routing for dozens of services inside your cluster, saving you money and giving you smart routing capabilities."

---

### Question 9: What is the purpose of kube-proxy?

**Answer:**
`kube-proxy` is a network proxy that runs on each node in your cluster. It implements the concept of Kubernetes Services.
It maintains network rules (using iptables or IPVS) on the host node, which allow network communication directly to the Pods from network sessions inside or outside of your cluster.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the purpose of kube-proxy?
**Your Response:** "kube-proxy is like the network traffic controller on every node. When you create a service with a stable IP, kube-proxy is what makes the magic happen on each server. It sets up network rules using iptables that say 'any traffic coming to this service IP should be forwarded to these actual pod IPs.' It's what enables services to work - without kube-proxy, the service IP would just be a dead end. It runs on every node and constantly updates the rules as pods come and go."

---

### Question 10: How do containers within the same pod communicate with each other?

**Answer:**
Because containers within the same Pod share the same network namespace, they can communicate with each other using standard inter-process communications (IPC) like SystemV semaphores or POSIX shared memory, and most commonly over **`localhost`**. For example, a web app container on port 8080 can talk to a sidecar container on port 9090 via `http://localhost:9090`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do containers within the same pod communicate with each other?
**Your Response:** "Containers in the same pod are like roommates sharing an apartment - they share everything! Most importantly, they share the same network namespace, which means they can talk to each other using localhost. It's like having two applications on the same computer - one on port 8080 can call the other on port 9090 just by using localhost:9090. They don't need to go through any complex networking because they're literally on the same network stack. This is why pods are so useful for tightly coupled applications."

---

## 🔹 Configs, Storage, & Troubleshooting (Questions 11-15)

### Question 11: How do you handle application configurations and sensitive data in K8s?

**Answer:**
Using **ConfigMaps** and **Secrets**.
- **ConfigMaps:** Used to store non-confidential data in key-value pairs (e.g., environment variables, settings files).
- **Secrets:** Used to store sensitive information like passwords, OAuth tokens, and SSH keys. They are base64-encoded.
Both can be injected into a Pod either as environment variables or mounted as files in a Volume.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle application configurations and sensitive data in K8s?
**Your Response:** "Kubernetes provides two main ways to handle configuration. For regular settings like database connection strings or feature flags, we use ConfigMaps - they're like configuration files that can be injected as environment variables or mounted as files. For sensitive data like passwords or API keys, we use Secrets - they're similar to ConfigMaps but are designed for sensitive information and have additional security features. Both can be attached to pods as environment variables or as files, so your application doesn't need to know about Kubernetes - it just reads from environment variables or files like normal."

---

### Question 12: Tell me the difference between a PersistentVolume (PV) and a PersistentVolumeClaim (PVC).

**Answer:**
- **PersistentVolume (PV):** Represents a piece of storage in the cluster (like an AWS EBS volume or NFS drive) provisioned by an administrator or dynamically. It is the actual physical resource.
- **PersistentVolumeClaim (PVC):** Represents a request for storage by a user/Pod. It claims a specific size and access mode from a PV. Pods mount a PVC, and Kubernetes binds the PVC to a suitable PV behind the scenes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Tell me the difference between a PersistentVolume (PV) and a PersistentVolumeClaim (PVC).
**Your Response:** "Think of it like renting a storage unit. The PersistentVolume is the actual storage unit - it's the physical disk space, whether it's an AWS EBS volume or an NFS share. The PersistentVolumeClaim is like your rental agreement - you say 'I need 10GB of storage with read-write access.' Kubernetes then matches your claim to an available storage unit. Your pods don't connect directly to the storage - they connect to the claim, and Kubernetes handles the plumbing behind the scenes. This separation lets storage admins manage the actual storage while developers just request what they need."

---

### Question 13: What happens if a worker node crashes?

**Answer:**
1. The `kube-controller-manager` (specifically the Node Controller) will notice the node isn't sending heartbeats.
2. The node is marked as `NotReady`.
3. After a timeout period (typically 5 minutes), the Pods on that node will be marked for eviction.
4. If those Pods are managed by a Deployment/ReplicaSet, new Pods will automatically be scheduled onto different, healthy worker nodes to maintain the desired replica count.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens if a worker node crashes?
**Your Response:** "Kubernetes is designed for failures. When a worker node crashes, the control plane notices it's not sending heartbeats anymore and marks it as NotReady. After about 5 minutes, Kubernetes assumes the node is dead and starts evicting the pods from it. If those pods are managed by a Deployment - which they should be - Kubernetes automatically creates replacement pods on other healthy nodes to maintain the desired number of replicas. This self-healing capability is one of Kubernetes' biggest strengths - applications recover automatically from node failures."

---

### Question 14: Explain Liveness and Readiness probes.

**Answer:**
Probes are used by the kubelet to understand the health of a container.
- **Readiness Probe:** Checks if the app is ready to accept traffic. If it fails, Kubernetes removes the Pod's IP from the Service endpoints (stops sending user traffic), but does not kill the container.
- **Liveness Probe:** Checks if the app is running smoothly. If the application is deadlocked or frozen and fails the liveness probe, Kubernetes will automatically restart the container.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Liveness and Readiness probes.
**Your Response:** "These are health checks that tell Kubernetes about your application's state. The readiness probe is like a 'ready for business' sign - it checks if your app is actually ready to handle traffic. If it fails, Kubernetes stops sending new requests to that pod but doesn't kill it. The liveness probe is like a pulse check - it verifies your app is still alive and not frozen. If the liveness probe fails, Kubernetes restarts the container. Both are crucial - readiness prevents sending traffic to apps that aren't ready, and liveness automatically recovers from crashes or deadlocks."

---

### Question 15: What are the common kubectl commands you use to troubleshoot a failing Pod?

**Answer:**
1. `kubectl get pods`: Check the status (is it CrashLoopBackOff, ImagePullBackOff, Pending?).
2. `kubectl describe pod <pod-name>`: Check the Events section at the bottom for scheduling errors or probe failures.
3. `kubectl logs <pod-name>`: View the standard output/error logs of the application.
4. `kubectl logs <pod-name> -p`: View logs from the *previous* crashed instance of the container.
5. `kubectl exec -it <pod-name> -- /bin/sh`: Open a shell inside the container to manually check network connectivity or configuration files.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the common kubectl commands you use to troubleshoot a failing Pod?
**Your Response:** "When a pod is failing, I follow a systematic approach. First, `kubectl get pods` to see the status - is it CrashLoopBackOff or Pending? Then `kubectl describe pod` to see the events section, which often tells you exactly what went wrong. Next, `kubectl logs` to see the application's own error messages. If the pod keeps crashing, I use `kubectl logs -p` to see logs from the previous crash. Finally, if I need to dig deeper, I'll `kubectl exec` into the container to manually check things like network connectivity or configuration files. This usually covers 90% of pod issues."

---
