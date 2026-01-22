## 🔹 Edge Cases & Troubleshooting Scenarios (Questions 751-760)

### Question 751: How do you recover a deleted namespace?

**Answer:**
Often namespaces get stuck in `Terminating`.
1.  Check `kubectl get ns my-ns -o json`.
2.  Look for `finalizers`.
3.  Manually remove the finalizer string via API call (`/finalize` endpoint) using `curl`.
4.  Warning: This orphans resources.

---

### Question 752: What happens if a container in a pod crashes continuously?

**Answer:**
- Pod status: `CrashLoopBackOff`.
- Kubelet restarts it with exponential backoff (10s, 20s, 40s, 5m).
- Check logs: `kubectl logs my-pod --previous`.

---

### Question 753: How do you prevent log flooding in a crash loop?

**Answer:**
- **App level:** Fix the bug.
- **Node level:** The container runtime rotates log files.
- **System:** `logging` agents (Fluentd) might get overwhelmed. Use rate limits in Fluentd.

---

### Question 754: What does `Terminating` status mean and how to debug it?

**Answer:**
Means `deletionTimestamp` is set.
- Pod is waiting for `preStop` hook?
- Pod is waiting for `terminationGracePeriod`?
- Pod handling of SIGTERM is stuck?
- Node is down?

---

### Question 755: What is a zombie pod and how do you clean it up?

**Answer:**
A pod that remains in API server even after node is gone.
- `kubectl delete pod my-pod --grace-period=0 --force`.
- Removes it from Etcd immediately.

---

### Question 756: How do you detect node pressure conditions?

**Answer:**
`kubectl describe node`.
- Conditions: `MemoryPressure`, `DiskPressure`, `PIDPressure`.
- Kubelet sets these when resources run low.

---

### Question 757: What causes a pod to be stuck in `ContainerCreating`?

**Answer:**
1.  **Mounting Volume:** CSI Driver timeout? (EBS stuck).
2.  **CNI Plugin:** IP Address exhaustion?
3.  **Secrets:** Missing Secret or ConfigMap referenced in PodSpec.

---

### Question 758: What is an image pull backoff error?

**Answer:**
Cannot pull image.
- Review `kubectl describe pod`.
- Check Secret (ImagePullSecret).
- Check Network (Node to Registry).
- Check Image Name/Tag.

---

### Question 759: How do you detect and clean orphaned resources?

**Answer:**
- **Orphaned Pods:** Pods without Owners.
- **Orphaned PVCs:** Released but not deleted.
- **Tools:** `kubectl get pods --all-namespaces` | grep/awk. Or `kor` (Kubernetes Orphan Resource cleaner).

---

### Question 760: How do you handle resource quota starvation?

**Answer:**
One team eats all CPU.
- **Immediate:** Delete their pods.
- **Fix:** Set `ResourceQuota` on their namespace.
- **Fix:** Set `LimitRange` to prevent massive pod requests.

---

## 🔹 Advanced Pod Lifecycle & Management (Questions 761-770)

### Question 761: What is the difference between liveness and readiness probes?

**Answer:**
- **Liveness:** "Restart me if I fail." (Deadlock).
- **Readiness:** "Don't send traffic unless I pass." (Startup/Overload).

---

### Question 762: How does a startup probe differ from a liveness probe?

**Answer:**
**Startup Probe:**
- Runs FIRST.
- Disables Liveness/Readiness checks until it passes.
- Use Case: Legacy app takes 3 minutes to boot. Liveness would kill it at 30s without this.

---

### Question 763: What are init containers used for?

**Answer:**
Setup scripts.
- `nslookup db-service` (Wait for DB).
- `download-assets`.
- `chown /data`.
- Run to completion before Main containers start.

---

### Question 764: How do sidecars affect the pod lifecycle?

**Answer:**
Native Sidecars (K8s 1.28+):
- `restartPolicy: Always` in initContainers.
- They start before main container, shut down after main container.
- Solves the "Job finishes but sidecar keeps Pod running" problem.

---

### Question 765: What is `terminationGracePeriodSeconds`?

**Answer:**
Time Kubelet waits after SIGTERM before SIGKILL.
- Default: 30s.
- Make sure your app drains connections within this time.

---

### Question 766: How do you ensure graceful shutdown of pods?

**Answer:**
- Handle OS Signals (`SIGTERM`).
- Stop accepting new TCP connections.
- Finish processing active requests.
- Close DB/File handles.
- Exit.

