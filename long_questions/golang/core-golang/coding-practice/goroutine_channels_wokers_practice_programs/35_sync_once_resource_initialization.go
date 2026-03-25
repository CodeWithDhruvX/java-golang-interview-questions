package main

import (
	"fmt"
	"sync"
	"time"
)

type DatabaseConnection struct {
	connected bool
	createdAt time.Time
}

func (dc *DatabaseConnection) Connect() error {
	if dc.connected {
		return fmt.Errorf("already connected")
	}
	
	// Simulate expensive connection setup
	time.Sleep(2 * time.Second)
	dc.connected = true
	dc.createdAt = time.Now()
	fmt.Println("Database connection established")
	return nil
}

func (dc *DatabaseConnection) Query(sql string) string {
	if !dc.connected {
		return "Error: not connected"
	}
	return fmt.Sprintf("Result for: %s", sql)
}

type Service struct {
	once     sync.Once
	db       *DatabaseConnection
}

func NewService() *Service {
	return &Service{
		db: &DatabaseConnection{},
	}
}

func (s *Service) getDatabase() *DatabaseConnection {
	s.once.Do(func() {
		fmt.Println("Initializing database connection...")
		s.db.Connect()
	})
	return s.db
}

func (s *Service) HandleRequest(requestID int) {
	db := s.getDatabase()
	result := db.Query(fmt.Sprintf("SELECT * FROM table WHERE id = %d", requestID))
	fmt.Printf("Request %d: %s\n", requestID, result)
}

func main() {
	service := NewService()
	wg := sync.WaitGroup{}

	fmt.Println("Starting multiple requests with lazy initialization")

	// Start multiple concurrent requests
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			service.HandleRequest(requestID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does sync.Once ensure resource initialization happens only once?

**Your Response:** sync.Once guarantees that a function is executed exactly once, no matter how many goroutines call it concurrently. I use it to implement lazy initialization of the database connection.

The getDatabase method calls s.once.Do() with an initialization function. The first goroutine to reach this point executes the initialization, while all other goroutines block until it completes. After the first execution, subsequent calls to once.Do() do nothing - they just return immediately.

This ensures the expensive database connection happens only once, even if multiple requests arrive simultaneously. It's thread-safe and more efficient than using a mutex because after initialization, there's no locking overhead.

The key insight is that sync.Once is perfect for one-time setup operations like loading configuration, establishing connections, or initializing caches. It combines the safety of proper synchronization with the performance of no overhead after initialization.
