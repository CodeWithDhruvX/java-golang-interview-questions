# 🔴 ConfigMaps, Secrets & Security

---

### 1. What is a ConfigMap?

"A ConfigMap stores **non-sensitive configuration data** as key-value pairs that can be injected into Pods as environment variables, command-line arguments, or mounted files.

It decouples configuration from your container images. The same image can run in dev, staging, and prod by just pointing to different ConfigMaps.

I use ConfigMaps for app settings, feature flags, and config files like `nginx.conf` or `application.properties`."

#### In Depth
ConfigMaps have a size limit of **1MB**. For larger config files, consider mounting them from a ConfigMap volume or using a config server. When a ConfigMap is mounted as a volume, updates to the ConfigMap are eventually reflected in the mounted files (with a delay based on kubelet's sync period, typically ~60s). Environment variables from ConfigMaps are NOT updated automatically — they require a pod restart.

---

### 2. What is a Secret?

"A Secret stores **sensitive data** like passwords, API keys, and TLS certificates.

It's Base64-encoded (not encrypted by default) and can be mounted as volumes or exposed as environment variables. Access is controlled via RBAC.

I always pair Secrets with RBAC policies that limit which ServiceAccounts can read them. For production, I go further: use **external secret management** like HashiCorp Vault or AWS Secrets Manager via the External Secrets Operator."

#### In Depth
The critical misconception: **Base64 is encoding, not encryption**. Anyone with `kubectl get secret -o yaml` can decode it trivially. For real security, enable **encryption at rest** in etcd using KMS integration (`--encryption-provider-config`) and restrict RBAC access to Secrets. The Kubernetes documentation explicitly warns against using Secrets as your only security control.

---

### 3. How do you mount a ConfigMap to a Pod?

"Two ways: as **environment variables** or as a **mounted volume**.

As env var: `envFrom: - configMapRef: name: my-config` injects all keys as environment variables.

As volume: mount the ConfigMap as a directory — each key becomes a file. I use the volume approach for config files like `application.yaml` because it's auto-updated when the ConfigMap changes (with some delay)."

#### In Depth
You can also mount **specific keys** using `items` in the volume spec, mapping a ConfigMap key to a specific file path. Using `subPath` mounts a single file without overwriting the entire directory. `subPath` mounts do NOT get updated when the ConfigMap is changed — a pod restart is required. This is a common footgun.

---

### 4. How are Secrets stored in Kubernetes?

"By default, Secrets are stored in **etcd as Base64-encoded, unencrypted data**.

This is why etcd security is critical: restrict access to the etcd endpoint, enable TLS for etcd communication, and most importantly — **enable encryption at rest**.

Kubernetes supports encrypting Secrets at rest using an `EncryptionConfiguration` that specifies a KMS provider (AWS KMS, GCP KMS, Azure Key Vault) or an AES key. With KMS, the data encryption key is itself encrypted by the KMS service."

#### In Depth
Even with encryption at rest, the API server decrypts secrets before serving them. If someone has `get secret` RBAC access, they get the cleartext value. True zero-trust secret management requires: (1) RBAC limiting who can `get` secrets, (2) audit logging of all secret access, (3) using a Secrets store CSI driver so secrets are mounted directly from Vault/AWS SM rather than stored in etcd at all.

---

### 5. How do you update a ConfigMap without restarting a Pod?

"For **volume-mounted ConfigMaps**, updates are propagated automatically within the kubelet sync period (~60s by default).

For **environment variable ConfigMaps**, you must restart the pods — env vars are set at container startup.

In practice, I use a combination: version the ConfigMap name (e.g., `app-config-v3`) and update the Deployment to reference it, triggering a controlled rolling restart."

#### In Depth
The automatic volume update uses the kubelet's `configmap` cache sync. The update uses an atomic swap — a symlink is updated to point to the new directory — so applications reading the files won't see a partial update. Applications that watch for file changes (using inotify or fsnotify) can get hot-reloaded config without restarts.

---

### 6. What is RBAC in Kubernetes?

"RBAC (Role-Based Access Control) controls **who can do what** in Kubernetes.

The key objects: **Role/ClusterRole** (define a set of permissions), **RoleBinding/ClusterRoleBinding** (grant a Role to a Subject — user, group, or ServiceAccount).

In production, I follow least-privilege: developers get `get`, `list`, `watch` in their namespace; CI service accounts get `apply` and `delete`; no one gets `cluster-admin` except break-glass accounts with MFA."

#### In Depth
ClusterRole vs. Role: a Role is namespace-scoped, a ClusterRole is cluster-wide. A ClusterRole can be bound in a specific namespace via RoleBinding — this is how you define a "standard developer" ClusterRole and bind it selectively to namespaces. ClusterRoleBinding grants the role cluster-wide. Aggregated ClusterRoles allow composing RBAC rules from multiple labeled roles — used by the Kubernetes built-in `view`, `edit`, `admin` roles.

---

### 7. What are ServiceAccounts?

"A ServiceAccount provides a **Kubernetes identity for processes running inside pods**.

By default, every pod gets the `default` ServiceAccount. Controllers, operators, and CI bots need specific ServiceAccounts with specific permissions.

I always create dedicated ServiceAccounts for applications that need API access (e.g., an app that reads ConfigMaps or creates Jobs). The `default` SA typically has no permissions in well-secured clusters."

#### In Depth
Since K8s 1.21, ServiceAccount tokens are **time-limited and audience-bound** (projected service account tokens). The old auto-mounted long-lived tokens are deprecated. For workload identity in cloud providers (AWS IRSA, GKE Workload Identity), you annotate the ServiceAccount with the cloud IAM role ARN/email, and the cloud SDK automatically exchanges the K8s token for cloud credentials.

---

### 8. How does Kubernetes handle authentication?

"Kubernetes supports multiple authentication mechanisms: **X.509 client certificates**, **Bearer tokens** (ServiceAccount or OIDC), **static token files**, and **authentication webhook**s.

In production, I use **OIDC** (via tools like Dex, Google Identity, Okta). Users authenticate with SSO, get an OIDC token, and configure kubectl with that token. The API server validates the token against the OIDC provider.

For CI systems, I use ServiceAccount tokens with fine-grained RBAC."

#### In Depth
K8s has no built-in user management — there's no `kubectl create user`. Users are identified by their certificate's CN field or OIDC claims. `kubeconfig` files contain the credentials. For multi-cluster access, tools like `kubie`, `kubectx`, or Rancher's Lens handle context switching. Kubernetes audit logs record every authenticated request — essential for compliance.

---

### 9. What are NetworkPolicies used for?

"NetworkPolicies act as **L3/L4 firewalls for pods** — controlling which pods can communicate with which other pods or external services based on labels and ports.

Without NetworkPolicies, any pod can reach any other pod in the cluster. This is unacceptable for multi-tenant environments or services handling sensitive data.

A production pattern I always implement: namespace-level default-deny, then explicit allow rules for known communication paths. For example, only the API gateway pods can reach the auth-service on port 8080."

#### In Depth
NetworkPolicies are enforced by the CNI plugin — Calico, Cilium, AWS VPC CNI (with Calico for policy). They are **additive and stateful**: if you allow TCP ingress on port 80, the return traffic is automatically allowed. Policies select pods via `podSelector` (within namespace), `namespaceSelector` (from other namespaces), or `ipBlock` (for external CIDRs). There's no deny syntax in native NetworkPolicy — use Calico's GlobalNetworkPolicy for deny rules.

---

### 10. How do you restrict access to the API server?

"Multiple layers of defense:

1. **Private endpoint**: Put the API server in a private subnet, accessible only via VPN or bastion.
2. **Authorized networks**: `--apiserver-authorized-networks` in managed K8s to whitelist CIDRs.
3. **Authentication**: Disable anonymous auth, enforce OIDC or certificates.
4. **RBAC**: Never use cluster-admin in day-to-day operations.
5. **Audit logging**: Log all API server requests to a SIEM for anomaly detection.

I've also set up API server audit webhooks that automatically alert on suspicious patterns like accessing secrets at odd hours."

#### In Depth
The API server's `--anonymous-auth=false` flag is critical. By default, anonymous requests get the `system:anonymous` user with `system:unauthenticated` group — if any ClusterRoleBinding grants permissions to this group, you have an unauthenticated attack surface. Also restrict `kubectl exec` and `port-forward` permissions via RBAC — these are powerful escape hatches.

---

### 11. What is PodSecurityAdmission (PSA)?

"PSA replaced PodSecurityPolicies in K8s 1.25+. It enforces security standards at the namespace level via labels.

Three enforcement levels:
- **Privileged**: No restrictions. Only for K8s system components.
- **Baseline**: Minimum restrictions, blocks known privilege escalations. Most workloads.
- **Restricted**: Hardened, disables most attack vectors. Best for sensitive workloads.

I label production namespaces with `pod-security.kubernetes.io/enforce: restricted` and `pod-security.kubernetes.io/warn: restricted` on staging namespaces."

#### In Depth
PSA can run in three modes per level: `enforce` (reject), `audit` (log but allow, visible in audit logs), `warn` (user-visible warning but allow). A safe migration path is: set `audit` and `warn` first to see what would break, fix the violating pod specs, then switch to `enforce`. For more fine-grained policy than PSA offers, layer Kyverno or OPA Gatekeeper on top.

---

### 12. What is container security context?

"The security context defines **privilege and access control settings** for a pod or individual container.

Key settings I always configure:
- `runAsNonRoot: true` — prevents running as root
- `runAsUser: 1000` — specific non-root UID
- `readOnlyRootFilesystem: true` — no writes to container filesystem
- `allowPrivilegeEscalation: false` — no sudo, no setuid binaries

These settings reduce the blast radius if a container is compromised."

#### In Depth
`capabilities` allow fine-grained Linux capability control. Instead of running as root (which has all capabilities), drop all capabilities and add back only what's needed: `capabilities: {drop: [ALL], add: [NET_BIND_SERVICE]}` for a web server that needs to bind port 80. `seccompProfile: RuntimeDefault` applies a seccomp filter that blocks unusual syscalls — this is a strong defense against kernel exploits.

---

### 13. What are sealed secrets?

"Sealed Secrets (Bitnami's solution) allows you to **safely commit encrypted secrets to Git**.

You run `kubeseal` to encrypt a Secret using the cluster's public key. The encrypted SealedSecret resource is safe to commit. The Sealed Secrets controller in the cluster decrypts it using the private key and creates the actual K8s Secret.

This is the native-K8s approach to GitOps secret management. The alternative is to not store secrets in Git at all, using External Secrets Operator with Vault or AWS Secrets Manager instead."

#### In Depth
The sealed secret is tied to a specific cluster by default (the encryption uses the cluster's certificate). If you rotate the cluster or move to a new one, you need to re-seal all secrets. You can configure `--scope=cluster-wide` to use a shared key for easier portability. A risk: if the controller's private key is compromised, all sealed secrets are exposed — rotate the key regularly.

---

### 14. What is HashiCorp Vault and how does it integrate with Kubernetes?

"Vault is an enterprise-grade **secrets management and PKI system**. It stores secrets with versioning, dynamic secret generation (short-lived DB credentials), and detailed audit logging.

Integration options:
1. **Vault Agent Sidecar**: An init container fetches secrets from Vault and writes them to a shared volume. The app reads files.
2. **Secrets Store CSI Driver**: Mounts Vault secrets directly as pod volumes without needing a sidecar.
3. **External Secrets Operator**: Syncs Vault secrets to K8s Secrets automatically.

I use the CSI approach in production — secrets never touch etcd, and rotation is handled automatically."

#### In Depth
Vault's **dynamic secrets** are a game-changer for databases: Vault generates a unique username/password on demand with a 1-hour TTL. After the TTL, Vault revokes the credentials. This means even if credentials leak, the attack window is tiny. Vault's Kubernetes auth method uses the pod's ServiceAccount token to authenticate and get a Vault token back.

---

### 15. How do you enforce non-root containers?

"Three layers of enforcement:

1. **Security context at the Deployment level**: `securityContext.runAsNonRoot: true`
2. **PSA**: namespace label `pod-security.kubernetes.io/enforce: restricted` — which requires non-root.
3. **OPA Gatekeeper/Kyverno policy**: A policy that rejects any pod spec running as UID 0.

I test images at build time too: tools like `hadolint` warn on `USER root` in Dockerfiles, and Trivy scans for images that default to running as root."

#### In Depth
Even with `runAsNonRoot: true`, if your container image's `USER` directive isn't set, the spec is rejected at runtime with: `container has runAsNonRoot and image will run as root`. The fix is to set `runAsUser: 1000` explicitly. Also, files in the container's filesystem need to be readable by the non-root user — permission issues are a common gotcha when enforcing non-root.

---
