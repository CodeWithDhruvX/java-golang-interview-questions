import java.util.Arrays;

public class RotateArray {
    
    // 189. Rotate Array
    // Time: O(N), Space: O(1)
    public static void rotate(int[] nums, int k) {
        int n = nums.length;
        k = k % n; // Handle k > n
        
        if (k == 0) return;
        
        // Reverse the entire array
        reverse(nums, 0, n - 1);
        // Reverse first k elements
        reverse(nums, 0, k - 1);
        // Reverse remaining elements
        reverse(nums, k, n - 1);
    }
    
    private static void reverse(int[] nums, int left, int right) {
        while (left < right) {
            int temp = nums[left];
            nums[left] = nums[right];
            nums[right] = temp;
            left++;
            right--;
        }
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 2, 3, 4, 5, 6, 7}, 3},
            {new int[]{-1, -100, 3, 99}, 2},
            {new int[]{1, 2, 3, 4, 5}, 7},
            {new int[]{1}, 0},
            {new int[]{1, 2}, 1},
            {new int[]{1, 2, 3}, 4},
            {new int[]{1, 2, 3, 4, 5, 6}, 2},
            {new int[]{1, 2, 3, 4, 5, 6}, 3}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = (int[]) testCases[i][0];
            int k = (int) testCases[i][1];
            int[] original = nums.clone();
            
            rotate(nums, k);
            System.out.printf("Test Case %d: %s rotated by %d -> %s\n", 
                i + 1, Arrays.toString(original), k, Arrays.toString(nums));
        }
    }
}
