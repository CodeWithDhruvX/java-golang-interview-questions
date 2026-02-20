# 🌐 **Docker Advanced Networking Internals (351–380)**

---

### 351. How does Docker's iptables integration work?
"Docker automatically manages iptables to handle container networking. When you publish a port or create a network, Docker adds iptables rules.

Key chains Docker creates:
- `DOCKER`: rules for published port NAT
- `DOCKER-USER`: user-defined rules (run before DOCKER rules)
- `DOCKER-ISOLATION-STAGE-1/2`: prevent inter-network traffic

Example rules for `docker run -p 8080:80`:
```
-A DOCKER -d 172.17.0.2/32 ! -i docker0 -o docker0 -p tcp -m tcp --dport 80 -j ACCEPT
-A DOCKER-ISOLATION-STAGE-1 -i docker0 ! -o docker0 -j DOCKER-ISOLATION-STAGE-2
```

`DOCKER-USER` chain runs first — add custom firewall rules there, never edit DOCKER chain directly."

#### In-depth
Critical operational detail: Docker manipulates iptables at runtime. If you run `iptables -F` (flush all), you'll break all container networking until Docker daemon restarts. In production, never flush iptables without accounting for Docker's rules. For UFW (Ubuntu Firewall) users: disable UFW's management of iptables forwarding rules (`DEFAULT_FORWARD_POLICY="ACCEPT"` in `/etc/default/ufw`) or UFW will block container traffic. The DOCKER-USER chain is the safe place for custom rules — it persists across Docker daemon restarts.

---

### 352. What is IPVS and how does Docker Swarm use it?
"IPVS (IP Virtual Server) is a Linux kernel-level load balancer used by Docker Swarm for service VIP (Virtual IP) load balancing.

When a Swarm service is created, it gets a VIP (a virtual IP address). IPVS maintains a table mapping the VIP to the real task IPs (replicas). Packets destined for the VIP are rewritten by IPVS to a selected backend IP using the chosen algorithm (default: round-robin).

This all happens in the kernel — faster than userspace proxies like Nginx or HAProxy.

Check IPVS table: `ipvsadm -Ln` (run on a Swarm node)."

#### In-depth
IPVS supports multiple load balancing algorithms: round-robin (`rr`), weighted round-robin (`wrr`), least connections (`lc`), weighted least connections (`wlc`), source hashing (`sh`) for session affinity. Swarm uses `rr` by default. For stateful services requiring session affinity, use `sh` (source IP hashing directs the same client IP to the same backend). IPVS operates at Layer 4 (transport) — it doesn't inspect HTTP headers. For HTTP-level routing (path, header-based), you need a Layer 7 proxy like Traefik or Nginx.

---

### 353. How does Docker implement inter-container DNS resolution?
"Docker's embedded DNS resolver runs at `127.0.0.11:53` inside each container (for user-defined networks).

When a container queries `http://api`, the resolver:
1. Checks if `api` is a container/service name on the network
2. Returns the container's IP (bridge network) or the service VIP (Swarm overlay)
3. Falls back to host's `/etc/resolv.conf` for external lookups

For Swarm services: `tasks.api` resolves to all task IPs (A records per replica), `api` resolves to the service VIP.

The bridge networks use Docker's resolver; the host network bypasses it entirely."

#### In-depth
The embedded DNS resolver is what makes `depends_on` service-name communication work in Compose — containers can reference each other by service name because the DNS resolver automatically maps service names to container IPs. The resolver is also responsible for hot-updating when container IPs change (container restart gets a new IP — updated in DNS immediately). This avoids the DNS TTL staleness problem. For debugging DNS: `docker run --rm --network mynet busybox nslookup api` verifies DNS resolution inside the network context.

---

### 354. What is MacVLAN networking in Docker?
"MacVLAN assigns each container its own MAC address and IP on the physical network — containers appear as real network devices to the external network (switches, routers).

```bash
docker network create -d macvlan \
  --subnet=192.168.1.0/24 \
  --gateway=192.168.1.1 \
  --ip-range=192.168.1.128/25 \
  -o parent=eth0 \
  mymacvlan

docker run --network mymacvlan --ip 192.168.1.130 nginx
```

The nginx container is reachable at `192.168.1.130` from any host on the local network — no port publishing needed."

