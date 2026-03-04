# System Design (HLD) - Centralized Notification System

## Problem Statement
Design a centralized Notification System that can be used by an entire enterprise (multiple internal microservices) to send Emails, SMS, and Push Notifications to users.

## 1. Requirements Clarification
### Functional Requirements
*   **Send Notifications:** Accept requests to send Emails, SMS, and Push Notifications.
*   **Templating:** Support dynamic templates (e.g., "Hello {{name}}, your OTP is {{otp}}").
*   **Preferences:** Users can opt-out of certain notification types (e.g., "Don't send SMS, only Email").
*   **Tracking:** Track Delivery and Read statuses.

### Non-Functional Requirements
*   **Highly Available:** Don't drop notification requests.
*   **Scalable:** Can handle millions of notifications per day.
*   **Asynchronous:** The system must not block the calling microservices.

## 2. High-Level Architecture

```text
[ Microservice A ] --+
[ Microservice B ] --+---> [ API Gateway ] ---> [ Notification API ]
[ Microservice C ] --+                                |
                                            [ Message Queue (Kafka) ]
                                                      |
                                    +-----------------+-----------------+
                                    |                 |                 |
                             [ Email Worker ]    [ SMS Worker ]   [ Push Worker ]
                                    |                 |                 |
                             (SendGrid/AWS SES) (Twilio/Plivo) (APNs/FCM)
```

## 3. Component Deep Dive

### A. The Ingestion API
Internal microservices send an HTTP POST request to the Notification API.
*   **Payload Example:** `{ "user_id": 123, "template_id": "OTP_LOGIN", "channel": "SMS", "params": {"otp": "1234"} }`
*   The API performs basic validation, saves the request to a database with status `PENDING`, and drops it into a Kafka topic (e.g., `notifications-topic`). It immediately returns `202 Accepted` to the caller.

### B. The Message Broker (Kafka/RabbitMQ)
Using a queue is essential to absorb sudden spikes (e.g., a batch job sending 1 million marketing emails at 10 AM). The queue acts as a buffer so the API and Workers don't crash.

### C. The Worker Nodes
We have specific worker microservices for each channel (Email workers, SMS workers).
*   **User Preferences & Anti-Spam:** Before sending, the worker checks the User DB to see if `user_id 123` has opted out of SMS. If so, it drops the message.
*   **Rate Limiting:** Workers ensure they don't exceed the third-party API rate limits (e.g., Twilio only allows X messages/second).
*   **Templating Engine:** The worker fetches the `OTP_LOGIN` template from the DB, replaces `{{otp}}` with `"1234"`, and constructs the final text.
*   **Third-Party Dispatch:** It makes the API call to Twilio or SendGrid.

## 4. Database Design
*   **Notification Log DB (Cassandra or PostgreSQL):** Stores the history of every notification for auditing and debugging.
    *   `id`, `user_id`, `type`, `status` (PENDING, SENT, FAILED, BOUNCED), `created_at`.
*   **Template DB (NoSQL/SQL):** Stores the HTML/Text templates.

## 5. Retry and Failure Handling
What if Twilio's API is down?
*   Worker catches the HTTP 500 error from Twilio.
*   It places the message into a **Retry Queue** (or Dead Letter Queue - DLQ) with exponential backoff (retry in 1 min, then 5 mins, then 15 mins).
*   After 5 retries, mark it as `FAILED` in the database.

## 6. Follow-up Questions for Candidate
1.  How do you ensure a user doesn't receive the exact same SMS twice if a worker crashes midway? (Idempotency keys and checking the DB before sending).
2.  How would you design an analytics dashboard for this? (Use an ELK stack: Elasticsearch, Logstash, Kibana by piping worker logs directly into Logstash).
