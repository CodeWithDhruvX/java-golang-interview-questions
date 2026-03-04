# System Design (HLD) - Service Based Companies

This section deals with System Design and Architecture questions from the perspective of Service-Based companies (e.g., Accenture, Cognizant, TCS, Capgemini, Infosys) or for enterprise-scale B2B application roles.

While FAANG style system design focuses on massive horizontally scalable public-facing networks (Millions of concurrent users, NoSQL, eventual consistency), **Enterprise / Service-Based HLD focuses heavily on the integration of disparate internal systems, security, ACID properties, API architecture, and reliable logging/monitoring.**

## Topics Covered:
1.  [Centralized Notification System](01_hld_notification_system.md)
2.  [Centralized Logging & Monitoring (ELK Stack)](02_hld_logging_monitoring.md)
3.  [Payment Gateway Integration Layer](03_hld_payment_gateway.md)
4.  [API Gateway & Service Mesh](04_hld_api_gateway.md)
5.  [Content Management System (CMS) / Document Platform](05_hld_cms_system.md)
6.  [Distributed Job Scheduler](06_hld_distributed_job_scheduler.md)
7.  [Time Series Database (Metrics System)](07_hld_metrics_time_series.md)

## Success Criteria for Enterprise System Design Interviews
1.  **Microservices Patterns:** Clear understanding of API Gateways, Service Discovery (Eureka/Consul), Circuit Breakers, and the Saga Pattern.
2.  **Consistency First:** Relational databases (PostgreSQL/Oracle) often reign supreme here due to strict financial or compliance requirements. You must understand ACID properties.
3.  **Observability:** How do you debug 50 decoupled microservices when something fails? (ELK Stack, Zipkin/Jaeger Trace IDs, Prometheus).
4.  **Security & Identity:** Solid understanding of OAuth2.0, JWT, API Rate Limiting, and WAFs.
5.  **Handling Third-Party Integrations:** Using Kafka, Webhooks, Idempotency keys, and retry queues when interacting with external APIs (like Twilio, Stripe, Salesforce).
