# 🛡️ Security Architecture — Questions 1–10

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Razorpay, PhonePe, CRED, Amazon, Google — any company handling sensitive data or payments

---

### 1. What is zero-trust security architecture?

"Zero trust is a security model based on the principle: **'never trust, always verify'**. It rejects the traditional perimeter-based model where everything inside the corporate network is trusted. In zero trust, no user, device, or service is trusted by default — regardless of whether it's inside or outside the network.

The traditional model is like a castle with a moat: once you cross the drawbridge (VPN), you're trusted everywhere inside. Zero trust treats every access request as if it came from an untrusted network, requiring authentication and authorization for every resource access — even from inside the corporate network.

Real-world implementation: Google's BeyondCorp replaced VPNs. Engineers access internal tools from any network using strong identity verification (hardware security key + identity certificates). Access is granted based on device posture + user identity, not network location."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Security-focused companies, BFSI sector

#### Indepth
Zero trust pillars:
1. **Verify explicitly:** Always authenticate and authorize based on all available data (identity, location, device health, service, data classification)
2. **Use least privilege:** Limit user and service access to the minimum needed. JIT (just-in-time) access for privileged operations.
3. **Assume breach:** Minimize blast radius. Segment access. End-to-end encryption. Full observability to detect anomalies.

Implementation components:
- **Identity Provider (IdP):** Okta, Azure AD — single source of truth for identity
- **Device trust:** MDM (Mobile Device Management) verifies device health before granting access
- **mTLS:** All service-to-service communication requires mutual TLS (both sides present certificates)
- **Microsegmentation:** Network policies prevent lateral movement (compromised Service A can't reach Service C's port)
- **Policy engine:** OPA (Open Policy Agent) evaluates access based on identity, context, and policy

#### 🗣️ How to Explain in Interview
**Interviewer:** What is zero-trust security architecture?
**Your Response:** "Zero Trust is a strategic shift from the old 'castle-and-moat' perimeter model to a **'never trust, always verify'** approach. In a traditional network, once you’re on the VPN, you’re often trusted implicitly. In Zero Trust, we treat every request—whether it’s from an employee in the office or a remote contractor—as potentially malicious. 

We verify the user's identity through strong MFA, but we also check the **device posture**—is the laptop encrypted? Is the antivirus running? Only when all signals (identity, device, location, and the sensitivity of the data) align do we grant access, and even then, only to that specific resource. It’s about assuming that the network is already compromised and building the system to contain the 'blast radius' of any single breach."

---

### 2. What is OAuth 2.0 and how does it work?

"OAuth 2.0 is an **authorization framework** that enables a third-party application to access a user's resources on another service, without exposing the user's credentials to the third-party.

Classic example: 'Login with Google' on a third-party app. You click it, get redirected to Google's consent screen, approve, and Google gives the third-party app an access token. The app can now access your Google profile — never seeing your Google password.

This is authorization (what can this app access?), not authentication (who is this user?). For authentication, OpenID Connect (OIDC) is built on top of OAuth 2.0 and adds an ID token that contains user identity claims."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company with third-party integrations or SSO

#### Indepth
OAuth 2.0 grant types:
1. **Authorization Code (with PKCE):** Web & mobile apps. User is redirected to auth server, gets code, server exchanges code for tokens. Most secure.
2. **Client Credentials:** Machine-to-machine (service-to-service). No user involved. Service authenticates with client_id + client_secret.
3. **Implicit:** (Deprecated) Browser-only flow, tokens in URL. Replaced by Auth Code + PKCE.
4. **Device Code:** Smart TVs, CLI tools. Device shows a code, user authenticates on another device.

Authorization Code flow:
```
Client → Auth Server: GET /authorize?client_id=&redirect_uri=&scope=
Auth Server → User: Display consent screen
User → Auth Server: Approve
Auth Server → Client: Redirect to redirect_uri?code=abc
Client → Auth Server: POST /token { code=abc, client_secret }
Auth Server → Client: { access_token, refresh_token, expires_in }
Client → Resource Server: GET /user/profile
                Header: Authorization: Bearer {access_token}
```

