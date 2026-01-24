## 🟢 RabbitMQ (Questions 1-20)

### Question 1: What is RabbitMQ and why is it used?

**Answer:**
RabbitMQ is an open-source message-broker software (sometimes called message-oriented middleware) that originally implemented the Advanced Message Queuing Protocol (AMQP) and has since been extended with a plug-in architecture to support Streaming Text Oriented Messaging Protocol (STOMP), MQTT, and other protocols.

*   **Why it is used:**
    *   **Decoupling:** Separates the producers (data creators) from consumers (data processors).
    *   **Buffering:** Helps control the flow of data by buffering messages when the consumer is unable to process them immediately.
    *   **Asynchronous Processing:** Allows tasks to be processed in the background, improving the responsiveness of the main application.
    *   **Scalability:** Distributes the workload across multiple consumers.
    *   **Reliability:** Ensures message delivery even if a component fails (using persistence and acknowledgments).

### Question 2: Explain the basic architecture of RabbitMQ.

**Answer:**
The RabbitMQ architecture revolves around the following flow:
1.  **Producer** generates a message and publishes it to an **Exchange**.
2.  The **Exchange** receives the message and routes it to one or more **Queues** based on specific rules (Bindings/Routing Keys).
3.  The **Queue** stores the message until a consumer is ready to process it.
4.  **Consumer** subscribes to the Queue and receives messages.

*   **Key Concept:** Producers never send messages directly to queues. They always send to an Exchange.

### Question 3: What are the main components of RabbitMQ (Producer, Consumer, Queue, Exchange, Binding)?

**Answer:**
*   **Producer:** User application that sends messages.
*   **Exchange:** A message routing agent. It accepts messages from producers and pushes them to queues depending on routing rules.
*   **Queue:** A buffer that stores messages. It resides inside RabbitMQ.
*   **Binding:** A link between a queue and an exchange. It tells the exchange which queue to send messages to.
*   **Consumer:** User application that waits for and processes messages.

### Question 4: What is an Exchange in RabbitMQ? Name the different types of Exchanges.

**Answer:**
An Exchange is the entry point for messages in RabbitMQ. It decides where a message should go.
*   **Direct Exchange:** Delivers messages to queues based on the message routing key.
*   **Fanout Exchange:** Broadcasts messages to all bound queues indiscriminately.
*   **Topic Exchange:** Routes messages based on matching the routing key with a pattern.
*   **Headers Exchange:** Routes messages based on message header attributes instead of the routing key.

### Question 5: Explain the Direct Exchange.

**Answer:**
A Direct exchange delivers messages to queues based on the **exact match** of the message routing key.
*   **Scenario:** A logging system.
*   **Bindings:**
    *   Queue A bound with key `error`.
    *   Queue B bound with key `info`.
*   **Behavior:** A message published with routing key `error` goes **only** to Queue A.

### Question 6: Explain the Topic Exchange.

**Answer:**
Topic exchanges route messages to one or many queues based on matching a routing key against a pattern.
*   **Wildcards:**
    *   `*` (star) can substitute for exactly one word.
    *   `#` (hash) can substitute for zero or more words.
*   **Example:** `user.created`, `user.deleted.eu`, `order.created`.
    *   Binding `user.*` matches `user.created` but not `user.deleted.eu`.
    *   Binding `user.#` matches both.

### Question 7: Explain the Fanout Exchange.

**Answer:**
A Fanout exchange copies and routes a received message to **all** queues that are bound to it, regardless of the routing key or pattern.
*   **Use Case:** Mass notifications, updating a scoreboard, or sending logs to multiple systems (one for disk storage, one for real-time analysis).
*   **Performance:** It is the fastest exchange type because it ignores routing logic.

### Question 8: Explain the Headers Exchange.

**Answer:**
A Headers exchange routes messages based on arguments containing headers and optional values. It ignores the routing key.
*   **Arguments:**
    *   `x-match = all`: All header pairs must match.
    *   `x-match = any`: At least one header pair must match.
*   **Use Case:** Complex routing where multiple attributes (format=pdf, type=report, encrypted=true) determine the destination.

### Question 9: What is a Binding in RabbitMQ?

**Answer:**
A binding is a "rule" that RabbitMQ uses to route messages from an exchange to a queue.
*   **Metaphor:** If the Exchange is a Mail Sorting Office and the Queue is a Mailbox, the Binding is the Address logic ("Send all mail for Zip 90210 to this box").
*   It can optionally take a **routing key** to filter messages.

### Question 10: What is a Dead Letter Exchange (DLX)?

**Answer:**
A Dead Letter Exchange is a normal exchange that RabbitMQ automatically routes "dead" messages to.
*   **When is a message "dead"?**
    1.  The message is rejected (`basic.reject` or `basic.nack`) with `requeue=false`.
    2.  The message TTL (Time To Live) has expired.
    3.  The queue length limit is exceeded.
