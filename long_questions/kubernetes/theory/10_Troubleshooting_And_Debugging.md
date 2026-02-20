# 🔴 Troubleshooting & Debugging

---

### 1. How do you troubleshoot a CrashLoopBackOff error?

"CrashLoopBackOff means a container is **crashing repeatedly** and K8s is applying exponential backoff before restarting it.

My debugging sequence:
1. `kubectl logs <pod> --previous` — see the last crash's output before the container died
2. `kubectl describe pod <pod>` — check events and exit code
3. `kubectl exec -it <pod> -- sh` (if I can get in before it crashes) — check the app itself
4. Check liveness probe configuration — an overly aggressive liveness probe is a common false killer

Most common causes: app crashes on startup (missing env var, bad config), OOMKilled (hit memory limit), liveness probe too aggressive."

#### In Depth
The backoff sequence: 10s → 20s → 40s → 80s → 160s → 300s (max). If the exit code is 137 (`SIGKILL`, memory limit exceeded) or 139 (`SIGSEGV`), that narrows the cause immediately. Exit code 1 is generic app failure — you need the logs. For containers that exit too fast to exec into, use `kubectl debug` to create a debug container that stays alive:
`kubectl debug -it <pod> --image=busybox --copy-to=debug-pod`

---

### 2. What command do you use to view Pod logs?

"`kubectl logs <pod-name>` for current container output.

`kubectl logs <pod-name> --previous` for the last terminated container's logs — essential for CrashLoopBackOff debugging.

`kubectl logs <pod-name> -c <container-name>` for multi-container pods.

`kubectl logs <pod-name> -f` to follow/stream live.

`kubectl logs -l app=my-app --all-containers=true` for aggregate logs across all pods with that label."

#### In Depth
Kubernetes stores container logs on the node (in `/var/log/pods/`). By default, only the current and previous container's logs are available. Set up a centralized logging system (Loki, Elasticsearch) to retain logs beyond pod lifetime. In cluster, `stern` is a powerful CLI tool for following logs from multiple pods matching a regex with color-coding. `kubetail` is another popular option.

---

### 3. How do you debug a Kubernetes node issue?

"First, check overall node health: `kubectl get nodes` — look for `NotReady` or issues.

Then diagnose: `kubectl describe node <node-name>` — look at Conditions (DiskPressure, MemoryPressure, PIDPressure), events, and `Allocatable` vs `Requested` resources.

For a `NotReady` node, SSH into it and check: `systemctl status kubelet`, `journalctl -u kubelet -n 100`, disk space (`df -h`), and memory (`free -m`). Usually it's a kubelet crash, disk full, or kernel issue."

#### In Depth
The node problem detector (NPD) DaemonSet automatically exposes node-level issues (bad disk I/O, kernel panics, NTP failure) as node conditions that Prometheus can alert on. For draining a problematic node: `kubectl drain <node> --ignore-daemonsets --delete-emptydir-data` — this safely evicts pods and marks the node `SchedulingDisabled`. After fixing the issue, `kubectl uncordon <node>` to allow scheduling again.

---

### 4. What are common reasons for Pending Pods?

"A pod stays in `Pending` when the scheduler **can't find a suitable node**.

Most common reasons:
1. **Insufficient resources**: No node has enough CPU/memory — check `kubectl describe pod` events for 'Insufficient cpu'.
2. **Node affinity/taint mismatch**: Pod requires a label or toleration no available node provides.
3. **PVC not bound**: The referenced PVC doesn't exist or no matching PV. 
4. **Node selector**: `nodeSelector` points to a label that no node has.
5. **Topology spread constraints**: Constraints can't be satisfied.

Running `kubectl describe pod <pod>` and reading the `Events` section tells you exactly which filter rejected it."

