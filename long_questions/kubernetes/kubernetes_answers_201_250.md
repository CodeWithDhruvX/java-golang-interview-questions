## 🔹 Core Concepts Deep Dive (Questions 201-210)

### Question 201: What is a pod lifecycle?

**Answer:**
The sequence of states a Pod goes through:
1.  **Pending:** Accepted by API, but container image not pulled yet.
2.  **Running:** Bound to node, all containers created, at least one running.
3.  **Succeeded:** All containers terminated with exit code 0 (Completed).
4.  **Failed:** All terminated, at least one with error.
5.  **Unknown:** State cannot be obtained.

---

### Question 202: What are the phases of a pod?

**Answer:**
Same as lifecycle phases: **Pending, Running, Succeeded, Failed, Unknown**.
- Note: `Terminating` is not a Phase, it's a condition where `deletionTimestamp` is set.

---

### Question 203: What is the difference between `status.phase` and `status.conditions`?

**Answer:**
- **Phase:** High-level summary (e.g., "Running").
- **Conditions:** Array of detailed boolean states.
  - `PodScheduled`: True
  - `Initialized`: True (Init containers finished)
  - `ContainersReady`: True
  - `Ready`: True (Passed readiness probe)
A pod can be Phase=Running but Ready=False.

---

### Question 204: How does `kubectl get events` help in debugging?

**Answer:**
Events show the chronological log of state changes.
- "Scheduled to Node X"
- "Pulling Image"
- "Failed to Pull Image (Auth Error)"
- "Liveness Probe Failed"
It tells you *why* a pod is in a certain state (Pending/CrashLoop).

---

### Question 205: How does Kubernetes handle pod eviction?

**Answer:**
**Eviction** = Forcing a pod to terminate.
**Triggers:**
1.  **Node Pressure:** Node is out of RAM/Disk.
2.  **Taint Based:** Node is tainted (e.g., NetworkUnavailable) and pod doesn't tolerate.
3.  **API Eviction:** User runs `kubectl drain`.

---

### Question 206: What triggers a pod to be terminated?

**Answer:**
1.  User deletes the Pod/Deployment.
2.  Job completes.
3.  Liveness Probe fails repeatedly (Restart policy determines if it stops).
4.  OOMKilled.
5.  Node Eviction.

---

### Question 207: What happens during pod termination?

**Answer:**
1.  API sets `deletionTimestamp`.
2.  Endpoint Controller removes Pod IP from Service Endpoints (Stops new traffic).
3.  Kubelet runs `preStop` hook.
4.  Kubelet sends `SIGTERM` to main process (PID 1).
5.  Waits for `terminationGracePeriod` (default 30s).
6.  Force kills with `SIGKILL`.

---

### Question 208: How does graceful shutdown work in Kubernetes?

**Answer:**
The application must handle `SIGTERM`.
- Code: `process.on('SIGTERM', () => server.close())`.
- It should stop accepting new requests, finish processing current in-flight requests, close DB connections, then exit.

---

### Question 209: What is `terminationGracePeriodSeconds`?

**Answer:**
A setting in PodSpec (default 30s).
- The time K8s waits between sending `SIGTERM` and `SIGKILL`.
- Can be extended (`60s`) for apps that need long cleanup times.

---

### Question 210: What is a liveness probe vs readiness probe?

**Answer:**
- **Liveness Probe:** "Are you alive?"
  - Fail -> **Restart Pod**.
  - Use case: Deadlock, Crash.
- **Readiness Probe:** "Can you accept traffic?"
  - Fail -> **Remove from Service**.
  - Use case: Warming up cache, Loading large data.

---

## 🔹 Controllers & Workloads (Questions 211-220)

### Question 211: What is a DeploymentSet? Is it a real object?

**Answer:**
No, there is no object called "DeploymentSet".
- There is **Deployment**.
- There is **DaemonSet**.
- There is **StatefulSet**.
- There is **ReplicaSet**.
*Trick question.*

---

### Question 212: What happens when a Deployment is deleted?

**Answer:**
By default, the Deployment Controller deletes the **ReplicaSets** and **Pods** it manages (Cascading deletion).
- Use `--cascade=orphan` to leave pods running.

---

### Question 213: How does the ReplicaSet know to scale?

**Answer:**
It looks at the `spec.replicas` field.
- If Current < Desired -> Create Pods.
- If Current > Desired -> Delete Pods.
- It doesn't "know" load/traffic. HPA modifies the `spec.replicas` field externally.

---

### Question 214: What is a rollout history in Kubernetes?

**Answer:**
Deployments keep old **ReplicaSets**.
- Each Revision matches an old RS.
- `revisionHistoryLimit` (default 10) determines how many old RS are kept for rollback.

---

### Question 215: How can you resume a paused Deployment?

**Answer:**
`kubectl rollout resume deployment/my-app`.
- This triggers the controller to reconcile the new state (apply the changes that accumulated while paused).

