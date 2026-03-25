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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you expose a service running in a custom namespace to the internet?
**Your Response:** "I'd use either an Ingress or LoadBalancer service. For cost efficiency, I prefer Ingress - I create a ClusterIP service inside the custom namespace to group the pods, then create an Ingress resource in that same namespace pointing to the service. The cluster needs an Ingress Controller running, which detects the Ingress and sets up the external routing. This way I can expose multiple services through one load balancer using different paths or hostnames, which is much cheaper than giving every service its own load balancer."

---

### Question 32: What is the difference between an API Gateway and a Kubernetes Ingress?

**Answer:**
- **Kubernetes Ingress:** Acts purely as an entry point into the cluster routing HTTP/HTTPS traffic based on hostname or URL path to internal Services safely.
- **API Gateway (e.g., Kong, Apigee):** A much heavier component. While it routes traffic like an ingress, it is built to manage the APIs themselves: implementing API keys, rate-limiting, monetize API payloads, structure payload transformation, and provide deep analytics. 
*(Note: Some modern ingress controllers like Kong Ingress do act as both!)*

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between an API Gateway and a Kubernetes Ingress?
**Your Response:** "Kubernetes Ingress is like a basic traffic cop - it just routes HTTP requests based on hostnames or paths to the right services. An API Gateway is like a sophisticated security guard and bouncer combined - it does routing but also handles API keys, rate limiting, authentication, and can even transform requests and responses. Ingress is pure L7 routing, while API Gateway is about managing the APIs themselves with security and business logic. Some tools like Kong blur the line by doing both, but traditionally they serve different purposes."

---

### Question 33: Explain Kubernetes Network Policies.

**Answer:**
By default, K8s holds a **"flat" network design**—all Pods can talk to all other Pods freely, even across namespaces. 
A **NetworkPolicy** acts like a firewall (implemented by the CNI, e.g. Calico) to control traffic.
You apply a policy to a subset of pods (using `podSelector`), and specify `Ingress` (incoming) or `Egress` (outgoing) rules defining what IP blocks or other pod labels are allowed connection access. If a policy is applied, all other traffic not defined is implicitly dropped.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Kubernetes Network Policies.
**Your Response:** "By default, Kubernetes is like an open office - anyone can talk to anyone. Network Policies are like putting up firewalls between teams. I can create rules that say 'only pods with label web can talk to pods with label database on port 5432.' The policies are implemented by the network plugin like Calico, and they follow a default deny approach - if traffic isn't explicitly allowed, it's blocked. This is crucial for security in multi-tenant environments where I need to isolate applications from each other."

---

### Question 34: What is `kube-dns` or CoreDNS?

**Answer:**
It is the built-in DNS service for the cluster. When you abstract a set of pods behind a service named `my-database` in namespace `dev`, CoreDNS ensures any pod inside the cluster can resolve `my-database.dev.svc.cluster.local` predictably into the Service's virtual IP (ClusterIP). 
Without CoreDNS, apps would have to securely hardcode external configurations or rely on fixed IPs, breaking K8s dynamic scaling benefits.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `kube-dns` or CoreDNS?
**Your Response:** "CoreDNS is like the phone directory of the Kubernetes cluster. When I create a service called 'my-database' in the 'dev' namespace, CoreDNS makes sure that any pod can simply connect to 'my-database.dev.svc.cluster.local' and it automatically resolves to the service's IP address. Without CoreDNS, applications would need to hardcode IP addresses or configuration files, which would break every time pods get rescheduled. CoreDNS enables service discovery - applications can find each other by name no matter where they're running in the cluster."

---

### Question 35: How could you secure the namespace so that no pods inside it can communicate with other namespaces?

**Answer:**
You create a specific **"Default Deny"** NetworkPolicy scoped to that namespace.
You specify an empty `podSelector: {}` (matching all pods in the namespace) and an empty `ingress: []`. This instructs the CNI to drop all incoming TCP/UDP traffic unless you apply a secondary specific policy allowing traffic from exact sources explicitly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How could you secure the namespace so that no pods inside it can communicate with other namespaces?
**Your Response:** "I'd create a 'default deny' NetworkPolicy in that namespace. I set the pod selector to match all pods in the namespace and leave the ingress rules empty. This tells the network plugin to block all incoming traffic from other namespaces. Then I can create additional policies that explicitly allow only the specific traffic I want - like allowing pods from the web namespace to talk to pods in the api namespace on specific ports. It's a whitelist approach - everything is denied by default, and I only open up what's absolutely necessary."

---

## 🔹 Persistent Data & Storage Management (PV/PVCs) (Questions 36-40)

### Question 36: What is a StorageClass in K8s?

