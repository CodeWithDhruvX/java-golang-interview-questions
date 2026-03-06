## 🟢 Messaging Systems: Scenarios & Code Iterations (Questions 61-65)

### Question 61: Scenario - Designing a Real-Time Notification System

**Scenario:** You need to design a notification system where users receive alerts (email, SMS, push) immediately after an event occurs (e.g., an order is placed). You have millions of users and events. Would you choose RabbitMQ or Kafka, and why?

**Answer:**
*   **Choice:** Both can work, but **RabbitMQ** is often the stronger choice for traditional routing of individual notifications, while **Kafka** is better if notifications need to be replayed or analyzed as a stream.
*   **Why RabbitMQ:**
    *   **Complex Routing:** You can use a `Topic Exchange` (e.g., `notification.email.important`, `notification.sms.marketing`) to route messages to specific queues handled by dedicated workers.
    *   **Per-Message Acknowledgement:** Important for ensuring *every single* email was sent successfully. If an email worker fails, RabbitMQ efficiently requeues that exact message.
    *   **Dead Letter Queues:** Easy to route failed notifications to a DLQ for retry.
*   **Why Kafka:**
    *   If the exact same "Order Placed" event needs to be consumed by the Notification Service, the Analytics Service, and the Fraud Detection Service independently (Pub-Sub with high retention).
    *   If you need massive throughput (millions of events per second) and horizontal scalability via partitions.

### Question 62: Scenario - Handling "Poison Pill" Messages

**Scenario:** A producer sent a malformed message. Your consumer reads it, throws an exception while processing, and crashes or rejects the message. The broker requeues it, the consumer reads it again, crashes again. This is a "Poison Pill". How do you resolve this?

**Answer:**
*   **RabbitMQ Approach:**
    *   Configure a **Dead Letter Exchange (DLX)**.
    *   Set a configuration so that if a message is rejected (`basic.nack` or `basic.reject` with `requeue=false`), it goes to the DLX/DLQ.
    *   An admin or a separate worker monitors the DLQ to investigate the bad payload.
*   **Kafka Approach:**
    *   Kafka doesn't have built-in DLQs in the same way RabbitMQ does (except when using Kafka Connect).
    *   **Manual Handling:** In the consumer application (Java/Go), catch the deserialization or processing exception (`try-catch`).
    *   Instead of crashing, log the error, optionally write the raw bad message to a separate "error-topic" (your custom DLQ), and **commit the offset** so the consumer can move on to the next message.

### Question 63: Scenario - Dealing with Traffic Spikes (Consumer Lag)

**Scenario:** It's Black Friday. Your producers are publishing 10,000 messages/sec, but your consumers can only process 2,000 messages/sec. How do you handle this lag?

**Answer:**
*   **Kafka:**
    *   **Increase Consumers:** If your topic has 50 partitions but only 10 consumers in the group, spin up 40 more consumer instances.
    *   **Increase Partitions (Careful):** If you already have 10 consumers for 10 partitions, you must increase the number of partitions (e.g., to 30) AND spin up more consumers.
    *   **Batch Processing:** Modify the consumer to process messages in batches instead of one-by-one to improve database write speeds.
*   **RabbitMQ:**
    *   **Add More Workers:** Simply add more consumer instances listening to the same queue. RabbitMQ will round-robin the messages.
    *   **Adjust Prefetch Count:** Ensure `basic.qos` (prefetch) is set appropriately. If consumers are idle waiting for network IO, increasing prefetch might help.

### Question 64: Code snippet - Kafka Consumer code in Java (Spring Boot)

**Answer:**
In Java / Spring Boot, Kafka consumption is usually handled using the `@KafkaListener` annotation. 

```java
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.support.Acknowledgment;
import org.springframework.stereotype.Service;

@Service
public class OrderNotificationConsumer {

    // Listens to the 'order-events' topic, part of the 'notification-group'
    @KafkaListener(topics = "order-events", groupId = "notification-group")
    public void consumeOrderEvent(String message, Acknowledgment acknowledgment) {
        try {
            System.out.println("Received Order Event: " + message);
            // 1. Parse message
            // 2. Send Email/SMS
            
            // MANUAL ACKNOWLEDGEMENT: Commits the offset only after successful processing
            // Requires spring.kafka.listener.ack-mode=MANUAL context property
            acknowledgment.acknowledge(); 
            
        } catch (Exception e) {
            System.err.println("Error processing message. Sending to DLQ topic...");
            // Handle Poison Pill (e.g., send to error topic)
            // Acknowledge anyway so we don't get stuck in a loop
            acknowledgment.acknowledge(); 
        }
    }
}
```

### Question 65: Code snippet - Kafka Consumer code in Golang (Segmentio or Confluent)

**Answer:**
In Go, `github.com/segmentio/kafka-go` or `github.com/confluentinc/confluent-kafka-go` are commonly used. Here is an example using `kafka-go` with manual offset commitments.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 1. Initialize the reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-events",
		GroupID: "notification-group", // Required for offset tracking
		MaxBytes:  10e6, // 10MB
	})
	defer r.Close()

	fmt.Println("Started consuming...")

	for {
		// 2. Fetch the message (DOES NOT commit offset yet)
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			log.Printf("Error fetching message: %v\n", err)
			break
		}

		fmt.Printf("Received message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		// 3. Process the message (Simulated)
		err = processMessage(m.Value)
		
		if err != nil {
			log.Printf("Failed to process message (Poison Pill): %s. Moving on.", err)
			// Depending on logic, send to DLQ topic here
		} 

		// 4. Manually commit the offset AFTER processing
		if err := r.CommitMessages(context.Background(), m); err != nil {
			log.Printf("Failed to commit messages: %v\n", err)
		} else {
             fmt.Println("Offset cleanly committed.")
        }
	}
}

func processMessage(payload []byte) error {
	// Simulate work
	time.Sleep(10 * time.Millisecond)
	return nil // Return error if processing fails
}
```
