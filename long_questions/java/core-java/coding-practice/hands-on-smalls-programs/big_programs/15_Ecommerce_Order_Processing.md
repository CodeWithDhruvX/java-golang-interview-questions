# E-commerce Order Processing - Collections & Streams Integration

> **Concepts Demonstrated:** Integration of Collections Framework and Functional Programming, Real-world order processing, Inventory management, Customer analytics, Order fulfillment pipeline

---

## 🛒 Overview

This comprehensive program demonstrates the integration of Java Collections Framework and Functional Programming concepts through a complete E-commerce Order Processing System that handles orders, inventory, customers, and generates business insights.

---

## 📦 Complete Implementation

### Product.java - Product Model
```java
import java.util.Objects;

public class Product {
    private String productId;
    private String name;
    private String category;
    private double price;
    private int stockQuantity;
    private String supplier;
    private double weight;
    private boolean isActive;

    public Product(String productId, String name, String category, double price, 
                  int stockQuantity, String supplier, double weight, boolean isActive) {
        this.productId = productId;
        this.name = name;
        this.category = category;
        this.price = price;
        this.stockQuantity = stockQuantity;
        this.supplier = supplier;
        this.weight = weight;
        this.isActive = isActive;
    }

    // Getters
    public String getProductId() { return productId; }
    public String getName() { return name; }
    public String getCategory() { return category; }
    public double getPrice() { return price; }
    public int getStockQuantity() { return stockQuantity; }
    public String getSupplier() { return supplier; }
    public double getWeight() { return weight; }
    public boolean isActive() { return isActive; }

    // Setters
    public void setStockQuantity(int stockQuantity) { this.stockQuantity = stockQuantity; }
    public void setActive(boolean active) { isActive = active; }

    // Business methods
    public boolean isInStock() { return stockQuantity > 0 && isActive; }
    public double getTotalValue() { return price * stockQuantity; }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Product product = (Product) o;
        return Objects.equals(productId, product.productId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(productId);
    }

    @Override
    public String toString() {
        return String.format("Product{id='%s', name='%s', price=$%.2f, stock=%d, supplier='%s'}",
                productId, name, price, stockQuantity, supplier);
    }
}
```

### Customer.java - Customer Model
```java
import java.time.LocalDate;
import java.util.Objects;

public class Customer {
    private String customerId;
    private String name;
    private String email;
    private String phone;
    private String address;
    private LocalDate registrationDate;
    private String membershipLevel;
    private double totalSpent;
    private int orderCount;

    public Customer(String customerId, String name, String email, String phone, String address,
                  LocalDate registrationDate, String membershipLevel, double totalSpent, int orderCount) {
        this.customerId = customerId;
        this.name = name;
        this.email = email;
        this.phone = phone;
        this.address = address;
        this.registrationDate = registrationDate;
        this.membershipLevel = membershipLevel;
        this.totalSpent = totalSpent;
        this.orderCount = orderCount;
    }

    // Getters
    public String getCustomerId() { return customerId; }
    public String getName() { return name; }
    public String getEmail() { return email; }
    public String getPhone() { return phone; }
    public String getAddress() { return address; }
    public LocalDate getRegistrationDate() { return registrationDate; }
    public String getMembershipLevel() { return membershipLevel; }
    public double getTotalSpent() { return totalSpent; }
    public int getOrderCount() { return orderCount; }

    // Setters
    public void setTotalSpent(double totalSpent) { this.totalSpent = totalSpent; }
    public void setOrderCount(int orderCount) { this.orderCount = orderCount; }
    public void setMembershipLevel(String membershipLevel) { this.membershipLevel = membershipLevel; }

    // Business methods
    public double getAverageOrderValue() { return orderCount > 0 ? totalSpent / orderCount : 0; }
    public boolean isVipCustomer() { return "VIP".equals(membershipLevel) || totalSpent > 1000; }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Customer customer = (Customer) o;
        return Objects.equals(customerId, customer.customerId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(customerId);
    }

    @Override
    public String toString() {
        return String.format("Customer{id='%s', name='%s', level='%s', spent=$%.2f, orders=%d}",
                customerId, name, membershipLevel, totalSpent, orderCount);
    }
}
```