#### In Depth
For resource-based Pending, check the gap: `kubectl describe nodes | grep -A3 Allocatable` vs `kubectl describe nodes | grep -A3 Requests`. The Cluster Autoscaler (if installed) should fix Pending-due-to-resources by adding nodes. If CA is running but the pod is still Pending, check CA logs — the pod might have an unsatisfiable constraint that prevents CA from finding any node group that would work.

---

### 5. How to check resource usage per Pod?

"`kubectl top pods` shows CPU and memory usage for all pods in the current namespace.

`kubectl top pods -A` for all namespaces.

`kubectl top pods --sort-by=memory` to find memory hogs.

`kubectl top nodes` for node-level resource consumption.

These use the Metrics Server — if it's not installed, top commands fail. In production, I prefer Grafana with kube-state-metrics + node-exporter for historical trends, not just current snapshots."

#### In Depth
`kubectl top` shows the **current instant** usage, not historical. For right-sizing, look at P95/P99 over 7-14 days in Grafana. The VPA recommendations tool is excellent for this — it analyzes historical usage and recommends requests/limits. Never set requests=limits (it makes the pod Guaranteed QoS class but prevents flexibility) or limits extremely high (it wastes allocatable space and misleads the scheduler).

---

### 6. How do you debug a hung or stuck pod?

"A pod that's not terminating is usually caused by:

1. **Finalizers**: Something added a finalizer and the controller managing it is broken. Fix: manually edit the pod and remove the finalizer (`kubectl patch pod <pod> -p '{"metadata":{"finalizers":[]}}'`) — with caution.
2. **Graceful shutdown timeout**: Container is taking longer than `terminationGracePeriodSeconds`. Check what the container is doing.
3. **Zombie process**: The main process exited but a child process survived and holds the PID namespace.

For force-deletion: `kubectl delete pod <pod> --grace-period=0 --force` — the pod object is deleted from etcd immediately, but the actual container may still be running on the node."

#### In Depth
The `--force` flag tells the API server to delete the pod object without waiting for the kubelet to confirm teardown. This can lead to **ghost pods** — the container is still running on the node but K8s doesn't know about it. Always follow up by SSHing into the node and confirming the container is gone: `crictl ps | grep <pod-name>`. In multi-master HA clusters, a pod stuck terminating after a node failure is often just waiting for the node to be marked as failed and evicted.

---

### 7. What does `kubectl describe pod` show?

"`kubectl describe pod` is my **first diagnostic command** for any issue.

Key sections:
- **Events**: Recent events from schedulers, kubelet, controllers — error messages, image pull failures, scheduling failures.
- **Conditions**: PodScheduled, Initialized, ContainersReady, Ready — which phase the pod is stuck in.
- **Container statuses**: Last exit code, restart count, current state (Running/Waiting/Terminated).
- **Node**: Which node the pod landed on.
- **Volumes**: Whether PVCs are mounted.
- **Resource requests/limits**: What the pod asked for."

#### In Depth
The **Events section** is pure gold. Kubernetes emits extremely specific events:
- `FailedScheduling: 0/3 nodes available: 3 Insufficient cpu.` — exactly which constraint failed
- `Failed to pull image "docker.io/foo:latest": rpc error...` — image pull issue with details
- `Readiness probe failed: Get "http://10.0.1.2:8080/health": dial tcp: connect: connection refused` — the actual HTTP response

Events are stored in the `kube-system` or target namespace and expire after ~1 hour by default. For persistent event storage, use tools like Eventrouter.

---

### 8. What does `ImagePullBackOff` mean?

"`ImagePullBackOff` means K8s is failing to pull the container image and applying exponential backoff before retrying.

Causes:
1. **Image doesn't exist**: Wrong tag, wrong registry path, image deleted.
2. **Auth failure**: imagePullSecret not configured or expired. Check: `kubectl describe pod` → `Events: Failed to pull image: unauthorized`.
3. **Registry unreachable**: Network connectivity issue from the node to the registry.
4. **Rate limiting**: Docker Hub's free tier rate limits unauthenticated pulls (100 pulls/hour per IP).

