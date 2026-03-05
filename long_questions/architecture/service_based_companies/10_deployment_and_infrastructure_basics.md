# Deployment and Infrastructure Basics (Service-Based Companies)

## 1. What is the difference between Virtual Machines (VMs) and Containers (Docker)?

**Expected Answer:**
Both are used to isolate applications, but at different levels.

*   **Virtual Machines (VMs):**
    *   *Virtualize the Hardware.* A hypervisor creates multiple independent machines, each with its own full overarching Guest Operating System (e.g., Windows, Linux).
    *   Heavyweight, slow to boot (minutes), consumes significant RAM and CPU just to run the OS. Provides strong isolation.
*   **Containers (Docker):**
    *   *Virtualize the Operating System.* Containers share the host machine's OS kernel. They only package the application code, runtime, system tools, and libraries.
    *   Lightweight, extremely fast to boot (milliseconds to seconds), highly portable ("works on my machine, works in production").

## 2. Explain the stages of a standard CI/CD pipeline.

**Expected Answer:**
CI/CD (Continuous Integration / Continuous Deployment) automates the software delivery process.

1.  **Source/Code:** Developer pushes code to Git (GitHub/Bitbucket). A webhook triggers the CI server (Jenkins, GitLab CI).
2.  **Build (CI):** The code is compiled. Dependencies are downloaded (npm install, mvn clean install). Docker images are built.
3.  **Test (CI):** Automated Unit Tests and Integration Tests run. Code quality scans (SonarQube) and security vulnerability checks happen here. If tests fail, the pipeline stops.
4.  **Publish:** The built artifact (JAR file, Docker Image) is pushed to an Artifact Repository (Nexus, AWS ECR, Docker Hub).
5.  **Deploy (CD):** The artifact is deployed to an environment (Dev -> QA -> Staging -> Production). Often pushed to Kubernetes or AWS ECS.

## 3. Compare Blue-Green Deployment vs. Canary Deployment.

**Expected Answer:**
These are advanced deployment strategies to minimize downtime and risk during releases.

*   **Blue-Green Deployment:**
    *   Maintain two identical production environments (Blue and Green).
    *   Currently, Blue is live routing 100% of user traffic.
    *   Deploy the new version to Green. Run smoke tests on Green internally.
    *   When ready, flip the load balancer switch to route 100% of traffic to Green.
    *   *Pros:* Zero downtime, instant rollback (just flip the switch back to Blue).
    *   *Cons:* Expensive (requires double the infrastructure).
*   **Canary Deployment:**
    *   Gradually shift traffic to the new version in increments.
    *   Deploy new version to a small set of servers. Route 5% of user traffic to it (the "canary").
    *   Monitor logs, errors, and performance for that 5%.
    *   If successful, gradually increase to 10%, 25%, 50%, 100%. If errors jump, automate an immediate rollback to the old version.
    *   *Pros:* Minimizes risk to the overall user base. Cheaper than Blue/Green.

## 4. What is Infrastructure as Code (IaC)? Provide examples.

**Expected Answer:**
*   **What it is:** The practice of managing and provisioning computing infrastructure (servers, networks, databases) through machine-readable definition files (code), rather than manual hardware configuration or clicking through cloud console UI portals.
*   **Benefits:**
    *   *Version Control:* Infrastructure changes are tracked in Git, PRs can be reviewed.
    *   *Reproducibility:* You can spin up a perfect copy of Production in QA with one command.
    *   *Automation:* Eliminates manual human error in configuration "drift."
*   **Popular Tools:**
    *   **Terraform:** Cloud-agnostic, uses declarative HCL (HashiCorp Configuration Language).
    *   **AWS CloudFormation:** AWS-specific native tool using JSON/YAML.
    *   **Ansible / Chef / Puppet:** Often used alongside Terraform for Configuration Management (installing software on the servers once they are provisioned).

## 5. What role does a Load Balancer play in a web architecture?

**Expected Answer:**
A Load Balancer sits in front of a group of application servers and acts as a traffic cop.

*   **Traffic Distribution:** It distributes incoming HTTP requests evenly across available backend servers to prevent any single server from being overwhelmed.
*   **High Availability:** It constantly performs "Health Checks" on backend servers. If a server crashes or stops responding, the Load Balancer stops sending traffic to it until it recovers.
*   **SSL Termination:** It can handle decrypting HTTPS traffic, reducing the CPU load on the backend application servers.
*   *Types:* Layer 4 (Network LB, acts purely on TCP/IP) vs. Layer 7 (Application LB, acts on HTTP headers and URLs).

## 6. What is API Gateway, and why use it instead of exposing Microservices directly?

**Expected Answer:**
An API Gateway provides a single point of entry for all clients (Web, Mobile) into a microservices architecture. Exposing microservices directly to the outside world is an anti-pattern.

*   **Why use it:**
    *   **Routing:** Maps a single external API domain (e.g., `api.example.com/orders`) to the internal IP address of the Order microservice.
    *   **Cross-Cutting Concerns:** Centralizes logic that every service needs, such as Authentication/Authorization (validating JWTs), Rate Limiting, CORS headers, and Logging.
    *   **Protocol Translation:** The external client talks HTTP/REST to the gateway, but the gateway might talk gRPC to the internal services.
    *   **Aggregation:** A client asks for "User Profile", and the gateway gathers data from the User Service, Billing Service, and Notification Service, aggregating it into one single response payload.
