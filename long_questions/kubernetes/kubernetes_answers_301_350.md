## 🔹 Pods & Containers – Advanced (Questions 301-310)

### Question 301: What’s the difference between a container restart and a pod restart?

**Answer:**
- **Container Restart:** The pod stays (IP, ID same). Only the container process restarts (e.g., PID 1 crashes, Kubelet Relaunches). Usage count increases.
- **Pod Restart:** Usually means "Recreation". The old pod is deleted. A new pod with new IP and new ID is scheduled.

---

### Question 302: How do you handle init containers in Kubernetes?

**Answer:**
Defined in `spec.initContainers`.
- They run sequentially (one by one).
- All must succeed before the main app starts.
- If one fails, the Pod restarts (depending on policy).
- **Note:** Probes (Liveness/Readiness) do NOT apply to init containers.

---

### Question 303: What are the use cases of sidecar containers?

**Answer:**
1.  **Log Forwarding:** Read logs from shared volume, send to Splunk.
2.  **Proxy:** Act as local proxy (Service Mesh Envoy, Cloud SQL Proxy).
3.  **Config Watcher:** Pull new config and signal main app to reload.

---

### Question 304: How do you detect a memory leak in a container?

**Answer:**
monitor the `container_memory_usage_bytes` metric over time.
- If the "Sawtooth pattern" (Growth -> OOMKill -> Drop -> Growth) appears.
- Use `kubectl top pod` to see current usage vs limit.

---

### Question 305: What are the pros/cons of running multiple containers in one pod?

**Answer:**
- **Pros:** Tight coupling (Share localhost/volumes/IPC), Single deployment unit.
- **Cons:** Shared fate (If one crashes pod *might* be unstable), Resource contention, Complex resource calculation.
- **Best Practice:** One container per pod, unless tightly coupled helper pattern.

---

### Question 306: How do you share files between containers in the same pod?

**Answer:**
Use an `emptyDir` volume.
```yaml
volumes:
- name: shared-data
  emptyDir: {}
```
Mount `shared-data` to both containers. Container A writes, Container B reads.

---

### Question 307: Can you run privileged containers in Kubernetes?

**Answer:**
Yes, `securityContext: privileged: true`.
- Giving the container access to **all devices** on the host.
- **Risk:** High (Root on host). Avoid unless necessary (e.g., storage drivers, CNI).

---

### Question 308: How do you mount a hostPath volume?

**Answer:**
```yaml
volumes:
- name: test-volume
  hostPath:
    path: /data
    type: Directory
```
- **Warning:** Pod is now tied to that specific node's filesystem.

---

### Question 309: What is the lifecycle hook `preStop` used for?

**Answer:**
Executes before the container terminates.
- **Use Case:** Graceful shutdown.
  - `preStop: exec: command: ["/bin/sh", "-c", "sleep 10"]`.
  - Gives the LoadBalancer time to spot the "Terminating" state and remove the IP before the app actually stops listening.

---

### Question 310: How do you test pod startup and shutdown behavior?

**Answer:**
1.  **Shutdown:** Run `kubectl delete pod` and capture logs simultaneously to ensure it catches SIGTERM.
2.  **Startup:** Define `startupProbe` to delay liveness checks if app takes 60s to boot.

---

## 🔹 Deployments & Rollouts (Questions 311-320)

### Question 311: What is a canary deployment strategy?

**Answer:**
Releasing a change to a small subset (e.g., 5%) of users.
- Verify metrics (Error rate).
- If good, roll out to 100%.
- Kubernetes does not support this natively via Deployment object (Deployments handle RollingUpdate). Requires **Ingress/Service Mesh** or tools like **Argo Rollouts**.

---

### Question 312: How do you pause a deployment in Kubernetes?

**Answer:**
`kubectl rollout pause deployment/my-dep`.
- Freezes the state.
- You can make multiple edit (cpu, image, labels).
- `kubectl rollout resume` triggers the actual single update.

---

### Question 313: How do you manually trigger a rollout?

**Answer:**
If no config changed but you want to restart pods (e.g., to pick up new External Secret or clear cache):
`kubectl rollout restart deployment/my-dep`.

---

### Question 314: What is a surge vs unavailable setting in rolling updates?

**Answer:**
- **maxSurge:** How many extra pods can we create? (e.g., 25% -> 125% capacity).
- **maxUnavailable:** How many pods can be deleted? (e.g., 25% -> 75% capacity).
- Tune these for speed vs stability.

---

### Question 315: What happens when you update a deployment manifest?

