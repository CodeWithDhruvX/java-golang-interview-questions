# Full-Stack Integration (Service-Based Companies)

Service-based companies often ask integration questions to ensure you understand how the frontend, backend, and database communicate, handle security, and manage user sessions.

## Connecting Frontend to Backend

### 1. What is CORS and why is it necessary?
CORS (Cross-Origin Resource Sharing) is a security mechanism built into modern web browsers. It restricts web pages from making requests to a different domain than the one that served the web page.
*   **Why necessary**: It prevents malicious scripts from making unauthorized requests on behalf of a user to another site where they might be authenticated.
*   **How to handle it**: The backend server must explicitly allow requests from the frontend origin by setting specific HTTP headers (e.g., `Access-Control-Allow-Origin: *` or a specific URL).
*   **In Express**: Use the `cors` middleware: `app.use(cors({ origin: 'http://localhost:4200' }));`

### 2. Explain the typical request-response cycle in a MEAN/MERN application.
1.  **User Action**: The user interacts with the UI (e.g., clicks a "Submit" button on a React/Angular form).
2.  **Frontend Request**: An HTTP request (e.g., POST with JSON payload) is initiated using `fetch`, `axios`, or Angular's `HttpClient` to the Node.js backend URL.
3.  **Backend Receiving**: The Express server receives the request at the corresponding route (e.g., `app.post('/api/users')`).
4.  **Middleware Execution**: The request passes through middleware (parsing JSON, checking CORS, authentication).
5.  **Controller Logic**: The route handler processes the data and interacts with Mongoose models.
6.  **Database Interaction**: Mongoose translates the Node.js commands into MongoDB operations (e.g., `User.create()`).
7.  **Database Response**: MongoDB returns the result (success or error) back to Mongoose/Node.js.
8.  **Backend Response**: Express formulates an HTTP response (status code, JSON data/error message) and sends it back to the frontend.
9.  **Frontend Handling**: The frontend receives the response (handled within a `.then()` or `await`) and updates the UI accordingly (e.g., showing a success toast or redirecting).

## Authentication & Authorization

### 3. What is JWT (JSON Web Token)? How does it work?
JWT is an open standard that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed.
*   **Structure**: Header (algorithm), Payload (data like user ID or roles), Signature (created using a secret key on the server).
*   **Workflow**:
    1.  User logs in with credentials.
    2.  Server verifies credentials and creates a JWT signed with a secret key.
    3.  Server sends the JWT to the client.
    4.  Client stores the JWT (usually in `localStorage` or a secure `HttpOnly` cookie).
    5.  For subsequent requests, the client sends the JWT in the `Authorization` header (`Bearer <token>`).
    6.  The server validates the signature. If valid, it grants access to protected routes.

### 4. What is the difference between Authentication and Authorization?
*   **Authentication**: Verifying *who* a user is (e.g., logging in with email and password).
*   **Authorization**: Verifying *what* a user is allowed to do (e.g., an Admin can delete users, a regular User can only view them).

### 5. Where should you store JWTs on the client side? What are the risks?
*   **`localStorage` / `sessionStorage`**:
    *   *Pros*: Easy to access via JavaScript, persists across page reloads.
    *   *Cons*: Vulnerable to **XSS (Cross-Site Scripting)** attacks. If an attacker injects a malicious script, they can easily read `localStorage` and steal the token.
*   **`HttpOnly` Cookies**:
    *   *Pros*: The browser prevents client-side JavaScript from accessing the cookie, mitigating XSS risks.
    *   *Cons*: Vulnerable to **CSRF (Cross-Site Request Forgery)** attacks unless mitigated with anti-CSRF tokens or `SameSite` cookie attributes. This is generally the recommended approach for high-security applications.

## General Integration Concepts

### 6. How do you handle environment variables in a Node.js project?
Environment variables store sensitive configuration data (database URIs, secret keys, port numbers) outside of the codebase to prevent hardcoding them in version control.
*   In Node.js, they are accessed via `process.env`.
*   A common tool is the `dotenv` package (`require('dotenv').config()`) which loads variables from a `.env` file into `process.env` during development.

### 7. What is package-lock.json (or yarn.lock)?
These lock files are automatically generated when you install node modules. They record the exact version of every installed dependency and sub-dependency.
*   **Purpose**: Ensures that the project has the exact same dependency tree installed across different environments (development, staging, production) and among team members, preventing "it works on my machine" issues caused by minor version updates of nested packages.
