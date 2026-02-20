# 🟡 Pods, ReplicaSets & Deployments

---

### 1. What is a ReplicaSet?

"A ReplicaSet ensures that a **specified number of pod replicas** are running at any given time.

If a pod dies or is deleted, the ReplicaSet controller creates a new one. If there are too many, it deletes extras. It's the primitive that gives Kubernetes its self-healing capability for stateless workloads.

That said, I almost never create ReplicaSets directly — Deployments manage them for me and add rollout history on top."

#### In Depth
A ReplicaSet uses a **label selector** to identify which pods it manages. This means if you manually create pods with matching labels, the ReplicaSet will count them toward its desired replica count. This is powerful but can cause confusion — always use Deployments to avoid managing ReplicaSets manually.

---

### 2. How is a ReplicaSet different from a ReplicationController?

"A ReplicationController is the older, deprecated version of ReplicaSet.

The key difference is the **label selector**: ReplicationController uses equality-based selectors (`app=frontend`), while ReplicaSet supports **set-based selectors** (`app in (frontend, web)`), giving you more expressive matching.

In modern clusters, ReplicationControllers are essentially legacy. All new workloads should use Deployments, which internally use ReplicaSets."

#### In Depth
ReplicationControllers were deprecated in Kubernetes 1.11. The API still exists for backward compatibility, but there isn't a functional difference large enough to ever prefer them. The migration path is straightforward: replace RC with Deployment and you gain rollout history, pause/resume, and much richer update strategies.

---

### 3. How do you scale Pods in Kubernetes?

"There are three ways:

1. **Manual scaling**: `kubectl scale deployment my-app --replicas=5` — instant but not automatic.
2. **Horizontal Pod Autoscaler (HPA)**: Automatically scales based on CPU, memory, or custom metrics. My go-to for production.
3. **Vertical Pod Autoscaler (VPA)**: Automatically adjusts resource requests/limits — useful when you're not sure what a container really needs.

I always set up HPA with a reasonable min/max range and tie it to custom business metrics like request rate for better responsiveness than CPU alone."

#### In Depth
When HPA scales down, it respects `minReplicas` and the `PodDisruptionBudget`. Scale-down is deliberately slower than scale-up to avoid thrashing — the default stabilization window is 5 minutes. For stateful workloads, be very careful with HPA; StatefulSets have ordering constraints that HPA doesn't account for.

---

### 4. What is a multi-container Pod?

"A multi-container Pod hosts more than one container that share the same network and storage.

The most common patterns are **sidecar** (aux process alongside the main app — like a log shipper), **ambassador** (proxy handling external communication — like a connection pooler), and **adapter** (transforming output format — like metrics normalization).

I've used the sidecar pattern extensively in service mesh architectures where Envoy/Linkerd sidecars are injected automatically into every pod."

#### In Depth
Containers in a Pod communicate via `localhost` since they share the same network namespace. They also share `emptyDir` volumes for file sharing. However, they have **separate process namespaces** by default — you can't `kill` a process in one container from another unless `shareProcessNamespace: true` is set. Init containers are a special case — they run and complete before the main containers start.

---

### 5. How do Pods communicate with each other?

"Every Pod gets a unique IP address. Pods can communicate **directly by IP** — there's no NAT between pods in the same cluster (that's the K8s networking model promise).

For service-to-service communication, I use **Kubernetes Services** which provide a stable DNS name and virtual IP that round-robins to backing pods.

In production, I enforce NetworkPolicies to restrict which pods can talk to which — the default allow-all is too permissive for sensitive services like databases."

#### In Depth
The Kubernetes networking model requires: (1) pods can communicate with all other pods without NAT, (2) nodes can communicate with all pods without NAT, (3) pod IP is the same from inside and outside the pod. CNI plugins implement this — Calico uses BGP, Flannel uses VXLAN overlay, Cilium uses eBPF. Each has different performance and feature characteristics.

---

### 6. What is a Deployment?

"A Deployment is the standard way to manage **stateless application workloads** in Kubernetes.

It wraps a ReplicaSet and adds rollout management: rolling updates, rollback, pause/resume, and revision history.

