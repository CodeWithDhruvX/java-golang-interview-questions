# Security & Authentication Interview Questions (101-110)

## Security Fundamentals

### 101. What is authentication vs authorization?
"**Authentication (AuthN)** is about identity: 'Who are you?' It’s proving you are who you say you are—usually via username/password, OTP, or fingerprint.

**Authorization (AuthZ)** is about permissions: 'What can you do?' Once we know who you are, do you have access to delete this user? Or view this report?

I like to analogy: Authentication is your ID card at the building entrance. Authorization is the key card that only opens specific doors inside."

**Spoken Format:**
"Authentication and Authorization are like the two-step security process at a secure building.

**Authentication** is like showing your ID card at the front door. The security guard checks your photo, verifies your identity, and decides if you are who you claim to be. It answers 'Who are you?'

**Authorization** is like the key card system inside the building. Once you're inside, different doors require different keys. Your ID card might get you through the front door, but you need a special key to access the server room, executive floor, or safe.

The key difference: Authentication happens once at the entrance, authorization happens continuously as you move through the building.

In web apps: Authentication is logging in with username/password. Authorization is the permissions check that happens for every API call - 'Can this user delete posts? Can they access admin features?'

Both are essential - authentication proves identity, authorization proves permissions!"

### 102. How does JWT work internally?
"A JWT (JSON Web Token) is just a base64 encoded string with three parts separated by dots.

1.  **Header**: Describes the algorithm (e.g., HS256).
2.  **Payload**: The data claims (User ID, Role, Expiration time).
3.  **Signature**: The cryptographic proof. It’s a hash of (Header + Payload + Secret Key).

When the server receives a JWT, it recalculates the signature using its secret key. If it matches the signature in the token, we know the data hasn't been tampered with. It’s stateless, which makes it great for scaling."

**Spoken Format:**
"JWT is like a digital passport that proves your identity.

Imagine you're traveling to a foreign country. You need a passport that proves your identity and citizenship. JWT works similarly:

1. **Header**: The passport cover page with your name and photo.
2. **Payload**: The visa page with your travel details and permissions.
3. **Signature**: The official stamp that proves the passport is genuine.

When you arrive at the destination, the immigration officer checks your passport by recalculating the signature. If it matches, you're allowed to enter. If not, you're denied entry.

JWT is stateless, meaning the server doesn't store any information about you. It's like a self-contained passport that proves your identity and permissions."

### 103. Where should JWT be stored on the client and why?
"This is a huge debate.

**LocalStorage**: Easy to implement, but vulnerable to XSS (Cross-Site Scripting). If malicious JS runs on your page, it can steal the token.

**HttpOnly Cookie**: Safer. The cookie is automatically sent with requests, but JavaScript cannot read it, so XSS attacks can't steal it. However, this makes you vulnerable to CSRF (Cross-Site Request Forgery), so you need CSRF protection.

For high-security apps (like banking), I strictly use **HttpOnly Cookies** with `SameSite=Strict`. For lesser risks, LocalStorage is... acceptable, but not ideal."

**Spoken Format:**
"Storing JWT is like deciding where to keep your house keys - each choice has different security implications.

**LocalStorage** is like hiding your keys under the doormat - convenient but risky. If malicious JavaScript runs on your website, it can reach under the doormat and steal your JWT token.

**HttpOnly Cookie** is like keeping your keys in a special locked box that only your website can access. JavaScript can't reach into this box, making it much safer.

For high-security apps like banking, I always use HttpOnly cookies with SameSite=Strict - it's like having maximum security for your keys.

The tradeoff is convenience vs. security. LocalStorage is easier to implement but vulnerable to XSS. HttpOnly cookies are slightly more complex but much more secure.

Choose based on your security requirements - for banking apps, always choose the most secure option!"

### 104. What are common security vulnerabilities in REST APIs?
"The OWASP Top 10 lists the big ones.

1.  **Broken Object Level Authorization (BOLA/IDOR)**: I change the ID in the URL (`/users/5` -> `/users/6`) and view data I shouldn't see.
2.  **Broken Authentication**: Weak passwords, no rate limiting on login (brute force), allowing weak JWT secrets.
3.  **Injection**: SQL Injection or Command Injection.
4.  **Mass Assignment**: Sending a JSON with `{"role": "admin"}` and the API blindly updating the user object to be an admin."

