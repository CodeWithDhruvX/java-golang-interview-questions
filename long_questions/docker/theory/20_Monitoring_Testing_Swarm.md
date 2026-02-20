# 📊 **Docker Monitoring, Testing & Swarm Deep Dive (221–250)**

---

### 221. How do you monitor Docker containers with Prometheus?
"Deploy **cAdvisor** alongside your containers — it exports per-container metrics in Prometheus format.

```yaml
services:
  cadvisor:
    image: gcr.io/cadvisor/cadvisor:latest
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:ro
      - /sys:/sys:ro
      - /var/lib/docker/:/var/lib/docker:ro
    ports:
      - '8083:8080'
```

Key metrics: `container_cpu_usage_seconds_total`, `container_memory_working_set_bytes`, `container_network_receive_bytes_total`, `container_fs_reads_bytes_total`.

Add Grafana with the Docker Overview dashboard (ID: 893) for visualization."

#### In-depth
For production Docker monitoring: deploy the Prometheus + cAdvisor + Grafana stack as sidecar services or on a dedicated monitoring host. Set alert rules: `rate(container_cpu_usage_seconds_total[5m]) > 0.9 * container_spec_cpu_quota / container_spec_cpu_period` alerts when CPU is at 90% of limit. Add `container_oom_events_total > 0` for OOM kills. These two alerts catch the most impactful container issues.

---

### 222. How do you collect container metrics using cAdvisor?
"cAdvisor (Container Advisor) automatically discovers running Docker containers and collects resource usage metrics.

Access metrics at `http://localhost:8083/metrics` in Prometheus format.
Web UI at `http://localhost:8083` shows per-container real-time graphs.

Key panels available: CPU usage (cores), memory (RSS, cache, working set), network I/O (bytes/s), disk I/O (IOPS), file system usage.

For Kubernetes: cAdvisor is built into the kubelet — no separate deployment needed. Metrics exposed via `/api/v1/nodes/<node>/proxy/metrics/cadvisor`."

#### In-depth
cAdvisor collects metrics by reading from kernel interfaces: `/sys/fs/cgroup` for CPU and memory, `/proc/net/dev` for network, and `/dev/disk` for I/O. It needs privileged access to these paths — hence the volume mounts in the Compose file. On read-only container filesystems, cAdvisor itself needs to be excluded from the security policy. For multi-host environments, run cAdvisor on each Docker host and configure Prometheus federation or remote write.

---

### 223. What is the difference between health-based routing and plain load balancing?
"**Plain load balancing**: distributes traffic across all containers in rotation (round-robin, least connections). No regard for whether a container is ready or healthy.

**Health-based routing**: distributes traffic only to containers whose health checks are passing. An unhealthy container is automatically removed from the routing pool.

Docker Swarm + internal DNS VIP uses health-based routing: only **running and healthy** tasks receive traffic. A task whose health check fails is restart by Swarm and removed from DNS until it recovers."

#### In-depth
Health-based routing is essential for zero-downtime deployments. Without it: during a rolling update, the load balancer sends requests to newly started containers before they're ready (DB connections established, caches warmed), causing errors. With health-based routing: new containers enter the pool only after passing health checks. The `start_period` in health check config provides a grace period for slow-starting apps before failures count.

---

### 224. How do you set up log aggregation for multiple containers?
"Central logging architecture: containers → log driver → aggregator → storage.

**With Fluentd** (Compose sidecar):
```yaml
logging:
  driver: fluentd
  options:
    fluentd-address: localhost:24224
    tag: '{{.Name}}'
```

**With AWS CloudWatch**:
```yaml
logging:
  driver: awslogs
  options:
    awslogs-region: us-east-1
    awslogs-group: /docker/myapp
```

**With ELK (Elasticsearch + Logstash + Kibana)**: Filebeat agent on each host ships container logs from `/var/lib/docker/containers/` to Logstash/Elasticsearch."

