# 🌐 Docker Networking & Docker Compose Deep Dive (Service-Based Companies)

This document covers Docker networking concepts and Docker Compose patterns frequently tested at service-based company interviews.

---

### Q1: How does Docker networking work? Explain the default network types.

**Answer:**
Docker has several built-in network drivers. Understanding them is key for multi-service applications and troubleshooting connectivity issues.

**Default Network Types:**

**1. Bridge (default for standalone containers)**
- Docker creates a virtual switch called `docker0` on the host.
- Each container gets a virtual ethernet interface (`veth`) connected to this bridge.
- Containers on the same bridge network can communicate using **container names or IP addresses**.
- External traffic is NAT-ed through the host's IP.
```bash
docker network create myapp-net    # Custom bridge network
docker run --network myapp-net --name web myapp:latest
docker run --network myapp-net --name db postgres:15
# 'web' can reach 'db' via hostname 'db'
```

**2. Host**
- Removes network isolation — container uses the **host's network stack directly**.
- No port mapping needed (container port IS host port).
- Faster (no NAT overhead) but higher security risk.
```bash
docker run --network host nginx
# nginx on port 80 inside container is directly on host port 80
```

**3. None**
- Container has **no network access** — completely isolated.
- Useful for security-sensitive batch jobs that only write to volumes.
```bash
docker run --network none myapp:latest
```

**4. Overlay**
- Used for **multi-host** Docker Swarm networking.
- Tunnels container traffic between hosts using VXLAN.
- Containers on different Docker hosts communicate as if on the same LAN.

**5. Macvlan**
- Assigns a unique MAC address to a container — it appears as a **physical network device**.
- Used for legacy applications that expect to be on the actual network segment.

---

### Q2: How does container-to-container communication work in Docker Compose?

**Answer:**
Docker Compose automatically creates a **default bridge network** for all services in the same `compose.yml`. Services can reach each other by their **service name** as the hostname — Docker's embedded DNS handles name resolution.

**Example:**
```yaml
services:
  backend:
    build: .
    environment:
      - DB_HOST=database      # 'database' resolves to the db container's IP
      - REDIS_HOST=cache      # 'cache' resolves to the redis container's IP
    
  database:
    image: postgres:15
    
  cache:
    image: redis:7-alpine
```

- `backend` can reach `database` at hostname `database`, port `5432`
- `backend` can reach `cache` at hostname `cache`, port `6379`
- `database` and `cache` **cannot** reach each other if they're on separate composed stacks

**Custom networks in Compose (best practice for isolation):**
```yaml
services:
  frontend:
    networks: [public-net]
    
  backend:
    networks: [public-net, private-net]
    
  database:
    networks: [private-net]  # Only reachable by backend, not frontend

networks:
  public-net:
  private-net:
    internal: true  # No external internet access
```

---

### Q3: How does port mapping work in Docker and what is the difference between `EXPOSE` and `-p`?

**Answer:**

**`EXPOSE` (in Dockerfile):**
- A **documentation directive** — it tells other developers and tools which port the application inside the container listens on.
- **Does NOT actually open or publish the port** to the host.
- Acts as metadata only.

```dockerfile
EXPOSE 8080  # Documents that the app listens on 8080 — doesn't open it
```

**`-p / --publish` (in `docker run`):**
- **Actually maps** a host port to a container port.
- Format: `-p <host_port>:<container_port>`
- This creates a NAT rule using `iptables` on the host.

```bash
docker run -p 3000:8080 myapp:latest
# Host port 3000 → Container port 8080
# Visit http://localhost:3000 to hit the container's port 8080
```

**`-P` (uppercase — publish all exposed ports):**
- Publishes all ports declared with `EXPOSE` to **random available host ports**.
```bash
docker run -P myapp:latest
docker ps  # Shows: 0.0.0.0:49153->8080/tcp
```

**Binding to specific host interface:**
```bash
docker run -p 127.0.0.1:3000:8080 myapp:latest
# Only accessible from localhost, not from external IPs (security best practice)
```

---

### Q4: What is a Health Check in Docker and how do you configure one?

**Answer:**
A **Docker health check** is a command that Docker runs periodically inside a container to determine if the application inside is functioning correctly. A container can be `running` but `unhealthy` if the app has crashed internally or entered a deadlock.

