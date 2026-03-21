import java.util.*;

public class RemoveDuplicates {

    // Brute Force: Using HashSet
    public static int removeDuplicatesBruteForce(int[] arr) {
        Set<Integer> unique = new LinkedHashSet<>();
        for (int num : arr) {
            unique.add(num);
        }
        
        int index = 0;
        for (int num : unique) {
            arr[index++] = num;
        }
        return unique.size();
    }
    
    // Optimized: Two pointers
    public static int removeDuplicatesOptimized(int[] arr) {
        if (arr.length == 0) return 0;
        
        int i = 0;
        for (int j = 1; j < arr.length; j++) {
            if (arr[j] != arr[i]) {
                i++;
                arr[i] = arr[j];
            }
        }
        return i + 1;
    }

    public static void main(String[] args) {
        int[] sortedArray = {1, 1, 2, 2, 2, 3, 4, 4, 5};
        
        System.out.println("Remove Duplicates from Sorted Array");
        System.out.println("Original: " + Arrays.toString(sortedArray));
        
        int[] bruteArray = sortedArray.clone();
        int bruteLength = removeDuplicatesBruteForce(bruteArray);
        System.out.println("Brute Force: " + Arrays.toString(Arrays.copyOf(bruteArray, bruteLength)));
        
        int[] optArray = sortedArray.clone();
        int optLength = removeDuplicatesOptimized(optArray);
        System.out.println("Optimized: " + Arrays.toString(Arrays.copyOf(optArray, optLength)));
    }
}
