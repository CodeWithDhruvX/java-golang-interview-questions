# ☁️ Cloud-Native Architecture — Questions 1–10

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon, Google, Flipkart, Zepto, Razorpay — any company running on cloud infrastructure

---

### 1. What is cloud-native architecture?

"Cloud-native architecture is a design approach that fully leverages **cloud computing models** — elasticity, on-demand provisioning, managed services, and pay-per-use — rather than lifting-and-shifting traditional on-premise designs to the cloud.

A cloud-native app is built to: **scale horizontally** on commodity servers, **fail gracefully** (assumes failures will happen and designs for them), **update continuously** with no downtime (rolling deployments, canary releases), and **configure through the environment** (not hardcoded config files).

The CNCF (Cloud Native Computing Foundation) defines it as: using microservices, containers, service meshes, immutable infrastructure, and declarative APIs. Kubernetes is the de facto cloud-native runtime."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company with cloud infrastructure

#### Indepth
Cloud-native vs cloud-enabled:
- **Cloud-enabled:** Traditional monolith running on a VM in the cloud. You moved it, but didn't redesign it. You're paying cloud prices for on-premise architecture.
- **Cloud-native:** Microservices in containers, orchestrated by Kubernetes, auto-scaling, health checks, service discovery — designed specifically to leverage cloud capabilities.

Key cloud-native principles:
1. **Containers (Docker):** Consistent environment from dev to prod
2. **Orchestration (Kubernetes):** Automate deployment, scaling, healing
3. **Microservices:** Independently deployable, aligned to business domains
4. **DevOps + CI/CD:** Code to production in minutes, not weeks
5. **Observability:** Metrics, logs, traces built-in from day 1
6. **12-factor methodology:** Configuration, statelessness, disposability

#### 🗣️ How to Explain in Interview
**Interviewer:** What is cloud-native architecture?
**Your Response:** "Cloud-native is a design philosophy focused on building applications that fully leverage the power of the cloud—specifically **elasticity, scale, and managed services**. It’s the opposite of a 'lift-and-shift' approach. 

In a cloud-native architecture, we move away from 'Pet' servers that we manually patch and toward **'Cattle'**—immutable, containerized instances that can be destroyed and recreated at any time. We focus on **statelessness**, so our application can scale horizontally in seconds, and we use **declarative APIs** like Kubernetes to manage the complexity. My goal is always to build a system that is 'plugged into' the cloud's native capabilities, making it much more resilient to hardware failures and traffic spikes."

---

### 2. What is containerization and how does Docker work?

"Containerization packages an application and all its dependencies into a **container** — a lightweight, portable, isolated unit that runs consistently across any environment.

Docker uses Linux kernel features: **namespaces** (isolate processes, network, file system — each container has its own view of the system) and **cgroups** (limit CPU, memory, and I/O resources for each container).

The result: the same Docker image runs identically on a developer's MacBook, a CI server, and a production Kubernetes cluster. Eliminates 'it works on my machine' — the container is the machine."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All companies using modern infrastructure

#### Indepth
Docker layered filesystem (UnionFS):
- Each `RUN`/`COPY`/`ADD` instruction in a Dockerfile creates a new layer
- Layers are cached — only changed layers are rebuilt. A `RUN apt-get install` layer is cached until the Dockerfile line above it changes.
- Final image = base layer + all instruction layers stacked
- Containers add a thin writable layer on top of the read-only image layers

Multi-stage builds (best practice):
```dockerfile
# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

# Run stage (minimal image)
FROM gcr.io/distroless/base
COPY --from=builder /app/server /server
CMD ["/server"]
# Final image: ~10MB vs 800MB with Go toolchain
```

Container vs VM: Containers share the host OS kernel (lightweight, millisecond startup). VMs have their own OS kernel (full isolation, second-level startup). For most microservices workloads, containers are sufficiently isolated and far more efficient.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is containerization and how does Docker work?
**Your Response:** "Containerization involves bundling an application with its entire runtime environment—every library, config file, and dependency it needs. I use Docker for this because it ensures **environmental consistency**. 

Architecturally, Docker uses Linux **namespaces and cgroups** to create an isolated 'sandbox' for the application. The big 'win' for a dev team is that it eliminates 'it works on my machine'—the image that runs on my laptop is the exact same binary and environment that runs on a multi-node production cluster. It also enables us to run dozens of microservices on a single server without them interfering with each other's dependencies, which is a massive boost for resource efficiency."

---

### 3. What is Kubernetes and what problems does it solve?

