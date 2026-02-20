# 🔁 **Docker CI/CD Integration & GitOps (141–150)**

---

### 141. How do you version Docker images in a CI pipeline?
"I use multiple tags on the same image for different purposes:

- **Git SHA tag**: `app:a3f5c82` — immutable, traceable to exact commit
- **Semantic version tag**: `app:v2.3.1` — human-readable release identifier
- **Branch/environment tag**: `app:main`, `app:staging` — mutable, always points to latest for that branch
- **Latest**: `app:latest` — points to the most recent main build

In GitHub Actions:
```yaml
tags: |
  myregistry/app:latest
  myregistry/app:${{ github.sha }}
  myregistry/app:v${{ github.run_number }}
```"

#### In-depth
The combination of SHA (for traceability) + semantic version (for human readability) + `latest` (for easy exploration) covers all use cases. For deployment, always reference SHA or specific version tags — never `latest` — because latest can be rewritten. Kubernetes `imagePullPolicy: IfNotPresent` with `latest` is particularly dangerous: it may use a stale cached image with no notification.

---

### 142. What is the best strategy to tag Docker images?
"Follow the **single source of truth** principle — the image tag must be traceable to a commit.

**For production**: `registry/app:v1.2.3-sha-abc1234`. The semver communicates what changed; the SHA pins it to an exact commit.

**For development branches**: `registry/app:branch-name-sha`.

**For releases**: create a Git tag → CI builds `registry/app:v1.2.3` → deploy by version.

Never reuse tags for different images (don't overwrite `v1.2.3` after it's deployed). Use semantic versioning for packages, never for deployment artifacts."

#### In-depth
**Immutable tags** are the gold standard. Some registries (ECR with `imageTagMutability=IMMUTABLE`) enforce this. An immutable `v1.2.3` tag can only be pushed once. If you need to hotfix, it becomes `v1.2.4`. This prevents the 'same tag, different image' problem that has caused production incidents when teams accidentally overwrote a deployed tag with a broken build.

---

### 143. How do you implement GitOps with Docker images?
"GitOps: **Git is the single source of truth for both application code and infrastructure configuration**.

Workflow:
1. Developer pushes code → CI builds and pushes `app:sha-abc123`
2. CI opens a PR to a **config repo** updating the image tag in the deployment manifest (Kubernetes YAML or Compose file)
3. Code review on the config PR
4. Merge to config repo → GitOps operator (Argo CD, Flux) detects the change and applies it

The config repo contains exact image tags — no `latest`. Changes to production go through Git history, with PR reviews and audit trail."

#### In-depth
**Argo CD Image Updater** automates step 2: it watches the registry for new image versions matching a tag policy and auto-creates commits in the config repo. This closes the loop — code push → image build → config update → deploy, all automated. The Git commit history becomes a deployment audit log: who changed what image to which version and when, with full approval workflow.

---

### 144. How do you ensure rollback safety in Docker-based deployments?
"Rollback safety requires: **preserved old image tags**, **versioned config**, and **tested rollback procedure**.

Technical measures:
- Keep old images in the registry (lifecycle policy: keep last 10 versions)
- Use Swarm's built-in rollback: `docker service rollback myservice`
- In Kubernetes: `kubectl rollout undo deployment/myapp`
- In GitOps: revert the config repo commit → Argo CD/Flux applies the old image tag

Non-technical: define rollback criteria upfront (error rate >1% → rollback), assign an on-call rotation to monitor deployments, and test rollbacks in staging regularly."

#### In-depth
The often-overlooked rollback scenario: **schema migrations**. If v2.0 runs a database migration that v1.9 can't handle, rolling back v2.0 may corrupt the database. Solutions: **expand-contract pattern** (v2.0 only adds new columns, never drops old ones), **forward-only migrations** (never rollback the schema), and **blue-green deployments** with separate database environments. Design for rollback at the schema level, not just at the image level.

---

### 145. How do you automate security scanning in CI for Docker images?
"I integrate scanning at multiple pipeline stages:

**Pre-build**: lint Dockerfile with Hadolint in the PR check.
**Post-build**: `trivy image --exit-code 1 --severity CRITICAL myapp:$SHA`. Pipeline fails on critical CVEs.
**Post-push**: Docker Scout or Snyk container monitor scans images in the registry and alerts on new CVEs (even after deployment).

```yaml
- name: Scan image
  run: |
    trivy image --format sarif --output results.sarif myapp:${{ github.sha }}
- name: Upload scan results
  uses: github/codeql-action/upload-sarif@v2
  with:
    sarif_file: results.sarif
```"

#### In-depth
Uploading SARIF results to GitHub Security tab makes vulnerabilities visible in the PR — developers see issues before merging. The philosophy: **shift security left**. Catching a critical CVE during code review is faster (fix: update the base image tag in Dockerfile) than catching it in post-deployment scanning. Establish a vulnerability SLA: critical CVEs patched within 1 business day, high within a week, medium within a sprint.

