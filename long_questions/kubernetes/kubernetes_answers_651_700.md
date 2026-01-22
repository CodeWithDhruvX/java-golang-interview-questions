## 🔹 Cluster Lifecycle & Tools (Questions 651-660)

### Question 651: What is kubeadm and when should you use it?

**Answer:**
The standard tool to bootstrap a K8s cluster.
- `kubeadm init`: Starts master.
- `kubeadm join`: Joins workers.
- **Use Case:** Bare metal, VM-based installs, Learning. Not for Managed K8s (EKS).

---

### Question 652: How does Rancher simplify Kubernetes management?

**Answer:**
A management plane for multiple clusters.
- **GUI:** Visualizes workloads.
- **Provisioning:** Can launch clusters on AWS, vSphere, Bare metal (RKE).
- **Auth:** Centralized AD/LDAP proxy for all clusters.

---

### Question 653: What is kops and how does it differ from kubeadm?

**Answer:**
"Kubernetes Operations".
- Focuses on **Cloud** (AWS).
- It provisions the **Infrastructure** (EC2, VPC, ELB) + K8s.
- Kubeadm assumes you already have a machine; Kops creates the machine.

---

### Question 654: What is kind (Kubernetes in Docker)?

**Answer:**
Local dev cluster.
- Nodes are Docker Containers.
- Super fast startup.
- Good for CI testing.

---

### Question 655: How does Minikube work?

**Answer:**
Local single-node cluster.
- Virtualizes (VM/Container).
- Includes addons (Ingress, Dashboard).
- The classic learning tool.

---

### Question 656: What are common tools for creating local K8s clusters?

**Answer:**
- **Kind:** Best for CI / Multi-node simulation on laptop.
- **Minikube:** Best for features/addons.
- **K3d:** Wrapper around K3s in docker. Fast.
- **Docker Desktop:** Built-in K8s single node.

---

### Question 657: How do you use Terraform to manage Kubernetes clusters?

**Answer:**
Two providers:
1.  **AWS/Azure/Google Provider:** To create the cluster (EKS).
2.  **Kubernetes/Helm Provider:** To deploy apps *into* the cluster.
Best practice: Separate these states.

---

### Question 658: What is Crossplane?

**Answer:**
Control Plane to manage Cloud Infrastructure (`kind: RDSInstance`).
- "K8s as a Universal Control Plane".
- Replaces Terraform for some. Applying YAML creates AWS RDS.

---

### Question 659: What are the benefits of GitOps with Crossplane?

**Answer:**
You store infrastructure definition (RDS YAML) in Git alongside App YAML.
- App + Database provisioned together in one sync.
- No "Terraform apply" step needed externally.

---

### Question 660: What is Cluster API (CAPI)?

**Answer:**
Declarative K8s cluster creation.
- Define `Cluster`, `MachineSet` content.
- A Management Cluster provisions distinct Workload Clusters.
- Supports AWS, Azure, vSphere, etc.

---

## 🔹 Advanced Ingress Strategies (Questions 661-670)

### Question 661: How do you manage multi-tenant ingress controllers?

**Answer:**
- **Shared:** One Nginx Controller. Annotations lock down which hostnames are allowed.
- **Dedicated:** Run one Controller Deployment per Tenant Namespace (`--watch-namespace=tenant-a`). Safest isolation.

---

### Question 662: What is the difference between ingressClass and annotation-based ingress?

**Answer:**
- **Annotation:** `kubernetes.io/ingress.class: nginx`. (Deprecated).
- **IngressClass:** Resource API. `spec.ingressClassName: nginx`. Clean way to handle multiple controllers (e.g., Internal vs External LB).

---

### Question 663: What is a canary ingress?

**Answer:**
Nginx/Istio can split traffic.
- `nginx.ingress.kubernetes.io/canary: "true"`
- `nginx.ingress.kubernetes.io/canary-weight: "10"`
- Sends 10% traffic to this specific Ingress object (which points to new service).

---

### Question 664: How do you configure rate limiting at the ingress level?

**Answer:**
Annotations on the Ingress object.
- `nginx.ingress.kubernetes.io/limit-rps: "5"`
- Or global config in the Controller ConfigMap (e.g., uses basic memory table or Redis).

---

### Question 665: What is SSL passthrough in ingress?

**Answer:**
Ingress does **not** terminate TLS.
- Passes encrypted bytes to Backend Pod.
- Pod checks cert.
- Required for mTLS authentication at Pod level.

