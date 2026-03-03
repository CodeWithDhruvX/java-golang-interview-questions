# 🏗️ Kafka — Service-Based Companies Fundamentals

> **Level:** 🟢 Junior to 🟡 Intermediate
> **Asked at:** TCS, Infosys, Wipro, Accenture, Cognizant

---

## Q1. What is Apache Kafka and how does its fundamental architecture work?

"Apache Kafka is a distributed event streaming platform primarily utilized for building real-time data pipelines and streaming applications. It functions as a highly scalable, fault-tolerant publish-subscribe message broker.

Its core architecture consists of four main components:
1. **Producers**: The applications that publish (write) data to Kafka topics.
2. **Consumers**: The applications that subscribe to (read) data from topics.
3. **Brokers**: The individual servers in a Kafka cluster that receive, store, and serve the data.
4. **ZooKeeper / KRaft**: The overarching management systems that maintain cluster configuration, manage broker metadata, and facilitate leader elections."

#### 🏢 Company Context
**Level:** 🟢 Junior | **Asked at:** TCS, Infosys — standard foundational question to test elementary distributed system knowledge.

#### Indepth
**Topic & Partition Mechanics:** A Topic is a logical channel where messages are held. However, topics are physically broken into **Partitions**. Partitions allow a single topic's data to be split across multiple brokers, drastically enhancing the system's ability to parallelize processing.

---

## Q2. What is a Consumer Group and why is it important in Kafka?

"A Consumer Group is a collection of consumers that collectively cooperate to consume messages from a specific topic. 

It is Kafka's foundational mechanism for achieving massive read scalability. Each individual partition within a given topic can only be assigned to exactly *one* consumer within that Consumer Group. This guarantees that every message in the topic will be processed by only one consumer in the group, preventing any duplicate processing of data without complex locking.

If you have 4 partitions and 4 consumers in a group, each gets 1 partition. If you add a 5th consumer, it will sit completely idle because there are no available partitions left."

#### 🏢 Company Context
**Level:** 🟢 Junior to 🟡 Intermediate | **Asked at:** Accenture, Wipro — assessing how horizontally scalable read operations function.

#### Indepth
**Fan-out Pattern:** If you want multiple independent applications to read the exact same data stream (e.g., an Analytics service and an Emailing service reading new user signups), you configure them with **different** Consumer Group IDs. This results in Kafka handing a full copy of the data stream to both groups independently.

---

## Q3. Explain the difference between Kafka and RabbitMQ.

"While both can act as message brokers, they belong to entirely different architectural paradigms:

**1. Message Handling and Retention:**
**RabbitMQ:** Traditional message broker using the smart-broker/dumb-consumer model. After a consumer reads an acknowledged message, RabbitMQ immediately deletes it.
**Kafka:** Functions as an immutable commit log (dumb-broker/smart-consumer). Messages are appended to a log and retained regardless of consumption until the defined retention period (e.g., 7 days) expires. Consumers independently track what they have read using 'offsets'.

**2. Routing vs. Streaming:**
**RabbitMQ:** Exceptional at complex routing logic, utilizing complex exchange types and binding rules to precisely route messages to queues.
**Kafka:** Used ideally for high-throughput event streaming where data needs to be retained, re-played, or analyzed stream-by-stream."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Cognizant, Tech Mahindra — to evaluate the candidate's intuition heavily focused on picking the right tool for an architecture.

#### Indepth
**Push vs Pull:** RabbitMQ generally uses a Push architecture—the broker pushes the message actively to the consumer. Kafka employs a Pull architecture, where consumers poll for batches of new data, preventing consumer exhaustion during large bursts of traffic (backpressure control).

---

## Q4. How do you monitor a Kafka cluster and debug common issues like Consumer Lag?

"Consumer lag is the delta between the latest message appended to a partition by a producer (Log End Offset) and the last message successfully processed by the consumer (Current Offset). 

In production environments, monitoring consumer lag is critical. I generally use tools like **Prometheus and Grafana** mapped via JMX metrics from the JVM (e.g., `records-lag-max`). Alternatively, native tools like `kafka-consumer-groups.sh` expose this value in the terminal.

If consumer lag is consistently spiking, it implies the consumer is processing slower than the producer writes. To resolve it, I would fundamentally scale the consumer group. Because one consumer can only read one partition max, I would have to proportionally increase the number of partitions on the topic, and then spin up corresponding consumer instances to balance the elevated load."

#### 🏢 Company Context
**Level:** 🟡 Intermediate | **Asked at:** Infosys, Wipro — assessing practical system maintenance and support capability.

#### Indepth
**Garbage Collection Pauses:** Sometimes lag isn't down to code logic but JVM tuning. Excessive GC pauses stall the consumer's event loop, causing it to randomly miss its heartbeat intervals (`heartbeat.interval.ms`), forcing rebalances and resulting in erratic artificial lag spikes. 
---
