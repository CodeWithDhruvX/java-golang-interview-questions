public class FindMinMax {

    // Brute Force: Sort and pick first/last
    public static int[] findMinMaxBruteForce(int[] arr) {
        if (arr.length == 0) return new int[]{};
        
        // Create a copy and sort it
        int[] sorted = arr.clone();
        java.util.Arrays.sort(sorted);
        
        // First element is min, last is max
        return new int[]{sorted[0], sorted[sorted.length - 1]};
    }
    
    // Optimized: Single pass with two variables
    public static int[] findMinMaxOptimized(int[] arr) {
        if (arr.length == 0) return new int[]{};
        
        int minVal = arr[0];
        int maxVal = arr[0];
        
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > maxVal) {
                maxVal = arr[i];
            }
            if (arr[i] < minVal) {
                minVal = arr[i];
            }
        }
        return new int[]{minVal, maxVal};
    }

    public static void main(String[] args) {
        int[] testArray = {5, 2, 8, 1, 9, 3};
        
        System.out.println("Find Min and Max Element");
        System.out.println("Array: " + java.util.Arrays.toString(testArray));
        
        int[] bruteResult = findMinMaxBruteForce(testArray);
        System.out.println("Brute Force - Min: " + bruteResult[0] + ", Max: " + bruteResult[1]);
        
        int[] optResult = findMinMaxOptimized(testArray);
        System.out.println("Optimized - Min: " + optResult[0] + ", Max: " + optResult[1]);
    }
}
