# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 7 — Final)
> **Target:** TCS, Wipro, Infosys, Cognizant, HCL, Tech Mahindra, IBM, Capgemini, etc.
> **Focus:** `kubectl debug` & Ephemeral Containers, ConfigMap patterns, HPA with custom metrics, Event monitoring, essential kubectl productivity tips, and K8s networking fundamentals recap.

---

## 🔹 Advanced Debugging Techniques (Questions 76-80)

### Question 76: What is an Ephemeral Container and how does `kubectl debug` use it?

**Answer:**
An **Ephemeral Container** is a temporary container added to a running pod *without* restarting it, primarily for debugging. Unlike regular containers, they cannot be restarted and have no resource limits or probes.

**Problem it solves:**
Production images are often built `FROM scratch` or `FROM distroless` — they have no shell (`sh`, `bash`), no debugging tools (`curl`, `netstat`, `ps`). You can't exec into them to troubleshoot.

**`kubectl debug` adds an ephemeral container with your chosen debug image:**
```bash
# Attach a debug container to a running pod using a full-featured image
kubectl debug -it payment-api-7d9f-xz2p \
  --image=busybox:latest \
  --target=payment-api \         # Share the process namespace of THIS container
  --container=debugger           # Name for the ephemeral container

# Inside the ephemeral container — you can now inspect the main app's processes
# ls /proc/1/fd          — inspect open file descriptors
# cat /proc/1/net/tcp    — check network connections
# wget http://localhost:8080/health  — test internal endpoints
```

**Copy + debug a crashing pod (without the crash loop):**
```bash
# Creates a copy of the crashing pod with the command overridden
kubectl debug payment-api-7d9f-xz2p \
  --copy-to=debug-payment \
  --container=payment-api \
  -- /bin/sh                    # Override entrypoint with a shell

# Now manually run the startup command inside the shell to see the real error
```

**Debug a node directly:**
```bash
# Spin up a privileged pod on a specific node with host access
kubectl debug node/node-worker-1 \
  -it \
  --image=ubuntu \
  -- nsenter -t 1 -m -u -i -n -p -- bash
# Full host access for node-level debugging (network, processes, disk)
```

---

### Question 77: A pod claims to be healthy but users report errors. How do you differentiate between application bugs vs Kubernetes infrastructure issues?

**Answer:**
Structured approach to isolate the layer:

**Step 1 — Verify Kubernetes infrastructure is sound:**
```bash
# Is the pod actually Ready? (readiness probe passing = yes)
kubectl get pod payment-api -o jsonpath='{.status.conditions}'

# Is the Service sending traffic to this pod? (check endpoints)
kubectl get endpoints payment-api
# If empty — label mismatch between Service selector and Pod labels

# Are all replicas receiving balanced traffic?
kubectl top pod -l app=payment-api
```

**Step 2 — Check if traffic is actually reaching the pod:**
```bash
# Use kubectl port-forward to bypass the Service/Ingress entirely
kubectl port-forward pod/payment-api-7d9f-xz2p 8080:8080

# Now call the pod directly
curl http://localhost:8080/api/payment/status
# If this works but the Service doesn't → kube-proxy / iptables issue
# If this also fails → application bug
```

**Step 3 — Inspect live traffic (with Istio/Envoy):**
```bash
# Tail Envoy access logs for the specific pod
kubectl logs payment-api-7d9f-xz2p -c istio-proxy | tail -100
# Access logs show: status codes, latency, upstream cluster, response flags
```

**Step 4 — Check resource pressure (pod throttled by CPU limits):**
```bash
kubectl top pod payment-api-7d9f-xz2p
# If CPU usage = CPU limit → pod is being throttled → latency spikes
# Solution: Increase CPU limit OR reduce CPU requests to not trigger throttling
```

---

### Question 78: Explain how Kubernetes Events work and how you monitor them for cluster health.

**Answer:**
**Events** are API objects that record noteworthy occurrences in the cluster. They have a TTL of ~1 hour by default (stored in etcd, then purged).

**Viewing events:**
```bash
# All events in a namespace (sorted by time)
kubectl get events -n production --sort-by='.lastTimestamp'

# Events for a specific resource
kubectl describe pod payment-api-7d9f-xz2p
# The "Events:" section at the bottom is the most useful part of describe

# Watch events in real-time
kubectl get events -n production -w

# Filter only warning events
kubectl get events -n production --field-selector type=Warning
```

**Common warning events and their meanings:**

