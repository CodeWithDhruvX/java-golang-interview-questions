import java.util.Arrays;

public class ProductExceptSelf {
    
    // 238. Product of Array Except Self
    // Time: O(N), Space: O(1) (excluding output)
    public static int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];
        
        // First pass: Calculate left products
        result[0] = 1;
        for (int i = 1; i < n; i++) {
            result[i] = result[i - 1] * nums[i - 1];
        }
        
        // Second pass: Calculate right products and multiply with left
        int right = 1;
        for (int i = n - 1; i >= 0; i--) {
            result[i] = result[i] * right;
            right *= nums[i];
        }
        
        return result;
    }

    public static void main(String[] args) {
        int[][] testCases = {
            {1, 2, 3, 4},
            {-1, 1, 0, -3, 3},
            {1, 2, 3, 4, 5},
            {0, 1, 2, 3},
            {1, 0, 3, 4},
            {1, 2, 0, 4},
            {1, 2, 3, 0},
            {2, 3, 4, 5},
            {-1, -2, -3, -4},
            {1, 1, 1, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i];
            int[] result = productExceptSelf(nums.clone());
            
            System.out.printf("Test Case %d: %s -> %s\n", 
                i + 1, Arrays.toString(nums), Arrays.toString(result));
        }
    }
}