#### In-depth
MacVLAN bypasses the Linux bridge entirely — packets go directly from the container's virtual NIC to the physical NIC at line speed. This makes it ideal for: network appliances (containers that need to be reachable by IP on the LAN), legacy apps that use hardcoded IPs, or environments where network admins assign specific IPs to containers. Limitation: containers on MacVLAN cannot communicate with the host directly (the host's physical NIC can't talk to its MacVLAN children). Use IPvlan mode 2 (L3) if host-to-container communication is needed.

---

### 355. What are Docker network namespaces?
"Each container gets its own network namespace — an isolated network stack including: network interfaces, IP addresses, routing tables, iptables rules, and socket tables.

The host OS has its own network namespace. When a container starts, Docker creates a new network namespace, creates a veth pair (virtual ethernet pair — two connected virtual NICs), puts one end in the container's namespace and the other on the Docker bridge (in the host namespace).

The container sees only its own network interface (`eth0`) and loopback — completely isolated from other containers' network stacks."

#### In-depth
Network namespaces are the foundation of container networking isolation. `ip netns` command manages Linux network namespaces. Docker-created namespaces are in `/var/run/docker/netns/` (not the standard `/var/run/netns/` path). To inspect a container's network namespace: `nsenter --net=/var/run/docker/netns/<container-ns-id> ip addr` — runs `ip addr` inside the container's network namespace without entering the container. This is useful for debugging networking issues at the kernel level.

---

### 356. How do you create custom Docker network plugins?
"Docker's CNM (Container Network Model) has a plugin API. Custom network plugins implement:

- `NetworkCreate/Delete`: create/delete network resources
- `EndpointCreate/Delete`: add/remove a container from the network
- `Join/Leave`: connect container's network namespace to the plugin's network
- `DiscoverNew/Delete`: service discovery events

Implement via REST API over a Unix socket:
```
POST /NetworkDriver.CreateNetwork
POST /NetworkDriver.CreateEndpoint
POST /NetworkDriver.Join
```

Example plugins: Weave Net, Calico (as Docker plugin), Project Flannel."

#### In-depth
Building a custom network plugin gives you full control over container networking: custom IP addressing schemes, integration with existing SDN infrastructure, policy-based routing specific to your environment. The plugin runs as a daemon and registers with Docker via the plugin manifest. In practice, most custom networking needs are satisfied by existing plugins (Calico, Weave, Cilium) or by Docker's built-in drivers (bridge, overlay, macvlan). Writing a custom plugin is reserved for organizations with very specific networking requirements (custom SDN controllers, integration with proprietary network hardware).

---

### 357. What is the difference between Docker's host and bridge networking modes?
"**Bridge networking** (default): container gets its own network namespace. Docker creates a virtual bridge (`docker0`), assigns the container a private IP on the bridge subnet (172.17.0.0/16). Traffic to the outside world is NAT'd through the host's IP. Published ports create iptables rules.

**Host networking** (`--network=host`): container shares the host's network namespace directly. No separate network stack — the container uses the host's network interfaces, IPs, and ports directly. No NAT overhead.

Host mode performance: best possible network performance (no overhead). Risk: container's port 80 IS host's port 80 — port conflicts possible."

#### In-depth
Host networking gives the best throughput for network-intensive workloads (Kafka brokers, DPDK applications). The performance difference vs bridge: 10-20% throughput improvement and significant latency reduction by eliminating virtual NIC, bridging, and NAT steps. Security trade-off: a container in host networking with a vulnerability can access all host ports and listen on host interfaces. Never use host networking for user-facing services or untrusted code. Appropriate for: monitoring agents (cAdvisor needs host network for k8s metrics), performance benchmarking, and specific HPC workloads.

---

### 358. How do you implement network segmentation with Docker?
"Network segmentation isolates groups of containers, limiting blast radius of a compromise.

**Multi-network pattern**:
```yaml
networks:
  frontend:   # Public-facing containers
  backend:    # Internal app services
  db:         # Database tier (most restricted)

services:
  nginx:
    networks: [frontend]
  api:
    networks: [frontend, backend]  # Bridge between tiers
  postgres:
    networks: [db]
  worker:
    networks: [backend, db]
```

Nginx cannot reach postgres directly (no shared network). Only api and worker can reach the db network."

