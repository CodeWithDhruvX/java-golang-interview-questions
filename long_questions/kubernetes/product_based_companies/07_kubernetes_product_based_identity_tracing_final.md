# 🚀 Kubernetes Interview Questions - Product-Based Companies (Part 7 — Final)
> **Target:** Google, Amazon, Microsoft, Uber, Stripe, Datadog, Cloudflare, etc.
> **Focus deep-dive:** Workload Identity (IRSA/GKE WI), Cluster API, VCluster, Distributed Tracing, OpenTelemetry Operator, Velero Disaster Recovery, and Kubernetes Gateway API.

---

## 🔹 Workload Identity — Cloud IAM Without Secrets (Questions 73-76)

### Question 73: What is AWS IRSA (IAM Roles for Service Accounts) and why is it better than node-level IAM roles?

**Answer:**
**IRSA** allows individual Kubernetes Pods to assume specific AWS IAM roles without sharing credentials through environment variables or EC2 instance profile permissions.

**Problem with node-level IAM roles:**
All pods on a node share the same EC2 instance profile. If your `S3-reader` pod and your `RDS-admin` pod run on the same node, BOTH get the combined permissions, violating the principle of least privilege.

**How IRSA works:**
1. Create an IAM role with a **trust policy** that allows a specific Kubernetes `ServiceAccount` to assume it.
2. Annotate the `ServiceAccount` with the IAM role ARN.
3. EKS's **OIDC provider** issues a signed JWT token (projected into the pod volume).
4. When the AWS SDK in your pod calls `STS AssumeRoleWithWebIdentity`, it presents the JWT.
5. AWS validates the JWT against the cluster's OIDC endpoint and issues short-lived credentials.

```yaml
# Step 1: Create annotated ServiceAccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: s3-reader-sa
  namespace: data-pipeline
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::123456789:role/s3-reader-role

---
# Step 2: Use the ServiceAccount in your Pod
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: s3-reader-sa
  containers:
    - name: app
      image: my-app:latest
      # AWS SDK auto-discovers credentials from the projected token volume
```

**Key advantages:**
- **Credential isolation:** Each pod gets credentials scoped to exactly its ServiceAccount's role.
- **No static secrets:** Tokens are short-lived (15 minutes by default) and auto-refreshed.
- **Automatic rotation:** If a pod is compromised, the token expires and cannot be reused across pod boundaries.

---

### Question 74: How does GKE Workload Identity differ from AWS IRSA?

**Answer:**
Both achieve the same goal (pod-level cloud IAM without static secrets), but the mechanism differs:

| | **AWS IRSA** | **GKE Workload Identity** |
|---|---|---|
| **Token type** | OIDC JWT projected as a volume | GKE-managed service account binding |
| **Trust mechanism** | IAM trust policy trusts K8s OIDC JWKS endpoint | GCP IAM `roles/iam.workloadIdentityUser` binding |
| **Config** | Annotate K8s ServiceAccount with IAM role ARN | Annotate K8s SA + bind GCP SA via `gcloud` |
| **Metadata server** | AWS metadata endpoint + STS | GKE Metadata Server (replaces the instance metadata server) |

**GKE Workload Identity setup:**
```bash
# 1. Enable Workload Identity on the cluster
gcloud container clusters update my-cluster \
  --workload-pool=my-project.svc.id.goog

# 2. Create a GCP Service Account
gcloud iam service-accounts create gcs-reader-sa

# 3. Bind K8s ServiceAccount to GCP SA
gcloud iam service-accounts add-iam-policy-binding \
  gcs-reader-sa@my-project.iam.gserviceaccount.com \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:my-project.svc.id.goog[data-pipeline/k8s-gcs-reader]"

# 4. Annotate the K8s ServiceAccount
kubectl annotate serviceaccount k8s-gcs-reader \
  iam.gke.io/gcp-service-account=gcs-reader-sa@my-project.iam.gserviceaccount.com \
  -n data-pipeline
```

---

### Question 75: How do you securely access Kubernetes API from within a Pod running in the same cluster?

**Answer:**
Every Pod automatically gets a `ServiceAccount` token and the cluster CA certificate mounted at:
```
/var/run/secrets/kubernetes.io/serviceaccount/token
/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```

**Secure in-cluster API access pattern:**
```go
// Go — using the in-cluster config (auto-detects token + CA)
config, err := rest.InClusterConfig()
clientset, err := kubernetes.NewForConfig(config)

// List pods in current namespace
pods, err := clientset.CoreV1().Pods("my-namespace").List(context.Background(), metav1.ListOptions{})
```

