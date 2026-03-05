# 💰 Microservices — Cost Optimization & FinOps on Kubernetes

> **Level:** 🔴 Senior – Principal
> **Asked at:** CRED, Zepto, Razorpay, Flipkart, Swiggy (Engineering leadership, L6+)

> **Interview Reality:** At scaling Indian startups (CRED, Zepto), engineering managers and VPs routinely ask senior engineers about cloud costs. A question like *"Our K8s bill grew 3x in 6 months — what do you look at?"* is being asked in final rounds.

---

## Q1. How do you right-size pods in Kubernetes? Walk me through the process.

**"Right-sizing means setting `requests` and `limits` accurately — not too low (risking OOM kills or throttling) and not too high (wasting money by reserving idle capacity).**

**The problem with over-provisioning:**  
If `memory: 2Gi` is requested but the pod uses only 400MB peak, the 1.6Gi difference is wasted on that node. Multiply by 50 services × 3 replicas = 240Gi of paid but unused RAM.

**Step-by-step right-sizing process:**

**Step 1: Instrument and observe (2–4 weeks)**
- Deploy with generous initial limits. Let the service run real production traffic.
- Query Prometheus for actual consumption:
```promql
# MAX CPU used by order-service pods over 2 weeks
max_over_time(
  rate(container_cpu_usage_seconds_total{container="order-service"}[5m])[2w:5m]
)

# P99 memory used by order-service pods over 2 weeks  
quantile_over_time(0.99,
  container_memory_working_set_bytes{container="order-service"}[2w:5m]
)
```

**Step 2: Apply the right-sizing formula**
```
CPU Request    = P90 actual usage × 1.25 (25% headroom)
CPU Limit      = P99 actual usage × 1.5  (spike headroom)
Memory Request = P99 actual usage × 1.2  (20% headroom)
Memory Limit   = P99 actual usage × 1.4  (OOM buffer)
```

**Step 3: Use VPA (Vertical Pod Autoscaler) in Recommendation mode**
```yaml
apiVersion: autoscaling.k8s.io/v1
kind: VerticalPodAutoscaler
metadata:
  name: order-service-vpa
spec:
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: order-service
  updatePolicy:
    updateMode: "Off"  # Recommendation only — do NOT auto-apply in production
```
```bash
kubectl describe vpa order-service-vpa
# Shows: Lower Bound, Target (recommended), Upper Bound for CPU + Memory
```

**Step 4: Apply gradually to staging → production**"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** CRED, Zepto — startups where cloud costs hit $500K+/month at moderate scale. This is a leadership question, not just a technical one.

#### Indepth
**VPA Modes:**
- `Off` — only recommendations, you apply manually (safest for production)
- `Initial` — sets resources at pod creation, no live updates
- `Auto` — applies recommendations live (restarts pods) — risky, use with caution

**Goldilocks:** A popular open-source tool (Fairwinds) that runs VPA in recommendation mode and provides a nice UI showing cost savings by namespace.

---

## Q2. How do you use Spot/Preemptible instances for microservices on EKS?

**"Spot instances offer up to 90% cost savings over On-Demand, but AWS can reclaim them with 2 minutes notice. The architecture must tolerate this.**

**Node Group Strategy (EKS):**
```
┌────────────────────────────────────────────────────────┐
│  EKS Cluster                                          │
│                                                       │
│  On-Demand Node Group (20% of capacity):              │
│  • Core infrastructure: ArgoCD, Istio, Prometheus     │
│  • Stateful workloads                                 │
│  • Minimum: 2 nodes always running                    │
│                                                       │
│  Spot Node Group (80% of capacity):                   │
│  • All stateless microservices (Order, Payment, API)  │
│  • Multiple instance families (m5, m5a, m4, m5d)      │
│    to maximize spot availability                      │
│  • Karpenter auto-provisions new nodes on demand      │
└────────────────────────────────────────────────────────┘
```

**Kubernetes tolerations to schedule on spot nodes:**
```yaml
# In your Deployment spec
spec:
  template:
    spec:
      tolerations:
      - key: "spot"
        operator: "Equal"
        value: "true"
        effect: "NoSchedule"

      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 80
            preference:
              matchExpressions:
              - key: node-type
                operator: In
                values: ["spot"]
```

**Handling Spot Interruption (2-minute warning):**
1. **AWS Node Termination Handler** — a DaemonSet that listens for EC2 spot interruption events
2. When a node gets the termination notice → it cordons the node → K8s evicts pods gracefully → pods schedule on healthy nodes
3. Your service must have `PodDisruptionBudget` to ensure continuity:
```yaml
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: order-service-pdb
spec:
  minAvailable: 2  # At least 2 pods must be running during disruption
  selector:
    matchLabels:
      app: order-service
```

