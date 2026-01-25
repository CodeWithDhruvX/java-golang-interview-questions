# Pipeline Pattern

## 🟢 What is it?
The **Pipeline Pattern** is a series of stages connected by channels, where each stage is a group of goroutines running the same function. In each stage, the workers:
1.  Receive values from upstream via an inbound channel.
2.  Perform some function on that data (process/transform).
3.  Send values downstream via an outbound channel.

This creates a "Streaming" architecture where data flows through the system.

---

## 🏛️ Real World Analogy
**Car Manufacturing Assembly Line**:
*   **Stage 1 (Chassis)**: Robots weld the frame.
*   **Stage 2 (Paint)**: Robots spray paint the frame.
*   **Stage 3 (Engine)**: Humans install the engine.
*   The Chassis team doesn't wait for the Engine to be installed. They just finish their job and pass it to the Paint team. All stages happen simultaneously on different cars.

---

## 🎯 Strategy to Implement

1.  **Stage Generator**: A function that converts a list of data into a channel (the source).
2.  **Processing Stages**: Functions that take an input channel (`<-chan Type`) and return an output channel (`<-chan Type`).
    *   They start a goroutine internally.
    *   They loop over input, process, and send to output.
    *   They close the output channel when input closes.
3.  **Chaining**: In `main`, you connect them: `out1 := Stage1(source)`; `out2 := Stage2(out1)`.

---

## 💻 Code Example

```go
package main

import (
	"fmt"
)

// Stage 1: Generator - Converts a list of numbers to a channel
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Square - Squares the input numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Stage 3: AddOne - Adds 1 to the input numbers
func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline
	// [ 2, 3, 4 ] -> [ generate ] -> [ square ] -> [ addOne ] -> Result
	
	// 1. Source
	c1 := generate(2, 3, 4)

	// 2. Stage A (Square)
	c2 := square(c1)

	// 3. Stage B (Add One)
	c3 := addOne(c2)

	// 4. Sink (Consume results)
	for result := range c3 {
		fmt.Println(result)
	}
	
	// Output Trace:
	// 2 -> 4 -> 5
	// 3 -> 9 -> 10
	// 4 -> 16 -> 17
}
```

---

## ✅ When to use?

*   **ETL Jobs (Extract, Transform, Load)**: Reading a file (Stage 1), parsing JSON (Stage 2), filtering bad data (Stage 3), saving to DB (Stage 4).
*   **Stream Processing**: Manipulating video/audio frames or large logs where you can't load everything into memory.
*   **Modular logic**: It separates concerns cleanly (Squaring logic is totally separate from Generation logic).
