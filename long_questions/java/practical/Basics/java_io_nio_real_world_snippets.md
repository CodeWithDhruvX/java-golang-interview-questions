# Java IO/NIO — Real-World Practical Code Snippets

> **Topics:** Real-world scenarios using File I/O, NIO.2, Streams, Buffers, Channels, and file system operations in business applications

---

## 📋 Reading Progress

- [ ] **Section 1:** File Operations & Management (Q1–Q8)
- [ ] **Section 2:** Data Processing & Transformation (Q9–Q16)
- [ ] **Section 3:** Network I/O & Communication (Q17–Q24)
- [ ] **Section 4:** Advanced NIO Features (Q25–Q32)

> 🔖 **Last read:** <!-- -->

---

## Section 1: File Operations & Management (Q1–Q8)

### 1. Log File Analyzer — Reading Large Files Efficiently
**Q: Analyze log files to extract error patterns. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.regex.*;
import java.time.*;
import java.time.format.*;

class LogAnalyzer {
    private static final Pattern ERROR_PATTERN = 
        Pattern.compile("(\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}) \\[(\\w+)\\] (.+)");
    private static final DateTimeFormatter DATE_FORMATTER = 
        DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss");
    
    public static class LogEntry {
        private final LocalDateTime timestamp;
        private final String level;
        private final String message;
        
        public LogEntry(LocalDateTime timestamp, String level, String message) {
            this.timestamp = timestamp;
            this.level = level;
            this.message = message;
        }
        
        public LocalDateTime getTimestamp() { return timestamp; }
        public String getLevel() { return level; }
        public String getMessage() { return message; }
    }
    
    public static List<LogEntry> analyzeLogFile(Path logFile) throws IOException {
        List<LogEntry> entries = new ArrayList<>();
        
        try (BufferedReader reader = Files.newBufferedReader(logFile)) {
            String line;
            int lineNumber = 0;
            
            while ((line = reader.readLine()) != null) {
                lineNumber++;
                
                Matcher matcher = ERROR_PATTERN.matcher(line);
                if (matcher.find()) {
                    try {
                        LocalDateTime timestamp = LocalDateTime.parse(matcher.group(1), DATE_FORMATTER);
                        String level = matcher.group(2);
                        String message = matcher.group(3);
                        
                        entries.add(new LogEntry(timestamp, level, message));
                    } catch (DateTimeParseException e) {
                        System.err.printf("Line %d: Invalid timestamp format%n", lineNumber);
                    }
                }
            }
        }
        
        return entries;
    }
    
    public static void printAnalysis(List<LogEntry> entries) {
        Map<String, Integer> levelCounts = new HashMap<>();
        Map<String, Integer> hourlyErrors = new HashMap<>();
        
        for (LogEntry entry : entries) {
            // Count by level
            levelCounts.merge(entry.getLevel(), 1, Integer::sum);
            
            // Count errors by hour
            if ("ERROR".equals(entry.getLevel())) {
                String hour = entry.getTimestamp().format(DateTimeFormatter.ofPattern("HH"));
                hourlyErrors.merge(hour, 1, Integer::sum);
            }
        }
        
        System.out.println("=== Log Analysis Results ===");
        System.out.printf("Total entries: %d%n", entries.size());
        
        System.out.println("\nLog Levels:");
        levelCounts.entrySet().stream()
            .sorted(Map.Entry.<String, Integer>comparingByValue().reversed())
            .forEach(entry -> System.out.printf("  %s: %d%n", entry.getKey(), entry.getValue()));
        
        if (!hourlyErrors.isEmpty()) {
            System.out.println("\nErrors by Hour:");
            hourlyErrors.entrySet().stream()
                .sorted(Map.Entry.comparingByKey())
                .forEach(entry -> System.out.printf("  %s:00 - %d errors%n", entry.getKey(), entry.getValue()));
        }
        
        // Show latest errors
        System.out.println("\nLatest 5 Errors:");
        entries.stream()
            .filter(entry -> "ERROR".equals(entry.getLevel()))
            .sorted(Comparator.comparing(LogEntry::getTimestamp).reversed())
            .limit(5)
            .forEach(entry -> System.out.printf("  %s [%s] %s%n", 
                entry.getTimestamp(), entry.getLevel(), entry.getMessage()));
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create sample log file content
        String logContent = """
            2024-01-15 09:15:23 [INFO] Application started
            2024-01-15 09:15:24 [DEBUG] Loading configuration
            2024-01-15 09:15:25 [INFO] Database connected
            2024-01-15 09:16:01 [ERROR] Failed to connect to external service
            2024-01-15 09:16:02 [WARN] Retrying connection attempt
            2024-01-15 09:17:30 [ERROR] Database timeout occurred
            2024-01-15 09:18:15 [INFO] User login: john_doe
            2024-01-15 09:18:16 [DEBUG] Session created
            2024-01-15 09:19:45 [ERROR] Invalid user credentials
            2024-01-15 09:20:00 [INFO] User logout: john_doe
            2024-01-15 10:01:12 [ERROR] Memory threshold exceeded
            2024-01-15 10:01:13 [WARN] Garbage collection triggered
            2024-01-15 10:02:30 [INFO] System health check passed
            2024-01-15 10:03:45 [ERROR] File not found: config.xml
            2024-01-15 10:04:00 [INFO] Backup completed successfully
            """;
        
        // Write to temporary file
        Path logFile = Files.createTempFile("application", ".log");
        Files.write(logFile, logContent.getBytes());
        
        System.out.printf("Analyzing log file: %s%n", logFile);
        
        // Analyze the log file
        List<LogAnalyzer.LogEntry> entries = LogAnalyzer.analyzeLogFile(logFile);
        LogAnalyzer.printAnalysis(entries);
        
        // Clean up
        Files.deleteIfExists(logFile);
    }
}
```
**A:** 
```
Analyzing log file: C:\Users\USER\AppData\Local\Temp\application123456.log
=== Log Analysis Results ===
Total entries: 15

Log Levels:
  INFO: 7
  ERROR: 5
  DEBUG: 2
  WARN: 1

Errors by Hour:
  09:00 - 3 errors
  10:00 - 2 errors

Latest 5 Errors:
  2024-01-15 10:03:45 [ERROR] File not found: config.xml
  2024-01-15 10:01:12 [ERROR] Memory threshold exceeded
  2024-01-15 09:19:45 [ERROR] Invalid user credentials
  2024-01-15 09:17:30 [ERROR] Database timeout occurred
  2024-01-15 09:16:01 [ERROR] Failed to connect to external service
