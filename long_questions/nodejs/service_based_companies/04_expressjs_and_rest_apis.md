# 🌐 04 — Express.js & REST APIs
> **Most Asked in Service-Based Companies** | 🌐 Difficulty: Easy to Medium

---

## 🔑 Must-Know Topics
- Creating a basic Express application
- Routing (`GET`, `POST`, `PUT`, `DELETE`)
- Middleware functions and their types
- Error handling middleware
- Handling Request Data (`req.body`, `req.params`, `req.query`)
- REST API principles
- Templating engines (basics)

---

## ❓ Frequently Asked Questions

### Q1. What is Express.js? Why is it practically ubiquitous with Node.js?

**Answer:**
Express.js is a fast, unopinionated, minimalist web framework for Node.js. While Node's core `http` module can create servers, it requires a lot of boilerplate code to handle routing, serving static files, and parsing request bodies. Express simplifies all of this by providing:

1. **Routing:** A robust and easy-to-use routing mechanism.
2. **Middleware:** A flexible pipeline to process requests before they hit the route handler.
3. **Simplicity:** Greatly reduces boilerplate code.
4. **Integration:** Easy integration with template engines (EJS, Pug) and databases.

---

### Q2. Explain Routing in Express.js.

**Answer:**
Routing refers to how an application responds to a client request to a particular endpoint, which is a URI (or path) and a specific HTTP request method (GET, POST, etc.).

```javascript
const express = require('express');
const app = express();

// GET request
app.get('/', (req, res) => {
    res.send('Home Page');
});

// POST request
app.post('/api/users', (req, res) => {
    res.status(201).json({ message: 'User created' });
});

app.listen(3000, () => console.log('Server running on port 3000'));
```

**Route Parameters vs Query Parameters:**
- **Route Params (`req.params`):** Used to identify a specific resource. `app.get('/users/:id', ...)`
- **Query Params (`req.query`):** Used to sort/filter resources. `/users?sort=asc`

---

### Q3. What is Middleware in Express.js? Give an example.

**Answer:**
Middleware functions are functions that have access to the request object (`req`), the response object (`res`), and the `next` function in the application’s request-response cycle.

**Middleware can:**
- Execute any code.
- Make changes to the request and the response objects.
- End the request-response cycle.
- Call the next middleware in the stack (`next()`).

If the current middleware does not end the request-response cycle, it *must* call `next()` to pass control to the next middleware function. Otherwise, the request will be left hanging.

**Example:**
```javascript
const express = require('express');
const app = express();

// Application level middleware
const logger = (req, res, next) => {
    console.log(`${req.method} request to ${req.url}`);
    next(); // Pass control to the next function
};

app.use(logger);

app.get('/about', (req, res) => {
    res.send('About Us');
});
```

---

### Q4. What are the different types of middleware in Express?

**Answer:**
1. **Application-level middleware:** Bound to an instance of the `express` object using `app.use()`. Runs for every request that hits the app (or a specific path).
2. **Router-level middleware:** Works the same as application-level middleware, but it is bound to an instance of `express.Router()`.
3. **Error-handling middleware:** Always takes four arguments: `(err, req, res, next)`.
4. **Built-in middleware:** Provided by Express, like `express.json()`, `express.urlencoded()`, and `express.static()`.
5. **Third-party middleware:** e.g., `helmet` (security), `cors`, `morgan` (logging).

---

### Q5. How do you parse incoming request bodies in Express?

**Answer:**
By default, Express does not parse the incoming request body. You must use built-in middleware to process the incoming payload.

1. **For JSON payloads (`application/json`):**
   ```javascript
   app.use(express.json());
   ```
2. **For URL-encoded payloads (form submissions):**
   ```javascript
   app.use(express.urlencoded({ extended: true }));
   ```

```javascript
app.post('/login', (req, res) => {
    // Without express.json(), req.body would be undefined
    const { username, password } = req.body;
    res.send(`Logging in ${username}`);
});
```

---

### Q6. How do you implement Error Handling in Express?

**Answer:**
Express comes with a default error handler, but you should create a custom one to format error responses consistently. Error-handling middleware is defined just like regular middleware, except it takes *four* arguments instead of three: `(err, req, res, next)`.

**It must be placed at the very end of all `app.use()` and route calls.**

```javascript
const express = require('express');
const app = express();

app.get('/', (req, res) => {
    throw new Error('Something broke!'); // Synchronous errors are caught automatically
});

app.get('/async', async (req, res, next) => {
    try {
        const data = await someAsyncFunction(); // If this throws, you MUST catch it
    } catch (err) {
        next(err); // Pass async errors to the error handler manually
    }
});

// Custom Error Handler (MUST BE LAST)
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.status(500).json({ error: 'Internal Server Error', message: err.message });
});
```

*(Note: Starting in Express 5.0, route handlers that return a Promise will automatically call `next(err)` when they reject. But for Express 4.x, you must use `try/catch` and `next(err)`).*
