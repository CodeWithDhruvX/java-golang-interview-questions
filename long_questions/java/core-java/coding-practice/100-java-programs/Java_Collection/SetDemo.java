import java.util.*;

/**
 * Demonstrates Set Interface methods and implementations.
 * 
 * Implementations: HashSet, LinkedHashSet, TreeSet
 * 
 * Key Characteristics:
 * - No duplicates allowed
 * - Specific ordering depends on implementation
 */
public class SetDemo {

    public static void main(String[] args) {
        System.out.println("=== 1. HashSet Implementation ===");
        // Unordered, allows null, fast access (O(1))
        Set<String> hashSet = new HashSet<>();
        hashSet.add("Zebra");
        hashSet.add("Ant");
        hashSet.add("Lion");
        hashSet.add("Ant"); // Duplicate ignored
        hashSet.add(null); // Allows null

        System.out.println("HashSet (Unordered): " + hashSet);

        // Check containment
        System.out.println("Contains 'Lion': " + hashSet.contains("Lion"));

        System.out.println("\n=== 2. LinkedHashSet Implementation ===");
        // Ordered by insertion, slower than HashSet
        Set<String> linkedHashSet = new LinkedHashSet<>();
        linkedHashSet.add("Zebra");
        linkedHashSet.add("Ant");
        linkedHashSet.add("Lion");
        linkedHashSet.add("Ant"); // Duplicate ignored

        System.out.println("LinkedHashSet (Insertion Order): " + linkedHashSet);

        System.out.println("\n=== 3. TreeSet Implementation ===");
        // Sorted (Natural order or Comparator), slower (O(log n)), no nulls
        try {
            TreeSet<String> treeSet = new TreeSet<>();
            treeSet.add("Zebra");
            treeSet.add("Ant");
            treeSet.add("Lion");
            // treeSet.add(null); // Throws NullPointerException

            System.out.println("TreeSet (Sorted): " + treeSet);

            // NavigableSet methods (TreeSet implements NavigableSet)
            System.out.println("First: " + treeSet.first());
            System.out.println("Last: " + treeSet.last());
            System.out.println("Higher than 'Ant': " + treeSet.higher("Ant"));
            System.out.println("Lower than 'Lion': " + treeSet.lower("Lion"));

        } catch (Exception e) {
            e.printStackTrace();
        }

        System.out.println("\n=== Set Operations (Mathematical) ===");
        Set<Integer> a = new HashSet<>(Arrays.asList(1, 3, 2, 4, 8, 9, 0));
        Set<Integer> b = new HashSet<>(Arrays.asList(1, 3, 7, 5, 4, 0, 7, 5));

        // Union
        Set<Integer> union = new HashSet<>(a);
        union.addAll(b);
        System.out.println("Union: " + union);

        // Intersection
        Set<Integer> intersection = new HashSet<>(a);
        intersection.retainAll(b);
        System.out.println("Intersection: " + intersection);

        // Difference
        Set<Integer> difference = new HashSet<>(a);
        difference.removeAll(b);
        System.out.println("Difference (a - b): " + difference);
    }
}
