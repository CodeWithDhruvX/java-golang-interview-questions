# 🟢 **76–90: Security**

### 76. How to secure microservices?
"Securing microservices requires a multi-layered approach ('Defense in Depth'). The massive attack surface demands both edge security and internal zero-trust.

At the Edge, I mandate an API Gateway to handle rate-limiting, SSL termination, and initial authentication (usually rejecting requests lacking a valid JWT). 

Internally, I never assume the network is 'safe'. I enforce mTLS (Mutual TLS) for service-to-service communication to prevent internal wiretapping. Finally, I restrict access at the infrastructure level using Kubernetes Network Policies, ensuring that the Order pod can talk to Payment, but the Invoice pod definitively cannot."

#### Indepth
Common weak points are exposed management interfaces (Spring Boot Actuator endpoints like `/env` or `/heapdump`). If these are accidentally exposed publicly without strict Spring Security constraints, it leads to instantaneous catastrophic credentials leaks.

---

### 77. What is OAuth2?
"OAuth2 is a massive industry-standard authorization framework. It allows an application to obtain limited access to user accounts on an HTTP service without exposing user credentials.

When you see 'Log in with Google', that's OAuth2. The client application asks the Authorization Server (Google) for a token. Google authentically verifies the user, asks for consent, and hands back an Access Token. The client uses this token to fetch your profile picture.

I implement it because it completely offloads the terrifying responsibility of password management—hashing, storing, and resetting—away from my microservices and onto a dedicated Identity Provider like Okta or Keycloak."

#### Indepth
OAuth2 defines four primary 'Grant Types': Authorization Code, Implicit, Client Credentials, and Resource Owner Password Credentials. The Authorization Code flow (with PKCE for mobile/SPAs) is the most secure and universally recommended approach today.

---

### 78. What is OpenID Connect?
"OpenID Connect (OIDC) is an identity layer built beautifully on top of the OAuth2 protocol. 

Where OAuth2 hands you an Access Token (an opaque string simply granting *access*), OIDC hands you an additional token: the **ID Token** (always a JWT).

The Access Token tells an API what you are *allowed* to do. The ID Token tells the client application *who* you actually are (containing claims like `email`, `name`, and `picture`). I use OIDC when my frontend application needs to display 'Welcome, Dhruv!' natively."

#### Indepth
Because an Access Token can theoretically just be a massive random string meaningless to everyone except the backend, the ID Token is mathematically structured (JSON Web Token) so the frontend React/Angular app can parse it locally without needing a continuous network connection to an identity server.

---

### 79. What is JWT?
"A JWT (JSON Web Token) is a standardized, extremely compact, and self-contained way to securely transmit information between parties as a JSON object.

