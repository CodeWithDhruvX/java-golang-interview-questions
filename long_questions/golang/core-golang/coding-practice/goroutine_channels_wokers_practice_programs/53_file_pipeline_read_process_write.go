package main

import (
	"fmt"
	"sync"
	"time"
)

type FileData struct {
	Filename string
	Content  string
	Stage    string
}

type FilePipeline struct {
	readChan   chan FileData
	processChan chan FileData
	writeChan  chan FileData
	wg         sync.WaitGroup
}

func NewFilePipeline() *FilePipeline {
	return &FilePipeline{
		readChan:    make(chan FileData, 10),
		processChan: make(chan FileData, 10),
		writeChan:   make(chan FileData, 10),
	}
}

func (fp *FilePipeline) Start() {
	fp.wg.Add(3)
	go fp.reader()
	go fp.processor()
	go fp.writer()
}

func (fp *FilePipeline) reader() {
	defer fp.wg.Done()
	defer close(fp.readChan)
	
	// Simulate reading files
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}
	
	for _, filename := range files {
		time.Sleep(200 * time.Millisecond) // Simulate I/O
		
		data := FileData{
			Filename: filename,
			Content:  fmt.Sprintf("Original content of %s", filename),
			Stage:    "read",
		}
		
		fp.readChan <- data
		fmt.Printf("Reader: read %s\n", filename)
	}
}

func (fp *FilePipeline) processor() {
	defer fp.wg.Done()
	defer close(fp.processChan)
	
	for data := range fp.readChan {
		time.Sleep(300 * time.Millisecond) // Simulate processing
		
		// Process the content
		data.Content = fmt.Sprintf("PROCESSED: %s", data.Content)
		data.Stage = "processed"
		
		fp.processChan <- data
		fmt.Printf("Processor: processed %s\n", data.Filename)
	}
}

func (fp *FilePipeline) writer() {
	defer fp.wg.Done()
	
	for data := range fp.processChan {
		time.Sleep(150 * time.Millisecond) // Simulate I/O
		
		// Write the processed data
		fmt.Printf("Writer: wrote %s - %s\n", data.Filename, data.Content)
	}
}

func (fp *FilePipeline) Stop() {
	fp.wg.Wait()
}

func main() {
	pipeline := NewFilePipeline()
	
	fmt.Println("=== File Pipeline: Read -> Process -> Write ===")
	
	startTime := time.Now()
	
	// Start pipeline
	pipeline.Start()
	
	// Wait for pipeline to complete
	pipeline.Stop()
	
	totalTime := time.Since(startTime)
	fmt.Printf("\nPipeline completed in %v\n", totalTime)
	fmt.Printf("Sequential would take: %v\n", 
		5*(200+300+150)*time.Millisecond)
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a file processing pipeline with read, process, and write stages?

**Your Response:** I implement a three-stage pipeline where each stage is a separate goroutine connected by channels. The pipeline processes files concurrently through the stages: reader → processor → writer.

The reader stage simulates reading files and sends data to the process channel. The processor stage transforms the content and passes it to the write channel. The writer stage handles the final output.

Each stage works independently - while the processor is working on file 1, the reader can be reading file 2, and the writer can be writing file 3. This creates efficient pipeline parallelism where all stages can be busy simultaneously.

The key insight is that pipeline stages are decoupled by channels, allowing each to work at its own pace. The channels act as buffers that smooth out rate differences between stages.

This pattern is extremely useful for data processing workflows, ETL operations, image processing, or any scenario where data needs to flow through multiple processing steps. The performance gain comes from overlapping the work of different stages instead of doing them sequentially.
