# Security & Authentication Interview Questions (101-110)

## Security Fundamentals

### 101. What is authentication vs authorization?
"**Authentication (AuthN)** is about identity: 'Who are you?' It’s proving you are who you say you are—usually via username/password, OTP, or fingerprint.

**Authorization (AuthZ)** is about permissions: 'What can you do?' Once we know who you are, do you have access to delete this user? Or view this report?

I like the analogy: Authentication is your ID card at the building entrance. Authorization is the key card that only opens specific doors inside."

### 102. How does JWT work internally?
"A JWT (JSON Web Token) is just a base64 encoded string with three parts separated by dots.

1.  **Header**: Describes the algorithm (e.g., HS256).
2.  **Payload**: The data claims (User ID, Role, Expiration time).
3.  **Signature**: The cryptographic proof. It’s a hash of (Header + Payload + Secret Key).

When the server receives a JWT, it recalculates the signature using its secret key. If it matches the signature in the token, we know the data hasn't been tampered with. It’s stateless, which makes it great for scaling."

### 103. Where should JWT be stored on the client and why?
"This is a huge debate.

**LocalStorage**: Easy to implement, but vulnerable to XSS (Cross-Site Scripting). If malicious JS runs on your page, it can steal the token.

**HttpOnly Cookie**: Safer. The cookie is automatically sent with requests, but JavaScript cannot read it, so XSS attacks can't steal it. However, this makes you vulnerable to CSRF (Cross-Site Request Forgery), so you need CSRF protection.

For high-security apps (like banking), I strictly use **HttpOnly Cookies** with `SameSite=Strict`. For lesser risks, LocalStorage is... acceptable, but not ideal."

### 104. What are common security vulnerabilities in REST APIs?
"The OWASP Top 10 lists the big ones.

1.  **Broken Object Level Authorization (BOLA/IDOR)**: I change the ID in the URL (`/users/5` -> `/users/6`) and view data I shouldn't see.
2.  **Broken Authentication**: Weak passwords, no rate limiting on login (brute force), allowing weak JWT secrets.
3.  **Injection**: SQL Injection or Command Injection.
4.  **Mass Assignment**: Sending a JSON with `{"role": "admin"}` and the API blindly updating the user object to be an admin."

### 105. What is CORS and how does it work?
"CORS (Cross-Origin Resource Sharing) is a browser security feature. By default, a browser running a site on `domain-a.com` prevents it from making API calls to `domain-b.com`. This is the Same-Origin Policy.

When you legitimately need to call an external API, the browser sends a pre-flight `OPTIONS` request. The server must respond with specific headers like `Access-Control-Allow-Origin: domain-a.com`.

If the headers match, the browser allows the actual request. If not, you get that red CORS error in the console."

### 106. Difference between OAuth2 and JWT?
"They are apples and oranges.

**JWT** is a token format. It defines *how* data is packaged and signed.

**OAuth2** is an authorization framework/protocol. It defines the *process* of how a user grants a third-party app access to their data (like 'Log in with Google').

OAuth2 often *uses* JWTs as the Access Token format, but it doesn't have to. You can use OAuth2 with random string tokens too."

### 107. How does Spring Security filter chain work?
"Spring Security is essentially a chain of Servlet Filters.

When a request comes in, it passes through this chain *before* reaching your DispatcherServlet (Controller).
There are filters for:
-   Checking headers (Basic Auth, Bearer Token).
-   Handling CORS.
-   CSRF protection.
-   Exception translation (turning 403s into redirects to login page).

If any filter rejects the request (throws AuthenticationException), the request stops there and never reaches your business logic."

### 108. What is CSRF and how do you prevent it?
"CSRF (Cross-Site Request Forgery) is when a malicious site (attacker.com) tricks your browser into making a request to a site where you are logged in (bank.com). Because your browser automatically sends cookies, bank.com thinks *you* made the request.

We prevent it using **CSRF Tokens**. The server generates a random token and embeds it in the HTML form. When you submit, the server checks if the token matches. Since attacker.com can't read the token from your bank page (due to Same-Origin Policy), they can't forge the request."

### 109. How do you secure internal microservice communication?
"We don't just rely on 'being inside the VPC/Firewall'. That’s zero trust.

1.  **mTLS (Mutual TLS)**: Both services present certificates to prove their identity and encrypt traffic. Usually handled by a Service Mesh like Istio.
2.  **JWT Propagation**: Passing the user's JWT from the Gateway down to Service A -> Service B, so Service B knows *who* triggered the action.
3.  **Service-to-Service Auth**: Using `Client Credentials Flow` (OAuth2) where Service A gets its own token to call Service B."

### 110. What is password hashing and salting?
"You never store plain text passwords. You store a hash.

**Hashing** transforms 'password123' into a fixed string 'a1b2c3...'. It’s one-way; you can't reverse it.

**Salting** adds a random string to the password *before* hashing (`hash("password123" + "random_salt")`). This prevents **Rainbow Table attacks**, where attackers have pre-computed hashes for common passwords. Even if two users have the same password, they will have different salts, so their stored hashes will look completely different.

I use **BCrypt** or **Argon2** because they are intentionally slow, making brute-force attacks expensive."
