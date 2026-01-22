# Kubernetes Interview Questions & Answers

## 🔹 1. Kubernetes Basics (Questions 1-10)

**Q1: What is Kubernetes?**
Kubernetes (K8s) is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

**Q2: Why do we use Kubernetes?**
To manage large numbers of containers efficiently, ensuring high availability, scalability, self-healing, and automated deployment/rollout.

**Q3: What is a Pod in Kubernetes?**
The smallest deployable unit in K8s. It represents a single instance of a running process and can hold one or more containers sharing network/storage.

**Q4: What is a Node?**
A worker machine (VM or physical) in a cluster that runs Pods. It contains the Kubelet, Kube-proxy, and container runtime.

**Q5: What is a Cluster?**
A set of grouped nodes (master and workers) that run containerized applications managed by Kubernetes.

**Q6: What is a Namespace?**
A virtual cluster inside a physical cluster. It provides isolation for resources, users, and policies (e.g., `dev`, `prod`).

**Q7: What is the difference between Docker and Kubernetes?**
Docker is a container runtime (creates containers). Kubernetes is a container orchestrator (manages/scales containers across nodes).

**Q8: What is Minikube?**
A lightweight tool that runs a single-node Kubernetes cluster locally inside a VM or container for development and testing.

**Q9: What is kubelet?**
An agent running on every node that ensures containers are running in a Pod as per the instructions from the control plane.

**Q10: What is kubectl?**
The command-line tool used to communicate with the Kubernetes API Server to manage correctly the cluster resources.

---

## 🔹 2. Architecture (Questions 11-20)

**Q11: Explain Kubernetes architecture.**
It consists of a Control Plane (Master) and Worker Nodes. Master manages the cluster (API Server, Scheduler, Controller Manager, etcd). Workers run apps (Kubelet, Kube-proxy, Container Runtime).

**Q12: What are the components of the Master Node?**
API Server, etcd, Scheduler, Controller Manager, and Cloud Controller Manager.

**Q13: What is the role of etcd?**
A highly available key-value store that keeps the cluster's persistent state (config, secrets, metadata). Source of truth for K8s.

**Q14: What is the Scheduler?**
Watches for newly created Pods with no assigned node and selects the best node for them based on resources/constraints.

**Q15: What is the Controller Manager?**
Runs controller processes (Node Controller, Replication Controller) that regulate the state of the cluster towards the desired state.

**Q16: What is the API Server?**
The front-end of the K8s control plane. It exposes the REST API (port 6443) and validates/processes all requests.

**Q17: What is a worker node?**
A machine that runs application workloads. It hosts Pods and reports status to the Master.

**Q18: What are DaemonSets?**
A controller that ensures a copy of a specific Pod runs on *all* (or selected) nodes (e.g., logging agents, monitoring).

**Q19: What is kube-proxy?**
A network proxy running on each node that maintains network rules and allows communication to/from Pods (Service load balancing).

**Q20: How does service discovery work in Kubernetes?**
Using DNS (CoreDNS) and Environment Variables. Services get a stable IP/DNS name. Pods look up services by name.

---

## 🔹 3. Pods & ReplicaSets (Questions 21-25)

**Q21: What is a ReplicaSet?**
Ensures a specified number of Pod replicas are running at any given time. Replaces the older ReplicationController.

**Q22: How is a ReplicaSet different from a ReplicationController?**
ReplicaSets support "selector broadcasting" (set-based selectors: `In`, `NotIn`) while ReplicationControllers only support equity-based selectors.

**Q23: How do you scale Pods in Kubernetes?**
`kubectl scale deployment <name> --replicas=5` or declaratively by updating the `replicas` field in the manifest.

**Q24: What is a multi-container Pod?**
A Pod with more than one container. They share the same IP, Lifecycle, and Storage (e.g., App container + Sidecar logger).

**Q25: How do Pods communicate with each other?**
If in the same Pod: `localhost`. If different Pods: via their Pod IPs (flat network) or via Services.

---

## 🔹 4. Deployments & Rollouts (Questions 26-30)

**Q26: What is a Deployment?**
A higher-level object that manages ReplicaSets and Pods. It enables declarative updates, rollbacks, and scaling.

**Q27: How do you roll back a Deployment?**
`kubectl rollout undo deployment <name>`. Reverts to the previous revision.

**Q28: What is a Rolling Update?**
The default deployment strategy. Updates Pods incrementally (killing old, starting new) with zero downtime.

**Q29: How do you pause and resume Deployments?**
`kubectl rollout pause deployment <name>` and `kubectl rollout resume deployment <name>`.

**Q30: What is a canary deployment in Kubernetes?**
Rolling out a new version to a small subset of users (traffic) first, verifying it, then rolling out to everyone.

---

## 🔹 5. Services & Networking (Questions 31-40)

**Q31: What is a Service in Kubernetes?**
An abstraction that defines a logical set of Pods and a policy to access them (provides a stable IP/DNS).

**Q32: What are the types of Kubernetes Services?**
ClusterIP (default), NodePort, LoadBalancer, and ExternalName.

**Q33: What is a ClusterIP service?**
Exposes the service on a cluster-internal IP. Reachable only within the cluster.

**Q34: What is a NodePort service?**
Exposes the service on each Node's IP at a static port (30000-32767). External traffic access: `<NodeIP>:<NodePort>`.

**Q35: What is a LoadBalancer service?**
Exposes the service externally using a cloud provider's Load Balancer (AWS ELB, GCP LB).

**Q36: What is an Ingress?**
An API object that manages external access to services (HTTP/HTTPS), typically providing load balancing, SSL termination, and name-based virtual hosting.

**Q37: Difference between Ingress and LoadBalancer?**
LoadBalancer = One IP per service (expensive). Ingress = One IP for *many* services (smart routing layer).

**Q38: How does DNS work in Kubernetes?**
CoreDNS watches the API server for new Services and creates DNS records (`my-svc.my-ns.svc.cluster.local`).

**Q39: What is a NetworkPolicy?**
A firewall rule for Pods. Controls which Pods can communicate with each other (default is allow-all).

**Q40: How do you expose a service to the outside world?**
Use NodePort, LoadBalancer, Ingress, or `kubectl port-forward` (for debugging).

---

## 🔹 6. ConfigMaps & Secrets (Questions 41-45)

**Q41: What is a ConfigMap?**
An API object to store non-confidential data in key-value pairs. Decouples config from container image.

**Q42: What is a Secret?**
Similar to ConfigMap but for confidential data (passwords, tokens). Stored base64-encoded (not encrypted by default until configured).

**Q43: How do you mount a ConfigMap to a Pod?**
As environment variables (`envFrom`) or as a volume (`volumes` mount) to a file path.

**Q44: How are Secrets stored in Kubernetes?**
Stored in etcd. By default, they are base64 encoded strings. In production, enable Encryption at Rest.

**Q45: How do you update a ConfigMap without restarting a Pod?**
If mounted as a Volume, the file updates automatically (eventually). If used as Env Var, Pod restart is required.

---

## 🔹 7. Volumes & Storage (Questions 46-50)

**Q46: What is a Volume in Kubernetes?**
A directory containing data, accessible to containers in a specific Pod. Solves the ephemeral nature of container filesystems.

**Q47: Difference between emptyDir and hostPath?**
`emptyDir`: Temporary storage tied to Pod lifespan. `hostPath`: Mounts a file/dir from the Node's filesystem (persists on Node).

**Q48: What is PersistentVolume (PV)?**
Storage resource in the cluster (like an EBS volume or NFS share) provisioned by admin or dynamically. Independent of Pod lifecycle.

**Q49: What is PersistentVolumeClaim (PVC)?**
A request for storage by a user. It consumes PV resources (like a Pod consumes Node resources).

**Q50: What is StorageClass?**
Defines the "class" of storage (e.g., standard, ssd-fast) and allows dynamic provisioning of PVs.

---

## 🔹 8. Security (Questions 51-55)

**Q51: What is RBAC?**
Role-Based Access Control. Regulates access to K8s resources based on the roles of individual users within the organization.

**Q52: What are ServiceAccounts?**
Identities for processes running in Pods (non-human users) to interact with the API Server.

**Q53: How does Kubernetes handle authentication?**
It doesn't have a user database. Uses modules: Client Certs, Bearer Tokens, OIDC, Webhooks, or Impersonation.

**Q54: What are NetworkPolicies used for?**
To secure traffic flow between Pods. "Deny all" by default, then whitelist specific paths.

**Q55: How do you restrict access to the API server?**
Use RBAC, enable Audit Logging, disable anonymous access, and firewall the API server endpoint.

---

## 🔹 9. Helm & Operators (Questions 56-60)

**Q56: What is Helm?**
The package manager for Kubernetes. Helps define, install, and upgrade complex K8s apps.

**Q57: What are Helm Charts?**
Packages of pre-configured K8s resources. Equivalent to `apt` or `rpm` packages but for clusters.

**Q58: How do you install applications using Helm?**
`helm install <release-name> <chart-name>`.

**Q59: What is a Helm repository?**
An HTTP server that houses an `index.yaml` and chart packages (tarballs).

**Q60: What is a Kubernetes Operator?**
Software extensions to K8s that make use of custom resources to manage applications and their components (automating human ops tasks).

---

## 🔹 10. Monitoring & Logging (Questions 61-65)

**Q61: How do you monitor a Kubernetes cluster?**
Using Metrics Server (for basic stats) and full stack tools like Prometheus + Grafana.

**Q62: What is Prometheus?**
An open-source monitoring system that scrapes metrics from K8s nodes/pods via HTTP endpoints.

**Q63: What is Grafana?**
A visualization tool used to create dashboards from Prometheus data.

**Q64: How do you collect logs in Kubernetes?**
`kubectl logs`. For centralized logging, run a log agent (Fluentd/Filebeat) on nodes to ship logs to ELK/Splunk.

**Q65: What is Fluentd?**
A unified logging layer. Collects and transforms logs from containers and sends them to a backend (Elasticsearch).

---

## 🔹 11. Autoscaling (Questions 66-70)

**Q66: What is Horizontal Pod Autoscaler (HPA)?**
Automatically scales the number of Pods up/down based on CPU/Memory usage.

**Q67: What is Vertical Pod Autoscaler (VPA)?**
Automatically adjusts the CPU/Memory requests/limits for containers in a Pod (restarts Pods to apply).

**Q68: What is Cluster Autoscaler?**
Scales the number of Nodes in the cluster. Adds a node if Pods are pending; removes a node if underutilized.

**Q69: How does HPA work internally?**
It queries the Metrics Server periodically (default 15s) and calculates generic usage vs target.

**Q70: Can you autoscale based on custom metrics?**
Yes, using the Prometheus Adapter to expose custom metrics (like queue length) to the custom metrics API.

---

## 🔹 12. Troubleshooting (Questions 71-75)

**Q71: How do you troubleshoot a CrashLoopBackOff error?**
`kubectl logs <pod>`, `kubectl describe pod <pod>` (check events), check Liveness Probe config.

**Q72: What command do you use to view Pod logs?**
`kubectl logs <pod-name> [-c <container-name>]`.

**Q73: How do you debug a Kubernetes node issue?**
`kubectl describe node <node>`. SSH into node. Check `journalctl -u kubelet`. Check disk/memory pressure.

**Q74: What are common reasons for Pending Pods?**
Insufficient resources (CPU/Mem), Taints/Tolerations mismatch, PVC not bound, or Scheduler failure.

**Q75: How to check resource usage per Pod?**
`kubectl top pod`. Requires Metrics Server to be installed.

---

## 🔹 13. CI/CD with Kubernetes (Questions 76-80)

**Q76: How do you implement CI/CD in Kubernetes?**
CI builds docker image -> Push to Registry -> CD tool (Argo/Flux/Jenkins) updates the deployment manifest to new tag.

**Q77: What tools integrate well with Kubernetes for CI/CD?**
Jenkins, GitLab CI, CircleCI, ArgoCD, FluxCD, Tekton.

**Q78: What is ArgoCD?**
A declarative, GitOps continuous delivery tool for Kubernetes. Syncs cluster state with Git repo.

**Q79: What is FluxCD?**
Another popular GitOps tool. Similar to ArgoCD but headless (CLI focused) and tightly integrated with Helm.

**Q80: How does GitOps work in Kubernetes?**
Git is the single source of truth. An agent in the cluster pulls changes from Git and applies them.

---

## 🔹 14. Advanced Concepts (Questions 81-90)

**Q81: What is a StatefulSet?**
Manages stateful apps (DBs). Guarantees stable network IDs (`web-0`, `web-1`), stable storage binding, and ordered deployment.

**Q82: What is a Job and CronJob?**
Job: Runs a Pod to completion (batch task). CronJob: deeply creates Jobs on a schedule (like crontab).

**Q83: What is taint and toleration?**
Taint: Applied to Node ("Don't schedule here unless..."). Toleration: Applied to Pod ("I can tolerate this taint").

**Q84: What is affinity and anti-affinity?**
Rules to prefer/require Pods to run on specific nodes (Node Affinity) or near/far from other Pods (Pod Affinity/Anti-Affinity).

**Q85: What is a sidecar container?**
A helper container running alongside the main app container in the same Pod (e.g., logging agent, proxy).

**Q86: What is a headless service?**
A Service with `ClusterIP: None`. It doesn't load balance. DNS returns IPs of ALL backing Pods. Used for StatefulSets/Discovery.

**Q87: Explain init containers.**
Containers that run *before* the main app container starts. Used for setup (creating files, waiting for DB).

**Q88: What is a PodDisruptionBudget?**
Limits the number of Pods that can be down simultaneously during voluntary disruptions (e.g., node drain).

**Q89: What is custom resource definition (CRD)?**
Extends the K8s API allows you to define your own object types (like `PrometheusRule` or `MyDatabase`).

**Q90: How does Kubernetes perform self-healing?**
Restarting failed containers, rescheduling Pods when nodes die, and killing containers that don't respond to health checks.

---

## 🔹 15. Real-world & DevOps (Questions 91-100)

**Q91: How do you do blue-green deployments?**
Create a new Deployment (Green). Point Service selector to Green. If good, delete Blue.

**Q92: How do you do canary deployments?**
Replica split (90% stable, 10% canary) or advanced traffic splitting using Service Mesh (Istio/Linkerd) or Ingress.

**Q93: What are some best practices in Kubernetes security?**
Least privilege (RBAC), NetworkPolicies, Read-Only FS, Scan images, Secrets encryption, No root containers.

**Q94: How do you manage secrets in production?**
External Secret Stores (Vault, AWS Secrets Manager) synced via CSI drivers or operators (ExternalSecrets).

**Q95: How do you backup and restore etcd?**
`etcdctl snapshot save`. Critical for cluster recovery.

**Q96: What are some production monitoring strategies?**
USE method (Utilization, Saturation, Errors) for Nodes. RED method (Rate, Errors, Duration) for Services.

**Q97: What’s the difference between Kubernetes and OpenShift?**
OpenShift is RedHat's Enterprise K8s. Adds Developer UI, Source-to-Image (S2I), stricter security by default, and integrated registry.

**Q98: What cloud providers support Kubernetes?**
AWS (EKS), Google (GKE), Azure (AKS), DigitalOcean, Linode, etc.

**Q99: What is kubeadm?**
A tool to bootstrap a best-practice Kubernetes cluster. `kubeadm init` (master) and `kubeadm join` (workers).

**Q100: What is the future of Kubernetes?**
More abstraction (Serverless K8s), Edge computing (K3s), better security defaults, and AI/ML workload orchestration.

---

## 🔹 1. Cluster Management & Node Operations (Questions 101-110)

**Q101: How do you upgrade a Kubernetes cluster?**
Using `kubeadm upgrade plan` and `kubeadm upgrade apply` on the control plane, then `kubeadm upgrade node` on workers. Managed services (EKS/GKE) have 1-click upgrades.

**Q102: How do you drain a node safely?**
`kubectl drain <node-name> --ignore-daemonsets`. Evicts all Pods so the node can be maintained or removed.

**Q103: What happens during a `kubectl cordon`?**
Marks the node as unschedulable. No new Pods will be scheduled on it, but existing Pods continue running.

**Q104: How to add a node to an existing cluster?**
Run `kubeadm token create --print-join-command` on the master, then run the output command on the new node.

**Q105: How to safely remove a node from a Kubernetes cluster?**
Drain it (`kubectl drain`), delete it (`kubectl delete node <name>`), then reset the node (`kubeadm reset`).

**Q106: What is the difference between cordon, drain, and delete node?**
Cordon: Stop scheduling new pods. Drain: Cordon + Evict existing pods. Delete: Remove node object from API server.

**Q107: How do you handle node failures in Kubernetes?**
K8s marks node as `NotReady`. After 5m (default), Pods are evicted and rescheduled to healthy nodes.

**Q108: How does Kubernetes know a node is unhealthy?**
Node Controller checks heartbeat signals from Kubelet. If no signal for `node-monitor-grace-period`, it marks `NotReady`.

**Q109: What are node conditions?**
Status fields like `Ready`, `MemoryPressure`, `DiskPressure`, `PIDPressure`, and `NetworkUnavailable`.

**Q110: What is the purpose of a kubeconfig file?**
Stores cluster connection info, user credentials, and context (cluster + user + namespace combo) for `kubectl`.

---

## 🔹 2. Performance & Optimization (Questions 111-120)

**Q111: How do you optimize resource requests and limits?**
Measure actual usage (metrics), set Requests slightly above average usage, and Limits near peak acceptable usage to prevent OOM.

**Q112: What happens if a container exceeds its memory limit?**
The kernel OOM-kills the process (container restarts with `OOMKilled` reason).

**Q113: What is the effect of setting CPU limits too low?**
CPU throttling. The app runs slower (latency increases) but doesn't crash (like memory OOM).

**Q114: How do you monitor pod resource utilization?**
`kubectl top pod` or Prometheus metrics (`container_memory_usage_bytes`, `container_cpu_usage_seconds_total`).

**Q115: What is the difference between CPU throttling and CPU limits?**
Limit is the configuration cap. Throttling is the kernel action enforcing that cap by pausing the process.

**Q116: What tools can you use for cluster performance tuning?**
Prometheus/Grafana, Goldilocks (VPA recommender), Kube-cost, and stress-test tools (k6, Locust).

**Q117: How does Kubernetes scheduler select a node?**
Filters (fits requirements?) -> Scores (which fits best?) -> Selects highest score. Factors: Resources, Affinity, Taints.

**Q118: What is bin packing in scheduling?**
Packing Pods tightly onto fewer nodes to maximize utilization and save costs (often enables Cluster Autoscaler to downscale empty nodes).

**Q119: How to avoid noisy neighbor issues in Kubernetes?**
Set Resource Requests/Limits for *all* pods. Use ResourceQuotas and ranges. Isolate critical workloads on dedicated nodes/pools.

**Q120: How do resource requests affect scheduling?**
The scheduler uses "Requests" (not Limits) to decide if a Node has enough free capacity to fit the Pod.

---

## 🔹 3. Networking - Advanced (Questions 121-130)

**Q121: What is CNI (Container Network Interface)?**
A standard interface (spec) used by K8s to call network plugins (Calico, Flannel) to configure Pod networking.

**Q122: How does Kubernetes networking differ from Docker networking?**
K8s assumes a flat IP space (every Pod can talk to every Pod/Node without NAT). Docker uses bridge networks with NAT default.

**Q123: What are common CNI plugins?**
Calico (Network Policy), Flannel (Simple Overlay), Weave, Cilium (eBPF), AWS VPC CNI.

**Q124: What is Calico and how does it work?**
A CNI plugin supporting BGP routing and advanced Network Policies. Often used for security and performance.

**Q125: What is Flannel in Kubernetes?**
A simple, easy-to-configure Layer 3 overlay network (VXLAN). Ideally for smaller/simpler clusters.

**Q126: What is the role of iptables in Kubernetes?**
Used by `kube-proxy` in iptables mode to forward traffic from Service IPs to backend Pod IPs using NAT rules.

**Q127: How does service mesh work in Kubernetes?**
Injects a sidecar proxy (Envoy) into every Pod. The proxies intercept all traffic to handle mTLS, routing, and observability.

**Q128: What is Istio and why use it?**
A popular Service Mesh. Use it for Traffic Management (Canary), Security (mTLS), and Observability (Tracing) without changing app code.

**Q129: Explain the envoy proxy in Kubernetes.**
A high-performance L7 proxy used as the data plane in Service Meshes (Istio) or Ingress Controllers (Contour/Gloo).

**Q130: How is network isolation enforced in Kubernetes?**
Using `NetworkPolicy` resources (implemented by CNI like Calico) to deny/allow traffic between Pods based on labels.

---

## 🔹 4. Storage - Advanced (Questions 131-140)

**Q131: What is dynamic volume provisioning?**
Allows storage volumes (PVs) to be created on-demand when a user requests them via a PVC and StorageClass.

