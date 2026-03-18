# Java Collections Framework — Real-World Practical Code Snippets

> **Topics:** Real-world scenarios using List, Set, Map, Queue, and their practical applications in business contexts

---

## 📋 Reading Progress

- [ ] **Section 1:** E-commerce & Shopping Cart (Q1–Q8)
- [ ] **Section 2:** User Management & Authentication (Q9–Q16)
- [ ] **Section 3:** Inventory & Product Management (Q17–Q24)
- [ ] **Section 4:** Order Processing & Queuing (Q25–Q32)

> 🔖 **Last read:** <!-- -->

---

## Section 1: E-commerce & Shopping Cart (Q1–Q8)

### 1. Shopping Cart — Remove Out-of-Stock Items
**Q: You have a shopping cart with items, some are out of stock. What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> cart = new LinkedHashMap<>();
        cart.put("Laptop", 1);
        cart.put("Mouse", 2);
        cart.put("Keyboard", 1);
        cart.put("Monitor", 0); // Out of stock
        
        Set<String> outOfStock = new HashSet<>();
        cart.forEach((item, quantity) -> {
            if (quantity == 0) outOfStock.add(item);
        });
        
        outOfStock.forEach(cart::remove);
        System.out.println(cart);
    }
}
```
**A:** `{Laptop=1, Mouse=2, Keyboard=1}`. Items with 0 quantity are removed from cart.

---

### 2. Product Recommendations — Find Similar Products
**Q: Based on user's purchase history, find recommended products. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Set<String> userPurchases = Set.of("Laptop", "Mouse", "Keyboard");
        Map<String, Set<String>> productRelations = Map.of(
            "Laptop", Set.of("Mouse", "Keyboard", "Monitor"),
            "Mouse", Set.of("Laptop", "Keyboard", "Mousepad"),
            "Keyboard", Set.of("Laptop", "Mouse", "Monitor")
        );
        
        Set<String> recommendations = userPurchases.stream()
            .flatMap(product -> productRelations.getOrDefault(product, Set.of()).stream())
            .filter(product -> !userPurchases.contains(product))
            .collect(Collectors.toSet());
        
        System.out.println(recommendations);
    }
}
```
**A:** `[Mousepad, Monitor]`. Products frequently bought together but not yet purchased.

---

### 3. Price Comparison — Find Best Deal
**Q: Compare prices from multiple vendors for the same product. What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> vendorPrices = new HashMap<>();
        vendorPrices.put("Amazon", 999);
        vendorPrices.put("BestBuy", 949);
        vendorPrices.put("Walmart", 979);
        vendorPrices.put("Target", 989);
        
        String bestVendor = vendorPrices.entrySet().stream()
            .min(Map.Entry.comparingByValue())
            .map(Map.Entry::getKey)
            .orElse("Unknown");
        
        int bestPrice = vendorPrices.getOrDefault(bestVendor, Integer.MAX_VALUE);
        System.out.println("Best deal: " + bestVendor + " at $" + bestPrice);
    }
}
```
**A:** `Best deal: BestBuy at $949`. Finds the vendor with lowest price.

---

### 4. Order History — Group by Month
**Q: Group orders by month for analytics. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    public static void main(String[] args) {
        List<LocalDate> orderDates = Arrays.asList(
            LocalDate.of(2024, 1, 15),
            LocalDate.of(2024, 1, 20),
            LocalDate.of(2024, 2, 5),
            LocalDate.of(2024, 2, 15),
            LocalDate.of(2024, 3, 10)
        );
        
        Map<YearMonth, Long> ordersByMonth = orderDates.stream()
            .collect(Collectors.groupingBy(
                YearMonth::from,
                Collectors.counting()
            ));
        
        ordersByMonth.forEach((month, count) -> 
            System.out.println(month + ": " + count + " orders"));
    }
}
```
**A:** 
```
2024-01: 2 orders
2024-02: 2 orders
2024-03: 1 orders
```

---