A JWT has three parts: Header, Payload (where the user's ID and roles live), and Signature. 

When my Authorization Server generates a JWT, it digitally signs it using a private key (like RS256). When a user sends that JWT to my Order microservice, the Order service verifies the signature using the Authorization Server's public key. If math checks out, my service trusts the payload data completely and instantly, *without* needing to make a database or network call to verify the user."

#### Indepth
Because they are self-contained and stateless, JWTs are the bedrock of microservice scalability—they eliminate the central stateful "Session Database" bottleneck. However, the inability to instantly invalidate/revoke a compromised stateless JWT is its most heavily criticized flaw, usually mitigated by extremely short shelf lives (e.g., 5-minute expirations).

---

### 80. What is token refresh strategy?
"Because Access Tokens (JWTs) cannot be easily revoked, they are designed to expire extremely quickly—often in just 5 to 15 minutes. 

To prevent the user from having to violently log back in every 15 minutes, the Authorization Server provides a long-lived **Refresh Token** alongside the initial Access Token.

When the Access Token expires, the frontend API call fails with a 401 Unauthorized. The frontend seamlessly intercepts this, silently sends the Refresh Token to the auth server, receives a brand new Access Token, and retries the original API call. The user notices nothing."

#### Indepth
Refresh Tokens are long-lived (days or months) and are inherently stateful. Unlike JWTs, they are stored securely in a central database by the Auth server. If a user's laptop is stolen, the admin can simply "revoke" the Refresh Token in the database. In 14 minutes, the thief's Access Token dies, and they will be permanently locked out.

---

### 81. What is mTLS?
"mTLS stands for Mutual Transport Layer Security. Standard TLS (HTTPS) involves the client verifying the server's certificate to ensure it's not talking to an imposter. 

In mTLS, the verification goes both ways: the server *also* demands and verifies the client's cryptographic certificate.

In a microservice cluster, I use mTLS to guarantee that when the Payment service receives a request, it cryptographically proves that the request came exclusively from the Order service, not from a rogue container secretly injected into my network."

#### Indepth
Setting up mTLS manually involves generating massive amounts of X.509 certificates and managing complex Certificate Authorities, which is an operational nightmare. Consequently, developers rely on Service Meshes (like Istio or Linkerd) which transparently handle certificate generation, rotation, and enforcement across the entire cluster dynamically.

---

### 82. What is service-to-service authentication?
"Service-to-service authentication answers the question: 'How does Service A prove its identity to Service B?'

The network is flat. Anyone inside the VPC can hit any IP address. 

Beyond mTLS (which handles the network transport layer), I implement this via the OAuth2 **Client Credentials Flow**. The Order service hits the Auth server, identifies itself using a hardcoded `client_id` and `client_secret`, and receives an Access Token. It passes this token to the Payment service in the Authorization header. Payment verifies the token."

#### Indepth
Relying entirely on perimeter security (the 'castle-and-moat' approach) is obsolete. A single compromised frontend container gives the attacker unrestricted HTTP access to the entire backend. Service-to-service authentication strictly enforces the Zero Trust principle at the application logic layer.

---

### 83. What is RBAC?
"RBAC stands for Role-Based Access Control. Instead of assigning individual permissions to users (which is unscalable), permissions are grouped into Roles (e.g., 'ADMIN', 'VIEWER', 'CASHIER'). 

Users are assigned a role. When a user requests my microservice, their JWT usually contains a `roles: ["ADMIN"]` claim. 

In Spring Boot, I use the `@PreAuthorize("hasRole('ADMIN')")` annotation directly over my controller method. The framework magically intercepts the request, checks the JWT roles, and throws an HTTP 403 Forbidden if the user isn't an ADMIN, enforcing authorization effortlessly."

#### Indepth
While RBAC is great for generic access, it fails at granular data ownership. An 'AUTHOR' role lets me edit articles, but RBAC cannot easily restrict me to editing *only my own* articles. For that, systems eventually evolve to ABAC (Attribute-Based Access Control) which evaluates runtime contexts (User ID == Article Owner ID).

---

### 84. What is API key authentication?
"API Key authentication is a simple security scheme where the client attaches a pre-generated string (the API Key) to every request, usually as an HTTP header (`x-api-key: a1b2c3d4`).

I use it heavily for B2B (Business-to-Business) server-side integrations, or to identify programmatic clients rather than human users.

Unlike OAuth2, it has no complex flows or token expirations. The client just passes the string, and my API Gateway checks if the string exists in the database. However, it's inherently inflexible and highly dangerous if exposed in frontend JavaScript code."

#### Indepth
Because they rarely expire automatically, API Keys must never be used in Single Page Applications (SPAs) or mobile apps where hackers can reverse-engineer the code to steal them. Furthermore, the backend should securely hash API keys in the database exactly like passwords—never store them in plain-text.

---

### 85. What is zero-trust architecture?
"Zero-Trust Architecture is a philosophical IT security paradigm summarized by the motto: 'Never Trust, Always Verify'.

Traditionally, firewalls protected the edge. Once you VPN'd inside the network, you were 'trusted' and could talk to any server unimpeded. Zero-Trust destroys this assumption. It operates on the premise that the internal network is *already compromised*.

I implement this by requiring strong authentication, mTLS encryption, and strict RBAC authorization on every single internal microservice endpoint, regardless of whether the call originated externally or from another internal backend service."

#### Indepth
It shifts the defensive perimeter from the broad network boundaries directly down to the individual users, devices, and application workloads.

---

### 86. How to secure secrets?
"A 'secret' is anything critical: database passwords, API keys, Kafka connection strings, or JWT signing certificates. 

The cardinal rule is: I absolutely *never* hardcode secrets in source code, and I never commit them to Git repositories.

Instead, I heavily utilize external Secret Managers like HashiCorp Vault, AWS Secrets Manager, or Kubernetes Secrets. My Spring Boot application boots up blank, securely connects to the Vault using an IAM role, and pulls its database password purely into temporary memory at runtime."

#### Indepth
If a GitHub repo goes public containing `application.yml` passwords, scrapers compromise the database in seconds. Secret Management tools not only securely store data but also provide "Secret Rotation"—automatically changing the database password every 30 days without human intervention or application downtime.

---

### 87. What is CSRF?
"Cross-Site Request Forgery (CSRF) is an attack where a malicious website tricks a user's browser into executing an unwanted action on a *different* site where the user is currently authenticated.

If a user logs into their bank using a Cookie-based session, and then opens a malicious site in another tab, the malicious site can secretly send a `POST /transfer-funds` request entirely in the background. The browser, trying to be helpful, automatically attaches the user's authentic Bank Cookie to this malicious request. 

The bank honors the transaction because it possesses a valid session cookie."

#### Indepth
I mitigate CSRF by issuing unpredictable Anti-CSRF tokens to the legitimate frontend application, which requires the frontend to consciously append it to every POST. Alternately, by utilizing local-storage JWTs instead of Cookies, CSRF is mathematically impossible because the browser does not automatically append local-storage tokens to cross-site background requests.

---

### 88. What is CORS?
"Cross-Origin Resource Sharing (CORS) is a browser security mechanism that restricts web pages from making API requests to a different domain than the one that served the web page.

If my frontend runs on `myapp.com`, but my API runs on `api.myapp.com`, the browser strictly blocks the request by default. 

To fix this, I configure my API server (or API Gateway) to explicitly return specific CORS headers (like `Access-Control-Allow-Origin: https://myapp.com`). The browser reads this and permits the Javascript to execute the call safely."

#### Indepth
Before a `POST` or `PUT`, browsers transparently fire an `OPTIONS` request (a "Preflight" request) to ask the server for permission. Servers must be heavily optimized to handle CORS requests, as naive implementations effectively double the network traffic volume by continuously computing preflight authorizations.

---

### 89. How to implement rate limiting?
"Rate limiting protects an API by capping the number of requests a given client can make within a time window (e.g., 50 requests per minute).

I always implement this at the very edge of my infrastructure—the API Gateway.

Using an algorithm like the Token Bucket or Sliding Window, the Gateway tracks the client's IP address or JWT User ID in a high-speed distributed cache like Redis. If the client makes request #51 in a minute, the Gateway immediately drops the request and returns the HTTP status `429 Too Many Requests`."

#### Indepth
Rate limiting prevents noisy-neighbor problems in multi-tenant systems and stops elementary brute-force attacks (like guessing passwords 5,000 times a second). A well-designed system also returns headers like `X-RateLimit-Reset` to tell honest clients exactly how many seconds to pause before trying again.

---

### 90. How to prevent DDoS?
"A Distributed Denial of Service (DDoS) attack involves compromising thousands of computers worldwide and commanding them to simultaneously flood my servers with fake traffic, exhausting my bandwidth or CPU so real users cannot connect.

No application code can survive a massive 500-Gigabit-per-second volumetric DDoS attack. 

To prevent it, I place a robust CDN (Content Delivery Network) or Cloud Edge security layer—like Cloudflare or AWS Shield—in front of my infrastructure. These edge networks absorb the massive traffic spikes globally, mathematically scrub away the malicious packets, and only forward clean, legitimate HTTP requests down to my backend API Gateway."

#### Indepth
DDoS attacks operate at different OSI layers. Layer 3/4 attacks flood raw UDP or SYN packets. Layer 7 attacks are slower but infinitely more devastating: they simulate thousands of real HTTP clients asking for the most computationally expensive JSON report in the app, crippling the database servers incredibly quickly. Application-level throttling is required alongside edge scrubbing.