#### In-depth
Choosing a log driver: `json-file` (default) is simplest but requires separate log shipping. `fluentd` and `awslogs` ship directly from the Docker daemon without intermediate files. The trade-off: driver-based shipping is faster but creates a hard dependency — if fluentd is down, containers fail to start (unless `fluentd-async=true` is set). File-based shipping (Filebeat reads json-file logs) is more resilient but adds an agent to manage.

---

### 225. How do you run integration tests using Docker Compose?
"Pattern: all services start, test runner service runs tests against them, Pipeline exits based on test runner's exit code.

```yaml
services:
  app:
    build: .
    depends_on:
      db:
        condition: service_healthy
  db:
    image: postgres:15
    healthcheck:
      test: ['CMD-SHELL', 'pg_isready']
      interval: 2s
      retries: 10
  test:
    image: mytest-runner
    depends_on:
      app:
        condition: service_healthy
    command: pytest tests/integration/
```

CI command: `docker compose up --abort-on-container-exit --exit-code-from test`"

#### In-depth
The `--abort-on-container-exit` flag stops the entire Compose stack when ANY service exits (including the test runner). Combined with `--exit-code-from test`, the exit code of the `test` service determines whether the CI step passes. After tests: `docker compose logs test > test-output.txt && docker compose down -v` — capture logs before cleanup, and use `-v` to remove volumes so the next test run starts clean.

---

### 226. How do you implement service mesh patterns with Docker?
"A service mesh manages service-to-service communication: mTLS encryption, traffic routing, retries, circuit breaking, observability.

**Lightweight approach with Consul + Envoy sidecars** in Swarm:
Each service gets an Envoy sidecar proxy. All inter-service traffic routes through proxies. mTLS certificates are managed by Consul.

**In Compose (dev mesh simulation)**:
```yaml
services:
  api:
    image: myapi
  api-proxy:
    image: envoyproxy/envoy
    volumes:
      - ./envoy.yaml:/etc/envoy/envoy.yaml
    depends_on: [api]
```"

#### In-depth
Full service mesh (Istio, Linkerd) is designed for Kubernetes and adds significant complexity to Docker Compose setups. For Docker Swarm or Compose, a practical alternative is **Traefik** as a reverse proxy with circuit breaker middleware, **Consul** for service discovery, and application-level retries. For development environments, simulating mesh-like features (mTLS, traffic splitting) requires significant manual configuration — most teams use a simpler Compose setup locally and only deploy the mesh in Kubernetes staging/production.

---

### 227. What is Docker Swarm's overlay network and how does it work?
"Swarm overlay networks enable containers on **different Docker hosts** to communicate as if they're on the same network.

Under the hood: VXLAN (Virtual Extensible LAN) tunnels. Each Swarm node creates a VXLAN endpoint. When a container sends a packet to another node's container, the sending node encapsulates the packet in a VXLAN UDP packet and sends it to the target node's VXLAN endpoint, which decapsulates and delivers it.

Keys: port 4789 (VXLAN UDP), port 7946 (Swarm gossip), port 2377 (cluster management)."

#### In-depth
VXLAN uses 24-bit VNI (VXLAN Network Identifier) — up to 16 million overlay networks. Each Swarm overlay network gets its own VNI. The overlay driver embeds an ARP cache (learning switches) so nodes don't need to broadcast ARP for every packet. Overlay network performance: approximately 10-15% overhead vs. bare metal networking due to encapsulation. For latency-sensitive workloads, consider placing tightly-coupled services on the same Swarm node using placement constraints.

---

### 228. What is the Raft consensus algorithm in Docker Swarm?
"Swarm manager nodes use **Raft consensus** to maintain a consistent cluster state. Raft ensures all managers agree on the current state (services, tasks, networks, secrets) even if some managers fail.

