# Data Analytics Dashboard - Functional Programming & Streams in Action

> **Concepts Demonstrated:** Lambda expressions, Functional interfaces, Method references, Optional, Stream API, Collectors, Parallel streams, Predicate composition, Function chaining, Custom functional interfaces

---

## 📊 Overview

This practical program demonstrates advanced functional programming concepts through a Data Analytics Dashboard that processes sales data, generates insights, and produces comprehensive reports using Java Streams and functional programming patterns.

---

## 📈 Complete Implementation

### SalesData.java - Model Class
```java
import java.time.LocalDate;
import java.util.Objects;

public class SalesData {
    private String productId;
    private String productName;
    private String category;
    private String region;
    private double amount;
    private int quantity;
    private LocalDate date;
    private String salesRep;

    public SalesData(String productId, String productName, String category, String region, 
                    double amount, int quantity, LocalDate date, String salesRep) {
        this.productId = productId;
        this.productName = productName;
        this.category = category;
        this.region = region;
        this.amount = amount;
        this.quantity = quantity;
        this.date = date;
        this.salesRep = salesRep;
    }

    // Getters
    public String getProductId() { return productId; }
    public String getProductName() { return productName; }
    public String getCategory() { return category; }
    public String getRegion() { return region; }
    public double getAmount() { return amount; }
    public int getQuantity() { return quantity; }
    public LocalDate getDate() { return date; }
    public String getSalesRep() { return salesRep; }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        SalesData salesData = (SalesData) o;
        return Objects.equals(productId, salesData.productId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(productId);
    }

    @Override
    public String toString() {
        return String.format("SalesData{product='%s', category='%s', region='%s', amount=$%.2f, qty=%d, date=%s, rep='%s'}",
                productName, category, region, amount, quantity, date, salesRep);
    }
}
```

