# 🔒 Security — Advanced Interview Questions (Product-Based Companies)

This document covers advanced security concepts for product-based company interviews (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay). Targeted at 3–10 years of experience rounds.

---

### Q1: Explain OAuth 2.0 and OpenID Connect (OIDC). What is the PKCE extension?

**Answer:**
**OAuth 2.0** is an **authorization** framework (delegated access).
**OpenID Connect (OIDC)** is an **authentication** layer built on top of OAuth 2.0.

**OAuth 2.0 Authorization Code Flow (standard):**
1. User clicks "Login with Google".
2. App redirects user to Google's Authorization Server.
3. User authenticates and grants permission.
4. Google redirects back to App with an **Authorization Code**.
5. App (backend) exchanges the Code + Client Secret for an **Access Token**.
6. App uses Access Token to call Google APIs on behalf of the user.

**Why the "Code" step?** 
If Google simply returned the Access Token in the URL redirect (Implicit Flow), it could be intercepted via browser history or referer headers. The back-channel exchange (Step 5) keeps the token secure.

**OpenID Connect (OIDC):**
During Step 5, OIDC adds an **ID Token** (a JWT) alongside the Access Token. The ID Token proves *who* the user is (authentication), while the Access Token grants API access (authorization).

**PKCE (Proof Key for Code Exchange) Extension:**
Required for Mobile Apps and SPAs (Single Page Apps) that cannot securely store a Client Secret.
- Client generates a random `code_verifier` and its hash `code_challenge`.
- Sends `code_challenge` in Step 2.
- Sends `code_verifier` in Step 5.
- The Authorization Server verifies the hash. This prevents an attacker who intercepts the Authorization Code from exchanging it, because the attacker doesn't know the original `code_verifier`.

---

### Q2: What is mTLS (Mutual TLS) and why is it essential for microservices?

**Answer:**
Standard TLS (HTTPS) physically authenticates the **server** to the **client**. The client verifies the server's certificate.

**Mutual TLS (mTLS)** requires **both** parties to authenticate each other:
1. Server verifies Client's certificate.
2. Client verifies Server's certificate.

**Why it's essential for Microservices (Zero-Trust Architecture):**
- In traditional networks, the perimeter is secured (firewall), but internal traffic is unencrypted (HTTP). Once an attacker breaches the perimeter, they can move laterally without restriction.
- **Zero-Trust** assumes the internal network is compromised.
- mTLS ensures that every service-to-service call is:
  1. **Encrypted** (preventing sniffing).
  2. **Authenticated** (Service A mathematically proves it is Service A to Service B).
- Usually managed transparently by a **Service Mesh** (like Istio), which automatically issues, rotates, and validates certificates for every pod.

---

### Q3: Contrast CSRF and XSS. How does SameSite cookie attribute mitigate CSRF?

**Answer:**

**XSS (Cross-Site Scripting):**
- Attacker runs malicious JavaScript in the victim's browser on the **trusted site**.
- Focuses on stealing data (cookies, tokens) or performing local actions.
- **Defense:** Output encoding, CSP, HttpOnly cookies.

**CSRF (Cross-Site Request Forgery):**
- Attacker tricks the victim's browser into making an unwanted request to a **trusted site** where the user is already authenticated.
- Focuses on state-changing actions (e.g., transferring money, changing password).
- Example: Attacker's site has `<img src="https://bank.com/transfer?amount=1000&to=attacker">`. Browser automatically attaches the victim's `bank.com` cookie to the request.

**Defense against CSRF:**
1. **Anti-CSRF Tokens**: A hidden, random token in forms that the attacker cannot guess. The server verifies it matches the user's session.
2. **SameSite Cookie Attribute (Modern default):**
   - `SameSite=None`: Cookie sent with all cross-site requests (requires `Secure`).
   - `SameSite=Lax`: (Default) Cookie NOT sent for cross-site POST requests. Sent for top-level navigations (GET).
   - `SameSite=Strict`: Cookie is ONLY sent if the request originates from the same domain. Extremely effective against CSRF.

---

### Q4: What are timing attacks and replay attacks? How do you prevent them?

**Answer:**

**Timing Attack:**
A side-channel attack where the attacker measures how long the server takes to respond to infer secret data (e.g., passwords, HMACs).
*Example:* A naive string comparison checks character by character. If the first character of the password matches, it takes slightly longer to fail on the second character. Attacker brute-forces one character at a time based on response time.
*Prevention:* Always use **constant-time comparison** functions (e.g., `MessageDigest.isEqual()` in Java, `hmac.Compare()` in Go).

**Replay Attack:**
An attacker intercepts a valid, encrypted request (e.g., "Transfer $100") and blindly sends it again (replays it) without decrypting it, hoping the server processes it a second time.
*Prevention:*
1. **Nonces (Numbers used once):** Client includes a unique, random string. Server rejects requests with previously seen nonces.
2. **Timestamps:** Request includes a timestamp and signature. Server rejects requests older than a few minutes.
3. **Idempotency keys:** As discussed in distributed systems.