### 5. Customer Segmentation — Categorize by Purchase Amount
**Q: Categorize customers based on their total spending. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Double> customerSpending = new HashMap<>();
        customerSpending.put("Alice", 1500.0);
        customerSpending.put("Bob", 250.0);
        customerSpending.put("Charlie", 5000.0);
        customerSpending.put("Diana", 800.0);
        
        Map<String, List<String>> segments = customerSpending.entrySet().stream()
            .collect(Collectors.groupingBy(entry -> {
                double amount = entry.getValue();
                if (amount >= 1000) return "Premium";
                else if (amount >= 500) return "Regular";
                else return "Basic";
            }, Collectors.mapping(Map.Entry::getKey, Collectors.toList())));
        
        System.out.println(segments);
    }
}
```
**A:** `{Premium=[Alice, Charlie], Regular=[Diana], Basic=[Bob]}`. Customers categorized by spending.

---

### 6. Inventory Alert — Low Stock Detection
**Q: Identify products that need restocking. What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> inventory = new HashMap<>();
        inventory.put("Laptop", 5);
        inventory.put("Mouse", 2);
        inventory.put("Keyboard", 15);
        inventory.put("Monitor", 3);
        
        int threshold = 5;
        List<String> lowStock = inventory.entrySet().stream()
            .filter(entry -> entry.getValue() < threshold)
            .map(Map.Entry::getKey)
            .sorted()
            .toList();
        
        System.out.println("Restock needed: " + lowStock);
    }
}
```
**A:** `Restock needed: [Monitor, Mouse]`. Products below threshold need restocking.

---

### 7. Shopping Cart Merge — Combine Multiple Carts
**Q: Merge shopping carts from different sessions. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> cart1 = Map.of("Laptop", 1, "Mouse", 2);
        Map<String, Integer> cart2 = Map.of("Mouse", 1, "Keyboard", 1);
        Map<String, Integer> cart3 = Map.of("Laptop", 1, "Monitor", 1);
        
        Map<String, Integer> mergedCart = Stream.of(cart1, cart2, cart3)
            .flatMap(map -> map.entrySet().stream())
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue,
                Integer::sum
            ));
        
        System.out.println(mergedCart);
    }
}
```
**A:** `{Laptop=2, Mouse=3, Keyboard=1, Monitor=1}`. Quantities summed for same items.

---

### 8. Product Search — Filter by Multiple Criteria
**Q: Find products matching multiple search criteria. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Product(String name, String category, double price, int rating) {}
    
    public static void main(String[] args) {
        List<Product> products = Arrays.asList(
            new Product("Laptop", "Electronics", 999.99, 4),
            new Product("Mouse", "Electronics", 29.99, 5),
            new Product("Book", "Education", 19.99, 4),
            new Product("Keyboard", "Electronics", 79.99, 3)
        );
        
        List<Product> searchResults = products.stream()
            .filter(p -> p.category.equals("Electronics"))
            .filter(p -> p.price < 100)
            .filter(p -> p.rating >= 4)
            .sorted(Comparator.comparingDouble(Product::price))
            .toList();
        
        searchResults.forEach(p -> 
            System.out.println(p.name + " - $" + p.price));
    }
}
```
**A:** `Mouse - $29.99`. Electronics under $100 with rating 4+.

---

## Section 2: User Management & Authentication (Q9–Q16)

### 9. User Permissions — Role-Based Access Control
**Q: Check if user has required permissions. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Set<String> userPermissions = Set.of("READ", "WRITE", "DELETE");
        Set<String> requiredPermissions = Set.of("READ", "WRITE");
        
        boolean hasAllPermissions = requiredPermissions.stream()
            .allMatch(userPermissions::contains);
        
        Set<String> missingPermissions = requiredPermissions.stream()
            .filter(perm -> !userPermissions.contains(perm))
            .collect(Collectors.toSet());
        
        System.out.println("Access granted: " + hasAllPermissions);
        if (!hasAllPermissions) {
            System.out.println("Missing: " + missingPermissions);
        }
    }
}
```
**A:** `Access granted: true`. User has all required permissions.

---

### 10. Active Sessions — Track Logged-in Users
**Q: Track and clean up expired user sessions. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    public static void main(String[] args) {
        Map<String, LocalDateTime> activeSessions = new HashMap<>();
        activeSessions.put("user1", LocalDateTime.now().minusHours(2));
        activeSessions.put("user2", LocalDateTime.now().minusMinutes(30));
        activeSessions.put("user3", LocalDateTime.now().minusDays(1));
        
        LocalDateTime cutoff = LocalDateTime.now().minusHours(1);
        List<String> expiredUsers = activeSessions.entrySet().stream()
            .filter(entry -> entry.getValue().isBefore(cutoff))
            .map(Map.Entry::getKey)
            .toList();
        
        expiredUsers.forEach(activeSessions::remove);
        
        System.out.println("Expired sessions: " + expiredUsers);
        System.out.println("Active users: " + activeSessions.keySet());
    }
}
```
**A:** 
```
Expired sessions: [user1, user3]
Active users: [user2]
```

---