---

### 146. What's the difference between ephemeral containers and long-lived ones in CI?
"**Ephemeral containers** are created fresh for each CI job and destroyed afterwards. They have no state between runs. This ensures **reproducibility** — every build starts from the same clean image state.

**Long-lived containers** persist between CI jobs. State accumulates: residual files, leftover processes, disk usage buildup. They're faster (no startup time) but risk contamination between builds.

Best practice: ephemeral for builds and tests (clean, reproducible), long-lived only for shared infrastructure (the CI server itself, shared artifact caches)."

#### In-depth
GitHub Actions, CircleCI, and GitLab CI all use ephemeral containers by default for job runners. The trade-off: container startup time (~1-5 seconds) is paid per job. Docker layer cache mitigates this — even on a fresh container, cached layers reduce build times significantly. **Docker cache-from** with registry cache ensures even truly cold ephemeral containers get fast builds via registry-stored cache layers.

---

### 147. What is an image promotion strategy in Docker pipelines?
"Image promotion means using the **same image** across environments with environment-specific config — not rebuilding for each environment.

Pipeline: Build → Tag `app:$SHA` → Push to dev registry → Test in dev → **Promote** the same image to staging registry → Test in staging → **Promote** to prod registry → Deploy.

Promotion is a pure metadata operation: re-tag the verified image with a prod tag or copy it to the production registry. No rebuild, no risk of inconsistency. The image tested in dev is exactly what runs in prod."

#### In-depth
Registry replication tools (skopeo, crane) make promotion efficient: `skopeo copy docker://dev-registry/app:sha docker://prod-registry/app:sha` copies the image between registries without pulling/pushing locally — the layers are transferred directly between registries server-side. This is much faster than `docker pull && docker tag && docker push`. For cross-account AWS ECR promotion, `aws ecr batch-check-layer-availability` + layer upload is further optimized.

---

### 148. How do you use Docker in a monorepo setup?
"In a monorepo (multiple services in one repo), each service has its own Dockerfile.

**Challenge**: rebuilding all services on every commit is wasteful.

**Solution**: detect which services changed and build only those:
```bash
git diff --name-only HEAD~1 HEAD | grep '^services/api/' && docker build services/api/
```

Or use **Turborepo**, **Nx**, or **Bazel** which understand monorepo dependencies — if `service-a` depends on `lib-shared`, only rebuild `service-a` when `lib-shared` changes.

Docker BuildKit's multi-stage builds work well here — shared code in a common stage, service-specific code in later stages."

#### In-depth
Monorepo Docker builds benefit greatly from **BuildKit caching**. With a shared base image stage (installing common deps) referenced by multiple service stages, the shared stage is built once and cached. Services that didn't change reuse the cache completely. The key: treat the monorepo as a dependency graph, not a flat list of services. Build tooling that understands this graph (Nx, Bazel, Turborepo) provides accurate build scoping.

---

### 149. How do you test multi-service containers before pushing to prod?
"I use Docker Compose for multi-service integration testing:

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
      test: pg_isready
  test-runner:
    image: postman/newman
    command: run /tests/api.json --environment /tests/env.json
    depends_on:
      app:
        condition: service_healthy
```

Run: `docker compose up --abort-on-container-exit --exit-code-from test-runner`. Pipeline fails if test-runner exits with non-zero code."

#### In-depth
The `--exit-code-from` flag makes the exit code of the entire Compose run equal to the exit code of the specified service — enabling CI pass/fail based on test results. Clean up after: `docker compose down -v` to remove volumes and prevent state leakage. For parallel CI on multiple branches: use Compose's `--project-name` to namespace resources — `docker compose -p pr-$PR_NUMBER up/down`, isolating each PR's containers, networks, and volumes.

---

### 150. How do you manage secrets in CI/CD for Dockerized apps?
"Layered approach for CI/CD secrets:

**Build-time secrets (in Dockerfile)**: Use BuildKit secret mounts — `RUN --mount=type=secret,id=npm_token npm install`. Pass via `docker build --secret id=npm_token,src=.npmrc`. The secret is never in any image layer.

**Runtime secrets (for deployed containers)**: External secret managers (AWS SSM, Vault) → injected at container start via entrypoint script. Never in Compose files or K8s YAML.

**CI pipeline secrets**: Store in CI system's secret store (GitHub Actions secrets, GitLab CI variables) — injected as env vars only for the pipeline run, never logged."

#### In-depth
The most common CI secret mistake: printing secrets in pipeline logs. `echo $MY_SECRET` or any command that outputs the value will expose it in the logs if the CI system doesn't redact it. Modern CI systems redact known secrets from logs, but this isn't foolproof. Best practice: never echo secrets, use `docker login --password-stdin` for registry auth, and use `--secret` flags instead of `--build-arg` for any sensitive values (build args persist in image metadata).

---
