import java.util.*;

public class DailyTemperatures {
    
    // 739. Daily Temperatures
    // Time: O(N), Space: O(N)
    public static int[] dailyTemperatures(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Store indices
        
        for (int i = 0; i < n; i++) {
            // While current temperature is greater than temperature at stack top
            while (!stack.isEmpty() && temperatures[i] > temperatures[stack.peek()]) {
                int prevIndex = stack.pop();
                result[prevIndex] = i - prevIndex;
            }
            stack.push(i);
        }
        
        // Remaining indices in stack have no warmer day (result is already 0)
        return result;
    }

    // Alternative approach using reverse iteration
    public static int[] dailyTemperaturesReverse(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        Deque<Integer> stack = new ArrayDeque<>(); // Store indices
        
        // Process from right to left
        for (int i = n - 1; i >= 0; i--) {
            // Remove indices that are not warmer than current temperature
            while (!stack.isEmpty() && temperatures[i] >= temperatures[stack.peek()]) {
                stack.pop();
            }
            
            if (!stack.isEmpty()) {
                result[i] = stack.peek() - i;
            }
            
            stack.push(i);
        }
        
        return result;
    }

    // Brute force approach for comparison (O(N^2))
    public static int[] dailyTemperaturesBruteForce(int[] temperatures) {
        int n = temperatures.length;
        int[] result = new int[n];
        
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (temperatures[j] > temperatures[i]) {
                    result[i] = j - i;
                    break;
                }
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {73, 74, 75, 71, 69, 72, 76, 73},
            {30, 40, 50, 60},
            {30, 60, 90},
            {90, 80, 70, 60},
            {55, 60, 65, 70, 75},
            {65, 70, 65, 60, 65},
            {50},
            {},
            {73, 73, 73, 73},
            {30, 40, 30, 50, 30, 60, 30}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] result1 = dailyTemperatures(testCases[i]);
            int[] result2 = dailyTemperaturesReverse(testCases[i]);
            int[] result3 = dailyTemperaturesBruteForce(testCases[i]);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Forward: %s\n", Arrays.toString(result1));
            System.out.printf("  Reverse: %s\n", Arrays.toString(result2));
            System.out.printf("  Brute:   %s\n\n", Arrays.toString(result3));
        }
    }
}
