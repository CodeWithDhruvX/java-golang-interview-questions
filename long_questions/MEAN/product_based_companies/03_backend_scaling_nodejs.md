# Backend Scaling and Architecture (Node.js/Express)

Product-based backend interviews focus on your ability to scale a Node.js application from a single instance to a globally distributed system handling millions of requests, managing state, background jobs, and optimizing performance.

## Scaling Node.js

### 1. Node.js is single-threaded. How do you scale it to utilize a multi-core CPU?
We use the **Cluster module** or process managers like **PM2**.
*   The `cluster` module allows the creation of child processes (workers) that run simultaneously and share the same server port.
*   The master process acts as a load balancer, distributing incoming connections to the workers using a Round-Robin algorithm (by default).
*   Typically, you spawn one worker per CPU core to maximize utilization without incurring excessive context-switching overhead.
```javascript
const cluster = require('cluster');
const os = require('os');
const express = require('express');

if (cluster.isPrimary) {
    const numCPUs = os.cpus().length;
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork();
    }
} else {
    // Worker processes have a http server.
    const app = express();
    app.listen(3000);
}
```

### 2. How do Worker Threads differ from the Cluster module? When would you use them?
*   **Cluster/Child Processes**: Spawns entirely separate V8 instances and Node.js event loops. Processes communicate via IPC (Inter-Process Communication), which is relatively slow. Good for **Scaling I/O-intensive web servers** horizontally across cores.
*   **Worker Threads (`worker_threads` module)**: Threads run within a single Node.js process and share the same memory space (using `SharedArrayBuffer` for very fast communication).
*   **Use Case for Worker Threads**: **CPU-intensive tasks** (image processing, video encoding, complex cryptography, heavy math). Running these on the main event loop would block it, halting all other requests. Offloading them to a worker thread keeps the main thread free for I/O.

## High Availability and Architecture

### 3. What is the API Gateway pattern in a microservices architecture?
An API Gateway acts as a single entry point for all clients into the microservices ecosystem.
**Responsibilities:**
*   **Request Routing**: Routing traffic to the correct underlying microservice.
*   **Composition/Aggregation**: Calling multiple microservices and aggregating the results into a single response (reducing round trips for the client).
*   **Cross-Cutting Concerns**: Handling Authentication, SSL termination, Rate Limiting, Cross-Origin Resource Sharing (CORS), and logging at the edge rather than inside every individual service.
*(Tools used: NGINX, AWS API Gateway, Kong).*

### 4. How do you handle session state in a horizontally scaled Node.js application?
If you have 5 instances of a Node API behind a Load Balancer, a user logging into Instance A will fail subsequent requests if routed to Instance B, because Instance B doesn't share memory with A.
*   **Solution 1: Statelessness (JWT)**. Sessions aren't stored on the server at all. The token contains the data. The server just verifies the signature statelessly.
*   **Solution 2: Distributed Cache (Redis / Memcached)**. The session ID is stored in a cookie. When Instance B receives the session ID, it looks up the session data in a centralized, blazing-fast, in-memory Redis cluster that all Node instances have access to.

## Asynchronous Processing and Messaging

### 5. Explain how and why you would use a Message Queue (like RabbitMQ) or an Event Stream (like Kafka) with Node.js.
When a user registers, you might need to: Save to DB, Send an Email, Generate a PDF report, and Notify a data warehouse. Doing this sequentially within the HTTP request will make the response extremely slow.
*   **Message Brokers (RabbitMQ, Redis Pub/Sub, Kafka)** decouple these processes.
*   The Node.js API server drops an "User Registered" event into the queue and immediately responds `200 OK` to the user.
*   Separate background Node.js worker services ("consumers") listen to the queue, pick up the event, and process the heavy tasks (sending email, generating PDF) at their own pace.
*   This ensures the API remains highly responsive and creates a resilient system (if the email service crashes, messages wait in the queue until it comes back online).

### 6. Node.js Streams vs. Buffers: When handling a massive file upload, which do you use and why?
*   **Buffers**: Read the entire file into server RAM before processing it. If 1,000 users upload a 100MB file simultaneously, the server needs 100GB of RAM and will likely crash (V8 memory limit exceeded).
*   **Streams**: Read data chunk by chunk (e.g., 64KB at a time). The server acts as a pipe, directly streaming the incoming HTTP request chunks into a file system write stream or an S3 upload stream. Memory usage remains constant and minimal, regardless of file size.

### 7. How do you implement robust Graceful Shutdown in a Node API?
When deploying new code or scaling down, the process receives a `SIGTERM` signal. If you just kill the process immediately, any users currently downloading data or paying for items will have their connections severed abruptly.
**Graceful Shutdown steps:**
1.  Listen for `SIGTERM` and `SIGINT`.
2.  Tell the HTTP Server to stop accepting *new* connections (`server.close()`).
3.  Wait for existing in-flight requests to finish. Node.js `server.close()` takes a callback when all connections are done.
4.  Once HTTP traffic clears, close database connections (Mongoose/Redis).
5.  Call `process.exit(0)`.