**JWT (JSON Web Token) as access token:**
- Header: algorithm
- Payload: claims (sub, iat, exp, scope, roles)
- Signature: HMAC-SHA256(header.payload, secret)

The Resource Server verifies the signature locally — no call to Auth Server needed for every request. Token expiry (15-60 min typical) limits the damage if a token is stolen.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is OAuth 2.0 and how does it work?
**Your Response:** "OAuth 2.0 is the industry standard for **delegated authorization**. It allows a user to grant a third-party application access to their data without ever sharing their actual password. This is done through an **Access Token**.

For example, when you see a 'Login with Google' button, you aren't giving the app your Google password. Instead, Google redirects you to their own secure page, you approve the access, and Google sends back a token that the app can use to fetch your name and email. In modern web and mobile apps, we always use the **Authorization Code flow with PKCE** to prevent malicious actors from intercepting those codes. OAuth 2.0 is great because it decouples the service that owns the data from the app that wants to use it, keeping the user's credentials safe."

---

### 3. What is the difference between authentication and authorization?

"**Authentication** answers 'Who are you?' — verifying the identity of the user or service. Methods: username/password, JWT, API keys, certificates, biometrics, hardware tokens.

**Authorization** answers 'What are you allowed to do?' — verifying if the authenticated identity has permission to perform the requested action. Systems: RBAC, ABAC, ACL.

Both must happen on every request. Common mistake: validating the JWT (authentication) but not checking if this user has permission to access this specific resource (authorization). A valid JWT proves you're logged in — it doesn't prove you can read another user's private data."

#### 🏢 Company Context
**Level:** 🟢 Junior – 🟡 Mid | **Asked at:** All companies

#### Indepth
Authentication methods:
- **Password + MFA:** Password verifies identity, TOTP/SMS OTP provides second factor
- **JWT:** Stateless, self-contained token. Verify signature, check expiry, extract claims.
- **API Keys:** Long-lived credentials for service-to-service or developer API access. Scoped to a project.
- **mTLS (mutual TLS):** Both client and server present certificates. Common for service mesh.
- **SSO (SAML/OIDC):** Centralized auth across multiple services. One login for all internal apps.

Authorization models:
- **RBAC (Role-Based):** Permissions are assigned to roles; roles to users. `admin` role → all permissions. Simple to manage.
- **ABAC (Attribute-Based):** Policies based on attributes of user, resource, and environment. "User can access document if user.department == document.department AND time is 9am-6pm". More flexible, more complex.
- **ReBAC (Relationship-Based, used by Google Zanzibar):** "Can user U access document D?" is answered by checking the relationship graph. Google Docs' permission model. Implemented by Authzed, OPA.
- **ACL (Access Control List):** Per-resource list of who can access. Simple for small systems, hard to manage at scale.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the difference between authentication and authorization?
**Your Response:** "The simplest way to put it is: **Authentication** is 'Who are you?', and **Authorization** is 'What are you allowed to do?'. 

Think of it like an airport. Showing your passport to a security officer is **Authentication**—it proves the identity. Your boarding pass is the **Authorization**—it proves you have permission to get on a specific plane and sit in a specific seat. In our software systems, we often see developers validate a JWT and think they’re done. But that just proves the user is logged in. You still have to perform an authorization check on every single request to make sure that 'User A' isn't trying to access 'User B's' private documents."

---

### 4. What is JWT and what are its security considerations?

"JWT (JSON Web Token) is a compact, self-contained token for securely transmitting information between parties as a JSON object. It's digitally signed (HMAC-SHA256 or RSA) so the recipient can verify it hasn't been tampered with.

Structure: `base64(header).base64(payload).signature`. The payload contains claims — `sub` (subject/user ID), `exp` (expiry), `iat` (issued at), and custom claims like `role`, `scope`.

The key point: JWT validation is **stateless** — the resource server verifies the signature using the public key (for RSA) or shared secret (for HMAC). No database lookup needed. This makes it horizontally scalable — any instance can validate any token."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** All companies with user authentication

