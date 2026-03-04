# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 3)
> **Target:** TCS, Wipro, Infosys, Cognizant, IBM, Tech Mahindra, etc.
> **Focus:** Services logic, Networking & Network Policies, Storage Management (PV/PVCs), and Scaling.

---

## 🔹 Services & Ingress Networking (Questions 31-35)

### Question 31: How do you expose a service running in a custom namespace to the internet?

**Answer:**
You would use an **Ingress** resource or a **LoadBalancer** Service.
For cost-efficiency and layer-7 routing:
1. Create a `ClusterIP` Service to group the application Pods in the custom namespace.
2. Create an `Ingress` resource in that *same* namespace, pointing to the ClusterIP Service.
3. The cluster must have an Ingress Controller (like Nginx configured with cloud load balancing) running. It will detect the Ingress object and configure external routing.

---

### Question 32: What is the difference between an API Gateway and a Kubernetes Ingress?

**Answer:**
- **Kubernetes Ingress:** Acts purely as an entry point into the cluster routing HTTP/HTTPS traffic based on hostname or URL path to internal Services safely.
- **API Gateway (e.g., Kong, Apigee):** A much heavier component. While it routes traffic like an ingress, it is built to manage the APIs themselves: implementing API keys, rate-limiting, monetize API payloads, structure payload transformation, and provide deep analytics. 
*(Note: Some modern ingress controllers like Kong Ingress do act as both!)*

---

### Question 33: Explain Kubernetes Network Policies.

**Answer:**
By default, K8s holds a **"flat" network design**—all Pods can talk to all other Pods freely, even across namespaces. 
A **NetworkPolicy** acts like a firewall (implemented by the CNI, e.g. Calico) to control traffic.
You apply a policy to a subset of pods (using `podSelector`), and specify `Ingress` (incoming) or `Egress` (outgoing) rules defining what IP blocks or other pod labels are allowed connection access. If a policy is applied, all other traffic not defined is implicitly dropped.

---

### Question 34: What is `kube-dns` or CoreDNS?

**Answer:**
It is the built-in DNS service for the cluster. When you abstract a set of pods behind a service named `my-database` in namespace `dev`, CoreDNS ensures any pod inside the cluster can resolve `my-database.dev.svc.cluster.local` predictably into the Service's virtual IP (ClusterIP). 
Without CoreDNS, apps would have to securely hardcode external configurations or rely on fixed IPs, breaking K8s dynamic scaling benefits.

---

### Question 35: How could you secure the namespace so that no pods inside it can communicate with other namespaces?

**Answer:**
You create a specific **"Default Deny"** NetworkPolicy scoped to that namespace.
You specify an empty `podSelector: {}` (matching all pods in the namespace) and an empty `ingress: []`. This instructs the CNI to drop all incoming TCP/UDP traffic unless you apply a secondary specific policy allowing traffic from exact sources explicitly.

---

## 🔹 Persistent Data & Storage Management (PV/PVCs) (Questions 36-40)

### Question 36: What is a StorageClass in K8s?

**Answer:**
A **StorageClass** provides a way for administrators to describe the "classes" of storage they offer (e.g., "fast" SSDs vs "slow" HDD block storage). 
Most importantly, a StorageClass allows for **Dynamic Provisioning**. Instead of an admin manually creating PersistentVolume (PV) objects beforehand, when a developer creates a PVC requesting a specific `storageClassName`, Kubernetes automatically reaches out to the cloud provider (like AWS) and provisions the actual drive dynamically.

---

### Question 37: Explain the main Persistent Volume "Access Modes".

**Answer:**
1. **ReadWriteOnce (RWO):** The volume can be mounted as read-write by a *single node*. (Standard for Databases like MySQL).
2. **ReadOnlyMany (ROX):** The volume can be mounted read-only by *many nodes*. (Standard for static web content).
3. **ReadWriteMany (RWX):** The volume can be mounted as read-write by *many nodes* simultaneously. (Typically requires an NFS share like AWS EFS, not raw block storage like EBS).

