# Web Security Deep Dive (Product-Based Companies)

Product companies handle massive amounts of user data, making security a top priority. You must understand common attack vectors and how to mitigate them in a Node.js/Frontend environment.

## Cross-Site Scripting (XSS)

### 1. What is XSS and how do you prevent it in a MEAN/MERN app?
**XSS** occurs when an attacker injects malicious client-side JavaScript into a web page viewed by other users. If successful, the script can steal session cookies, tokens, or perform actions on behalf of the user.
*   **Stored XSS**: Malicious script is saved in the DB (e.g., in a blog comment) and executed when users view the page.
*   **Reflected XSS**: Malicious script is embedded in a link/URL query parameter and executed when the victim clicks the link.
**Prevention**:
*   **React/Angular do this automatically**: By default, data binding (`{data}` in React, `{{data}}` in Angular) escapes HTML entities. `<script>` payload becomes safe text.
*   **Never use Dangerous APIs blindly**: Avoid `dangerouslySetInnerHTML` in React or bypassing security trust in Angular unless strictly necessary and after running the input through a sanitization library like `DOMPurify`.
*   **Input Sanitization**: Use libraries like `validator.js` or `xss-clean` on the Node.js backend to scrub incoming data *before* saving it to MongoDB.

## Cross-Site Request Forgery (CSRF)

### 2. What is CSRF and how do you mitigate it?
**CSRF** tricks an authenticated user into executing unwanted actions on a web application in which they are currently authenticated. Imagine user A is logged into their bank. Attacker sends user A an email with an invisible `<img src="http://bank.com/transfer?amount=1000&to=attacker">`. The browser automatically includes the user's session cookies in the request, and the bank transfers the money.
**Prevention:**
*   **Anti-CSRF Tokens (Synchronizer Token Pattern)**: The server generates a unique, random token for the session and sends it to the client. The client must include this token in hidden form fields or custom HTTP headers for any state-changing requests (POST, PUT, DELETE). The attacker cannot guess this token.
*   **SameSite Cookie Attribute**: Set `SameSite=Strict` or `SameSite=Lax` on session cookies. This tells the browser *not* to send the cookie along with cross-site requests (like the image tag in the attacker's email).
*   **If using JWTs in LocalStorage**: CSRF is not naturally possible because the attacker's site cannot access the victim's LocalStorage to attach the `Authorization: Bearer <token>` header. (However, LocalStorage is vulnerable to XSS).

## Backend Security

### 3. What is NoSQL Injection and how do you prevent it in MongoDB?
Similar to SQL injection, NoSQL injection occurs when user input is used to construct a database query without sanitization, allowing the attacker to alter the query logic.
**Example Attack**:
An attacker passes `{"$gt": ""}` as the password field during login.
`db.users.find({ username: "admin", password: { $gt: "" } })`. Since all strings are "greater than" an empty string, this query evaluates to true, and the attacker logs in without the password.
**Prevention**:
*   **Use Mongoose**: Mongoose schemes cast data strictly. If a field is defined as a `String`, passing an object `{$gt: ""}` will throw a CastError before ever hitting the database.
*   **Sanitize input**: Use middleware like `express-mongo-sanitize` which recursively removes any keys starting with `$` or `.` from `req.body`, `req.query`, and `req.params`.

### 4. What is `helmet.js` and why should you use it in Express?
Helmet is a collection of 15 smaller middleware functions that set secure HTTP response headers. By default, Express leaks information (like the `X-Powered-By: Express` header) and lacks basic security headers.
**Key headers set by Helmet:**
*   `Content-Security-Policy (CSP)`: Highly effective against XSS. Prevents loading scripts, images, or iframes from unauthorized domains.
*   `Strict-Transport-Security (HSTS)`: Forces browsers to communicate with your server exclusively over HTTPS.
*   `X-Frame-Options`: Prevents Clickjacking by disabling your site from being rendered inside an `<iframe>` on an attacker's site.

### 5. How do you securely hash passwords in Node.js?
**Never store plain text passwords.**
*   Use a strong hashing algorithm like **Bcrypt** or **Argon2**.
*   These algorithms include a **Salt** (random string appended to the password before hashing) to defeat Rainbow Table attacks (pre-computed hash lists).
*   They are intentionally **slow** (variable computational "work factor" or "rounds"). This makes brute-force attacks computationally unfeasible for hackers.
```javascript
const bcrypt = require('bcrypt');
const saltRounds = 10;
const hashedPassword = await bcrypt.hash(plainTextPassword, saltRounds);
// Later, to compare during login:
const match = await bcrypt.compare(plainTextPassword, hashedPassword);
```