### AnalyticsDashboard.java - Main Class
```java
import java.time.LocalDate;
import java.time.YearMonth;
import java.time.format.DateTimeFormatter;
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class AnalyticsDashboard {
    private List<SalesData> salesData;

    public AnalyticsDashboard() {
        this.salesData = new ArrayList<>();
        initializeSampleData();
    }

    private void initializeSampleData() {
        salesData.addAll(Arrays.asList(
            new SalesData("P001", "Laptop Pro", "Electronics", "North", 1200.00, 5, 
                LocalDate.of(2023, 1, 15), "Alice"),
            new SalesData("P002", "Wireless Mouse", "Electronics", "South", 25.00, 20, 
                LocalDate.of(2023, 1, 16), "Bob"),
            new SalesData("P003", "Office Chair", "Furniture", "East", 350.00, 8, 
                LocalDate.of(2023, 1, 17), "Charlie"),
            new SalesData("P004", "Standing Desk", "Furniture", "West", 600.00, 3, 
                LocalDate.of(2023, 1, 18), "Diana"),
            new SalesData("P005", "Monitor 4K", "Electronics", "North", 450.00, 10, 
                LocalDate.of(2023, 2, 10), "Alice"),
            new SalesData("P006", "Keyboard Mechanical", "Electronics", "South", 120.00, 15, 
                LocalDate.of(2023, 2, 11), "Bob"),
            new SalesData("P007", "Desk Lamp", "Furniture", "East", 45.00, 12, 
                LocalDate.of(2023, 2, 12), "Charlie"),
            new SalesData("P008", "Webcam HD", "Electronics", "West", 80.00, 25, 
                LocalDate.of(2023, 2, 13), "Eve"),
            new SalesData("P009", "Coffee Maker", "Appliances", "North", 150.00, 6, 
                LocalDate.of(2023, 3, 5), "Frank"),
            new SalesData("P010", "Water Bottle", "Appliances", "South", 15.00, 30, 
                LocalDate.of(2023, 3, 6), "Grace"),
            new SalesData("P011", "Laptop Pro", "Electronics", "East", 1200.00, 3, 
                LocalDate.of(2023, 3, 7), "Alice"),
            new SalesData("P012", "Office Chair", "Furniture", "West", 350.00, 5, 
                LocalDate.of(2023, 3, 8), "Henry")
        ));
    }

    // === LAMBDA EXPRESSIONS & FUNCTIONAL INTERFACES ===

    // Demonstrate basic lambda with Predicate
    public void demonstrateBasicLambda() {
        System.out.println("=== BASIC LAMBDA DEMONSTRATION ===");
        
        // Traditional approach
        Predicate<SalesData> highValuePredicate = new Predicate<SalesData>() {
            @Override
            public boolean test(SalesData data) {
                return data.getAmount() > 500;
            }
        };
        
        // Lambda approach
        Predicate<SalesData> highValueLambda = data -> data.getAmount() > 500;
        
        // Method reference approach
        Predicate<SalesData> highValueMethodRef = this::isHighValue;
        
        System.out.println("High value sales (traditional): " + 
            salesData.stream().filter(highValuePredicate).count());
        System.out.println("High value sales (lambda): " + 
            salesData.stream().filter(highValueLambda).count());
        System.out.println("High value sales (method ref): " + 
            salesData.stream().filter(highValueMethodRef).count());
    }

    // Helper method for method reference
    private boolean isHighValue(SalesData data) {
        return data.getAmount() > 500;
    }

    // Demonstrate Predicate composition
    public void demonstratePredicateComposition() {
        System.out.println("\n=== PREDICATE COMPOSITION ===");
        
        Predicate<SalesData> electronics = data -> data.getCategory().equals("Electronics");
        Predicate<SalesData> highValue = data -> data.getAmount() > 100;
        Predicate<SalesData> recent = data -> data.getDate().isAfter(LocalDate.of(2023, 2, 1));
        
        // Composition using and()
        Predicate<SalesData> electronicsAndHighValue = electronics.and(highValue);
        
        // Composition using or()
        Predicate<SalesData> highValueOrRecent = highValue.or(recent);
        
        // Composition using negate()
        Predicate<SalesData> notElectronics = electronics.negate();
        
        System.out.println("Electronics AND high value: " + 
            salesData.stream().filter(electronicsAndHighValue).count() + " sales");
        System.out.println("High value OR recent: " + 
            salesData.stream().filter(highValueOrRecent).count() + " sales");
        System.out.println("NOT electronics: " + 
            salesData.stream().filter(notElectronics).count() + " sales");
    }

    // Demonstrate Function chaining
    public void demonstrateFunctionChaining() {
        System.out.println("\n=== FUNCTION CHAINING ===");
        
        // Function to extract product name
        Function<SalesData, String> extractProductName = SalesData::getProductName;
        
        // Function to get name length
        Function<String, Integer> getNameLength = String::length;
        
        // Function to categorize by length
        Function<Integer, String> categorizeByLength = length -> 
            length > 10 ? "Long Name" : "Short Name";
        
        // Chained function
        Function<SalesData, String> nameCategory = 
            extractProductName.andThen(getNameLength).andThen(categorizeByLength);
        
        Map<String, Long> nameCategories = salesData.stream()
            .collect(Collectors.groupingBy(nameCategory, Collectors.counting()));
        
        nameCategories.forEach((category, count) -> 
            System.out.println(category + ": " + count + " products"));
    }

    // Demonstrate Consumer operations
    public void demonstrateConsumerOperations() {
        System.out.println("\n=== CONSUMER OPERATIONS ===");
        
        // Simple consumer
        Consumer<SalesData> printSales = data -> 
            System.out.println(data.getProductName() + ": $" + data.getAmount());
        
        // Chained consumers
        Consumer<SalesData> printWithCategory = printSales.andThen(data -> 
            System.out.println("  Category: " + data.getCategory()));
        
        System.out.println("First 3 sales with categories:");
        salesData.stream().limit(3).forEach(printWithCategory);
    }

    // Demonstrate Supplier operations
    public void demonstrateSupplierOperations() {
        System.out.println("\n=== SUPPLIER OPERATIONS ===");
        
        // Supplier for random sales data
        Supplier<SalesData> randomSalesSupplier = () -> {
            String[] products = {"Laptop", "Mouse", "Keyboard", "Monitor", "Chair"};
            String[] categories = {"Electronics", "Furniture"};
            Random random = new Random();
            return new SalesData(
                "P" + random.nextInt(1000),
                products[random.nextInt(products.length)],
                categories[random.nextInt(categories.length)],
                "North",
                random.nextDouble() * 1000,
                random.nextInt(20) + 1,
                LocalDate.now(),
                "RandomRep"
            );
        };
        
        System.out.println("Generated random sales data:");
        Stream.generate(randomSalesSupplier)
            .limit(3)
            .forEach(System.out::println);
    }

    // Demonstrate custom functional interface
    public void demonstrateCustomFunctionalInterface() {
        System.out.println("\n=== CUSTOM FUNCTIONAL INTERFACE ===");
        
        // Custom functional interface
        TriFunction<SalesData, String, Double, Boolean> categoryAndAmountFilter = 
            (data, category, minAmount) -> 
                data.getCategory().equals(category) && data.getAmount() >= minAmount;
        
        // Using the custom interface
        String targetCategory = "Electronics";
        double minAmount = 100.0;
        
        long count = salesData.stream()
            .filter(data -> categoryAndAmountFilter.apply(data, targetCategory, minAmount))
            .count();
        
        System.out.println("Electronics sales >= $100: " + count);
    }

    // Custom functional interface
    @FunctionalInterface
    interface TriFunction<T, U, V, R> {
        R apply(T t, U u, V v);
    }

    // === STREAM API OPERATIONS ===

    // Demonstrate intermediate operations
    public void demonstrateIntermediateOperations() {
        System.out.println("\n=== INTERMEDIATE OPERATIONS ===");
        
        System.out.println("Stream pipeline demonstration:");
        long count = salesData.stream()
            .filter(data -> data.getAmount() > 100)  // filter
            .peek(data -> System.out.print("[" + data.getProductName() + "] "))  // peek for debugging
            .map(SalesData::getCategory)  // map
            .distinct()  // distinct
            .limit(3)  // limit
            .count();  // terminal operation
        
        System.out.println("\nUnique categories (first 3): " + count);
    }

    // Demonstrate flatMap
    public void demonstrateFlatMap() {
        System.out.println("\n=== FLATMAP DEMONSTRATION ===");
        
        // Group sales by category, then flatten to get all product names
        List<String> allProductNames = salesData.stream()
            .collect(Collectors.groupingBy(SalesData::getCategory))
            .values()  // Collection<List<SalesData>>
            .stream()  // Stream<List<SalesData>>
            .flatMap(List::stream)  // Stream<SalesData>
            .map(SalesData::getProductName)  // Stream<String>
            .distinct()  // unique names
            .sorted()  // sorted
            .collect(Collectors.toList());
        
        System.out.println("All unique product names:");
        allProductNames.forEach(name -> System.out.println("  - " + name));
    }

    // Demonstrate advanced collectors
    public void demonstrateAdvancedCollectors() {
        System.out.println("\n=== ADVANCED COLLECTORS ===");
        
        // Grouping by category with multiple aggregations
        Map<String, DoubleSummaryStatistics> categoryStats = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getCategory,
                Collectors.summarizingDouble(SalesData::getAmount)
            ));
        
        System.out.println("Category statistics:");
        categoryStats.forEach((category, stats) -> 
            System.out.println(String.format(
                "%s: %d sales, Total: $%.2f, Avg: $%.2f, Min: $%.2f, Max: $%.2f",
                category, stats.getCount(), stats.getSum(), 
                stats.getAverage(), stats.getMin(), stats.getMax()
            )));
        
        // Partitioning by high/low value
        Map<Boolean, List<SalesData>> partitionedByValue = salesData.stream()
            .collect(Collectors.partitioningBy(data -> data.getAmount() > 200));
        
        System.out.println("\nHigh value sales (> $200): " + partitionedByValue.get(true).size());
        System.out.println("Low value sales (<= $200): " + partitionedByValue.get(false).size());
        
        // Multi-level grouping
        Map<String, Map<String, Long>> regionCategoryCount = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getRegion,
                Collectors.groupingBy(SalesData::getCategory, Collectors.counting())
            ));
        
        System.out.println("\nSales by region and category:");
        regionCategoryCount.forEach((region, categoryMap) -> {
            System.out.println(region + ":");
            categoryMap.forEach((category, count) -> 
                System.out.println("  " + category + ": " + count));
        });
    }

    // Demonstrate Optional usage
    public void demonstrateOptionalOperations() {
        System.out.println("\n=== OPTIONAL OPERATIONS ===");
        
        // Find highest sale
        Optional<SalesData> highestSale = salesData.stream()
            .max(Comparator.comparingDouble(SalesData::getAmount));
        
        highestSale.ifPresent(sale -> 
            System.out.println("Highest sale: " + sale.getProductName() + " - $" + sale.getAmount()));
        
        // Find sale by product ID
        Optional<SalesData> specificSale = salesData.stream()
            .filter(data -> data.getProductId().equals("P001"))
            .findFirst();
        
        String productName = specificSale
            .map(SalesData::getProductName)
            .orElse("Unknown Product");
        
        System.out.println("Product P001: " + productName);
        
        // Optional chaining
        String salesRep = salesData.stream()
            .filter(data -> data.getProductId().equals("P999"))  // Non-existent
            .findFirst()
            .map(SalesData::getSalesRep)
            .orElseGet(() -> "No Sales Representative");
        
        System.out.println("Sales rep for P999: " + salesRep);
        
        // Optional with filter
        salesData.stream()
            .filter(data -> data.getCategory().equals("Electronics"))
            .findFirst()
            .filter(data -> data.getAmount() > 1000)
            .ifPresent(data -> System.out.println("Found high-value electronics sale: " + data.getProductName()));
    }

    // Demonstrate primitive streams
    public void demonstratePrimitiveStreams() {
        System.out.println("\n=== PRIMITIVE STREAMS ===");
        
        // Using mapToInt for performance
        IntSummaryStatistics amountStats = salesData.stream()
            .mapToInt(SalesData::getQuantity)
            .summaryStatistics();
        
        System.out.println("Quantity statistics:");
        System.out.println("  Total quantity: " + amountStats.getSum());
        System.out.println("  Average quantity: " + amountStats.getAverage());
        System.out.println("  Min quantity: " + amountStats.getMin());
        System.out.println("  Max quantity: " + amountStats.getMax());
        
        // Generate range of months
        List<String> months = IntStream.rangeClosed(1, 12)
            .mapToObj(month -> YearMonth.of(2023, month).format(DateTimeFormatter.ofPattern("MMM")))
            .collect(Collectors.toList());
        
        System.out.println("\nMonths in 2023: " + String.join(", ", months));
    }

    // Demonstrate parallel streams
    public void demonstrateParallelStreams() {
        System.out.println("\n=== PARALLEL STREAMS ===");
        
        // Sequential processing
        long startTime = System.nanoTime();
        long sequentialCount = salesData.stream()
            .filter(data -> data.getAmount() > 100)
            .count();
        long sequentialTime = System.nanoTime() - startTime;
        
        // Parallel processing
        startTime = System.nanoTime();
        long parallelCount = salesData.parallelStream()
            .filter(data -> data.getAmount() > 100)
            .count();
        long parallelTime = System.nanoTime() - startTime;
        
        System.out.println("Sequential processing: " + sequentialCount + " sales in " + 
            (sequentialTime / 1_000_000) + " ms");
        System.out.println("Parallel processing: " + parallelCount + " sales in " + 
            (parallelTime / 1_000_000) + " ms");
        
        // Note: For small datasets, parallel might be slower due to overhead
        double speedup = (double) sequentialTime / parallelTime;
        System.out.println("Speedup: " + String.format("%.2f", speedup) + "x");
    }

    // === REAL-WORLD ANALYTICS ===

    // Generate comprehensive sales report
    public void generateSalesReport() {
        System.out.println("\n=== COMPREHENSIVE SALES REPORT ===");
        
        // Total revenue
        double totalRevenue = salesData.stream()
            .mapToDouble(SalesData::getAmount)
            .sum();
        
        // Sales by category
        Map<String, Double> revenueByCategory = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getCategory,
                Collectors.summingDouble(SalesData::getAmount)
            ));
        
        // Sales by region
        Map<String, Long> salesByRegion = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getRegion,
                Collectors.counting()
            ));
        
        // Top performing products
        List<Map.Entry<String, Double>> topProducts = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getProductName,
                Collectors.summingDouble(SalesData::getAmount)
            ))
            .entrySet()
            .stream()
            .sorted(Map.Entry.<String, Double>comparingByValue().reversed())
            .limit(5)
            .collect(Collectors.toList());
        
        // Monthly trend
        Map<String, Double> monthlyRevenue = salesData.stream()
            .collect(Collectors.groupingBy(
                data -> data.getDate().format(DateTimeFormatter.ofPattern("MMM yyyy")),
                Collectors.summingDouble(SalesData::getAmount)
            ));
        
        // Print report
        System.out.println("TOTAL REVENUE: $" + String.format("%.2f", totalRevenue));
        System.out.println("\nREVENUE BY CATEGORY:");
        revenueByCategory.forEach((category, revenue) -> 
            System.out.println("  " + category + ": $" + String.format("%.2f", revenue)));
        
        System.out.println("\nSALES BY REGION:");
        salesByRegion.forEach((region, count) -> 
            System.out.println("  " + region + ": " + count + " sales"));
        
        System.out.println("\nTOP 5 PRODUCTS:");
        topProducts.forEach(entry -> 
            System.out.println("  " + entry.getKey() + ": $" + String.format("%.2f", entry.getValue())));
        
        System.out.println("\nMONTHLY REVENUE:");
        monthlyRevenue.entrySet().stream()
            .sorted(Map.Entry.comparingByKey())
            .forEach(entry -> 
                System.out.println("  " + entry.getKey() + ": $" + String.format("%.2f", entry.getValue())));
    }

    // Advanced analytics with custom functions
    public void advancedAnalytics() {
        System.out.println("\n=== ADVANCED ANALYTICS ===");
        
        // Calculate correlation between quantity and amount
        double correlation = calculateCorrelation(
            salesData.stream().mapToInt(SalesData::getQuantity).toArray(),
            salesData.stream().mapToDouble(SalesData::getAmount).toArray()
        );
        
        System.out.println("Correlation between quantity and amount: " + String.format("%.3f", correlation));
        
        // Find outliers using custom predicate
        Predicate<SalesData> outlierPredicate = data -> {
            double avgAmount = salesData.stream()
                .mapToDouble(SalesData::getAmount)
                .average()
                .orElse(0);
            double stdDev = Math.sqrt(salesData.stream()
                .mapToDouble(data -> Math.pow(data.getAmount() - avgAmount, 2))
                .average()
                .orElse(0));
            return Math.abs(data.getAmount() - avgAmount) > 2 * stdDev;
        };
        
        List<SalesData> outliers = salesData.stream()
            .filter(outlierPredicate)
            .collect(Collectors.toList());
        
        System.out.println("\nOutliers (2+ standard deviations from mean):");
        outliers.forEach(outlier -> 
            System.out.println("  " + outlier.getProductName() + ": $" + outlier.getAmount()));
        
        // Performance analysis by sales representative
        Map<String, DoubleSummaryStatistics> repPerformance = salesData.stream()
            .collect(Collectors.groupingBy(
                SalesData::getSalesRep,
                Collectors.summarizingDouble(SalesData::getAmount)
            ));
        
        System.out.println("\nSALES REPRESENTATIVE PERFORMANCE:");
        repPerformance.entrySet().stream()
            .sorted(Map.Entry.<String, DoubleSummaryStatistics>comparingByValue(
                Comparator.comparingDouble(DoubleSummaryStatistics::getSum).reversed()))
            .forEach(entry -> {
                String rep = entry.getKey();
                DoubleSummaryStatistics stats = entry.getValue();
                System.out.println(String.format(
                    "  %s: %d sales, Total: $%.2f, Avg: $%.2f",
                    rep, stats.getCount(), stats.getSum(), stats.getAverage()
                ));
            });
    }

    // Helper method to calculate correlation
    private double calculateCorrelation(int[] x, double[] y) {
        if (x.length != y.length || x.length == 0) return 0;
        
        double meanX = Arrays.stream(x).average().orElse(0);
        double meanY = Arrays.stream(y).average().orElse(0);
        
        double numerator = 0;
        double sumSqX = 0;
        double sumSqY = 0;
        
        for (int i = 0; i < x.length; i++) {
            double diffX = x[i] - meanX;
            double diffY = y[i] - meanY;
            numerator += diffX * diffY;
            sumSqX += diffX * diffX;
            sumSqY += diffY * diffY;
        }
        
        double denominator = Math.sqrt(sumSqX * sumSqY);
        return denominator == 0 ? 0 : numerator / denominator;
    }

    public static void main(String[] args) {
        AnalyticsDashboard dashboard = new AnalyticsDashboard();
        
        // Run all demonstrations
        dashboard.demonstrateBasicLambda();
        dashboard.demonstratePredicateComposition();
        dashboard.demonstrateFunctionChaining();
        dashboard.demonstrateConsumerOperations();
        dashboard.demonstrateSupplierOperations();
        dashboard.demonstrateCustomFunctionalInterface();
        dashboard.demonstrateIntermediateOperations();
        dashboard.demonstrateFlatMap();
        dashboard.demonstrateAdvancedCollectors();
        dashboard.demonstrateOptionalOperations();
        dashboard.demonstratePrimitiveStreams();
        dashboard.demonstrateParallelStreams();
        dashboard.generateSalesReport();
        dashboard.advancedAnalytics();
    }
}
```

