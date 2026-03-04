# 🚀 08 — Performance & Deployment
> **Most Asked in Service-Based Companies** | 🚀 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Process Managers (PM2)
- Environment Variables (`.env`)
- Clustering and Worker Threads basics
- Managing Memory Leaks
- Caching (`Redis`)

---

## ❓ Frequently Asked Questions

### Q1. What are Environment Variables and why are they used?

**Answer:**
Environment Variables are configuration values passed to the Node.js application from the operating system environment. In Node.js, they are accessed via the `process.env` object.

**Why use them?**
- **Security:** Keep sensitive data like database passwords, API keys, and JWT secrets out of the source code (and out of version control).
- **Flexibility:** Allow the application to run differently depending on the environment (e.g., Development vs. QA vs. Production) without changing the code. For example, changing the `PORT` or the database connection string.

Typically, in development, developers use the `dotenv` package to load variables from a `.env` file into `process.env`.
```javascript
require('dotenv').config();
console.log(process.env.DB_HOST);
```

---

### Q2. What is PM2 and what advantages does it offer?

**Answer:**
PM2 is a Production Process Manager for Node.js applications with a built-in load balancer. 

When you run an app with `node app.js` and the app crashes, it stops completely. PM2 ensures your app stays online.

**Key Advantages:**
- **Auto-Restart:** Automatically restarts the application if it crashes.
- **Background Execution:** Runs the application as a daemon in the background.
- **Log Management:** Aggregates, manages, and rotates logs.
- **Cluster Mode:** Allows you to scale a single Node.js application across all available CPU cores seamlessly.
- **Zero Downtime Reloads:** You can update your application and reload it without dropping any active connections.

**Common Commands:**
```bash
pm2 start app.js    # Start the app
pm2 ls              # List running apps
pm2 logs            # View logs
pm2 stop app        # Stop app
pm2 restart app     # Restart app
```

---

### Q3. Since Node.js is single-threaded, how do you take advantage of multi-core systems?

**Answer:**
A single instance of Node.js runs in a single thread. To take advantage of multi-core systems, you need to launch a **cluster** of Node.js processes.

The **`cluster` module** allows you to easily create child processes (workers) that all share server ports.
The master process listens for incoming connections and distributes them across the worker processes using a round-robin algorithm.

```javascript
const cluster = require('cluster');
const http = require('http');
const numCPUs = require('os').cpus().length;

if (cluster.isMaster) {
    console.log(`Master ${process.pid} is running`);
    // Fork workers for each CPU core.
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork();
    }
    cluster.on('exit', (worker, code, signal) => {
        console.log(`Worker ${worker.process.pid} died. Restarting...`);
        cluster.fork(); 
    });
} else {
    // Workers can share any TCP connection (e.g., an HTTP server)
    http.createServer((req, res) => {
        res.writeHead(200);
        res.end(`Hello from Worker ${process.pid}\n`);
    }).listen(8000);
    console.log(`Worker ${process.pid} started`);
}
```
*(Note: PM2's cluster mode does exactly this under the hood without requiring you to write the boilerplate code).*

---

### Q4. What is the difference between the `cluster` module and `worker_threads`?

**Answer:**
- **`cluster` module:** Spins up completely separate Node.js processes. They don't share memory. They communicate via IPC (Inter-Process Communication). This is primarily used for **scaling web servers** horizontally across CPU cores.
- **`worker_threads` module:** Threads run within the *same* Node.js process and can share memory (using `SharedArrayBuffer`). This is primarily used for **CPU-intensive tasks** (like complex calculations or cryptography) to avoid blocking the main Event Loop.

---

### Q5. How does caching work in Node.js? What is Redis?

**Answer:**
Caching temporarily stores frequently accessed data in memory so future requests for that data can be served much faster than fetching it from the primary database.

**Redis** (Remote Dictionary Server) is an open-source, in-memory data structure store, used as a database, cache, and message broker.

**How it works with Node.js:**
1. A client requests data.
2. The Node.js server checks Redis to see if the data is already cached.
3. **Cache Hit:** If found in Redis, it immediately returns the data to the client (extremely fast).
4. **Cache Miss:** If not found, Node.js queries the primary database (e.g., MongoDB), returns the data to the client, and *saves* a copy to Redis for future requests.

By using caching, you drastically reduce database load and improve response times for read-heavy applications.
