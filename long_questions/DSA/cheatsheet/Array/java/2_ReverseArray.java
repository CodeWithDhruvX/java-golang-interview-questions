public class ReverseArray {

    // Brute Force: Using extra array
    public static int[] reverseBruteForce(int[] arr) {
        int n = arr.length;
        int[] reversed = new int[n];
        
        for (int i = 0; i < n; i++) {
            reversed[i] = arr[n - 1 - i];
        }
        return reversed;
    }
    
    // Optimized: Two pointers swapping
    public static void reverseOptimized(int[] arr) {
        int left = 0;
        int right = arr.length - 1;
        
        while (left < right) {
            // Swap elements
            int temp = arr[left];
            arr[left] = arr[right];
            arr[right] = temp;
            
            left++;
            right--;
        }
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 3, 4, 5};
        
        System.out.println("Reverse Array In-Place");
        System.out.println("Original: " + java.util.Arrays.toString(testArray));
        
        int[] bruteResult = reverseBruteForce(testArray.clone());
        System.out.println("Brute Force: " + java.util.Arrays.toString(bruteResult));
        
        int[] optArray = testArray.clone();
        reverseOptimized(optArray);
        System.out.println("Optimized: " + java.util.Arrays.toString(optArray));
    }
}
