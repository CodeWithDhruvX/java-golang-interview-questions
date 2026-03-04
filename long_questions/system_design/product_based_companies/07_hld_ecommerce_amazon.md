# System Design (HLD) - E-commerce Platform (Amazon/Flipkart)

## Problem Statement
Design an E-commerce store like Amazon supporting millions of concurrent users browsing, searching, adding items to carts, and checking out.

## 1. Requirements Clarification
### Functional Requirements
*   **Search/Catalog:** Users can browse and search for products by name or category.
*   **Shopping Cart:** Add, remove, update items.
*   **Checkout & Order Processing:** Secure payment and inventory decrement.
*   **Reviews & Ratings:** Users can rate and review purchased items.

### Non-Functional Requirements
*   **High Availability:** The Catalog/Search must be highly available (100% read uptime).
*   **Strong Consistency:** Checkout and Inventory management must be Strongly Consistent to avoid selling an item you don't actually have in stock (Overselling).
*   **Scalability:** Must handle events like "Black Friday" where traffic spikes by 10x.

## 2. High-Level Architecture (Microservices)

An e-commerce giant is deeply decomposed into isolated microservices communicating via HTTP APIs, gRPC, and Asynchronous Message Brokers.

```text
[ Client App ] ---> [ API Gateway ]
                          |
             +------------+------------+---------------+
             |            |            |               |
     [ Search Svc ] [ Cart Svc ] [ Checkout Svc ] [ Rating Svc ]
             |            |            |               |
     [ Elasticsearch] [ Redis ]  [ RDBMS PostgreSQL ] [ NoSQL ]
```

## 3. Component Deep Dive

### A. The Search & Catalog Service
*   **Database:** You cannot query a massive SQL database for "Red shoes size 10" with `LIKE` clauses fast enough. We use a full-text search engine like **Elasticsearch** (or Apache Solr).
*   **Data Ingestion:** When a seller adds a product to the main SQL DB, an event is emitted to a Kafka topic. A consumer reads it and indexed the document in Elasticsearch.

### B. The Shopping Cart Service
*   **Storage:** The cart is ephemeral but must persist across sessions. We use an in-memory key-value store like **Redis**.
    *   *Key:* User ID or Session ID.
    *   *Value:* Hash map of `ProductID` -> `Quantity`.
*   **Availability:** If a Redis node dies, the user loses their cart. We mitigate this with Redis Cluster replication, but occasionally a cart is lost—this is acceptable compared to the Checkout database failing.

### C. The Checkout & Inventory Service (Crucial)
This is where **Consistency > Availability** (CAP Theorem).
*   **Database:** A relational database (RDBMS) like MySQL or PostgreSQL, because ACID transactions are mandatory.
*   **The Overselling Problem:** Two users try to buy the last iPhone at the exact same millisecond.
    *   *Naive implementation:* Get stock (returns 1). If > 0, decrement by 1. *RACE CONDITION! Both users buy it.*
    *   *Solution:* We must lock the row during the transaction. Often this involves:
        *   **Pessimistic Locking:** `SELECT * FROM inventory WHERE item_id = 99 FOR UPDATE;` (Database locks the row until transaction commit).
        *   **Distributed Lock:** Acquire a Redis or Zookeeper lock on `item_id:99` before hitting the DB.

## 4. Payment Gateway and Idempotency
When we charge a user's credit card, network errors happen.
*   The client clicks 'Pay', the server contacts Stripe/PayPal. Stripe charges the card but the server crashes before telling the client.
*   The client clicks 'Pay' again.
*   **Idempotency Key:** The client generates a unique UUID for the checkout attempt. The backend stores this UUID. If the client retries with the same UUID, the server knows it's a duplicate and doesn't charge the card twice.

## 5. Review & Recommendation System (Big Data)
*   **Storage:** A wide-column store like Cassandra is excellent for massive append-only review data.
*   **Recommendations:** User clicks and purchases are streamed into Kafka -> Hadoop/Spark -> Machine Learning Pipelines -> Generated models stored in Redis for fast user-profile lookup at runtime.

## 6. Follow-up Questions for Candidate
1.  How do you design the architecture for "Flash Sales" (e.g., 1000 iPhones at $1)? (Queueing all requests in Kafka immediately to protect the DB, processing them asynchronously, and dropping excess requests at the API Gateway level).
2.  Explain the Saga Pattern or 2-Phase Commit when an Order involves updating the Order DB, Payment DB, and Inventory DB sequentially.