---

### Question 767: How do you debug pod termination issues?

**Answer:**
- Check logs closely as it shuts down.
- Use `preStop` hook to log "Stopping..." to a file on hostPath to verify it runs.

---

### Question 768: How does `preStop` hook work?

**Answer:**
Kubelet executes this command inside container.
- **Blocking call.**
- SIGTERM is NOT sent until preStop completes.
- Useful to delay shutdown: `sleep 5` (Allow LoadBalancer to update).

---

### Question 769: What happens if a pod never becomes Ready?

**Answer:**
- It is removed from Service Endpoints.
- Deployment Controller waits.
- Rollout pauses/times out.
- Old pods continue running (safe).

---

### Question 770: How do you prioritize pod scheduling with `priorityClassName`?

**Answer:**
1.  Create `PriorityClass` (value: 1000).
2.  Assign to Pod via `priorityClassName`.
3.  If cluster full, Scheduler preempts (evicts) lower priority pods to make room.

---

## 🔹 Advanced Job & Batch Workloads (Questions 771-780)

### Question 771: What is the difference between Job and DaemonSet?

**Answer:**
- **Job:** Runs to completion (Exit 0).
- **DaemonSet:** Runs forever (Services) on every node.

---

### Question 772: How does a CronJob handle missed schedules?

**Answer:**
`startingDeadlineSeconds`.
- If controller is down at 10:00 (Scheduled time).
- It comes up at 10:02.
- If deadline > 120s, it starts the job. Else, it skips it as "Missed".

---

### Question 773: What is `concurrencyPolicy` in a CronJob?

**Answer:**
- **Allow:** Concurrent runs allowed.
- **Forbid:** If previous job running, skip new one.
- **Replace:** If previous job running, kill it, start new one.

---

### Question 774: How do you avoid duplicated job runs?

**Answer:**
K8s Jobs provide "At least once" guarantee.
- Rarely, duplicates happen.
- **App Logic:** Must be idempotent. (DB Unique constraints).

---

### Question 775: How can jobs fail silently and how to detect it?

**Answer:**
- `backoffLimit` reached. Job marked Failed.
- **Monitoring:** Prometheus `kube_job_status_failed`. Alert on this!

---

### Question 776: How do you implement retry logic in batch jobs?

**Answer:**
`backoffLimit` field (default 6).
- K8s retries the pod with exponential backoff on failure.

---

### Question 777: What is a parallel job with indexed completion?

**Answer:**
`completionMode: Indexed`.
- Creates pods: `job-0`, `job-1`, `job-2`.
- Inside pod, `$JOB_COMPLETION_INDEX` env var is available.
- Use case: Sharding data processing (Pod 0 processes Part 0).

---

### Question 778: What are suspend/resume features in Jobs?

**Answer:**
`suspend: true`.
- Job exists but controller doesn't create pods.
- Useful for queuing or pausing expensive operations.

---

### Question 779: How do you clean up old Jobs automatically?

**Answer:**
`ttlSecondsAfterFinished: 100`.
- Controller deletes the Job (and Pods) 100s after completion.
- Keeps etcd clean.

---

### Question 780: How do you monitor and alert on Job status?

**Answer:**
- **Kube-state-metrics:** `kube_job_status_succeeded`.
- **Alert:** `kube_job_status_failed > 0`.

---

## 🔹 Real-time & Low Latency Workloads (Questions 781-790)

### Question 781: How do you run low-latency workloads on Kubernetes?

**Answer:**
- **CPU Manager:** Static policy (Pinning).
- **HugePages:** Pre-allocate RAM.
- **Realtime Kernel:** Tune OS.
- **Networking:** SR-IOV / DPDK.

---

### Question 782: What is CPU pinning and how is it achieved?

**Answer:**
Prevents CPU context switching across cores.
- Kubelet config: `--cpu-manager-policy=static`.
- Pod Request: Must be Integer (`cpu: 2`) and Request==Limit (Guaranteed QoS).
- Kubelet assigns exclusive cores.

---

### Question 783: How do you configure guaranteed QoS class?

**Answer:**
set `requests` == `limits` for both CPU and Memory.
- Highest priority.
- Exempt from Eviction (mostly).
- Gets exclusive CPUs (if configured).

---

### Question 784: What are real-time kernel considerations?