"Kubernetes (K8s) is an **open-source container orchestration platform** that automates the deployment, scaling, self-healing, and management of containerized applications.

The problems it solves: You have 100 containers across 50 services. How do you decide which host each container runs on? What happens when a container crashes? How do you roll out a new version without downtime? How do services find each other? How do you scale out when traffic spikes? Kubernetes answers all of these.

Core capabilities: **Scheduling** (places containers on appropriate nodes based on resource requests), **Self-healing** (restarts crashed containers, replaces unhealthy nodes), **Horizontal scaling** (HPA scales pods based on CPU/memory), **Service discovery + load balancing** (kube-proxy routes traffic to healthy pods), **Rolling deployments** (zero-downtime updates)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company running microservices

#### Indepth
Key Kubernetes objects:
- **Pod:** Smallest deployable unit. One or more containers sharing network + storage.
- **Deployment:** Manages Pod replicas. Handles rolling updates, rollbacks.
- **Service:** Stable network endpoint for a set of Pods. Types: ClusterIP (internal), NodePort (fixed port on node), LoadBalancer (cloud LB).
- **Ingress:** HTTP routing rules (path, host). Entry point from internet to services.
- **ConfigMap:** Non-sensitive configuration data (environment variables, config files).
- **Secret:** Sensitive data (passwords, API keys). Base64 encoded (not encrypted).
- **HPA (Horizontal Pod Autoscaler):** Scale pods based on CPU/memory/custom metrics.
- **PersistentVolume:** Storage that outlives pods (for stateful apps).

K8s control loop: desired state → actual state reconciliation. You declare "I want 5 replicas of service A". K8s continuously works to make the actual state match. If a pod dies, K8s creates a new one. If node fails, K8s reschedules pods on healthy nodes. This is declarative infrastructure.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is Kubernetes and what problems does it solve?
**Your Response:** "I like to think of Kubernetes as the **'operating system' for a distributed cluster**. When you have hundreds of microservices, you can’t manually decide where to put every container or how to route traffic to them. Kubernetes takes over that 'undifferentiated heavy lifting.'

It’s a **declarative system**: I define the 'desired state'—like 'I want 10 replicas of my checkout service'—and Kubernetes continuously monitors the cluster. If one pod crashes or a server dies, it automatically reschedules a new one to keep the promise I made in my configuration. It solves the hardest parts of distributed systems: service discovery, load balancing, automated rollouts, and self-healing."

---

### 4. What is serverless architecture and when to use it?

"Serverless is a cloud execution model where you **write functions, the cloud provider manages all infrastructure** — no servers to provision, patch, or scale. You're billed per invocation and execution time, not for idle servers.

AWS Lambda runs your function in response to events (HTTP request via API Gateway, S3 file upload, Kafka message, scheduled timer). The function executes, returns, and the runtime is torn down. You pay only for the 100ms execution, nothing when idle.

When serverless shines: Event-driven processing (image thumbnailing when a photo is uploaded), periodic jobs (daily report generation), lightweight APIs with variable traffic. When it doesn't: Long-running processes, stateful workloads, high-throughput systems where per-invocation overhead adds up."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** AWS-heavy companies, startup teams

#### Indepth
Serverless trade-offs:
| Aspect | Pros | Cons |
|--------|------|------|
| Scaling | Auto-scales to zero and to millions | Cold starts (100ms-3s latency spike) |
| Cost | Pay per use | Expensive at sustained high traffic |
| Operations | No server management | Limited debugging, less control |
| Vendor lock | Managed by cloud | Hard to migrate across providers |
| Execution | Up to 15 min (Lambda) | Not for long-running processes |

Cold start problem: When a Lambda hasn't been invoked recently (or scales beyond warm instances), the runtime must initialize (load code, start runtime, establish connections). Solutions: **Provisioned concurrency** (keep N instances always warm — costs money), **keep-alive pings** (invoke a no-op every 5 minutes), **minimize package size** (smaller ZIP = faster cold start), **use lightweight runtimes** (Go/Rust < Node.js < Python < Java).

**Serverless frameworks:** AWS SAM, Serverless Framework, Pulumi, CDK — abstract CloudFormation complexity and enable local development simulation.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is serverless architecture and when to use it?
**Your Response:** "Serverless, like **AWS Lambda**, is the ultimate 'pay-as-you-go' model. It's an architecture where you focus purely on the business logic—writing small, event-driven functions—and the cloud provider manages the entire underlying infrastructure, from scaling to OS patching.

