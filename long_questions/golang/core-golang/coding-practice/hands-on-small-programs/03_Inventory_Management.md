# Mini-Project 3: Inventory Management System

**Goal**: Demonstrate logic building, calculations, and struct management.

## Features
1.  Add Products
2.  Update Stock
3.  Calculate Total Inventory Value

## Code Implementation

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Product Struct
type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(id int, name string, price float64, quantity int) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func (p *Product) GetTotalValue() float64 {
	return p.Price * float64(p.Quantity)
}

func (p *Product) String() string {
	return fmt.Sprintf("ID: %d | %-10s | Price: $%.2f | Qty: %d | Value: $%.2f",
		p.ID, p.Name, p.Price, p.Quantity, p.GetTotalValue())
}

type Inventory struct {
	products []*Product
}

func NewInventory() *Inventory {
	// Pre-populating some data
	return &Inventory{
		products: []*Product{
			NewProduct(101, "Laptop", 800.0, 5),
			NewProduct(102, "Mouse", 20.0, 50),
			NewProduct(103, "Keyboard", 50.0, 20),
		},
	}
}

func (inv *Inventory) AddProduct(product *Product) {
	inv.products = append(inv.products, product)
}

func (inv *Inventory) DisplayAll() {
	fmt.Println("Current Inventory:")
	for _, product := range inv.products {
		fmt.Println(product)
	}
}

func (inv *Inventory) GetTotalValue() float64 {
	total := 0.0
	for _, product := range inv.products {
		total += product.GetTotalValue()
	}
	return total
}

func (inv *Inventory) FindProduct(id int) *Product {
	for _, product := range inv.products {
		if product.ID == id {
			return product
		}
	}
	return nil
}

func (inv *Inventory) UpdateStock(id int, quantity int) bool {
	product := inv.FindProduct(id)
	if product == nil {
		return false
	}
	
	newQuantity := product.Quantity + quantity
	if newQuantity < 0 {
		return false // Can't have negative stock
	}
	
	product.Quantity = newQuantity
	return true
}

