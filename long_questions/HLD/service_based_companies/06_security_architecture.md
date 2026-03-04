# High-Level Design (HLD): Security Architecture

Service-based companies build enterprise B2B software where strict security, auditing, and compliance are paramount. 

## 1. Explain Authentication vs. Authorization.
**Answer:**
*   **Authentication (AuthN):** Validating *who* the user is. (e.g., Logging in with Username/Password, biometrics, or relying on a trusted IdP like Google). Result: The system knows your identity.
*   **Authorization (AuthZ):** Validating *what* the user is allowed to do. (e.g., Are you an Admin? Can you delete this document?). Occurs *after* authentication.

## 2. Explain OAuth 2.0 and OpenID Connect (OIDC).
**Answer:**
Enterprises rarely build their own login systems from scratch; they delegate to identity providers using these protocols.
*   **OAuth 2.0:** An *authorization* framework. It allows an application (e.g., a photo printing app) to access resources on behalf of a user from another server (e.g., Facebook Photos) *without* the user giving their Facebook password to the printing app. It grants an **Access Token**.
*   **OpenID Connect (OIDC):** An identity layer built directly *on top* of the OAuth 2.0 protocol. It adds *authentication*. It allows clients to verify the identity of the end-user based on the authentication performed by an Authorization Server, returning an **ID Token** (usually a JWT).

## 3. How do JSON Web Tokens (JWT) work?
**Answer:**
A JWT is a compact, URL-safe means of representing claims to be transferred between two parties.
*   **Structure:** `Header.Payload.Signature`
    *   *Header:* Contains token type (JWT) and signing algorithm (e.g., RS256).
    *   *Payload:* Contains the claims (user ID, roles, expiration time). *Note: This is just Base64 encoded, NOT encrypted. Never put passwords here.*
    *   *Signature:* Created by taking the encoded header, encoded payload, a secret key, and the algorithm specified in the header.
*   **How it works:** 
    1.  User logs in. Server generates a signed JWT and sends it to the client.
    2.  Client stores it (usually in an `HttpOnly` secure cookie).
    3.  Client sends JWT in the `Authorization: Bearer <token>` header on every request.
    4.  Server verifies the *signature* cryptographically. Because it's signed, the server trusts the data inside it without needing to query the database. (Stateless Authentication).

## 4. How do you secure Web APIs against common attacks? (OWASP Top 10)
**Answer:**
*   **SQL Injection:** An attacker sends SQL commands via input fields. *Solution:* Always use ORMs, Parameterized Queries, or Prepared Statements. Never concatenate strings into SQL queries.
*   **Cross-Site Scripting (XSS):** An attacker injects malicious JavaScript into forms, which is later executed on other users' browsers. *Solution:* Sanitize and escape all user input before rendering it in the UI (React/Angular do this automatically by default).
*   **Cross-Site Request Forgery (CSRF):** An attacker tricks a logged-in user's browser into executing an unwanted action on a trusted site. *Solution:* Use Anti-CSRF tokens injected into forms, or rely on `SameSite` cookie attributes.
*   **DDoS (Distributed Denial of Service):** *Solution:* Use a Web Application Firewall (WAF) like Cloudflare or AWS WAF, implement aggressive API Rate Limiting, and Auto-Scaling to absorb spikes.

## 5. What is IAM (Identity and Access Management) and RBAC?
**Answer:**
*   **IAM:** The overarching framework of policies and technologies for ensuring the right users have the appropriate access to technology resources. Often centralized using Active Directory, Okta, or AWS IAM.
*   **RBAC (Role-Based Access Control):** 
    *   Instead of assigning specific permissions directly to a user (e.g., `User A can edit Document B`), you create Roles.
    *   You assign permissions to Roles (e.g., `Role: Editor -> [create_doc, edit_doc]`).
    *   You assign Users to Roles (e.g., `User A has Role: Editor`).
    *   *Benefit:* Immensely simplifies enterprise management. When a user changes departments, you just change their role, rather than auditing hundreds of individual permissions.
