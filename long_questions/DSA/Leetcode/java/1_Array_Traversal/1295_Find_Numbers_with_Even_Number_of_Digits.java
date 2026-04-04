import java.util.Arrays;

public class FindNumbersWithEvenNumberOfDigits {
    
    // 1295. Find Numbers with Even Number of Digits
    // Time: O(N), Space: O(1)
    public static int findNumbers(int[] nums) {
        int count = 0;
        
        for (int num : nums) {
            if (hasEvenDigits(num)) {
                count++;
            }
        }
        
        return count;
    }

    // Helper function to check if a number has even number of digits
    private static boolean hasEvenDigits(int num) {
        if (num == 0) {
            return false;
        }
        
        int digitCount = 0;
        while (num > 0) {
            num /= 10;
            digitCount++;
        }
        
        return digitCount % 2 == 0;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {12, 345, 2, 6, 7896},
            {555, 901, 482, 1771},
            {1, 22, 333, 4444},
            {0, 10, 100, 1000},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = findNumbers(testCases[i]);
            System.out.printf("Test Case %d: %s -> Numbers with even digits: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
