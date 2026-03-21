public class CountEvenOdd {

    // Brute Force: Simple iteration (already optimal)
    public static int[] countEvenOddBruteForce(int[] arr) {
        int evenCount = 0;
        int oddCount = 0;
        
        for (int num : arr) {
            if (num % 2 == 0) {
                evenCount++;
            } else {
                oddCount++;
            }
        }
        return new int[]{evenCount, oddCount};
    }
    
    // Optimized: Same as brute force (already optimal)
    public static int[] countEvenOddOptimized(int[] arr) {
        return countEvenOddBruteForce(arr);
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        
        System.out.println("Count Even and Odd Elements");
        System.out.println("Array: " + java.util.Arrays.toString(testArray));
        
        int[] bruteResult = countEvenOddBruteForce(testArray);
        System.out.println("Brute Force - Even: " + bruteResult[0] + ", Odd: " + bruteResult[1]);
        
        int[] optResult = countEvenOddOptimized(testArray);
        System.out.println("Optimized - Even: " + optResult[0] + ", Odd: " + optResult[1]);
    }
}