---

### Question 216: What are the alternatives to StatefulSets?

**Answer:**
1.  **Deployment + Single Replica + PVC:** Works if you only need 1 instance.
2.  **Operators:** Use a specific DB Operator (Postgres Operator) which might manage Pods directly for complex logic (failover/leader election).

---

### Question 217: Why would you use a DaemonSet instead of a Deployment?

**Answer:**
If you need **Exactly One** pod per node.
- **Deployment:** "Run 5 copies, I don't care where."
- **DaemonSet:** "Run copy on Node A, Node B, Node C..."
- Essential for infrastructure (logs, monitoring, storage drivers).

---

### Question 218: What happens to DaemonSet pods when a node is drained?

**Answer:**
They are **ignored** (unless drain command has `--ignore-daemonsets`).
- They are not evicted because they are local to the node.
- If node is actually removed/deleted, the DS pod is deleted.

---

### Question 219: How can you run a one-time job on a schedule?

**Answer:**
Use a **CronJob**.
- Defines a `jobTemplate`.
- At schedule time, K8s creates a **Job**.
- The Job creates a **Pod**.
- Pod runs to completion.

---

### Question 220: Can you use Jobs for parallel processing?

**Answer:**
Yes.
1.  `completions: 10`: Run until 10 succeed.
2.  `parallelism: 3`: Run 3 pods at a time.
Useful for processing a queue of work items.

---

## 🔹 Custom Resources & Extensibility (Questions 221-230)

### Question 221: What is a CRD?

**Answer:**
**Custom Resource Definition.**
- Extends K8s API with your own types.
- Example: Defining a `Pizzas` resource.
- `kubectl get pizzas`.

---

### Question 222: How do you create a CRD?

**Answer:**
Create a YAML file with `kind: CustomResourceDefinition`.
- Define names (`singular`, `plural`, `kind`).
- Define schema (`openAPIV3Schema`) for validation (e.g., "size" must be "small/medium/large").

---

### Question 223: How do you manage CRDs with Helm?

**Answer:**
Helm has a special `crds/` folder.
- CRDs in this folder are installed **before** the other templates.
- **Limitation:** Helm generally does NOT upgrade/delete CRDs automatically to prevent data loss.

---

### Question 224: What is a Custom Controller?

**Answer:**
A piece of code (Go/Python) running in the cluster that watches your Custom Resource.
- **CRD** = Data Model.
- **Controller** = Logic.
- Without a controller, a CRD is just a text record in etcd.

---

### Question 225: What are the use cases of an Operator?

**Answer:**
1.  **Database Management:** Backups, upgrades, failover (Prometheus, Postgres).
2.  **Service Mesh:** Istio Operator.
3.  **App Lifecycle:** "Day 2" operations.

---

### Question 226: What is Kubebuilder?

**Answer:**
An SDK (framework) for building Kubernetes APIs and Controllers using Go.
- Scaffolds the project structure.
- Generates YAMLs.
- Built on top of `controller-runtime`.

---

### Question 227: What is the difference between a controller and an operator?

**Answer:**
- **Controller:** Generic term. Any control loop (Native or Custom).
- **Operator:** A specific pattern. A Custom Controller that uses CRDs to manage **Application-Specific Operational Knowledge**.
- All Operators are Controllers. Not all Controllers are Operators.

---

### Question 228: How do you version your CRDs?

**Answer:**
CRDs support multiple versions (`v1alpha1`, `v1beta1`, `v1`).
- You can serve multiple versions simultaneously.
- Use **Conversion Webhooks** to translate data between versions on the fly.

---

### Question 229: What is apiextensions.k8s.io?

**Answer:**
The API Group where CRDs live.
- `apiVersion: apiextensions.k8s.io/v1`.

---

### Question 230: How do you validate CRD schemas?

**Answer:**
Inside the CRD YAML, use `openAPIV3Schema`.
- `type: integer`
- `minimum: 1`
- `enum: ["blue", "green"]`
The API Server rejects requests that don't match.

---

## 🔹 Scheduling & Affinities (Questions 231-240)

### Question 231: What is inter-pod affinity?

**Answer:**
Rules about how pods should be placed **relative to other pods**.
- "Run Pod A only on nodes where Pod B is running."
- Use Case: Latency sensitive communication between App and Cache.

---

### Question 232: What is pod anti-affinity?

**Answer:**
Rules to keep pods apart.
- "Do not run Pod A on a node if Pod A is already there."
- Use Case: High Availability (Spreading replicas across nodes/zones).

---

### Question 233: What are node affinity types?

**Answer:**
1.  **requiredDuringSchedulingIgnoredDuringExecution:** (Hard). Must match or Pod stays Pending.
2.  **preferredDuringSchedulingIgnoredDuringExecution:** (Soft). Try to match, but schedule anyway if not possible.

