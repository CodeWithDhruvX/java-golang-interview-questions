# API Gateway vs Load Balancer

## 1. The Confusion
Both sit in front of servers and route traffic. But they serve very different purposes.

*   **Load Balancer (LB)**: "I distribute traffic."
*   **API Gateway (GW)**: "I manage APIs."

## 2. Load Balancer (The Traffic Cop)
*   **Core Function**: Distributes network traffic across multiple servers to ensure no single server is overwhelmed.
*   **Layer**: L4 (TCP/UDP) or L7 (HTTP).
*   **Features**: Health checks, SSL Termination.
*   **Analogy**: A receptionist who points you to the next available teller.

## 3. API Gateway (The Facade)
*   **Core Function**: A reverse proxy that acts as the single entry point for all Microservices. It routes requests to various services (Product Service, User Service).
*   **Layer**: L7 (Application).
*   **Features**:
    *   **Authentication/Authorization**: Validates JWTs before request hits backend.
    *   **Rate Limiting**: Throttles distinct users (LB usually throttles IPs).
    *   **Protocol Translation**: Converts HTTP REST -> gRPC or SOAP.
    *   **Request Aggregation**: Calls Service A and Service B, combines results, returns 1 JSON to client.
*   **Analogy**: A Project Manager who takes your complex request, talks to 5 different departments, and gives you one consolidated report.

## 4. Where they fit
Usually, you use **BOTH**.

```
Client -> [Load Balancer] -> [API Gateway 1, 2] -> [Internal Load Balancers] -> [Microservices]
```
1.  **Public LB**: Distributes traffic to multiple instances of API Gateway.
2.  **API Gateway**: Handles Auth, Routing, etc.
3.  **Service Discovery**: API Gateway asks Registry (e.g., Eureka/Consul) "Where is Product Service?" and forwards traffic.

## 5. Comparison Table

| Feature | Load Balancer | API Gateway |
| :--- | :--- | :--- |
| **Primary Goal** | High Availability, Scale | API Management, Security |
| **Routing Logic** | Round Robin, Least Conn | Path (/api/v1/users), Headers |
| **Auth** | Minimal (SSL) | Deep (OAuth2, JWT, Scopes) |
| **Payload** | Pass-through (mostly) | Can modify (Transformation) |
| **Examples** | AWS ALB, NGINX | Kong, AWS API Gateway, Zuul |

## 6. Interview Questions
1.  **Can NGINX be an API Gateway?**
    *   *Ans*: Yes. While NGINX is a Load Balancer, with modules (Lua) or NGINX Plus, it can do Auth, Rate Limiting, and Transformation, effectively acting as a Gateway.
2.  **Pattern: Backends for Frontends (BFF)**
    *   *Ans*: Instead of one giant API Gateway, create separate Gateways for separate clients.
        *   `Mobile-Gateway`: Returns small JSON (save data).
        *   `Web-Gateway`: Returns rich JSON.
