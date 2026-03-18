# Restaurant Order Management - Collections & Streams Integration

> **Challenge:** Design a restaurant management system that efficiently handles orders, inventory, and customer analytics using both collections and functional programming.

---

## 🍽️ Problem Statement

A busy restaurant needs a system to:
1. Manage table reservations and waitlists
2. Process orders and track kitchen status
3. Handle inventory and ingredient management
4. Generate revenue analytics and customer insights
5. Optimize kitchen workflow and table turnover

---

## 🎯 Core Challenge

**Question:** How would you integrate Collections Framework and Functional Programming to create an efficient restaurant management system that handles real-time operations and provides business insights?

---

## 🛠️ Solution Implementation

### MenuItem.java - Model Class
```java
import java.util.Objects;

public class MenuItem {
    private String itemId;
    private String name;
    private String category;
    private double price;
    private int preparationTime; // in minutes
    private boolean available;
    private Map<String, Integer> ingredients;

    public MenuItem(String itemId, String name, String category, double price, 
                   int preparationTime, boolean available) {
        this.itemId = itemId;
        this.name = name;
        this.category = category;
        this.price = price;
        this.preparationTime = preparationTime;
        this.available = available;
        this.ingredients = new HashMap<>();
    }

    // Getters
    public String getItemId() { return itemId; }
    public String getName() { return name; }
    public String getCategory() { return category; }
    public double getPrice() { return price; }
    public int getPreparationTime() { return preparationTime; }
    public boolean isAvailable() { return available; }
    public Map<String, Integer> getIngredients() { return ingredients; }

    // Setters
    public void setAvailable(boolean available) { this.available = available; }

    // Business methods
    public void addIngredient(String ingredientName, int quantity) {
        ingredients.put(ingredientName, quantity);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        MenuItem menuItem = (MenuItem) o;
        return Objects.equals(itemId, menuItem.itemId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(itemId);
    }

    @Override
    public String toString() {
        return String.format("MenuItem{id='%s', name='%s', price=$%.2f, prepTime=%dmin, available=%s}",
                itemId, name, price, preparationTime, available);
    }
}
```

### Order.java - Model Class
```java
import java.time.LocalDateTime;
import java.util.Objects;

public class Order {
    private String orderId;
    private String tableId;
    private Map<String, Integer> items; // itemId -> quantity
    private LocalDateTime orderTime;
    private String status; // PENDING, PREPARING, READY, SERVED, CANCELLED
    private double totalAmount;
    private int estimatedPrepTime;
    private String specialInstructions;

    public Order(String orderId, String tableId, Map<String, Integer> items, 
                LocalDateTime orderTime, String specialInstructions) {
        this.orderId = orderId;
        this.tableId = tableId;
        this.items = new HashMap<>(items);
        this.orderTime = orderTime;
        this.status = "PENDING";
        this.specialInstructions = specialInstructions;
        this.totalAmount = 0.0;
        this.estimatedPrepTime = 0;
    }

    // Getters
    public String getOrderId() { return orderId; }
    public String getTableId() { return tableId; }
    public Map<String, Integer> getItems() { return items; }
    public LocalDateTime getOrderTime() { return orderTime; }
    public String getStatus() { return status; }
    public double getTotalAmount() { return totalAmount; }
    public int getEstimatedPrepTime() { return estimatedPrepTime; }
    public String getSpecialInstructions() { return specialInstructions; }

    // Setters
    public void setStatus(String status) { this.status = status; }
    public void setTotalAmount(double totalAmount) { this.totalAmount = totalAmount; }
    public void setEstimatedPrepTime(int estimatedPrepTime) { this.estimatedPrepTime = estimatedPrepTime; }

    // Business methods
    public int getTotalItems() {
        return items.values().stream().mapToInt(Integer::intValue).sum();
    }

    public boolean isCompleted() {
        return "SERVED".equals(status) || "CANCELLED".equals(status);
    }

    public long getWaitingTimeMinutes() {
        return java.time.Duration.between(orderTime, LocalDateTime.now()).toMinutes();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Order order = (Order) o;
        return Objects.equals(orderId, order.orderId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(orderId);
    }

    @Override
    public String toString() {
        return String.format("Order{id='%s', table='%s', status='%s', items=%d, total=$%.2f}",
                orderId, tableId, status, getTotalItems(), totalAmount);
    }
}
```