func main() {
	inventory := NewInventory()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Inventory Menu ---")
		fmt.Println("1. View Inventory")
		fmt.Println("2. Add Product")
		fmt.Println("3. Update Stock")
		fmt.Println("4. Total Asset Value")
		fmt.Println("5. Exit")
		fmt.Print("Select: ")

		if !scanner.Scan() {
			break
		}

		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		switch choice {
		case 1:
			inventory.DisplayAll()

		case 2:
			fmt.Print("ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID!")
				continue
			}

			fmt.Print("Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Price: ")
			scanner.Scan()
			price, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid price!")
				continue
			}

			fmt.Print("Quantity: ")
			scanner.Scan()
			qty, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid quantity!")
				continue
			}

			// Check for duplicate ID
			if inventory.FindProduct(id) != nil {
				fmt.Println("Product with this ID already exists!")
				continue
			}

			inventory.AddProduct(NewProduct(id, name, price, qty))
			fmt.Println("Product Added.")

		case 3:
			fmt.Print("Product ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID!")
				continue
			}

			fmt.Print("Quantity Change (+/-): ")
			scanner.Scan()
			qty, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid quantity!")
				continue
			}

			if inventory.UpdateStock(id, qty) {
				fmt.Println("Stock updated successfully.")
			} else {
				fmt.Println("Failed to update stock. Product not found or insufficient stock.")
			}

		case 4:
			totalValue := inventory.GetTotalValue()
			fmt.Printf("\nTotal Inventory Value: $%.2f\n", totalValue)

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
```

## Key Code Concepts Used
*   **Formatted String**: `fmt.Sprintf` for clean table-like output.
*   **Struct Methods**: Methods attached to Product struct for business logic.
*   **Composition**: Inventory struct contains slice of Product pointers.
*   **Pointer Semantics**: Using pointers for efficient memory usage and modifications.

---

## 📋 Interview Questions

### **Design & Architecture Questions**

**Q1: Why is the `Product` struct designed as a simple data carrier?**
**A**: "I designed Product as a simple data carrier struct because it represents a real-world entity with minimal complexity. Go encourages this approach - simple structs with clear responsibilities. The Product struct just holds data (id, name, price, quantity) and provides basic business logic (GetTotalValue) and a nice String representation. This approach keeps the code simple, maintainable, and follows Go's philosophy of favoring composition over complex inheritance hierarchies."

**Q2: Why are the Product fields public instead of private?**
**A**: "In Go, field visibility is controlled by capitalization. I made the fields public (uppercase) because this is a simple example where all code is in the same package. For production code, I might make them private (lowercase) and provide getter/setter methods to follow encapsulation principles. However, Go's approach to encapsulation is package-based rather than class-based, so private fields would still be accessible within the same package."

**Q3: What's the benefit of having a `GetTotalValue()` method in Product?**
**A**: "The GetTotalValue() method encapsulates the business logic of calculating the total value of a product's inventory. Instead of putting the calculation logic in the main struct, I keep it in the Product struct where it belongs. This follows the principle of keeping related data and behavior together. It also makes the code more reusable - I can call this method from anywhere without duplicating the price * quantity calculation."

**Q4: Why use `[]*Product` (slice of pointers) instead of `[]Product`?**
**A**: "I used a slice of pointers because it's more memory-efficient when dealing with large collections. When you have a slice of structs, each operation like appending or passing elements copies the entire struct. With pointers, you only copy the pointer (8 bytes on 64-bit systems), regardless of how large the struct is. This is especially important for inventory systems that might handle thousands of products. Pointers also allow for easier modification of individual products without needing to reassign them."

### **Go-Specific Programming Questions**

**Q5: How does Go's method receiver syntax compare to Java's methods?**
**A**: "Go's method receiver syntax `(p *Product)` is similar to instance methods in Java, but with some key differences. The receiver can be either a value receiver `(p Product)` or a pointer receiver `(p *Product)`. I used pointer receivers because they allow modification of the struct and avoid copying. Unlike Java where methods are always called on objects, Go methods can be called on both values and pointers, and the compiler handles the dereferencing automatically."

**Q6: Why use `fmt.Sprintf` instead of string concatenation in the String() method?**
**A**: "I used `fmt.Sprintf` because it provides clean, formatted output with proper alignment and fixed decimal places for currency values. The format specifiers like `%-10s` for left-aligned strings and `%.2f` for two decimal places create a table-like appearance that's easy to read. This is much better than string concatenation for displaying tabular data and makes the output look professional and consistent."

**Q7: What is the purpose of the `String() string` method?**
**A**: "The `String() string` method implements Go's `Stringer` interface, which is the Go equivalent of Java's `toString()` method. When I use `fmt.Println(product)`, Go automatically calls the `String()` method. This makes the objects work naturally with Go's formatting and printing functions. It's the idiomatic way to provide string representations of structs in Go."

**Q8: How does Go's error handling in this system compare to Java's exceptions?**
**A**: "Go uses explicit error returns instead of exceptions. Functions that can fail return an error as their last return value, and callers must check this error. In the inventory system, I use boolean return values for simple success/failure cases, but for more complex operations, I'd return `(result, error)`. This approach makes error handling more visible and explicit in the code, unlike Java's try-catch blocks where errors can be caught far from where they occur."

### **Data Structure & Performance Questions**

**Q9: Why use a slice for inventory storage instead of a map?**
**A**: "I chose a slice because it provides good performance for this inventory management use case. Slices offer O(1) access time for displaying products when iterating sequentially, and efficient append operations for adding new products. Since inventory operations typically involve adding new products and displaying all products (sequential access), a slice is optimal. If I needed frequent random access by ID, a map[int]*Product would be better for O(1) lookups."

**Q10: How would this perform with 1 million products?**
**A**: "With 1 million products, the current approach might face memory and performance challenges. The slice would consume significant memory, and the linear search in `FindProduct` (O(n)) would become slow. For large datasets, I'd consider: 1) Using a map for O(1) product lookups, 2) Database storage with proper indexing, 3) Pagination for displaying products, 4) Caching frequently accessed products. The current linear iteration for total value calculation would still be O(n), but that's unavoidable since we need to visit each product."

**Q11: What's the time complexity of adding a product vs calculating total value?**
**A**: "Adding a product is O(1) amortized time for slice append - it's very fast unless resizing is needed. Finding a product is O(n) where n is the number of products, since I need to iterate through the slice. Calculating total value is also O(n) since I need to visit each product once. For frequent product lookups, I might maintain an additional map[int]*Product for O(1) access by ID."

**Q12: How would you implement more efficient product searching?**
**A**: "For efficient searching, I'd maintain a map[int]*Product indexed by product ID alongside the slice. This would give O(1) lookups by ID. For name-based searching, I could maintain a map[string][]*Product for products with the same name prefix. For complex queries, I'd move to a database with proper indexing. I could also implement binary search if I keep the slice sorted by ID, but that would require maintaining sort order."

### **Business Logic & Validation Questions**

**Q13: How would you add validation for product data?**
**A**: "I'd add validation in the `NewProduct` constructor or in the input handling code. For example, ensure price is positive and quantity is non-negative. I could return an error instead of the product for invalid data: `if price <= 0 { return nil, fmt.Errorf("price must be positive") }`. For the name, I'd check it's not empty and maybe validate length. This prevents invalid data from entering the system and maintains data integrity."

**Q14: How would you handle stock updates (increment/decrement) more robustly?**
**A**: "I'd add more sophisticated methods like `AddStock(id, quantity)` and `RemoveStock(id, quantity)` that include better validation. For example, check if removing stock would result in negative quantity and return an error. I could also add a `IsInStock()` method to check if quantity > 0. For inventory management, I might want to track stock movements with timestamps and reasons, which would require additional structs and methods."

**Q15: How would you add support for product categories?**
**A**: "I'd add a `Category` field to the Product struct and maybe create a `Category` type (string or enum). For better organization, I might use a `map[Category][]*Product` to group products by category. This would allow me to easily calculate inventory value by category or display products by category. I could also add methods like `GetProductsByCategory()` to the Inventory struct for category-based operations."

### **User Interface & Experience Questions**

**Q16: How would you implement search functionality to find products by name?**
**A**: "I'd add a search method that iterates through the products slice: `for _, product := range inv.products { if strings.Contains(strings.ToLower(product.Name), strings.ToLower(searchTerm)) { results = append(results, product) } }`. For better performance with large inventories, I might maintain an additional map structure for name-based lookups, or use a database with proper indexing."

**Q17: How would you implement pagination for displaying large inventories?**
**A**: "I'd add methods like `DisplayProducts(page, pageSize)` that use slice operations: `start := page * pageSize; end := start + pageSize; if end > len(inv.products) { end = len(inv.products) }; for i := start; i < end; i++ { fmt.Println(inv.products[i]) }`. This would allow users to navigate through large inventories without overwhelming the console. I'd also show total pages and current page information."

**Q18: How would you add sorting capabilities for the inventory display?**
**A**: "I'd add sorting methods using Go's `sort` package. For example: `sort.Slice(inv.products, func(i, j int) bool { return inv.products[i].Name < inv.products[j].Name })`. I could provide different sorting options - by name, price, quantity, or total value. I might add a menu option to choose the sort criteria and maintain the sorted order for subsequent displays."

### **Data Persistence & Storage Questions**

**Q19: How would you add data persistence to save inventory?**
**A**: "I'd implement file I/O similar to the Student Management system. I could use JSON encoding with `json.Marshal()` and `json.Unmarshal()` to save and load the entire products slice to a file. For production, I'd use a database with proper tables for products. I'd also add auto-save functionality and maybe implement a change journal to track all modifications for audit purposes."

**Q20: What are the limitations of in-memory storage for inventory?**
**A**: "In-memory storage has several limitations: 1) Data is lost when the application closes, 2) Limited by available RAM, 3) No concurrent access support, 4) No backup or recovery mechanisms, 5) Security risks as data is not encrypted at rest. For a real inventory system, I'd need persistent storage like a database that can handle large datasets, multiple users, and provide data integrity guarantees."

**Q21: How would you implement concurrent access to the inventory?**
**A**: "For thread safety, I'd add a `sync.RWMutex` to the Inventory struct and use `Lock()/Unlock()` for write operations (add product, update stock) and `RLock()/RUnlock()` for read operations (display, calculate total). This ensures that multiple goroutines can read simultaneously, but writes are exclusive. This prevents race conditions where multiple operations could modify the inventory simultaneously."

### **Advanced Features & Extensions Questions**

**Q22: How would you add support for low stock alerts?**
**A**: "I'd add a `ReorderLevel` field to Product and a method like `IsLowStock()` that checks if quantity <= reorderLevel. I could implement a monitoring function that periodically checks all products and generates alerts. For real-time alerts, I could use Go channels to notify interested parties when stock levels change. This would help with inventory management and prevent stockouts."

**Q23: How would you implement inventory tracking across multiple warehouses?**
**A**: "I'd create a more complex data model with Warehouse and InventoryItem structs. Each Product might have a map[Warehouse]int to track quantities by location. I'd add methods like `GetTotalQuantity()` and `GetQuantityByWarehouse()`. For reporting, I could provide warehouse-specific views and transfer operations to move stock between locations."

**Q24: How would you add support for product suppliers and purchase orders?**
**A**: "I'd create Supplier and PurchaseOrder structs with relationships to Product. The Product might have a []*Supplier for approved suppliers and methods for reordering. I could implement automatic reordering when stock falls below threshold, generating purchase orders with recommended quantities based on historical usage patterns. This would make the system more comprehensive for real business use."

**Q25: How would you implement audit logging for inventory changes?**
**A**: "I'd create an AuditLog struct to track all changes with timestamps, user information, and old/new values. Every operation that modifies inventory would create an audit entry. I could use the Command pattern to wrap operations and ensure logging happens automatically. This would provide traceability for compliance and help with troubleshooting inventory discrepancies."

**Q26: How would you add reporting capabilities for sales and inventory analysis?**
**A**: "I'd implement a reporting module that provides various analytics. For example, `GetTopSellingProducts()`, `GetInventoryTurnover()`, or `GetDeadStock()`. I could generate reports in different formats (console, CSV, JSON) and add filtering by date ranges or categories. This would provide business insights for inventory optimization and purchasing decisions."

**Q27: How would you implement a REST API for this inventory system?**
**A**: "I'd use Go's `net/http` package to create HTTP endpoints. I'd implement handlers for GET /products (list), POST /products (create), GET /products/{id} (get one), PUT /products/{id} (update), PATCH /products/{id}/stock (update stock), and DELETE /products/{id} (delete). I'd use JSON for request/response bodies and implement proper HTTP status codes and error handling."

---
