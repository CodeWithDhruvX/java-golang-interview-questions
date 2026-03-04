# System Design (HLD) - Distributed Job Scheduler

## Problem Statement
Design a Distributed Job Scheduler (like Quartz, Airflow, or AWS EventBridge/CloudWatch Events). The system should execute arbitrary code or trigger HTTP endpoints at specific times or cron intervals.

## 1. Requirements Clarification
### functional Requirements
*   **Schedule Jobs:** Users can submit jobs to run once at a specific time (e.g., "Run on Friday at 3 PM") or on a recurring schedule (e.g., "Run every day at midnight").
*   **Execute Jobs:** The system triggers the job (e.g., sends a webhook, publishes a Kafka message, or runs a Docker container).
*   **Dashboard:** View job history, success/failure logs.

### Non-Functional Requirements
*   **High Availability:** The scheduler cannot go down; otherwise, critical business jobs are missed.
*   **Fault Tolerance:** If a worker executing a job crashes, the job should failover to another worker and retry.
*   **Scale:** Support millions of registered jobs and thousands running concurrently.

## 2. High-Level Architecture

```text
[ Client Svc ] ---> [ Scheduler API Gateway ]
                             |
                     [ Scheduler Service ] ---> [ RDBMS (Job Metadata) ]
                             |
                             V
                 [ Job Execution Queue (Kafka/SQS) ]
                             |
         +-------------------+-------------------+
         |                   |                   |
    [ Worker 1 ]        [ Worker 2 ]        [ Worker 3 ]
         |                   |                   |
    (Runs Docker / Hits Webhooks / Writes to DB)
```

## 3. Core Components

### A. The Database (Metadata)
An RDBMS (MySQL/PostgreSQL) is ideal for storing job definitions because we need ACID guarantees when updating job states from "PENDING" to "RUNNING".
*   **Table: `Jobs`** -> `job_id`, `cron_expression`, `action_type`, `payload`, `status`.
*   **Table: `Job_History`** -> Tracks every execution for the dashboard.

### B. The Scheduler (The "Tick" mechanism)
How does the system know it's time to run a job?
*   **Option 1: DB Polling (Bad for scale).** A thread runs every minute: `SELECT * FROM jobs WHERE next_run_time <= NOW()`. This locks the DB and causes massive latency if there are 1 million jobs.
*   **Option 2: Time-Wheel / Priority Queue (In-Memory).**
    *   The Scheduler service pulls upcoming jobs (e.g., for the next 10 minutes) from the database into an in-memory Min-Heap (ordered by execution time) or a Hashed Wheel Timer.
    *   A fast-ticking thread checks the top/current slot of the data structure. If it's time, it blasts the `job_id` into a message queue.

### C. The Message Queue (Decoupling)
The Scheduler thread should *never* actually execute the job logic itself. It simply signals that it is time to run.
*   Scheduler pushes `{"job_id": 123, "action": "POST /url"}` to Kafka or AWS SQS.

### D. The Workers (Execution)
Thousands of worker nodes listen to the queue.
*   They pick up the message.
*   They update the DB `Job_History` to "RUNNING".
*   They execute the heavy lifting (e.g., making a REST call to a billing service).
*   They update the DB to "COMPLETED" or "FAILED".

## 4. Fault Tolerance & Deadlines
What if Worker 2 picks up Job 123, starts executing, and the server catches on fire?
*   The system needs a heartbeat or timeout mechanism.
*   When a Worker picks up a message from SQS, the message becomes "Invisible" to other workers for say, 5 minutes.
*   If the worker successfully finishes, it explicitly deletes the message from SQS.
*   If the worker burns down, it never deletes the message. After 5 minutes, SQS makes the message visible again, and Worker 3 picks it up.

## 5. Follow-up Questions for Candidate
1.  How do you guarantee a job runs *exactly* once? (It is nearly impossible in distributed systems. We guarantee "At Least Once" delivery via the message queue, and require the *execution logic itself* to be Idempotent so it is safe to run twice).
2.  If you have highly unbalanced jobs (Job A takes 1 second, Job B takes 2 hours), how do you prevent the fast jobs from getting stuck behind slow jobs? (Partition the queues by job duration: FastQueue and SlowQueue, with dedicated workers for each).