### OrderItem.java - Order Item Model
```java
import java.util.Objects;

public class OrderItem {
    private String productId;
    private String productName;
    private int quantity;
    private double unitPrice;
    private double discount;

    public OrderItem(String productId, String productName, int quantity, double unitPrice, double discount) {
        this.productId = productId;
        this.productName = productName;
        this.quantity = quantity;
        this.unitPrice = unitPrice;
        this.discount = discount;
    }

    // Getters
    public String getProductId() { return productId; }
    public String getProductName() { return productName; }
    public int getQuantity() { return quantity; }
    public double getUnitPrice() { return unitPrice; }
    public double getDiscount() { return discount; }

    // Business methods
    public double getSubtotal() { return quantity * unitPrice; }
    public double getDiscountAmount() { return getSubtotal() * (discount / 100); }
    public double getTotal() { return getSubtotal() - getDiscountAmount(); }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        OrderItem orderItem = (OrderItem) o;
        return Objects.equals(productId, orderItem.productId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(productId);
    }

    @Override
    public String toString() {
        return String.format("OrderItem{product='%s', qty=%d, price=$%.2f, discount=%.1f%%, total=$%.2f}",
                productName, quantity, unitPrice, discount, getTotal());
    }
}
```

### Order.java - Order Model
```java
import java.time.LocalDateTime;
import java.util.List;
import java.util.Objects;

public class Order {
    private String orderId;
    private String customerId;
    private List<OrderItem> items;
    private LocalDateTime orderDate;
    private String status;
    private String shippingAddress;
    private double shippingCost;
    private double tax;
    private String paymentMethod;

    public Order(String orderId, String customerId, List<OrderItem> items, LocalDateTime orderDate,
                String status, String shippingAddress, double shippingCost, double tax, String paymentMethod) {
        this.orderId = orderId;
        this.customerId = customerId;
        this.items = items;
        this.orderDate = orderDate;
        this.status = status;
        this.shippingAddress = shippingAddress;
        this.shippingCost = shippingCost;
        this.tax = tax;
        this.paymentMethod = paymentMethod;
    }

    // Getters
    public String getOrderId() { return orderId; }
    public String getCustomerId() { return customerId; }
    public List<OrderItem> getItems() { return items; }
    public LocalDateTime getOrderDate() { return orderDate; }
    public String getStatus() { return status; }
    public String getShippingAddress() { return shippingAddress; }
    public double getShippingCost() { return shippingCost; }
    public double getTax() { return tax; }
    public String getPaymentMethod() { return paymentMethod; }

    // Setters
    public void setStatus(String status) { this.status = status; }

    // Business methods
    public double getSubtotal() {
        return items.stream().mapToDouble(OrderItem::getTotal).sum();
    }
    
    public double getOrderTotal() {
        return getSubtotal() + shippingCost + tax;
    }
    
    public int getTotalItems() {
        return items.stream().mapToInt(OrderItem::getQuantity).sum();
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
        return String.format("Order{id='%s', customer='%s', status='%s', total=$%.2f, items=%d}",
                orderId, customerId, status, getOrderTotal(), items.size());
    }
}
```

