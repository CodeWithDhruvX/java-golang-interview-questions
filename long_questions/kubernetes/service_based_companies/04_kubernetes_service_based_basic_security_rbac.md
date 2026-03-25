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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What acts as the central security checkpoint in a Kubernetes cluster?
**Your Response:** "The Kube-API Server is the security gatekeeper of the cluster. Every request goes through three security stages: First, Authentication proves who you are using certificates or tokens. Second, Authorization checks what you're allowed to do using RBAC rules. Third, Admission Control validates and potentially modifies the request before it's saved to etcd. Think of it like airport security - first they check your ID, then they check what you're allowed to bring, and finally they inspect your bags before letting you through."

---

### Question 47: Explain the difference between a User Account and a Service Account.

**Answer:**
- **User Account:** For humans. Kubernetes does not technically store User Accounts natively (no `User` objects in API). It relies on external identity providers (Active Directory, OIDC, or raw TLS certificates `CN/O` values) to verify human identities.
- **Service Account:** For Pods/Machine processes. Kube natively manages these APIs. When a Pod runs, it mounts a token associated with a specific ServiceAccount so the Pod itself can talk back to the API Server locally (like an Ingress Controller requesting endpoint updates).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Explain the difference between a User Account and a Service Account.
**Your Response:** "User accounts are for humans - developers, admins, etc. Kubernetes doesn't actually store user accounts internally; it relies on external systems like Active Directory or LDAP to authenticate people. Service accounts are for pods and applications - they're native Kubernetes objects that give applications identity so they can talk to the API server. When a pod needs to access cluster resources, it gets a service account token mounted automatically. Think of user accounts as employee badges and service accounts as service IDs for automated processes."

---

### Question 48: What is Role-Based Access Control (RBAC)? Evaluate Role vs ClusterRole.

**Answer:**
RBAC is the mechanism K8s uses to strictly regulate access to API resources based on user roles.
- **Role:** Sets permissions *within a specific single namespace*. (Example: App Developer role in `Dev` namespace can delete pods, but cannot touch `Prod`).
- **ClusterRole:** Sets permissions *cluster-wide* non-namespaced resources (like Nodes, PVs) or extends access across *all namespaces* simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Role-Based Access Control (RBAC)? Evaluate Role vs ClusterRole.
**Your Response:** "RBAC is how Kubernetes controls who can do what. It's like giving different people different levels of access. A Role is namespace-scoped - it only gives permissions within one namespace, like allowing developers to manage pods in the dev namespace but not in production. A ClusterRole is cluster-wide - it can access non-namespaced resources like nodes and persistent volumes, or grant permissions across all namespaces. I use Roles for application-specific permissions and ClusterRoles for administrative tasks that need cluster-level access."

---

### Question 49: How do you bind a Role to a User or Service Account?

**Answer:**
Using a **RoleBinding** or a **ClusterRoleBinding**.
- **RoleBinding:** Binds a Role (or ClusterRole) to a user/group strictly within the confines of a specific Namespace.
- **ClusterRoleBinding:** Binds a ClusterRole to a user/group at a global level (bypassing namespace boundary limits).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you bind a Role to a User or Service Account?
**Your Response:** "I use RoleBindings and ClusterRoleBindings. A RoleBinding connects a role to a user or service account within a specific namespace - like giving a developer access to the dev namespace only. A ClusterRoleBinding is global - it connects a cluster role to a user across the entire cluster. The key difference is scope: RoleBindings are namespace-scoped, ClusterRoleBindings are cluster-wide. I always prefer RoleBindings for the principle of least privilege - only use ClusterRoleBindings when absolutely necessary for cluster-wide operations."

---

### Question 50: An application pod requires access to the Kubernetes API to read Secrets across the cluster. Is this safe?

**Answer:**
No, giving a Pod a ServiceAccount bound to a `ClusterRole` with `get/list Secrets` permissions across all namespaces grants it "Cluster Admin" equivalent potential privilege. If that single Pod is compromised, the attacker can extract every credential spanning the entire cluster. You must adhere strictly to the principle of least privilege.

### How to Explain in Interview (Spoken style format)
**Interviewer:** An application pod requires access to the Kubernetes API to read Secrets across the cluster. Is this safe?
**Your Response:** "Absolutely not! Giving a pod a ClusterRole that can read secrets across all namespaces is like giving it the keys to the entire kingdom. If that single pod gets compromised, an attacker can steal every secret in the cluster - database passwords, API keys, certificates, everything. I'd follow the principle of least privilege: give the pod access only to the specific secrets it needs in its own namespace, or even better, use a secret management solution like HashiCorp Vault. Cluster-wide secret access is a major security risk."

---

## 🔹 Securing Pods & Basic CI/CD (Questions 51-55)

### Question 51: What is a Pod Security Context?

