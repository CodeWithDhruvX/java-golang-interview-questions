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