Fix: verify the image name/tag, check imagePullSecrets, and test from the node: `crictl pull <image>`."

#### In Depth
`imagePullSecrets` at the pod level or on the ServiceAccount. For private registries in production, I create a Secret of type `kubernetes.io/dockerconfigjson` and attach it to the `default` ServiceAccount in the namespace — then all pods in that namespace can pull automatically without specifying `imagePullSecrets` in each spec. Use `imagePullPolicy: IfNotPresent` (default for tagged images) in production — not `Always`, which adds latency and registry traffic.

---

### 9. How do you trace network issues in Kubernetes?

"Systematic approach:

1. **DNS resolution**: `kubectl exec <pod> -- nslookup <service-name>` — verify the service name resolves.
2. **Service reachability**: `kubectl exec <pod> -- curl http://<service>:<port>` — can the pod reach the service?
3. **Direct pod access**: `kubectl exec <pod> -- curl http://<pod-ip>:<port>` — bypass service routing, test direct connectivity.
4. **Check endpoints**: `kubectl get endpoints <service>` — are pods listed? Empty endpoints = no ready pods.
5. **NetworkPolicy**: Is there a policy blocking the traffic?

Tool: `kubectl run debug --rm -it --image=nicolaka/netshoot -- bash` gives you a debug pod with tcpdump, curl, dig, traceroute, nmap."

#### In Depth
For advanced network debugging: `kubectl exec <pod> -- ss -tlnp` to see what the container is actually listening on (common issue: app binds to 127.0.0.1 instead of 0.0.0.0). `tcpdump` inside a container (requires `hostNetwork: true` or a debug container) to see actual packet flows. Cilium's Hubble provides L7-aware network observability — you can see which flows are dropped by NetworkPolicy without running any debug pods.

---

### 10. How do you handle OOMKilled containers?

"OOMKilled (exit code 137) means the container was killed by the Linux OOM killer because it exceeded its **memory limit**.

Fixes:
1. **Increase the memory limit**: The quick fix. Monitor actual usage and set to P99 + 20% buffer.
2. **Fix the memory leak**: Profile the app (Go pprof, Java heap dump) and fix the root cause.
3. **Set appropriate JVM heap**: Java apps often need `-Xmx` set to < (container limit - 200m for overhead).
4. **Use VPA**: Let Vertical Pod Autoscaler recommend the right limits based on historical usage.

Alert on OOMKills proactively: `kube_pod_container_status_last_terminated_reason == oomkilled`."

#### In Depth
The Linux kernel's OOM killer doesn't respect container boundaries — it kills a process in the container's cgroup when that cgroup's memory limit is reached. Setting `requests.memory = limits.memory` (Guaranteed QoS) makes the process the LAST to be killed during node pressure (lower OOM score). Setting requests < limits (Burstable QoS) makes it more likely to be killed when the node is under pressure. Know your app's memory profile and set realistic limits.

---

### 11. How do you debug a stuck terminating pod?

"A pod stuck in `Terminating` for more than 30 seconds is usually a **finalizer issue** or a **node communication failure**.

Debug steps:
1. `kubectl describe pod <pod>` → check `Finalizers:` in the metadata section. If there are finalizers, a controller that manages them might be down.
2. Check if the node is reachable: `kubectl get node <node>`. If `NotReady`, the pod will stay terminating until the node comes back or the `nodeEvictionTimeout` expires.
3. Check kubelet logs on the node: `journalctl -u kubelet | tail -50`

Force delete: `kubectl delete pod <pod> --grace-period=0 --force` — use this only as a last resort, and verify the container is actually gone on the node."

#### In Depth
The **finalizer pattern**: finalizers are strings in `metadata.finalizers`. Kubernetes waits to delete the object until all finalizers are removed. The controller that manages each finalizer is responsible for removing it after cleanup. If the controller is dead, the pod will never progress. Safe resolution: fix the controller. Nuclear option: patch the pod to remove the finalizers manually, which lets K8s delete it but skips the cleanup logic.

