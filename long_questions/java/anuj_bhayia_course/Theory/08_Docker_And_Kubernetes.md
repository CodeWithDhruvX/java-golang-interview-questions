# Docker and Kubernetes - Interview Questions and Answers

## 1. What is Docker, and what problem does it solve in Microservices?
**Answer:**
**Docker** is an open-source platform that automates the deployment, scaling, and management of applications using containerization.

**Problem Solved in Microservices:**
Before Docker, developers built applications locally ("it works on my machine"), but deploying them to production often failed due to different operating systems, missing dependencies, or conflicting library versions. With 50 microservices, managing 50 different runtime environments is impossible.

Docker solves this by packaging an application and all its dependencies (Java Runtime Environment, system libraries, configuration files) into a single, standardized unit called a **Container**.
- **Consistency:** A Docker container runs exactly the same way on a developer's Mac, a QA testing server, and an AWS production server.
- **Isolation:** Each container is isolated from others. You can run a service requiring Java 8 and another requiring Java 17 on the exact same virtual machine without interference.
- **Efficiency:** Containers share the host OS kernel, making them much lighter, faster to boot, and less resource-intensive than traditional Virtual Machines (VMs).

## 2. Explain the process of Dockerizing a Spring Boot application step-by-step.
**Answer:**
**Step 1: Write a Dockerfile:**
A `Dockerfile` is a text document containing instructions on how to build a Docker image.
```dockerfile
# Use an official OpenJDK base image from Docker Hub
FROM eclipse-temurin:17-jre-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the packaged Spring Boot executable JAR file into the container
COPY target/my-spring-app.jar /app/my-spring-app.jar

# Expose the port the app runs on (informational)
EXPOSE 8080

# Specify the command to run when the container starts
ENTRYPOINT ["java", "-jar", "my-spring-app.jar"]
```

**Step 2: Build the Docker Image:**
Run the Docker CLI command in the directory containing the `Dockerfile` to build the image and tag it.
`docker build -t my-company/my-spring-app:1.0 .`

**Step 3: Run the Docker Container:**
Start a container from the newly created image, mapping the host port to the container port.
`docker run -p 8080:8080 -d --name my-app-container my-company/my-spring-app:1.0`

## 3. What is Docker Compose, and why is it useful?
**Answer:**
While the `docker run` command is fine for starting a single container, a typical Spring Boot microservice requires a database (like PostgreSQL), a cache (like Redis), and maybe a message broker (like Kafka). Starting all of these individually, linking them together, and setting up shared networks using long CLI commands is tedious and error-prone.

**Docker Compose:**
- It is a tool for defining and running multi-container Docker applications.
- **How it works:** You define all the services, their images, environment variables, ports, and network dependencies in a single, declarative YAML file (`docker-compose.yml`).
- **Execution:** You bring up the entire environment with a single command: `docker-compose up -d`. This is incredibly useful for local development, allowing developers to spin up the entire microservice ecosystem instantly.

## 4. What is Kubernetes (K8s) and what is Container Orchestration?
**Answer:**
While Docker spins up containers, what happens when a container crashes in production? Who restarts it? If traffic spikes, how do you spin up 10 more containers of the Order Service? How do you route traffic to them? Docker alone cannot handle this natively at an enterprise scale.

**Container Orchestration:** The automated management of lifecycle, scaling, networking, deployment, and health-checking of containers across a cluster of servers.

**Kubernetes:** It is the industry-standard, open-source container orchestration platform originally developed by Google.
- **Role:** Kubernetes automatically deploys your Docker containers across a cluster of machines. It monitors their health, restarts them if they fail (Self-Healing), scales them up or down based on CPU/Memory usage (Auto-scaling), and performs zero-downtime rolling updates.

## 5. Explain the core components of the Kubernetes Architecture (Master and Worker Nodes).
**Answer:**
A Kubernetes cluster consists of two main parts:

**1. The Control Plane (Master Node):** The brain of the cluster making global decisions.
- **kube-apiserver:** The front-end API; all communication (CLI tools like `kubectl` or users via dashboards) goes through it.
- **etcd:** A highly available key-value store holding all cluster data, configuration, and state.
- **kube-scheduler:** Watches for newly created Pods that have no assigned node and selects a worker node for them to run on based on resource requirements.
- **kube-controller-manager:** Runs controller processes (like Node Controller noticing when nodes go down, or ReplicaSet Controller ensuring exactly 3 copies of a Pod are running).

**2. Worker Nodes:** The machines (physical or virtual) that actually run the containerized applications.
- **kubelet:** An agent running on each worker node ensuring the containers described in the PodSpecs are actually running and healthy.
- **kube-proxy:** Manages network routing rules on each node, allowing communication to the Pods from inside or outside the cluster.
- **Container Runtime:** The software that runs the containers (e.g., Docker, containerd).

## 6. What are the key Kubernetes Objects (Pod, ReplicaSet, Deployment, Service)?
**Answer:**
You declare the desired state of your applications using YAML files referencing these objects. K8s works continuously to make reality match your desired state.

