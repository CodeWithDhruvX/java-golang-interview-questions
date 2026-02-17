import java.util.*;
import java.util.stream.*;

/**
 * Demonstrates Stream API Methods.
 * 
 * Methods covered:
 * - filter, map, flatMap
 * - sorted, distinct, limit, skip
 * - forEach, collect, reduce
 * - count, anyMatch, allMatch, noneMatch
 * - findFirst, findAny
 */
public class StreamDemo {

    public static void main(String[] args) {
        List<String> names = Arrays.asList("Alice", "Bob", "Charlie", "David", "Edward", "Alice");

        System.out.println("Original List: " + names);

        System.out.println("\n--- Intermediate Operations ---");

        // Filter: Names starting with 'A'
        List<String> startingWithA = names.stream()
                .filter(name -> name.startsWith("A"))
                .collect(Collectors.toList());
        System.out.println("Starts with 'A': " + startingWithA);

        // Map: Convert to Uppercase
        List<String> upperCase = names.stream()
                .map(String::toUpperCase)
                .collect(Collectors.toList());
        System.out.println("Uppercase: " + upperCase);

        // Distinct & Sorted
        List<String> distinctSorted = names.stream()
                .distinct()
                .sorted()
                .collect(Collectors.toList());
        System.out.println("Distinct & Sorted: " + distinctSorted);

        // Limit & Skip
        List<String> limitSkip = names.stream()
                .sorted()
                .skip(1)
                .limit(2)
                .collect(Collectors.toList());
        System.out.println("Skip 1, Limit 2 (Sorted): " + limitSkip);

        System.out.println("\n--- Terminal Operations ---");

        // Count
        long count = names.stream().filter(n -> n.length() > 3).count();
        System.out.println("Count (length > 3): " + count);

        // AnyMatch / AllMatch
        boolean anyStartWithZ = names.stream().anyMatch(n -> n.startsWith("Z"));
        System.out.println("Any start with 'Z': " + anyStartWithZ);

        // FindFirst
        Optional<String> first = names.stream().sorted().findFirst();
        first.ifPresent(name -> System.out.println("First sorted name: " + name));

        // Reduce (Concatenate all)
        String concatenated = names.stream().reduce("", (s1, s2) -> s1 + s2);
        System.out.println("Reduced (Concatenated): " + concatenated);

        System.out.println("\n--- FlatMap Example ---");
        List<List<String>> nestedList = Arrays.asList(
                Arrays.asList("One", "Two"),
                Arrays.asList("Three", "Four"));

        List<String> flatList = nestedList.stream()
                .flatMap(List::stream)
                .collect(Collectors.toList());
        System.out.println("Flattened List: " + flatList);
    }
}
