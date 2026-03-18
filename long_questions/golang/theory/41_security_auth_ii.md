# 🟢 Go Theory Questions: 801–820 Security & Authentication II

## 801. How do you use JWT securely in Go APIs?

**Answer:**
We use `golang-jwt/jwt`.
Key security steps:
1.  **Algorithm**: Enforce `HS256` or `RS256` in the parser. `jwt.Parse(token, func(t *jwt.Token) { if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { return nil, fmt.Errorf("bad alg") } })`.
2.  **Secret Management**: Load secret from ENV, never hardcode.
3.  **Claims**: Check `exp` (Expiry) and `iss` (Issuer) strictly.
4.  **Rotate**: Support key rotation by using a `KeyFunc` that looks up the current valid key ID (kid).

---

## 802. How do you manage CSRF protection in a Go web app?

**Answer:**
(See Q 504).
Middleware `gorilla/csrf`.
It injects a randomized token into headers/cookies.
For **Single Page Apps (SPA)**:
Cookie: `XSRF-TOKEN` (readable by JS).
Header: `X-XSRF-TOKEN` (sent by JS).
The backend validates that the Header == Cookie. The browser's Same-Origin Policy forbids attacker sites from reading the cookie, so they can't construct the header.

---

## 803. How do you handle XSS prevention in Go templates?

**Answer:**
`html/template` is safer than `text/template`.
It uses **Context-Aware Escaping**.
`{{ .Var }}`.
If Var is `<script>alert(1)</script>`, Go renders it as `&lt;script&gt;...`.
**Risk**: `template.HTML`. Only use this type if the content was run through a strict sanitizer like **Bluemonday** to strip dangerous tags while keeping formatting (b, i, u).

---

## 804. How do you implement OAuth 2.0 flows in Go?

**Answer:**
Standards: `golang.org/x/oauth2`.
Flow: **Authorization Code Flow**.
1.  Redirect user to Provider (Google).
2.  User approves. Provider redirects back with `?code=xyz`.
3.  Go server swaps `code` for `Access Token` via backend channel (Client Secret).
4.  Use Access Token to fetch User Info.
We store the User Info in a session and issue our own App JWT.

---

## 805. How do you encrypt/decrypt sensitive data in Go?

**Answer:**
Standard: **AES-GCM**.
```go
block, _ := aes.NewCipher(key)
gcm, _ := cipher.NewGCM(block)
nonce := make([]byte, gcm.NonceSize())
io.ReadFull(rand.Reader, nonce)
ciphertext := gcm.Seal(nonce, nonce, data, nil)
```
The `Seal` function encrypts AND signs (Auth Tag). We prepend the Nonce to the ciphertext so we can decrypt it later (`Open`).

---

## 806. What’s the use of `crypto/rand` vs `math/rand`?

**Answer:**
`math/rand`: Seeded PRNG. Deterministic. Fast. Use for simulations, fuzzing, games.
`crypto/rand`: CSPRNG. Reads OS entropy (`/dev/urandom`). Slow. Blocking (theoretically).
**Security Rule**: ALWAYS use `crypto/rand` for anything related to Auth, Keys, Salts, or IDs. If an attacker predicts your "Random" session ID, they account takeover your users.

---

## 807. How do you manage TLS certs in Go servers?

**Answer:**
Development: `GenerateCert` (Self-signed).
Production: **Let's Encrypt** (ACME).
Go has `golang.org/x/crypto/acme/autocert`.
```go
m := &autocert.Manager{
    Cache:      autocert.DirCache("certs"),
    HostPolicy: autocert.HostWhitelist("example.com"),
}
s := &http.Server{TLSConfig: m.TLSConfig()}
```
This automatically negotiates, downloads, and renews SSL certificates from Let's Encrypt with zero manual intervention.

---

## 808. How do you validate tokens in Go microservices?

**Answer:**
Service A calls Service B.
Method 1: **Introspection**. B calls Identity Provider (IdP): "Is this token valid?". (Slow).
Method 2: **Local Validation** (JWKS).
B downloads the Public Keys (JWK Set) from IdP once (caches them).
B validates the JWT signature locally using the RSA Public Key. This is fast (0 latency) and stateless.

