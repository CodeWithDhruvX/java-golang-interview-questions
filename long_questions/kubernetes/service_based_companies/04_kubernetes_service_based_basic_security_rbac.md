# 🏢 Kubernetes Interview Questions - Service-Based Companies (Part 4)
> **Target:** TCS, Wipro, Infosys, Cognizant, IBM, Tech Mahindra, etc.
> **Focus:** Security Concepts, RBAC (Role-Based Access Control), Resource Isolation, and basic CI/CD.

---

## 🔹 Authentication & Authorization (RBAC) (Questions 46-50)

### Question 46: What acts as the central security checkpoint in a Kubernetes cluster?

**Answer:**
The **Kube-API Server**. Before any request (like `kubectl create pod` or an internal component creating a pod) is executed, the API server passes the request through three stages:
1. **Authentication:** Who are you? (Using TLS certs, bearer tokens, OIDC).
2. **Authorization:** What are you allowed to do? (Verified via RBAC, ABAC, Webhooks).
3. **Admission Control:** How should this request be modified or validated? (Mutating/Validating Webhooks prior to writing to etcd).

---

### Question 47: Explain the difference between a User Account and a Service Account.

**Answer:**
- **User Account:** For humans. Kubernetes does not technically store User Accounts natively (no `User` objects in API). It relies on external identity providers (Active Directory, OIDC, or raw TLS certificates `CN/O` values) to verify human identities.
- **Service Account:** For Pods/Machine processes. Kube natively manages these APIs. When a Pod runs, it mounts a token associated with a specific ServiceAccount so the Pod itself can talk back to the API Server locally (like an Ingress Controller requesting endpoint updates).

---

### Question 48: What is Role-Based Access Control (RBAC)? Evaluate Role vs ClusterRole.

**Answer:**
RBAC is the mechanism K8s uses to strictly regulate access to API resources based on user roles.
- **Role:** Sets permissions *within a specific single namespace*. (Example: App Developer role in `Dev` namespace can delete pods, but cannot touch `Prod`).
- **ClusterRole:** Sets permissions *cluster-wide* non-namespaced resources (like Nodes, PVs) or extends access across *all namespaces* simultaneously.

---

### Question 49: How do you bind a Role to a User or Service Account?

**Answer:**
Using a **RoleBinding** or a **ClusterRoleBinding**.
- **RoleBinding:** Binds a Role (or ClusterRole) to a user/group strictly within the confines of a specific Namespace.
- **ClusterRoleBinding:** Binds a ClusterRole to a user/group at a global level (bypassing namespace boundary limits).

---

### Question 50: An application pod requires access to the Kubernetes API to read Secrets across the cluster. Is this safe?

**Answer:**
No, giving a Pod a ServiceAccount bound to a `ClusterRole` with `get/list Secrets` permissions across all namespaces grants it "Cluster Admin" equivalent potential privilege. If that single Pod is compromised, the attacker can extract every credential spanning the entire cluster. You must adhere strictly to the principle of least privilege.

---

## 🔹 Securing Pods & Basic CI/CD (Questions 51-55)

### Question 51: What is a Pod Security Context?

**Answer:**
The `securityContext` field allows you to define privileges and access control directly at the Pod or Container level when it executes.
It allows you to specify things like:
- Which User ID (UID/GID) the container should run as (e.g., `runAsUser: 1000` to avoid running as root).
- Setting `allowPrivilegeEscalation: false`.
- Denying root file system modification (`readOnlyRootFilesystem: true`).

---

### Question 52: By default, are Pods allowed to mount the host node's underlying filesystem?

**Answer:**
Technically yes, using the `hostPath` volume. However, because mounting sensitive areas (like `/var/run/docker.sock` or `/etc`) can completely compromise a worker node, Administrators typically employ **Pod Security Admission (PSA)** plugins to categorically reject any Pod YAML requesting privileged access or hostPaths.

---

### Question 53: What is the typical flow of deploying an application to K8s using a CI/CD pipeline like Jenkins or GitLab?

**Answer:**
1. **Continuous Integration (CI):** Developer pushes code to Git. Pipeline compiles the code, runs unit tests, builds a Docker image, and tags that image with a build ID.
2. **Push:** The new image gets pushed to a Container Registry (DockerHub, AWS ECR, ACR).
3. **Continuous Deployment/Delivery (CD):** The pipeline updates the Kubernetes YAML manifests (usually via a Helm variable or Kustomize patch) to point to the new image tag.
4. **Apply:** Pipeline runs `kubectl apply` (or uses an operator like ArgoCD) to apply the updated manifest to K8s, triggering a rolling deployment.

---

### Question 54: What is GitOps and how does it relate to Kubernetes?

**Answer:**
GitOps is a methodology where a Git repository serves as the *single source of truth* for your declarative cluster infrastructure and applications.
Instead of a CI tool "pushing" commands via `kubectl`, an operator inside the K8s cluster (like ArgoCD or Flux) constantly pulls the Git repository's main branch and applies any changes it detects to the cluster automatically.

---

### Question 55: If a container gets hacked, how do you prevent the hacker from reaching the Kube-API from inside the container?

**Answer:**
Every pod automatically mounts a default ServiceAccount token into `/var/run/secrets/kubernetes.io/serviceaccount/`. An attacker can use this token to curl the internal API server (`kubernetes.default.svc`). 
To prevent this, unless the application explicitly needs to speak to the K8s API (which most standard web apps do not), you should set `automountServiceAccountToken: false` in the Pod or ServiceAccount YAML definition.
