# 📘 04 — Azure Kubernetes Service (AKS) Advanced
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Advanced

---

## 🔑 Must-Know Topics
- AKS Networking (Kubenet vs Azure CNI)
- AKS Auto-scaling (HPA vs Cluster Autoscaler vs KEDA)
- AKS Security (Azure AD Integration, Network Policies)
- Persistent Storage in AKS (CSI drivers)

---

## ❓ Most Asked Questions

### Q1. Compare Kubenet vs Azure CNI in AKS networking.

| Feature | Kubenet (Basic) | Azure CNI (Advanced) |
|---------|-----------------|----------------------|
| **IP Allocation** | Nodes get an Azure VNet IP. Pods get an IP from a logically different address space natively managed by Kubernetes. | Every Pod gets a first-class IP address from the Azure VNet Subnet. |
| **VNet Integration**| Traffic from Pods to outside the VNet uses NAT (Network Address Translation) through the Node IP. | Pods can communicate directly with on-premise networks or ExpressRoute without NAT. |
| **IP Exhaustion** | Conserves VNet IP addresses (only nodes need them). | High risk of VNet IP exhaustion (1 pod = 1 VNet IP). |
| **Complexity** | Simple, standard Kubernetes networking. | Required for advanced Azure integrations, Windows nodes, and Virtual Nodes. |

---

### Q2. How do you scale workloads in AKS?
AKS scaling happens at two levels: the Pod level and the Node level.

1. **HPA (Horizontal Pod Autoscaler):** Watches CPU/Memory metrics (or custom metrics). If CPU exceeds 80%, HPA increases the number of Pod replicas.
2. **KEDA (Kubernetes Event-Driven Autoscaling):** Allows you to scale Pods based on the length of an Azure Service Bus Queue or Azure blob events, pushing scaling down to 0 replicas.
3. **Cluster Autoscaler:** If HPA creates more Pods and the Nodes run out of CPU/Memory to schedule them, the Pods go into a `Pending` state. The Cluster Autoscaler notices this and automatically provisions new underlying Azure VMs (Nodes) to the cluster.

---

### Q3. How do you secure Pod-to-Pod communication within AKS?
By default, all Pods in a Kubernetes cluster can communicate with all other Pods.
To restrict this, you use **Network Policies**.
- You must enable Network Policies during AKS creation (Azure Network Policies or Calico).
- You write YAML manifests that act as firewall rules. For example, explicitly defining that the `Frontend` Pods can only communicate with the `Backend` Pods on port `8080`, and the `Backend` Pods can only talk to `Database` Pods.

---

### Q4. Explain how Azure Active Directory (Entra ID) integrates with AKS.
Integrating AD with AKS provides fine-grained RBAC.
- Instead of using a generic `kubeconfig` admin file, developers must `az login` and authenticate against Azure AD.
- You map Azure AD Groups (e.g., "Backend-Devs") to Kubernetes `RoleBinding` objects.
- If a developer leaves the company and is disabled in Azure AD, they instantly lose access to the AKS cluster via `kubectl`.

---

### Q5. How do you handle persistent storage in AKS for stateful applications?
Kubernetes Pods are ephemeral. If a Pod dies, its data dies.
To persist data, you use **Volumes**:
1. **Azure Disks:** Good for single-pod read/write. Best for databases running in AKS. They mount a Managed Disk to a specific node/pod.
2. **Azure Files:** Good for read/write by multiple pods simultaneously (e.g., a shared assets folder for multiple web servers).

**How it works:**
The developer creates a `PersistentVolumeClaim` (PVC) in YAML. The Azure CSI (Container Storage Interface) driver automatically intercepts the claim, creates the requested Azure Disk in the resource group, and mounts it into the Pod.

---

### Q6. What is the Azure Application Gateway Ingress Controller (AGIC)?
Normally, traffic enters an AKS cluster through an internal NGINX Ingress Controller.
AGIC allows you to bypass NGINX and use **Azure Application Gateway** as the native ingress for the cluster.
**Benefits:**
- Traffic bypasses the cluster's nodes for routing and goes straight from App Gateway to the Pod.
- Offloads SSL TLS termination and WAF (Web Application Firewall) processing to the App Gateway, saving CPU cycles on the AKS worker nodes.