---

## 809. How do you securely store API keys in Go apps?

**Answer:**
We **Hash** them, just like passwords.
When we issue a key `sk_live_123`, we show it once to the user.
We store `sha256("sk_live_123")` in the DB.
On request, we hash the incoming key and compare.
This ensures that if our DB is leaked, the attacker cannot use the keys to call our API.

---

## 810. How do you create and validate secure cookies?

**Answer:**
Use `gorilla/securecookie`.
It handles **Encryption** (AES) and **Signing** (HMAC).
`s := securecookie.New(hashKey, blockKey)`
`s.Encode("session", value)` -> produce opaque string.
`s.Decode("session", cookieVal, &value)`.
This prevents users from tampering with cookie data (e.g., changing `admin=false` to `admin=true`) and hides the contents.

---

## 811. How do you implement role-based access control in Go?

**Answer:**
Middleware + Claims.
JWT contains `roles: ["audit", "viewer"]`.
Middleware `RequireRole("admin")`.
Checks: `if !contains(claims.Roles, "admin") { return 403 }`.
For complex policies (ABAC), we use **Open Policy Agent (OPA)** or **Casbin**.
`e.Enforce(sub, obj, act)` -> Casbin checks `policy.csv` to see if User can Edit Document.

---

## 812. How do you generate a secure random token in Go?

**Answer:**
Bit of entropy.
1.  Read 32 bytes from `crypto/rand`.
2.  Encode to String.
`base64.RawURLEncoding.EncodeToString(bytes)`.
This gives a ~43 char URL-safe string.
Do not use `uuid.New()` (v4) for *secrets* (session tokens) as UUIDs are designed for uniqueness, not unpredictability (though v4 is random, a 32-byte CSPRNG string has more entropy).

---

## 813. How do you prevent replay attacks with Go?

**Answer:**
(See Q 518).
Require `timestamp` + `signature` headers.
Server checks:
1.  Signature is valid (Auth).
2.  Timestamp is fresh (< 5s old).
3.  Nonce (optional) hasn't been seen in Redis.
This prevents an attacker from grabbing a valid HTTP packet off the wire and resending it later to repeat an action (e.h. "Pay $50").

---

## 814. How do you audit Go applications for security issues?

**Answer:**
1.  **Static Analysis**: `gosec`.
    `gosec ./...` finds hardcoded credentials, weak crypto, unhandled errors.
2.  **Dependency Scan**: `govulncheck`. Finds CVEs in imported modules.
3.  **Fuzzing**: Go Fuzzing (1.18+) to crash parsers with garbage input.

---

## 815. How do you apply security headers in Go HTTP servers?

**Answer:**
Middleware.
```go
w.Header().Set("X-Content-Type-Options", "nosniff")
w.Header().Set("X-Frame-Options", "DENY")
w.Header().Set("Content-Security-Policy", "default-src 'self'")
```
We usually use a library like `unrolled/secure` which sets these standard hardening headers (HSTS, CSP, Referrer) automatically with safe defaults.

---

## 816. How do you secure gRPC endpoints in Go?

**Answer:**
(See Q 570 - Interceptors).
1.  **TLS**: Mandatory encryption.
2.  **Auth Interceptor**: Extract `authorization` metadata. Validate Token.
3.  **Auditing**: Log "Who did what" in the interceptor.
4.  **Rate Limiting**: Per-user limits to prevent DoS.

---

## 817. How do you mock HTTP clients in Go tests?

**Answer:**
We mock the **Transport**, not the Client.
```go
type MockTransport struct{}
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    return &http.Response{StatusCode: 200, Body: ...}, nil
}
client := &http.Client{Transport: &MockTransport{}}
```
This intercepts the network call at the lowest level, preventing any real socket connection.

---

## 818. How do you achieve high test coverage in Go?

