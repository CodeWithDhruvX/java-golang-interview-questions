# Full-Stack System Design (Product-Based Companies)

These open-ended questions test your ability to take a vague problem involving both frontend and backend and synthesize a complete, scalable MEAN/MERN architecture.

## common Full-Stack Scenarios

### 1. How would you design a Real-Time Collaborative Document Editor (like Google Docs) using the MEAN/MERN stack?
**Key Challenges**: Real-time sync, conflict resolution, offline support, connection management.
*   **Frontend (React/Angular)**:
    *   State management (Redux/Zustand) for the document content.
    *   Use WebSockets (Socket.io) for real-time, bi-directional communication with the Node backend.
    *   Implement **Operational Transformation (OT)** or **Conflict-Free Replicated Data Types (CRDTs)** algorithms in the frontend to handle concurrent edits without overwriting others' work.
*   **Backend (Node.js/Express)**:
    *   A WebSocket server to broadcast document changes to all connected clients editing the same doc.
    *   Redis Pub/Sub to scale WebSockets across multiple Node.js instances (so a user connected to Server A sees edits from a user connected to Server B).
    *   A robust queuing system (Kafka or BullMQ) to batch document state changes and write them to the primary database asynchronously to avoid overwhelming it with small keystroke updates.
*   **Database (MongoDB)**:
    *   Store the "source of truth" document.
    *   Alternatively, store an array of document "deltas" or operations, allowing historical playback and version history.

### 2. Design a URL Shortening Service (like Bitly).
**Key Challenges**: Extremely high read throughput, generating unique short codes without collisions, analytics tracking.
*   **Frontend**: Simple UI to input a URL and receive the shortened version. Dashboard for analytics.
*   **Backend Architecture**:
    *   Node.js stateless API scaling horizontally.
    *   **Short Code Generation**:
        *   Approach A: Base62 encode a unique auto-incrementing ID from a database.
        *   Approach B: Hash the long URL (MD5) and take the first 7 characters. Deal with collisions by appending a counter.
        *   Approach C (Best for high scale): Key Generation Service. A background worker pre-generates unique 7-character strings (e.g., using ZooKeeper for distributed counters) and stores them in Redis. The Node API simply pops a pre-generated key from Redis—virtually instant.
*   **Database & Caching**:
    *   MongoDB to map `shortCode -> longUrl`.
    *   **Critical**: Use a distributed cache (Redis/Memcached) acting as a read-through layer for the Node servers. Over 99% of requests will hit the cache, not MongoDB.
*   **Analytics**: Asynchronously push click events (IP, timestamp, user agent) to a message queue (Kafka) for processing by a data warehouse, avoiding blocking the redirection logic.

### 3. You are building an E-Commerce application. How do you handle Flash Sales (Millions of requests in 5 minutes for limited stock)?
**Key Challenges**: Database locking, massive concurrency, over-selling (race conditions).
*   **The Problem with MongoDB/SQL in a Flash Sale**: If 100,000 users try to decrement the inventory counter for "PS5" simultaneously, row/document level locks will cause massive latency spikes, timeouts, and potentially database crashes. Over-selling can occur if transactions are poorly isolated.
*   **The Architecture (Redis is the hero)**:
    1.  Before the sale, load total inventory counts into Redis (e.g., `SET ps5_stock 1000`).
    2.  When a user clicks "Buy", the Node server executes a Redis atomic decrement: `DECR ps5_stock`.
    3.  If the result is `>= 0`, the sale is valid. The Node server immediately responds "Success - Order Processing" to the user.
    4.  Node.js places an "Order Created" message onto a Message Broker (RabbitMQ/Kafka) containing user and product details.
    5.  If the result of `DECR` is `< 0`, the stock is gone. Reject the request instantly.
    6.  Background worker services pull messages from the queue at a steady, sustainable rate to handle the heavy lifting: charging the credit card via Stripe, inserting the definitive order record into MongoDB (which now only sees a slow, steady stream of inserts), and sending confirmation emails.

### 4. How would you handle Pagination on an endpoint returning millions of records?
*   **Offset Pagination (The bad way)**: `db.collection.find().skip(500000).limit(20)`.
    *   As the offset grows, performance degrades exponentially because the database must scan all 500,000 previous documents just to find the 20 you want.
*   **Cursor-Based / Keyset Pagination (The right way for large datasets)**:
    *   Instead of an offset page number, the client sends the `_id` (or timestamp) of the *last* item they received.
    *   Query: `db.collection.find({ _id: { $gt: last_seen_id } }).limit(20)`.
    *   Because `_id` is indexed, MongoDB instantly jumps to that node in the B-tree index and scans the next 20 items. Performance remains consistent O(1) whether on page 1 or page 10,000.