1. **Pod:** The smallest, most basic deployable object in Kubernetes. It represents a single instance of a running process. A Pod usually encapsulates a single Docker container (though it can hold multiple tightly coupled containers sharing network/storage).
2. **ReplicaSet:** Ensures that a specified number of Pod replicas are running at any given time for high availability.
3. **Deployment (Usually what developers use):** A higher-level abstraction that manages ReplicaSets. It provides declarative updates to Pods. If you want to update from version 1.0 to 1.1 of your Spring Boot app, you update the Deployment, and it automatically orchestrates a rolling update, spinning up v1.1 Pods and terminating v1.0 Pods without downtime.
4. **Service:** Pods are ephemeral; they die and get assigned new IP addresses. A Service defines a logical set of Pods and a strictly stable DNS name and IP address policy by which to access them (Load Balancing). Example: A `ClusterIP` Service allows frontend Pods to connect to database Pods reliably, regardless of which specific database Pod is alive at that exact millisecond.

## 7. How do you expose a Spring Boot microservice running in Kubernetes to the outside world?
**Answer:**
By default, Pods and standard Services (`ClusterIP`) are only accessible from within the internal Kubernetes cluster network. To expose a service to external internet users, you configure specific K8s objects:

1. **NodePort Service:** Exposes the service on a specific, static port (e.g., 30005) on the IP address of *every* Worker Node in the cluster. It's simple but generally not recommended for production due to port limitations and security.
2. **LoadBalancer Service:** The standard way in cloud environments (AWS, GCP, Azure). When you define a Service as type `LoadBalancer`, Kubernetes talks to the cloud provider to provision an external cloud load balancer (e.g., an AWS ALB). Traffic hits the cloud ALB and is routed into your cluster.
3. **Ingress:** The most powerful approach for routing HTTP/HTTPS traffic. An Ingress is an API object that manages external access (routing rules) to multiple internal services based on the URL path or hostname (e.g., `api.example.com/users` routes to User Service, `api.example.com/orders` routes to Order Service). It requires an Ingress Controller (like NGINX) running in the cluster to implement the rules.

## 8. What is the Kubernetes Dashboard, and how is it used?
**Answer:**
The **Kubernetes Dashboard** is a general-purpose, web-based user interface for Kubernetes clusters.

**Use Cases:**
- Instead of using only the `kubectl` command-line tool, developers and administrators can visually monitor the health of the cluster.
- View and manage cluster resources: You can view details about Deployments, Pods, Services, and ConfigMaps.
- Check Logs: You can directly read the output logs of specific containers running inside a Pod from the UI.
- Execute commands: You can open an interactive shell terminal into a running container directly from the browser for debugging.
## 9. What are Docker Volumes, and why are they necessary for databases?
**Answer:**
By default, all files created inside a Docker container are stored on a writable container layer. When the container is deleted, **all data is permanently lost**.

This is fine for a stateless Spring Boot API, but disastrous for a database container (like PostgreSQL or MySQL).

**Docker Volumes:**
- Volumes are the preferred mechanism for persisting data generated by and used by Docker containers.
- They are directories stored completely *outside* the container's virtual file system, physically residing on the host machine's hard drive (managed by Docker in `/var/lib/docker/volumes/`).
- **Benefits:**
  - Data outlives the container: If you `docker rm` your database container and start a new one pointing to the same volume, all your tables and records are still there.
  - Performance: Volumes do not increase the size of the containers using them, and the volume's contents exist outside the lifecycle of a given container.
  - Sharing: Volumes can be shared and mapped into multiple containers simultaneously.

## 10. Explain the concept of Volumes and PersistentVolumeClaims (PVCs) in Kubernetes.
**Answer:**
Similarly to Docker, the disk files inside a Kubernetes Pod are ephemeral. If a Node crashes and the K8s Control Plane spins up a replacement Pod on a different Node, it starts with a clean slate.

To run stateful applications (like databases or message brokers) in K8s, you need to detach storage from the ephemeral Pod lifecycle.

**1. PersistentVolume (PV):**
A piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned via Storage Classes. It is a resource in the cluster, just like a Node is a cluster resource. It physical abstracts the storage technology (e.g., an AWS EBS volume, a Google Cloud Persistent Disk, or an NFS drive).

**2. PersistentVolumeClaim (PVC):**
A request for storage by a developer/user. Just as a Pod consumes Node resources (CPU/Memory), a PVC consumes PV resources (Disk Size/Access Modes).
- The developer creates a YAML file defining a PVC: "I need 50GB of ReadWriteOnce storage."
- Kubernetes automatically looks for an available PV that matches these criteria (or dynamically spins one up) and **binds** the PVC to the PV.
- Finally, in the Deployment/Pod YAML, the developer mounts the PVC into the container's file system (e.g., mapping it to `/var/lib/postgresql/data`).
- Even if the Pod crashes and is rescheduled to an entirely different physical Worker Node, Kubernetes guarantees the EBS volume (the PV) detaches from the old node and re-attaches to the new node, ensuring zero data loss.