---

### Q5: How do passwords get compromised? Explain salt, pepper, and modern hashing algorithms.

**Answer:**

**Naive approach (MD5 / SHA-256):**
Extremely fast. Attackers use Rainbow Tables (precomputed hashes of common passwords) or GPUs to crack billions of hashes per second.

**Salt:**
A random string unique to each user, appended to the password before hashing: `hash(password + salt)`.
- **Purpose:** Prevents Rainbow Table attacks and ensures two users with the same password have different hashes. Salt is stored in plaintext next to the hash.

**Pepper:**
A global secret added to the password, but **never stored in the database** (stored in an environment variable or KMS).
- **Purpose:** If the database is stolen, the attacker still cannot crack the passwords without also compromising the application servers or KMS to get the pepper.

**Modern Iterative / Memory-Hard Hashing:**
These algorithms are intentionally **slow** and require high CPU/RAM, resisting GPU brute-forcing.
1. **PBKDF2**: CPU-hard. Standard in many older frameworks.
2. **Bcrypt**: CPU-hard, limits maximum password length. Standard in Spring Security and modern web frameworks.
3. **Argon2 / Scrypt**: **Memory-hard**. Requires large amounts of RAM to calculate the hash, making it extremely expensive and infeasible for GPUs/ASICs to crack in parallel. (Argon2 is the current industry gold standard).

---

### Q6: Compare AES-GCM vs ChaCha20-Poly1305. When would you choose each?

**Answer:**

**AES-GCM (Advanced Encryption Standard - Galois/Counter Mode):**
- **Hardware acceleration** on most modern CPUs (AES-NI)
- **128-bit** block size, **256-bit** key max
- **96-bit** nonce recommended (12 bytes)
- **Performance**: Excellent on servers with AES-NI
- **Adoption**: TLS 1.3 standard, widely supported

**ChaCha20-Poly1305:**
- **Software-optimized** stream cipher
- **256-bit** key, **96-bit** nonce (12 bytes)
- **Performance**: Consistently fast on all hardware (mobile, IoT)
- **Adoption**: TLS 1.3 alternative, mobile networks (5G), QUIC

**Decision Matrix:**

| Scenario | Choose | Why |
|---|---|---|
| **Server-side encryption** | **AES-GCM** | Hardware acceleration makes it 3-5x faster |
| **Mobile apps** | **ChaCha20-Poly1305** | Consistent performance, no hardware dependency |
| **IoT devices** | **ChaCha20-Poly1305** | No AES hardware, power-efficient |
| **High-throughput APIs** | **AES-GCM** | Leverages server CPU acceleration |
| **Cross-platform compatibility** | **ChaCha20-Poly1305** | Uniform performance across devices |

**Implementation (Go):**
```go
// AES-GCM (server-side)
func encryptAESGCM(key, plaintext []byte) ([]byte, error) {
    block, _ := aes.NewCipher(key)
    gcm, _ := cipher.NewGCM(block)
    nonce := make([]byte, gcm.NonceSize())
    io.ReadFull(rand.Reader, nonce)
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// ChaCha20-Poly1305 (mobile/IoT)
func encryptChaCha20(key, plaintext []byte) ([]byte, error) {
    aead, _ := chacha20poly1305.NewX(key) // XChaCha20 for larger nonce
    nonce := make([]byte, aead.NonceSize())
    io.ReadFull(rand.Reader, nonce)
    return aead.Seal(nonce, nonce, plaintext, nil), nil
}
```

---

### Q7: What are the security implications of nonce reuse in GCM/CTR modes?

**Answer:**
**Nonce reuse is catastrophic** for counter-based modes (GCM, CTR). It completely breaks confidentiality.

**Why it's so dangerous:**
- Both modes use XOR: `ciphertext = plaintext ⊕ keystream`
- Keystream = `encrypt(counter + nonce)`
- **Same nonce + same counter = same keystream**
- Attacker can XOR two ciphertexts to eliminate keystream:
  ```
  C1 ⊕ C2 = (P1 ⊕ K) ⊕ (P2 ⊕ K) = P1 ⊕ P2
  ```

**Real-world attack scenario:**
```
Two messages encrypted with same nonce:
C1 = P1 ⊕ K  (known: "GET /users/123")
C2 = P2 ⊕ K  (unknown: secret API key)

Attacker computes:
P1 ⊕ P2 = C1 ⊕ C2
Since P1 is known, attacker derives P2!
```

**Prevention strategies:**

