# 🧩 **Docker Compose Advanced (151–160)**

---

### 151. How do you override Docker Compose configs for staging vs prod?
"I use **multiple Compose files** layered with the `-f` flag.

Base file: `docker-compose.yml` — defines all services and common config.
Override file: `docker-compose.prod.yml` — sets production-specific values.

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

`docker-compose.prod.yml` only contains overrides — it's merged with the base. Example: base sets `image: app:dev`, prod override sets `image: registry/app:v1.2.3` and adds resource limits.

For staging: `docker compose -f docker-compose.yml -f docker-compose.staging.yml up`."

#### In-depth
The merge behavior of Compose files: for scalars (image, command), the later file wins. For lists (volumes, ports), values are appended. For mappings (environment, labels), keys from the later file override, new keys are added. This merge logic enables clean separation: never repeat configs, only declare differences. The `COMPOSE_FILE` environment variable sets the default file list — useful in deployment scripts.

---

### 152. What is the purpose of `.env` in Docker Compose?
"The `.env` file provides **default values for variables** used in `docker-compose.yml`.

```dotenv
# .env
APP_PORT=3000
IMAGE_TAG=latest
DB_PASSWORD=localpass
```

These are substituted in the Compose file: `image: myapp:${IMAGE_TAG}` becomes `myapp:latest`. It avoids hardcoding values and enables easy environment switching.

**Important**: `.env` is for Docker Compose file variable substitution. It's NOT automatically injected as environment variables into containers (that's `env_file:` under the service)."

#### In-depth
The difference between `.env` (Compose variable substitution) and `env_file:` (container environment injection) confuses many users. Test with `docker compose config` — it shows the rendered Compose file after all variable substitution is applied. Never commit `.env` files with real secrets to git — add them to `.gitignore` and provide a `.env.example` with dummy values as documentation for team members.

---

### 153. How do you handle environment-specific containers in Compose?
"Use **Compose profiles** — tag services with profile names and only start services matching the active profile.

```yaml
services:
  api:
    image: myapp
  db:
    image: postgres
    profiles: [dev, test]
  db-prod:
    image: postgres
    profiles: [prod]
    deploy:
      replicas: 3
```

Start dev: `docker compose --profile dev up`.
Start prod: `docker compose --profile prod up`.

Services without a profile always start. Services with profiles only start when that profile is active."

#### In-depth
Profiles solve the multi-environment config sprawl problem. Alternative patterns before profiles existed: separate Compose files per environment (maintenance burden) or `docker compose --scale db=0 up` to prevent certain services from starting (hacky). Profiles are the clean solution. The `COMPOSE_PROFILES` environment variable sets active profiles: `COMPOSE_PROFILES=dev,debug docker compose up` — useful in CI where you can't always pass CLI flags.

---

### 154. How do you connect external containers to Docker Compose network?
"Two approaches:

**1. Join the Compose network from outside**:
```bash
# Get the network name
docker network ls | grep projectname
# Connect an external container
docker network connect projectname_default external-container
```

**2. Declare the network as external in Compose**:
```yaml
networks:
  shared:
    external: true
    name: existing-network-name
```

Now both your Compose services and pre-existing containers can be on `existing-network-name`."

#### In-depth
The external network pattern is essential for multi-Compose-project setups — when microservices are managed by separate Compose files but need to communicate. Create a shared network once: `docker network create shared-infra`. Each project's `docker-compose.yml` declares it as external. All projects share the same DNS-discoverable network without merging their Compose files.

---

### 155. Can you define multiple Docker Compose files?
"Yes — this is the recommended pattern for environment-specific configuration.

**Using `-f` flag**: `docker compose -f base.yml -f override.yml up`

**Using `COMPOSE_FILE` variable**: `export COMPOSE_FILE=docker-compose.yml:docker-compose.dev.yml`

**Common structure**:
- `docker-compose.yml` — base service definitions
- `docker-compose.override.yml` — auto-loaded in dev (Docker Compose loads this by default)
- `docker-compose.prod.yml` — production overrides (explicit `-f`)
- `docker-compose.test.yml` — CI/testing specific services

`docker-compose.override.yml` is the Dev convention — auto-loaded when `docker compose up` is run without `-f`."

#### In-depth
The auto-loading of `docker-compose.override.yml` is a powerful convention. Keep production-specific settings out of the base file entirely — it should be deployment-agnostic. The override file adds dev conveniences (volume mounts for hot reload, expose extra ports, add debugging tools). This way new developers get a productive local environment by default, and prod deployments explicitly specify a clean prod override — no accidental dev settings in production.

---

