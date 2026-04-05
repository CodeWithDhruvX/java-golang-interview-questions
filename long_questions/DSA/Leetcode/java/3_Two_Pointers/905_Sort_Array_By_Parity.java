import java.util.*;

public class SortArrayByParity {
    
    // 905. Sort Array By Parity
    // Time: O(N), Space: O(1)
    public static int[] sortArrayByParity(int[] nums) {
        if (nums == null || nums.length <= 1) {
            return nums;
        }
        
        int left = 0;
        int right = nums.length - 1;
        
        while (left < right) {
            // Move left pointer to find odd number
            while (left < right && nums[left] % 2 == 0) {
                left++;
            }
            
            // Move right pointer to find even number
            while (left < right && nums[right] % 2 == 1) {
                right--;
            }
            
            // Swap odd on left with even on right
            if (left < right) {
                int temp = nums[left];
                nums[left] = nums[right];
                nums[right] = temp;
                left++;
                right--;
            }
        }
        
        return nums;
    }

    public static void main(String[] args) {
        int[][] testCases = {
            {3, 1, 2, 4},
            {0},
            {1},
            {2, 4, 6, 8},
            {1, 3, 5, 7},
            {1, 2, 3, 4, 5, 6},
            {6, 5, 4, 3, 2, 1},
            {0, 1, 0, 1, 0, 1},
            {2, 2, 2, 2},
            {1, 1, 1, 1},
            {1, 2, 3},
            {3, 2, 1},
            {1, 0},
            {0, 1},
            {1, 2, 3, 4, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i].clone();
            int[] original = nums.clone();
            
            sortArrayByParity(nums);
            System.out.printf("Test Case %d: %s -> %s\n", 
                i + 1, Arrays.toString(original), Arrays.toString(nums));
        }
    }
}