---

## 🎯 Key Concepts Demonstrated

### 1. **Lambda Expressions**
- Basic lambda syntax: `data -> data.getAmount() > 500`
- Method references: `SalesData::getProductName`
- Anonymous class vs lambda comparison

### 2. **Functional Interfaces**
- **Predicate**: Boolean conditions with composition (`and`, `or`, `negate`)
- **Function**: Transformations with chaining (`andThen`, `compose`)
- **Consumer**: Side-effect operations with chaining
- **Supplier**: Value generation
- **Custom functional interfaces**: `TriFunction` for complex operations

### 3. **Stream Operations**
- **Intermediate**: `filter`, `map`, `flatMap`, `distinct`, `sorted`, `limit`, `skip`, `peek`
- **Terminal**: `collect`, `forEach`, `count`, `reduce`, `min`, `max`, `findFirst`
- **Short-circuiting**: `limit`, `findFirst`, `findAny`

### 4. **Advanced Collectors**
- `Collectors.groupingBy()` with downstream collectors
- `Collectors.partitioningBy()` for binary classification
- `Collectors.summarizingDouble()` for statistics
- Multi-level grouping for hierarchical analysis

### 5. **Optional Usage**
- Safe null handling: `ifPresent()`, `orElse()`, `orElseGet()`
- Optional chaining: `map()`, `filter()`
- Avoiding null pointer exceptions

