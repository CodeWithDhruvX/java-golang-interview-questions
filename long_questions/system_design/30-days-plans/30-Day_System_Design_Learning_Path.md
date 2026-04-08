# 30-Day System Design Learning Path

> **Complete roadmap from beginner to interview-ready system designer**
> 
> **Duration:** 30 Days | **Level:** Beginner to Advanced | **Target:** Product & Service Company Interviews

---

## **Phase 1: Introduction (Day 1-4)**

### **Day 1: What is System Design?**
**Topics:**
- High-level overview of system design
- Why System Design matters in interviews
- Components of systems (databases, caches, load balancers, CDNs)

**Reading Materials:**
- [Basic Concepts - Questions 1-10](../theory/01_Basic_Concepts.md)
- Focus on: "What is system design?" and "HLD vs LLD"

**Hands-on Exercise:**
- Draw a simple system diagram for a personal blog
- Identify components: web server, database, CDN

**Practice Question:**
- "Explain system design in your own words" (See interview format in 01_Basic_Concepts.md)

---

### **Day 2: System Design in Interviews**
**Topics:**
- Interviewer expectations
- How to structure your answer (4-step framework)
- Common frameworks (requirements clarification, estimation, design, deep dive)

**Reading Materials:**
- [Interview Strategy - Advanced](../theory/14_Interview_Strategy_Advanced.md)
- Focus on: Interview frameworks and communication patterns

**Hands-on Exercise:**
- Practice the 4-step framework with a simple problem (design a URL shortener)
- Record yourself explaining the approach

**Practice Question:**
- Mock interview: "Design a simple to-do list application"

---

### **Day 3: High-Level vs Low-Level Design**
**Topics:**
- HLD concepts: services, databases, APIs
- LLD concepts: classes, schemas, algorithms
- When to use each approach

**Reading Materials:**
- [Basic Concepts - Questions 1-10](../theory/01_Basic_Concepts.md)
- Focus on: "Difference between high-level and low-level design"

**Hands-on Exercise:**
- Create HLD diagram for a chat application
- Design LLD for the message entity (class diagram)

**Practice Question:**
- "When would you choose HLD over LLD in an interview?"

---

### **Day 4: How to Think Like a System Designer**
**Topics:**
- Trade-offs (performance vs cost, consistency vs availability)
- Design decision frameworks
- Common pitfalls to avoid

**Reading Materials:**
- [Architecture Patterns](../theory/05_Architecture_Patterns.md)
- Focus on: Trade-offs and design patterns

**Hands-on Exercise:**
- Analyze 3 different approaches to caching
- List pros/cons for each approach

**Practice Question:**
- "What are the trade-offs between SQL and NoSQL databases?"

---

## **Phase 2: Fundamentals (Day 5-13)**

### **Day 5: Scalability, Latency, Throughput**
**Topics:**
- Vertical vs Horizontal scaling
- Latency metrics and optimization
- Throughput measurement and improvement

**Reading Materials:**
- [Scalability & Availability](../theory/06_Scalability_Availability.md)
- Focus on: Scaling strategies and performance metrics

**Hands-on Exercise:**
- Calculate QPS for a system with 1M users
- Design scaling strategy for an e-commerce site

**Practice Question:**
- "How would you scale a web application from 1000 to 1M users?"

---

### **Day 6: Client-Server Architecture**
**Topics:**
- Protocols (HTTP, TCP, WebSockets)
- Communication patterns
- API design principles

**Reading Materials:**
- [Networking Protocols](../theory/11_Networking_Protocols.md)
- Focus on: Client-server communication patterns

**Hands-on Exercise:**
- Design API endpoints for a user management system
- Choose appropriate protocols for different use cases

**Practice Question:**
- "When would you use WebSockets vs HTTP?"

---

### **Day 7: HTTP, REST, RPC**
**Topics:**
- REST API design principles
- gRPC and GraphQL
- API versioning strategies

**Reading Materials:**
- [Networking Protocols](../theory/11_Networking_Protocols.md)
- Focus on: API design and communication protocols

**Hands-on Exercise:**
- Design REST API for a blogging platform
- Compare REST vs GraphQL for a mobile app

**Practice Question:**
- "What are the advantages of GraphQL over REST?"

---

### **Day 8: Load Balancers**
**Topics:**
- L4 vs L7 load balancers
- Algorithms (Round Robin, Weighted, Sticky Session)
- Health checks and failover