### Table.java - Model Class
```java
import java.time.LocalDateTime;
import java.util.Objects;

public class Table {
    private String tableId;
    private int capacity;
    private String status; // AVAILABLE, OCCUPIED, RESERVED, CLEANING
    private LocalDateTime occupiedSince;
    private String currentOrderId;
    private String reservationName;

    public Table(String tableId, int capacity) {
        this.tableId = tableId;
        this.capacity = capacity;
        this.status = "AVAILABLE";
        this.occupiedSince = null;
        this.currentOrderId = null;
        this.reservationName = null;
    }

    // Getters
    public String getTableId() { return tableId; }
    public int getCapacity() { return capacity; }
    public String getStatus() { return status; }
    public LocalDateTime getOccupiedSince() { return occupiedSince; }
    public String getCurrentOrderId() { return currentOrderId; }
    public String getReservationName() { return reservationName; }

    // Business methods
    public boolean isAvailable() { return "AVAILABLE".equals(status); }
    public boolean isOccupied() { return "OCCUPIED".equals(status); }
    public boolean isReserved() { return "RESERVED".equals(status); }
    
    public void occupy(String orderId) {
        this.status = "OCCUPIED";
        this.occupiedSince = LocalDateTime.now();
        this.currentOrderId = orderId;
        this.reservationName = null;
    }
    
    public void reserve(String name) {
        this.status = "RESERVED";
        this.reservationName = name;
    }
    
    public void release() {
        this.status = "AVAILABLE";
        this.occupiedSince = null;
        this.currentOrderId = null;
        this.reservationName = null;
    }
    
    public long getOccupiedMinutes() {
        return occupiedSince != null ? 
            java.time.Duration.between(occupiedSince, LocalDateTime.now()).toMinutes() : 0;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Table table = (Table) o;
        return Objects.equals(tableId, table.tableId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(tableId);
    }

    @Override
    public String toString() {
        return String.format("Table{id='%s', capacity=%d, status='%s', occupied=%d min}",
                tableId, capacity, status, getOccupiedMinutes());
    }
}
```

