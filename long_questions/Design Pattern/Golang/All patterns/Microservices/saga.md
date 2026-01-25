# SAGA Pattern

## 🟢 What is it?
The **SAGA Pattern** manages data consistency across microservices in distributed transactions. Since you cannot use a traditional ACID database transaction across two different services (and databases), a SAGA breaks the transaction into a sequence of local transactions.

If one step fails, the SAGA executes **Compensating Transactions** (Undo operations) to roll back the changes made by the preceding steps.

---

## 🏛️ Real World Analogy
**Booking a Trip**:
1.  **Book Flight** (Success) -> "Flight Confirmed"
2.  **Book Hotel** (Success) -> "Hotel Confirmed"
3.  **Book Rental Car** (Failed - No cars available)

Since Step 3 failed:
*   **Undo Booking Hotel** (Cancel Reservation)
*   **Undo Booking Flight** (Refund Ticket)
*   **End State**: Clean slate, no money lost.

---

## 🎯 Strategy to Implement

There are two main approaches:
1.  **Choreography**: Events based. Service A emits "Order Created", Service B listens and emits "Payment Processed". Decentralized.
2.  **Orchestration**: A central Coordinator (State Machine) tells Service A what to do, waits for result, then tells Service B. Centralized.

**Orchestration is easier to reason about.**

---

## 💻 Code Example (Simple Orchestration)

```go
package main

import (
	"errors"
	"fmt"
)

// The Orchestrator
func ProcessOrderSaga(orderID int) {
	fmt.Printf("Starting SAGA for Order %d\n", orderID)

	// Step 1: Inventory Service
	err := ReserveInventory(orderID)
	if err != nil {
		fmt.Println("Inventory Failed. Stop.")
		return
	}

	// Step 2: Payment Service
	err = ChargeCreditCard(orderID)
	if err != nil {
		fmt.Println("Payment Failed. Compensating Inventory...")
		CompensateInventory(orderID) // UNDO Step 1
		return
	}

	// Step 3: Shipping Service
	err = ScheduleShipping(orderID)
	if err != nil {
		fmt.Println("Shipping Failed. Compensating Payment & Inventory...")
		CompensatePayment(orderID)   // UNDO Step 2
		CompensateInventory(orderID) // UNDO Step 1
		return
	}

	fmt.Println("SAGA Completed Successfully!")
}

// --- Services ---

func ReserveInventory(id int) error {
	fmt.Println("[1] Inventory Reserved")
	return nil
}
func CompensateInventory(id int) {
	fmt.Println("[!1] Inventory Released (Undo)")
}

func ChargeCreditCard(id int) error {
	// Simulate Failure for demo
	// return errors.New("insufficient funds")
	fmt.Println("[2] Credit Card Charged")
	return nil
}
func CompensatePayment(id int) {
	fmt.Println("[!2] Payment Refunded (Undo)")
}

func ScheduleShipping(id int) error {
	// Simulate Failure
	return errors.New("courier unavailable")
}

func main() {
	ProcessOrderSaga(101)
}
```

---

## ✅ When to use?

*   **E-Commerce Checkout**: Order Service -> Inventory Service -> Payment Service.
*   **Bank Transfers**: Debit Account A -> Credit Account B (Different banks/services).
*   **Long Running Processes**: Video processing pipelines where an error in the final encoding step requires cleaning up temporary files created in step 1.
