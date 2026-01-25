# API Gateway Pattern

## 🟢 What is it?
The **API Gateway Pattern** sits between clients and microservices. Ideally, clients (Mobile, Web) should not call microservices directly (Service A, Service B, Service C). Instead, they call a single entry point—the API Gateway—which forwards the request to the appropriate service.

It acts as a **Front Door** for your backend.

---

## 🏛️ Real World Analogy
**Hotel Receptionist**:
*   You don't just walk into the kitchen to order food, or run to the laundry room to ask for towels.
*   You call the **Receptionist** (Gateway).
*   The reception routes your request: "Room Service" -> Kitchen, "Maintenance" -> Engineer.
*   The receptionist also checks if you are actually a guest (Authentication) before helping you.

---

## 🎯 Strategy to Implement

1.  **Reverse Proxy**: The core function. Read `req.Path`, determine destination.
2.  **Aggregation**: Call Service A (User Profile) and Service B (Orders), combine them into one JSON response to save the client round-trips (Backend For Frontend - BFF).
3.  **Cross-Cutting Concerns**:
    *   **Auth**: Validate JWT tokens here so services don't have to.
    *   **Rate Limiting**: Block abuse here.
    *   **SSL Termination**: Handle HTTPS here.

---

## 💻 Code Example (Reverse Proxy)

Golang's `net/http/httputil` makes this incredibly easy.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy creates a simple reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}

func main() {
	// Identify where our microservices live
	userServiceUrl := "http://localhost:8081"
	orderServiceUrl := "http://localhost:8082"

	// Create proxies
	userProxy, _ := NewProxy(userServiceUrl)
	orderProxy, _ := NewProxy(orderServiceUrl)

	// Main Router (The Gateway)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if len(path) >= 5 && path[:5] == "/user" {
			fmt.Println("Gateway: Routing to User Service")
			// Strip prefix if needed, or just forward
			userProxy.ServeHTTP(w, r)
		} else if len(path) >= 6 && path[:6] == "/order" {
			fmt.Println("Gateway: Routing to Order Service")
			orderProxy.ServeHTTP(w, r)
		} else {
			http.Error(w, "404 Not Found", http.StatusNotFound)
		}
	})

	fmt.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## ✅ When to use?

*   **Security Barrier**: You want to hide your internal network topology (IP addresses, ports) from the public internet.
*   **Protocol Translation**: External clients speak HTTP/REST, but internal services speak gRPC. The Gateway translates between them.
*   **Request Aggregation**: Mobile app needs "Home Page" data. Gateway calls 5 services and returns 1 JSON.