**Answer:**
The `securityContext` field allows you to define privileges and access control directly at the Pod or Container level when it executes.
It allows you to specify things like:
- Which User ID (UID/GID) the container should run as (e.g., `runAsUser: 1000` to avoid running as root).
- Setting `allowPrivilegeEscalation: false`.
- Denying root file system modification (`readOnlyRootFilesystem: true`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a Pod Security Context?
**Your Response:** "The security context is like setting security rules for each container. It lets me control things like what user the container runs as - I can specify a non-root user ID instead of running as root. I can prevent privilege escalation, make the filesystem read-only, and set other security constraints. It's crucial for hardening containers because by default many containers run as root, which is dangerous. The security context is my first line of defense for making sure containers run with minimal privileges."

---

### Question 52: By default, are Pods allowed to mount the host node's underlying filesystem?

**Answer:**
Technically yes, using the `hostPath` volume. However, because mounting sensitive areas (like `/var/run/docker.sock` or `/etc`) can completely compromise a worker node, Administrators typically employ **Pod Security Admission (PSA)** plugins to categorically reject any Pod YAML requesting privileged access or hostPaths.

### How to Explain in Interview (Spoken style format)
**Interviewer:** By default, are Pods allowed to mount the host node's underlying filesystem?
**Your Response:** "Technically yes with hostPath volumes, but this is extremely dangerous and usually blocked. A pod could mount sensitive system directories like `/etc` or even the Docker socket, which would give it complete control over the worker node. That's why most production clusters use Pod Security Admission policies to reject any pod trying to use hostPath or other privileged features. It's like giving someone the keys to the server room - technically possible but a huge security risk that should be avoided unless absolutely necessary."

---

### Question 53: What is the typical flow of deploying an application to K8s using a CI/CD pipeline like Jenkins or GitLab?

**Answer:**
1. **Continuous Integration (CI):** Developer pushes code to Git. Pipeline compiles the code, runs unit tests, builds a Docker image, and tags that image with a build ID.
2. **Push:** The new image gets pushed to a Container Registry (DockerHub, AWS ECR, ACR).
3. **Continuous Deployment/Delivery (CD):** The pipeline updates the Kubernetes YAML manifests (usually via a Helm variable or Kustomize patch) to point to the new image tag.
4. **Apply:** Pipeline runs `kubectl apply` (or uses an operator like ArgoCD) to apply the updated manifest to K8s, triggering a rolling deployment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the typical flow of deploying an application to K8s using a CI/CD pipeline like Jenkins or GitLab?
**Your Response:** "The flow starts when a developer pushes code to Git. The CI pipeline kicks in - compiles the code, runs tests, builds a Docker image, and pushes it to a registry with a unique tag. Then the CD part takes over - the pipeline updates the Kubernetes manifests to point to the new image tag, usually using Helm or Kustomize. Finally, it runs `kubectl apply` to deploy the changes, triggering a rolling update. The whole process is automated from code commit to production deployment."

---

### Question 54: What is GitOps and how does it relate to Kubernetes?

**Answer:**
GitOps is a methodology where a Git repository serves as the *single source of truth* for your declarative cluster infrastructure and applications.
Instead of a CI tool "pushing" commands via `kubectl`, an operator inside the K8s cluster (like ArgoCD or Flux) constantly pulls the Git repository's main branch and applies any changes it detects to the cluster automatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is GitOps and how does it relate to Kubernetes?
**Your Response:** "GitOps treats Git as the source of truth for everything in your Kubernetes cluster. Instead of having CI pipelines push changes to the cluster, you have operators inside the cluster like ArgoCD or Flux that constantly pull from Git and apply any changes they find. This means the Git repo always reflects what's actually running in the cluster. It's like making Git the brain and the cluster the body - the body always does what the brain says. This gives you audit trails, version control, and the ability to restore your entire cluster configuration from Git."

---

### Question 55: If a container gets hacked, how do you prevent the hacker from reaching the Kube-API from inside the container?

**Answer:**
Every pod automatically mounts a default ServiceAccount token into `/var/run/secrets/kubernetes.io/serviceaccount/`. An attacker can use this token to curl the internal API server (`kubernetes.default.svc`). 
To prevent this, unless the application explicitly needs to speak to the K8s API (which most standard web apps do not), you should set `automountServiceAccountToken: false` in the Pod or ServiceAccount YAML definition.

### How to Explain in Interview (Spoken style format)
**Interviewer:** If a container gets hacked, how do you prevent the hacker from reaching the Kube-API from inside the container?
**Your Response:** "By default, every pod gets a service account token mounted automatically, which an attacker could use to talk to the Kubernetes API. To prevent this, I set `automountServiceAccountToken: false` unless the application actually needs to access the API. Most web applications don't need to talk to the Kubernetes API, so there's no reason to give them that capability. It's like removing the master key from a hotel room - if someone breaks in, they can't use the room key to access other parts of the hotel. This is a simple but effective security hardening step."
