# 🟣 **266–280: Advanced Kubernetes & Deployment**

### 266. How does Kubernetes manage container networking?
"In Kubernetes, every Pod gets its own unique, routable IP address within the cluster network. This completely eliminates the need for managing dynamic port mapping (like Docker `-p 8080:8080`), because Pod A can talk to Pod B directly via their IPs.

When a pod dies and is recreated, its IP changes. To solve this, Kubernetes uses a **Service** object. A Service provides a stable, static virtual IP (or DNS name) that sits in front of the dynamically changing Pod IPs. 

When a microservice needs to call the `Payment` service, it literally sends an HTTP request to `http://payment-service:8080`. The Kubernetes internal DNS (CoreDNS) resolves this name to the Service's virtual IP, and `kube-proxy` then load balances the request to a healthy backing Pod."

#### Indepth
The networking is physically implemented by a Container Network Interface (CNI) plugin (like Flannel, Calico, or Cilium). These plugins handle the underlying complexity of assigning IPs and configuring routing tables across physical EC2/VM nodes so that a container on Server 1 can ping a container on Server 2 seamlessly.

---

### 267. What are Readiness and Liveness Probes?
"When Kubernetes starts my Spring Boot or Go microservice container, the JVM might take 5-10 seconds to load all beans before it can accept API traffic.

**Readiness Probe:** K8s sends an HTTP request to an endpoint I define (e.g., `/actuator/health`). If the probe fails, K8s *refuses* to send live user traffic to that Pod. Once it succeeds, the Pod is added to the load balancer pool. This prevents routing traffic to a container that is still booting up.

**Liveness Probe:** This checks if the running application has become deadlocked or permanently crashed. K8s pings the endpoint periodically. If it fails 3 times in a row, K8s forcibly kills the Pod and automatically restarts a fresh one to self-heal the system."

#### Indepth
You must design probes carefully. A Readiness probe usually checks downstream dependencies (e.g., "Can I currently connect to the Database? No? Then remove me from the load balancer"). A Liveness probe should *never* check downstream dependencies. If the master Database goes offline for 5 seconds, you do not want Kubernetes to aggressively kill and restart all 500 of your application pods simultaneously.

---

### 268. What is a StatefulSet vs a Deployment?
"A **Deployment** manages stateless applications (like a standard REST API microservice). If I request 3 replicas, K8s spins them up in random order, gives them random hash names (like `payment-5d6f7...`), and mounts identical, ephemeral storage. They are completely interchangeable cattle.

A **StatefulSet** manages stateful applications (like deploying a Cassandra or Kafka cluster inside K8s). It provides guarantees about the ordering and uniqueness of these Pods.
1. Pods get sticky, sequential names (`kafka-0`, `kafka-1`, `kafka-2`).
2. They are started strictly in order (0 must be healthy before 1 starts).
3. Most importantly, they get **Persistent Network Identity and Storage**. If `kafka-1` crashes, when K8s restarts it, it retains the exact same name `kafka-1` and automatically reconnects to the exact same Persistent Volume disk it used before."

#### Indepth
Deploying databases inside Kubernetes is historically controversial due to the complexity of distributed storage (CSI drivers) and dealing with node failures corrupting persistent states. Many enterprises still prefer running stateless microservices in K8s, but use managed cloud services (like AWS RDS or DynamoDB) for the actual databases.

---

### 269. How do you implement Zero-Downtime Deployments in K8s?
"I use the **Rolling Update** strategy, which is the default for K8s Deployments.

When I update the Docker image version from v1 to v2 in the deployment manifest, K8s does not kill all the v1 pods at once. 

Instead, it spins up a new v2 pod. Once the v2 pod passes its *Readiness Probe* (proving it can handle traffic), K8s gracefully terminates one v1 pod. It repeats this process one by one until all v1 pods are replaced by v2 pods. At no point is the application unavailable to the user."

#### Indepth
Graceful termination is critical here. When K8s sends a `SIGTERM` signal to a v1 pod, the application (e.g., Spring Boot) must stop accepting *new* HTTP requests immediately, but wait (usually up to 30 seconds) to finish processing any *in-flight* transactions before physically shutting down the JVM process.

---

### 270. What is an Ingress Controller?
"While a K8s 'Service' load balances traffic *inside* the cluster, an **Ingress** manages network traffic coming from the *outside* internet into the cluster.

Instead of provisioning an expensive AWS Application Load Balancer (ALB) for every single microservice (Payment, Order, User), I provision exactly ONE cloud Load Balancer that points to the K8s Ingress Controller (like Nginx-Ingress or Traefik).

The Ingress Controller reads routing rules I define:
- `myapp.com/api/payment` -> routes to internal Payment Service
- `myapp.com/api/order` -> routes to internal Order Service

It acts as a consolidated API Gateway, handling host/path-based routing, SSL termination, and rate limiting at the cluster edge."

#### Indepth
The Kubernetes Gateway API is the modern evolution of the classic Ingress API. It provides a more expressive, role-oriented standard for defining L4/L7 routing models within the cluster, allowing the Infrastructure team to define the Gateway, and product teams to dynamically attach HTTPRoutes to it.

---

### 271. What is Blue-Green Deployment vs Canary?
"Both are deployment strategies aimed at reducing risk.

**Blue-Green Deployment:** You run two completely identical production environments (Blue and Green). Blue is live, serving 100% of user traffic. You deploy v2 to Green and run heavy QA testing. Once verified, you flip the load balancer router to instantly send 100% of traffic to Green. If something breaks, you flip the router instantly back to Blue. It is safe but requires doubling your infrastructure bill.

**Canary Deployment:** You deploy v2 to a tiny subset of production servers and route only 5% of live user traffic to it. You monitor the success rates and error logs of the v2 'canary' closely using Datadog/Prometheus. If errors spike, the deployment automatically rolls back. If successful, you gradually dial the traffic up (20%, 50%, 100%). It mitigates blast radius significantly."

#### Indepth
Canary deployments in K8s are often automated using tools like **Argo Rollouts** or **Flagger**. These tools integrate directly with your monitoring system (Prometheus), checking HTTP 500 error rates every minute, and mathematically deciding whether to promote the canary to 10% traffic or instantly abort the release without human intervention.
