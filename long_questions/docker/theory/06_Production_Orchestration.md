# ☁️ **Docker in Production & Orchestration (71–80)**

---

### 71. How do you deploy Docker containers in production?
"My production deployment process:

1. **Build & tag** with Git SHA in CI: `docker build -t myregistry/app:$GIT_SHA`
2. **Test, scan, push** the image
3. **Update the service** with the new image tag — via Kubernetes (`kubectl set image`), Docker Swarm (`docker service update --image`), or Compose (`docker compose up -d`)
4. **Health checks** verify the deployment succeeded
5. **Rollback capability** — previous image tag is kept in the registry

I never run `docker run` manually in production — infrastructure-as-code (Helm, Compose, Terraform) manages the desired state."

#### In-depth
The key pattern is **immutable deployments**: never modify a running container. Instead, replace it with a new container from a new image. This ensures reproducibility — if production fails, there's an exact image to investigate. Combined with Git SHA tags, you have full traceability from production container → registry image → source commit → team member who authored the change.

---

### 72. What is the difference between Docker Swarm and Kubernetes?
"Both orchestrate containers, but at very different scales of complexity.

**Docker Swarm**: simpler, native Docker, uses Docker Compose syntax, deploys in minutes, smaller feature set. Good for small-medium teams already using Docker. Built-in with Docker Engine — no extra installation.

**Kubernetes**: complex, massive ecosystem, requires dedicated ops expertise, richer features (auto-scaling, RBAC, CRDs, policy engines). The industry standard for production at scale.

My heuristic: Swarm for teams of 1-10 devs/simpler workloads, Kubernetes for larger organizations or complex requirements."

#### In-depth
Kubernetes won the container orchestration wars (vs Swarm, Mesos, Nomad) primarily due to cloud vendor adoption (EKS, GKE, AKS). But its complexity is real — many teams use managed Kubernetes services (EKS, GKE) specifically to offload the control plane management. Docker Swarm remains production-viable and dramatically simpler for teams that don't need Kubernetes' advanced features.

---

### 73. How do you monitor containers in production?
"I monitor at three levels:

**Container level**: CPU, memory, I/O per container — via `docker stats` or cAdvisor + Prometheus.
**Application level**: request rates, error rates, latency (RED metrics) — via APM tools (Datadog, New Relic) or OpenTelemetry.
**Infrastructure level**: host CPU, memory, disk — via node exporters.

Alert on: container crash loops, memory approaching limits (OOM risk), high CPU saturation, unhealthy health check failures."

#### In-depth
The **cAdvisor + Prometheus + Grafana** stack is the open-source standard for container monitoring. cAdvisor exposes container metrics (CPU, memory, network, filesystem) in Prometheus format. Grafana dashboards (pre-built Docker dashboards are available on grafana.com) visualize trends. Set up **recording rules** in Prometheus for expensive queries to ensure dashboard responsiveness at scale.

---

### 74. What tools are used for Docker monitoring?
"My production monitoring stack:

- **cAdvisor**: container-level metrics, native Prometheus support
- **Prometheus**: metrics collection and storage
- **Grafana**: visualization and alerting
- **Datadog**: all-in-one commercial option with excellent Docker auto-discovery
- **Portainer**: lightweight Docker management UI with basic monitoring
- **Netdata**: lightweight, real-time container metrics

For logs: **Grafana Loki** (log aggregation) + Promtail (log shipper from containers)."

#### In-depth
Datadog's Docker integration auto-discovers container labels for service naming, custom metrics, and log tagging. Labeling containers with `com.datadoghq.ad.check_names` enables automatic monitoring setup. Open-source alternative: the Prometheus ecosystem with **Alertmanager** for alerting (PagerDuty, Slack) — requires more setup but zero licensing cost.

---

### 75. What is container orchestration?
"Container orchestration is **automated management of containers at scale** — scheduling, scaling, networking, health, and lifecycle management.

Without orchestration: manually SSH into servers to run/stop containers, manual load balancing, no automatic recovery from failures.

With orchestration (Swarm/Kubernetes): declare desired state ('3 replicas of my API'), and the platform ensures that state continuously — restarting failed containers, rescheduling on failed nodes, scaling up under load, rolling deployments without downtime."

