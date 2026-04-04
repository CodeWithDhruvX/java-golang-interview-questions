import java.util.Arrays;

public class MoveZeroes {
    
    // 283. Move Zeroes
    // Time: O(N), Space: O(1)
    public static void moveZeroes(int[] nums) {
        int lastNonZeroFoundAt = 0;
        
        // Move all non-zero elements to the front
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                nums[lastNonZeroFoundAt] = nums[i];
                lastNonZeroFoundAt++;
            }
        }
        
        // Fill the remaining positions with zeros
        for (int i = lastNonZeroFoundAt; i < nums.length; i++) {
            nums[i] = 0;
        }
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {0, 1, 0, 3, 12},
            {0},
            {1, 2, 3, 4},
            {0, 0, 1, 0, 2, 0, 3},
            {4, 0, 5, 0, 3, 0, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i].clone();
            int[] original = nums.clone();
            
            moveZeroes(nums);
            System.out.printf("Test Case %d: %s -> After moving zeroes: %s\n", 
                i + 1, Arrays.toString(original), Arrays.toString(nums));
        }
    }
}
