public class SecondLargest {

    // Brute Force: Sort and pick second last
    public static int secondLargestBruteForce(int[] arr) {
        if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
        
        int[] sorted = arr.clone();
        java.util.Arrays.sort(sorted);
        
        // Handle duplicates - find the first element different from the largest
        for (int i = sorted.length - 2; i >= 0; i--) {
            if (sorted[i] != sorted[sorted.length - 1]) {
                return sorted[i];
            }
        }
        throw new IllegalArgumentException("No second largest element (all elements are equal)");
    }
    
    // Optimized: Single pass with two variables
    public static int secondLargestOptimized(int[] arr) {
        if (arr.length < 2) throw new IllegalArgumentException("Array must have at least 2 elements");
        
        int largest = Integer.MIN_VALUE;
        int secondLargest = Integer.MIN_VALUE;
        
        for (int num : arr) {
            if (num > largest) {
                secondLargest = largest;
                largest = num;
            } else if (num > secondLargest && num != largest) {
                secondLargest = num;
            }
        }
        
        if (secondLargest == Integer.MIN_VALUE) {
            throw new IllegalArgumentException("No second largest element (all elements are equal)");
        }
        return secondLargest;
    }

    public static void main(String[] args) {
        int[] testArray = {5, 2, 8, 1, 9, 3, 8};
        
        System.out.println("Find Second Largest Element");
        System.out.println("Array: " + java.util.Arrays.toString(testArray));
        
        try {
            int bruteResult = secondLargestBruteForce(testArray);
            System.out.println("Brute Force: " + bruteResult);
            
            int optResult = secondLargestOptimized(testArray);
            System.out.println("Optimized: " + optResult);
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }
    }
}