#### In-depth
Orchestration platforms implement a **control loop** (reconciliation loop): desired state is stored in the system's database (etcd for Kubernetes, Raft log for Swarm). Controllers continuously compare actual state to desired state and take corrective actions. This is the fundamental pattern behind all modern infrastructure automation — from containers to serverless to Kubernetes operators.

---

### 76. How do you perform zero-downtime deployments with Docker?
"Zero-downtime requires: **health checks**, **rolling updates**, and a **load balancer**.

In Swarm: `docker service update --image myapp:v2 --update-parallelism 1 --update-delay 10s myservice`. Swarm stops one task, starts it with the new image, waits for health check to pass, then moves to the next.

In Compose (no Swarm): run the new image alongside the old one, let a reverse proxy (Traefik, nginx) shift traffic, then remove the old container.

Health checks are the key — without them, Swarm may route traffic to a starting container."

#### In-depth
**Blue-green deployment** is an alternative zero-downtime strategy: bring up the new version (green) completely, switch the load balancer from old (blue) to new all at once, then tear down blue. Advantages: instant rollback (switch back to blue), no mixed-version traffic. Disadvantage: requires 2x resources during the cutover. Rolling updates are more resource-efficient but expose both versions simultaneously.

---

### 77. How do you handle configuration management in Docker?
"Configuration flows from external sources into containers at runtime — never baked into images.

**Simple**: environment variables passed via `-e` or `--env-file`.
**Files**: config files mounted via volumes or bind mounts.
**Production secrets**: Docker secrets (Swarm), or external secret managers (Vault, AWS SSM, GCP Secret Manager).
**Dynamic config**: config servers (Consul, etcd) polled by the app at startup or runtime.

The principle: one image artifact, behavior controlled by external config per environment."

#### In-depth
AWS Parameter Store + the `envsubst` pattern: at container startup, an entrypoint script fetches parameters from SSM (`aws ssm get-parameter`) and injects them as environment variables before starting the main process. This avoids storing secrets anywhere in the container image while keeping them centrally managed and audited in AWS.

---

### 78. What is container health check?
"A health check is a command Docker runs **inside the container on a schedule** to determine if it's healthy.

If the health check fails repeatedly (after `retries`), the container status becomes `unhealthy`. Orchestrators (Swarm, Kubernetes) use this signal to stop routing traffic and schedule a replacement.

In Docker, health status is: `starting` (during the initial delay), `healthy`, or `unhealthy`. You query it with `docker inspect --format='{{.State.Health.Status}}' container`."

#### In-depth
Health checks are distinct from liveness checks in Kubernetes (though they serve the same purpose). The key parameters: `--interval` (how often to check), `--timeout` (how long to wait for the command), `--retries` (failures before marking unhealthy), `--start-period` (grace period for slow-starting apps). Tuning these prevents false positives during startup and false negatives during transient issues.

---

### 79. How do you implement a health check in Dockerfile?
"I define it with the `HEALTHCHECK` instruction:

```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --start-period=30s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```

For apps without `curl`, I use a custom binary or the app's own health command:

```dockerfile
HEALTHCHECK --interval=15s CMD [\"./healthcheck\"]
```

The command should return exit code 0 for healthy, 1 for unhealthy."

#### In-depth
A good health endpoint (`/health` or `/healthz`) checks the application's dependencies: database connectivity, cache connectivity, circuit breaker state. A **deep health check** fails when any dependency fails — preventing traffic to a broken instance. A **shallow health check** only checks if the process is running — less useful for detecting application-level failures like broken DB connections.

---

### 80. How does Docker handle service discovery?
"Docker's embedded **DNS server** (127.0.0.11) handles service discovery.

Within a Docker network, containers can reach each other by **service name** (in Compose) or **container name**. DNS resolves service names to the current container IPs. For Swarm services with multiple replicas, the DNS returns a **Virtual IP (VIP)** that load-balances across all healthy tasks.

For more complex scenarios (cross-network, external services), I use service meshes (Consul Connect, Istio) or reverse proxies (Traefik) that integrate with Docker labels."

#### In-depth
Docker's VIP load balancing uses Linux IPVS under the hood. IPVS operates in kernel space, making it significantly faster than user-space proxies for high-throughput services. Docker supports an alternative `dnsrr` endpoint mode (DNS round-robin) that returns actual container IPs instead of a VIP — better for clients that do their own load balancing (like Consul-aware apps).

---