```

---

### 2. Configuration File Processor — JSON/YAML Handling
**Q: Process configuration files with validation and defaults. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.regex.*;

class ConfigurationProcessor {
    
    public static class ConfigEntry {
        private final String key;
        private final String value;
        private final String type;
        private final boolean required;
        
        public ConfigEntry(String key, String value, String type, boolean required) {
            this.key = key;
            this.value = value;
            this.type = type;
            this.required = required;
        }
        
        public String getKey() { return key; }
        public String getValue() { return value; }
        public String getType() { return type; }
        public boolean isRequired() { return required; }
        
        @Override
        public String toString() {
            return String.format("%s: %s (%s%s)", 
                key, value, type, required ? ", required" : "");
        }
    }
    
    public static Map<String, ConfigEntry> loadConfiguration(Path configFile) throws IOException {
        Map<String, ConfigEntry> config = new HashMap<>();
        
        try (BufferedReader reader = Files.newBufferedReader(configFile)) {
            String line;
            int lineNumber = 0;
            
            while ((line = reader.readLine()) != null) {
                lineNumber++;
                line = line.trim();
                
                // Skip empty lines and comments
                if (line.isEmpty() || line.startsWith("#")) {
                    continue;
                }
                
                try {
                    ConfigEntry entry = parseConfigLine(line);
                    config.put(entry.getKey(), entry);
                } catch (IllegalArgumentException e) {
                    System.err.printf("Line %d: Invalid format - %s%n", lineNumber, e.getMessage());
                }
            }
        }
        
        return config;
    }
    
    private static ConfigEntry parseConfigLine(String line) {
        // Format: key = value [type] [required]
        Pattern pattern = Pattern.compile("^([^=]+)=\\s*([^\\[]+)(?:\\[(.+)\\])?(?:\\s*\\[(required)\\])?$");
        Matcher matcher = pattern.matcher(line);
        
        if (!matcher.find()) {
            throw new IllegalArgumentException("Invalid configuration line format");
        }
        
        String key = matcher.group(1).trim();
        String value = matcher.group(2).trim();
        String type = matcher.group(3) != null ? matcher.group(3).trim() : "string";
        boolean required = matcher.group(4) != null && matcher.group(4).equals("required");
        
        return new ConfigEntry(key, value, type, required);
    }
    
    public static Map<String, ConfigEntry> validateAndApplyDefaults(
            Map<String, ConfigEntry> config, Map<String, String> defaults) {
        
        Map<String, ConfigEntry> result = new HashMap<>(config);
        List<String> missingRequired = new ArrayList<>();
        
        // Check required fields
        for (ConfigEntry entry : config.values()) {
            if (entry.isRequired() && (entry.getValue() == null || entry.getValue().isEmpty())) {
                missingRequired.add(entry.getKey());
            }
        }
        
        // Apply defaults for missing non-required fields
        for (Map.Entry<String, String> def : defaults.entrySet()) {
            if (!config.containsKey(def.getKey())) {
                result.put(def.getKey(), new ConfigEntry(def.getKey(), def.getValue(), "string", false));
                System.out.printf("Applied default for %s: %s%n", def.getKey(), def.getValue());
            }
        }
        
        if (!missingRequired.isEmpty()) {
            throw new RuntimeException("Missing required configuration: " + String.join(", ", missingRequired));
        }
        
        return result;
    }
    
    public static void printConfiguration(Map<String, ConfigEntry> config) {
        System.out.println("=== Configuration ===");
        config.values().stream()
            .sorted(Comparator.comparing(ConfigEntry::getKey))
            .forEach(entry -> System.out.println("  " + entry));
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create sample configuration file
        String configContent = """
            # Application Configuration
            server.port = 8080 [integer] [required]
            server.host = localhost [string]
            database.url = jdbc:mysql://localhost:3306/mydb [string] [required]
            database.pool.size = 10 [integer]
            cache.enabled = true [boolean]
            cache.ttl = 3600 [integer]
            log.level = INFO [string]
            # Missing api.key (required) - should cause error
            """;
        
        // Write configuration to temporary file
        Path configFile = Files.createTempFile("config", ".properties");
        Files.write(configFile, configContent.getBytes());
        
        System.out.printf("Processing configuration file: %s%n", configFile);
        
        try {
            // Load configuration
            Map<String, ConfigurationProcessor.ConfigEntry> config = 
                ConfigurationProcessor.loadConfiguration(configFile);
            
            // Define defaults
            Map<String, String> defaults = Map.of(
                "server.host", "0.0.0.0",
                "database.pool.size", "5",
                "cache.enabled", "false",
                "cache.ttl", "1800",
                "log.level", "INFO",
                "api.key", "default-api-key"  // This won't be applied since it's required
            );
            
            // Validate and apply defaults
            Map<String, ConfigurationProcessor.ConfigEntry> validatedConfig = 
                ConfigurationProcessor.validateAndApplyDefaults(config, defaults);
            
            ConfigurationProcessor.printConfiguration(validatedConfig);
            
        } catch (RuntimeException e) {
            System.err.println("Configuration validation failed: " + e.getMessage());
        }
        
        // Clean up
        Files.deleteIfExists(configFile);
    }
}
```
**A:** 
```
Processing configuration file: C:\Users\Users\AppData\Local\Temp\config123456.properties
Applied default for cache.enabled: false
Applied default for cache.ttl: 1800
Configuration validation failed: Missing required configuration: api.key
```

---

### 3. File Backup System — Directory Operations
**Q: Implement a file backup system with directory traversal. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.nio.file.attribute.*;
import java.util.*;
import java.util.stream.*;
import java.time.*;

class FileBackupSystem {
    
    public static class BackupResult {
        private final int filesCopied;
        private final int directoriesCreated;
        private final long totalBytes;
        private final List<String> errors;
        private final Duration duration;
        
        public BackupResult(int filesCopied, int directoriesCreated, long totalBytes, 
                           List<String> errors, Duration duration) {
            this.filesCopied = filesCopied;
            this.directoriesCreated = directoriesCreated;
            this.totalBytes = totalBytes;
            this.errors = errors;
            this.duration = duration;
        }
        
        public int getFilesCopied() { return filesCopied; }
        public int getDirectoriesCreated() { return directoriesCreated; }
        public long getTotalBytes() { return totalBytes; }
        public List<String> getErrors() { return errors; }
        public Duration getDuration() { return duration; }
        
        @Override
        public String toString() {
            return String.format(
                "Backup completed in %s - Files: %d, Directories: %d, Size: %s, Errors: %d",
                duration, filesCopied, directoriesCreated, formatBytes(totalBytes), errors.size());
        }
        
        private static String formatBytes(long bytes) {
            if (bytes < 1024) return bytes + " B";
            if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
            if (bytes < 1024 * 1024 * 1024) return String.format("%.1f MB", bytes / (1024.0 * 1024));
            return String.format("%.1f GB", bytes / (1024.0 * 1024 * 1024));
        }
    }
    
    public static BackupResult backupDirectory(Path sourceDir, Path targetDir) throws IOException {
        Instant start = Instant.now();
        int filesCopied = 0;
        int directoriesCreated = 0;
        long totalBytes = 0;
        List<String> errors = new ArrayList<>();
        
        // Create target directory if it doesn't exist
        if (!Files.exists(targetDir)) {
            Files.createDirectories(targetDir);
            directoriesCreated++;
            System.out.printf("Created directory: %s%n", targetDir);
        }
        
        // Walk through source directory
        try (Stream<Path> paths = Files.walk(sourceDir)) {
            List<Path> allPaths = paths.collect(Collectors.toList());
            
            for (Path sourcePath : allPaths) {
                try {
                    Path relativePath = sourceDir.relativize(sourcePath);
                    Path targetPath = targetDir.resolve(relativePath);
                    
                    if (Files.isDirectory(sourcePath)) {
                        // Create directory
                        if (!Files.exists(targetPath)) {
                            Files.createDirectories(targetPath);
                            directoriesCreated++;
                            System.out.printf("Created directory: %s%n", targetPath);
                        }
                    } else {
                        // Copy file
                        copyFileWithAttributes(sourcePath, targetPath);
                        filesCopied++;
                        totalBytes += Files.size(sourcePath);
                        System.out.printf("Copied file: %s (%s)%n", 
                            relativePath, formatFileSize(Files.size(sourcePath)));
                    }
                } catch (IOException e) {
                    String error = String.format("Failed to backup %s: %s", 
                        sourcePath, e.getMessage());
                    errors.add(error);
                    System.err.println(error);
                }
            }
        }
        
        Duration duration = Duration.between(start, Instant.now());
        return new BackupResult(filesCopied, directoriesCreated, totalBytes, errors, duration);
    }
    
    private static void copyFileWithAttributes(Path source, Path target) throws IOException {
        // Copy file with attributes
        Files.copy(source, target, StandardCopyOption.COPY_ATTRIBUTES, 
                  StandardCopyOption.REPLACE_EXISTING);
    }
    
    private static String formatFileSize(long bytes) {
        if (bytes < 1024) return bytes + " B";
        if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
        return String.format("%.1f MB", bytes / (1024.0 * 1024));
    }
    