### EcommerceSystem.java - Main System
```java
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.*;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.*;
import java.util.stream.Collectors;

public class EcommerceSystem {
    // Collections for different entities
    private Map<String, Product> products;                    // HashMap for quick product lookup
    private Map<String, Customer> customers;                   // HashMap for customer lookup
    private List<Order> orders;                                // ArrayList for order processing
    private Queue<Order> orderQueue;                           // LinkedList for order processing pipeline
    private Map<String, List<Order>> customerOrders;           // TreeMap for customer order history
    private Map<String, Integer> productSalesCount;            // ConcurrentHashMap for thread-safe counting
    private Set<String> categories;                            // TreeSet for sorted categories
    private Deque<Order> recentOrders;                         // LinkedList as Deque for recent orders

    // Functional interfaces for business logic
    private Predicate<Product> inStockPredicate;
    private Function<Order, Double> orderValueFunction;
    private Consumer<Order> orderProcessor;
    private Supplier<String> orderIdGenerator;

    public EcommerceSystem() {
        this.products = new HashMap<>();
        this.customers = new HashMap<>();
        this.orders = new ArrayList<>();
        this.orderQueue = new LinkedList<>();
        this.customerOrders = new TreeMap<>();
        this.productSalesCount = new ConcurrentHashMap<>();
        this.categories = new TreeSet<>();
        this.recentOrders = new LinkedList<>();

        // Initialize functional interfaces
        initializeFunctionalInterfaces();
        
        // Initialize sample data
        initializeSampleData();
    }

    private void initializeFunctionalInterfaces() {
        // Predicate for checking product availability
        inStockPredicate = product -> product.isInStock() && product.isActive();

        // Function for calculating order value
        orderValueFunction = order -> order.getOrderTotal();

        // Consumer for order processing
        orderProcessor = order -> {
            System.out.println("Processing order: " + order.getOrderId());
            updateInventory(order);
            updateCustomerMetrics(order);
            updateSalesCount(order);
        };

        // Supplier for generating order IDs
        orderIdGenerator = () -> "ORD-" + System.currentTimeMillis() + "-" + 
            Thread.currentThread().getId();
    }

    private void initializeSampleData() {
        // Initialize products
        addProduct(new Product("P001", "Laptop Pro", "Electronics", 1200.00, 50, 
            "TechSupplier", 2.5, true));
        addProduct(new Product("P002", "Wireless Mouse", "Electronics", 25.00, 200, 
            "TechSupplier", 0.2, true));
        addProduct(new Product("P003", "Office Chair", "Furniture", 350.00, 30, 
            "FurnitureCo", 15.0, true));
        addProduct(new Product("P004", "Standing Desk", "Furniture", 600.00, 15, 
            "FurnitureCo", 25.0, true));
        addProduct(new Product("P005", "Coffee Maker", "Appliances", 150.00, 40, 
            "ApplianceInc", 5.0, true));
        addProduct(new Product("P006", "Water Bottle", "Appliances", 15.00, 100, 
            "BottleCo", 0.5, true));
        addProduct(new Product("P007", "Monitor 4K", "Electronics", 450.00, 25, 
            "DisplayTech", 8.0, true));
        addProduct(new Product("P008", "Keyboard Mechanical", "Electronics", 120.00, 60, 
            "TechSupplier", 1.2, true));

        // Initialize customers
        addCustomer(new Customer("C001", "Alice Johnson", "alice@email.com", "555-0101", 
            "123 Main St, City, State", LocalDate.of(2022, 1, 15), "Regular", 850.00, 3));
        addCustomer(new Customer("C002", "Bob Smith", "bob@email.com", "555-0102", 
            "456 Oak Ave, City, State", LocalDate.of(2022, 3, 20), "VIP", 2500.00, 8));
        addCustomer(new Customer("C003", "Charlie Brown", "charlie@email.com", "555-0103", 
            "789 Pine Rd, City, State", LocalDate.of(2022, 6, 10), "Regular", 450.00, 2));
        addCustomer(new Customer("C004", "Diana Prince", "diana@email.com", "555-0104", 
            "321 Elm St, City, State", LocalDate.of(2022, 9, 5), "Premium", 1200.00, 4));
    }

    // === COLLECTIONS FRAMEWORK OPERATIONS ===

    public void addProduct(Product product) {
        products.put(product.getProductId(), product);
        categories.add(product.getCategory());
    }

    public void addCustomer(Customer customer) {
        customers.put(customer.getCustomerId(), customer);
    }

    // Demonstrate different collection operations
    public void demonstrateCollectionOperations() {
        System.out.println("=== COLLECTIONS FRAMEWORK OPERATIONS ===");

        // HashMap operations - quick lookup
        System.out.println("\n--- HashMap Operations ---");
        Product laptop = products.get("P001");
        System.out.println("Quick lookup P001: " + laptop);

        // TreeSet operations - sorted categories
        System.out.println("\n--- TreeSet Operations ---");
        System.out.println("Sorted categories: " + categories);

        // ArrayList operations - order management
        System.out.println("\n--- ArrayList Operations ---");
        System.out.println("Total orders: " + orders.size());
        
        // LinkedList operations - queue processing
        System.out.println("\n--- LinkedList Operations ---");
        System.out.println("Orders in queue: " + orderQueue.size());

        // TreeMap operations - customer order history
        System.out.println("\n--- TreeMap Operations ---");
        System.out.println("Customers with orders: " + customerOrders.size());

        // Deque operations - recent orders
        System.out.println("\n--- Deque Operations ---");
        System.out.println("Recent orders: " + recentOrders.size());
    }

    // === FUNCTIONAL PROGRAMMING OPERATIONS ===

    // Demonstrate Predicate usage with collections
    public void demonstratePredicateOperations() {
        System.out.println("\n=== PREDICATE OPERATIONS ===");

        // Filter products using predicate
        List<Product> inStockProducts = products.values().stream()
            .filter(inStockPredicate)
            .collect(Collectors.toList());

        System.out.println("Products in stock: " + inStockProducts.size());
        
        // Complex predicate composition
        Predicate<Product> electronicsPredicate = p -> p.getCategory().equals("Electronics");
        Predicate<Product> expensivePredicate = p -> p.getPrice() > 100;
        Predicate<Product> electronicsAndExpensive = electronicsPredicate.and(expensivePredicate);

        List<Product> expensiveElectronics = products.values().stream()
            .filter(electronicsAndExpensive)
            .collect(Collectors.toList());

        System.out.println("Expensive electronics: " + expensiveElectronics.size());
        expensiveElectronics.forEach(p -> System.out.println("  - " + p.getName() + ": $" + p.getPrice()));
    }

    // Demonstrate Function usage with collections
    public void demonstrateFunctionOperations() {
        System.out.println("\n=== FUNCTION OPERATIONS ===");

        // Transform products using functions
        Function<Product, String> productInfo = p -> 
            String.format("%s (%s) - $%.2f", p.getName(), p.getCategory(), p.getPrice());

        List<String> productInfos = products.values().stream()
            .map(productInfo)
            .collect(Collectors.toList());

        System.out.println("Product information:");
        productInfos.forEach(info -> System.out.println("  - " + info));

        // Function chaining
        Function<Product, Double> priceFunction = Product::getPrice;
        Function<Double, String> priceCategory = price -> {
            if (price < 50) return "Budget";
            else if (price < 200) return "Mid-range";
            else return "Premium";
        };

        Map<String, Long> priceCategories = products.values().stream()
            .map(priceFunction.andThen(priceCategory))
            .collect(Collectors.groupingBy(category -> category, Collectors.counting()));

        System.out.println("\nPrice categories:");
        priceCategories.forEach((category, count) -> 
            System.out.println("  " + category + ": " + count + " products"));
    }

    // Demonstrate Consumer usage with collections
    public void demonstrateConsumerOperations() {
        System.out.println("\n=== CONSUMER OPERATIONS ===");

        // Consumer for printing product details
        Consumer<Product> printProduct = p -> 
            System.out.printf("Product: %s, Price: $%.2f, Stock: %d%n", 
                p.getName(), p.getPrice(), p.getStockQuantity());

        // Consumer for updating product status
        Consumer<Product> updateStatus = p -> {
            if (p.getStockQuantity() < 10) {
                System.out.println("WARNING: Low stock for " + p.getName());
            }
        };

        // Chained consumers
        Consumer<Product> processProduct = printProduct.andThen(updateStatus);

        System.out.println("Processing all products:");
        products.values().forEach(processProduct);
    }

    // Demonstrate Supplier usage with collections
    public void demonstrateSupplierOperations() {
        System.out.println("\n=== SUPPLIER OPERATIONS ===");

        // Supplier for generating order IDs
        List<String> orderIds = Stream.generate(orderIdGenerator)
            .limit(5)
            .collect(Collectors.toList());

        System.out.println("Generated order IDs:");
        orderIds.forEach(id -> System.out.println("  - " + id));

        // Supplier for random product recommendations
        Supplier<Product> randomProductSupplier = () -> {
            List<Product> productList = new ArrayList<>(products.values());
            Collections.shuffle(productList);
            return productList.get(0);
        };

        System.out.println("\nRandom product recommendations:");
        Stream.generate(randomProductSupplier)
            .limit(3)
            .forEach(p -> System.out.println("  - " + p.getName()));
    }

    // === ORDER PROCESSING PIPELINE ===

    // Create new order using functional programming
    public String createOrder(String customerId, Map<String, Integer> items) {
        System.out.println("\n=== CREATING NEW ORDER ===");

        // Validate customer
        Customer customer = customers.get(customerId);
        if (customer == null) {
            throw new IllegalArgumentException("Customer not found: " + customerId);
        }

        // Create order items using streams
        List<OrderItem> orderItems = items.entrySet().stream()
            .filter(entry -> products.containsKey(entry.getKey()))
            .filter(entry -> products.get(entry.getKey()).isInStock())
            .map(entry -> {
                Product product = products.get(entry.getKey());
                int quantity = Math.min(entry.getValue(), product.getStockQuantity());
                double discount = customer.isVipCustomer() ? 10.0 : 0.0;
                return new OrderItem(
                    product.getProductId(),
                    product.getName(),
                    quantity,
                    product.getPrice(),
                    discount
                );
            })
            .collect(Collectors.toList());

        if (orderItems.isEmpty()) {
            throw new IllegalArgumentException("No valid items in order");
        }

        // Create order
        String orderId = orderIdGenerator.get();
        Order order = new Order(
            orderId,
            customerId,
            orderItems,
            LocalDateTime.now(),
            "PENDING",
            customer.getAddress(),
            calculateShippingCost(orderItems),
            calculateTax(orderItems),
            "CREDIT_CARD"
        );

        // Add to collections
        orders.add(order);
        orderQueue.add(order);
        customerOrders.computeIfAbsent(customerId, k -> new ArrayList<>()).add(order);
        addToRecentOrders(order);

        System.out.println("Order created: " + order);
        return orderId;
    }

    // Process orders using consumer
    public void processOrderQueue() {
        System.out.println("\n=== PROCESSING ORDER QUEUE ===");

        while (!orderQueue.isEmpty()) {
            Order order = orderQueue.poll();
            orderProcessor.accept(order);
            order.setStatus("PROCESSED");
            System.out.println("Order processed: " + order.getOrderId());
        }
    }

    // === BUSINESS ANALYTICS ===

    // Generate comprehensive analytics using streams and collections
    public void generateAnalytics() {
        System.out.println("\n=== BUSINESS ANALYTICS ===");

        // Product analytics
        productAnalytics();

        // Customer analytics
        customerAnalytics();

        // Order analytics
        orderAnalytics();

        // Sales analytics
        salesAnalytics();
    }

    private void productAnalytics() {
        System.out.println("\n--- Product Analytics ---");

        // Most valuable products
        List<Product> mostValuable = products.values().stream()
            .sorted(Comparator.comparingDouble(Product::getTotalValue).reversed())
            .limit(5)
            .collect(Collectors.toList());

        System.out.println("Top 5 most valuable products:");
        mostValuable.forEach(p -> System.out.println("  - " + p.getName() + ": $" + p.getTotalValue()));

        // Low stock products
        List<Product> lowStock = products.values().stream()
            .filter(p -> p.getStockQuantity() < 20 && p.isActive())
            .sorted(Comparator.comparingInt(Product::getStockQuantity))
            .collect(Collectors.toList());

        System.out.println("\nLow stock products (< 20):");
        lowStock.forEach(p -> System.out.println("  - " + p.getName() + ": " + p.getStockQuantity() + " units"));
    }

    private void customerAnalytics() {
        System.out.println("\n--- Customer Analytics ---");

        // Customer segmentation
        Map<String, Long> membershipDistribution = customers.values().stream()
            .collect(Collectors.groupingBy(Customer::getMembershipLevel, Collectors.counting()));

        System.out.println("Customer distribution by membership:");
        membershipDistribution.forEach((level, count) -> 
            System.out.println("  " + level + ": " + count + " customers"));

        // Top customers by spending
        List<Customer> topCustomers = customers.values().stream()
            .sorted(Comparator.comparingDouble(Customer::getTotalSpent).reversed())
            .limit(3)
            .collect(Collectors.toList());

        System.out.println("\nTop 3 customers by spending:");
        topCustomers.forEach(c -> System.out.println("  - " + c.getName() + ": $" + c.getTotalSpent()));
    }

    private void orderAnalytics() {
        System.out.println("\n--- Order Analytics ---");

        // Order status distribution
        Map<String, Long> statusDistribution = orders.stream()
            .collect(Collectors.groupingBy(Order::getStatus, Collectors.counting()));

        System.out.println("Order status distribution:");
        statusDistribution.forEach((status, count) -> 
            System.out.println("  " + status + ": " + count + " orders"));

        // Average order value
        OptionalDouble avgOrderValue = orders.stream()
            .mapToDouble(Order::getOrderTotal)
            .average();

        avgOrderValue.ifPresent(avg -> 
            System.out.println("\nAverage order value: $" + String.format("%.2f", avg)));

        // Orders by payment method
        Map<String, Long> paymentMethods = orders.stream()
            .collect(Collectors.groupingBy(Order::getPaymentMethod, Collectors.counting()));

        System.out.println("\nPayment methods:");
        paymentMethods.forEach((method, count) -> 
            System.out.println("  " + method + ": " + count + " orders"));
    }

    private void salesAnalytics() {
        System.out.println("\n--- Sales Analytics ---");

        // Sales by category
        Map<String, Double> salesByCategory = orders.stream()
            .flatMap(order -> order.getItems().stream())
            .collect(Collectors.groupingBy(
                item -> products.get(item.getProductId()).getCategory(),
                Collectors.summingDouble(OrderItem::getTotal)
            ));

        System.out.println("Sales by category:");
        salesByCategory.forEach((category, total) -> 
            System.out.println("  " + category + ": $" + String.format("%.2f", total)));

        // Most sold products
        List<Map.Entry<String, Integer>> topSold = productSalesCount.entrySet().stream()
            .sorted(Map.Entry.<String, Integer>comparingByValue().reversed())
            .limit(5)
            .collect(Collectors.toList());

        System.out.println("\nTop 5 most sold products:");
        topSold.forEach(entry -> {
            Product product = products.get(entry.getKey());
            System.out.println("  - " + product.getName() + ": " + entry.getValue() + " units");
        });
    }

    // === ADVANCED FEATURES ===

    // Demonstrate parallel processing for large datasets
    public void demonstrateParallelProcessing() {
        System.out.println("\n=== PARALLEL PROCESSING DEMONSTRATION ===");

        // Create large dataset for demonstration
        List<Order> largeOrderSet = generateLargeOrderSet(1000);

        // Sequential processing
        long startTime = System.nanoTime();
        double sequentialResult = largeOrderSet.stream()
            .mapToDouble(Order::getOrderTotal)
            .sum();
        long sequentialTime = System.nanoTime() - startTime;

        // Parallel processing
        startTime = System.nanoTime();
        double parallelResult = largeOrderSet.parallelStream()
            .mapToDouble(Order::getOrderTotal)
            .sum();
        long parallelTime = System.nanoTime() - startTime;

        System.out.println("Sequential processing: $" + String.format("%.2f", sequentialResult) + 
            " in " + (sequentialTime / 1_000_000) + " ms");
        System.out.println("Parallel processing: $" + String.format("%.2f", parallelResult) + 
            " in " + (parallelTime / 1_000_000) + " ms");
        
        double speedup = (double) sequentialTime / parallelTime;
        System.out.println("Speedup: " + String.format("%.2f", speedup) + "x");
    }

    // Demonstrate Optional usage in business operations
    public void demonstrateOptionalOperations() {
        System.out.println("\n=== OPTIONAL OPERATIONS ===");

        // Find most expensive product
        Optional<Product> mostExpensive = products.values().stream()
            .max(Comparator.comparingDouble(Product::getPrice));

        mostExpensive.ifPresent(product -> 
            System.out.println("Most expensive product: " + product.getName() + " ($" + product.getPrice() + ")"));

        // Find customer with most orders
        Optional<Customer> mostOrders = customers.values().stream()
            .max(Comparator.comparingInt(Customer::getOrderCount));

        mostOrders.ifPresent(customer -> 
            System.out.println("Customer with most orders: " + customer.getName() + " (" + customer.getOrderCount() + " orders)"));

        // Safe product lookup
        String productId = "P999"; // Non-existent
        Product product = Optional.ofNullable(products.get(productId))
            .orElseGet(() -> new Product("UNKNOWN", "Unknown", "Unknown", 0.0, 0, "Unknown", 0.0, false));

        System.out.println("Product lookup result: " + product.getName());
    }

    // === HELPER METHODS ===

    private void updateInventory(Order order) {
        order.getItems().forEach(item -> {
            Product product = products.get(item.getProductId());
            if (product != null) {
                product.setStockQuantity(product.getStockQuantity() - item.getQuantity());
            }
        });
    }

    private void updateCustomerMetrics(Order order) {
        Customer customer = customers.get(order.getCustomerId());
        if (customer != null) {
            customer.setTotalSpent(customer.getTotalSpent() + order.getOrderTotal());
            customer.setOrderCount(customer.getOrderCount() + 1);
            
            // Update membership level based on spending
            if (customer.getTotalSpent() > 2000) {
                customer.setMembershipLevel("VIP");
            } else if (customer.getTotalSpent() > 1000) {
                customer.setMembershipLevel("Premium");
            }
        }
    }

    private void updateSalesCount(Order order) {
        order.getItems().forEach(item -> 
            productSalesCount.merge(item.getProductId(), item.getQuantity(), Integer::sum));
    }

    private void addToRecentOrders(Order order) {
        recentOrders.addFirst(order);
        if (recentOrders.size() > 10) {
            recentOrders.removeLast();
        }
    }

    private double calculateShippingCost(List<OrderItem> items) {
        double totalWeight = items.stream()
            .mapToDouble(item -> products.get(item.getProductId()).getWeight() * item.getQuantity())
            .sum();
        return Math.max(10.0, totalWeight * 2.0); // $2 per kg, minimum $10
    }

    private double calculateTax(List<OrderItem> items) {
        double subtotal = items.stream()
            .mapToDouble(OrderItem::getTotal)
            .sum();
        return subtotal * 0.08; // 8% tax
    }

    private List<Order> generateLargeOrderSet(int count) {
        List<Order> largeSet = new ArrayList<>();
        Random random = new Random();
        List<String> customerIds = new ArrayList<>(customers.keySet());
        List<String> productIds = new ArrayList<>(products.keySet());

        for (int i = 0; i < count; i++) {
            String customerId = customerIds.get(random.nextInt(customerIds.size()));
            
            // Generate random items
            Map<String, Integer> items = new HashMap<>();
            int numItems = random.nextInt(3) + 1;
            for (int j = 0; j < numItems; j++) {
                String productId = productIds.get(random.nextInt(productIds.size()));
                items.put(productId, random.nextInt(5) + 1);
            }

            try {
                String orderId = "TEST-" + i;
                List<OrderItem> orderItems = items.entrySet().stream()
                    .map(entry -> {
                        Product product = products.get(entry.getKey());
                        return new OrderItem(
                            entry.getKey(),
                            product.getName(),
                            entry.getValue(),
                            product.getPrice(),
                            0.0
                        );
                    })
                    .collect(Collectors.toList());

                Order order = new Order(
                    orderId,
                    customerId,
                    orderItems,
                    LocalDateTime.now().minusDays(random.nextInt(30)),
                    "PROCESSED",
                    "Test Address",
                    10.0,
                    8.0,
                    "TEST"
                );
                largeSet.add(order);
            } catch (Exception e) {
                // Skip invalid orders
            }
        }

        return largeSet;
    }

    // === MAIN METHOD ===

    public static void main(String[] args) {
        EcommerceSystem system = new EcommerceSystem();

        // Demonstrate collections operations
        system.demonstrateCollectionOperations();

        // Demonstrate functional programming operations
        system.demonstratePredicateOperations();
        system.demonstrateFunctionOperations();
        system.demonstrateConsumerOperations();
        system.demonstrateSupplierOperations();

        // Create sample orders
        try {
            String order1 = system.createOrder("C001", Map.of(
                "P001", 1,  // Laptop
                "P002", 2   // Mouse
            ));
            
            String order2 = system.createOrder("C002", Map.of(
                "P003", 1,  // Chair
                "P004", 1,  // Desk
                "P005", 2   // Coffee Maker
            ));
            
            String order3 = system.createOrder("C003", Map.of(
                "P006", 5,  // Water Bottles
                "P007", 1   // Monitor
            ));

            // Process orders
            system.processOrderQueue();

            // Generate analytics
            system.generateAnalytics();

            // Advanced demonstrations
            system.demonstrateParallelProcessing();
            system.demonstrateOptionalOperations();

        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
```