### 11. User Activity — Most Active Users
**Q: Find users with highest activity counts. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Integer> userActivity = new HashMap<>();
        userActivity.put("Alice", 150);
        userActivity.put("Bob", 75);
        userActivity.put("Charlie", 200);
        userActivity.put("Diana", 120);
        
        List<String> topUsers = userActivity.entrySet().stream()
            .sorted(Map.Entry.<String, Integer>comparingByValue().reversed())
            .limit(3)
            .map(Map.Entry::getKey)
            .toList();
        
        topUsers.forEach(user -> 
            System.out.println(user + ": " + userActivity.get(user) + " actions"));
    }
}
```
**A:** 
```
Charlie: 200 actions
Alice: 150 actions
Diana: 120 actions
```

---

### 12. Password Policy — Validate Complexity
**Q: Check if passwords meet security requirements. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, String> userPasswords = new HashMap<>();
        userPasswords.put("alice", "Password123!");
        userPasswords.put("bob", "weak");
        userPasswords.put("charlie", "StrongPass456@");
        
        Set<String> weakPasswords = userPasswords.entrySet().stream()
            .filter(entry -> {
                String password = entry.getValue();
                return password.length() < 8 || 
                       !password.matches(".*[A-Z].*") ||
                       !password.matches(".*[a-z].*") ||
                       !password.matches(".*\\d.*") ||
                       !password.matches(".*[!@#$%^&*].*");
            })
            .map(Map.Entry::getKey)
            .collect(Collectors.toSet());
        
        System.out.println("Users with weak passwords: " + weakPasswords);
    }
}
```
**A:** `Users with weak passwords: [bob]`. Only Bob's password fails security checks.

---

### 13. User Groups — Find Common Members
**Q: Find users who belong to multiple groups. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Set<String> adminGroup = Set.of("alice", "bob", "charlie");
        Set<String> developerGroup = Set.of("bob", "charlie", "diana");
        Set<String> testerGroup = Set.of("charlie", "diana", "eve");
        
        Set<String> multiRoleUsers = Arrays.asList(adminGroup, developerGroup, testerGroup)
            .stream()
            .flatMap(Set::stream)
            .collect(Collectors.groupingBy(user -> user, Collectors.counting()))
            .entrySet().stream()
            .filter(entry -> entry.getValue() > 1)
            .map(Map.Entry::getKey)
            .collect(Collectors.toSet());
        
        System.out.println("Users in multiple groups: " + multiRoleUsers);
    }
}
```
**A:** `Users in multiple groups: [bob, charlie, diana]`. These users belong to 2+ groups.

---

### 14. Login Attempts — Detect Suspicious Activity
**Q: Detect users with unusual login patterns. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    public static void main(String[] args) {
        Map<String, List<LocalDateTime>> loginAttempts = new HashMap<>();
        loginAttempts.put("alice", Arrays.asList(
            LocalDateTime.now().minusHours(2),
            LocalDateTime.now().minusHours(1)
        ));
        loginAttempts.put("bob", Arrays.asList(
            LocalDateTime.now().minusMinutes(5),
            LocalDateTime.now().minusMinutes(3),
            LocalDateTime.now().minusMinutes(1),
            LocalDateTime.now()
        ));
        
        Set<String> suspiciousUsers = loginAttempts.entrySet().stream()
            .filter(entry -> {
                List<LocalDateTime> attempts = entry.getValue();
                if (attempts.size() < 3) return false;
                
                LocalDateTime first = attempts.get(0);
                LocalDateTime last = attempts.get(attempts.size() - 1);
                return Duration.between(first, last).toMinutes() < 10;
            })
            .map(Map.Entry::getKey)
            .collect(Collectors.toSet());
        
        System.out.println("Suspicious login patterns: " + suspiciousUsers);
    }
}
```
**A:** `Suspicious login patterns: [bob]`. Bob had 4 logins within 5 minutes.

---

