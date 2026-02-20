# 🟡 Helm & GitOps

---

### 1. What is Helm?

"Helm is the **package manager for Kubernetes** — it lets you define, install, and upgrade complex K8s applications as reusable units called **Charts**.

Instead of maintaining separate YAML files for each environment, a Helm chart templatizes them with values. You pass different `values.yaml` files for dev vs. prod.

I use Helm for everything — internal services, third-party tools (Prometheus, cert-manager, Istio), and application deployments. It's the most widely used K8s package format."

#### In Depth
Helm 3 removed Tiller (the server-side component from Helm 2) — Helm now communicates directly with the K8s API using your kubeconfig. Release state is stored in Kubernetes Secrets in the target namespace. This was a significant security improvement since Tiller ran with cluster-admin privileges by default.

---

### 2. What are Helm Charts?

"A Helm Chart is a **collection of YAML templates and a default values file** that together define a K8s application.

Chart structure:
```
my-chart/
  Chart.yaml       # metadata: name, version, dependencies
  values.yaml      # default values
  templates/       # Go templated K8s manifests
  charts/          # subcharts (dependencies)
```

When you run `helm install`, Helm renders the templates with your values and applies them to the cluster."

#### In Depth
Chart versioning follows SemVer. The `appVersion` tracks the app version, `version` tracks the chart version — they can differ (e.g., chart v3.2.0 packages app v1.9.0). Chart dependencies are declared in `Chart.yaml` under `dependencies`, downloaded with `helm dependency update`, and stored in `charts/`. Helm OCI support (GA in Helm 3.8) allows storing charts in OCI registries like ECR or GCR alongside container images.

---

### 3. What is the difference between Helm 2 and Helm 3?

"The major change: Helm 3 removed **Tiller**, the server-side component.

In Helm 2, Tiller ran in the cluster with cluster-admin permissions — a massive security risk (called 'Tiller hell'). Helm 3 uses your local kubeconfig to authenticate directly to the K8s API.

Other improvements in Helm 3: 3-way strategic merge patch (better diff handling), release namespacing (different releases can share names if in different namespaces), JSON schema validation for values, and library charts."

#### In Depth
Helm 2 is end-of-life. The migration from Helm 2 to Helm 3 was generally smooth — the `helm-2to3` plugin handles state migration. One subtle difference: Helm 3 stores release state in `helm.sh/release`-typed Secrets, not a ConfigMap. This matters for RBAC — a CI account needs `get/list/create/update/delete` on Secrets to manage releases.

---

### 4. What is a values.yaml file?

"`values.yaml` is the **central configuration file for a Helm chart**. It provides default values for all template variables.

Override values at install/upgrade time: `helm install my-app ./chart --set image.tag=1.2.3` or `-f production.values.yaml`.

I maintain separate values files per environment: `values-dev.yaml`, `values-staging.yaml`, `values-prod.yaml`. Each one overrides only the differences from the base `values.yaml`."

#### In Depth
Helm merges values files in order — later files override earlier ones. This allows layering: `helm upgrade app ./chart -f base.yaml -f prod.yaml -f hotfix.yaml`. JSON Schema for values (`values.schema.json`) validates values before rendering — catches typos and invalid types early, before Kubernetes rejects the manifest. Always add schema validation to shared charts used across teams.

---

### 5. How do you roll back a Helm deployment?

"`helm rollback my-release <REVISION>` rolls back to a specific release revision.

Find the revision with `helm history my-release`. Rollback creates a new release revision (not overwrite) — so rollback is itself auditable.

For automated rollback on health check failure, I integrate Helm with `helm-monitor` or use ArgoCD which can automatically sync back to the last healthy state."

#### In Depth
Helm's rollback does a `helm upgrade` with the previous release's chart and values. The `--cleanup-on-fail` flag (on upgrade) removes newly created resources if an upgrade fails. For completeness, also set `--atomic` during upgrades — it automatically rolls back if the upgrade fails (waiting for pods to be Ready within `--timeout`).

---

### 6. What is Helm templating?

"Helm uses Go's `text/template` with Sprig functions for templating.

Templates read values with `{{ .Values.image.tag }}`, loop with `{{ range }}`, conditionally render with `{{ if }}`, and join with `{{ include }}` for reusable partials.

I use the `_helpers.tpl` file to define shared template fragments like a standard label set or resource name formatter. This DRYs up the templates significantly."

#### In Depth
Common gotchas: YAML indentation in templates is sensitive — use `{{ . | nindent 8 }}` for multi-line string insertion. `{{ - }}` strips whitespace before the action, `{{ - }}` after. Helm's `required` function fails fast: `{{ required "image.tag is required!" .Values.image.tag }}` gives a clear error instead of an opaque K8s rejection. Use `tpl` to re-render values that themselves contain Go templates.

---

### 7. What is GitOps?

"GitOps is an **operational model where Git is the single source of truth** for your infrastructure and application configuration.

Any change to the cluster goes through a PR → merge → automated sync. The GitOps operator (ArgoCD, FluxCD) watches Git and continuously reconciles the cluster to match the desired state in Git.

The main benefits: full audit trail (every change is a Git commit with author and message), easy rollback (revert the PR), and drift detection (any manual change to the cluster is automatically corrected)."

#### In Depth
GitOps has two workflow variants: **push model** (CI pipeline runs `kubectl apply` after merge — simple but not self-healing) and **pull model** (an in-cluster operator continuously pulls from Git — self-healing, no cluster credentials in CI). Modern GitOps uses the pull model exclusively. The GitOps principles are defined at opengitops.dev: declarative, versioned, automated, continuously reconciled.

