## 🔹 Kubernetes & GPUs / ML Workloads (Questions 601-610)

### Question 601: How do you schedule GPU workloads in Kubernetes?

**Answer:**
Requires **Device Plugins**.
- Install NVIDIA Device Plugin DaemonSet.
- It exposes resource `nvidia.com/gpu` on nodes.
- Pod Request:
```yaml
resources:
  limits:
    nvidia.com/gpu: 1
```

---

### Question 602: What is the NVIDIA device plugin?

**Answer:**
A software component running on nodes.
- Talks to Kubelet (via Device Plugin API).
- Advertises GPU availability (Count, Memory) to the Scheduler.
- Mounts GPU drivers/libraries into the container at start.

---

### Question 603: How do you limit GPU usage to specific namespaces?

**Answer:**
**ResourceQuota.**
```yaml
apiVersion: v1
kind: ResourceQuota
spec:
  hard:
    nvidia.com/gpu: "4"
```

---

### Question 604: What is the `nvidia.com/gpu` resource?

**Answer:**
Extended Resource.
- Unlike CPU/RAM (Core types), this is an integer-only resource.
- Containers cannot share fractional GPUs (unless using Multi-Instance GPU (MIG) tech).

---

### Question 605: How do you manage mixed CPU/GPU workloads?

**Answer:**
Use **Taints and Tolerations**.
- Taint GPU nodes: `accelerator=nvidia:NoSchedule`.
- Only ML pods tolerate it.
- Prevents NGINX from wasting expensive GPU nodes.

---

### Question 606: What is Kubeflow?

**Answer:**
The ML Toolkit for Kubernetes.
- Bundles Jupyter Notebooks, Training Operators (TFJob), Serving (KServe), and Pipelines.
- Standardizes "MLOps" on K8s.

---

### Question 607: What’s the difference between TensorFlowJob and PyTorchJob?

**Answer:**
CRDs provided by **Kubeflow Training Operator**.
- They manage distributed training.
- **TFJob:** Parameter Server strategy.
- **PyTorchJob:** Master/Worker rank strategy.

---

### Question 608: How do you manage ML model versioning in Kubernetes?

**Answer:**
- **Model Registry:** MLFlow / S3.
- **Serving:** KServe (ModelMesh).
- **InferenceService CRD:** `predictor: sklearn, storageUri: s3://models/v1`.

---

### Question 609: What is KServe (formerly KFServing)?

**Answer:**
Model Serving framework.
- Abstracts complexity of "Deployment + Service + Ingress + Autoscaling".
- Provides **Scale-to-Zero** (Serverless inference).

---

### Question 610: How do you scale ML inference workloads with K8s?

**Answer:**
**KEDA or KServe.**
- Scale based on `request_per_second` or `gpu_utilization`.
- Unlike basic apps, model loading takes time, so predictive scaling is often needed.

---

## 🔹 GitOps & DevOps Practices (Questions 611-620)

### Question 611: How do you roll back changes in ArgoCD?

**Answer:**
- **UI/CLI:** `argocd app rollback <appname> <revision>`.
  - *Warning:* If auto-sync is on, it will immediately revert to what is in Git.
  - **True Fix:** `git revert <commit-id>` in the repo.

---

### Question 612: What is auto-sync vs manual sync in GitOps?

**Answer:**
- **Manual:** Human verification required. Good for Prod.
- **Auto:** Continuous Deployment. Good for Dev/Stage.
  - Can configure "Auto-Sync with Manual Prune" for safety.

---

### Question 613: How do you manage secrets in a GitOps pipeline?

**Answer:**
**Sealed Secrets, SOPS, External Secrets.**
- Encrypted file in Git -> Controller decrypts in cluster.
- Reference ID in Git -> Controller fetches from Vault.

---

### Question 614: What is Kustomize and how is it used in GitOps?

**Answer:**
Native configuration management.
- **Base:** Common YAMLs.
- **Overlays:** Patches for specific environments.
- ArgoCD natively supports `kustomization.yaml` to build the final manifest.

---

### Question 615: How does Helm differ from Kustomize in GitOps?

**Answer:**
- **Helm:** Templating (Text replacement). Great for packaging/distribution.
- **Kustomize:** Patching (Structured editing). Great for configuration variants.
- ArgoCD supports both.

---

### Question 616: How do you set up CI triggers with GitOps?