### 15. User Preferences — Merge Settings
**Q: Merge user preferences with default settings. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, String> defaultSettings = Map.of(
            "theme", "light",
            "language", "en",
            "notifications", "enabled",
            "fontSize", "medium"
        );
        
        Map<String, String> userSettings = Map.of(
            "theme", "dark",
            "fontSize", "large"
        );
        
        Map<String, String> mergedSettings = Stream.of(defaultSettings, userSettings)
            .flatMap(map -> map.entrySet().stream())
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue,
                (userValue, defaultValue) -> userValue // User settings override defaults
            ));
        
        mergedSettings.forEach((key, value) -> 
            System.out.println(key + ": " + value));
    }
}
```
**A:** 
```
theme: dark
language: en
notifications: enabled
fontSize: large
```

---

### 16. User Analytics — Calculate Engagement Metrics
**Q: Calculate user engagement scores based on activity. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Map<String, Integer>> userMetrics = new HashMap<>();
        userMetrics.put("alice", Map.of("logins", 10, "posts", 5, "comments", 20));
        userMetrics.put("bob", Map.of("logins", 5, "posts", 2, "comments", 8));
        userMetrics.put("charlie", Map.of("logins", 15, "posts", 12, "comments", 35));
        
        Map<String, Double> engagementScores = userMetrics.entrySet().stream()
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                entry -> {
                    Map<String, Integer> metrics = entry.getValue();
                    return metrics.getOrDefault("logins", 0) * 1.0 +
                           metrics.getOrDefault("posts", 0) * 5.0 +
                           metrics.getOrDefault("comments", 0) * 2.0;
                }
            ))
            .entrySet().stream()
            .sorted(Map.Entry.<String, Double>comparingByValue().reversed())
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                Map.Entry::getValue,
                (e1, e2) -> e1,
                LinkedHashMap::new
            ));
        
        engagementScores.forEach((user, score) -> 
            System.out.printf("%s: %.1f points%n", user, score));
    }
}
```
**A:** 
```
charlie: 115.0 points
alice: 70.0 points
bob: 31.0 points
```

---

## Section 3: Inventory & Product Management (Q17–Q24)

### 17. Product Catalog — Category Organization
**Q: Organize products by category and subcategory. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Product(String name, String category, String subcategory) {}
    
    public static void main(String[] args) {
        List<Product> products = Arrays.asList(
            new Product("Laptop", "Electronics", "Computers"),
            new Product("Mouse", "Electronics", "Accessories"),
            new Product("Keyboard", "Electronics", "Accessories"),
            new Product("Novel", "Books", "Fiction"),
            new Product("Textbook", "Books", "Education")
        );
        
        Map<String, Map<String, List<String>>> catalog = products.stream()
            .collect(Collectors.groupingBy(
                Product::category,
                Collectors.groupingBy(
                    Product::subcategory,
                    Collectors.mapping(Product::name, Collectors.toList())
                )
            ));
        
        catalog.forEach((category, subcategories) -> {
            System.out.println(category + ":");
            subcategories.forEach((subcat, items) -> 
                System.out.println("  " + subcat + ": " + items));
        });
    }
}
```
**A:** 
```
Electronics:
  Accessories: [Mouse, Keyboard]
  Computers: [Laptop]
Books:
  Education: [Textbook]
  Fiction: [Novel]
```

---

### 18. Stock Management — FIFO Inventory Tracking
**Q: Track inventory using FIFO (First In, First Out) principle. What is the output?**
```java
import java.util.*;
public class Main {
    public static void main(String[] args) {
        Queue<Map.Entry<String, LocalDate>> inventory = new LinkedList<>();
        inventory.add(Map.entry("Batch1", LocalDate.of(2024, 1, 1)));
        inventory.add(Map.entry("Batch2", LocalDate.of(2024, 1, 15)));
        inventory.add(Map.entry("Batch3", LocalDate.of(2024, 2, 1)));
        
        System.out.println("Initial inventory: " + 
            inventory.stream().map(Map.Entry::getKey).toList());
        
        // Sell items (FIFO)
        Map.Entry<String, LocalDate> sold = inventory.poll();
        System.out.println("Sold: " + sold.getKey() + " (from " + sold.getValue() + ")");
        
        System.out.println("Remaining inventory: " + 
            inventory.stream().map(Map.Entry::getKey).toList());
    }
}
```
**A:** 
```
Initial inventory: [Batch1, Batch2, Batch3]
Sold: Batch1 (from 2024-01-01)
Remaining inventory: [Batch2, Batch3]
```

---

### 19. Price History — Track Price Changes
**Q: Track product price changes over time. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    public static void main(String[] args) {
        Map<String, NavigableMap<LocalDate, Double>> priceHistory = new HashMap<>();
        
        NavigableMap<LocalDate, Double> laptopPrices = new TreeMap<>();
        laptopPrices.put(LocalDate.of(2024, 1, 1), 999.99);
        laptopPrices.put(LocalDate.of(2024, 2, 1), 949.99);
        laptopPrices.put(LocalDate.of(2024, 3, 1), 899.99);
        priceHistory.put("Laptop", laptopPrices);
        
        LocalDate queryDate = LocalDate.of(2024, 2, 15);
        Map.Entry<LocalDate, Double> priceAtQuery = priceHistory.get("Laptop")
            .floorEntry(queryDate);
        
        System.out.println("Price on " + queryDate + ": $" + priceAtQuery.getValue());
        
        // Calculate price drop
        Double firstPrice = laptopPrices.firstEntry().getValue();
        Double currentPrice = laptopPrices.lastEntry().getValue();
        double drop = firstPrice - currentPrice;
        
        System.out.printf("Total price drop: $%.2f (%.1f%%)%n", 
            drop, (drop / firstPrice) * 100);
    }
}
```
**A:** 
```
Price on 2024-02-15: $949.99
Total price drop: $100.00 (10.0%)
```