#### Indepth
JWT security pitfalls:
1. **Algorithm confusion attack:** The token header specifies the algorithm. An attacker changes `alg: RS256` to `alg: none` in a forged token. Fix: Server MUST explicitly specify the expected algorithm, never trust the header's `alg` field.
2. **Weak secret:** HMAC tokens with weak secrets can be brute-forced. Use minimum 256-bit random secrets.
3. **Token leakage:** If stored in localStorage, XSS can steal it. Store access tokens in memory; store refresh tokens in HttpOnly cookies (inaccessible to JavaScript).
4. **Missing expiry:** Tokens without `exp` are valid forever. Always set short expiry (15-60 min).
5. **No revocation (inherent limitation):** JWTs can't be invalidated before expiry. If a token is stolen, it's valid until it expires. Solutions: short TTL + refresh tokens, token blacklist (compromises statelessness), opaque tokens for high-security contexts.

**Token rotation:** Access token expires every 15 minutes. Refresh token (7-30 days) is used to get a new access token silently. Refresh token rotation (invalidate old refresh token on use) limits token theft damage.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is JWT and what are its security considerations?
**Your Response:** "JWT is a compact, self-contained way for two parties to exchange verified data as a JSON object. Since it's digitally signed, we can trust the 'claims' inside it, like the user's ID or their roles. The big architectural advantage is that it’s **stateless**—your microservices can verify the token locally using a public key, which avoids millions of database lookups for session data.

However, from a security standpoint, the main trade-off is that you cannot 'revoke' a JWT easily once it's issued. If a token is stolen, it's valid until it naturally expires. To mitigate this, I always recommend **short-lived access tokens** (like 15 minutes) paired with more secure refresh tokens. We also store access tokens in **HttpOnly cookies** to protect them from XSS attacks, which is a major common vulnerability in modern single-page apps."

---

### 5. What is encryption at rest and encryption in transit?

"**Encryption in transit** (data in motion): Data is encrypted while moving between systems — client to server, service to service, service to database. TLS (Transport Layer Security) is the standard. HTTPS = HTTP + TLS. All modern APIs should require HTTPS. gRPC uses TLS by default.

**Encryption at rest** (data at storage): Data is encrypted when stored on disk — in databases, S3 buckets, backups. AWS RDS enables AES-256 encryption at the storage layer with one checkbox. Encrypting at rest protects against physical disk theft and cloud provider data breaches.

Both are required for compliance (PCI-DSS, GDPR, SOC 2). Both together mean data is protected throughout its lifecycle."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Any company handling PII or payment data

