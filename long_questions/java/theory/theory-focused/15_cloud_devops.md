# Cloud, DevOps & Deployment - Interview Answers

> 🎯 **Focus:** Product companies care deeply about *how* you ship and run code, not just how you write it.

### 1. How do you deploy a Spring Boot application to Kubernetes (K8s)?
"It’s a 3-step process.
1. **Containerize**: I build a Docker image of the JAR and push it to a registry (like ECR or Docker Hub).
2. **Define Resources**: I write a `deployment.yaml` to define *how many* replicas I want (e.g., 3 pods) and a `service.yaml` to expose them (LoadBalancer).
3. **Apply**: I run `kubectl apply -f deployment.yaml`. K8s pulls the image and schedulers the pods across the nodes."

---

### 2. What is Blue-Green Deployment?
"It’s a strategy to release software with zero downtime and instant rollback.
We have two identical environments: **Blue** (Live) and **Green** (Idle).
1. We deploy the new version to Green.
2. We test Green thoroughly.
3. We switch the Load Balancer to point to Green.
Now Green is Live. If anything goes wrong, we just flip the switch back to Blue instantly. No panic."

---

### 3. How do you handle Secrets (Passwords/API Keys) in Production?
"We never, ever hardcode them in `application.yml` or commit them to Git.
In Kubernetes, we use **K8s Secrets**. The secret is stored encoded in etcd and mounted into the Pod as an environment variable.
In AWS, we might use **AWS Secrets Manager**. Ideally, the app fetches the secret at runtime so it's not even visible in the environment variables."

---

### 4. CI/CD Pipeline flow?
"I typically use Jenkins or GitHub Actions.
The flow is:
1. **Commit**: Developer pushes code to Git.
2. **Build**: CI server detects the change, runs `mvn clean install`.
3. **Test**: It runs Unit Tests (`surefire`) and Integration Tests. If any fail, the pipeline stops.
4. **Scan**: SonarQube scans for vulnerabilities.
5. **Package**: It builds the Docker image.
6. **Deploy**: It pushes the image to Artifactory and triggers a deployment to the Dev/Staging cluster."

---

### 5. What is a "Sidecar" container?
"It’s a helper container that runs alongside the main application container in the *same* Pod.
A classic example is a **Logging Agent** (like Fluentd). The main app writes logs to a file, and the sidecar reads that file and pushes it to Splunk/ELK.
Another example is **Istio Envoy Proxy** for handling service mesh networking. They share the same lifecycle and network."

---

### 6. Horizontal Pod Autoscaling (HPA)?
"HPA automatically scales the number of pods up or down based on CPU or Memory usage.
If I set a target of 70% CPU, and a traffic spike hits causing CPU to go to 90%, K8s spins up new pods to handle the load. When traffic drops, it kills the extra pods to save money.
It’s the magic of the cloud."

---

### 7. What is Serverless (AWS Lambda)?
"It’s 'Function as a Service'. You focus purely on the code (the function).
You upload a Java function, and AWS handles the servers, OS, and scaling.
It’s event-driven. A file lands in S3 -> triggers Lambda -> processes file.
You pay only for the milliseconds the code runs. It's great for sporadic tasks but can have 'Cold Start' latency issues."

---

### 8. Docker vs Virtual Machine (VM)?
"A **VM** virtualizes the *Hardware*. It runs a full heavy OS (Guest OS) on top of the Host OS. It takes minutes to boot.
**Docker** virtualizes the *OS*. Containers share the Host OS kernel but have their own filesystem (bins/libs). They are lightweight (MBs, not GBs) and start in milliseconds.
That's why we can pack dozens of containers on a single machine."

---

### 9. How do you monitor a production Spring Boot app?
"We use the **ELK Stack** (Elasticsearch, Logstash, Kibana) for logs. All app logs are centralized there for searching.
We use **Prometheus & Grafana** for metrics. The app exposes `/actuator/prometheus`, Prometheus scrapes it, and Grafana visualizes dashboards (CPU, Heap, Request Latency).
We use **Zipkin/Jaeger** for distributed tracing to see how a request flows through microservices."

---

### 10. Breaking Monolith to Microservices - Strategy?
"I use the **Strangler Fig Pattern**.
I don't rewrite the whole thing at once. I identify one domain (e.g., 'Shipping').
1. I build a new 'Shipping Microservice'.
2. I point the API Gateway to route `/shipping` requests to the new service, and everything else to the old Monolith.
3. I repeat this module by module until the Monolith is empty (strangled)."