**Q132: What is the reclaim policy of a PersistentVolume?**
Determines what happens to PV when PVC is deleted: `Retain` (keep data), `Delete` (remove volume), `Recycle` (clean scrub - deprecated).

**Q133: How do StatefulSets handle persistent storage?**
Using `volumeClaimTemplates`. Each replica gets its own unique PVC (e.g., `data-web-0`, `data-web-1`).

**Q134: How do you back up Kubernetes volumes?**
Volume Snapshots (CSI), or tools like Velero (fs-backup), or application-level dumps (pg_dump) to S3.

**Q135: How do you resize a PVC?**
Edit the PVC spec (`resources.requests.storage`) and apply. Requires `allowVolumeExpansion: true` in StorageClass.

**Q136: What happens if a pod uses a PVC that no longer exists?**
Pod stays checking for the volume and fails to start (`ContainerCreating` or `VolumeMismatch` error).

**Q137: What is volume binding mode?**
Controls *when* volume is provisioned. `WaitForFirstConsumer` delays provisioning until Pod is scheduled (ensures volume is in same zone as node).

**Q138: How do you share storage between multiple pods?**
Use `ReadWriteMany` (RWX) access mode. Requires storage backend support like NFS, CephFS, or EFS.

**Q139: What is ephemeral storage?**
Local storage on the Node used by Pods for logs/temp files/container images. Cleared when Pod is evicted.

**Q140: Difference between CSI and in-tree volume plugins?**
In-tree: Code inside K8s binary (hard to update). CSI: Standard interface, plugins run as Pods (decoupled from K8s release cycle).

---

## 🔹 5. Authentication & Authorization (Questions 141-150)

**Q141: What is the difference between RBAC and ABAC?**
RBAC: Access based on User Roles. ABAC: Access based on policies combining attributes (User, Resource, State). K8s mostly uses RBAC.

**Q142: What is a ClusterRole vs Role?**
Role: Scoped to a Namespace. ClusterRole: Scoped to the entire Cluster (can access non-namespaced resources like Nodes).

**Q143: How do you create a custom ClusterRole?**
Define a `ClusterRole` YAML listing `rules` (apiGroups, resources, verbs) and apply it.

**Q144: How to bind a Role to a user?**
Create a `RoleBinding` (or `ClusterRoleBinding`) referencing the `Role` and the `User`/`Group`.

**Q145: What is a ServiceAccount token?**
A JWT token mounted in Pods (`/var/run/secrets/...`) allowing them to authenticate to the API Server.

**Q146: How do you secure access to the Kubernetes Dashboard?**
Expose via `kubectl proxy` (localhost only) or Ingress with Authentication/Oauth2-Proxy. Never expose publicly without auth.

**Q147: What is OIDC authentication in Kubernetes?**
OpenID Connect. Allows K8s to trust an external Identity Provider (Google, Okta, Keycloak) for user login.

**Q148: How do you rotate API server certificates?**
Regenerate certs on Master node and restart API server. Kubeadm handles this with `kubeadm certs renew`.

**Q149: What is impersonation in Kubernetes?**
A feature allowing an admin user to act as another user/group (`kubectl get pods --as=jane`). Useful for testing RBAC permissions.

**Q150: How do audit logs work in Kubernetes?**
Records every request to the API Server (Who, What, When). Configured via `audit-policy-file` flag on API Server.

---

## 🔹 6. Ingress & Traffic Control (Questions 151-160)

**Q151: What is the difference between Ingress Controller and Ingress Resource?**
Resource: The YAML definition (rules). Controller: The software (e.g., NGINX Pod) that implements those rules.

**Q152: What is the NGINX Ingress Controller?**
A popular implementation using NGINX as a reverse proxy to route traffic based on Ingress resources.

**Q153: How do you implement TLS termination in Ingress?**
Add `tls` section to Ingress spec referencing a Secret containing `tls.crt` and `tls.key`.

**Q154: What is path-based routing in Ingress?**
Routing traffic to different services based on URL path (e.g., `/api` -> Service A, `/web` -> Service B).

**Q155: How to handle rate limiting in Ingress?**
Use annotations specific to the controller (e.g., `nginx.ingress.kubernetes.io/limit-rps: "5"`).

**Q156: What is sticky session in Kubernetes?**
Session Affinity. routing requests from same client (cookie/IP) to same Pod. Configured in Service (`sessionAffinity`) or Ingress annotations.

**Q157: How to configure custom headers in Ingress?**
Commonly possible via annotations (e.g., `configuration-snippet` in Nginx) or ConfigMap options.

**Q158: What is external DNS in Kubernetes?**
A tool that syncs K8s Service/Ingress changes to external DNS providers (Route53, Cloudflare).

**Q159: How do you manage HTTPS certificates in Kubernetes?**
Use `cert-manager` to automate issuance/renewal from Let's Encrypt or other CAs.

**Q160: What is cert-manager?**
A native K8s controller that watches Certificate CRDs and manages the lifecycle (request, renew) of SSL certs.

---

## 🔹 7. Logging & Monitoring - Advanced (Questions 161-170)

**Q161: What are the best practices for centralized logging in Kubernetes?**
Use a DaemonSet (Flux/Filebeat) to collect node/container logs and ship to a central store (ES, Loki, Splunk).

**Q162: What is the EFK stack?**
Elasticsearch (Storage), Fluentd (Collector), Kibana (Visualization). A common logging stack.

**Q163: What is Loki in Kubernetes logging?**
A log aggregation system by Grafana Labs. Like Prometheus but for logs (indexes labels, not content). efficient and lightweight.

**Q164: What is the difference between container logs and system logs?**
Container: Stdout/Stderr from apps. System: Logs from Kubelet, Docker, Kernel (`/var/log/syslog`).

**Q165: How to monitor etcd health?**
Check `/health` endpoint. Monitor metrics `etcd_server_has_leader`. Use `etcdctl endpoint health`.

**Q166: How do you implement alerting in Kubernetes?**
Prometheus AlertManager. Define rules (`PrometheusRule`), group alerts, and route to receivers (Slack, PagerDuty).

**Q167: What is kube-state-metrics?**
A service that talks to API Server and generates metrics about the state of objects (Deployment status, Pod restarts, etc.).

**Q168: What are some metrics to monitor for production clusters?**
Node CPU/Mem/Disk, Pod Restart Count, API Server Latency, Etcd Leader changes, PVC usage.

**Q169: How to handle high disk I/O alerts in Kubernetes?**
Identify noisy pod (iotop). Check for aggressive logging or heavy DB ops. Move workload to high-IOPS node/storage.

**Q170: What are application-level vs infrastructure-level metrics?**
Infra: CPU, Mem, Network (Node/Pod level). App: Request count, Error rate, Latency, Business metrics (exposed via `/metrics`).

---

## 🔹 8. CI/CD - Advanced (Questions 171-180)

**Q171: How do you implement GitOps in Kubernetes?**
Install ArgoCD/Flux. Store K8s manifests in Git. Configure tool to sync Git -> Cluster.

**Q172: What is the role of Argo Workflows?**
A workflow engine for K8s. Used for complex job orchestration, data pipelines, or CI steps native to K8s.

**Q173: What is a Kubernetes webhook?**
HTTP callback. api-server calls it to validate or mutate a request (Admission Webhook) or handle CRD conversion.

**Q174: How do you trigger a deployment pipeline on config change?**
CI/CD system webhook -> builds new image -> updates Git repo (GitOps) OR runs `kubectl set image`.

**Q175: What are pre-deployment and post-deployment hooks?**
Jobs/Scripts that run before/after a sync/upgrade (e.g., DB migration schema update). Helm and ArgoCD support these hooks.

**Q176: What is a Helm lifecycle hook?**
Special annotations in templates (e.g., `helm.sh/hook: pre-install`) to run jobs at specific points in release lifecycle.

**Q177: How do you use Kustomize in CI/CD?**
Keep base manifest in repo. CI applies overlays (prod/dev) using `kustomize build . | kubectl apply -f -`.

**Q178: How do you rollback failed deployments automatically?**
Helm (`--atomic`), ArgoCD (auto-sync with self-heal disabled or custom analysis), or `kubectl rollout undo` via script.

**Q179: What is Kubernetes-native CI/CD?**
CI/CD tools that run *on* K8s and use K8s resources (Tekton, Argo Workflows, Jenkins X).

**Q180: How do you manage secrets in GitOps?**
Don't commit raw secrets. Use Sealed Secrets (Bitnami), SOPS (Mozilla), or External Secrets Operator (references Vault/AWS).

---

## 🔹 9. Kubernetes Internals (Questions 181-190)

**Q181: What is a container runtime?**
Software that executes containers. Examples: containerd, CRI-O, Docker Engine (legacy), gVisor.

**Q182: What runtimes are supported in Kubernetes?**
Any runtime implementing CRI (Container Runtime Interface). Default is often containerd.

**Q183: What is CRI and why is it important?**
Interface allowing Kubelet to talk to different runtimes without recompiling. Decoupled Docker from K8s core.

**Q184: How do controllers work internally?**
Control Loop: Watch API Server for changes -> Compare Current State vs Desired State -> Act to reconcile.

**Q185: How does the Kubernetes Scheduler work?**
Informer watches Pods. Predicates (Filter) nodes. Priorities (Score) nodes. Bind Pod to best Node.

**Q186: What is a reconcile loop?**
The core logic of controllers. `func Reconcile()` runs repeatedly to ensure the world matches the spec.

**Q187: What is a finalizer in Kubernetes?**
A key in metadata that blocks deletion of an object until a controller performs cleanup and removes the key.

**Q188: What is a taint effect NoExecute?**
Evicts existing pods on the node if they don't tolerate the taint. (vs `NoSchedule` which only blocks new ones).

**Q189: How do events propagate inside Kubernetes?**
Component updates API Server -> API Server stores in etcd -> Watchers (other components) are notified via established connections.

**Q190: What are API groups in Kubernetes?**
Logical categorization of resources (e.g., `apps/v1`, `batch/v1`, `rbac.authorization.k8s.io/v1`). Helps versioning.

---

## 🔹 10. Real-world & Production Readiness (Questions 191-200)

**Q191: How do you prepare a cluster for production?**
HA Master, RBAC tight, NetPolicies on, Monitoring/Logging setup, Secrets encryption, Backup enabled.

**Q192: What are anti-patterns in Kubernetes deployment?**
Using `latest` tag, Baking config/secrets into images, privileged containers, no resource limits, gigantic clusters.

**Q193: How do you manage multiple environments (dev/stage/prod)?**
Namespaces (for small scale) or Separate Clusters (isolation). Manage config variance via Helm/Kustomize.

**Q194: How do you handle image versioning?**
SemVer tagging (v1.0.1). Use SHA digest for absolute immutability in critical prod deployments.

**Q195: How do you manage secrets securely in CI/CD pipelines?**
Inject at runtime via masked variables. Use OIDC to auth with cloud providers for fetching secrets.

**Q196: What is the best way to do blue/green deployment in Kubernetes?**
Use Service selector switching (point to deployment-blue then deployment-green). Or use Argo Rollouts for automated traffic shifting.

**Q197: How to implement canary deployments using Istio?**
VirtualService with weight rules (`weight: 90` destination v1, `weight: 10` destination v2).

**Q198: How do you secure inter-pod communication?**
Wait for NetworkPolicy (L3/L4) and Service Mesh mTLS (L7, identity-based encryption).

**Q199: What is multi-cluster Kubernetes?**
Running workloads across multiple distinct clusters for reliability, locality, or scale. Managed via federation or GitOps.

**Q200: What is federation in Kubernetes?**
Syncing resources across multiple clusters (KubeFed). Complex, giving way to GitOps-based multi-cluster management.

---

## 🔹 1. Core Concepts Deep Dive (Questions 201-210)

**Q201: What is a pod lifecycle?**
The phases a Pod passes through: Pending -> Running -> Succeeded/Failed. Also includes conditions (Initialized, Ready).

**Q202: What are the phases of a pod?**
`Pending` (scheduling/downloading), `Running` (at least one container up), `Succeeded` (exit 0), `Failed` (non-zero), `Unknown`.

**Q203: What is the difference between `status.phase` and `status.conditions`?**
Phase is high-level state (e.g., Running). Conditions is array of specific states (PodScheduled, Initialized, ContainersReady, Ready).

**Q204: How does `kubectl get events` help in debugging?**
Shows recent cluster activities (Scheduling failure, ImagePullBackOff, OOMKill, Probe failure) tied to resources.

**Q205: How does Kubernetes handle pod eviction?**
If node is under pressure, Kubelet kills pods based on QoS (BestEffort first, then Burstable). Scheduler moves them elsewhere.

**Q206: What triggers a pod to be terminated?**
Deleting the Pod, scaling down Deployment, Job completion, Failed health checks, or Node eviction/drain.

**Q207: What happens during pod termination?**
K8s sets `DeletionTimestamp`. Kubelet sends SIGTERM to containers. Waits `terminationGracePeriod`. Sends SIGKILL if still running.

**Q208: How does graceful shutdown work in Kubernetes?**
App traps SIGTERM, stops accepting new connections, finishes current requests, and exits. PreStop hooks can also run.

**Q209: What is `terminationGracePeriodSeconds`?**
Time to wait between SIGTERM and SIGKILL. Default 30s.

**Q210: What is a liveness probe vs readiness probe?**
Liveness: Restarts container if it fails (deadlock). Readiness: Removes Pod from Service endpoints if it fails (not ready for traffic).

---

## 🔹 2. Controllers & Workloads (Questions 211-220)

**Q211: What is a DeploymentSet? Is it a real object?**
No, "DeploymentSet" is not a standard object. You likely mean `DaemonSet` or `StatefulSet`.

**Q212: What happens when a Deployment is deleted?**
The associated ReplicaSets and Pods are typically garbage collected (deleted) unless `--cascade=orphan` is used.

**Q213: How does the ReplicaSet know to scale?**
It watches the Pods counting those matching its selector. If (Current < Desired), it creates. If (Current > Desired), it deletes.

**Q214: What is a rollout history in Kubernetes?**
Stored revisions of ReplicaSets belonging to a Deployment. Allows `kubectl rollout undo`.

**Q215: How can you resume a paused Deployment?**
`kubectl rollout resume deployment/<name>`. It continues processing the new ReplicaSet updates.

**Q216: What are the alternatives to StatefulSets?**
Deployments (for stateless), DaemonSets (per-node), or using Operators for complex stateful logic (databases).

**Q217: Why would you use a DaemonSet instead of a Deployment?**
If you need exactly one pod per node (e.g., logs, monitoring, CNI). Auto-scales with node count.

**Q218: What happens to DaemonSet pods when a node is drained?**
They are ignored by default. You must use `--ignore-daemonsets` to drain successfully. They are deleted but immediately recreated if node stays Ready.

**Q219: How can you run a one-time job on a schedule?**
Use a `CronJob` resource defined with a specific cron schedule `@once` (not standard, usually you set a specific date or use `Job`).

**Q220: Can you use Jobs for parallel processing?**
Yes, use `.spec.parallelism` (number of pods running at once) and `.spec.completions` (total successes needed).

---

## 🔹 3. Custom Resources & Extensibility (Questions 221-230)

**Q221: What is a CRD?**
Custom Resource Definition. Extends Kubernetes API with your own objects.

**Q222: How do you create a CRD?**
Write a YAML with `kind: CustomResourceDefinition`, specify group/version/kind (e.g., `MyApp`), and apply it.