Key properties:
- **Leader election**: one manager is the leader — it handles all cluster changes. If the leader fails, remaining managers elect a new leader.
- **Quorum required**: majority of managers must be healthy to elect a leader and accept changes. `(N/2)+1` nodes needed.
- **Log replication**: leader replicates every state change to followers before committing

With 3 managers: can lose 1. With 5 managers: can lose 2."

#### In-depth
The quorum requirement has an important operational implication: **never deploy 2 or 4 managers**. With 2: a single failure creates a split-brain (both nodes think they should lead, neither has quorum). With 4: same quorum as 3 (lose 1), with double the overhead. Recommended: 3 managers for small clusters, 5 for production. If you need to reduce managers below quorum: `docker swarm update --autolock` and `docker node demote` carefully, or the cluster will become read-only.

---

### 229. How do you drain and remove a Swarm node?
"**Drain** a node before maintenance — tasks migrate off it:
```bash
docker node update --availability drain worker-1
# All tasks on worker-1 migrate to other nodes
# worker-1 stops receiving new tasks
```

**Verify** all tasks moved: `docker node ps worker-1`

**Remove after maintenance**:
```bash
docker node update --availability active worker-1  # Restore for normal ops
# OR remove permanently:
docker node rm worker-1  # On manager, after node left the swarm
# On the node itself:
docker swarm leave
```"

#### In-depth
The drain operation is graceful — running tasks are rescheduled, not killed immediately. The Swarm scheduler finds nodes with capacity, starts new tasks, waits for them to pass health checks, then the old tasks on the draining node stop. This enables zero-downtime node maintenance. For rolling OS updates across a Swarm cluster: drain node 1, update OS, reactivate, wait for all tasks healthy, then drain node 2, and so on.

---

### 230. How do you deploy a global service in Docker Swarm?
"Global mode runs **exactly one replica on every Swarm node** (that matches placement constraints).

```bash
docker service create \
  --mode global \
  --name log-collector \
  fluentd:latest
```

Or in Compose (v3+):
```yaml
services:
  log-collector:
    image: fluentd
    deploy:
      mode: global
```

Use cases: log collectors, monitoring agents (cAdvisor, Prometheus node exporter), security scanners — anything that should run on every host without count management."

#### In-depth
Global services automatically deploy to new nodes as they join the Swarm — unattended deployment of node-level infrastructure. When a node leaves/is removed, its global service task is simply not replaced (no replacement target). Global services with placement constraints: `--constraint node.role==worker` excludes managers from the deployment — useful for avoiding monitoring agents on management-only nodes. Combine with resource limits to ensure the agent doesn't impact workloads.

---

### 231. How do you use Docker Swarm configs?
"Swarm configs store **non-sensitive** configuration data in the Swarm's Raft store — available to all services without bind-mounts or environment variables.

```bash
docker config create nginx.conf ./nginx.conf
docker service create \
  --name nginx \
  --config source=nginx.conf,target=/etc/nginx/nginx.conf \
  nginx:latest
```

Update a config (create a new version, then update the service):
```bash
docker config create nginx.conf.v2 ./nginx-v2.conf
docker service update --config-rm nginx.conf --config-add source=nginx.conf.v2,target=/etc/nginx/nginx.conf nginx
```"

#### In-depth
Swarm configs are stored encrypted in the Raft log (same as secrets). They're mounted as tempfs inside the container, readable by the running service. The immutability of configs (you can't update a config's content, only create a new version) is a feature: it provides audit trail and rollback capability. For nginx, updating configs typically requires a service rolling restart — the nginx `-s reload` signal approach doesn't work automatically in Swarm without additional tooling.

---

### 232. What are placement constraints in Swarm?
"Placement constraints restrict which nodes a service's tasks can run on.

By node role: `--constraint node.role==manager` or `==worker`
By node label: `--constraint node.labels.gpu==true`
By node hostname: `--constraint node.hostname==worker-gpu-1`
By OS/arch: `--constraint node.platform.os==linux`