**Answer:**
- **Push:** CI builds docker image -> Pushes to Registry -> Updates Git Repo (`image: v1` -> `image: v2`).
- **Pull:** ArgoCD detects git change -> Deploys.

---

### Question 617: How do you visualize GitOps deployments?

**Answer:**
**ArgoCD UI.**
- Visualizes tree of resources (App -> Service -> Pod).
- Colors indicate Health (Green) and Sync Status (Synced/OutOfSync).

---

### Question 618: What’s the risk of Git drift and how to detect it?

**Answer:**
**Drift:** Someone changed cluster manually (`kubectl edit`).
- Argo/Flux detects this plain as day: "OutOfSync".
- Can configure "SelfHeal" to immediately overwrite manual changes.

---

### Question 619: How do you manage multiple tenants using GitOps?

**Answer:**
**App-of-Apps Pattern.**
- `Root App` watches a folder containing `Team-A-App`, `Team-B-App` manifests.
- Enforce strict RBAC on who can commit to which folder in Git.

---

### Question 620: How does GitOps promote compliance and auditability?

**Answer:**
Git commit history is the Audit Trail.
- "Who changed the firewall rule?" -> `git blame`.
- "When was it changed?" -> Timestamp.
- "Why?" -> Commit message + PR description.

---

## 🔹 Admission Controllers & Policy Enforcement (Questions 621-630)

### Question 621: What is an admission controller in Kubernetes?

**Answer:**
Code that sits in the API path.
- **Mutating:** Defaults values.
- **Validating:** Rejects bad requests.
- Compiled into kube-apiserver binary (like `NodeRestriction`) or Webhook-based (OPA).

---

### Question 622: What is the difference between validating and mutating webhooks?

**Answer:**
- **Mutating:** "User sent pod. I will add a sidecar to it."
- **Validating:** "User sent pod. Does it have a sidecar? No? Reject."

---

### Question 623: What is OPA Gatekeeper?

**Answer:**
The standard policy engine.
- Uses **ConstraintTemplates** (logic in Rego) and **Constraints** (parameters).
- Example: "All namespaces must have a label `cost-center`".

---

### Question 624: How do you write policies with Rego in OPA?

**Answer:**
Rego is a query language.
```rego
violation[{"msg": msg}] {
  input.request.kind.kind == "Pod"
  not input.request.object.metadata.labels.owner
  msg := "Owner label missing"
}
```

---

### Question 625: What is Kyverno?

**Answer:**
Kubernetes-native policy engine.
- No Rego!
- Policies are simple YAML resources.
- Easier learning curve than OPA.

---

### Question 626: How does Kyverno differ from OPA Gatekeeper?

**Answer:**
- **OPA:** General purpose (can be used for Terraform, Linux, etc). Rego is complex.
- **Kyverno:** K8s only. YAML simplicity. Built-in Mutation/Generation capabilities (OPA strictly validating mostly).

---

### Question 627: How do you enforce image registry policies?

**Answer:**
Policy:
- `input.request.object.spec.containers[_].image` must start with `mycompany.jfrog.io/`.
- Reject `docker.io` or `quay.io`.

---

### Question 628: How do you prevent privilege escalation in policies?

**Answer:**
Block `securityContext.allowPrivilegeEscalation: true`.
- Block `privileged: true`.
- Block `capabilities: ["SYS_ADMIN"]`.

---

### Question 629: How do you test policies before enforcing?

**Answer:**
**Dry Run / Audit Mode.**
- Configure Gatekeeper/Kyverno to `action: audit` (not deny).
- Check logs for "Violation" events.
- Once clean, switch to `enforce`.

---

### Question 630: What is a dry-run policy evaluation?

**Answer:**
Evaluating the policy against existing resources to see what *would* break.
- Gatekeeper provides an "Audit" controller that updates `Constraint` status with list of existing violations.

---

## 🔹 Resource Scheduling & Placement (Questions 631-640)

### Question 631: How does Kubernetes score nodes for pod scheduling?

**Answer:**
Two steps:
1.  **Filtering:** (Ram ok? Port conflict? Taint matched?)
2.  **Scoring:** (Balanced? Affinity? Image Locality?)
- Weighted sum of scores. Highest wins.

---

### Question 632: What is inter-pod affinity vs anti-affinity?

**Answer:**
- **Affinity:** "I like Pod B. Put me near him." (Communication).
- **Anti-Affinity:** "I hate Pod A. Put me away from him." (High Availability).

