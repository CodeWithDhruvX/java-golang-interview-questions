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
