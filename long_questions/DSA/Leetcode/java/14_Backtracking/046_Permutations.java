import java.util.*;

public class Permutations {
    
    // 46. Permutations
    // Time: O(N * N!), Space: O(N!) for result + O(N) for recursion
    public static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        boolean[] used = new boolean[nums.length];
        int[] current = new int[nums.length];
        
        backtrackPermute(nums, used, current, 0, result);
        return result;
    }

    private static void backtrackPermute(int[] nums, boolean[] used, int[] current, 
                                   int pos, List<List<Integer>> result) {
        if (pos == nums.length) {
            // Make a copy of current permutation
            List<Integer> temp = new ArrayList<>();
            for (int num : current) {
                temp.add(num);
            }
            result.add(temp);
            return;
        }
        
        for (int i = 0; i < nums.length; i++) {
            if (!used[i]) {
                used[i] = true;
                current[pos] = nums[i];
                
                backtrackPermute(nums, used, current, pos + 1, result);
                
                used[i] = false;
            }
        }
    }

    // Alternative approach using swapping
    public static List<List<Integer>> permuteSwap(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrackSwap(nums, 0, result);
        return result;
    }

    private static void backtrackSwap(int[] nums, int start, List<List<Integer>> result) {
        if (start == nums.length) {
            // Make a copy of current permutation
            List<Integer> temp = new ArrayList<>();
            for (int num : nums) {
                temp.add(num);
            }
            result.add(temp);
            return;
        }
        
        for (int i = start; i < nums.length; i++) {
            swap(nums, start, i);
            backtrackSwap(nums, start + 1, result);
            swap(nums, start, i);
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3},
            {0, 1},
            {1},
            {},
            {1, 2, 3, 4},
            {1, 1, 2},
            {1, 2, 2},
            {5, 6, 7},
            {1, 2, 3, 4, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            // Make copies for both methods
            int[] nums1 = testCases[i].clone();
            int[] nums2 = testCases[i].clone();
            
            List<List<Integer>> result1 = permute(nums1);
            List<List<Integer>> result2 = permuteSwap(nums2);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Used array: %d permutations\n", result1.size());
            System.out.printf("  Swap method: %d permutations\n\n", result2.size());
        }
    }
}
