package main

import (
	"fmt"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Source    string
}

type LogProcessor struct {
	input     chan LogEntry
	output    chan string
	wg        sync.WaitGroup
}

func NewLogProcessor() *LogProcessor {
	return &LogProcessor{
		input:  make(chan LogEntry, 1000),
		output: make(chan string, 1000),
	}
}

func (lp *LogProcessor) Start() {
	lp.wg.Add(1)
	go lp.process()
}

func (lp *LogProcessor) process() {
	defer lp.wg.Done()
	defer close(lp.output)
	
	// Buffer logs for ordering
	logBuffer := make([]LogEntry, 0)
	
	for entry := range lp.input {
		logBuffer = append(logBuffer, entry)
		
		// Process buffer when it has enough entries or input channel is empty
		if len(logBuffer) >= 10 {
			lp.flushLogs(&logBuffer)
		}
	}
	
	// Flush remaining logs
	lp.flushLogs(&logBuffer)
}

func (lp *LogProcessor) flushLogs(buffer *[]LogEntry) {
	if len(*buffer) == 0 {
		return
	}
	
	// Sort logs by timestamp for ordered output
	for i := 0; i < len(*buffer)-1; i++ {
		for j := i + 1; j < len(*buffer); j++ {
			if (*buffer)[i].Timestamp.After((*buffer)[j].Timestamp) {
				(*buffer)[i], (*buffer)[j] = (*buffer)[j], (*buffer)[i]
			}
		}
	}
	
	// Send to output in order
	for _, entry := range *buffer {
		formatted := fmt.Sprintf("[%v] %s [%s]: %s", 
			entry.Timestamp.Format("15:04:05.000"), 
			entry.Level, 
			entry.Source, 
			entry.Message)
		lp.output <- formatted
	}
	
	*buffer = (*buffer)[:0] // Clear buffer
}

func (lp *LogProcessor) Write(entry LogEntry) {
	lp.input <- entry
}

func (lp *LogProcessor) GetOutput() <-chan string {
	return lp.output
}

func (lp *LogProcessor) Stop() {
	close(lp.input)
	lp.wg.Wait()
}

func logWriter(source string, processor *LogProcessor, wg *sync.WaitGroup) {
	defer wg.Done()
	
	levels := []string{"INFO", "WARN", "ERROR", "DEBUG"}
	messages := []string{
		"User login successful",
		"Database connection slow",
		"Authentication failed",
		"Cache updated",
		"Network timeout",
		"Service started",
	}
	
	for i := 0; i < 10; i++ {
		entry := LogEntry{
			Timestamp: time.Now().Add(time.Duration(i) * time.Millisecond),
			Level:     levels[i%len(levels)],
			Message:   messages[i%len(messages)],
			Source:    source,
		}
		
		processor.Write(entry)
		time.Sleep(time.Duration(source[7]-'0') * 50 * time.Millisecond) // Different rates
	}
}

func main() {
	processor := NewLogProcessor()
	processor.Start()
	
	fmt.Println("=== Log Processor: Multiple Writers, Ordered Output ===")
	
	// Start multiple log writers
	wg := sync.WaitGroup{}
	sources := []string{"Service1", "Service2", "Service3"}
	
	for _, source := range sources {
		wg.Add(1)
		go logWriter(source, processor, &wg)
	}
	
	// Collect and print ordered output
	go func() {
		for logLine := range processor.GetOutput() {
			fmt.Println(logLine)
		}
	}()
	
	// Wait for all writers to finish
	wg.Wait()
	
	// Stop processor
	processor.Stop()
	
	fmt.Println("Log processing completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle multiple log writers but maintain ordered output?

**Your Response:** I implement a log processor that collects logs from multiple sources and outputs them in timestamp order. Multiple services write to a single input channel concurrently.

The processor uses a buffering strategy - it collects logs until it has enough entries, then sorts them by timestamp before outputting. This ensures chronological order despite concurrent writes from different sources.

Each log writer sends entries with timestamps to the processor. The processor maintains a buffer, sorts entries by timestamp, and sends formatted log lines to the output channel. This creates a centralized logging system with ordered output.

The key insight is separating concurrent collection from ordered processing. Writers can send logs at different rates without blocking each other, while the processor ensures the final output is properly ordered.

This pattern is essential for real-world logging systems where multiple services generate logs concurrently but need unified, chronological output for debugging and auditing. It demonstrates understanding of concurrent data collection with ordered processing requirements.