**Answer:**
A **StorageClass** provides a way for administrators to describe the "classes" of storage they offer (e.g., "fast" SSDs vs "slow" HDD block storage). 
Most importantly, a StorageClass allows for **Dynamic Provisioning**. Instead of an admin manually creating PersistentVolume (PV) objects beforehand, when a developer creates a PVC requesting a specific `storageClassName`, Kubernetes automatically reaches out to the cloud provider (like AWS) and provisions the actual drive dynamically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a StorageClass in K8s?
**Your Response:** "StorageClass is like the menu of storage options available in the cluster. I can define different classes like 'fast-ssd' for databases or 'slow-hdd' for backups. The magic is dynamic provisioning - instead of admins manually creating storage volumes beforehand, when a developer requests storage with a specific storage class, Kubernetes automatically calls the cloud provider and creates the volume on demand. It's like ordering storage from a catalog instead of having to pre-provision everything."

---

### Question 37: Explain the main Persistent Volume "Access Modes".

**Answer:**
1. **ReadWriteOnce (RWO):** The volume can be mounted as read-write by a *single node*. (Standard for Databases like MySQL).
2. **ReadOnlyMany (ROX):** The volume can be mounted read-only by *many nodes*. (Standard for static web content).
3. **ReadWriteMany (RWX):** The volume can be mounted as read-write by *many nodes* simultaneously. (Typically requires an NFS share like AWS EFS, not raw block storage like EBS).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the main Persistent Volume "Access Modes".
**Your Response:** "There are three access modes. ReadWriteOnce means only one node can mount it as read-write - perfect for databases where you don't want multiple servers writing to the same data. ReadOnlyMany lets many nodes mount it read-only - great for sharing website content across multiple web servers. ReadWriteMany is the special one - many nodes can read and write simultaneously, but this usually requires network file systems like NFS or cloud solutions like AWS EFS, not regular block storage like EBS volumes."

---

### Question 38: What does the Storage Reclaim Policy do when a PVC is deleted?

**Answer:**
It defines what happens to the underlying PV when the user deletes the PVC:
- **Retain:** The PV is kept, but it is marked as "Released" and requires manual cleanup. Data is kept safe.
- **Delete:** Both the K8s PV object and the underlying physical cloud asset (like the EBS Volume) are permanently deleted.
- **Recycle (Deprecated):** Wipes the data (`rm -rf`) and makes the volume available again.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does the Storage Reclaim Policy do when a PVC is deleted?
**Your Response:** "The reclaim policy decides what happens to the storage when the claim is deleted. Retain is the safe option - the storage stays around but needs manual cleanup, so your data isn't accidentally deleted. Delete is the automatic option - both the Kubernetes object and the actual cloud storage get deleted, which is convenient but risky if you have important data. There used to be a Recycle policy that wiped the data and reused the volume, but it's deprecated now. Most production environments use Retain to prevent accidental data loss."

---

### Question 39: How do you share the exact same storage data between two separate pod configurations running in different Nodes?

**Answer:**
You must use a persistent volume that supports the **ReadWriteMany (RWX)** or **ReadOnlyMany** access mode.
Usually, this is achieved by deploying an NFS server (or using a cloud native distributed filesystem like AWS EFS/Azure Files). You configure a PV and PVC for it, and both Pods mount the exact same PVC claim name inside their deployments.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you share the exact same storage data between two separate pod configurations running in different Nodes?
**Your Response:** "To share storage between pods on different nodes, I need a volume that supports ReadWriteMany access mode. Regular block storage like EBS won't work because it can only be mounted by one node at a time. I'd use either an NFS server or a cloud distributed filesystem like AWS EFS. I create a PersistentVolume and PersistentVolumeClaim for the shared storage, and both pods mount the exact same PVC claim. This way both pods can read and write to the same data regardless of which node they're running on."

---

### Question 40: Can a Pod attach a local hard drive from a worker node directly?

**Answer:**
Yes, by using a `hostPath` volume. This directly exposes a file or directory from the node's filesystem to the Pod.
**Warning:** This is generally considered a severe security risk and bad practice, because if the pod gets destroyed and spun up on a *different* worker node, it loses access to that physical hard drive path! It's usually restricted by Pod Security Admission.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can a Pod attach a local hard drive from a worker node directly?
**Your Response:** "Technically yes, using hostPath volumes, but this is generally a bad idea. HostPath directly exposes a directory from the worker node's filesystem to the pod. The problem is that if the pod gets rescheduled to a different node, it loses access to that physical drive completely. It's also a security risk because the pod could access sensitive system files. Most production clusters disable or heavily restrict hostPath volumes with Pod Security Admission policies. It's better to use proper persistent storage solutions."

---

## 🔹 Auto-Scaling & Daily Maintenance (Questions 41-45)

### Question 41: How do you handle auto-scaling of pods based on CPU consumption?