**Reading Materials:**
- [Load Balancing](../theory/04_Load_Balancing.md)
- Focus on: Load balancing algorithms and strategies

**Hands-on Exercise:**
- Design load balancing strategy for a microservices architecture
- Implement health check logic

**Practice Question:**
- "How do load balancers improve system reliability?"

---

### **Day 9: Caching (Redis, CDN)**
**Topics:**
- Cache invalidation strategies
- Write-through / Write-behind patterns
- Eviction policies (LRU, LFU)

**Reading Materials:**
- [Caching & Performance](../theory/03_Caching_Performance.md)
- Focus on: Caching strategies and patterns

**Hands-on Exercise:**
- Design caching strategy for an e-commerce product catalog
- Implement cache invalidation logic

**Practice Question:**
- "When would you use write-through vs write-behind caching?"

---

### **Day 10: SQL vs NoSQL, CAP Theorem**
**Topics:**
- ACID vs BASE properties
- CAP theorem explained
- Database selection criteria

**Reading Materials:**
- [Database Design](../theory/02_Database_Design.md)
- Focus on: SQL vs NoSQL comparison and CAP theorem

**Hands-on Exercise:**
- Choose appropriate database for 3 different scenarios
- Design data models for each choice

**Practice Question:**
- "Explain CAP theorem with real-world examples"

---

### **Day 11: Indexing, Sharding, Replication**
**Topics:**
- B-Tree indexing basics
- Consistent hashing for sharding
- Master-Slave replication

**Reading Materials:**
- [Database Deep Dive](../theory/13_Database_Deep_Dive.md)
- Focus on: Indexing, sharding, and replication strategies

**Hands-on Exercise:**
- Design sharding strategy for a user database
- Create index plan for query optimization

**Practice Question:**
- "How does consistent hashing help in database sharding?"

---

### **Day 12: Consistency Models**
**Topics:**
- Strong Consistency
- Eventual Consistency
- Causal Consistency

**Reading Materials:**
- [Advanced Distributed Systems](../theory/10_Advanced_Distributed_Systems.md)
- Focus on: Consistency models and trade-offs

**Hands-on Exercise:**
- Analyze consistency requirements for different systems
- Design consistency strategy for a social media platform

**Practice Question:**
- "When is eventual consistency acceptable?"

---

### **Day 13: CDN + Global Load Delivery**
**Topics:**
- Edge servers and content distribution
- Push vs Pull CDN strategies
- Global content delivery optimization

**Reading Materials:**
- [Caching & Performance](../theory/03_Caching_Performance.md)
- Focus on: CDN concepts and global delivery

**Hands-on Exercise:**
- Design CDN strategy for a video streaming platform
- Calculate cache hit rates and cost optimization

**Practice Question:**
- "How do CDNs improve user experience globally?"

---

## **Phase 3: Core System Design Components (Day 14-17)**

### **Day 14: Message Queues & Pub/Sub**
**Topics:**
- Kafka vs RabbitMQ
- Service decoupling patterns
- At-least-once delivery guarantees

**Reading Materials:**
- [Architecture Patterns](../theory/05_Architecture_Patterns.md)
- Focus on: Message queues and event-driven architecture

**Hands-on Exercise:**
- Design message queue system for order processing
- Implement pub/sub pattern for notifications

**Practice Question:**
- "How do message queues improve system scalability?"

---

### **Day 15: Microservices vs Monolith**
**Topics:**
- Service discovery mechanisms
- API Gateway patterns
- Pros/Cons analysis

**Reading Materials:**
- [Architecture Patterns](../theory/05_Architecture_Patterns.md)
- Focus on: Microservices architecture patterns

**Hands-on Exercise:**
- Break down a monolith into microservices
- Design service discovery mechanism

**Practice Question:**
- "When would you choose monolith over microservices?"

---

### **Day 16: Rate Limiting & Throttling**
**Topics:**
- Token Bucket algorithm
- Leaky Bucket algorithm
- Fixed Window vs Sliding Window

**Reading Materials:**
- [Rate Limiter HLD](../product_based_companies/05_hld_rate_limiter.md)
- Focus on: Rate limiting algorithms and implementation

**Hands-on Exercise:**
- Implement token bucket rate limiter
- Design distributed rate limiting system

