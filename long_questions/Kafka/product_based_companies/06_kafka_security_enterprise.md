# 🏗️ Kafka — Enterprise Security & Access Control

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Goldman Sachs, JPMorgan, Amazon, PayPal, Razorpay, Zepto

---

## Q1. Walk me through a fully secured Kafka cluster setup for a banking application.

"A production-grade banking Kafka cluster requires all three layers of security working together: **Encryption**, **Authentication**, and **Authorization**.

**Layer 1 — Encryption in Transit (TLS):**
Every connection — producers to brokers, consumers to brokers, brokers to each other (inter-broker replication) — must be encrypted using TLS. This prevents man-in-the-middle attacks and packet sniffing.

```properties
# broker server.properties
listeners=SSL://0.0.0.0:9093
advertised.listeners=SSL://broker1.bank.internal:9093
ssl.keystore.location=/etc/kafka/ssl/kafka.server.keystore.jks
ssl.keystore.password=<secret>
ssl.truststore.location=/etc/kafka/ssl/kafka.server.truststore.jks
ssl.truststore.password=<secret>
ssl.client.auth=required      # enforces mTLS — client must present a cert
```

**Layer 2 — Mutual TLS (mTLS) Authentication:**
Both sides (client AND server) present certificates. The broker validates the client certificate against its truststore. The client's `CN` (Common Name) from the cert becomes its Kafka principal: `User:CN=payments-service`.

**Layer 3 — ACL Authorization:**
After identity is established, ACLs define what that identity can do:
```bash
# Allow payments-service to ONLY write to the payments topic
kafka-acls.sh --add \
  --allow-principal User:CN=payments-service \
  --operation Write \
  --topic payments
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Natively consumes JKS or PEM stores provided via `spring.kafka.ssl.*` properties to automatically handle both Transport TLS and mTLS layers without explicit Java code.
* **Golang:** `kafka-go` requires populating a `tls.Config` struct with loaded `x509.Certificate` key pairs and providing it to the `Dialer` or `Transport`, rather than relying purely on config strings.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** Goldman Sachs, JPMorgan, PayPal — banking and fintech companies enforce mTLS as a non-negotiable security baseline. Candidates are expected to demonstrate end-to-end setup knowledge.

#### Indepth
**Certificate Rotation Zero Downtime:** When TLS certificates expire (typically yearly), rotating them on a live cluster without downtime requires a rolling restart strategy. The new certificate is added to the truststore FIRST (so both old and new certs are trusted simultaneously). Each broker is then restarted one at a time. Only after all brokers are running on the new cert is the old cert removed from truststores.

---

## Q2. Explain SASL mechanisms in Kafka: PLAIN, SCRAM, and OAUTHBEARER. When do you use each?

"SASL (Simple Authentication and Security Layer) provides pluggable authentication without requiring certificates for every client.

**1. SASL/PLAIN:**
A simple username and password are sent in clear text (so MUST be used with TLS). Credentials are stored in the broker's `jaas.conf` as static config — no runtime updates possible without a broker restart.
```
Use case: Simple internal dev/staging environments. Never in production without change management for credential rotation.
```

**2. SASL/SCRAM (Salted Challenge Response Authentication Mechanism):**
Passwords are stored in ZooKeeper/KRaft as salted, iterated hashes (SHA-256 or SHA-512). Credentials can be created, updated, and deleted at runtime without broker restarts.
```bash
# Create a SCRAM user at runtime
kafka-configs.sh --alter \
  --entity-type users --entity-name reporting-service \
  --add-config 'SCRAM-SHA-256=[iterations=8192,password=s3cr3t]'
```
```
Use case: Internal microservices authentication where username/password model is acceptable and zero-downtime credential rotation is needed.
```

**3. SASL/OAUTHBEARER:**
Kafka clients obtain a short-lived JWT token from a corporate Identity Provider (IdP) like Keycloak or Okta via OAuth2 client credentials flow. The token is presented to the broker, which validates it via the IdP's public key.
```
Use case: Enterprise environments with a centralized identity platform where service credentials are managed through IAM rather than Kafka config files.
```"

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Handled extensively by passing JAAS (Java Authentication and Authorization Service) configurations directly into the underlying Spring application properties.
* **Golang:** `confluent-kafka-go` utilizes `sasl.mechanism` mapped directly in the client map. For `kafka-go`, the `Dialer` accepts a `SASLMechanism` interface, with native support for PLAIN, SCRAM, and external plugins for OAuth.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon, Razorpay, Zepto — evaluating security maturity and understanding of the progression from static credentials to dynamic token-based identity.

#### Indepth
**SASL + TLS Together:** SASL provides identity; TLS provides encryption. They are independent. SASL/PLAIN over raw TCP sends the password as base64 (easily decoded). Always use `SASL_SSL` as the listener protocol, never `SASL_PLAINTEXT` in production.

---

## Q3. How do you implement audit logging and govern topic access in a large multi-team Kafka cluster?

"In enterprises with hundreds of topics and dozens of teams, plain ACLs become unmanageable. A layered governance model is required.

**Authorizer-Level Audit Logging:**
Kafka's default ACL authorizer logs every authorization decision to the broker log at `WARN` level for denials. For production auditing, replace the default authorizer with a custom one (or use Confluent's implementation) that emits structured audit events to a dedicated `_kafka_audits` topic. Each event captures: `principal`, `operation`, `resource`, `result`, `timestamp`, `clientId`, `clientHost`.

**Topic Governance Model:**

```
Namespace: team.domain.topic_name
Example:  payments.transactions.order-events
          payments.transactions.refund-events  
          logistics.tracking.shipment-events
```

Each team owns a namespace. ACLs are assigned at the prefix level:
```bash
# Grant payments-team full access to all their topics
kafka-acls.sh --add \
  --allow-principal User:payments-service \
  --operation All \
  --topic payments. \
  --resource-pattern-type prefixed
```

**Confluent RBAC (Role-Based Access Control):**
Instead of raw ACLs, Confluent Platform's RBAC assigns pre-defined roles (`DeveloperRead`, `DeveloperWrite`, `ResourceOwner`) to users or groups managed by LDAP. This integrates Kafka security with existing corporate directory services (Active Directory), making access management auditable through standard IAM tooling."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot & Golang:** Cross-cluster RBAC and namespace isolation are entirely broker-side. The Go/Java application interacts strictly via its authenticated Principal with whatever prefixes it has access to. If the connection hits `TopicAuthorizationException`, it means the client attempted operation outside its authenticated bounds.

#### 🏢 Company Context
**Level:** 🟣 Architect | **Asked at:** JPMorgan, Amazon — large organizations with GDPR, SOC2, and PCI-DSS compliance requirements are expected to have governance mechanisms beyond simple ACLs.

#### Indepth
**Schema-Level Access Control:** Beyond topic ACLs, enterprise environments also control who can register or evolve schemas in the Schema Registry. A fintech might allow `payment-producer` to register new Avro schemas for `payments.*` subjects but restrict breaking schema changes (`FULL` compatibility) to reviewed, CI-gated deployments only.
---
