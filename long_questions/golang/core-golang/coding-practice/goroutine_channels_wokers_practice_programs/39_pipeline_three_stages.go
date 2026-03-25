package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value   int
	Stage   string
	History []string
}

func generator(output chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for i := 1; i <= 10; i++ {
		data := Data{
			Value:   i,
			Stage:   "generated",
			History: []string{"generated"},
		}
		output <- data
		fmt.Printf("Generator: produced %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func processor(input <-chan Data, output chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for data := range input {
		// Process the data
		processed := data.Value * 2
		data.Value = processed
		data.Stage = "processed"
		data.History = append(data.History, "processed")
		
		output <- data
		fmt.Printf("Processor: %d -> %d\n", processed/2, processed)
		time.Sleep(300 * time.Millisecond)
	}
}

func outputter(input <-chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for data := range input {
		fmt.Printf("Output: final value %d (stages: %v)\n", 
			data.Value, data.History)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	wg := sync.WaitGroup{}
	
	// Pipeline channels
	genToProc := make(chan Data, 3)
	procToOut := make(chan Data, 3)

	fmt.Println("Starting 3-stage pipeline: generator -> processor -> outputter")

	// Start pipeline stages
	wg.Add(3)
	go generator(genToProc, &wg)
	go processor(genToProc, procToOut, &wg)
	go outputter(procToOut, &wg)

	wg.Wait()
	fmt.Println("Pipeline completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain the three-stage pipeline pattern and how data flows through it?

**Your Response:** The pipeline pattern processes data through multiple stages, with each stage performing a specific transformation. I implement it as three connected goroutines with channels between them.

Stage 1 (generator) creates initial data and sends it to the processor. Stage 2 (processor) transforms the data by doubling the value and passes it to the outputter. Stage 3 (outputter) handles the final results.

The channels act as conveyor belts between stages - genToProc carries data from generator to processor, procToOut carries data from processor to outputter. Each stage uses the range pattern to automatically handle channel closure.

The key insight is that pipeline stages can work concurrently - while the processor is working on item 3, the generator can be producing item 4, and the outputter can be handling item 2. This creates efficient flow-through processing.

This pattern is extremely useful for data processing workflows, ETL pipelines, or any scenario where you need to apply multiple transformations to data streams. It demonstrates understanding of concurrent data flow patterns.
