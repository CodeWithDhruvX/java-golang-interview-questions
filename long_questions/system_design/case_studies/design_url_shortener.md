# Design URL Shortener (TinyURL)

## 1. Requirements
*   **Functional**:
    *   Input: Long URL (https://www.google.com/search?q=design).
    *   Output: Short URL (http://tiny.url/Ab72x).
    *   Redirection: Clicking Short URL takes user to Long URL.
*   **Non-Functional**:
    *   **High Read Heavy**: 100:1 Read:Write ratio.
    *   **Low Latency**: Redirection must be instant.

## 2. API Design
*   `create(api_key, long_url) -> short_url`
*   `get(short_url) -> 301 Redirect (long_url)`

## 3. Core Algorithm (Generating the Alias)

### Approach A: Hashing (MD5/SHA256)
*   Hash the Long URL. `MD5("google.com") = 5d4140...`
*   Take first 7 characters.
*   **Problem**: Collisions. Two different Long URLs might share same first 7 chars.

### Approach B: Base62 Conversion (Preferred)
*   Use a **Counter** (Database Auto Increment ID or Distributed ID Generator like Snowflake).
*   Convert the ID to Base62 (a-z, A-Z, 0-9).
*   *Example*:
    *   ID = 1000000
    *   Base62(1000000) = "4c92"
    *   Short URL = `http://tiny.url/4c92`
*   **Constraint**: 7 chars of Base62 = 62^7 = ~3.5 Trillion combinations. Enough for years.

## 4. Components
1.  **Frontend**: Simple box for input.
2.  **API Server**: Handles requests.
3.  **Token Service (KGS - Key Generation Service)**:
    *   To speed up writes, pre-generate millions of 7-char keys and store in a DB/Redis.
    *   When user asks, just pop one key. (Avoids realtime encoding).
4.  **Database**:
    *   Table: `id(PK), short_key, long_url, created_at`
    *   NoSQL (DynamoDB/Cassandra) is better due to massive scale and simple KV structure.

## 5. 301 vs 302 Redirect
*   **301 (Moved Permanently)**:
    *   Browser caches the redirection. Future requests go directly to Long URL without hitting your server.
    *   *Pros*: Reduced server load.
    *   *Cons*: No analytics (You can't count clicks).
*   **302 (Found / Temporary)**:
    *   Browser hits your server every time.
    *   *Pros*: Accurate Analytics.
    *   *Cons*: Higher server load.

## 6. Interview Questions
1.  **How to detect duplicates?**
    *   *Ans*: If user enters google.com again, should we give new URL or same URL?
    *   Check DB. If exists, return existing.
2.  **What if the pre-generated key DB runs out?**
    *   *Ans*: It's a "Key Generation Service". It should have a background worker adding new keys when pool size < 20%.
