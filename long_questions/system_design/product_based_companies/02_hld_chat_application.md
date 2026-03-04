# System Design (HLD) - Chat Application (WhatsApp/Telegram)

## Problem Statement
Design a global instant messaging application like WhatsApp, Facebook Messenger, or Telegram.

## 1. Requirements Clarification
### Functional Requirements
*   **1-on-1 Messaging:** Users can send text messages to each other.
*   **Online/Offline Status:** Users can see if their contacts are online, offline, or "last seen".
*   **Message Status:** Sent, Delivered, and Read receipts.
*   **Push Notifications:** Users should receive notifications for messages when offline.

*Out of scope for initial MVP: Group chats, Voice/Video calls, Media uploads.*

### Non-Functional Requirements
*   **Low Latency:** Real-time message delivery is critical.
*   **High Availability:** The system should rarely go down.
*   **Scalability:** Must support 500 million daily active users (DAU).

## 2. Protocol Selection
Standard HTTP is not suitable for realtime chat because the client has to constantly poll the server.
*   **WebSockets:** The preferred protocol. It establishes a persistent, bi-directional connection between the client and the server, allowing the server to push messages to the client instantly.

## 3. High-Level Design (Architecture)

The architecture is split into three main parts:
1.  **Stateless Services:** Handling login, signup, user profiles, and friend lists (Standard HTTP REST APIs).
2.  **Stateful Services:** The Chat/WebSocket Servers that maintain persistent connections with active users.
3.  **Third-Party Integration:** Push Notifications (APNs for iOS, FCM for Android).

```text
[ User A ] <--(WebSocket)--> [ Chat Server 1 ] <---(RPC/Message Bus)---> [ Chat Server 2 ] <--(WebSocket)--> [ User B ]
                                     |                                         |
                                     +-----> [ Redis / Key-Value Store ] <-----+
                                                  (Presence Service)
```

## 4. Detailed Component Design

### The Chat Flow (Sending a message)
1. User A sends a message to User B via their open WebSocket connection to `Chat Server 1`.
2. `Chat Server 1` receives the message. It asks the **Presence Service** (Redis): "Which Chat Server is User B connected to?"
3. If User B is online and connected to `Chat Server 2`, `Chat Server 1` forwards the message to `Chat Server 2` over the internal network (or via a message broker like Kafka/RabbitMQ).
4. `Chat Server 2` drops the message down its WebSocket connection to User B.
5. If User B is offline, the system stores the message in the Database and triggers a Push Notification.

### Presence Service (Online Status)
How to know if a user is online?
*   When a user connects, their status in Redis is set to `Online` along with their connected `Chat Server ID`.
*   Users send "heartbeat" pings every few seconds.
*   If a heartbeat is missed for (e.g., 30 seconds), the user is marked `Offline`.

## 5. Database Design
*   **Key-Value Store (Cassandra/HBase):** Excellent for message history because of fast write speeds and efficient range queries (e.g., ordering messages by timestamp).
*   **Schema (Messages Table):**
    *   `message_id` (Primary Key, Snowflake ID)
    *   `chat_id` (Partition Key, derived from user A and user B IDs)
    *   `sender_id`
    *   `created_at` (Sort Key)
    *   `content`

## 6. Bottlenecks and Considerations
*   **Handling Millions of Connections:** A single modern server can hold roughly 1 million concurrent WebSocket connections (tuning Linux file descriptors and Epoll). A load balancer must distribute incoming connections evenly.
*   **Message Ordering:** Relies on a distributed ID generator (e.g., Snowflake) to generate sortable `message_id`s rather than relying strictly on clustered timestamps.

## 7. Follow-up Questions for Candidate
1.  How do you implement End-to-End Encryption (E2EE)? (Exchange of public keys via a Key Server; servers cannot decrypt payloads).
2.  How do you handle Group Chats which require fanning out one message to 500 users?