---

### Question 666: How do you enable backend protocol upgrades (e.g., WebSockets)?

**Answer:**
Commonly supported out of box.
- Nginx needs specific headers (`Upgrade`, `Connection`).
- Annotations: `nginx.org/websocket-services: "my-service"`.

---

### Question 667: How do you handle client IP preservation?

**Answer:**
- **NodePort:** Source NAT obscures IP.
- **Fix:** `externalTrafficPolicy: Local`. Packets only arrive on nodes running the pod. No NAT.
- **Proxy Protocol:** Enable on ELB and Nginx.

---

### Question 668: What’s the impact of large headers on ingress performance?

**Answer:**
Buffers.
- If headers exceed `client-header-buffer-size`, Nginx returns 413/414 or buffers to disk (slow).
- Tune via ConfigMap.

---

### Question 669: What is the difference between L7 and L4 ingress controllers?

**Answer:**
- **L7 (Nginx/ALB):** Understands HTTP path/host.
- **L4 (MetalLB/NLB):** Forwards TCP packets. No path routing.

---

### Question 670: How do you handle gRPC services through ingress?

**Answer:**
gRPC uses HTTP/2.
- Ingress must support HTTP/2 backend.
- Annotation: `nginx.ingress.kubernetes.io/backend-protocol: "GRPC"`.

---

## 🔹 Energy Efficiency & Green Kubernetes (Questions 671-680)

### Question 671: What is energy-aware scheduling in Kubernetes?

**Answer:**
Scheduling pods based on carbon footprint.
- "Schedule batch job in Region A because wind energy is high there now."
- "Pack nodes tighter to shut down unused nodes."

---

### Question 672: What tools help measure energy consumption in clusters?

**Answer:**
- **Kepler:** Uses eBPF to estimate power usage of a Container/Pod.
- **Scaphandre:** Power consumption metrics.

---

### Question 673: How can auto-scaling be optimized for power saving?

**Answer:**
- **Scale to Zero:** Kill dev envs at 6 PM.
- **Downscale:** Use `HPA` aggressively on non-critical workloads.

---

### Question 674: What is Kube-green?

**Answer:**
An operator that automatically shuts down (sets replicas to 0) namespaces during non-business hours (Sleep/Wake).

---

### Question 675: What is idle pod hibernation?

**Answer:**
Suspending the process rather than killing it. (Not natively supported well yet, mostly scale-to-zero).
- **Checkpoint/Restore (CRIU):** Emerging tech allows freezing a pod state to disk.

---

### Question 676: How can you use spot/preemptible instances to reduce cost & energy?

**Answer:**
Spot instances utilize "spare" capacity (which is already powered on/idle in the datacenter).
- Using them is more efficient for the cloud provider and cheaper for you.

---

### Question 677: What is node hibernation and how to automate it?

**Answer:**
Cluster Autoscaler deletes empty nodes.
- **Automation:** Ensure batch jobs run sequentially to allow full cleanup between runs.

---

### Question 678: How do you offload non-critical workloads during peak energy hours?

**Answer:**
**CronJob time windows.**
- API to check "Carbon Intensity". If High, pause/suspend the CronJob.

---

### Question 679: What are the trade-offs in aggressive resource packing?

**Answer:**
- **Green:** Less servers.
- **Risk:** CPU contention, OOM Kills, Latency spikes.

---

### Question 680: What are container-native power optimizations?

**Answer:**
- Compiled languages (Rust/Go) consume less CPU than interpreted (Python/Ruby).
- Smaller images = Less network transfer energy.

---

## 🔹 Day-2 Ops & Maintenance (Questions 681-690)

### Question 681: How do you detect and mitigate resource starvation?

**Answer:**
- **Detection:** High "Pending" pod count. High "Evicted" count.
- **Mitigation:** ResourceQuotas, PriorityClasses (Kill low priority), Add Nodes.

---

### Question 682: What are signs of etcd stress or failure?

**Answer:**
- **Logs:** "took too long to execute" (fsync latency).
- **Metrics:** `etcd_disk_wal_fsync_duration_seconds` > 10ms.
- **Symptoms:** `kubectl` timeouts.

---

### Question 683: How do you recover from etcd data corruption?

**Answer:**
1.  Stop API Server.
2.  Stop Etcd.
3.  Remove data dir.
4.  `etcdctl snapshot restore backup.db`.
5.  Start Etcd.

---

### Question 684: How do you rotate kubelet certificates?