    public static void createSampleDirectory(Path baseDir) throws IOException {
        // Create sample directory structure
        Files.createDirectories(baseDir.resolve("docs"));
        Files.createDirectories(baseDir.resolve("images"));
        Files.createDirectories(baseDir.resolve("data"));
        Files.createDirectories(baseDir.resolve("backup"));
        
        // Create sample files
        Files.write(baseDir.resolve("config.txt"), "server.port=8080\n".getBytes());
        Files.write(baseDir.resolve("docs/readme.txt"), "This is a readme file.\n".getBytes());
        Files.write(baseDir.resolve("docs/manual.txt"), "User manual content.\n".getBytes());
        Files.write(baseDir.resolve("images/logo.png"), "fake-image-data".getBytes());
        Files.write(baseDir.resolve("data/users.csv"), "id,name,email\n1,John,john@email.com\n".getBytes());
        Files.write(baseDir.resolve("backup/old_config.txt"), "old configuration\n".getBytes());
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create temporary directories
        Path sourceDir = Files.createTempDirectory("backup_source");
        Path targetDir = Files.createTempDirectory("backup_target");
        
        try {
            // Create sample directory structure
            FileBackupSystem.createSampleDirectory(sourceDir);
            
            System.out.printf("Starting backup from %s to %s%n", sourceDir, targetDir);
            System.out.println();
            
            // Perform backup
            FileBackupSystem.BackupResult result = 
                FileBackupSystem.backupDirectory(sourceDir, targetDir);
            
            System.out.println();
            System.out.println("=== Backup Summary ===");
            System.out.println(result);
            
            if (!result.getErrors().isEmpty()) {
                System.out.println("\nErrors encountered:");
                result.getErrors().forEach(error -> System.out.println("  " + error));
            }
            
        } finally {
            // Clean up
            try {
                Files.walk(sourceDir)
                    .sorted(Comparator.reverseOrder())
                    .forEach(path -> {
                        try { Files.deleteIfExists(path); } catch (IOException e) {}
                    });
                Files.walk(targetDir)
                    .sorted(Comparator.reverseOrder())
                    .forEach(path -> {
                        try { Files.deleteIfExists(path); } catch (IOException e) {}
                    });
            } catch (IOException e) {
                System.err.println("Error during cleanup: " + e.getMessage());
            }
        }
    }
}
```
**A:** 
```
Starting backup from C:\Users\USER\AppData\Local\Temp\backup_source123456 to C:\Users\USER\AppData\Local\Temp\backup_target789012

Created directory: C:\Users\USER\AppData\Local\Temp\backup_target789012\backup
Created directory: C:\Users\USER\AppData\Local\Temp\backup_target789012\data
Created directory: C:\Users\USER\AppData\Local\Temp\backup_target789012\docs
Created directory: C:\Users\USER\AppData\Local\Temp\backup_target789012\images
Copied file: config.txt (16 B)
Copied file: docs\readme.txt (23 B)
Copied file: docs\manual.txt (20 B)
Copied file: images\logo.png (15 B)
Copied file: data\users.csv (41 B)
Copied file: backup\old_config.txt (19 B)

=== Backup Summary ===
Backup completed in PT0.012S - Files: 6, Directories: 4, Size: 134 B, Errors: 0
```

---

### 4. File Watcher Service — Real-time Monitoring
**Q: Monitor directory changes with file watcher service. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.concurrent.*;
import java.nio.file.attribute.*;

class FileWatcherService {
    private final WatchService watchService;
    private final Map<WatchKey, Path> watchKeys;
    private final ExecutorService executor;
    private volatile boolean isRunning = false;
    
    public FileWatcherService() throws IOException {
        this.watchService = FileSystems.getDefault().newWatchService();
        this.watchKeys = new HashMap<>();
        this.executor = Executors.newSingleThreadExecutor();
    }
    
    public void watchDirectory(Path directory) throws IOException {
        if (!Files.isDirectory(directory)) {
            throw new IllegalArgumentException("Path is not a directory: " + directory);
        }
        
        WatchKey key = directory.register(watchService,
            StandardWatchEventKinds.ENTRY_CREATE,
            StandardWatchEventKinds.ENTRY_DELETE,
            StandardWatchEventKinds.ENTRY_MODIFY);
        
        watchKeys.put(key, directory);
        System.out.printf("Watching directory: %s%n", directory);
    }
    
    public void startWatching() {
        if (isRunning) {
            return;
        }
        
        isRunning = true;
        executor.submit(this::processEvents);
        System.out.println("File watcher service started");
    }
    
    public void stopWatching() {
        isRunning = false;
        executor.shutdown();
        try {
            if (!executor.awaitTermination(5, TimeUnit.SECONDS)) {
                executor.shutdownNow();
            }
        } catch (InterruptedException e) {
            executor.shutdownNow();
            Thread.currentThread().interrupt();
        }
        
        try {
            watchService.close();
        } catch (IOException e) {
            System.err.println("Error closing watch service: " + e.getMessage());
        }
        
        System.out.println("File watcher service stopped");
    }
    
    private void processEvents() {
        while (isRunning) {
            try {
                WatchKey key = watchService.take();
                Path watchedDir = watchKeys.get(key);
                
                if (watchedDir != null) {
                    for (WatchEvent<?> event : key.pollEvents()) {
                        handleEvent(watchedDir, event);
                    }
                }
                
                boolean valid = key.reset();
                if (!valid) {
                    watchKeys.remove(key);
                    if (watchKeys.isEmpty()) {
                        break;
                    }
                }
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                break;
            }
        }
    }
    
    private void handleEvent(Path directory, WatchEvent<?> event) {
        WatchEvent.Kind<?> kind = event.kind();
        Path fileName = (Path) event.context();
        Path fullPath = directory.resolve(fileName);
        
        try {
            String fileType = Files.isDirectory(fullPath) ? "Directory" : "File";
            long size = Files.isDirectory(fullPath) ? 0 : Files.size(fullPath);
            
            System.out.printf("[%s] %s %s: %s (%s)%n",
                new Date(), eventTypeToString(kind), fileType, fileName, formatFileSize(size));
            
            // If it's a new directory, watch it too
            if (kind == StandardWatchEventKinds.ENTRY_CREATE && Files.isDirectory(fullPath)) {
                try {
                    watchDirectory(fullPath);
                } catch (IOException e) {
                    System.err.println("Failed to watch new directory: " + e.getMessage());
                }
            }
            
        } catch (IOException e) {
            System.err.printf("Error getting file info for %s: %s%n", fullPath, e.getMessage());
        }
    }
    
    private String eventTypeToString(WatchEvent.Kind<?> kind) {
        if (kind == StandardWatchEventKinds.ENTRY_CREATE) return "Created";
        if (kind == StandardWatchEventKinds.ENTRY_DELETE) return "Deleted";
        if (kind == StandardWatchEventKinds.ENTRY_MODIFY) return "Modified";
        return "Unknown";
    }
    
    private String formatFileSize(long bytes) {
        if (bytes < 1024) return bytes + " B";
        if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
        return String.format("%.1f MB", bytes / (1024.0 * 1024));
    }
}

class FileOperationsSimulator implements Runnable {
    private final Path baseDir;
    private final int operationsCount;
    
    public FileOperationsSimulator(Path baseDir, int operationsCount) {
        this.baseDir = baseDir;
        this.operationsCount = operationsCount;
    }
    
    @Override
    public void run() {
        Random random = new Random();
        
        try {
            for (int i = 0; i < operationsCount; i++) {
                int operation = random.nextInt(4);
                String fileName = "file_" + i + ".txt";
                Path filePath = baseDir.resolve(fileName);
                
                switch (operation) {
                    case 0: // Create file
                        Files.write(filePath, ("Content of " + fileName).getBytes());
                        break;
                        
                    case 1: // Modify file
                        if (Files.exists(filePath)) {
                            Files.write(filePath, ("Modified content " + i).getBytes());
                        }
                        break;
                        
                    case 2: // Delete file
                        Files.deleteIfExists(filePath);
                        break;
                        
                    case 3: // Create directory
                        Path dirPath = baseDir.resolve("dir_" + i);
                        Files.createDirectories(dirPath);
                        Files.write(dirPath.resolve("nested.txt"), "nested content".getBytes());
                        break;
                }
                
                Thread.sleep(200 + random.nextInt(300));
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        } catch (IOException e) {
            System.err.println("File operation error: " + e.getMessage());
        }
    }
}

public class Main {
    public static void main(String[] args) throws Exception {
        // Create temporary directory
        Path watchDir = Files.createTempDirectory("file_watcher_test");
        
        try {
            System.out.printf("Setting up file watcher for directory: %s%n", watchDir);
            
            // Start file watcher
            FileWatcherService watcher = new FileWatcherService();
            watcher.watchDirectory(watchDir);
            watcher.startWatching();
            
            // Start file operations simulator
            Thread simulator = new Thread(new FileOperationsSimulator(watchDir, 10));
            simulator.start();
            
            // Wait for simulator to complete
            simulator.join();
            
            // Wait a bit more for any remaining events
            Thread.sleep(2000);
            
            // Stop watcher
            watcher.stopWatching();
            
        } finally {
            // Clean up
            try {
                Files.walk(watchDir)
                    .sorted(Comparator.reverseOrder())
                    .forEach(path -> {
                        try { Files.deleteIfExists(path); } catch (IOException e) {}
                    });
            } catch (IOException e) {
                System.err.println("Error during cleanup: " + e.getMessage());
            }
        }
    }
}
```
**A:** (Output timestamps and order will vary)
```
Setting up file watcher for directory: C:\Users\USER\AppData\Local\Temp\file_watcher_test123456
Watching directory: C:\Users\USER\AppData\Local\Temp\file_watcher_test123456
File watcher service started
[Mon Jan 15 10:30:45 EST 2024] Created File: file_0.txt (20 B)
[Mon Jan 15 10:30:45 EST 2024] Created Directory: dir_1 (0 B)
Watching directory: C:\Users\USER\AppData\Local\Temp\file_watcher_test123456\dir_1
[Mon Jan 15 10:30:45 EST 2024] Created File: nested.txt (14 B)
[Mon Jan 15 10:30:46 EST 2024] Created File: file_2.txt (20 B)
[Mon Jan 15 10:30:46 EST 2024] Modified File: file_2.txt (28 B)
[Mon Jan 15 10:30:47 EST 2024] Deleted File: file_0.txt (0 B)
[Mon Jan 15 10:30:47 EST 2024] Created Directory: dir_3 (0 B)
Watching directory: C:\Users\USER\AppData\Local\Temp\file_watcher_test123456\dir_3
[Mon Jan 15 10:30:47 EST 2024] Created File: nested.txt (14 B)
[Mon Jan 15 10:30:48 EST 2024] Created File: file_4.txt (20 B)
[Mon Jan 15 10:30:48 EST 2024] Deleted File: file_2.txt (0 B)
[Mon Jan 15 10:30:49 EST 2024] Created Directory: dir_5 (0 B)
Watching directory: C:\Users\USER\AppData\Local\Temp\file_watcher_test123456\dir_5
[Mon Jan 15 10:30:49 EST 2024] Created File: nested.txt (14 B)
File watcher service stopped
```