### 156. How do you use `build` vs `image` in Compose?
"`build:` tells Compose to build the image from a Dockerfile. `image:` tells Compose to use an existing image from a registry (or local Docker cache).

```yaml
# Build from local Dockerfile
api:
  build:
    context: ./api
    dockerfile: Dockerfile.prod
    args:
      - BUILD_ENV=production

# Use pre-built image from registry
db:
  image: postgres:15-alpine
```

You can specify both: `image: myregistry/app:latest` + `build: .`. Then `docker compose build` builds and tags as `myregistry/app:latest`, and `docker compose push` pushes it."

#### In-depth
In production, always use `image:` with a specific tag in Compose — never `build:`. Building in production means the build environment (Dockerfile context, dependencies) affects what runs. The correct production workflow: build in CI, push to registry, reference the specific tag in the production Compose file. The production host should never have source code or a Docker build context — only the registry and the Compose file.

---

### 157. How do you restart services automatically in Compose?
"Use the `restart` policy per service:

```yaml
services:
  api:
    image: myapp
    restart: unless-stopped
```

Restart policy options:
- `no` — never restart (default)
- `always` — always restart, even after daemon restart
- `on-failure` — restart only on non-zero exit code
- `on-failure:3` — restart up to 3 times, then give up
- `unless-stopped` — restart always unless manually stopped

For production: `unless-stopped` or `always`. `on-failure:3` good for crash detection (stops retrying after 3 failures, surfacing the issue)."

#### In-depth
`restart: unless-stopped` vs `restart: always`: the difference appears after a Docker daemon restart (e.g., after a host reboot). `always` restarts the container automatically after daemon restart. `unless-stopped` only restarts it if it wasn't manually stopped before the daemon restarted. For production servers that must survive reboots: use `always`. For development where you want to explicitly control which containers run: use `unless-stopped`.

---

### 158. How do you define health checks in Compose?
"The `healthcheck` key in a service definition:

```yaml
services:
  api:
    image: myapp
    healthcheck:
      test: ['CMD', 'curl', '-f', 'http://localhost:8080/health']
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
  db:
    image: postgres
    healthcheck:
      test: ['CMD-SHELL', 'pg_isready -U postgres']
      interval: 5s
      timeout: 5s
      retries: 5
```

The `start_period` grace period prevents false health failures during slow container startup."

#### In-depth
Once health checks are defined, you can use `condition: service_healthy` in `depends_on` — the depending service won't start until the health check passes. This properly solves the DB-not-ready startup race condition. Without health check conditions, you need hacky solutions like `sleep 5 && start_app.sh` or retry loops in entrypoint scripts. Health checks make the orchestration explicit and reliable.

---

### 159. What are Compose profiles and how are they used?
"Profiles allow **conditional service activation** within a single Compose file.

```yaml
services:
  app:
    image: myapp      # Always starts
  redis:
    image: redis
    profiles: [cache]  # Only starts with --profile cache
  mailhog:
    image: mailhog/mailhog
    profiles: [dev]   # Only starts with --profile dev
  prometheus:
    image: prom/prometheus
    profiles: [monitoring]
```

Commands:
- `docker compose up` — only starts `app`
- `docker compose --profile dev up` — starts `app` + `mailhog`
- `docker compose --profile dev --profile cache up` — starts `app` + `mailhog` + `redis`"

#### In-depth
Compose profiles are particularly valuable for optional infrastructure services. In a large team, some developers need email testing (MailHog), some need caching (Redis), some need monitoring (Prometheus). With profiles, the single `docker-compose.yml` serves all needs without everyone starting every service. CI pipelines use specific profiles: `--profile test` for testing infrastructure, `--profile ci` for CI-specific mocks.

---

### 160. How do you upgrade a running Compose stack with zero downtime?
"Zero-downtime Compose upgrades require a **load balancer that can route between old and new containers**.

**Pattern with Traefik**:
1. Scale up the new version alongside the old: `docker compose up --scale api=2 --no-recreate`
2. Wait for new containers to pass health checks
3. Traefik automatically adds new containers to its routing (via Docker labels)
4. Stop old containers: `docker compose stop api_old`

**Without Traefik**: Native Compose doesn't support zero-downtime upgrades. Use Docker Swarm or Kubernetes for built-in rolling update support.

For simple single-service updates: accept brief downtime — `docker compose up -d api` recreates the container in a few seconds."

#### In-depth
Traefik's Docker provider watches for container events and auto-configures its routing table in real time. When a container with the right labels starts, Traefik adds it as a backend. When it stops, Traefik removes it. Combined with health check conditions and `depends_on`, you can achieve true zero-downtime updates with plain Docker Compose — no Swarm or Kubernetes required for simple services.

---