**Answer:**
1.  Deployment Controller notices change (Hash of PodTemplate changes).
2.  Creates NEW ReplicaSet (Scale 0 -> 1).
3.  Scales OLD ReplicaSet (10 -> 9).
4.  Repeat until New=10, Old=0.

---

### Question 316: How do you perform A/B testing in Kubernetes?

**Answer:**
Similar to Canary, but routing is based on **User Feature Flags** or Headers, not random %.
- Requires an Ingress Controller or Service Mesh that can route `Header: X-Test-User` to Service B.

---

### Question 317: How does rollout history differ from rollout status?

**Answer:**
- **History:** List of previous ReplicaSets (Revisions).
- **Status:** Current progress (e.g., "2 of 4 updated replicas are available").

---

### Question 318: How can you automate rollback on failure?

**Answer:**
Native Kubernetes Deployments do **not** auto-rollback. They hang.
- **Helm:** `helm upgrade --atomic` (Rolls back if fails).
- **Flagger/Argo:** advanced metrics-based rollback.
- **Script:** Watch status, if timeout -> rollback.

---

### Question 319: How do you check the reason for deployment failure?

**Answer:**
1.  `kubectl rollout status deployment/my-dep` (Will hang/error).
2.  `kubectl get rs` (See if new RS is stuck at 0 ready).
3.  `kubectl describe pod <new-pod-hash>` (Readiness probe failing?).

---

### Question 320: What is progressive delivery?

**Answer:**
An umbrella term for advanced deployment patterns (Canary, Blue/Green, A/B) + Automation.
- Moving away from "Big Bang" deployments.

---

## 🔹 Namespaces & Isolation (Questions 321-330)

### Question 321: How do namespaces help with resource isolation?

**Answer:**
1.  **Visibility:** "Get Pods" only shows your team's pods.
2.  **ResourceQuotas:** Limit CPU/RAM per namespace.
3.  **Name collision:** Both `Dev` and `Prod` can have a service named `db` without conflict.

---

### Question 322: What is the default namespace?

**Answer:**
`default`.
If you run `kubectl get pods` without `-n`, it looks here.
Ideally, don't use it for production. Create explicit namespaces.

---

### Question 323: How can you apply policies per namespace?

**Answer:**
- **LimitRange:** Default CPU/Mem requests for new pods in NS.
- **ResourceQuota:** Hard ceiling for NS.
- **NetworkPolicy:** Firewall rules for NS.
- **RBAC:** RoleBindings for that NS.

---

### Question 324: How can you restrict access to a namespace?

**Answer:**
Create a **Role** (Allow read/write) and a **RoleBinding** linking User `Bob` to `Namespace A`.
Bob cannot see Namespace B.

---

### Question 325: How do you list all resources in a namespace?

**Answer:**
`kubectl get all -n my-ns`.
*Note: `get all` is misleading, it doesn't show ConfigMaps/Secrets/Ingresses, only workloads/services.*

---

### Question 326: What is the `kube-system` namespace used for?

**Answer:**
Holds the "System" components created by Kubernetes system.
- kube-dns, kube-proxy, metrics-server, etcd (sometimes).
- **Do not modify** unless you know what you are doing.

---

### Question 327: What is a namespace quota?

**Answer:**
Defined by `ResourceQuota` object.
```yaml
spec:
  hard:
    pods: "10"
    secrets: "5"
```
Prevents creation if limit exceeded.

---

### Question 328: How do you clean up all resources in a namespace?

**Answer:**
`kubectl delete namespace my-ns`.
- This deletes the namespace AND everything inside it (Cascading delete).

---

### Question 329: Can you have the same pod name in different namespaces?

**Answer:**
Yes.
Namespace provides a scope for names.
`dev/nginx` and `prod/nginx` are distinct objects.

---

### Question 330: How does RBAC differ when applied cluster-wide vs namespace-wide?

**Answer:**
- **RoleBinding:** Grants permissions in **specific namespace**.
- **ClusterRoleBinding:** Grants permissions in **all namespaces** (entire cluster).

---

## 🔹 Cluster Configuration & Architecture (Questions 331-340)

### Question 331: What is the role of the Kubernetes control plane?

**Answer:**
To maintain the **Desired State** of the cluster.
- Decides where to run pods.
- Detects failures.
- Responds to API requests.

---

### Question 332: What is the kubelet responsible for?

**Answer:**
The "Captain" of the Node.
- Registers node with API server.
- Watches PodSpecs assigned to node.
- Tells Docker/Containerd: "Run this image".
- Reports health back to Master.

---

### Question 333: What is the purpose of the kube-proxy?

**Answer:**
Handles Service Networking.
- Programs iptables/IPVS on the node.
- Ensures a request to `ClusterIP` gets routed to one of the backing Pod IPs.

