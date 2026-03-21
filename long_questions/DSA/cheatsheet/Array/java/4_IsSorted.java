public class IsSorted {

    // Brute Force: Sort and compare
    public static boolean isSortedBruteForce(int[] arr) {
        int[] sorted = arr.clone();
        java.util.Arrays.sort(sorted);
        return java.util.Arrays.equals(arr, sorted);
    }
    
    // Optimized: Single pass adjacent comparison
    public static boolean isSortedOptimized(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] > arr[i + 1]) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        int[] sortedArray = {1, 2, 3, 4, 5};
        int[] unsortedArray = {1, 3, 2, 4, 5};
        
        System.out.println("Check if Array is Sorted");
        
        System.out.println("\nTest 1 - Sorted Array: " + java.util.Arrays.toString(sortedArray));
        System.out.println("Brute Force: " + isSortedBruteForce(sortedArray));
        System.out.println("Optimized: " + isSortedOptimized(sortedArray));
        
        System.out.println("\nTest 2 - Unsorted Array: " + java.util.Arrays.toString(unsortedArray));
        System.out.println("Brute Force: " + isSortedBruteForce(unsortedArray));
        System.out.println("Optimized: " + isSortedOptimized(unsortedArray));
    }
}