#### In-depth
The correct model: network membership defines security perimeters. Containers on separate networks cannot communicate even on the same host — the bridge/overlay driver enforces this. An attacker who compromises the Nginx container cannot directly reach the database (Nginx has no route to the `db` network). Only services that explicitly need cross-tier communication (api, worker) have dual network membership. This minimizes lateral movement paths. Combine with read-only volume mounts and non-root containers for defense-in-depth.

---

### 359. What is the VXLAN encapsulation overhead in Docker Swarm?
"VXLAN (Virtual Extensible LAN) encapsulates Layer 2 frames in UDP packets for overlay networking.

Overhead per packet:
- Outer Ethernet header: 14 bytes
- Outer IP header: 20 bytes
- Outer UDP header: 8 bytes
- VXLAN header: 8 bytes
- Inner Ethernet header: 14 bytes
Total overhead: ~64 bytes per packet

On a 1500-byte MTU network: VXLAN reduces effective payload to ~1436 bytes. This increases packet count for large transfers and causes fragmentation for near-MTU packets."

#### In-depth
MTU misconfiguration is the #1 cause of mysterious networking issues in Docker Swarm overlay networks. If the physical network MTU is 1500 and containers send 1500-byte packets, the VXLAN-encapsulated packet becomes ~1564 bytes — exceeding the MTU and causing fragmentation or packet drops. Fix: set Docker's overlay MTU to 1450 (leaves room for VXLAN overhead) in `daemon.json`: `{"mtu": 1450}`. For AWS Elastic Fabric Adapter or Jumbo Frame networks (MTU 9000+), set the overlay MTU to 8950 — large payloads in a single packet dramatically improve throughput.

---

### 360. How does Docker handle IPv6 networking?
"Docker supports IPv6 with explicit configuration.

Enable in `daemon.json`:
```json
{
  "ipv6": true,
  "fixed-cidr-v6": "2001:db8:1::/64"
}
```

Create IPv6 network:
```bash
docker network create --ipv6 \
  --subnet 2001:db8:2::/64 \
  ipv6net
```

Containers on IPv6 networks get both IPv4 and IPv6 addresses (dual-stack). External IPv6 publishing:
```bash
docker run -p [::]:80:80 nginx  # Publish on IPv6 only
docker run -p 80:80 nginx       # Both IPv4 and IPv6 (if dual-stack enabled)
```"

#### In-depth
IPv6 in Docker is often configured incorrectly because Docker's default bridge only has IPv4. For full dual-stack: enable `ipv6: true` in daemon.json AND create your user-defined networks with `--ipv6` AND configure NDP (Neighbor Discovery Protocol) proxying if IPv6 routing across hosts is needed. In Kubernetes, IPv6 dual-stack requires the `--feature-gates=IPv6DualStack=true` flag on kubeadm (beta in K8s 1.21+, stable in 1.23+). Most cloud providers now offer IPv6 by default on EKS, AKS, and GKE clusters.

---

### 361. What is the `none` network mode in Docker?
"`--network=none` gives the container only a loopback interface — no external network connectivity at all.

```bash
docker run --network=none myimage
# Inside container: only lo (127.0.0.1) available
# Cannot reach the internet, other containers, or the host
```

Use cases:
1. **Maximum isolation**: data processing containers that should never have network access
2. **Security scanning**: scan potentially malicious content in a fully isolated container
3. **Build steps**: compilation steps that shouldn't phone home
4. **Testing network failure handling**: test application behavior with no network"

#### In-depth
`--network=none` is the strictest network isolation in Docker — stronger than a custom network (which still has DNS and peer connectivity) or even seccomp filtering (which can still make network syscalls). For true air-gap isolation of container execution, `none` combined with `--read-only` filesystem and `--cap-drop ALL` creates the most restrictive container environment possible without resorting to Kata Containers or gVisor. Use in CI for untrusted code execution analysis.

---

### 362. How do you debug Docker networking issues?
"Systematic Docker network debugging:

```bash
# 1. Check container's network config
docker inspect container_name | jq '.[0].NetworkSettings'

# 2. Test connectivity from inside the container
docker exec -it container_name ping other_container
docker exec -it container_name curl -v http://api:8080/health

# 3. Check if they're on the same network
docker network inspect mynetwork

# 4. Check DNS resolution
docker exec -it container_name nslookup api

# 5. Use netshoot (network troubleshooting container)
docker run --rm --network container:target_container \
  nicolaka/netshoot tcpdump -i eth0
```"