---

### Question 38: What does the Storage Reclaim Policy do when a PVC is deleted?

**Answer:**
It defines what happens to the underlying PV when the user deletes the PVC:
- **Retain:** The PV is kept, but it is marked as "Released" and requires manual cleanup. Data is kept safe.
- **Delete:** Both the K8s PV object and the underlying physical cloud asset (like the EBS Volume) are permanently deleted.
- **Recycle (Deprecated):** Wipes the data (`rm -rf`) and makes the volume available again.

---

### Question 39: How do you share the exact same storage data between two separate pod configurations running in different Nodes?

**Answer:**
You must use a persistent volume that supports the **ReadWriteMany (RWX)** or **ReadOnlyMany** access mode.
Usually, this is achieved by deploying an NFS server (or using a cloud native distributed filesystem like AWS EFS/Azure Files). You configure a PV and PVC for it, and both Pods mount the exact same PVC claim name inside their deployments.

---

### Question 40: Can a Pod attach a local hard drive from a worker node directly?

**Answer:**
Yes, by using a `hostPath` volume. This directly exposes a file or directory from the node's filesystem to the Pod.
**Warning:** This is generally considered a severe security risk and bad practice, because if the pod gets destroyed and spun up on a *different* worker node, it loses access to that physical hard drive path! It's usually restricted by Pod Security Admission.

---

## 🔹 Auto-Scaling & Daily Maintenance (Questions 41-45)

### Question 41: How do you handle auto-scaling of pods based on CPU consumption?

**Answer:**
By deploying a **Horizontal Pod Autoscaler (HPA)**.
1. Ensure the `metrics-server` is installed in the cluster.
2. Ensure the pods have resource Memory and CPU `requests` explicitly defined (HPA calculates against requests).
3. Create the HPA targeting your Deployment: `kubectl autoscale deployment my-app --cpu-percent=70 --min=2 --max=10`. 
When traffic spikes CPU usage over 70%, HPA increases the replica count automatically up to 10.

---

### Question 42: What is the primary difference between `kubectl apply` and `kubectl create`?

**Answer:**
- **Create:** Imperative approach. It tells K8s to create the resource exactly as specified. If the resource already exists, it throws an error immediately entirely failing out.
- **Apply:** Declarative approach. It compares the given YAML file with the live object running in the cluster. If it doesn't exist, it creates it. If it does exist, it simply calculates the delta (differential) and intelligently patches the object safely to match the new desired state. 

---

### Question 43: A developer accidentally deployed a pod with `image: wrongnginx:latest`. It shows `ImagePullBackOff`. How do you quickly find the exact error message?

**Answer:**
By running `kubectl describe pod <pod-name>`.
You scroll down to the "Events" section at the very bottom. It will display chronological actions performed by the kubelet and scheduler, showing the exact error output returned from the container registry (like "repository does not exist or unauthorized").

---

### Question 44: How does a Kubelet know when to evict a Pod?

**Answer:**
The Kubelet continuously monitors the node's resources (like memory, disk space, and PID pressure). 
If available node memory dips critically below a certain threshold (e.g., 100Mi), the kubelet enters **Pressure** state. It then proactively kills Pods to reclaim resources, starting with pods that are exceeding their requested limits, or those running "BestEffort" without requests entirely.

---

### Question 45: What command can you run to temporarily redirect a local port on your laptop securely to a Pod deep inside the cluster database without exposing it via Service?

**Answer:**
You employ `kubectl port-forward`. 
Example: `kubectl port-forward pod/my-database-pod 5432:5432`.
Run locally on your laptop, this intercepts your local `localhost:5432` traffic, tunnels it via the Kube-API server, securely down to the target pod's port `5432`, bypassing generic internet exposition requirements for safe debugging.