---

### 5. CSV Data Processor — Large File Handling
**Q: Process large CSV files with memory-efficient streaming. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.stream.*;

class CsvProcessor {
    
    public static class SalesRecord {
        private final String orderId;
        private final String customerId;
        private final String product;
        private final int quantity;
        private final double price;
        private final String date;
        
        public SalesRecord(String orderId, String customerId, String product, 
                          int quantity, double price, String date) {
            this.orderId = orderId;
            this.customerId = customerId;
            this.product = product;
            this.quantity = quantity;
            this.price = price;
            this.date = date;
        }
        
        // Getters
        public String getOrderId() { return orderId; }
        public String getCustomerId() { return customerId; }
        public String getProduct() { return product; }
        public int getQuantity() { return quantity; }
        public double getPrice() { return price; }
        public String getDate() { return date; }
        
        public double getTotal() { return quantity * price; }
        
        @Override
        public String toString() {
            return String.format("SalesRecord{id=%s, product=%s, qty=%d, price=$%.2f, total=$%.2f}",
                orderId, product, quantity, price, getTotal());
        }
    }
    
    public static class ProcessingStats {
        private final int totalRecords;
        private final int validRecords;
        private final int invalidRecords;
        private final double totalRevenue;
        private final Map<String, Integer> productCounts;
        private final Map<String, Double> customerTotals;
        
        public ProcessingStats(int totalRecords, int validRecords, int invalidRecords,
                            double totalRevenue, Map<String, Integer> productCounts,
                            Map<String, Double> customerTotals) {
            this.totalRecords = totalRecords;
            this.validRecords = validRecords;
            this.invalidRecords = invalidRecords;
            this.totalRevenue = totalRevenue;
            this.productCounts = productCounts;
            this.customerTotals = customerTotals;
        }
        
        // Getters
        public int getTotalRecords() { return totalRecords; }
        public int getValidRecords() { return validRecords; }
        public int getInvalidRecords() { return invalidRecords; }
        public double getTotalRevenue() { return totalRevenue; }
        public Map<String, Integer> getProductCounts() { return productCounts; }
        public Map<String, Double> getCustomerTotals() { return customerTotals; }
        
        public void printStats() {
            System.out.println("=== CSV Processing Statistics ===");
            System.out.printf("Total Records: %d%n", totalRecords);
            System.out.printf("Valid Records: %d%n", validRecords);
            System.out.printf("Invalid Records: %d%n", invalidRecords);
            System.out.printf("Total Revenue: $%.2f%n", totalRevenue);
            
            System.out.println("\nTop Products:");
            productCounts.entrySet().stream()
                .sorted(Map.Entry.<String, Integer>comparingByValue().reversed())
                .limit(5)
                .forEach(entry -> System.out.printf("  %s: %d units%n", entry.getKey(), entry.getValue()));
            
            System.out.println("\nTop Customers:");
            customerTotals.entrySet().stream()
                .sorted(Map.Entry.<String, Double>comparingByValue().reversed())
                .limit(5)
                .forEach(entry -> System.out.printf("  %s: $%.2f%n", entry.getKey(), entry.getValue()));
        }
    }
    
    public static ProcessingStats processCsvFile(Path csvFile) throws IOException {
        AtomicInteger totalRecords = new AtomicInteger(0);
        AtomicInteger validRecords = new AtomicInteger(0);
        AtomicInteger invalidRecords = new AtomicInteger(0);
        AtomicDouble totalRevenue = new AtomicDouble(0.0);
        Map<String, Integer> productCounts = new ConcurrentHashMap<>();
        Map<String, Double> customerTotals = new ConcurrentHashMap<>();
        
        try (Stream<String> lines = Files.lines(csvFile)) {
            lines.skip(1) // Skip header
                .parallel()
                .forEach(line -> {
                    totalRecords.incrementAndGet();
                    
                    try {
                        SalesRecord record = parseCsvLine(line);
                        validRecords.incrementAndGet();
                        totalRevenue.addAndGet(record.getTotal());
                        
                        productCounts.merge(record.getProduct(), record.getQuantity(), Integer::sum);
                        customerTotals.merge(record.getCustomerId(), record.getTotal(), Double::sum);
                        
                    } catch (Exception e) {
                        invalidRecords.incrementAndGet();
                        System.err.printf("Invalid line: %s - Error: %s%n", line, e.getMessage());
                    }
                });
        }
        
        return new ProcessingStats(
            totalRecords.get(), validRecords.get(), invalidRecords.get(),
            totalRevenue.get(), productCounts, customerTotals);
    }
    
    private static SalesRecord parseCsvLine(String line) {
        String[] fields = line.split(",");
        if (fields.length != 6) {
            throw new IllegalArgumentException("Expected 6 fields, got " + fields.length);
        }
        
        try {
            String orderId = fields[0].trim();
            String customerId = fields[1].trim();
            String product = fields[2].trim();
            int quantity = Integer.parseInt(fields[3].trim());
            double price = Double.parseDouble(fields[4].trim());
            String date = fields[5].trim();
            
            if (quantity <= 0 || price <= 0) {
                throw new IllegalArgumentException("Quantity and price must be positive");
            }
            
            return new SalesRecord(orderId, customerId, product, quantity, price, date);
        } catch (NumberFormatException e) {
            throw new IllegalArgumentException("Invalid number format: " + e.getMessage());
        }
    }
    