Example — GPU workloads only on GPU nodes:
```bash
docker node update --label-add gpu=true worker-gpu-1
docker service create --constraint node.labels.gpu==true --name ml-train myml:latest
```

Placement preferences (spread, not restrict): `--placement-pref spread=node.labels.zone`"

#### In-depth
Placement constraints are essential for: **hardware affinity** (GPU services on GPU nodes), **zone spreading** (place replicas in different failure domains), **compliance** (data sovereignty — some services must only run on specific-region nodes), and **resource tiering** (heavy services on large nodes, light services on small nodes). Combine constraints with resource limits (`--reserve-cpu`, `--reserve-memory`) for complete capacity planning in mixed-workload Swarm clusters.

---

### 233. How do you expose a Swarm service to external traffic?
"Option 1: **Published ports** (routing mesh):
```bash
docker service create --name web --publish 80:8080 --replicas 3 nginx
```
Any Swarm node's IP at port 80 routes to one of the 3 replicas (IPVS-based routing mesh).

Option 2: **Host mode publishing** (bypass routing mesh):
```bash
docker service create --publish mode=host,target=8080,published=80 nginx
```
Only the node running a replica serves traffic on port 80.

Option 3: **Ingress controller** (Traefik, NGINX): deploys as a global service, routes via overlay network using DNS service discovery."

#### In-depth
The Swarm routing mesh (`ingress` mode) is convenient but adds latency — a request to node A for a service running only on node B makes an extra hop via IPVS. Host mode eliminates this hop but requires an external load balancer to know which nodes are running replicas. For production: deploy Traefik as a global service (one per node), all traffic hits any node → Traefik, Traefik routes via Swarm overlay DNS to the service's VIP, IPVS load-balances across replicas on the overlay network.

---

### 234. How do you update Swarm services with zero downtime?
"Docker Swarm's built-in rolling update:
```bash
docker service update \
  --image myapp:v2.0 \
  --update-parallelism 1 \
  --update-delay 30s \
  --update-failure-action rollback \
  --update-monitor 60s \
  myservice
```

- `--update-parallelism 1`: update 1 task at a time
- `--update-delay 30s`: wait 30 seconds between each task update
- `--update-failure-action rollback`: if a task fails to start healthy, auto-rollback the entire service
- `--update-monitor 60s`: how long Swarm monitors each updated task before considering it healthy"

#### In-depth
The `--update-monitor` duration must be longer than your health check interval × retries. If health checks take 30s to pass (start_period=30s), set monitor to at least 60s. Otherwise Swarm declares a task healthy before the health check has a chance to fail — and proceeds to update the next task, potentially cascading a bad deployment. The `--update-failure-action rollback` is the safety net — always set it in production.

---

### 235. How do Swarm and Kubernetes differ in orchestration model?
"**Docker Swarm**: simpler, Docker-native, integrated with Docker CLI/Compose. Easier setup. Less features.

**Kubernetes**: more complex, broader ecosystem, cloud-native standard. Much more powerful.

Key differences:
| Feature | Swarm | Kubernetes |
|---------|-------|-----------|
| Setup complexity | Minutes | Hours/days |
| Scaling | Simple (`--replicas N`) | Advanced (HPA, VPA, KEDA) |
| Networking | Overlay + routing mesh | CNI plugins (Calico, Cilium, etc.) |
| Storage | Volume plugins | CSI drivers, PV/PVC |
| RBAC | Basic | Fine-grained |
| Ecosystem | Small | Massive (CNCF) |
| Learning curve | Low | High |"

#### In-depth
Swarm is appropriate for: small teams, simple deployments, quick setup, legacy systems. Kubernetes is appropriate for: production microservices at scale, complex networking, compliance requirements, multi-cloud portability. The industry has clearly moved to Kubernetes — most job postings, tooling, and investment are Kubernetes-focused. However, Docker Compose + Swarm is still a valid, productive choice for small organizations or teams just starting containerization.

