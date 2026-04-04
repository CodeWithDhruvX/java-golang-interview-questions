import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class ContainsDuplicate {
    
    // 217. Contains Duplicate
    // Time: O(N), Space: O(N)
    public static boolean containsDuplicate(int[] nums) {
        Set<Integer> numSet = new HashSet<>();
        
        for (int num : nums) {
            if (!numSet.add(num)) {
                return true;
            }
        }
        
        return false;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 1},
            {1, 2, 3, 4},
            {1, 1, 1, 3, 2, 2, 2},
            {},
            {0},
            {-1, -2, -3, -1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = containsDuplicate(testCases[i]);
            System.out.printf("Test Case %d: %s -> Contains duplicate: %b\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
