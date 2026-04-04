import java.util.Arrays;

public class RemoveDuplicatesFromSortedArray {
    
    // 26. Remove Duplicates from Sorted Array
    // Time: O(N), Space: O(1)
    public static int removeDuplicates(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        int uniqueIndex = 0;
        
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[uniqueIndex]) {
                uniqueIndex++;
                nums[uniqueIndex] = nums[i];
            }
        }
        
        return uniqueIndex + 1;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 1, 2},
            {0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
            {1, 2, 3, 4, 5},
            {1, 1, 1, 1},
            {},
            {2}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i].clone();
            int[] original = nums.clone();
            
            int length = removeDuplicates(nums);
            int[] result = Arrays.copyOf(nums, length);
            
            System.out.printf("Test Case %d: %s -> Length: %d, Unique elements: %s\n", 
                i + 1, Arrays.toString(original), length, Arrays.toString(result));
        }
    }
}
