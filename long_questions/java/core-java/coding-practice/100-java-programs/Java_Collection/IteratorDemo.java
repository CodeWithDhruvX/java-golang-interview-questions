import java.util.*;

/**
 * Demonstrates Iterator Interface methods.
 * 
 * Methods covered:
 * - hasNext()
 * - next()
 * - remove()
 * - forEachRemaining()
 */
public class IteratorDemo {

    public static void main(String[] args) {
        List<String> fruits = new ArrayList<>();
        fruits.add("Apple");
        fruits.add("Banana");
        fruits.add("Cherry");
        fruits.add("Date");

        System.out.println("Original List: " + fruits);

        System.out.println("--- Iterating with Iterator ---");
        Iterator<String> it = fruits.iterator();

        while (it.hasNext()) {
            String fruit = it.next();
            System.out.println("Processing: " + fruit);

            // remove() - Safe removal during iteration
            if (fruit.equals("Banana")) {
                it.remove();
                System.out.println(" (Removed Banana)");
            }
        }
        System.out.println("List after removal: " + fruits);

        System.out.println("\n--- forEachRemaining (Java 8+) ---");
        List<Integer> nums = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        Iterator<Integer> numIt = nums.iterator();

        // Consume first two
        numIt.next();
        numIt.next();

        // Consume the rest
        numIt.forEachRemaining(n -> System.out.println("Remaining: " + n));
    }
}
