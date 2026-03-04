# High-Level Design (HLD): API Design and Integration

Enterprise software relies heavily on how disparate systems communicate. API design is a paramount topic for service-based company HLD rounds.

## 1. What are the key principles of RESTful APIs?
**Answer:**
REST (Representational State Transfer) is an architectural style for distributed hypermedia systems.
*   **Client-Server Separation:** The client handles UI/UX, the server handles data/logic. They evolve independently.
*   **Statelessness:** Every request from client to server must contain all information needed to understand the request. The server holds no session state about the client.
*   **Standardized Interfaces (Uniform Interface):** Use standard HTTP methods consistently:
    *   `GET` - Retrieve a resource.
    *   `POST` - Create a new resource.
    *   `PUT` - Completely update/replace a resource.
    *   `PATCH` - Partially update a resource.
    *   `DELETE` - Delete a resource.
*   **Resource-Based URLs:** URLs should identify resources (nouns), not actions (verbs). E.g., `/users/123/orders` NOT `/getOrdersForUser?id=123`.
*   **Cacheability:** Responses must define themselves as cacheable or not to improve network efficiency.

## 2. REST vs. SOAP: When would you use which?
**Answer:**
*   **SOAP (Simple Object Access Protocol):** An older XML-based protocol.
    *   *Features:* Strict contracts (WSDL), built-in standards for security (WS-Security) and transactional reliability (WS-AtomicTransaction). Message payload is always heavy XML.
    *   *When to use:* Legacy enterprise systems (banking, telecom) where formal contracts, distributed transactions, and highly structured security protocols are mandatory.
*   **REST:** An architectural style (not a strict protocol, though usually over HTTP/JSON).
    *   *Features:* Lightweight, faster parsing (JSON), highly scalable, uses standard HTTP infrastructure.
    *   *When to use:* Modern web and mobile applications, microservices, public-facing APIs where bandwidth and simplicity matter.

## 3. How do you version an API? Why is it necessary?
**Answer:**
API Versioning is required because once a service is published to clients, breaking changes (changing a property name, removing a field) will cause client applications to crash.
**Versioning Strategies:**
1.  **URI Versioning (Most Common):** `https://api.example.com/v1/users`
    *   *Pros:* Very visible, cache-friendly. *Cons:* Pollutes the URL space.
2.  **Query Parameter Versioning:** `https://api.example.com/users?version=1`
    *   *Pros:* Easy to implement.
3.  **Custom Header Versioning:** Passing a custom header like `X-API-Version: 1`
    *   *Pros:* Keeps URLs clean. *Cons:* Harder to test simply via browser.
4.  **Content Negotiation (Accept Header):** `Accept: application/vnd.example.v1+json`
    *   *Pros:* The most REST-pure approach.

## 4. How do you implement secure APIs?
**Answer:**
*   **HTTPS/TLS:** Data must always be encrypted in transit. Never use plain HTTP.
*   **Authentication & Authorization:**
    *   Use OAuth 2.0 or OpenID Connect for delegating access.
    *   Use JWT (JSON Web Tokens) to transmit stateless identity and roles.
*   **Input Validation:** Validate everything on the server side to prevent SQL Injection (use ORMs/Prepared Statements) and XSS (sanitize HTML).
*   **Rate Limiting:** Protect APIs from DoS attacks and brute force attempts by restricting the number of calls per IP/User.
*   **CORS (Cross-Origin Resource Sharing):** Define strict CORS policies to allow only specific trusted front-end domains from calling the API.

## 5. What is the API Gateway Pattern and why do enterprises use it?
**Answer:**
Instead of external clients directly calling hundreds of different internal backend microservices, an API Gateway provides a single, unified entry point for all clients.
*   **Benefits for Enterprises:**
    *   **Security & Centralized Auth:** Authentication happens once at the Gateway, offloading this from individual backend services.
    *   **Protocol Translation:** Gateway can accept REST from a mobile client and translate it into strict SOAP/XML for a legacy backend mainframe.
    *   **Rate Limiting & Throttling/Billing:** Enforcing SLA limits per enterprise client.
    *   **Aggregation:** Making multiple internal calls and bundling the results into one single JSON response for the client.