---

### Question 633: How does `preferredDuringSchedulingIgnoredDuringExecution` work?

**Answer:**
Soft constraint.
- Scheduler *tries* to fulfill it.
- If impossible, schedules anyway.
- Once running, if condition violated later (node label changes), pod keeps running.

---

### Question 634: What is nodeSelector vs node affinity?

**Answer:**
- **nodeSelector:** Simple equality (`disk: ssd`).
- **Affinity:** Expressive (`key: zone, operator: In, values: [us-east-1, us-east-2]`). Supports Soft/Hard rules.

---

### Question 635: What is a topologySpreadConstraint?

**Answer:**
"Make sure pods are spread evenly across Zones".
- `maxSkew: 1`.
- `topologyKey: topology.kubernetes.io/zone`.
- Better than anti-affinity for balancing.

---

### Question 636: How do you pin a pod to a specific NUMA node?

**Answer:**
**Topology Manager** (Kubelet feature).
- `topology-manager-policy: single-numa-node`.
- Ensures CPU and Memory allocated are on same local bus for performance.

---

### Question 637: What is custom scheduler in Kubernetes?

**Answer:**
You can write your own scheduler binary.
- `schedulerName: my-scheduler` in PodSpec.
- Default scheduler ignores these pods.
- Your scheduler binds them.
- Use Case: Batch jobs, AI workloads.

---

### Question 638: What’s the difference between default-scheduler and kube-scheduler?

**Answer:**
They are the same thing. `kube-scheduler` IS the default scheduler implementation.

---

### Question 639: How do you write scheduling extenders?

**Answer:**
Instead of replacing the whole scheduler, config it to call an HTTP webhook (Filter/Prioritize).
- `FilterVerb`, `PrioritizeVerb`.
- Allows plugging in custom logic without recompiling K8s.

---

### Question 640: How do taints and tolerations impact scheduler decisions?

**Answer:**
They act as a **Filter** step.
- If Node has Taint X, and Pod no Toleration X -> Node Rejected.

---

## 🔹 Controller Manager Deep Dive (Questions 641-650)

### Question 641: What are the responsibilities of kube-controller-manager?

**Answer:**
Monolithic binary running all standard controllers.
- Node, Job, Endpoints, ServiceAccount, Token, etc.
- Run loop: Watch -> Diff -> Act.

---

### Question 642: What is the replication controller vs replica set?

**Answer:**
- **ReplicationController:** (Old/Deprecated). Only equality selector (`env=prod`).
- **ReplicaSet:** (New). Set-based selector (`env in (prod, stage)`).

---

### Question 643: What is the horizontal pod autoscaler controller?

**Answer:**
Calculates desired replicas.
- Fetches metrics.
- Updates `scale` subresource of Deployment.

---

### Question 644: How does the endpoint controller work?

**Answer:**
Watches **Services** and **Pods**.
- If Service Selector matches Pod Labels + Pod is Ready:
- Add Pod IP to `Endpoints` object.

---

### Question 645: What is the garbage collector controller?

**Answer:**
Handles **OwnerReferences**.
- If I delete deployment, this controller sees the RS has `owner: deployment`.
- It deletes the RS.
- It sees Pods have `owner: RS`.
- It deletes Pods.

---

### Question 646: What does the namespace controller do?

**Answer:**
Handles deletion.
- User deletes Namespace.
- Controller updates status `Terminating`.
- Deletes all resources in that NS.
- Once empty, deletes the Namespace object.

---

### Question 647: How does service account controller operate?

**Answer:**
Ensures every Namespace has a `default` ServiceAccount.
- Also manages TokenSecrets for ServiceAccounts (legacy behavior).

---

### Question 648: What is the job controller responsible for?

**Answer:**
Watches Job objects.
- Creates Pods to satisfy `completions`.
- Tracks success/failure.
- Backoff handling.

---

### Question 649: How do controllers detect stale state?

**Answer:**
They assume state is essentially stale/eventual.
- `resyncPeriod`: Even if no events happened, Lister re-queues all objects every X hours to ensure nothing slipped through.

---

### Question 650: What happens when a controller crashes?

**Answer:**
It restarts (Pod/Systemd).
- Since it is stateless (state is in Etcd), it reads Etcd, rebuilds cache, and resumes where it left off.
- Temporary delay in reconciliation.
