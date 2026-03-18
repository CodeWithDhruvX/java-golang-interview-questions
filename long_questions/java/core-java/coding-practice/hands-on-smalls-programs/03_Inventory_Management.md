# Mini-Project 3: Inventory Management System

**Goal**: Demonstrate logic building, calculations, and POJO management.

## Features
1.  Add Products
2.  Update Stock
3.  Calculate Total Inventory Value

## Code Implementation

```java
import java.util.*;

class Product {
    int id;
    String name;
    double price;
    int quantity;

    public Product(int id, String name, double price, int quantity) {
        this.id = id;
        this.name = name;
        this.price = price;
        this.quantity = quantity;
    }

    public double getTotalValue() {
        return price * quantity;
    }

    @Override
    public String toString() {
        return String.format("ID: %d | %-10s | Price: $%.2f | Qty: %d | Value: $%.2f", 
                id, name, price, quantity, getTotalValue());
    }
}

public class InventoryApp {
    private static List<Product> inventory = new ArrayList<>();

    public static void main(String[] args) {
        // Pre-populating some data
        inventory.add(new Product(101, "Laptop", 800.0, 5));
        inventory.add(new Product(102, "Mouse", 20.0, 50));
        inventory.add(new Product(103, "KeyBoard", 50.0, 20));

        Scanner sc = new Scanner(System.in);
        while(true) {
            System.out.println("\n--- Inventory Menu ---");
            System.out.println("1. View Inventory");
            System.out.println("2. Add Product");
            System.out.println("3. Total Asset Value");
            System.out.println("4. Exit");
            System.out.print("Select: ");
            
            int choice = sc.nextInt();
            if (choice == 4) break;

            switch(choice) {
                case 1:
                    System.out.println("\nCurrent Inventory:");
                    inventory.forEach(System.out::println);
                    break;
                case 2:
                    System.out.print("ID: "); int id = sc.nextInt();
                    System.out.print("Name: "); String name = sc.next();
                    System.out.print("Price: "); double price = sc.nextDouble();
                    System.out.print("Qty: "); int qty = sc.nextInt();
                    inventory.add(new Product(id, name, price, qty));
                    System.out.println("Product Added.");
                    break;
                case 3:
                    double totalValue = inventory.stream()
                        .mapToDouble(Product::getTotalValue)
                        .sum();
                    System.out.printf("\nTotal Inventory Value: $%.2f\n", totalValue);
                    break;
            }
        }
    }
}
```

## Key Code Concepts Used
*   **formatted String**: `String.format` for clean table-like output.
*   **Stream API**: `mapToDouble().sum()` for efficient calculation.
*   **POJO Class**: Simple data carrier class `Product`.

---

## 📋 Interview Questions

### **Design & Architecture Questions**

**Q1: Why is the `Product` class designed as a POJO (Plain Old Java Object)?**
**A**: "I designed Product as a POJO because it's a simple data carrier class that represents a real-world entity with minimal complexity. POJOs don't extend any framework classes or implement special interfaces, making them lightweight and easy to test. The Product class just holds data (id, name, price, quantity) and provides basic business logic (getTotalValue) and a nice toString representation. This approach keeps the code simple, maintainable, and framework-independent."

**Q2: Why are the Product fields `package-private` instead of `private`?**
**A**: "In this simple example, I made them package-private for brevity since all the code is in the same package. But in production code, I should make them private and provide proper getter/setter methods. Private fields would follow encapsulation principles and allow me to add validation logic - like ensuring price is positive or quantity is not negative. Package-private access is acceptable for learning examples but not recommended for production systems."

**Q3: What's the benefit of having a `getTotalValue()` method in Product?**
**A**: "The getTotalValue() method encapsulates the business logic of calculating the total value of a product's inventory. Instead of putting the calculation logic in the main class, I keep it in the Product class where it belongs. This follows the principle of keeping related data and behavior together. It also makes the code more reusable - I can call this method from anywhere without duplicating the price * quantity calculation."

