# System Design (HLD) - API Gateway Design

## Problem Statement
Design an API Gateway for a sprawling microservices architecture (e.g., 50+ internal microservices). The gateway will be the single entry point for all client applications (Web, iOS, Android).

## 1. Requirements Clarification
### functional Requirements
*   **Routing:** Route incoming HTTP requests to the correct internal microservice based on the URL path.
*   **Authentication & Authorization:** Validate JWT tokens before letting requests pass through to internal networks.
*   **Rate Limiting / Throttling:** Protect internal services from DDoS.
*   **Protocol Translation:** (Optional) Convert REST HTTP from the outside to gRPC on the inside.

### Non-Functional Requirements
*   **High Performance:** Adds minimal latency (Sub-millisecond overhead).
*   **High Availability:** Deployed across multiple Availability Zones.
*   **Extensibility:** Easy to add custom plugins (e.g., Logging, Caching, WAF).

## 2. Why an API Gateway? (The Before & After)
*   **Without Gateway:** The iOS APP has to know the IP address of 50 different microservices. If the `UserSvc` changes IP, the iOS app breaks until an update is pushed to the App Store. Furthermore, every single microservice must implement JWT validation, Rate Limiting, and CORS headers. This is repeated code and a security nightmare.
*   **With Gateway:** The iOS APP only talks to `api.company.com/v1/*`. The Gateway handles Security/Auth centrally, then seamlessly proxies the request over the private subnet to the `UserSvc`.

## 3. High-Level Architecture

```text
       [ iOS App ]     [ Website ]
            \              /
             \            /
       (HTTPS over public internet)
              \          /
     +--------------------------------+
     |        Load Balancer           | (AWS ALB / Nginx)
     +--------------------------------+
                    |
     +--------------------------------+
     |         API Gateway Range      |
     |  (Kong, Nginx, AWS API Gateway)|
     |--------------------------------|
     | - Auth (JWT Verification)      |
     | - Rate Limiting (Redis backed) |
     | - Request Tracing (Zipkin/B3)  |
     | - Request / Response Logging   |
     +--------------------------------+
                    |
           (Internal Private VPC)
         /          |          \
   [Auth Svc]  [Product Svc] [Order Svc]
```

## 4. Gateway Components / Request Lifecycle

1. **Firewall / WAF:** Filters out SQL injections, malformed HTTP packets.
2. **Authentication Filter:**
   * Extracts the `Authorization: Bearer <JWT>` header.
   * Validates the cryptographic signature of the JWT using the Auth Service's Public Key.
   * Extracts the `User_ID` and injects it as an internal header (e.g., `X-User-ID: 123`). This means backend microservices *never* have to parse JWTs—they simply trust the `X-User-ID` header coming from the internal Gateway.
3. **Rate Limiting Filter:**
   * Checks Redis for `rate_limit:123` to ensure the user hasn't exceeded 100 req/min.
4. **Tracing Filter:**
   * Generates a unique `X-B3-TraceId` and passes it downstream.
5. **Reverse Proxy / Routing:**
   * Looks up the URI `/v1/products` in its routing table and forwards it to the private IP `10.0.1.55:8080`.

## 5. Caching at the Edge
*   The API Gateway can also cache responses! If 10,000 users ask for the homepage products `/v1/home-feed`, the Gateway caches the JSON response in Redis for 60 seconds.
*   This means the 9,999 subsequent requests never even reach the `Product Svc`, saving immense backend compute resources.

## 6. Follow-up Questions for Candidate
1.  How do you handle API versioning? (Using path `/v1/` vs `/v2/` or Accept Headers application/vnd.company.v1+json).
2.  What happens if the internal `Order Svc` dies? Does the Gateway hang forever waiting for a response? (No, implement Timeouts and Circuit Breakers in the Gateway to instantly return HTTP 503 instead of tying up threads).
3.  Kong vs NGINX vs Spring Cloud Gateway: Compare them.