**Answer:**
Enable `RotateKubeletServerCertificate` feature gate.
- Kubelet automatically requests new certs from API (CSR).
- Controller Manager approves them.

---

### Question 685: What are common upgrade pitfalls in managed services?

**Answer:**
- **Deprecated APIs:** App fails to deploy.
- **Addons:** CNI/CoreDNS versions must be upgraded compatible with K8s version.
- **PDBs:** Aggressive PDBs blocking node drain.

---

### Question 686: How do you automate node OS patching?

**Answer:**
- **Kured (Kubernetes Reboot Daemon):** Watches for `/var/run/reboot-required` flag on host. Drains node, reboots, uncordons.
- **Snooze:** Schedule it.

---

### Question 687: How do you audit failed pod startups?

**Answer:**
Query `kube_pod_container_status_waiting_reason`.
- Look for `ImagePullBackOff`, `CrashLoopBackOff`.
- Alert on > 5% failure rate.

---

### Question 688: How do you debug stuck terminating pods?

**Answer:**
Usually a Finalizer issue.
1.  Check `metadata.finalizers`.
2.  If resource attached (Volume) is stuck, Pod waits.
3.  **Force:** `kubectl delete pod x --grace-period=0 --force` (Risky!).

---

### Question 689: How do you manage orphaned PVs?

**Answer:**
- PVs with status `Released` but not `Available`.
- Check Policy (`Retain`?).
- Manual cleanup or script to identify unclaims.

---

### Question 690: How do you update a running container’s environment variables?

**Answer:**
You cannot.
- **Immutable Infrastructure.**
- Must update Deployment spec -> New ReplicaSet -> New Pods.

---

## 🔹 Real-World Design & Trade-offs (Questions 691-700)

### Question 691: How do you decide between monorepo vs multirepo in Kubernetes GitOps?

**Answer:**
- **Monorepo:** Easier shared libs, atomic global refactors. Risk: Huge CI builds, permission creep.
- **Multirepo:** Clear ownership, independent lifecycles. Risk: Dependency hell, fragmentation.

---

### Question 692: What are trade-offs of large multi-tenant clusters?

**Answer:**
- **Pros:** Efficient bin-packing, less management (1 upgrade vs 10).
- **Cons:** Blast radius (Upgrade breaks everyone), Complex RBAC/NetworkPolicy needed, noisy neighbors.

---

### Question 693: What is the best way to manage 1000+ microservices on K8s?

**Answer:**
- **Service Mesh:** Mandatory for observability/traffic control.
- **GitOps:** Mandatory for config management.
- **Platform Team:** Provides "Golden Path" templates/scaffolding.

---

### Question 694: When should you avoid Kubernetes?

**Answer:**
- Simple monolithic app (Use Heroku/Render).
- Static site (Use Vercel/S3).
- Small team with no Ops experience.

---

### Question 695: How do you architect for zero-downtime releases?

**Answer:**
- **RollingUpdates:** With health probes.
- **Database:** Backward compatible schema changes (Add column -> Deploy App -> Remove column).
- **Graceful Shutdown:** `preStop` hooks.

---

### Question 696: What’s the difference between application and infrastructure Helm charts?

**Answer:**
- **App Chart:** Deploys Deployment+Service.
- **Infra Chart:** Deploys Redis/Kafka/Consul. Usually stateful, complex, rarely updated.

---

### Question 697: How do you handle secrets across multiple cloud providers?

**Answer:**
**External Secrets Operator.**
- Provider A (AWS SM) -> Secret A.
- Provider B (Azure KV) -> Secret B.
- ESO syncs both to K8s secrets. App doesn't care.

---

### Question 698: What are cost implications of overprovisioning in K8s?

**Answer:**
- Requests reserve capacity.
- If Request=4CPU but Usage=0.1CPU, you pay for 4CPU node capacity.
- **Waste.** Use VPA or Kubecost to find rightsizing.

---

### Question 699: When is a serverless solution better than Kubernetes?

**Answer:**
- **Spiky traffic:** (0 to 1 million in seconds). Lambda scales faster than K8s/Nodes.
- **Event-driven:** S3 trigger -> Process.
- **Glue code.**

---

### Question 700: How do you future-proof your Kubernetes platform?

**Answer:**
- Stick to standard APIs (Ingress vs proprietary CRDs).
- Use CNCF Graduated projects.
- Automate upgrades.
- Don't build in-house what you can buy/adopt (Don't write your own Scheduler).