**Security hardening:**
1. **Minimal RBAC:** Create a `Role` with only the exact verbs required (e.g., `get` on `pods` only, not `list` or `watch`).
2. **Namespace-scoped:** Use `Role` + `RoleBinding`, not `ClusterRole` + `ClusterRoleBinding`.
3. **Disable auto-mount when unused:** `automountServiceAccountToken: false` in the pod spec stops the default token from being injected for apps that don't need API access.
4. **Bound tokens:** Modern K8s projects short-lived tokens via `projected` volumes with audience and expiration constraints — far safer than the legacy long-lived secrets.

---

### Question 76: What is Cluster API (CAPI) and how does it change how you manage Kubernetes clusters?

**Answer:**
**Cluster API** is a Kubernetes sub-project that brings the **declarative, controller-based approach** of K8s to cluster lifecycle management itself.

Instead of using cloud consoles or custom scripts to create clusters, CAPI lets you define clusters as Kubernetes resources and a **Management Cluster** reconciles them.

**Core CRDs:**
```yaml
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: production-us-east
spec:
  clusterNetwork:
    pods:
      cidrBlocks: ["192.168.0.0/16"]
  controlPlaneRef:
    kind: KubeadmControlPlane
    name: production-us-east-cp
  infrastructureRef:
    kind: AWSCluster        # Provider-specific (AWS, GCP, Azure, vSphere)
    name: production-us-east
```

**Workflow:**
```
[Git: cluster.yaml] → [ArgoCD] → [Management Cluster] → [CAPI Controllers] → [AWS/GCP APIs → New K8s Cluster]
```

**Key benefits:**
- Cluster creation, scaling, and upgrades become GitOps-managed.
- **ClusterClass** (CAPI v1.2+): Define a cluster template once; stamp out 50 identical clusters with variable overrides.
- Cluster upgrades are a simple `spec.version` change — CAPI handles the rolling upgrade orchestration.
- Enables true multi-cluster fleet management from a single management plane.

---

## 🔹 VCluster, Distributed Tracing & OpenTelemetry (Questions 77-80)

### Question 77: What is a vCluster and how does it provide stronger multi-tenancy than Namespaces?

**Answer:**
A **vCluster** (by Loft Labs) is a fully functional Kubernetes control plane running *inside* a namespace of a "host" cluster as a pod.

**Architecture:**
```
Host Cluster
└── namespace: tenant-acme
    ├── vcluster-pod (contains: k3s control plane + etcd)
    └── syncer (maps vcluster workloads → host namespace pods)
```

**What tenants get:**
- Their own **kube-apiserver** — they can create CRDs, RBAC, cluster-scoped resources.
- Full `kubectl` access with their own kubeconfig.
- Isolation from other tenants at the API level.

**How it works under the hood:**
- The vCluster runs its own etcd and API server inside the pod.
- A **Syncer** component watches the vCluster's API for Pods, Services, etc. and **translates** them into actual resources inside the host namespace.
- Worker nodes are shared — actual containers run on the host cluster's nodes (no VM overhead).

**vCluster vs Namespace isolation:**

| | **Namespace** | **vCluster** |
|---|---|---|
| Own CRDs | ❌ Shares cluster CRDs | ✅ Full CRD isolation |
| Own RBAC | ❌ Scoped but shared | ✅ Tenant-owned RBAC tree |
| Cluster-scoped resources | ❌ Cannot create | ✅ Full access (Nodes, PV, etc.) |
| Overhead | None | ~100MB RAM for control plane |
| Use case | Internal team isolation | External customer / CI isolation |

---

### Question 78: How does distributed tracing work in a Kubernetes microservices environment with Istio and Jaeger?

**Answer:**
Distributed tracing follows a request as it hops across multiple services, building a **trace** (timeline) of every span.

**Istio's automatic trace propagation:**
When Istio's Envoy sidecar forwards a request to another service, it:
1. Generates a unique `X-B3-TraceId` and `X-B3-SpanId` header on the first request.
2. On subsequent hops, passes the `TraceId` unchanged and generates new `SpanId` values.
3. Reports all spans to the configured tracing backend (Jaeger, Zipkin, Tempo).

**Critical requirement — Header forwarding:**
Istio generates the first span automatically, but your **application code must forward the trace headers** to downstream calls. Otherwise, each service appears as a separate, disconnected trace.