It’s my go-to for **event-driven tasks** or sporadic workloads. For example, if we need to process an image upload or run a daily cleanup job, serverless is incredibly cost-efficient because you pay zero dollars when the code isn't running. The trade-off you have to manage is **'cold start' latency**—the time it takes for a new instance to boot up after being idle. So for high-traffic, latency-sensitive core APIs, I might stick with containers, but for everything else, serverless is a huge productivity booster."

---

### 5. What is infrastructure as code (IaC)?

"Infrastructure as Code is the practice of **managing and provisioning infrastructure through code** — version-controlled, repeatable, and reviewable — instead of manual processes or interactive configuration tools.

Instead of clicking through the AWS console to create a VPC, subnet, EC2 instances, and security groups, I write Terraform or CloudFormation code that declares the desired infrastructure state. Running `terraform apply` creates the exact same infrastructure every time, in any environment, with no human error.

IaC enables: reproducibility (dev/staging/prod environments are identical), auditability (infrastructure changes go through code review like any other code change), disaster recovery (rebuild from scratch in minutes), and automation (CI/CD pipeline creates and destroys environments automatically)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** DevOps/platform engineering roles, senior backend roles

#### Indepth
IaC tools comparison:
| Tool | Provider | Language | State Management |
|------|----------|----------|-----------------|
| Terraform | Cloud-agnostic | HCL | Remote state (S3, Terraform Cloud) |
| CloudFormation | AWS only | YAML/JSON | AWS-managed |
| CDK | Cloud-agnostic | TypeScript, Python, Java | Transforms to CFN |
| Pulumi | Cloud-agnostic | General purpose (TS, Python, Go) | Remote state |
| Ansible | Config management | YAML (declarative DSL) | Stateless (procedural) |
| Helm | Kubernetes | YAML templates | Chart versioning |

Terraform workflow:
1. `terraform init` → Download providers
2. `terraform plan` → Show diff between current and desired state (review before applying)
3. `terraform apply` → Apply changes to cloud
4. `terraform destroy` → Teardown all resources

**GitOps:** IaC + Git as the single source of truth. Any change to infrastructure goes through a PR. Merge to main → CI pipeline runs `terraform apply`. ArgoCD (for Kubernetes) continuously syncs cluster state to Git repo.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is infrastructure as code (IaC)?
**Your Response:** "Infrastructure as Code, or IaC, is about moving away from 'click-ops' in a cloud console and toward **version-controlled, repeatable definitions** of your environment. Using tools like **Terraform**, I can define my entire network, database, and cluster in a Git repo.

This has three major advantages: First, **consistency**—our staging and production environments are guaranteed to be identical. Second, **auditability**—infrastructure changes go through code review just like features. Third, **velocity**—we can spin up an entire disaster recovery environment in minutes with a single command. It turns infrastructure from a manual bottleneck into an automated pipeline."

---

### 6. What is a service registry and configuration management?

"Service registry is a dynamic directory of **all running service instances and their network locations** — used for service discovery. Configuration management handles **application settings** (feature flags, DB connection strings, API keys) across environments.

Service Registry tools: Consul (discovery + health check + KV store), etcd (Kubernetes's backbone), Eureka (Netflix's registry). Kubernetes replaces the need for a separate registry with its DNS-based discovery.

Configuration management: **Environment variables** are the 12-factor approach (simple, universal). **Config servers** like Spring Cloud Config or AWS AppConfig centralize config and push updates without redeployment. **Secrets managers** (AWS Secrets Manager, HashiCorp Vault) securely store and rotate credentials."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Platform engineering, DevOps-heavy roles

#### Indepth
Configuration hierarchy (from highest to lowest precedence):
1. Command-line flags
2. Environment variables
3. Config files (application.yaml, config.json)
4. Default values in code

**Feature flags vs config:** Configuration is application settings (DB host, cache TTL). Feature flags are boolean toggles that enable/disable features at runtime without redeployment. Tools: LaunchDarkly, Unleash, AWS AppConfig. Use cases: canary releases (enable feature for 5% of users), A/B testing, instant kill switches for buggy features.

**Dynamic configuration update:** Apps that support hot reload can pick up config changes without restarting. ConfigMap changes in Kubernetes can be automatically pushed to pods via projected volumes. AWS AppConfig supports gradual rollout of configuration changes with automatic rollback on error rate increase.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is a service registry and configuration management?
**Your Response:** "In a distributed system, services are constantly being created and destroyed. A **Service Registry** is the dynamic directory that keeps track of where everything is currently living. Without it, you'd be stuck hard-coding IP addresses, which is impossible at scale.