---

### 20. Supplier Management — Find Best Suppliers
**Q: Evaluate suppliers based on multiple criteria. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Supplier(String name, double price, int reliability, int deliveryDays) {}
    
    public static void main(String[] args) {
        List<Supplier> suppliers = Arrays.asList(
            new Supplier("SupplierA", 95.0, 8, 3),
            new Supplier("SupplierB", 90.0, 6, 5),
            new Supplier("SupplierC", 98.0, 9, 2),
            new Supplier("SupplierD", 92.0, 7, 4)
        );
        
        List<Supplier> rankedSuppliers = suppliers.stream()
            .sorted(Comparator
                .comparingDouble((Supplier s) -> s.price) // Lower price is better
                .thenComparing(Comparator.comparingInt((Supplier s) -> s.reliability).reversed()) // Higher reliability
                .thenComparingInt(Supplier::getDeliveryDays)) // Lower delivery time
            )
            .toList();
        
        rankedSuppliers.forEach(s -> System.out.printf(
            "%s: $%.1f, Reliability: %d/10, Delivery: %d days%n",
            s.name, s.price, s.reliability, s.deliveryDays));
    }
}
```
**A:** 
```
SupplierB: $90.0, Reliability: 6/10, Delivery: 5 days
SupplierD: $92.0, Reliability: 7/10, Delivery: 4 days
SupplierA: $95.0, Reliability: 8/10, Delivery: 3 days
SupplierC: $98.0, Reliability: 9/10, Delivery: 2 days
```

---

### 21. Product Search — Autocomplete Suggestions
**Q: Generate autocomplete suggestions for product search. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        List<String> productNames = Arrays.asList(
            "Laptop", "Laptop Stand", "Laptop Bag", "Mouse", "Mouse Pad", 
            "Keyboard", "Monitor", "Monitor Stand", "Webcam"
        );
        
        String query = "Lap";
        List<String> suggestions = productNames.stream()
            .filter(name -> name.toLowerCase().startsWith(query.toLowerCase()))
            .sorted(Comparator.comparingInt(String::length)) // Shorter matches first
            .limit(5)
            .toList();
        
        System.out.println("Suggestions for '" + query + "': " + suggestions);
    }
}
```
**A:** `Suggestions for 'Lap': [Laptop, Laptop Bag, Laptop Stand]`

---

### 22. Bundle Pricing — Calculate Package Deals
**Q: Calculate optimal bundle pricing for product packages. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Double> individualPrices = Map.of(
            "Laptop", 999.99,
            "Mouse", 29.99,
            "Keyboard", 79.99,
            "Monitor", 299.99
        );
        
        List<List<String>> bundles = Arrays.asList(
            Arrays.asList("Laptop", "Mouse", "Keyboard"),
            Arrays.asList("Laptop", "Monitor"),
            Arrays.asList("Mouse", "Keyboard"),
            Arrays.asList("Keyboard", "Monitor")
        );
        
        Map<List<String>, Double> bundlePrices = bundles.stream()
            .collect(Collectors.toMap(
                bundle -> bundle,
                bundle -> {
                    double total = bundle.stream()
                        .mapToDouble(product -> individualPrices.getOrDefault(product, 0.0))
                        .sum();
                    return total * 0.9; // 10% bundle discount
                }
            ));
        
        bundlePrices.forEach((bundle, price) -> {
            double individualTotal = bundle.stream()
                .mapToDouble(product -> individualPrices.getOrDefault(product, 0.0))
                .sum();
            double savings = individualTotal - price;
            System.out.printf("%s: $%.2f (save $%.2f)%n", 
                bundle, price, savings);
        });
    }
}
```
**A:** 
```
[Laptop, Mouse, Keyboard]: $999.97 (save $111.11)
[Laptop, Monitor]: $1179.96 (save $129.99)
[Mouse, Keyboard]: $98.97 (save $10.92)
[Keyboard, Monitor]: $341.96 (save $37.91)
```

---

### 23. Warehouse Locations — Nearest Warehouse Finder
**Q: Find nearest warehouse for customer orders. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Warehouse(String name, double x, double y) {
        double distanceTo(double customerX, double customerY) {
            return Math.sqrt(Math.pow(x - customerX, 2) + Math.pow(y - customerY, 2));
        }
    }
    
    public static void main(String[] args) {
        List<Warehouse> warehouses = Arrays.asList(
            new Warehouse("North", 10, 20),
            new Warehouse("South", 30, 5),
            new Warehouse("East", 50, 15),
            new Warehouse("West", 5, 10)
        );
        
        double customerX = 25, customerY = 12;
        
        Warehouse nearest = warehouses.stream()
            .min(Comparator.comparingDouble(w -> w.distanceTo(customerX, customerY)))
            .orElse(null);
        
        System.out.printf("Nearest warehouse: %s (%.2f units away)%n", 
            nearest.name, nearest.distanceTo(customerX, customerY));
    }
}
```
**A:** `Nearest warehouse: South (6.40 units away)`