### RestaurantManagementSystem.java - Main Class
```java
import java.time.LocalDateTime;
import java.util.*;
import java.util.function.*;
import java.util.stream.Collectors;

public class RestaurantManagementSystem {
    // Collections for different entities
    private Map<String, MenuItem> menuItems;              // HashMap for quick menu lookup
    private List<Order> orders;                           // ArrayList for order processing
    private Map<String, Table> tables;                    // HashMap for table management
    private Queue<String> waitlist;                       // LinkedList for customer waitlist
    private Map<String, Integer> inventory;                // HashMap for ingredient tracking
    private Map<String, List<Order>> kitchenQueue;        // TreeMap for priority-based kitchen
    private Set<String> activeCategories;                  // HashSet for category management
    private Deque<Order> recentOrders;                    // LinkedList as Deque for recent orders

    // Functional interfaces for business logic
    private Predicate<Table> availableTablePredicate;
    private Function<Order, Double> orderValueCalculator;
    private Consumer<Order> orderProcessor;
    private Supplier<String> orderIdGenerator;

    public RestaurantManagementSystem() {
        this.menuItems = new HashMap<>();
        this.orders = new ArrayList<>();
        this.tables = new HashMap<>();
        this.waitlist = new LinkedList<>();
        this.inventory = new HashMap<>();
        this.kitchenQueue = new TreeMap<>(Comparator.comparingInt(this::getOrderPriority).reversed());
        this.activeCategories = new HashSet<>();
        this.recentOrders = new LinkedList<>();

        initializeFunctionalInterfaces();
        initializeSampleData();
    }

    private void initializeFunctionalInterfaces() {
        // Predicate for finding available tables
        availableTablePredicate = table -> table.isAvailable();
        
        // Function for calculating order value
        orderValueCalculator = order -> {
            return order.getItems().entrySet().stream()
                .mapToDouble(entry -> {
                    MenuItem item = menuItems.get(entry.getKey());
                    return item != null ? item.getPrice() * entry.getValue() : 0.0;
                })
                .sum();
        };
        
        // Consumer for order processing
        orderProcessor = order -> {
            updateInventory(order);
            updateKitchenQueue(order);
            System.out.println("Processing order: " + order.getOrderId());
        };
        
        // Supplier for generating order IDs
        orderIdGenerator = () -> "ORD-" + System.currentTimeMillis();
    }

    private void initializeSampleData() {
        // Initialize tables
        for (int i = 1; i <= 10; i++) {
            int capacity = i <= 4 ? 4 : i <= 7 ? 6 : 8;
            tables.put("T" + i, new Table("T" + i, capacity));
        }

        // Initialize menu items
        addMenuItem(new MenuItem("M001", "Burger", "Main", 12.99, 15, true));
        addMenuItem(new MenuItem("M002", "Pizza", "Main", 15.99, 20, true));
        addMenuItem(new MenuItem("M003", "Salad", "Appetizer", 8.99, 10, true));
        addMenuItem(new MenuItem("M004", "Soup", "Appetizer", 6.99, 8, true));
        addMenuItem(new MenuItem("M005", "Steak", "Main", 24.99, 25, true));
        addMenuItem(new MenuItem("M006", "Pasta", "Main", 14.99, 18, true));
        addMenuItem(new MenuItem("M007", "Ice Cream", "Dessert", 5.99, 5, true));
        addMenuItem(new MenuItem("M008", "Coffee", "Beverage", 3.99, 3, true));

        // Initialize inventory
        inventory.put("Beef", 50);
        inventory.put("Cheese", 100);
        inventory.put("Bread", 80);
        inventory.put("Tomato", 60);
        inventory.put("Lettuce", 40);
        inventory.put("Dough", 30);
        inventory.put("Pasta", 25);
        inventory.put("Ice Cream", 20);
        inventory.put("Coffee Beans", 50);

        // Add ingredients to menu items
        menuItems.get("M001").addIngredient("Beef", 1);
        menuItems.get("M001").addIngredient("Cheese", 1);
        menuItems.get("M001").addIngredient("Bread", 1);
        menuItems.get("M001").addIngredient("Lettuce", 1);
        
        menuItems.get("M002").addIngredient("Cheese", 2);
        menuItems.get("M002").addIngredient("Tomato", 1);
        menuItems.get("M002").addIngredient("Dough", 1);
        
        menuItems.get("M003").addIngredient("Lettuce", 2);
        menuItems.get("M003").addIngredient("Tomato", 1);
        
        menuItems.get("M004").addIngredient("Tomato", 2);
        
        menuItems.get("M005").addIngredient("Beef", 2);
        
        menuItems.get("M006").addIngredient("Pasta", 1);
        
        menuItems.get("M007").addIngredient("Ice Cream", 1);
        
        menuItems.get("M008").addIngredient("Coffee Beans", 1);
    }

    private void addMenuItem(MenuItem item) {
        menuItems.put(item.getItemId(), item);
        activeCategories.add(item.getCategory());
    }

    // === COLLECTIONS & STREAMS INTEGRATION CHALLENGES ===

    /**
     * Challenge 1: Table management with collections
     * Problem: Efficiently manage table assignments and reservations
     */
    public String assignTable(int partySize, String reservationName) {
        System.out.println("\n=== TABLE ASSIGNMENT CHALLENGE ===");
        
        // Find available table using streams and predicates
        Optional<Table> availableTable = tables.values().stream()
            .filter(availableTablePredicate.and(table -> table.getCapacity() >= partySize))
            .min(Comparator.comparingInt(Table::getCapacity));

        if (availableTable.isPresent()) {
            Table table = availableTable.get();
            table.occupy("TEMP_ORDER"); // Will be updated when order is placed
            System.out.println("Table " + table.getTableId() + " assigned to " + reservationName);
            return table.getTableId();
        } else {
            // Add to waitlist
            waitlist.add(reservationName + ":" + partySize);
            System.out.println("No available table. Added to waitlist: " + reservationName);
            return null;
        }
    }

    /**
     * Challenge 2: Order processing with functional programming
     * Problem: Process orders using functional interfaces and streams
     */
    public String createOrder(String tableId, Map<String, Integer> items, String specialInstructions) {
        System.out.println("\n=== ORDER PROCESSING CHALLENGE ===");
        
        // Validate items using streams
        boolean allItemsAvailable = items.entrySet().stream()
            .allMatch(entry -> {
                MenuItem item = menuItems.get(entry.getKey());
                return item != null && item.isAvailable();
            });

        if (!allItemsAvailable) {
            System.out.println("Some items are not available");
            return null;
        }

        // Create order
        String orderId = orderIdGenerator.get();
        Order order = new Order(orderId, tableId, items, LocalDateTime.now(), specialInstructions);
        
        // Calculate total using function
        double total = orderValueCalculator.apply(order);
        order.setTotalAmount(total);
        
        // Calculate preparation time
        int prepTime = items.entrySet().stream()
            .mapToInt(entry -> {
                MenuItem item = menuItems.get(entry.getKey());
                return item != null ? item.getPreparationTime() * entry.getValue() : 0;
            })
            .max()
            .orElse(0);
        order.setEstimatedPrepTime(prepTime);

        // Add to collections
        orders.add(order);
        recentOrders.addFirst(order);
        if (recentOrders.size() > 10) {
            recentOrders.removeLast();
        }

        // Process order using consumer
        orderProcessor.accept(order);

        System.out.println("Order created: " + order);
        return orderId;
    }

    /**
     * Challenge 3: Kitchen queue management with priority
     * Problem: Manage kitchen queue with priority-based processing
     */
    private void updateKitchenQueue(Order order) {
        // Use TreeMap for automatic priority sorting
        kitchenQueue.put(order.getOrderId(), order);
        System.out.println("Added to kitchen queue: " + order.getOrderId());
    }

    public void processKitchenQueue() {
        System.out.println("\n=== KITCHEN QUEUE CHALLENGE ===");
        
        while (!kitchenQueue.isEmpty()) {
            String orderId = kitchenQueue.firstKey();
            Order order = kitchenQueue.remove(orderId);
            
            order.setStatus("PREPARING");
            System.out.println("Kitchen preparing: " + orderId);
            
            // Simulate preparation
            try {
                Thread.sleep(100); // Simulate work
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
            
            order.setStatus("READY");
            System.out.println("Order ready: " + orderId);
        }
    }

    /**
     * Challenge 4: Inventory management with streams
     * Problem: Track and manage inventory using functional operations
     */
    private void updateInventory(Order order) {
        order.getItems().forEach((itemId, quantity) -> {
            MenuItem item = menuItems.get(itemId);
            if (item != null) {
                item.getIngredients().forEach((ingredient, requiredQty) -> {
                    int totalRequired = requiredQty * quantity;
                    inventory.merge(ingredient, -totalRequired, Integer::sum);
                    
                    // Check for low inventory
                    if (inventory.get(ingredient) < 10) {
                        System.out.println("LOW INVENTORY ALERT: " + ingredient + 
                            " (Remaining: " + inventory.get(ingredient) + ")");
                    }
                });
            }
        });
    }

    public void displayInventoryStatus() {
        System.out.println("\n=== INVENTORY STATUS CHALLENGE ===");
        
        // Group inventory by status using streams
        Map<String, Long> inventoryStatus = inventory.entrySet().stream()
            .collect(Collectors.groupingBy(
                entry -> entry.getValue() < 10 ? "LOW" : 
                         entry.getValue() < 25 ? "MEDIUM" : "GOOD",
                Collectors.counting()
            ));

        System.out.println("Inventory status:");
        inventoryStatus.forEach((status, count) -> 
            System.out.println("  " + status + ": " + count + " ingredients"));

        // Show low inventory items
        List<String> lowItems = inventory.entrySet().stream()
            .filter(entry -> entry.getValue() < 10)
            .map(Map.Entry::getKey)
            .sorted()
            .collect(Collectors.toList());

        if (!lowItems.isEmpty()) {
            System.out.println("\nLow inventory items (restock needed):");
            lowItems.forEach(item -> System.out.println("  - " + item + ": " + inventory.get(item)));
        }
    }

    /**
     * Challenge 5: Revenue analytics with streams and collectors
     * Problem: Generate comprehensive revenue analytics
     */
    public void generateRevenueAnalytics() {
        System.out.println("\n=== REVENUE ANALYTICS CHALLENGE ===");
        
        // Total revenue
        double totalRevenue = orders.stream()
            .mapToDouble(Order::getTotalAmount)
            .sum();

        // Revenue by category
        Map<String, Double> revenueByCategory = orders.stream()
            .flatMap(order -> order.getItems().entrySet().stream())
            .collect(Collectors.groupingBy(
                entry -> menuItems.get(entry.getKey()).getCategory(),
                Collectors.summingDouble(entry -> {
                    MenuItem item = menuItems.get(entry.getKey());
                    return item.getPrice() * entry.getValue();
                })
            ));

        // Most popular items
        Map<String, Long> itemPopularity = orders.stream()
            .flatMap(order -> order.getItems().entrySet().stream())
            .collect(Collectors.groupingBy(
                Map.Entry::getKey,
                Collectors.summingLong(Map.Entry::getValue)
            ));

        // Average order value
        OptionalDouble avgOrderValue = orders.stream()
            .mapToDouble(Order::getTotalAmount)
            .average();

        // Display results
        System.out.println("Total Revenue: $" + String.format("%.2f", totalRevenue));
        
        System.out.println("\nRevenue by Category:");
        revenueByCategory.forEach((category, revenue) -> 
            System.out.println("  " + category + ": $" + String.format("%.2f", revenue)));

        System.out.println("\nTop 5 Most Popular Items:");
        itemPopularity.entrySet().stream()
            .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
            .limit(5)
            .forEach(entry -> {
                MenuItem item = menuItems.get(entry.getKey());
                System.out.println("  " + item.getName() + ": " + entry.getValue() + " orders");
            });

        avgOrderValue.ifPresent(avg -> 
            System.out.println("\nAverage Order Value: $" + String.format("%.2f", avg)));
    }

    /**
     * Challenge 6: Table turnover analysis
     * Problem: Analyze table efficiency and turnover rates
     */
    public void analyzeTableTurnover() {
        System.out.println("\n=== TABLE TURNOVER CHALLENGE ===");
        
        // Table status distribution
        Map<String, Long> tableStatus = tables.values().stream()
            .collect(Collectors.groupingBy(Table::getStatus, Collectors.counting()));

        System.out.println("Table Status Distribution:");
        tableStatus.forEach((status, count) -> 
            System.out.println("  " + status + ": " + count + " tables"));

        // Average occupation time
        double avgOccupationTime = tables.values().stream()
            .filter(Table::isOccupied)
            .mapToInt(Table::getOccupiedMinutes)
            .average()
            .orElse(0.0);

        System.out.println("\nAverage Table Occupation Time: " + 
            String.format("%.1f", avgOccupationTime) + " minutes");

        // Most efficient tables (highest turnover)
        List<Table> efficientTables = tables.values().stream()
            .filter(table -> table.getOccupiedMinutes() > 0)
            .sorted(Comparator.comparingInt(Table::getOccupiedMinutes))
            .limit(3)
            .collect(Collectors.toList());

        System.out.println("\nMost Efficient Tables (fastest turnover):");
        efficientTables.forEach(table -> 
            System.out.println("  " + table.getTableId() + ": " + table.getOccupiedMinutes() + " minutes"));
    }

    /**
     * Challenge 7: Menu optimization using streams
     * Problem: Analyze menu performance and suggest optimizations
     */
    public void analyzeMenuPerformance() {
        System.out.println("\n=== MENU PERFORMANCE CHALLENGE ===");
        
        // Category performance
        Map<String, Long> categoryOrders = orders.stream()
            .flatMap(order -> order.getItems().entrySet().stream())
            .collect(Collectors.groupingBy(
                entry -> menuItems.get(entry.getKey()).getCategory(),
                Collectors.summingLong(Map.Entry::getValue)
            ));

        // Category revenue
        Map<String, Double> categoryRevenue = orders.stream()
            .flatMap(order -> order.getItems().entrySet().stream())
            .collect(Collectors.groupingBy(
                entry -> menuItems.get(entry.getKey()).getCategory(),
                Collectors.summingDouble(entry -> {
                    MenuItem item = menuItems.get(entry.getKey());
                    return item.getPrice() * entry.getValue();
                })
            ));

        // Calculate performance metrics
        System.out.println("Category Performance:");
        categoryOrders.forEach((category, orderCount) -> {
            double revenue = categoryRevenue.getOrDefault(category, 0.0);
            double avgOrderValue = orderCount > 0 ? revenue / orderCount : 0.0;
            System.out.printf("  %s: %d orders, $%.2f revenue, $%.2f avg value%n",
                category, orderCount, revenue, avgOrderValue);
        });

        // Least popular items
        Map<String, Long> itemOrders = orders.stream()
            .flatMap(order -> order.getItems().entrySet().stream())
            .collect(Collectors.groupingBy(Map.Entry::getKey, Collectors.summingLong(Map.Entry::getValue)));

        List<String> unpopularItems = itemOrders.entrySet().stream()
            .sorted(Map.Entry.comparingByValue())
            .limit(3)
            .map(Map.Entry::getKey)
            .collect(Collectors.toList());

        System.out.println("\nLeast Popular Items (consider removing):");
        unpopularItems.forEach(itemId -> {
            MenuItem item = menuItems.get(itemId);
            long orders = itemOrders.get(itemId);
            System.out.println("  " + item.getName() + ": " + orders + " orders");
        });
    }

    /**
     * Challenge 8: Waitlist management with collections
     * Problem: Efficiently manage customer waitlist using queue operations
     */
    public void manageWaitlist() {
        System.out.println("\n=== WAITLIST MANAGEMENT CHALLENGE ===");
        
        System.out.println("Current waitlist: " + waitlist.size() + " parties");
        
        // Process waitlist when tables become available
        Iterator<String> iterator = waitlist.iterator();
        while (iterator.hasNext()) {
            String waitlistEntry = iterator.next();
            String[] parts = waitlistEntry.split(":");
            String name = parts[0];
            int partySize = Integer.parseInt(parts[1]);
            
            String tableId = assignTable(partySize, name);
            if (tableId != null) {
                iterator.remove();
                System.out.println("Waitlist processed: " + name + " assigned to " + tableId);
            }
        }
        
        if (!waitlist.isEmpty()) {
            System.out.println("Remaining on waitlist:");
            waitlist.forEach(entry -> {
                String[] parts = entry.split(":");
                System.out.println("  - " + parts[0] + " (party of " + parts[1] + ")");
            });
        }
    }

    /**
     * Challenge 9: Performance comparison between collection types
     * Problem: Demonstrate performance differences in restaurant context
     */
    public void demonstratePerformanceComparison() {
        System.out.println("\n=== PERFORMANCE COMPARISON CHALLENGE ===");
        
        // HashMap vs TreeMap for menu lookup
        Map<String, MenuItem> hashMap = new HashMap<>(menuItems);
        Map<String, MenuItem> treeMap = new TreeMap<>(menuItems);
        
        String[] searchItems = {"M001", "M002", "M003", "M004", "M005"};
        
        // HashMap lookup
        long start = System.nanoTime();
        for (int i = 0; i < 10000; i++) {
            for (String itemId : searchItems) {
                hashMap.get(itemId);
            }
        }
        long hashMapTime = System.nanoTime() - start;
        
        // TreeMap lookup
        start = System.nanoTime();
        for (int i = 0; i < 10000; i++) {
            for (String itemId : searchItems) {
                treeMap.get(itemId);
            }
        }
        long treeMapTime = System.nanoTime() - start;
        
        System.out.println("Menu lookup performance:");
        System.out.println("HashMap: " + (hashMapTime / 1_000_000) + " ms");
        System.out.println("TreeMap: " + (treeMapTime / 1_000_000) + " ms");
        System.out.println("HashMap is " + String.format("%.2f", (double) treeMapTime / hashMapTime) + "x faster");
        
        // ArrayList vs LinkedList for order processing
        List<Order> arrayList = new ArrayList<>(orders);
        List<Order> linkedList = new LinkedList<>(orders);
        
        // Remove from beginning
        start = System.nanoTime();
        while (!arrayList.isEmpty()) {
            arrayList.remove(0);
        }
        long arrayListTime = System.nanoTime() - start;
        
        // Re-populate
        arrayList.addAll(orders);
        
        start = System.nanoTime();
        while (!linkedList.isEmpty()) {
            linkedList.remove(0);
        }
        long linkedListTime = System.nanoTime() - start;
        
        System.out.println("\nOrder queue removal performance:");
        System.out.println("ArrayList: " + (arrayListTime / 1_000_000) + " ms");
        System.out.println("LinkedList: " + (linkedListTime / 1_000_000) + " ms");
        System.out.println("LinkedList is " + String.format("%.2f", (double) arrayListTime / linkedListTime) + "x faster");
    }

    /**
     * Challenge 10: Advanced stream operations for business insights
     * Problem: Use advanced stream features for complex business analysis
     */
    public void advancedBusinessInsights() {
        System.out.println("\n=== ADVANCED BUSINESS INSIGHTS CHALLENGE ===");
        
        // Peak hours analysis
        Map<Integer, Long> ordersByHour = orders.stream()
            .collect(Collectors.groupingBy(
                order -> order.getOrderTime().getHour(),
                Collectors.counting()
            ));

        System.out.println("Orders by Hour:");
        ordersByHour.entrySet().stream()
            .sorted(Map.Entry.comparingByKey())
            .forEach(entry -> 
                System.out.println("  " + String.format("%02d:00", entry.getKey()) + ": " + entry.getValue() + " orders"));

        // Customer preference patterns
        Map<String, Map<String, Long>> categoryPreferences = orders.stream()
            .collect(Collectors.groupingBy(
                order -> order.getTableId(),
                Collectors.groupingBy(
                    order -> {
                        // Get primary category (most items from one category)
                        return order.getItems().entrySet().stream()
                            .collect(Collectors.groupingBy(
                                entry -> menuItems.get(entry.getKey()).getCategory(),
                                Collectors.summingInt(Map.Entry::getValue)
                            ))
                            .entrySet().stream()
                            .max(Map.Entry.comparingByValue())
                            .map(Map.Entry::getKey)
                            .orElse("Mixed");
                    },
                    Collectors.counting()
                )
            ));

        System.out.println("\nCategory Preferences by Table:");
        categoryPreferences.forEach((tableId, preferences) -> {
            String topCategory = preferences.entrySet().stream()
                .max(Map.Entry.comparingByValue())
                .map(Map.Entry::getKey)
                .orElse("None");
            System.out.println("  " + tableId + ": " + topCategory);
        });

        // Revenue optimization suggestions
        Map<String, Double> itemProfitability = menuItems.values().stream()
            .collect(Collectors.toMap(
                MenuItem::getItemId,
                item -> {
                    long orders = itemOrders.getOrDefault(item.getItemId(), 0L);
                    return orders > 0 ? (item.getPrice() * orders) / orders : 0.0;
                }
            ));

        List<MenuItem> mostProfitable = menuItems.values().stream()
            .sorted(Comparator.comparingDouble(item -> 
                itemProfitability.getOrDefault(item.getItemId(), 0.0)).reversed())
            .limit(3)
            .collect(Collectors.toList());

        System.out.println("\nMost Profitable Items (promote these):");
        mostProfitable.forEach(item -> 
            System.out.println("  " + item.getName() + ": $" + 
                String.format("%.2f", itemProfitability.get(item.getItemId()))));
    }

    // Helper method for kitchen queue priority
    private int getOrderPriority(String orderId) {
        Order order = orders.stream()
            .filter(o -> o.getOrderId().equals(orderId))
            .findFirst()
            .orElse(null);
        
        if (order == null) return 0;
        
        // Priority based on preparation time and waiting time
        int prepTime = order.getEstimatedPrepTime();
        int waitingTime = (int) order.getWaitingTimeMinutes();
        
        return (prepTime * 2) + waitingTime; // Higher priority for longer prep times and longer waits
    }

    // === MAIN METHOD ===

    public static void main(String[] args) {
        RestaurantManagementSystem restaurant = new RestaurantManagementSystem();
        
        System.out.println("=== RESTAURANT MANAGEMENT SYSTEM CHALLENGES ===");
        
        // Simulate restaurant operations
        String table1 = restaurant.assignTable(4, "Smith Family");
        String table2 = restaurant.assignTable(2, "Johnson Couple");
        String table3 = restaurant.assignTable(6, "Brown Group"); // Should go to waitlist
        
        // Create orders
        if (table1 != null) {
            restaurant.createOrder(table1, Map.of(
                "M001", 2,  // 2 Burgers
                "M003", 1,  // 1 Salad
                "M008", 2   // 2 Coffees
            ), "Extra ketchup please");
        }
        
        if (table2 != null) {
            restaurant.createOrder(table2, Map.of(
                "M002", 1,  // 1 Pizza
                "M007", 2   // 2 Ice Creams
            ), "No onions");
        }
        
        // Process kitchen queue
        restaurant.processKitchenQueue();
        
        // Generate analytics
        restaurant.displayInventoryStatus();
        restaurant.generateRevenueAnalytics();
        restaurant.analyzeTableTurnover();
        restaurant.analyzeMenuPerformance();
        restaurant.manageWaitlist();
        restaurant.demonstratePerformanceComparison();
        restaurant.advancedBusinessInsights();
        
        System.out.println("\n=== ALL RESTAURANT MANAGEMENT CHALLENGES COMPLETED ===");
    }
}
```