When I update a Deployment's pod template (e.g., new image), K8s creates a new ReplicaSet and gradually shifts pods from old to new. If the new pods are unhealthy, I can `kubectl rollout undo` to go back instantly."

#### In Depth
A Deployment keeps a configurable number of old ReplicaSets for rollback (`revisionHistoryLimit`, default 10). The actual rollout behavior is controlled by `strategy.rollingUpdate.maxSurge` (extra pods during update) and `strategy.rollingUpdate.maxUnavailable` (pods that can be down during update). Setting both to 0% is not allowed — one must allow at least some variance.

---

### 7. How do you roll back a Deployment?

"`kubectl rollout undo deployment/my-app` — this is the most common way.

It reverts to the **previous revision**. I can also specify an exact revision: `kubectl rollout undo deployment/my-app --to-revision=3`.

Before rolling back, I always check `kubectl rollout history deployment/my-app` to see what each revision was and `kubectl rollout status` to confirm rollback completion."

#### In Depth
Rollback works by updating the Deployment's spec to match the previous ReplicaSet's pod template. The old ReplicaSet is scaled back up, and the current one is scaled down — using the same rolling update strategy. If you need to rollback from a bad blue-green or canary, the mechanism depends on the tool (Argo Rollouts has separate rollback semantics).

---

### 8. What is a Rolling Update?

"A Rolling Update replaces Pods one (or a few) at a time so the application remains **available throughout the update**.

During a rolling update, the new ReplicaSet scales up while the old one scales down, constrained by `maxSurge` and `maxUnavailable` settings.

I always configure a health check (readiness probe) so that the rollout pauses if the new pods aren't ready. Without that, you can shift traffic to broken pods before they crash."

#### In Depth
Rolling updates are the default Deployment strategy. The alternative is `Recreate` — which terminates all old pods before creating new ones. I use `Recreate` only when the app can't run two versions simultaneously (e.g., DB schema migration that requires exclusive access). It causes downtime, so it's rare in production.

---

### 9. How do you pause and resume Deployments?

"`kubectl rollout pause deployment/my-app` halts a rollout mid-way.

This is useful for **canary-style testing**: push an update, pause the rollout, observe the small percentage of new pods in production, and then resume if metrics look good.

`kubectl rollout resume deployment/my-app` continues from where it left off. This gives you manual control over the rollout pace without needing a full canary deployment framework."

#### In Depth
While a deployment is paused, changes to the PodSpec (e.g., changing the image again) are accumulated but not applied. All changes are batched and applied atomically when you resume. This is useful for making multiple config changes at once without triggering multiple rollouts.

---

### 10. What is a canary deployment in Kubernetes?

"A canary deployment sends a **small percentage of traffic** to a new version before rolling it out to everyone.

In native K8s, you implement this by running two Deployments — `app-stable` with 90 replicas and `app-canary` with 10 — both selected by the same Service. 10% of traffic goes to the canary.

For proper traffic-percentage control (not pod-count based), I use **Argo Rollouts** or **Istio traffic shifting**, which give precise percentage control via VirtualServices."

#### In Depth
The native K8s approach is coarse-grained — traffic split is proportional to pod count. Argo Rollouts with Istio allows you to send exactly 5% of traffic to the canary regardless of replica count. It also integrates with Prometheus for **analysis-based promotion**: automatically promote or rollback based on error rate thresholds.

---

### 11. What is a pod lifecycle?

"A pod goes through these phases: **Pending → Running → Succeeded/Failed**.

- **Pending**: Scheduled but containers not yet started (image pulling, volume mounting)
- **Running**: At least one container is running
- **Succeeded**: All containers exited 0 (typical for Jobs)
- **Failed**: At least one container exited non-zero
- **Unknown**: Node communication lost

In debugging, `CrashLoopBackOff` is not a phase — it's a status condition where containers keep restarting with exponential backoff."

#### In Depth
Beyond phases, pods have `conditions`: `PodScheduled`, `Initialized`, `ContainersReady`, and `Ready`. The `Ready` condition controls whether the pod receives traffic from Services. The transition from `ContainersReady` to `Ready` can be delayed by custom readiness gates — useful for service mesh scenarios where the sidecar needs to be ready first.

---

### 12. What are the phases of a pod?

