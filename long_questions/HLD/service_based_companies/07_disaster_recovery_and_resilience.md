# High-Level Design (HLD): Disaster Recovery and Resilience

Service-based companies often have strict Service Level Agreements (SLAs). If an application goes down, the company loses money or pays penalties. Designing for failure is critical.

## 1. What are RTO and RPO in Disaster Recovery?
**Answer:**
These are foundational business metrics that dictate the technical architecture of backups.
*   **RPO (Recovery Point Objective):** The maximum acceptable amount of data loss measured in time. 
    *   *Example:* If RPO is 1 hour, your backups must happen at least every hour. If the system crashes at 10:59, you lose 59 minutes of data. For banks, RPO is often near 0.
*   **RTO (Recovery Time Objective):** The maximum acceptable amount of time the application can be offline.
    *   *Example:* If RTO is 4 hours, your team has 4 hours to spin up new servers, restore the database from backups, and route traffic back.

## 2. Explain the different Disaster Recovery Strategies for Cloud.
**Answer:**
Ranging from cheapest/slowest to most expensive/fastest:
1.  **Backup and Restore:** (RTO: Hours, RPO: Hours)
    *   Data and configurations are backed up to S3. If primary region fails, you manually spin up a completely new environment in a new region and restore the DB.
2.  **Pilot Light:** (RTO: Minutes, RPO: Minutes)
    *   Data is continuously replicated to a secondary region. The core database is running in the secondary region, but application servers are turned off to save cost. During disaster, you boot the app servers and switch traffic.
3.  **Warm Standby:** (RTO: Minutes, RPO: Seconds)
    *   A scaled-down, fully functional version of the application is always running in the secondary region. It handles a tiny amount of traffic or sits idle. During disaster, you just scale it out.
4.  **Multi-Region Active-Active:** (RTO: Instant, RPO: Zero)
    *   Full production apps running in multiple global regions simultaneously. Traffic is routed based on location. If one region dies, traffic instantly reroutes to others. Extremely complex due to bidirectional database synchronization.

## 3. What is a Circuit Breaker and the Retry Pattern?
**Answer:**
Resilience patterns for microservices to prevent catastrophic cascading failures.
*   **Retry Pattern with Exponential Backoff:** If Service A calls Service B and it fails (e.g., due to a temporary network blip), automatically retry. To prevent overloading a struggling Service B, wait exponentially longer between retries (e.g., 1s, 2s, 4s, 8s). Add "Jitter" (randomness) so not all clients retry at the exact same millisecond.
*   **Circuit Breaker:** If Service B is completely down, repeated retries will exhaust threads in Service A. A circuit breaker monitors failures. If failures pass a threshold, the circuit "opens," short-circuiting calls instantly and returning a default value or error, giving Service B time to recover.

## 4. How do you design an application with no Single Point of Failure (SPOF)?
**Answer:**
Every tier of the architecture must have redundancy.
*   **DNS:** Use highly available global providers (Route53).
*   **Edge/Load Balancing:** Deploy multi-AZ load balancers.
*   **App Tier:** Deploy stateless microservices across multiple Availability Zones in auto-scaling groups. If one server dies, another takes its place.
*   **Database Tier:** Use Master-Slave replication. Ensure the cloud provider automatically promotes the standby to master if the primary hardware fails.
*   **Storage:** Use distributed object storage (S3) which duplicates data across multiple facilities.

## 5. What is Chaos Engineering?
**Answer:**
The discipline of experimenting on a software system in production in order to build confidence in the system's capability to withstand turbulent and unexpected conditions.
*   Instead of waiting for an outage, you intentionally inject failures into production (e.g., Netflix's *Chaos Monkey* randomly terminates virtual machine instances).
*   This forces engineering teams to build fully automated resilience, self-healing, and failover mechanisms.