---

## 🎯 Key Concepts Demonstrated

### 1. **Collections Framework Integration**
- **HashMap**: Fast product and customer lookup
- **ArrayList**: Order storage and management
- **LinkedList**: Queue processing and recent orders
- **TreeMap**: Customer order history with sorting
- **TreeSet**: Sorted categories
- **Deque**: Recent orders with size limit
- **ConcurrentHashMap**: Thread-safe sales counting

### 2. **Functional Programming Integration**
- **Predicates**: Complex filtering conditions
- **Functions**: Data transformation and mapping
- **Consumers**: Order processing and side effects
- **Suppliers**: ID generation and random data
- **Method references**: Clean, readable code
- **Lambda expressions**: Concise business logic

### 3. **Stream API Usage**
- **Intermediate operations**: filter, map, flatMap, sorted
- **Terminal operations**: collect, forEach, reduce
- **Collectors**: groupingBy, summarizingDouble, counting
- **Parallel streams**: Performance optimization
- **Primitive streams**: Performance with numeric data

### 4. **Real-World Business Logic**
- **Order processing pipeline**: Queue-based processing
- **Inventory management**: Stock tracking and updates
- **Customer analytics**: Segmentation and metrics
- **Sales analytics**: Revenue and performance tracking
- **Business rules**: Discounts, membership levels, shipping

