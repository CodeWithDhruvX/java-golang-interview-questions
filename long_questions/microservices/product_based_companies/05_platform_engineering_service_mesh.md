# 🔧 Microservices — Platform Engineering & Istio Service Mesh

> **Level:** 🔴 Senior – Principal
> **Asked at:** Google, Uber, Flipkart, Amazon, Swiggy, Razorpay (Platform/Infrastructure teams)

---

## Q1. What is Platform Engineering? How is it different from DevOps?

**"Platform Engineering is building an Internal Developer Platform (IDP) — a self-service layer that application teams use to deploy, manage, and observe their services, without needing to understand the underlying infrastructure.**

The core shift:

| DevOps | Platform Engineering |
|--------|---------------------|
| Embedding ops engineers in each team | Central platform team builds tooling for ALL teams |
| Each team writes their own CI/CD pipelines | Golden path: one opinionated pipeline that all teams use |
| Teams manage their own K8s YAML | Abstractions like Backstage or custom CLIs hide K8s complexity |
| Best practice is individual team's responsibility | Platform enforces best practices via policy-as-code (OPA Gatekeeper) |

**The 'Golden Path' concept:** You build one blessed, battle-tested way to deploy a microservice. The path includes auto-injected sidecar proxies, default HPA configs, Prometheus scraping, log shipping to ELK, and namespace-level network policies. A new team gets all of this for free just by following the golden path."

#### 🏢 Company Context
**Level:** 🔴 Principal | **Asked at:** Google SRE interviews, Uber's infrastructure platform teams, Flipkart for L6+ engineers. If you're interviewing for a Platform Engineer role, this is Q1.

#### Indepth
**Key tools in a modern IDP Stack:**
- **Backstage (Spotify):** Developer portal — service catalog, scaffolding new services from templates, documentation
- **ArgoCD:** GitOps-based continuous delivery — Git is the source of truth for K8s state
- **Crossplane:** Provision cloud resources (RDS, S3) using K8s CRDs from Git — no Terraform needed
- **OPA Gatekeeper:** Policy-as-code — prevents teams from deploying without resource limits, liveness probes, or with `latest` image tags
- **Tekton / GitHub Actions:** CI pipelines standardized across all teams

---

## Q2. What is a Service Mesh? Explain Istio's architecture.

**"A Service Mesh is a dedicated infrastructure layer that handles all service-to-service communication concerns — without changing a single line of application code.**

Every microservice gets a **sidecar proxy** (Envoy) automatically injected alongside it.  
All network traffic (inbound and outbound) passes through this proxy, not directly between services.

**What the mesh handles for free:**
- **mTLS** — every service connection is mutually authenticated and encrypted
- **Traffic management** — canary routing, circuit breaking, retries, timeouts
- **Observability** — golden signals (latency, errors, traffic, saturation) for every service pair
- **Policy** — rate limits, access control between services

**Istio Architecture:**

```
┌─────────────────────────────────────────────────────┐
│  CONTROL PLANE: istiod                              │
│  ┌────────────┐ ┌─────────────┐ ┌───────────────┐  │
│  │   Pilot    │ │   Citadel   │ │    Galley     │  │
│  │ (routing   │ │ (cert mgmt/ │ │ (config       │  │
│  │  config)   │ │ mTLS certs) │ │  validation)  │  │
│  └────────────┘ └─────────────┘ └───────────────┘  │
└─────────────────────────────────────────────────────┘
         │ xDS API (pushes config to all proxies)
         ▼
┌──────────────────────────┐   ┌──────────────────────────┐
│  Order Service Pod       │   │  Payment Service Pod     │
│  ┌────────────────────┐  │   │  ┌────────────────────┐  │
│  │   App Container    │  │   │  │   App Container    │  │
│  │   (Spring Boot)    │  │   │  │   (Spring Boot)    │  │
│  └────────────────────┘  │   │  └────────────────────┘  │
│  ┌────────────────────┐  │   │  ┌────────────────────┐  │
│  │  Envoy Sidecar     │◄─┼───┼─►│  Envoy Sidecar     │  │
│  │  (istio-proxy)     │  │   │  │  (istio-proxy)     │  │
│  └────────────────────┘  │   │  └────────────────────┘  │
└──────────────────────────┘   └──────────────────────────┘
         mTLS encrypted communication
```"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Swiggy, Razorpay — companies that run large Kubernetes clusters where managing mTLS and observability at the application level becomes an operationally unsustainable nightmare.

#### Indepth
**Istio vs Linkerd vs Consul:**
| Feature | Istio | Linkerd | Consul Connect |
|---------|-------|---------|----------------|
| Proxy | Envoy (powerful, complex) | Linkerd2-proxy (Rust, lightweight) | Envoy |
| Learning Curve | High | Low | Medium |
| Resource overhead | Higher (~50MB/sidecar) | Lower (~10MB/sidecar) | Medium |
| Multi-cluster | Excellent | Good | Excellent |
| Best for | Feature-rich, large orgs | Simplicity, small-medium orgs | Multi-cloud/multi-DC |

---

## Q3. How do you implement Canary Deployments using Istio VirtualService?

**"Without Istio, canary deployment means running two Deployments (stable + canary) and carefully controlling replica ratios. With Istio, you can split traffic with percent precision regardless of pod count, using VirtualService:**

