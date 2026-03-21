public class ArraySum {

    // Brute Force: Simple iteration (already optimal)
    public static long sumBruteForce(int[] arr) {
        long sum = 0;
        for (int num : arr) {
            sum += num;
        }
        return sum;
    }
    
    // Optimized: Same as brute force (already optimal)
    public static long sumOptimized(int[] arr) {
        return sumBruteForce(arr);
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        
        System.out.println("Find Sum of All Elements");
        System.out.println("Array: " + java.util.Arrays.toString(testArray));
        
        long bruteResult = sumBruteForce(testArray);
        System.out.println("Brute Force: " + bruteResult);
        
        long optResult = sumOptimized(testArray);
        System.out.println("Optimized: " + optResult);
    }
}