#### Indepth
TLS handshake (simplified):
1. Client hello (supported TLS versions, cipher suites)
2. Server hello + certificate (public key)
3. Client verifies certificate (is it signed by a trusted CA, is the domain correct, is it expired?)
4. Key exchange (client generates session key, encrypts with server's public key, sends)
5. Both sides now have the session key → symmetric encryption for the session

Encryption key management:
- **AWS KMS (Key Management Service):** Create/manage/rotate encryption keys. Keys never leave KMS — KMS encrypts/decrypts via API. Audit log of every key usage in CloudTrail.
- **Envelope encryption:** Generate a data encryption key (DEK) to encrypt data, then encrypt the DEK with a key encryption key (KEK) stored in KMS. Only the encrypted DEK is stored alongside data.
- **Key rotation:** Rotate master keys annually (required for compliance). Old data remains decryptable with old key versions.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is encryption at rest and encryption in transit?
**Your Response:** "In any secure architecture, you have to protect data at every stage of its lifecycle. **Encryption in transit** protects data as it travels across the network—usually via TLS. It ensures that no one can 'eavesdrop' on my users' passwords or credit card numbers while they're being sent to our backend.

**Encryption at rest**, on the other hand, protects the data while it’s sitting on a physical disk—in our database or an S3 bucket. This handles the scenario where the cloud provider's physical hardware is compromised or a backup disk is stolen. I generally use managed services like **AWS KMS** to handle the heavy lifting of key rotation and access control, ensuring that only the specific services that need the data have the 'keys to the kingdom'."

---

### 6. What is the principle of least privilege?

"Least privilege means **giving users, services, and processes only the minimum access required** to perform their function — nothing more.

An order service should be able to read/write the orders table and nothing else — not the users table, not the payments table, not the admin functions. If the order service is compromised, the blast radius is limited to what it can access.

In AWS IAM: Each Lambda function, EC2 instance, or service gets a role with only the specific permissions it needs. Not `*:*` on all resources. A payment processor Lambda needs `dynamodb:PutItem` on one specific table. That's it."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Security-focused companies, BFSI, any AWS-heavy system

#### Indepth
Least privilege in practice:
1. **Database users:** Create one DB user per service with only the tables and operations it needs. Never use root/admin for application connections.
2. **AWS IAM:** Write SCP (Service Control Policies) that deny sensitive permissions by default. Require MFA for all human users. Use roles for services, not long-lived access keys.
3. **Kubernetes RBAC:** Limit what each service account can do within the cluster. The frontend service should not be able to read Secrets.
4. **Secret access:** Use Vault or AWS Secrets Manager with fine-grained policies. A database password should only be accessible by the service that needs it.

**JIT (Just-in-Time) access:** For privileged operations, grant access temporarily for a specific task rather than permanently. "Need to debug production DB? Request access → manager approves → 4-hour access window → automatically revoked." Tools: Teleport, Vault with TTL leases.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is the principle of least privilege?
**Your Response:** "The Principle of Least Privilege is a fundamental security rule: give a user, a service, or an application only the **absolute minimum** access it needs to perform its task, and nothing more. 

For instance, a microservice that generates PDF reports shouldn't have access to the 'User Passwords' table. If that service is ever compromised through a library vulnerability, the attacker is 'stuck' inside that limited scope. In my designs, I use dedicated **IAM roles** for every service and fine-grained database users. It's about 'defense in depth'—if one layer of security fails, the Principle of Least Privilege ensures that the damage is contained and the 'blast radius' is as small as possible."

---

### 7. What is SQL injection and how do you prevent it architecturally?

"SQL injection is an attack where **malicious SQL code is injected through user input**, causing the database to execute unintended commands — accessing, modifying, or deleting data.

Classic example: Login form with username `admin' OR '1'='1`. If the query is `SELECT * FROM users WHERE username = '` + input + `'`, the final query becomes `WHERE username = 'admin' OR '1'='1'` — which returns all users. Attacker bypasses authentication.

This isn't just a code bug — it's an architectural failure. Defense must be at multiple layers: parameterized queries (never concatenate user input into SQL), ORMs (they use parameterized queries by default), input validation, WAF (Web Application Firewall), and database user privileges (the app user can't DROP TABLE)."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Backend roles at any company

#### Indepth
Prevention layers (defense in depth):
1. **Parameterized queries (primary defense):**
```go
// VULNERABLE:
db.QueryRow("SELECT * FROM users WHERE email = '" + email + "'")

// SAFE (parameterized):
db.QueryRow("SELECT * FROM users WHERE email = ?", email)
```

2. **ORM with parameterized bindings:** GORM, Hibernate, SQLAlchemy all use parameterized queries internally. Don't bypass the ORM with raw string concatenation.

3. **Input validation:** Validate and sanitize all input. Whitelist expected characters. Reject anything unusual.

4. **Least privilege DB user:** App's DB user can only SELECT/INSERT/UPDATE on specific tables. Cannot DROP, ALTER, or access system tables.

5. **WAF (Web Application Firewall):** AWS WAF, Cloudflare WAF detect common SQL injection patterns in HTTP requests and block them at the edge — before they reach your application.

6. **Stored procedures:** Parameterized stored procedures prevent injection, but modern apps prefer parameterized queries.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is SQL injection and how do you prevent it architecturally?
**Your Response:** "SQL injection happens when an attacker tries to trick your database by injecting malicious SQL code into an input field, like a login form. If you're building queries by concatenating strings, you're opening the door for them to bypass passwords or even delete your entire database.

To prevent this architecturally, I ensure we always use **parameterized queries** or a mature ORM like Gorm or Hibernate. We also apply 'Defense in Depth' by using a **Web Application Firewall (WAF)** to block common attack patterns at the edge. Finally, we make sure the database user our app uses doesn't have permissions to 'DROP' tables or access system schemas. By securing the code, the network, and the database itself, we create multiple layers of protection."

---

### 8. What is rate limiting from a security perspective?

"From a security perspective, rate limiting **prevents brute-force attacks, credential stuffing, DDoS, and API abuse** by limiting how many requests a client can make in a time window.

Specific security applications: Login endpoint rate limiting (prevent brute-force password attacks — max 5 attempts per 5 minutes per account), OTP/SMS endpoint (prevent OTP bombing — max 3 SMS per 10 minutes), password reset (prevent account enumeration and abuse), payment API (prevent card testing attacks).

Security-focused rate limiting differs from performance rate limiting: you need IP-level blocking, not just per-user. An attacker running a botnet uses different IPs for each attempt."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Razorpay, PhonePe, CRED, any fintech

#### Indepth
Security rate limiting patterns:
1. **Account lockout:** Lock account after N failed login attempts. Risk: denial-of-service (attacker locks out legitimate users). Mitigation: exponential backoff instead of hard lockout, captcha after threshold.
2. **IP-based rate limiting:** Block IPs exceeding threshold. Risk: shared IPs (NAT). Mitigation: progressive penalties (slow down) rather than hard blocks.
3. **Device fingerprinting:** Rate limit by device fingerprint, not just IP. Harder to bypass.
4. **CAPTCHA challenges:** Trigger CAPTCHA after suspicious activity rather than hard blocking.

**OWASP API Security Top 10** — rate limiting covers:
- API2: Broken Authentication (brute-force prevention)
- API4: Lack of Resources & Rate Limiting (DDoS prevention)
- API7: Security Misconfiguration (no rate limiting on sensitive endpoints)

#### 🗣️ How to Explain in Interview
**Interviewer:** What is rate limiting from a security perspective?
**Your Response:** "From a security standpoint, rate limiting is a shield against automated abuse—specifically **brute-force attacks** and **credential stuffing**. If a hacker is trying to crack a user's password, they’ll use bots to try thousands of combinations in seconds. 

By setting an aggressive rate limit on sensitive endpoints—like only 5 failed login attempts per account every 10 minutes—we break their business model. I also implement rate limiting on 'expensive' resources like SMS OTPs to prevent cost abuse. Architecturally, I prefer to handle this far 'upstream' at the API Gateway or a CDN like Cloudflare, so my backend servers don't even have to process the malicious traffic."

---

### 9. What is OWASP Top 10 and how does it affect architecture?

"OWASP Top 10 is a standard awareness document listing the **10 most critical web application security risks**. The 2021 list includes Injections, Broken Authentication, Security Misconfiguration, Insecure Design, Vulnerable Components, and more.

From an architecture perspective: these aren't individual bugs to patch — they're organizational and design failures. **Insecure Design** (A04:2021) is explicitly about architecture — designing systems without security modeling leads to systemic vulnerabilities.

Architectural mitigations: Threat modeling in the design phase (not after), defense-in-depth (multiple security layers), secure defaults (HTTPS enforced, secure cookies, CSP headers), dependency scanning in CI/CD pipeline, secrets management."

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Security audits, product companies with compliance requirements

#### Indepth
Top 5 with architectural implications:
1. **A01 Broken Access Control:** Missing authorization checks. Architectural fix: Centralize authorization checks in middleware/gateway; use OPA for policy enforcement; test authorization at every API boundary.
2. **A02 Cryptographic Failures:** Sensitive data in plaintext, weak algorithms. Fix: Enforce HTTPS, encrypt PII at rest (Vault, KMS), use strong algorithms (AES-256, SHA-256), never roll your own crypto.
3. **A03 Injection (SQLi, NoSQLi, XSS):** User input treated as code. Fix: Never concatenate user input; use parameterized queries, input validation, output encoding, WAF.
4. **A04 Insecure Design:** Threat modeling not done. Fix: Threat model during design phase (STRIDE methodology), establish security requirements as NFRs, design review checklist.
5. **A06 Vulnerable Components:** Outdated dependencies with known CVEs. Fix: Dependency scanning in CI (Snyk, Dependabot), automated dependency updates, software composition analysis (SCA).

#### 🗣️ How to Explain in Interview
**Interviewer:** What is OWASP Top 10 and how does it affect architecture?
**Your Response:** "The OWASP Top 10 is the definitive list of the most critical security risks facing web applications today. As a senior engineer, I use it not just as a 'bug checklist,' but as a guide for **Security-by-Design**. 

For example, knowing that 'Broken Access Control' is the #1 risk, I make sure we have a centralized authorization middleware rather than letting every developer write their own logic. We also integrate tools into our CI/CD pipeline to scan for 'Vulnerable Components' and ensure our 'Secrets Management' is handled correctly. It’s about building a culture where security is a shared responsibility throughout the entire development lifecycle, rather than just an audit that happens at the end."

---

### 10. What is data masking and tokenization?

"**Data masking** replaces sensitive data with a realistic but fictitious equivalent. It's used in non-production environments — developers get a full-size copy of the production database, but all PII is replaced with fake data. Customer name 'Dhruv Shah' becomes 'Ramesh Kumar', phone `9876543210` becomes `9812345678`.

**Tokenization** replaces sensitive data with a **non-sensitive placeholder (token)** that maps back to the original in a secure token vault. Unlike encryption (which can be reversed with the key), tokenization stores the mapping in a separate, access-controlled vault. Used heavily in payments — a credit card number is replaced with a token. The merchant stores the token; the actual card number is only in the payment processor's vault.

PCI-DSS mandates tokenization or encryption for card data (PANs) — merchants who tokenize don't need to store raw card data at all, dramatically reducing their PCI scope."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Fintech companies (Razorpay, PhonePe), healthcare, any PCI/GDPR-regulated company

#### Indepth
Data masking types:
- **Static masking:** Done once, produces a masked copy (for dev/test environments). Production → mask → test DB.
- **Dynamic masking:** Applied at query time — different roles see different levels of data. A call center agent sees `****-****-****-1234`, a fraud analyst sees the full number, a regular employee sees `[REDACTED]`.

Tokenization vs Encryption comparison:
| Aspect | Encryption | Tokenization |
|--------|-----------|--------------|
| Reversibility | Possible with key | Only via token vault lookup |
| Key exposure | Key theft = data breach | Key stolen but vault not accessed = safe |
| Format | Changed (ciphertext) | Preserved (token looks like card number) |
| DB scope | In-place (same DB) | Cross-service (vault is separate) |
| PCI scope reduction | Partial | Significant (raw card never touches merchant) |

GDPR right to erasure: Tokenization enables "pseudonymous deletion" — you delete the token-to-PII mapping from the vault, making all stored tokens meaningless. Simpler than hunting down PII across many tables.

#### 🗣️ How to Explain in Interview
**Interviewer:** What is data masking and tokenization?
**Your Response:** "Data masking and tokenization are critical for reducing the exposure of sensitive PII. **Data masking** is primarily for our developers—it replaces real customer names and emails with fake but realistic data so we can test our software without ever seeing real private info. 

**Tokenization**, however, is a security powerhouse for production systems. It replaces a high-value piece of data, like a credit card number, with a random 'token.' The actual card data is stored in a separate, ultra-secure 'vault' that only the payment processor can talk to. This means that if our main database were ever breached, the hacker would only find useless tokens, not our customers' credit cards. It’s the single most effective way to reduce our 'PCI compliance scope' and protect our users."