For configuration, I follow the **12-factor app** principles by using environment variables or dedicated secret managers. This allows us to keep sensitive data like API keys out of our source code. More importantly, it lets us change application behavior—like adjusting a timeout or toggling a feature flag—without needing to rebuild and redeploy the entire application. It’s about keeping the 'brain' of the app separate from its settings."

---

### 7. What is autoscaling and what types exist?

"Autoscaling is the ability to **automatically add or remove compute resources** based on demand — ensuring you have enough capacity during peaks and don't waste money during troughs.

Three types: **Horizontal (HPA in K8s)** — add/remove instances of a service. Most common for stateless services. **Vertical (VPA in K8s)** — increase/decrease CPU and memory limits of existing instances. Good for stateful services that can't scale horizontally. **Cluster autoscaler** — add/remove nodes from the Kubernetes cluster when pods can't be scheduled (scale the infrastructure itself).

My standard setup: HPA for stateless services (CPU > 70% → scale out), Cluster Autoscaler for the underlying node pool (pending pods → provision new nodes). Combined, this handles 10x traffic spikes automatically."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company running cloud infrastructure

#### Indepth
HPA configuration example:
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
spec:
  minReplicas: 2
  maxReplicas: 20
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

**KEDA (Kubernetes Event-Driven Autoscaling):** Extends HPA to scale based on external metrics: Kafka lag, SQS queue depth, custom metrics. If a Kafka consumer group has 100K unprocessed messages, KEDA can scale from 2 to 50 consumer pods automatically, then back to 2 when the lag is consumed. Essential for event-driven microservices.

Scale-down considerations: **Scale-down is slower than scale-up** (to avoid flutter — rapid scaling up and down). Default stabilization window: 5 minutes before scaling down. Pod disruption budgets (PDB) ensure a minimum number of pods remain available during scale-down.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is autoscaling and what types exist?
**Your Response:** "Autoscaling is the 'secret sauce' that makes cloud economics work. **Horizontal scaling** is the most common—adding more 'replicas' of a service when the load increases. I usually set this up based on CPU or memory utilization.

But for more complex systems, I like to scale based on **custom metrics**, like the number of messages sitting in a Kafka queue. If we have a sudden burst of data, Kubernetes can scale our consumers from 2 to 50 pods in seconds to chew through that lag, and then scale back down to zero when the work is done. It ensures we always have the capacity to serve our users without paying for idle servers in the middle of the night."

---

### 8. What is blue-green deployment vs canary deployment?

"Both are zero-downtime deployment strategies, but they serve different risk profiles:

**Blue-Green:** Run two identical environments (blue = current prod, green = new version). Deploy to green, test it, then switch all traffic from blue to green in one instant flip. Instant and safe — rollback is just switching traffic back to blue. Cost: you need 2x infrastructure during the switch.

**Canary:** Gradually shift traffic from the old version to the new one — start with 1%, then 5%, 25%, 50%, 100%. If the canary version shows elevated error rates at 5%, stop and rollback. Catch issues at small blast radius before they affect all users. Cost: more complex traffic management, need good observability to know when to proceed."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company with CI/CD pipelines

#### Indepth
| Aspect | Blue-Green | Canary |
|--------|------------|--------|
| Risk | High if green has bugs (100% instant) | Low (gradual exposure) |
| Rollback | Instant (flip back) | Gradual reversal |
| Cost | 2x infrastructure | Slightly more complex routing |
| Observability needs | Lower (quick sanity check) | High (must detect errors at 5% traffic) |
| Traffic control | DNS flip or LB switch | Weighted routing (Nginx, Istio, Argo Rollouts) |
| Use case | Predictable migration, DB schema changes | Risky feature releases, performance changes |

**Argo Rollouts** (Kubernetes): Automates canary deployments with automatic pause, analysis, and promotion. Configure success metrics (error rate < 1%, p99 latency < 200ms). If metrics pass at each stage (5% → 25% → 50% → 100%), deployment proceeds automatically. If metrics fail, instant rollback.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is blue-green deployment vs canary deployment?
**Your Response:** "These are both strategies for **zero-downtime deployments**, but they handle risk differently. **Blue-Green** is like having two identical versions of your whole infrastructure; you deploy to the new one, test it, and then instantly flip a toggle to send everyone there. It’s great for fast rollbacks but it’s expensive because you're paying for 2x servers during the switch.

**Canary Deployment** is more of a gradual 'soak test.' You send just 5% of users to the new version and watch the error logs. If the 'canary' survives (meaning the metrics look good), you move to 25%, 50%, and finally 100%. I prefer Canary for complex microservice environments because it limits the **'blast radius'**—if there's a subtle bug, only a handful of users are affected before we catch it."

