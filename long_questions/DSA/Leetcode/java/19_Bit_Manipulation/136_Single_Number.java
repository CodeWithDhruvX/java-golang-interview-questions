import java.util.*;

public class SingleNumber {
    
    // 136. Single Number
    // Time: O(N), Space: O(1)
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num;
        }
        return result;
    }

    // Alternative approach using hash map (O(N) space)
    public static int singleNumberHash(int[] nums) {
        Map<Integer, Integer> count = new HashMap<>();
        for (int num : nums) {
            count.put(num, count.getOrDefault(num, 0) + 1);
        }
        
        for (Map.Entry<Integer, Integer> entry : count.entrySet()) {
            if (entry.getValue() == 1) {
                return entry.getKey();
            }
        }
        
        return -1; // Should not reach here for valid input
    }

    // XOR properties explanation
    public static int singleNumberWithExplanation(int[] nums) {
        System.out.printf("XOR Properties:\n");
        System.out.printf("1. a ^ a = 0 (XOR of same numbers is 0)\n");
        System.out.printf("2. a ^ 0 = a (XOR with 0 is the number itself)\n");
        System.out.printf("3. XOR is commutative and associative\n\n");
        
        System.out.printf("Processing: ");
        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > 0) {
                System.out.printf(" ^ ");
            }
            System.out.printf("%d", nums[i]);
            result ^= nums[i];
        }
        System.out.printf(" = %d\n\n", result);
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
            {1},
            {2, 2, 3, 3, 4},
            {0, 1, 1},
            {-1, -1, -2},
            {99, 99, 100},
            {5, 5, 6, 6, 7, 7, 8},
            {10, 10, 10, 10, 15}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            
            if (i == 0) {
                // Show detailed explanation for first test case
                int result = singleNumberWithExplanation(testCases[i]);
                System.out.printf("Single number: %d\n\n", result);
            } else {
                int result1 = singleNumber(testCases[i]);
                int result2 = singleNumberHash(testCases[i]);
                System.out.printf("XOR: %d, Hash: %d\n\n", result1, result2);
            }
        }
    }
}
