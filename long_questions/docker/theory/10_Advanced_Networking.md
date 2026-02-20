# ⚙️ **Docker Networking Advanced (111–120)**

---

### 111. What is a MACVLAN network in Docker?
"MACVLAN gives each container its own **MAC address and IP address directly on the physical network**.

The container appears as a full physical device on the network — DHCP servers, routers, and other devices see it as a separate machine (not a container). This is useful when you need containers to receive IPs from an existing DHCP server or when network policies require each service to have a distinct IP.

Create it: `docker network create -d macvlan --subnet=192.168.1.0/24 --gateway=192.168.1.1 -o parent=eth0 mymacvlan`."

#### In-depth
MACVLAN has a significant limitation: containers on a MACVLAN network cannot communicate with the **host itself** — because the physical NIC switches from host mode to promiscuous mode, and the host's kernel doesn't receive frames destined for the container's virtual MAC. The fix: create a MACVLAN interface on the host too and assign the host an IP on that MACVLAN. This is an advanced networking scenario typically used for legacy apps that require L2 network presence.

---

### 112. How does bridge networking work in Docker?
"The Docker bridge driver creates a **virtual Ethernet bridge** (`docker0` by default) on the host.

When containers join the bridge network, Docker creates a **veth pair** (virtual Ethernet cable) — one end in the container's network namespace, the other attached to the bridge. The bridge acts like a network switch connecting all containers.

Packets from a container flow through the veth to the bridge. The bridge routes to other containers directly (same subnet) or to the host's default route for external traffic. NAT (iptables masquerade) is applied for outbound internet traffic."

#### In-depth
The default `docker0` bridge has a limitation: no automatic DNS. Containers must discover each other by IP, which changes every restart. User-defined bridge networks solve this: Docker creates an embedded DNS server (127.0.0.11) that answers queries using container names. This is the #1 reason to always create custom networks instead of using the default bridge.

---

### 113. What is host networking, and when should you use it?
"Host networking (`--network=host`) removes all network isolation between the container and the host. The container uses the **host's network stack directly** — same IP, same ports, same routing table.

There's no port mapping needed (or possible) — the app binds directly to host ports. Performance is near-native since there's no NAT translation.

I use it for: **network monitoring tools** (need to see host interfaces), **high-performance networking** where NAT overhead matters (gaming, trading), and **troubleshooting** when diagnosing network issues without container networking in the way."

#### In-depth
Host networking on macOS and Windows Docker Desktop (which runs a Linux VM) has a different meaning: containers join the Linux VM's network, not the actual macOS/Windows host network. This trips up many developers. On Linux (bare metal or native Docker), `--network=host` truly shares the host kernel's network stack — no VM in the middle.

---

### 114. How do you create a custom Docker network?
"Simple: `docker network create mynet`.

With options: `docker network create --driver bridge --subnet 10.20.0.0/16 --gateway 10.20.0.1 --ip-range 10.20.5.0/24 --opt com.docker.network.bridge.name=mybridge mynet`.

For overlay (Swarm): `docker network create -d overlay --attachable mynet`. The `--attachable` flag allows standalone containers (not just Swarm services) to connect to the overlay network."

#### In-depth
Custom bridge networks support **internal isolation**: `docker network create --internal mynet` creates a network with no external connectivity. Containers on this network can talk to each other but have no internet access or host access. Useful for databases or caches that should have zero external exposure — they communicate only with the app containers on the same network.

---

### 115. How do you inspect a Docker network?
"I use `docker network inspect <network-name>`.

It shows: driver, subnet, gateway, connected containers and their IPs, and network-specific options. The output is JSON — I use `--format` to extract specific fields:

`docker network inspect mynet --format='{{range .Containers}}{{.Name}}: {{.IPv4Address}}{{println}}{{end}}'`

To see which networks a container is on: `docker inspect container --format='{{json .NetworkSettings.Networks}}' | jq .`."

#### In-depth
Network inspection is essential for debugging connectivity. The key field is `Containers` — it maps container IDs to their IPs on this network. If a container that should be on a network isn't showing up there, that's your connectivity problem. Also check `Options` for custom MTU settings — MTU mismatches between overlay network MTU and physical network MTU cause silent packet drops that are notoriously hard to debug.

---

