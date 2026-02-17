import java.util.*;

/**
 * Demonstrates Arrays Utility Class methods.
 * 
 * Methods covered:
 * - sort, parallelSort
 * - binarySearch, equals, fill
 * - copyOf, asList
 * - toString, deepToString
 */
public class ArraysUtilityDemo {

    public static void main(String[] args) {
        int[] arr = { 5, 2, 9, 1, 6 };

        System.out.println("Original: " + Arrays.toString(arr));

        // Sort
        Arrays.sort(arr);
        System.out.println("Sorted: " + Arrays.toString(arr));

        // Binary Search
        int index = Arrays.binarySearch(arr, 9);
        System.out.println("Index of 9: " + index);

        // Fill
        int[] filled = new int[5];
        Arrays.fill(filled, 7);
        System.out.println("Filled with 7: " + Arrays.toString(filled));

        // CopyOf
        int[] copy = Arrays.copyOf(arr, 10); // Copies and pads with 0
        System.out.println("Copy (Length 10): " + Arrays.toString(copy));

        // asList (Bridge between Array and Collection)
        String[] strArr = { "A", "B", "C" };
        List<String> list = Arrays.asList(strArr);
        System.out.println("asList: " + list);
        // list.add("D"); // Exception: Arrays.asList returns fixed-size list

        // DeepToString (For multidimensional arrays)
        int[][] matrix = { { 1, 2 }, { 3, 4 } };
        System.out.println("Matrix (deepToString): " + Arrays.deepToString(matrix));
    }
}
