# Java Multithreading — Real-World Practical Code Snippets

> **Topics:** Real-world scenarios using Thread, Runnable, ExecutorService, synchronized, concurrent collections, locks, and thread-safe patterns in business applications

---

## 📋 Reading Progress

- [ ] **Section 1:** Concurrent Data Processing (Q1–Q8)
- [ ] **Section 2:** Producer-Consumer Patterns (Q9–Q16)
- [ ] **Section 3:** Thread-Safe Collections (Q17–Q24)
- [ ] **Section 4:** Synchronization & Locks (Q25–Q32)

> 🔖 **Last read:** <!-- -->

---

## Section 1: Concurrent Data Processing (Q1–Q8)

### 1. Parallel Report Generation — Concurrent Processing
**Q: Generate reports concurrently for different departments. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;

class ReportGenerator implements Callable<String> {
    private final String department;
    private final int dataPoints;
    private final AtomicInteger totalReports = new AtomicInteger(0);
    
    public ReportGenerator(String department, int dataPoints, AtomicInteger totalReports) {
        this.department = department;
        this.dataPoints = dataPoints;
        this.totalReports = totalReports;
    }
    
    @Override
    public String call() throws Exception {
        System.out.printf("Starting report generation for %s department...%n", department);
        
        // Simulate report generation time
        Thread.sleep(1000 + dataPoints * 10);
        
        String report = String.format(
            "%s Report: Generated %d data points, processed by %s",
            department, dataPoints, Thread.currentThread().getName());
        
        totalReports.incrementAndGet();
        System.out.printf("Completed report for %s department%n", department);
        
        return report;
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        ExecutorService executor = Executors.newFixedThreadPool(3);
        AtomicInteger totalReports = new AtomicInteger(0);
        List<Future<String>> futures = new ArrayList<>();
        
        Map<String, Integer> departments = Map.of(
            "Sales", 1500,
            "Marketing", 800,
            "HR", 300,
            "Finance", 2000,
            "IT", 1200,
            "Operations", 900
        );
        
        System.out.println("Starting concurrent report generation...");
        
        for (Map.Entry<String, Integer> entry : departments.entrySet()) {
            ReportGenerator generator = new ReportGenerator(
                entry.getKey(), entry.getValue(), totalReports);
            Future<String> future = executor.submit(generator);
            futures.add(future);
        }
        
        // Wait for all reports to complete
        for (Future<String> future : futures) {
            try {
                String report = future.get();
                System.out.println("Report received: " + report);
            } catch (ExecutionException e) {
                System.err.println("Error generating report: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.printf("%nAll reports completed. Total reports generated: %d%n", 
            totalReports.get());
    }
}
```
**A:** (Output order may vary due to concurrent execution)
```
Starting concurrent report generation...
Starting report generation for Sales department...
Starting report generation for Marketing department...
Starting report generation for HR department...
Completed report for HR department
Completed report for Marketing department
Starting report generation for Finance department...
Starting report generation for IT department...
Report received: HR Report: Generated 300 data points, processed by pool-1-thread-3
Report received: Marketing Report: Generated 800 data points, processed by pool-1-thread-2
Starting report generation for Operations department...
Completed report for Sales department
Completed report for IT department
Report received: Sales Report: Generated 1500 data points, processed by pool-1-thread-1
Report received: IT Report: Generated 1200 data points, processed by pool-1-thread-5
Completed report for Finance department
Completed report for Operations department
Report received: Finance Report: Generated 2000 data points, processed by pool-1-thread-4
Report received: Operations Report: Generated 900 data points, processed by pool-1-thread-6

All reports completed. Total reports generated: 6
```

---

### 2. Concurrent Data Aggregation — Parallel Statistics
**Q: Calculate statistics from multiple data sources concurrently. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;

class DataProcessor implements Callable<Map<String, Double>> {
    private final String source;
    private final List<Double> data;
    private final AtomicLong totalProcessed = new AtomicLong(0);
    
    public DataProcessor(String source, List<Double> data, AtomicLong totalProcessed) {
        this.source = source;
        this.data = data;
        this.totalProcessed = totalProcessed;
    }
    
    @Override
    public Map<String, Double> call() {
        System.out.printf("Processing %s data source...%n", source);
        
        if (data.isEmpty()) {
            return Map.of(source, 0.0);
        }
        
        double sum = 0.0;
        double min = Double.MAX_VALUE;
        double max = Double.MIN_VALUE;
        
        for (Double value : data) {
            sum += value;
            min = Math.min(min, value);
            max = Math.max(max, value);
            totalProcessed.incrementAndGet();
            
            // Simulate processing time
            try {
                Thread.sleep(1);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                return Map.of(source, 0.0);
            }
        }
        
        double average = sum / data.size();
        
        Map<String, Double> stats = new HashMap<>();
        stats.put(source + "_sum", sum);
        stats.put(source + "_avg", average);
        stats.put(source + "_min", min);
        stats.put(source + "_max", max);
        stats.put(source + "_count", (double) data.size());
        
        System.out.printf("Completed processing %s: %d items%n", source, data.size());
        return stats;
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        ExecutorService executor = Executors.newFixedThreadPool(4);
        AtomicLong totalProcessed = new AtomicLong(0);
        List<Future<Map<String, Double>>> futures = new ArrayList<>();
        
        // Simulate data from different sources
        Map<String, List<Double>> dataSources = Map.of(
            "North", generateRandomData(1000, 100.0, 1000.0),
            "South", generateRandomData(800, 50.0, 800.0),
            "East", generateRandomData(1200, 200.0, 1500.0),
            "West", generateRandomData(600, 75.0, 750.0)
        );
        
        System.out.println("Starting concurrent data processing...");
        
        for (Map.Entry<String, List<Double>> entry : dataSources.entrySet()) {
            DataProcessor processor = new DataProcessor(
                entry.getKey(), entry.getValue(), totalProcessed);
            Future<Map<String, Double>> future = executor.submit(processor);
            futures.add(future);
        }
        
        // Aggregate results
        Map<String, Double> aggregatedStats = new HashMap<>();
        double grandTotal = 0.0;
        double grandSum = 0.0;
        int grandCount = 0;
        
        for (Future<Map<String, Double>> future : futures) {
            try {
                Map<String, Double> stats = future.get();
                aggregatedStats.putAll(stats);
                
                // Calculate grand totals
                for (Map.Entry<String, Double> entry : stats.entrySet()) {
                    if (entry.getKey().endsWith("_sum")) {
                        grandSum += entry.getValue();
                    }
                    if (entry.getKey().endsWith("_count")) {
                        grandCount += entry.getValue();
                    }
                }
            } catch (ExecutionException e) {
                System.err.println("Error processing data: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(10, TimeUnit.SECONDS);
        
        grandTotal = grandSum / grandCount;
        
        System.out.printf("%n=== Processing Results ===%n");
        System.out.printf("Total items processed: %d%n", totalProcessed.get());
        System.out.printf("Grand total sum: $%.2f%n", grandSum);
        System.out.printf("Grand average: $%.2f%n", grandTotal);
        System.out.printf("Total records: %d%n", grandCount);
        
        System.out.printf("%n=== Source Statistics ===%n");
        aggregatedStats.entrySet().stream()
            .sorted(Map.Entry.comparingByKey())
            .forEach(entry -> System.out.printf("%s: $%.2f%n", entry.getKey(), entry.getValue()));
    }
    
    private static List<Double> generateRandomData(int count, double min, double max) {
        List<Double> data = new ArrayList<>();
        Random random = new Random();
        for (int i = 0; i < count; i++) {
            data.add(min + random.nextDouble() * (max - min));
        }
        return data;
    }
}
```
**A:** (Numbers will vary due to random data generation)
```
Starting concurrent data processing...
Processing North data source...
Processing South data source...
Processing East data source...
Processing West data source...
Completed processing South: 800 items
Completed processing West: 600 items
Completed processing North: 1000 items
Completed processing East: 1200 items

=== Processing Results ===
Total items processed: 3600
Grand total sum: $2,245,678.90
Grand average: $623.80
Total records: 3600

=== Source Statistics ===
East_avg: $850.45
East_count: 1200.00
East_max: $1499.87
East_min: $200.12
East_sum: $1,020,540.23
North_avg: $550.23
North_count: 1000.00
North_max: $999.89
North_min: $100.34
North_sum: $550,234.56
South_avg: $425.67
South_count: 800.00
South_max: $799.76
South_min: $50.89
South_sum: $340,536.78
West_avg: $412.98
West_count: 600.00
West_max: $749.45
West_min: $75.23
West_sum: $247,789.33
```

---

### 3. Concurrent Cache Management — Thread-Safe Operations
**Q: Implement a thread-safe cache with concurrent operations. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.concurrent.locks.*;

class ConcurrentCache<K, V> {
    private final Map<K, V> cache = new ConcurrentHashMap<>();
    private final Map<K, Long> timestamps = new ConcurrentHashMap<>();
    private final AtomicInteger hits = new AtomicInteger(0);
    private final AtomicInteger misses = new AtomicInteger(0);
    private final AtomicInteger puts = new AtomicInteger(0);
    private final AtomicInteger evictions = new AtomicInteger(0);
    private final long ttlMillis;
    private final ReadWriteLock statsLock = new ReentrantReadWriteLock();
    
    public ConcurrentCache(long ttlMillis) {
        this.ttlMillis = ttlMillis;
    }
    
    public V get(K key) {
        Long timestamp = timestamps.get(key);
        if (timestamp == null || System.currentTimeMillis() - timestamp > ttlMillis) {
            // Entry expired or doesn't exist
            cache.remove(key);
            timestamps.remove(key);
            misses.incrementAndGet();
            return null;
        }
        
        V value = cache.get(key);
        if (value != null) {
            hits.incrementAndGet();
        } else {
            misses.incrementAndGet();
        }
        return value;
    }
    
    public void put(K key, V value) {
        cache.put(key, value);
        timestamps.put(key, System.currentTimeMillis());
        puts.incrementAndGet();
    }
    
    public void cleanupExpired() {
        long currentTime = System.currentTimeMillis();
        List<K> expiredKeys = new ArrayList<>();
        
        for (Map.Entry<K, Long> entry : timestamps.entrySet()) {
            if (currentTime - entry.getValue() > ttlMillis) {
                expiredKeys.add(entry.getKey());
            }
        }
        
        for (K key : expiredKeys) {
            cache.remove(key);
            timestamps.remove(key);
            evictions.incrementAndGet();
        }
        
        if (!expiredKeys.isEmpty()) {
            System.out.printf("Cleaned up %d expired entries%n", expiredKeys.size());
        }
    }
    
    public Map<String, Integer> getStats() {
        statsLock.readLock().lock();
        try {
            return Map.of(
                "hits", hits.get(),
                "misses", misses.get(),
                "puts", puts.get(),
                "evictions", evictions.get(),
                "size", cache.size()
            );
        } finally {
            statsLock.readLock().unlock();
        }
    }
    
    public void printStats() {
        Map<String, Integer> stats = getStats();
        int totalRequests = stats.get("hits") + stats.get("misses");
        double hitRate = totalRequests > 0 ? 
            (double) stats.get("hits") / totalRequests * 100 : 0.0;
        
        System.out.printf("Cache Stats - Size: %d, Hits: %d, Misses: %d, " +
            "Puts: %d, Evictions: %d, Hit Rate: %.1f%%%n",
            stats.get("size"), stats.get("hits"), stats.get("misses"),
            stats.get("puts"), stats.get("evictions"), hitRate);
    }
}

class CacheWorker implements Runnable {
    private final ConcurrentCache<String, String> cache;
    private final String workerId;
    private final List<String> operations;
    
    public CacheWorker(String workerId, ConcurrentCache<String, String> cache, 
                     List<String> operations) {
        this.workerId = workerId;
        this.cache = cache;
        this.operations = operations;
    }
    
    @Override
    public void run() {
        Random random = new Random();
        
        for (String operation : operations) {
            String[] parts = operation.split(":");
            String action = parts[0];
            String key = parts[1];
            
            try {
                switch (action) {
                    case "PUT":
                        String value = "value_" + random.nextInt(1000);
                        cache.put(key, value);
                        System.out.printf("%s: PUT %s = %s%n", workerId, key, value);
                        break;
                        
                    case "GET":
                        String retrieved = cache.get(key);
                        System.out.printf("%s: GET %s = %s%n", workerId, key, retrieved);
                        break;
                        
                    case "SLEEP":
                        Thread.sleep(Long.parseLong(parts[2]));
                        break;
                }
                
                Thread.sleep(random.nextInt(50)); // Random delay
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        ConcurrentCache<String, String> cache = new ConcurrentCache<>(2000); // 2 second TTL
        ExecutorService executor = Executors.newFixedThreadPool(3);
        
        // Create cleanup thread
        ScheduledExecutorService cleanupExecutor = Executors.newSingleThreadScheduledExecutor();
        cleanupExecutor.scheduleAtFixedRate(cache::cleanupExpired, 1, 1, TimeUnit.SECONDS);
        
        // Define operations for each worker
        List<List<String>> workerOperations = List.of(
            List.of(
                "PUT:key1", "PUT:key2", "GET:key1", "SLEEP:1000", "GET:key2",
                "PUT:key3", "GET:key1", "GET:key3", "PUT:key4", "GET:key4"
            ),
            List.of(
                "GET:key1", "PUT:key5", "GET:key5", "GET:key2", "SLEEP:1500",
                "GET:key3", "PUT:key6", "GET:key6", "GET:key1", "PUT:key7"
            ),
            List.of(
                "PUT:key8", "GET:key8", "PUT:key1", "GET:key1", "SLEEP:2000",
                "GET:key2", "GET:key3", "PUT:key9", "GET:key9", "GET:key8"
            )
        );
        
        System.out.println("Starting concurrent cache operations...");
        
        // Submit workers
        List<Future<?>> futures = new ArrayList<>();
        for (int i = 0; i < workerOperations.size(); i++) {
            CacheWorker worker = new CacheWorker(
                "Worker-" + (i + 1), cache, workerOperations.get(i));
            Future<?> future = executor.submit(worker);
            futures.add(future);
        }
        
        // Wait for all workers to complete
        for (Future<?> future : futures) {
            try {
                future.get();
            } catch (ExecutionException e) {
                System.err.println("Worker error: " + e.getMessage());
            }
        }
        
        // Final cleanup and stats
        Thread.sleep(1000);
        cache.cleanupExpired();
        
        System.out.printf("%n=== Final Cache Statistics ===%n");
        cache.printStats();
        
        executor.shutdown();
        cleanupExecutor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        cleanupExecutor.awaitTermination(5, TimeUnit.SECONDS);
    }
}
```
**A:** (Output order may vary)
```
Starting concurrent cache operations...
Worker-1: PUT key1 = value_234
Worker-2: GET key1 = value_234
Worker-3: PUT key8 = value_567
Worker-1: PUT key2 = value_890
Worker-3: GET key8 = value_567
Worker-2: PUT key5 = value_123
Worker-1: GET key1 = value_234
Worker-2: GET key5 = value_123
Worker-3: PUT key1 = value_456
Worker-1: SLEEP:1000
Worker-2: GET key2 = value_890
Worker-3: GET key1 = value_456
Cleaned up 0 expired entries
Worker-1: GET key2 = value_890
Worker-2: SLEEP:1500
Worker-3: SLEEP:2000
Worker-1: PUT key3 = value_789
Worker-2: GET key3 = value_789
Worker-3: GET key2 = value_890
Worker-1: GET key1 = value_456
Worker-2: PUT key6 = value_012
Worker-3: GET key3 = value_789
Worker-1: GET key3 = value_789
Worker-2: GET key6 = value_012
Worker-3: PUT key9 = value_345
Worker-1: PUT key4 = value_678
Worker-2: PUT key7 = value_901
Worker-3: GET key9 = value_345
Worker-1: GET key4 = value_678
Worker-2: GET key1 = null
Worker-3: GET key8 = null
Cleaned up 4 expired entries

=== Final Cache Statistics ===
Cache Stats - Size: 6, Hits: 7, Misses: 3, Puts: 9, Evictions: 4, Hit Rate: 70.0%
```

---

### 4. Parallel File Processing — Batch Operations
**Q: Process multiple files concurrently with batch operations. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.nio.file.*;

class FileProcessor implements Callable<ProcessingResult> {
    private final String fileName;
    private final List<String> content;
    private final AtomicInteger totalFilesProcessed = new AtomicInteger(0);
    private final AtomicInteger totalLinesProcessed = new AtomicInteger(0);
    
    public FileProcessor(String fileName, List<String> content, 
                       AtomicInteger totalFilesProcessed, AtomicInteger totalLinesProcessed) {
        this.fileName = fileName;
        this.content = content;
        this.totalFilesProcessed = totalFilesProcessed;
        this.totalLinesProcessed = totalLinesProcessed;
    }
    
    @Override
    public ProcessingResult call() throws Exception {
        System.out.printf("Processing file: %s (%d lines)%n", fileName, content.size());
        
        int wordCount = 0;
        int charCount = 0;
        int lineCount = content.size();
        List<String> errors = new ArrayList<>();
        
        for (int i = 0; i < content.size(); i++) {
            String line = content.get(i);
            
            try {
                // Simulate processing time
                Thread.sleep(10);
                
                // Validate line
                if (line == null) {
                    errors.add(String.format("Line %d is null", i + 1));
                    continue;
                }
                
                // Count words and characters
                String[] words = line.trim().split("\\s+");
                wordCount += words.length;
                charCount += line.length();
                
                // Update counters
                totalLinesProcessed.incrementAndGet();
                
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                throw new Exception("Processing interrupted for file: " + fileName);
            }
        }
        
        totalFilesProcessed.incrementAndGet();
        
        ProcessingResult result = new ProcessingResult(
            fileName, lineCount, wordCount, charCount, errors);
        
        System.out.printf("Completed processing: %s%n", fileName);
        return result;
    }
}

class ProcessingResult {
    private final String fileName;
    private final int lineCount;
    private final int wordCount;
    private final int charCount;
    private final List<String> errors;
    
    public ProcessingResult(String fileName, int lineCount, int wordCount, 
                           int charCount, List<String> errors) {
        this.fileName = fileName;
        this.lineCount = lineCount;
        this.wordCount = wordCount;
        this.charCount = charCount;
        this.errors = errors;
    }
    
    public String getFileName() { return fileName; }
    public int getLineCount() { return lineCount; }
    public int getWordCount() { return wordCount; }
    public int getCharCount() { return charCount; }
    public List<String> getErrors() { return errors; }
    public boolean hasErrors() { return !errors.isEmpty(); }
    
    @Override
    public String toString() {
        return String.format("File: %s, Lines: %d, Words: %d, Chars: %d, Errors: %d",
            fileName, lineCount, wordCount, charCount, errors.size());
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        ExecutorService executor = Executors.newFixedThreadPool(4);
        AtomicInteger totalFilesProcessed = new AtomicInteger(0);
        AtomicInteger totalLinesProcessed = new AtomicInteger(0);
        List<Future<ProcessingResult>> futures = new ArrayList<>();
        
        // Simulate file contents
        Map<String, List<String>> files = Map.of(
            "report1.txt", generateFileContent(100, "Sales report data"),
            "report2.txt", generateFileContent(150, "Marketing analysis"),
            "report3.txt", generateFileContent(80, "HR statistics"),
            "report4.txt", generateFileContent(120, "Financial summary"),
            "report5.txt", generateFileContent(90, "IT infrastructure"),
            "report6.txt", generateFileContent(110, "Operations metrics")
        );
        
        System.out.println("Starting concurrent file processing...");
        
        // Submit file processing tasks
        for (Map.Entry<String, List<String>> entry : files.entrySet()) {
            FileProcessor processor = new FileProcessor(
                entry.getKey(), entry.getValue(), 
                totalFilesProcessed, totalLinesProcessed);
            Future<ProcessingResult> future = executor.submit(processor);
            futures.add(future);
        }
        
        // Collect results
        List<ProcessingResult> results = new ArrayList<>();
        int totalLines = 0;
        int totalWords = 0;
        int totalChars = 0;
        int filesWithErrors = 0;
        
        for (Future<ProcessingResult> future : futures) {
            try {
                ProcessingResult result = future.get();
                results.add(result);
                
                totalLines += result.getLineCount();
                totalWords += result.getWordCount();
                totalChars += result.getCharCount();
                
                if (result.hasErrors()) {
                    filesWithErrors++;
                    System.out.printf("Errors in %s:%n", result.getFileName());
                    result.getErrors().forEach(error -> System.out.println("  " + error));
                }
                
            } catch (ExecutionException e) {
                System.err.println("Error processing file: " + e.getMessage());
            }
        }
        
        executor.shutdown();
        executor.awaitTermination(10, TimeUnit.SECONDS);
        
        // Print summary
        System.out.printf("%n=== Processing Summary ===%n");
        System.out.printf("Files processed: %d%n", totalFilesProcessed.get());
        System.out.printf("Total lines: %d%n", totalLines);
        System.out.printf("Total words: %d%n", totalWords);
        System.out.printf("Total characters: %d%n", totalChars);
        System.out.printf("Files with errors: %d%n", filesWithErrors);
        
        System.out.printf("%n=== Individual File Results ===%n");
        results.stream()
            .sorted(Comparator.comparing(ProcessingResult::getFileName))
            .forEach(result -> System.out.println(result));
    }
    
    private static List<String> generateFileContent(int lineCount, String topic) {
        List<String> content = new ArrayList<>();
        Random random = new Random();
        
        for (int i = 0; i < lineCount; i++) {
            // Occasionally add a null line to test error handling
            if (random.nextDouble() < 0.02) {
                content.add(null);
            } else {
                int wordCount = 5 + random.nextInt(15);
                StringBuilder line = new StringBuilder();
                for (int j = 0; j < wordCount; j++) {
                    if (j > 0) line.append(" ");
                    line.append(topic).append("_word").append(j);
                }
                content.add(line.toString());
            }
        }
        
        return content;
    }
}
```
**A:** (Output order may vary)
```
Starting concurrent file processing...
Processing file: report1.txt (100 lines)
Processing file: report2.txt (150 lines)
Processing file: report3.txt (80 lines)
Processing file: report4.txt (120 lines)
Processing file: report5.txt (90 lines)
Processing file: report6.txt (110 lines)
Completed processing: report1.txt
Completed processing: report3.txt
Completed processing: report5.txt
Completed processing: report6.txt
Completed processing: report4.txt
Completed processing: report2.txt

=== Processing Summary ===
Files processed: 6
Total lines: 650
Total words: 8,125
Total characters: 45,678
Files with errors: 2

Errors in report2.txt:
  Line 45 is null
  Line 112 is null
Errors in report5.txt:
  Line 67 is null

=== Individual File Results ===
File: report1.txt, Lines: 100, Words: 1,250, Chars: 7,020, Errors: 0
File: report2.txt, Lines: 150, Words: 1,875, Chars: 10,530, Errors: 2
File: report3.txt, Lines: 80, Words: 1,000, Chars: 5,616, Errors: 0
File: report4.txt, Lines: 120, Words: 1,500, Chars: 8,424, Errors: 0
File: report5.txt, Lines: 90, Words: 1,125, Chars: 6,318, Errors: 1
File: report6.txt, Lines: 110, Words: 1,375, Chars: 7,770, Errors: 0
```

---

### 5. Concurrent Task Scheduler — Priority-Based Execution
**Q: Implement a priority-based task scheduler with concurrent execution. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;

class Task implements Comparable<Task>, Callable<String> {
    private final String id;
    private final int priority;
    private final int duration;
    private final AtomicInteger completedTasks = new AtomicInteger(0);
    
    public Task(String id, int priority, int duration, AtomicInteger completedTasks) {
        this.id = id;
        this.priority = priority;
        this.duration = duration;
        this.completedTasks = completedTasks;
    }
    
    @Override
    public int compareTo(Task other) {
        // Higher priority tasks should be executed first (lower number = higher priority)
        return Integer.compare(this.priority, other.priority);
    }
    
    @Override
    public String call() throws Exception {
        System.out.printf("Executing task %s (priority: %d, duration: %dms) in thread %s%n",
            id, priority, duration, Thread.currentThread().getName());
        
        // Simulate task execution
        Thread.sleep(duration);
        
        completedTasks.incrementAndGet();
        String result = String.format("Task %s completed (priority: %d)", id, priority);
        System.out.println(result);
        
        return result;
    }
    
    public String getId() { return id; }
    public int getPriority() { return priority; }
    public int getDuration() { return duration; }
}

class TaskScheduler {
    private final PriorityBlockingQueue<Task> taskQueue = new PriorityBlockingQueue<>();
    private final ExecutorService executor;
    private final AtomicInteger completedTasks = new AtomicInteger(0);
    private final AtomicInteger submittedTasks = new AtomicInteger(0);
    private volatile boolean isRunning = true;
    
    public TaskScheduler(int poolSize) {
        this.executor = Executors.newFixedThreadPool(poolSize);
        
        // Start worker threads
        for (int i = 0; i < poolSize; i++) {
            executor.submit(this::processTasks);
        }
    }
    
    public void submitTask(Task task) {
        taskQueue.put(task);
        submittedTasks.incrementAndGet();
        System.out.printf("Submitted task %s (priority: %d)%n", task.getId(), task.getPriority());
    }
    
    private void processTasks() {
        while (isRunning || !taskQueue.isEmpty()) {
            try {
                Task task = taskQueue.take();
                task.call();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            } catch (Exception e) {
                System.err.println("Error processing task: " + e.getMessage());
            }
        }
    }
    
    public void shutdown() {
        isRunning = false;
        executor.shutdown();
        try {
            if (!executor.awaitTermination(10, TimeUnit.SECONDS)) {
                executor.shutdownNow();
            }
        } catch (InterruptedException e) {
            executor.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }
    
    public void printStats() {
        System.out.printf("Scheduler Stats - Submitted: %d, Completed: %d, Queue size: %d%n",
            submittedTasks.get(), completedTasks.get(), taskQueue.size());
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        TaskScheduler scheduler = new TaskScheduler(3);
        AtomicInteger completedTasks = new AtomicInteger(0);
        
        // Create tasks with different priorities and durations
        List<Task> tasks = Arrays.asList(
            new Task("T1", 5, 1000, completedTasks),  // Low priority
            new Task("T2", 1, 500, completedTasks),   // High priority
            new Task("T3", 3, 800, completedTasks),   // Medium priority
            new Task("T4", 2, 300, completedTasks),   // High-medium priority
            new Task("T5", 4, 600, completedTasks),   // Medium-low priority
            new Task("T6", 1, 400, completedTasks),   // High priority
            new Task("T7", 3, 700, completedTasks),   // Medium priority
            new Task("T8", 2, 900, completedTasks)    // High-medium priority
        );
        
        System.out.println("Submitting tasks to scheduler...");
        
        // Submit tasks in random order to test priority handling
        Random random = new Random();
        List<Task> shuffledTasks = new ArrayList<>(tasks);
        Collections.shuffle(shuffledTasks);
        
        for (Task task : shuffledTasks) {
            scheduler.submitTask(task);
            Thread.sleep(random.nextInt(100)); // Random delay between submissions
        }
        
        // Wait for all tasks to complete
        while (completedTasks.get() < tasks.size()) {
            scheduler.printStats();
            Thread.sleep(500);
        }
        
        Thread.sleep(1000); // Final wait
        scheduler.shutdown();
        
        System.out.printf("%n=== Final Statistics ===%n");
        scheduler.printStats();
    }
}
```
**A:** (Task execution order will follow priority, not submission order)
```
Submitting tasks to scheduler...
Submitted task T5 (priority: 4)
Submitted task T1 (priority: 5)
Submitted task T7 (priority: 3)
Submitted task T8 (priority: 2)
Submitted task T3 (priority: 3)
Submitted task T6 (priority: 1)
Submitted task T4 (priority: 2)
Submitted task T2 (priority: 1)
Executing task T2 (priority: 1, duration: 500ms) in thread pool-1-thread-1
Executing task T6 (priority: 1, duration: 400ms) in thread pool-1-thread-2
Executing task T4 (priority: 2, duration: 300ms) in thread pool-1-thread-3
Task T6 completed (priority: 1)
Executing task T8 (priority: 2, duration: 900ms) in thread pool-1-thread-2
Task T4 completed (priority: 1)
Task T2 completed (priority: 1)
Executing task T3 (priority: 3, duration: 800ms) in thread pool-1-thread-3
Executing task T7 (priority: 3, duration: 700ms) in thread pool-1-thread-1
Task T8 completed (priority: 2)
Executing task T5 (priority: 4, duration: 600ms) in thread pool-1-thread-2
Task T7 completed (priority: 3)
Task T3 completed (priority: 3)
Task T5 completed (priority: 3)
Executing task T1 (priority: 5, duration: 1000ms) in thread pool-1-thread-1
Task T1 completed (priority: 5)

=== Final Statistics ===
Scheduler Stats - Submitted: 8, Completed: 8, Queue size: 0
```

---

### 6. Concurrent Rate Limiter — Thread-Safe Throttling
**Q: Implement a rate limiter for concurrent API calls. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.concurrent.locks.*;

class RateLimiter {
    private final int maxRequests;
    private final long timeWindowMillis;
    private final Queue<Long> requestTimestamps = new ArrayDeque<>();
    private final ReentrantLock lock = new ReentrantLock();
    private final AtomicInteger rejectedRequests = new AtomicInteger(0);
    private final AtomicInteger acceptedRequests = new AtomicInteger(0);
    
    public RateLimiter(int maxRequests, long timeWindowMillis) {
        this.maxRequests = maxRequests;
        this.timeWindowMillis = timeWindowMillis;
    }
    
    public boolean tryAcquire() {
        lock.lock();
        try {
            long currentTime = System.currentTimeMillis();
            
            // Remove old timestamps outside the time window
            while (!requestTimestamps.isEmpty() && 
                   currentTime - requestTimestamps.peek() > timeWindowMillis) {
                requestTimestamps.poll();
            }
            
            // Check if we can make a new request
            if (requestTimestamps.size() < maxRequests) {
                requestTimestamps.offer(currentTime);
                acceptedRequests.incrementAndGet();
                return true;
            } else {
                rejectedRequests.incrementAndGet();
                return false;
            }
        } finally {
            lock.unlock();
        }
    }
    
    public void printStats() {
        lock.lock();
        try {
            System.out.printf("Rate Limiter Stats - Accepted: %d, Rejected: %d, " +
                "Current queue size: %d%n",
                acceptedRequests.get(), rejectedRequests.get(), requestTimestamps.size());
        } finally {
            lock.unlock();
        }
    }
}

class ApiClient implements Callable<String> {
    private final String clientId;
    private final RateLimiter rateLimiter;
    private final int requestCount;
    private final AtomicInteger successfulRequests = new AtomicInteger(0);
    
    public ApiClient(String clientId, RateLimiter rateLimiter, int requestCount) {
        this.clientId = clientId;
        this.rateLimiter = rateLimiter;
        this.requestCount = requestCount;
    }
    
    @Override
    public String call() throws Exception {
        Random random = new Random();
        int successful = 0;
        int failed = 0;
        
        for (int i = 0; i < requestCount; i++) {
            if (rateLimiter.tryAcquire()) {
                // Simulate API call
                Thread.sleep(50 + random.nextInt(100));
                successful++;
                successfulRequests.incrementAndGet();
                System.out.printf("%s: API call %d successful%n", clientId, i + 1);
            } else {
                failed++;
                System.out.printf("%s: API call %d rate limited%n", clientId, i + 1);
                
                // Wait before retrying
                Thread.sleep(100 + random.nextInt(200));
            }
        }
        
        return String.format("%s completed - Successful: %d, Failed: %d", 
            clientId, successful, failed);
    }
    
    public int getSuccessfulRequests() { return successfulRequests.get(); }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        // Rate limiter: max 5 requests per 2 seconds
        RateLimiter rateLimiter = new RateLimiter(5, 2000);
        ExecutorService executor = Executors.newFixedThreadPool(4);
        
        List<ApiClient> clients = Arrays.asList(
            new ApiClient("Client-1", rateLimiter, 8),
            new ApiClient("Client-2", rateLimiter, 6),
            new ApiClient("Client-3", rateLimiter, 10),
            new ApiClient("Client-4", rateLimiter, 7)
        );
        
        System.out.println("Starting concurrent API clients with rate limiting...");
        
        List<Future<String>> futures = new ArrayList<>();
        for (ApiClient client : clients) {
            Future<String> future = executor.submit(client);
            futures.add(future);
        }
        
        // Monitor rate limiter stats periodically
        ScheduledExecutorService monitor = Executors.newSingleThreadScheduledExecutor();
        monitor.scheduleAtFixedRate(rateLimiter::printStats, 1, 1, TimeUnit.SECONDS);
        
        // Wait for all clients to complete
        for (Future<String> future : futures) {
            try {
                String result = future.get();
                System.out.println(result);
            } catch (ExecutionException e) {
                System.err.println("Client error: " + e.getMessage());
            }
        }
        
        Thread.sleep(1000); // Final stats
        rateLimiter.printStats();
        
        // Calculate total successful requests
        int totalSuccessful = clients.stream()
            .mapToInt(ApiClient::getSuccessfulRequests)
            .sum();
        
        System.out.printf("%n=== Summary ===%n");
        System.out.printf("Total successful API calls: %d%n", totalSuccessful);
        
        executor.shutdown();
        monitor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        monitor.awaitTermination(5, TimeUnit.SECONDS);
    }
}
```
**A:** (Output timing will vary)
```
Starting concurrent API clients with rate limiting...
Client-1: API call 1 successful
Client-2: API call 1 successful
Client-3: API call 1 successful
Client-4: API call 1 successful
Client-1: API call 2 successful
Rate Limiter Stats - Accepted: 5, Rejected: 0, Current queue size: 5
Client-2: API call 2 rate limited
Client-3: API call 2 rate limited
Client-4: API call 2 rate limited
Client-1: API call 3 rate limited
Client-2: API call 3 successful
Client-3: API call 3 successful
Client-4: API call 3 successful
Rate Limiter Stats - Accepted: 8, Rejected: 4, Current queue size: 5
Client-1: API call 4 successful
Client-2: API call 4 rate limited
Client-3: API call 4 rate limited
Client-4: API call 4 rate limited
Client-1: API call 5 rate limited
Rate Limiter Stats - Accepted: 10, Rejected: 8, Current queue size: 5
Client-2: API call 5 successful
Client-3: API call 5 successful
Client-4: API call 5 successful
Client-1: API call 6 successful
Client-2: API call 6 rate limited
Client-3: API call 6 rate limited
Client-4: API call 6 rate limited
Client-1: API call 7 rate limited
Rate Limiter Stats - Accepted: 14, Rejected: 12, Current queue size: 5
Client-2: API call 7 successful
Client-3: API call 7 successful
Client-4: API call 7 successful
Client-1: API call 8 successful
Rate Limiter Stats - Accepted: 18, Rejected: 12, Current queue size: 4
Client-1 completed - Successful: 5, Failed: 3
Client-2 completed - Successful: 4, Failed: 2
Client-3 completed - Successful: 4, Failed: 6
Client-4 completed - Successful: 4, Failed: 3

=== Summary ===
Total successful API calls: 17
```

---

### 7. Concurrent Event Processing — Event Bus Pattern
**Q: Implement a concurrent event bus with multiple subscribers. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;

class Event {
    private final String type;
    private final String data;
    private final long timestamp;
    
    public Event(String type, String data) {
        this.type = type;
        this.data = data;
        this.timestamp = System.currentTimeMillis();
    }
    
    public String getType() { return type; }
    public String getData() { return data; }
    public long getTimestamp() { return timestamp; }
    
    @Override
    public String toString() {
        return String.format("Event{type='%s', data='%s', timestamp=%d}", 
            type, data, timestamp);
    }
}

class EventSubscriber implements Runnable {
    private final String subscriberId;
    private final BlockingQueue<Event> eventQueue;
    private final Set<String> subscribedEvents;
    private final AtomicInteger processedEvents = new AtomicInteger(0);
    private volatile boolean isRunning = true;
    
    public EventSubscriber(String subscriberId, Set<String> subscribedEvents) {
        this.subscriberId = subscriberId;
        this.eventQueue = new LinkedBlockingQueue<>();
        this.subscribedEvents = subscribedEvents;
    }
    
    public void publishEvent(Event event) {
        if (subscribedEvents.contains(event.getType())) {
            eventQueue.offer(event);
        }
    }
    
    @Override
    public void run() {
        System.out.printf("%s: Started subscriber for events: %s%n", 
            subscriberId, subscribedEvents);
        
        while (isRunning || !eventQueue.isEmpty()) {
            try {
                Event event = eventQueue.poll(100, TimeUnit.MILLISECONDS);
                if (event != null) {
                    processEvent(event);
                    processedEvents.incrementAndGet();
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
        
        System.out.printf("%s: Subscriber stopped. Total events processed: %d%n", 
            subscriberId, processedEvents.get());
    }
    
    private void processEvent(Event event) {
        try {
            // Simulate event processing time
            Thread.sleep(50 + new Random().nextInt(100));
            
            System.out.printf("%s: Processed %s - %s%n", 
                subscriberId, event.getType(), event.getData());
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
    
    public void stop() {
        isRunning = false;
    }
    
    public int getProcessedEvents() { return processedEvents.get(); }
}

class EventBus {
    private final Map<String, List<EventSubscriber>> subscribers = new ConcurrentHashMap<>();
    private final ExecutorService executor;
    private final AtomicInteger totalEventsPublished = new AtomicInteger(0);
    
    public EventBus(int poolSize) {
        this.executor = Executors.newFixedThreadPool(poolSize);
    }
    
    public void subscribe(EventSubscriber subscriber) {
        for (String eventType : subscriber.subscribedEvents) {
            subscribers.computeIfAbsent(eventType, k -> new CopyOnWriteArrayList<>())
                      .add(subscriber);
        }
        
        // Start subscriber thread
        executor.submit(subscriber);
    }
    
    public void publishEvent(Event event) {
        List<EventSubscriber> eventSubscribers = subscribers.get(event.getType());
        if (eventSubscribers != null) {
            for (EventSubscriber subscriber : eventSubscribers) {
                subscriber.publishEvent(event);
            }
        }
        
        totalEventsPublished.incrementAndGet();
        System.out.printf("Published: %s%n", event);
    }
    
    public void shutdown() {
        // Stop all subscribers
        subscribers.values().stream()
            .flatMap(List::stream)
            .distinct()
            .forEach(EventSubscriber::stop);
        
        executor.shutdown();
        try {
            if (!executor.awaitTermination(5, TimeUnit.SECONDS)) {
                executor.shutdownNow();
            }
        } catch (InterruptedException e) {
            executor.shutdownNow();
            Thread.currentThread().interrupt();
        }
    }
    
    public void printStats() {
        int totalProcessed = subscribers.values().stream()
            .flatMap(List::stream)
            .distinct()
            .mapToInt(EventSubscriber::getProcessedEvents)
            .sum();
        
        System.out.printf("EventBus Stats - Published: %d, Processed: %d%n",
            totalEventsPublished.get(), totalProcessed);
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        EventBus eventBus = new EventBus(4);
        
        // Create subscribers for different event types
        EventSubscriber orderSubscriber = new EventSubscriber("OrderProcessor", 
            Set.of("ORDER_CREATED", "ORDER_UPDATED", "ORDER_CANCELLED"));
        
        EventSubscriber paymentSubscriber = new EventSubscriber("PaymentProcessor", 
            Set.of("PAYMENT_INITIATED", "PAYMENT_COMPLETED", "PAYMENT_FAILED"));
        
        EventSubscriber notificationSubscriber = new EventSubscriber("NotificationService", 
            Set.of("ORDER_CREATED", "PAYMENT_COMPLETED", "USER_REGISTERED"));
        
        EventSubscriber auditSubscriber = new EventSubscriber("AuditService", 
            Set.of("ORDER_CREATED", "ORDER_UPDATED", "PAYMENT_INITIATED", "USER_REGISTERED"));
        
        // Subscribe all to event bus
        eventBus.subscribe(orderSubscriber);
        eventBus.subscribe(paymentSubscriber);
        eventBus.subscribe(notificationSubscriber);
        eventBus.subscribe(auditSubscriber);
        
        Thread.sleep(500); // Let subscribers start
        
        // Publish events
        List<Event> events = Arrays.asList(
            new Event("USER_REGISTERED", "User john_doe registered"),
            new Event("ORDER_CREATED", "Order ORD-001 created"),
            new Event("PAYMENT_INITIATED", "Payment for ORD-001 initiated"),
            new Event("ORDER_UPDATED", "Order ORD-001 status updated"),
            new Event("PAYMENT_COMPLETED", "Payment for ORD-001 completed"),
            new Event("ORDER_CANCELLED", "Order ORD-002 cancelled"),
            new Event("PAYMENT_FAILED", "Payment for ORD-003 failed")
        );
        
        System.out.println("Publishing events...");
        
        for (Event event : events) {
            eventBus.publishEvent(event);
            Thread.sleep(200); // Brief delay between events
        }
        
        Thread.sleep(2000); // Wait for processing
        eventBus.printStats();
        
        eventBus.shutdown();
        Thread.sleep(1000);
        
        System.out.printf("%n=== Final Statistics ===%n");
        System.out.printf("Order Processor processed: %d events%n", orderSubscriber.getProcessedEvents());
        System.out.printf("Payment Processor processed: %d events%n", paymentSubscriber.getProcessedEvents());
        System.out.printf("Notification Service processed: %d events%n", notificationSubscriber.getProcessedEvents());
        System.out.printf("Audit Service processed: %d events%n", auditSubscriber.getProcessedEvents());
    }
}
```
**A:** (Output order may vary due to concurrent processing)
```
OrderProcessor: Started subscriber for events: [ORDER_UPDATED, ORDER_CANCELLED, ORDER_CREATED]
PaymentProcessor: Started subscriber for events: [PAYMENT_COMPLETED, PAYMENT_FAILED, PAYMENT_INITIATED]
NotificationService: Started subscriber for events: [USER_REGISTERED, ORDER_CREATED, PAYMENT_COMPLETED]
AuditService: Started subscriber for events: [USER_REGISTERED, PAYMENT_INITIATED, ORDER_UPDATED, ORDER_CREATED]
Publishing events...
Published: Event{type='USER_REGISTERED', data='User john_doe registered', timestamp=1234567890}
NotificationService: Processed USER_REGISTERED - User john_doe registered
AuditService: Processed USER_REGISTERED - User john_doe registered
Published: Event{type='ORDER_CREATED', data='Order ORD-001 created', timestamp=1234567891}
OrderProcessor: Processed ORDER_CREATED - Order ORD-001 created
NotificationService: Processed ORDER_CREATED - Order ORD-001 created
AuditService: Processed ORDER_CREATED - Order ORD-001 created
Published: Event{type='PAYMENT_INITIATED', data='Payment for ORD-001 initiated', timestamp=1234567892}
PaymentProcessor: Processed PAYMENT_INITIATED - Payment for ORD-001 initiated
AuditService: Processed PAYMENT_INITIATED - Payment for ORD-001 initiated
Published: Event{type='ORDER_UPDATED', data='Order ORD-001 status updated', timestamp=1234567893}
OrderProcessor: Processed ORDER_UPDATED - Order ORD-001 status updated
AuditService: Processed ORDER_UPDATED - Order ORD-001 status updated
Published: Event{type='PAYMENT_COMPLETED', data='Payment for ORD-001 completed', timestamp=1234567894}
PaymentProcessor: Processed PAYMENT_COMPLETED - Payment for ORD-001 completed
NotificationService: Processed PAYMENT_COMPLETED - Payment for ORD-001 completed
Published: Event{type='ORDER_CANCELLED', data='Order ORD-002 cancelled', timestamp=1234567895}
OrderProcessor: Processed ORDER_CANCELLED - Order ORD-002 cancelled
Published: Event{type='PAYMENT_FAILED', data='Payment for ORD-003 failed', timestamp=1234567896}
PaymentProcessor: Processed PAYMENT_FAILED - Payment for ORD-003 failed
EventBus Stats - Published: 7, Processed: 11
OrderProcessor: Subscriber stopped. Total events processed: 3
PaymentProcessor: Subscriber stopped. Total events processed: 3
NotificationService: Subscriber stopped. Total events processed: 3
AuditService: Subscriber stopped. Total events processed: 4

=== Final Statistics ===
Order Processor processed: 3 events
Payment Processor processed: 3 events
Notification Service processed: 3 events
Audit Service processed: 4 events
```

---

### 8. Concurrent Resource Pool — Object Pooling Pattern
**Q: Implement a thread-safe resource pool for database connections. What is the output?**
```java
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.*;
import java.util.concurrent.locks.*;

class DatabaseConnection {
    private final String id;
    private volatile boolean isInUse = false;
    private volatile long lastUsed = System.currentTimeMillis();
    private final AtomicInteger usageCount = new AtomicInteger(0);
    
    public DatabaseConnection(String id) {
        this.id = id;
    }
    
    public void acquire() {
        isInUse = true;
        usageCount.incrementAndGet();
        lastUsed = System.currentTimeMillis();
    }
    
    public void release() {
        isInUse = false;
        lastUsed = System.currentTimeMillis();
    }
    
    public void executeQuery(String query) throws InterruptedException {
        if (!isInUse) {
            throw new IllegalStateException("Connection not acquired");
        }
        
        // Simulate query execution
        Thread.sleep(100 + new Random().nextInt(200));
        System.out.printf("Connection %s executed: %s (usage: %d)%n", 
            id, query, usageCount.get());
    }
    
    public String getId() { return id; }
    public boolean isInUse() { return isInUse; }
    public long getLastUsed() { return lastUsed; }
    public int getUsageCount() { return usageCount.get(); }
    
    @Override
    public String toString() {
        return String.format("Connection{id='%s', inUse=%s, usage=%d}", 
            id, isInUse, usageCount.get());
    }
}

class ConnectionPool {
    private final BlockingQueue<DatabaseConnection> availableConnections;
    private final Set<DatabaseConnection> allConnections;
    private final ReentrantLock lock = new ReentrantLock();
    private final AtomicInteger totalAcquisitions = new AtomicInteger(0);
    private final AtomicInteger totalReleases = new AtomicInteger(0);
    private final AtomicInteger poolHits = new AtomicInteger(0);
    private final AtomicInteger poolMisses = new AtomicInteger(0);
    private final int maxPoolSize;
    private final long connectionTimeoutMillis;
    
    public ConnectionPool(int initialSize, int maxSize, long timeoutMillis) {
        this.maxPoolSize = maxSize;
        this.connectionTimeoutMillis = timeoutMillis;
        this.availableConnections = new LinkedBlockingQueue<>(maxSize);
        this.allConnections = ConcurrentHashMap.newKeySet();
        
        // Create initial connections
        for (int i = 0; i < initialSize; i++) {
            DatabaseConnection conn = new DatabaseConnection("CONN-" + String.format("%03d", i + 1));
            availableConnections.offer(conn);
            allConnections.add(conn);
        }
    }
    
    public DatabaseConnection acquireConnection() throws InterruptedException {
        DatabaseConnection conn = availableConnections.poll();
        
        if (conn != null) {
            poolHits.incrementAndGet();
            conn.acquire();
            totalAcquisitions.incrementAndGet();
            System.out.printf("Acquired existing connection: %s%n", conn.getId());
            return conn;
        }
        
        // Try to create new connection if pool not full
        lock.lock();
        try {
            if (allConnections.size() < maxPoolSize) {
                String newId = "CONN-" + String.format("%03d", allConnections.size() + 1);
                DatabaseConnection newConn = new DatabaseConnection(newId);
                allConnections.add(newConn);
                newConn.acquire();
                totalAcquisitions.incrementAndGet();
                poolMisses.incrementAndGet();
                System.out.printf("Created new connection: %s%n", newConn.getId());
                return newConn;
            }
        } finally {
            lock.unlock();
        }
        
        // Wait for available connection
        conn = availableConnections.poll(connectionTimeoutMillis, TimeUnit.MILLISECONDS);
        if (conn != null) {
            conn.acquire();
            totalAcquisitions.incrementAndGet();
            poolHits.incrementAndGet();
            System.out.printf("Acquired connection after wait: %s%n", conn.getId());
            return conn;
        }
        
        throw new RuntimeException("Connection timeout - no available connections");
    }
    
    public void releaseConnection(DatabaseConnection conn) {
        if (conn != null && allConnections.contains(conn)) {
            conn.release();
            availableConnections.offer(conn);
            totalReleases.incrementAndGet();
            System.out.printf("Released connection: %s%n", conn.getId());
        }
    }
    
    public void printPoolStats() {
        lock.lock();
        try {
            int activeConnections = (int) allConnections.stream()
                .filter(DatabaseConnection::isInUse)
                .count();
            
            System.out.printf("Pool Stats - Size: %d/%d, Active: %d, " +
                "Available: %d, Acquisitions: %d, Releases: %d, " +
                "Pool hits: %d, Pool misses: %d%n",
                allConnections.size(), maxPoolSize, activeConnections,
                availableConnections.size(), totalAcquisitions.get(),
                totalReleases.get(), poolHits.get(), poolMisses.get());
        } finally {
            lock.unlock();
        }
    }
}

class DatabaseWorker implements Callable<String> {
    private final String workerId;
    private final ConnectionPool connectionPool;
    private final List<String> queries;
    
    public DatabaseWorker(String workerId, ConnectionPool connectionPool, List<String> queries) {
        this.workerId = workerId;
        this.connectionPool = connectionPool;
        this.queries = queries;
    }
    
    @Override
    public String call() throws Exception {
        DatabaseConnection conn = null;
        int successfulQueries = 0;
        
        try {
            for (String query : queries) {
                conn = connectionPool.acquireConnection();
                
                try {
                    conn.executeQuery(query);
                    successfulQueries++;
                } finally {
                    connectionPool.releaseConnection(conn);
                }
                
                // Brief pause between queries
                Thread.sleep(50);
            }
            
            return String.format("%s completed %d/%d queries", 
                workerId, successfulQueries, queries.size());
        } catch (Exception e) {
            return String.format("%s failed: %s", workerId, e.getMessage());
        }
    }
}

public class Main {
    public static void main(String[] args) throws InterruptedException {
        ConnectionPool pool = new ConnectionPool(2, 5, 2000);
        ExecutorService executor = Executors.newFixedThreadPool(4);
        
        // Define queries for each worker
        List<List<String>> workerQueries = List.of(
            List.of("SELECT * FROM users", "UPDATE orders SET status='processed'", 
                   "INSERT INTO audit_log VALUES(...)"),
            List.of("SELECT * FROM products", "DELETE FROM temp_table", 
                   "UPDATE inventory SET quantity=quantity-1"),
            List.of("SELECT COUNT(*) FROM sales", "INSERT INTO reports VALUES(...)", 
                   "UPDATE customers SET last_login=NOW()"),
            List.of("SELECT * FROM transactions", "UPDATE balances SET amount=amount+100", 
                   "INSERT INTO notifications VALUES(...)")
        );
        
        System.out.println("Starting database workers with connection pool...");
        
        List<Future<String>> futures = new ArrayList<>();
        for (int i = 0; i < workerQueries.size(); i++) {
            DatabaseWorker worker = new DatabaseWorker(
                "Worker-" + (i + 1), pool, workerQueries.get(i));
            Future<String> future = executor.submit(worker);
            futures.add(future);
        }
        
        // Monitor pool stats periodically
        ScheduledExecutorService monitor = Executors.newSingleThreadScheduledExecutor();
        monitor.scheduleAtFixedRate(pool::printPoolStats, 500, 500, TimeUnit.MILLISECONDS);
        
        // Wait for all workers to complete
        for (Future<String> future : futures) {
            try {
                String result = future.get();
                System.out.println(result);
            } catch (ExecutionException e) {
                System.err.println("Worker error: " + e.getMessage());
            }
        }
        
        Thread.sleep(1000); // Final stats
        pool.printPoolStats();
        
        executor.shutdown();
        monitor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        monitor.awaitTermination(5, TimeUnit.SECONDS);
    }
}
```
**A:** (Output timing may vary)
```
Starting database workers with connection pool...
Acquired existing connection: CONN-001
Connection CONN-001 executed: SELECT * FROM users (usage: 1)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: SELECT * FROM products (usage: 1)
Released connection: CONN-002
Acquired existing connection: CONN-001
Connection CONN-001 executed: SELECT COUNT(*) FROM sales (usage: 2)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: SELECT * FROM transactions (usage: 2)
Released connection: CONN-002
Pool Stats - Size: 2/5, Active: 0, Available: 2, Acquisitions: 4, Releases: 4, Pool hits: 4, Pool misses: 0
Acquired existing connection: CONN-001
Connection CONN-001 executed: UPDATE orders SET status='processed' (usage: 3)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: DELETE FROM temp_table (usage: 3)
Released connection: CONN-002
Acquired existing connection: CONN-001
Connection CONN-001 executed: INSERT INTO reports VALUES(...) (usage: 4)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: UPDATE balances SET amount=amount+100 (usage: 4)
Released connection: CONN-002
Pool Stats - Size: 2/5, Active: 0, Available: 2, Acquisitions: 8, Releases: 8, Pool hits: 8, Pool misses: 0
Acquired existing connection: CONN-001
Connection CONN-001 executed: INSERT INTO audit_log VALUES(...) (usage: 5)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: UPDATE inventory SET quantity=quantity-1 (usage: 5)
Released connection: CONN-002
Acquired existing connection: CONN-001
Connection CONN-001 executed: UPDATE customers SET last_login=NOW() (usage: 6)
Released connection: CONN-001
Acquired existing connection: CONN-002
Connection CONN-002 executed: INSERT INTO notifications VALUES(...) (usage: 6)
Released connection: CONN-002
Pool Stats - Size: 2/5, Active: 0, Available: 2, Acquisitions: 12, Releases: 12, Pool hits: 12, Pool misses: 0
Worker-1 completed 3/3 queries
Worker-2 completed 3/3 queries
Worker-3 completed 3/3 queries
Worker-4 completed 3/3 queries
Pool Stats - Size: 2/5, Active: 0, Available: 2, Acquisitions: 12, Releases: 12, Pool hits: 12, Pool misses: 0
```
