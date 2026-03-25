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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a ResourceQuota and why is it important in a shared cluster?
**Your Response:** "ResourceQuota is like giving each team a budget for cluster resources. It sets hard limits on how much CPU, memory, and number of pods each namespace can consume. This prevents one team from accidentally using all the resources and starving other teams. When someone tries to create a pod that would exceed their quota, Kubernetes immediately rejects it with a clear error message. It's essential for multi-tenant clusters where different teams share the same infrastructure - it ensures fair resource allocation and prevents noisy neighbor problems."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a LimitRange and how does it differ from ResourceQuota?
**Your Response:** "ResourceQuota and LimitRange work together but at different scopes. ResourceQuota is like the total budget for the entire namespace - it caps the total CPU, memory, and pods the team can use. LimitRange is like setting rules for individual containers - it defines default values and maximum limits per container. The key difference is that ResourceQuota prevents teams from exceeding their total budget, while LimitRange ensures individual containers don't misbehave. In fact, if you have ResourceQuota, you must have LimitRange too, otherwise pods without explicit limits get rejected!"

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you monitor resource usage in a Kubernetes cluster from the command-line?
**Your Response:** "The main tool is `kubectl top`, which shows real-time CPU and memory usage. I use `kubectl top nodes` to see cluster-wide utilization and identify saturated nodes. For pod-level details, I use `kubectl top pods --sort-by=memory` to find memory-hungry pods. I also supplement this with commands like `kubectl describe node` to check node pressure conditions, `kubectl describe resourcequota` to see quota consumption, and `kubectl get pods -w` to watch pod restarts in real-time. These commands give me a complete picture of cluster health from the command line."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** A deployment is consuming far more memory than expected. Walk through your investigation process.
**Your Response:** "I start with `kubectl top pods --sort-by=memory` to identify the memory-hungry pod. Then I check `kubectl describe pod` to see its limits and if it's been OOMKilled recently. I look at the 'Last State' section for exit code 137, which indicates memory kills. Next I check the previous pod logs with `kubectl logs --previous` to see application-level memory errors. Based on what I find, I either increase the memory limit if it's legitimate usage, investigate for memory leaks if it's abnormal, or adjust the limits to match actual usage patterns. The key is systematically checking from cluster level down to application level."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain Horizontal vs Vertical Pod Autoscaling. Which one should you choose for a Java Spring Boot API?
**Your Response:** "HPA scales horizontally by adding more pod replicas, while VPA scales vertically by increasing CPU/memory limits on existing pods. For a Spring Boot API, I'd choose HPA because Spring Boot apps are typically stateless web services that benefit from scaling out to handle more traffic. HPA adds replicas without restarting, which is perfect for traffic spikes. VPA requires pod restarts to apply new resource limits, which would cause downtime. I'd use VPA for things like batch jobs or single-instance services where I want to automatically right-size the JVM heap based on actual usage patterns. The key rule is never use VPA and HPA together on the same CPU metric - they'll conflict with each other."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Blue-Green deployment strategy and how do you implement it in Kubernetes?
**Your Response:** "Blue-Green is like having two identical restaurants side by side. Blue is serving all customers while Green is ready but closed. When Green is fully tested and ready, I instantly switch all traffic from Blue to Green by changing the service selector from 'version: blue' to 'version: green'. The switch is instant - either all traffic goes to Blue or all to Green. The advantage is zero-downtime deployment and instant rollback. The disadvantage is it costs twice as much because both environments run simultaneously. I implement it with two deployments with different version labels and one service that I patch to switch between them."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the difference between Blue-Green and Canary deployments? Which is safer?
**Your Response:** "Blue-Green is like flipping a light switch - either all traffic goes to the old version or all to the new version. Canary is like dimming a light gradually - you send 5% of traffic to the new version, monitor it, then gradually increase to 25%, 50%, and eventually 100%. Canary is generally safer because if there's a bug, only a small percentage of users are affected. With Blue-Green, if the green version has a critical bug, 100% of users get impacted immediately. The trade-off is Canary requires more careful monitoring and gradual rollout, while Blue-Green gives you instant rollback but higher risk."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** You deployed a new version via rolling update but noticed a spike in 5xx errors. How do you rollback immediately?
**Your Response:** "I'd immediately run `kubectl rollout undo deployment/payment-api` to rollback to the previous version. Then I'd verify the rollback is progressing with `kubectl rollout status` and confirm the correct image is running with `kubectl describe deployment`. If I need to rollback to a specific older version, I'd check the history first with `kubectl rollout history` and then use `--to-revision` to target a specific version. The key is acting fast - one command stops the bleeding and gets back to the last known good state. I always annotate deployments with change causes so the history is readable."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a canary release without Argo Rollouts using only native Kubernetes?
**Your Response:** "I use the native Kubernetes pattern with two deployments and one service. I create a 'stable' deployment with 9 replicas running the current version, and a 'canary' deployment with 1 replica running the new version. Both deployments have the same 'app' label so the service routes to both. Since there are 9 stable pods and 1 canary pod, about 10% of traffic goes to the canary. I monitor the canary's performance, and if it looks good, I gradually increase canary replicas while decreasing stable replicas. This gives me a simple canary deployment using only built-in Kubernetes features."

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Walk me through a full CrashLoopBackOff root cause analysis.
**Your Response:** "CrashLoopBackOff means a pod keeps starting and crashing. I start by identifying the pod with `kubectl get pods`, then check the logs from the previous crash with `kubectl logs --previous` to see the application error. Next I run `kubectl describe pod` to check the exit code - 137 means OOMKilled, 1 means application error, 0 means the app exited immediately. Based on the exit code and logs, I can identify the root cause: memory limits, wrong environment variables, port conflicts, or bad image references. If needed, I use `kubectl debug` to create a copy with a shell to manually run the startup command and see the exact error. The key is systematically checking logs, exit codes, and pod events to narrow down the issue."
