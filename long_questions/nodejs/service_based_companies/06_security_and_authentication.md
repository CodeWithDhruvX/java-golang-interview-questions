# 🔒 06 — Security & Authentication
> **Most Asked in Service-Based Companies** | 🔒 Difficulty: Medium to Hard

---

## 🔑 Must-Know Topics
- JWT (JSON Web Tokens) Basics
- Authentication vs Authorization
- Bcrypt and Password Hashing
- CORS (Cross-Origin Resource Sharing)
- Helmet.js and Security Headers

---

## ❓ Frequently Asked Questions

### Q1. What is the difference between Authentication and Authorization?

**Answer:**
- **Authentication:** Verifying *who* a user is. (e.g., Logging in with a username and password). It checks identity.
- **Authorization:** Verifying *what* a user is allowed to do. (e.g., Can this logged-in user delete a post? Are they an admin?). It checks permissions.

---

### Q2. What is a JWT (JSON Web Token) and how does it work?

**Answer:**
A JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object.

**Structure (Three parts separated by dots `.`):**
1. **Header:** Contains the token type (JWT) and the signing algorithm (e.g., HMAC SHA256 or RSA).
2. **Payload:** Contains the claims (statements about an entity and additional data like user ID or role). *Note: The payload is Base64 encoded, not encrypted. Do not put sensitive data like passwords here!*
3. **Signature:** Created using the encoded header, the encoded payload, a secret, and the specified algorithm. It verifies the token wasn't altered in transit.

**How it works:**
1. User logs in with credentials.
2. Server verifies and generates a JWT, signing it with a secret key.
3. Server sends the JWT to the client.
4. Client stores the JWT (e.g., in HttpOnly cookies or localStorage) and sends it in the `Authorization` header (`Bearer <token>`) for subsequent requests.
5. Server verifies the signature using the secret key to authenticate the request.

---

### Q3. Why should you hash passwords before storing them? How do you do it in Node.js?

**Answer:**
Storing passwords in plain text is a massive security risk. If a database is compromised, attackers would instantly have all user passwords. 

**Hashing** is a one-way mathematical function. You cannot easily reverse a hash to the original password.
We use **Salting** (adding random data to the password before hashing) to defend against dictionary attacks and rainbow tables.

**Using `bcrypt` in Node.js:**
```javascript
const bcrypt = require('bcrypt');
const saltRounds = 10;

// Hashing a password (e.g., during Registration)
const hashPassword = async (plainPassword) => {
    const hash = await bcrypt.hash(plainPassword, saltRounds);
    return hash; 
};

// Verifying a password (e.g., during Login)
const checkPassword = async (plainPassword, storedHash) => {
    const match = await bcrypt.compare(plainPassword, storedHash);
    return match; // true or false
};
```

---

### Q4. What is CORS and how do you enable it in Express?

**Answer:**
**CORS (Cross-Origin Resource Sharing)** is a browser security feature that restricts cross-origin HTTP requests. A web application executing in one origin (domain, protocol, and port) cannot request resources from a different origin unless the server explicitly allows it using specific HTTP headers (`Access-Control-Allow-Origin`).

If CORS is not configured, the browser will block requests made via `fetch` or `XMLHttpRequest` to a different domain.

**Enabling CORS in Express:**
```javascript
const express = require('express');
const cors = require('cors'); // Third-party middleware
const app = express();

// Allow all origins (Default)
app.use(cors());

// Or, allow specific origins only
const corsOptions = {
    origin: 'https://mycoolfrontend.com',
    optionsSuccessStatus: 200
};
app.use(cors(corsOptions));
```

---

### Q5. What is Helmet.js?

**Answer:**
Helmet is a popular middleware package for Express.js that helps secure Node.js applications by setting various HTTP headers. It protects against well-known web vulnerabilities.

By simply adding `app.use(helmet());`, it automatically configures security headers like:
- `Content-Security-Policy` (Mitigates XSS attacks)
- `X-Frame-Options` (Mitigates clickjacking)
- `Strict-Transport-Security` (Enforces secure HTTPS connections)
- `X-Powered-By` (Removes the `X-Powered-By: Express` header to hide the technology stack)
