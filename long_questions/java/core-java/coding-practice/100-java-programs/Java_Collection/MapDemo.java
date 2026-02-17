import java.util.*;

/**
 * Demonstrates Map Interface methods and implementations.
 * Note: Map does NOT extend Collection.
 * 
 * Implementations: HashMap, LinkedHashMap, TreeMap, Hashtable,
 * ConcurrentHashMap
 * 
 * Methods covered:
 * - put, putAll, get, getOrDefault
 * - remove, replace
 * - containsKey, containsValue
 * - keySet, values, entrySet
 * - compute, merge
 */
public class MapDemo {

    public static void main(String[] args) {
        System.out.println("=== 1. HashMap Implementation ===");
        // Unordered, allows one null key
        Map<String, Integer> hashMap = new HashMap<>();
        hashMap.put("Java", 100);
        hashMap.put("Python", 90);
        hashMap.put("C++", 80);
        hashMap.put(null, 0);

        System.out.println("HashMap: " + hashMap);

        // get / getOrDefault
        System.out.println("Get 'Java': " + hashMap.get("Java"));
        System.out.println("Get 'Go' (Default): " + hashMap.getOrDefault("Go", 0));

        // contains
        System.out.println("Contains Key 'Python': " + hashMap.containsKey("Python"));
        System.out.println("Contains Value 100: " + hashMap.containsValue(100));

        // Iteration
        System.out.println("--- Iterating EntrySet ---");
        for (Map.Entry<String, Integer> entry : hashMap.entrySet()) {
            System.out.println(entry.getKey() + " -> " + entry.getValue());
        }

        System.out.println("\n=== 2. LinkedHashMap Implementation ===");
        // Maintains insertion order
        Map<String, Integer> linkedMap = new LinkedHashMap<>();
        linkedMap.put("First", 1);
        linkedMap.put("Second", 2);
        linkedMap.put("Third", 3);

        System.out.println("LinkedHashMap (Ordered): " + linkedMap);

        System.out.println("\n=== 3. TreeMap Implementation ===");
        // Sorted by keys
        Map<String, Integer> treeMap = new TreeMap<>();
        treeMap.put("Zebra", 10);
        treeMap.put("Apple", 20);
        treeMap.put("Mango", 30);

        System.out.println("TreeMap (Sorted Keys): " + treeMap);

        System.out.println("\n=== 4. Advanced Map Methods (Java 8+) ===");
        // computeIfAbsent
        hashMap.computeIfAbsent("Rust", k -> 50);
        System.out.println("After computeIfAbsent 'Rust': " + hashMap);

        // merge
        // If key exists, add 10 to current value. If not, set to 10.
        hashMap.merge("Java", 10, (oldVal, newVal) -> oldVal + newVal);
        System.out.println("After merge 'Java' (+10): " + hashMap);
    }
}