#### In-depth
`nicolaka/netshoot` is the Swiss army knife for container network debugging — it contains tcpdump, wireshark, ping, traceroute, netstat, ss, iperf3, nmap, and dozens of other networking tools. Running it with `--network container:TARGET` shares the TARGET container's network namespace, letting you capture traffic exactly as the target container sees it. This is invaluable for debugging TLS issues (see the actual packets), diagnosing connection refused vs timeout errors, and verifying service discovery DNS resolution from the exact container's perspective.

---

### 363. What is Calico and how does it work with Docker?
"Calico is a container networking plugin providing Layer 3 networking and network policy for Docker and Kubernetes.

Unlike overlay networks (which encapsulate packets), Calico uses direct IP routing — each container gets a routable IP and routes are distributed via BGP (Border Gateway Protocol).

```bash
# Install Calico CNI for Kubernetes
kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.27.0/manifests/calico.yaml
```

Benefits:
- No encapsulation overhead (VXLAN/IPIP optional for cross-subnet)
- Standard network debugging (packets look normal)
- Network policies with L7 support (HTTP method, path)
- Scales to thousands of nodes"

#### In-depth
Calico's BGP model treats each Kubernetes node as a router. Route reflectors (special BGP peers) distribute pod routes across all nodes. In a large cluster, every node knows the IP range of every other node's pods and routes packets directly — no VXLAN encapsulation, no overlay. The result: full line-rate networking with zero encapsulation overhead. For clusters spanning multiple AZs or subnets (where direct routing isn't possible), Calico automatically falls back to IP-in-IP encapsulation. This hybrid mode provides the best possible performance in each scenario.

---

### 364. How do you implement service mesh without Kubernetes?
"Service mesh capabilities for Docker without Kubernetes by combining tools:

**mTLS between services**: SPIRE + Envoy sidecar per service. SPIRE issues X.509 SVIDs; Envoy handles TLS termination and mTLS initiation.

**Traffic management**: Envoy's cluster/listener config manages retries, timeouts, circuit breaking.

**Observability**: Envoy's stats exported to Prometheus, Jaeger for distributed tracing.

**Service discovery**: Consul provides SD; Envoy integrates via xDS (Consul's Envoy integration).

This is 80% of a service mesh without Kubernetes, but requires significant operational investment compared to Istio/Linkerd in K8s."

#### In-depth
The complexity of a DIY service mesh is why most organizations adopt Kubernetes + Istio rather than building mesh capabilities for Docker Swarm. The moving parts: SPIRE server cluster (HA for certificate issuance), SPIRE node agents on every host, Envoy sidecar templates per service, Consul cluster (HA), Prometheus for Envoy metrics, and Jaeger for traces. Each component requires its own HA configuration, monitoring, and operational runbooks. In K8s, Istio installs in 10 minutes and handles all of this automatically.

---

### 365. How do you configure Docker for high-bandwidth environments?
"Optimizing Docker networking for high-bandwidth:

1. **Jumbo frames on physical network**: MTU 9000. Set Docker MTU to 8950: `{"mtu": 8950}` in daemon.json.

2. **Enable kernel network acceleration**:
```bash
sysctl -w net.core.rmem_max=134217728
sysctl -w net.core.wmem_max=134217728
sysctl -w net.ipv4.tcp_rmem='4096 87380 67108864'
sysctl -w net.ipv4.tcp_wmem='4096 65536 67108864'
sysctl -w net.core.netdev_max_backlog=250000
```

3. **Host networking mode**: avoid NAT/bridge overhead for high-throughput containers.

4. **DPDK containers**: for <100µs latency requirements, use SR-IOV with Multus CNI to pass physical NIC directly to container."

#### In-depth
For network-intensive workloads (Kafka, high-frequency trading, video streaming), the container networking stack can be the bottleneck. SR-IOV (Single Root I/O Virtualization) bypasses the Linux networking stack entirely: a physical NIC creates virtual functions (VFs), each assigned to a container. The container's network traffic goes directly to hardware, bypassing the host kernel's networking stack. This achieves near-bare-metal throughput and latency. In Kubernetes, the Multus CNI plugin + the SR-IOV device plugin provision SR-IOV NICs to pods that request them via resource annotations.

---

### 366. What is the difference between `docker network connect` and `--network` flag?
"`--network` at container start: assigns the container to a network on creation.
```bash
docker run --network mynet nginx
```

`docker network connect`: adds a running container to an additional network without stopping it.
```bash
docker network connect another-net nginx-container
```

A container can be on multiple networks simultaneously — it gets a separate IP per network. Traffic to each network flows through a dedicated virtual interface.

`docker network disconnect mynet nginx-container`: removes from a network while running."

#### In-depth
The `docker network connect` command is useful for: temporarily connecting a debugging container to a production service's network (inspect traffic, run queries), connecting a proxy container to multiple backend networks, and dynamically changing a container's network access without restart. In Swarm, services are connected to networks at service creation — dynamic network membership changes require a service update with rolling restart. For blue-green in Swarm: `docker service update --network-add` and `--network-rm` to shift traffic between service versions.

---

### 367. How does Docker networking work in Windows containers?
"Windows containers have different networking drivers than Linux:

**nat**: Windows equivalent of Docker bridge. Default for Windows containers. Uses Windows Hyper-V Virtual Switch with NAT.

**transparent**: containers connect to physical network (like Linux macvlan). Gets IPs from physical DHCP.

**l2bridge**: L2 bridge network. Containers on the same host can communicate via L2 (MAC-level). External traffic goes through the host.

**overlay**: Windows Server 2019+ supports overlay for Docker Swarm across Windows hosts.

Windows networking backend: Host Networking Service (HNS) — Windows equivalent of netfilter/iptables."

#### In-depth
Windows container networking has significant limitations compared to Linux: overlay networking requires Windows Server 2019 (datacenter edition), not all features work on Windows 10/11. Mixed Linux/Windows Swarm clusters have networking restrictions — Windows tasks and Linux tasks can't be on the same overlay network. For mixed workloads, use separate services on separate networks and communicate via published VIPs. HNS rules (equivalent to iptables) can be inspected with `Get-HNSNetwork` and `Get-HNSEndpoint` PowerShell commands.

---

### 368. How do you implement traffic mirroring with Docker?
"Traffic mirroring (port mirroring/packet mirroring) copies live traffic to a monitoring endpoint for analysis.

**With iptables (Linux)**: mirror to another container:
```bash
# Mirror all traffic on docker0 to tshark container's IP
iptables -t mangle -A PREROUTING -i docker0 \
  -j TEE --gateway 172.18.0.100
```

**With Envoy sidecar**: configure request mirroring in Envoy:
```yaml
route:
  cluster: primary_service
  request_mirror_policy:
    cluster: shadow_service  # Mirror to shadow without affecting response
    runtime_fraction:
      default_value:
        numerator: 100  # Mirror 100% of traffic
```"

#### In-depth
Traffic mirroring is used in shadow testing (also called dark launches): deploy a new version of a service as a shadow, mirror production traffic to it, compare responses with the production service. If the shadow responds identically, the new version is safe to promote. Diffy (open source) and Envoy's request mirroring enable this. The shadow receives real traffic but its responses are discarded — users are unaffected. This catches bugs in new code that weren't caught by unit/integration tests, using real production traffic patterns.

---

### 369. What is the Docker Compose `healthcheck` merge behavior with overrides?
"When using multiple Compose files (base + override), `healthcheck` merges at the key level.

**Base**:
```yaml
healthcheck:
  test: ['CMD', 'curl', '-f', 'http://localhost/health']
  interval: 30s
```

**Override**:
```yaml
healthcheck:
  interval: 10s  # Only changes interval
```

**Result**: test stays from base, interval becomes 10s. Keys not specified in override retain base values.

To disable a health check in override: `healthcheck: {disable: true}` overrides all parent settings."

#### In-depth
Health check merge behavior is subtle — unlike most Compose key overrides (which replace the entire value), healthcheck merges at individual key level. This allows overrides to only change specific settings (make checks more frequent in CI without duplicating the test command). To completely replace a health check in an override, specify the `test` key too. To disable in dev (where health checks slow startup): `healthcheck: disable: true` in `docker-compose.override.yml` — development environment starts immediately without waiting for health checks.

---

### 370. How do you optimize Docker networking for low latency?
"Low-latency Docker networking optimizations:

1. **Avoid NAT**: use host networking or macvlan to eliminate NAT translation overhead.

2. **Co-locate communicating containers on the same host**: intra-host communication only uses the virtual bridge (kernel memory copy), not network hardware. Sub-millisecond latency vs. ~0.5ms for network.

3. **CPU affinity**: pin containers to specific CPUs, disable CPU frequency scaling: `performance` governor.

4. **Disable Nagle's algorithm in app**: `setsockopt(TCP_NODELAY)` — prevents 40ms Nagle delays for small packets.

5. **Use unix sockets** for same-host IPC: `/var/run/myservice.sock` — significantly lower latency than localhost TCP."

#### In-depth
Unix domain sockets provide the lowest-latency IPC between co-located containers (share the socket via a Docker volume): Nginx → PHP-FPM via unix socket instead of TCP `127.0.0.1:9000`. The elimination of TCP SYN/ACK/FIN handshake and TCP buffer management drops latency from ~0.1ms to ~0.01ms for small requests. For real-time or latency-sensitive apps (trading, gaming, media), co-location + unix sockets + CPU pinning + `TCP_NODELAY` is the optimization stack.

---

### 371. How does Docker handle multicast networking?
"Docker's bridge networking doesn't support multicast by default — packets sent to multicast addresses (224.0.0.0/4) aren't forwarded to other containers on the bridge.

**Enable multicast on a bridge**:
```bash
docker network create --driver bridge \
  -o 'com.docker.network.bridge.enable_ip_masquerade=false' \
  multicast-net
```
Then configure the bridge with multicast: `ip link set docker0 multicast on`

**For Swarm overlay**: multicast is not supported — overlay uses unicast VXLAN. Use application-level multicast (service discovery via Consul or mDNS-over-unicast)."

#### In-depth
Multicast in containers is niche but required for: mDNS-based service discovery (Bonjour/Avahi for IoT), ZeroConf networking, some media streaming protocols (IGMP multicast), and JGroups clustering (JBoss, Infinispan use IP multicast for cluster formation). The most common use in containers: running JBoss EAP in Docker where JGroups default multicast-based cluster discovery fails. Fix: configure JGroups to use **TCPPING** (unicast with static host list) or **JDBC_PING** (discovery via shared database) instead of multicast discovery.

---

### 372. What is the cni-plugins project and how it relates to Docker?
"CNI (Container Network Interface) is the standard networking interface for Kubernetes and other container orchestrators (NOT Docker's native networking model — Docker uses CNM).

CNI plugins are executables that implement network setup/teardown for container namespaces. They're called by the container runtime (containerd via CRI) when pods/containers start.

Standard CNI plugins (reference implementations): bridge, host-local IPAM, loopback, macvlan, ipvlan, portmap, bandwidth, firewall.

3rd-party CNI plugins: Calico, Cilium, Weave, Flannel — each provides its own networking model while conforming to the CNI API."

#### In-depth
Docker uses CNM (Container Networking Model) internally, not CNI. When Docker transitioned to containerd, containerd added a CRI plugin that maps Kubernetes pod networking to containerd's internal model using CNI. This is why Kubernetes can use the same containerd as Docker but with pluggable CNI networking. When debugging K8s networking issues, understand which CNI plugin is in use — `ls /etc/cni/net.d/` on a node shows the CNI config. The config tells you whether it's Calico, Flannel, Cilium, etc., and each has different debugging tools and log locations.

---

### 373. How do you implement Kubernetes Ingress with Docker containers?
"Kubernetes Ingress routes external HTTP/HTTPS traffic to internal services based on hostname and path.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
    - host: api.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: api-service
                port:
                  number: 80
  tls:
    - hosts: [api.example.com]
      secretName: api-tls-cert
```

Deploy the Ingress controller: `kubectl apply -f https://raw.githubusercontent.com/.../nginx/deploy.yaml`"

#### In-depth
The Ingress controller is itself a Docker container (nginx, Traefik, Envoy) running in Kubernetes. It watches Ingress resources and reconfigures itself dynamically when they change — no restarts needed. Ingress is not a built-in K8s feature (it's a spec that needs an implementation). Cert-manager automatically provisions Let's Encrypt TLS certificates for Ingress resources, making HTTPS automatic: annotate your Ingress with `cert-manager.io/cluster-issuer: letsencrypt-prod` and the certificate is provisioned and renewed automatically.

---

### 374. What is the difference between ClusterIP, NodePort, and LoadBalancer in Kubernetes?
"Three Kubernetes Service types for network access:

**ClusterIP** (default): internal cluster IP only. Only reachable from within the cluster. Used for internal service-to-service communication.

**NodePort**: exposes the service on each node's IP at a static port (30000-32767). `<NodeIP>:30080` reaches the service from outside the cluster.

**LoadBalancer**: provisions a cloud load balancer (AWS ALB/NLB, GCP Load Balancer) pointing to the node ports. Production-grade external access. Works only on cloud providers.

Mapping to Docker:
- ClusterIP ≈ Docker internal network service DNS
- NodePort ≈ `docker run -p host_port:container_port`
- LoadBalancer ≈ cloud provider fronting Docker's published ports"

#### In-depth
The cost of LoadBalancer services: each creates a separate cloud load balancer (~$15-50/month on AWS). For APIs with many microservices, you'd create 50 LoadBalancers (50× the cost). Solution: use ClusterIP for internal services, one LoadBalancer Ingress controller, and Ingress resources for routing rules. All external traffic flows through one cloud LB → Nginx Ingress Controller → ClusterIP Services. This pattern reduces external LB count from N services to 1 (per ingress class).

---

### 375. How do you implement network flow logging in Docker environments?
"Network flow logging captures metadata about network connections (src/dst IP, port, protocol, bytes, packets) without capturing payload.

**On Linux with nftables**:
```bash
nft add table filter
nft add chain filter logging {type filter hook forward priority 0}
nft add rule filter logging ct state new log prefix '"DOCKER-FLOW: "' group 0
```

**With eBPF/Hubble** (Cilium): `hubble observe --output json` provides real-time flow logs with service-level context (which pod sent what to which pod).

**AWS VPC Flow Logs**: if containers are on ECS Fargate, AWS captures flow logs at the ENI level."

#### In-depth
Hubble (Cilium's eBPF observability tool) is the most powerful network flow logging for Kubernetes: it captures L4 and L7 flows with service identity (not just pod IPs which change frequently). `hubble observe --pod myapp --port 8080` shows all connections to the myapp pod on port 8080 in real-time. Hubble Relay and UI provide cluster-wide network topology visualization. For compliance (PCI-DSS requirement 10.2, SOC 2), network flow logs prove which services communicated and when — a critical audit artifact.

---

### 376. How do you benchmark Docker network performance?
"Network performance benchmarking tools and methods:

**iperf3 test** (container-to-container throughput):
```bash
# Server container:
docker run --rm -d --name iperf-server networkstatic/iperf3 -s

# Client container (same host):
docker run --rm networkstatic/iperf3 -c iperf-server -t 30
# Expected: 20+Gbps for bridge mode
```

**Latency test**:
```bash
docker run --rm alpine ping -c 100 iperf-server
# Expected: <0.1ms for same-host bridge, 0.2-0.5ms for overlay
```

**HTTP benchmark** with wrk:
```bash
docker run --rm williamyeh/wrk -t12 -c400 -d30s http://api:8080
```"

#### In-depth
Benchmark baselines for planning: same-host bridge networking should achieve near-kernel memory bandwidth (20-40 Gbps on modern hardware). Cross-host overlay networking achieves physical network speed minus VXLAN overhead (~1 Gbps on GbE, ~9 Gbps on 10GbE). Host networking matches bare-metal. If your benchmark results are significantly lower: check MTU configuration (fragmentation kills throughput), check CPU usage (encryption or NAT may be CPU-bound), check for network policy rules adding per-packet overhead. iperf3 with `-P 8` (parallel streams) better tests maximum throughput under realistic conditions.

---

### 377. What is Flannel and how does it compare to other CNI plugins?
"Flannel is the simplest Kubernetes CNI plugin — a basic overlay network with minimal features.

**Architecture**: flanneld daemon on each node exchanges routing info (UDP by default, or VXLAN). Each node gets a subnet (/24) from the cluster's pod CIDR. Flannel routes traffic between node subnets.

**Flannel vs Calico**:
- Flannel: simpler, no network policies, UDP/VXLAN backends, less complexity
- Calico: BGP routing (faster), network policies, L7 awareness, IPAM flexibility

**Flannel vs Cilium**:
- Flannel: no network policies, no observability
- Cilium: eBPF-based, L7 policies, full observability, higher complexity"

#### In-depth
Flannel is appropriate for: small clusters, development environments, or clusters where simplicity trumps features. Its lack of NetworkPolicy support is a significant security limitation for production — you cannot enforce pod-to-pod access controls. k3s (lightweight Kubernetes) ships with flannel as default, but supports Calico as an alternative. For any production cluster with security requirements: Calico or Cilium. The migration path (Flannel → Calico/Cilium) is complex and typically requires cluster rebuild — choose the right CNI at cluster creation.

---

### 378. How do you handle DNS in Docker Swarm with external services?
"Docker Swarm services often need to reach external services (databases outside the cluster, SaaS APIs).

**External service aliasing** — create a Swarm service that proxies to the external endpoint:
```bash
docker service create \
  --name ext-postgres \
  --network mynet \
  -p 5432:5432 \
  --env POSTGRES_REMOTE_HOST=prod.database.example.com \
  haproxy:alpine
```

Or use Docker network external DNS aliases:
```bash
# Add to /etc/hosts on each node (not ideal)
# Better: use ExternalName in Kubernetes (no Swarm equivalent)
```

**Recommended**: use real DNS entries that resolve to external services, inject via environment variables."

#### In-depth
Swarm's lack of ExternalName service (a K8s feature that creates a DNS alias for an external hostname) is a gap. Kubernetes `ExternalName` services make `http://my-db` DNS inside the cluster resolve to `prod.database.example.com` transparently — service-internal code doesn't know it's an external service. In Swarm, the cleanest workaround: configure the application to use a hostname injected via environment variable (`DATABASE_HOST=prod.database.example.com`) rather than a fixed service-name. This makes external services configurable without modifying internal DNS.

---

### 379. How do you implement traffic shaping for Docker containers?
"Traffic shaping controls bandwidth, latency, and packet loss at the network level.

**Using tc (traffic control)** on the host interface:
```bash
# Limit container eth0 to 100Mbps
docker_pid=$(docker inspect -f '{{.State.Pid}}' mycontainer)
nsenter -t $docker_pid -n tc qdisc add dev eth0 root tbf \
  rate 100mbit burst 32kbit latency 400ms

# Add 50ms latency + 10% packet loss (chaos testing)
nsenter -t $docker_pid -n tc qdisc add dev eth0 root netem \
  delay 50ms loss 10%
```

**Docker Compose via CNI bandwidth plugin**:
```yaml
networks:
  limited:
    driver: bridge
    driver_opts:
      com.docker.network.driver.mtu: '1450'
```"

#### In-depth
Traffic shaping is invaluable for: **chaos engineering** (simulate poor network conditions to test resilience), **rate limiting bandwidth-heavy containers** (prevent one container from saturating the host NIC), and **QoS prioritization** (give critical services higher bandwidth share). Tools like Pumba use the `tc netem` approach internally. For Kubernetes: the `bandwidth` CNI meta-plugin applies `tc` rules to pod interfaces based on annotations: `kubernetes.io/ingress-bandwidth: 10M` — a simple, declarative way to set per-pod rate limits.

---

### 380. How do you use Consul service mesh with Docker containers?
"Consul Connect provides service mesh capabilities via sidecar proxies registered in Consul.

**Setup per service**:
1. Start Consul agent on each host: `consul agent -dev`
2. Register service with Consul Connect:
```json
{
  'name': 'api',
  'connect': {
    'sidecar_service': {
      'proxy': {
        'upstreams': [{'destination_name': 'db', 'local_bind_port': 5432}]
      }
    }
  }
}
```
3. Start Envoy sidecar: `consul connect envoy -sidecar-for api`
4. App talks to db via `localhost:5432` (the local Envoy upstream)"

#### In-depth
Consul Connect's mTLS works via certificates issued by Consul's built-in CA (or an external CA like Vault PKI). Each service's Envoy sidecar presents its certificate for identity verification. The certificates encode the service identity (spiffe://dc1.consul/ns/default/dc/dc1/svc/api). Access control via Consul intentions: `consul intention create -allow api db` allows api to connect to db; without an allow intention, Envoy rejects the connection. This is the zero-trust model: no implicit trust based on network location, every connection requires an explicit allow rule.

---
