# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 5)
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, IBM, Capgemini, etc.
> **Focus:** Resource Quotas, Canary/Blue-Green Deployments, `kubectl top` monitoring, and advanced CrashLoopBackOff troubleshooting.

---

## 🔹 Resource Quotas & Namespace Isolation (Questions 56-60)

### Question 56: What is a ResourceQuota and why is it important in a shared cluster?

**Answer:**
A **ResourceQuota** object restricts the total amount of resources that can be consumed in a given Namespace. This is critical in shared or multi-team clusters to prevent one team from accidentally (or intentionally) consuming all cluster resources and starving other teams.

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: team-a-quota
  namespace: team-a
spec:
  hard:
    requests.cpu: "10"          # Team A can never request more than 10 CPU cores total
    requests.memory: 20Gi       # Total memory requests capped at 20Gi
    limits.cpu: "20"
    limits.memory: 40Gi
    pods: "50"                  # Max 50 pods in this namespace
    persistentvolumeclaims: "10"
    services.loadbalancers: "2"
```

When a team tries to create a Pod that would exceed the quota, the API server **rejects the creation immediately** with a `403 Forbidden` response and a clear message about which quota was violated.

---

### Question 57: What is a LimitRange and how does it differ from ResourceQuota?

**Answer:**

| | **ResourceQuota** | **LimitRange** |
|---|---|---|
| **Scope** | Entire Namespace total | Per individual Pod/Container |
| **Purpose** | Caps total namespace consumption | Sets default and max values per container |
| **Effect** | Rejects resources that breach total budget | Injects defaults; rejects over-sized containers |

**LimitRange example** — if a developer forgets to set resource limits, this auto-injects them:
```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: default-limits
  namespace: team-a
spec:
  limits:
    - type: Container
      default:
        cpu: "500m"
        memory: 512Mi
      defaultRequest:
        cpu: "200m"
        memory: 256Mi
      max:
        cpu: "2"                # No single container can ask for > 2 CPU
        memory: 2Gi
```

**Key rule:** If a Namespace has a **ResourceQuota** set on CPU/memory, every Pod in that namespace **must** explicitly define `requests` and `limits`. Otherwise, the API server rejects the pod. This is where LimitRange helps by injecting defaults automatically.

---

### Question 58: How do you monitor resource usage in a Kubernetes cluster from the command-line?

**Answer:**
The primary tool for real-time resource consumption is **`kubectl top`**, powered by the `metrics-server` add-on.

**Node-level usage:**
```bash
kubectl top nodes
# NAME          CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
# node-1        350m         17%    2048Mi          64%
# node-2        980m         49%    3512Mi          87%   ← saturation
```

**Pod-level usage:**
```bash
kubectl top pods -n production --sort-by=memory
# NAME                    CPU(cores)   MEMORY(bytes)
# payment-api-7d9f-xz2p   120m         890Mi
# user-service-6b4-pq91    45m          210Mi
```

**Other important monitoring commands:**
```bash
# Check if nodes are under pressure
kubectl describe node node-2 | grep -A5 Conditions

# Check resource quota consumption in a namespace
kubectl describe resourcequota -n team-a

# See running vs requested replicas across all deployments
kubectl get deployment -A

# Watch pod restarts in real time
kubectl get pods -n production -w
```

---

### Question 59: A deployment is consuming far more memory than expected. Walk through your investigation process.

**Answer:**
Step-by-step memory investigation:

1. **Identify the offending pod:**
```bash
kubectl top pods -n production --sort-by=memory
```

2. **Check if it's nearing its limit (about to be OOMKilled):**
```bash
kubectl describe pod <pod-name> -n production
# Look for: Limits, Requests, and the "Last State" section for OOMKilled reasons
```

3. **Check recent OOMKill history:**
```bash
kubectl get pod <pod-name> -o jsonpath='{.status.containerStatuses[0].lastState}'
# Will show: reason: OOMKilled, exitCode: 137
```

4. **Check actual application logs for memory-related errors:**
```bash
kubectl logs <pod-name> --previous   # logs from the crashed instance
```

5. **Decision:**
   - If legitimate spike: Increase `limits.memory` and/or optimize the application.
   - If leak: Enable profiling endpoints (e.g., Go's `pprof`, Java's heap dump) and capture a heap dump.
   - If wrong limit set: Adjust `requests` and `limits` to match `kubectl top` baseline data.

---

### Question 60: Explain Horizontal vs Vertical Pod Autoscaling. Which one should you choose for a Java Spring Boot API?

**Answer:**

| | **HPA (Horizontal)** | **VPA (Vertical)** |
|---|---|---|
| **What it scales** | Number of Pod replicas | CPU/Memory Requests & Limits on an existing pod |
| **Restart required** | No | Yes (VPA restarts pod to apply new resources) |
| **Best for** | Stateless APIs (web, REST, gRPC) | Stateful apps, batch jobs, single-instance workloads |
| **Conflict risk** | Safe alone | Do NOT use VPA with HPA on same CPU metric |

**For a Java Spring Boot API → Use HPA:**
- Spring Boot apps are typically stateless behind a load balancer.
- HPA scales replicas in response to traffic spikes with zero restarts.
- Configure it based on CPU or custom metrics (e.g., requests-per-second from Ingress):

```bash
kubectl autoscale deployment spring-boot-api \
  --cpu-percent=60 \
  --min=2 \
  --max=20