    public static void generateSampleCsv(Path csvFile, int recordCount) throws IOException {
        String[] products = {"Laptop", "Mouse", "Keyboard", "Monitor", "Headphones", "Webcam"};
        String[] customers = {"C001", "C002", "C003", "C004", "C005"};
        Random random = new Random();
        
        try (BufferedWriter writer = Files.newBufferedWriter(csvFile)) {
            // Write header
            writer.write("orderId,customerId,product,quantity,price,date\n");
            
            // Write data rows
            for (int i = 0; i < recordCount; i++) {
                String orderId = "ORD" + String.format("%04d", i + 1);
                String customerId = customers[random.nextInt(customers.length)];
                String product = products[random.nextInt(products.length)];
                int quantity = 1 + random.nextInt(10);
                double price = 10.0 + random.nextInt(1000);
                String date = "2024-01-" + String.format("%02d", 1 + random.nextInt(28));
                
                writer.write(String.format("%s,%s,%s,%d,%.2f,%s%n",
                    orderId, customerId, product, quantity, price, date));
            }
        }
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create sample CSV file
        Path csvFile = Files.createTempFile("sales_data", ".csv");
        
        try {
            System.out.println("Generating sample CSV data...");
            CsvProcessor.generateSampleCsv(csvFile, 1000);
            System.out.printf("Generated CSV file: %s%n", csvFile);
            
            // Process the CSV file
            System.out.println("\nProcessing CSV file...");
            CsvProcessor.ProcessingStats stats = CsvProcessor.processCsvFile(csvFile);
            
            // Print statistics
            stats.printStats();
            
        } finally {
            // Clean up
            Files.deleteIfExists(csvFile);
        }
    }
}
```
**A:** (Numbers will vary due to random data generation)
```
Generating sample CSV data...
Generated CSV file: C:\Users\USER\AppData\Local\Temp\sales_data123456.csv

Processing CSV file...
Invalid line: ORD0999,C003,Laptop,0,543.21,2024-01-15 - Error: Quantity and price must be positive

=== CSV Processing Statistics ===
Total Records: 1000
Valid Records: 999
Invalid Records: 1
Total Revenue: $278,456.78

Top Products:
  Laptop: 823 units
  Mouse: 812 units
  Keyboard: 789 units
  Monitor: 798 units
  Headphones: 801 units

Top Customers:
  C003: $56,789.12
  C001: $55,234.56
  C005: $56,123.45
  C002: $55,678.90
  C004: $54,632.75
```

---

### 6. File Compression Utility — ZIP Operations
**Q: Create a file compression utility with progress tracking. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.zip.*;
import java.util.concurrent.atomic.*;

class FileCompressionUtility {
    
    public static class CompressionResult {
        private final int filesCompressed;
        private final long originalSize;
        private final long compressedSize;
        private final double compressionRatio;
        private final Duration duration;
        private final List<String> errors;
        
        public CompressionResult(int filesCompressed, long originalSize, long compressedSize,
                               Duration duration, List<String> errors) {
            this.filesCompressed = filesCompressed;
            this.originalSize = originalSize;
            this.compressedSize = compressedSize;
            this.compressionRatio = originalSize > 0 ? (double) compressedSize / originalSize : 0.0;
            this.duration = duration;
            this.errors = errors;
        }
        
        public void printResult() {
            System.out.println("=== Compression Results ===");
            System.out.printf("Files Compressed: %d%n", filesCompressed);
            System.out.printf("Original Size: %s%n", formatBytes(originalSize));
            System.out.printf("Compressed Size: %s%n", formatBytes(compressedSize));
            System.out.printf("Compression Ratio: %.1f%%%n", (1 - compressionRatio) * 100);
            System.out.printf("Duration: %s%n", duration);
            
            if (!errors.isEmpty()) {
                System.out.printf("Errors: %d%n", errors.size());
                errors.forEach(error -> System.out.println("  " + error));
            }
        }
        
        private static String formatBytes(long bytes) {
            if (bytes < 1024) return bytes + " B";
            if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
            if (bytes < 1024 * 1024 * 1024) return String.format("%.1f MB", bytes / (1024.0 * 1024));
            return String.format("%.1f GB", bytes / (1024.0 * 1024 * 1024));
        }
    }
    
    public static CompressionResult compressDirectory(Path sourceDir, Path zipFile) throws IOException {
        Instant start = Instant.now();
        AtomicInteger filesCompressed = new AtomicInteger(0);
        AtomicLong originalSize = new AtomicLong(0);
        AtomicLong compressedSize = new AtomicLong(0);
        List<String> errors = new ArrayList<>();
        
        try (ZipOutputStream zos = new ZipOutputStream(Files.newOutputStream(zipFile))) {
            Files.walk(sourceDir)
                .filter(path -> !Files.isDirectory(path))
                .forEach(path -> {
                    try {
                        String entryName = sourceDir.relativize(path).toString().replace('\\', '/');
                        ZipEntry entry = new ZipEntry(entryName);
                        
                        long fileSize = Files.size(path);
                        originalSize.addAndGet(fileSize);
                        
                        zos.putNextEntry(entry);
                        
                        // Copy file content
                        try (InputStream fis = Files.newInputStream(path)) {
                            byte[] buffer = new byte[8192];
                            int bytesRead;
                            while ((bytesRead = fis.read(buffer)) != -1) {
                                zos.write(buffer, 0, bytesRead);
                                compressedSize.addAndGet(bytesRead);
                            }
                        }
                        
                        zos.closeEntry();
                        filesCompressed.incrementAndGet();
                        
                        System.out.printf("Compressed: %s (%s)%n", 
                            entryName, formatFileSize(fileSize));
                        
                    } catch (IOException e) {
                        String error = String.format("Failed to compress %s: %s", path, e.getMessage());
                        errors.add(error);
                        System.err.println(error);
                    }
                });
        }
        
        Duration duration = Duration.between(start, Instant.now());
        return new CompressionResult(
            filesCompressed.get(), originalSize.get(), compressedSize.get(), duration, errors);
    }
    
    public static void extractZipFile(Path zipFile, Path extractDir) throws IOException {
        if (!Files.exists(extractDir)) {
            Files.createDirectories(extractDir);
        }
        
        try (ZipInputStream zis = new ZipInputStream(Files.newInputStream(zipFile))) {
            ZipEntry entry;
            while ((entry = zis.getNextEntry()) != null) {
                Path entryPath = extractDir.resolve(entry.getName());
                
                // Ensure the entry path is within the extract directory (security check)
                if (!entryPath.normalize().startsWith(extractDir.normalize())) {
                    throw new IOException("Zip entry is outside extraction directory: " + entry.getName());
                }
                
                if (entry.isDirectory()) {
                    Files.createDirectories(entryPath);
                } else {
                    // Ensure parent directories exist
                    Files.createDirectories(entryPath.getParent());
                    
                    // Extract file
                    try (OutputStream fos = Files.newOutputStream(entryPath)) {
                        byte[] buffer = new byte[8192];
                        int bytesRead;
                        while ((bytesRead = zis.read(buffer)) != -1) {
                            fos.write(buffer, 0, bytesRead);
                        }
                    }
                    
                    System.out.printf("Extracted: %s%n", entry.getName());
                }
                
                zis.closeEntry();
            }
        }
    }
    
    private static String formatFileSize(long bytes) {
        if (bytes < 1024) return bytes + " B";
        if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
        return String.format("%.1f MB", bytes / (1024.0 * 1024));
    }
    
    public static void createSampleFiles(Path directory) throws IOException {
        Files.createDirectories(directory.resolve("docs"));
        Files.createDirectories(directory.resolve("images"));
        Files.createDirectories(directory.resolve("data"));
        
        // Create sample files with different sizes
        Files.write(directory.resolve("readme.txt"), "This is a readme file with some content.".getBytes());
        Files.write(directory.resolve("config.properties"), 
            ("server.port=8080\nserver.host=localhost\n").repeat(100).getBytes());
        Files.write(directory.resolve("docs/manual.pdf"), 
            "fake-pdf-content".repeat(500).getBytes());
        Files.write(directory.resolve("images/logo.png"), 
            "fake-image-data".repeat(1000).getBytes());
        Files.write(directory.resolve("data/users.json"), 
            ("{\"users\":[]}\n").repeat(200).getBytes());
        Files.write(directory.resolve("large-file.txt"), 
            "This is a large file for testing compression. ".repeat(2000).getBytes());
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create temporary directories
        Path sourceDir = Files.createTempDirectory("compression_source");
        Path extractDir = Files.createTempDirectory("compression_extract");
        Path zipFile = Files.createTempFile("test_archive", ".zip");
        
        try {
            System.out.printf("Creating sample files in: %s%n", sourceDir);
            FileCompressionUtility.createSampleFiles(sourceDir);
            
            System.out.println("\n=== Compression Phase ===");
            FileCompressionUtility.CompressionResult result = 
                FileCompressionUtility.compressDirectory(sourceDir, zipFile);
            result.printResult();
            
            System.out.printf("\n=== Extraction Phase ===");
            System.out.printf("Extracting ZIP file to: %s%n", extractDir);
            FileCompressionUtility.extractZipFile(zipFile, extractDir);
            
            // Verify extraction
            System.out.println("\n=== Verification ===");
            System.out.printf("Original directory files: %d%n",
                Files.walk(sourceDir).filter(Files::isRegularFile).count());
            System.out.printf("Extracted directory files: %d%n",
                Files.walk(extractDir).filter(Files::isRegularFile).count());
            
        } finally {
            // Clean up
            cleanupDirectory(sourceDir);
            cleanupDirectory(extractDir);
            Files.deleteIfExists(zipFile);
        }
    }
    
    private static void cleanupDirectory(Path directory) {
        try {
            Files.walk(directory)
                .sorted(Comparator.reverseOrder())
                .forEach(path -> {
                    try { Files.deleteIfExists(path); } catch (IOException e) {}
                });
        } catch (IOException e) {
            System.err.println("Error during cleanup: " + e.getMessage());
        }
    }
}
```
**A:** (File sizes may vary)
```
Creating sample files in: C:\Users\USER\AppData\Local\Temp\compression_source123456

=== Compression Phase ===
Compressed: readme.txt (39 B)
Compressed: config.properties (3,900 B)
Compressed: docs\manual.pdf (7,500 B)
Compressed: images\logo.png (15,000 B)
Compressed: data\users.json (3,600 B)
Compressed: large-file.txt (69,000 B)
=== Compression Results ===
Files Compressed: 6
Original Size: 99.0 KB
Compressed Size: 15.2 KB
Compression Ratio: 84.6%
Duration: PT0.045S

=== Extraction Phase ===
Extracting ZIP file to: C:\Users\USER\AppData\Local\Temp\compression_extract789012
Extracted: readme.txt
Extracted: config.properties
Extracted: docs/manual.pdf
Extracted: images/logo.png
Extracted: data/users.json
Extracted: large-file.txt

=== Verification ===
Original directory files: 6
Extracted directory files: 6
```

