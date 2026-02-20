# 🟡 Infrastructure, Cloud & DevOps Concepts — Questions 111–120

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Amazon (AWS roles and solutions architect), Google Cloud, Azure, DevOps/SRE roles at Swiggy, Meesho, Razorpay

---

### 111. What is cloud computing?
"Cloud computing is the delivery of **computing services — servers, storage, databases, networking, software, analytics, intelligence — over the internet ('the cloud')** on a pay-as-you-go basis.

Instead of buying and maintaining physical hardware in an on-premises data center, I provision a VM on AWS in minutes, use it for what I need, and pay only for what I consume. This fundamentally changes economics: a startup can access the same infrastructure as Netflix without any upfront capital expenditure.

The three delivery models: **IaaS** (raw VMs, storage — AWS EC2), **PaaS** (managed platform for deploying code — AWS Elastic Beanstalk, Google App Engine), **SaaS** (fully managed software — Google Workspace, Salesforce)."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** AWS/GCP/Azure certification discussions, any company migrating to cloud — TCS, Infosys cloud centers, Swiggy, Meesho

#### Indepth
Cloud computing benefits vs challenges:
- **Benefits:** Elasticity (scale in minutes), global reach (deploy in 25+ regions), no hardware procurement, pay-per-use, managed services (let AWS manage RDS, not you managing Postgres on EC2).
- **Challenges:** Vendor lock-in (AWS proprietary services like DynamoDB, SQS are hard to migrate away from), cost optimization (easy to over-provision and overpay), shared responsibility model (cloud provider secures infrastructure; you secure your applications and data).

**Shared Responsibility Model (AWS example):**
- AWS is responsible for: Hardware, physical data center, hypervisor, networking (fiber, routers), foundational cloud services
- Customer is responsible for: Guest OS patches, application security, network configuration (security groups, VPCs), data encryption, IAM policies, access management

**Multi-cloud strategy:** Companies like Spotify use both AWS and GCP to avoid vendor lock-in and leverage specific strengths. Challenge: higher operational complexity. Most companies are better served by going deep with one provider.

---

### 112. What is Kubernetes?
"Kubernetes (K8s) is an **open-source container orchestration platform** that automates the deployment, scaling, and management of containerized applications.

The core problem it solves: you have 100 Docker containers that need to run across a cluster of 10 servers. How do you schedule them? What happens when a container crashes? How do you update the containers without downtime? How do you scale from 10 to 100 containers on traffic spikes? K8s answers all of these.

Key abstractions: **Pod** (one or more containers that share network and storage, the smallest deployable unit), **Deployment** (manages a set of identical pods, handles rolling updates and rollbacks), **Service** (stable network endpoint for a set of pods), **ConfigMap/Secret** (configuration and sensitive data)."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Any company running microservices on K8s — Swiggy, Meesho, Razorpay, Zomato, Google, Amazon (EKS team)

#### Indepth
Kubernetes core components:
- **Control Plane:** kube-apiserver (all cluster API), etcd (all cluster state), kube-scheduler (decides which node runs which pod), kube-controller-manager (ensures desired state matches actual state — runs ReplicaSet controller, Deployment controller, etc.)
- **Node (worker):** kubelet (agent on each node, manages pods on that node), kube-proxy (manages iptables rules for Service networking), Container runtime (containerd or CRI-O, runs actual containers)

Rolling deployment workflow:
```
kubectl apply -f deployment.yaml
→ Deployment controller sees new desired state (v2 image)
→ Creates new ReplicaSet for v2
→ Gradually scales up v2 pods (maxSurge=1)
→ Gradually scales down v1 pods (maxUnavailable=0)
→ Health checks (readiness probe) must pass before routing traffic
→ Old ReplicaSet kept (enables rollback: kubectl rollout undo)
```

**K8s auto-scaling:**
- **HPA (Horizontal Pod Autoscaler):** Scales pod count based on CPU/memory/custom metrics. `kubectl autoscale deployment api --cpu-percent=70 --min=3 --max=50`.
- **VPA (Vertical Pod Autoscaler):** Adjusts CPU/memory *requests* per pod. Useful for rightsizing.
- **Cluster Autoscaler:** Adds/removes nodes from the cluster based on pod scheduling needs. Integrates with AWS Auto Scaling Groups.