---

### 9. What is observability in cloud-native systems?

"Observability is the ability to **understand the internal state of a system by examining its outputs**. In distributed systems, you can't know if the system is healthy without comprehensive telemetry.

The three pillars: **Metrics** (what is the system doing at a macro level — CPU, error rate, requests per second → Prometheus + Grafana), **Logs** (what happened, with context — structured JSON logs → ELK stack or Loki), **Traces** (how did a specific request flow through the system → Jaeger, Tempo).

Modern addition: **Profiles** (CPU/memory flamegraphs for performance debugging) and **Events** (significant occurrences like deployments, incidents).

Observability is what separates a debuggable production system from a guesswork one."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All companies with production systems

#### Indepth
OpenTelemetry (OTel) — the standard:
- Vendor-agnostic instrumentation SDK for metrics, traces, and logs
- Auto-instrumentation for HTTP clients, DB drivers, gRPC — zero code changes for common frameworks
- Collector: Receives OTel data from services, exports to backend (Prometheus, Jaeger, Datadog, etc.)

Prometheus + Grafana stack:
```
App exposes /metrics endpoint → Prometheus scrapes every 15s → Grafana queries Prometheus → Dashboard
                                                              → Alertmanager → PagerDuty on threshold breach
```

**SRE Golden Signals (Google):**
1. **Latency:** How long requests take (p50, p95, p99)
2. **Traffic:** Requests per second
3. **Errors:** Rate of failed requests (4xx, 5xx)
4. **Saturation:** How close to capacity (CPU %, memory %, disk %)

Alert on golden signal deviations. Example: Error rate > 1% → alert. p99 latency > 500ms → alert. This is more meaningful than alerting on raw CPU%.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is observability in cloud-native systems?
**Your Response:** "Observability is the next evolution beyond simple monitoring. While monitoring tells you *if* a system is down, observability helps you understand *why* it's behaving a certain way. I build this using the 'three pillars': **Metrics** for the high-level health (CPU, error rates), **Logs** for deep context when things go wrong, and **Distributed Tracing** to follow a single user's request as it jumps through 10 different microservices.

My goal is to monitor the **'Golden Signals'**—latency, traffic, errors, and saturation. By setting up dashboards in Grafana, we can see in real-time if a specific database query is slowing down or if a recent deployment caused an error spike, often before the customers even notice."

---

### 10. What is GitOps?

"GitOps is an operational framework where **Git is the single source of truth for both application code and infrastructure configuration**. All changes — to code, to Kubernetes manifests, to Terraform configs — go through Git PRs. The actual state of the system is continuously reconciled to match what's declared in Git.

Pull-based model (vs push-based CI/CD): A GitOps operator (ArgoCD, Flux) runs inside the cluster and continuously watches a Git repository. When it detects a difference between the repo and the cluster state, it applies the change. No external system needs access to the cluster to deploy — the cluster pulls changes itself.

Benefits: Every change is auditable (Git log), rollbacks are `git revert`, security is improved (no external push access needed), and disaster recovery is just reapplying the repo to a new cluster."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Platform engineering, SRE roles

#### Indepth
GitOps tools:
- **ArgoCD:** Kubernetes-native GitOps. Watches a Git repo, syncs Kubernetes manifests. UI shows drift (what's in Git vs what's in the cluster). Auto-sync or manual sync mode.
- **Flux:** GitOps toolkit. More configuration-driven than ArgoCD. Supports Helm, Kustomize.

GitOps workflow:
1. Developer pushes code to main branch
2. CI pipeline builds and pushes Docker image (tagged with commit SHA)
3. CI pipeline opens PR to GitOps repo: updates `image.tag: sha-abc123` in deployment manifest
4. Platform review + merge
5. ArgoCD detects Git change → applies new deployment to cluster
6. Kubernetes does rolling update

**Environment promotion:** dev → staging → prod is now just merging changes from one Git branch/directory to another (with approvals). No more "who deployed what to staging?" — Git history tells you everything.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is GitOps?
**Your Response:** "GitOps is an operational framework where **Git is the one and only source of truth** for your system's state. Instead of manually running commands to deploy, you push your Kubernetes manifests or Terraform code to a Git repo.

A tool like **ArgoCD** runs inside the cluster and acts as a 'sync loop.' It watches the Git repo and the live cluster, and if it detects any **'drift'**—like someone manually changing a setting—it automatically reverts it to match the code. It makes rollbacks as simple as a `git revert` and ensures that our production environment is 100% reproducible and auditable."