**Answer:**
By deploying a **Horizontal Pod Autoscaler (HPA)**.
1. Ensure the `metrics-server` is installed in the cluster.
2. Ensure the pods have resource Memory and CPU `requests` explicitly defined (HPA calculates against requests).
3. Create the HPA targeting your Deployment: `kubectl autoscale deployment my-app --cpu-percent=70 --min=2 --max=10`. 
When traffic spikes CPU usage over 70%, HPA increases the replica count automatically up to 10.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle auto-scaling of pods based on CPU consumption?
**Your Response:** "I use the Horizontal Pod Autoscaler (HPA). First, I need to make sure the metrics-server is installed to collect CPU metrics. Then I ensure all my pods have CPU and memory requests defined - HPA uses these as the baseline. I create an HPA targeting my deployment with a target CPU percentage, like 70%, and set min and max replicas. When CPU usage goes above 70%, HPA automatically adds more pods up to my max limit. When traffic drops, it scales back down. It's like having an automatic throttle for your application based on actual load."

---

### Question 42: What is the primary difference between `kubectl apply` and `kubectl create`?

**Answer:**
- **Create:** Imperative approach. It tells K8s to create the resource exactly as specified. If the resource already exists, it throws an error immediately entirely failing out.
- **Apply:** Declarative approach. It compares the given YAML file with the live object running in the cluster. If it doesn't exist, it creates it. If it does exist, it simply calculates the delta (differential) and intelligently patches the object safely to match the new desired state. 

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the primary difference between `kubectl apply` and `kubectl create`?
**Your Response:** "Create is like giving exact orders - it tries to create exactly what you specify, and if something already exists, it just fails. Apply is smarter - it's like showing Kubernetes what you want and letting it figure out how to get there. Apply compares your YAML with what's actually running and makes only the necessary changes. If the resource doesn't exist, it creates it. If it does exist, it updates it. Apply is the preferred approach for GitOps and managing changes over time." 

---

### Question 43: A developer accidentally deployed a pod with `image: wrongnginx:latest`. It shows `ImagePullBackOff`. How do you quickly find the exact error message?

**Answer:**
By running `kubectl describe pod <pod-name>`.
You scroll down to the "Events" section at the very bottom. It will display chronological actions performed by the kubelet and scheduler, showing the exact error output returned from the container registry (like "repository does not exist or unauthorized").

### How to Explain in Interview (Spoken style format)
**Interviewer:** A developer accidentally deployed a pod with `image: wrongnginx:latest`. It shows `ImagePullBackOff`. How do you quickly find the exact error message?
**Your Response:** "I'd run `kubectl describe pod <pod-name>` and scroll down to the Events section at the bottom. The Events section is like a log of everything that happened to the pod - it shows the exact error message from the container registry, like 'image not found' or 'access denied'. This is usually the fastest way to diagnose ImagePullBackOff issues because it gives you the precise reason Kubernetes couldn't pull the image, whether it's a typo in the name, wrong tag, or authentication problem."

---

### Question 44: How does a Kubelet know when to evict a Pod?

**Answer:**
The Kubelet continuously monitors the node's resources (like memory, disk space, and PID pressure). 
If available node memory dips critically below a certain threshold (e.g., 100Mi), the kubelet enters **Pressure** state. It then proactively kills Pods to reclaim resources, starting with pods that are exceeding their requested limits, or those running "BestEffort" without requests entirely.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a Kubelet know when to evict a Pod?
**Your Response:** "The kubelet is like a building manager who constantly monitors available resources. When memory gets critically low - like under 100MB - the kubelet goes into pressure mode and starts evicting pods to free up space. It's smart about which pods to kill first - it starts with pods using more than they requested, then 'best effort' pods that didn't specify any resource requests. This prevents the whole node from crashing when resources run low. It's Kubernetes' self-protection mechanism at the node level."

---

### Question 45: What command can you run to temporarily redirect a local port on your laptop securely to a Pod deep inside the cluster database without exposing it via Service?

**Answer:**
You employ `kubectl port-forward`. 
Example: `kubectl port-forward pod/my-database-pod 5432:5432`.
Run locally on your laptop, this intercepts your local `localhost:5432` traffic, tunnels it via the Kube-API server, securely down to the target pod's port `5432`, bypassing generic internet exposition requirements for safe debugging.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What command can you run to temporarily redirect a local port on your laptop securely to a Pod deep inside the cluster database without exposing it via Service?
**Your Response:** "I'd use `kubectl port-forward`. For example, `kubectl port-forward pod/my-database-pod 5432:5432` creates a secure tunnel from my laptop's port 5432 directly to the pod's port 5432 through the API server. It's perfect for debugging because I can connect to localhost on my machine and reach a pod deep inside the cluster without exposing it to the internet. The traffic goes through the Kubernetes API server, so it's secure and doesn't require creating any services or changing the cluster configuration."