**Answer:**
1.  **Table Driven Tests**: easy to add edge cases.
2.  **Interface Injection**: Mock dependencies.
3.  **Integration Tests**: For DB/API layers.
4.  **Gate**: Fail CI if coverage < 80%.
We focus on **Branch Coverage** (paths taken) rather than just line coverage.

---

## 819. How do you test race conditions in Go?

**Answer:**
`go test -race ./...`.
The race detector instruments memory accesses.
It crashes the test if two goroutines access the same variable concurrently and at least one is a Write.
It is **Runtime Detection** (dynamic), not static. It only finds races that actually happen during the test execution, so running the test 100 times (`-count=100`) increases confidence.

---

## 820. How do you benchmark functions in Go?

**Answer:**
(See Q 521).
`func BenchmarkX(b *testing.B)`.
Tips:
1.  **ResetTimer**: Call `b.ResetTimer()` after expensive setup.
2.  **StopTimer/StartTimer**: Pause during non-measured work.
3.  **RunParallel**: `b.RunParallel(func(pb *testing.PB) { ... })` to saturate all CPU cores and test contention.
Measurement: `ns/op` (Time) and `B/op` (Allocations).

---

## 821. What encryption modes are available in Go's crypto package? How do you implement them?

**Answer:**
Go's `crypto/cipher` package provides several encryption modes for block ciphers like AES.

**Available Modes in Go:**

| Mode | Package | Use Case | Security Level |
|---|---|---|---|
| **GCM** | `crypto/cipher` | **Modern standard** - authenticated encryption | ✅ Excellent |
| **CBC** | `crypto/cipher` | Legacy compatibility | ⚠️ Needs HMAC |
| **CTR** | `crypto/cipher` | Stream cipher, high performance | ⚠️ Needs HMAC |
| **CFB** | `crypto/cipher` | Stream mode, self-synchronizing | ⚠️ Needs HMAC |
| **OFB** | `crypto/cipher` | Stream mode, error propagation | ⚠️ Needs HMAC |

**Implementation Examples:**

**1. AES-GCM (Recommended):**
```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func EncryptAESGCM(key, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    // Prepend nonce to ciphertext
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func DecryptAESGCM(key, ciphertext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}
```

**2. AES-CBC with HMAC (Legacy):**
```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/hmac"
    "crypto/sha256"
    "crypto/rand"
    "encoding/binary"
)

func EncryptAESCBC(key, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    // Add PKCS#7 padding
    padding := aes.BlockSize - len(plaintext)%aes.BlockSize
    plaintext = append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)
    
    // Generate random IV
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }
    
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
    
    // Calculate HMAC
    mac := hmac.New(sha256.New, key)
    mac.Write(ciphertext)
    return append(ciphertext, mac.Sum(nil)...), nil
}
```

**3. CTR Mode for High Performance:**
```go
func EncryptAESCTR(key, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    // Generate IV (counter)
    iv := make([]byte, aes.BlockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }
    
    ciphertext := make([]byte, len(plaintext))
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(ciphertext, plaintext)
    
    // Prepend IV to ciphertext
    return append(iv, ciphertext...), nil
}
```

**Best Practices in Go:**
- **Always use GCM** for new code
- **Never reuse IVs/nonce** with the same key
- **Use `crypto/rand`** for generating random values
- **Handle authentication** separately for non-AEAD modes
- **Prefer `crypto/aes`** over third-party implementations

---

## 822. How do you handle key rotation in Go applications?

**Answer:**
Key rotation is critical for security. Here's a production-ready approach:

**1. Key Versioning Strategy:**
```go
type EncryptionKey struct {
    ID        string    `json:"id"`
    Version   int       `json:"version"`
    Key       []byte    `json:"-"` // Never serialize
    Algorithm string    `json:"algorithm"`
    CreatedAt time.Time `json:"created_at"`
    ExpiresAt time.Time `json:"expires_at"`
}

type KeyManager struct {
    currentKey   *EncryptionKey
    keys         map[string]*EncryptionKey // version -> key
    rotationTime time.Duration
    mutex        sync.RWMutex
}
```