```go
// Go — forward B3 trace headers to downstream services
func callDownstream(ctx context.Context, r *http.Request) {
    req, _ := http.NewRequest("GET", "http://inventory-service/items", nil)
    // Forward trace headers from incoming request
    for _, h := range []string{"x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid", "x-b3-sampled", "x-request-id"} {
        if v := r.Header.Get(h); v != "" {
            req.Header.Set(h, v)
        }
    }
    http.DefaultClient.Do(req)
}
```

**Jaeger deployment in K8s:**
```bash
helm install jaeger jaegertracing/jaeger \
  --set collector.enabled=true \
  --set storage.type=elasticsearch
```

**Istio tracing config:**
```yaml
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  meshConfig:
    enableTracing: true
    defaultConfig:
      tracing:
        zipkin:
          address: jaeger-collector.monitoring:9411
        sampling: 10.0    # Sample 10% of requests (100% is too expensive at scale)
```

---

### Question 79: What is the OpenTelemetry Operator and how does it simplify observability in Kubernetes?

**Answer:**
The **OpenTelemetry Operator** is a Kubernetes operator that manages `OpenTelemetryCollector` and `Instrumentation` CRDs, enabling standardized telemetry (metrics, traces, logs) collection.

**Key capabilities:**

**1. Deploy a centralized OTel Collector as a DaemonSet:**
```yaml
apiVersion: opentelemetry.io/v1alpha1
kind: OpenTelemetryCollector
metadata:
  name: otel-daemonset
spec:
  mode: daemonset    # One collector per node
  config: |
    receivers:
      otlp:
        protocols:
          grpc:        # Apps send traces to localhost:4317
    processors:
      batch:
    exporters:
      jaeger:
        endpoint: jaeger-collector:14250
      prometheus:
        endpoint: "0.0.0.0:8889"
    service:
      pipelines:
        traces:
          receivers: [otlp]
          processors: [batch]
          exporters: [jaeger]
        metrics:
          receivers: [otlp]
          exporters: [prometheus]
```

**2. Auto-instrument applications without code changes:**
```yaml
apiVersion: opentelemetry.io/v1alpha1
kind: Instrumentation
metadata:
  name: java-auto-instrument
spec:
  java:
    image: ghcr.io/open-telemetry/opentelemetry-operator/autoinstrumentation-java:latest
  exporter:
    endpoint: http://otel-daemonset-collector:4317
```

The operator injects the OTel Java agent as an init container into any pod annotated with:
```yaml
annotations:
  instrumentation.opentelemetry.io/inject-java: "true"
```

App gets full traces, metrics, and logs forwarded to your backend — without a single line of instrumentation code.

---

### Question 80: Perform a Velero backup and restore. How do you recover from accidental namespace deletion?

**Answer:**
**Velero** backs up K8s resources (API objects) to object storage (S3/GCS) and optionally snapshots PVCs.

**Setup:**
```bash
# Install Velero with AWS S3 backend
velero install \
  --provider aws \
  --plugins velero/velero-plugin-for-aws:v1.8.0 \
  --bucket k8s-velero-backups \
  --backup-location-config region=us-east-1 \
  --snapshot-location-config region=us-east-1 \
  --secret-file ./credentials-velero
```

**Scheduled backup:**
```bash
# Daily backup of all namespaces at midnight, retained for 30 days
velero schedule create daily-cluster-backup \
  --schedule="0 0 * * *" \
  --ttl 720h

# Backup a specific namespace before a risky change
velero backup create pre-migration-backup \
  --include-namespaces payments \
  --wait
```

**Disaster recovery — restore after accidental `kubectl delete namespace payments`:**
```bash
# Step 1: List available backups
velero backup get
# NAME                        STATUS     CREATED                         EXPIRES
# daily-cluster-backup-2024   Completed  2024-10-15 00:00:00 UTC         29d

# Step 2: Describe the backup to verify it contains the namespace
velero backup describe daily-cluster-backup-2024 --details

# Step 3: Restore the specific namespace
velero restore create \
  --from-backup daily-cluster-backup-2024 \
  --include-namespaces payments \
  --wait

# Step 4: Verify restore
velero restore describe <restore-name>
kubectl get all -n payments
```

**PVC restore with volume snapshots:**
```bash
# Backup with volume snapshot (creates EBS snapshot)
velero backup create payments-with-data \
  --include-namespaces payments \
  --snapshot-volumes=true

# Restore including PVC data
velero restore create --from-backup payments-with-data \
  --restore-volumes=true
```

---

## 🔹 Kubernetes Gateway API — The Future of Ingress (Questions 81-83)

### Question 81: What is the Kubernetes Gateway API and why is it replacing Ingress?

