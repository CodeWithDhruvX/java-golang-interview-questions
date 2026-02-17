import java.util.*;

/**
 * Demonstrates Collection Interface common methods.
 * 
 * Methods covered:
 * - add(E e), addAll(Collection c)
 * - remove(Object o), removeAll(Collection c), retainAll(Collection c)
 * - clear(), size(), isEmpty()
 * - contains(Object o), containsAll(Collection c)
 * - toArray()
 * - stream(), parallelStream()
 */
public class CollectionDemo {

    public static void main(String[] args) {
        // Using ArrayList as the Collection implementation
        Collection<String> collection = new ArrayList<>();

        System.out.println("--- Adding Elements ---");
        // add(E e)
        collection.add("Apple");
        collection.add("Banana");
        collection.add("Cherry");
        System.out.println("Collection after add: " + collection);

        // addAll(Collection c)
        Collection<String> moreFruits = new ArrayList<>();
        moreFruits.add("Date");
        moreFruits.add("Elderberry");
        collection.addAll(moreFruits);
        System.out.println("Collection after addAll: " + collection);

        System.out.println("\n--- Checking Status ---");
        // size()
        System.out.println("Size: " + collection.size());
        // isEmpty()
        System.out.println("Is Empty: " + collection.isEmpty());
        // contains(Object o)
        System.out.println("Contains 'Banana': " + collection.contains("Banana"));
        // containsAll(Collection c)
        System.out.println("Contains all 'moreFruits': " + collection.containsAll(moreFruits));

        System.out.println("\n--- Removing Elements ---");
        // remove(Object o)
        collection.remove("Apple");
        System.out.println("Removed 'Apple': " + collection);

        // removeAll(Collection c)
        Collection<String> toRemove = new ArrayList<>();
        toRemove.add("Date");
        collection.removeAll(toRemove);
        System.out.println("Removed 'Date' (removeAll): " + collection);

        // retainAll(Collection c) (Intersection)
        Collection<String> toRetain = new ArrayList<>();
        toRetain.add("Banana");
        toRetain.add("Cherry");
        collection.retainAll(toRetain);
        System.out.println("Retained only 'Banana' and 'Cherry': " + collection);

        System.out.println("\n--- Conversion ---");
        // toArray()
        Object[] array = collection.toArray();
        System.out.println("Converted to Array: " + Arrays.toString(array));

        System.out.println("\n--- Streams ---");
        // stream()
        System.out.println("Stream processing:");
        collection.stream()
            .filter(s -> s.startsWith("B"))
            .forEach(System.out::println);

        System.out.println("\n--- Clearing ---");
        // clear()
        collection.clear();
        System.out.println("Collection after clear: " + collection);
        System.out.println("Is Empty now: " + collection.isEmpty());
    }
}