**Spoken Format:**
"REST API vulnerabilities are like leaving doors unlocked in your house - attackers can walk right in if they know where to look.

The OWASP Top 10 are the most common mistakes developers make:

**Broken Object Level Authorization** is like having a filing system where anyone can change any folder number. If I'm user 123 but can access `/users/456` data by just changing the URL, that's a huge security hole.

**Broken Authentication** is like having a password policy that allows '123456' - attackers can just guess common passwords.

**Injection attacks** are like leaving a suggestion box where attackers can write their own commands. If you build SQL queries by concatenating user input, attackers can make your database run their code.

**Mass Assignment** is like giving every new employee an admin badge by default - one malicious user can take over the whole system.

The key is to never trust user input - always validate, sanitize, and check permissions before taking any action!"

### 105. What is CORS and how does it work?
"CORS (Cross-Origin Resource Sharing) is a browser security feature. By default, a browser running a site on `domain-a.com` prevents it from making API calls to `domain-b.com`. This is the Same-Origin Policy.

When you legitimately need to call an external API, the browser sends a pre-flight `OPTIONS` request. The server must respond with specific headers like `Access-Control-Allow-Origin: domain-a.com`."

**Spoken Format:**
"CORS is like having a bouncer at each door who only lets people from certain buildings enter.

Imagine your website is on `domain-a.com` and you want to call an API on `domain-b.com`. By default, browsers block this - it's a security feature called Same-Origin Policy.

CORS is the system that lets you override this policy:

1. Your browser sends a pre-flight request asking 'Can I make a call to domain-b.com?'
2. The server responds with headers saying 'Yes, domain-a.com is allowed'
3. Only then does the actual API call happen

If the server doesn't respond properly, you get the famous CORS error in your console.

It's like having to get permission from multiple bouncers before you can visit another building - annoying for developers but essential for security!

Modern frameworks handle most of this automatically, but understanding CORS helps debug cross-origin issues quickly."

### 106. Difference between OAuth2 and JWT?
"They are apples and oranges.

**JWT** is a token format. It defines *how* data is packaged and signed.

**OAuth2** is an authorization framework/protocol. It defines the *process* of how a user grants a third-party app access to their data (like 'Log in with Google').

OAuth2 often *uses* JWTs as the Access Token format, but it doesn't have to. You can use OAuth2 with random string tokens too."

**Spoken Format:**
"OAuth2 and JWT are like two different types of VIP passes.

**JWT** is like a physical VIP pass - it contains information about who you are (your identity, permissions, expiration). It's a standardized format that many systems can understand.

**OAuth2** is like the entire VIP system - it defines the process of how you get the pass, where you can use it, and what happens when you don't need it anymore.

The key difference: JWT is the format of the pass itself, OAuth2 is the system that issues and manages the passes.

OAuth2 often uses JWT as the pass format, but it could use any format - it's the protocol that matters.

Think of it this way: JWT is what your VIP pass looks like, OAuth2 is the whole VIP management system that creates and validates those passes."

### 107. How does Spring Security filter chain work?
"Spring Security is essentially a chain of Servlet Filters.

When a request comes in, it passes through this chain *before* reaching your DispatcherServlet (Controller).
There are filters for:
-   Checking headers (Basic Auth, Bearer Token).
-   Handling CORS.
-   CSRF protection.
-   Exception translation (turning 403s into redirects to login page).

If any filter rejects the request (throws AuthenticationException), the request stops there and never reaches your business logic."

**Spoken Format:**
"Spring Security filters are like having multiple security checkpoints that every request must pass through.

Imagine a request entering your application - it has to go through a series of security guards before reaching the main area.

The filter chain works like this:

1. **Header Check** - First guard checks if you have a ticket (authentication token)
2. **CORS Check** - Second guard verifies you're coming from an allowed location
3. **CSRF Check** - Third guard ensures you're not being tricked into making requests
4. **Exception Handling** - If any guard rejects you, you're redirected to the entrance (error page)