"Officially, Kubernetes defines 5 pod phases: **Pending, Running, Succeeded, Failed, and Unknown**.

- `Pending`: Pod accepted but not all containers started. Often waiting for image pull or scheduling.
- `Running`: Pod bound to node, all containers created, at least one running.
- `Succeeded`: All containers have exited successfully (exit code 0). Typical for batch jobs.
- `Failed`: All containers have terminated, at least one with non-zero exit.
- `Unknown`: Cannot get pod status (usually node failure or network issue).

I frequently troubleshoot pods stuck in `Pending` — usually due to insufficient resources, missing PVCs, or taint mismatches."

#### In Depth
Phase is stored in `status.phase`. You should not rely on phase for fine-grained state — use `status.conditions` instead. For example, a pod can be `Running` but not `Ready` if its readiness probe is failing, meaning it won't receive Service traffic. This distinction is critical for zero-downtime deployments.

---

### 13. What is a liveness probe vs readiness probe?

"They serve different purposes:

- **Liveness probe**: 'Is this container alive?' If it fails, Kubernetes restarts the container. Used to recover from deadlocks.
- **Readiness probe**: 'Is this container ready to serve traffic?' If it fails, the pod is removed from Service endpoints. Used during startup and rolling updates.

I always configure both separately. A common mistake is using `/healthz` for both — but `/healthz` returning 200 doesn't mean the app is fully initialized and ready to handle requests."

#### In Depth
There's a third probe: **Startup probe**. It delays liveness checks until the app has finished booting — useful for slow-starting applications. Without it, a liveness probe that runs before the app is ready would cause a crash loop. `failureThreshold * periodSeconds` gives you the total startup window.

---

### 14. What triggers a pod to be terminated?

"Pods are terminated for several reasons:

1. **Manual deletion**: `kubectl delete pod`
2. **Node failure** or eviction (due to resource pressure)
3. **Deployment rolling update** (old pods terminated as new ones are ready)
4. **Liveness probe failure** (container restart, not full pod deletion)
5. **Preemption** — a higher-priority pod needs the resources

In production, I watch for unexpected pod restarts in Grafana dashboards and alert on high restart counts — they usually indicate liveness probe misconfiguration or memory leaks."

#### In Depth
The termination sequence is: SIGTERM sent → `terminationGracePeriodSeconds` countdown (default 30s) → SIGKILL if still running. The `preStop` lifecycle hook runs before SIGTERM, useful for draining connection pools or deregistering from a service registry. In Istio environments, you should delay SIGTERM by a few seconds using `preStop` sleep to let the sidecar finish draining.

---

### 15. What happens during pod termination?

"Kubernetes terminates a pod through an orderly shutdown sequence:

1. Pod is set to `Terminating` state and removed from Service endpoints
2. `preStop` hook runs (if defined)
3. SIGTERM is sent to all containers
4. K8s waits `terminationGracePeriodSeconds` (default 30s)
5. If containers haven't exited, SIGKILL is sent forcefully

I always implement SIGTERM handlers in my apps. For Go HTTP servers, I use `http.Server.Shutdown()` which drains in-flight requests before closing."

#### In Depth
The pod removal from Service endpoints and the SIGTERM signal happen **concurrently**, not sequentially. This means traffic can still arrive after SIGTERM is sent. The `preStop` sleep hack (`preStop: exec: command: [sleep, '5']`) gives the endpoint controller time to propagate the removal before the app starts refusing connections.

---

### 16. What does `terminationGracePeriodSeconds` do?

"`terminationGracePeriodSeconds` gives containers a window to **shut down gracefully** after receiving SIGTERM.

The default is 30 seconds. After this period, SIGKILL is sent — no more cleanup.

For apps with long-running requests (like video processing or batch jobs), I increase this value. For fast-startup microservices, I sometimes reduce it to speed up rolling updates. The right value depends on your P99 request duration."

#### In Depth
If `terminationGracePeriodSeconds` is set to 0, the pod is force-killed immediately. This is useful for stuck `Terminating` pods where the kubelet is fighting with a finalizer. You can override it per-deletion: `kubectl delete pod my-pod --grace-period=0 --force` — but use this only for debugging, never in automation.

---
