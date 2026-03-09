# High-Level Design (HLD): Food Delivery System (Swiggy/Zomato)

## Problem Statement
Design a scalable food delivery system like Swiggy/Zomato that handles millions of orders daily across Indian cities, with real-time tracking, multiple payment methods, and surge pricing.

## Requirements Analysis
*   **Scale:** 10M+ daily orders, 50M+ active users, 100K+ restaurant partners
*   **Real-time:** Live order tracking, driver location updates, ETA calculations
*   **Performance:** <2s response time for restaurant search, <500ms for order placement
*   **Availability:** 99.99% uptime, especially during peak hours (12-2pm, 7-10pm)
*   **Geographic Coverage:** Tier-1 to Tier-3 cities with varying network conditions

## System Architecture

### 1. Microservices Architecture
```
Client Apps → API Gateway → Microservices → Message Queue → Data Stores
```

**Core Microservices:**
- **User Service:** Authentication, profiles, addresses, preferences
- **Restaurant Service:** Menu management, inventory, operating hours
- **Order Service:** Order lifecycle, status tracking, coordination
- **Delivery Service:** Driver management, assignment algorithms, tracking
- **Payment Service:** UPI, cards, wallets, COD processing
- **Search Service:** Restaurant discovery, relevance ranking
- **Notification Service:** Push notifications, SMS, email
- **Analytics Service:** Real-time metrics, business intelligence

### 2. Data Flow Architecture

#### Order Placement Flow
```
1. User browses restaurants → Search Service (Elasticsearch)
2. User selects items → Restaurant Service (Redis cache)
3. User places order → Order Service (MySQL)
4. Payment processing → Payment Gateway (Redis + PostgreSQL)
5. Driver assignment → Delivery Service (Matching Algorithm)
6. Real-time tracking → WebSocket + Location Service
```

### 3. Database Design

#### Primary Databases
- **MySQL (Master-Slave):** Orders, users, transactions (ACID compliance)
- **PostgreSQL:** Restaurant data, menus, analytics (JSON support)
- **Redis:** Session management, caching, real-time data
- **Elasticsearch:** Restaurant search, indexing, autocomplete
- **MongoDB:** Driver location logs, audit trails

#### Data Partitioning Strategy
- **User Data:** Shard by user_id (consistent hashing)
- **Order Data:** Shard by order_date + user_id (temporal locality)
- **Restaurant Data:** Geo-partition by city/region
- **Location Data:** Time-series partition by date

### 4. Caching Strategy

#### Multi-Level Caching
```
L1: Application Cache (Caffeine/Guava) - Hot data
L2: Redis Cluster - Session, user preferences
L3: CDN (CloudFront) - Restaurant images, static assets
```

**Cache Patterns:**
- **Read-Through:** Restaurant menus, user profiles
- **Write-Through:** Order status updates
- **Cache-Aside:** Search results, popular items
- **Write-Behind:** Analytics data, user activity logs

### 5. Real-Time Architecture

#### Location Tracking System
```
Driver App → MQTT Broker → Location Processor → Redis Geo → Client Apps
```

**Components:**
- **MQTT Broker:** High-throughput message handling (100K+ concurrent connections)
- **Location Processor:** Kafka streams for real-time processing
- **Redis Geo:** Geospatial indexing for nearby drivers
- **WebSocket Gateway:** Real-time updates to clients

#### Order Status Updates
```
Order Service → Kafka → Notification Service → Push Gateway → Client Apps
```

### 6. Search & Discovery Architecture

#### Restaurant Search Pipeline
```
User Query → API Gateway → Search Service → Elasticsearch → Ranking Engine → Results
```

**Search Features:**
- **Geo-based Search:** Location-aware restaurant ranking
- **Cuisine Filtering:** Multi-cuisine classification
- **Price Range:** Budget-based filtering
- **Rating & Popularity:** Dynamic ranking algorithms
- **Availability:** Real-time restaurant status

### 7. Delivery Matching Algorithm

#### Driver Assignment Strategy
```
1. Geo-proximity: Drivers within 3km radius
2. Availability: Currently active drivers
3. Capacity: Current order load (<3 orders)
4. Rating: Driver performance score
5. ETA: Estimated pickup time
```

**Algorithm Selection:**
- **Hungarian Algorithm:** Optimal assignment for small sets
- **Greedy with Constraints:** Scalable for large driver pools
- **Machine Learning:** Predictive ETA and matching accuracy

### 8. Payment Architecture

#### Payment Processing Flow
```
Order Service → Payment Service → Payment Gateway → Bank/NPCI → Response
```

**Payment Methods:**
- **UPI Integration:** Direct bank-to-bank transfers
- **Card Processing:** PCI DSS compliant tokenization
- **Wallet Integration:** PayTM, PhonePe, Amazon Pay
- **Cash on Delivery:** Driver payment collection