| Event Reason | Component | Meaning |
|---|---|---|
| `FailedScheduling` | Scheduler | No suitable node found |
| `BackOff` | Kubelet | Container is crashlooping |
| `OOMKilling` | Kubelet | Container exceeded memory limit |
| `FailedMount` | Kubelet | PVC couldn't be mounted |
| `Unhealthy` | Kubelet | Liveness/Readiness probe failed |
| `Evicted` | Kubelet | Pod evicted due to node pressure |
| `FailedCreate` | ReplicaSet | Can't create pod (quota hit) |

**Persistent event monitoring with `kube-state-metrics`:**
```promql
# Alert if events show pods failing scheduling for > 5 min
count(kube_event_count{reason="FailedScheduling", type="Warning"}) > 0
```

---

### Question 79: How do you implement HPA with custom metrics from Prometheus?

**Answer:**
The default HPA only supports CPU and memory. For custom metrics (e.g., requests-per-second, queue depth), you need an adapter.

**Architecture:**
```
[Prometheus] → [Prometheus Adapter] → [Custom Metrics API] → [HPA Controller]
```

**Step 1 — Install Prometheus Adapter:**
```bash
helm install prometheus-adapter prometheus-community/prometheus-adapter \
  --set prometheus.url=http://prometheus-operated.monitoring.svc \
  --set rules.default=true
```

**Step 2 — Configure a custom metric rule in the adapter:**
```yaml
# In prometheus-adapter ConfigMap
rules:
  custom:
    - seriesQuery: 'http_requests_total{namespace!="",pod!=""}'
      resources:
        overrides:
          namespace: {resource: "namespace"}
          pod: {resource: "pod"}
      name:
        matches: "^(.*)_total$"
        as: "${1}_per_second"
      metricsQuery: 'rate(<<.Series>>{<<.LabelMatchers>>}[2m])'
```

**Step 3 — Create HPA using the custom metric:**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: payment-api-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: payment-api
  minReplicas: 2
  maxReplicas: 20
  metrics:
    - type: Pods
      pods:
        metric:
          name: http_requests_per_second    # Our custom metric
        target:
          type: AverageValue
          averageValue: "500"              # Scale when avg rps per pod > 500
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 70           # Also scale on CPU
```

**Verify the custom metric is visible:**
```bash
kubectl get --raw /apis/custom.metrics.k8s.io/v1beta1 | jq .
kubectl describe hpa payment-api-hpa
# Shows current metric values and HPA's scaling decisions
```

---

### Question 80: What is the difference between `kubectl apply`, `kubectl patch`, `kubectl replace`, and `kubectl edit`?

**Answer:**

| Command | Mechanism | Use Case | Risk |
|---|---|---|---|
| `kubectl apply` | 3-way strategic merge using `last-applied-config` annotation | Standard declarative updates from YAML file | Low — safe default |
| `kubectl patch` | Inline partial update (JSON Merge Patch or Strategic Merge Patch) | Quick targeted field changes without a full file | Medium — easy to mistype |
| `kubectl replace` | Complete replacement — deletes old, creates new from file | Major structural changes not possible with apply | High — causes downtime for non-replaceable resources |
| `kubectl edit` | Opens current live spec in `$EDITOR`, apply on save | Ad-hoc emergency edits | Medium — changes aren't in Git |

**`kubectl patch` examples:**
```bash
# Patch a single field — add a label
kubectl patch deployment payment-api \
  -p '{"metadata":{"labels":{"env":"prod"}}}'

# Strategic merge patch — only update the image of one specific container
kubectl patch deployment payment-api --type='json' \
  -p='[{"op": "replace", "path": "/spec/template/spec/containers/0/image", "value":"payment-api:v2.1"}]'

# Add a toleration
kubectl patch node node-1 \
  -p '{"spec":{"taints":[{"effect":"NoSchedule","key":"maintenance","value":"true"}]}}'
```

**`kubectl diff` — preview changes before applying (safety check):**
```bash
# Show what 'kubectl apply' WOULD change without actually changing it
kubectl diff -f payment-deployment.yaml
# Output shows red (removed) / green (added) lines — like a git diff
```

---

## 🔹 ConfigMap & Secret Patterns (Questions 81-83)

### Question 81: What are the different ways to consume a ConfigMap in a Pod, and what are the trade-offs?

**Answer:**

**Method 1 — Environment Variables (most common, simplest):**
```yaml
envFrom:
  - configMapRef:
      name: app-config    # Injects ALL keys as env vars
```
```yaml
env:
  - name: DB_HOST
    valueFrom:
      configMapKeyRef:
        name: app-config
        key: database.host   # Inject a single key
```
- ✅ Simple, universally supported
- ❌ **No hot reload** — pod must restart to see config changes
- ❌ Env vars exposed to all child processes

**Method 2 — Volume Mount (recommended for files):**
```yaml
volumes:
  - name: config-vol
    configMap:
      name: app-config