**Q223: How do you manage CRDs with Helm?**
Put them in `crds/` folder in Chart. Helm installs them first (but generally doesn't upgrade/delete them safely).

**Q224: What is a Custom Controller?**
Code creating a loop that watches your Custom Resource and performs logic (the "Operator" pattern).

**Q225: What are the use cases of an Operator?**
Manage complex stateful apps (DB backups, recovery), automate infrastructure provisioning, or inject sidecars.

**Q226: What is Kubebuilder?**
A framework/SDK for building Kubernetes APIs and Controllers using Go.

**Q227: What is the difference between a controller and an operator?**
Operator = Controller + Domain Knowledge (usually for a specific app). Controller is the generic pattern.

**Q228: How do you version your CRDs?**
Use `versions` list in CRD spec (v1alpha1, v1beta1, v1). Use Conversion Webhooks to migrate data between versions.

**Q229: What is apiextensions.k8s.io?**
The API group that manages CRDs.

**Q230: How do you validate CRD schemas?**
Use OpenAPI v3 schema validation inside the CRD definition to enforce field types and required fields.

---

## 🔹 4. Scheduling & Affinities (Questions 231-240)

**Q231: What is inter-pod affinity?**
"Schedule this Pod ONLY on nodes that already have Pod X running" (Co-location).

**Q232: What is pod anti-affinity?**
"Do NOT schedule this Pod on nodes that have Pod X running" (Spread/HA).

**Q233: What are node affinity types?**
`requiredDuringScheduling...` (Hard rule), `preferredDuringScheduling...` (Soft rule).

**Q234: What does `preferredDuringSchedulingIgnoredDuringExecution` mean?**
Try to put the pod on a matching node/zone. If not possible, any node is fine. If node changes later, don't evict.

**Q235: How does taint-based eviction work?**
If a node gets a Taint (e.g., `NoExecute`), all Pods without a matching Toleration are evicted immediately.

**Q236: What are tolerations used for?**
Allowing a Pod to be scheduled on a specific Tainted node (e.g., GPU nodes, Master nodes).

**Q237: How can you ensure a pod always runs on a specific node?**
Use `nodeSelector` (simple) or `nodeAffinity` (flexible) matching unique node labels (hostname).

**Q238: What is a nodeSelector?**
Simple key-value map in Pod spec. `nodeSelector: { diskType: ssd }`.

**Q239: What is a pod topology spread constraint?**
Controls how Pods are spread across failure-domains (zones, regions, nodes) to ensure high availability and balance.

**Q240: How does the scheduler prioritize nodes?**
Feasible nodes are scored based on functions (LeastRequested, BalancedResourceAllocation, TaintToleration, ImageLocality).

---

## 🔹 5. Config & Secrets Advanced (Questions 241-250)

**Q241: What are subPaths in volume mounts?**
Allows mounting a single file (or subdir) from a Volume into the container, rather than mounting the whole root of the volume.

**Q242: How do you use environment variables from ConfigMaps?**
`envFrom: - configMapRef: name: my-config` (Imports all keys) OR `valueFrom: configMapKeyRef: ...` (Specific key).

**Q243: How do you update a Secret without restarting pods?**
If mounted as a Volume, it updates. Application must watch the file for changes (hot reload). If Env Var, must restart.

**Q244: What is the difference between stringData and data in Secrets?**
`data`: Base64 encoded values. `stringData`: Plain text (convenience for writing YAML). Converted to base64 by API.

**Q245: How can you encrypt Secrets at rest?**
Enable `EncryptionConfiguration` on API server pointing to a provider (local key, KMS, AES-CBC).

**Q246: What is KMS integration in Kubernetes?**
Key Management Service. API Server calls external KMS (AWS KMS, Google KMS) to generate Data Encryption Keys for Secrets.

**Q247: How do you manage secrets in Git securely?**
Never commit plain YAML. Use SealedSecrets (encrypts locally, cluster decrypts) or ExternalSecrets (fetch from Vault).

**Q248: What are sealed secrets?**
Controller + CLI. You encrypt secret with Public Key. K8s Controller decrypts with Private Key only it knows.

**Q249: What is HashiCorp Vault and how does it integrate with Kubernetes?**
A secret manager. Integrates via Agent Injector (sidecar) or CSI Provider to inject secrets into Pods.

**Q250: What is external-secrets?**
An Operator that reads secrets from AWS/GCP/Azure/Vault and creates native Kubernetes Secret objects.

---

## 🔹 6. Networking Deep Dive (Questions 251-260)

**Q251: What are endpoints in Kubernetes?**
List of IP:Port pairs for a Service (the actual backend Pods). Managed by EndpointController.

**Q252: How does the EndpointSlice resource differ from Endpoints?**
Scalable version of Endpoints. Splits large user-backends into multiple smaller slice objects to improve performance.

**Q253: What is kube-dns vs CoreDNS?**
kube-dns was legacy. CoreDNS is the current standard, plugin-based, incredibly flexible DNS server.

**Q254: What is the role of DNSPolicy?**
Determines how Pod resolves DNS. `ClusterFirst` (default), `Default` (node's /etc/resolv.conf), `None` (custom).

**Q255: What is headless service discovery?**
Service with no ClusterIP. DNS returns A-records for *all* Pod IPs directly. Client chooses which to connect to.

**Q256: What are readiness gates?**
Extra conditions in Pod Status that must be true for the Pod to be considered ready for Load Balancer traffic.

**Q257: What is ExternalName service?**
Maps a K8s Service to an external CNAME (e.g., `my-db.rds.amazonaws.com`). No proxying, just DNS alias.

**Q258: How do services resolve within a namespace?**
Just usage the name: `my-service`. (Resolves to `my-service.my-ns.svc.cluster.local`).

**Q259: How do you troubleshoot DNS issues in Kubernetes?**
Start a debug pod (`dnsutils`). Check `/etc/resolv.conf`. Run `nslookup kubernetes.default`. Check CoreDNS logs.

**Q260: What is `dnsPolicy: ClusterFirstWithHostNet`?**
For Pods running with `hostNetwork: true`. Ensures they still use the Cluster's DNS (CoreDNS) instead of the Node's DNS.

---

## 🔹 7. Security & Compliance (Questions 261-270)

**Q261: What is a PodSecurityPolicy (PSP)?**
Deprecated. Cluster-level resource controlling security sensitive aspects of pod specification (e.g. running as root).

**Q262: What replaced PSPs in newer Kubernetes versions?**
Pod Security Admission (PSA) / Pod Security Standards (PSS). Built-in labels on namespaces (`pod-security.kubernetes.io/enforce`).

**Q263: What is PodSecurityAdmission?**
Admission controller that enforces PSS levels (Privileged, Baseline, Restricted) based on namespace labels.

**Q264: How does Open Policy Agent (OPA) work?**
Generic policy engine. Used via Gatekeeper in K8s. Validates every API request against Rego policies (e.g. "Image must come from X").

**Q265: What is Kyverno?**
K8s-native policy engine. Policies are Kubernetes Resources (YAML), not Rego. Easier for K8s admins.

**Q266: What is container security context?**
`securityContext` field in Pod/Container spec. Controls User (UID), Capabilities, SELinux, ReadOnlyRoot.

**Q267: What is `runAsUser` and `fsGroup`?**
`runAsUser`: UID to run process. `fsGroup`: GID owning mounted volumes.

**Q268: How do you enforce non-root containers?**
Set `runAsNonRoot: true` in SecurityContext. Pod fails if image tries to run as UID 0.

**Q269: How to limit container capabilities?**
`capabilities: { drop: ["ALL"], add: ["NET_BIND_SERVICE"] }`. Least privilege principle.

**Q270: What are seccomp and AppArmor?**
Linux kernel security modules. Seccomp restricts syscalls (e.g., prevent `reboot`). AppArmor restricts file/network access.

---

## 🔹 8. Backup & Disaster Recovery (Questions 271-280)

**Q271: What is Velero?**
Standard tool for K8s backup/restore. Backs up Object Manifests to S3 and PVs (via Restic/Snapshots).

**Q272: How do you back up etcd?**
`etcdctl snapshot save backup.db`. Must utilize ca/cert/key flags for authentication.

**Q273: What happens if etcd is corrupted?**
Cluster API stops working. State is lost. Must restore from snapshot to recover cluster.

**Q274: How do you restore a Kubernetes cluster?**
Restore etcd snapshot on master nodes. Restart API server. Re-apply manifests if using GitOps.

**Q275: How do you perform point-in-time recovery for etcd?**
Only possible if you have periodic snapshots. Restore the specific snapshot file fitting the timeframe.

**Q276: What are etcd snapshots?**
Valid, consistent copies of the etcd database state at a moment in time.

**Q277: How to take a manual etcd backup?**
Exec into etcd pod or run `etcdctl` on the host. `etcdctl snapshot save /tmp/snapshot.db`.

**Q278: How do you monitor etcd disk usage?**
Metrics `etcd_mvcc_db_total_size_in_bytes`. Etcd has a quota limit (default 2GB or 8GB). Compact and defrag if full.

**Q279: What’s the impact of etcd latency?**
Slow API responses. Controller loops lag. Watch timeouts. Cluster instability.

**Q280: How do you back up Persistent Volumes?**
Use CSI VolumeSnapshots `kubectl create -f snapshot.yaml`. Or file-level backup using sidecars.

---

## 🔹 9. Cluster Federation & Multi-tenancy (Questions 281-290)

**Q281: What is Kubernetes federation?**
Managing multiple K8s clusters as a single entity. Syncing resources across them.

**Q282: How does kubefed work?**
Installs a controller that pushes "Federated" resources to member clusters. (Mostly inactive project now).

**Q283: What are the limitations of Kubernetes federation?**
Complexity, API version mismatch between clusters, differing network topologies, limited support.

**Q284: What is multi-tenancy in Kubernetes?**
Sharing a cluster between multiple teams/users. Hard Multi-tenancy (Zero trust, different NetworkPolicies). Soft (Namespaces).

**Q285: How do you isolate teams in Kubernetes?**
Namespaces, RBAC (RoleBindings), NetworkPolicies (deny cross-ns), ResourceQuotas, PSA (Restricted).

**Q286: What are Virtual Clusters?**
Running a full Control Plane (API/Controller/Scheduler) inside a Pod in a host cluster. Tenant thinks they are root.

**Q287: What is vCluster and its use case?**
A tool to create virtual clusters. Great for dev environments (cheap) or strict isolation.

**Q288: What is hierarchical namespace controller (HNC)?**
Allows Namespaces to have parents/children. Policy/RBAC propagates down the tree.

**Q289: How do you enforce network policies across namespaces?**
Default Deny-All policy. Whitelist specific namespaces using `namespaceSelector`.

**Q290: How do you implement resource quotas per namespace?**
Create `ResourceQuota` object in the namespace limits CPU/Mem/Pods.

---

## 🔹 10. Real-world Scenarios & Misc (Questions 291-300)

**Q291: How do you manage Kubernetes in air-gapped environments?**
Mirror all docker images to internal registry. Host local Helm charts. Download all binaries. No internet access allowed.

**Q292: What is OpenEBS and how does it help with local storage?**
A CAS (Container Attached Storage) solution. Turns local disk into dynamic PVs. Good for on-prem K8s.

**Q293: What is Portworx?**
Enterprise storage solution for K8s. HA, Replication, Encryption, Disaster Recovery across nodes/regions.

**Q294: How do you monitor API server performance?**
Prometheus metrics: `apiserver_request_duration_seconds`, `apiserver_request_total`. Watch 4xx/5xx rates.

**Q295: What is kube-bench?**
Tool that checks whether your K8s deployment meets the CIS Kubernetes Benchmark security guidelines.

**Q296: What is the Kubernetes conformance test suite?**
Sonobuoy. Runs standardized tests to verify a cluster works as per Official K8s specs. Required for "Certified K8s" logo.

**Q297: What is Cluster API (CAPI)?**
Declarative K8s-style API to create, configure, and manage Kubernetes clusters themselves (Clusters as Resources).

**Q298: What is a managed Kubernetes service?**
Cloud provider handles the Control Plane (Master). You just manage Nodes/Workloads. (EKS, AKS, GKE).

**Q299: What is the role of the CSI driver in Kubernetes?**
Connects K8s to storage vendors (AWS EBS, NetApp). Implements Attach/Mount logic.

**Q300: How would you secure Kubernetes clusters on a public cloud?**
Private Endpoints (no public API), minimal IAM roles for Nodes, Private Nodes (no public IP), Audit Logs enabled.

---

## 🔹 1. Pods & Containers Advanced (Questions 301-310)

**Q301: What’s the difference between a container restart and a pod restart?**
K8s doesn't "restart" Pods. If a Pod restarts (e.g. by ReplicaSet), it's a *new* Pod (new IP, new UID). A container restart happens *inside* the same Pod (same IP).

**Q302: How do you handle init containers in Kubernetes?**
Define them in `spec.initContainers`. Useful for ordered startup. They run sequentially and must strict succeed before the main app starts.

**Q303: What are the use cases of sidecar containers?**
Log shipping, Proxy/Mesh (Istio), Monitoring agent (Datadog), or Data seeding (git-sync).

**Q304: How do you detect a memory leak in a container?**
Monitor `container_memory_usage_bytes`. If it grows linearly without flats, investigate code. Check `OOMKilled` events.

**Q305: What are the pros/cons of running multiple containers in one pod?**
Pros: Local communication (localhost), shared volume. Cons: Tightly coupled scaling (must scale together), larger failure domain.

**Q306: How do you share files between containers in the same pod?**
Use an `emptyDir` volume mounted to both containers.

**Q307: Can you run privileged containers in Kubernetes?**
Yes, `securityContext.privileged: true`. Dangerous as it gives root access to the host. Avoid in production.

**Q308: How do you mount a hostPath volume?**
In `volumes:` section: `hostPath: { path: /var/log }`. Pod is tied to that specific node's filesystem.

**Q309: What is the lifecycle hook `preStop` used for?**
Graceful cleanup. E.g., executing a script to save state or drain connections before SIGTERM is sent.

**Q310: How do you test pod startup and shutdown behavior?**
Use `kubectl create` and `kubectl delete`. Watch events. Logs timestamps. Use stress test tools to kill pods randomly.

---

## 🔹 2. Deployments & Rollouts (Questions 311-320)

**Q311: What is a canary deployment strategy?**
Release new version to small % of traffic. Monitor metrics (Error rate). Increase traffic if healthy.

**Q312: How do you pause a deployment in Kubernetes?**
`kubectl rollout pause deployment/my-app`. Allows making multiple changes (image, env, resources) without triggering immediate rollouts.

**Q313: How do you manually trigger a rollout?**
`kubectl rollout restart deployment/my-app`. Forces a new ReplicaSet creation even if spec hasn't changed.

**Q314: What is a surge vs unavailable setting in rolling updates?**
`maxSurge`: How many extra pods over desire count. `maxUnavailable`: How many pods can be down during update.

**Q315: What happens when you update a deployment manifest?**
Controller detects diff. Creates new ReplicaSet. Scales up new RS and scales down old RS according to Strategy.

**Q316: How do you perform A/B testing in Kubernetes?**
Two separate Deployments (vA, vB). Ingress/Service Mesh splits traffic based on Headers (e.g., User-Agent or Cookie).

**Q317: How does rollout history differ from rollout status?**
`history`: List of previous revisions (ConfigMaps). `status`: Current detailed state (Progressing, Available, Failed).

**Q318: How can you automate rollback on failure?**
Native K8s Deployments don't auto-rollback on metrics. Need tools like Argo Rollouts or Flagger.

**Q319: How do you check the reason for deployment failure?**
`kubectl describe deployment`. Check conditions (`ProgressDeadlineExceeded`). Check Pod logs.

**Q320: What is progressive delivery?**
Advanced deployment (Canary/Blue-Green) controlled by automated metrics analysis (Prometheus queries) rather than manual approval.

---

## 🔹 3. Namespaces & Isolation (Questions 321-330)

**Q321: How do namespaces help with resource isolation?**
Grouping objects. Allowing ResourceQuotas (cpu/mem limits) per namespace. Scoping RBAC permissions.

**Q322: What is the default namespace?**
The namespace used when none is specified. `default`.

**Q323: How can you apply policies per namespace?**
NetworkPolicies, LimitRanges, ResourceQuotas, and RBAC RoleBindings are all namespaced resources.

**Q324: How can you restrict access to a namespace?**
Create a Role (not ClusterRole) providing access. Bind user to that Role in that Namespace only.

**Q325: How do you list all resources in a namespace?**
`kubectl api-resources --namespaced=true -o name | xargs -n 1 kubectl get --show-kind --ignore-not-found -n <ns>` (Tricky! `kubectl get all` misses many).

**Q326: What is the `kube-system` namespace used for?**
Core cluster components: API Server pods, Scheduler, Controller Manager, CNI, DNS, Proxy.

**Q327: What is a namespace quota?**
`ResourceQuota` object. Limits total CPU/Memory/Pod-count usable by all Pods combined in that namespace.

**Q328: How do you clean up all resources in a namespace?**
`kubectl delete namespace <name>`. Deletes the NS and everything inside it via garbage collection.

**Q329: Can you have the same pod name in different namespaces?**
Yes. Metadata names must be unique only *within* a namespace.

**Q330: How does RBAC differ when applied cluster-wide vs namespace-wide?**
ClusterRoleBinding grants permissions across *all* namespaces. RoleBinding allows permissions *only* in that namespace.

---

## 🔹 4. Cluster Configuration & Architecture (Questions 331-340)

**Q331: What is the role of the Kubernetes control plane?**
Make global decisions (scheduling), detecting/responding to events (starting new pods). Brain of the cluster.

**Q332: What is the kubelet responsible for?**
Node-agent. Registers node. Watches API server for Pods assigned to it. Tells runtime to start containers. Reports status.

**Q333: What is the purpose of the kube-proxy?**
Maintains network rules on nodes. Implements Service abstraction (VIPs) by forwarding traffic to backend pods.

**Q334: How does Kubernetes maintain desired state?**
Controllers run reconciliation loops. `Current != Desired` -> Take Action -> Repeat.

**Q335: What are admission controllers?**
Plugins in API Server that intercept requests *after* auth but *before* persistence. Can Validate or Mutate objects.

**Q336: What are validating vs mutating admission webhooks?**
Mutating: Modifies object (e.g., injects sidecar). Validating: Accepts/Rejects object (e.g., checks policies). Mutating runs first.

**Q337: What is the function of the API server?**
REST interface for K8s. Stateless. Validates data. Only component that talks to etcd.

**Q338: How does controller-manager differ from scheduler?**
scheduler: Decides *where* a pod goes. Controller-manager: Ensures workloads (Deployments, replicas) actually *exist*.

**Q339: What is a lease in Kubernetes?**
A mechanism for distributed locking. Used by Control Plane components for high-availability leader election.

**Q340: How does leader election work in control plane components?**
Components try to acquire a `Lease` object in `kube-system`. The one who holds the lock is active; others standby.

---

## 🔹 5. Autoscaling (Questions 341-350)

**Q341: How does the Horizontal Pod Autoscaler (HPA) work?**
Checks metrics (CPU/Mem). Calculates `desiredReplicas = currentReplicas * (currentMetric / desiredMetric)`. Updates Deployment scale.

**Q342: What is the Vertical Pod Autoscaler (VPA)?**
Analyzes historical usage. Suggests or Apply new CPU/Mem requests to Pods. Helps right-sizing.

**Q343: What metrics does HPA use?**
Standard: CPU, Memory (from Metrics Server). Custom: Throughput, Queue Depth (from Prometheus/Adapter).

**Q344: How does the Cluster Autoscaler work?**
Checks for Pending Pods (unschedulable). Adds Node. Checks for underutilized Nodes. Removes Node.

**Q345: What’s the minimum requirement for HPA to function?**
Metrics Server must be installed (provides `metrics.k8s.io` API). Requests must be set in Pod spec.

**Q346: How can you autoscale based on custom metrics?**
Deploy Prometheus Adapter. Configure it to translate Prometheus queries into Custom Metrics API format for HPA.

**Q347: What are the challenges with autoscaling StatefulSets?**
Data replication/sharding takes time. Scaling down requires careful data-handoff. HPA on StatefulSets is risky for databases.

**Q348: What is KEDA (Kubernetes Event-driven Autoscaling)?**
CNCF project. Autoscales based on events (Kafka topic offset, SQS queue length) even to zero replicas.

**Q349: What is the downside of aggressive autoscaling?**
Thrashing (rapid scale up/down). Increased cost due to minimum billing times. Cold start latency for users.

**Q350: How do you configure cooldown periods for autoscaling?**
HPA `behavior` field. `scaleDown.stabilizationWindowSeconds` (default 300s) prevents removing pods too quickly.

---

## 🔹 6. Helm – Package Management (Questions 351-360)

**Q351: What is Helm and how does it help in Kubernetes?**
Helps manage complexity by bundling manifests into Charts. Template variables allow reusing charts for different envs.

**Q352: What are Helm charts?**
The package format (collection of files: templates, values.yaml, Chart.yaml) describing the app.

**Q353: How do you create a custom Helm chart?**
`helm create <mychart>`. Generates standard directory structure.

**Q354: What is the difference between Helm 2 and Helm 3?**
Helm 2 used Tiller (server-side component) - security risk. Helm 3 is client-only, uses Secrets for state, simpler.

**Q355: What is a values.yaml file used for?**
Default configuration values for templates. Users override these during installation.

**Q356: What is Helm templating?**
Go text/template language. Uses `{{ .Values.key }}` syntax to inject values into YAML files dynamically.

**Q357: How do you pass values to a Helm chart at install time?**
`helm install --set image.tag=v1 -f production-values.yaml ...`.

**Q358: How do you upgrade an existing Helm release?**
`helm upgrade <release> <chart>`. Applies diff patches.

**Q359: What is a Helm release?**
An instance of a Chart running in the cluster. (Chart + Config = Release).

**Q360: How do you roll back a Helm deployment?**
`helm rollback <release> [revision]`. Reverts to previous release version.

---

## 🔹 7. Troubleshooting & Debugging (Questions 361-370)

**Q361: How do you debug CrashLoopBackOff errors?**
Check logs `kubectl logs`. Check exit code `kubectl describe pod`. Often config error, missing secret, or panicking app.

**Q362: What does `kubectl describe pod` show?**
Events (Scheduler decisions, Pull errors, Health check failures), Status, IP, Node assignment.

**Q363: How do you trace network issues in Kubernetes?**
Check NetworkPolicies. Ping/Curl from inside Pod. Check Service Endpoints. Use `netshoot` debug container.

**Q364: What logs do you check for API server issues?**
Control plane logs. On KCA: `/var/log/pods/...kube-apiserver...`. On Cloud: Cloud Logging options.

**Q365: How do you diagnose a hung pod?**
Exec into it `kubectl exec -it ...`. Check `top`, process list. Strace the process. Check for deadlocks.

**Q366: What does the `kubectl top` command show?**
Current CPU and Memory usage for Pods or Nodes (Requires Metrics Server).

**Q367: How do you find out which node a pod is scheduled on?**
`kubectl get pod -o wide` shows the NODE column.

**Q368: What are common reasons for pod pending state?**
Resource insufficient (CPU/Mem request too high), Taints, PVC pending, Node Selector mismatch.

**Q369: What is the difference between a warning and an error event?**
Warning: Something might be wrong (ImagePullBackOff). Normal: State change (Scheduled, Pulled). (Events have Type: Normal/Warning).

**Q370: What does `ImagePullBackOff` mean?**
Kubelet failed to pull image (Auth error, Typo, Network). It is backing off (waiting longer between retries).

---

## 🔹 8. Advanced Networking & DNS (Questions 371-380)

**Q371: How does service discovery happen in Kubernetes?**
DNS. Service `foo` in NS `bar` gets DNS record `foo.bar`. Pods search domains in `/etc/resolv.conf`.

**Q372: How do you create an internal-only service?**
Default ClusterIP service. Accessible only inside cluster.

**Q373: What is kube-proxy and how does it route traffic?**
Runs on every node. Watches Services/Endpoints. Updates iptables/IPVS rules to redirect Service IP traffic to Pod IPs.

**Q374: What is hairpin NAT in Kubernetes?**
Allows a Pod to access itself via its Service IP. (Traffic goes out of Pod to Bridge, loops back to Pod).

**Q375: What is ipVS mode in kube-proxy?**
Uses Linux IPVS database instead of iptables. Faster, scalable (O(1) lookup), supports better LBs (LeastConn).

**Q376: What is an overlay network?**
Virtual network (VXLAN/IPIP) built on top of physical node network. Allows Pod IPs to be routable across nodes.

**Q377: What is MetalLB?**
LoadBalancer implementation for bare metal clusters (where no AWS/GCP LB exists). Uses BGP or L2 (ARP).

**Q378: What is a NodePort service and its limitations?**
Opens port on *every* node. Security risk. Port range limited (30000+). Client must know Node IPs.

**Q379: What is the use of service type ExternalName?**
No ClusterIP/Selector. Returns a CNAME record. Redirects to external services (e.g., legacy DB).

**Q380: How do you expose a Kubernetes service to the internet?**
LoadBalancer Service (Cloud), Ingress (HTTP/LB), or NodePort (Not recommended for prod).

---

## 🔹 9. Cloud & Infrastructure (Questions 381-390)

**Q381: What is the difference between GKE, EKS, and AKS?**
Managed Service by Google, AWS, Azure. Differ in Control Plane management, CNI plugins default, and integrations.

**Q382: What is the cloud controller manager?**
Component linking K8s to Cloud API. Handles Node lifecycle (deleting K8s node if VM deleted) and LBs.

**Q383: What’s the impact of a cloud provider plugin?**
Legacy: built-in. Modern: External CCM. Allows K8s to create LBs, Volumes, Routes in the cloud account.

**Q384: What is CSI vs cloud storage drivers?**
CSI is the standard interface. Cloud drivers implement CSI to provision EBS/PD/AzureDisk.

**Q385: How does load balancing work in cloud-managed clusters?**
Service type `LoadBalancer` triggers CCM to call Cloud API -> Creates AWS ELB/ALB -> Points to Node Ports -> Kube-proxy -> Pod.

**Q386: What is the recommended way to use IAM roles with Kubernetes pods?**
Workload Identity (GKE), IRSA (EKS - IAM Roles for Service Accounts). Maps K8s SA to Cloud IAM Role via OIDC.

**Q387: What is Fargate in EKS?**
Serverless Compute. You don't manage EC2 nodes. Each Pod runs in its own isolated MicroVM environment.

**Q388: What is workload identity in GKE?**
Allows K8s ServiceAccount to act as a Google Service Account (GSA) to access GCP APIs (GCS, BigQuery).

**Q389: How do you enable autoscaling in managed clusters?**
Enable Cluster Autoscaler (configured in node pool settings). It adjusts the Auto Scaling Group (ASG) size.

**Q390: What is node pool management?**
Grouping nodes with same config (instance type, zone). Allows mixed operations (GPU pool, Spot pool).

---

## 🔹 10. Best Practices & Real Scenarios (Questions 391-400)

**Q391: What are best practices for multi-tenant Kubernetes clusters?**
Namespaces, NetworkPolicy isolation, ResourceQuotas, RBAC constraints, Pod Security Standards.

**Q392: How do you secure a cluster running in production?**
Private API Endpoint. Minimal Node OS. Image Scanning. GitOps. mTLS. Audit Logs.

**Q393: How do you isolate staging and production workloads?**
Separate Clusters is best (Strong isolation). Separate Namespaces if cost-constrained (Weak isolation).

**Q394: What are some cost-saving techniques in Kubernetes?**
Spot Instances for stateless. Right-sizing requests (VPA/Goldilocks). Downscaling dev envs at night (Kube-green).

**Q395: How do you version control Kubernetes manifests?**
GitOps repository. Structure by Env (dev/prod) or App. Use Helm/Kustomize to prevent duplicate YAML.

**Q396: How do you ensure reliability in Kubernetes upgrades?**
Blue/Green cluster upgrade or Rolling node pool upgrade. Test in Staging first. Back up etcd.

**Q397: What are the risks of running outdated Kubernetes versions?**
Security CVEs. Incompatibility with newer APIs/Tools. EOL (no patch support).

**Q398: How do you ensure compliance in a Kubernetes environment?**
Policy engines (OPA Gatekeeper/Kyverno). Enforce restrictions (Registry whitelist, No root, Required labels).

**Q399: What are anti-patterns in Kubernetes monitoring?**
Alerting on everything (Alert fatigue). Storing high-cardinality metrics (expensive). Not monitoring "Golden Signals".

**Q400: What is GitOps and how does it relate to Kubernetes?**
Ops by Pull Request. Git is the source of truth. Controller in cluster ensures usage matches Git. (ArgoCD/Flux).

---

## 🔹 1. GitOps & CI/CD in Kubernetes (Questions 401-410)

**Q401: What is GitOps in the context of Kubernetes?**
A methodology where Git is the single source of truth for K8s manifests. Agents (Argo/Flux) sync the cluster state to match the Git repo.

**Q402: How does ArgoCD work?**
A GitOps controller. It watches a Git repo (Helm/YAML) and compares it to the live cluster state. It syncs changes automatically or manually.

**Q403: What is FluxCD?**
Another GitOps tool. Flux v2 uses the GitOps Toolkit (Source Controller, Kustomize Controller, Helm Controller) to reconcile state.

**Q404: How does GitOps differ from traditional CI/CD?**
Traditional: CI pipeline pushes changes ("kubectl apply") to cluster. GitOps: CI pushes to Git; Cluster *pulls* changes (Pull interaction).

**Q405: What are the benefits of GitOps for K8s environments?**
Audit trail (Git commit history), Rollback (revert commit), Drift detection (alert if cluster changes manually), Disaster recovery.

**Q406: What are health checks in ArgoCD?**
Scripts (Lua) that assess the health of resources (e.g., "Is Deployment Progressing?"). Determines if a sync was successful.

**Q407: How do you manage secrets securely with GitOps?**
Store encrypted secrets in Git (Sealed Secrets, SOPS). The cluster-side controller has the private key to decrypt them.

**Q408: How do you implement multi-environment GitOps?**
Use directories (`/dev`, `/prod`) or branches (`main`, `prod`). Use Kustomize overlays or Helm Values to differentiate envs.

**Q409: What is the sync policy in ArgoCD?**
`Manual` (click button) or `Automated` (sync immediately). Options include `Prune` (delete orphan resources) and `SelfHeal` (fix drift).

**Q410: How do you track drift between Git and cluster state?**
ArgoCD UI shows "OutOfSync" status. `kubectl diff` can also show differences.

---

## 🔹 2. Chaos Engineering & Resilience (Questions 411-420)

**Q411: What is chaos engineering in Kubernetes?**
Proactively testing system resilience by injecting failures (killing pods, failing network) to ensure the system recovers gracefully.

**Q412: How does LitmusChaos work?**
A framework for defining Chaos Experiments as Custom Resources (`ChaosEngine`). Operators execute the chaos injection.

**Q413: What are the most common chaos experiments for K8s?**
Pod Delete (randomly kill pods), Node Drain, Network Latency injection, Packet Drop, CPU Stress.

**Q414: How do you simulate node failure?**
Stop the VM in the cloud console, or run `kubectl drain --ignore-daemonsets --delete-emptydir-data`.

**Q415: What is a pod-kill scenario?**
A test to verify that the ReplicaSet/Deployment controller restarts the pod and the application recovers connection.

**Q416: What metrics determine resiliency of a K8s app?**
MTTR (Mean Time To Recovery), Availability (Uptime during chaos), Error Rate during failure injection.

**Q417: How can you test for cascading failure in Kubernetes?**
Stress one service (inject latency) and observe if dependent services fail (Circuit Breaker testing).

**Q418: What is network partition chaos testing?**
Blocking traffic between two sets of pods (using NetworkPolicies or iptables) to test partition tolerance.

**Q419: What is the role of probes in resilience?**
Liveness probes restart hung apps. Readiness probes cut traffic to overloaded/failing apps to prevent 500 errors.

**Q420: How do you test for auto-healing behavior?**
Check if `kubectl get interactions` or controller events confirm the system detected the failure and remediated it without human intervention.

---

## 🔹 3. Observability & Monitoring (Questions 421-430)

**Q421: What is the role of Prometheus in Kubernetes?**
Metric collection via Pull model. Scrapes targets. Stores time-series data. Evaluates Alert rules.

**Q422: How does kube-state-metrics help in monitoring?**
It converts the state of K8s objects (Deployments, Nodes, PVCs) into metrics (e.g., `kube_deployment_status_replicas_available`).

**Q423: What is Grafana used for in K8s environments?**
Visualizing data from Prometheus (and others) via dashboards (Cluster Stats, Pod Usage, Network I/O).

**Q424: What is the difference between whitebox and blackbox monitoring?**
Whitebox: Metrics from *inside* the app (HTTP requests processed). Blackbox: Checking from *outside* (Health check ping, DNS resolution).

**Q425: How do you monitor node health?**
Node Exporter (OS metrics like CPU/Disk) + Kubelet metrics + Node Problem Detector (kernel issues).

**Q426: What is the role of Alertmanager?**
Receives alerts from Prometheus. Deduplicates, groups, routes them (Slack/Email/PagerDuty), and handles silence periods.

**Q427: How do you define custom metrics for HPA?**
Your app exposes metrics (e.g. `/metrics` queue_size). Prometheus scrapes it. Prometheus Adapter exposes it to Custom Metrics API.

**Q428: What is the difference between logs and metrics?**
Metrics: Aggregatable numbers (counters, gauges) for trends. Logs: Text events for debugging specific errors/context.

**Q429: What is OpenTelemetry?**
A vendor-neutral standard for ensuring apps generate traces, metrics, and logs in a unified format for any backend.

**Q430: How do you collect traces in Kubernetes?**
Instrument code with OTel SDK. Deploy Jaeger or Tempo. Send traces to the collector -> Backend.

---

## 🔹 4. Logging & Auditing (Questions 431-440)

**Q431: What is Fluentd and how does it integrate with Kubernetes?**
A log collector. Runs as DaemonSet. Reads `/var/log/containers/*.log`. Parses (JSON), filters, and buffers logs to ES/S3.

**Q432: How do you implement centralized logging?**
Use "Node Logging Agent" pattern (Fluentd/Bit). Ship all logs to a central cluster (Elastic/Splunk). Do not exec into pods to read logs.

**Q433: What is EFK/ELK stack?**
Elasticsearch (Search/Storage), Fluentd/Logstash (Collection), Kibana (UI). The standard open-source logging stack.

**Q434: How can you collect logs from pods?**
`kubectl logs`. Or sidecar container (if app logs to file, sidecar tails file to stdout). Or DaemonSet agents.

**Q435: What is the audit log in Kubernetes?**
A security record of every call to the API server (who did what). Essential for compliance (SOC2/PCI).

**Q436: How do you configure Kubernetes audit policies?**
Pass `--audit-policy-file=policy.yaml` to API server. Define levels (`None`, `Metadata`, `Request`, `RequestResponse`).

**Q437: How do you ensure logs are tamper-proof?**
Ship them immediately to a Write-Once-Read-Many (WORM) storage (e.g., S3 Object Lock) or separate restricted logging account.

**Q438: How do you identify suspicious activity in logs?**
Look for `403 Forbidden` spikes, accessing secrets, exec-ing into pods, or creating privileged pods.

**Q439: How do you rotate and manage logs in Kubernetes?**
Kubelet/Container Runtime handles log rotation on the node (`max-size`, `max-files`). Log Agent just tails the current file.

**Q440: What is Loki and how does it work with Grafana?**
Log aggregation system inspired by Prometheus. It labels log streams (using K8s labels) instead of indexing full text. Very cheap storage.

---

## 🔹 5. Storage Deep Dive (Questions 441-450)

**Q441: What is CSI (Container Storage Interface)?**
Specification for storage drivers. Allows 3rd party storage vendors to write plugins without modifying K8s core code.

**Q442: How does dynamic volume provisioning work?**
User creates PVC. StorageClass defines the provisioner (e.g., `ebs.csi.aws.com`). Controller creates the Volume (PV) automatically.

**Q443: What are volume reclaim policies?**
`Retain`: PV stays after PVC delete (Manual cleanup). `Delete`: PV & Data deleted. `Recycle`: `rm -rf` (deprecated).

**Q444: What is the difference between Retain, Delete, and Recycle?**
See above. Cloud volumes typically default to `Delete`. NFS defaults to `Retain`.

**Q445: What is ReadWriteOnce vs ReadWriteMany?**
RWO: Bound to a single Node (e.g., EBS). RWX: Shared by many Nodes (e.g., NFS/EFS). (Note: RWO can be used by multiple pods on *same* node).

**Q446: How do you mount a PVC across multiple pods?**
For RWO: Pods must be on same node. For RWX: Pods can be anywhere.

**Q447: How do StatefulSets work with persistent storage?**
Every replica gets a unique PVC from a `volumeClaimTemplate`. `pod-0` gets `pvc-0`. If `pod-0` restarts, it reattaches `pvc-0`.

**Q448: What are StorageClasses?**
Profiles for storage. Define `provisioner`, `parameters` (type=gp3, iops=3000), and `reclaimPolicy`.

**Q449: What’s the difference between ephemeral and persistent volumes?**
Ephemeral (EmptyDir): Lifetime = Pod Lifetime. Persistent (PV): Lifetime = Cluster/Storage System Lifetime.

**Q450: How do you backup data in PVCs?**
VolumeSnapshots (CSI). Or application-level backups (SQL dump, S3 sync).

---

## 🔹 6. Authentication & Authorization (Questions 451-460)

**Q451: How does Kubernetes handle authentication?**
It validates credentials (Cert, Token). If valid, extracts Username+Groups. Does NOT manage users itself.

**Q452: What are authentication plugins?**
Modules in API server. X509 Client Certs, Static Token File, Bootstrap Token, Service Account Token, OpenID Connect, Webhook.

**Q453: What is service account token projection?**
Mounting ServiceAccount tokens as Projected Volumes. Allows audience (`aud`) claim validation and token rotation.

**Q454: How do you authenticate with OIDC?**
User logs in to IDP (Google). Gets ID Token. Sends token to API Server (`Authorization: Bearer <token>`). API Server verifies signature.

**Q455: What is impersonation in Kubernetes?**
A user acts as another using headers (`Impersonate-User`). Requires `impersonate` verb on `users` resource in RBAC.

**Q456: What is ABAC and is it still supported?**
Attribute-Based Access Control. Using a static policy file. Hard to manage/update. Replaced by RBAC. Still supported but rarely used.

**Q457: How can you limit a user to read-only access?**
Create a Role with rule: `verbs: ["get", "list", "watch"]`. Bind user to this Role.

**Q458: How do you audit RBAC permissions?**
Tools like `kubectl auth can-i`, `rback`, or `rakkess` to visualize who can do what.

**Q459: What is RBAC escalation and how to prevent it?**
A user creating a Role with permissions *they don't have* (and binding it to themselves). K8s prevents this automatically (stops escalation).

**Q460: What are roles vs clusterroles?**
Role = Namespace scoped. ClusterRole = Cluster scoped (non-namespaced resources or default permissions for all namespaces).

---

## 🔹 7. Performance & Optimization (Questions 461-470)

**Q461: How do you profile CPU and memory usage of pods?**
`kubectl top`. Prometheus metrics. Or `kubectl profile` (plugin) or install `pprof` handlers in Go apps.

**Q462: What are best practices for container resource limits?**
Always set Requests (for scheduling). Set Limits (to prevent node starvation). Keep limits reasonable to avoid OOMKill.

**Q463: What happens if you don’t set resource limits?**
Pod is `BestEffort`. Can consume all CPU/Mem on node. First to be killed if node runs out of resources.

**Q464: How does the scheduler choose nodes based on resources?**
It filters out nodes where `NodeCapacity < (Sum of existing requests + New Pod request)`.

**Q465: What is the eviction threshold for memory pressure?**
Kubelet setting (e.g., `memory.available < 100Mi`). When crossed, Kubelet evicts pods to reclaim memory.

**Q466: How do you minimize pod startup time?**
Small images (distroless/alpine). Lazy loading. Set correct probes (don't wait too long). Use Image Pull Cache (Nodes already have image).

**Q467: What are tools to benchmark Kubernetes clusters?**
K6, Locust (App load). Kubemark (Control plane stress). Etcd-benchmark.

**Q468: How do you reduce image pull latency?**
Use smaller images. Use a local registry mirror. Pre-pull images on nodes (DaemonSet or ImagePuller).

**Q469: What is `imagePullPolicy: IfNotPresent` vs `Always`?**
`IfNotPresent`: Use local cache if available (Fast). `Always`: Check registry for new digest every time (Secure/Up-to-date, but requires auth/network).

**Q470: How does the pod preemption mechanism work?**
If high-priority pod is Pending, Scheduler evicts lower-priority pods to make room on a node.

---

## 🔹 8. Upgrades & Versioning (Questions 471-480)

**Q471: How do you upgrade a Kubernetes cluster safely?**
Upgrade Master components first. Then upgrade Nodes one by one (Drain -> Upgrade Kubelet -> Uncordon).

**Q472: What is the upgrade path for minor/patch versions?**
Skip minor versions? No (must go 1.25 -> 1.26 -> 1.27). Patch versions can jump (1.25.1 -> 1.25.5).

**Q473: What is the role of kubeadm in cluster upgrades?**
`kubeadm upgrade apply` upgrades the static pods manifests (API, Controller) and certificates on the node.

**Q474: How do you perform a zero-downtime upgrade?**
Your apps must have multiple replicas. PDB (PodDisruptionBudget) ensures enough replicas stay up while nodes drain.

**Q475: How do you upgrade CRDs?**
Apply the new YAML. Generally compatible but check for deprecated fields in the schema.

**Q476: How do you check for deprecated APIs before upgrading?**
`kubent` (Kube No Trouble) or `pluto`. Scans cluster for objects using API versions removed in next release (e.g. `v1beta1`).

**Q477: What’s the role of feature gates in Kubernetes upgrades?**
Flags to enable/disable alpha/beta features. New versions might graduate features to Beta/GA or remove them.

**Q478: How do you verify cluster health after an upgrade?**
Check `kubectl get nodes`. Check system pods `kubectl get pods -n kube-system`. Run conformance tests (Sonobuoy).

**Q479: What is the impact of upgrading etcd?**
Master/API downtime during update (unless HA). Critical to have backups.

**Q480: What’s the difference between kubelet and control plane upgrades?**
Control Plane: API uptime. Kubelet: Node functionality. Control plane version cannot be lower than Kubelet version.

---

## 🔹 9. Production Readiness (Questions 481-490)

**Q481: What are readiness criteria for going to production in Kubernetes?**
Resources set? Probes configured? HPA enabled? Logging/Monitoring active? Backup strategy tested? Security scanned?

**Q482: How do you enforce policies for resource limits?**
Namespace `LimitRange` (default request/limits) and `ResourceQuota`. PSA (Restricted).

**Q483: What is a pod disruption budget?**
Ensures application availability during maintenance. "Always keep at least 2 pods running".

**Q484: How do you manage blue/green deployment in production?**
Duplicate the stack. Switch the Ingress or Service LoadBalancer to point to the new stack.

**Q485: How do you ensure HA of critical workloads?**
Multiple replicas. Anti-affinity (spread across nodes). Topology constraint (spread across zones). PDB.

**Q486: What is the impact of control plane downtime?**
You can't change cluster state (deploy/scale). Existing apps Continue Running fine. Agents (Kubelet) keep running current state.

**Q487: What’s the difference between availability and resilience?**
Availability: % of time system is up. Resilience: Ability to recover from faults (resilience leads to high availability).

**Q488: How do you manage production secrets?**
External Secret Store (Vault/KMS). Never in Git. Rotate keys regularly.

**Q489: What are common production anti-patterns in Kubernetes?**
No limits/requests, manually applying changes (no GitOps), using `latest` tag, large shared clusters without isolation.

**Q490: How do you scale Kubernetes clusters in production?**
Cluster Autoscaler for infrastructure. GitOps for config. HPA for apps. Federation for multi-region.

---

## 🔹 10. Edge, Hybrid, & Specialized Use Cases (Questions 491-500)

**Q491: What is K3s and where is it used?**
Lightweight K8s distribution (binary <100MB). IoT, Edge, Dev environments. Removes legacy cloud providers/drivers.

**Q492: What is KubeEdge?**
CNCF project extending K8s to Edge. Supports offline autonomy (nodes work when connection to master is lost).

**Q493: What are microK8s?**
Canonical's zero-ops K8s. Single command install (`snap install microk8s`). Great for Ubuntu/local/IoT.

**Q494: How does Kubernetes support edge computing?**
By running lighter distros (K3s), managing low bandwidth (request coalescing), and tolerating disconnects.

**Q495: How can Kubernetes be used for IoT workloads?**
To deploy containerized processing logic (AI models) to edge gateways closer to sensors.

**Q496: How do you handle intermittent connectivity in edge K8s?**
Edge nodes cache config. Continued operation without API server connectivity. Sync when back online.

**Q497: What is the difference between bare-metal and cloud K8s?**
Bare-metal: You manage OS, Networking (BGP/MetalLB), Storage (Rook/Ceph). Cloud: Provider manages LB, Volume, Node Provisioning.

**Q498: How does multi-cloud Kubernetes deployment work?**
Abstraction layers (Crossplane/KubeFed) or independent clusters managed by a central GitOps plane.

**Q499: What is Anthos and how does it relate to Kubernetes?**
Google's platform to manage K8s clusters anywhere (GCP, AWS, Azure, On-Prem) with a unified control pane.

**Q500: What’s the role of WASM (WebAssembly) in Kubernetes?**
Running WASM modules instead of containers (via Krustlet/WasmEdge). Faster startup, smaller size, sandboxed.

---

## 🔹 1. Kubernetes Internals & Architecture (Questions 501-510)

**Q501: How does Kubernetes maintain consistency across distributed components?**
Uses etcd as the single source of truth (strongly consistent) and Watch mechanisms to notify controllers of changes.

**Q502: What is the role of etcd in Kubernetes?**
A distributed, reliable key-value store for critical cluster data (config, secrets, state). It implements RAFT consensus.

**Q503: How is data stored and secured in etcd?**
Stored as keys (e.g. `/registry/pods/default/mypod`). Secured via TLS (client certs) and Encryption at Rest.

**Q504: What is the informer pattern in Kubernetes controllers?**
A client-go pattern. Instead of polling API server, it caches objects locally and gets notified of events (Add/Update/Delete). Efficient.

**Q505: How do reconciliation loops work?**
`for { current = get_state(); target = get_spec(); if (current != target) make_changes(); sleep(interval); }`

**Q506: What is optimistic concurrency in K8s resource updates?**
Uses `resourceVersion`. If you update an object, but `resourceVersion` has changed on server (someone else updated), your update is rejected (409 Conflict). Retries needed.

**Q507: What is leader election in Kubernetes controllers?**
Ensures only one replica of the Controller Manager or Scheduler is active (Active/Passive HA). Uses Lease API.

**Q508: How does Kubernetes handle clock drift across nodes?**
It doesn't handle it well. Use NTP. Large drift causes TLS verification failures, token expiries, and anomalous behavior.

**Q509: How does Kubernetes achieve eventual consistency?**
Components retry indefinitely until state matches. It doesn't guarantee instant updates, but guarantees convergence over time.

**Q510: What is the function of a finalizer in resource deletion?**
Blocks the actual deletion of the object until the logic defined by the finalizer (clean up S3 bucket, LB) is executed and the finalizer key removed.

---

## 🔹 2. CRDs & Operator Pattern (Questions 511-520)

**Q511: What is a Custom Resource Definition (CRD)?**
Defines a new API type in the cluster. It doesn't do anything by itself until you write a Controller for it.

**Q512: How do you create a CRD?**
YAML manifest with `kind: CustomResourceDefinition`. Defines names (`myapps`), group (`example.com`), and schema (validation).

**Q513: What is an operator in Kubernetes?**
A software pattern: CRD + Custom Controller. Encodes human operational knowledge (how to upgrade MySQL) into code.

**Q514: What is the Operator SDK?**
A toolkit (Ansible, Helm, Go) to bootstrap and build Operators easily. Provides code scaffolding and testing tools.

**Q515: How does an operator differ from a controller?**
Controller is the mechanism. Operator is the application of that mechanism to a specific domain (like a "Kafka Operator").

**Q516: What are the common use cases for writing an operator?**
Stateful Apps (DBs), Auto-managing external SaaS (AWS RDS), Complex lifecycle management (Upgrades/Backups).

**Q517: What is the reconciliation logic in a custom controller?**
The `Reconcile` function. Receives `Namespace/Name`. Fetches object. Checks world. Fixes world. Updates Status.

**Q518: How do you watch for changes in custom resources?**
In Controller code (e.g., Kubebuilder), you set up a `Watch`. `ctrl.NewControllerManagedBy(mgr).For(&MyKind{})...`

**Q519: What are conversion webhooks in CRD versioning?**
Webhooks that convert object data between API versions (e.g. v1alpha1 -> v1beta1) on the fly during API requests.

**Q520: How do you handle backward compatibility in CRDs?**
Serve multiple versions (v1, v2). Use Conversion Webhooks. Deprecate fields slowly.

---

## 🔹 3. Multi-Cluster & Federation (Questions 521-530)

**Q521: What is Kubernetes Federation (KubeFed)?**
The official (but struggling) project to propagate K8s resources to member clusters.

**Q522: What are the use cases for multi-cluster deployment?**
Disaster Recovery (Region Failover), Latency (Edge), Isolation (Hard Multi-tenancy), Scale (exceeding 5k nodes).

**Q523: How do you share secrets across clusters?**
External Secret Store (Vault). GitOps (Sealed Secrets per cluster). Or replication tools like Reflector.

**Q524: What is a service mesh in a multi-cluster setup?**
Connects services across clusters transparently (e.g., Istio Multi-cluster). Pod in Cluster A calls `svc.ns.svc.cluster-b`.

**Q525: What are the challenges of multi-cluster observability?**
Aggregating metrics/logs from disparate sources. Use Federated Prometheus (Thanos/Cortex) and centralized Logging (Loki).

**Q526: How do you manage DNS across multiple clusters?**
`ExternalDNS` (syncs to Route53). Or `CoreDNS` with plugin for cross-cluster service discovery (`app.cluster-b.local`).

**Q527: How does ArgoCD support multi-cluster deployment?**
A single ArgoCD instance can register multiple external clusters (target contexts) and deploy apps to them from one UI.

**Q528: What is a cluster registry?**
A concept/API storing a list of available clusters and their connection endpoints.

**Q529: How do you enforce policies across clusters?**
Use Policy engines (OPA/Kyverno) deployed to all clusters via GitOps. Enforce standardized configuration.

**Q530: What is the difference between geo-replication and multi-cluster?**
Geo-replication usually strictly data (DBs). Multi-cluster means Compute + Config + Networking logic is distributed.

---

## 🔹 4. Advanced Networking Plugins (Questions 531-540)

**Q531: What is CNI in Kubernetes?**
Container Network Interface. Library for writing plugins to configure network interfaces in Linux containers.

**Q532: How does Calico differ from Flannel?**
Flannel: Simple L3 VXLAN overlay. Calico: BGP (Performance), NetworkPolicies (Security), No encapsulation options.

**Q533: What is the role of Canal in networking?**
Historical combination of Flannel (Data plane/Networking) + Calico (Control plane/Policies).

**Q534: What is Cilium and why is it popular?**
CNI plugin based on eBPF (kernel bytecode). Highly performant, observable, and provides advanced L7 policy/load-balancing without sidecars.

**Q535: How does eBPF enhance Kubernetes networking?**
Allows programming the kernel network stack without modules. Enables super-fast packet filtering, routing, and observability.

**Q536: What is a NetworkPolicy and how is it enforced?**
YAML rule: "Allow traffic from Label A to Label B". Enforced by CNI (Calico/Cilium). If CNI doesn't support it (e.g. standard Flannel), policies are ignored.

**Q537: What are ingress and egress policies?**
Ingress: Incoming traffic to Pod. Egress: Outgoing traffic from Pod.

**Q538: What is a network plugin vs network proxy?**
Plugin (CNI): Sets up interface/IP. Proxy (Kube-proxy/Envoy): Handles Service routing/Load balancing.

**Q539: What is kube-router and when would you use it?**
Slim alternative Kube-proxy + CNI replacement. Uses LVS/IPVS and BGP. Use for simplicity and performance.

**Q540: How does IPAM (IP Address Management) work in K8s?**
CNI plugins call IPAM plugins (host-local, dhcp) to assign unique IPs to Pods from the node's CIDR range.

---

## 🔹 5. Ingress Controllers & Service Mesh (Questions 541-550)

**Q541: What is the difference between an ingress and ingress controller?**
Ingress: Traffic rules (Resource). Controller: The Load Balancer enforcing rules (Software).

**Q542: What is the NGINX ingress controller and how does it work?**
Based on NGINX. Watches Ingress objects. Regenerates `nginx.conf` + Reloads (or uses Lua) to route traffic.

**Q543: What is Istio and how does it relate to Kubernetes?**
Service Mesh. Adds sidecar proxies. Manages traffic between microservices (L7) inside K8s (which usually only does L4).

**Q544: What is Linkerd and how is it different from Istio?**
Lighter, simpler, faster ("Ultralight service mesh"). Written in Rust. Less features than Istio, but easier to operate.

**Q545: How does Envoy proxy work in a service mesh?**
Intercepts all inbound/outbound traffic of a Pod. Performs Retry, Timeout, Circuit Breaking, mTLS encryption.

**Q546: How do mutual TLS (mTLS) connections work in Istio?**
Sidecar A talks to Sidecar B. They authenticate certificates (issued by Istio CA). Traffic is encrypted transparently to app.

**Q547: How do sidecars help with observability and security?**
They see all traffic. They generate metrics (latency, error rate) and access logs automatically without app changes.

**Q548: What is a virtual service in Istio?**
CRD to configure routing rules (Canary w/ weights, Retries, Fault Injection) for a destination.

**Q549: How do you implement traffic shifting with a service mesh?**
Define subsets (v1, v2). Assign weights (90/10). Mesh distributes traffic accordingly.

**Q550: How do you manage circuit breaking in Istio?**
Configure `DestinationRule`. E.g., "Eject this pod if it returns 5xx errors 3 times in 10s".

---

## 🔹 6. Kubernetes at Scale (Questions 551-560)

**Q551: What is the maximum number of pods per node?**
Default is 110. Can be configured (`kubelet --max-pods`). Limited by IP range node CIDR (often /24 = 254 IPs).

**Q552: What are best practices for scaling to thousands of nodes?**
Use larger IP CIDRs (/16 or /12). Optimize etcd (NVMe, separate cluster). Use NodeLocal DNSCache. Split into multiple clusters.

**Q553: How does Kubernetes schedule high-density workloads?**
Use `Descheduler` to rebalance. Adjust `scheduler` throughput params (percentageOfNodesToScore).

**Q554: How do you isolate workloads in large clusters?**
NodePools with Taints/Tolerations. Namespaces. NetworkPolicies.

**Q555: What is resource bin-packing in Kubernetes?**
Efficiently filling nodes to 100% capacity. Handled by Scheduler scoring functions (`LeastRequested` implies spread. `MostRequested` implies bin-pack).

**Q556: How do you configure taints and tolerations at scale?**
Use standardized labels/taints across node groups. Automate via Cluster Autoscaler or Terraform.

**Q557: What is topology-aware scheduling?**
Scheduling pods close to their dependencies (Service/Volume) or spreading them across distinct zones.

**Q558: How do you scale CRDs and operators in large clusters?**
Watch out for "watch" fatigue on API server. Limit operator scope (Namespace-scoped vs Cluster-scoped).

**Q559: What is a horizontal pod autoscaler’s impact at scale?**
Can cause thundering herd on Metrics Server or API Server if thousands of HPAs query every 15s. Increase sync period.

**Q560: How do you avoid API server overload?**
Rate Limiting (API Priority & Fairness). Caching (Informers). Reduce `LIST` calls. Upgrade etcd.

---

## 🔹 7. Security – Advanced Concepts (Questions 561-570)

**Q561: What is the Kubernetes Pod Security Admission (PSA)?**
The standard, built-in replacement for PSP. Controls security level per namespace via labels.

**Q562: What are seccomp profiles and how are they used?**
JSON files defining allowed syscalls. Pods reference them. `RuntimeDefault` is good practice.

**Q563: What is AppArmor in Kubernetes?**
Profile loaded into kernel restricting program capabilities (file access). Annotated on Pod. (Linux only).

**Q564: How do you configure SELinux with Kubernetes?**
Via `securityContext.seLinuxOptions`. Maps container processes to SELinux labels on the host.

**Q565: What are PodSecurityPolicies (PSP) and their alternatives?**
PSP usage is removed in v1.25+. Alternatives: PSA (Built-in), OPA Gatekeeper, Kyverno.

**Q566: What is a container escape and how can it happen?**
Process breaking out of isolation to access Host OS. Caused by Privileged mode, Kernel exploits (Dirty Cow), or Mounting Docker socket.

**Q567: How do you enforce read-only root filesystem?**
`securityContext.readOnlyRootFilesystem: true`. Forces app to write only to mounted Volumes (emptyDir/tmp). Excellent security.

**Q568: What is a minimal base image and why use it?**
Distroless or Alpine. Reduces attack surface (no shell, no package manager) and image size.

**Q569: How do you scan container images for vulnerabilities?**
Trivy, Clair, Anchore. Scan in CI pipeline AND inside the registry/cluster (Admission Controller blocks critical CVEs).

**Q570: What is Kube-bench and how does it help with security?**
Runs checks against CIS Benchmark (Center for Internet Security) for K8s. outputs Pass/Fail report.

---

## 🔹 8. Edge & IoT Integration (Questions 571-580)

**Q571: How does Kubernetes manage edge workloads with limited resources?**
Uses K3s (low footprint). Tolerates flaky networks. Uses static pods or lightweight runtimes.

**Q572: What is the role of MQTT in edge applications with K8s?**
Messaging protocol for IoT. K8s hosts MQTT Broker (Mosquitto/VerneMQ) to ingest sensor data.

**Q573: How can you use K3s on Raspberry Pi devices?**
It ships as a single binary. Supports ARM64. `curl -sfL https://get.k3s.io | sh -`.

**Q574: What is the role of cloud sync in edge clusters?**
Edge clusters run autonomously but sync aggregated data/logs to Cloud for analysis (KubeEdge or custom controllers).

**Q575: What are the challenges in autoscaling edge applications?**
Limited hardware capacity (cannot add nodes). Scale-up might be impossible. Must prioritize workloads.

**Q576: How do you implement offline-first apps with Kubernetes?**
Local container registry. Local volume storage. App logic handles retry/buffering until connectivity restores.

**Q577: What is the role of WASM in IoT applications?**
Extremely lightweight. IoT devices with tiny RAM can run WASM modules managed by K8s instead of heavy Docker containers.

**Q578: What is a lightweight ingress solution for edge clusters?**
Traefik (bundled with K3s) or NGINX (minimal config).

**Q579: How do you implement OTA (Over-the-Air) updates using Kubernetes?**
Simply `kubectl apply` new image tag. Kubelet pulls new image when bandwidth allows.

**Q580: How does 5G/Edge integration affect Kubernetes architecture?**
MEC (Multi-access Edge Computing). K8s clusters sits in Telco tower. Requires ultra-low latency networking plugins.

---

## 🔹 9. WebAssembly (WASM) & Modern Workloads (Questions 581-590)

**Q581: What is WebAssembly (WASM) in cloud-native applications?**
Binary instruction format. Portable compilation target. Runs code near-native speed safely in sandbox.

**Q582: How does WASM differ from containers?**
Containers = OS Virtualization (Namespace/Cgroup). WASM = Application Virtualization (Process sandbox). Faster start, no OS overhead.

**Q583: What is Krustlet?**
A specialized Kubelet implementation (in Rust) that listens to K8s API but runs WASM payloads instead of Docker containers.

**Q584: How do you run WASM workloads in Kubernetes?**
Use a Runtime Class (WasmEdge/WasmTime) with containerd shims. Or use Krustlet.

**Q585: What are the benefits of WASM in K8s?**
Security (Memory safe, capability-based). Speed (Sub-millisecond startup). Size (Kilobytes image).

**Q586: What are the limitations of WASM today in production?**
No full socket/thread support (WASI is evolving). Debugging tools immature. Not all languages compile to WASM easily.

**Q587: How does WASI relate to Kubernetes?**
WebAssembly System Interface. Standardizes how WASM accesses Files/Network/Env. Essential for running server-side on K8s.

**Q588: What are use cases for WASM in microservices?**
Serverless functions, Side filters (Envoy/Istio plugins), Edge computing.

**Q589: How do you secure WASM modules?**
Signed binaries (Wasm signature). Sandbox provides strong isolation by default (Deny-all capabilities).

**Q590: What is SpinKube?**
Project integration Fermyon Spin (WASM framework) with Kubernetes.

---

## 🔹 10. Miscellaneous Deep-Dive Topics (Questions 591-600)

**Q591: What is the difference between a Job and a CronJob?**
Job: Runs once to completion. CronJob: Creates Jobs on a schedule.

**Q592: How do you manage third-party software lifecycle in Kubernetes?**
Helm Charts or Operators (OLM - Operator Lifecycle Manager).

**Q593: What are Kubernetes Plugins vs Extensions?**
Plugins: Extend core components (CNI, CSI, CRI). Extensions: Extend API (CRDs, Aggregated API).

**Q594: How do you configure cluster DNS resolution?**
ConfigMap `coredns` in `kube-system`. Or `nodelocaldns` for caching.

**Q595: What is the difference between external-dns and CoreDNS?**
CoreDNS: Internal Service Discovery. ExternalDNS: Programs Public DNS (Route53/Google DNS) for external access.

**Q596: How do you implement immutable infrastructure with K8s?**
Nodes are ephemeral (never ssh/patch, just replace). Pods are immutable (restart to update config).

**Q597: What is kube-burner and how is it used?**
Tool to stress test K8s clusters. Creates thousands of objects to burn/test control plane limits.

**Q598: How do you test high availability in HA Kubernetes clusters?**
Kill a Master node. Kill etcd leader. Verify API functionality continues.

**Q599: What is the future of Kubernetes in AI/ML workflows?**
Dynamic Resource Allocation (DRA) for GPUs. Ray on K8s. Kueue (Job queueing system).

**Q600: What is the role of the CNCF in Kubernetes development?**
Cloud Native Computing Foundation. Hosts the project. Governance. Marketing. Neutral home preventing vendor lock-in.

---

## 🔹 1. Kubernetes & GPUs / ML Workloads (Questions 601-610)

**Q601: How do you schedule GPU workloads in Kubernetes?**
Use specific Node Selectors/Taints. Set `resources.limits.nvidia.com/gpu: 1`. Requires NVIDIA Device Plugin running on nodes.

**Q602: What is the NVIDIA device plugin?**
DaemonSet that registers GPU capacity with Kubelet. Allows K8s to see GPUs as allocatable resources.

**Q603: How do you limit GPU usage to specific namespaces?**
Use `ResourceQuota` with `requests.nvidia.com/gpu`.

**Q604: What is the `nvidia.com/gpu` resource?**
The extended resource name exposed by the device plugin. Used in Pod spec `resources: active`.

**Q605: How do you manage mixed CPU/GPU workloads?**
Use Taints on GPU nodes (`sku=gpu:NoSchedule`) so only GPU pods (with tolerations) can run there.

**Q606: What is Kubeflow?**
Machine Learning Toolkit for K8s. Orchestrate complex ML pipelines (Jupyter, Training, Serving) on K8s.

**Q607: What’s the difference between TensorFlowJob and PyTorchJob?**
CRDs provided by Kubeflow (Training Operator). They manage distributed training specific to the framework's architecture (Parameter Servers vs AllReduce).

**Q608: How do you manage ML model versioning in Kubernetes?**
Use Model Registry (MLflow) or container image tags. Store weights in Object Storage (S3) referenced by Pod.

**Q609: What is KServe (formerly KFServing)?**
Standard for Model Inference on K8s. Autoscales (even to zero), supports Canary rollouts, and handles Protocol (HTTP/gRPC/V2).

**Q610: How do you scale ML inference workloads with K8s?**
HPA custom metrics (requests per second). Or KEDA (Kafka lag). Or KServe (automatically handles concurrency).

---

## 🔹 2. GitOps & DevOps Practices (Questions 611-620)

**Q611: How do you roll back changes in ArgoCD?**
Click "Rollback" in UI (reverts live state to previous sync). Fix Git repo immediately after, or it will sync back to bad state.

**Q612: What is auto-sync vs manual sync in GitOps?**
Auto: Cluster strictly follows Git automatically. Manual: Admin approves changes (Click Sync). Safer for Prod.

**Q613: How do you manage secrets in a GitOps pipeline?**
SealedSecrets (Encrypted in Git). SOPS (PGP/KMS encrypted). ExternalSecrets (Reference to Vault).

**Q614: What is Kustomize and how is it used in GitOps?**
Template-free configuration management. Uses `base` and `overlays`. Built into ArgoCD/Flux. Excellent for environment variance.

**Q615: How does Helm differ from Kustomize in GitOps?**
Helm: Templating engine (flexible reuse). Kustomize: Patching engine (specific overrides). Often used together (Helm calls Kustomize).

**Q616: How do you set up CI triggers with GitOps?**
CI builds Docker Image -> Pushes to Registry -> Updates image tag in Git Repo (manifest) -> GitOps Controller pulls change.

**Q617: How do you visualize GitOps deployments?**
ArgoCD UI. Shows graph of Git Commit -> App -> Resources -> Pod Status.

**Q618: What’s the risk of Git drift and how to detect it?**
Risk: Manual `kubectl edit` makes cluster different from Git. Detect: Argo/Flux alerts on "OutOfSync".

**Q619: How do you manage multiple tenants using GitOps?**
AppProject CRD (Argo). Restrict which Git repos can deploy to which Namespaces/Clusters.

**Q620: How does GitOps promote compliance and auditability?**
Git History = Audit Log (Who changed what, when). Pull Request = Change Approval process.

---

## 🔹 3. Admission Controllers & Policy Enforcement (Questions 621-630)

**Q621: What is an admission controller in Kubernetes?**
Interceptors in API Server. Can Mutate (change default values) or Validate (deny invalid requests) objects before creation.

**Q622: What is the difference between validating and mutating webhooks?**
Mutating runs first (e.g. set default imagePullPolicy). Validating runs second (e.g. check image signature).

**Q623: What is OPA Gatekeeper?**
Policy Controller. Uses Rego language to define rules ("All images must come from quay.io"). Enforces them via Validating Webhook.

**Q624: How do you write policies with Rego in OPA?**
Declarative query language. `deny[msg] { input.request.kind.kind == "Pod"; ... }`. Steep learning curve.

**Q625: What is Kyverno?**
Policy engine that uses YAML for policies. Easier than OPA. Supports Mutation and Validation.

**Q626: How does Kyverno differ from OPA Gatekeeper?**
Kyverno is K8s-native (YAML policies). OPA is general purpose (Rego). Kyverno generates Policy Reports natively.

**Q627: How do you enforce image registry policies?**
Policy: `input.request.object.spec.containers[_].image startswith "my-registry.com/"`. Block others.

**Q628: How do you prevent privilege escalation in policies?**
Block `securityContext.allowPrivilegeEscalation: true` and `privileged: true`.

**Q629: How do you test policies before enforcing?**
Run in `Warn` or `Audit` mode (Dry Run). Check logs for violations without blocking deployments.

**Q630: What is a dry-run policy evaluation?**
Gatekeeper/Kyverno allows setting policies to "audit" action. It creates violation events but allows the Admission.

---

## 🔹 4. Resource Scheduling & Placement (Questions 631-640)

**Q631: How does Kubernetes score nodes for pod scheduling?**
Filters (Resources available?) -> Priority Functions (Image Locality, Least Requested, Taint Toleration) -> Weighted Sum.

**Q632: What is inter-pod affinity vs anti-affinity?**
Affinity: Run close to Pod X (Performance). Anti-Affinity: Run far from Pod X (HA/Spread).

**Q633: How does `preferredDuringSchedulingIgnoredDuringExecution` work?**
Scheduler tries to find matching node. If none, picks any valid node. "Best Effort" placement.

**Q634: What is nodeSelector vs node affinity?**
Selector: Exact match label (Hard). Affinity: Expressive rules (In, NotIn, Gt), Hard/Soft modes.

**Q635: What is a topologySpreadConstraint?**
Controls skew of Pods across topology keys (zone, hostname). Ensures even distribution (maxSkew).

**Q636: How do you pin a pod to a specific NUMA node?**
Use Topology Manager feature in Kubelet (`--topology-manager-policy=single-numa-node`). aligned CPU/Memory.

**Q637: What is custom scheduler in Kubernetes?**
A secondary scheduler binary. You can instruct specific Pods to use it via `schedulerName: my-scheduler`.

**Q638: What’s the difference between default-scheduler and kube-scheduler?**
Same thing. kube-scheduler is the binary name. Default-scheduler is the name of the profile it runs.

**Q639: How do you write scheduling extenders?**
HTTP webhooks called by the main scheduler to filter/prioritize nodes based on external logic.

**Q640: How do taints and tolerations impact scheduler decisions?**
Filter step. If Node has Taint X, and Pod lacks Toleration X, Node is discarded immediately.

---

## 🔹 5. Controller Manager Deep Dive (Questions 641-650)

**Q641: What are the responsibilities of kube-controller-manager?**
Single binary running all standard controllers (Node, ReplicaSet, Endpoint, Namespace, ServiceAccount, etc.).

**Q642: What is the replication controller vs replica set?**
RC is legacy. RS is current (supports set-based selectors). Deployments manage RS.

**Q643: What is the horizontal pod autoscaler controller?**
Runs in controller-manager. Periodically (15s) queries Metrics Server and updates `replicas` field of Scale targets.

**Q644: How does the endpoint controller work?**
Watches Services and Pods. Updates `Endpoints` object with IPs of Pods matching Service Selector.

**Q645: What is the garbage collector controller?**
Deletes dependent objects (orphans) when owner is deleted. (e.g. Deletes RS when Deployment is deleted).

**Q646: What does the namespace controller do?**
When Namespace is deleted, it deletes all resources inside it. Once empty, it finalizes the Namespace deletion.

**Q647: How does service account controller operate?**
Ensures every Namespace has a `default` ServiceAccount. Creates API tokens (secrets) for SAs (in older K8s).

**Q648: What is the job controller responsible for?**
Watches Job objects. Creates Pods to run task. Tracks completions. Retries failures up to backoff limit.

**Q649: How do controllers detect stale state?**
They list/watch API server. If connection breaks, they re-list (resync) to ensure local cache is accurate.

**Q650: What happens when a controller crashes?**
Kube-controller-manager Pod restarts (managed by Static Pod/Kubelet). HA setup has standby instance take over lease.

---

## 🔹 6. Cluster Lifecycle & Tools (Questions 651-660)

**Q651: What is kubeadm and when should you use it?**
Official bootstrapper. Use for building vanilla clusters on VM/Bare-metal. Best practice ref impl.

**Q652: How does Rancher simplify Kubernetes management?**
Web UI to manage multiple clusters (EKS, GKE, RKE). Handles Auth, Upgrades, Monitoring, catalog.

**Q653: What is kops and how does it differ from kubeadm?**
Kops provisions Infrastructure (AWS EC2/VPC) AND K8s. Kubeadm only installs K8s on existing machines.

**Q654: What is kind (Kubernetes in Docker)?**
Runs K8s nodes as Docker containers. Fast startup. Great for CI/local testing.

**Q655: How does Minikube work?**
Runs single-node cluster in a VM (VirtualBox/KVM) or Container. Good for learning.

**Q656: What are common tools for creating local K8s clusters?**
Kind, Minikube, K3d (K3s in Docker), Docker Desktop, MicroK8s.

**Q657: How do you use Terraform to manage Kubernetes clusters?**
Provision Cloud Infra (EKS module). Then use Kubernetes Provider to manage generic resources (Namespaces, Quotas).

**Q658: What is Crossplane?**
Framework to manage Cloud Resources (S3, RDS) using K8s Manifests. "Infrastructure as Data".

**Q659: What are the benefits of GitOps with Crossplane?**
Unified workflow. App deployment (K8s) and Infra provisioning (Cloud) both managed via Git commits.

**Q660: What is Cluster API (CAPI)?**
K8s operator to provision K8s clusters. Define `Cluster` and `Machine` resources in YAML.

---

## 🔹 7. Advanced Ingress Strategies (Questions 661-670)

**Q661: How do you manage multi-tenant ingress controllers?**
Deploy one Ingress Class per tenant. Or share one Controller but use strict Ingress/Host validation rules.

**Q662: What is the difference between ingressClass and annotation-based ingress?**
Annotation (`kubernetes.io/ingress.class`) is deprecated. `IngressClass` resource is the standard way to select implementation.

**Q663: What is a canary ingress?**
Splitting traffic percentage using annotations (NGINX) or CRDs (Istio). "10% to service-v2".

**Q664: How do you configure rate limiting at the ingress level?**
Annotations. `nginx.ingress.kubernetes.io/limit-rps`. Redis is often used for global counting.

**Q665: What is SSL passthrough in ingress?**
Traffic is passed encrypted to the backend Pod. Ingress does not terminate TLS. (No L7 inspection possible).

**Q666: How do you enable backend protocol upgrades (e.g., WebSockets)?**
Most controllers support it out-of-box. NGINX: `proxy_http_version 1.1`, Upgrade headers managed automatically.

**Q667: How do you handle client IP preservation?**
`externalTrafficPolicy: Local` on Service. Or receive PROXY protocol from LoadBalancer.

**Q668: What’s the impact of large headers on ingress performance?**
Buffer overflows or 413 Errors. Must tune `client_header_buffer_size` in config.

**Q669: What is the difference between L7 and L4 ingress controllers?**
L7 (Nginx/Traefik): Understands HTTP/Path/Host. L4 (MetalLB/BGP): Just forwards TCP/UDP packets.

**Q670: How do you handle gRPC services through ingress?**
Requires HTTP/2 support. Ingress needs specific annotations (`nginx.ingress.kubernetes.io/backend-protocol: "GRPC"`).

---

## 🔹 8. Energy Efficiency & Green Kubernetes (Questions 671-680)

**Q671: What is energy-aware scheduling in Kubernetes?**
Experiments to schedule pods in zones with low Carbon Intensity or cooler datacenters. Kepler project helps.

**Q672: What tools help measure energy consumption in clusters?**
Kepler (Kubernetes-based Efficient Power Level Exporter). Scrapes RAPL/ACPI metrics via eBPF.

**Q673: How can auto-scaling be optimized for power saving?**
Scale to zero (KEDA) for idle apps. Bin-pack tightly to shut down nodes.

**Q674: What is Kube-green?**
A controller that sleeps (scales down) your dev namespaces at night/weekends to save CO2/Cost.

**Q675: What is idle pod hibernation?**
Stopping the container but keeping state (not native yet, achievable with Checkpoint/Restore or Scale-to-zero).

**Q676: How can you use spot/preemptible instances to reduce cost & energy?**
Use them for batch jobs or fault-tolerant stateless apps. Cloud providers reuse idle hardware (efficient).

**Q677: What is node hibernation and how to automate it?**
Cluster Autoscaler removes empty nodes. Karpenter can consolidate workloads aggressively.

**Q678: How do you offload non-critical workloads during peak energy hours?**
CronJobs to pause non-essential deployments. Carbon-aware scheduling plugins (carbon-aware-sdk).

**Q679: What are the trade-offs in aggressive resource packing?**
Higher energy efficiency but risk of Performance interference (Noisy neighbor), Throttling, and Cascading failures.

**Q680: What are container-native power optimizations?**
Building slim apps (compiled languages). Using ARM processers (Graviton/Ampere) which are more power-efficient.

---

## 🔹 9. Day-2 Ops & Maintenance (Questions 681-690)

**Q681: How do you detect and mitigate resource starvation?**
Monitor `kube_pod_container_resource_limits_cpu_cores` vs usage. Use ResourceQuotas. Prioritize Critical Pods.

**Q682: What are signs of etcd stress or failure?**
High latency (p99 > 100ms). Leader election failures. "Database space exceeded" errors.

**Q683: How do you recover from etcd data corruption?**
Stop etcd instances. Restore from last known good snapshot (`etcdctl snapshot restore`). Start new cluster.

**Q684: How do you rotate kubelet certificates?**
Enable `RotateKubeletClientCertificate` in Kubelet config. It requests new certs from API. Controller Manager approves them.

**Q685: What are common upgrade pitfalls in managed services?**
Deprecated APIs (v1beta1). PDBs blocking node drain. Admission Webhooks failing (blocking all upgrades).

**Q686: How do you automate node OS patching?**
Kured (Kubernetes Reboot Daemon). Watches for `/var/run/reboot-required`, drains node, reboots, uncordons.

**Q687: How do you audit failed pod startups?**
Query `kube_pod_container_status_waiting_reason`. Dashboard or ELK logs filtering for events "FailedCreate".

**Q688: How do you debug stuck terminating pods?**
Check if Finalizer is hung (`kubectl get pod -o yaml`). Force delete (`--grace-period=0 --force`) if necessary (Caution!).

**Q689: How do you manage orphaned PVs?**
List PVs with status `Released`. Check ReclaimPolicy. Manually delete or archive data.

**Q690: How do you update a running container’s environment variables?**
You cannot. Must modify Pod Spec (Deployment) and recreate the Pod. Env vars are injected at start.

---

## 🔹 10. Real-World Design & Trade-offs (Questions 691-700)

**Q691: How do you decide between monorepo vs multirepo in Kubernetes GitOps?**
Monorepo: Easier shared configs/refactoring. Multirepo: Strict isolation, clear ownership. Argo/Flux support both.

**Q692: What are trade-offs of large multi-tenant clusters?**
Pros: Efficient resource pooling, less admin overhead. Cons: "Blast radius" of master failure is huge. Complex security isolation.

**Q693: What is the best way to manage 1000+ microservices on K8s?**
Service Mesh (Observability). GitOps (Automation). Developer Platform (Backstage) to abstract YAML.

**Q694: When should you avoid Kubernetes?**
Simple monolithic app. Small startup team (Ops overhead). Stateful legacy app not designed for ephemeral infra.

**Q695: How do you architect for zero-downtime releases?**
Rolling Updates + Health Probes. PDBs. Graceful Shutdown handling in app code. Database backward compatibility.

**Q696: What’s the difference between application and infrastructure Helm charts?**
App: Deploys Deployment/Service. Infra: Deploys Datadog, Ingress Controller, CertManager. Managed separately.

**Q697: How do you handle secrets across multiple cloud providers?**
External Secrets Operator fetching from AWS Secrets Manager AND Azure KeyVault based on where it runs.

**Q698: What are cost implications of overprovisioning in K8s?**
Paying for idle "Requests". Cluster Autoscaler won't scale down if Requests fill the node, even if Usage is 0%.

**Q699: When is a serverless solution better than Kubernetes?**
Event-driven, sporadic traffic, simple function logic. (Lambda/CloudRun). Less Ops than maintaining K8s nodes.

**Q700: How do you future-proof your Kubernetes platform?**
Stick to standard APIs (Ingress, Gateway API). Avoid vendor-proprietary annotations where possible. Use CNCF graduated projects.

---

## 🔹 1. Kubernetes Storage & CSI Drivers (Questions 701-710)

**Q701: What is a Container Storage Interface (CSI)?**
Standard interface decoupling storage plugins from K8s core code. Allows vendors (EBS, Ceph) to release updates independently.

**Q702: How do you install a CSI driver in Kubernetes?**
Usually a Helm Chart or Manifests. Deploys Controller (StatefulSet) and Node Drivers (DaemonSet).

**Q703: What is dynamic volume provisioning?**
When a PVC is created, K8s automatically talks to the Storage Provider to create the physical volume (PV) without admin intervention.

**Q704: What is the difference between ReadWriteOnce and ReadWriteMany?**
RWO: Mountable by pods on ONE node. RWX: Mountable by pods on MANY nodes (Shared FS like NFS).

**Q705: How do you mount the same volume across multiple pods?**
If pods are on same node, RWO works. If different nodes, must use RWX (NFS/EFS/Ceph) or ReadOnlyMany.

**Q706: What is a VolumeSnapshot and how is it used?**
CRD to trigger a storage-level snapshot (backup). Referenced by a VolumeSnapshotClass.

**Q707: How do you clone a PVC in Kubernetes?**
Create a new PVC and specify `dataSource: { kind: PersistentVolumeClaim, name: source-pvc }`. CSI driver handles cloning.

**Q708: What’s the difference between ephemeral and persistent volumes?**
Ephemeral: Tied to Pod (emptyDir, secret, configMap). Persistent: Independent lifecycle (EBS, NFS).

**Q709: How do you implement volume resizing?**
Allow `volumeExpansion` in StorageClass. Edit PVC size. File system resizes automatically on next Pod restart (or online if supported).

**Q710: What is storageClass and how is it used?**
Template for creating PVs. Defines `provisioner`, `reclaimPolicy`, and parameters (e.g. disk type `gp3`).

---

## 🔹 2. Logging, Monitoring & Observability (Questions 711-720)

**Q711: What is the recommended logging architecture for Kubernetes?**
Node-level logging agent (DaemonSet) reading `/var/log/containers` and forwarding to central backend.

**Q712: How does Fluent Bit differ from Fluentd?**
Fluent Bit is lightweight (C language), lower CPU/RAM. Fluentd is feature-rich (Ruby). Fluent Bit is preferred for K8s nodes.

**Q713: What is Loki and how does it integrate with Promtail?**
Loki is log store (indexed by label). Promtail is the agent (DaemonSet) that tails logs and pushes into Loki. Efficient stack.

**Q714: How do you collect application logs in Kubernetes?**
Apps write to STDOUT/STDERR. Container runtime captures them. Log agent reads the capture files.

**Q715: What is the ELK stack and how does it relate to K8s?**
Elastic-Logstash-Kibana. Traditional logging. Heavy on resources (Java). Being replaced by EFK (Fluentd) or PLG (Promtail-Loki-Grafana).

**Q716: How do you forward logs to external providers?**
Configure Fluentd/FluentBit "Output" plugins to send to Datadog, Splunk, CloudWatch, or S3.

**Q717: How does Prometheus service discovery work?**
It queries K8s API for Endpoints/Pods/Nodes. It dynamically builds list of targets to scrape based on annotations/labels.

**Q718: What are Prometheus relabeling rules?**
Logic to rewrite/filter metrics labels before storage. E.g., drop high-cardinality labels or map `__meta_kubernetes_pod_name` to `pod`.

**Q719: What is a blackbox exporter in Prometheus?**
Probes endpoints (HTTP, ICMP, TCP, DNS) from outside. Checks connectivity and latency (e.g. "Is Google.com reachable from this pod?").

**Q720: How does OpenTelemetry enhance Kubernetes observability?**
Provides a unified way to collect Metrics, Logs, and Traces from apps and infra, reducing vendor lock-in.

---

## 🔹 3. Helm Advanced Usage (Questions 721-730)

**Q721: What is a Helm release lifecycle?**
Install -> Upgrade -> Rollback -> Uninstall. Tracks state via Secrets/ConfigMaps in namespace.

**Q722: How do you upgrade a Helm chart without downtime?**
Ensure Deployment has `rollingUpdate` strategy. Pre-check values. Use `helm upgrade --atomic` to auto-rollback on failure.

**Q723: How do you manage secrets with Helm charts?**
Don't put secrets in values.yaml. Use an external secret provider (Vault, AWS Secrets) or Helm Secrets plugin (SOPS).

**Q724: What are Helm hooks and how do they work?**
Job/Pod that runs at specific phase (`pre-install`, `post-upgrade`). Useful for DB migrations.

**Q725: What is the difference between `values.yaml`, `secrets.yaml`, and `Chart.yaml`?**
Values: Default config. Secrets: (User convention) encrypted config. Chart: Metadata (version, name, dependencies).

**Q726: How do you package and distribute a custom Helm chart?**
`helm package mychart/` -> `.tgz`. Upload to OCI registry (or http repo). `helm push`.

**Q727: What are subcharts in Helm?**
Dependencies used by a parent chart. Referenced in `Chart.yaml`. Can override their values from parent.

**Q728: What is `helm lint` and how is it used?**
Checks Chart for syntax errors and best practices before packaging.

**Q729: How do you rollback a Helm release?**
`helm rollback <name> <revision>`. Re-deploys the manifest from that historic revision.

**Q730: How do you test Helm charts using Helm unittest?**
Plugin allowing you to write unit tests for templates (e.g. "If I set X=true, verify deployment has label Y").

---

## 🔹 4. Platform Engineering with Kubernetes (Questions 731-740)

**Q731: What is an Internal Developer Platform (IDP)?**
Layer above K8s (Backstage, Port). Abstracts complexity. Devs click "Create Microservice" -> Platform handles K8s/Cloud details.

**Q732: How does Backstage integrate with Kubernetes?**
Plugin helps devs view Pod status, logs, and errors for their service directly in Backstage portal.

**Q733: What is score.dev and how does it help platform teams?**
Workload specification. Devs write `score.yaml` (agnostic). It translates to K8s/Helm/Docker-Compose.

**Q734: What’s the role of a Kubernetes platform team?**
Provide K8s as a Service. Manage Upgrades, Security, Policies, Base Images, and Observability for product teams.

**Q735: How do you define golden paths for developers?**
Pre-approved templates (Helm Charts) and CI pipelines ensuring Security/Standards are met by default.

**Q736: What tools help automate environment creation?**
Crossplane, Terraform, ArgoCD ApplicationSets, vCluster.

**Q737: What is a self-service deployment portal?**
GUI where devs trigger deployments or provision DBs without raising tickets or needing Kubectl access.

**Q738: How do you secure multi-tenant platform services?**
NetworkPolicies (deny-all cross-tenant). OIDC Authentication. RBAC isolation. Quotas.

**Q739: What is GitHub Copilot for Kubernetes manifests?**
AI helper in IDE. Auto-completes YAML. Can explain complex regex or configuration errors.

**Q740: What is Kratix and how is it used?**
Platform framework. Allows defining "Promises" (Service offerings) and associating them with Workflows.

---

## 🔹 5. CNCF Ecosystem Awareness (Questions 741-750)

**Q741: What is the CNCF Landscape?**
Massive directory of Cloud Native projects (Service Mesh, Storage, Observability) grouped by functionality.

**Q742: How do tools like ArgoCD and Flux differ?**
Argo: UI-centric, Application concept. Flux: Controller-centric, OCI support, GitOps Toolkit.

**Q743: What is Dapr and how does it work with Kubernetes?**
Distributed Application Runtime. Sidecar providing APIs for PubSub, State Store, Secrets. Decouples app from infra code.

**Q744: What is OpenFunction?**
FaaS (Function as a Service) platform on K8s. Uses Knative, KEDA, and Dapr.

**Q745: What is LitmusChaos used for?**
Cloud-native Chaos Engineering. Fits into Jenkins/GitLab to run chaos tests in pipelines.

**Q746: What is Keptn and how does it enable SLO-based delivery?**
Life-cycle orchestration. "Quality Gates". If Deployment fails SLO (latency > 200ms), Keptn stops promotion.

**Q747: What is OpenKruise?**
Extended workload controllers (Advanced StatefulSet, BroadcastJob, SidecarSet) for large-scale management.

**Q748: What is Harbor and how does it help?**
Open-source Registry. Supports Image Scanning, Signing (Notary), and Replication.

**Q749: What is Crossplane’s role in infrastructure-as-code?**
Extends K8s to manage Cloud (AWS S3, RDS). Eliminates need for separate Terraform state file. Uses K8s database (etcd).

**Q750: What is KEDA (Kubernetes Event-Driven Autoscaling)?**
Scaler. Watches external source (RabbitMQ, Kafka, SQL). Scales deployment 0 -> 1 -> N based on lag.

---

## 🔹 6. Edge Cases & Troubleshooting Scenarios (Questions 751-760)

**Q751: How do you recover a deleted namespace?**
If stuck in `Terminating`: Remove the `finalizers` from the Namespace JSON manually (via raw API call).

**Q752: What happens if a container in a pod crashes continuously?**
Pod status `CrashLoopBackOff`. Kubelet keeps restarting it with exponential backoff delay (10s, 20s, ... 5m).

**Q753: How do you prevent log flooding in a crash loop?**
Fix the app. Or ensure log rotation is aggressive. Or use logging agent that drops duplicate lines.

**Q754: What does `Terminating` status mean and how to debug it?**
Pod is shutting down. Waiting for preStop hook or Finalizer. Check `kubectl describe pod`.

**Q755: What is a zombie pod and how do you clean it up?**
Pod exists in API but node is gone. Force delete `kubectl delete pod <name> --force --grace-period=0`.

**Q756: How do you detect node pressure conditions?**
Status conditions `MemoryPressure`, `DiskPressure`. Check Kernel logs (dmesg) for OOM.

**Q757: What causes a pod to be stuck in `ContainerCreating`?**
Mounting volume failed, Pulling image failed, CNI plugin failed (IP allocation), or Sandbox creation error.

**Q758: What is an image pull backoff error?**
K8s cannot pull image. Check Image name, Tag, Secret (Auth), Network connectivity.

**Q759: How do you detect and clean orphaned resources?**
Tools like `kubectl-janitor`, `fissile`, or simply checking resources without `OwnerReferences`.

**Q760: How do you handle resource quota starvation?**
Increase Quota. Or delete lower priority pods. Or optimized resource requests.

---

## 🔹 7. Advanced Pod Lifecycle & Management (Questions 761-770)

**Q761: What is the difference between liveness and readiness probes?**
Liveness: "Am I alive?" (Restart if no). Readiness: "Can I serve traffic?" (Remove from LB if no).

**Q762: How does a startup probe differ from a liveness probe?**
Startup: Runs ONLY at start. Disables Liveness until it passes. Good for slow-starting legacy apps.

**Q763: What are init containers used for?**
Prerequisites. "Wait for DB". "Download Config". "Fix permissions". Runs before app starts.

**Q764: How do sidecars affect the pod lifecycle?**
They start/stop alongside main container. K8s 1.29+ introduced SidecarContainer feature (starts before main, stops after).

**Q765: What is `terminationGracePeriodSeconds`?**
Time Kubelet gives an app to handle SIGTERM (clean shutdown) before force killing (SIGKILL).

**Q766: How do you ensure graceful shutdown of pods?**
Handle SIGTERM in code. Close connections. Finish requests. Use preStop hook if needed (e.g. Nginx sleep).

**Q767: How do you debug pod termination issues?**
Logs (do you see "Shutting down"?). Timestamps (does it hit 30s timeout?).

**Q768: How does `preStop` hook work?**
Executed *inside* container *before* SIGTERM is sent. Synchronous (blocks termination).

**Q769: What happens if a pod never becomes Ready?**
Deployment will pause rolling updates (if `progressDeadline` exceeded). Old pods serve traffic.

**Q770: How do you prioritize pod scheduling with `priorityClassName`?**
Assign PriorityClass (Value 1000000). High priority pods evict lower priority pods if cluster full.

---

## 🔹 8. Advanced Job & Batch Workloads (Questions 771-780)

**Q771: What is the difference between Job and DaemonSet?**
Job: Runs to completion (Batch). DaemonSet: Runs indefinitely on every node (Service/Agent).

**Q772: How does a CronJob handle missed schedules?**
`startingDeadlineSeconds`. If job missed its window (cluster down), K8s checks if it's within deadline. If not, skips.

**Q773: What is `concurrencyPolicy` in a CronJob?**
`Allow` (Concurrent runs), `Forbid` (Skip if previous running), `Replace` (Kill old, start new).

**Q774: How do you avoid duplicated job runs?**
Hard to guarantee 100%. App should be idempotent. Use `concurrencyPolicy: Forbid`.

**Q775: How can jobs fail silently and how to detect it?**
If `backoffLimit` reached. Check Job Status `Failed`. Alert on `kube_job_status_failed`.

**Q776: How do you implement retry logic in batch jobs?**
`backoffLimit: 6`. K8s restarts pod with exponential delay.

**Q777: What is a parallel job with indexed completion?**
Job runs N pods. Each gets index (0, 1, 2...). Useful for sharded processing (Pod 0 does chunk 0).

**Q778: What are suspend/resume features in Jobs?**
Set `.spec.suspend: true`. Pauses creation of pods. Useful for debugging or flow control.

**Q779: How do you clean up old Jobs automatically?**
`ttlSecondsAfterFinished: 100`. K8s deletes Job object (and pods) 100s after completion.

**Q780: How do you monitor and alert on Job status?**
kube-state-metrics (`kube_job_status_succeeded`). Alert if `kube_job_status_failed > 0`.

---

## 🔹 9. Real-time & Low Latency Workloads (Questions 781-790)

**Q781: How do you run low-latency workloads on Kubernetes?**
Dedicated Nodes. CPU Pinning. HugePages. High-perf CNI (Calico/Cilium). Bypass Kube-proxy (HostNetwork).

**Q782: What is CPU pinning and how is it achieved?**
Assign exclusive CPU cores to container. `CPU Manager Policy: static`. Pod implies Guaranteed QoS (Requests=Limits).

**Q783: How do you configure guaranteed QoS class?**
Set `resources.limits` equal to `resources.requests` for both CPU and Memory.

**Q784: What are real-time kernel considerations?**
Use PREEMPT_RT kernel patches on Node OS. Tune interrupts. Isolate CPUs.

**Q785: How do you ensure predictable scheduling latency?**
Use high-priority classes. Pre-pull images. Avoid heavy Bin-packing.

**Q786: What are HugePages and how are they configured?**
Large memory pages (2MB/1GB). Reduces TLB misses. Set `resources.limits.hugepages-2Mi` in Pod.

**Q787: How does Kubernetes support DPDK applications?**
SR-IOV CNI plugin. Passes physical NIC VF directly to Pod (bypassing kernel network stack).

**Q788: What’s the difference between guaranteed vs best-effort pod QoS?**
Guaranteed: Exclusive resources. Last to be evicted. BestEffort: No requests/limits. First to be evicted.

**Q789: How do you run audio/video processing workloads?**
Use GPU nodes. Mount `/dev/shm` (Shared memory) for fast IPC.

**Q790: How do you reduce cold start times for latency-sensitive pods?**
Over-provision pool (Pause pods). Use Image Puller. Use lighter runtimes (WASM) or snapshots.

---

## 🔹 10. Design Principles & Industry Patterns (Questions 791-800)

**Q791: What is GitOps vs ChatOps in Kubernetes operations?**
GitOps: Declarative (Git). ChatOps: Imperative (Slack bot command). GitOps is safer/auditable.

**Q792: What are 12-factor app considerations in K8s?**
Config in Env/Secrets. Disposability (Fast startup/shutdown). Logs to stdout. Dev/Prod parity.

**Q793: How do you enforce resource limits as best practice?**
Default `LimitRange` in every namespace. `ResourceQuota` to cap aggregate.

**Q794: What are common anti-patterns in Helm usage?**
"Super Chart" (One chart for EVERYTHING). Hardcoded values in templates. No versioning.

**Q795: What is an opinionated vs flexible Kubernetes platform?**
Opinionated (OpenShift/Rancher): Batteries included, rigid. Flexible (EKS/Vanilla): Build your own lego.

**Q796: What are design considerations for multi-cloud clusters?**
Latency between clouds. Egress costs. Data gravity (Database location). Unified Identity.

**Q797: How do you manage regional failover in Kubernetes?**
Global Load Balancer (DNS/Anycast) pointing to Region A and Region B clusters. Apps must be stateless or have DB replication.

**Q798: What is the pet vs cattle analogy in Kubernetes?**
Pets (VMs): Hand-raised, named, nursed. Cattle (Pods): Numbers, identical, replaced if sick. K8s treats everything as cattle.

**Q799: When is it better to use serverless over Kubernetes?**
Short-lived tasks, unpredictable spikes, zero ops capability. "Scale to zero" requirement.

**Q800: How do you align Kubernetes architecture with DORA metrics?**
Deployment Freq (GitOps). Lead Time (Auto-provisioning). MTTR (Self-healing). Change Failure (Canary analysis).

---

## 🔹 1. Kubernetes API & Etcd Internals (Questions 801-810)

**Q801: What is the structure of a Kubernetes API request?**
Group (batch), Version (v1), Resource (jobs). `POST /apis/batch/v1/namespaces/default/jobs`.

**Q802: How does `kubectl` find the API server?**
Reads `~/.kube/config`. Uses `clusters.server` URL and `users` credentials (cert/token) to authenticate.

**Q803: What is the role of the Aggregation Layer in API Server?**
Allows extending the API (e.g., Metrics Server) without modifying core code. Proxies requests to extension apiservers.

**Q804: How does etcd ensure strong consistency?**
Raft consensus algorithm. Writes must be acknowledged by majority (quorum) of nodes before committed.

**Q805: What happens if etcd loses quorum?**
Cluster becomes Read-Only (Safe). Cannot schedule new pods or change state until quorum is restored (2 out of 3).

**Q806: How do you debug slow API responses?**
Metrics (`apiserver_request_duration_seconds`). Check etcd disk latency. Check high CPU on API server.

**Q807: What is the difference between `Watch` and `List`?**
List: Get all objects (expensive). Watch: Open stream, get updates only (efficient). Controllers use `List` once then `Watch`.

**Q808: How does Kubernetes handle optimistic locking?**
Client sends `resourceVersion`. If server has newer version, it rejects request. Client must re-fetch and re-apply.

**Q809: What is a MutatingAdmissionWebhook?**
HTTP callback that receives object (Pod) and can modify it (Inject sidecar, set defaults) before saving to etcd.

**Q810: How do you backup an etcd cluster?**
`etcdctl snapshot save db.db`. Must include certificates if TLS enabled.

---

## 🔹 2. Advanced Cluster Networking (Questions 811-820)

**Q811: What is the CNI (Container Network Interface) specification?**
Standard for configuring network interfaces. Plugins (Calico, Cilium) implement: `ADD` (Setup IP), `DEL` (Teardown).

**Q812: How does BGP (Border Gateway Protocol) work in Calico?**
Nodes exchange routing info. "To reach Pod CIDR A, send to Node A". No overlay needed (high performance).

**Q813: What is the difference between Overlay and Underlay networks?**
Overlay: VXLAN/IPIP tunnel (easy setup, slight overhead). Underlay: Native routing directly on physical network (fast, complex setup).

**Q814: How does IPVS mode improve kube-proxy performance?**
Uses Linux Kernel IPVS (Hash tables) instead of iptables (Linear search). Scales better to 5000+ services.

**Q815: What is a `NetworkPolicy` 'deny-all' rule?**
`podSelector: {}` `policyTypes: [Ingress]`. Blocks all traffic unless allowed. Best practice for zero-trust.

**Q816: How do you troubleshoot DNS resolution issues in Pods?**
`nslookup kubernetes.default`. Check `/etc/resolv.conf`. Check CoreDNS logs. Check NetworkPolicy blocking UDP/53.

**Q817: What is `externalTrafficPolicy: Local`?**
Preserves Client IP. Drops traffic if Pod not on receiving node (no hopping). Can cause load imbalance.

**Q818: How does Cilium use eBPF for networking?**
Bypasses iptables entirely. Programs kernel to route packets directly from socket. Supports L7 visibility efficiently.

**Q819: What is Multus CNI?**
Meta-plugin. Allows attaching multiple network interfaces to a single Pod (e.g. eth0 for K8s, net1 for DPDK/Data).

**Q820: How do you handle IPv6 in Kubernetes?**
Dual-stack clusters. Pods get IPv4 and IPv6 addresses. Service definitions typically match Traffic Family.

---

## 🔹 3. Service Mesh Deep Dive (Istio/Linkerd) (Questions 821-830)

**Q821: What is the sidecar pattern in Service Mesh?**
Deploying a proxy (Envoy) alongside app container. Intercepts all traffic. Adds logic (retry, mTLS) transparently.

**Q822: How does Istio perform mTLS?**
Citadel (CA) issues certs to sidecars. Sidecars handshake. App thinks it's plain HTTP. Network sees encrypted TCP.

**Q823: What is Traffic Splitting in Istio?**
`VirtualService` rule: send 90% traffic to `subset: v1`, 10% to `subset: v2`. Used for Canary releases.

**Q824: How does Linkerd differ from Istio?**
"Sidecar-less" mode (Ambient mesh) available in newer versions. Or Rust-based micro-proxy. Focus on simplicity/speed over features.

**Q825: What is Gateway API vs Ingress API?**
Gateway API is the new standard. Role-based (GatewayClass, Gateway, HTTPRoute). More expressive than Ingress.

**Q826: How do you debug 503 errors in a Service Mesh?**
Check if Sidecar is ready. Check `DestinationRule` (mTLS mismatch?). Check backend application health.

**Q827: What is centralized vs distributed tracing?**
Distributed: Requests carry Trace ID across microservices (Jaeger). Centralized: Aggregating these traces to visualize flow.

**Q828: How do you inject faults using Istio?**
`fault: abort: httpStatus: 500` or `delay: fixedDelay: 5s`. Test app resilience without changing code.

**Q829: What are the performance costs of a Service Mesh?**
Increased latency (proxy hop). CPU/RAM overhead for sidecars. Complexity in debugging.

**Q830: What is Ambient Mesh?**
Sidecar-less architecture (Istio). Layer 4 proxy (ztunnel) per node + Layer 7 proxy (Waypoints) per namespace.

---

## 🔹 4. Secret Management & Compliance (Questions 831-840)

**Q831: Why are Kubernetes Secrets not secure by default?**
Stored as Base64 encoded strings in etcd. Anyone with API access to Namespace can read them.

**Q832: How do you enable Encryption at Rest for Secrets?**
Configure API Server `EncryptionConfiguration`. Use a provider (aescbc, kms) to encrypt etcd values.

**Q833: What is the External Secrets Operator (ESO)?**
Syncs secrets from external robust stores (AWS Secrets Manager, Vault, Azure KV) into K8s Secrets.

**Q834: How does HashiCorp Vault integrate with Kubernetes?**
Sidecar injector (Agent) authenticates against Vault using K8s ServiceAccount. Injects secrets into shared memory/file.

**Q835: What is CSI Secret Store driver?**
Mounts secrets from Vault/AWS directly as Volumes. No "K8s Secret" object is created (more secure, memory only).

**Q836: How do you rotate secrets in Kubernetes?**
Update external source. Operator Syncs. App checks file change (hot reload) or Restarts (if using Env vars).

**Q837: What is kyverno's role in security compliance?**
Enforce policies: "Require all pods to have non-root user". "Require all images to be signed".

**Q838: How do you audit "who read this secret"?**
Enable Kubernetes Audit Logs. Filter for UserAgent, Verb=Get/List, Resource=Secrets.

**Q839: What is Cert-Manager?**
Automates issuance/renewal of TLS certificates (LetsEncrypt, HashiCorp Vault) for Ingress/Pod.

**Q840: How do you scan for secret leaks in manifests?**
Tools like `git-secrets`, `trufflehog`, or check Git history before applying.

---

## 🔹 5. Serverless on Kubernetes (Knative/FaaS) (Questions 841-850)

**Q841: What is Knative Serving?**
Scale-to-zero, request-driven autoscaling, and revision management for stateless containers.

**Q842: How does Knative scale to zero?**
Route traffic to "Activator". If no pods, Activator holds request, scales up Deployment, forwards request.

**Q843: What is Knative Eventing?**
Abstracts delivery of events (CloudEvents) from producers (Github, Kafka) to consumers (Knative Services, triggers).

**Q844: What is OpenFaaS?**
Function-as-a-Service framework. simple UI/CLI. Docker containers as functions.

**Q845: What is the difference between FaaS and PaaS on K8s?**
FaaS: Event-driven, short-lived, ephemeral (Knative). PaaS: Deploy and run long-running apps easily (Heroku-like, e.g. Acorn/Rancher).

**Q846: What are the latency implications of Serverless K8s?**
Cold start issue. When scaling from 0 to 1, first request waits for Pod download + startup (seconds).

**Q847: How does KEDA enable serverless-like scaling?**
Scales standard Deployments based on event queue depth (RabbitMQ > 10 msgs -> Scale Up).

**Q848: What is Cloud Run (Google)?**
Managed Knative. You give a container; Google runs it. No cluster management visible.

**Q849: How do you handle long-running tasks in Serverless?**
Not suitable for FaaS (timeouts). Use Kubernetes Jobs or Async queues with worker pods.

**Q850: What is the advantage of K8s-based Serverless vs Lambda?**
No vendor lock-in. Run same functions on-prem, AWS, or GCP. Standard Docker images.

---

## 🔹 6. Multi-Tenancy & Isolation Patterns (Questions 851-860)

**Q851: What is Soft vs Hard Multi-tenancy?**
Soft: Shared cluster, Namespaces, trusting tenants (Internal teams). Hard: Zero-trust, VM isolation (Kata), different risks (SaaS customers).

**Q852: How does `Hierarchical Namespaces` (HNC) help?**
Allows syncing Roles/ResourceQuotas from Parent Namespace to Child Namespaces. tree structure.

**Q853: How do you isolate networking between tenants?**
default-deny NetworkPolicy. Or CNI isolation (Calico/Cilium). Or Service Mesh authorization policies.

**Q854: What is vCluster (Virtual Cluster)?**
Running a full Control Plane (API, Controller) *inside* a Pod in a host cluster. Tenant thinks they are root.

**Q855: How do you manage fair resource usage in multi-tenant clusters?**
ResourceQuotas (Hard limits). LimitRanges (Default container sizes). PriorityClasses (Prevent preemption).

**Q856: What are Kata Containers?**
Lightweight VMs (microVMs) running as pods. Kernel isolation. High security for Hard Multi-tenancy.

**Q857: How do you restrict node access per tenant?**
Taints (tenant=A:NoSchedule) and Tolerations. Or Node Affinity.

**Q858: What is the role of RBAC in multi-tenancy?**
Bind groups to Roles (Namespace scope), NOT ClusterRoles. Prevent access to Cluster-wide resources (Nodes, PVs).

**Q859: How do you handle storage isolation?**
StorageClasses that provision separate encrypted volumes per claim. Verify policies prevent mounting hostPath.

**Q860: What is namespace lifecycle management?**
Onboarding automation (create NS + quotas + RBAC). Offboarding (delete NS). Tools like Kiosk or Capsule.

---

## 🔹 7. Chaos Engineering & Reliability (Questions 861-870)

**Q861: What is the 'Blast Radius' in chaos engineering?**
The impact scope of a failure. Goal: minimize it. (e.g. 1 pod crash vs entire cluster crash).

**Q862: What is Chaos Mesh?**
CNCF project. Chaos operator. Usage: `kubectl apply -f chaos.yaml` (NetworkDelay, PodKill, KernelPanic).

**Q863: How do you test DNS failures?**
Block UDP/53 in NetworkPolicy or use Chaos Mesh `DNSChaos` to return errors/random IPs.

**Q864: What is a Game Day?**
Planned team event to execute chaos experiments in Staging/Prod to validate runbooks and alert coverage.

**Q865: How do you automate chaos in pipelines?**
LitmusChaos step in GitLab CI / Argo Workflows. "Deploy -> Run Pod Delete -> Check Success".

**Q866: What is a Steady State Hypothesis?**
"The system is normal if 200 OK rate > 99%". Chaos validates if system returns to steady state after injection.

**Q867: How do you test for split-brain scenarios?**
Network Partition between control plane nodes (Etcd leaders). Verify cluster enters Read-Only safe mode.

**Q868: What is resource exhaustion testing?**
Stress test (Stress-ng). consume 100% CPU/Memory. Verify Autoscaler kicks in or Priority-based eviction saves critical apps.

**Q869: How do you ensure safety during chaos testing?**
Abort conditions. If Error Rate > 5%, Stop Experiment automatically. Run only on non-critical / canary first.

**Q870: What is the benefit of random pod deletion (Simian Army)?**
Forces developers to build stateless, resilient apps that handle restarts gracefully at any time.

---

## 🔹 8. Supply Chain Security (Questions 871-880)

**Q871: What is a Software Bill of Materials (SBOM)?**
Inventory of all libraries/packages in your image. "Ingredients list". Standard formats: SPDX, CycloneDX.

**Q872: How do you generate an SBOM for a container?**
Use tools like `syft` or `trivy`. `syft image:master > sbom.json`.

**Q873: What is Sigstore / Cosign?**
Tools for signing containers (keyless). `cosign sign user/demo`. Verifies that image is exactly what CI built.

**Q874: How do you verify image signatures in K8s?**
Admission Controller (Kyverno/Gatekeeper). "Allow Pod only if Image is signed by Key XYZ".

**Q875: What is the SLSA framework?**
Supply-chain Levels for Software Artifacts. Security checklist (Source control, Build server isolation, Provenance).

**Q876: How do you detect malicious images?**
Scanner (Clair/Trivy) checks for Known Malware/CVEs. Behavioral analysis (Falco) checks runtime activity.

**Q877: What is Chainguard Images?**
Vendor providing hardened, minimal, low-CVE base images (Wolfi Linux) implementation.

**Q878: How do you secure your build pipeline?**
Run builds in isolated ephemeral runners (Tekton). Don't use `privileged`. Scan artifacts before push.

**Q879: What is The Update Framework (TUF)?**
Security framework ensures updates (like fetching image layers/metadata) are not tampered with. Used by Sigstore.

**Q880: Why should you avoid `latest` tag in production?**
Not immutable. Supply chain attack can overwrite `latest` with malicious code. Breaks rollback ability.

---

## 🔹 9. Operator Pattern & Controller Development (Questions 881-890)

**Q881: What implies "Level 5" Operator Capability?**
Autopilot. Handles Install (1), Upgrade (2), Lifecycle (3), Metrics (4), and *Auto-Scaling/Tuning/Healing* (5).

**Q882: What is Kubebuilder?**
SDK for building Operators in Go. Scaffolds Project, API types, Controller logic, Tests.

**Q883: How do you test Kubernetes Controllers?**
Unit Test (Fake client). Integration Test (EnvTest - local control plane). E2E Test (Kind cluster).

**Q884: What is the `Spec` vs `Status` design pattern?**
Spec: User Desires (Replica: 3). Status: Controller Reports (Current: 1). Controller works to make Status == Spec.

**Q885: How do to handle external resources in a Controller?**
Add Finalizer. On delete: Run external cleanup logic (API call to AWS), then remove Finalizer.

**Q886: What is a sidecar injection controller?**
Mutating Webhook. Intercepts Pod Create. Patches YAML to add `containers: - name: sidecar ...`.

**Q887: How do you avoid "fighting controllers"?**
Ensure ownership is clear. If 2 controllers edit same field, they will infinite loop.

**Q888: What is client-go `workqueue`?**
Queue used in controllers. Events (Add/Update) go in queue. Workers process items. Retries handles failures.

**Q889: What is bootstrapping an operator?**
Using a small initial operator to install/manage the complex main operator.

**Q890: Why use Go for Kubernetes Controllers?**
Native K8s language. Access to high-quality libraries (`client-go`, `controller-runtime`). Performance/Concurrency.

---

## 🔹 10. Kubernetes Cost Management & FinOps (Questions 891-900)

**Q891: What is Kubecost?**
Tool providing visibility into K8s spend. Allocates cost by Namespace/Label based on Cloud Billing integration.

**Q892: How do you chargeback costs to teams?**
Tag resources (Labels `team=frontend`). Use Kubecost to generate reports per label.

**Q893: What is the difference between Request cost and Usage cost?**
Request: You pay for what you reserve (ResourceQuota). Usage: What you actually burn. Optimize the gap (Waste).

**Q894: How do Spot Instances reduce costs?**
Up to 90% discount. Risk of interruption. Use for stateless, fault-tolerant Batch/API nodes.

**Q895: How does correct sizing (Rightsizing) impact bill?**
Reduces requested resources to match reality. Allows Cluster Autoscaler to consolidate nodes.

**Q896: What is OpenCost?**
Open-source core of Kubecost. CNCF sandbox. Standard spec for cost monitoring.

**Q897: How do you manage data transfer costs in K8s?**
Avoid cross-zone traffic (Topology Aware Routing). Keep Chatty microservices in same availability zone.

**Q898: What is the cost impact of unmanaged storage (zombie PVs)?**
You pay for Disks even if Pod is deleted. Need to implement policy to Delete or Snapshot-then-delete PVs.

**Q899: How does Karpenter improve cost efficiency?**
Just-in-time Node provisioning. Picks the exact cheapest instance type that fits the pending pod (instead of fixed AutoScalingGroups).

**Q900: What is FinOps in the context of Kubernetes?**
Cultural practice. Engineering + Finance + Business. Optimized spend. Accountability (Devs see cost of their deployments).

---

## 🔹 1. Multi-Cluster Networking & Federation (Questions 901-910)

**Q901: What is the concept of a "ClusterSet" in Multi-Cluster Services?**
A group of clusters that share services. A Service exported in Cluster A can be imported in Cluster B seamlessly.

**Q902: How does Submariner work?**
Connects overlay networks of different clusters. Opens a secure tunnel (IPSec/WireGuard) between gateway nodes.

**Q903: What is Project KubeFed (Federation v2)?**
Allows defining "Federated" resources (`FederatedDeployment`) which the control plane propagates to multiple member clusters.

**Q904: How do you handle ingress traffic for multi-region clusters?**
Global Anycast Load Balancer (CloudFlare/AWS Global Accelerator). Routes user to nearest healthy cluster.

**Q905: What is Istio Multi-Primary architecture?**
Each cluster has its own Istio Control Plane (istiod). Sidecars can talk to sidecars in other clusters (via East-West Gateway).

**Q906: What are the challenges of data consistency in multi-cluster?**
CAP theorem. High latency between regions makes strong consistency (for DBs) slow. Usually rely on Eventual Consistency.

**Q907: How do you implement global service discovery?**
CoreDNS plugin `k8s_external`. Or ExternalDNS syncing Service IPs to a global Zone (e.g. `mysvc.global.example.com`).

**Q908: What is a stretch cluster?**
Single K8s cluster spanning multiple zones (or regions). High latency for etcd. Not recommended for multi-region.

**Q909: How do you manage CI/CD for 50+ clusters?**
Pull-based GitOps (ArgoCD) is essential. Central Dashboard. "ApplicationSets" to target list of clusters dynamically.

**Q910: what is the "Hub and Spoke" pattern in multi-cluster?**
One Management Cluster (Hub) runs ArgoCD, Vault, Observability. Many Workload Clusters (Spokes) run the apps.

---

## 🔹 2. Advanced Authentication (OIDC, IRSA) (Questions 911-920)

**Q911: How does OIDC works with Kubernetes?**
User -> Identity Provider (Google) -> ID Token (JWT). User sends Token to K8s. K8s verifies signature and `iss`. Trust established.

**Q912: What is "IAM Roles for Service Accounts" (IRSA) on AWS?**
Pod gets a Web Identity Token. Injects into SDK. AWS STS assumes an IAM Role. No hardcoded AWS Keys in secrets!

**Q913: What is "Workload Identity" on GKE?**
Maps a K8s ServiceAccount (KSA) to a Google ServiceAccount (GSA). Pod runs as GSA.

**Q914: How do you perform group-based RBAC with OIDC?**
ID Token contains `groups: ["admins"]`. ClusterRoleBinding binds `group: admins` to `cluster-admin`.

**Q915: What is Dex?**
Open-source OIDC provider/proxy. Connects K8s to LDAP, SAML, GitHub, etc. Acts as the "glue".

**Q916: What is the downside of using Client Certificates for users?**
Hard to revoke. If a laptop is stolen, you must rotate the entire CA or re-issue all certs. Prefer OIDC.

**Q917: How does `kubectl` cache tokens?**
Does not refresh tokens automatically by default. Use `exec` plugins (aws-iam-authenticator / kubelogin) to fetch fresh tokens.

**Q918: What is the `--oidc-issuer-url` flag?**
Configures API Server to trust a specific OIDC provider (e.g. `https://accounts.google.com`).

**Q919: How do you handle multi-tenancy auth?**
Specific groups per tenant. `dev-team-a` group gets `edit` role in `ns-team-a` only.

**Q920: What is SPIFFE and SPIRE?**
Standards for workload identity. SPIRE issues short-lived X.509 certs (SVIDs) to workloads tailored for mTLS across heterogeneous environments.

---

## 🔹 3. Disaster Recovery & Backup Strategies (Questions 921-930)

**Q921: What is the difference between High Availability (HA) and Disaster Recovery (DR)?**
HA: Surviving a node/zone failure (Automatic). DR: Surviving a Region/Data-center wipeout (Manual/Orchestrated recovery).

**Q922: How does Velero work?**
Backs up API Objects (YAML) to S3. Backs up Persistent Volumes (Restic/Snapshots) to S3. Restores them on demand.

**Q923: What is RPO and RTO in Kubernetes DR?**
RPO (Point Objective): How much data loss? (Snapshots every hour = 1hr RPO). RTO (Time Objective): How long to restore?

**Q924: How do you backup etcd manually?**
`etcdctl snapshot save`. Store securely offsite. Test restore procedures regularly!

**Q925: What is Active-Passive DR strategy?**
Prod Cluster (Active). DR Cluster (scaled down/empty). Data replication matches RPO. On failover, scale up DR cluster.

**Q926: What is Active-Active DR?**
Both clusters handle traffic. Requires complex data replication (Multi-master DBs) and Geo-LB.

**Q927: How do you handle non-cloud PV backups?**
Use Velero's generic backup tool (Restic/Kopia) which copies filesystem bits to Object Storage.

**Q928: Why is "Configuration Drift" dangerous for DR?**
If DR cluster config doesn't match Prod (missing secrets/CRDs), restore will fail. Use GitOps to sync both.

**Q929: How do you test a DR plan?**
Schedule a "Fire Drill". Spin up new cluster from backups. Verify app functionality. Measure actual RTO.

**Q930: What is Kasten K10?**
Enterprise backup solution. supports application-aware backups (Blueprints) for Databases (flushing buffer before freeze).

---

## 🔹 4. Compliance & Policy as Code (Questions 931-940)

**Q931: What is Policy as Code?**
Defining security/compliance rules (No Root, Limit Ranges) as version-controlled code (Rego/YAML) enforced automatically.

**Q932: How do you allow exceptions to policies (e.g. for System Pods)?**
Exclude namespaces (`kube-system`, `monitoring`). Or Exclude based on User (`system:serviceaccount:...`).

**Q933: What is a Pod Security Standard (PSS)?**
Official define levels: `Privileged` (Unrestricted), `Baseline` (Minimal restrictions), `Restricted` (Hardening).

**Q934: What is the difference between enforcement vs auditing mode?**
Enforcement: Block request (Deploy fails). Audit: Allow request, but log it as a violation (for monitoring/cleanup).

**Q935: How does NIST 800-190 apply to Kubernetes?**
Security Guide for Container apps. Covers: Image Vulns, Registry Trust, Orchestrator Security, Runtime Security.

**Q936: How do you restrict container capabilities?**
`securityContext.capabilities.drop: ["ALL"]`. Only add back what is strictly needed (e.g. `NET_BIND_SERVICE`).

**Q937: What is drift detection in compliance?**
Noticing when a resource (deployed via kubectl edit) violates the Git-defined policy or runtime constraints.

**Q938: How do you limit HostPath usage?**
Policy: Block `hostPath` volumes. They bypass storage isolation and allow node filesystem access.

**Q939: What is the "least privilege" principle in K8s RBAC?**
Grant only required verbs on required resources. Use `Role` instead of `ClusterRole`. Use `ServiceAccount` per app.

**Q940: How do you ensure CIS Benchmark compliance?**
Run `kube-bench` regularly. Remediate failed checks (e.g. file permissions on etcd certs).

---

## 🔹 5. Microservices Patterns on Kubernetes (Questions 941-950)

**Q941: What is the Ambassador pattern?**
A sidecar container acts as a proxy for the main app to external world (e.g. Auth proxy, Connection pool).

**Q942: What is the Adapter pattern?**
Sidecar normalizes output. e.g. App logs in random format -> Adapter converts to JSON -> Logger reads JSON.

**Q943: What is the Sidecar pattern?**
Helper container extending functionality (Log shipping, Proxy, Config reloading) without changing app code.

**Q944: What is the Leader Election pattern for microservices?**
Only one replica does the "Work" (Cron/Writer). Use `leaderelection` library to lock a text record in K8s.

**Q945: How do you handle distributed transactions (Saga)?**
Not K8s specific. Orchestrator (Temporal/Cadence) manages sequence of calls. Retries/Compensation via Pods.

**Q946: What is Bulkhead pattern?**
Isolating pools of resources. Separating ThreadPools. If Service A is slow, it doesn't starve Service B.

**Q947: What is Circuit Breaker pattern?**
Stop calling a failing service to allow it to recover. Implemented in App code (Hystrix) or Service Mesh.

**Q948: What is Service Discovery pattern?**
Using K8s DNS (`my-svc`) instead of hardcoded IPs. Decoupling location from consumption.

**Q949: What is the Self-Healing pattern?**
Liveness probes restart hung apps. Deployment controller replaces deleted pods.

**Q950: How do you implement the "BFF" (Backend for Frontend) pattern?**
Deploy distinct API Services for Mobile vs Web, aggregated by Ingress or Gateway.

---

## 🔹 6. Developer Experience (DevX) tools (Questions 951-960)

**Q951: What is Skaffold?**
Google tool. Watches code. On save: Builds Image, Pushes, Deploys to K8s (Kind/Remote). Fast feedback loop.

**Q952: What is Telepresence?**
CNCF tool. Connects your local machine to the cluster network. "Swa" an existing deployment with your local process.

**Q953: How does Tilt help developers?**
Dev environment tool. "Tiltfile" defines how to build/deploy multiple microservices updates instantly (Live Update).

**Q954: What is the downside of remote development (ssh)?**
Latency. Network dropouts. Lack of local IDE responsiveness (though VSCode Remote helps).

**Q955: What is a "Preview Environment"?**
Ephemeral namespace created automatically per Pull Request. Validates changes in isolation. Deleted on Merge.

**Q956: What is Lens IDE?**
Popular desktop dashboard for K8s. Easy viewing of multiple clusters, logs, shell access.

**Q957: How do you debug a pod that crashes instantly?**
`kubectl run -it --rm --image=<img-name> debug -- sh`. Or `kubectl debug` ephemeral containers.

**Q958: What is "Okteto"?**
Platform providing remote cloud development environments. "Code locally, run in cloud container".

**Q959: How do you manage local ports vs cluster ports?**
`kubectl port-forward`. Maps localhost:8080 -> Pod:80. Quick and dirty access.

**Q960: What is K9s?**
Terminal UI (TUI) for K8s. Vim-like shortcuts. Fast, efficient access to logs/shell/yamls.

---

## 🔹 7. Debugging & Performance Tuning (eBPF) (Questions 961-970)

**Q961: What is eBPF?**
Extended Berkeley Packet Filter. Sandbox in Linux Kernel. Safely run custom bytecode (observability/security) without recompiling kernel.

**Q962: How does Pixie use eBPF?**
Auto-telemetry. Instantly sees HTTP bodies, SQL queries, Latency maps without manual instrumentation (sidecars).

**Q963: How do you debug high DNS latency?**
Check `conntrack` table (full?). Check `nodelocaldns` cache hits. Use `dnstap` to trace queries.

**Q964: What is "Perf" tool/Flamegraphs?**
Profiling CPU usage. Collect samples of stack traces. Visualize where CPU time is spent (e.g. Garbage Collection).

**Q965: How do you analyze Network Throttling?**
Check if `dropped` packets increase. Check `bandwidth` CNI plugin limits. Check Cloud Provider Limits (e.g. EC2 network credits).

**Q966: What is CPU Throttling?**
When Pod uses quota (`requests` to `limits`). CFS Scheduler pauses the process. Causes latency even if Node is idle.

**Q967: How do you detect OOM Kills at node level?**
`dmesg | grep -i kill`. System OOM killer sacrifices processes to save Kernel.

**Q968: How do you troubleshoot a slow persistent volume?**
`iostat`. Check PV IOPS limits. Cloud Disk burst balance. FIO benchmark inside pod.

**Q969: What is "kubectl debug"?**
New command. Attaches an "Ephemeral Container" (with debugging tools like curl/dig) to a running Pod (distroless).

**Q970: How do you trace system calls?**
`strace -p <pid>`. Shows every file open, network connect, read/write. High overhead! Use sparingly.

---

## 🔹 8. Advanced Monitoring & Alerting (Questions 971-980)

**Q971: What are "Golden Signals" of monitoring?**
Latency, Traffic, Errors, Saturation. (Google SRE book). Monitor these for every service.

**Q972: What is Prometheus Federation?**
Hierarchical scraping. Global Prometheus scrapes only "aggregated" metrics from Leaf Prometheus servers. Saves storage/bandwidth.

**Q973: What is Thanos?**
Extension for Prometheus. Unlimited storage (Obj Store). Global query view (Query component). HA (Deduplication).

**Q974: What is Cortex/Mimir?**
Scalable, multi-tenant, long-term storage for Prometheus metrics. Push-based architecture.

**Q975: How do you alert on "Burn Rate"?**
Error budget. "If we burn 2% of our monthly error budget in 1 hour -> Critical Alert".

**Q976: What is "recording rule" in Prometheus?**
Pre-compute expensive queries (`sum(rate(...))`) and save as new timeseries. Speeds up Dashboards.

**Q977: How do you monitor certificate expiry?**
Blackbox exporter (probing https). Or `cert-manager` metrics (`certmanager_certificate_expiration_timestamp_seconds`).

**Q978: How do you reduce metric cardinality?**
Drop high-cardinality labels (IPs, user_ids, request_ids). Use histogram buckets wisely.

**Q979: What is Opni?**
AIOps tool for K8s. Uses Log/Metric anomaly detection to suppress noise and find root causes.

**Q980: How do you monitor etcd performance?**
Metrics: `etcd_disk_wal_fsync_duration_seconds` (Must be <10ms). `etcd_server_has_leader`.

---

## 🔹 9. Future Trends (Mainframe modernization, WASM) (Questions 981-990)

**Q981: How is Kubernetes modernizing Mainframes?**
"Lift and Shift" legacy apps to containers. Or managing Mainframe resources via K8s connectors (IBM z/OS).

**Q982: What is the trend of "Cluster Sprawl"?**
Running many small clusters vs few large ones. Driven by isolation needs and managed K8s ease (EKS).

**Q983: What is "AI-Ops" in K8s?**
Using ML to detect anomalies, predict scaling needs, and auto-tune resource requests (Goldilocks/VPA).

**Q984: What is the "Gateway API" significance?**
Successor to Ingress. Role-oriented. Standardizes API Gateways/Service Mesh configuration.

**Q985: What is "Sustainability" in K8s trends?**
Carbon-aware scheduling. Scaling to zero. ARM adoption. Measuring energy strictly.

**Q986: What is the role of eBPF in future security?**
Runtime enforcement without sidecars. Deep visibility into syscalls. Replacing traditional AV agents.

**Q987: How is Kubernetes moving to the Edge?**
K3s, MicroK8s, KubeEdge. Managing satellites, factories, retail stores with limited connectivity.

**Q988: What is "Platform Engineering"?**
Shift from "DevOps" (Everyone does ops) to "Platform Team" building abstractions (IDP) for developers.

**Q989: What is "No-Code/Low-Code" for K8s?**
Tools like Portainer or Yacht. Visual management reducing CLI need for non-experts.

**Q990: What is the future of "Serverless Containers"?**
Abstraction of the "Node". Virtual Kubelet. Fargate. You just pay for vCPU/RAM per Pod.

---

## 🔹 10. Community & Contribution (Questions 991-1000)

**Q991: What is a SIG in Kubernetes?**
Special Interest Group. Group of contributors focused on specific area (SIG-Network, SIG-Node, SIG-Auth).

**Q992: What is a KEP (Kubernetes Enhancement Proposal)?**
Design document required for any significant change/feature. Needs approval from SIG leads + Release team.

**Q993: How does the Kubernetes release cycle work?**
3 releases per year (approx every 4 months). Alpha -> Beta -> Stable.

**Q994: What is a "Triage" role?**
Classifying GitHub issues. "Is it a bug? Feature? Support request?". Assigning priorities.

**Q995: How can you contribute to Kubernetes?**
Code (Go). Docs (Markdown). Testing (E2E). Triage issues. Blog posts.

**Q996: What is the Code of Conduct?**
Rules ensuring respectful, inclusive behavior in the community. Enforced by Code of Conduct Committee.

**Q997: What is KubeCon?**
Flagship conference. CNCF hosted. Networking, technical talks, vendor showcase.

**Q998: Who maintains the Kubernetes documentation?**
SIG-Docs. Open source. Anyone can submit PRs to `kubernetes/website`.

**Q999: What is the "Lighthouse" project?**
No, referring to Prow. Prow is the CI/CD system based on K8s that builds K8s (ChatOps `/test all`).

**Q1000: Why has Kubernetes won the orchestrator war?**
Google heritage (Borg). Open Governance (CNCF). Massive Ecosystem. Portability. Flexibility (CRDs).







