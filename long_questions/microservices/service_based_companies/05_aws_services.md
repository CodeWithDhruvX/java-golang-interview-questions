# ☁️ Microservices — AWS Services Deep Dive

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** TCS (AWS SkillBuild roles), Infosys (AWS certified projects), Wipro, Cognizant, Accenture, Capgemini — and all product companies on AWS (Razorpay, Swiggy, CRED, Zepto)

---

## Q1. What is SQS? How is it different from Kafka?

**"Amazon SQS (Simple Queue Service) is a fully managed message queuing service. You don't manage servers, brokers, or partitions — AWS handles all of it.**

**When I use SQS:**
- Fire-and-forget tasks: send email, process image, trigger a report
- Decoupling microservices without operational overhead
- When I don't need event replay (messages are deleted after consumption)
- When the team doesn't have Kafka expertise

**Core SQS Concepts:**

| Concept | Description |
|---------|-------------|
| **Standard Queue** | At-least-once delivery, best-effort ordering. Highest throughput. |
| **FIFO Queue** | Exactly-once processing, strict ordering. 3000 TPS limit. |
| **Visibility Timeout** | After a consumer receives a message, it's hidden from others for X seconds. If not deleted in time, it becomes visible again (retry). Default: 30s. |
| **Dead Letter Queue (DLQ)** | Messages that fail processing N times are moved here for investigation. |
| **Long Polling** | Consumer waits up to 20 seconds for a message, reducing empty responses and cost. |

```java
// Spring Boot — Sending to SQS (AWS SDK v2)
@Service
public class NotificationService {

    @Autowired private SqsAsyncClient sqsClient;

    private static final String QUEUE_URL = "https://sqs.ap-south-1.amazonaws.com/123/notification-queue";

    public void sendNotification(NotificationRequest request) {
        SendMessageRequest sendMsg = SendMessageRequest.builder()
            .queueUrl(QUEUE_URL)
            .messageBody(objectMapper.writeValueAsString(request))
            .messageGroupId("notifications")   // Required for FIFO queues
            .messageDeduplicationId(request.getOrderId()) // Idempotency for FIFO
            .build();

        sqsClient.sendMessage(sendMsg);
    }
}

// Consuming from SQS using Spring Cloud AWS
@SqsListener("notification-queue")
public void onNotification(NotificationRequest request) {
    emailService.send(request.getEmail(), request.getMessage());
    // Returning normally = SQS auto-deletes the message (acknowledges)
    // Throwing exception = message becomes visible again after Visibility Timeout
}
```

**SQS vs Kafka — When to use which:**

| | SQS | Kafka |
|--|-----|-------|
| **Managed** | Fully managed (AWS) | Self-managed or Confluent Cloud |
| **Event replay** | ❌ No (messages deleted) | ✅ Yes (configurable retention) |
| **Ordering** | Limited (FIFO only) | Per-partition ordering |
| **Consumers** | Each message consumed by ONE consumer | Multiple consumer groups, each sees all messages |
| **Throughput** | Very high (Standard), 3K TPS (FIFO) | Near unlimited with partitions |
| **Cost** | Pay per request (~$0.40/million) | Fixed cost of cluster |
| **Best for** | Task queues, decoupling, simple workloads | Event streaming, audit logs, multi-consumer fan-out |"

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS (AWS projects), Infosys, and at product companies choosing between SQS and Kafka. The SQS vs Kafka distinction is a classic interview question.

---

## Q2. What is SNS? Explain the SNS + SQS Fan-out Pattern.

**"Amazon SNS (Simple Notification Service) is a pub/sub service. One SNS Topic → multiple subscribers. A single message published to SNS is delivered to all subscribed endpoints simultaneously (fan-out).**

**SNS Subscribers can be:**
- SQS Queues — for reliable, durable async processing
- Lambda Functions — for event-triggered serverless
- HTTP/HTTPS endpoints
- Email/SMS — for alerts

**The Classic Fan-out Pattern (SNS + SQS):**

