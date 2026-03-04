# System Design (HLD) - API Rate Limiter

## Problem Statement
Design a massive-scale API Rate Limiter to protect back-end services from being overwhelmed by too many requests (e.g., DDoS attacks, scraping, buggy client code).

## 1. Requirements Clarification
### functional Requirements
*   **Limit Requests:** Ensure that a single user, IP, or API Key cannot exceed a predefined number of requests within a given time window (e.g., 100 requests per minute).
*   **Response:** If a user exceeds the limit, return a standard HTTP `429 Too Many Requests` status code.

### Non-Functional Requirements
*   **Extremely Low Latency:** The rate limiter must intercept every single API call without adding noticeable latency (ideally < 2ms).
*   **High Availability:** It shouldn't be a single point of failure. If the rate limiter cluster dies, it should "fail-open" (allow requests to pass through) rather than taking down the entire API.

## 2. Defining Rate Limiting Algorithms
Before architectural design, you must understand the algorithms:
*   **Token Bucket:** A bucket holds `N` tokens and refills at a constant rate `R`. Every request takes a token. *Most common, allows burst traffic.*
*   **Leaky Bucket:** Requests enter a queue and are processed at a constant rate. Smooths out traffic.
*   **Fixed Window Counter:** A counter increments for a specific time window (e.g., 12:00:00 to 12:01:00). Very fast, but subject to "bursts" at window edges.
*   **Sliding Window Log:** Stores timestamps of all API requests in a sorted set. Highly accurate, but terrible memory footprint for high volume.
*   **Sliding Window Counter:** A hybrid combining fixed windows with weighted proportions of the previous window. Best balance of speed and accuracy.

## 3. High-Level Integration Where does it live?
The Rate Limiter is typically built as a middleware component embedded inside the **API Gateway** (e.g., Kong, NGINX, AWS API Gateway) because it runs at the edge of the network.

## 4. Centralized vs Decentralized State
If an API Gateway has 50 nodes receiving traffic from the same IP address, how do they keep a synchronized count?

*   **Option 1: Sticky Sessions (Bad)**
    Traffic from an IP always hits the same Gateway node. Terrible for load balancing.
*   **Option 2: Centralized Cache (Good but risky)**
    All Gateway nodes ask a central Redis cluster (`INCR user_id_key`). Add network latency.
*   **Option 3: Eventual Consistency / Gossip Protocol (Best for massive scale)**
    Each node keeps a local count in-memory and asynchronously shares its count with the other nodes over UDP. (e.g., Redis Enterprise's CRDT implementation).

## 5. Detailed Design using Redis (Token Bucket)

*   **Key:** `rate_limit_token:user_id:1234`
*   **Value:** `{"tokens_left": 80, "last_refill_timestamp": 1699990000}`

**Flow:**
1. Request arrives. We fetch the key from Redis.
2. We calculate how many tokens should have been refilled since the `last_refill_timestamp`.
3. If `tokens_left` + `refilled` > 1:
   * Subtract 1 token.
   * Update the `tokens_left` and `timestamp` in Redis.
   * Allow request.
4. Else:
   * Block request (`429`).

**Optimization (Lua Script):**
Doing a `GET`, calculating, and then a `SET` in Redis causes a Race Condition if 5 concurrent requests hit simultaneously. We bundle this logic into a **Redis Lua Script**, which Redis guarantees will execute atomically.

## 6. Follow-up Questions for Candidate
1.  How do you inform the client how many requests they have left? (Injecting HTTP headers: `X-RateLimit-Remaining`, `X-RateLimit-Reset`).
2.  If the Redis cluster is down, do you block all traffic? (No, fail-open policy; relying on backend natural limits until it recovers).