**Practice Question:**
- "How does rate limiting protect your system?"

---

### **Day 17: Monitoring, Logging & Observability**
**Topics:**
- Prometheus metrics
- Grafana dashboards
- ELK Stack for logging
- Distributed tracing

**Reading Materials:**
- [Monitoring & DevOps](../theory/09_Monitoring_DevOps.md)
- Focus on: Monitoring and observability patterns

**Hands-on Exercise:**
- Design monitoring dashboard for a microservice
- Implement distributed tracing

**Practice Question:**
- "What are the three pillars of observability?"

---

## **Phase 4: Learn by Mini Projects (Day 18-23)**

### **Day 18: TinyURL Project**
**Topics:**
- Hashing strategies for URL shortening
- Database mapping and storage
- Expiration and analytics

**Reading Materials:**
- [URL Shortener HLD](../product_based_companies/01_hld_url_shortener.md)
- [URL Shortener Case Study](../case_studies/design_url_shortener.md)

**Hands-on Exercise:**
- Design URL shortening algorithm
- Implement basic URL shortener service
- Add analytics and expiration

**Practice Question:**
- "How would you handle hash collisions in a URL shortener?"

---

### **Day 19: Rate Limiter Project**
**Topics:**
- Distributed rate limiting
- Redis-based caching
- Token bucket implementation

**Reading Materials:**
- [Rate Limiter HLD](../product_based_companies/05_hld_rate_limiter.md)

**Hands-on Exercise:**
- Implement distributed rate limiter using Redis
- Add different rate limiting algorithms
- Test with high concurrency

**Practice Question:**
- "How do you handle rate limiting in a distributed system?"

---

### **Day 20: Chat App Project**
**Topics:**
- WebSockets for real-time communication
- Message ordering guarantees
- Presence detection and user status

**Reading Materials:**
- [Chat Application HLD](../product_based_companies/02_hld_chat_application.md)

**Hands-on Exercise:**
- Design WebSocket architecture for chat
- Implement message persistence
- Add online/offline status tracking

**Practice Question:**
- "How do you ensure message ordering in a distributed chat system?"

---

### **Day 21: Notification System Project**
**Topics:**
- Push/SMS/Email delivery
- Event queue design
- Fan-out pattern for scalability

**Reading Materials:**
- [Architecture Patterns - Message Queues](../theory/05_Architecture_Patterns.md)

**Hands-on Exercise:**
- Design notification system architecture
- Implement fan-out pattern
- Add delivery tracking and retry logic

**Practice Question:**
- "How do you handle notification failures and retries?"

---

### **Day 22: File Storage (Drive Clone)**
**Topics:**
- Block vs Object Storage
- S3 architecture principles
- Upload chunking and metadata management

**Reading Materials:**
- [Infrastructure & Cloud](../theory/12_Infrastructure_Cloud.md)

**Hands-on Exercise:**
- Design file upload system with chunking
- Implement file versioning
- Add access control and sharing

**Practice Question:**
- "How do you handle large file uploads reliably?"

---

### **Day 23: E-commerce System Project**
**Topics:**
- Payment gateway integration
- Inventory management
- Flash sale handling

**Reading Materials:**
- [E-commerce Amazon HLD](../product_based_companies/07_hld_ecommerce_amazon.md)

**Hands-on Exercise:**
- Design e-commerce architecture
- Implement shopping cart and checkout
- Add flash sale capacity planning

**Practice Question:**
- "How do you handle inventory consistency during flash sales?"

---

## **Phase 5: Advanced Scaling Concepts (Day 24-26)**

### **Day 24: Distributed Systems Basics**
**Topics:**
- Logical clocks and time synchronization
- Consensus algorithms (Paxos & Raft)
- Distributed coordination

**Reading Materials:**
- [Advanced Distributed Systems](../theory/10_Advanced_Distributed_Systems.md)

**Hands-on Exercise:**
- Implement simple consensus algorithm
- Design distributed lock service
- Analyze failure scenarios

**Practice Question:**
- "Explain Raft consensus algorithm in simple terms"

---

### **Day 25: Partitioning & Replication**
**Topics:**
- Leader election algorithms
- Quorum reads/writes
- Data replication lag management

**Reading Materials:**
- [Database Deep Dive](../theory/13_Database_Deep_Dive.md)

**Hands-on Exercise:**
- Design partitioning strategy for social network
- Implement quorum-based replication
- Handle network partitions

