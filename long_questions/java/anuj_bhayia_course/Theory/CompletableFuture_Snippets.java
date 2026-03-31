import java.util.concurrent.*;
import java.util.function.Supplier;

/**
 * Small CompletableFuture snippets for interview preparation
 */
public class CompletableFuture_Snippets {

    public static void main(String[] args) throws Exception {
        
        // SNIPPET 1: Basic CompletableFuture with supplyAsync
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
            return "Hello World!";
        });
        
        System.out.println(future.get()); // Blocking get
        
        // SNIPPET 2: Non-blocking with thenAccept
        CompletableFuture.supplyAsync(() -> "Task Result")
            .thenAccept(result -> System.out.println("Result: " + result));
        
        // SNIPPET 3: Chain operations with thenApply
        CompletableFuture.supplyAsync(() -> 5)
            .thenApply(num -> num * 2)
            .thenApply(result -> "Double of 5 is: " + result)
            .thenAccept(System.out::println);
        
        // SNIPPET 4: Exception handling
        CompletableFuture.supplyAsync(() -> {
            throw new RuntimeException("Something went wrong!");
        })
        .exceptionally(ex -> "Error handled: " + ex.getMessage())
        .thenAccept(System.out::println);
        
        // SNIPPET 5: Combine two futures
        CompletableFuture<String> future1 = CompletableFuture.supplyAsync(() -> "Hello");
        CompletableFuture<String> future2 = CompletableFuture.supplyAsync(() -> "World");
        
        future1.thenCombine(future2, (s1, s2) -> s1 + " " + s2)
               .thenAccept(System.out::println);
        
        // SNIPPET 6: Run multiple tasks in parallel
        CompletableFuture<Void> task1 = CompletableFuture.runAsync(() -> {
            System.out.println("Task 1 running");
        });
        
        CompletableFuture<Void> task2 = CompletableFuture.runAsync(() -> {
            System.out.println("Task 2 running");
        });
        
        CompletableFuture.allOf(task1, task2)
            .thenRun(() -> System.out.println("All tasks completed!"));
        
        // Wait to see async results
        Thread.sleep(1000);
    }
}