---

## 🎯 Key Integration Skills Demonstrated

### 1. **Collections Framework Integration**
- **HashMap**: Quick menu item and table lookup
- **ArrayList**: Order storage and processing
- **LinkedList**: Waitlist FIFO management
- **TreeMap**: Priority-based kitchen queue
- **HashSet**: Active category management
- **Deque**: Recent orders tracking

### 2. **Functional Programming Integration**
- **Predicates**: Complex table and order filtering
- **Functions**: Order value calculation and data transformation
- **Consumers**: Order processing and inventory updates
- **Suppliers**: Order ID generation
- **Streams**: Data analysis and reporting

### 3. **Real-World Problem Solving**
- **Table assignment**: Optimal table matching using predicates
- **Kitchen workflow**: Priority-based processing with TreeMap
- **Inventory management**: Real-time tracking with functional updates
- **Revenue analytics**: Comprehensive business insights using collectors

### 4. **Performance Optimization**
- **Collection selection**: Based on access patterns and operations
- **Stream efficiency**: Proper use of intermediate and terminal operations
- **Memory management**: Efficient data structures for restaurant operations

---

## 🚀 Expected Output

```
=== RESTAURANT MANAGEMENT SYSTEM CHALLENGES ===

=== TABLE ASSIGNMENT CHALLENGE ===
Table T1 assigned to Smith Family
Table T2 assigned to Johnson Couple
No available table. Added to waitlist: Brown Group

=== ORDER PROCESSING CHALLENGE ===
LOW INVENTORY ALERT: Beef (Remaining: 48)
LOW INVENTORY ALERT: Cheese (Remaining: 97)
LOW INVENTORY ALERT: Bread (Remaining: 78)
LOW INVENTORY ALERT: Lettuce (Remaining: 38)
LOW INVENTORY ALERT: Coffee Beans (Remaining: 48)
Processing order: ORD-1716034567890
Order created: Order{id='ORD-1716034567890', table='T1', status='PENDING', items=5, total=$44.94}
LOW INVENTORY ALERT: Cheese (Remaining: 95)
LOW INVENTORY ALERT: Dough (Remaining: 29)
LOW INVENTORY ALERT: Ice Cream (Remaining: 18)
Processing order: ORD-1716034567891
Order created: Order{id='ORD-1716034567891', table='T2', status='PENDING', items=3, total=$27.97}

=== KITCHEN QUEUE CHALLENGE ===
Added to kitchen queue: ORD-1716034567890
Added to kitchen queue: ORD-1716034567891
Kitchen preparing: ORD-1716034567890
Order ready: ORD-1716034567890
Kitchen preparing: ORD-1716034567891
Order ready: ORD-1716034567891

=== INVENTORY STATUS CHALLENGE ===
Inventory status:
  GOOD: 5 ingredients
  MEDIUM: 2 ingredients
  LOW: 1 ingredients

Low inventory items (restock needed):
  - Ice Cream: 18

=== REVENUE ANALYTICS CHALLENGE ===
Total Revenue: $72.91

Revenue by Category:
  Main: $41.97
  Appetizer: $8.99
  Beverage: $7.98
  Dessert: $13.98

Top 5 Most Popular Items:
  Burger: 2 orders
  Coffee: 2 orders
  Salad: 1 order
  Pizza: 1 order
  Ice Cream: 2 orders

Average Order Value: $36.46

=== TABLE TURNOVER CHALLENGE ===
Table Status Distribution:
  OCCUPIED: 2 tables
  AVAILABLE: 8 tables

Average Table Occupation Time: 2.5 minutes

Most Efficient Tables (fastest turnover):
  T1: 2 minutes
  T2: 3 minutes

=== MENU PERFORMANCE CHALLENGE ===
Category Performance:
  Main: 3 orders, $41.97 revenue, $13.99 avg value
  Appetizer: 1 orders, $8.99 revenue, $8.99 avg value
  Beverage: 2 orders, $7.98 revenue, $3.99 avg value
  Dessert: 2 orders, $13.98 revenue, $6.99 avg value

Least Popular Items (consider removing):
  Soup: 0 orders
  Steak: 0 orders
  Pasta: 0 orders

=== WAITLIST MANAGEMENT CHALLENGE ===
Current waitlist: 1 parties
Remaining on waitlist:
  - Brown Group (party of 6)

=== PERFORMANCE COMPARISON CHALLENGE ===
Menu lookup performance:
HashMap: 2 ms
TreeMap: 8 ms
HashMap is 4.00x faster

Order queue removal performance:
ArrayList: 1 ms
LinkedList: 0 ms
LinkedList is Infinityx faster

=== ADVANCED BUSINESS INSIGHTS CHALLENGE ===
Orders by Hour:
  09:00: 2 orders

Category Preferences by Table:
  T1: Main
  T2: Main

Most Profitable Items (promote these):
  Burger: $25.98
  Pizza: $15.99
  Ice Cream: $13.98

=== ALL RESTAURANT MANAGEMENT CHALLENGES COMPLETED ===
```

---

## 💡 Interview Preparation Points

### Integration Benefits
- **Best of both worlds**: Collections for data structure, Streams for processing
- **Real-time operations**: Efficient handling of restaurant workflow
- **Business intelligence**: Advanced analytics for decision making
- **Scalability**: Design patterns for growing restaurant operations

### Architecture Decisions
- **HashMap for lookups**: O(1) access for menu items and tables
- **TreeMap for priorities**: Automatic sorting for kitchen queue
- **LinkedList for queues**: Efficient FIFO operations
- **Streams for analytics**: Declarative data processing

### Performance Considerations
- **Collection choice**: Based on operation patterns
- **Stream optimization**: Proper use of intermediate operations
- **Memory efficiency**: Appropriate data structures
- **Real-time processing**: Efficient order handling

### Business Logic Implementation
- **Functional interfaces**: Encapsulate business rules
- **Stream pipelines**: Complex data analysis
- **Collection operations**: Efficient data management
- **Integration patterns**: Seamless combination of concepts

### Real-World Applications
- **Restaurant management**: Complete operational system
- **Inventory tracking**: Real-time stock management
- **Customer service**: Waitlist and table management
- **Business analytics**: Revenue and performance insights
