import java.util.*;

/**
 * Demonstrates Collections Utility Class methods.
 * 
 * Methods covered:
 * - sort, reverse, shuffle
 * - binarySearch, min, max, frequency
 * - copy, fill
 * - unmodifiableList, synchronizedList, emptyList, singleton
 */
public class CollectionsUtilityDemo {

    public static void main(String[] args) {
        List<Integer> nums = new ArrayList<>(Arrays.asList(5, 2, 9, 1, 6));

        System.out.println("Original: " + nums);

        // Sort
        Collections.sort(nums);
        System.out.println("Sorted: " + nums);

        // Reverse
        Collections.reverse(nums);
        System.out.println("Reversed: " + nums);

        // Shuffle
        Collections.shuffle(nums);
        System.out.println("Shuffled: " + nums);

        // Min / Max
        System.out.println("Min: " + Collections.min(nums));
        System.out.println("Max: " + Collections.max(nums));

        // Binary Search (List must be sorted first)
        Collections.sort(nums);
        int index = Collections.binarySearch(nums, 5); // Assuming 5 is in list
        System.out.println("Sorted for Search: " + nums);
        System.out.println("Index of 5: " + index);

        // Frequency
        nums.add(5);
        System.out.println("Frequency of 5: " + Collections.frequency(nums, 5));

        // Unmodifiable View
        List<Integer> readOnly = Collections.unmodifiableList(nums);
        try {
            readOnly.add(10); // Exception
        } catch (UnsupportedOperationException e) {
            System.out.println("Cannot modify unmodifiable list.");
        }

        // Singleton (Immutable set containing one element)
        Set<String> oneItem = Collections.singleton("One");
        System.out.println("Singleton Set: " + oneItem);
    }
}