**Key rule:** Never run fewer than 2 replicas of a service on spot nodes. K8s can evict all pods from one node simultaneously."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Swiggy, Zepto — companies doing real-time delivery with thousands of pods where compute costs are a major P&L concern.

---

## Q3. What is FinOps? What is your strategy to reduce cloud costs by 40%?

**"FinOps is the practice of bringing financial accountability to the variable spend model of cloud. It's about giving engineering teams visibility into their costs and making cost a first-class engineering metric alongside latency and reliability.**

**FinOps Maturity Model:**
1. **Crawl:** Tag every resource. Show each team their costs in Grafana/Cost Explorer
2. **Walk:** Rightsizing, reservations for predictable workloads
3. **Run:** Real-time cost anomaly alerting, per-PR cost estimation

**Typical 40% saving strategy:**

| Action | Typical Savings |
|--------|----------------|
| Right-size over-provisioned pods (VPA recommendations) | 15–20% |
| Move stateless workloads to Spot instances | 20–25% |
| Reserved Instances for databases, core nodes | 30–40% off those specific resources |
| Turn off dev/staging environments at night (K8s KEDA CronScaler) | 5–10% |
| Delete orphaned resources (unattached EBS, old snapshots, unused LBs) | 3–5% |
| Optimize data transfer cost (keep cross-AZ traffic internal, use VPC endpoints) | 2–5% |

**Tools I use:**
- **AWS Cost Explorer** — break down costs by tag (team:payments, env:prod)
- **Kubecost** — real-time cost per namespace, per pod, with savings recommendations
- **Infracost** — runs in CI/CD and comments on PRs with estimated infrastructure cost change
- **AWS Compute Optimizer** — cross-service rightsizing recommendations

**Organizational change:** Cost reviews in sprint retros. Each team owns a cost budget. Dashboard in Grafana showing current month spend vs budget visible to all engineers."

#### 🏢 Company Context
**Level:** 🔴 Principal | **Asked at:** CRED (cloud cost efficiency is a core engineering value), Flipkart at Staff+ levels. This question reveals technical depth AND business acumen.

---

## Q4. How does Karpenter differ from Cluster Autoscaler? When would you use each?

**"Both scale Kubernetes nodes, but with fundamentally different philosophies.**

| | Cluster Autoscaler | Karpenter |
|--|------------------|-----------|
| **How it works** | Watches unschedulable pods, scales predefined Node Groups | Watches unschedulable pods, provisions ANY EC2 instance type that fits |
| **Node types** | Fixed — you define which instance types in Node Groups | Dynamic — picks the cheapest available instance that satisfies pod requirements |
| **Spot handling** | Manual — you configure multiple node groups for different instance types | Automatic — falls back to next-cheapest spot or on-demand if spot unavailable |
| **Scale-down** | Slow — 10-minute cool-down by default | Fast — consolidates nodes, can bin-pack and remove underutilized nodes in minutes |
| **Setup complexity** | Higher — manage multiple node group configs | Lower — one Provisioner CRD |

**When to use Karpenter:**
- AWS-native workloads on EKS (Karpenter is AWS-developed)
- When you want maximum cost efficiency — Karpenter continuously consolidates nodes
- When your workload has heterogeneous resource needs (some pods need GPU, some need high-memory)

**When to use Cluster Autoscaler:**
- Multi-cloud or on-prem (Karpenter only supports AWS/Azure)
- When you need predictable node types for compliance reasons"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Zepto, Swiggy — rapidly scaling companies where cluster size can triple in hours during high-traffic events (midnight discount sales, IPL cricket matches).

---

## Q5. How do you optimize inter-service network costs?

**"In AWS, data transfer between Availability Zones (AZs) costs $0.01/GB. At scale (Swiggy processes millions of orders), cross-AZ traffic alone can be a significant cost.**

**Topology Aware Hints (K8s 1.21+):**
```yaml
# Service that prefers to route within the same AZ
apiVersion: v1
kind: Service
metadata:
  name: payment-service
  annotations:
    service.kubernetes.io/topology-mode: Auto
spec:
  selector:
    app: payment-service
  ports:
  - port: 8080
```
K8s will prefer routing to pods in the same AZ. Reduces cross-AZ latency AND cost.

**Other network cost optimizations:**
- **VPC Endpoints** — S3 and DynamoDB traffic stays within AWS private network (free) vs going through the internet gateway ($0.09/GB egress)
- **Compression** — gzip API responses and Kafka messages; reduce payload sizes
- **CDN for static assets** — CloudFront caches and servers from edge; reduces origin data transfer to near zero
- **Keep heavy data processing in same region as storage** — never stream 100GB of S3 data to a different region; run EMR/Glue in the same region"
