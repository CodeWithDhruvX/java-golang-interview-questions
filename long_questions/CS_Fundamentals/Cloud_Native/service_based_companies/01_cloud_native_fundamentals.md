# ☁️ Cloud Native Fundamentals — Interview Questions (Service-Based Companies)

This document covers cloud native concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL. Targeted at 1–5 years of experience rounds.

---

### Q1: What does "Cloud Native" mean? How is it different from "Cloud Hosted"?

**Answer:**
**Cloud Native** refers to an approach for building and running applications that fully exploit the advantages of the cloud computing delivery model.

- **Cloud Hosted (Lift and Shift):** Taking an existing monolithic application and deploying it on an AWS EC2 instance. It runs in the cloud, but it is not optimized for it. If traffic spikes, you must manually spin up a massive new VM.
- **Cloud Native:** The application is architected *for* the cloud. It is typically built as microservices, packaged in containers, dynamically orchestrated (Kubernetes), and utilizes continuous delivery (CI/CD). It scales horizontally automatically and heals itself when nodes fail.

**Pillars of Cloud Native:**
1. Microservices architecture
2. Containers (Docker)
3. Orchestration (Kubernetes)
4. DevOps & CI/CD

---

### Q2: What are the Twelve-Factor App principles?

**Answer:**
The **12-Factor App** is a methodology for building modern, scalable, and maintainable SaaS applications, highly relevant to Cloud Native.

Some of the most important factors for interviews:

- **I. Codebase:** One codebase tracked in revision control (Git), many deploys (dev, staging, prod).
- **III. Config:** Store configuration in the **environment** (env vars), NOT in the code.
- **IV. Backing services:** Treat attached resources (databases, message queues) as interchangeable, network-attached services.
- **VI. Processes:** Execute the app as one or more **stateless** processes. Any data that needs to persist must be stored in a stateful backing service (database).
- **IX. Disposability:** Maximize robustness with fast startup and graceful shutdown. Don't assume a container will live forever.
- **XI. Logs:** Treat logs as event streams. Write to `stdout`/`stderr` and let the execution environment (EFK/ELK stack) route them.

---

### Q3: What is a Container? How is it different from a VM?

**Answer:**
A **Container** is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another.

| Feature | Virtual Machine (VM) | Container |
|---|---|---|
| Virtualizes | The entire hardware | The OS (user space) |
| Hypervisor/Kernel | Runs on a Hypervisor, has its own OS kernel | Runs on a container engine (Docker), shares the Host OS kernel |
| Boot Time | Minutes | Seconds (or milliseconds) |
| Size | Gigabytes | Megabytes |
| Isolation | Strong hardware-level isolation | Process-level isolation (Namespaces & Cgroups) |

**Use case:** Use containers for stateless APIs and microservices. Use VMs for workloads requiring strict security isolation or specific kernel versions.

---

### Q4: What is Kubernetes (K8s) and why do we need it?

**Answer:**
**Kubernetes** is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications.

**Why Docker isn't enough:**
Docker can run a container on a single laptop. But in production, you might have 100 containers spreading across 10 servers.
- What happens if a server dies?
- How do containers talk to each other securely?
- How do you scale from 3 instances to 20 when traffic spikes?

**What Kubernetes provides:**
- **Auto-scaling:** Spins up more pods based on CPU/RAM metrics.
- **Self-healing:** Restarts failed containers, replaces and reschedules containers when nodes die.
- **Service discovery & load balancing:** Gives containers distinct IPs and a single DNS name to balance load across them.
- **Storage orchestration:** Automatically mounts local or cloud storage (AWS EBS, NFS).
- **Rollouts & Rollbacks:** Update applications with zero downtime.

---

### Q5: Define Pod, Service, and Deployment in Kubernetes.

**Answer:**

**1. Pod:**
- The smallest deployable unit in Kubernetes.
- Encapsulates one or more containers (usually one) that share storage, network IP, and localhost.
- Pods are ephemeral — they can die and be replaced, getting a new IP address every time.

**2. Deployment:**
- A controller that provides declarative updates for Pods.
- You tell the Deployment "I want 3 replicas of the Nginx v1.14 pod."
- If a pod dies, the Deployment ensures a new one is created to maintain the desired state (3 replicas).
- Manages Rolling Updates (swapping old pods for new pods gradually).

**3. Service:**
- An abstract way to expose an application running on a set of Pods as a network service.
- Because Pod IPs change constantly, a Service provides a **stable IP address and DNS name** for a group of Pods.
- It acts as an internal load balancer, directing traffic to healthy running pods.

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