### 116. Can two containers on different networks communicate?
"Not by default — Docker networks are **isolated** from each other. Containers on network A cannot reach containers on network B.

To allow communication: **attach a container to multiple networks**. `docker network connect networkB container-on-networkA`. Now that container has an interface on both networks and can bridge them.

This is useful for a reverse proxy container: connect it to a public-facing network AND the internal app network. It routes traffic from outside to services inside the internal isolated network."

#### In-depth
This multi-network pattern is common in security-conscious architectures. An API gateway container is attached to: `external-net` (accessible from internet), `app-net` (connected to API services), and `db-net` isn't accessible at all from the gateway — only app services connect to the database network. This is network segmentation without firewalls — pure Docker networking.

---

### 117. What is port mapping, and how is it done?
"Port mapping publishes a container port on the host, making the container service accessible from outside.

`docker run -p 8080:80 nginx` maps host port 8080 to container port 80. External traffic hits host:8080 → iptables NAT → container:80.

Multiple mappings: `-p 8080:80 -p 8443:443`. Bind to specific host IP: `-p 127.0.0.1:8080:80` (localhost only). Random host port: `-p 80` (Docker assigns an available host port). Check the assigned port: `docker port container 80`."

#### In-depth
Docker implements port mapping via **iptables DNAT rules** on the host. `docker run -p 8080:80` adds an iptables rule: packets arriving at host port 8080 are destination-NAT'd to the container's IP:80. Docker manages these rules in the `DOCKER` chain. Be careful with host firewalls: UFW, firewalld can conflict with Docker's iptables rules — Docker bypasses UFW by default, which can expose ports you meant to block.

---

### 118. How do you debug Docker network issues?
"My step-by-step debug process:

1. `docker network inspect <net>` — verify containers are connected with correct IPs
2. `docker exec container ping <other-container-ip>` — test L3 connectivity
3. `docker exec container nslookup <other-container-name>` — test DNS
4. `docker exec container curl http://<other>:port` — test application port
5. `docker exec container netstat -tlnp` — verify the service is listening
6. `docker exec container ss -tuln` — alternative to netstat
7. If ping works but curl doesn't: the service isn't listening on 0.0.0.0 (loopback-only binding)"

#### In-depth
For deep network debugging, tcpdump is powerful. `docker run --rm --network container:target nicolaka/netshoot tcpdump -i any -n host <ip>` attaches the netshoot debugging image to the target container's network namespace and runs packet capture. Netshoot contains every network debugging tool you need in one container (tcpdump, curl, dig, nmap, iperf, etc.) without modifying the production container.

---

### 119. How do you restrict container-to-container communication?
"By default, containers on the same network can communicate freely (ICC — inter-container communication).

**Method 1**: Separate networks — put services that shouldn't talk in different networks.
**Method 2**: Disable ICC on the bridge: `docker network create -o com.docker.network.bridge.enable_icc=false mynet`. With ICC disabled, containers on the same bridge cannot talk to each other — only to services published via host port mapping.
**Method 3**: Use internal networks with no external access for isolated service groups."

#### In-depth
The `--icc=false` daemon option disables ICC on the default `docker0` bridge globally. With it enabled (default), any container on the bridge can connect to any other container's ports — no explicit allow rules needed. Disabling ICC is a security best practice for multi-tenant environments. But it breaks service discovery — you then need explicit `--link` flags (legacy) or separate networks for each communicating pair.

---

### 120. How do you implement DNS-based service discovery in Docker?
"Docker's embedded DNS (127.0.0.11) automatically provides DNS-based service discovery for user-defined networks.

In a custom network, every container is reachable by its **container name** as a DNS A record. In Compose, every service is reachable by its **service name** as well. In Swarm, every service has a **VIP** returned for its name, with IPVS load-balancing across replicas.

No configuration needed — just create a user-defined network and connect containers to it. Standard DNS queries (`nslookup api`, `curl http://api:8080`) resolve automatically."

#### In-depth
Docker's DNS also supports **aliases** — a container can register multiple DNS names on a network: `docker network connect --alias frontend mynet mycontainer`. This is useful for blue-green deployments: both old and new containers register as `frontend`, DNS round-robins between them during the transition. Or for zero-downtime alias cutover: remove the alias from old, add to new.

---
