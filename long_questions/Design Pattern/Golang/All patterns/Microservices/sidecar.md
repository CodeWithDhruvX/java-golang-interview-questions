# Sidecar Pattern

## 🟢 What is it?
The **Sidecar Pattern** involves deploying a helper service of its own container/process alongside the main application service (the "Sidecar" attached to the "Motorcycle").

The sidecar handles peripheral tasks such as logging, monitoring, configuration synchronization, or network proxying, allowing the main application to focus purely on business logic.

---

## 🏛️ Real World Analogy
**Motorcycle Sidecar**:
*   The Motorcycle (Main App) has the engine and steering (Business Logic).
*   The Sidecar (Helper) carries the passenger or extra luggage (Infrastructure concerns).
*   They are separate units but travel together and share the same lifecycle.

---

## 🎯 Strategy to Implement

In a Kubernetes Context:
1.  **Pod Definition**: Define a Pod with **two containers**.
2.  **Shared Volume**: Use a shared `emptyDir` volume if they need to share files (e.g., logs or configs).
3.  **Localhost Networking**: They share the same network namespace, so they can talk via `localhost:port`.

In a Golang Code Context (Simulating the behavior):
You can run a "Sidecar" goroutine or process that proxies traffic.

---

## 💻 Code Example (Simulating Log Collection)

Imagine a main app writing logs to a file, and a sidecar reading them and sending to a server.

```go
package main

import (
	"fmt"
	"os"
	"time"
)

// 1. The Main Application (Motorcycle)
func mainApp(logFile string) {
	f, _ := os.Create(logFile)
	defer f.Close()

	for i := 0; i < 5; i++ {
		// Business Logic
		msg := fmt.Sprintf("Processing Order %d\n", i)
		fmt.Println("App: Writing log...")
		f.WriteString(msg)
		time.Sleep(1 * time.Second)
	}
}

// 2. The Sidecar (Helper)
func sidecar(logFile string) {
	// Waits for file to exist
	for {
		if _, err := os.Stat(logFile); err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	file, _ := os.Open(logFile)
	defer file.Close()

	// Tailing the file (Simplified)
	buffer := make([]byte, 1024)
	for {
		n, _ := file.Read(buffer)
		if n > 0 {
			content := string(buffer[:n])
			// Helper Logic: Determine if we should alert
			fmt.Printf("[Sidecar] Shipping logs to Splunk: %s", content)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	logPath := "./shared_logs.txt"
	
	// Start Sidecar in background
	go sidecar(logPath)

	// Start Main App
	mainApp(logPath)
	
	// Keep alive for demo
	time.Sleep(2 * time.Second)
}
```

---

## ✅ When to use?

*   **Service Mesh (Istio/Linkerd)**: The "Envoy Proxy" sidecar intercepts all network traffic to handle TLS, retry logic, and tracing without changing the app code.
*   **Log Aggregation**: Application writes to `stdout` or a file; Sidecar (Filebeat/Fluentd) reads it and ships to Elasticsearch.
*   **Config Hot-Reloading**: A sidecar watches a remote Config Server (Consul/Etcd) and updates a local JSON file that the app reads.
