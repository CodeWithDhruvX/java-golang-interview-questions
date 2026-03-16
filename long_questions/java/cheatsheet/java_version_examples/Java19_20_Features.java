import java.util.concurrent.*;

public class Java19_20_Features {
    public static void main(String[] args) throws Exception {
        System.out.println("Java 19 & 20 Features Demo (Previews)");

        // Java 19: Virtual Threads (Preview) - The foundation of Project Loom
        // Virtual threads are lightweight threads that reduce memory overhead.
        try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
            Future<String> future = executor.submit(() -> {
                return "Running in a Virtual Thread! (Name: " + Thread.currentThread() + ")";
            });
            System.out.println(future.get());
        }

        // Java 19/20: Record Patterns (Preview)
        // Allows deconstructing records in pattern matching
        record Point(int x, int y) {}
        Object obj = new Point(10, 20);
        
        if (obj instanceof Point(int x, int y)) {
            System.out.println("Deconstructed Point: x=" + x + ", y=" + y);
        }

        System.out.println("Java 20 also introduced Scoped Values (Incubator) and structured concurrency updates.");
    }
}
