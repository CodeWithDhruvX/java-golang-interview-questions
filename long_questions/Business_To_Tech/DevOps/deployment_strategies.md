# Deployment Strategies

## 1. Rolling Deployment (Default in K8s)
Gradually replaces instances of the old version with the new version.

*   **Mechanism**:
    *   Running: V1 (10 pods).
    *   Deployment starts.
    *   Creates 2 pods of V2. Waits for health check.
    *   Terminates 2 pods of V1.
    *   Repeats until 10 V2 pods and 0 V1 pods.
*   **Pros**: No downtime.
*   **Cons**: Slow rollout. Difficult to rollback instantly.

## 2. Blue-Green Deployment
Running two identical environments simultaneously.

*   **Mechanism**:
    *   **Blue (Live)**: Running V1. Serving 100% traffic.
    *   **Green (Idle)**: Deploy V2. Test it thoroughly (Integration tests).
    *   **Switch**: Update Load Balancer to point to Green. V2 is now Live.
*   **Pros**: Instant switch. Instant rollback (just switch LB back to Blue).
*   **Cons**: Expensive. Requires 2x resources cost.

## 3. Canary Deployment
Releasing the change to a small subset of users first.

*   **Mechanism**:
    *   Route 95% traffic to V1.
    *   Route 5% traffic to V2 (The "Canary").
    *   Monitor metrics (Error rate, Latency).
    *   If healthy, gradually increase to 10%, 50%, 100%.
    *   If unhealthy, kill V2 instantly.
*   **Pros**: Lowest risk. Real-world testing.
*   **Cons**: Complex traffic routing setup (Istio / Mesh).

## 4. Comparison Table

| Strategy | Downtime | Rollback Duration | Cost | Risk |
| :--- | :--- | :--- | :--- | :--- |
| **Recreate** | Yes (High) | Slow | Low | High |
| **Rolling** | No | Slow | Low | Medium |
| **Blue-Green** | No | Instant | High (2x) | Low |
| **Canary** | No | Instant | Low | Lowest |

## 5. Interview Questions
1.  **Which deployment strategy for a Mission Critical Payment Service?**
    *   *Ans*: **Canary**. You cannot afford to break payments for everyone. Test with 1% of users first.
2.  **How to handle Database changes in Blue-Green?**
    *   *Ans*: Hardest part. The Database is shared by both Blue (V1) and Green (V2).
    *   *Rule*: DB changes must be **Backward Compatible**. (e.g., Add a nullable column, don't rename a column).
