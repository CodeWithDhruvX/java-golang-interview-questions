package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Name     string
	Interval time.Duration
	LastRun  time.Time
}

type Scheduler struct {
	tasks    chan Task
	stop     chan struct{}
	wg       sync.WaitGroup
	running  bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(chan Task, 100),
		stop:  make(chan struct{}),
	}
}

func (s *Scheduler) AddTask(id int, name string, interval time.Duration) {
	task := Task{
		ID:       id,
		Name:     name,
		Interval: interval,
		LastRun:  time.Time{},
	}
	s.tasks <- task
	fmt.Printf("Scheduled task %d: %s (every %v)\n", id, name, interval)
}

func (s *Scheduler) Start() {
	if s.running {
		return
	}
	s.running = true
	
	s.wg.Add(1)
	go s.run()
	fmt.Println("Task scheduler started")
}

func (s *Scheduler) run() {
	defer s.wg.Done()
	
	taskMap := make(map[int]Task)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case task := <-s.tasks:
			taskMap[task.ID] = task
		case <-ticker.C:
			now := time.Now()
			for id, task := range taskMap {
				if task.LastRun.IsZero() || now.Sub(task.LastRun) >= task.Interval {
					go s.executeTask(task)
					task.LastRun = now
					taskMap[id] = task
				}
			}
		case <-s.stop:
			fmt.Println("Scheduler stopping...")
			return
		}
	}
}

func (s *Scheduler) executeTask(task Task) {
	fmt.Printf("Executing task %d: %s at %v\n", 
		task.ID, task.Name, time.Now().Format("15:04:05"))
	
	// Simulate task work
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("Completed task %d: %s\n", task.ID, task.Name)
}

func (s *Scheduler) Stop() {
	if !s.running {
		return
	}
	
	close(s.stop)
	s.wg.Wait()
	s.running = false
	fmt.Println("Task scheduler stopped")
}

func main() {
	scheduler := NewScheduler()
	
	fmt.Println("=== Task Scheduler with Intervals ===")
	
	// Add tasks with different intervals
	scheduler.AddTask(1, "Database Cleanup", 2*time.Second)
	scheduler.AddTask(2, "Log Rotation", 3*time.Second)
	scheduler.AddTask(3, "Health Check", 1*time.Second)
	scheduler.AddTask(4, "Metrics Collection", 1500*time.Millisecond)
	
	// Start scheduler
	scheduler.Start()
	
	// Let it run for 8 seconds
	time.Sleep(8 * time.Second)
	
	// Stop scheduler
	scheduler.Stop()
	fmt.Println("Scheduler demonstration completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a task scheduler that runs jobs at different intervals?

**Your Response:** I implement a task scheduler using a ticker-based approach with a task map that tracks the last run time of each task.

The scheduler uses a ticker that fires every 100ms to check which tasks are ready to run. Each task has its own interval, and I compare the current time with the last run time to determine if it's time to execute.

Tasks are executed in separate goroutines so they don't block the scheduler's main loop. The scheduler maintains a map of tasks by ID and updates their last run time after execution.

The key insight is separating task scheduling from task execution. The scheduler only decides when to run tasks, while the actual work happens concurrently. This ensures the scheduler remains responsive even if tasks take time to complete.

I use channels for task submission and a stop channel for graceful shutdown. This pattern is commonly used in real systems for cron jobs, periodic maintenance tasks, health checks, and any scheduled operations that need to run at regular intervals.