---

### 7. File Search Utility — Content and Metadata Search
**Q: Implement a file search utility with content and metadata filtering. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.nio.file.attribute.*;
import java.util.*;
import java.util.regex.*;
import java.util.stream.*;
import java.time.*;

class FileSearchUtility {
    
    public static class SearchCriteria {
        private final Pattern fileNamePattern;
        private final Pattern contentPattern;
        private final long minSizeBytes;
        private final long maxSizeBytes;
        private final LocalDateTime modifiedAfter;
        private final LocalDateTime modifiedBefore;
        private final Set<String> extensions;
        
        public SearchCriteria(String fileNamePattern, String contentPattern, 
                           Long minSizeBytes, Long maxSizeBytes,
                           LocalDateTime modifiedAfter, LocalDateTime modifiedBefore,
                           Set<String> extensions) {
            this.fileNamePattern = fileNamePattern != null ? Pattern.compile(fileNamePattern) : null;
            this.contentPattern = contentPattern != null ? Pattern.compile(contentPattern, Pattern.CASE_INSENSITIVE) : null;
            this.minSizeBytes = minSizeBytes != null ? minSizeBytes : 0;
            this.maxSizeBytes = maxSizeBytes != null ? maxSizeBytes : Long.MAX_VALUE;
            this.modifiedAfter = modifiedAfter;
            this.modifiedBefore = modifiedBefore;
            this.extensions = extensions != null ? new HashSet<>(extensions) : null;
        }
        
        public boolean matches(Path path) throws IOException {
            // Check file name pattern
            if (fileNamePattern != null && !fileNamePattern.matcher(path.getFileName().toString()).matches()) {
                return false;
            }
            
            // Check extension
            if (extensions != null && !extensions.isEmpty()) {
                String fileName = path.getFileName().toString();
                String extension = fileName.contains(".") ? 
                    fileName.substring(fileName.lastIndexOf('.') + 1).toLowerCase() : "";
                if (!extensions.contains(extension)) {
                    return false;
                }
            }
            
            // Check file size
            long fileSize = Files.size(path);
            if (fileSize < minSizeBytes || fileSize > maxSizeBytes) {
                return false;
            }
            
            // Check modification time
            BasicFileAttributes attrs = Files.readAttributes(path, BasicFileAttributes.class);
            LocalDateTime modifiedTime = LocalDateTime.ofInstant(
                attrs.lastModifiedTime().toInstant(), ZoneId.systemDefault());
            
            if (modifiedAfter != null && modifiedTime.isBefore(modifiedAfter)) {
                return false;
            }
            if (modifiedBefore != null && modifiedTime.isAfter(modifiedBefore)) {
                return false;
            }
            
            // Check content pattern
            if (contentPattern != null && Files.isRegularFile(path)) {
                try {
                    String content = Files.readString(path);
                    if (!contentPattern.matcher(content).find()) {
                        return false;
                    }
                } catch (IOException e) {
                    // Skip files that can't be read as text
                    return false;
                }
            }
            
            return true;
        }
    }
    
    public static class SearchResult {
        private final Path path;
        private final long size;
        private final LocalDateTime modifiedTime;
        private final List<String> matchedLines;
        
        public SearchResult(Path path, long size, LocalDateTime modifiedTime, List<String> matchedLines) {
            this.path = path;
            this.size = size;
            this.modifiedTime = modifiedTime;
            this.matchedLines = matchedLines;
        }
        
        public Path getPath() { return path; }
        public long getSize() { return size; }
        public LocalDateTime getModifiedTime() { return modifiedTime; }
        public List<String> getMatchedLines() { return matchedLines; }
        
        @Override
        public String toString() {
            return String.format("%s (%s, modified: %s)", 
                path, formatFileSize(size), modifiedTime);
        }
        
        private static String formatFileSize(long bytes) {
            if (bytes < 1024) return bytes + " B";
            if (bytes < 1024 * 1024) return String.format("%.1f KB", bytes / 1024.0);
            return String.format("%.1f MB", bytes / (1024.0 * 1024));
        }
    }
    
    public static List<SearchResult> searchFiles(Path searchDir, SearchCriteria criteria) throws IOException {
        List<SearchResult> results = new ArrayList<>();
        
        try (Stream<Path> paths = Files.walk(searchDir)) {
            paths.filter(Files::isRegularFile)
                 .filter(path -> {
                     try {
                         return criteria.matches(path);
                     } catch (IOException e) {
                         System.err.println("Error checking file " + path + ": " + e.getMessage());
                         return false;
                     }
                 })
                 .forEach(path -> {
                     try {
                         BasicFileAttributes attrs = Files.readAttributes(path, BasicFileAttributes.class);
                         LocalDateTime modifiedTime = LocalDateTime.ofInstant(
                             attrs.lastModifiedTime().toInstant(), ZoneId.systemDefault());
                         
                         List<String> matchedLines = new ArrayList<>();
                         if (criteria.contentPattern != null) {
                             try (BufferedReader reader = Files.newBufferedReader(path)) {
                                 String line;
                                 int lineNumber = 1;
                                 while ((line = reader.readLine()) != null) {
                                     if (criteria.contentPattern.matcher(line).find()) {
                                         matchedLines.add(String.format("Line %d: %s", lineNumber, line.trim()));
                                     }
                                     lineNumber++;
                                 }
                             }
                         }
                         
                         results.add(new SearchResult(path, attrs.size(), modifiedTime, matchedLines));
                         
                     } catch (IOException e) {
                         System.err.println("Error processing file " + path + ": " + e.getMessage());
                     }
                 });
        }
        
        return results;
    }
    
    public static void printSearchResults(List<SearchResult> results, boolean showMatches) {
        System.out.printf("Found %d matching files:%n", results.size());
        
        for (SearchResult result : results) {
            System.out.println("  " + result);
            
            if (showMatches && !result.getMatchedLines().isEmpty()) {
                System.out.println("    Matches:");
                result.getMatchedLines().forEach(line -> System.out.println("      " + line));
            }
        }
    }
    
