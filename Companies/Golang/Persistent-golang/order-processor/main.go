package main

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Config
const (
	MaxWorkerPoolSize = 5
	MaxJobQueueSize   = 100
	Port              = ":8080"
)

// Metrics (using expvar for thread-safe counters, standard library)
var (
	ordersProcessed = expvar.NewInt("orders_processed")
	ordersFailed    = expvar.NewInt("orders_failed")
	queueDepth      = expvar.NewInt("queue_depth")
)

// Order represents the incoming request payload
type Order struct {
	ID    string  `json:"id"`
	Value float64 `json:"value"`
}

// Job represents the unit of work for the worker
type Job struct {
	Order Order
}

func main() {
	// 1. Setup Worker Pool
	// Buffered channel acts as a job queue with defined capacity
	jobQueue := make(chan Job, MaxJobQueueSize)
	var wg sync.WaitGroup

	// Create a context for graceful shutdown of workers
	ctx, cancel := context.WithCancel(context.Background())

	// Start Workers (Goroutines)
	log.Printf("Starting %d workers...", MaxWorkerPoolSize)
	for i := 1; i <= MaxWorkerPoolSize; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i, jobQueue)
	}

	// 2. Setup HTTP Server
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", submitHandler(jobQueue))

	// Native Expvar metrics (JSON)
	mux.Handle("/debug/vars", expvar.Handler())

	// Custom Prometheus-style endpoint (Demonstrating JD requirement)
	mux.HandleFunc("/metrics", prometheusHandler)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	// 3. Start Server in a Goroutine
	go func() {
		log.Printf("Server listening on %s", Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 4. Graceful Shutdown Logic
	// Wait for interrupt signal (Ctrl+C or SIGTERM from K8s)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("\nShutdown signal received...")

	// Create a deadline to wait for current requests to complete
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Stop accepting new HTTP requests
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Signal workers to stop processing *new* jobs from the queue
	// Note: In a real world scenario, you might close the channel instead,
	// but context cancellation allows for immediate stop of idle workers.
	log.Println("Stopping workers...")
	cancel()

	// Wait for workers to finish their cleanup
	wg.Wait()
	log.Println("Server exited properly")
}

// worker is the consumer function running in a Goroutine
func worker(ctx context.Context, wg *sync.WaitGroup, id int, jobs <-chan Job) {
	defer wg.Done()
	log.Printf("Worker %d ready", id)

	for {
		select {
		case <-ctx.Done():
			// Context cancelled, stop worker
			return
		case job := <-jobs:
			log.Printf("[Worker %d] Processing Key: %s Val: %.2f", id, job.Order.ID, job.Order.Value)

			// Simulate processing time
			time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond)

			// Simulate Random Failure
			if rand.Float32() < 0.1 {
				ordersFailed.Add(1)
				log.Printf("[Worker %d] Failed processing %s", id, job.Order.ID)
			} else {
				ordersProcessed.Add(1)
			}

			// Decrement queue depth
			queueDepth.Add(-1)
		}
	}
}

// submitHandler receives HTTP requests and enqueues them
func submitHandler(jobQueue chan<- Job) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var order Order
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
			http.Error(w, "Invalid body", http.StatusBadRequest)
			return
		}

		// Non-blocking send to handle Backpressure
		// If the queue is full, we reject the request immediately (503 Service Unavailable)
		// This is critical for microservices stability.
		select {
		case jobQueue <- Job{Order: order}:
			queueDepth.Add(1)
			w.WriteHeader(http.StatusAccepted) // 202 Accepted
			json.NewEncoder(w).Encode(map[string]string{
				"status": "queued",
				"id":     order.ID,
			})
		default:
			// Queue is full
			http.Error(w, "Server busy - Queue full", http.StatusServiceUnavailable)
		}
	}
}

// prometheusHandler manually formats metrics to Prometheus text format
// This avoids importing third-party libraries for this simple demo
func prometheusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")

	fmt.Fprintf(w, "# HELP orders_processed Total valid orders processed\n")
	fmt.Fprintf(w, "# TYPE orders_processed counter\n")
	fmt.Fprintf(w, "orders_processed %v\n", ordersProcessed.String())

	fmt.Fprintf(w, "# HELP orders_failed Total failed orders\n")
	fmt.Fprintf(w, "# TYPE orders_failed counter\n")
	fmt.Fprintf(w, "orders_failed %v\n", ordersFailed.String())

	fmt.Fprintf(w, "# HELP queue_depth Current items in job queue\n")
	fmt.Fprintf(w, "# TYPE queue_depth gauge\n")
	fmt.Fprintf(w, "queue_depth %v\n", queueDepth.String())
}