**Answer:**
`PREEMPT_RT` patch on Linux.
- Reduces kernel latency.
- Required for Telco/HFT applications.

---

### Question 785: How do you ensure predictable scheduling latency?

**Answer:**
- Use **Dedicated Nodes** (Taints/Tolerations).
- Pre-pull images.
- Ensure efficient scheduler config.

---

### Question 786: What are HugePages and how are they configured?

**Answer:**
Large memory pages (2MB/1GB instead of 4KB).
- Reduces TLB cache misses.
- Node must have HugePages enabled.
- Pod: `resources: limits: hugepages-2Mi: 100Mi`.

---

### Question 787: How does Kubernetes support DPDK applications?

**Answer:**
Data Plane Development Kit (Fast packet processing).
- Uses SR-IOV CNI.
- Pod gets direct access to NIC hardware (bypassing Kernel network stack).

---

### Question 788: What’s the difference between guaranteed vs best-effort pod QoS?

**Answer:**
- **Guaranteed:** Request=Limit. Committed resources.
- **Burstable:** Request < Limit. Sharing pool.
- **BestEffort:** No Request. Scavenger.

---

### Question 789: How do you run audio/video processing workloads?

**Answer:**
- High CPU/GPU requirements.
- Likely Job based.
- Use **Node Affinity** to land on compute-optimized instances.

---

### Question 790: How do you reduce cold start times for latency-sensitive pods?

**Answer:**
- **Pause Containers:** Keep a pool of warm "Pause" containers? (Complex).
- **Node Caching:** Image already on node.
- **WASM:** Faster application start.

---

## 🔹 Design Principles & Industry Patterns (Questions 791-800)

### Question 791: What is GitOps vs ChatOps in Kubernetes operations?

**Answer:**
- **GitOps:** Repo is source of truth. Automated.
- **ChatOps:** Slack bot (`/deploy my-app`). Trigger based. often triggers GitOps or CI pipeline under hood.

---

### Question 792: What are 12-factor app considerations in K8s?

**Answer:**
- **Config:** Env Vars (ConfigMaps).
- **Backing Services:** Attached resources (Secrets/ExternalNames).
- **Logs:** Stdout (Fluentd).
- **Concurrency:** Scale out (Replicas).

---

### Question 793: How do you enforce resource limits as best practice?

**Answer:**
Admission Controller (Kyverno/Gatekeeper).
- "Deny Pod if no Requests/Limits set".
- "Deny Pod if Memory Request > 8GB".

---

### Question 794: What are common anti-patterns in Helm usage?

**Answer:**
- **Super-Chart:** One giant chart ("The Monolith") for 50 services. Hard to maintain.
- **Hardcoding:** putting environments in `if/else` logic inside templates instead of `values.yaml`.

---

### Question 795: What is an opinionated vs flexible Kubernetes platform?

**Answer:**
- **Opinionated:** (OpenShift). Decisions made for you (Router, Registry, UI included).
- **Flexible:** (EKS + DIY). You pick Ingress, Monitoring, etc.

---

### Question 796: What are design considerations for multi-cloud clusters?

**Answer:**
- **Data Gravity:** Data transfer costs ($$$).
- **Latency:** DB in AWS, App in Azure? No.
- **Abstraction:** Need common API (Crossplane) to provision generic resources.

---

### Question 797: How do you manage regional failover in Kubernetes?

**Answer:**
DNS layer (Global Traffic Manager).
- Health check fails for `us-east`.
- DNS points user to `eu-west`.
- Clusters are independent.

---

### Question 798: What is the pet vs cattle analogy in Kubernetes?

**Answer:**
- **Pets:** (VMs). Named (db-01), Hand-nursed, Unique.
- **Cattle:** (Pods). Numbers (app-x9d1), Replaceable, identical.
- K8s treats everything as Cattle (mostly).

---

### Question 799: When is it better to use serverless over Kubernetes?

**Answer:**
- Low operational overhead needed.
- Event-driven.
- Scale-to-zero required.
- **Cost:** Low traffic is cheaper on Serverless. High traffic is cheaper on K8s (Flat fee nodes).

---

### Question 800: How do you align Kubernetes architecture with DORA metrics?

**Answer:**
- **Deployment Frequency:** Easy with K8s/GitOps.
- **Lead Time:** Fast CI/CD.
- **MTTR:** Self-healing pods / Rolling Rollbacks.
- **Change Failure Rate:** Canary deployments reduce this.