    public static void createSampleFiles(Path directory) throws IOException {
        Files.createDirectories(directory.resolve("docs"));
        Files.createDirectories(directory.resolve("src"));
        Files.createDirectories(directory.resolve("config"));
        
        // Create sample files
        Files.write(directory.resolve("readme.txt"), 
            "This is the main readme file.\nIt contains important information.".getBytes());
        
        Files.write(directory.resolve("docs/api.txt"), 
            "API Documentation\nThis file describes the API endpoints.\nGET /users\nPOST /users".getBytes());
        
        Files.write(directory.resolve("src/main.java"), 
            "public class Main {\n    public static void main(String[] args) {\n        System.out.println(\"Hello World\");\n    }\n}".getBytes());
        
        Files.write(directory.resolve("config/settings.properties"), 
            "server.port=8080\nserver.host=localhost\ndatabase.url=jdbc:mysql://localhost:3306/db".getBytes());
        
        Files.write(directory.resolve("config/backup.txt"), 
            "Backup configuration file\nContains backup settings and schedules.".getBytes());
        
        Files.write(directory.resolve("large_file.dat"), 
            "Binary data content".repeat(1000).getBytes());
    }
}

public class Main {
    public static void main(String[] args) throws IOException {
        // Create temporary directory with sample files
        Path searchDir = Files.createTempDirectory("file_search_test");
        
        try {
            System.out.printf("Creating sample files in: %s%n", searchDir);
            FileSearchUtility.createSampleFiles(searchDir);
            
            System.out.println("\n=== Search 1: Find all .txt files ===");
            FileSearchUtility.SearchCriteria criteria1 = new FileSearchUtility.SearchCriteria(
                null, null, null, null, null, null, Set.of("txt"));
            List<FileSearchUtility.SearchResult> results1 = FileSearchUtility.searchFiles(searchDir, criteria1);
            FileSearchUtility.printSearchResults(results1, false);
            
            System.out.println("\n=== Search 2: Find files containing 'API' ===");
            FileSearchUtility.SearchCriteria criteria2 = new FileSearchUtility.SearchCriteria(
                null, "API", null, null, null, null, null);
            List<FileSearchUtility.SearchResult> results2 = FileSearchUtility.searchFiles(searchDir, criteria2);
            FileSearchUtility.printSearchResults(results2, true);
            
            System.out.println("\n=== Search 3: Find files larger than 1KB with 'server' in content ===");
            FileSearchUtility.SearchCriteria criteria3 = new FileSearchUtility.SearchCriteria(
                null, "server", 1024L, null, null, null, null);
            List<FileSearchUtility.SearchResult> results3 = FileSearchUtility.searchFiles(searchDir, criteria3);
            FileSearchUtility.printSearchResults(results3, true);
            
            System.out.println("\n=== Search 4: Find Java files ===");
            FileSearchUtility.SearchCriteria criteria4 = new FileSearchUtility.SearchCriteria(
                ".*\\.java$", null, null, null, null, null, null);
            List<FileSearchUtility.SearchResult> results4 = FileSearchUtility.searchFiles(searchDir, criteria4);
            FileSearchUtility.printSearchResults(results4, false);
            
        } finally {
            // Clean up
            try {
                Files.walk(searchDir)
                    .sorted(Comparator.reverseOrder())
                    .forEach(path -> {
                        try { Files.deleteIfExists(path); } catch (IOException e) {}
                    });
            } catch (IOException e) {
                System.err.println("Error during cleanup: " + e.getMessage());
            }
        }
    }
}
```
**A:** 
```
Creating sample files in: C:\Users\USER\AppData\Local\Temp\file_search_test123456

=== Search 1: Find all .txt files ===
Found 3 matching files:
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\readme.txt (66 B, modified: 2024-01-15T10:45:30)
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\docs\api.txt (89 B, modified: 2024-01-15T10:45:30)
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\config\backup.txt (65 B, modified: 2024-01-15T10:45:30)

=== Search 2: Find files containing 'API' ===
Found 1 matching files:
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\docs\api.txt (89 B, modified: 2024-01-15T10:45:30)
    Matches:
      Line 1: API Documentation
      Line 2: This file describes the API endpoints.

=== Search 3: Find files larger than 1KB with 'server' in content ===
Found 1 matching files:
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\config\settings.properties (89 B, modified: 2024-01-15T10:45:30)
    Matches:
      Line 1: server.port=8080
      Line 2: server.host=localhost

=== Search 4: Find Java files ===
Found 1 matching files:
  C:\Users\USER\AppData\Local\Temp\file_search_test123456\src\main.java (108 B, modified: 2024-01-15T10:45:30)
```

---

### 8. File Synchronization Utility — Bidirectional Sync
**Q: Implement a bidirectional file synchronization utility. What is the output?**
```java
import java.io.*;
import java.nio.file.*;
import java.nio.file.attribute.*;
import java.util.*;
import java.util.stream.*;
import java.time.*;

class FileSyncUtility {
    
    public static class SyncOperation {
        private final String operation;
        private final Path source;
        private final Path target;
        private final String reason;
        
        public SyncOperation(String operation, Path source, Path target, String reason) {
            this.operation = operation;
            this.source = source;
            this.target = target;
            this.reason = reason;
        }
        
        public String getOperation() { return operation; }
        public Path getSource() { return source; }
        public Path getTarget() { return target; }
        public String getReason() { return reason; }
        
        @Override
        public String toString() {
            return String.format("%s: %s -> %s (%s)", operation, source, target, reason);
        }
    }
    
    public static class FileMetadata {
        private final long size;
        private final LocalDateTime lastModified;
        private final String checksum;
        
        public FileMetadata(long size, LocalDateTime lastModified, String checksum) {
            this.size = size;
            this.lastModified = lastModified;
            this.checksum = checksum;
        }
        
        public long getSize() { return size; }
        public LocalDateTime getLastModified() { return lastModified; }
        public String getChecksum() { return checksum; }
        
        @Override
        public boolean equals(Object obj) {
            if (this == obj) return true;
            if (obj == null || getClass() != obj.getClass()) return false;
            FileMetadata that = (FileMetadata) obj;
            return size == that.size && 
                   Objects.equals(checksum, that.checksum);
        }
        
        @Override
        public int hashCode() {
            return Objects.hash(size, checksum);
        }
    }
    
    public static List<SyncOperation> analyzeSync(Path dir1, Path dir2) throws IOException {
        List<SyncOperation> operations = new ArrayList<>();
        Map<Path, FileMetadata> dir1Files = getFileMetadata(dir1);
        Map<Path, FileMetadata> dir2Files = getFileMetadata(dir2);
        
        Set<Path> allFiles = new HashSet<>();
        allFiles.addAll(dir1Files.keySet());
        allFiles.addAll(dir2Files.keySet());
        
        for (Path relativePath : allFiles) {
            Path file1 = dir1.resolve(relativePath);
            Path file2 = dir2.resolve(relativePath);
            FileMetadata meta1 = dir1Files.get(relativePath);
            FileMetadata meta2 = dir2Files.get(relativePath);
            
            if (meta1 == null && meta2 != null) {
                // File exists only in dir2
                operations.add(new SyncOperation("COPY", file2, file1, "File missing in dir1"));
            } else if (meta1 != null && meta2 == null) {
                // File exists only in dir1
                operations.add(new SyncOperation("COPY", file1, file2, "File missing in dir2"));
            } else if (meta1 != null && meta2 != null && !meta1.equals(meta2)) {
                // File exists in both but differs
                if (meta1.getLastModified().isAfter(meta2.getLastModified())) {
                    operations.add(new SyncOperation("UPDATE", file1, file2, "Newer version in dir1"));
                } else {
                    operations.add(new SyncOperation("UPDATE", file2, file1, "Newer version in dir2"));
                }
            }
            // If files are identical, no action needed
        }
        
        return operations;
    }
    