**2. Automatic Rotation:**
```go
func (km *KeyManager) StartRotation(ctx context.Context) {
    ticker := time.NewTicker(km.rotationTime)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            km.rotateKey()
        }
    }
}

func (km *KeyManager) rotateKey() {
    km.mutex.Lock()
    defer km.mutex.Unlock()
    
    // Create new key
    newKey := &EncryptionKey{
        ID:        uuid.New().String(),
        Version:   km.currentKey.Version + 1,
        Key:       generateKey(),
        Algorithm: "AES-256-GCM",
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(90 * 24 * time.Hour), // 90 days
    }
    
    // Archive old key but keep it for decryption
    km.keys[km.currentKey.ID] = km.currentKey
    km.currentKey = newKey
    
    log.Info("Key rotated", "new_version", newKey.Version)
}
```

**3. Encryption with Key Versioning:**
```go
type EncryptedData struct {
    KeyVersion string `json:"key_version"`
    Algorithm  string `json:"algorithm"`
    Nonce      []byte `json:"nonce"`
    Ciphertext []byte `json:"ciphertext"`
}

func (km *KeyManager) Encrypt(data []byte) (*EncryptedData, error) {
    km.mutex.RLock()
    key := km.currentKey
    km.mutex.RUnlock()
    
    encrypted, err := encryptWithKey(data, key)
    if err != nil {
        return nil, err
    }
    
    return &EncryptedData{
        KeyVersion: key.ID,
        Algorithm:  key.Algorithm,
        Nonce:      encrypted.Nonce,
        Ciphertext: encrypted.Ciphertext,
    }, nil
}

func (km *KeyManager) Decrypt(encrypted *EncryptedData) ([]byte, error) {
    km.mutex.RLock()
    key, exists := km.keys[encrypted.KeyVersion]
    km.mutex.RUnlock()
    
    if !exists {
        return nil, errors.New("key version not found")
    }
    
    return decryptWithKey(encrypted, key)
}
```

**4. Integration with KMS:**
```go
type KMSKeyManager struct {
    client     *kms.Client
    keyID      string
    cache      map[string][]byte // encrypted data key -> plaintext key
    cacheMutex sync.RWMutex
    cacheTTL   time.Duration
}

func (k *KMSKeyManager) GetDataKey(ctx context.Context) ([]byte, error) {
    // Generate data key in KMS
    result, err := k.client.GenerateDataKey(ctx, &kms.GenerateDataKeyInput{
        KeyId:   &k.keyID,
        KeySpec: types.DataKeySpecAes256,
    })
    if err != nil {
        return nil, err
    }
    
    // Cache plaintext key briefly (5 minutes)
    k.cacheMutex.Lock()
    k.cache[string(result.CiphertextBlob)] = result.Plaintext
    k.cacheMutex.Unlock()
    
    // Schedule cache cleanup
    go func() {
        time.Sleep(k.cacheTTL)
        k.cacheMutex.Lock()
        delete(k.cache, string(result.CiphertextBlob))
        k.cacheMutex.Unlock()
    }()
    
    return result.Plaintext, nil
}
```

**5. Monitoring and Alerting:**
```go
type KeyRotationMetrics struct {
    RotationCount    prometheus.Counter
    DecryptionErrors prometheus.Counter
    KeyAge           prometheus.Gauge
}

func (km *KeyManager) collectMetrics() {
    km.currentKeyAge.Set(time.Since(km.currentKey.CreatedAt).Hours())
    
    // Alert if key is older than rotation period
    if time.Since(km.currentKey.CreatedAt) > km.rotationTime {
        log.Warn("Key rotation overdue", 
            "key_age", time.Since(km.currentKey.CreatedAt),
            "rotation_period", km.rotationTime)
    }
}
```

**Production Checklist:**
- ✅ **Automated rotation** every 90 days
- ✅ **Key versioning** for backward compatibility
- ✅ **Secure storage** (KMS, HashiCorp Vault)
- ✅ **Audit logging** of all key operations
- ✅ **Graceful degradation** during rotation
- ✅ **Monitoring** for rotation failures
- ✅ **Backup and recovery** procedures