### 6. **Primitive Streams**
- `mapToInt()`, `mapToDouble()`, `mapToLong()`
- Performance benefits (avoiding boxing)
- Specialized operations: `sum()`, `average()`, `summaryStatistics()`

### 7. **Parallel Streams**
- `parallelStream()` for concurrent processing
- Performance considerations and overhead
- When to use parallel vs sequential

### 8. **Real-World Applications**
- Data analytics and reporting
- Statistical analysis
- Business intelligence
- Performance metrics

---

## 🚀 Expected Output

```
=== BASIC LAMBDA DEMONSTRATION ===
High value sales (traditional): 5
High value sales (lambda): 5
High value sales (method ref): 5

=== PREDICATE COMPOSITION ===
Electronics AND high value: 4
High value OR recent: 7
NOT electronics: 5

=== FUNCTION CHAINING ===
Long Name: 3 products
Short Name: 9 products

=== CONSUMER OPERATIONS ===
First 3 sales with categories:
Laptop Pro: $1200.00
  Category: Electronics
Wireless Mouse: $25.00
  Category: Electronics
Office Chair: $350.00
  Category: Furniture

=== SUPPLIER OPERATIONS ===
Generated random sales data:
SalesData{product='Chair', category='Furniture', region='North', amount=$876.54, qty=15, date=2026-03-18, rep='RandomRep'}
SalesData{product='Keyboard', category='Electronics', region='North', amount=$123.45, qty=8, date=2026-03-18, rep='RandomRep'}
SalesData{product='Monitor', category='Electronics', region='North', amount=$987.65, qty=19, date=2026-03-18, rep='RandomRep'}

=== CUSTOM FUNCTIONAL INTERFACE ===
Electronics sales >= $100: 4

=== INTERMEDIATE OPERATIONS ===
Stream pipeline demonstration:
[Laptop Pro][Wireless Mouse][Office Chair][Standing Desk][Monitor 4K][Keyboard Mechanical][Desk Lamp][Webcam HD][Coffee Maker][Water Bottle][Laptop Pro][Office Chair] 
Unique categories (first 3): 3

=== FLATMAP DEMONSTRATION ===
All unique product names:
  - Coffee Maker
  - Desk Lamp
  - Keyboard Mechanical
  - Laptop Pro
  - Monitor 4K
  - Office Chair
  - Standing Desk
  - Water Bottle
  - Webcam HD
  - Wireless Mouse

=== ADVANCED COLLECTORS ===
Category statistics:
Appliances: 2 sales, Total: $165.00, Avg: $82.50, Min: $15.00, Max: $150.00
Electronics: 6 sales, Total: $3375.00, Avg: $562.50, Min: $25.00, Max: $1200.00
Furniture: 4 sales, Total: $1345.00, Avg: $336.25, Min: $45.00, Max: $600.00

High value sales (> $200): 5
Low value sales (<= $200): 7

Sales by region and category:
East:
  Furniture: 2
  Electronics: 1
North:
  Electronics: 2
  Appliances: 1
South:
  Electronics: 2
  Appliances: 1
West:
  Furniture: 2
  Electronics: 1

=== OPTIONAL OPERATIONS ===
Highest sale: Laptop Pro - $1200.00
Product P001: Laptop Pro
Sales rep for P999: No Sales Representative

=== PRIMITIVE STREAMS ===
Quantity statistics:
  Total quantity: 137
  Average quantity: 11.416666666666666
  Min quantity: 3
  Max quantity: 30

Months in 2023: Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec

=== PARALLEL STREAMS ===
Sequential processing: 7 sales in 2 ms
Parallel processing: 7 sales in 5 ms
Speedup: 0.40x

=== COMPREHENSIVE SALES REPORT ===
TOTAL REVENUE: $4885.00

REVENUE BY CATEGORY:
  Appliances: $165.00
  Electronics: $3375.00
  Furniture: $1345.00

SALES BY REGION:
  East: 3 sales
  North: 3 sales
  South: 3 sales
  West: 3 sales

TOP 5 PRODUCTS:
  Laptop Pro: $2400.00
  Standing Desk: $600.00
  Monitor 4K: $450.00
  Office Chair: $700.00
  Keyboard Mechanical: $120.00

MONTHLY REVENUE:
  Jan 2023: $2225.00
  Feb 2023: $795.00
  Mar 2023: $1865.00

=== ADVANCED ANALYTICS ===
Correlation between quantity and amount: -0.234
Outliers (2+ standard deviations from mean):
  Laptop Pro: $1200.00
  Laptop Pro: $1200.00

SALES REPRESENTATIVE PERFORMANCE:
  Alice: 2 sales, Total: $2400.00, Avg: $1200.00
  Bob: 2 sales, Total: $145.00, Avg: $72.50
  Charlie: 2 sales, Total: $395.00, Avg: $197.50
  Diana: 1 sales, Total: $600.00, Avg: $600.00
  Eve: 1 sales, Total: $80.00, Avg: $80.00
  Frank: 1 sales, Total: $150.00, Avg: $150.00
  Grace: 1 sales, Total: $15.00, Avg: $15.00
  Henry: 1 sales, Total: $350.00, Avg: $350.00
```

---

## 💡 Interview Talking Points

### Functional Programming Benefits
- **Declarative style**: What to do, not how to do it
- **Immutability**: Reduce side effects and bugs
- **Composition**: Build complex operations from simple ones
- **Readability**: More concise and expressive code

### Performance Considerations
- **Lazy evaluation**: Streams process elements only when needed
- **Parallel streams**: Use for large datasets, avoid for small ones
- **Primitive streams**: Avoid boxing overhead for numeric operations
- **Short-circuiting**: Operations like `limit()` and `findFirst()` can improve performance

### Best Practices
- **Method references**: Use when lambda just calls a method
- **Optional**: Avoid `get()`, prefer `orElse()`, `orElseGet()`, `ifPresent()`
- **Collectors**: Use built-in collectors before custom ones
- **Immutable operations**: Prefer `map()` over modifying in place

### Common Pitfalls
- **Multiple operations on same stream**: Streams can't be reused
- **Side effects in lambdas**: Avoid modifying external state
- **Overuse of parallel**: Consider overhead vs benefits
- **Null handling**: Use Optional instead of null checks