**Configuring in Dockerfile:**
```dockerfile
FROM node:20-alpine
...

HEALTHCHECK --interval=30s \
            --timeout=10s \
            --start-period=15s \
            --retries=3 \
            CMD curl -f http://localhost:3000/health || exit 1
```

**Parameters:**
| Parameter | Meaning | Default |
|---|---|---|
| `--interval` | How often to run the check | 30s |
| `--timeout` | Max time for one check | 30s |
| `--start-period` | Grace period after container start (gives app time to boot) | 0s |
| `--retries` | Failures before marking `unhealthy` | 3 |

**Configuring in Docker Compose:**
```yaml
services:
  api:
    image: myapi:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 15s
```

**Health states:**
- `starting` — Within `start-period` (not yet evaluated)
- `healthy` — All recent checks passed
- `unhealthy` — Reached `retries` failures

**`depends_on` with health check (Compose):**
```yaml
services:
  web:
    depends_on:
      db:
        condition: service_healthy   # Wait until db is healthy, not just started
  db:
    image: postgres:15
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      retries: 5
```

---

### Q5: How do you connect multiple Docker Compose files or share services across projects?

**Answer:**

**Scenario 1: Override files (common dev vs prod configs)**
```bash
# Base config
docker compose -f compose.yml -f compose.override.yml up

# compose.yml — base (checked into git)
# compose.override.yml — local dev overrides (in .gitignore)
# compose.prod.yml — production-specific settings
```

**Scenario 2: External networks (sharing services across Compose projects)**
A common pattern is a shared `monitoring` or `database` project and multiple app projects connecting to it.

```yaml
# In the database project (docker-compose.yml)
networks:
  shared-db-net:
    name: shared-db-net  # Explicit name makes it predictable

services:
  postgres:
    image: postgres:15
    networks: [shared-db-net]
```

```yaml
# In your app project (compose.yml)
networks:
  shared-db-net:
    external: true   # Join the already-existing network

services:
  app:
    build: .
    networks: [shared-db-net]
    environment:
      - DB_HOST=postgres  # Resolves because they're on the same network
```

**Scenario 3: Profiles (run subset of services)**
```yaml
services:
  app:
    build: .
  
  db:
    image: postgres:15
    profiles: ["dev", "testing"]   # Only starts with --profile dev or --profile testing
  
  mock-server:
    image: mockoon/cli
    profiles: ["testing"]
```
```bash
docker compose --profile dev up        # starts app + db
docker compose --profile testing up    # starts app + db + mock-server
docker compose up                      # starts only app
```

---

### Q6: How do you inspect and debug Docker networking problems?

**Answer:**
Network debugging is a common interview topic and a critical real-world skill.

**Useful commands:**
```bash
# List all networks
docker network ls

# Inspect a network (see connected containers, their IPs, subnet)
docker network inspect bridge         # Default bridge details
docker network inspect myapp_default  # Compose default network

# Check container's IP and network
docker inspect <container> --format '{{json .NetworkSettings.Networks}}' | python -m json.tool

# Test connectivity between containers
docker exec web ping db               # By service name (Compose)
docker exec web nslookup db           # DNS resolution check
docker exec web curl http://db:5432   # Port reachability

# Check port bindings on the host
docker port <container_id>
# OR
netstat -tlnp | grep docker
```

**Common problems and solutions:**

| Problem | Symptom | Fix |
|---|---|---|
| Can't reach container by name | `ping: bad address 'db'` | Ensure both are on the same **custom** network (not default bridge — default bridge doesn't support DNS) |
| Port not accessible from host | Connection refused on mapped port | Check `-p` flag binding, verify `EXPOSE` is set, check app actually binds to `0.0.0.0` not `127.0.0.1` |
| Container can't reach internet | DNS queries fail | Check host DNS, try `--dns 8.8.8.8` |
| Two Compose projects can't talk | Cross-project hostname fails | Use external named networks |

**Quick diagnostic container:**
```bash
# Spin up a temporary netshoot container for debugging
docker run --rm -it --network <target-network> nicolaka/netshoot
# Has: curl, dig, nslookup, ping, tcpdump, ss, iperf, etc.
```

---

*Prepared for technical screening and L1/L2 rounds at service-based companies.*
