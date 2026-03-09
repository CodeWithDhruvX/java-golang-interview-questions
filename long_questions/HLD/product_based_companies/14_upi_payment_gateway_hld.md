# High-Level Design (HLD): UPI Payment Gateway System

## Problem Statement
Design a highly scalable and secure UPI payment gateway system that processes millions of transactions daily, integrates with multiple banks and NPCI, and ensures regulatory compliance for the Indian market.

## Requirements Analysis
*   **Scale:** 100M+ daily transactions, 10K+ TPS peak load
*   **Latency:** <200ms response time for authorization
*   **Availability:** 99.99% uptime (financial system requirements)
*   **Security:** PCI DSS compliance, end-to-end encryption
*   **Compliance:** RBI regulations, NPCI guidelines, data localization
*   **Settlement:** T+1 settlement to merchant accounts

## System Architecture

### 1. Core Components Architecture
```
Merchant Apps → API Gateway → Payment Gateway → NPCI → Banks
                    ↓              ↓
                Auth Service → Settlement Engine
                    ↓              ↓
                Risk Engine → Reconciliation System
```

**Key Components:**
- **API Gateway:** Request routing, rate limiting, authentication
- **Payment Orchestrator:** Transaction coordination and flow management
- **Bank Integration Layer:** Adapter pattern for multiple bank APIs
- **NPCI Interface:** UPI protocol implementation
- **Risk Management:** Fraud detection and prevention
- **Settlement Engine:** Daily reconciliation and fund transfers

### 2. Transaction Flow Architecture

#### UPI Payment Flow
```
1. Merchant Initiation → Payment Gateway
2. Customer Authentication → Bank/NPCI
3. Authorization Response → Payment Gateway
4. Transaction Confirmation → Merchant
5. Settlement Processing → NPCI → Banks
```

**Detailed Flow:**
1. **Initiation:** Merchant creates payment request with amount, VPA
2. **Routing:** Gateway routes to appropriate bank based on VPA
3. **Authentication:** Customer authenticates via UPI app/bank
4. **Authorization:** Bank validates funds and authorizes transaction
5. **Confirmation:** Gateway receives success/failure response
6. **Settlement:** Batch processing for fund transfers

### 3. Database Architecture

#### Primary Data Stores
- **PostgreSQL (Cluster):** Transaction records, user accounts (ACID compliance)
- **Redis Cluster:** Session management, rate limiting, caching
- **MongoDB:** Audit logs, transaction history (time-series)
- **Elasticsearch:** Search, analytics, monitoring data
- **Object Storage:** Reports, statements, archived data

#### Data Partitioning Strategy
- **Transaction Data:** Shard by transaction_id (UUID)
- **User Data:** Shard by customer_id (consistent hashing)
- **Merchant Data:** Geo-partition by region/state
- **Audit Data:** Time-series partition by date

### 4. Security Architecture

#### Multi-Layer Security
```
Network Layer → Application Layer → Data Layer → Integration Layer
     ↓              ↓              ↓              ↓
  WAF/DMZ        Encryption    Tokenization    API Security
```

**Security Measures:**
- **Encryption:** AES-256 for data at rest, TLS 1.3 for in-transit
- **Tokenization:** Replace sensitive data with non-sensitive equivalents
- **Authentication:** OAuth 2.0, JWT tokens, 2FA enforcement
- **Authorization:** RBAC, least privilege principle
- **Audit Logging:** Complete transaction audit trail

#### PCI DSS Compliance
- **Cardholder Data:** Never store full card numbers
- **Encryption:** Strong cryptography for sensitive data
- **Access Control:** Strict authentication and authorization
- **Monitoring:** Real-time security monitoring and alerting
- **Testing:** Regular penetration testing and vulnerability scans

### 5. Integration Architecture

#### Bank Integration Patterns
```
Payment Gateway → Adapter Layer → Bank APIs
                    ↓
                Protocol Translation
                    ↓
                Error Handling & Retry
```

**Integration Strategies:**
- **Adapter Pattern:** Standardize different bank API formats
- **Circuit Breaker:** Prevent cascade failures from bank outages
- **Retry Logic:** Exponential backoff with jitter
- **Dead Letter Queue:** Handle failed transactions
- **Health Monitoring:** Real-time bank service status

#### NPCI Integration
- **UPI Protocol:** Implement NPCI UPI 2.0 specifications
- **API Endpoints:** Collect, Verify, Pay, Status Check
- **Webhook Handling:** Asynchronous status updates
- **Compliance:** Regulatory reporting and data formats

### 6. Risk Management System

#### Fraud Detection Pipeline
```
Transaction Input → Risk Scoring → Decision Engine → Action
                      ↓
                Machine Learning Models
                      ↓
                Rule Engine
```

**Risk Assessment Factors:**
- **Transaction Patterns:** Amount, frequency, timing
- **User Behavior:** Device, location, historical patterns
- **Merchant Risk:** Category, transaction history
- **Geolocation:** IP address, device location
- **Velocity Checks:** Multiple transactions in short time

#### Machine Learning Models
- **Anomaly Detection:** Identify unusual transaction patterns
- **Behavioral Analysis:** User-specific spending patterns
- **Network Analysis:** Transaction relationship mapping
- **Real-time Scoring:** Sub-millisecond risk assessment

### 7. Settlement Engine

#### Daily Settlement Flow
```
1. Transaction Aggregation → Batch Processing
2. NPCI Reconciliation → Fund Transfer
3. Merchant Settlement → Account Credit
4. Fee Deduction → Revenue Recognition
5. Reporting → Financial Statements
```

