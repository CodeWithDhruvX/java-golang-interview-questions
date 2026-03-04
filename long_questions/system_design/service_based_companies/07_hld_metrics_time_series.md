# System Design (HLD) - Time Series Database (Metrics System)

## Problem Statement
Design a massive-scale Time Series Metrics System (like Datadog, Prometheus, or AWS CloudWatch) that can ingest billions of data points per minute from servers globally and allow developers to query charts instantly.

## 1. Requirements Clarification
### functional Requirements
*   **Ingestion:** Accept telemetry data (CPU%, Memory usage, HTTP requests) tagged with metadata (Host, Region, AppId).
*   **Querying:** Allow fast retrieval of data over time ranges (e.g., "Show me avg CPU for AppId X over the last 7 days grouped by hour").
*   **Alerting:** Trigger alerts if thresholds are breached.

### Non-Functional Requirements
*   **Extreme Write Load:** Time-series is inherently incredibly write-heavy (99% writes, 1% reads).
*   **Read Latency:** Queries for dashboards must return in under a second.

## 2. Understanding Time Series Data
A time-series data point looks like this:
`{ "metric": "cpu.usage.idle", "tags": {"host": "server-1", "region": "us-east"}, "timestamp": 1690000000, "value": 85.5 }`

Traditional RDBMS (SQL) or Document stores (MongoDB) will collapse under the sheer volume of billions of row-inserts and the specific aggregation math required for querying this data.

## 3. High-Level Architecture

```text
[ Application Servers / Agents ] 
               |
         (Push / Pull)
               V
       [ Ingestion API ] ---> [ Stream Processor (Kafka / Flink) ] ---> (Real-time Alerting Engine)
               |
               +--------------------------------------+
                                                      |
                                           [ Time Series Database ]
                                           (InfluxDB / Cassandra)
```

## 4. Database Selection & Design
The core of this interview question is the internal design of the TSDB. Wide-column stores like **Apache Cassandra** or specialized databases like **InfluxDB** are preferred.

**Why Cassandra?**
*   It handles massive write throughput easily.
*   Data is written Sequentially (LSM Trees). Disks love sequential writes.

**Data Modeling in Cassandra for TSDB:**
We partition the data by the metric name + tags, and we order the columns by timestamp.
*   `Partition Key:` Hash of (MetricName + Tags). This ensures all data for a specific server's CPU goes to the same physical hard drive in the Cassandra cluster.
*   `Clustering Key:` Timestamp. This ensures the data is stored sorted chronologically on the disk.
When a user asks for data between `T1` and `T2`, Cassandra goes to the exact server (Partition), finds the start of `T1` on disk, and sequentially reads up to `T2`. Highly efficient!

## 5. Rollups & Downsampling (Crucial Optimization)
Storing 1 data point per second means 60/min -> 3600/hour -> 86,400/day. If someone queries a 1-year chart, scanning billions of points will timeout.
*   **Downsampling:** A background batch job (e.g., Spark or an internal Cron) continuously takes the raw 1-second data and compresses it.
    *   It calculates the Min, Max, and Avg for every minute, and writes it to a new `metrics_1min` table.
    *   It does the same for 1-hour intervals (`metrics_1hour` table).
*   When the UI requests a 1-year chart, the API queries the `metrics_1day` table (returning only 365 rows), rendering instantly.

## 6. Real-time Alerting
Alerting cannot wait for the data to be safely written to Cassandra and then polled.
*   The Ingestion API drops all data points into **Kafka**.
*   A stream processing engine (like **Apache Flink** or **Spark Streaming**) consumes the Kafka topic in real-time.
*   It holds a "sliding window" in RAM (e.g., the last 5 minutes of CPU data). If the average crosses 90%, it instantly fires a webhook to PagerDuty.

## 7. Follow-up Questions for Candidate
1.  How do you handle late-arriving data? (e.g., A mobile phone was offline, and uploads 3 days of metrics when it connects to Wifi). (The database handles it naturally due to Timestamp clustering, but the Alerting engine needs a watermarking policy to discard extremely old alerts).
2.  Data compression is key since TSDBs generate petabytes quickly. How is it compressed? (Delta-of-Delta encoding - e.g., Gorilla compression. Since timestamps increment predictably (+10s, +10s), we just store the delta (10), and if the value changes by tiny amounts, we store bit-shifted deltas, reducing a 64-bit float to 1-2 bits).
