package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// How to run:
// 1. go get github.com/onsi/ginkgo/v2
// 2. go get github.com/onsi/gomega
// 3. go test ./...

func TestOrderProcessor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order Processor Suite")
}

var _ = Describe("Order Processor API", func() {
	var (
		jobQueue chan Job
		handler  http.HandlerFunc
	)

	BeforeEach(func() {
		jobQueue = make(chan Job, 10) // Small buffered channel for testing
		handler = submitHandler(jobQueue)
	})

	Describe("Submit Handler", func() {
		Context("When a valid POST request is made", func() {
			It("Should enqueue the job and return 202 Accepted", func() {
				order := Order{ID: "TEST-1", Value: 100.50}
				body, _ := json.Marshal(order)
				req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				handler.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusAccepted))

				// Check response body
				var resp map[string]string
				json.Unmarshal(w.Body.Bytes(), &resp)
				Expect(resp["status"]).To(Equal("queued"))
				Expect(resp["id"]).To(Equal("TEST-1"))

				// Check if job was enqueued
				var job Job
				Eventually(jobQueue).Should(Receive(&job))
				Expect(job.Order.ID).To(Equal("TEST-1"))
				Expect(job.Order.Value).To(Equal(100.50))
			})
		})

		Context("When the request method is not POST", func() {
			It("Should return 405 Method Not Allowed", func() {
				req := httptest.NewRequest(http.MethodGet, "/submit", nil)
				w := httptest.NewRecorder()

				handler.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})

		Context("When the request body is invalid", func() {
			It("Should return 400 Bad Request", func() {
				req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString("invalid-json"))
				w := httptest.NewRecorder()

				handler.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("When the queue is full", func() {
			It("Should return 503 Service Unavailable", func() {
				// Fill the queue
				for i := 0; i < 10; i++ {
					jobQueue <- Job{Order: Order{ID: "Filler"}}
				}

				order := Order{ID: "Overflow", Value: 50.0}
				body, _ := json.Marshal(order)
				req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBuffer(body))
				w := httptest.NewRecorder()

				handler.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusServiceUnavailable))
			})
		})
	})

	Describe("Metrics Handler", func() {
		It("Should return Prometheus formatted metrics", func() {
			req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
			w := httptest.NewRecorder()

			prometheusHandler(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			body := w.Body.String()

			Expect(body).To(ContainSubstring("# HELP orders_processed"))
			Expect(body).To(ContainSubstring("# TYPE orders_processed counter"))
			Expect(body).To(ContainSubstring("orders_processed"))

			// Verify basic structure of other metrics
			Expect(body).To(ContainSubstring("orders_failed"))
			Expect(body).To(ContainSubstring("queue_depth"))
		})
	})
})