---

### 236. How do you perform rolling restarts in Swarm?
"Trigger a rolling restart (no image change) using `--force`:
```bash
docker service update --force myservice
```

This recreates every task according to the current update configuration (parallelism, delay), achieving a controlled rolling restart even with no spec changes.

Or use `--image` with the same tag to force a repull and restart:
```bash
docker service update --image myapp:latest --force myservice
```

For scheduled restarts (e.g., weekly night restart): use a cron job running the above command."

#### In-depth
`docker service update --force` is the Swarm equivalent of Kubernetes' `kubectl rollout restart`. Use cases: apply updated Swarm configs/secrets (the service must restart to pick up new mounted configs), refresh TLS certs, or force restart after host kernel updates. The rolling nature ensures zero downtime — same parallelism and delay settings as normal updates apply.

---

### 237. How do you limit resources in Docker Swarm services?
"Resource limits and reservations in Swarm:

```bash
docker service create \
  --limit-cpu 0.5 \
  --limit-memory 512m \
  --reserve-cpu 0.25 \
  --reserve-memory 256m \
  myapp
```

Or in Compose:
```yaml
deploy:
  resources:
    limits:
      cpus: '0.5'
      memory: 512M
    reservations:
      cpus: '0.25'
      memory: 256M
```

**Limit**: hard cap — if the container uses more than `limit-memory`, it's OOM-killed.
**Reservation**: guaranteed resource — Swarm won't schedule this task on a node with less available than the reservation."

#### In-depth
The distinction between limits and reservations is important for capacity planning. Reservations are used by the Swarm scheduler for bin-packing decisions — it uses reservations (not actual usage) to determine if a node has capacity. If all containers are well below their limits (common), the node may have plenty of actual capacity while the scheduler thinks it's full based on reservations. Set reservations = expected average usage, limits = expected peak usage + buffer.

---

### 238. What are Swarm nodes and managers?
"A **Swarm node** is any Docker Engine that's part of the Swarm cluster.

**Manager nodes**: run the Swarm control plane — scheduling, state management via Raft consensus, API handling, task assignment. Also run workloads (unless you restrict with placement constraints).

**Worker nodes**: accept and run tasks assigned by managers. No cluster state or scheduling involved. Most nodes in a large cluster are workers.

Command overview:
```bash
docker swarm init  # Create cluster, current node becomes manager
docker swarm join-token worker  # Show join token for workers
docker swarm join-token manager  # Show join token for managers
docker node ls  # List all nodes and their roles
```"

#### In-depth
Security implication: manager nodes have full cluster control — they can execute tasks on any node and access all secrets. In production, dedicate manager nodes to cluster management (no workloads). Add placement constraint `node.role==worker` to all production services. This prevents a compromised workload container on a manager from accessing the Raft state or modifying cluster config. Limit manager node count to 3 or 5; more isn't more fault-tolerant (quorum math doesn't benefit from 7+ managers).

---

### 239. How do you use Docker Swarm secrets for production?
"Create a secret: `echo "my-db-password" | docker secret create db_password -`
Or from file: `docker secret create tls_cert ./server.crt`

Use in a service:
```bash
docker service create \
  --name myapp \
  --secret db_password \
  myapp:latest
# Secret available at /run/secrets/db_password inside container
```

In Compose:
```yaml
secrets:
  db_password:
    external: true
services:
  app:
    secrets:
      - db_password
```"

#### In-depth
Swarm secrets are encrypted at rest (in the Raft log with AES-256) and in transit (TLS). They're mounted as `tmpfs` files inside the container — in memory only, not on disk. Only services that explicitly reference a secret receive it. Secret rotation: create a new secret version (secrets are immutable), update the service to use the new secret and remove the old one — the service rolls to pick up the change. Secrets can be 500KB max — suitable for certificates, not for large binary blobs.

