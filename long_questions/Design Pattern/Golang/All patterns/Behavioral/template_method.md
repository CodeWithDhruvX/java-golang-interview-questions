# Template Method Pattern

## 🟢 What is it?
The **Template Method Pattern** defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps.

Think of it like **Building a House**: Foundation -> Walls -> Roof (Standard). Walls can be Wood or Brick (Variable).

---

## 🎯 Strategy to Implement

Go doesn't have class inheritance. We use **Composition** and **Interfaces**.
1.  **Interface**: Define the steps that vary.
2.  **Template Function**: A function (or method on a base struct) that takes the Interface and executes the steps in order.

---

## 💻 Code Example

```go
package main

import "fmt"

// 1. Interface for steps that vary
type DataMinerOps interface {
    ExtractData()
    ParseData()
    AnalyzeData() // Hook
}

// 2. Template Logic
// We can use a struct to hold common logic if needed
type DataMiner struct {}

func (d *DataMiner) Mine(ops DataMinerOps) {
    d.Connect()
    ops.ExtractData()
    ops.ParseData()
    ops.AnalyzeData()
    d.Disconnect()
}

func (d *DataMiner) Connect() { fmt.Println("Connecting to data source...") }
func (d *DataMiner) Disconnect() { fmt.Println("Disconnecting.") }

// 3. Concrete Implementations
type PDFMiner struct {}
func (p *PDFMiner) ExtractData() { fmt.Println("Extracting text from PDF...") }
func (p *PDFMiner) ParseData() { fmt.Println("Parsing PDF structure...") }
func (p *PDFMiner) AnalyzeData() { fmt.Println("Basic analysis.") }

type CSVMiner struct {}
func (c *CSVMiner) ExtractData() { fmt.Println("Reading CSV lines...") }
func (c *CSVMiner) ParseData() { fmt.Println("Splitting CSV commas...") }
func (c *CSVMiner) AnalyzeData() { fmt.Println("Running statistical analysis.") }

func main() {
    miner := &DataMiner{}

    fmt.Println("--- Processing PDF ---")
    miner.Mine(&PDFMiner{})

    fmt.Println("\n--- Processing CSV ---")
    miner.Mine(&CSVMiner{})
}
```

---

## ✅ When to use?

*   **Standard Process**: Fixed sequence, variable steps.
*   **Frameworks**: Lifecycle hooks.