---

### Question 234: What does `preferredDuringSchedulingIgnoredDuringExecution` mean?

**Answer:**
It is a "Best Effort" placement.
- If Node A matches the preference, score it higher.
- If no node matches, schedule on Node B anyway.
- "IgnoredDuringExecution" means if the node changes later (label removed), the pod is **not evicited**.

---

### Question 235: How does taint-based eviction work?

**Answer:**
Using `NoExecute` effect.
- If a node gets tainted (e.g., `node.kubernetes.io/unreachable`), pods that do not tolerate this taint are evicted.
- Enables automatic cleaning of broken nodes.

---

### Question 236: What are tolerations used for?

**Answer:**
To allow (but not require) a Pod to schedule onto a node with a matching Taint.
- Taint: "Dedicated=GPU".
- Normal pods: Blocked.
- GPU pods (with Toleration): Allowed.

---

### Question 237: How can you ensure a pod always runs on a specific node?

**Answer:**
1.  **nodeSelector:** (Simple) `nodeSelector: { hostname: node1 }`.
2.  **Node Affinity (Required):** (Advanced) `key: kubernetes.io/hostname, operator: In, values: [node1]`.

---

### Question 238: What is a nodeSelector?

**Answer:**
Simplest form of node constraints.
- Map of Key-Value pairs.
- Node must have these labels to host the pod.

---

### Question 239: What is a pod topology spread constraint?

**Answer:**
Newer API (`topologySpreadConstraints`).
- Controls how Pods are spread across your cluster among failure-domains such as regions, zones, nodes.
- "Max distinct difference" (skew).
- Replaces complex Anti-Affinity rules for spreading.

---

### Question 240: How does the scheduler prioritize nodes?

**Answer:**
After filtering, it runs **Scoring Functions**.
- `ImageLocality`: Node has image? (+Score).
- `LeastRequested`: Node is empty? (+Score for spread).
- `NodeAffinity`: Matches preference? (+Score).
Sums detailed scores and picks winner.

---

## 🔹 Config & Secrets Advanced (Questions 241-250)

### Question 241: What are subPaths in volume mounts?

**Answer:**
Allows sharing a volume for multiple mount points.
- Mount `vol-A/mysql` to `/var/lib/mysql`.
- Mount `vol-A/api` to `/var/log/api`.
- Without subPath, mounting a volume usually obscures the whole directory.

---

### Question 242: How do you use environment variables from ConfigMaps?

**Answer:**
```yaml
envFrom:
  - configMapRef:
      name: my-config
```
Injects all key-values in ConfigMap as environment variables.

---

### Question 243: How do you update a Secret without restarting pods?

**Answer:**
If mounted as a **Volume**, K8s updates the file content automatically.
- The Application must watch the file for changes (Hot Reload).
- If ConfigMap is used as ENV VAR, you MUST restart the pod.

---

### Question 244: What is the difference between stringData and data in Secrets?

**Answer:**
- **data:** Requires Base64 encoded strings.
- **stringData:** Write-only field accepting Plain Text.
  - When submitted, API Server automatically converts it to Base64 and moves it to `data` field. Convenience feature.

---

### Question 245: How can you encrypt Secrets at rest?

**Answer:**
Configure **EncryptionConfiguration** in the API Server arguments.
- Use a provider: `aescbc`, `secretbox`, or `kms` (Key Management Service).
- Ensures raw data in etcd is unreadable.

---

### Question 246: What is KMS integration in Kubernetes?

**Answer:**
Using an external Key Management Service (AWS KMS, Google Cloud KMS) to encrypt the encryption keys.
- K8s asks KMS to decrypt the DEK (Data Encryption Key) to read secrets.
- Most secure method (Envelope Encryption).

---

### Question 247: How do you manage secrets in Git securely?

**Answer:**
**Sealed Secrets (Bitnami)**.
- Public Key is in cluster.
- Developer encrypts secret with Public Key offline.
- Commits "SealedSecret" CRD to Git.
- Controller in cluster decrypts with Private Key.

---

### Question 248: What are sealed secrets?

**Answer:**
A custom resource (`kind: SealedSecret`) that contains encrypted secret data. Only the controller running in the cluster can decrypt it.

---

### Question 249: What is HashiCorp Vault and how does it integrate with Kubernetes?

**Answer:**
Enterprise Secret Management.
- **Integration:** Vault Agent Injector.
- It injects a sidecar that authenticates with Vault and renders secrets to a shared memory volume (`/vault/secrets`).
- App reads file. No K8s Secrets used.

---

### Question 250: What is external-secrets?

**Answer:**
**External Secrets Operator (ESO).**
- Syncs secrets from external APIs (AWS Secrets Manager, Azure Key Vault) **into** Kubernetes Secrets.
- Useful if you want to use native K8s Secrets (env vars) but store master data in Cloud Manager.

---