---

### 113. What is Docker?
"Docker is a **containerization platform** that packages an application and all its dependencies (runtime, libraries, config) into a portable container — a standardized unit that runs identically anywhere.

The fundamental problem Docker solves: 'it works on my machine'. With Docker, the application runs in an isolated environment that's identical in development, staging, and production. The developer, CI/CD pipeline, and production server all use the exact same container image.

A Docker image is built from a `Dockerfile` — instructions to build the environment step by step. Images are layered. Common layers (Ubuntu base, Node.js runtime) are shared across images, saving disk space. You push images to a registry (Docker Hub, ECR) and pull them for deployment."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any modern engineering team — Swiggy, Razorpay, Flipkart, Uber engineering

#### Indepth
Docker vs Virtual Machines:
| Feature | VM | Container (Docker) |
|---|---|---|
| Isolation | Full OS virtualization (hypervisor) | Process-level (shared kernel) |
| Startup time | Minutes (boot full OS) | Milliseconds (process start) |
| Size | GBs (full OS image) | MBs (app + libs only) |
| Overhead | High (each VM has full OS) | Low (shared kernel) |
| Security | Strongest isolation | Process + namespace isolation |
| Use case | Full environment isolation | App packaging and deployment |

Docker architecture:
- **Docker daemon (dockerd):** Background service managing containers
- **Docker CLI:** User interface (`docker run`, `docker build`, `docker push`)
- **Docker Hub / ECR:** Container image registry
- **containerd:** OCI-compliant container runtime (what K8s uses under the hood)

Multi-stage build (key optimization):
```dockerfile
# Build stage — install compiler, build binary
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/server .

# Final stage — minimal runtime image (2MB scratch vs 800MB Go image)
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
```

---

### 114. What is CI/CD?
"CI/CD (Continuous Integration / Continuous Delivery) is the practice of automating the steps from code commit to production deployment.

**Continuous Integration (CI):** Every code commit automatically triggers a pipeline: code checkout → build → unit tests → integration tests → static analysis → security scan. If any step fails, the PR is blocked. This catches bugs at commit time, not weeks later in a big-bang release.

**Continuous Delivery (CD):** After a CI pipeline passes, the artifact (Docker image) is automatically deployed to staging. With full CD (Continuous Deployment), it's automatically deployed to production after passing automated validation. This is how Netflix deploys hundreds of times per day."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** DevOps interviews, engineering culture questions at Netflix, Razorpay, Swiggy, any modern product company

#### Indepth
CI/CD pipeline stages:
```
Code commit → Push to Git
              ↓
   [CI Pipeline - GitHub Actions / Jenkins / CircleCI]
   1. Build Docker image
   2. Run unit tests (must pass)
   3. Run integration tests (in Docker Compose)
   4. Static analysis (SonarQube, golangci-lint)
   5. Security scan (Snyk, Trivy for container CVEs)
   6. Push image to ECR with git SHA tag
              ↓
   [CD Pipeline - ArgoCD / Spinnaker / AWS CodeDeploy]
   7. Auto-deploy to staging (Kubernetes HelmRelease update)
   8. Run automated E2E tests on staging
   9. Canary deploy to 1% production (Flagger)
   10. Gradual promotion: 1% → 10% → 100%
```

