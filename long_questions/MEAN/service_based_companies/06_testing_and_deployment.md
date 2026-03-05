# Testing and Deployment Basics (Service-Based Companies)

Service-based companies want to know that you can write code that works, prove it works via tests, and deploy it to a standard hosting environment.

## Testing Fundamentals

### 1. What are the different types of testing in a typical MEAN/MERN application?
*   **Unit Testing**: Testing individual components or functions in isolation (e.g., testing that a React button renders correctly, or a Node.js utility function returns the right value). Tools: Jest, Jasmine, Mocha.
*   **Integration Testing**: Testing how different parts of the system work together (e.g., testing an Express route that connects to a test MongoDB database to fetch data). Tools: Supertest + Jest/Mocha.
*   **End-to-End (E2E) Testing**: Simulating a real user clicking through the application in a headless browser to ensure the entire system (frontend + backend + DB) works as expected. Tools: Cypress, Playwright, Selenium.

### 2. How do you test an Express API route? (Integration Testing)
You use a library like `supertest` along with a test runner like `jest`. `supertest` allows you to make HTTP requests to your Express app without actually starting the server on a port.
```javascript
const request = require('supertest');
const app = require('../app'); // Your express app

describe('GET /users', () => {
    it('should return a list of users with 200 status', async () => {
        const response = await request(app).get('/users');
        expect(response.status).toBe(200);
        expect(Array.isArray(response.body)).toBeTruthy();
    });
});
```

### 3. What is Mocking, and why is it important in Unit Testing?
Mocking is replacing real objects or functions with fake versions (mocks or stubs).
*   **Why it's important**: When testing a component that fetches data from an API, you don't want the test to actually hit the real API over the network (it makes tests slow and flaky if the network goes down). Instead, you *mock* the fetch call to return a hardcoded JSON response immediately.

## Deployment Basics

### 4. What is PM2 and why is it used for Node.js applications?
Node.js runs on a single thread. If an uncaught exception occurs, the Node.js process crashes, and your website goes offline.
**PM2** is a production process manager for Node.js.
*   It automatically restarts the application if it crashes.
*   It allows you to run multiple instances of your app (Clustering) to utilize all CPU cores.
*   It runs your app in the background (daemon mode) so it stays alive after you log out of the server SSH session.

### 5. Explain the role of NGINX in front of a Node.js server.
NGINX is typically used as a **Reverse Proxy** in front of Node.js.
*   **Security/Port Binding**: Node.js shouldn't run directly on port 80 or 443 as root. NGINX listens on ports 80/443 and forwards traffic to Node running on a higher port (like 3000).
*   **Static Asset Serving**: NGINX is infinitely faster at serving static files (images, CSS, the React/Angular built frontend) than Node.js is.
*   **SSL Termination**: NGINX handles the SSL certificates (HTTPS encryption/decryption) so Node doesn't have to waste CPU cycles doing it.
*   **Load Balancing**: Distributing incoming traffic across multiple PM2 instances or multiple servers.

### 6. How do you deploy a React or Angular application to production?
Frontend apps are just static files (HTML, JS, CSS) once built. You don't need a Node.js server to run them in production.
1.  **Build**: Run `npm run build` (or `ng build`). This minifies and bundles the code.
2.  **Host**: Take the generated `build/` or `dist/` folder and place it on a static file hosting service.
    *   Examples: AWS S3 combined with CloudFront, Netlify, Vercel, or a simple NGINX directory on a VPS.
3.  **Routing configuration**: Because they are SPAs, direct links to nested routes (like `/users/profile`) will throw a 404 from the hosting server. You must configure NGINX or the hosting provider to redirect all 404 traffic back to `index.html` so the JS router can take over.
