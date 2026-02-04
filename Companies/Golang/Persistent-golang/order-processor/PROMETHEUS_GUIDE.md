# Prometheus Setup Guide for Order Processor

This guide explains how to visualize the metrics exposed by your application using Prometheus.

## 1. Prerequisites

-   Your **Order Processor** application must be running (normally on port `8080`).
-   You need **Docker** installed (easiest way).

## 2. Verify Your App's Metrics

Ensure your application is running and exposing metrics:

1.  Open your browser or use `curl`.
2.  Go to: `http://localhost:8080/metrics`
3.  You should see text output starting with `# HELP orders_processed ...`.

## 3. Run Prometheus (Using Docker)

We have created a `prometheus.yml` file in this directory. Run the following command in your terminal while inside the `order-processor` folder:

```powershell
# 1. First, remove any existing container with the same name (if you get a conflict error)
docker rm -f prometheus-order-processor

# 2. Run Prometheus with the EXACT path to your config file
docker run --name prometheus-order-processor -d -p 9090:9090 -v "c:\Users\dhruv\Downloads\personal_projects\golang-java-interview-questions\java-golang-interview-questions\Companies\Persistent-golang\order-processor\prometheus.yml:/etc/prometheus/prometheus.yml" prom/prometheus
```

> **Troubleshooting:**
> *   **Conflict Error:** Run `docker rm -f prometheus-order-processor` to clean up the old failed container.
> *   **Path Error:** Ensure the path above matches exactly where your `prometheus.yml` is located.

## 4. View Dashboard

1.  Open `http://localhost:9090` in your browser.
2.  Click on **Status** -> **Targets** to ensure `order-processor` is "UP".
3.  Click on **Graph**.
4.  Type one of your metrics into the expression bar and click **Execute**:
    -   `orders_processed`
    -   `orders_failed`
    -   `queue_depth`
5.  Switch to the **Graph** tab to see the values change over time.

## 5. Generate Traffic

To see the graph move, submit some orders using a separate terminal:

```powershell
# Send a valid order
curl -X POST http://localhost:8080/submit -d "{\"id\": \"1\", \"value\": 100}"

# Send many orders to see the queue fill up
for ($i=0; $i -lt 50; $i++) { curl -X POST http://localhost:8080/submit -d "{\"id\": \"$i\", \"value\": 10}" }
```