---

### 24. Product Recommendations — Collaborative Filtering
**Q: Recommend products based on similar users' purchases. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Set<String>> userPurchases = Map.of(
            "User1", Set.of("Laptop", "Mouse", "Keyboard"),
            "User2", Set.of("Laptop", "Monitor", "Webcam"),
            "User3", Set.of("Mouse", "Keyboard", "Monitor"),
            "User4", Set.of("Laptop", "Mouse", "Webcam")
        );
        
        Set<String> targetUser = Set.of("Laptop", "Mouse");
        
        Map<String, Double> similarityScores = userPurchases.entrySet().stream()
            .filter(entry -> !entry.getKey().equals("TargetUser"))
            .collect(Collectors.toMap(
                Map.Entry::getKey,
                entry -> {
                    Set<String> purchases = entry.getValue();
                    long intersection = purchases.stream()
                        .filter(targetUser::contains)
                        .count();
                    long union = purchases.size() + targetUser.size() - intersection;
                    return union == 0 ? 0.0 : (double) intersection / union;
                }
            ));
        
        Set<String> recommendations = userPurchases.entrySet().stream()
            .filter(entry -> similarityScores.getOrDefault(entry.getKey(), 0.0) > 0.3)
            .flatMap(entry -> entry.getValue().stream())
            .filter(product -> !targetUser.contains(product))
            .collect(Collectors.toSet());
        
        System.out.println("Recommended products: " + recommendations);
    }
}
```
**A:** `Recommended products: [Keyboard, Monitor, Webcam]`

---

## Section 4: Order Processing & Queuing (Q25–Q32)

### 25. Order Queue — Process Orders by Priority
**Q: Process orders based on priority levels. What is the output?**
```java
import java.util.*;
public class Main {
    record Order(String id, String priority, double amount) {}
    
    public static void main(String[] args) {
        Queue<Order> orderQueue = new PriorityQueue<>(Comparator
            .comparing((Order o) -> o.priority)
            .thenComparing(Order::amount));
        
        orderQueue.add(new Order("ORD001", "HIGH", 500.0));
        orderQueue.add(new Order("ORD002", "LOW", 100.0));
        orderQueue.add(new Order("ORD003", "HIGH", 300.0));
        orderQueue.add(new Order("ORD004", "MEDIUM", 200.0));
        
        System.out.println("Processing orders by priority:");
        while (!orderQueue.isEmpty()) {
            Order order = orderQueue.poll();
            System.out.printf("%s: %s - $%.2f%n", 
                order.id, order.priority, order.amount);
        }
    }
}
```
**A:** 
```
Processing orders by priority:
ORD003: HIGH - $300.00
ORD001: HIGH - $500.00
ORD004: MEDIUM - $200.00
ORD002: LOW - $100.00
```

---

### 26. Order Fulfillment — Batch Processing
**Q: Group orders for efficient batch processing. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Order(String id, String destination, double weight) {}
    
    public static void main(String[] args) {
        List<Order> orders = Arrays.asList(
            new Order("ORD001", "NY", 2.5),
            new Order("ORD002", "NY", 1.8),
            new Order("ORD003", "CA", 3.2),
            new Order("ORD004", "NY", 0.9),
            new Order("ORD005", "CA", 2.1)
        );
        
        Map<String, List<Order>> batches = orders.stream()
            .collect(Collectors.groupingBy(Order::destination));
        
        batches.forEach((destination, batchOrders) -> {
            double totalWeight = batchOrders.stream()
                .mapToDouble(Order::weight)
                .sum();
            System.out.printf("Batch to %s: %d orders, %.1f kg total%n",
                destination, batchOrders.size(), totalWeight);
        });
    }
}
```
**A:** 
```
Batch to NY: 3 orders, 5.2 kg total
Batch to CA: 2 orders, 5.3 kg total
```