*   **Purpose:** Allows you to capture and debug problematic messages instead of losing them silently.

### Question 11: How does RabbitMQ handle message durability?

**Answer:**
To ensure messages survive a broker restart or crash, two things are required:
1.  **Durable Queue:** The queue itself must be declared as `durable` (so the queue definition survives).
2.  **Persistent Messages:** The producer must mark the message delivery mode as `2` (persistent).
*   **Trade-off:** Persistence impacts performance because RabbitMQ must write the message to the disk.

### Question 12: What is the difference between transient and durable queues?

**Answer:**
*   **Durable Queue:** Its metadata is stored on disk. If RabbitMQ restarts, the queue will still exist.
*   **Transient Queue:** Metadata is stored in memory. If RabbitMQ restarts, the queue is deleted.
*   **Note:** Durable queues do **not** imply that the messages inside them are persistent. That is a separate property of the message itself.

### Question 13: How can you ensure message delivery in RabbitMQ?

**Answer:**
1.  **Publisher Confirms:** The broker sends an acknowledgment to the producer once the message is safely accepted (and persisted, if applicable).
2.  **Consumer Acknowledgments:** The consumer sends an ack to the broker only after it has successfully processed the message (`basic.ack`).
3.  **Transactions:** (Older method, heavier) Wrap publish operations in a transaction committed to the broker.

### Question 14: What are consumer acknowledgments?

**Answer:**
RabbitMQ needs to know when to delete a message from the queue.
*   **Automatic Ack (`autoAck=true`):** Message is deleted immediately after being sent to the consumer (fire-and-forget). Risk of data loss if consumer crashes.
*   **Manual Ack (`autoAck=false`):** Application explicitly sends an ack (`channel.basicAck`) after processing. If the consumer dies before acking, RabbitMQ re-queues the message.

### Question 15: What is prefetch count in RabbitMQ?

**Answer:**
Prefetch count (`basic.qos`) defines how many unacknowledged messages a consumer can hold at one time.
*   **Purpose:** Prevents a fast producer from overwhelming a slow consumer.
*   **Mechanism:** If prefetch is set to 1, RabbitMQ waits for the consumer to ack the previous message before sending a new one. This ensures fair load balancing among multiple consumers.

### Question 16: How does RabbitMQ support clustering?

**Answer:**
RabbitMQ clustering connects multiple nodes to form a single logical broker.
*   **State Sharing:** All nodes share users, virtual hosts, queues, exchanges, and bindings.
*   **Queue Location:** By default, a queue resides on the node where it was created. Metadata is visible everywhere, but the actual message data is on one node.
*   **Mirrored Queues (Classic):** Copies messages to other nodes for High Availability (HA). *Deprecated in favor of Quorum Queues.*
*   **Quorum Queues:** New HA queue type based on Raft consensus algorithm.

### Question 17: What is VHost (Virtual Host) in RabbitMQ?

**Answer:**
A Virtual Host (vhost) provides logical grouping and separation of resources.
*   It is similar to a virtual machine or a Kubernetes Namespace.
*   **Function:** Separate apps can share the same RabbitMQ instance without naming collisions or permission issues.
*   **Security:** Permissions are applied per vhost.

### Question 18: What is the Erlang Cookie and why is it important in RabbitMQ clustering?

**Answer:**
RabbitMQ is built on Erlang. Erlang nodes communicate with each other using a shared secret called the **Erlang Cookie**.
*   **Importance:** For two nodes to join a cluster, they must have the exact same Erlang Cookie string (usually stored in `.erlang.cookie` file).
*   If cookies don't match, the nodes will refuse to communicate.

### Question 19: How do you monitor RabbitMQ?

**Answer:**
1.  **RabbitMQ Management Plugin:** A built-in web UI (port 15672) showing connections, queues, rates, and consumers.
2.  **`rabbitmqctl`:** Command-line tool for admin tasks and status checks.
3.  **Prometheus & Grafana:** Enable the `rabbitmq_prometheus` plugin to export metrics (queue depth, message rates) to Prometheus.

### Question 20: Compare RabbitMQ with Kafka.

**Answer:**

| Feature | RabbitMQ | Apache Kafka |
| :--- | :--- | :--- |
| **Model** | Push-based (Smart Broker). | Pull-based (Dumb Broker, Smart Consumer). |
| **Routing** | Complex (Exchanges, Bindings). | Simple (Topics, Partitions). |
| **Message Order** | Guaranteed per queue (mostly). | Guaranteed per partition only. |
| **Persistence** | Messages deleted after consumption. | Log-based storage (retention period). |
| **Throughput** | Moderate (4k-10k msgs/sec). | Very High (100k-Millions msgs/sec). |
| **Use Case** | Complex routing, immediate processing. | Stream processing, event sourcing, replayability. |