```
                    OrderService
                        │
                        │ Publish "OrderPlaced"
                        ▼
              ┌─────────────────┐
              │   SNS Topic     │
              │  (order-events) │
              └─────────────────┘
          ┌──────────┼──────────┐
          ▼          ▼          ▼
   ┌──────────┐ ┌──────────┐ ┌──────────┐
   │  SQS     │ │  SQS     │ │  SQS     │
   │Inventory │ │ Shipping │ │ Notif.   │
   │  Queue   │ │  Queue   │ │  Queue   │
   └──────────┘ └──────────┘ └──────────┘
        │              │            │
   InventoryService ShipService NotifService
        consumes       consumes    consumes
```

**Why SQS as the subscriber (not direct service endpoint)?**
- If the service is down, SQS holds the message until it's back
- SQS provides retry, DLQ, and visibility timeout
- Without SQS buffer, if the endpoint is down, SNS delivery fails and message is lost

```java
// Publishing to SNS in Spring Boot
@Service
public class OrderEventPublisher {

    @Autowired private SnsAsyncClient snsClient;

    private static final String ORDER_TOPIC_ARN = "arn:aws:sns:ap-south-1:123:order-events";

    public void publishOrderPlaced(OrderPlacedEvent event) {
        PublishRequest request = PublishRequest.builder()
            .topicArn(ORDER_TOPIC_ARN)
            .message(objectMapper.writeValueAsString(event))
            .subject("OrderPlaced")
            .messageAttributes(Map.of(
                "eventType", MessageAttributeValue.builder()
                    .dataType("String")
                    .stringValue("ORDER_PLACED")
                    .build()
            ))
            .build();

        snsClient.publish(request);
        // Message delivered to ALL SNS subscribers simultaneously
    }
}
```

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** TCS AWS projects, Infosys cloud-native roles, Accenture. The SNS+SQS fan-out pattern is the standard AWS architecture for event-driven microservices.

---

## Q3. ALB vs NLB — When do you use each?

**"Both are AWS load balancers, but they operate at different OSI layers with very different use cases:**

| | ALB (Application LB) | NLB (Network LB) |
|--|---------------------|-------------------|
| **OSI Layer** | Layer 7 (Application — HTTP/HTTPS) | Layer 4 (Transport — TCP/UDP) |
| **Routing** | By URL path, hostname, headers, query params, cookies | By IP + Port only |
| **Latency** | Slightly higher (inspects HTTP packets) | Ultra-low (~100µs) |
| **Protocol** | HTTP, HTTPS, WebSocket, gRPC | TCP, UDP, TLS |
| **Target types** | EC2, ECS tasks, Lambda, K8s pods (via Ingress) | EC2, IP addresses |
| **SSL termination** | Yes — decrypts at load balancer | Yes, OR pass-through to backend |
| **Connection** | Opens new connection to backend | Preserves source IP end-to-end |

**Use ALB when:**
- Routing different paths to different services: `/api/orders` → OrderService, `/api/payments` → PaymentService
- Need to inspect HTTP headers (route premium users to faster instances)
- Hosting multiple services on one load balancer (host-based routing: `api.example.com` vs `admin.example.com`)
- WebSocket connections
- **Kubernetes Ingress** — AWS ALB Ingress Controller is standard for EKS

**Use NLB when:**
- Need static IP or Elastic IP for whitelisting by clients (banks, compliance)
- Ultra-low latency applications (financial trading, gaming)
- Non-HTTP protocols (game servers using UDP, IoT using MQTT)
- Need to preserve the original client IP at the backend (ALB replaces it with its own IP)

**Typical Microservices Architecture:**
```
Internet
    │
    ▼
┌──────────────────┐
│  ALB             │  path-based routing to services
│  (api.example.com)│
└──────────────────┘
    │          │
    ▼          ▼
OrderService  PaymentService

          ┌──────────────────┐
          │  NLB             │  for internal service with static IP
          │  (internal,      │  requirement (e.g., B2B partner connections)
          │   static IP)     │
          └──────────────────┘
```"

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Wipro (AWS infrastructure roles), Infosys. Standard AWS architecture knowledge expected for any backend engineer deploying microservices on AWS.

---

## Q4. ECS vs EKS — When do you use each?

**"Both run Docker containers on AWS, but with fundamentally different operational philosophies:**