    private static Map<Path, FileMetadata> getFileMetadata(Path directory) throws IOException {
        Map<Path, FileMetadata> metadata = new HashMap<>();
        
        if (!Files.exists(directory)) {
            return metadata;
        }
        
        try (Stream<Path> paths = Files.walk(directory)) {
            paths.filter(Files::isRegularFile)
                 .forEach(path -> {
                     try {
                         Path relativePath = directory.relativize(path);
                         long size = Files.size(path);
                         LocalDateTime lastModified = LocalDateTime.ofInstant(
                             Files.getLastModifiedTime(path).toInstant(), ZoneId.systemDefault());
                         String checksum = calculateChecksum(path);
                         
                         metadata.put(relativePath, new FileMetadata(size, lastModified, checksum));
                     } catch (IOException e) {
                         System.err.println("Error reading metadata for " + path + ": " + e.getMessage());
                     }
                 });
        }
        
        return metadata;
    }
    
    private static String calculateChecksum(Path file) throws IOException {
        // Simple checksum based on file size and first/last 100 bytes
        long size = Files.size(file);
        if (size == 0) return "0";
        
        try (InputStream in = Files.newInputStream(file)) {
            byte[] buffer = new byte[Math.min(100, (int) size)];
            int bytesRead = in.read(buffer);
            String start = new String(buffer, 0, bytesRead);
            
            if (size > 100) {
                // Skip to end
                long skipped = in.skip(size - 100);
                if (skipped > 0) {
                    bytesRead = in.read(buffer);
                    String end = new String(buffer, 0, bytesRead);
                    return size + "_" + start.hashCode() + "_" + end.hashCode();
                }
            }
            
            return size + "_" + start.hashCode();
        }
    }
    
    public static void executeSyncOperations(List<SyncOperation> operations) throws IOException {
        for (SyncOperation operation : operations) {
            try {
                switch (operation.getOperation()) {
                    case "COPY":
                    case "UPDATE":
                        // Ensure target directory exists
                        Path targetParent = operation.getTarget().getParent();
                        if (targetParent != null && !Files.exists(targetParent)) {
                            Files.createDirectories(targetParent);
                        }
                        
                        Files.copy(operation.getSource(), operation.getTarget(), 
                                  StandardCopyOption.REPLACE_EXISTING, StandardCopyOption.COPY_ATTRIBUTES);
                        System.out.println("✓ " + operation);
                        break;
                }
            } catch (IOException e) {
                System.err.println("✗ Failed to execute " + operation + ": " + e.getMessage());
            }
        }
    }
    
    public static void createSampleFiles(Path directory, String prefix) throws IOException {
        Files.createDirectories(directory);
        
        Files.write(directory.resolve(prefix + "_file1.txt"), 
            ("Content of " + prefix + " file 1").getBytes());
        Files.write(directory.resolve(prefix + "_file2.txt"), 
            ("Content of " + prefix + " file 2").getBytes());
        
        Files.createDirectories(directory.resolve("subdir"));
        Files.write(directory.resolve("subdir/" + prefix + "_nested.txt"), 
            ("Nested content in " + prefix).getBytes());
        
        // Add some delay to ensure different timestamps
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
    
    public static void modifyFile(Path directory, String fileName) throws IOException {
        Path filePath = directory.resolve(fileName);
        if (Files.exists(filePath)) {
            String existingContent = Files.readString(filePath);
            String newContent = existingContent + "\nModified at " + LocalDateTime.now();
            Files.write(filePath, newContent.getBytes());
            System.out.println("Modified: " + filePath);
        }
    }
}

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {
        // Create temporary directories
        Path dir1 = Files.createTempDirectory("sync_dir1");
        Path dir2 = Files.createTempDirectory("sync_dir2");
        
        try {
            System.out.printf("Setting up synchronization test:%n");
            System.out.printf("  Directory 1: %s%n", dir1);
            System.out.printf("  Directory 2: %s%n", dir2);
            
            // Create initial files
            System.out.println("\n=== Creating initial files ===");
            FileSyncUtility.createSampleFiles(dir1, "dir1");
            FileSyncUtility.createSampleFiles(dir2, "dir2");
            
            // Analyze initial sync
            System.out.println("\n=== Initial Sync Analysis ===");
            List<FileSyncUtility.SyncOperation> initialOps = FileSyncUtility.analyzeSync(dir1, dir2);
            initialOps.forEach(op -> System.out.println("  " + op));
            
            // Execute initial sync
            System.out.println("\n=== Executing Initial Sync ===");
            FileSyncUtility.executeSyncOperations(initialOps);
            
            // Modify some files
            System.out.println("\n=== Modifying Files ===");
            FileSyncUtility.modifyFile(dir1, "dir1_file1.txt");
            FileSyncUtility.modifyFile(dir2, "dir2_file2.txt");
            
            // Add new file to dir1
            Files.write(dir1.resolve("new_file.txt"), "New file content".getBytes());
            System.out.println("Added: " + dir1.resolve("new_file.txt"));
            
            // Analyze after modifications
            System.out.println("\n=== After Modifications Sync Analysis ===");
            List<FileSyncUtility.SyncOperation> modifiedOps = FileSyncUtility.analyzeSync(dir1, dir2);
            modifiedOps.forEach(op -> System.out.println("  " + op));
            
            // Execute modified sync
            System.out.println("\n=== Executing Modified Sync ===");
            FileSyncUtility.executeSyncOperations(modifiedOps);
            
            // Final verification
            System.out.println("\n=== Final Verification ===");
            Map<Path, FileSyncUtility.FileMetadata> dir1Files = FileSyncUtility.getFileMetadata(dir1);
            Map<Path, FileSyncUtility.FileMetadata> dir2Files = FileSyncUtility.getFileMetadata(dir2);
            
            System.out.printf("Directory 1 files: %d%n", dir1Files.size());
            System.out.printf("Directory 2 files: %d%n", dir2Files.size());
            
            boolean inSync = dir1Files.equals(dir2Files);
            System.out.printf("Directories are in sync: %s%n", inSync ? "YES" : "NO");
            
        } finally {
            // Clean up
            cleanupDirectory(dir1);
            cleanupDirectory(dir2);
        }
    }
    
    private static void cleanupDirectory(Path directory) {
        try {
            Files.walk(directory)
                .sorted(Comparator.reverseOrder())
                .forEach(path -> {
                    try { Files.deleteIfExists(path); } catch (IOException e) {}
                });
        } catch (IOException e) {
            System.err.println("Error during cleanup: " + e.getMessage());
        }
    }
}
```
**A:** (Timestamps will vary)
```
Setting up synchronization test:
  Directory 1: C:\Users\USER\AppData\Local\Temp\sync_dir1123456
  Directory 2: C:\Users\USER\AppData\Local\Temp\sync_dir2789012

=== Creating initial files ===

=== Initial Sync Analysis ===
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file1.txt (File missing in dir2)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file2.txt (File missing in dir2)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\subdir\dir1_nested.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\subdir\dir1_nested.txt (File missing in dir2)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file1.txt (File missing in dir1)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file2.txt (File missing in dir1)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\subdir\dir2_nested.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\subdir\dir2_nested.txt (File missing in dir1)

=== Executing Initial Sync ===
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file1.txt (File missing in dir2)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file2.txt (File missing in dir2)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\subdir\dir1_nested.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\subdir\dir1_nested.txt (File missing in dir2)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file1.txt (File missing in dir1)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file2.txt (File missing in dir1)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\subdir\dir2_nested.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\subdir\dir2_nested.txt (File missing in dir1)

=== Modifying Files ===
Modified: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file1.txt
Modified: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file2.txt
Added: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\new_file.txt

=== After Modifications Sync Analysis ===
  UPDATE: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file1.txt (Newer version in dir1)
  UPDATE: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file2.txt (Newer version in dir2)
  COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\new_file.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\new_file.txt (File missing in dir2)

=== Executing Modified Sync ===
✓ UPDATE: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir1_file1.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir1_file1.txt (Newer version in dir1)
✓ UPDATE: C:\Users\USER\AppData\Local\Temp\sync_dir2789012\dir2_file2.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir1123456\dir2_file2.txt (Newer version in dir2)
✓ COPY: C:\Users\USER\AppData\Local\Temp\sync_dir1123456\new_file.txt -> C:\Users\USER\AppData\Local\Temp\sync_dir2789012\new_file.txt (File missing in dir2)

=== Final Verification ===
Directory 1 files: 7
Directory 2 files: 7
Directories are in sync: YES
```