```yaml
# Step 1: Two Kubernetes Deployments with different labels
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-service-stable
spec:
  replicas: 3
  selector:
    matchLabels:
      app: payment-service
      version: stable
  template:
    metadata:
      labels:
        app: payment-service
        version: stable
    spec:
      containers:
      - name: payment-service
        image: payment-service:v1.0

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-service-canary
spec:
  replicas: 1
  selector:
    matchLabels:
      app: payment-service
      version: canary
  template:
    metadata:
      labels:
        app: payment-service
        version: canary
    spec:
      containers:
      - name: payment-service
        image: payment-service:v1.1  # new version

---
# Step 2: Istio DestinationRule — defines subsets
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: payment-service-dr
spec:
  host: payment-service
  subsets:
  - name: stable
    labels:
      version: stable
  - name: canary
    labels:
      version: canary

---
# Step 3: VirtualService — controls traffic split
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: payment-service-vs
spec:
  hosts:
  - payment-service
  http:
  - route:
    - destination:
        host: payment-service
        subset: stable
      weight: 90
    - destination:
        host: payment-service
        subset: canary
      weight: 10  # 10% traffic goes to v1.1
```

**Progression:** Start at 10%, watch error rates in Grafana → move to 50% → 100% → delete stable.  
**Instant rollback:** Change `weight: 0` for canary. No pod restarts needed."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Uber, Razorpay — companies doing multiple production deployments per day where a bad release must be caught and rolled back within minutes.

---

## Q4. How does Istio implement mTLS? What is SPIFFE/SPIRE?

**"mTLS (Mutual TLS) means BOTH sides of a connection authenticate each other — the server proves it's the real PaymentService, and the client proves it's the real OrderService. Without mTLS, a rogue pod inside your cluster could impersonate any service.**

**Istio's mTLS flow:**
1. `istiod` (Citadel component) acts as a Certificate Authority (CA)
2. When a pod starts, its Envoy sidecar sends a Certificate Signing Request (CSR) to `istiod`
3. `istiod` issues a short-lived X.509 certificate encoding the SPIFFE identity: `spiffe://cluster.local/ns/default/sa/payment-service`
4. When OrderService calls PaymentService, the two Envoy sidecars perform the mTLS handshake — both verify each other's certificates
5. Certificates rotate automatically every 24 hours (configurable)

**SPIFFE (Secure Production Identity Framework For Everyone):** A spec that defines a universal identity format for workloads. SPIRE is the implementation.  
`spiffe://trust-domain/path/to/workload` — a workload identity that works across Kubernetes, VMs, and cloud providers.

**PeerAuthentication policy — enforce strict mTLS:**
```yaml
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: production
spec:
  mtls:
    mode: STRICT  # Reject any non-mTLS traffic
```"

#### 🏢 Company Context
**Level:** 🔴 Senior/Principal | **Asked at:** Razorpay, PhonePe, and any fintech where PCI-DSS compliance mandates encryption of all internal traffic. This is a must-know for companies handling payment data.

---

## Q5. What is an Internal Developer Platform (IDP)? Design one.

**System Design Prompt:** *"Your company has 200 engineers and 50 microservices. Design an Internal Developer Platform that lets a new engineer deploy their first service within 1 hour without knowing Kubernetes."*

**"I'd design the platform around a self-service portal backed by a GitOps pipeline:**

```
Developer Experience Layer:
┌─────────────────────────────────────────────────┐
│  Backstage Portal (React UI)                    │
│  • Service Catalog: browse all 50 services      │
│  • Scaffolder: create new service from template │
│  • TechDocs: auto-rendered docs from repo       │
│  • Deploy button → triggers GitOps workflow     │
└─────────────────────────────────────────────────┘
              │
              ▼  (Backstage creates a GitHub PR)
┌─────────────────────────────────────────────────┐
│  GitOps Layer (ArgoCD)                          │
│  • Watches the config Git repo                  │
│  • Any merged PR = deploy to Kubernetes         │
│  • Drift detection: K8s state must match Git    │
└─────────────────────────────────────────────────┘
              │
              ▼
┌─────────────────────────────────────────────────┐
│  Policy Layer (OPA Gatekeeper)                  │
│  • All K8s resources validated against policies │
│  • Blocks: missing resource limits, latest tag  │
│  • Enforces: liveness probes, pod disruption    │
└─────────────────────────────────────────────────┘
              │
              ▼
┌─────────────────────────────────────────────────┐
│  Infrastructure Layer (Crossplane)              │
│  • Provision RDS via K8s CRD (no Terraform!)    │
│  • S3 bucket → K8s object → real AWS resource  │
└─────────────────────────────────────────────────┘
```

**The 'Golden Path' new service gets automatically:**
- Namespace with NetworkPolicy (default: only allow ingress from API gateway)
- Resource requests/limits pre-configured  
- Prometheus ServiceMonitor (auto-scraped metrics)
- Fluentbit DaemonSet (auto-shipped logs to ELK)
- HPA with default min=2, max=10 replicas
- Istio sidecar injection enabled for mTLS

**A new engineer clones the template, writes business logic, merges a PR — and everything above is auto-provisioned.**"

#### 🏢 Company Context
**Level:** 🔴 Principal | **Asked at:** Google and Uber's platform engineering interviews. This design question reveals whether you think product-first (developer experience) vs ops-first (technical correctness).
