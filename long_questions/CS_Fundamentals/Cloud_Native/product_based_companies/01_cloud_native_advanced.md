# ☁️ Cloud Native — Advanced Interview Questions (Product-Based Companies)

This document covers advanced cloud native/infrastructure concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto). Targeted at 3–10 years of experience rounds.

---

### Q1: How do Containers actually work under the hood in Linux? Explain namespaces and cgroups.

**Answer:**
"Containers" don't actually exist as a first-class object in the Linux kernel. They are an illusion created by combining two core kernel features: **Namespaces** and **Cgroups**, along with a **Union Filesystem**.

**1. Namespaces (Isolation):**
Namespaces restrict what a process can **see**.
- **PID Namespace:** Isolates the process ID tree. A process might be PID 1 inside the container, but PID 45982 on the host.
- **NET Namespace:** Isolates network interfaces. The container has its own `eth0`, distinct routing tables, and ports.
- **MNT Namespace:** Isolates mount points (filesystems).
- **UTS:** Isolates hostname.
- **USER:** Maps root in the container to a non-root user on the host.

**2. Cgroups / Control Groups (Resource Limiting):**
Cgroups restrict what a process can **use**.
- Limits memory, CPU, disk I/O, and network bandwidth.
- Prevents one container from consuming 100% of the host CPU (Noisy Neighbor problem).

**3. Union Filesystem (OverlayFS):**
- Layers filesystems transparently. Forms the basis of Docker images (Image layers are read-only; the top container layer is read-write).

---

### Q2: What is a Service Mesh? How does Istio use the Sidecar pattern?

**Answer:**
A **Service Mesh** is a dedicated infrastructure layer for managing service-to-service communication. It provides observability, traffic management, and security, abstracting these concerns away from the application code.

**The Sidecar Pattern:**
Instead of writing retry, circuit breaker, or mTLS logic inside your Java/Go application via libraries, a proxy container (e.g., Envoy) is deployed alongside your application container inside the **same Kubernetes Pod**. This is the "sidecar."

**How it works (Data Plane vs Control Plane):**
- **Data Plane (Envoy proxies):** Every request leaving or entering the application container passes through the Envoy proxy. The proxies handle routing, mTLS encryption, and metrics collection.
- **Control Plane (Istio / istiod):** Manages and configures the proxies. Pushes routing rules and issues TLS certificates to the sidecars.

**Benefits:**
- **Decoupled:** Developers write pure business logic. DevOps/Infra engineers configure traffic rules centrally.
- **Language Agnostic:** Works exactly the same for a Java pod, a Node.js pod, and a Rust pod.
- **Features:** A/B testing, Canary deployments, Distributed Tracing header injection, mTLS by default.

---

### Q3: Explain Kubernetes Operators. How do they differ from basic Deployments?

**Answer:**
A **Kubernetes Operator** is a method of packaging, deploying, and managing a **stateful application** by extending the Kubernetes API using Custom Resource Definitions (CRDs) and custom Controllers.

**The limitation of standard resources:**
Standard controllers (Deployments, StatefulSets) understand generic pods. They know how to restart a pod if it dies. But they know **nothing about the application domain**.
- If a primary database node dies, Kubernetes will restart the pod, but it doesn't know how to run the database-specific command to promote a replica to primary, update configuration, and elect a new leader.

**The Operator Pattern:**
- An Operator encodes the **human operational knowledge** (that an SRE would write in a runbook) into code.
- It constantly watches the state of its Custom Resources (e.g., a `PostgresCluster` resource).
- **Example:** You apply a YAML manifest that says `kind: PostgresCluster, replicas: 3`. 
- The Postgres Operator detects this, spins up a StatefulSet, initializes the primary, initializes replicas, wires up replication, sets up connection pooling, and configures backup cronjobs. It handles failovers natively.

---

### Q4: What is Serverless vs FaaS? What are the challenges with cold starts?

**Answer:**
**Serverless** is an operational model where infrastructure management is completely abstracted away, and billing is scaled to zero (pay-per-request).
**FaaS (Function-as-a-Service)** is a specific subset of serverless (e.g., AWS Lambda) where you deploy discrete functions instead of full applications.

**The Cold Start Problem:**
When a Lambda function hasn't been invoked recently, the cloud provider spins down the micro-VM/container holding it to save resources. When a new request arrives, a **Cold Start** occurs:
1. Provider provisions a new micro-VM (Firecracker).
2. Downloads your code/container image.
3. Bootstraps the runtime (JVM, Node environment).
4. Executes your initialization code.
5. Finally processes the event.

**Latency impact:** Can take hundreds of milliseconds to several seconds (especially for Java/.NET).

**Mitigation strategies:**
1. **Provisioned Concurrency:** Pay a baseline fee to keep N instances permanently warm.
2. **Runtime selection:** Go, Rust, and Node.js have significantly faster cold starts than Java.
3. **GraalVM/Ahead-of-Time compilation:** Compiling Java into a native binary drastically reduces startup time.
4. **Reduce package size:** Avoid heavy frameworks (e.g., avoid full Spring Boot in Lambda). Include fewer dependencies.

---

### Q5: What is eBPF? Why is it revolutionizing cloud native networking and observability?

**Answer:**
**eBPF (Extended Berkeley Packet Filter)** is a revolutionary technology that allows running sandboxed programs **within the OS kernel** without changing kernel source code or loading kernel modules.

**Why it matters:**
Traditionally, to inspect packets or trace system calls, tools had to switch context between user-space and kernel-space (which is slow), or require risky custom kernel modules.

**How eBPF changes Cloud Native:**
- eBPF programs are securely compiled and attached to kernel hooks (network events, syscalls, tracepoints).
- **Networking (Cilium):** Instead of relying on complex `iptables` rules (which degrade exponentially as a Kubernetes cluster grows), eBPF bypasses iptables entirely, routing packets directly at the socket level. Massive performance boost.
- **Observability (Hubble / Pixie):** Allows capturing granular network metrics, CPU profiling, and even HTTP/encrypted traffic directly from the kernel without instrumenting application code or injecting sidecars.
- **Security:** Granular process and network filtering directly at the kernel boundary.

eBPF is essentially to the kernel what JavaScript is to the web browser — a safe way to inject logic without rebuilding the core engine.

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