**1. Counter-based nonce generation:**
```go
// Atomic counter ensures uniqueness
var nonceCounter uint64
func nextNonce() []byte {
    n := atomic.AddUint64(&nonceCounter, 1)
    return make([]byte, 12) // 96-bit nonce
}
```

**2. Random with collision detection:**
```go
func generateRandomNonce() []byte {
    nonce := make([]byte, 12)
    for {
        rand.Read(nonce)
        if !isNonceUsed(nonce) { // Check in-memory set
            markNonceUsed(nonce)
            return nonce
        }
    }
}
```

**3. Use XChaCha20-Poly1305:**
- 192-bit nonce makes collision practically impossible
- Can use random nonce safely

**Production best practices:**
- **Never reuse keys** with same nonce
- **Use key rotation** to limit damage
- **Monitor nonce usage** in distributed systems
- **Prefer AEADs** with larger nonce spaces

---

### Q8: Design an encryption service for a microservices architecture. What are the key considerations?

**Answer:**

**Architecture Overview:**
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Service A     │    │   Service B     │    │   Service C     │
│                 │    │                 │    │                 │
│  encrypt(data)──┼────┼──encrypt(data)──┼────┼──encrypt(data)──┼─────►
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                    ┌─────────────────────┐
                    │  Encryption Service │
                    │                     │
                    │  • Key Management   │
                    │  • Algorithm Choice │
                    │  • Audit Logging    │
                    └─────────────────────┘
```

**Key Design Considerations:**

**1. Key Management:**
- **Never store keys in application code**
- Use **KMS** (AWS KMS, Azure Key Vault, GCP KMS)
- **Envelope encryption**: KMS encrypts data keys, data keys encrypt payload
- **Key rotation**: Automatic every 90 days
- **Key separation**: Different keys per service/data type

**2. Algorithm Selection:**
```yaml
production:
  algorithm: "AES-256-GCM"
  key_size: 256
  nonce_size: 12
  
legacy_compatibility:
  algorithm: "AES-256-CBC"
  additional_protection: "HMAC-SHA256"
  
mobile_iot:
  algorithm: "XChaCha20-Poly1305"
  key_size: 256
```

**3. API Design:**
```go
type EncryptionRequest struct {
    Data        []byte            `json:"data"`
    KeyID       string            `json:"key_id"`
    Algorithm   string            `json:"algorithm"`
    Context     map[string]string `json:"context"` // For audit
}

type EncryptionResponse struct {
    Ciphertext  []byte `json:"ciphertext"`
    KeyID       string `json:"key_id"`
    Algorithm   string `json:"algorithm"`
    Timestamp   int64  `json:"timestamp"`
    RequestID   string `json:"request_id"`
}
```

**4. Security Controls:**
- **mTLS** between services
- **Rate limiting** to prevent DoS
- **Audit logging** of all encryption operations
- **Input validation** (max payload size, format checks)
- **Circuit breaker** for KMS failures

**5. Performance Optimizations:**
- **Connection pooling** to KMS
- **Local caching** of encrypted data keys (TTL: 1 hour)
- **Batch operations** for bulk encryption
- **Async processing** for non-critical operations

**6. Monitoring & Alerting:**
- **Encryption latency** (P95 < 50ms)
- **Key rotation failures**
- **Unusual access patterns**
- **Error rates** by service/algorithm

**Implementation Example:**
```go
type EncryptionService struct {
    kmsClient    kms.Client
    keyCache     *sync.Map
    auditLogger  logger.Logger
    metrics      metrics.Collector
}

func (s *EncryptionService) Encrypt(ctx context.Context, req EncryptionRequest) (*EncryptionResponse, error) {
    // 1. Validate request
    if err := s.validateRequest(req); err != nil {
        return nil, err
    }
    
    // 2. Get or create data key
    dataKey, err := s.getDataKey(ctx, req.KeyID)
    if err != nil {
        return nil, err
    }
    
    // 3. Encrypt data
    ciphertext, err := s.encryptWithKey(req.Data, dataKey, req.Algorithm)
    if err != nil {
        return nil, err
    }
    
    // 4. Audit log
    s.auditLogger.Info("encryption_completed",
        "service", req.Context["service"],
        "key_id", req.KeyID,
        "algorithm", req.Algorithm,
        "data_size", len(req.Data),
    )
    
    return &EncryptionResponse{
        Ciphertext: ciphertext,
        KeyID:      req.KeyID,
        Algorithm:  req.Algorithm,
        Timestamp:  time.Now().Unix(),
        RequestID:  uuid.New().String(),
    }, nil
}
```

**Deployment Considerations:**
- **Multi-region** replication for global services
- **Blue-green** deployments for key rotation
- **Health checks** including KMS connectivity
- **Graceful degradation** during KMS outages

---

*Prepared for technical rounds at product-based companies (Google, Meta, Amazon, Flipkart, Swiggy, CRED, Zepto, Razorpay, Groww).*
