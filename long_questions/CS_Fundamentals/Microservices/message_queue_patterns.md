# Message Queue Patterns

## 1. Why use a Message Queue?
*   **Decoupling**: Services don't need to know about each other.
*   **Async Processing**: Fire and forget. (e.g., Send Email).
*   **Buffering (Throttling)**: If Consumers are slow, the Queue acts as a buffer so Producers don't crash.

## 2. Common Patterns

### A. Point-to-Point (Work Queue)
*   **Mechanism**: One Producer, One Queue, Multiple Consumers.
*   **Behavior**: Each message is consumed by **exactly one** consumer.
*   **Use Case**: Task distribution. Competing Consumers pattern. (RabbitMQ Default).

### B. Publish-Subscribe (Fan-out)
*   **Mechanism**: One Producer, Multiple Queues (one per subscriber).
*   **Behavior**: Every message matches multiple subscriptions and is copied to each queue. Each subscriber gets a copy.
*   **Use Case**: User uploads a video -> 1. Encode Service, 2. Notify Service, 3. Analytics Service. (SNS -> SQS, Kafka Topics).

## 3. Advanced Patterns

### C. Dead Letter Queue (DLQ)
*   **Problem**: What if a message is malformed and crashes the consumer? The consumer restarts, reads it again, and crashes loop.
*   **Solution**: If a message fails processing `N` times, move it to a separate "Dead Letter Queue" for manual inspection. Don't block the main queue.

### D. Delayed Queue
*   **Mechanism**: Messages are hidden for `X` seconds before becoming visible to consumers.
*   **Use Case**: "Remind me in 15 mins", "Retry this failed payment in 1 hour" (Exponential Backoff).

## 4. Kafka vs RabbitMQ

| Feature | RabbitMQ (Push) | Kafka (Pull) |
| :--- | :--- | :--- |
| **Model** | Smart Broker, Dumb Consumer | Dumb Broker, Smart Consumer |
| **Persistence** | Queue usually empty (processed) | Log (Store events for days) |
| **Throughput** | High (20k/sec) | Extreme (100k+/sec) |
| **Order** | Not strictly guaranteed | Guaranteed within a partition |
| **Use Case** | Complex routing, Task Queue | Event Streaming, Logs, Analytics |

## 5. Interview Questions
1.  **What is "At-Least-Once" delivery?**
    *   *Ans*: The system guarantees the message arrives, but it *might* arrive twice (e.g., if consumer crashes after processing but before sending ACK). Consumer must be **Idempotent**.
2.  **How to ensure strict ordering?**
    *   *Ans*: In Kafka, use the same **Partition Key** (e.g., UserID). All events for that UserID go to the same partition, which is consumed by one consumer in order.