**Answer:**
The **Gateway API** is the official successor to the `Ingress` resource, addressing its limitations:

**Ingress limitations:**
- Single resource handles everything → vendor-specific annotations proliferate (e.g., `nginx.ingress.kubernetes.io/rate-limit`).
- No native support for TCP/UDP routing.
- No role separation — developers and infrastructure teams share the same resource.
- Cannot express traffic weighting, header manipulation, or retries natively.

**Gateway API's role-separated model:**
```
Infrastructure Team                  Developer Team
─────────────────────────            ─────────────
GatewayClass (defines infra)  →  Gateway (binds to class)  →  HTTPRoute (defines routing)
```

**GatewayClass** — cluster-scoped, owned by infra team (defines the controller to use):
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: GatewayClass
metadata:
  name: nginx-gateway
spec:
  controllerName: gateway.nginx.org/nginx-gateway-controller
```

**Gateway** — cluster or namespace-scoped, defines the listener:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: prod-gateway
spec:
  gatewayClassName: nginx-gateway
  listeners:
    - name: https
      protocol: HTTPS
      port: 443
      tls:
        certificateRefs:
          - name: prod-tls-secret
```

**HTTPRoute** — namespace-scoped, owned by the developer team:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: payment-routes
spec:
  parentRefs:
    - name: prod-gateway
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /api/payments
      filters:
        - type: RequestHeaderModifier
          requestHeaderModifier:
            add:
              - name: "X-Internal-Request"
                value: "true"
      backendRefs:
        - name: payment-service
          port: 8080
          weight: 90
        - name: payment-service-canary
          port: 8080
          weight: 10    # Native 90/10 traffic split — no annotations needed!
```

---

### Question 82: What types of routes does the Gateway API support beyond HTTP?

**Answer:**
The Gateway API has protocol-specific route types, solving the biggest Ingress limitation (HTTP-only):

| Route Type | Protocol | Use Case |
|---|---|---|
| `HTTPRoute` | HTTP/HTTPS | Web APIs, REST, gRPC (via HTTP/2) |
| `TCPRoute` | Raw TCP | Database connections, MQTT brokers |
| `TLSRoute` | TLS (SNI-based) | Route TLS traffic by SNI hostname without terminating |
| `GRPCRoute` | gRPC | Native gRPC method-level routing (e.g., route by `.package.Service/Method`) |
| `UDPRoute` | UDP | DNS servers, game servers, media streaming |

**GRPCRoute example** (native gRPC routing — no annotations needed):
```yaml
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: GRPCRoute
metadata:
  name: payments-grpc
spec:
  parentRefs:
    - name: prod-gateway
  rules:
    - matches:
        - method:
            service: payments.v1.PaymentService
            method: ProcessPayment        # Route only this specific RPC method
      backendRefs:
        - name: payment-grpc-service
          port: 9090
```

---

### Question 83: How do you handle cross-cutting concerns like rate limiting and authentication in the Gateway API?

**Answer:**
Gateway API handles cross-cutting concerns via **`Policies`** that attach to Gateway/Route resources.

**`BackendTLSPolicy`** — enforce TLS to the backend pod (end-to-end encryption):
```yaml
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: BackendTLSPolicy
metadata:
  name: backend-tls
spec:
  targetRef:
    kind: Service
    name: payment-service
  tls:
    caCertificateRefs:
      - name: internal-ca
    hostname: payment-service.internal
```

**Controller-specific extension policies** (example: NGINX Gateway Fabric):
```yaml
apiVersion: gateway.nginx.org/v1alpha1
kind: ObservabilityPolicy
metadata:
  name: rate-limit-policy
spec:
  targetRef:
    kind: HTTPRoute
    name: payment-routes
  tracing:
    strategy: ratio
    ratio: 10      # Trace 10% of requests
---
apiVersion: gateway.nginx.org/v1alpha1
kind: ClientSettingsPolicy
spec:
  targetRef:
    kind: Gateway
    name: prod-gateway
  keepAlive:
    requests: 1000
    time: 1h
```

**Auth via `ExtensionRef` (request authentication):**
```yaml
# HTTPRoute filter — call an external auth service before routing
filters:
  - type: ExtensionRef
    extensionRef:
      group: gateway.nginx.org
      kind: RequestHeaderModifier
```

> **Interview key point:** The Gateway API's `Policy` attachment model is intentional — it separates "who defines the route" from "who enforces cross-cutting concerns", enabling infrastructure and developer teams to work independently without merge conflicts on a single YAML file.