**Failure Handling:**
- **Idempotency:** Prevent duplicate charges
- **Retry Logic:** Exponential backoff for failures
- **Fallback:** Alternative payment methods
- **Reconciliation:** Daily settlement processing

### 9. Surge Pricing Algorithm

#### Dynamic Pricing Model
```
Base Price = Distance Rate × Distance + Time Rate × Time + Base Fee
Surge Multiplier = f(Demand/Supply Ratio, Weather, Time of Day)
Final Price = Base Price × Surge Multiplier
```

**Factors Considered:**
- **Demand/Supply Ratio:** Orders vs available drivers
- **Weather Conditions:** Rain, extreme heat
- **Peak Hours:** Lunch, dinner times
- **Events:** Concerts, matches, festivals
- **Location:** Premium areas, airports

### 10. Scalability & Performance

#### Horizontal Scaling Strategy
- **Stateless Services:** Auto-scaling based on CPU/memory
- **Database Sharding:** Read replicas for scaling
- **Load Balancing:** Round-robin with health checks
- **Circuit Breakers:** Fault isolation between services

#### Performance Optimization
- **Database Indexing:** Composite indexes for common queries
- **Connection Pooling:** Reduce database connection overhead
- **Async Processing:** Kafka for non-blocking operations
- **CDN Integration:** Static content delivery

### 11. Monitoring & Observability

#### Key Metrics
- **Business Metrics:** Order volume, delivery time, customer satisfaction
- **Technical Metrics:** API response times, error rates, system health
- **Infrastructure Metrics:** CPU, memory, network utilization

#### Monitoring Stack
- **Prometheus:** Metrics collection and alerting
- **Grafana:** Visualization and dashboards
- **ELK Stack:** Log aggregation and analysis
- **Jaeger:** Distributed tracing

### 12. Disaster Recovery

#### High Availability Design
- **Multi-AZ Deployment:** Primary and backup regions
- **Database Replication:** Cross-region replication
- **Failover Mechanism:** Automatic service failover
- **Data Backup:** Regular snapshots and point-in-time recovery

#### Business Continuity
- **Graceful Degradation:** Limited functionality during outages
- **Offline Mode:** Basic ordering without tracking
- **Communication:** Status updates for users and partners

## Indian Market Specific Considerations

### 1. Network Conditions
- **2G/3G Support:** Lightweight mobile apps
- **Offline Mode:** Basic functionality without internet
- **Progressive Loading:** Critical content first

### 2. Payment Preferences
- **UPI Dominance:** 60%+ transactions via UPI
- **Cash on Delivery:** Still significant in Tier-2/3 cities
- **Wallet Integration:** Multiple wallet options

### 3. Geographic Diversity
- **Address Systems:** Support for Indian address formats
- **Language Support:** Multi-language interfaces
- **Regional Cuisines:** Local food preferences

### 4. Regulatory Compliance
- **FSSAI Regulations:** Food safety compliance
- **GST Integration:** Tax calculation and reporting
- **Data Localization:** Data residency requirements

## Capacity Estimation

### Traffic Estimates
- **Daily Orders:** 10M orders
- **Peak Hour Traffic:** 1M orders/hour
- **Concurrent Users:** 500K active users
- **API Requests:** 50K requests/second

### Storage Requirements
- **Order Data:** 1TB/day (including metadata)
- **Location Data:** 100GB/day (driver tracking)
- **User Data:** 500GB (profiles, preferences)
- **Restaurant Data:** 200GB (menus, images)

### Infrastructure Sizing
- **Application Servers:** 1000+ instances (auto-scaling)
- **Database Nodes:** 50+ nodes (sharded cluster)
- **Redis Cluster:** 100+ nodes (caching layer)
- **Kafka Brokers:** 20+ brokers (message processing)

## Interview Success Tips

### Key Discussion Points
1. **Scalability:** How to handle 10x growth in orders
2. **Real-time Processing:** WebSocket vs polling strategies
3. **Database Design:** SQL vs NoSQL for different use cases
4. **Caching Strategy:** Multi-level caching approach
5. **Microservices:** Service boundaries and communication
6. **Payment Processing:** Handling Indian payment ecosystem
7. **Location Tracking:** Efficient geospatial algorithms

### Common Follow-up Questions
- How would you handle driver fraud?
- What happens during network outages?
- How do you ensure order consistency?
- How would you optimize for Tier-2 cities?
- What's your approach to A/B testing?

### Technical Deep Dives
- **Consistency Models:** Eventual vs strong consistency
- **CAP Theorem:** Trade-offs in distributed systems
- **Rate Limiting:** Preventing abuse and ensuring fairness
- **Circuit Breakers:** Fault tolerance patterns
- **Event Sourcing:** Audit trails and replay capabilities