---

### Question 334: How does Kubernetes maintain desired state?

**Answer:**
Through **Control Loops** (Reconciliation Loops).
- Current State: What is running.
- Desired State: YAML in etcd.
- Controller calculates Diff and applies patches (Create/Delete).

---

### Question 335: What are admission controllers?

**Answer:**
Plugins that intercept requests to the API server **after** Authentication/Authorization but **before** persistence.
- **Example:** AlwaysPullImages, LimitRanger, MutatingWebhook (Inject sidecar).

---

### Question 336: What are validating vs mutating admission webhooks?

**Answer:**
- **Mutating:** Touches request *before* validation. Can modify it. (e.g., set default requests if missing).
- **Validating:** Checks final request. Can reject it. (e.g., Reject if "privileged: true").

---

### Question 337: What is the function of the API server?

**Answer:**
The only component that talks to etcd.
- It validates data.
- It serves the REST API.
- It is the central coordination point.

---

### Question 338: How does controller-manager differ from scheduler?

**Answer:**
- **Scheduler:** Focuses on ONE thing: Assigning `nodeName` to pending pods.
- **Controller-Manager:** Focuses on EVERYTHING ELSE: Replica counts, Endpoints, Token creation, Service Accounts.

---

### Question 339: What is a lease in Kubernetes?

**Answer:**
A mechanism for distributed locking (Leader Election).
- Controller Manager High Availability uses Leases to decide which instance is "Leader".
- `kubectl get leases -n kube-system`.

---

### Question 340: How does leader election work in control plane components?

**Answer:**
- Multiple replicas (e.g., 3 Scheduler pods) run.
- They race to acquire/renew a `Lease` object in API server.
- The one holding the lease is Active. Others verify lease timestamp and wait (Passive).

---

## 🔹 Autoscaling (Questions 341-350)

### Question 341: How does the Horizontal Pod Autoscaler (HPA) work?

**Answer:**
Adjusts `replicas` count.
- Source: Metrics Server (CPU/Mem).
- Logic: `Current / Target` ratio.
- Loop: Runs every 15s.

---

### Question 342: What is the Vertical Pod Autoscaler (VPA)?

**Answer:**
Adjusts `client-go` Requests/Limits.
- Used for stateful apps or monoliths that can't scale horizontally.
- Requires restart to apply CPU/Mem changes (mostly).

---

### Question 343: What metrics does HPA use?

**Answer:**
- **Resource Metrics:** CPU, Memory (from Metrics Server).
- **Custom Metrics:** QPS, Queue Depth (from Prometheus Adaptor).
- **External Metrics:** Metrics unrelated to K8s objects (e.g., SQS Queue Size).

---

### Question 344: How does the Cluster Autoscaler work?

**Answer:**
Watches for **Pending Pods** failed due to "Insufficient Resources".
- If true -> Calls Cloud Provider (AWS ASG) -> Adds Node.
- Also scales down empty nodes.

---

### Question 345: What’s the minimum requirement for HPA to function?

**Answer:**
1.  **Metrics Server:** Must be installed.
2.  **Requests:** Pods MUST have resource requests defined (otherwise HPA cannot calculate percentage).

---

### Question 346: How can you autoscale based on custom metrics?

**Answer:**
Install **Prometheus Adapter**.
- It translates Prometheus Query -> Kubernetes Custom Metric API.
- HPA spec: `type: Pods`, `metricName: requests_per_second`.

---

### Question 347: What are the challenges with autoscaling StatefulSets?

**Answer:**
- Scaling down is risky (Risk of data loss if not replicated correctly).
- Sharding/Rebalancing data is application specific (cannot be handled by generic K8s logic).

---

### Question 348: What is KEDA (Kubernetes Event-driven Autoscaler)?

**Answer:**
A CNCF project to simplify event-based scaling.
- Scaler for Kafka, RabbitMQ, Azure ServiceBus, etc.
- It translates external events into HPA metrics automatically.
- Allows scaling to **Zero** (Serverless feel).

---

### Question 349: What is the downside of aggressive autoscaling?

**Answer:**
**Thrashing (Flapping).**
- Scale up -> Load drops -> Scale down -> Load spikes -> Scale up.
- Wastes compute startup time.
- Fix: Use `stabilizationWindowSeconds` (Hysteresis).

---

### Question 350: How do you configure cooldown periods for autoscaling?

**Answer:**
In HPA v2beta2+:
```yaml
behavior:
  scaleDown:
    stabilizationWindowSeconds: 300 # Wait 5m before scaling down
```
Prevents removing pods during temporary dips in traffic.