| | ECS (Elastic Container Service) | EKS (Elastic Kubernetes Service) |
|--|--------------------------------|----------------------------------|
| **Orchestrator** | AWS proprietary | Kubernetes (open-source standard) |
| **Learning curve** | Low — AWS Console/CLI familiar | High — full K8s knowledge needed |
| **Portability** | AWS-only | Multi-cloud (K8s APIs are standard) |
| **Managed control plane** | Fully managed (free) | Managed, but you pay for control plane ($0.10/hr ~$73/month) |
| **Serverless containers** | ✅ Fargate — zero server management | ✅ Fargate on EKS (more complex) |
| **Ecosystem** | AWS-native integrations (CloudWatch, IAM) | Huge K8s ecosystem (Helm, ArgoCD, Istio) |
| **Service mesh** | AWS App Mesh (proprietary) | Istio, Linkerd (portable, feature-rich) |
| **Auto-scaling** | Service Auto Scaling (simple) | HPA, VPA, KEDA, Karpenter (powerful) |

**Use ECS when:**
- Small team, quick time-to-market (startup)
- Don't need K8s ecosystem (Helm charts, custom controllers, Istio)
- Want simplest possible operations — just run containers
- Starting fresh on AWS with limited DevOps experience

**Use EKS when:**
- Large team with K8s expertise (or plan to build it)
- Need portability — might migrate to GKE or Azure AKS later
- Need advanced scheduling, custom operators, or service mesh
- Regulatory or architectural requirements demand K8s features (pod security policies, network policies)

**ECS Fargate vs EKS Fargate:**
- ECS Fargate: Simplest: `aws ecs run-task` → container running, no node management ever
- EKS Fargate: K8s API but no node management — good middle ground if you know K8s but hate managing nodes

**My recommendation for interview answers:**
*'For a new project with a small team, I'd start with ECS Fargate — it's run-and-forget. As the team scales, adopts GitOps (ArgoCD), and service mesh becomes necessary, migrating to EKS is justified. The decision is driven by team maturity and operational needs, not technology preference.'*"

#### 🏢 Company Context
**Level:** 🟡 Mid | **Asked at:** Infosys (AWS division), TCS (AWS projects), Accenture, and any company with significant AWS footprint. This is becoming a standard mid-level question.

---

## Q5. How do SQS, SNS, and EventBridge fit together? What is EventBridge?

**"These three services form a complete event-driven architecture toolkit on AWS:**

```
SQS  = Message Queue   (point-to-point, durable, task processing)
SNS  = Pub/Sub         (one-to-many fan-out)
EventBridge = Event Bus (sophisticated event routing with rules, schema registry, SaaS integration)
```

**Amazon EventBridge (formerly CloudWatch Events):**
- A serverless event bus that routes events from AWS services, custom apps, or SaaS apps
- Rules can filter events and route them to SQS, SNS, Lambda, Step Functions, etc.
- Built-in Schema Registry — discovers and documents event shapes
- Partner integrations: Shopify, Zendesk, Datadog events can flow directly into your EventBridge

```
Example Rule:
IF event.source == "order-service"
AND event.detail.status == "PAYMENT_FAILED"
AND event.detail.amount > 10000
THEN → route to → fraud-investigation Lambda

IF event.source == "order-service"
AND event.detail.status == "PAYMENT_FAILED"
THEN → route to → retry-payment SQS Queue
```

**When to use each:**

| Use Case | Service |
|----------|---------|
| Background task processing (send email, resize image) | SQS |
| Fan-out: one event → multiple consumers | SNS → SQS |
| Complex routing based on event content | EventBridge |
| AWS service events (S3 upload → process file) | EventBridge + Lambda |
| Real-time streaming, event replay, multi-consumer | Kafka (MSK) |

**Full Example Architecture:**
```
Order Placed
    │
    ▼ (custom event to EventBridge)
EventBridge
    │
    ├─ [amount > 50000] → SQS Manual Review Queue → Human Agent
    ├─ [normal] → SNS OrderTopic
    │                   ├─ SQS InventoryQueue → InventoryService
    │                   ├─ SQS ShippingQueue → ShippingService
    │                   └─ Lambda → Real-time Analytics
    └─ [fraud signals] → Step Functions → Fraud Investigation Workflow
```"

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** TCS digital AWS roles, Infosys cloud architecture rounds, and product companies doing full AWS-native architecture. Understanding the distinctions between these three services shows real cloud architecture maturity.