---

### 27. Order Tracking — Status Updates
**Q: Track order status through different stages. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    public static void main(String[] args) {
        Map<String, Map<String, LocalDateTime>> orderStatus = new LinkedHashMap<>();
        
        Map<String, LocalDateTime> order1Status = new LinkedHashMap<>();
        order1Status.put("Placed", LocalDateTime.now().minusDays(2));
        order1Status.put("Processing", LocalDateTime.now().minusDays(1));
        order1Status.put("Shipped", LocalDateTime.now().minusHours(6));
        orderStatus.put("ORD001", order1Status);
        
        Map<String, LocalDateTime> order2Status = new LinkedHashMap<>();
        order2Status.put("Placed", LocalDateTime.now().minusHours(12));
        order2Status.put("Processing", LocalDateTime.now().minusHours(6));
        orderStatus.put("ORD002", order2Status);
        
        orderStatus.forEach((orderId, statusMap) -> {
            String currentStatus = statusMap.keySet().stream()
                .reduce((first, second) -> second)
                .orElse("Unknown");
            LocalDateTime lastUpdate = statusMap.get(currentStatus);
            
            System.out.printf("%s: %s (updated %s ago)%n",
                orderId, currentStatus, 
                Duration.between(lastUpdate, LocalDateTime.now()).toHours() + " hours");
        });
    }
}
```
**A:** 
```
ORD001: Shipped (updated 6 hours ago)
ORD002: Processing (updated 6 hours ago)
```

---

### 28. Order Validation — Check Order Completeness
**Q: Validate orders have all required information. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Order(String id, String customer, String address, List<String> items) {}
    
    public static void main(String[] args) {
        List<Order> orders = Arrays.asList(
            new Order("ORD001", "Alice", "123 Main St", Arrays.asList("Laptop", "Mouse")),
            new Order("ORD002", "Bob", null, Arrays.asList("Keyboard")),
            new Order("ORD003", "Charlie", "456 Oak Ave", Arrays.asList()),
            new Order("ORD004", "Diana", "789 Pine Rd", Arrays.asList("Monitor"))
        );
        
        List<String> invalidOrders = orders.stream()
            .filter(order -> order.customer == null || order.customer.trim().isEmpty())
            .filter(order -> order.address == null || order.address.trim().isEmpty())
            .filter(order -> order.items == null || order.items.isEmpty())
            .map(Order::id)
            .toList();
        
        List<String> incompleteOrders = orders.stream()
            .filter(order -> order.items == null || order.items.isEmpty())
            .map(Order::id)
            .toList();
        
        System.out.println("Invalid orders: " + invalidOrders);
        System.out.println("Orders without items: " + incompleteOrders);
    }
}
```
**A:** 
```
Invalid orders: []
Orders without items: [ORD003]
```

---

### 29. Order Analytics — Calculate Metrics
**Q: Calculate order processing metrics. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Order(String id, double amount, String category, LocalDateTime date) {}
    
    public static void main(String[] args) {
        List<Order> orders = Arrays.asList(
            new Order("ORD001", 500.0, "Electronics", LocalDateTime.now().minusDays(5)),
            new Order("ORD002", 200.0, "Books", LocalDateTime.now().minusDays(3)),
            new Order("ORD003", 750.0, "Electronics", LocalDateTime.now().minusDays(2)),
            new Order("ORD004", 150.0, "Books", LocalDateTime.now().minusDays(1))
        );
        
        double totalRevenue = orders.stream()
            .mapToDouble(Order::amount)
            .sum();
        
        Map<String, Double> revenueByCategory = orders.stream()
            .collect(Collectors.groupingBy(
                Order::category,
                Collectors.summingDouble(Order::amount)
            ));
        
        double averageOrderValue = orders.stream()
            .mapToDouble(Order::amount)
            .average()
            .orElse(0.0);
        
        System.out.printf("Total Revenue: $%.2f%n", totalRevenue);
        System.out.println("Revenue by Category:");
        revenueByCategory.forEach((cat, rev) -> 
            System.out.printf("  %s: $%.2f%n", cat, rev));
        System.out.printf("Average Order Value: $%.2f%n", averageOrderValue);
    }
}
```
**A:** 
```
Total Revenue: $1600.00
Revenue by Category:
  Electronics: $1250.00
  Books: $350.00
