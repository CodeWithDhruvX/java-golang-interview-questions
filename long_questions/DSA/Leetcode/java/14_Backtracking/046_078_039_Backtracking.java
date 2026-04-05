import java.util.*;

public class BacktrackingProblems {
    
    // 46. Permutations
    // Time: O(N * N!), Space: O(N)
    public static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        if (nums == null || nums.length == 0) {
            return result;
        }
        
        boolean[] used = new boolean[nums.length];
        List<Integer> current = new ArrayList<>();
        backtrackPermute(nums, used, current, result);
        return result;
    }
    
    private static void backtrackPermute(int[] nums, boolean[] used, 
                                       List<Integer> current, List<List<Integer>> result) {
        if (current.size() == nums.length) {
            result.add(new ArrayList<>(current));
            return;
        }
        
        for (int i = 0; i < nums.length; i++) {
            if (used[i]) {
                continue;
            }
            
            used[i] = true;
            current.add(nums[i]);
            
            backtrackPermute(nums, used, current, result);
            
            current.remove(current.size() - 1);
            used[i] = false;
        }
    }

    // 78. Subsets
    // Time: O(N * 2^N), Space: O(N)
    public static List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        if (nums == null || nums.length == 0) {
            result.add(new ArrayList<>());
            return result;
        }
        
        List<Integer> current = new ArrayList<>();
        backtrackSubsets(nums, 0, current, result);
        return result;
    }
    
    private static void backtrackSubsets(int[] nums, int start, 
                                       List<Integer> current, List<List<Integer>> result) {
        result.add(new ArrayList<>(current));
        
        for (int i = start; i < nums.length; i++) {
            current.add(nums[i]);
            backtrackSubsets(nums, i + 1, current, result);
            current.remove(current.size() - 1);
        }
    }

    // 39. Combination Sum
    // Time: O(N^target), Space: O(target)
    public static List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> result = new ArrayList<>();
        if (candidates == null || candidates.length == 0) {
            return result;
        }
        
        Arrays.sort(candidates);
        List<Integer> current = new ArrayList<>();
        backtrackCombinationSum(candidates, target, 0, current, result);
        return result;
    }
    
    private static void backtrackCombinationSum(int[] candidates, int target, int start,
                                               List<Integer> current, List<List<Integer>> result) {
        if (target == 0) {
            result.add(new ArrayList<>(current));
            return;
        }
        
        for (int i = start; i < candidates.length; i++) {
            if (candidates[i] > target) {
                break;
            }
            
            current.add(candidates[i]);
            backtrackCombinationSum(candidates, target - candidates[i], i, current, result);
            current.remove(current.size() - 1);
        }
    }

    public static void main(String[] args) {
        // Test cases for permute
        int[][] testCases1 = {
            {1, 2, 3},
            {0, 1},
            {1},
            {},
            {1, 2, 3, 4},
            {1, 1, 2},
            {2, 2, 2},
            {1, 2, 2},
            {3, 3, 3},
            {1, 2, 3, 4, 5}
        };
        
        // Test cases for subsets
        int[][] testCases2 = {
            {1, 2, 3},
            {0},
            {},
            {1},
            {1, 2},
            {1, 2, 3, 4},
            {1, 1, 2},
            {2, 3, 4},
            {1, 2, 2},
            {1, 2, 3, 4, 5}
        };
        
        // Test cases for combinationSum
        Object[][] testCases3 = {
            {new int[]{2, 3, 6, 7}, 7},
            {new int[]{2, 3, 5}, 8},
            {new int[]{2}, 1},
            {new int[]{1}, 1},
            {new int[]{1, 2}, 3},
            {new int[]{2, 3, 6, 7}, 1},
            {new int[]{3, 5, 7}, 8},
            {new int[]{2, 4, 6, 8}, 8},
            {new int[]{1, 3, 5, 7}, 7},
            {new int[]{2, 3, 5, 8}, 11}
        };
        
        System.out.println("Permutations:");
        for (int i = 0; i < testCases1.length; i++) {
            int[] nums = testCases1[i];
            List<List<Integer>> result = permute(nums);
            System.out.printf("Test Case %d: %s -> %d permutations\n", 
                i + 1, Arrays.toString(nums), result.size());
        }
        
        System.out.println("\nSubsets:");
        for (int i = 0; i < testCases2.length; i++) {
            int[] nums = testCases2[i];
            List<List<Integer>> result = subsets(nums);
            System.out.printf("Test Case %d: %s -> %d subsets\n", 
                i + 1, Arrays.toString(nums), result.size());
        }
        
        System.out.println("\nCombination Sum:");
        for (int i = 0; i < testCases3.length; i++) {
            int[] candidates = (int[]) testCases3[i][0];
            int target = (int) testCases3[i][1];
            List<List<Integer>> result = combinationSum(candidates, target);
            System.out.printf("Test Case %d: %s, target=%d -> %d combinations\n", 
                i + 1, Arrays.toString(candidates), target, result.size());
        }
    }
}