```

**VPA is better suited for:** Background Java batch workers (single-instance jobs) where you need to right-size JVM heap memory automatically based on actual usage patterns over time.

---

## 🔹 Canary & Blue-Green Deployments (Questions 61-65)

### Question 61: What is a Blue-Green deployment strategy and how do you implement it in Kubernetes?

**Answer:**
In a **Blue-Green deployment**, you maintain two complete, identical environments:
- **Blue** = currently serving 100% of production traffic (old version).
- **Green** = the new version deployed and tested alongside, but receiving no traffic yet.

When you're confident in Green, you switch ALL traffic from Blue → Green instantly by patching the Service selector.

**Implementation:**
```yaml
# Blue Deployment (currently live)
metadata:
  name: payment-api-blue
  labels:
    app: payment-api
    version: blue

# Green Deployment (new version, ready and tested)
metadata:
  name: payment-api-green
  labels:
    app: payment-api
    version: green
```

```yaml
# Service - currently pointing to BLUE
spec:
  selector:
    app: payment-api
    version: blue       # ← Change this to "green" to switch ALL traffic
```

**Switch command:**
```bash
kubectl patch service payment-api -p '{"spec":{"selector":{"version":"green"}}}'
```

**Advantage:** Instant switch-over. Rollback = re-patch the selector back to `blue` (old pods are still running).
**Disadvantage:** Requires **2x the resources** to run both environments simultaneously.

---

### Question 62: What is the difference between Blue-Green and Canary deployments? Which is safer?

**Answer:**

| | **Blue-Green** | **Canary** |
|---|---|---|
| **Traffic switch** | Instant 0% → 100% | Gradual (e.g., 5% → 25% → 100%) |
| **Risk** | If green has a bug, 100% of users are affected immediately | Only a small % of users see the bug |
| **Rollback speed** | Instant (patch selector back) | Progressively scale down canary |
| **Resource cost** | 2x (both environments fully up) | ~1.1x (1 extra canary pod) |
| **Best for** | Schema migrations, major version upgrades | Feature flag testing, A/B testing |

**Canary is generally safer** for user-facing features because issues are caught with minimal blast radius.

---

### Question 63: You deployed a new version via rolling update but noticed a spike in 5xx errors. How do you rollback immediately?

**Answer:**
```bash
# 1. Immediately rollback to the previous stable revision
kubectl rollout undo deployment/payment-api

# 2. Verify rollback is progressing
kubectl rollout status deployment/payment-api

# 3. Confirm new pods are from the previous version
kubectl describe deployment payment-api | grep Image

# 4. Check revision history to rollback to a specific older version
kubectl rollout history deployment/payment-api
# REVISION  CHANGE-CAUSE
# 1         Initial deploy
# 2         v1.5 feature release
# 3         v1.6 (BAD - currently rolling back from this)

kubectl rollout undo deployment/payment-api --to-revision=2
```

**Pro tip:** Always annotate your deployments with a change cause for readable history:
```bash
kubectl annotate deployment payment-api kubernetes.io/change-cause="v1.6 feature X release"
```

---

### Question 64: How do you implement a canary release without Argo Rollouts using only native Kubernetes?

**Answer:**
Using **two Deployments + one Service** (the native K8s canary pattern):

**Step 1 — Stable (v1): 9 replicas**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-api-stable
spec:
  replicas: 9
  selector:
    matchLabels:
      app: payment-api
      track: stable
  template:
    metadata:
      labels:
        app: payment-api     # Service selects on this
        track: stable
    spec:
      containers:
        - name: api
          image: payment-api:v1.0
```

**Step 2 — Canary (v2): 1 replica**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-api-canary
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: payment-api     # Same label — Service routes to this too!
        track: canary
    spec:
      containers:
        - name: api
          image: payment-api:v2.0
```

**Step 3 — Service routes to BOTH:**
```yaml
spec:
  selector:
    app: payment-api   # Matches both stable and canary pods
```

**Traffic split:** 9 stable pods + 1 canary pod = **~10% to canary**. Monitor error rates. If healthy, gradually increase canary replicas and decrease stable replicas.

---

### Question 65: Walk me through a full CrashLoopBackOff root cause analysis.

**Answer:**
`CrashLoopBackOff` means the container starts, crashes, and K8s keeps trying to restart it (with exponential delay: 10s → 20s → 40s → ... → 5min max).

**Step 1 — Identify the pod:**
```bash
kubectl get pods -n production
# NAME                    READY   STATUS             RESTARTS   AGE
# payment-api-6b4d-xz2p   0/1     CrashLoopBackOff   8          12m
```

**Step 2 — Read logs from the crashed instance:**
```bash
kubectl logs payment-api-6b4d-xz2p --previous
# This shows stdout from the LAST crash, not the current (restarting) container
```

**Step 3 — Inspect the exit code:**
```bash
kubectl describe pod payment-api-6b4d-xz2p
# Last State: Terminated
#   Reason:   OOMKilled    ← exit code 137 = OOMKill
#   Reason:   Error        ← exit code 1 = Application error
#   Reason:   Error        ← exit code 2 = Misuse of shell command
#   Reason:   Completed    ← exit code 0 = App exited normally (wrong image?)
```

**Common root causes and their fixes:**

| Exit Code / Symptom | Root Cause | Fix |
|---|---|---|
| `OOMKilled` (exit 137) | Memory limit too low | Increase `limits.memory` |
| Exit 1 + DB connect error in logs | Wrong DB env var / secret | Check ConfigMap/Secret values |
| Exit 1 + port bind error | Another container using same port | Fix port config |
| Exit 0 | App completes instantly (wrong image/command) | Check `command`/`args` in spec |
| Exit 1 during init | InitContainer failing | Fix init container first |
| `ImagePullBackOff` before crash | Wrong image name/tag | Fix image reference |

**Step 4 — Temporarily override the entrypoint for debugging:**
```bash
# Override the crashing command to open a shell instead
kubectl debug pod/payment-api-6b4d-xz2p -it --copy-to=debug-pod \
  --container=api -- /bin/sh
```
Inside the shell, manually run the application startup command to see the actual error interactively.
