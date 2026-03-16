import java.util.*;
import java.util.concurrent.*;

public class Java21_Features {
    public static void main(String[] args) throws Exception {
        // Java 21: Virtual Threads, Sequenced Collections, Record Patterns
        
        // 1. Sequenced Collections
        List<String> list = new ArrayList<>(List.of("First", "Middle", "Last"));
        System.out.println("Sequenced Access - First: " + list.getFirst());
        System.out.println("Sequenced Access - Last: " + list.getLast());

        // 2. Virtual Threads (Lightweight threads)
        try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
            Future<String> future = executor.submit(() -> {
                Thread.sleep(100);
                return "Completed in Virtual Thread: " + Thread.currentThread();
            });
            System.out.println(future.get());
        }

        System.out.println("Java 21 Features: Virtual Threads, Sequenced Collections, Record Patterns, Pattern Matching for Switch.");
    }
}