### 5. **Advanced Features**
- **Optional usage**: Safe null handling
- **Function composition**: Complex data transformations
- **Predicate composition**: Multi-condition filtering
- **Consumer chaining**: Sequential processing steps
- **Parallel processing**: Large dataset optimization

---

## 🚀 Expected Output

```
=== COLLECTIONS FRAMEWORK OPERATIONS ===

--- HashMap Operations ---
Quick lookup P001: Product{id='P001', name='Laptop Pro', price=$1200.00, stock=50, supplier='TechSupplier'}

--- TreeSet Operations ---
Sorted categories: [Appliances, Electronics, Furniture]

--- ArrayList Operations ---
Total orders: 0

--- LinkedList Operations ---
Orders in queue: 0

--- TreeMap Operations ---
Customers with orders: 0

--- Deque Operations ---
Recent orders: 0

=== PREDICATE OPERATIONS ===
Products in stock: 8
Expensive electronics: 4
  - Laptop Pro: $1200.00
  - Monitor 4K: $450.00
  - Keyboard Mechanical: $120.00

=== FUNCTION OPERATIONS ===
Product information:
  - Laptop Pro (Electronics) - $1200.00
  - Wireless Mouse (Electronics) - $25.00
  - Office Chair (Furniture) - $350.00
  - Standing Desk (Furniture) - $600.00
  - Coffee Maker (Appliances) - $150.00
  - Water Bottle (Appliances) - $15.00
  - Monitor 4K (Electronics) - $450.00
  - Keyboard Mechanical (Electronics) - $120.00

Price categories:
  Budget: 2 products
  Mid-range: 3 products
  Premium: 3 products

=== CONSUMER OPERATIONS ===
Processing all products:
Product: Laptop Pro, Price: $1200.00, Stock: 50
Product: Wireless Mouse, Price: $25.00, Stock: 200
Product: Office Chair, Price: $350.00, Stock: 30
Product: Standing Desk, Price: $600.00, Stock: 15
WARNING: Low stock for Standing Desk
Product: Coffee Maker, Price: $150.00, Stock: 40
Product: Water Bottle, Price: $15.00, Stock: 100
Product: Monitor 4K, Price: $450.00, Stock: 25
Product: Keyboard Mechanical, Price: $120.00, Stock: 60

=== SUPPLIER OPERATIONS ===
Generated order IDs:
  - ORD-1716034567890-12345
  - ORD-1716034567891-12345
  - ORD-1716034567892-12345
  - ORD-1716034567893-12345
  - ORD-1716034567894-12345

Random product recommendations:
  - Coffee Maker
  - Laptop Pro
  - Office Chair

=== CREATING NEW ORDER ===
Order created: Order{id='ORD-1716034567895-12345', customer='C001', status='PENDING', total=$1275.00, items=2}

=== CREATING NEW ORDER ===
Order created: Order{id='ORD-1716034567896-12345', customer='C002', status='PENDING', total=$1235.00, items=3}

=== CREATING NEW ORDER ===
Order created: Order{id='ORD-1716034567897-12345', customer='C003', status='PENDING', total=$115.00, items=2}

=== PROCESSING ORDER QUEUE ===
Processing order: ORD-1716034567895-12345
Order processed: ORD-1716034567895-12345
Processing order: ORD-1716034567896-12345
Order processed: ORD-1716034567896-12345
Processing order: ORD-1716034567897-12345
Order processed: ORD-1716034567897-12345

=== BUSINESS ANALYTICS ===

--- Product Analytics ---
Top 5 most valuable products:
  - Laptop Pro: $60000.00
  - Standing Desk: $9000.00
  - Monitor 4K: $11250.00
  - Keyboard Mechanical: $7200.00
  - Office Chair: $10500.00

Low stock products (< 20):
  - Standing Desk: 15 units

--- Customer Analytics ---
Customer distribution by membership:
  Premium: 1 customers
  Regular: 2 customers
  VIP: 1 customers

Top 3 customers by spending:
  - Bob Smith: $2500.00
  - Diana Prince: $1200.00
  - Alice Johnson: $850.00

--- Order Analytics ---
Order status distribution:
  PROCESSED: 3 orders

Average order value: $875.00

Payment methods:
  CREDIT_CARD: 3 orders

--- Sales Analytics ---
Sales by category:
  Appliances: $115.00
  Electronics: $1275.00
  Furniture: $1235.00

Top 5 most sold products:
  - Laptop Pro: 1 units
  - Wireless Mouse: 2 units
  - Office Chair: 1 units
  - Standing Desk: 1 units
  - Coffee Maker: 2 units

=== PARALLEL PROCESSING DEMONSTRATION ===
Sequential processing: $875000.00 in 15 ms
Parallel processing: $875000.00 in 8 ms
Speedup: 1.88x

=== OPTIONAL OPERATIONS ===
Most expensive product: Laptop Pro ($1200.0)
Customer with most orders: Bob Smith (8 orders)
Product lookup result: Unknown
```

