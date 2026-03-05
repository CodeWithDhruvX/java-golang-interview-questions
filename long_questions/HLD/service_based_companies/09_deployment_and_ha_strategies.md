# High-Level Design (HLD): Deployment and High Availability (HA) Strategies

Service-based interviews often focus on operational realities and how code is safely moved from development to production without breaking existing clients.

## 1. Explain the differences between Blue-Green, Canary, and Rolling Deployments.
**Answer:**
These are deployment strategies designed to eliminate downtime during releases and minimize risk if a new version contains a critical bug.
*   **Rolling Update:** The default in Kubernetes. Instances of the old version (V1) are gradually replaced by instances of the new version (V2).
    *   *Pros:* No downtime, does not require double the infrastructure.
    *   *Cons:* During the rollout, both V1 and V2 are serving traffic simultaneously. This can cause issues if V1 and V2 expect different database schemas. Rollbacks take time.
*   **Blue-Green Deployment:** You maintain two identical production environments. Blue is currently live (V1). You deploy V2 completely to the Green environment. You run internal tests on Green. Once verified, you flip the switch at the Load Balancer level to route 100% of traffic to Green.
    *   *Pros:* Instant rollback (just flip the switch back to Blue). No "mixed version" traffic.
    *   *Cons:* Very expensive. Requires double the infrastructure costs for the duration of the deployment.
*   **Canary Release:** Similar to a miner's canary, used for risk mitigation. You route 99% of traffic to V1, and just 1% of traffic to V2. You monitor V2 closely for errors. If it looks healthy, you increase it to 5%, then 20%, then 100%.
    *   *Pros:* Lowest risk. If V2 contains a bug, only 1% of users are affected.
    *   *Cons:* Complex to script and monitor. Requires a smart load balancer/service mesh (like Istio) to handle percentage-based routing.

## 2. Active-Active vs. Active-Passive Failover Architectures
**Answer:**
When building a Highly Available (HA) system, you must design for the eventuality that a whole data center or AWS Region might go offline.
*   **Active-Passive (Hot-Standby):** The primary data center handles 100% of user traffic. The secondary data center sits idle, simply receiving asynchronous database replication logs. If the primary goes down, the DNS is manually or automatically updated to point all traffic to the secondary.
    *   *Pros:* Easier to build, no complex multi-region data conflict resolution needed.
    *   *Cons:* You are paying for a massive secondary data center that sits doing nothing 99.9% of the time. The failover switch might result in a few minutes of downtime or data loss (if async replication hasn't caught up).
*   **Active-Active:** Both data centers are active simultaneously and both handle portions of the global traffic (routed based on geography via Route 53 or Azure Traffic Manager).
    *   *Pros:* Zero waste of resources. Zero downtime if one region dies (the other just takes the full load).
    *   *Cons:* Extremely difficult to build. Requires multi-master database replication (like Cassandra or Spanner) and perfect conflict resolution algorithms if a user connected to Europe edits data simultaneously with a user connected to US East.

## 3. How do you manage Database schema changes (Migrations) with Zero Downtime?
**Answer:**
When deploying a Rolling Update, old code and new code run at the same time. You cannot simply drop a column or rename a table, as the old code will instantly crash.
**The Expand and Contract Pattern:**
1.  **Phase 1 (Expand):** Add the new column `full_name` to the DB. Do not remove the old columns `first_name` and `last_name`. Deploy V2 of the app. V2 reads/writes to `full_name` AND writes to `first_name/last_name` to keep data backwards compatible for V1.
2.  **Phase 2 (Data Migration):** Run a background script to backfill historical data. Concatenate old `first_name` and `last_name` into the new `full_name` column for all existing rows.
3.  **Phase 3 (Contract):** Once all instances. Deploy V3 of the application that only references `full_name`. Finally, run a DB script to `DROP` the `first_name` and `last_name` columns safely.