**Practice Question:**
- "How do you handle split-brain scenarios in distributed systems?"

---

### **Day 26: Scaling for Millions**
**Topics:**
- Bottleneck identification
- Horizontal scaling strategies
- Database sharding at scale

**Reading Materials:**
- [Scalability & Availability](../theory/06_Scalability_Availability.md)

**Hands-on Exercise:**
- Analyze system bottlenecks
- Design scaling roadmap
- Plan database sharding strategy

**Practice Question:**
- "What are the common bottlenecks when scaling to millions of users?"

---

## **Phase 6: Real Interview Projects (Day 27-29)**

### **Day 27: Uber Backend Design**
**Topics:**
- Quadtree geospatial indexing
- Real-time driver matching
- Location service at scale

**Reading Materials:**
- [Ride Sharing HLD](../product_based_companies/04_hld_ride_sharing.md)
- [Uber Backend Case Study](../case_studies/design_uber_backend.md)

**Hands-on Exercise:**
- Design geospatial indexing system
- Implement driver matching algorithm
- Handle real-time location updates

**Practice Question:**
- "How do you efficiently find nearby drivers in real-time?"

---

### **Day 28: YouTube/Netflix Design**
**Topics:**
- Video encoding pipeline
- CDN distribution strategies
- Adaptive bitrate streaming

**Reading Materials:**
- [Video Streaming HLD](../product_based_companies/03_hld_video_streaming.md)

**Hands-on Exercise:**
- Design video processing pipeline
- Plan CDN distribution strategy
- Implement adaptive streaming logic

**Practice Question:**
- "How do you handle video streaming for millions of concurrent users?"

---

### **Day 29: Instagram/Threads Design**
**Topics:**
- News Feed algorithms
- Fan-out on write vs read
- Media storage scaling

**Reading Materials:**
- [Popular System Designs](../theory/08_Popular_System_Designs.md)

**Hands-on Exercise:**
- Design news feed generation system
- Implement fan-out pattern
- Plan media storage strategy

**Practice Question:**
- "How do you generate personalized news feeds efficiently?"

---

## **Phase 7: Interview Mastery & Resources (Day 30)**

### **Day 30: Mock Interview, Resume & LinkedIn, Resources**
**Topics:**
- Communication best practices
- Behavioral questions integration
- Resume projects section optimization
- LinkedIn profile enhancement

**Reading Materials:**
- [Interview Strategy - Advanced](../theory/14_Interview_Strategy_Advanced.md)
- [Mock Interview Templates](mock-interview-templates.md)
- [Behavioral Questions Guide](behavioral-questions-guide.md)

**Hands-on Exercise:**
- Conduct mock interviews with peers
- Update resume with system design projects
- Optimize LinkedIn profile

**Practice Question:**
- Full mock interview: Design a system of your choice

---

## **Daily Schedule Recommendation**

### **Weekday Structure (2-3 hours):**
- **30 min:** Reading theory
- **45 min:** Hands-on coding exercise
- **30 min:** Practice question explanation
- **15 min:** Review and reflection

### **Weekend Structure (4-5 hours):**
- **1 hour:** Project work
- **30 min:** Mock interview practice
- **30 min:** Review week's topics
- **30 min:** Plan next week

---

## **Assessment Criteria**

### **Weekly Milestones:**
- **Week 1:** Can explain basic system design concepts
- **Week 2:** Understand fundamental components and trade-offs
- **Week 3:** Can design small systems end-to-end
- **Week 4:** Handle complex distributed systems

### **Final Assessment:**
- Complete 3 full system design interviews
- Score 80%+ on practice questions
- Build 2 complete mini-projects

---

## **Additional Resources**

### **Coding Exercises:**
- [Daily Coding Exercises](daily-coding-exercises.md)
- [Project Implementations](../product_based_companies/)

### **Interview Preparation:**
- [Mock Interview Templates](mock-interview-templates.md)
- [Behavioral Questions Guide](behavioral-questions-guide.md)
- [Common Pitfalls to Avoid](common-pitfalls.md)

### **Community & Support:**
- Join system design study groups
- Participate in mock interview sessions
- Contribute to open source projects

---

**Remember:** System design is a skill that improves with practice. Focus on understanding the "why" behind each design decision, not just memorizing solutions. Good luck on your 30-day journey!
