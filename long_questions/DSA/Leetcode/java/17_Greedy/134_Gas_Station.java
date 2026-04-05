import java.util.*;

public class GasStation {
    
    // 134. Gas Station
    // Time: O(N), Space: O(1)
    public static int canCompleteCircuit(int[] gas, int[] cost) {
        if (gas == null || cost == null || gas.length != cost.length) {
            return -1;
        }
        
        int totalGas = 0;
        int totalCost = 0;
        int currentGas = 0;
        int startIndex = 0;
        
        for (int i = 0; i < gas.length; i++) {
            totalGas += gas[i];
            totalCost += cost[i];
            currentGas += gas[i] - cost[i];
            
            // If currentGas is negative, we can't start from current startIndex
            if (currentGas < 0) {
                startIndex = i + 1;
                currentGas = 0;
            }
        }
        
        return totalGas >= totalCost ? startIndex : -1;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 2, 3, 4, 5}, new int[]{3, 4, 5, 1, 2}},
            {new int[]{2, 3, 4}, new int[]{3, 4, 3}},
            {new int[]{5, 1, 2, 3, 4}, new int[]{4, 4, 1, 5, 1}},
            {new int[]{3, 3, 4}, new int[]{3, 4, 4}},
            {new int[]{1}, new int[]{1}},
            {new int[]{2}, new int[]{1}},
            {new int[]{1}, new int[]{2}},
            {new int[]{4, 5, 6, 7, 8}, new int[]{5, 6, 7, 8, 9}},
            {new int[]{5, 5, 5, 5}, new int[]{4, 4, 4, 4}},
            {new int[]{1, 2, 3, 4, 5}, new int[]{1, 2, 3, 4, 5}},
            {new int[]{6, 7, 8, 9, 10}, new int[]{5, 6, 7, 8, 9}},
            {new int[]{2, 4, 6}, new int[]{3, 5, 7}},
            {new int[]{3, 1, 2}, new int[]{2, 2, 2}},
            {new int[]{5, 8, 2, 8}, new int[]{6, 5, 6, 3}},
            {new int[]{1, 2, 3, 4, 5, 6}, new int[]{2, 3, 4, 5, 6, 1}}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] gas = (int[]) testCases[i][0];
            int[] cost = (int[]) testCases[i][1];
            
            int result = canCompleteCircuit(gas, cost);
            System.out.printf("Test Case %d: gas=%s, cost=%s -> start=%d\n", 
                i + 1, Arrays.toString(gas), Arrays.toString(cost), result);
        }
    }
}
