# High-Level Design (HLD): API Security and IAM

Service-based interviews, especially for FinTech and HealthCare clients, prioritize security architecture over absolute scalability.

## 1. OAuth 2.0 vs OpenID Connect (OIDC)
**Answer:**
*   **OAuth 2.0:** is an **Authorization** framework. It provides a way for a client application (Spotify) to access resources hosted on another server (Facebook photos) on behalf of a user, without the user giving their password to Spotify. It yields an `Access Token`.
*   **OpenID Connect (OIDC):** is an **Authentication** layer built *on top* of OAuth 2.0. It allows the client application to know *who* the user is. Instead of just an Access Token, OIDC also returns an `ID Token` (which is always a JWT) containing basic profile information (email, name, picture). "Sign in with Google" uses OIDC.

## 2. JWT (JSON Web Token) vs. Opaque Tokens
**Answer:**
A system needs to maintain sessions. How do we pass the session identity?
*   **Opaque Tokens (Reference Tokens):** A random string (e.g., `d7a8f9b...`) stored in a database or Redis.
    *   *How it works:* Client sends the token. The API Gateway receives it, goes to the Database/Redis to look up "who owns this token?".
    *   *Pros:* Immediate revocation. If an admin clicks "Logout All Devices," the record in Redis is deleted. The next request fails immediately.
    *   *Cons:* Not stateless. Every API request results in a network hop to the Redis/Database to validate the token, adding latency.
*   **JWT (Value/Stateless Tokens):** A Base64 encoded JSON object containing the user data, cryptographically signed by the Authorization server (using a secret key or RSA private key).
    *   *How it works:* The API Gateway receives the JWT. It looks at the signature and runs a quick mathematical verification using a public key. If the signature is valid, it implicitly trusts the payload inside the JWT. No DB lookup required!
    *   *Pros:* Completely stateless. Massively scalable.
    *   *Cons:* Very hard to revoke before expiration. If a token has an expiration (TTL) of 24 hours and a hacker steals it, there is no easy way to invalidate it instantly because no DB is checked. (Solution: Keep JWT TTL very short, e.g., 5 minutes, and use Refresh Tokens).

## 3. RBAC vs. ABAC for Authorization
**Answer:**
Once Identity is confirmed, how do we determine permissions?
*   **RBAC (Role-Based Access Control):** Broad strokes. You assign permissions to a "Role" (e.g., `Admin`, `Editor`, `Viewer`), and you assign users to those roles. (User X is an Editor. Editors can Delete posts. Therefore, User X can delete posts).
*   **ABAC (Attribute-Based Access Control):** Fine-grained. Access is granted based on attributes of the user, the resource, and the environment. 
    *   *Example:* "A user can edit a document IF the user's `department` equals the document's `owning_department`, AND the time is between `9AM-5PM`."
    *   *Use Case:* Complex enterprise or healthcare systems where "Doctor" role isn't enough; they can only see files of *their specific* patients.

## 4. How do you secure an internal Microservice to Microservice communication?
**Answer:**
It's an anti-pattern to assume the internal corporate network (VPC) is safe. This requires a **Zero-Trust Architecture**.
*   **Mutual TLS (mTLS):** Normal HTTPS ensures the client trusts the server. mTLS ensures the Server *also* requires the Client to present a valid cryptographic certificate to prove who it is.
*   **Service Mesh:** Rather than hardcoding certificates into every application's Java/Node code, use a Service Mesh (like Istio/Linkerd). A sidecar proxy (Envoy) sits next to the application. When Service A wants to talk to Service B, it just talks unencrypted HTTP over localhost to the sidecar. The sidecar intercepts it, establishes an mTLS connection with B's sidecar, and forwards the data securely. The application code remains completely ignorant of the security layer.