### **Stream API & Functional Programming Questions**

**Q4: How does the Stream API calculation work for total inventory value?**
**A**: "The stream API provides a clean, declarative way to calculate the total value. `inventory.stream()` creates a stream of Product objects, `mapToDouble(Product::getTotalValue)` converts each Product to its total value using a method reference, and `sum()` adds all the values together. This approach is more readable and less error-prone than a traditional for loop with manual accumulation. It also allows for easy parallelization if needed with `parallelStream()`."

**Q5: Why use `mapToDouble()` instead of `map()` in the stream?**
**A**: "I used `mapToDouble()` because it returns a `DoubleStream` which has specialized numeric operations like `sum()`, `average()`, `min()`, and `max()`. If I used regular `map()`, I'd get a `Stream<Double>` and would need to use `reduce(0.0, Double::sum)` or convert to a primitive stream anyway. `mapToDouble()` is more efficient because it avoids boxing/unboxing overhead and provides the numeric operations I need directly."

**Q6: How would you modify this to sort products by value?**
**A**: "I could use the stream API with sorting: `inventory.stream().sorted(Comparator.comparing(Product::getTotalValue).reversed()).forEach(System.out::println)`. This would sort products by their total value in descending order. For more complex sorting, I could create a custom Comparator or use `thenComparing()` for multi-level sorting. The stream API makes sorting operations very expressive and readable."

### **Data Structure & Performance Questions**

**Q7: Why use `ArrayList` for inventory storage?**
**A**: "I chose ArrayList because it provides good performance for this inventory management use case. ArrayList offers O(1) access time for displaying products, and efficient iteration for calculations. Since inventory operations typically involve adding new products and displaying all products (sequential access), ArrayList is optimal. If I needed frequent insertions/deletions in the middle of the list, LinkedList might be better, but that's not common in inventory scenarios."

**Q8: How would this perform with 1 million products?**
**A**: "With 1 million products, the current approach might face memory and performance challenges. The ArrayList would consume significant memory, and the stream calculation might become slow. For large datasets, I'd consider: 1) Database storage with proper indexing, 2) Pagination for displaying products, 3) Caching frequently accessed products, 4) Batch processing for calculations. The stream approach is still efficient, but I'd need to optimize memory usage and possibly use parallel streams."

**Q9: What's the time complexity of adding a product vs calculating total value?**
**A**: "Adding a product is O(1) amortized time for ArrayList - it's very fast unless resizing is needed. Calculating total value is O(n) where n is the number of products, since I need to visit each product once. The stream API makes this operation efficient, but it's still linear time. For frequent total value calculations, I might consider maintaining a running total that gets updated when products are added or modified."

### **Business Logic & Validation Questions**

**Q10: How would you add validation for product data?**
**A**: "I'd add validation in the Product constructor or setter methods. For example, ensure price is positive and quantity is non-negative. I could throw `IllegalArgumentException` for invalid values: `if (price <= 0) throw new IllegalArgumentException("Price must be positive")`. For the name, I'd check it's not empty and maybe validate length. This prevents invalid data from entering the system and maintains data integrity."

**Q11: How would you handle stock updates (increment/decrement)?**
**A**: "I'd add methods like `addStock(int quantity)` and `removeStock(int quantity)` in the Product class. These methods would validate that the quantity doesn't go negative and maybe throw a custom `InsufficientStockException`. I could also add a `isInStock()` method to check if quantity > 0. This encapsulates the stock management logic in the Product class where it belongs, rather than scattering it throughout the code."

**Q12: How would you add support for product categories?**
**A**: "I'd add a `category` field to the Product class and maybe create a `Category` enum. For better organization, I might use a `Map<Category, List<Product>>` to group products by category. This would allow me to easily calculate inventory value by category or display products by category. I could also add methods like `getProductsByCategory()` to the main class for category-based operations."