Average Order Value: $400.00
```

---

### 30. Order Routing — Find Optimal Route
**Q: Find optimal route for order deliveries. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record Delivery(String orderId, double distance, int priority) {}
    
    public static void main(String[] args) {
        List<Delivery> deliveries = Arrays.asList(
            new Delivery("ORD001", 5.2, 1),
            new Delivery("ORD002", 2.8, 3),
            new Delivery("ORD003", 8.1, 2),
            new Delivery("ORD004", 3.5, 1)
        );
        
        List<Delivery> optimizedRoute = deliveries.stream()
            .sorted(Comparator
                .comparingInt((Delivery d) -> d.priority) // Higher priority first
                .reversed()
                .thenComparingDouble(Delivery::distance)) // Then shorter distance
            .toList();
        
        System.out.println("Optimized delivery route:");
        optimizedRoute.forEach(d -> 
            System.out.printf("%s: Priority %d, %.1f km%n", 
                d.orderId, d.priority, d.distance));
    }
}
```
**A:** 
```
Optimized delivery route:
ORD002: Priority 3, 2.8 km
ORD003: Priority 2, 8.1 km
ORD001: Priority 1, 5.2 km
ORD004: Priority 1, 3.5 km
```

---

### 31. Order Cancellation — Handle Refunds
**Q: Process order cancellations and calculate refunds. What is the output?**
```java
import java.util.*;
import java.time.*;
public class Main {
    record Order(String id, double amount, LocalDateTime orderDate, boolean cancelled) {}
    
    public static void main(String[] args) {
        List<Order> orders = Arrays.asList(
            new Order("ORD001", 500.0, LocalDateTime.now().minusDays(10), true),
            new Order("ORD002", 200.0, LocalDateTime.now().minusDays(5), false),
            new Order("ORD003", 750.0, LocalDateTime.now().minusDays(2), true),
            new Order("ORD004", 150.0, LocalDateTime.now().minusDays(1), false)
        );
        
        Map<String, Double> refunds = orders.stream()
            .filter(Order::cancelled)
            .collect(Collectors.toMap(
                Order::id,
                order -> {
                    long daysSinceOrder = Duration.between(order.orderDate, LocalDateTime.now()).toDays();
                    double refundAmount = order.amount;
                    if (daysSinceOrder > 7) {
                        refundAmount *= 0.8; // 20% penalty after 7 days
                    }
                    return refundAmount;
                }
            ));
        
        double totalRefunds = refunds.values().stream()
            .mapToDouble(Double::doubleValue)
            .sum();
        
        System.out.println("Refunds to process:");
        refunds.forEach((orderId, amount) -> 
            System.out.printf("%s: $%.2f%n", orderId, amount));
        System.out.printf("Total refund amount: $%.2f%n", totalRefunds);
    }
}
```
**A:** 
```
Refunds to process:
ORD001: $400.00
ORD003: $750.00
Total refund amount: $1150.00
```

---

### 32. Order Forecasting — Predict Future Orders
**Q: Predict order volume based on historical data. What is the output?**
```java
import java.util.*;
import java.util.stream.*;
public class Main {
    record OrderData(String month, int orderCount) {}
    
    public static void main(String[] args) {
        List<OrderData> historicalData = Arrays.asList(
            new OrderData("Jan", 120),
            new OrderData("Feb", 135),
            new OrderData("Mar", 142),
            new OrderData("Apr", 158),
            new OrderData("May", 165),
            new OrderData("Jun", 172)
        );
        
        double averageGrowthRate = historicalData.stream()
            .skip(1) // Skip first month
            .mapToDouble(data -> {
                int previousIndex = historicalData.indexOf(data) - 1;
                int previousCount = historicalData.get(previousIndex).orderCount;
                return ((double) (data.orderCount - previousCount) / previousCount) * 100;
            })
            .average()
            .orElse(0.0);
        
        int lastMonthCount = historicalData.get(historicalData.size() - 1).orderCount;
        double predictedNextMonth = lastMonthCount * (1 + averageGrowthRate / 100);
        
        System.out.printf("Average monthly growth: %.1f%%%n", averageGrowthRate);
        System.out.printf("Predicted orders for next month: %.0f%n", predictedNextMonth);
    }
}
```
**A:** 
```
Average monthly growth: 7.4%
Predicted orders for next month: 185
```
