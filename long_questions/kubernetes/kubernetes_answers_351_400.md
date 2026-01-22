## 🔹 Helm – Package Management (Questions 351-360)

### Question 351: What is Helm and how does it help in Kubernetes?

**Answer:**
Helm is the package manager for Kubernetes.
- **Problem:** K8s YAMLs are static and hard to manage for complex apps.
- **Solution:** Helm packages them into **Charts**, allowing parameterization (templating).
- Helps in Sharing, Versioning, and Rollbacks.

---

### Question 352: What are Helm charts?

**Answer:**
A bundle of information necessary to create an instance of a Kubernetes application.
- `Chart.yaml`: Metadata (Version, Name).
- `values.yaml`: Default configuration.
- `templates/`: Go templates that generate manifest files.

---

### Question 353: How do you create a custom Helm chart?

**Answer:**
Command: `helm create my-chart`.
- Creates standard directory structure.
- Edit `templates/` to match your Deployment/Service.
- Edit `values.yaml` to expose variables (Image tag, Replicas).

---

### Question 354: What is the difference between Helm 2 and Helm 3?

**Answer:**
**Tiller** (The server-side component) was removed in Helm 3.
- Helm 2: Client -> Tiller (Pod with massive rights) -> API Server. (Security risk).
- Helm 3: Client -> API Server directly. (Uses user's RBAC).

---

### Question 355: What is a values.yaml file used for?

**Answer:**
Defines the default values for the templates.
- Users override these values at install time.
- `replicaCount: 1` in values.yaml.
- User runs `--set replicaCount=3`.

---

### Question 356: What is Helm templating?

**Answer:**
Uses Go Template language (`{{ .Values.key }}`).
- Logic: `{{ if .Values.enabled }}`.
- Loops: `{{ range .Values.ingress.hosts }}`.
Generates valid YAML string which is then sent to K8s.

---

### Question 357: How do you pass values to a Helm chart at install time?

**Answer:**
1.  **CLI Flag:** `helm install --set image.tag=v2`.
2.  **Values File:** `helm install -f production-values.yaml`.
3.  **Environment:** `helm install` (uses logic inside chart).

---

### Question 358: How do you upgrade an existing Helm release?

**Answer:**
`helm upgrade <release-name> <chart-name>`.
- Helm calculates the diff.
- Patches the K8s resources.
- Updates the release version history (Secret).

---

### Question 359: What is a Helm release?

**Answer:**
An instance of a chart running in a Kubernetes cluster.
- One Chart (Redis) can be installed twice to create two Releases (`redis-cache`, `redis-queue`).

---

### Question 360: How do you roll back a Helm deployment?

**Answer:**
`helm rollback <release-name> <revision>`.
- `helm rollback my-app 1`.
- Reverts the K8s resources to the state defined in Revision 1.

---

## 🔹 Troubleshooting & Debugging (Questions 361-370)

### Question 361: How do you debug CrashLoopBackOff errors?

**Answer:**
1.  **Logs:** `kubectl logs pod`.
2.  **Exit Code:** `kubectl describe pod` -> Look for `Exit Code`.
    - 0: App finished (but restarted because restartPolicy=Always).
    - 1: App Error.
    - 137: OOMKilled.
    - 139: Segfault.
3.  **Liveness Probe:** Is the probe killing it?

---

### Question 362: What does `kubectl describe pod` show?

**Answer:**
- **Events:** Most useful (Scheduling failures, Image Pull errors).
- **Status:** Phase, Conditions.
- **Config:** IPs, Mounts, Environment vars.
- **Controllers:** Who manages this pod (ReplicaSet).

---

### Question 363: How do you trace network issues in Kubernetes?

**Answer:**
1.  **Connectivity:** `kubectl exec -it podA -- curl -v podB`.
2.  **DNS:** `nslookup my-service`.
3.  **NetworkPolicies:** Check if traffic is blocked.
4.  **Kube-proxy:** Check iptables rules on node.
5.  **Tools:** `netshoot` container.

---

### Question 364: What logs do you check for API server issues?

**Answer:**
On the Master Node:
- `/var/log/kube-apiserver.log`.
- `journalctl -u kube-apiserver`.
- Or looking for the Static Pod logs in `/var/log/containers/kube-apiserver*`.

---

### Question 365: How do you diagnose a hung pod?

**Answer:**
1.  **Logs:** Check for deadlock messages.
2.  **Metrics:** Is CPU at 0%?
3.  **Exec:** `kubectl exec` and install `strace` or `gdb` (if permitted) or take a Thread Dump (Java/Go).

---

### Question 366: What does the `kubectl top` command show?

**Answer:**
Current resource usage.
- `kubectl top nodes`: CPU/RAM of nodes.
- `kubectl top pods`: CPU/RAM of pods.
- Requires `metrics-server`.

---

### Question 367: How do you find out which node a pod is scheduled on?

**Answer:**
`kubectl get pod -o wide`.
Column `NODE` shows the assignment.

---

### Question 368: What are common reasons for pod pending state?

**Answer:**
1.  **Insufficent Resources:** Cluster is full.
2.  **Taints:** Node has taint, pod has no toleration.
3.  **Affinity:** Constraints cannot be met.
4.  **PVC:** Storage not ready.

---

### Question 369: What is the difference between a warning and an error event?

**Answer:**
- **Warning:** Something went wrong (Probe failed, Image pull failed, Node disk pressure). Needs attention.
- **Normal (Info):** State change (Started, Pulled, Scheduled).

---

### Question 370: What does `ImagePullBackOff` mean?

**Answer:**
Kubelet failed to pull the image and is retrying with exponential backoff (10s, 20s, 40s...).
- **Causes:** Typo in image name, Tag doesn't exist, Authentication failed (Private Registry), Network down.

---

## 🔹 Advanced Networking & DNS (Questions 371-380)

### Question 371: How does service discovery happen in Kubernetes?

**Answer:**
Primarily via **DNS**.
- CoreDNS watches API for new Services.
- Adds A Record `my-svc.my-ns.svc.cluster.local`.
- Pods query this name.

---

### Question 372: How do you create an internal-only service?

**Answer:**
By default, `ClusterIP` services are internal-only.
They are not reachable from outside the cluster unless you create Ingress or change type to NodePort/LoadBalancer.

---

### Question 373: What is kube-proxy and how does it route traffic?

**Answer:**
It runs on every node.
- **Iptables Mode:** Randomly picks backend pod. Sets up DNAT rule.
- **IPVS Mode:** Uses Linux IPVS kernel module (Hash table). Better scale (O(1)).

---

### Question 374: What is hairpin NAT in Kubernetes?

**Answer:**
Allows a Pod to access itself via the Service IP.
- Pod A -> Service IP -> Pod A.
- Requires "Hairpin Mode" (Promiscuous bridge mode) to loop packet back to same interface.

---

### Question 375: What is ipVS mode in kube-proxy?

**Answer:**
IP Virtual Server.
- Provides transport-layer load balancing.
- Supports different algorithms (Round Robin, Least Connection).
- Faster rule syncing than iptables for large clusters (10k+ services).

---

### Question 376: What is an overlay network?

**Answer:**
A virtual network built on top of physical network.
- **VXLAN/IPIP:** Encapsulates Ethernet frames inside UDP packets.
- Allows Pods to see a flat L3 network (10.x.x.x) even if nodes are on different subnets (192.x.x.x).

---

### Question 377: What is MetalLB?

**Answer:**
A LoadBalancer implementation for bare metal (On-Prem).
- On Cloud, AWS gives you an LB.
- On-Prem, you have none. MetalLB simulates it by assigning a "External IP" from a pool to the Service and announcing it via ARP/BGP.

---

### Question 378: What is a NodePort service and its limitations?

**Answer:**
- **Limit 1:** Non-standard ports (30000+).
- **Limit 2:** Performance (Hairpinning).
- **Limit 3:** Security (Opening ports on every node).
- **Limit 4:** One service per port.

---

### Question 379: What is the use of service type ExternalName?

**Answer:**
DNS Alias (CNAME).
- `ExternalName: google.com`.
- Lookup internal service -> returns CNAME `google.com`.
- Routing happens at DNS level, no proxying.

---

### Question 380: How do you expose a Kubernetes service to the internet?

**Answer:**
1.  **Ingress:** (HTTP/HTTPS) - Recommended.
2.  **LoadBalancer:** (TCP/UDP) - One public IP per service.
3.  **NodePort:** (Development).

---

## 🔹 Cloud & Infrastructure (Questions 381-390)

### Question 381: What is the difference between GKE, EKS, and AKS?

**Answer:**
Managed K8s offerings.
- **GKE (Google):** Most mature, fastest spin-up, auto-upgrades.
- **EKS (AWS):** Deep AWS integration, manual upgrades mostly, highly configurable.
- **AKS (Azure):** Good integration with Active Directory, aggressive updates.

---

### Question 382: What is the cloud controller manager?

**Answer:**
The specialized controller that talks to Cloud Provider API.
- **Route Controller:** Configures routes in VPC.
- **Service Controller:** Creates Cloud Load Balancers.
- **Node Controller:** Checks if deleted node VM is gone from cloud.

---

### Question 383: What’s the impact of a cloud provider plugin?

**Answer:**
It allows K8s to control physical resources.
- `type: LoadBalancer` works because of this plugin.
- `StorageClass: gp2` works because of this plugin.

---

### Question 384: What is CSI vs cloud storage drivers?

**Answer:**
Legacy drivers were hardcoded.
**CSI:** Decoupled. AWS EBS driver is a Pod running in cluster. It provisions volumes via CSI API.

---

### Question 385: How does load balancing work in cloud-managed clusters?

**Answer:**
When you create a Service `LoadBalancer`:
1.  CCM (Cloud Controller Manager) calls AWS API.
2.  AWS provisions ELB.
3.  ELB targets are set to Worker Node's NodePorts.
4.  Traffic: ELB -> NodePort -> IPTables -> Pod.

---

### Question 386: What is the recommended way to use IAM roles with Kubernetes pods?

**Answer:**
**Workload Identity (GKE) / IRSA (EKS).**
- Instead of putting AWS Keys in Secret.
- You annotate the ServiceAccount (`eks.amazonaws.com/role-arn: ...`).
- K8s injects a JWT token.
- AWS STS verifies token and allows access.

---

### Question 387: What is Fargate in EKS?

**Answer:**
Serverless Compute for EKS.
- You don't manage EC2 nodes.
- You schedule a Pod.
- AWS provisions a Micro-VM just for that Pod.
- You pay per vCPU/RAM.

---

### Question 388: What is workload identity in GKE?

**Answer:**
Binds a Kubernetes ServiceAccount to a Google Cloud IAM Service Account.
- Allows Pods to access Google Cloud Storage/Spanner/PubSub securely.

---

### Question 389: How do you enable autoscaling in managed clusters?

**Answer:**
Enable the **Cluster Autoscaler** feature in the console (or terraform).
- It manages the Auto Scaling Group (ASG) size.

---

### Question 390: What is node pool management?

**Answer:**
Grouping nodes with similar properties.
- Pool A: "General Purpose" (t3.medium).
- Pool B: "GPU" (p3.2xlarge).
- You schedule pods to pools using NodeSelectors/Affinity.

---

## 🔹 Best Practices & Real Scenarios (Questions 391-400)

### Question 391: What are best practices for multi-tenant Kubernetes clusters?

**Answer:**
1.  **Namespaces** for isolation.
2.  **NetworkPolicies** deny-all by default.
3.  **RBAC** strictly scoped.
4.  **ResourceQuotas** to prevent noisy neighbors.
5.  **PodSecurityAdmission** to prevent root access.

---

### Question 392: How do you secure a cluster running in production?

**Answer:**
- **Private API Endpoint:** No public internet access to Master.
- **Bastion Host:** Access kubectl via Bastion.
- **Image Scanning:** Fail CI if critical CVEs found.
- **Audit Logs:** Enable and monitor.
- **Least Privilege:** Minimal RBAC.

---

### Question 393: How do you isolate staging and production workloads?

**Answer:**
**Best:** Separate Clusters.
- Eliminates risk of "accidental deletion", "upgrades breaking API", "shared kernel exploits".
**Acceptable:** Separate Namespaces with strict Node Affinity (Dedicate specific nodes to Prod).

---

### Question 394: What are some cost-saving techniques in Kubernetes?

**Answer:**
1.  **Spot Instances:** Use for stateless/batch workloads (up to 90% cheaper).
2.  **Rightsizing:** Use VPA/Goldilocks to set accurate requests.
3.  **autoscaling:** Scale down to 0 at night (if dev env).
4.  **Bin Packing:** Use fewer large nodes.

---

### Question 395: How do you version control Kubernetes manifests?

**Answer:**
**GitOps.**
- Repo structure: `/base`, `/overlays/dev`, `/overlays/prod`.
- Changes are PRs.
- History is Git Commit log.

---

### Question 396: How do you ensure reliability in Kubernetes upgrades?

**Answer:**
1.  **Blue/Green Cluster:** Build new cluster, switch traffic. (Expensive/Safe).
2.  **Rolling Upgrade:** Upgrade one node at a time (Standard).
   - Ensure `PodDisruptionBudgets` are set so app stays up.
3.  **Backup Etcd** before starting.

---

### Question 397: What are the risks of running outdated Kubernetes versions?

**Answer:**
- **Security:** Missing patches for container breakouts/CVEs.
- **Compatibility:** APIs get deprecated (e.g., PSP removed). Helm charts stop working.
- **Support:** Cloud providers force upgrade or charge penalties.

---

### Question 398: How do you ensure compliance in a Kubernetes environment?

**Answer:**
Using Policy Engines (OPA Gatekeeper / Kyverno / Datree).
- "All images must come from trusted registry."
- "All pods must have labels."
- "No LoadBalancers allowed in Dev."

---

### Question 399: What are anti-patterns in Kubernetes monitoring?

**Answer:**
1.  **High Cardinality:** Logging unique IDs (IPs, SessionIDs) as Metric Labels. explodes TSDB.
2.  **Alert Fatigue:** Alerting on "Pod Restart" (Normal K8s behavior) instead of "Service Down".

---

### Question 400: What is GitOps and how does it relate to Kubernetes?

**Answer:**
GitOps applies DevOps best practices to Infrastructure Management.
- Git is the "Single Source of Truth".
- Kubernetes applies that truth.
- `kubectl` is for read-only debugging, not for changing state.
