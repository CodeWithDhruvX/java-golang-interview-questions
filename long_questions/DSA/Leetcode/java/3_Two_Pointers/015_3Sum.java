import java.util.*;

public class ThreeSum {
    
    // 15. 3Sum
    // Time: O(N^2), Space: O(1) (ignoring output space)
    public static List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();
        int n = nums.length;
        
        for (int i = 0; i < n - 2; i++) {
            // Skip duplicates for the first element
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            
            // Two pointers approach for the remaining two elements
            int left = i + 1, right = n - 1;
            int target = -nums[i];
            
            while (left < right) {
                int sum = nums[left] + nums[right];
                
                if (sum == target) {
                    result.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    
                    // Skip duplicates for the second element
                    while (left < right && nums[left] == nums[left + 1]) {
                        left++;
                    }
                    // Skip duplicates for the third element
                    while (left < right && nums[right] == nums[right - 1]) {
                        right--;
                    }
                    
                    left++;
                    right--;
                } else if (sum < target) {
                    left++;
                } else {
                    right--;
                }
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {-1, 0, 1, 2, -1, -4},
            {0, 1, 1},
            {0, 0, 0},
            {-2, 0, 1, 1, 2},
            {-1, -2, -3, -4, -5},
            {1, 2, -2, -1},
            {3, -2, 1, 0, -1, 2, -3},
            {},
            {0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<List<Integer>> result = threeSum(testCases[i]);
            System.out.printf("Test Case %d: %s -> Triplets: %s\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
