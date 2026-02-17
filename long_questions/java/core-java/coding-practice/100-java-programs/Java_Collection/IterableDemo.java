import java.util.*;
import java.util.function.Consumer;

/**
 * Demonstrates Iterable Interface methods.
 * 
 * Methods covered:
 * - iterator()
 * - forEach()
 * - spliterator()
 */
public class IterableDemo {
    public static void main(String[] args) {
        // Create a list (which implements Iterable)
        List<String> names = new ArrayList<>();
        names.add("Alice");
        names.add("Bob");
        names.add("Charlie");
        names.add("David");

        System.out.println("--- 1. iterator() ---");
        Iterator<String> iterator = names.iterator();
        while (iterator.hasNext()) {
            System.out.println("Name: " + iterator.next());
        }

        System.out.println("\n--- 2. forEach() (Java 8+) ---");
        // Using lambda expression
        names.forEach(name -> System.out.println("Name: " + name));
        
        // Using method reference
        System.out.println("(Method Reference version):");
        names.forEach(System.out::println);

        System.out.println("\n--- 3. spliterator() (Java 8+) ---");
        Spliterator<String> spliterator = names.spliterator();
        System.out.println("Characteristics: " + spliterator.characteristics());
        System.out.println("Estimate Size: " + spliterator.estimateSize());
        
        System.out.println("Traversing with tryAdvance:");
        spliterator.tryAdvance(name -> System.out.println("Processed: " + name));
        
        System.out.println("Traversing remaining with forEachRemaining:");
        spliterator.forEachRemaining(name -> System.out.println("Remaining: " + name));
    }
}
