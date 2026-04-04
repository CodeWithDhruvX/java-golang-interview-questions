public class NumberOfDigitOne {
    
    // 233. Number of Digit One - Digit DP
    // Time: O(N * D * 9), Space: O(N * D) where N is number length, D is digit count
    public int countDigitOne(int n) {
        if (n <= 0) {
            return 0;
        }
        
        // Convert to string to get length
        String s = String.valueOf(n);
        int length = s.length();
        
        // Count digit ones for numbers with less digits
        int count = 0;
        for (int i = 1; i < length; i++) {
            count += countDigitOneHelper(i - 1);
        }
        
        // Count for numbers with same number of digits
        int firstDigit = s.charAt(0) - '0';
        int remaining = n - firstDigit * (int) Math.pow(10, length - 1);
        
        if (firstDigit == 1) {
            count += remaining + 1;
        } else {
            count += (int) Math.pow(10, length - 1) + 
                     countDigitOneHelper(firstDigit - 1) * (int) Math.pow(10, length - 1);
        }
        
        return count;
    }
    
    private int countDigitOneHelper(int n) {
        if (n <= 0) {
            return 0;
        }
        
        int count = 0;
        for (int i = 1; i <= n; i++) {
            count += countOnesInNumber(i);
        }
        return count;
    }
    
    private int countOnesInNumber(int num) {
        int count = 0;
        while (num > 0) {
            if (num % 10 == 1) {
                count++;
            }
            num /= 10;
        }
        return count;
    }
    
    // Digit DP with memoization
    public int countDigitOneDP(int n) {
        if (n <= 0) {
            return 0;
        }
        
        String s = String.valueOf(n);
        int length = s.length();
        
        // DP table: dp[pos][tight][hasOne]
        int[][][] dp = new int[length + 1][2][2];
        
        // Initialize with -1 for memoization
        for (int i = 0; i <= length; i++) {
            for (int j = 0; j < 2; j++) {
                for (int k = 0; k < 2; k++) {
                    dp[i][j][k] = -1;
                }
            }
        }
        
        return countDigitOneRecursive(s, 0, 0, 0, dp);
    }
    
    private int countDigitOneRecursive(String s, int pos, int tight, int hasOne, int[][][] dp) {
        if (pos == s.length()) {
            return hasOne;
        }
        
        if (dp[pos][tight][hasOne] != -1) {
            return dp[pos][tight][hasOne];
        }
        
        int limit = tight == 1 ? s.charAt(pos) - '0' : 9;
        int result = 0;
        
        for (int digit = 0; digit <= limit; digit++) {
            int newTight = (tight == 1 && digit == limit) ? 1 : 0;
            int newHasOne = (hasOne == 1 || digit == 1) ? 1 : 0;
            
            result += countDigitOneRecursive(s, pos + 1, newTight, newHasOne, dp);
        }
        
        dp[pos][tight][hasOne] = result;
        return result;
    }
    
    // Mathematical approach - more efficient
    public int countDigitOneMathematical(int n) {
        if (n <= 0) {
            return 0;
        }
        
        int count = 0;
        int factor = 1;
        
        while (factor <= n) {
            int lower = n - (n / factor) * factor;
            int current = (n / factor) % 10;
            int higher = n / (factor * 10);
            
            if (current == 0) {
                count += higher * factor;
            } else if (current == 1) {
                count += higher * factor + lower + 1;
            } else {
                count += (higher + 1) * factor;
            }
            
            factor *= 10;
        }
        
        return count;
    }
    
    // Optimized mathematical approach
    public int countDigitOneOptimized(int n) {
        if (n <= 0) {
            return 0;
        }
        
        int count = 0;
        long factor = 1; // Use long to prevent overflow
        
        while (factor <= n) {
            long lower = n - (n / (int) factor) * (int) factor;
            int current = (n / (int) factor) % 10;
            long higher = n / (factor * 10);
            
            count += higher * factor;
            
            if (current == 1) {
                count += lower + 1;
            } else if (current > 1) {
                count += factor;
            }
            
            factor *= 10;
        }
        
        return count;
    }
    
    // Version with detailed explanation
    public class CountResult {
        int count;
        java.util.List<String> explanation;
        
        CountResult(int count, java.util.List<String> explanation) {
            this.count = count;
            this.explanation = explanation;
        }
    }
    
    public CountResult countDigitOneDetailed(int n) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== Mathematical Approach to Count Digit One ===");
        explanation.add("Number: " + n);
        
        if (n <= 0) {
            explanation.add("Number <= 0, returning 0");
            return new CountResult(0, explanation);
        }
        
        int count = 0;
        int factor = 1;
        int step = 1;
        
        while (factor <= n) {
            int lower = n - (n / factor) * factor;
            int current = (n / factor) % 10;
            int higher = n / (factor * 10);
            
            explanation.add(String.format("Step %d: factor=%d", step++, factor));
            explanation.add(String.format("  higher=%d, current=%d, lower=%d", higher, current, lower));
            
            int stepCount = 0;
            if (current == 0) {
                stepCount = higher * factor;
                explanation.add(String.format("  current=0: count += %d * %d = %d", higher, factor, stepCount));
            } else if (current == 1) {
                stepCount = higher * factor + lower + 1;
                explanation.add(String.format("  current=1: count += %d * %d + %d + 1 = %d", higher, factor, lower, stepCount));
            } else {
                stepCount = (higher + 1) * factor;
                explanation.add(String.format("  current>1: count += (%d + 1) * %d = %d", higher, factor, stepCount));
            }
            
            count += stepCount;
            explanation.add(String.format("  Total count so far: %d", count));
            
            factor *= 10;
        }
        
        explanation.add("Final result: " + count);
        return new CountResult(count, explanation);
    }
    
    // Brute force approach for comparison
    public int countDigitOneBruteForce(int n) {
        if (n <= 0) {
            return 0;
        }
        
        int count = 0;
        for (int i = 1; i <= n; i++) {
            count += countOnesInNumber(i);
        }
        return count;
    }
    
    // Count occurrences of any digit (generalized version)
    public int countDigit(int n, int digit) {
        if (n <= 0 || digit < 0 || digit > 9) {
            return 0;
        }
        
        int count = 0;
        int factor = 1;
        
        while (factor <= n) {
            int lower = n - (n / factor) * factor;
            int current = (n / factor) % 10;
            int higher = n / (factor * 10);
            
            if (current == 0) {
                count += higher * factor;
            } else if (current == digit) {
                count += higher * factor + lower + 1;
            } else if (current > digit) {
                count += (higher + 1) * factor;
            } else {
                count += higher * factor;
            }
            
            factor *= 10;
        }
        
        return count;
    }
    
    // Count all digits from 0 to 9
    public int[] countAllDigits(int n) {
        int[] counts = new int[10];
        
        for (int digit = 0; digit <= 9; digit++) {
            counts[digit] = countDigit(n, digit);
        }
        
        return counts;
    }
    
    public static void main(String[] args) {
        NumberOfDigitOne solver = new NumberOfDigitOne();
        
        // Test cases
        int[] testCases = {
            13,
            0,
            1,
            10,
            11,
            99,
            100,
            101,
            111,
            999,
            1000,
            1234,
            9999,
            10000,
            12345
        };
        
        String[] descriptions = {
            "Small number",
            "Zero",
            "Single digit one",
            "Two digits",
            "Contains one",
            "Two digits max",
            "Three digits",
            "Contains ones",
            "All ones",
            "Three digits max",
            "Four digits",
            "Mixed digits",
            "Four digits max",
            "Five digits",
            "Complex case"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s (n=%d)\n", i + 1, descriptions[i], testCases[i]);
            
            int result1 = solver.countDigitOne(testCases[i]);
            int result2 = solver.countDigitOneMathematical(testCases[i]);
            int result3 = solver.countDigitOneOptimized(testCases[i]);
            int result4 = solver.countDigitOneBruteForce(testCases[i]);
            
            System.out.printf("  Original: %d\n", result1);
            System.out.printf("  Mathematical: %d\n", result2);
            System.out.printf("  Optimized: %d\n", result3);
            System.out.printf("  Brute Force: %d\n", result4);
            
            // Count all digits
            int[] allDigits = solver.countAllDigits(testCases[i]);
            System.out.printf("  All digits: %s\n", java.util.Arrays.toString(allDigits));
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        CountResult detailedResult = solver.countDigitOneDetailed(1234);
        System.out.printf("Result: %d\n", detailedResult.count);
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        int largeN = 1000000;
        
        long startTime = System.nanoTime();
        int largeResult1 = solver.countDigitOneOptimized(largeN);
        long endTime = System.nanoTime();
        
        System.out.printf("Large test (n=%d) - Optimized: %d (took %d ns)\n", 
            largeN, largeResult1, endTime - startTime);
        
        startTime = System.nanoTime();
        int largeResult2 = solver.countDigitOneMathematical(largeN);
        endTime = System.nanoTime();
        
        System.out.printf("Large test (n=%d) - Mathematical: %d (took %d ns)\n", 
            largeN, largeResult2, endTime - startTime);
        
        // Brute force for smaller number
        int mediumN = 10000;
        startTime = System.nanoTime();
        int mediumResult = solver.countDigitOneBruteForce(mediumN);
        endTime = System.nanoTime();
        
        System.out.printf("Medium test (n=%d) - Brute Force: %d (took %d ns)\n", 
            mediumN, mediumResult, endTime - startTime);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Negative number: %d\n", solver.countDigitOne(-1));
        System.out.printf("Very large number: %d\n", solver.countDigitOne(Integer.MAX_VALUE));
        System.out.printf("Count digit 7 in 777: %d\n", solver.countDigit(777, 7));
        System.out.printf("Count digit 0 in 1000: %d\n", solver.countDigit(1000, 0));
    }
}