volumeMounts:
  - name: config-vol
    mountPath: /etc/config
```
- ✅ Files update automatically when ConfigMap changes (kubelet syncs periodically, ~1-2 min)
- ✅ App can watch the file with `inotify` and reload config without restart
- ❌ App must implement file-watching logic

**Method 3 — Reloader (zero-code hot reload):**
Stakater **Reloader** watches ConfigMaps/Secrets for changes and automatically triggers a rolling restart:
```yaml
# Annotation on the Deployment
annotations:
  reloader.stakater.com/auto: "true"
```
- ✅ No app code changes needed
- ✅ Works with both env var and volume mount approaches
- ❌ Causes a rolling restart (slight traffic disruption)

---

### Question 82: How do you avoid base64 secrets leaking in GitOps workflows?

**Answer:**
Native Kubernetes Secrets are only base64-encoded, not encrypted. Committing them to Git is a **critical security vulnerability**.

**Three safe approaches for GitOps:**

**Option 1 — Sealed Secrets (Bitnami):**
```bash
# Encrypt a secret with the cluster's public key
kubeseal --format=yaml < my-secret.yaml > my-sealed-secret.yaml
# my-sealed-secret.yaml is safe to commit — only the cluster can decrypt it

# SealedSecret CRD (safe to store in Git)
apiVersion: bitnami.com/v1alpha1
kind: SealedSecret
metadata:
  name: db-credentials
spec:
  encryptedData:
    password: AgB3x9k... (encrypted ciphertext)
```

**Option 2 — SOPS (Mozilla):**
```bash
# Encrypt secrets file using KMS / age key (stored in .sops.yaml)
sops --encrypt secrets.yaml > secrets.enc.yaml
# Commit secrets.enc.yaml to Git
# CI/CD decrypts on-the-fly using the KMS key — plaintext never stored in Git
```

**Option 3 — External Secrets Operator (ESO):**
```yaml
# Only store an ExternalSecret reference in Git — NOT the value
kind: ExternalSecret
spec:
  secretStoreRef:
    name: aws-secretsmanager
  target:
    name: db-credentials
  data:
    - secretKey: password
      remoteRef:
        key: prod/db/password   # Reference only, value fetched at runtime
```

> **Best practice:** Use ESO with a central secrets vault (Vault/AWS SM) so secret rotation in the vault auto-syncs to the cluster without any Git changes.

---

### Question 83: What are the most important `kubectl` productivity commands every Kubernetes engineer should know?

**Answer:**

**Context & Namespace switching:**
```bash
# View all contexts
kubectl config get-contexts

# Switch cluster context
kubectl config use-context prod-cluster

# Set default namespace (avoid typing -n every time)
kubectl config set-context --current --namespace=production
```

**Output formatting:**
```bash
# Get pod IP and node placement
kubectl get pods -o wide

# Extract a specific field with jsonpath
kubectl get pod payment-api -o jsonpath='{.status.podIP}'
kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="ExternalIP")].address}'

# Custom columns output
kubectl get pods -o custom-columns="NAME:.metadata.name,STATUS:.status.phase,NODE:.spec.nodeName"

# Watch resource changes in real-time
kubectl get pods -w
```

**Labels & Selectors:**
```bash
# Filter by label
kubectl get pods -l app=payment-api,env=prod

# Add/remove labels on the fly
kubectl label pod payment-api-xyz debug=true
kubectl label pod payment-api-xyz debug-    # Remove label

# List all resources with a specific label across all namespaces
kubectl get all -A -l app=payment-api
```

**Resource introspection:**
```bash
# Explain a field's documentation inline
kubectl explain pod.spec.containers.resources.requests

# Get all API resources available in the cluster
kubectl api-resources

# Check your own RBAC permissions
kubectl auth can-i create pods -n production
kubectl auth can-i '*' '*'                  # Am I cluster-admin?
kubectl auth can-i get secrets --as=system:serviceaccount:default:my-sa
```

**Debugging shortcuts:**
```bash
# Get all non-running pods across all namespaces
kubectl get pods -A --field-selector=status.phase!=Running

# Force delete a stuck terminating pod
kubectl delete pod payment-api-xyz --force --grace-period=0

# Show resource usage sorted by memory across cluster
kubectl top pods -A --sort-by=memory

# Generate a YAML skeleton without creating the resource
kubectl create deployment test --image=nginx --dry-run=client -o yaml

# Apply with server-side dry-run (validates on the API server)
kubectl apply -f deployment.yaml --dry-run=server
```