---

### 8. How does ArgoCD work?

"ArgoCD runs in the cluster and **watches a Git repository**. When the Git state diverges from the cluster state, it shows the application as `OutOfSync`. It can automatically sync (apply changes) or wait for manual approval.

ArgoCD represents each application as an `Application` CRD, tracking which Git repo/path/branch to deploy to which cluster/namespace.

I like ArgoCD's UI — you can see the full deployment graph (what resources are deployed), sync status, and health for every resource at a glance."

#### In Depth
ArgoCD's **sync waves** allow ordering resource creation: `argocd.argoproj.io/sync-wave: "-1"` for CRDs and Namespaces that must exist before everything else. **Resource hooks** (PreSync, Sync, PostSync, SyncFail) allow running Jobs at specific sync stages — e.g., a database migration job before the app rolls out. **App of Apps** pattern: one ArgoCD Application that deploys other Applications — used for bootstrapping entire environment configurations from a single Git repo.

---

### 9. What is FluxCD?

"FluxCD is GitOps operator from Weaveworks, now a CNCF Graduated project.

Flux v2 is composed of specialized controllers: **Source Controller** (watches Git/OCI/Helm), **Kustomize Controller** (applies Kustomize resources), **Helm Controller** (manages Helm releases), **Notification Controller** (sends alerts/updates PR status).

I prefer ArgoCD for its UI, but Flux's composable controller model is more Kubernetes-native and easier to customize programmatically."

#### In Depth
Flux's **Image Automation** feature updates the Git repo automatically when a new container image is pushed. The image-reflector-controller scans registries, the image-automation-controller commits the new image tag back to Git, and the reconciliation loop deploys it. This enables fully automated continuous delivery where developers just push code — everything else is automated.

---

### 10. How does GitOps differ from traditional CI/CD?

"Traditional CI/CD: Pipeline pushes changes **to** the cluster using `kubectl` or `helm` commands. Cluster state may drift from what was deployed if someone makes a manual change.

GitOps: An operator **pulls** from Git continuously. Manual changes to the cluster are automatically reverted. The cluster state always converges to what's in Git.

The paradigm shift: with GitOps, you never deploy directly to production. You make a PR → it's reviewed → approved → merged → ArgoCD deploys automatically. Every production change has a Git commit trail."

#### In Depth
The practical implications for security: GitOps eliminates the need to give CI systems (GitHub Actions, Jenkins) direct access to the production cluster's credentials. The GitOps operator inside the cluster has access, but external CI only needs access to push to Git. This dramatically reduces the attack surface — compromising CI doesn't automatically compromise production.

---

### 11. What is Kustomize?

"Kustomize is a **configuration customization tool** built into `kubectl` that allows you to patch and overlay K8s manifests without templating.

Instead of Go templates, you have a base configuration and overlay directories for each environment. You apply environment-specific patches (change image tag, adjust replicas, add a ConfigMap entry) without touching the base.

I use Kustomize for simple multi-environment configs. For complex apps with many parameters, I use Helm. ArgoCD supports both natively."

#### In Depth
Kustomize uses strategic merge patches and JSON patches. The `kustomization.yaml` file is the entry point — it lists resources, patches, images (image tag override), name prefixes/suffixes, and labels to inject. The `images` field is particularly useful: `images: [{name: nginx, newTag: 1.25.0}]` updates the tag in all manifests that reference the `nginx` image without touching the manifests directly.

---

### 12. How do you manage secrets in GitOps?

"Three approaches:

1. **Sealed Secrets**: Encrypt secrets with cluster's public key. The encrypted SealedSecret YAML is safe to commit. Sealed Secrets controller decrypts it.
2. **External Secrets Operator**: Stores secrets in Vault/AWS SM/GCP SM. The operator syncs them to K8s Secrets. Git only has the ExternalSecret CRD reference, not the actual secret value.
3. **SOPS + age/GPG**: Encrypt secrets in Git using Mozilla SOPS. Flux has native SOPS support.

I prefer External Secrets Operator in production — secrets are managed by the proper secrets management system with access logging, rotation, and versioning."

#### In Depth
The fundamental rule: **never store plaintext secrets in Git**. Even private repos have insider threat and leak risks. The External Secrets Operator is the most production-grade approach — it keeps secrets in purpose-built systems (Vault, AWS SM) that have automatic rotation, audit logging, and fine-grained access control. The K8s Secret created by ESO is ephemeral — if the external secret is deleted, the K8s Secret is garbage-collected.

---

### 13. What are Helm hooks?

"Helm hooks are **lifecycle event listeners** — they run Jobs at specific points during a Helm release lifecycle.

Examples:
- `pre-install`: Run a migration Job before installing the chart
- `post-install`: Run a smoke test after installation
- `pre-upgrade`: Take a database backup before upgrading
- `post-delete`: Clean up external resources after uninstall

I use `pre-upgrade` hooks heavily for database migrations — run the migration Job (which must succeed) before rolling out the new app version."

#### In Depth
Hook resources are deleted by default after completion (`"helm.sh/hook-delete-policy": hook-succeeded`). You can retain them with `before-hook-creation` (only delete when a new hook is run). Hook weights (integer annotations) control execution order when multiple hooks are at the same lifecycle stage. Note: hooks are **not managed** by Helm revisions — they're run-and-forget. This means a failed hook can leave a Job around that blocks future hooks.

---
