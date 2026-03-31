import java.util.concurrent.*;
import java.util.function.Supplier;

/**
 * Basic Java program demonstrating Runnable vs Callable with Future and CompletableFuture
 * Each line is commented for easy understanding
 */
public class Runnable_Callable_Example {

    public static void main(String[] args) throws Exception {
        
        // 1. Create an ExecutorService to manage thread pool
        // ExecutorService manages creation and lifecycle of threads
        ExecutorService executor = Executors.newFixedThreadPool(2);
        
        System.out.println("=== Demonstrating Runnable (no return value) ===");
        
        // 2. Create a Runnable task using lambda expression
        // Runnable cannot return a value and cannot throw checked exceptions
        Runnable runnableTask = () -> {
            System.out.println("Thread: " + Thread.currentThread().getName() + " - Starting Runnable task");
            
            try {
                // Simulate some work with sleep
                Thread.sleep(1000);
                System.out.println("Thread: " + Thread.currentThread().getName() + " - Runnable task completed");
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt(); // Restore interrupt status
            }
        };
        
        // 3. Submit Runnable to executor (no return value)
        // submit() returns a Future<?> where ? means unknown/void
        Future<?> runnableFuture = executor.submit(runnableTask);
        
        // 4. Wait for Runnable to complete (blocking call)
        runnableFuture.get(); // This blocks until Runnable finishes
        System.out.println("Runnable task finished\n");
        
        System.out.println("=== Demonstrating Callable with Future ===");
        
        // 5. Create a Callable task that returns an Integer
        // Callable can return a value and throw checked exceptions
        Callable<Integer> callableTask = () -> {
            System.out.println("Thread: " + Thread.currentThread().getName() + " - Starting Callable task");
            
            // Simulate computation
            int sum = 0;
            for (int i = 1; i <= 5; i++) {
                sum += i;
                Thread.sleep(200); // Simulate work
            }
            
            System.out.println("Thread: " + Thread.currentThread().getName() + " - Callable task completed");
            return sum; // Return the computed result
        };
        
        // 6. Submit Callable to executor (returns Future<Integer>)
        // Future represents the result of an asynchronous computation
        Future<Integer> callableFuture = executor.submit(callableTask);
        
        // 7. Get the result from Future (blocking call)
        // .get() blocks until the task is complete and returns the result
        Integer result = callableFuture.get();
        System.out.println("Callable returned result: " + result + "\n");
        
        System.out.println("=== Demonstrating CompletableFuture (non-blocking) ===");
        
        // 8. Create a CompletableFuture using supplyAsync
        // supplyAsync runs a Supplier asynchronously and returns CompletableFuture
        CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> {
            System.out.println("Thread: " + Thread.currentThread().getName() + " - Starting CompletableFuture task");
            
            try {
                Thread.sleep(1500); // Simulate work
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
            
            return "Hello from CompletableFuture!";
        }, executor); // Pass executor to use our thread pool
        
        // 9. Chain callbacks without blocking
        // thenAccept() consumes the result when it's ready (non-blocking)
        completableFuture.thenAccept(resultStr -> {
            System.out.println("Callback received result: " + resultStr);
            System.out.println("Callback thread: " + Thread.currentThread().getName());
        });
        
        // 10. Handle exceptions with exceptionally
        completableFuture.exceptionally(ex -> {
            System.out.println("Error occurred: " + ex.getMessage());
            return "Default value on error";
        });
        
        System.out.println("Main thread continues without blocking...");
        
        // 11. Wait a bit to see CompletableFuture complete
        Thread.sleep(2000);
        
        // 12. Shutdown the executor (important!)
        executor.shutdown();
        executor.awaitTermination(5, TimeUnit.SECONDS);
        
        System.out.println("Program completed!");
    }
}
