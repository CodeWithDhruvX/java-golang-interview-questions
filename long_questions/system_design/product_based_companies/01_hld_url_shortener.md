# System Design (HLD) - URL Shortener (TinyURL)

## Problem Statement
Design a URL Shortening service like TinyURL or Bitly. The service will provide a short alias for a long URL and redirect users who access the short URL to the original long URL.

## 1. Requirements Clarification
### functional Requirements
*   **Shortening:** Given a long URL, return a short URL.
*   **Redirection:** Given a short URL, redirect to the original long URL.
*   **Custom URLs:** Users can optionally specify a custom short link.
*   **Expiration:** Links should expire after a standard duration (e.g., 5 years) or a user-specified time.

### Non-Functional Requirements
*   **Highly Available:** Redirection must succeed 99.99% of the time.
*   **Low Latency:** Redirection should be extremely fast.
*   **Scale:** The system must handle a heavy read load (Read-heavy system, typically 100:1 read-to-write ratio).

## 2. Capacity Estimation (Back-of-the-envelope)
*   **Traffic:** 100 million new URLs generated per month.
*   **Reads:** 10 billion redirections per month.
*   **Storage:** 100M URLs * 12 months * 5 years * 500 bytes/URL ≈ 3 TB of storage.

## 3. High-Level Design (Architecture)

```text
[ Client ] ---> [ Load Balancer ] ---> [ Web Servers (API) ]
                                            |   |
                   +---(Cache Miss/Write)---+   +---(Cache Hit)---> [ Redis Cache ]
                   |
             [ Database (NoSQL) ]
```

## 4. Database Design
A NoSQL database like **Cassandra** or **DynamoDB** is ideal because:
*   We need to store billions of rows.
*   There are no complex joins.
*   High availability and scaling out (horizontal scaling) are extremely important.

**Table Schema (URLMap):**
*   `hash` (Partition Key, varchar 7)
*   `original_url` (varchar)
*   `creation_date` (datetime)
*   `expiration_date` (datetime)
*   `user_id` (integer)

## 5. Core Component Design (The Encoding Logic)
How do we generate a unique short URL?
*   **Base62 Encoding:** Base62 uses `a-z`, `A-Z`, `0-9` (26+26+10 = 62 characters). A 7-character Base62 string provides $62^7 \approx 3.5$ trillion unique permutations.
*   **Algorithm 1 (Hashing):** MD5(LongURL) -> take the first 7 chars. *Problem:* Collisions can occur.
*   **Algorithm 2 (Counter - Best Approach):** Use an auto-incrementing ID generator (like Snowflake or a centralized Zookeeper counter). Convert the generated base-10 ID to Base62.
    *   Example: ID `125` -> Base62 `cb`.
    *   Every new URL gets the next integer ID, guaranteeing uniqueness.

## 6. Bottlenecks and Optimizations
*   **Database Read Bottleneck:** Introduce a caching layer like Redis or Memcached. Cache the most frequently accessed URLs (apply LRU eviction policy).
*   **ID Generator Single Point of Failure:** Provide a distributed ID generator (e.g., Twitter Snowflake) or use Zookeeper to hand out blocks of IDs to different application servers.

## 7. Follow-up Questions for Candidate
1.  How do you prevent malicious users from exhausting all 3.5 trillion URLs? (Rate Limiting based on API keys/IP addresses).
2.  How would you design the analytics service to track click counts per link without slowing down the redirection? (Async processing using Kafka).
