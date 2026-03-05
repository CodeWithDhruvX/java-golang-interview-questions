# Node.js & Express Fundamentals (Service-Based Companies)

Service-based companies expect a solid understanding of backend API development, routing, middleware, and how Node.js handles requests concurrently despite being single-threaded.

## Core Node.js Concepts

### 1. What is Node.js, and why is it used?
Node.js is a runtime environment that executes JavaScript code outside a web browser, built on Chrome's V8 JavaScript engine.
It is primarily used for building fast, scalable network applications (like REST APIs, real-time chats, and streaming applications) due to its event-driven, non-blocking I/O model.

### 2. How does Node.js handle concurrency if it is single-threaded?
Node.js uses a single-threaded Event Loop. However, it delegates blocking I/O operations (like file system access, database queries, network requests) to the operating system or to a thread pool (libuv) running in the background. Once the operation is complete, a callback is placed into the task queue, which the event loop picks up and executes on the main thread. This non-blocking architecture allows it to handle thousands of concurrent requests without the overhead of context-switching between OS threads.

### 3. What is the role of the `package.json` file?
`package.json` is the manifest file of any Node.js project. It contains:
*   Project metadata (name, version, description, author).
*   **Dependencies**: Packages required to run the application in production (installed via `npm install <pkg>`).
*   **DevDependencies**: Packages needed only during development and testing (e.g., Nodemon, Jest, Typescript), not in production.
*   **Scripts**: Aliases for CLI commands (e.g., `"start": "node index.js"`, `"test": "jest"`).

### 4. What are streams in Node.js?
Streams are collections of data—just like arrays or strings—but they might not be available all at once, and they don't have to fit in memory. They allow you to read or write data chunk by chunk.
*   **Types**: Readable (e.g., `fs.createReadStream`), Writable (e.g., `res` object in HTTP), Duplex, Transform.
*   **Use case**: Processing large files (like uploading a video) without consuming large amounts of RAM.

## Express.js Fundamentals

### 5. What is Express.js? Why use it over the native Node.js `http` module?
Express.js is a minimal and flexible Node.js web application framework.
Using the native `http` module requires manually parsing URLs, extracting query parameters, handling different HTTP methods via boilerplate logic, and manually setting response headers. Express abstract all this by providing robust routing, middleware support, and simplified request/response handling.

### 6. What is Middleware in Express? How does it work?
Middleware functions are functions that have access to the request object (`req`), the response object (`res`), and the `next` function in the application’s request-response cycle.
**Functions of middleware**:
*   Execute any code.
*   Make changes to the request and the response objects.
*   End the request-response cycle.
*   Call the next middleware in the stack (by calling `next()`). If the current middleware doesn't end the cycle, it *must* call `next()`, otherwise the request will be left hanging.
**Common Examples**: `express.json()`, `cors()`, authentication checkers, error handlers.

### 7. How do you handle routing in Express?
Routing refers to determining how an application responds to a client request to a particular endpoint (URI) and a specific HTTP method (GET, POST, PUT, DELETE).
```javascript
const express = require('express');
const app = express();

app.get('/users', (req, res) => {
    res.send('Get all users');
});

app.post('/users', (req, res) => {
    res.send('Create a user');
});
```
For modularity, Express provides `express.Router()` to create mountable, modular route handlers.

### 8. Explain how to handle errors globally in Express.
Global error handling is done using a special middleware function with **four arguments**: `(err, req, res, next)`. This middleware must be placed *after* all other routes and middleware.

```javascript
// A normal route
app.get('/error-prone', (req, res, next) => {
    const error = new Error('Something went wrong!');
    error.status = 500;
    next(error); // Pass the error to the error handler
});

// Global Error Handler
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(err.status || 500).json({
        message: err.message,
        error: process.env.NODE_ENV === 'development' ? err : {}
    });
});
```

### 9. What are Route Parameters and Query Parameters? How do you access them?
*   **Route Parameters**: Named URL segments used to capture values dynamically.
    *   Route: `/users/:id`  -> Access via: `req.params.id`
*   **Query Parameters**: Key-value pairs appended after the `?` in the URL, usually for sorting, filtering, or pagination.
    *   URL: `/users?role=admin&page=2` -> Access via: `req.query.role` and `req.query.page`
