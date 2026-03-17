# Priority Category 3: Automation & Orchestration

### 1. How do you implement workflow orchestration with Conductor in Go?
**Your Response:**
"Using Netflix Conductor with Go is about separating the 'What' from the 'How.' 
I write **Workers** in Go that poll for specific tasks from the Conductor server. Each worker is just a focused piece of code that executes one step (like 'Process Payment'). 
The complex logic—like 'If the payment fails, retry 3 times and then notify human support'—is defined in a JSON workflow in Conductor. This keeps my Go code clean, stateless, and incredibly easy to test because each worker only does one thing."

### 2. Design self-healing mechanisms in Go microservices.
**Your Response:**
"Self-healing works at two levels. 
First, **Proactive**: I implement `context` with timeouts and circuit breakers so the service doesn't hang. 
Second, **Reactive**: If a service enters a corrupt state (like a lost DB connection that doesn't recover), I mark it as 'Unhealthy' via a health check endpoint. This triggers Kubernetes to kill the pod and start a fresh one. 
In Go, I use a 'Graceful Shutdown' pattern: when the app receives a SIGTERM, it finishes current jobs, closes DB connections, and exits cleanly, which ensures the 'healing' process doesn't cause data loss."

### 3. How do you implement task execution and recovery workflows?
**Your Response:**
"I rely on **Idempotency** and **Compensation**. 
1. **Idempotency**: I design every task so it can be run multiple times safely. For example, 'Create Order' checks if the order ID already exists before doing anything. 
2. **Recovery**: If a critical task fails after all retries, I trigger a 'Compensating Transaction' in Conductor (the Saga pattern). If we can't 'Fix' the error, we gracefully 'Reverse' it (like refunding a partial payment) to keep the system state consistent."

### 4. Explain automation patterns for microservices management.
**Your Response:**
"I follow the **Infrastructure as Code** (IaC) and **GitOps** patterns. 
Everything—from my Go API's deployment manifest to its Prometheus alerting rules—is stored in Git.
I use **Helm** for templating our Kubernetes manifests so we can automate deployments across Dev, Staging, and Prod by just changing a `values.yaml` file. This eliminates manual errors and ensures our Go services are managed consistently no matter where they run."