Popular tools:
- **CI:** GitHub Actions, GitLab CI, Jenkins, CircleCI, Buildkite
- **CD:** ArgoCD (GitOps K8s), Spinnaker (Netflix's open-source CD), AWS CodeDeploy, Flux
- **Artifact registry:** Docker Hub, AWS ECR, Google Artifact Registry
- **Feature flags:** LaunchDarkly, Unleash (decouple feature activation from deployment)

DORA metrics (measuring CI/CD effectiveness):
- **Deployment frequency:** How often do you deploy to production?
- **Lead time for changes:** Commit to production in how long?
- **Change failure rate:** What % of deployments cause a production incident?
- **Mean time to restore (MTTR):** How long to recover from an incident?
Elite teams deploy multiple times/day with <1 hour lead time and <5% failure rate.

---

### 115. What is serverless computing?
"Serverless computing means deploying code **without managing servers** — you provide a function, the cloud provider runs it on demand and scales it automatically.

AWS Lambda is the canonical example: write a function `handler(event)`, deploy it, and it runs only when triggered (HTTP request, S3 event, Kafka message, scheduled timer). You pay only for actual execution time (measured in milliseconds). If no one calls it, you pay zero.

The appeal: no server provisioning, no auto-scaling configuration, no patching — just code. For event-driven workloads, image processing, webhook handlers, and scheduled batch jobs, Lambda is often the most cost-effective solution."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** AWS Solutions Architect discussions, startups, Swiggy backend (Lambda for notifications), CRED

#### Indepth
Serverless trade-offs:
- **Cold start:** First request to an idle function incurs a start-up delay (100ms-3s depending on runtime). Subsequent requests (within warm period) are instant. Go and Rust have fastest cold starts (~100ms). Java/JVM is slowest (~8s). Mitigated by provisioned concurrency (AWS keeps N instances warm — costs more).
- **Execution limits:** Lambda has 15-minute max execution time. Not suitable for long-running processes.
- **Stateless:** No persistent data between invocations. State must go to external store (DynamoDB, RDS, Redis).
- **Concurrency model:** Lambda scales by instantiating more instances per concurrent request — each instance handles one request at a time. 10,000 concurrent requests = 10,000 Lambda instances. Watch out for DB connection explosion (Lambda at 10K concurrency × 1 DB connection each = 10K connections — use RDS Proxy).

**When to use serverless:**
- Event-driven functions (S3 upload → resize image)
- Webhook handlers (Stripe webhook → process payment event)
- Scheduled batch jobs (nightly report generation)
- APIs with highly variable traffic (0 to peak and back)

**When NOT to use serverless:**
- Long-running computations (ML training, video transcoding)
- Low-latency APIs where cold starts are unacceptable
- High-throughput stateful services (maintain persistent connections)

---

### 116. What is Infrastructure as Code (IaC)?
"Infrastructure as Code means managing and provisioning infrastructure **through declarative configuration files** (code) rather than manual UI clicks or ad-hoc scripts.

Terraform is the most popular IaC tool: I declare 'I want an AWS VPC with these subnets, an EKS cluster with these node groups, an RDS instance with this configuration' — all in `.tf` files. Terraform plans the changes before applying and tracks current state. The same codebase can be applied to create dev, staging, and production environments identically.

The key benefit: infrastructure changes go through code review (PR), are versioned in git (rollback = git revert), and are reproducible. No more 'who made that change to production last Tuesday?'"

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** DevOps/Platform Engineering roles at Swiggy, Razorpay, Meesho, Amazon (AWS), any cloud-native company

#### Indepth
IaC ecosystem:
- **Terraform (HashiCorp):** Provider-agnostic, works with AWS/GCP/Azure and hundreds of others. Declarative (describe what you want, not how to get there). HCL (HashiCorp Configuration Language). State file tracks real infrastructure. Terraform plan shows diff before apply.
- **AWS CloudFormation / CDK:** AWS-specific IaC. CloudFormation: YAML/JSON. CDK: define infra in TypeScript/Python — compiles to CloudFormation.
- **Pulumi:** Like CDK but provider-agnostic. Write infra in TypeScript, Python, Go. Full programming language (loops, conditionals) vs declarative DSL.
- **Ansible:** Configuration management (not infra provisioning). Manages software installation, configuration of existing servers via SSH. Complementary to Terraform.
- **GitOps (ArgoCD/Flux):** K8s configuration as code in git. ArgoCD continuously syncs cluster state to git repository. Any drift is auto-corrected.

Terraform state management:
- State stored in S3 (remote backend) for team use — everyone runs `terraform apply` against same state file
- DynamoDB state locking — prevents concurrent `terraform apply` from conflicting
- Terraform workspaces — separate state per environment (dev/staging/prod)

---

### 117. What is a VPC?
"A VPC (Virtual Private Cloud) is a **logically isolated private network** within a cloud provider (AWS, GCP, Azure) where you launch your cloud resources.

Without a VPC, your EC2 instances would be open to the internet by default. A VPC gives you a private IP range (e.g., 10.0.0.0/16), and you control all networking: which subnets are public (accessible from internet), which are private (no internet access), what traffic is allowed between components (security groups, NACLs).

For a typical web application: load balancers and NAT gateways in public subnets (need internet access), application servers and databases in private subnets (only accessible from within VPC). Databases MUST be in private subnets — exposing RDS instances directly to the internet is a critical security failure."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** AWS Solutions Architect discussions, cloud security reviews — Razorpay, PhonePe, Swiggy

#### Indepth
VPC architecture for a standard web app:
```
Internet → Internet Gateway → VPC (10.0.0.0/16)
                              │
           ┌──────────────────┼──────────────────┐
           │                  │                  │
   Public Subnet A    Public Subnet B    ...
   (10.0.1.0/24)     (10.0.2.0/24)
   [Load Balancer]   [Load Balancer]
   [NAT Gateway]     [NAT Gateway]
           │
   Private Subnet A  Private Subnet B
   (10.0.11.0/24)    (10.0.12.0/24)
   [App Servers]      [App Servers]
           │
   Private Subnet C  Private Subnet D
   (10.0.21.0/24)    (10.0.22.0/24)
   [RDS Primary]      [RDS Standby]
```

VPC components:
- **Subnets:** Subdivisions of the VPC IP range. Each tied to one AZ.
- **Route tables:** Control where network traffic is directed (internet gateway, NAT, VPC peering).
- **Internet Gateway (IGW):** Allows resources in public subnets to communicate with the internet.
- **NAT Gateway:** Allows resources in private subnets to make outbound internet calls (software updates) without being reachable from the internet.
- **Security Groups:** Stateful firewall at the ENI (network interface) level. Allow rules only.
- **NACLs (Network ACL):** Stateless firewall at the subnet level. Allow and deny rules. First line of defense.
- **VPC Peering / Transit Gateway:** Connect multiple VPCs (different AWS accounts/regions) privately.

---

### 118. What is auto-scaling?
"Auto-scaling automatically adjusts the number of compute resources (EC2 instances, K8s pods, Lambda concurrency) based on current traffic or resource utilization — scaling out when load increases and scaling in when it drops.

For AWS EC2: define an Auto Scaling Group with min=2, max=50 instances. AWS CloudWatch alarm triggers when CPU > 70%: add 2 instances. Another alarm when CPU < 30%: remove 1 instance. This ensures the application always has enough capacity without over-provisioning 24/7 for peak load.

The key metric is choosing the right **scaling policy**: target tracking (maintain a metric at a target — e.g., keep ALB request count per target at 1000) is the recommended modern approach."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Amazon (core AWS service), Swiggy (lunch/dinner surge), Hotstar (match day surge), Flipkart (Big Billion Days sale)

#### Indepth
Auto-scaling types:
- **Reactive scaling:** Adds capacity after a metric threshold is breached (e.g., CPU > 70%). There's a lag (5-10 min) between traffic spike and new instances being ready. For sudden spikes, this is too slow.
- **Predictive scaling (AWS Predictive Scaling):** ML-based forecasting uses historical patterns to proactively scale *before* the spike. If every Monday at 9am traffic doubles, AWS launches instances at 8:45am. Hotstar uses predictive scaling before IPL matches.
- **Scheduled scaling:** Manually defined schedule. "Scale out to 50 instances every day from 7am-11am." Used for predictable patterns (lunch rush for Swiggy).

K8s HPA detailed:
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
spec:
  minReplicas: 3
  maxReplicas: 100
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70  # Scale when avg CPU > 70%
  - type: External
    external:
      metric:
        name: kafka_consumer_lag  # Custom metric from Kafka
      target:
        type: AverageValue
        averageValue: 1000  # Scale when consumer lag > 1000 per pod
```

**Scale-in protection:** During auto-scale-in (removing instances), ensure active requests complete before terminating. Use connection draining (ALB) and graceful shutdown (SIGTERM handler that drains in-flight requests before exiting).

---

### 119. How to handle large-scale data migration?
"Large-scale data migration is one of the riskiest operations in backend engineering — you're moving live data in a production system with zero tolerance for data loss or extended downtime.

My approach: **three-phase migration**. Phase 1 (Dual-write): make the application write to both old and new storage simultaneously. Old DB remains authoritative. Phase 2 (Backfill): migrate historical data in batches using rate-limited background jobs. Phase 3 (Cutover): validate data equivalence → switch reads to new DB → make new DB authoritative → stop writing to old → drain and decommission old DB.

This approach allows rollback at any phase: if Phase 3 reveals discrepancies, fall back to reading from the old DB. Never do a big-bang cutover."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Flipkart (migrated from Oracle to MySQL/Cassandra), Swiggy (DB migrations during scaling), PhonePe (NoSQL adoption)

#### Indepth
Detailed migration playbook:

**1. Pre-migration validation:**
- Define data equivalence tests: count rows, checksum key fields, sample & compare
- Set acceptable error threshold (0% for financial data, <0.01% for analytics)

**2. Backfill safely:**
- Rate-limit backfill to not overload DB: `LIMIT 1000 per batch, SLEEP 100ms between batches`
- Track progress with a checkpoint table: restart from checkpoint if interrupted
- Validate in parallel: after backfill writes, read from new DB and compare checksums

**3. Online Schema Changes (for MySQL):**
- `ALTER TABLE users ADD COLUMN age INT` — blocks table for hours on 500M row table → CATASTROPHIC for production
- Use `gh-ost` (GitHub Online Schema Change) or `pt-online-schema-change` (Percona): these create a shadow table, copy data in background, then atomically swap — zero downtime

**4. Blue-green DB cut-over:**
- Set application to read-only mode briefly (seconds)
- Wait for replication to catch up to 0 lag
- Switch connection string to new DB
- Resume read-write

---

### 120. What is disaster recovery (DR)?
"Disaster Recovery is the **process of recovering IT infrastructure and data after a catastrophic failure** — data center fire, regional AWS outage, ransomware, major data corruption.

DR planning revolves around two metrics: **RTO** (Recovery Time Objective — how many hours until the system is restored) and **RPO** (Recovery Point Objective — how much data is acceptable to lose, measured in time).

For a payment company: RPO=0 (zero data loss acceptable) and RTO=15 minutes. For a content blog: RPO=24 hours and RTO=4 hours. The tighter the requirements, the more expensive the DR infrastructure."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** BFSI companies (Razorpay, PhonePe), Amazon, Google — business continuity planning is regulated in financial services

#### Indepth
DR strategies (increasing cost and decreasing RTO/RPO):

1. **Backup and Restore (RPO: hours, RTO: hours):**
   - Daily/hourly DB backups to S3 Glacier. Manual restore on disaster.
   - Cheapest. Slowest recovery. Acceptable for non-critical systems.

2. **Pilot Light (RPO: minutes, RTO: <1 hour):**
   - Minimal version of the system always running in DR region (just the DB with replication, no app servers).
   - On disaster: launch EC2 instances from AMIs + point them to DR DB.
   - Cost: Very low (DB + storage only, no compute).

3. **Warm Standby (RPO: seconds-minutes, RTO: <15 minutes):**
   - Full system running in DR region at reduced capacity (e.g., 20% of production size).
   - On disaster: scale up DR region to 100%, switch DNS.
   - Cost: Moderate.

4. **Active-Active Multi-Region (RPO: near-zero, RTO: seconds):**
   - Full production load in both regions simultaneously.
   - On disaster: other region absorbs 100% of traffic (must be provisioned for 2x load).
   - Cost: Doubles infrastructure cost.
   - Used by: Netflix, Razorpay (payment must not go down), Amazon.com.

**RDS automated backups:** AWS RDS takes automated backups every 5 minutes (transaction log backups), enabling point-in-time recovery to any second within the retention period (up to 35 days). Backups to S3 in the same region + optionally replicated cross-region for DR.