---

### 240. How do you enable Swarm mode and join nodes?
"Initialize on the first manager:
```bash
docker swarm init --advertise-addr <manager-ip>
```

This outputs join tokens. Copy and run the worker token on each worker:
```bash
docker swarm join --token SWMTKN-... <manager-ip>:2377
```

Add more managers:
```bash
# Get the manager join token:
docker swarm join-token manager
# Run the output command on each new manager node
```

Verify:
```bash
docker node ls  # Shows all nodes, roles, availability, status
```"

#### In-depth
`--advertise-addr` is critical for multi-host setups. It specifies which IP address other Swarm nodes should use to reach this manager. If your host has multiple interfaces (internal + external), set this to the internal/cluster network interface. Common mistake: advertising the public IP when nodes are on a private network — cluster traffic flows over the public internet unnecessarily. Use the internal VPC/private network IP for all Swarm communication.

---

### 241. How does Swarm service discovery work?
"Swarm provides DNS-based service discovery via an embedded DNS server on each node.

A service named `api` is reachable by all containers in the same overlay network at `http://api`. The DNS resolves to the service's **VIP** (Virtual IP — one IP per service). IPVS then load-balances connections to the VIP across all healthy task replicas.

For task-by-task resolution (not VIP): use `tasks.api` DNS name — returns A records for each individual task's IP. Useful for stateful services (database replicas) where per-replica targeting matters."

#### In-depth
The VIP model is preferable for stateless services — client-side DNS caching isn't an issue because they connect to the VIP, and IPVS handles distribution server-side. For stateful services (Redis Cluster, PostgreSQL replication), use `tasks.api` to discover each replica directly. The IPVS table updates instantly when a task is added or removed — much faster than DNS TTL-based propagation, eliminating the stale DNS issue that affects external load balancers.

---

### 242. How would you design a Swarm cluster for HA?
"High-availability Swarm design:

**Manager tier**: 3 or 5 manager-only nodes (use placement constraints to keep workloads off). Distributes across availability zones/failure domains.

**Worker tier**: as many as needed, distributed across zones/racks.

**Key design decisions**:
- Managers in separate AZs (cloud) or server racks (on-prem)
- Automated node health monitoring (cloud provider health checks)
- Registry mirror on the cluster network (fast image pulls, no rate limiting)
- Shared storage (NFS, Ceph, GlusterFS) or cloud EFS for stateful volumes
- Ingress tier: Traefik as global service + external load balancer

Network: open ports 2377 (TCP), 7946 (TCP/UDP), 4789 (UDP) between all nodes."

#### In-depth
One often-overlooked HA aspect: the **image registry**. If Docker Hub is unreachable during scaling (rate limits, outage), image pulls fail and Swarm can't start new tasks. Solution: ECR/GCR private registry in the same network (no rate limits, low latency), or a self-hosted Nexus/Harbor registry on the cluster. In AWS: ECR is in the same region — pull requests rarely fail. Also disable `imagePullAlways` behavior for stable images — reference specific digests to avoid unnecessary pulls.

---

### 243. What are the best practices for Swarm service definitions?
"Key best practices:

1. **Name services semantically**: `api`, `frontend`, `worker` not `service1`
2. **Set resource limits**: every service should have `--limit-memory` and `--limit-cpu`
3. **Set restart policies**: `--restart-condition any --restart-delay 5s --restart-max-attempts 10`
4. **Use rolling updates with `--update-failure-action rollback`**
5. **Set health checks on all services**
6. **Use secrets for all sensitive config** — never environment variables
7. **Use named volumes** for stateful services
8. **Pin image versions**: never use `latest` in production Swarm services
9. **Use placement constraints** for hardware-specific workloads"