---

## 823. How do you implement ChaCha20-Poly1305 in Go for mobile/IoT?

**Answer:**
ChaCha20-Poly1305 is ideal for mobile/IoT because it's software-optimized and doesn't require AES hardware acceleration.

**Implementation:**
```go
import (
    "golang.org/x/crypto/chacha20poly1305"
)

func EncryptChaCha20(key, plaintext []byte) ([]byte, error) {
    // Use XChaCha20 for larger nonce space (24 bytes)
    aead, err := chacha20poly1305.NewX(key)
    if err != nil {
        return nil, err
    }
    
    // Generate random nonce (24 bytes for XChaCha20)
    nonce := make([]byte, aead.NonceSize())
    if _, err := rand.Read(nonce); err != nil {
        return nil, err
    }
    
    // Encrypt and authenticate
    ciphertext := aead.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}

func DecryptChaCha20(key, ciphertext []byte) ([]byte, error) {
    aead, err := chacha20poly1305.NewX(key)
    if err != nil {
        return nil, err
    }
    
    nonceSize := aead.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return aead.Open(nil, nonce, ciphertext, nil)
}
```

**Performance Comparison:**
```go
func BenchmarkEncryptAESGCM(b *testing.B) {
    key := make([]byte, 32)
    plaintext := make([]byte, 1024)
    rand.Read(key)
    rand.Read(plaintext)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        EncryptAESGCM(key, plaintext)
    }
}

func BenchmarkEncryptChaCha20(b *testing.B) {
    key := make([]byte, 32)
    plaintext := make([]byte, 1024)
    rand.Read(key)
    rand.Read(plaintext)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        EncryptChaCha20(key, plaintext)
    }
}
```

**Mobile-Specific Optimizations:**
```go
// Battery-friendly encryption with adaptive chunking
type MobileEncryptor struct {
    chunkSize int // Adjust based on battery level
    algorithm string
}

func (m *MobileEncryptor) EncryptStream(data []byte) <-chan []byte {
    out := make(chan []byte)
    
    go func() {
        defer close(out)
        
        for i := 0; i < len(data); i += m.chunkSize {
            end := i + m.chunkSize
            if end > len(data) {
                end = len(data)
            }
            
            chunk := data[i:end]
            encrypted, _ := m.encryptChunk(chunk)
            out <- encrypted
            
            // Yield CPU to save battery
            runtime.Gosched()
        }
    }()
    
    return out
}

func (m *MobileEncryptor) adjustForBattery(batteryLevel float64) {
    if batteryLevel < 0.2 {
        m.chunkSize = 512 // Smaller chunks for low battery
    } else if batteryLevel < 0.5 {
        m.chunkSize = 1024
    } else {
        m.chunkSize = 4096 // Larger chunks for good battery
    }
}
```

**IoT-Specific Considerations:**
```go
// Memory-constrained encryption for IoT devices
type IoTEncryptor struct {
    keyPool sync.Pool // Reuse key buffers
    nonce   [24]byte  // Fixed nonce to avoid allocations
}

func (i *IoTEncryptor) EncryptMinimal(data []byte) ([]byte, error) {
    // Use fixed nonce with counter to avoid allocations
    binary.BigEndian.PutUint64(i.nonce[:8], atomic.AddUint64(&counter, 1))
    
    aead, _ := chacha20poly1305.NewX(i.keyPool.Get().([]byte))
    defer i.keyPool.Put(aead)
    
    return aead.Seal(nil, i.nonce[:], data, nil), nil
}
```

**When to Choose ChaCha20:**
- **Mobile apps** (iOS/Android)
- **IoT devices** (Raspberry Pi, ESP32)
- **Embedded systems** without AES-NI
- **Cross-platform** consistency
- **Low-power** environments

**When to Choose AES-GCM:**
- **Server-side** with AES-NI
- **High-throughput** APIs
- **Hardware acceleration** available
- **Compatibility** with existing systems
