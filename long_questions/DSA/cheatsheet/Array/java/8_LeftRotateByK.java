public class LeftRotateByK {

    // Brute Force: K times rotate by 1
    public static void leftRotateByKBruteForce(int[] arr, int k) {
        int n = arr.length;
        if (n == 0) return;
        
        k = k % n;
        
        for (int i = 0; i < k; i++) {
            leftRotateByOneOptimized(arr);
        }
    }
    
    // Optimized: Reversal algorithm
    public static void leftRotateByKOptimized(int[] arr, int k) {
        int n = arr.length;
        if (n == 0) return;
        
        k = k % n;
        if (k == 0) return;
        
        // Reverse first k elements
        reverse(arr, 0, k - 1);
        // Reverse remaining n-k elements
        reverse(arr, k, n - 1);
        // Reverse entire array
        reverse(arr, 0, n - 1);
    }
    
    private static void reverse(int[] arr, int start, int end) {
        while (start < end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
    }
    
    // Helper method for brute force
    private static void leftRotateByOneOptimized(int[] arr) {
        if (arr.length <= 1) return;
        
        int temp = arr[0];
        for (int i = 1; i < arr.length; i++) {
            arr[i - 1] = arr[i];
        }
        arr[arr.length - 1] = temp;
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 3, 4, 5};
        int k = 2;
        
        System.out.println("Left Rotate Array by K Positions (K=" + k + ")");
        System.out.println("Original: " + java.util.Arrays.toString(testArray));
        
        int[] bruteArray = testArray.clone();
        leftRotateByKBruteForce(bruteArray, k);
        System.out.println("Brute Force: " + java.util.Arrays.toString(bruteArray));
        
        int[] optArray = testArray.clone();
        leftRotateByKOptimized(optArray, k);
        System.out.println("Optimized: " + java.util.Arrays.toString(optArray));
    }
}