If any filter fails, the request never reaches your business logic - it's like being stopped at security checkpoint and never getting to the main event.

This layered approach ensures that every request is properly validated before your application processes it. It's like having multiple layers of security - if one fails, the whole system stays secure!"

### 108. What is CSRF and how do you prevent it?
"CSRF (Cross-Site Request Forgery) is when a malicious site (attacker.com) tricks your browser into making a request to a site where you are logged in (bank.com). Because your browser automatically sends cookies, bank.com thinks *you* made the request.

We prevent it using **CSRF Tokens**. The server generates a random token and embeds it in the HTML form. When you submit, the server checks if the token matches. Since attacker.com can't read the token from your bank page (due to Same-Origin Policy), they can't forge the request."

**Spoken Format:**
"CSRF protection is like having a secret handshake that proves you really intended to make a request.

Imagine you're logged into your bank, and an attacker tricks your browser into sending a transfer request to the bank. The browser automatically includes your login cookie, so the bank thinks it's you making the request.

**CSRF tokens prevent this** by:
1. When the bank page loads, it generates a random secret token
2. This token is embedded in the transfer form
3. When you submit the transfer, the bank checks if the token matches
4. Since the attacker can't read the token from their malicious site, they can't forge the request

It's like having a secret password that changes every time - only someone who saw the legitimate page knows the current password.

Modern frameworks handle this automatically, but understanding CSRF helps you appreciate why those hidden form fields are so important for security!"

### 109. How do you secure internal microservice communication?
"We don't just rely on 'being inside the VPC/Firewall'. That’s zero trust.

1.  **mTLS (Mutual TLS)**: Both services present certificates to prove their identity and encrypt traffic. Usually handled by a Service Mesh like Istio.
2.  **JWT Propagation**: Passing the user's JWT from the Gateway down to Service A -> Service B, so Service B knows *who* triggered the action.
3.  **Service-to-Service Auth**: Using `Client Credentials Flow` (OAuth2) where Service A gets its own token to call Service B."

**Spoken Format:**
"Securing microservice communication is like having secure diplomatic channels between different departments.

Instead of assuming internal networks are safe, we treat them like potentially hostile territories:

**mTLS (Mutual TLS)** is like departments having to show ID badges to each other before sharing sensitive information. Both services verify each other's certificates and encrypt all traffic. It's like having a diplomatic protocol.

**JWT Propagation** is like carrying your passport with you as you visit different departments. When Service A calls Service B, it includes your JWT so Service B knows who made the request.

**Service-to-Service Auth** is like departments having their own special relationship. Service A gets its own special access token to call Service B, without involving users at all.

The key insight: Zero trust is dangerous. Even internal services should authenticate and authorize each other properly. It's like having security checkpoints between every department, not just at the building entrance!"

### 110. What is password hashing and salting?
"You never store plain text passwords. You store a hash.

**Hashing** transforms 'password123' into a fixed string 'a1b2c3...'. It’s one-way; you can't reverse it.

**Salting** adds a random string to the password *before* hashing (`hash("password123" + "random_salt")`). This prevents **Rainbow Table attacks**, where attackers have pre-computed hashes for common passwords. Even if two users have the same password, they will have different salts, so their stored hashes will look completely different.

I use **BCrypt** or **Argon2** because they are intentionally slow, making brute-force attacks expensive."

**Spoken Format:**
"Password hashing is like using a one-way blender for your passwords.

**Plain text passwords** are like writing your password on a sticky note - anyone can read it.

**Hashing** is like putting the sticky note in the blender. Once blended, you can't un-blend it back to the original password. The result is always the same hash for the same password.

**Salting** is like adding a secret ingredient before blending. Even if two people have the same password, their hashes will be completely different because they used different secret ingredients.

**BCrypt/Argon2** are like using a very slow, expensive blender. This makes it impractical for attackers to try millions of password combinations - each attempt takes significant time and computer power.

The beauty is: even if attackers steal your hash database, they can't reverse it to get the original passwords. They can only try to guess passwords and check if the hashes match.

Modern hashing algorithms are designed to be slow on purpose - security through computational expense!""