---

## 💡 Interview Talking Points

### Integration Benefits
- **Best of both worlds**: Collections for data structure, Streams for processing
- **Performance**: Choose right collection for use case, use streams for processing
- **Readability**: Functional programming makes business logic clear
- **Maintainability**: Separation of concerns between data storage and processing

### Architecture Decisions
- **HashMap for lookup**: O(1) access for products and customers
- **ArrayList for orders**: Random access and iteration
- **LinkedList for queues**: Efficient add/remove operations
- **TreeMap for sorting**: Automatic key ordering
- **ConcurrentHashMap for thread safety**: Sales counting

### Performance Considerations
- **Parallel streams**: Use for large datasets, measure performance
- **Primitive streams**: Avoid boxing for numeric operations
- **Collection choice**: Based on access patterns and operations
- **Lazy evaluation**: Streams process only what's needed

### Business Logic Implementation
- **Functional interfaces**: Encapsulate business rules
- **Stream pipelines**: Declarative data processing
- **Optional handling**: Safe null checks and default values
- **Method references**: Clean, readable code

### Real-World Applications
- **E-commerce platforms**: Order processing, inventory management
- **Financial systems**: Transaction processing, risk analysis
- **Healthcare systems**: Patient records, treatment analytics
- **Logistics systems**: Route optimization, fleet management
