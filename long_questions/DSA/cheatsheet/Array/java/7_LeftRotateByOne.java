public class LeftRotateByOne {

    // Brute Force: Using extra array
    public static int[] leftRotateByOneBruteForce(int[] arr) {
        int n = arr.length;
        if (n == 0) return arr;
        
        int[] rotated = new int[n];
        for (int i = 0; i < n - 1; i++) {
            rotated[i] = arr[i + 1];
        }
        rotated[n - 1] = arr[0];
        return rotated;
    }
    
    // Optimized: In-place with temp variable
    public static void leftRotateByOneOptimized(int[] arr) {
        if (arr.length <= 1) return;
        
        int temp = arr[0];
        for (int i = 1; i < arr.length; i++) {
            arr[i - 1] = arr[i];
        }
        arr[arr.length - 1] = temp;
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 3, 4, 5};
        
        System.out.println("Left Rotate Array by 1 Position");
        System.out.println("Original: " + java.util.Arrays.toString(testArray));
        
        int[] bruteResult = leftRotateByOneBruteForce(testArray.clone());
        System.out.println("Brute Force: " + java.util.Arrays.toString(bruteResult));
        
        int[] optArray = testArray.clone();
        leftRotateByOneOptimized(optArray);
        System.out.println("Optimized: " + java.util.Arrays.toString(optArray));
    }
}