**Settlement Components:**
- **Batch Processor:** Aggregate transactions by merchant
- **Reconciliation Engine:** Match transactions with NPCI records
- **Fund Transfer:** NEFT/RTGS/IMPS integration
- **Fee Calculator:** Commission and processing fees
- **Reporting Engine:** Daily/monthly settlement reports

### 8. Monitoring & Observability

#### Real-time Monitoring
```
Metrics Collection → Processing → Alerting → Dashboard
       ↓               ↓          ↓         ↓
   Prometheus    Grafana   AlertManager  Custom UI
```

**Key Metrics:**
- **Transaction Metrics:** Volume, success rate, response time
- **Business Metrics:** Revenue, settlement amounts, merchant activity
- **Technical Metrics:** CPU, memory, network, database performance
- **Security Metrics:** Failed authentications, fraud attempts

#### Distributed Tracing
- **Request Flow:** End-to-end transaction tracing
- **Service Dependencies:** Microservice interaction mapping
- **Performance Analysis:** Bottleneck identification
- **Error Tracking:** Root cause analysis

### 9. High Availability Design

#### Redundancy Architecture
```
Primary Region (Mumbai)          Backup Region (Chennai)
       ↓                               ↓
  Load Balancer                Load Balancer
       ↓                               ↓
  Application Servers        Application Servers
       ↓                               ↓
  Database Cluster           Database Cluster
```

**HA Features:**
- **Multi-AZ Deployment:** Primary and backup availability zones
- **Database Replication:** Synchronous and asynchronous replication
- **Failover Automation:** Automatic service failover
- **Health Checks:** Comprehensive service health monitoring
- **Graceful Degradation:** Limited functionality during outages

### 10. Performance Optimization

#### Caching Strategy
```
L1: Application Cache (Hot data)
L2: Redis Cluster (Session, rate limiting)
L3: CDN (Static assets, API responses)
```

**Optimization Techniques:**
- **Connection Pooling:** Database connection reuse
- **Async Processing:** Non-blocking I/O operations
- **Batch Processing:** Bulk operations for efficiency
- **Index Optimization:** Database query optimization
- **Load Balancing:** Intelligent traffic distribution

### 11. Regulatory Compliance

#### RBI Compliance Requirements
- **Data Localization:** Store Indian transaction data in India
- **Audit Requirements:** Complete transaction audit trail
- **Reporting:** Regulatory reporting formats and schedules
- **Security Standards:** Mandated security controls
- **KYC/AML:** Know Your Customer and Anti-Money Laundering

#### NPCI Guidelines
- **UPI Compliance:** UPI 2.0 feature implementation
- **Interoperability:** Cross-bank transaction support
- **Settlement Cycles:** T+1 settlement requirements
- **Error Handling:** Standardized error codes and responses
- **Testing:** Certification and compliance testing

### 12. Capacity Planning

#### Traffic Estimation
- **Daily Transactions:** 100M transactions
- **Peak Load:** 10K TPS (peak hours)
- **Concurrent Users:** 1M active users
- **API Requests:** 50K requests/second

#### Infrastructure Sizing
- **Application Servers:** 200+ instances (auto-scaling)
- **Database Nodes:** 20+ nodes (sharded cluster)
- **Redis Cluster:** 50+ nodes (caching layer)
- **Message Queue:** 20+ brokers (async processing)

#### Storage Requirements
- **Transaction Data:** 10TB/day (including metadata)
- **Audit Logs:** 5TB/day (compliance requirements)
- **User Data:** 2TB (profiles, preferences)
- **Reports:** 1TB/month (settlement reports)

## Indian Market Specific Considerations

### 1. UPI Ecosystem
- **Multiple PSPs:** PhonePe, PayTM, Google Pay integration
- **Bank Partnerships:** 150+ bank connections
- **VPA Resolution:** Virtual Payment Address validation
- **App Integration:** Deep linking and app callbacks

### 2. Network Infrastructure
- **Variable Connectivity:** Support for 2G/3G networks
- **Timeout Handling:** Extended timeouts for slow networks
- **Retry Logic:** Intelligent retry for failed transactions
- **Offline Support:** Basic functionality without internet

### 3. Cultural Adaptation
- **Language Support:** Multi-language interfaces
- **Regional Banks:** Integration with regional banks
- **Festival Load:** High volume during festivals
- **Rural Areas:** Limited connectivity considerations

### 4. Business Models
- **Merchant Categories:** Different rates for different categories
- **Small Merchants:** Micro-transaction support
- **Government Payments:** Integration with government schemes
- **International Payments:** Cross-border transaction support

## Interview Success Tips

### Key Discussion Points
1. **Scalability:** How to handle 10x growth in transactions
2. **Security:** PCI DSS compliance implementation
3. **Integration:** Bank adapter pattern and NPCI integration
4. **Performance:** Sub-200ms response time achievement
5. **Reliability:** 99.99% uptime guarantee mechanisms
6. **Compliance:** RBI and NPCI regulatory requirements
7. **Risk Management:** Fraud detection and prevention strategies

### Common Follow-up Questions
- How would you handle a bank outage?
- What's your approach to transaction reconciliation?
- How do you ensure data consistency across systems?
- What happens during network partitions?
- How would you optimize for rural areas?

### Technical Deep Dives
- **Distributed Transactions:** Two-phase commit implementation
- **Event Sourcing:** Audit trail and replay capabilities
- **Circuit Breakers:** Fault tolerance patterns
- **Rate Limiting:** Preventing abuse and ensuring fairness
- **Message Queues:** Reliable message delivery patterns