### **User Interface & Experience Questions**

**Q13: Why use `String.format()` in the toString() method?**
**A**: "String.format() provides clean, formatted output with proper alignment and fixed decimal places for currency values. The format specifiers like `%-10s` for left-aligned strings and `%.2f` for two decimal places create a table-like appearance that's easy to read. This is much better than simple string concatenation for displaying tabular data and makes the output look professional and consistent."

**Q14: How would you add search functionality to find products by name?**
**A**: "I'd add a search method using streams: `inventory.stream().filter(p -> p.name.toLowerCase().contains(searchTerm.toLowerCase())).collect(Collectors.toList())`. For better performance with large inventories, I might maintain an additional Map structure for name-based lookups. I'd also add a menu option for search and implement case-insensitive searching to make it more user-friendly."

**Q15: How would you implement pagination for displaying large inventories?**
**A**: "I'd add methods like `displayProducts(int page, int pageSize)` that use stream operations: `inventory.stream().skip(page * pageSize).limit(pageSize).forEach(System.out::println)`. This would allow users to navigate through large inventories without overwhelming the console. I'd also show total pages and current page information, and add navigation options like next/previous page."

### **Data Persistence & Storage Questions**

**Q16: How would you add data persistence to save inventory?**
**A**: "I'd implement file I/O similar to the Student Management system. I could use serialization to save the entire ArrayList to a file, or use CSV/JSON format for better interoperability. For production, I'd use a database with proper tables for products. I'd also add auto-save functionality and maybe implement a change journal to track all modifications for audit purposes."

**Q17: What are the limitations of in-memory storage for inventory?**
**A**: "In-memory storage has several limitations: 1) Data is lost when the application closes, 2) Limited by available RAM, 3) No concurrent access support, 4) No backup or recovery mechanisms, 5) Security risks as data is not encrypted at rest. For a real inventory system, I'd need persistent storage like a database that can handle large datasets, multiple users, and provide data integrity guarantees."

### **Advanced Features & Extensions Questions**

**Q18: How would you add support for low stock alerts?**
**A**: "I'd add a `reorderLevel` field to Product and a method like `isLowStock()` that checks if quantity <= reorderLevel. I could implement a monitoring service that periodically checks all products and generates alerts. For real-time alerts, I could use the Observer pattern where interested parties get notified when stock levels change. This would help with inventory management and prevent stockouts."

**Q19: How would you implement inventory tracking across multiple warehouses?**
**A**: "I'd create a more complex data model with Warehouse and InventoryItem classes. Each Product would have a Map<Warehouse, Integer> to track quantities by location. I'd add methods like `getTotalQuantity()` and `getQuantityByWarehouse()`. For reporting, I could provide warehouse-specific views and transfer operations to move stock between locations."

**Q20: How would you add support for product suppliers and purchase orders?**
**A**: "I'd create Supplier and PurchaseOrder classes with relationships to Product. The Product might have a List<Supplier> for approved suppliers and methods for reordering. I could implement automatic reordering when stock falls below threshold, generating purchase orders with recommended quantities based on historical usage patterns. This would make the system more comprehensive for real business use."

**Q21: How would you implement audit logging for inventory changes?**
**A**: "I'd create an AuditLog class to track all changes with timestamps, user information, and old/new values. Every operation that modifies inventory would create an audit entry. I could use the Command pattern to wrap operations and ensure logging happens automatically. This would provide traceability for compliance and help with troubleshooting inventory discrepancies."

**Q22: How would you add reporting capabilities for sales and inventory analysis?**
**A**: "I'd implement a reporting module that uses the stream API for data analysis. For example, `getTopSellingProducts()`, `getInventoryTurnover()`, or `getDeadStock()`. I could generate reports in different formats (console, CSV, PDF) and add filtering by date ranges or categories. This would provide business insights for inventory optimization and purchasing decisions."