---

### 12. How do you diagnose API server slowness?

"Symptoms: slow `kubectl` commands, controllers lagging, dashboard timing out.

Diagnosis:
1. **Metrics**: Check `apiserver_request_duration_seconds` in Prometheus. Identify slow verb/resource combinations.
2. **etcd health**: `etcdctl endpoint health`. API server slowness often originates from slow etcd responses. Check etcd disk I/O.
3. **etcd size**: Large etcd databases slow down list operations. `etcdctl endpoint status --write-out=table` shows DB size. If > 2GB, run defrag.
4. **Too many watch connections**: Large clusters with many controllers can overwhelm the API server. Check `apiserver_current_inflight_requests`.
5. **Node pressure**: Is the node running the API server under CPU/memory pressure?"

#### In Depth
The `--watch-cache-enabled=true` (default) caches resource state in the API server — this is important for performance. Pagination (using `--chunk-size`) in list operations prevents memory spikes. For multi-cluster API server scaling, run behind an L4 load balancer with multiple replicas. The API server's `--max-requests-inflight` and `--max-mutating-requests-inflight` can be tuned — but increasing them without addressing root causes just delays the problem.

---

### 13. What causes long scheduling delays?

"Scheduling delay (time from pod creation to it being bound to a node) is usually caused by:

1. **Scheduler overload**: Too many pods in Pending state — the scheduler queue is backed up.
2. **Complex predicates**: Affinity/anti-affinity rules require the scheduler to evaluate many pods, slowing scoring.
3. **Unschedulable pods**: Thousands of pods with unsatisfiable constraints fill the scheduler queue, delaying schedulable pods.
4. **Unready nodes**: All matching nodes are `NotReady` or `cordoned`.
5. **API server latency**: Scheduler reads from API server — if it's slow, scheduling is slow.

Monitor: `scheduler_scheduling_latency_seconds` in Prometheus."

#### In Depth
The scheduler's cache accelerates filtering by maintaining an in-memory copy of node and pod state. Scheduling plugins can be slow — complex plugins with many pods take O(nodes * pods) time per scheduling cycle. In very large clusters (5000+ nodes), using **topology-aware volume provisioning** (`WaitForFirstConsumer`) can delay scheduling since the PV must be provisioned first. Set `--parallelism` on the scheduler (default 16) to allow parallel pod evaluations.

---

### 14. What are production anti-patterns in Kubernetes?

"From real experience, the most common failure patterns:

1. **No resource requests/limits**: Pods get placed on already-overloaded nodes or get OOMKilled without proper sizing.
2. **No liveness/readiness probes**: Traffic routed to broken containers, failed containers not restarted.
3. **Single replica deployments**: Any pod restart causes downtime.
4. **No PodDisruptionBudget**: Node maintenance takes down the entire service.
5. **Using latest image tag**: Unpredictable deployments — `latest` changes silently.
6. **Secrets in environment variables from ConfigMaps**: Sensitive data logged accidentally.
7. **No resource quotas in namespaces**: One team's job can starve another team's production service.
8. **Running as root**: Massive security risk for the entire cluster.

These 8 anti-patterns alone account for the majority of K8s production incidents I've seen."

#### In Depth
The **Production Readiness Checklist** I follow for every service: (1) Multi-replica with PDB, (2) Readiness/liveness probes configured, (3) Resource requests set based on profiling, (4) ImagePullPolicy: IfNotPresent, (5) Specific image tags (not latest), (6) Non-root security context, (7) Network policy restricting ingress/egress, (8) PodAntiAffinity across zones, (9) Graceful shutdown handler (SIGTERM → drain → exit), (10) Centralized log shipping and Prometheus metrics exposed.

---