#### In-depth
Service definition as a script (not ad-hoc CLI): store service definitions in shell scripts or Ansible playbooks that are version-controlled. A `deploy.sh` that runs `docker service create` or `docker service update` with all parameters is reproducible and reviewable. Even better: use Docker Compose with `docker stack deploy` — this lets you define the entire stack as a YAML file that's version-controlled, reviewed, and deployed atomically.

---

### 244. How do you configure load balancing in Swarm?
"Swarm has two built-in load balancing mechanisms:

**Routing Mesh (ingress)**: external traffic reaches any node on the published port → IPVS routes to one of the service's replicas. Uses TCP/UDP load balancing at the network level.

**Internal VIP**: container-to-container traffic via service DNS resolves to the service VIP. IPVS distributes connections across replicas. Algorithm: round-robin by default.

Configuration for session affinity: not natively supported in Swarm VIP mode (use application-level stickiness or Traefik sticky sessions). For stateful services: use `tasks.myservice` DNS for client-side selection."

#### In-depth
IPVS (IP Virtual Server) operates at Layer 4 (TCP/UDP) — it distributes connections, not requests. For HTTP, one TCP connection can carry many requests. Long-lived connections (WebSocket, HTTP/2 keep-alive) may cause uneven distribution — one connection stays on one replica for its entire life. For HTTP load balancing at Layer 7 (request-level, header-based): use Traefik or NGINX with Docker provider, which understands HTTP semantics and distributes individual requests evenly.

---

### 245. How do you troubleshoot Swarm scheduling failures?
"Systematic approach:

1. `docker service ps myservice --no-trunc` — see task history, failure reasons
2. `docker service inspect myservice` — check placement constraints, resource requirements
3. `docker node ls` — check node availability and status
4. Check Node resources: `docker node inspect node1 --format '{{.Description.Resources}}'`
5. Compare service requirements vs. node capacity — scheduling fails silently if no node satisfies constraints

Common failure messages:
- 'No nodes available': all nodes drained or don't match constraints
- 'Insufficient resources': no node has enough CPU/memory based on reservations
- 'Task placement failure': constraint mismatch"

#### In-depth
The Swarm scheduler tries to schedule a task for 5 minutes after it fails (exponential backoff). Pending tasks show up as `Pending` in `docker service ps`. If needs are just barely not met: a node with 1.8GB free won't accept a service with `--reserve-memory 2G`. Fix: either add nodes, reduce reservations, or drain an underutilized service. The `docker service ps` `--no-trunc` flag is critical — the default truncates the failure reason, often cutting off the most informative part.

---

### 246. How do you scale Docker Compose services horizontally?
"Scale with `--scale`:
```bash
docker compose up -d --scale api=5
```

Or define defaults in Compose:
```yaml
services:
  api:
    image: myapi
    deploy:
      replicas: 3  # Compose v3 deploy section
```

Then: `docker compose up -d` starts 3 replicas.

**Important**: scaling works only when services don't use fixed host port mappings. With `ports: - '8080:80'`, only 1 instance can bind to host port 8080. For scaled services, either omit host port (expose only on internal network) or use host port ranges: `ports: - '8080-8082:80'` for 3 instances."

#### In-depth
Docker Compose horizontal scaling is for development/testing — not production. All replicas run on the same host, limiting actual capacity scaling. For production scaling: Docker Swarm (`docker service scale api=10` scales across the cluster) or Kubernetes (HPA automatically adjusts replicas). Use Compose scaling to test that your application behaves correctly with multiple instances before deploying to production — particularly testing that stateless behavior, session handling, and distributed locks work with multiple app instances.

---

### 247. What are the common mistakes in Docker Compose files?
"Most common mistakes I see in code reviews:

1. **Hardcoding secrets**: `POSTGRES_PASSWORD=mysecretpassword` in the Compose file
2. **Using `latest` tag**: unpredictable, cache-busting, no rollback
3. **Missing health checks**: services start before dependencies are ready
4. **Missing `depends_on`**: race conditions on startup
5. **Host port conflicts**: multiple services trying to use the same port
6. **No resource limits**: a runaway container can starve the entire host
7. **Missing `.dockerignore`**: huge build contexts slow all rebuilds
8. **Writable volumes on databases in production**: should use managed DB services
9. **No log rotation**: logs fill disk within days/weeks"

