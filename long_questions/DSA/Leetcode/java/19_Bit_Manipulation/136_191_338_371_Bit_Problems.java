import java.util.*;

public class BitManipulation {
    
    // 136. Single Number
    // Time: O(N), Space: O(1)
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num;
        }
        return result;
    }

    // 191. Number of 1 Bits
    // Time: O(1) for 32-bit integers, Space: O(1)
    public static int hammingWeight(int n) {
        int count = 0;
        while (n != 0) {
            count++;
            n &= (n - 1); // Clear the least significant bit
        }
        return count;
    }

    // 338. Counting Bits
    // Time: O(N), Space: O(N)
    public static int[] countBits(int n) {
        int[] result = new int[n + 1];
        
        for (int i = 1; i <= n; i++) {
            // Number of bits in i = number of bits in (i >> 1) + (i & 1)
            result[i] = result[i >> 1] + (i & 1);
        }
        
        return result;
    }

    // 371. Sum of Two Integers
    // Time: O(1) for 32-bit integers, Space: O(1)
    public static int getSum(int a, int b) {
        while (b != 0) {
            int carry = a & b;
            a = a ^ b;
            b = carry << 1;
        }
        return a;
    }

    // 268. Missing Number
    // Time: O(N), Space: O(1)
    public static int missingNumber(int[] nums) {
        int n = nums.length;
        int result = n;
        
        for (int i = 0; i < n; i++) {
            result ^= i ^ nums[i];
        }
        
        return result;
    }

    // 190. Reverse Bits
    // Time: O(1) for 32-bit integers, Space: O(1)
    public static int reverseBits(int n) {
        int result = 0;
        
        for (int i = 0; i < 32; i++) {
            result <<= 1;
            result |= (n & 1);
            n >>= 1;
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases for singleNumber
        int[][] testCases1 = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
            {1},
            {2, 2, 3, 3, 4},
            {0, 1, 1},
            {1, 1, 2, 2, 3, 3, 4},
            {5, 5, 6, 6, 7, 7, 8},
            {9, 9, 9, 9, 10},
            {11, 12, 11, 13, 12},
            {14, 15, 14, 16, 15}
        };
        
        // Test cases for hammingWeight
        int[] testCases2 = {
            0b00000000000000000000000000001011,
            0b00000000000000000000000010000000,
            0b11111111111111111111111111111101,
            0,
            1,
            2,
            3,
            7,
            15,
            31
        };
        
        // Test cases for countBits
        int[] testCases3 = {2, 5, 10, 15, 20, 25, 30, 0, 1, 3};
        
        // Test cases for getSum
        int[][] testCases4 = {
            {1, 2},
            {2, 3},
            {3, 4},
            {0, 5},
            {-1, 1},
            {-2, -3},
            {10, -5},
            {100, 200},
            {-10, 20},
            {50, -30}
        };
        
        // Test cases for missingNumber
        int[][] testCases5 = {
            {3, 0, 1},
            {0, 1},
            {9, 6, 4, 2, 3, 5, 7, 8, 0, 1},
            {0},
            {1},
            {2, 0},
            {1, 2, 3, 5},
            {0, 2, 3, 4},
            {1, 0, 2, 4},
            {0, 1, 3, 4}
        };
        
        // Test cases for reverseBits
        int[] testCases6 = {
            0b00000010100101000001111010011100,
            0b00111101000010000000000000000000,
            0b11111111111111111111111111111101,
            0,
            1,
            0b10000000000000000000000000000000,
            0b00000000000000000000000000000001,
            0b10101010101010101010101010101010,
            0b11111111111111111111111111111111,
            0b10000000000000000000000000000001
        };
        
        System.out.println("Single Number:");
        for (int i = 0; i < testCases1.length; i++) {
            int[] nums = testCases1[i];
            int result = singleNumber(nums);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(nums), result);
        }
        
        System.out.println("\nHamming Weight:");
        for (int i = 0; i < testCases2.length; i++) {
            int n = testCases2[i];
            int result = hammingWeight(n);
            System.out.printf("Test Case %d: %d -> %d\n", 
                i + 1, n, result);
        }
        
        System.out.println("\nCounting Bits:");
        for (int i = 0; i < testCases3.length; i++) {
            int n = testCases3[i];
            int[] result = countBits(n);
            System.out.printf("Test Case %d: n=%d -> %s\n", 
                i + 1, n, Arrays.toString(result));
        }
        
        System.out.println("\nSum of Two Integers:");
        for (int i = 0; i < testCases4.length; i++) {
            int a = testCases4[i][0];
            int b = testCases4[i][1];
            int result = getSum(a, b);
            System.out.printf("Test Case %d: %d + %d = %d\n", 
                i + 1, a, b, result);
        }
        
        System.out.println("\nMissing Number:");
        for (int i = 0; i < testCases5.length; i++) {
            int[] nums = testCases5[i];
            int result = missingNumber(nums);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(nums), result);
        }
        
        System.out.println("\nReverse Bits:");
        for (int i = 0; i < testCases6.length; i++) {
            int n = testCases6[i];
            int result = reverseBits(n);
            System.out.printf("Test Case %d: %d -> %d\n", 
                i + 1, n, result);
        }
    }
}
