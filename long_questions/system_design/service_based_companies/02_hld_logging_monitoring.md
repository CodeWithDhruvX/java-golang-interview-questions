# System Design (HLD) - Centralized Logging & Monitoring

## Problem Statement
Design a Centralized Logging and Monitoring system for a microservices architecture. Developers need a dashboard to search through logs across 50 different microservices, and Operations needs alerts when CPU goes high or error rates spike.

## 1. Requirements Clarification
### functional Requirements
*   **Log Ingestion:** Collect application logs from hundreds of servers.
*   **Search & Visualization:** A UI to search logs by `trace_id`, `error_code`, or `timestamp`.
*   **Metrics:** Collect system metrics (CPU, Memory, JVM Heap, Request Latency).
*   **Alerting:** Send a Slack/Email alert if a metric crosses a threshold (e.g., HTTP 500s > 2%).

### Non-Functional Requirements
*   **Scalable:** Must handle Terabytes of log data per day.
*   **Decoupled:** Logging shouldn't slow down the main applications.

## 2. High-Level Architecture (The Standard Stacks)

The industry standard approach is actually two distinct stacks: One for **Logs** (ELK/EFK Stack) and one for **Metrics** (Prometheus + Grafana).

### A. Logging Stack (ELK / EFK)
```text
[ Microservice A ] -> (Writes to local disk: app.log)
       |
  [ Filebeat / Fluentd ] (Agent reading the log file line by line)
       |
     (Network)
       |
[ Kafka (Buffer) ] ---> [ Logstash ] ---> [ Elasticsearch ] <--- [ Kibana (UI) ]
```

*   **Fluentd / Filebeat:** A lightweight agent installed on every application VM/Container. It tails the log files and ships them out asynchronously.
*   **Kafka:** Buffer massive log spikes (e.g., during a system crash when everything is logging exceptions) to prevent Elasticsearch from dying under the write load.
*   **Logstash:** Parses the raw log text into structured JSON (extracting timestamp, log level, trace ID).
*   **Elasticsearch (ES):** A highly scalable search engine. Stores the parsed logs.
*   **Kibana:** The web UI for developers to perform full-text searches against ES.

### B. Metrics Stack (Prometheus + Grafana)
```text
[ Microservice A ]  <--- (Pulls /metrics endpoint) --- [ Prometheus Server ]
(Exposes /metrics)                                            |
                                                        [ Grafana (UI) ]
                                                              |
                                                      [ Alertmanager ] -> (Slack/Email)
```

*   **Prometheus:** A time-series database. Instead of apps *pushing* metrics, Prometheus *pulls* (scrapes) the `/metrics` endpoint from every server every 10 seconds.
*   **Grafana:** A dashboarding tool that queries Prometheus to draw CPU and Latency charts.
*   **Alertmanager:** Rules engine. "If Memory > 90% for 5 mins, trigger PagerDuty."

## 3. Distributed Tracing
With 50 microservices, a single user request might travel through 5 different services. If it fails, how do you track it?
*   **Trace ID:** When the API Gateway receives a request, it generates a unique `X-B3-TraceId` correlation ID.
*   This ID is passed as an HTTP Header to every downstream microservice.
*   Every microservice includes this `TraceId` in its log statements.
*   In Kibana, a developer searches for that single `TraceId` and instantly sees logs from all 5 microservices in chronological order. Tools like **Jaeger** or **Zipkin** are used for tracing visualization.

## 4. Log Retention Strategies
*   Elasticsearch storage is expensive.
*   **Hot Logs (1-7 days):** Keep in Elasticsearch on fast SSDs for immediate troubleshooting.
*   **Warm Logs (7-30 days):** Move to cheaper HDD nodes within Elasticsearch.
*   **Cold Logs (1 year+):** Compress and archive logs to **AWS S3** for compliance/auditing purposes, and delete them from Elasticsearch.

## 5. Follow-up Questions for Candidate
1.  Why use Prometheus (Pull model) instead of pushing metrics? (Prevents monitoring systems from being overwhelmed by a flood of pushes; easier to know if a target is down because the pull fails).
2.  How do you ensure PII (Personally Identifiable Information like Passwords/Credit Cards) doesn't end up in Elasticsearch? (Apply masking/redaction filters in Logstash before saving to ES).