#### In-depth
The missing health check + depends_on failure mode is the most common cause of 'it works sometimes' issues. `depends_on` without a health condition only waits for the container to start (the process to fork), not for the service to be ready. A `postgres` container can take 2-5 seconds to fully initialize after the PID starts. Any app that tries to connect in those first seconds gets a connection refused. Fix: add a `pg_isready` health check and `depends_on: db: condition: service_healthy`.

---

### 248. How do you do rolling deployments in Docker Swarm?
"Rolling deployments update tasks gradually while maintaining availability.

Configure via update policy:
```bash
docker service update \
  --update-parallelism 2 \
  --update-delay 30s \
  --update-order start-first \
  --update-failure-action rollback \
  --image myapp:v2 \
  myservice
```

`--update-order start-first`: start new tasks BEFORE stopping old ones (requires available capacity). Ensures no capacity drop during update.

`--update-order stop-first` (default): stop old task first, then start new. Avoids extra resource usage but creates a brief capacity reduction."

#### In-depth
`start-first` vs `stop-first` is a critical choice. For mission-critical services, use `start-first` — briefly run `parallelism × 2` tasks during the update window, then drop to target count. For resource-constrained clusters with no spare capacity, `start-first` can fail because there's no room for extra tasks. Solution: over-provision capacity by 20% to support rolling updates, or use `stop-first` and accept brief capacity reduction during deployments (mitigated with small parallelism and sufficient replicas).

---

### 249. How do you set up multiple Docker Compose environments?
"Use the override file pattern:

```
docker-compose.yml          # Base: all services defined
docker-compose.override.yml # Auto-loaded in dev: volume mounts, debug ports
docker-compose.staging.yml  # Staging: staging-specific image tags
docker-compose.prod.yml     # Production: resource limits, restart policies
```

Dev (auto): `docker compose up` (loads base + override)
Staging: `docker compose -f docker-compose.yml -f docker-compose.staging.yml up`
Production: `docker compose -f docker-compose.yml -f docker-compose.prod.yml up`

Each override file only contains the DIFF from the base — DRY principle."

#### In-depth
Environment-specific configurations that typically differ:
- **Dev**: bind mounts for hot reload, debug ports exposed, development database with no auth
- **Staging**: production-like images from registry, staging-specific env vars, limited resource requests
- **Production**: resource limits, restart policies, no debug ports, secrets from secure store, logging drivers, health checks with real endpoints

Keep secrets out of all Compose files — inject at runtime via environment variables or Docker secrets. The Compose file is version-controlled and readable by developers — no secrets should ever be committed here.

---

### 250. How do you manage dependencies between services in Compose?
"Use `depends_on` with health check conditions:

```yaml
services:
  api:
    image: myapi
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_healthy
  db:
    image: postgres:15
    healthcheck:
      test: ['CMD-SHELL', 'pg_isready -U postgres']
      interval: 5s
      timeout: 5s
      retries: 5
  redis:
    image: redis:7-alpine
    healthcheck:
      test: ['CMD', 'redis-cli', 'ping']
      interval: 5s
      retries: 3
```

`api` won't start until both `db` and `redis` pass their health checks."

#### In-depth
For complex dependency graphs: Compose resolves `depends_on` chains automatically. `api depends on db, db depends on migration-job, migration-job depends on db` creates a valid execution order. But avoid circular dependencies — they prevent any service from starting. Application-side resilience is also important: even with Compose dependency management, in production (Swarm/K8s) services may restart independently. Build retry logic into connection code (pgx connection pool backoff, redis retry config) — don't rely solely on startup ordering.

---
